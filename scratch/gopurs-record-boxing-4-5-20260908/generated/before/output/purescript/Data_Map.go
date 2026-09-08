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
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(Call_Data_Map_SemigroupMap(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](x_0_box)))}
})
	})
	return cache_Data_Map_SemigroupMap
}

var cache_Data_Map_traversableWithIndexSemigroupMap gopurs_runtime.Value
var once_Data_Map_traversableWithIndexSemigroupMap sync.Once
func Get_Data_Map_traversableWithIndexSemigroupMap() gopurs_runtime.Value {
	once_Data_Map_traversableWithIndexSemigroupMap.Do(func() {
		cache_Data_Map_traversableWithIndexSemigroupMap = gopurs_runtime.Value{Type: 9, IntVal: 2078610234, UnsafePtr: unsafe.Pointer(Rebox_Data_Map_4018835873_1812164904(Rebox_Data_Map_1812164904_4018835873(gopurs_runtime.CoerceToStruct[Constructor_Data_TraversableWithIndex_TraversableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]](Get_Data_Map_Internal_traversableWithIndexMap()))))}
	})
	return cache_Data_Map_traversableWithIndexSemigroupMap
}

var cache_Data_Map_traversableSemigroupMap gopurs_runtime.Value
var once_Data_Map_traversableSemigroupMap sync.Once
func Get_Data_Map_traversableSemigroupMap() gopurs_runtime.Value {
	once_Data_Map_traversableSemigroupMap.Do(func() {
		cache_Data_Map_traversableSemigroupMap = gopurs_runtime.Value{Type: 9, IntVal: 3941073978, UnsafePtr: unsafe.Pointer(Rebox_Data_Map_2141765991_3043886126(Rebox_Data_Map_3043886126_2141765991(gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]](Get_Data_Map_Internal_traversableMap()))))}
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
		cache_Data_Map_newtypeSemigroupMap = gopurs_runtime.Value{Type: 9, IntVal: 3322196858, UnsafePtr: unsafe.Pointer(Rebox_Data_Map_451298056_385277032((&Constructor_Data_Newtype_Newtype[*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value], *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{}
})})))}
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
// FALLBACK TCO: isLoop=false len=1
go__go_0_1_0 = gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t4 *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]
{
var __t_tag_2 *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](v_1)
if (__t_tag_2 == nil) {
__t4 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(nil))})
goto end_branch_4
} else {

}
}
{
var __t_tag_3 *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](v_1)
if (__t_tag_3 != nil) {
__t4 = (&Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V0, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V1, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V2, Get_Data_Unit_unit(), gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(go__go_0_1_0, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V4)})), gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(go__go_0_1_0, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V5)}))})
goto end_branch_4
} else {

}
}
{
__t4 = func() *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value] { panic("Failed pattern match") }()
}
end_branch_4:
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(__t4)}
})
// TAST (Let): __local_var_0_0 shape=LetRec(Other) bindingType=(Func [(ADT ["Data","Map","Internal","Map"] [(TypeVar k), (TypeVar v)])] (ADT ["Data","Map","Internal","Map"] [(TypeVar k), Unit]))
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
		cache_Data_Map_functorWithIndexSemigroupMap = gopurs_runtime.Value{Type: 9, IntVal: 4077743418, UnsafePtr: unsafe.Pointer(Rebox_Data_Map_1397444001_2412140840(Rebox_Data_Map_2412140840_1397444001(gopurs_runtime.CoerceToStruct[Constructor_Data_FunctorWithIndex_FunctorWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]](Get_Data_Map_Internal_functorWithIndexMap()))))}
	})
	return cache_Data_Map_functorWithIndexSemigroupMap
}

var cache_Data_Map_functorSemigroupMap gopurs_runtime.Value
var once_Data_Map_functorSemigroupMap sync.Once
func Get_Data_Map_functorSemigroupMap() gopurs_runtime.Value {
	once_Data_Map_functorSemigroupMap.Do(func() {
		cache_Data_Map_functorSemigroupMap = gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(Rebox_Data_Map_3281783655_2812149806(Rebox_Data_Map_2812149806_3281783655(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Data_Map_Internal_functorMap()))))}
	})
	return cache_Data_Map_functorSemigroupMap
}

