package Data_Map

import (
	pkg_Data_Function "gopurs/output/Data.Function"
	pkg_Data_Map_Internal "gopurs/output/Data.Map.Internal"
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
		cache_semigroupSemigroupMap = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_semigroupSemigroupMap(dictOrd_0_box)
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
		cache_monoidSemigroupMap = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_monoidSemigroupMap(dictOrd_0_box)
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

func Call_semigroupSemigroupMap(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
compare_1_0 := gopurs_runtime.RecordGet(dictOrd_0, "compare")
_ = compare_1_0
return gopurs_runtime.Func(func(dictSemigroup_2 gopurs_runtime.Value) gopurs_runtime.Value {
append_3_1 := gopurs_runtime.RecordGet(dictSemigroup_2, "append")
_ = append_3_1
return gopurs_runtime.RecordDict1("append", gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.UncurriedApp4(pkg_Data_Map_Internal.Get_unsafeUnionWith(), compare_1_0, append_3_1, v_4, v1_5).UnsafePtr))}
})
}))
})
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

func Call_monoidSemigroupMap(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
semigroupSemigroupMap1_1_0 := Call_semigroupSemigroupMap(dictOrd_0)
_ = semigroupSemigroupMap1_1_0
return gopurs_runtime.Func(func(dictSemigroup_2 gopurs_runtime.Value) gopurs_runtime.Value {
semigroupSemigroupMap2_3_1 := gopurs_runtime.Apply(semigroupSemigroupMap1_1_0, dictSemigroup_2)
_ = semigroupSemigroupMap2_3_1
return gopurs_runtime.RecordDict2("Semigroup0", "mempty", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return semigroupSemigroupMap2_3_1
}), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: nil}.UnsafePtr))})
})
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
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.UncurriedApp4(pkg_Data_Map_Internal.Get_unsafeIntersectionWith(), compare_1_0, pkg_Data_Map_Internal.Get_identity2(), m1_2, m2_3).UnsafePtr))}
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
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.UncurriedApp4(pkg_Data_Map_Internal.Get_unsafeUnionWith(), compare_1_0, pkg_Data_Function.Get_go__const(), m1_2, m2_3).UnsafePtr))}
})
}))
}


