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
		cache_Data_Map_traversableWithIndexSemigroupMap = gopurs_runtime.Value{Type: 9, IntVal: 2078610234, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_TraversableWithIndex_TraversableWithIndex](Get_Data_Map_Internal_traversableWithIndexMap()))}
	})
	return cache_Data_Map_traversableWithIndexSemigroupMap
}

var cache_Data_Map_traversableSemigroupMap gopurs_runtime.Value
var once_Data_Map_traversableSemigroupMap sync.Once
func Get_Data_Map_traversableSemigroupMap() gopurs_runtime.Value {
	once_Data_Map_traversableSemigroupMap.Do(func() {
		cache_Data_Map_traversableSemigroupMap = gopurs_runtime.Value{Type: 9, IntVal: 3941073978, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable](Get_Data_Map_Internal_traversableMap()))}
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
		cache_Data_Map_newtypeSemigroupMap = gopurs_runtime.Value{Type: 9, IntVal: 3322196858, UnsafePtr: unsafe.Pointer(&Constructor_Data_Newtype_Newtype{1, gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{}
})})}
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
var go__go_0_1_6 gopurs_runtime.Value
_ = go__go_0_1_6
go__go_0_1_6 = gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
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
__t2 = &Constructor_Data_Map_Internal_Node{1, (*Constructor_Data_Map_Internal_Node)(v_1.UnsafePtr).V0, (*Constructor_Data_Map_Internal_Node)(v_1.UnsafePtr).V1, (*Constructor_Data_Map_Internal_Node)(v_1.UnsafePtr).V2, Get_Data_Unit_unit(), gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.Apply(go__go_0_1_6, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v_1.UnsafePtr).V4)})), gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.Apply(go__go_0_1_6, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v_1.UnsafePtr).V5)}))}
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
__local_var_0_0 := go__go_0_1_6
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
		cache_Data_Map_functorWithIndexSemigroupMap = gopurs_runtime.Value{Type: 9, IntVal: 4077743418, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_FunctorWithIndex_FunctorWithIndex](Get_Data_Map_Internal_functorWithIndexMap()))}
	})
	return cache_Data_Map_functorWithIndexSemigroupMap
}

var cache_Data_Map_functorSemigroupMap gopurs_runtime.Value
var once_Data_Map_functorSemigroupMap sync.Once
func Get_Data_Map_functorSemigroupMap() gopurs_runtime.Value {
	once_Data_Map_functorSemigroupMap.Do(func() {
		cache_Data_Map_functorSemigroupMap = gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](Get_Data_Map_Internal_functorMap()))}
	})
	return cache_Data_Map_functorSemigroupMap
}

var cache_Data_Map_foldableWithIndexSemigroupMap gopurs_runtime.Value
var once_Data_Map_foldableWithIndexSemigroupMap sync.Once
func Get_Data_Map_foldableWithIndexSemigroupMap() gopurs_runtime.Value {
	once_Data_Map_foldableWithIndexSemigroupMap.Do(func() {
		cache_Data_Map_foldableWithIndexSemigroupMap = gopurs_runtime.Value{Type: 9, IntVal: 74250362, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_FoldableWithIndex_FoldableWithIndex](Get_Data_Map_Internal_foldableWithIndexMap()))}
	})
	return cache_Data_Map_foldableWithIndexSemigroupMap
}

var cache_Data_Map_foldableSemigroupMap gopurs_runtime.Value
var once_Data_Map_foldableSemigroupMap sync.Once
func Get_Data_Map_foldableSemigroupMap() gopurs_runtime.Value {
	once_Data_Map_foldableSemigroupMap.Do(func() {
		cache_Data_Map_foldableSemigroupMap = gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable](Get_Data_Map_Internal_foldableMap()))}
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
// TAST (Let): showArray_2_0 -> *Constructor_Data_Show_Show
showArray_2_0 := &Constructor_Data_Show_Show{1, gopurs_runtime.Apply(Get_Data_Show_showArrayImpl(), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str((((("(Tuple ") + (gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictShow_0, "show"), (*Constructor_Data_Tuple_Tuple)(v_2.UnsafePtr).V0).StrVal())) + (" ")) + (gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictShow1_1, "show"), (*Constructor_Data_Tuple_Tuple)(v_2.UnsafePtr).V1).StrVal())) + (")"))
}))}
_ = showArray_2_0
return gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(&Constructor_Data_Show_Show{1, gopurs_runtime.Func(func(as_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str((("(fromFoldable ") + (gopurs_runtime.Apply(gopurs_runtime.Box(showArray_2_0.V0), func() gopurs_runtime.Value {
					arr := func() []*Constructor_Data_Tuple_Tuple {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply6(Get_Data_Unfoldable_unfoldrArrayImpl(), Get_Data_Maybe_isNothing(), gopurs_runtime.Apply(Get_Partial_Unsafe__unsafePartial(), gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if (v_5.Type == 9 && v_5.IntVal == 930809136 && v_5.UnsafePtr != nil) {
__t1 = (*Constructor_Data_Maybe_Just)(v_5.UnsafePtr).V0
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
})), Get_Data_Tuple_fst(), Get_Data_Tuple_snd(), Get_Data_Map_Internal_stepUnfoldr(), gopurs_runtime.Value{Type: 9, IntVal: 2861335956, UnsafePtr: unsafe.Pointer(&Constructor_Data_Map_Internal_IterNode{1, gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](as_3), gopurs_runtime.Value{Type: 9, IntVal: 2509360378, UnsafePtr: unsafe.Pointer(nil)}})}).UnsafePtr)
					unboxed := make([]*Constructor_Data_Tuple_Tuple, len(arr))
					for i, v := range arr { unboxed[i] = gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple](v) }
					return unboxed
				}()
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v)} }
					return gopurs_runtime.Array(boxed)
				}()).StrVal())) + (")"))
})})}
}

func Call_Data_Map_semigroupSemigroupMap(dictOrd_0_loop gopurs_runtime.Value, dictSemigroup_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
var dictSemigroup_1 gopurs_runtime.Value = dictSemigroup_1_loop
_ = dictSemigroup_1
// TAST (Let): append_2_0 -> gopurs_runtime.Value
append_2_0 := gopurs_runtime.RecordGet(dictSemigroup_1, "append")
_ = append_2_0
return gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer(&Constructor_Data_Semigroup_Semigroup{1, gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeUnionWith(), gopurs_runtime.RecordGet(dictOrd_0, "compare"), append_2_0, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](v_3))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](v1_4))})))}
})
})})}
}