var cache_Data_Map_foldableWithIndexSemigroupMap gopurs_runtime.Value
var once_Data_Map_foldableWithIndexSemigroupMap sync.Once
func Get_Data_Map_foldableWithIndexSemigroupMap() gopurs_runtime.Value {
	once_Data_Map_foldableWithIndexSemigroupMap.Do(func() {
		cache_Data_Map_foldableWithIndexSemigroupMap = gopurs_runtime.Value{Type: 9, IntVal: 74250362, UnsafePtr: unsafe.Pointer(Rebox_Data_Map_1245395425_3725484264(Rebox_Data_Map_3725484264_1245395425(gopurs_runtime.CoerceToStruct[Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]](Get_Data_Map_Internal_foldableWithIndexMap()))))}
	})
	return cache_Data_Map_foldableWithIndexSemigroupMap
}

var cache_Data_Map_foldableSemigroupMap gopurs_runtime.Value
var once_Data_Map_foldableSemigroupMap sync.Once
func Get_Data_Map_foldableSemigroupMap() gopurs_runtime.Value {
	once_Data_Map_foldableSemigroupMap.Do(func() {
		cache_Data_Map_foldableSemigroupMap = gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(Rebox_Data_Map_3596835815_1680800814(Rebox_Data_Map_1680800814_3596835815(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_Map_Internal_foldableMap()))))}
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

func Call_Data_Map_SemigroupMap(x_0_loop *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]) *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value] {
var x_0 *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value] = x_0_loop
_ = x_0
return x_0
}

func Call_Data_Map_showSemigroupMap(dictShow_0_loop gopurs_runtime.Value, dictShow1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictShow_0 gopurs_runtime.Value = dictShow_0_loop
_ = dictShow_0
var dictShow1_1 gopurs_runtime.Value = dictShow1_1_loop
_ = dictShow1_1
return gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Data_Map_2638796135_1386611502(Rebox_Data_Map_1386611502_2638796135(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](gopurs_runtime.Apply2(Get_Data_Map_Internal_showMap(), dictShow_0, dictShow1_1)))))}
}

func Call_Data_Map_semigroupSemigroupMap(dictOrd_0_loop gopurs_runtime.Value, dictSemigroup_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
var dictSemigroup_1 gopurs_runtime.Value = dictSemigroup_1_loop
_ = dictSemigroup_1
// TAST (Let): append_2_0 shape=Other bindingType=(Func [(TypeVar v), (TypeVar v)] (TypeVar v))
append_2_0 := gopurs_runtime.RecordGet(dictSemigroup_1, "append")
_ = append_2_0
return gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer(Rebox_Data_Map_3854229351_4179793454((&Constructor_Data_Semigroup_Semigroup[*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]]{1, gopurs_runtime.Func2(func(v_3 gopurs_runtime.Value, v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeUnionWith(), gopurs_runtime.RecordGet(dictOrd_0, "compare"), append_2_0, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](v_3))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](v1_4))})))}
})})))}
}

func Call_Data_Map_plusSemigroupMap(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
return gopurs_runtime.Value{Type: 9, IntVal: 3709470893, UnsafePtr: unsafe.Pointer(Rebox_Data_Map_4212347312_3706288089(Rebox_Data_Map_3706288089_4212347312(gopurs_runtime.CoerceToStruct[Constructor_Control_Plus_Plus[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Map_Internal_plusMap(), dictOrd_0)))))}
}

func Call_Data_Map_ordSemigroupMap(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
// TAST (Let): ordMap__193435443_1_0 shape=App(Var) bindingType=Any
ordMap__193435443_1_0 := gopurs_runtime.Apply(Get_Data_Map_Internal_ordMap(), dictOrd_0)
_ = ordMap__193435443_1_0
return gopurs_runtime.Func(func(dictOrd1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(Rebox_Data_Map_1910448679_4177771502(Rebox_Data_Map_4177771502_1910448679(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](gopurs_runtime.Apply(ordMap__193435443_1_0, dictOrd1_2)))))}
})
}

func Call_Data_Map_ord1SemigroupMap(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
return gopurs_runtime.Value{Type: 9, IntVal: 1632188299, UnsafePtr: unsafe.Pointer(Rebox_Data_Map_849406934_3985601471(Rebox_Data_Map_3985601471_849406934(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord1[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Map_Internal_ord1Map(), dictOrd_0)))))}
}

