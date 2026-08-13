package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Data_Map_SemigroupMap gopurs_runtime.Value
var once_Data_Map_SemigroupMap sync.Once
func Get_Data_Map_SemigroupMap() gopurs_runtime.Value {
	once_Data_Map_SemigroupMap.Do(func() {
		cache_Data_Map_SemigroupMap = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Map_SemigroupMap(x_0_box)
})
	})
	return cache_Data_Map_SemigroupMap
}

var cache_Data_Map_traversableWithIndexSemigroupMap gopurs_runtime.Value
var once_Data_Map_traversableWithIndexSemigroupMap sync.Once
func Get_Data_Map_traversableWithIndexSemigroupMap() gopurs_runtime.Value {
	once_Data_Map_traversableWithIndexSemigroupMap.Do(func() {
		cache_Data_Map_traversableWithIndexSemigroupMap = Get_Data_Map_Internal_traversableWithIndexMap()
	})
	return cache_Data_Map_traversableWithIndexSemigroupMap
}

var cache_Data_Map_traversableSemigroupMap gopurs_runtime.Value
var once_Data_Map_traversableSemigroupMap sync.Once
func Get_Data_Map_traversableSemigroupMap() gopurs_runtime.Value {
	once_Data_Map_traversableSemigroupMap.Do(func() {
		cache_Data_Map_traversableSemigroupMap = Get_Data_Map_Internal_traversableMap()
	})
	return cache_Data_Map_traversableSemigroupMap
}

var cache_Data_Map_showSemigroupMap gopurs_runtime.Value
var once_Data_Map_showSemigroupMap sync.Once
func Get_Data_Map_showSemigroupMap() gopurs_runtime.Value {
	once_Data_Map_showSemigroupMap.Do(func() {
		cache_Data_Map_showSemigroupMap = gopurs_runtime.Func2(func(dictShow_0_box gopurs_runtime.Value, dictShow1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Map_showSemigroupMap(dictShow_0_box, dictShow1_1_box)
})
	})
	return cache_Data_Map_showSemigroupMap
}

var cache_Data_Map_semigroupSemigroupMap gopurs_runtime.Value
var once_Data_Map_semigroupSemigroupMap sync.Once
func Get_Data_Map_semigroupSemigroupMap() gopurs_runtime.Value {
	once_Data_Map_semigroupSemigroupMap.Do(func() {
		cache_Data_Map_semigroupSemigroupMap = gopurs_runtime.Func2(func(dictOrd_0_box gopurs_runtime.Value, dictSemigroup_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Map_semigroupSemigroupMap(dictOrd_0_box, dictSemigroup_1_box)
})
	})
	return cache_Data_Map_semigroupSemigroupMap
}

var cache_Data_Map_plusSemigroupMap gopurs_runtime.Value
var once_Data_Map_plusSemigroupMap sync.Once
func Get_Data_Map_plusSemigroupMap() gopurs_runtime.Value {
	once_Data_Map_plusSemigroupMap.Do(func() {
		cache_Data_Map_plusSemigroupMap = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Map_plusSemigroupMap(dictOrd_0_box)
})
	})
	return cache_Data_Map_plusSemigroupMap
}

var cache_Data_Map_ordSemigroupMap gopurs_runtime.Value
var once_Data_Map_ordSemigroupMap sync.Once
func Get_Data_Map_ordSemigroupMap() gopurs_runtime.Value {
	once_Data_Map_ordSemigroupMap.Do(func() {
		cache_Data_Map_ordSemigroupMap = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Map_ordSemigroupMap(dictOrd_0_box)
})
	})
	return cache_Data_Map_ordSemigroupMap
}

var cache_Data_Map_ord1SemigroupMap gopurs_runtime.Value
var once_Data_Map_ord1SemigroupMap sync.Once
func Get_Data_Map_ord1SemigroupMap() gopurs_runtime.Value {
	once_Data_Map_ord1SemigroupMap.Do(func() {
		cache_Data_Map_ord1SemigroupMap = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Map_ord1SemigroupMap(dictOrd_0_box)
})
	})
	return cache_Data_Map_ord1SemigroupMap
}

