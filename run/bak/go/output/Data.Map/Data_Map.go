package Data_Map

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_Map_Internal "gopurs/output/Data.Map.Internal"
	pkg_Data_Unit "gopurs/output/Data.Unit"
	pkg_Control_Category "gopurs/output/Control.Category"
	pkg_Data_Function "gopurs/output/Data.Function"
)

var SemigroupMap gopurs_runtime.Value
var once_SemigroupMap sync.Once
func Get_SemigroupMap() gopurs_runtime.Value {
	once_SemigroupMap.Do(func() {
		SemigroupMap = gopurs_runtime.Func(func(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}()
})
	})
	return SemigroupMap
}

var traversableWithIndexSemigroupMap gopurs_runtime.Value
var once_traversableWithIndexSemigroupMap sync.Once
func Get_traversableWithIndexSemigroupMap() gopurs_runtime.Value {
	once_traversableWithIndexSemigroupMap.Do(func() {
		traversableWithIndexSemigroupMap = pkg_Data_Map_Internal.Get_traversableWithIndexMap()
	})
	return traversableWithIndexSemigroupMap
}

var traversableSemigroupMap gopurs_runtime.Value
var once_traversableSemigroupMap sync.Once
func Get_traversableSemigroupMap() gopurs_runtime.Value {
	once_traversableSemigroupMap.Do(func() {
		traversableSemigroupMap = pkg_Data_Map_Internal.Get_traversableMap()
	})
	return traversableSemigroupMap
}

var showSemigroupMap gopurs_runtime.Value
var once_showSemigroupMap sync.Once
func Get_showSemigroupMap() gopurs_runtime.Value {
	once_showSemigroupMap.Do(func() {
		showSemigroupMap = gopurs_runtime.Func2(func(dictShow_0_box gopurs_runtime.Value, dictShow1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_showSemigroupMap(dictShow_0_box, dictShow1_1_box)
})
	})
	return showSemigroupMap
}

var semigroupSemigroupMap gopurs_runtime.Value
var once_semigroupSemigroupMap sync.Once
func Get_semigroupSemigroupMap() gopurs_runtime.Value {
	once_semigroupSemigroupMap.Do(func() {
		semigroupSemigroupMap = gopurs_runtime.Func(func(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
compare_1_0 := gopurs_runtime.RecordGet(dictOrd_0, "compare")
_ = compare_1_0
return gopurs_runtime.Func(func(dictSemigroup_2 gopurs_runtime.Value) gopurs_runtime.Value {
append_3_1 := gopurs_runtime.RecordGet(dictSemigroup_2, "append")
_ = append_3_1
return gopurs_runtime.RecordDict1("append", gopurs_runtime.Func2(func(v_4 gopurs_runtime.Value, v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.UncurriedApp4(pkg_Data_Map_Internal.Get_unsafeUnionWith(), compare_1_0, append_3_1, v_4, v1_5)
}))
})
}()
})
	})
	return semigroupSemigroupMap
}

var plusSemigroupMap gopurs_runtime.Value
var once_plusSemigroupMap sync.Once
func Get_plusSemigroupMap() gopurs_runtime.Value {
	once_plusSemigroupMap.Do(func() {
		plusSemigroupMap = gopurs_runtime.Func(func(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
return gopurs_runtime.Apply(pkg_Data_Map_Internal.Get_plusMap(), dictOrd_0)
}()
})
	})
	return plusSemigroupMap
}

var ordSemigroupMap gopurs_runtime.Value
var once_ordSemigroupMap sync.Once
func Get_ordSemigroupMap() gopurs_runtime.Value {
	once_ordSemigroupMap.Do(func() {
		ordSemigroupMap = gopurs_runtime.Func(func(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
return gopurs_runtime.Apply(pkg_Data_Map_Internal.Get_ordMap(), dictOrd_0)
}()
})
	})
	return ordSemigroupMap
}

var ord1SemigroupMap gopurs_runtime.Value
var once_ord1SemigroupMap sync.Once
func Get_ord1SemigroupMap() gopurs_runtime.Value {
	once_ord1SemigroupMap.Do(func() {
		ord1SemigroupMap = gopurs_runtime.Func(func(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
return gopurs_runtime.Apply(pkg_Data_Map_Internal.Get_ord1Map(), dictOrd_0)
}()
})
	})
	return ord1SemigroupMap
}