func Call_Data_Map_monoidSemigroupMap(dictOrd_0_loop gopurs_runtime.Value, dictSemigroup_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
var dictSemigroup_1 gopurs_runtime.Value = dictSemigroup_1_loop
_ = dictSemigroup_1
// TAST (Let): append_2_1 shape=Other bindingType=(Func [(TypeVar v), (TypeVar v)] (TypeVar v))
append_2_1 := gopurs_runtime.RecordGet(dictSemigroup_1, "append")
_ = append_2_1
// TAST (Let): semigroupSemigroupMap2_2_0 shape=Let(LitRecord) bindingType=(ADT ["Data","Semigroup","Semigroup"] [(ADT ["Data","Map","Internal","Map"] [(TypeVar k), (TypeVar v)])])
semigroupSemigroupMap2_2_0 := (&Constructor_Data_Semigroup_Semigroup[*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]]{1, gopurs_runtime.Func2(func(v_3 gopurs_runtime.Value, v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeUnionWith(), gopurs_runtime.RecordGet(dictOrd_0, "compare"), append_2_1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](v_3))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](v1_4))})))}
})})
_ = semigroupSemigroupMap2_2_0
return gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(Rebox_Data_Map_1291127239_1201789390((&Constructor_Data_Monoid_Monoid[*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(_dollar___unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer(Rebox_Data_Map_3854229351_4179793454(semigroupSemigroupMap2_2_0))}
}), (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(nil)})))}
}

func Call_Data_Map_eqSemigroupMap(dictEq_0_loop gopurs_runtime.Value, dictEq1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictEq_0 gopurs_runtime.Value = dictEq_0_loop
_ = dictEq_0
var dictEq1_1 gopurs_runtime.Value = dictEq1_1_loop
_ = dictEq1_1
return gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Data_Map_1444241223_3790796878(Rebox_Data_Map_3790796878_1444241223(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](gopurs_runtime.Apply2(Get_Data_Map_Internal_eqMap(), dictEq_0, dictEq1_1)))))}
}

func Call_Data_Map_eq1SemigroupMap(dictEq_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictEq_0 gopurs_runtime.Value = dictEq_0_loop
_ = dictEq_0
return gopurs_runtime.Value{Type: 9, IntVal: 1715248107, UnsafePtr: unsafe.Pointer(Rebox_Data_Map_3691144502_1766074591(Rebox_Data_Map_1766074591_3691144502(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq1[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Map_Internal_eq1Map(), dictEq_0)))))}
}

func Call_Data_Map_bindSemigroupMap(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(Rebox_Data_Map_1860323088_2748095225(Rebox_Data_Map_2748095225_1860323088(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Map_Internal_bindMap(), dictOrd_0)))))}
}

func Call_Data_Map_applySemigroupMap(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(Rebox_Data_Map_1628034448_3741347833(Rebox_Data_Map_3741347833_1628034448(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Map_Internal_applyMap(), dictOrd_0)))))}
}

func Call_Data_Map_altSemigroupMap(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
// TAST (Let): compare_1_0 shape=Other bindingType=(Func [(TypeVar k), (TypeVar k)] (ADT ["Data","Ordering","Ordering"] []))
compare_1_0 := gopurs_runtime.RecordGet(dictOrd_0, "compare")
_ = compare_1_0
return gopurs_runtime.Value{Type: 9, IntVal: 4060500237, UnsafePtr: unsafe.Pointer(Rebox_Data_Map_2686810384_3421983481((&Constructor_Control_Alt_Alt[*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(_dollar___unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(Rebox_Data_Map_3281783655_2812149806(Rebox_Data_Map_2812149806_3281783655(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Data_Map_Internal_functorMap()))))}
}), gopurs_runtime.Func2(func(m1_2 gopurs_runtime.Value, m2_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeUnionWith(), compare_1_0, Get_Data_Function_go__const(), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](m1_2))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](m2_3))})))}
})})))}
}

func Rebox_Data_Map_1245395425_3725484264(in *Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]]) *Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = in.V1
		out.V2 = in.V2
		out.V3 = in.V3
	return out
}

func Rebox_Data_Map_1291127239_1201789390(in *Constructor_Data_Monoid_Monoid[*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]]) *Constructor_Data_Monoid_Monoid[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(in.V1)}
	return out
}

func Rebox_Data_Map_1386611502_2638796135(in *Constructor_Data_Show_Show[gopurs_runtime.Value]) *Constructor_Data_Show_Show[*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]] {
	if in == nil { return nil }
	out := &Constructor_Data_Show_Show[*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]]{}
		out.V0 = in.V0
	return out
}