var cache_Data_Map_newtypeSemigroupMap gopurs_runtime.Value
var once_Data_Map_newtypeSemigroupMap sync.Once
func Get_Data_Map_newtypeSemigroupMap() gopurs_runtime.Value {
	once_Data_Map_newtypeSemigroupMap.Do(func() {
		cache_Data_Map_newtypeSemigroupMap = gopurs_runtime.RecordDict1("Coercible0", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{}
}))
	})
	return cache_Data_Map_newtypeSemigroupMap
}

var cache_Data_Map_monoidSemigroupMap gopurs_runtime.Value
var once_Data_Map_monoidSemigroupMap sync.Once
func Get_Data_Map_monoidSemigroupMap() gopurs_runtime.Value {
	once_Data_Map_monoidSemigroupMap.Do(func() {
		cache_Data_Map_monoidSemigroupMap = gopurs_runtime.Func2(func(dictOrd_0_box gopurs_runtime.Value, dictSemigroup_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Map_monoidSemigroupMap(dictOrd_0_box, dictSemigroup_1_box)
})
	})
	return cache_Data_Map_monoidSemigroupMap
}

var cache_Data_Map_keys gopurs_runtime.Value
var once_Data_Map_keys sync.Once
func Get_Data_Map_keys() gopurs_runtime.Value {
	once_Data_Map_keys.Do(func() {
		cache_Data_Map_keys = func() gopurs_runtime.Value {
var go__go_0_1_0 gopurs_runtime.Value
_ = go__go_0_1_0
go__go_0_1_0 = gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 *Constructor_Data_Map_Internal_Node
{
if (v_1.Type == 9 && v_1.IntVal == 324739070 && v_1.UnsafePtr == nil) {
__t2 = (*Constructor_Data_Map_Internal_Node)(nil)
goto end_branch_2
} else {

}
}
{
if (v_1.Type == 9 && v_1.IntVal == 324739070 && v_1.UnsafePtr != nil) {
__t2 = &Constructor_Data_Map_Internal_Node{1, (*Constructor_Data_Map_Internal_Node)(v_1.UnsafePtr).V0, (*Constructor_Data_Map_Internal_Node)(v_1.UnsafePtr).V1, (*Constructor_Data_Map_Internal_Node)(v_1.UnsafePtr).V2, Get_Data_Unit_unit(), gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.Apply(go__go_0_1_0, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v_1.UnsafePtr).V4)})), gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.Apply(go__go_0_1_0, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v_1.UnsafePtr).V5)}))}
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_2:
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(__t2)}
})
// TAST (Let): __local_var_0_0 -> gopurs_runtime.Value
__local_var_0_0 := go__go_0_1_0
_ = __local_var_0_0
return gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_0_0, x_1)
})
}()
	})
	return cache_Data_Map_keys
}

var cache_Data_Map_functorWithIndexSemigroupMap gopurs_runtime.Value
var once_Data_Map_functorWithIndexSemigroupMap sync.Once
func Get_Data_Map_functorWithIndexSemigroupMap() gopurs_runtime.Value {
	once_Data_Map_functorWithIndexSemigroupMap.Do(func() {
		cache_Data_Map_functorWithIndexSemigroupMap = Get_Data_Map_Internal_functorWithIndexMap()
	})
	return cache_Data_Map_functorWithIndexSemigroupMap
}

var cache_Data_Map_functorSemigroupMap gopurs_runtime.Value
var once_Data_Map_functorSemigroupMap sync.Once
func Get_Data_Map_functorSemigroupMap() gopurs_runtime.Value {
	once_Data_Map_functorSemigroupMap.Do(func() {
		cache_Data_Map_functorSemigroupMap = Get_Data_Map_Internal_functorMap()
	})
	return cache_Data_Map_functorSemigroupMap
}