var newtypeSemigroupMap gopurs_runtime.Value
var once_newtypeSemigroupMap sync.Once
func Get_newtypeSemigroupMap() gopurs_runtime.Value {
	once_newtypeSemigroupMap.Do(func() {
		newtypeSemigroupMap = gopurs_runtime.RecordDict1("Coercible0", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{}
}))
	})
	return newtypeSemigroupMap
}

var monoidSemigroupMap gopurs_runtime.Value
var once_monoidSemigroupMap sync.Once
func Get_monoidSemigroupMap() gopurs_runtime.Value {
	once_monoidSemigroupMap.Do(func() {
		monoidSemigroupMap = gopurs_runtime.Func(func(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
semigroupSemigroupMap1_1_0 := gopurs_runtime.Apply(Get_semigroupSemigroupMap(), dictOrd_0)
_ = semigroupSemigroupMap1_1_0
return gopurs_runtime.Func(func(dictSemigroup_2 gopurs_runtime.Value) gopurs_runtime.Value {
semigroupSemigroupMap2_3_1 := gopurs_runtime.Apply(semigroupSemigroupMap1_1_0, dictSemigroup_2)
_ = semigroupSemigroupMap2_3_1
return gopurs_runtime.RecordDict2("mempty", "Semigroup0", gopurs_runtime.Constructor0("Leaf"), gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return semigroupSemigroupMap2_3_1
}))
})
}()
})
	})
	return monoidSemigroupMap
}

var keys gopurs_runtime.Value
var once_keys sync.Once
func Get_keys() gopurs_runtime.Value {
	once_keys.Do(func() {
		keys = func() gopurs_runtime.Value {
var go__0_0 gopurs_runtime.Value
_ = go__0_0
go__0_0 = gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if gopurs_runtime.Bool(v_1.StrVal == "Leaf").IntVal != 0 {
__t1 = gopurs_runtime.Constructor0("Leaf")
goto end_branch_1
} else {

}
}
{
if gopurs_runtime.Bool(v_1.StrVal == "Node").IntVal != 0 {
__t1 = gopurs_runtime.Constructor("Node", []gopurs_runtime.Value{(*[1024]gopurs_runtime.Value)(v_1.UnsafePtr)[0], (*[1024]gopurs_runtime.Value)(v_1.UnsafePtr)[1], (*[1024]gopurs_runtime.Value)(v_1.UnsafePtr)[2], pkg_Data_Unit.Get_unit(), gopurs_runtime.Apply(go__0_0, (*[1024]gopurs_runtime.Value)(v_1.UnsafePtr)[4]), gopurs_runtime.Apply(go__0_0, (*[1024]gopurs_runtime.Value)(v_1.UnsafePtr)[5])})
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
return __t1
})
return gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(go__0_0, x_1)
})
}()
	})
	return keys
}

var functorWithIndexSemigroupMap gopurs_runtime.Value
var once_functorWithIndexSemigroupMap sync.Once
func Get_functorWithIndexSemigroupMap() gopurs_runtime.Value {
	once_functorWithIndexSemigroupMap.Do(func() {
		functorWithIndexSemigroupMap = pkg_Data_Map_Internal.Get_functorWithIndexMap()
	})
	return functorWithIndexSemigroupMap
}

var functorSemigroupMap gopurs_runtime.Value
var once_functorSemigroupMap sync.Once
func Get_functorSemigroupMap() gopurs_runtime.Value {
	once_functorSemigroupMap.Do(func() {
		functorSemigroupMap = pkg_Data_Map_Internal.Get_functorMap()
	})
	return functorSemigroupMap
}

var foldableWithIndexSemigroupMap gopurs_runtime.Value
var once_foldableWithIndexSemigroupMap sync.Once
func Get_foldableWithIndexSemigroupMap() gopurs_runtime.Value {
	once_foldableWithIndexSemigroupMap.Do(func() {
		foldableWithIndexSemigroupMap = pkg_Data_Map_Internal.Get_foldableWithIndexMap()
	})
	return foldableWithIndexSemigroupMap
}

var foldableSemigroupMap gopurs_runtime.Value
var once_foldableSemigroupMap sync.Once
func Get_foldableSemigroupMap() gopurs_runtime.Value {
	once_foldableSemigroupMap.Do(func() {
		foldableSemigroupMap = pkg_Data_Map_Internal.Get_foldableMap()
	})
	return foldableSemigroupMap
}