func Call_Data_Map_plusSemigroupMap(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
// TAST (Let): compare_1_1 -> gopurs_runtime.Value
compare_1_1 := gopurs_runtime.RecordGet(dictOrd_0, "compare")
_ = compare_1_1
// TAST (Let): altMap1_1_0 -> *Constructor_Control_Alt_Alt
altMap1_1_0 := &Constructor_Control_Alt_Alt{1, gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](Get_Data_Map_Internal_functorMap()))}
}), gopurs_runtime.Func(func(m1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m2_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeUnionWith(), compare_1_1, Get_Data_Function_go__const(), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](m1_2))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](m2_3))})))}
})
})}
_ = altMap1_1_0
return gopurs_runtime.Value{Type: 9, IntVal: 3709470893, UnsafePtr: unsafe.Pointer(&Constructor_Control_Plus_Plus{1, gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4060500237, UnsafePtr: unsafe.Pointer(altMap1_1_0)}
}), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(nil))}})}
}

func Call_Data_Map_ordSemigroupMap(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
// TAST (Let): __local_var_1_0 -> gopurs_runtime.Value
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd_0, "Eq0"), gopurs_runtime.Value{})
_ = __local_var_1_0
// TAST (Let): __local_var_2_1 -> gopurs_runtime.Value
__local_var_2_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd_0, "Eq0"), gopurs_runtime.Value{})
_ = __local_var_2_1
return gopurs_runtime.Func(func(dictOrd1_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_4_4 -> gopurs_runtime.Value
__local_var_4_4 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd1_3, "Eq0"), gopurs_runtime.Value{})
_ = __local_var_4_4
var go__go_5_5_0 gopurs_runtime.Value
go__go_5_5_0 = gopurs_runtime.Func(func(a_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_7_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var a_6_loop gopurs_runtime.Value = a_6_loop_val
var b_7_loop gopurs_runtime.Value = b_7_loop_val
go__go_5_5_0:
for {
if false { continue go__go_5_5_0 }
var a_6 gopurs_runtime.Value = a_6_loop
_ = a_6
var b_7 gopurs_runtime.Value = b_7_loop
_ = b_7
// TAST (Let): v_8_6 -> *Constructor_Data_Map_Internal_IterNext
v_8_6 := gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_IterNext](gopurs_runtime.Apply(Get_Data_Map_Internal_stepAsc(), a_6))
_ = v_8_6
var __t9 bool
{
if (v_8_6 != nil) {
// TAST (Let): v2_9_7 -> *Constructor_Data_Map_Internal_IterNext
v2_9_7 := gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_IterNext](gopurs_runtime.Apply(Get_Data_Map_Internal_stepAsc(), b_7))
_ = v2_9_7
var __t8 bool
{
if ((v2_9_7 != nil)) && (((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_0, "eq"), (v_8_6).V0, (v2_9_7).V0).IntVal) != (0)) && ((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_4_4, "eq"), (v_8_6).V1, (v2_9_7).V1).IntVal) != (0))) {
a_6_loop = (v_8_6).V2
b_7_loop = (v2_9_7).V2
continue go__go_5_5_0
__t8 = (gopurs_runtime.Value{}.IntVal) != (0)
goto end_branch_8
} else {

}
}
{
__t8 = false
}
end_branch_8:
__t9 = __t8
goto end_branch_9
} else {

}
}
{
if (v_8_6 == nil) {
__t9 = true
goto end_branch_9
} else {

}
}
{
__t9 = (func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal) != (0)
}
end_branch_9:
return gopurs_runtime.Bool(__t9)
}
}()
})
})
// TAST (Let): eqMapIter2_4_3 -> *Constructor_Data_Eq_Eq
eqMapIter2_4_3 := gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq](gopurs_runtime.RecordDict1("eq", go__go_5_5_0))
_ = eqMapIter2_4_3
var go__go_5_10_1 gopurs_runtime.Value
go__go_5_10_1 = gopurs_runtime.Func(func(a_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_7_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var a_6_loop gopurs_runtime.Value = a_6_loop_val
var b_7_loop gopurs_runtime.Value = b_7_loop_val
go__go_5_10_1:
for {
if false { continue go__go_5_10_1 }
var a_6 gopurs_runtime.Value = a_6_loop
_ = a_6
var b_7 gopurs_runtime.Value = b_7_loop
_ = b_7
// TAST (Let): v_8_11 -> *Constructor_Data_Map_Internal_IterNext
v_8_11 := gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_IterNext](gopurs_runtime.Apply(Get_Data_Map_Internal_stepAsc(), b_7))
_ = v_8_11
// TAST (Let): v1_9_12 -> *Constructor_Data_Map_Internal_IterNext
v1_9_12 := gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_IterNext](gopurs_runtime.Apply(Get_Data_Map_Internal_stepAsc(), a_6))
_ = v1_9_12
var __t19 uint32
{
if (v1_9_12 != nil) {
var __t17 uint32
{
if (v_8_11 != nil) {
// TAST (Let): v3_10_13 -> gopurs_runtime.Value
v3_10_13 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictOrd_0, "compare"), (v1_9_12).V0, (v_8_11).V0)
_ = v3_10_13
var __t16 uint32
{
if (uint32(v3_10_13.IntVal) == 902936544) {
// TAST (Let): v4_11_14 -> gopurs_runtime.Value
v4_11_14 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictOrd1_3, "compare"), (v1_9_12).V1, (v_8_11).V1)
_ = v4_11_14
var __t15 uint32
{
if (uint32(v4_11_14.IntVal) == 902936544) {
a_6_loop = (v1_9_12).V2
b_7_loop = (v_8_11).V2
continue go__go_5_10_1
__t15 = uint32(gopurs_runtime.Value{}.IntVal)
goto end_branch_15
} else {

}
}
{
__t15 = uint32(v4_11_14.IntVal)
}
end_branch_15:
__t16 = __t15
goto end_branch_16
} else {

}
}
{
__t16 = uint32(v3_10_13.IntVal)
}
end_branch_16:
__t17 = __t16
goto end_branch_17
} else {

}
}
{
if (v_8_11 == nil) {
__t17 = 380165415
goto end_branch_17
} else {

}
}
{
__t17 = uint32(func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal)
}
end_branch_17:
__t19 = __t17
goto end_branch_19
} else {

}
}
{
if (v1_9_12 == nil) {
var __t18 uint32
{
if (v_8_11 == nil) {
__t18 = 902936544
goto end_branch_18
} else {

}
}
{
__t18 = 1527465420
}
end_branch_18:
__t19 = __t18
goto end_branch_19
} else {

}
}
{
if (v_8_11 == nil) {
__t19 = 380165415
goto end_branch_19
} else {

}
}
{
__t19 = uint32(func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal)
}
end_branch_19:
return gopurs_runtime.Value{Type: 9, IntVal: int64(__t19), UnsafePtr: nil}
}
}()
})
})
// TAST (Let): ordMapIter2_4_2 -> *Constructor_Data_Ord_Ord
ordMapIter2_4_2 := &Constructor_Data_Ord_Ord{1, gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(eqMapIter2_4_3)}
}), go__go_5_10_1}
_ = ordMapIter2_4_2
// TAST (Let): __local_var_5_21 -> gopurs_runtime.Value
__local_var_5_21 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd1_3, "Eq0"), gopurs_runtime.Value{})
_ = __local_var_5_21
var go__go_6_23_2 gopurs_runtime.Value
go__go_6_23_2 = gopurs_runtime.Func(func(a_7_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_8_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var a_7_loop gopurs_runtime.Value = a_7_loop_val
var b_8_loop gopurs_runtime.Value = b_8_loop_val
go__go_6_23_2:
for {
if false { continue go__go_6_23_2 }
var a_7 gopurs_runtime.Value = a_7_loop
_ = a_7
var b_8 gopurs_runtime.Value = b_8_loop
_ = b_8
// TAST (Let): v_9_24 -> *Constructor_Data_Map_Internal_IterNext
v_9_24 := gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_IterNext](gopurs_runtime.Apply(Get_Data_Map_Internal_stepAsc(), a_7))
_ = v_9_24
var __t27 bool
{
if (v_9_24 != nil) {
// TAST (Let): v2_10_25 -> *Constructor_Data_Map_Internal_IterNext
v2_10_25 := gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_IterNext](gopurs_runtime.Apply(Get_Data_Map_Internal_stepAsc(), b_8))
_ = v2_10_25
var __t26 bool
{
if ((v2_10_25 != nil)) && (((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_1, "eq"), (v_9_24).V0, (v2_10_25).V0).IntVal) != (0)) && ((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_5_21, "eq"), (v_9_24).V1, (v2_10_25).V1).IntVal) != (0))) {
a_7_loop = (v_9_24).V2
b_8_loop = (v2_10_25).V2
continue go__go_6_23_2
__t26 = (gopurs_runtime.Value{}.IntVal) != (0)
goto end_branch_26
} else {

}
}
{
__t26 = false
}
end_branch_26:
__t27 = __t26
goto end_branch_27
} else {

}
}
{
if (v_9_24 == nil) {
__t27 = true
goto end_branch_27
} else {

}
}
{
__t27 = (func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal) != (0)
}
end_branch_27:
return gopurs_runtime.Bool(__t27)
}
}()
})
})
// TAST (Let): eqMapIter2_6_22 -> *Constructor_Data_Eq_Eq
eqMapIter2_6_22 := &Constructor_Data_Eq_Eq{1, go__go_6_23_2}
_ = eqMapIter2_6_22
// TAST (Let): eqMap2_5_20 -> *Constructor_Data_Eq_Eq
eqMap2_5_20 := &Constructor_Data_Eq_Eq{1, gopurs_runtime.Func(func(xs_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(ys_8 gopurs_runtime.Value) gopurs_runtime.Value {
var __t30 bool
{
if (xs_7.Type == 9 && xs_7.IntVal == 324739070 && xs_7.UnsafePtr == nil) {
var __t28 bool
{
if (ys_8.Type == 9 && ys_8.IntVal == 324739070 && ys_8.UnsafePtr == nil) {
__t28 = true
goto end_branch_28
} else {

}
}
{
__t28 = false
}
end_branch_28:
__t30 = __t28
goto end_branch_30
} else {

}
}
{
if (xs_7.Type == 9 && xs_7.IntVal == 324739070 && xs_7.UnsafePtr != nil) {
var __t29 bool
{
if ((ys_8.Type == 9 && ys_8.IntVal == 324739070 && ys_8.UnsafePtr != nil)) && (((*Constructor_Data_Map_Internal_Node)(xs_7.UnsafePtr).V1) == ((*Constructor_Data_Map_Internal_Node)(ys_8.UnsafePtr).V1)) {
__t29 = (gopurs_runtime.Apply2(gopurs_runtime.Box(eqMapIter2_6_22.V0), gopurs_runtime.Value{Type: 9, IntVal: 2861335956, UnsafePtr: unsafe.Pointer(&Constructor_Data_Map_Internal_IterNode{1, gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](xs_7), gopurs_runtime.Value{Type: 9, IntVal: 2509360378, UnsafePtr: unsafe.Pointer(nil)}})}, gopurs_runtime.Value{Type: 9, IntVal: 2861335956, UnsafePtr: unsafe.Pointer(&Constructor_Data_Map_Internal_IterNode{1, gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](ys_8), gopurs_runtime.Value{Type: 9, IntVal: 2509360378, UnsafePtr: unsafe.Pointer(nil)}})}).IntVal) != (0)
goto end_branch_29
} else {

}
}
{
__t29 = false
}
end_branch_29:
__t30 = __t29
goto end_branch_30
} else {

}
}
{
__t30 = (func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal) != (0)
}
end_branch_30:
return gopurs_runtime.Bool(__t30)
})
})}
_ = eqMap2_5_20
return gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(&Constructor_Data_Ord_Ord{1, gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(eqMap2_5_20)}
}), gopurs_runtime.Func(func(xs_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(ys_7 gopurs_runtime.Value) gopurs_runtime.Value {
var __t33 uint32
{
if (xs_6.Type == 9 && xs_6.IntVal == 324739070 && xs_6.UnsafePtr == nil) {
var __t32 uint32
{
if (ys_7.Type == 9 && ys_7.IntVal == 324739070 && ys_7.UnsafePtr == nil) {
__t32 = 902936544
goto end_branch_32
} else {

}
}
{
__t32 = 1527465420
}
end_branch_32:
__t33 = __t32
goto end_branch_33
} else {

}
}
{
var __t31 uint32
{
if (ys_7.Type == 9 && ys_7.IntVal == 324739070 && ys_7.UnsafePtr == nil) {
__t31 = 380165415
goto end_branch_31
} else {

}
}
{
__t31 = uint32(gopurs_runtime.Apply2(gopurs_runtime.Box(ordMapIter2_4_2.V1), gopurs_runtime.Value{Type: 9, IntVal: 2861335956, UnsafePtr: unsafe.Pointer(&Constructor_Data_Map_Internal_IterNode{1, gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](xs_6), gopurs_runtime.Value{Type: 9, IntVal: 2509360378, UnsafePtr: unsafe.Pointer(nil)}})}, gopurs_runtime.Value{Type: 9, IntVal: 2861335956, UnsafePtr: unsafe.Pointer(&Constructor_Data_Map_Internal_IterNode{1, gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](ys_7), gopurs_runtime.Value{Type: 9, IntVal: 2509360378, UnsafePtr: unsafe.Pointer(nil)}})}).IntVal)
}
end_branch_31:
__t33 = __t31
}
end_branch_33:
return gopurs_runtime.Value{Type: 9, IntVal: int64(__t33), UnsafePtr: nil}
})
})})}
})
}