var cache_Data_Map_foldableWithIndexSemigroupMap gopurs_runtime.Value
var once_Data_Map_foldableWithIndexSemigroupMap sync.Once
func Get_Data_Map_foldableWithIndexSemigroupMap() gopurs_runtime.Value {
	once_Data_Map_foldableWithIndexSemigroupMap.Do(func() {
		cache_Data_Map_foldableWithIndexSemigroupMap = Get_Data_Map_Internal_foldableWithIndexMap()
	})
	return cache_Data_Map_foldableWithIndexSemigroupMap
}

var cache_Data_Map_foldableSemigroupMap gopurs_runtime.Value
var once_Data_Map_foldableSemigroupMap sync.Once
func Get_Data_Map_foldableSemigroupMap() gopurs_runtime.Value {
	once_Data_Map_foldableSemigroupMap.Do(func() {
		cache_Data_Map_foldableSemigroupMap = Get_Data_Map_Internal_foldableMap()
	})
	return cache_Data_Map_foldableSemigroupMap
}

var cache_Data_Map_eqSemigroupMap gopurs_runtime.Value
var once_Data_Map_eqSemigroupMap sync.Once
func Get_Data_Map_eqSemigroupMap() gopurs_runtime.Value {
	once_Data_Map_eqSemigroupMap.Do(func() {
		cache_Data_Map_eqSemigroupMap = gopurs_runtime.Func2(func(dictEq_0_box gopurs_runtime.Value, dictEq1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Map_eqSemigroupMap(dictEq_0_box, dictEq1_1_box)
})
	})
	return cache_Data_Map_eqSemigroupMap
}

var cache_Data_Map_eq1SemigroupMap gopurs_runtime.Value
var once_Data_Map_eq1SemigroupMap sync.Once
func Get_Data_Map_eq1SemigroupMap() gopurs_runtime.Value {
	once_Data_Map_eq1SemigroupMap.Do(func() {
		cache_Data_Map_eq1SemigroupMap = gopurs_runtime.Func(func(dictEq_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Map_eq1SemigroupMap(dictEq_0_box)
})
	})
	return cache_Data_Map_eq1SemigroupMap
}

var cache_Data_Map_bindSemigroupMap gopurs_runtime.Value
var once_Data_Map_bindSemigroupMap sync.Once
func Get_Data_Map_bindSemigroupMap() gopurs_runtime.Value {
	once_Data_Map_bindSemigroupMap.Do(func() {
		cache_Data_Map_bindSemigroupMap = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Map_bindSemigroupMap(dictOrd_0_box)
})
	})
	return cache_Data_Map_bindSemigroupMap
}

var cache_Data_Map_applySemigroupMap gopurs_runtime.Value
var once_Data_Map_applySemigroupMap sync.Once
func Get_Data_Map_applySemigroupMap() gopurs_runtime.Value {
	once_Data_Map_applySemigroupMap.Do(func() {
		cache_Data_Map_applySemigroupMap = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Map_applySemigroupMap(dictOrd_0_box)
})
	})
	return cache_Data_Map_applySemigroupMap
}

var cache_Data_Map_altSemigroupMap gopurs_runtime.Value
var once_Data_Map_altSemigroupMap sync.Once
func Get_Data_Map_altSemigroupMap() gopurs_runtime.Value {
	once_Data_Map_altSemigroupMap.Do(func() {
		cache_Data_Map_altSemigroupMap = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Map_altSemigroupMap(dictOrd_0_box)
})
	})
	return cache_Data_Map_altSemigroupMap
}