func Rebox_Data_Map_1397444001_2412140840(in *Constructor_Data_FunctorWithIndex_FunctorWithIndex[gopurs_runtime.Value, *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]]) *Constructor_Data_FunctorWithIndex_FunctorWithIndex[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_FunctorWithIndex_FunctorWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = in.V1
	return out
}

func Rebox_Data_Map_1444241223_3790796878(in *Constructor_Data_Eq_Eq[*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]]) *Constructor_Data_Eq_Eq[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Eq_Eq[gopurs_runtime.Value]{}
		out.V0 = in.V0
	return out
}

func Rebox_Data_Map_1628034448_3741347833(in *Constructor_Control_Apply_Apply[*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]]) *Constructor_Control_Apply_Apply[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Control_Apply_Apply[gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = in.V1
	return out
}

func Rebox_Data_Map_1680800814_3596835815(in *Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]) *Constructor_Data_Foldable_Foldable[*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]] {
	if in == nil { return nil }
	out := &Constructor_Data_Foldable_Foldable[*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]]{}
		out.V0 = in.V0
		out.V1 = in.V1
		out.V2 = in.V2
	return out
}

func Rebox_Data_Map_1766074591_3691144502(in *Constructor_Data_Eq_Eq1[gopurs_runtime.Value]) *Constructor_Data_Eq_Eq1[*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]] {
	if in == nil { return nil }
	out := &Constructor_Data_Eq_Eq1[*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]]{}
		out.V0 = in.V0
	return out
}

func Rebox_Data_Map_1812164904_4018835873(in *Constructor_Data_TraversableWithIndex_TraversableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]) *Constructor_Data_TraversableWithIndex_TraversableWithIndex[gopurs_runtime.Value, *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]] {
	if in == nil { return nil }
	out := &Constructor_Data_TraversableWithIndex_TraversableWithIndex[gopurs_runtime.Value, *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]]{}
		out.V0 = in.V0
		out.V1 = in.V1
		out.V2 = in.V2
		out.V3 = in.V3
	return out
}

func Rebox_Data_Map_1860323088_2748095225(in *Constructor_Control_Bind_Bind[*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]]) *Constructor_Control_Bind_Bind[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Control_Bind_Bind[gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = in.V1
	return out
}

func Rebox_Data_Map_1910448679_4177771502(in *Constructor_Data_Ord_Ord[*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]]) *Constructor_Data_Ord_Ord[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Ord_Ord[gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = in.V1
	return out
}

func Rebox_Data_Map_2141765991_3043886126(in *Constructor_Data_Traversable_Traversable[*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]]) *Constructor_Data_Traversable_Traversable[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = in.V1
		out.V2 = in.V2
		out.V3 = in.V3
	return out
}

func Rebox_Data_Map_2412140840_1397444001(in *Constructor_Data_FunctorWithIndex_FunctorWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]) *Constructor_Data_FunctorWithIndex_FunctorWithIndex[gopurs_runtime.Value, *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]] {
	if in == nil { return nil }
	out := &Constructor_Data_FunctorWithIndex_FunctorWithIndex[gopurs_runtime.Value, *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]]{}
		out.V0 = in.V0
		out.V1 = in.V1
	return out
}

func Rebox_Data_Map_2638796135_1386611502(in *Constructor_Data_Show_Show[*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]]) *Constructor_Data_Show_Show[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Show_Show[gopurs_runtime.Value]{}
		out.V0 = in.V0
	return out
}

func Rebox_Data_Map_2686810384_3421983481(in *Constructor_Control_Alt_Alt[*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]]) *Constructor_Control_Alt_Alt[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Control_Alt_Alt[gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = in.V1
	return out
}

func Rebox_Data_Map_2748095225_1860323088(in *Constructor_Control_Bind_Bind[gopurs_runtime.Value]) *Constructor_Control_Bind_Bind[*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]] {
	if in == nil { return nil }
	out := &Constructor_Control_Bind_Bind[*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]]{}
		out.V0 = in.V0
		out.V1 = in.V1
	return out
}

func Rebox_Data_Map_2812149806_3281783655(in *Constructor_Data_Functor_Functor[gopurs_runtime.Value]) *Constructor_Data_Functor_Functor[*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]] {
	if in == nil { return nil }
	out := &Constructor_Data_Functor_Functor[*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]]{}
		out.V0 = in.V0
	return out
}