func Call_Data_Map_ord1SemigroupMap(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
// TAST (Let): __local_var_1_0 -> gopurs_runtime.Value
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd_0, "Eq0"), gopurs_runtime.Value{})
_ = __local_var_1_0
// TAST (Let): __local_var_2_2 -> gopurs_runtime.Value
__local_var_2_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd_0, "Eq0"), gopurs_runtime.Value{})
_ = __local_var_2_2
// TAST (Let): eq1Map1_2_1 -> *Constructor_Data_Eq_Eq1
eq1Map1_2_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq1](gopurs_runtime.RecordDict1("eq1", gopurs_runtime.Func(func(dictEq1_3 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_4_4_3 gopurs_runtime.Value
go__go_4_4_3 = gopurs_runtime.Func(func(a_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var a_5_loop gopurs_runtime.Value = a_5_loop_val
var b_6_loop gopurs_runtime.Value = b_6_loop_val
go__go_4_4_3:
for {
if false { continue go__go_4_4_3 }
var a_5 gopurs_runtime.Value = a_5_loop
_ = a_5
var b_6 gopurs_runtime.Value = b_6_loop
_ = b_6
// TAST (Let): v_7_5 -> *Constructor_Data_Map_Internal_IterNext
v_7_5 := gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_IterNext](gopurs_runtime.Apply(Get_Data_Map_Internal_stepAsc(), a_5))
_ = v_7_5
var __t8 bool
{
if (v_7_5 != nil) {
// TAST (Let): v2_8_6 -> *Constructor_Data_Map_Internal_IterNext
v2_8_6 := gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_IterNext](gopurs_runtime.Apply(Get_Data_Map_Internal_stepAsc(), b_6))
_ = v2_8_6
var __t7 bool
{
if ((v2_8_6 != nil)) && (((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_2, "eq"), (v_7_5).V0, (v2_8_6).V0).IntVal) != (0)) && ((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq1_3, "eq"), (v_7_5).V1, (v2_8_6).V1).IntVal) != (0))) {
a_5_loop = (v_7_5).V2
b_6_loop = (v2_8_6).V2
continue go__go_4_4_3
__t7 = (gopurs_runtime.Value{}.IntVal) != (0)
goto end_branch_7
} else {

}
}
{
__t7 = false
}
end_branch_7:
__t8 = __t7
goto end_branch_8
} else {

}
}
{
if (v_7_5 == nil) {
__t8 = true
goto end_branch_8
} else {

}
}
{
__t8 = (func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal) != (0)
}
end_branch_8:
return gopurs_runtime.Bool(__t8)
}
}()
})
})
// TAST (Let): eqMapIter2_4_3 -> *Constructor_Data_Eq_Eq
eqMapIter2_4_3 := &Constructor_Data_Eq_Eq{1, go__go_4_4_3}
_ = eqMapIter2_4_3
return gopurs_runtime.Func(func(xs_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(ys_6 gopurs_runtime.Value) gopurs_runtime.Value {
var __t11 bool
{
if (xs_5.Type == 9 && xs_5.IntVal == 324739070 && xs_5.UnsafePtr == nil) {
var __t9 bool
{
if (ys_6.Type == 9 && ys_6.IntVal == 324739070 && ys_6.UnsafePtr == nil) {
__t9 = true
goto end_branch_9
} else {

}
}
{
__t9 = false
}
end_branch_9:
__t11 = __t9
goto end_branch_11
} else {

}
}
{
if (xs_5.Type == 9 && xs_5.IntVal == 324739070 && xs_5.UnsafePtr != nil) {
var __t10 bool
{
if ((ys_6.Type == 9 && ys_6.IntVal == 324739070 && ys_6.UnsafePtr != nil)) && (((*Constructor_Data_Map_Internal_Node)(xs_5.UnsafePtr).V1) == ((*Constructor_Data_Map_Internal_Node)(ys_6.UnsafePtr).V1)) {
__t10 = (gopurs_runtime.Apply2(gopurs_runtime.Box(eqMapIter2_4_3.V0), gopurs_runtime.Value{Type: 9, IntVal: 2861335956, UnsafePtr: unsafe.Pointer(&Constructor_Data_Map_Internal_IterNode{1, gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](xs_5), gopurs_runtime.Value{Type: 9, IntVal: 2509360378, UnsafePtr: unsafe.Pointer(nil)}})}, gopurs_runtime.Value{Type: 9, IntVal: 2861335956, UnsafePtr: unsafe.Pointer(&Constructor_Data_Map_Internal_IterNode{1, gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](ys_6), gopurs_runtime.Value{Type: 9, IntVal: 2509360378, UnsafePtr: unsafe.Pointer(nil)}})}).IntVal) != (0)
goto end_branch_10
} else {

}
}
{
__t10 = false
}
end_branch_10:
__t11 = __t10
goto end_branch_11
} else {

}
}
{
__t11 = (func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal) != (0)
}
end_branch_11:
return gopurs_runtime.Bool(__t11)
})
})
})))
_ = eq1Map1_2_1
return gopurs_runtime.Value{Type: 9, IntVal: 1632188299, UnsafePtr: unsafe.Pointer(&Constructor_Data_Ord_Ord1{1, gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1715248107, UnsafePtr: unsafe.Pointer(eq1Map1_2_1)}
}), gopurs_runtime.Func(func(dictOrd1_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_4_14 -> gopurs_runtime.Value
__local_var_4_14 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd1_3, "Eq0"), gopurs_runtime.Value{})
_ = __local_var_4_14
var go__go_5_15_4 gopurs_runtime.Value
go__go_5_15_4 = gopurs_runtime.Func(func(a_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_7_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var a_6_loop gopurs_runtime.Value = a_6_loop_val
var b_7_loop gopurs_runtime.Value = b_7_loop_val
go__go_5_15_4:
for {
if false { continue go__go_5_15_4 }
var a_6 gopurs_runtime.Value = a_6_loop
_ = a_6
var b_7 gopurs_runtime.Value = b_7_loop
_ = b_7
// TAST (Let): v_8_16 -> *Constructor_Data_Map_Internal_IterNext
v_8_16 := gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_IterNext](gopurs_runtime.Apply(Get_Data_Map_Internal_stepAsc(), a_6))
_ = v_8_16
var __t19 bool
{
if (v_8_16 != nil) {
// TAST (Let): v2_9_17 -> *Constructor_Data_Map_Internal_IterNext
v2_9_17 := gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_IterNext](gopurs_runtime.Apply(Get_Data_Map_Internal_stepAsc(), b_7))
_ = v2_9_17
var __t18 bool
{
if ((v2_9_17 != nil)) && (((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_0, "eq"), (v_8_16).V0, (v2_9_17).V0).IntVal) != (0)) && ((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_4_14, "eq"), (v_8_16).V1, (v2_9_17).V1).IntVal) != (0))) {
a_6_loop = (v_8_16).V2
b_7_loop = (v2_9_17).V2
continue go__go_5_15_4
__t18 = (gopurs_runtime.Value{}.IntVal) != (0)
goto end_branch_18
} else {

}
}
{
__t18 = false
}
end_branch_18:
__t19 = __t18
goto end_branch_19
} else {

}
}
{
if (v_8_16 == nil) {
__t19 = true
goto end_branch_19
} else {

}
}
{
__t19 = (func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal) != (0)
}
end_branch_19:
return gopurs_runtime.Bool(__t19)
}
}()
})
})
// TAST (Let): eqMapIter2_4_13 -> *Constructor_Data_Eq_Eq
eqMapIter2_4_13 := gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq](gopurs_runtime.RecordDict1("eq", go__go_5_15_4))
_ = eqMapIter2_4_13
var go__go_5_20_5 gopurs_runtime.Value
go__go_5_20_5 = gopurs_runtime.Func(func(a_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_7_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var a_6_loop gopurs_runtime.Value = a_6_loop_val
var b_7_loop gopurs_runtime.Value = b_7_loop_val
go__go_5_20_5:
for {
if false { continue go__go_5_20_5 }
var a_6 gopurs_runtime.Value = a_6_loop
_ = a_6
var b_7 gopurs_runtime.Value = b_7_loop
_ = b_7
// TAST (Let): v_8_21 -> *Constructor_Data_Map_Internal_IterNext
v_8_21 := gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_IterNext](gopurs_runtime.Apply(Get_Data_Map_Internal_stepAsc(), b_7))
_ = v_8_21
// TAST (Let): v1_9_22 -> *Constructor_Data_Map_Internal_IterNext
v1_9_22 := gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_IterNext](gopurs_runtime.Apply(Get_Data_Map_Internal_stepAsc(), a_6))
_ = v1_9_22
var __t29 uint32
{
if (v1_9_22 != nil) {
var __t27 uint32
{
if (v_8_21 != nil) {
// TAST (Let): v3_10_23 -> gopurs_runtime.Value
v3_10_23 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictOrd_0, "compare"), (v1_9_22).V0, (v_8_21).V0)
_ = v3_10_23
var __t26 uint32
{
if (uint32(v3_10_23.IntVal) == 902936544) {
// TAST (Let): v4_11_24 -> gopurs_runtime.Value
v4_11_24 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictOrd1_3, "compare"), (v1_9_22).V1, (v_8_21).V1)
_ = v4_11_24
var __t25 uint32
{
if (uint32(v4_11_24.IntVal) == 902936544) {
a_6_loop = (v1_9_22).V2
b_7_loop = (v_8_21).V2
continue go__go_5_20_5
__t25 = uint32(gopurs_runtime.Value{}.IntVal)
goto end_branch_25
} else {

}
}
{
__t25 = uint32(v4_11_24.IntVal)
}
end_branch_25:
__t26 = __t25
goto end_branch_26
} else {

}
}
{
__t26 = uint32(v3_10_23.IntVal)
}
end_branch_26:
__t27 = __t26
goto end_branch_27
} else {

}
}
{
if (v_8_21 == nil) {
__t27 = 380165415
goto end_branch_27
} else {

}
}
{
__t27 = uint32(func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal)
}
end_branch_27:
__t29 = __t27
goto end_branch_29
} else {

}
}
{
if (v1_9_22 == nil) {
var __t28 uint32
{
if (v_8_21 == nil) {
__t28 = 902936544
goto end_branch_28
} else {

}
}
{
__t28 = 1527465420
}
end_branch_28:
__t29 = __t28
goto end_branch_29
} else {

}
}
{
if (v_8_21 == nil) {
__t29 = 380165415
goto end_branch_29
} else {

}
}
{
__t29 = uint32(func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal)
}
end_branch_29:
return gopurs_runtime.Value{Type: 9, IntVal: int64(__t29), UnsafePtr: nil}
}
}()
})
})
// TAST (Let): ordMapIter2_4_12 -> *Constructor_Data_Ord_Ord
ordMapIter2_4_12 := &Constructor_Data_Ord_Ord{1, gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(eqMapIter2_4_13)}
}), go__go_5_20_5}
_ = ordMapIter2_4_12
return gopurs_runtime.Func(func(xs_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(ys_6 gopurs_runtime.Value) gopurs_runtime.Value {
var __t32 uint32
{
if (xs_5.Type == 9 && xs_5.IntVal == 324739070 && xs_5.UnsafePtr == nil) {
var __t31 uint32
{
if (ys_6.Type == 9 && ys_6.IntVal == 324739070 && ys_6.UnsafePtr == nil) {
__t31 = 902936544
goto end_branch_31
} else {

}
}
{
__t31 = 1527465420
}
end_branch_31:
__t32 = __t31
goto end_branch_32
} else {

}
}
{
var __t30 uint32
{
if (ys_6.Type == 9 && ys_6.IntVal == 324739070 && ys_6.UnsafePtr == nil) {
__t30 = 380165415
goto end_branch_30
} else {

}
}
{
__t30 = uint32(gopurs_runtime.Apply2(gopurs_runtime.Box(ordMapIter2_4_12.V1), gopurs_runtime.Value{Type: 9, IntVal: 2861335956, UnsafePtr: unsafe.Pointer(&Constructor_Data_Map_Internal_IterNode{1, gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](xs_5), gopurs_runtime.Value{Type: 9, IntVal: 2509360378, UnsafePtr: unsafe.Pointer(nil)}})}, gopurs_runtime.Value{Type: 9, IntVal: 2861335956, UnsafePtr: unsafe.Pointer(&Constructor_Data_Map_Internal_IterNode{1, gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](ys_6), gopurs_runtime.Value{Type: 9, IntVal: 2509360378, UnsafePtr: unsafe.Pointer(nil)}})}).IntVal)
}
end_branch_30:
__t32 = __t30
}
end_branch_32:
return gopurs_runtime.Value{Type: 9, IntVal: int64(__t32), UnsafePtr: nil}
})
})
})})}
}