func Call_Data_Map_SemigroupMap(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_Data_Map_showSemigroupMap(dictShow_0_loop gopurs_runtime.Value, dictShow1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictShow_0 gopurs_runtime.Value = dictShow_0_loop
_ = dictShow_0
var dictShow1_1 gopurs_runtime.Value = dictShow1_1_loop
_ = dictShow1_1
return gopurs_runtime.Apply2(Get_Data_Map_Internal_showMap(), dictShow_0, dictShow1_1)
}

func Call_Data_Map_semigroupSemigroupMap(dictOrd_0_loop gopurs_runtime.Value, dictSemigroup_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
var dictSemigroup_1 gopurs_runtime.Value = dictSemigroup_1_loop
_ = dictSemigroup_1
// TAST (Let): append_2_0 -> gopurs_runtime.Value
append_2_0 := gopurs_runtime.RecordGet(dictSemigroup_1, "append")
_ = append_2_0
return gopurs_runtime.RecordDict1("append", gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeUnionWith(), gopurs_runtime.RecordGet(dictOrd_0, "compare"), append_2_0, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](v_3))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](v1_4))})))}
})
}))
}

func Call_Data_Map_plusSemigroupMap(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
return gopurs_runtime.Apply(Get_Data_Map_Internal_plusMap(), dictOrd_0)
}

func Call_Data_Map_ordSemigroupMap(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
// TAST (Let): ordMap_1_0 -> gopurs_runtime.Value
ordMap_1_0 := gopurs_runtime.Apply(Get_Data_Map_Internal_ordMap(), dictOrd_0)
_ = ordMap_1_0
return gopurs_runtime.Func(func(dictOrd1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(ordMap_1_0, dictOrd1_2)
})
}

func Call_Data_Map_ord1SemigroupMap(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
return gopurs_runtime.Apply(Get_Data_Map_Internal_ord1Map(), dictOrd_0)
}

func Call_Data_Map_monoidSemigroupMap(dictOrd_0_loop gopurs_runtime.Value, dictSemigroup_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
var dictSemigroup_1 gopurs_runtime.Value = dictSemigroup_1_loop
_ = dictSemigroup_1
// TAST (Let): append_2_1 -> gopurs_runtime.Value
append_2_1 := gopurs_runtime.RecordGet(dictSemigroup_1, "append")
_ = append_2_1
// TAST (Let): semigroupSemigroupMap2_2_0 -> gopurs_runtime.Value
semigroupSemigroupMap2_2_0 := gopurs_runtime.RecordDict1("append", gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeUnionWith(), gopurs_runtime.RecordGet(dictOrd_0, "compare"), append_2_1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](v_3))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](v1_4))})))}
})
}))
_ = semigroupSemigroupMap2_2_0
return gopurs_runtime.RecordDict2("Semigroup0", "mempty", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return semigroupSemigroupMap2_2_0
}), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(nil))})
}

func Call_Data_Map_eqSemigroupMap(dictEq_0_loop gopurs_runtime.Value, dictEq1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictEq_0 gopurs_runtime.Value = dictEq_0_loop
_ = dictEq_0
var dictEq1_1 gopurs_runtime.Value = dictEq1_1_loop
_ = dictEq1_1
return gopurs_runtime.Apply2(Get_Data_Map_Internal_eqMap(), dictEq_0, dictEq1_1)
}

func Call_Data_Map_eq1SemigroupMap(dictEq_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictEq_0 gopurs_runtime.Value = dictEq_0_loop
_ = dictEq_0
return gopurs_runtime.Apply(Get_Data_Map_Internal_eq1Map(), dictEq_0)
}

func Call_Data_Map_bindSemigroupMap(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
return gopurs_runtime.Apply(Get_Data_Map_Internal_bindMap(), dictOrd_0)
}

func Call_Data_Map_applySemigroupMap(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
return gopurs_runtime.Apply(Get_Data_Map_Internal_applyMap(), dictOrd_0)
}

func Call_Data_Map_altSemigroupMap(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
// TAST (Let): compare_1_0 -> gopurs_runtime.Value
compare_1_0 := gopurs_runtime.RecordGet(dictOrd_0, "compare")
_ = compare_1_0
return gopurs_runtime.RecordDict2("Functor0", "alt", gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Map_Internal_functorMap()
}), gopurs_runtime.Func(func(m1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m2_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeUnionWith(), compare_1_0, Get_Data_Function_go__const(), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](m1_2))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](m2_3))})))}
})
}))
}