var eqSemigroupMap gopurs_runtime.Value
var once_eqSemigroupMap sync.Once
func Get_eqSemigroupMap() gopurs_runtime.Value {
	once_eqSemigroupMap.Do(func() {
		eqSemigroupMap = gopurs_runtime.Func2(func(dictEq_0_box gopurs_runtime.Value, dictEq1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_eqSemigroupMap(dictEq_0_box, dictEq1_1_box)
})
	})
	return eqSemigroupMap
}

var eq1SemigroupMap gopurs_runtime.Value
var once_eq1SemigroupMap sync.Once
func Get_eq1SemigroupMap() gopurs_runtime.Value {
	once_eq1SemigroupMap.Do(func() {
		eq1SemigroupMap = gopurs_runtime.Func(func(dictEq_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictEq_0 gopurs_runtime.Value = dictEq_0_loop
_ = dictEq_0
return gopurs_runtime.RecordDict1("eq1", gopurs_runtime.Func(func(dictEq1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(gopurs_runtime.Apply2(pkg_Data_Map_Internal.Get_eqMap(), dictEq_0, dictEq1_1), "eq")
}))
}()
})
	})
	return eq1SemigroupMap
}

var bindSemigroupMap gopurs_runtime.Value
var once_bindSemigroupMap sync.Once
func Get_bindSemigroupMap() gopurs_runtime.Value {
	once_bindSemigroupMap.Do(func() {
		bindSemigroupMap = gopurs_runtime.Func(func(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
return gopurs_runtime.Apply(pkg_Data_Map_Internal.Get_bindMap(), dictOrd_0)
}()
})
	})
	return bindSemigroupMap
}

var applySemigroupMap gopurs_runtime.Value
var once_applySemigroupMap sync.Once
func Get_applySemigroupMap() gopurs_runtime.Value {
	once_applySemigroupMap.Do(func() {
		applySemigroupMap = gopurs_runtime.Func(func(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
compare_1_0 := gopurs_runtime.RecordGet(dictOrd_0, "compare")
_ = compare_1_0
return gopurs_runtime.RecordDict2("apply", "Functor0", gopurs_runtime.Func2(func(m1_2 gopurs_runtime.Value, m2_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.UncurriedApp4(pkg_Data_Map_Internal.Get_unsafeIntersectionWith(), compare_1_0, gopurs_runtime.RecordGet(pkg_Control_Category.Get_categoryFn(), "identity"), m1_2, m2_3)
}), gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Map_Internal.Get_functorMap()
}))
}()
})
	})
	return applySemigroupMap
}

var altSemigroupMap gopurs_runtime.Value
var once_altSemigroupMap sync.Once
func Get_altSemigroupMap() gopurs_runtime.Value {
	once_altSemigroupMap.Do(func() {
		altSemigroupMap = gopurs_runtime.Func(func(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
compare_1_0 := gopurs_runtime.RecordGet(dictOrd_0, "compare")
_ = compare_1_0
return gopurs_runtime.RecordDict2("alt", "Functor0", gopurs_runtime.Func2(func(m1_2 gopurs_runtime.Value, m2_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.UncurriedApp4(pkg_Data_Map_Internal.Get_unsafeUnionWith(), compare_1_0, pkg_Data_Function.Get_const_(), m1_2, m2_3)
}), gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Map_Internal.Get_functorMap()
}))
}()
})
	})
	return altSemigroupMap
}

func Call_showSemigroupMap(dictShow_0_loop gopurs_runtime.Value, dictShow1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictShow_0 gopurs_runtime.Value = dictShow_0_loop
_ = dictShow_0
var dictShow1_1 gopurs_runtime.Value = dictShow1_1_loop
_ = dictShow1_1
return gopurs_runtime.Apply2(pkg_Data_Map_Internal.Get_showMap(), dictShow_0, dictShow1_1)
}

func Call_eqSemigroupMap(dictEq_0_loop gopurs_runtime.Value, dictEq1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictEq_0 gopurs_runtime.Value = dictEq_0_loop
_ = dictEq_0
var dictEq1_1 gopurs_runtime.Value = dictEq1_1_loop
_ = dictEq1_1
return gopurs_runtime.Apply2(pkg_Data_Map_Internal.Get_eqMap(), dictEq_0, dictEq1_1)
}