func Call_Data_Map_monoidSemigroupMap(dictOrd_0_loop gopurs_runtime.Value, dictSemigroup_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
var dictSemigroup_1 gopurs_runtime.Value = dictSemigroup_1_loop
_ = dictSemigroup_1
// TAST (Let): append_2_1 -> gopurs_runtime.Value
append_2_1 := gopurs_runtime.RecordGet(dictSemigroup_1, "append")
_ = append_2_1
// TAST (Let): semigroupSemigroupMap2_2_0 -> *Constructor_Data_Semigroup_Semigroup
semigroupSemigroupMap2_2_0 := &Constructor_Data_Semigroup_Semigroup{1, gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeUnionWith(), gopurs_runtime.RecordGet(dictOrd_0, "compare"), append_2_1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](v_3))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](v1_4))})))}
})
})}
_ = semigroupSemigroupMap2_2_0
return gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(&Constructor_Data_Monoid_Monoid{1, gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer(semigroupSemigroupMap2_2_0)}
}), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(nil))}})}
}

func Call_Data_Map_eqSemigroupMap(dictEq_0_loop gopurs_runtime.Value, dictEq1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictEq_0 gopurs_runtime.Value = dictEq_0_loop
_ = dictEq_0
var dictEq1_1 gopurs_runtime.Value = dictEq1_1_loop
_ = dictEq1_1
var go__go_2_1_7 gopurs_runtime.Value
go__go_2_1_7 = gopurs_runtime.Func(func(a_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var a_3_loop gopurs_runtime.Value = a_3_loop_val
var b_4_loop gopurs_runtime.Value = b_4_loop_val
go__go_2_1_7:
for {
if false { continue go__go_2_1_7 }
var a_3 gopurs_runtime.Value = a_3_loop
_ = a_3
var b_4 gopurs_runtime.Value = b_4_loop
_ = b_4
// TAST (Let): v_5_2 -> *Constructor_Data_Map_Internal_IterNext
v_5_2 := gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_IterNext](gopurs_runtime.Apply(Get_Data_Map_Internal_stepAsc(), a_3))
_ = v_5_2
var __t5 bool
{
if (v_5_2 != nil) {
// TAST (Let): v2_6_3 -> *Constructor_Data_Map_Internal_IterNext
v2_6_3 := gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_IterNext](gopurs_runtime.Apply(Get_Data_Map_Internal_stepAsc(), b_4))
_ = v2_6_3
var __t4 bool
{
if ((v2_6_3 != nil)) && (((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq_0, "eq"), (v_5_2).V0, (v2_6_3).V0).IntVal) != (0)) && ((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq1_1, "eq"), (v_5_2).V1, (v2_6_3).V1).IntVal) != (0))) {
a_3_loop = (v_5_2).V2
b_4_loop = (v2_6_3).V2
continue go__go_2_1_7
__t4 = (gopurs_runtime.Value{}.IntVal) != (0)
goto end_branch_4
} else {

}
}
{
__t4 = false
}
end_branch_4:
__t5 = __t4
goto end_branch_5
} else {

}
}
{
if (v_5_2 == nil) {
__t5 = true
goto end_branch_5
} else {

}
}
{
__t5 = (func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal) != (0)
}
end_branch_5:
return gopurs_runtime.Bool(__t5)
}
}()
})
})
// TAST (Let): eqMapIter2_2_0 -> *Constructor_Data_Eq_Eq
eqMapIter2_2_0 := &Constructor_Data_Eq_Eq{1, go__go_2_1_7}
_ = eqMapIter2_2_0
return gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(&Constructor_Data_Eq_Eq{1, gopurs_runtime.Func(func(xs_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(ys_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t8 bool
{
if (xs_3.Type == 9 && xs_3.IntVal == 324739070 && xs_3.UnsafePtr == nil) {
var __t6 bool
{
if (ys_4.Type == 9 && ys_4.IntVal == 324739070 && ys_4.UnsafePtr == nil) {
__t6 = true
goto end_branch_6
} else {

}
}
{
__t6 = false
}
end_branch_6:
__t8 = __t6
goto end_branch_8
} else {

}
}
{
if (xs_3.Type == 9 && xs_3.IntVal == 324739070 && xs_3.UnsafePtr != nil) {
var __t7 bool
{
if ((ys_4.Type == 9 && ys_4.IntVal == 324739070 && ys_4.UnsafePtr != nil)) && (((*Constructor_Data_Map_Internal_Node)(xs_3.UnsafePtr).V1) == ((*Constructor_Data_Map_Internal_Node)(ys_4.UnsafePtr).V1)) {
__t7 = (gopurs_runtime.Apply2(gopurs_runtime.Box(eqMapIter2_2_0.V0), gopurs_runtime.Value{Type: 9, IntVal: 2861335956, UnsafePtr: unsafe.Pointer(&Constructor_Data_Map_Internal_IterNode{1, gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](xs_3), gopurs_runtime.Value{Type: 9, IntVal: 2509360378, UnsafePtr: unsafe.Pointer(nil)}})}, gopurs_runtime.Value{Type: 9, IntVal: 2861335956, UnsafePtr: unsafe.Pointer(&Constructor_Data_Map_Internal_IterNode{1, gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](ys_4), gopurs_runtime.Value{Type: 9, IntVal: 2509360378, UnsafePtr: unsafe.Pointer(nil)}})}).IntVal) != (0)
goto end_branch_7
} else {

}
}
{
__t7 = false
}
end_branch_7:
__t8 = __t7
goto end_branch_8
} else {

}
}
{
__t8 = (func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal) != (0)
}
end_branch_8:
return gopurs_runtime.Bool(__t8)
})
})})}
}