func Rebox_Data_Map_3043886126_2141765991(in *Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]) *Constructor_Data_Traversable_Traversable[*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]] {
	if in == nil { return nil }
	out := &Constructor_Data_Traversable_Traversable[*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]]{}
		out.V0 = in.V0
		out.V1 = in.V1
		out.V2 = in.V2
		out.V3 = in.V3
	return out
}

func Rebox_Data_Map_3281783655_2812149806(in *Constructor_Data_Functor_Functor[*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]]) *Constructor_Data_Functor_Functor[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Functor_Functor[gopurs_runtime.Value]{}
		out.V0 = in.V0
	return out
}

func Rebox_Data_Map_3596835815_1680800814(in *Constructor_Data_Foldable_Foldable[*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]]) *Constructor_Data_Foldable_Foldable[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = in.V1
		out.V2 = in.V2
	return out
}

func Rebox_Data_Map_3691144502_1766074591(in *Constructor_Data_Eq_Eq1[*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]]) *Constructor_Data_Eq_Eq1[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Eq_Eq1[gopurs_runtime.Value]{}
		out.V0 = in.V0
	return out
}

func Rebox_Data_Map_3706288089_4212347312(in *Constructor_Control_Plus_Plus[gopurs_runtime.Value]) *Constructor_Control_Plus_Plus[*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]] {
	if in == nil { return nil }
	out := &Constructor_Control_Plus_Plus[*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]]{}
		out.V0 = in.V0
		out.V1 = in.V1
	return out
}

func Rebox_Data_Map_3725484264_1245395425(in *Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]) *Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]] {
	if in == nil { return nil }
	out := &Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]]{}
		out.V0 = in.V0
		out.V1 = in.V1
		out.V2 = in.V2
		out.V3 = in.V3
	return out
}

func Rebox_Data_Map_3741347833_1628034448(in *Constructor_Control_Apply_Apply[gopurs_runtime.Value]) *Constructor_Control_Apply_Apply[*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]] {
	if in == nil { return nil }
	out := &Constructor_Control_Apply_Apply[*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]]{}
		out.V0 = in.V0
		out.V1 = in.V1
	return out
}

func Rebox_Data_Map_3790796878_1444241223(in *Constructor_Data_Eq_Eq[gopurs_runtime.Value]) *Constructor_Data_Eq_Eq[*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]] {
	if in == nil { return nil }
	out := &Constructor_Data_Eq_Eq[*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]]{}
		out.V0 = in.V0
	return out
}

func Rebox_Data_Map_3854229351_4179793454(in *Constructor_Data_Semigroup_Semigroup[*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]]) *Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]{}
		out.V0 = in.V0
	return out
}

func Rebox_Data_Map_3985601471_849406934(in *Constructor_Data_Ord_Ord1[gopurs_runtime.Value]) *Constructor_Data_Ord_Ord1[*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]] {
	if in == nil { return nil }
	out := &Constructor_Data_Ord_Ord1[*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]]{}
		out.V0 = in.V0
		out.V1 = in.V1
	return out
}

func Rebox_Data_Map_4018835873_1812164904(in *Constructor_Data_TraversableWithIndex_TraversableWithIndex[gopurs_runtime.Value, *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]]) *Constructor_Data_TraversableWithIndex_TraversableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_TraversableWithIndex_TraversableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = in.V1
		out.V2 = in.V2
		out.V3 = in.V3
	return out
}

func Rebox_Data_Map_4177771502_1910448679(in *Constructor_Data_Ord_Ord[gopurs_runtime.Value]) *Constructor_Data_Ord_Ord[*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]] {
	if in == nil { return nil }
	out := &Constructor_Data_Ord_Ord[*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]]{}
		out.V0 = in.V0
		out.V1 = in.V1
	return out
}

func Rebox_Data_Map_4212347312_3706288089(in *Constructor_Control_Plus_Plus[*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]]) *Constructor_Control_Plus_Plus[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Control_Plus_Plus[gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = in.V1
	return out
}

func Rebox_Data_Map_451298056_385277032(in *Constructor_Data_Newtype_Newtype[*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value], *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]]) *Constructor_Data_Newtype_Newtype[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Newtype_Newtype[gopurs_runtime.Value, gopurs_runtime.Value]{}
		out.V0 = in.V0
	return out
}

func Rebox_Data_Map_849406934_3985601471(in *Constructor_Data_Ord_Ord1[*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]]) *Constructor_Data_Ord_Ord1[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Ord_Ord1[gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = in.V1
	return out
}