func Call_Data_Map_eq1SemigroupMap(dictEq_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictEq_0 gopurs_runtime.Value = dictEq_0_loop
_ = dictEq_0
return gopurs_runtime.Value{Type: 9, IntVal: 1715248107, UnsafePtr: unsafe.Pointer(&Constructor_Data_Eq_Eq1{1, gopurs_runtime.Func(func(dictEq1_1 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_2_1_8 gopurs_runtime.Value
go__go_2_1_8 = gopurs_runtime.Func(func(a_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var a_3_loop gopurs_runtime.Value = a_3_loop_val
var b_4_loop gopurs_runtime.Value = b_4_loop_val
go__go_2_1_8:
for {
if false { continue go__go_2_1_8 }
var a_3 gopurs_runtime.Value = a_3_loop
_ = a_3
var b_4 gopurs_runtime.Value = b_4_loop
_ = b_4
// TAST (Let): v_5_2 -> *Constructor_Data_Map_Internal_IterNext
v_5_2 := gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_IterNext](gopurs_runtime.Apply(Get_Data_Map_Internal_stepAsc(), a_3))
_ = v_5_2
var __t5 bool
{
if (v_5_2 != nil) {
// TAST (Let): v2_6_3 -> *Constructor_Data_Map_Internal_IterNext
v2_6_3 := gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_IterNext](gopurs_runtime.Apply(Get_Data_Map_Internal_stepAsc(), b_4))
_ = v2_6_3
var __t4 bool
{
if ((v2_6_3 != nil)) && (((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq_0, "eq"), (v_5_2).V0, (v2_6_3).V0).IntVal) != (0)) && ((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq1_1, "eq"), (v_5_2).V1, (v2_6_3).V1).IntVal) != (0))) {
a_3_loop = (v_5_2).V2
b_4_loop = (v2_6_3).V2
continue go__go_2_1_8
__t4 = (gopurs_runtime.Value{}.IntVal) != (0)
goto end_branch_4
} else {

}
}
{
__t4 = false
}
end_branch_4:
__t5 = __t4
goto end_branch_5
} else {

}
}
{
if (v_5_2 == nil) {
__t5 = true
goto end_branch_5
} else {

}
}
{
__t5 = (func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal) != (0)
}
end_branch_5:
return gopurs_runtime.Bool(__t5)
}
}()
})
})
// TAST (Let): eqMapIter2_2_0 -> *Constructor_Data_Eq_Eq
eqMapIter2_2_0 := &Constructor_Data_Eq_Eq{1, go__go_2_1_8}
_ = eqMapIter2_2_0
return gopurs_runtime.Func(func(xs_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(ys_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t8 bool
{
if (xs_3.Type == 9 && xs_3.IntVal == 324739070 && xs_3.UnsafePtr == nil) {
var __t6 bool
{
if (ys_4.Type == 9 && ys_4.IntVal == 324739070 && ys_4.UnsafePtr == nil) {
__t6 = true
goto end_branch_6
} else {

}
}
{
__t6 = false
}
end_branch_6:
__t8 = __t6
goto end_branch_8
} else {

}
}
{
if (xs_3.Type == 9 && xs_3.IntVal == 324739070 && xs_3.UnsafePtr != nil) {
var __t7 bool
{
if ((ys_4.Type == 9 && ys_4.IntVal == 324739070 && ys_4.UnsafePtr != nil)) && (((*Constructor_Data_Map_Internal_Node)(xs_3.UnsafePtr).V1) == ((*Constructor_Data_Map_Internal_Node)(ys_4.UnsafePtr).V1)) {
__t7 = (gopurs_runtime.Apply2(gopurs_runtime.Box(eqMapIter2_2_0.V0), gopurs_runtime.Value{Type: 9, IntVal: 2861335956, UnsafePtr: unsafe.Pointer(&Constructor_Data_Map_Internal_IterNode{1, gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](xs_3), gopurs_runtime.Value{Type: 9, IntVal: 2509360378, UnsafePtr: unsafe.Pointer(nil)}})}, gopurs_runtime.Value{Type: 9, IntVal: 2861335956, UnsafePtr: unsafe.Pointer(&Constructor_Data_Map_Internal_IterNode{1, gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](ys_4), gopurs_runtime.Value{Type: 9, IntVal: 2509360378, UnsafePtr: unsafe.Pointer(nil)}})}).IntVal) != (0)
goto end_branch_7
} else {

}
}
{
__t7 = false
}
end_branch_7:
__t8 = __t7
goto end_branch_8
} else {

}
}
{
__t8 = (func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal) != (0)
}
end_branch_8:
return gopurs_runtime.Bool(__t8)
})
})
})})}
}

func Call_Data_Map_bindSemigroupMap(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
// TAST (Let): compare_1_1 -> gopurs_runtime.Value
compare_1_1 := gopurs_runtime.RecordGet(dictOrd_0, "compare")
_ = compare_1_1
// TAST (Let): applyMap1_1_0 -> *Constructor_Control_Apply_Apply
applyMap1_1_0 := &Constructor_Control_Apply_Apply{1, gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](Get_Data_Map_Internal_functorMap()))}
}), gopurs_runtime.Func(func(m1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m2_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeIntersectionWith(), compare_1_1, gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return x_4
}), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](m1_2))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](m2_3))})))}
})
})}
_ = applyMap1_1_0
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(&Constructor_Control_Bind_Bind{1, gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(applyMap1_1_0)}
}), gopurs_runtime.Func(func(m_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_4_2_9 gopurs_runtime.Value
_ = go__go_4_2_9
go__go_4_2_9 = gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
var __t10 *Constructor_Data_Map_Internal_Node
{
if (v_5.Type == 9 && v_5.IntVal == 324739070 && v_5.UnsafePtr == nil) {
__t10 = (*Constructor_Data_Map_Internal_Node)(nil)
goto end_branch_10
} else {

}
}
{
if (v_5.Type == 9 && v_5.IntVal == 324739070 && v_5.UnsafePtr != nil) {
// TAST (Let): __local_var_6_3 -> gopurs_runtime.Value
__local_var_6_3 := (*Constructor_Data_Map_Internal_Node)(v_5.UnsafePtr).V2
_ = __local_var_6_3
var go__go_7_5_10 gopurs_runtime.Value
go__go_7_5_10 = gopurs_runtime.Func(func(v_8_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_8_loop gopurs_runtime.Value = v_8_loop_val
go__go_7_5_10:
for {
if false { continue go__go_7_5_10 }
var v_8 gopurs_runtime.Value = v_8_loop
_ = v_8
var __t8 *Constructor_Data_Maybe_Just
{
if (v_8.Type == 9 && v_8.IntVal == 324739070 && v_8.UnsafePtr == nil) {
__t8 = (*Constructor_Data_Maybe_Just)(nil)
goto end_branch_8
} else {

}
}
{
if (v_8.Type == 9 && v_8.IntVal == 324739070 && v_8.UnsafePtr != nil) {
// TAST (Let): v1_9_6 -> gopurs_runtime.Value
v1_9_6 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictOrd_0, "compare"), __local_var_6_3, (*Constructor_Data_Map_Internal_Node)(v_8.UnsafePtr).V2)
_ = v1_9_6
var __t7 *Constructor_Data_Maybe_Just
{
if (uint32(v1_9_6.IntVal) == 1527465420) {
v_8_loop = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v_8.UnsafePtr).V4)}
continue go__go_7_5_10
__t7 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Value{})
goto end_branch_7
} else {

}
}
{
if (uint32(v1_9_6.IntVal) == 380165415) {
v_8_loop = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v_8.UnsafePtr).V5)}
continue go__go_7_5_10
__t7 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Value{})
goto end_branch_7
} else {

}
}
{
if (uint32(v1_9_6.IntVal) == 902936544) {
__t7 = &Constructor_Data_Maybe_Just{1, (*Constructor_Data_Map_Internal_Node)(v_8.UnsafePtr).V3}
goto end_branch_7
} else {

}
}
{
__t7 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_7:
__t8 = __t7
goto end_branch_8
} else {

}
}
{
__t8 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_8:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t8)}
}
}()
})
// TAST (Let): v2_7_4 -> gopurs_runtime.Value
v2_7_4 := gopurs_runtime.Apply(go__go_7_5_10, gopurs_runtime.Apply(f_3, (*Constructor_Data_Map_Internal_Node)(v_5.UnsafePtr).V3))
_ = v2_7_4
var __t9 *Constructor_Data_Map_Internal_Node
{
if (v2_7_4.Type == 9 && v2_7_4.IntVal == 930809136 && v2_7_4.UnsafePtr != nil) {
__t9 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeBalancedNode(), (*Constructor_Data_Map_Internal_Node)(v_5.UnsafePtr).V2, (*Constructor_Data_Maybe_Just)(v2_7_4.UnsafePtr).V0, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.Apply(go__go_4_2_9, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v_5.UnsafePtr).V4)})))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.Apply(go__go_4_2_9, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v_5.UnsafePtr).V5)})))}))
goto end_branch_9
} else {

}
}
{
if (v2_7_4.Type == 9 && v2_7_4.IntVal == 930809136 && v2_7_4.UnsafePtr == nil) {
__t9 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.UncurriedApp2(Get_Data_Map_Internal_unsafeJoinNodes(), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.Apply(go__go_4_2_9, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v_5.UnsafePtr).V4)})))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.Apply(go__go_4_2_9, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v_5.UnsafePtr).V5)})))}))
goto end_branch_9
} else {

}
}
{
__t9 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_9:
__t10 = __t9
goto end_branch_10
} else {

}
}
{
__t10 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_10:
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(__t10)}
})
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.Apply(go__go_4_2_9, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](m_2))})))}
})
})})}
}

func Call_Data_Map_applySemigroupMap(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
// TAST (Let): compare_1_0 -> gopurs_runtime.Value
compare_1_0 := gopurs_runtime.RecordGet(dictOrd_0, "compare")
_ = compare_1_0
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(&Constructor_Control_Apply_Apply{1, gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](Get_Data_Map_Internal_functorMap()))}
}), gopurs_runtime.Func(func(m1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m2_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeIntersectionWith(), compare_1_0, gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return x_4
}), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](m1_2))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](m2_3))})))}
})
})})}
}

func Call_Data_Map_altSemigroupMap(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
// TAST (Let): compare_1_0 -> gopurs_runtime.Value
compare_1_0 := gopurs_runtime.RecordGet(dictOrd_0, "compare")
_ = compare_1_0
return gopurs_runtime.Value{Type: 9, IntVal: 4060500237, UnsafePtr: unsafe.Pointer(&Constructor_Control_Alt_Alt{1, gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](Get_Data_Map_Internal_functorMap()))}
}), gopurs_runtime.Func(func(m1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m2_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeUnionWith(), compare_1_0, Get_Data_Function_go__const(), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](m1_2))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](m2_3))})))}
})
})})}
}


