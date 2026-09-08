package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Data_Map_Internal_identity gopurs_runtime.Value
var once_Data_Map_Internal_identity sync.Once
func Get_Data_Map_Internal_identity() gopurs_runtime.Value {
	once_Data_Map_Internal_identity.Do(func() {
		cache_Data_Map_Internal_identity = gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Category_Category[gopurs_runtime.Value]](Get_Control_Category_categoryFn()).V1)
	})
	return cache_Data_Map_Internal_identity
}

var cache_Data_Map_Internal_identity1 gopurs_runtime.Value
var once_Data_Map_Internal_identity1 sync.Once
func Get_Data_Map_Internal_identity1() gopurs_runtime.Value {
	once_Data_Map_Internal_identity1.Do(func() {
		cache_Data_Map_Internal_identity1 = gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Category_Category[gopurs_runtime.Value]](Get_Control_Category_categoryFn()).V1)
	})
	return cache_Data_Map_Internal_identity1
}

var cache_Data_Map_Internal_identity2 gopurs_runtime.Value
var once_Data_Map_Internal_identity2 sync.Once
func Get_Data_Map_Internal_identity2() gopurs_runtime.Value {
	once_Data_Map_Internal_identity2.Do(func() {
		cache_Data_Map_Internal_identity2 = gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Category_Category[gopurs_runtime.Value]](Get_Control_Category_categoryFn()).V1)
	})
	return cache_Data_Map_Internal_identity2
}

var cache_Data_Map_Internal_Leaf gopurs_runtime.Value
var once_Data_Map_Internal_Leaf sync.Once
func Get_Data_Map_Internal_Leaf() gopurs_runtime.Value {
	once_Data_Map_Internal_Leaf.Do(func() {
		cache_Data_Map_Internal_Leaf = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(nil))}
	})
	return cache_Data_Map_Internal_Leaf
}

var cache_Data_Map_Internal_Node gopurs_runtime.Value
var once_Data_Map_Internal_Node sync.Once
func Get_Data_Map_Internal_Node() gopurs_runtime.Value {
	once_Data_Map_Internal_Node.Do(func() {
		cache_Data_Map_Internal_Node = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(value1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(value2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(value3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(value4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(value5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((&Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]{1, value0.IntVal, value1.IntVal, value2, value3, gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](value4), gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](value5)}))}
})
})
})
})
})
})
	})
	return cache_Data_Map_Internal_Node
}

var cache_Data_Map_Internal_Node__4238775810 gopurs_runtime.Value
var once_Data_Map_Internal_Node__4238775810 sync.Once
func Get_Data_Map_Internal_Node__4238775810() gopurs_runtime.Value {
	once_Data_Map_Internal_Node__4238775810.Do(func() {
		cache_Data_Map_Internal_Node__4238775810 = gopurs_runtime.Func6(func(__eta_norm_1_unused_0_box gopurs_runtime.Value, __eta_norm_0_unused_1_box gopurs_runtime.Value, __eta_norm_5_2_box gopurs_runtime.Value, __eta_norm_4_3_box gopurs_runtime.Value, __eta_norm_3_4_box gopurs_runtime.Value, __eta_norm_2_5_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(Rebox_Data_Map_Internal_261879545_2487766124(Call_Data_Map_Internal_Node__4238775810(__eta_norm_1_unused_0_box.IntVal, __eta_norm_0_unused_1_box.IntVal, uint32(__eta_norm_5_2_box.IntVal), __eta_norm_4_3_box.FloatVal(), gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[uint32, float64]](__eta_norm_3_4_box), gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[uint32, float64]](__eta_norm_2_5_box))))}
})
	})
	return cache_Data_Map_Internal_Node__4238775810
}

var cache_Data_Map_Internal_IterLeaf gopurs_runtime.Value
var once_Data_Map_Internal_IterLeaf sync.Once
func Get_Data_Map_Internal_IterLeaf() gopurs_runtime.Value {
	once_Data_Map_Internal_IterLeaf.Do(func() {
		cache_Data_Map_Internal_IterLeaf = gopurs_runtime.Value{Type: 9, IntVal: 2509360378, UnsafePtr: unsafe.Pointer(nil)}
	})
	return cache_Data_Map_Internal_IterLeaf
}

var cache_Data_Map_Internal_IterEmit gopurs_runtime.Value
var once_Data_Map_Internal_IterEmit sync.Once
func Get_Data_Map_Internal_IterEmit() gopurs_runtime.Value {
	once_Data_Map_Internal_IterEmit.Do(func() {
		cache_Data_Map_Internal_IterEmit = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(value1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(value2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1343415489, UnsafePtr: unsafe.Pointer((&Constructor_Data_Map_Internal_IterEmit[gopurs_runtime.Value, gopurs_runtime.Value]{1, value0, value1, value2}))}
})
})
})
	})
	return cache_Data_Map_Internal_IterEmit
}

var cache_Data_Map_Internal_IterNode gopurs_runtime.Value
var once_Data_Map_Internal_IterNode sync.Once
func Get_Data_Map_Internal_IterNode() gopurs_runtime.Value {
	once_Data_Map_Internal_IterNode.Do(func() {
		cache_Data_Map_Internal_IterNode = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(value1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2861335956, UnsafePtr: unsafe.Pointer((&Constructor_Data_Map_Internal_IterNode[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](value0), value1}))}
})
})
	})
	return cache_Data_Map_Internal_IterNode
}

var cache_Data_Map_Internal_IterDone gopurs_runtime.Value
var once_Data_Map_Internal_IterDone sync.Once
func Get_Data_Map_Internal_IterDone() gopurs_runtime.Value {
	once_Data_Map_Internal_IterDone.Do(func() {
		cache_Data_Map_Internal_IterDone = gopurs_runtime.Value{Type: 9, IntVal: 953589075, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_IterNext[gopurs_runtime.Value, gopurs_runtime.Value])(nil))}
	})
	return cache_Data_Map_Internal_IterDone
}

var cache_Data_Map_Internal_IterNext gopurs_runtime.Value
var once_Data_Map_Internal_IterNext sync.Once
func Get_Data_Map_Internal_IterNext() gopurs_runtime.Value {
	once_Data_Map_Internal_IterNext.Do(func() {
		cache_Data_Map_Internal_IterNext = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(value1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(value2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 953589075, UnsafePtr: unsafe.Pointer((&Constructor_Data_Map_Internal_IterNext[gopurs_runtime.Value, gopurs_runtime.Value]{1, value0, value1, value2}))}
})
})
})
	})
	return cache_Data_Map_Internal_IterNext
}

var cache_Data_Map_Internal_Split gopurs_runtime.Value
var once_Data_Map_Internal_Split sync.Once
func Get_Data_Map_Internal_Split() gopurs_runtime.Value {
	once_Data_Map_Internal_Split.Do(func() {
		cache_Data_Map_Internal_Split = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(value1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(value2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3373277644, UnsafePtr: unsafe.Pointer((&Constructor_Data_Map_Internal_Split[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](value0), gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](value1), gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](value2)}))}
})
})
})
	})
	return cache_Data_Map_Internal_Split
}

var cache_Data_Map_Internal_SplitLast gopurs_runtime.Value
var once_Data_Map_Internal_SplitLast sync.Once
func Get_Data_Map_Internal_SplitLast() gopurs_runtime.Value {
	once_Data_Map_Internal_SplitLast.Do(func() {
		cache_Data_Map_Internal_SplitLast = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(value1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(value2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2668112006, UnsafePtr: unsafe.Pointer((&Constructor_Data_Map_Internal_SplitLast[gopurs_runtime.Value, gopurs_runtime.Value]{1, value0, value1, gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](value2)}))}
})
})
})
	})
	return cache_Data_Map_Internal_SplitLast
}

var cache_Data_Map_Internal_unsafeNode gopurs_runtime.Value
var once_Data_Map_Internal_unsafeNode sync.Once
func Get_Data_Map_Internal_unsafeNode() gopurs_runtime.Value {
	once_Data_Map_Internal_unsafeNode.Do(func() {
		cache_Data_Map_Internal_unsafeNode = gopurs_runtime.Func4(func(k_0_box gopurs_runtime.Value, __local_var_1_box gopurs_runtime.Value, __local_var_2_box gopurs_runtime.Value, __local_var_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Map_Internal_unsafeNode(k_0_box, __local_var_1_box, __local_var_2_box, __local_var_3_box)
})
	})
	return cache_Data_Map_Internal_unsafeNode
}

var cache_Data_Map_Internal_toMapIter gopurs_runtime.Value
var once_Data_Map_Internal_toMapIter sync.Once
func Get_Data_Map_Internal_toMapIter() gopurs_runtime.Value {
	once_Data_Map_Internal_toMapIter.Do(func() {
		cache_Data_Map_Internal_toMapIter = gopurs_runtime.Func(func(a_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Map_Internal_toMapIter(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](a_0_box))
})
	})
	return cache_Data_Map_Internal_toMapIter
}

var cache_Data_Map_Internal_stepWith gopurs_runtime.Value
var once_Data_Map_Internal_stepWith sync.Once
func Get_Data_Map_Internal_stepWith() gopurs_runtime.Value {
	once_Data_Map_Internal_stepWith.Do(func() {
		cache_Data_Map_Internal_stepWith = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, next_1_box gopurs_runtime.Value, done_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Map_Internal_stepWith(f_0_box, next_1_box, done_2_box)
})
	})
	return cache_Data_Map_Internal_stepWith
}

var cache_Data_Map_Internal_size gopurs_runtime.Value
var once_Data_Map_Internal_size sync.Once
func Get_Data_Map_Internal_size() gopurs_runtime.Value {
	once_Data_Map_Internal_size.Do(func() {
		cache_Data_Map_Internal_size = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int(Call_Data_Map_Internal_size(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](v_0_box)))
})
	})
	return cache_Data_Map_Internal_size
}

var cache_Data_Map_Internal_size__1004260195 gopurs_runtime.Value
var once_Data_Map_Internal_size__1004260195 sync.Once
func Get_Data_Map_Internal_size__1004260195() gopurs_runtime.Value {
	once_Data_Map_Internal_size__1004260195.Do(func() {
		cache_Data_Map_Internal_size__1004260195 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int(Call_Data_Map_Internal_size__1004260195(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[uint32, float64]](v_0_box)))
})
	})
	return cache_Data_Map_Internal_size__1004260195
}

var cache_Data_Map_Internal_singleton gopurs_runtime.Value
var once_Data_Map_Internal_singleton sync.Once
func Get_Data_Map_Internal_singleton() gopurs_runtime.Value {
	once_Data_Map_Internal_singleton.Do(func() {
		cache_Data_Map_Internal_singleton = gopurs_runtime.Func2(func(k_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(Call_Data_Map_Internal_singleton(k_0_box, v_1_box))}
})
	})
	return cache_Data_Map_Internal_singleton
}

var cache_Data_Map_Internal_singleton__981367005 gopurs_runtime.Value
var once_Data_Map_Internal_singleton__981367005 sync.Once
func Get_Data_Map_Internal_singleton__981367005() gopurs_runtime.Value {
	once_Data_Map_Internal_singleton__981367005.Do(func() {
		cache_Data_Map_Internal_singleton__981367005 = gopurs_runtime.Func2(func(k_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(Rebox_Data_Map_Internal_261879545_2487766124(Call_Data_Map_Internal_singleton__981367005(uint32(k_0_box.IntVal), v_1_box.FloatVal())))}
})
	})
	return cache_Data_Map_Internal_singleton__981367005
}

var cache_Data_Map_Internal_unsafeBalancedNode gopurs_runtime.Value
var once_Data_Map_Internal_unsafeBalancedNode sync.Once
func Get_Data_Map_Internal_unsafeBalancedNode() gopurs_runtime.Value {
	once_Data_Map_Internal_unsafeBalancedNode.Do(func() {
		cache_Data_Map_Internal_unsafeBalancedNode = gopurs_runtime.Func4(func(k_0_box gopurs_runtime.Value, __local_var_1_box gopurs_runtime.Value, __local_var_2_box gopurs_runtime.Value, __local_var_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Map_Internal_unsafeBalancedNode(k_0_box, __local_var_1_box, __local_var_2_box, __local_var_3_box)
})
	})
	return cache_Data_Map_Internal_unsafeBalancedNode
}

var cache_Data_Map_Internal_unsafeSplit gopurs_runtime.Value
var once_Data_Map_Internal_unsafeSplit sync.Once
func Get_Data_Map_Internal_unsafeSplit() gopurs_runtime.Value {
	once_Data_Map_Internal_unsafeSplit.Do(func() {
		cache_Data_Map_Internal_unsafeSplit = gopurs_runtime.Func3(func(comp_0_box gopurs_runtime.Value, __local_var_1_box gopurs_runtime.Value, __local_var_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Map_Internal_unsafeSplit(comp_0_box, __local_var_1_box, __local_var_2_box)
})
	})
	return cache_Data_Map_Internal_unsafeSplit
}

var cache_Data_Map_Internal_unsafeSplitLast gopurs_runtime.Value
var once_Data_Map_Internal_unsafeSplitLast sync.Once
func Get_Data_Map_Internal_unsafeSplitLast() gopurs_runtime.Value {
	once_Data_Map_Internal_unsafeSplitLast.Do(func() {
		cache_Data_Map_Internal_unsafeSplitLast = gopurs_runtime.Func4(func(k_0_box gopurs_runtime.Value, __local_var_1_box gopurs_runtime.Value, __local_var_2_box gopurs_runtime.Value, __local_var_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Map_Internal_unsafeSplitLast(k_0_box, __local_var_1_box, __local_var_2_box, __local_var_3_box)
})
	})
	return cache_Data_Map_Internal_unsafeSplitLast
}

var cache_Data_Map_Internal_unsafeJoinNodes gopurs_runtime.Value
var once_Data_Map_Internal_unsafeJoinNodes sync.Once
func Get_Data_Map_Internal_unsafeJoinNodes() gopurs_runtime.Value {
	once_Data_Map_Internal_unsafeJoinNodes.Do(func() {
		cache_Data_Map_Internal_unsafeJoinNodes = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, __local_var_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Map_Internal_unsafeJoinNodes(v_0_box, __local_var_1_box)
})
	})
	return cache_Data_Map_Internal_unsafeJoinNodes
}

var cache_Data_Map_Internal_unsafeDifference gopurs_runtime.Value
var once_Data_Map_Internal_unsafeDifference sync.Once
func Get_Data_Map_Internal_unsafeDifference() gopurs_runtime.Value {
	once_Data_Map_Internal_unsafeDifference.Do(func() {
		cache_Data_Map_Internal_unsafeDifference = gopurs_runtime.Func3(func(comp_0_box gopurs_runtime.Value, __local_var_1_box gopurs_runtime.Value, __local_var_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Map_Internal_unsafeDifference(comp_0_box, __local_var_1_box, __local_var_2_box)
})
	})
	return cache_Data_Map_Internal_unsafeDifference
}

var cache_Data_Map_Internal_unsafeIntersectionWith gopurs_runtime.Value
var once_Data_Map_Internal_unsafeIntersectionWith sync.Once
func Get_Data_Map_Internal_unsafeIntersectionWith() gopurs_runtime.Value {
	once_Data_Map_Internal_unsafeIntersectionWith.Do(func() {
		cache_Data_Map_Internal_unsafeIntersectionWith = gopurs_runtime.Func4(func(comp_0_box gopurs_runtime.Value, __local_var_1_box gopurs_runtime.Value, __local_var_2_box gopurs_runtime.Value, __local_var_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Map_Internal_unsafeIntersectionWith(comp_0_box, __local_var_1_box, __local_var_2_box, __local_var_3_box)
})
	})
	return cache_Data_Map_Internal_unsafeIntersectionWith
}

var cache_Data_Map_Internal_unsafeUnionWith gopurs_runtime.Value
var once_Data_Map_Internal_unsafeUnionWith sync.Once
func Get_Data_Map_Internal_unsafeUnionWith() gopurs_runtime.Value {
	once_Data_Map_Internal_unsafeUnionWith.Do(func() {
		cache_Data_Map_Internal_unsafeUnionWith = gopurs_runtime.Func4(func(comp_0_box gopurs_runtime.Value, __local_var_1_box gopurs_runtime.Value, __local_var_2_box gopurs_runtime.Value, __local_var_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Map_Internal_unsafeUnionWith(comp_0_box, __local_var_1_box, __local_var_2_box, __local_var_3_box)
})
	})
	return cache_Data_Map_Internal_unsafeUnionWith
}

var cache_Data_Map_Internal_unionWith gopurs_runtime.Value
var once_Data_Map_Internal_unionWith sync.Once
func Get_Data_Map_Internal_unionWith() gopurs_runtime.Value {
	once_Data_Map_Internal_unionWith.Do(func() {
		cache_Data_Map_Internal_unionWith = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Map_Internal_unionWith(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](dictOrd_0_box))
})
	})
	return cache_Data_Map_Internal_unionWith
}

var cache_Data_Map_Internal_unionWith__2195946919 gopurs_runtime.Value
var once_Data_Map_Internal_unionWith__2195946919 sync.Once
func Get_Data_Map_Internal_unionWith__2195946919() gopurs_runtime.Value {
	once_Data_Map_Internal_unionWith__2195946919.Do(func() {
		cache_Data_Map_Internal_unionWith__2195946919 = gopurs_runtime.Func3(func(__eta_norm_2_unused_0_box gopurs_runtime.Value, __eta_norm_1_1_box gopurs_runtime.Value, __eta_norm_0_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Map_Internal_unionWith__2195946919(__eta_norm_2_unused_0_box, __eta_norm_1_1_box, __eta_norm_0_2_box)
})
	})
	return cache_Data_Map_Internal_unionWith__2195946919
}

var cache_Data_Map_Internal_union gopurs_runtime.Value
var once_Data_Map_Internal_union sync.Once
func Get_Data_Map_Internal_union() gopurs_runtime.Value {
	once_Data_Map_Internal_union.Do(func() {
		cache_Data_Map_Internal_union = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Map_Internal_union(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](dictOrd_0_box))
})
	})
	return cache_Data_Map_Internal_union
}

var cache_Data_Map_Internal_update gopurs_runtime.Value
var once_Data_Map_Internal_update sync.Once
func Get_Data_Map_Internal_update() gopurs_runtime.Value {
	once_Data_Map_Internal_update.Do(func() {
		cache_Data_Map_Internal_update = gopurs_runtime.Func3(func(dictOrd_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, k_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Map_Internal_update(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](dictOrd_0_box), f_1_box, k_2_box)
})
	})
	return cache_Data_Map_Internal_update
}

var cache_Data_Map_Internal_showTree gopurs_runtime.Value
var once_Data_Map_Internal_showTree sync.Once
func Get_Data_Map_Internal_showTree() gopurs_runtime.Value {
	once_Data_Map_Internal_showTree.Do(func() {
		cache_Data_Map_Internal_showTree = gopurs_runtime.Func2(func(dictShow_0_box gopurs_runtime.Value, dictShow1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Map_Internal_showTree(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](dictShow_0_box), gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](dictShow1_1_box))
})
	})
	return cache_Data_Map_Internal_showTree
}

var cache_Data_Map_Internal_semigroupMap gopurs_runtime.Value
var once_Data_Map_Internal_semigroupMap sync.Once
func Get_Data_Map_Internal_semigroupMap() gopurs_runtime.Value {
	once_Data_Map_Internal_semigroupMap.Do(func() {
		cache_Data_Map_Internal_semigroupMap = gopurs_runtime.Func3(func(_dollar___unused_0_box gopurs_runtime.Value, dictOrd_1_box gopurs_runtime.Value, dictSemigroup_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Map_Internal_semigroupMap(_dollar___unused_0_box, dictOrd_1_box, dictSemigroup_2_box)
})
	})
	return cache_Data_Map_Internal_semigroupMap
}

var cache_Data_Map_Internal_semigroupMap1 gopurs_runtime.Value
var once_Data_Map_Internal_semigroupMap1 sync.Once
func Get_Data_Map_Internal_semigroupMap1() gopurs_runtime.Value {
	once_Data_Map_Internal_semigroupMap1.Do(func() {
		cache_Data_Map_Internal_semigroupMap1 = gopurs_runtime.Func2(func(dictOrd_0_box gopurs_runtime.Value, dictSemigroup_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Map_Internal_semigroupMap1(dictOrd_0_box, dictSemigroup_1_box)
})
	})
	return cache_Data_Map_Internal_semigroupMap1
}

var cache_Data_Map_Internal_pop gopurs_runtime.Value
var once_Data_Map_Internal_pop sync.Once
func Get_Data_Map_Internal_pop() gopurs_runtime.Value {
	once_Data_Map_Internal_pop.Do(func() {
		cache_Data_Map_Internal_pop = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				_v := Call_Data_Map_Internal_pop(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](dictOrd_0_box))
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}()
})
	})
	return cache_Data_Map_Internal_pop
}

var cache_Data_Map_Internal_member gopurs_runtime.Value
var once_Data_Map_Internal_member sync.Once
func Get_Data_Map_Internal_member() gopurs_runtime.Value {
	once_Data_Map_Internal_member.Do(func() {
		cache_Data_Map_Internal_member = gopurs_runtime.Func2(func(dictOrd_0_box gopurs_runtime.Value, k_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Map_Internal_member(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](dictOrd_0_box), k_1_box)
})
	})
	return cache_Data_Map_Internal_member
}

var cache_Data_Map_Internal_mapMaybeWithKey gopurs_runtime.Value
var once_Data_Map_Internal_mapMaybeWithKey sync.Once
func Get_Data_Map_Internal_mapMaybeWithKey() gopurs_runtime.Value {
	once_Data_Map_Internal_mapMaybeWithKey.Do(func() {
		cache_Data_Map_Internal_mapMaybeWithKey = gopurs_runtime.Func2(func(dictOrd_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Map_Internal_mapMaybeWithKey(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](dictOrd_0_box), f_1_box)
})
	})
	return cache_Data_Map_Internal_mapMaybeWithKey
}

var cache_Data_Map_Internal_mapMaybe gopurs_runtime.Value
var once_Data_Map_Internal_mapMaybe sync.Once
func Get_Data_Map_Internal_mapMaybe() gopurs_runtime.Value {
	once_Data_Map_Internal_mapMaybe.Do(func() {
		cache_Data_Map_Internal_mapMaybe = gopurs_runtime.Func2(func(dictOrd_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Map_Internal_mapMaybe(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](dictOrd_0_box), x_1_box)
})
	})
	return cache_Data_Map_Internal_mapMaybe
}

var cache_Data_Map_Internal_lookupLE gopurs_runtime.Value
var once_Data_Map_Internal_lookupLE sync.Once
func Get_Data_Map_Internal_lookupLE() gopurs_runtime.Value {
	once_Data_Map_Internal_lookupLE.Do(func() {
		cache_Data_Map_Internal_lookupLE = gopurs_runtime.Func2(func(dictOrd_0_box gopurs_runtime.Value, k_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				_v := Call_Data_Map_Internal_lookupLE(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](dictOrd_0_box), k_1_box)
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}()
})
	})
	return cache_Data_Map_Internal_lookupLE
}

var cache_Data_Map_Internal_lookupGE gopurs_runtime.Value
var once_Data_Map_Internal_lookupGE sync.Once
func Get_Data_Map_Internal_lookupGE() gopurs_runtime.Value {
	once_Data_Map_Internal_lookupGE.Do(func() {
		cache_Data_Map_Internal_lookupGE = gopurs_runtime.Func2(func(dictOrd_0_box gopurs_runtime.Value, k_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				_v := Call_Data_Map_Internal_lookupGE(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](dictOrd_0_box), k_1_box)
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}()
})
	})
	return cache_Data_Map_Internal_lookupGE
}

var cache_Data_Map_Internal_lookup gopurs_runtime.Value
var once_Data_Map_Internal_lookup sync.Once
func Get_Data_Map_Internal_lookup() gopurs_runtime.Value {
	once_Data_Map_Internal_lookup.Do(func() {
		cache_Data_Map_Internal_lookup = gopurs_runtime.Func2(func(dictOrd_0_box gopurs_runtime.Value, k_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				_v := Call_Data_Map_Internal_lookup(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](dictOrd_0_box), k_1_box)
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}()
})
	})
	return cache_Data_Map_Internal_lookup
}

var cache_Data_Map_Internal_lookup__634903124 gopurs_runtime.Value
var once_Data_Map_Internal_lookup__634903124 sync.Once
func Get_Data_Map_Internal_lookup__634903124() gopurs_runtime.Value {
	once_Data_Map_Internal_lookup__634903124.Do(func() {
		cache_Data_Map_Internal_lookup__634903124 = gopurs_runtime.Func2(func(k_unused_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				_v := Call_Data_Map_Internal_lookup__634903124(uint32(k_unused_0_box.IntVal), gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[uint32, float64]](__eta_norm_0_1_box))
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}()
})
	})
	return cache_Data_Map_Internal_lookup__634903124
}

var cache_Data_Map_Internal_iterMapU gopurs_runtime.Value
var once_Data_Map_Internal_iterMapU sync.Once
func Get_Data_Map_Internal_iterMapU() gopurs_runtime.Value {
	once_Data_Map_Internal_iterMapU.Do(func() {
		cache_Data_Map_Internal_iterMapU = gopurs_runtime.Func2(func(iter_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Map_Internal_iterMapU(iter_0_box, gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](v_1_box))
})
	})
	return cache_Data_Map_Internal_iterMapU
}

var cache_Data_Map_Internal_stepUnorderedCps gopurs_runtime.Value
var once_Data_Map_Internal_stepUnorderedCps sync.Once
func Get_Data_Map_Internal_stepUnorderedCps() gopurs_runtime.Value {
	once_Data_Map_Internal_stepUnorderedCps.Do(func() {
		cache_Data_Map_Internal_stepUnorderedCps = gopurs_runtime.Apply(Get_Data_Map_Internal_stepWith(), Get_Data_Map_Internal_iterMapU())
	})
	return cache_Data_Map_Internal_stepUnorderedCps
}

var cache_Data_Map_Internal_stepUnfoldrUnordered gopurs_runtime.Value
var once_Data_Map_Internal_stepUnfoldrUnordered sync.Once
func Get_Data_Map_Internal_stepUnfoldrUnordered() gopurs_runtime.Value {
	once_Data_Map_Internal_stepUnfoldrUnordered.Do(func() {
		cache_Data_Map_Internal_stepUnfoldrUnordered = Call_Data_Map_Internal_stepWith(Get_Data_Map_Internal_iterMapU(), gopurs_runtime.Func3(func(k_0 gopurs_runtime.Value, __local_var_1 gopurs_runtime.Value, __local_var_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_Map_Internal_2549197956_3094389156(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[*Constructor_Data_Tuple_Tuple[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value], gopurs_runtime.Value]]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer((&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer((&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, k_0, __local_var_1}))}, __local_var_2}))}, true}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())))}
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_Map_Internal_2549197956_3094389156(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[*Constructor_Data_Tuple_Tuple[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value], gopurs_runtime.Value]]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())))}
}))
	})
	return cache_Data_Map_Internal_stepUnfoldrUnordered
}

var cache_Data_Map_Internal_toUnfoldableUnordered gopurs_runtime.Value
var once_Data_Map_Internal_toUnfoldableUnordered sync.Once
func Get_Data_Map_Internal_toUnfoldableUnordered() gopurs_runtime.Value {
	once_Data_Map_Internal_toUnfoldableUnordered.Do(func() {
		cache_Data_Map_Internal_toUnfoldableUnordered = gopurs_runtime.Func(func(dictUnfoldable_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Map_Internal_toUnfoldableUnordered(gopurs_runtime.CoerceToStruct[Constructor_Data_Unfoldable_Unfoldable[gopurs_runtime.Value]](dictUnfoldable_0_box))
})
	})
	return cache_Data_Map_Internal_toUnfoldableUnordered
}

var cache_Data_Map_Internal_stepUnordered gopurs_runtime.Value
var once_Data_Map_Internal_stepUnordered sync.Once
func Get_Data_Map_Internal_stepUnordered() gopurs_runtime.Value {
	once_Data_Map_Internal_stepUnordered.Do(func() {
		cache_Data_Map_Internal_stepUnordered = Call_Data_Map_Internal_stepWith(Get_Data_Map_Internal_iterMapU(), gopurs_runtime.Func3(func(k_0 gopurs_runtime.Value, __local_var_1 gopurs_runtime.Value, __local_var_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 953589075, UnsafePtr: unsafe.Pointer((&Constructor_Data_Map_Internal_IterNext[gopurs_runtime.Value, gopurs_runtime.Value]{1, k_0, __local_var_1, __local_var_2}))}
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 953589075, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_IterNext[gopurs_runtime.Value, gopurs_runtime.Value])(nil))}
}))
	})
	return cache_Data_Map_Internal_stepUnordered
}

var cache_Data_Map_Internal_iterMapR gopurs_runtime.Value
var once_Data_Map_Internal_iterMapR sync.Once
func Get_Data_Map_Internal_iterMapR() gopurs_runtime.Value {
	once_Data_Map_Internal_iterMapR.Do(func() {
		cache_Data_Map_Internal_iterMapR = func() gopurs_runtime.Value {
var Call_local_Data_Map_Internal_go__go_0_0_10 func(gopurs_runtime.Value, *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value
_ = Call_local_Data_Map_Internal_go__go_0_0_10
var go__go_0_0_10 gopurs_runtime.Value
_ = go__go_0_0_10
Call_local_Data_Map_Internal_go__go_0_0_10 = func(iter_1_loop gopurs_runtime.Value, v_2_loop *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
go__go_0_0_10:
for {
if false { continue go__go_0_0_10 }
var iter_1 gopurs_runtime.Value = iter_1_loop
_ = iter_1
var v_2 *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value] = v_2_loop
_ = v_2
var __t3 gopurs_runtime.Value
{
if (v_2 == nil) {
__t3 = iter_1
goto end_branch_3
} else {

}
}
{
if (v_2 != nil) {
var __t2 gopurs_runtime.Value
{
var __t_tag_1 *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value] = (v_2).V5
if (__t_tag_1 == nil) {
iter_1_loop = gopurs_runtime.Value{Type: 9, IntVal: 1343415489, UnsafePtr: unsafe.Pointer((&Constructor_Data_Map_Internal_IterEmit[gopurs_runtime.Value, gopurs_runtime.Value]{1, (v_2).V2, (v_2).V3, iter_1}))}
v_2_loop = (v_2).V4
continue go__go_0_0_10
__t2 = func() gopurs_runtime.Value { panic("unreachable") }()
goto end_branch_2
} else {

}
}
{
iter_1_loop = gopurs_runtime.Value{Type: 9, IntVal: 1343415489, UnsafePtr: unsafe.Pointer((&Constructor_Data_Map_Internal_IterEmit[gopurs_runtime.Value, gopurs_runtime.Value]{1, (v_2).V2, (v_2).V3, gopurs_runtime.Value{Type: 9, IntVal: 2861335956, UnsafePtr: unsafe.Pointer((&Constructor_Data_Map_Internal_IterNode[gopurs_runtime.Value, gopurs_runtime.Value]{1, (v_2).V4, iter_1}))}}))}
v_2_loop = (v_2).V5
continue go__go_0_0_10
__t2 = func() gopurs_runtime.Value { panic("unreachable") }()
}
end_branch_2:
__t3 = __t2
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
}
go__go_0_0_10 = gopurs_runtime.Func(func(iter_1_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return Call_local_Data_Map_Internal_go__go_0_0_10(iter_1_loop_val, gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](v_2_loop_val))
})
})
return go__go_0_0_10
}()
	})
	return cache_Data_Map_Internal_iterMapR
}

var cache_Data_Map_Internal_stepDescCps gopurs_runtime.Value
var once_Data_Map_Internal_stepDescCps sync.Once
func Get_Data_Map_Internal_stepDescCps() gopurs_runtime.Value {
	once_Data_Map_Internal_stepDescCps.Do(func() {
		cache_Data_Map_Internal_stepDescCps = gopurs_runtime.Apply(Get_Data_Map_Internal_stepWith(), Get_Data_Map_Internal_iterMapR())
	})
	return cache_Data_Map_Internal_stepDescCps
}

var cache_Data_Map_Internal_stepDesc gopurs_runtime.Value
var once_Data_Map_Internal_stepDesc sync.Once
func Get_Data_Map_Internal_stepDesc() gopurs_runtime.Value {
	once_Data_Map_Internal_stepDesc.Do(func() {
		cache_Data_Map_Internal_stepDesc = Call_Data_Map_Internal_stepWith(Get_Data_Map_Internal_iterMapR(), gopurs_runtime.Func3(func(k_0 gopurs_runtime.Value, __local_var_1 gopurs_runtime.Value, __local_var_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 953589075, UnsafePtr: unsafe.Pointer((&Constructor_Data_Map_Internal_IterNext[gopurs_runtime.Value, gopurs_runtime.Value]{1, k_0, __local_var_1, __local_var_2}))}
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 953589075, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_IterNext[gopurs_runtime.Value, gopurs_runtime.Value])(nil))}
}))
	})
	return cache_Data_Map_Internal_stepDesc
}

var cache_Data_Map_Internal_iterMapL gopurs_runtime.Value
var once_Data_Map_Internal_iterMapL sync.Once
func Get_Data_Map_Internal_iterMapL() gopurs_runtime.Value {
	once_Data_Map_Internal_iterMapL.Do(func() {
		cache_Data_Map_Internal_iterMapL = func() gopurs_runtime.Value {
var Call_local_Data_Map_Internal_go__go_0_0_11 func(gopurs_runtime.Value, *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value
_ = Call_local_Data_Map_Internal_go__go_0_0_11
var go__go_0_0_11 gopurs_runtime.Value
_ = go__go_0_0_11
Call_local_Data_Map_Internal_go__go_0_0_11 = func(iter_1_loop gopurs_runtime.Value, v_2_loop *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
go__go_0_0_11:
for {
if false { continue go__go_0_0_11 }
var iter_1 gopurs_runtime.Value = iter_1_loop
_ = iter_1
var v_2 *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value] = v_2_loop
_ = v_2
var __t3 gopurs_runtime.Value
{
if (v_2 == nil) {
__t3 = iter_1
goto end_branch_3
} else {

}
}
{
if (v_2 != nil) {
var __t2 gopurs_runtime.Value
{
var __t_tag_1 *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value] = (v_2).V5
if (__t_tag_1 == nil) {
iter_1_loop = gopurs_runtime.Value{Type: 9, IntVal: 1343415489, UnsafePtr: unsafe.Pointer((&Constructor_Data_Map_Internal_IterEmit[gopurs_runtime.Value, gopurs_runtime.Value]{1, (v_2).V2, (v_2).V3, iter_1}))}
v_2_loop = (v_2).V4
continue go__go_0_0_11
__t2 = func() gopurs_runtime.Value { panic("unreachable") }()
goto end_branch_2
} else {

}
}
{
iter_1_loop = gopurs_runtime.Value{Type: 9, IntVal: 1343415489, UnsafePtr: unsafe.Pointer((&Constructor_Data_Map_Internal_IterEmit[gopurs_runtime.Value, gopurs_runtime.Value]{1, (v_2).V2, (v_2).V3, gopurs_runtime.Value{Type: 9, IntVal: 2861335956, UnsafePtr: unsafe.Pointer((&Constructor_Data_Map_Internal_IterNode[gopurs_runtime.Value, gopurs_runtime.Value]{1, (v_2).V5, iter_1}))}}))}
v_2_loop = (v_2).V4
continue go__go_0_0_11
__t2 = func() gopurs_runtime.Value { panic("unreachable") }()
}
end_branch_2:
__t3 = __t2
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
}
go__go_0_0_11 = gopurs_runtime.Func(func(iter_1_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return Call_local_Data_Map_Internal_go__go_0_0_11(iter_1_loop_val, gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](v_2_loop_val))
})
})
return go__go_0_0_11
}()
	})
	return cache_Data_Map_Internal_iterMapL
}

var cache_Data_Map_Internal_stepAscCps gopurs_runtime.Value
var once_Data_Map_Internal_stepAscCps sync.Once
func Get_Data_Map_Internal_stepAscCps() gopurs_runtime.Value {
	once_Data_Map_Internal_stepAscCps.Do(func() {
		cache_Data_Map_Internal_stepAscCps = gopurs_runtime.Apply(Get_Data_Map_Internal_stepWith(), Get_Data_Map_Internal_iterMapL())
	})
	return cache_Data_Map_Internal_stepAscCps
}

var cache_Data_Map_Internal_stepAsc gopurs_runtime.Value
var once_Data_Map_Internal_stepAsc sync.Once
func Get_Data_Map_Internal_stepAsc() gopurs_runtime.Value {
	once_Data_Map_Internal_stepAsc.Do(func() {
		cache_Data_Map_Internal_stepAsc = Call_Data_Map_Internal_stepWith(Get_Data_Map_Internal_iterMapL(), gopurs_runtime.Func3(func(k_0 gopurs_runtime.Value, __local_var_1 gopurs_runtime.Value, __local_var_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 953589075, UnsafePtr: unsafe.Pointer((&Constructor_Data_Map_Internal_IterNext[gopurs_runtime.Value, gopurs_runtime.Value]{1, k_0, __local_var_1, __local_var_2}))}
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 953589075, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_IterNext[gopurs_runtime.Value, gopurs_runtime.Value])(nil))}
}))
	})
	return cache_Data_Map_Internal_stepAsc
}

var cache_Data_Map_Internal_eqMapIter gopurs_runtime.Value
var once_Data_Map_Internal_eqMapIter sync.Once
func Get_Data_Map_Internal_eqMapIter() gopurs_runtime.Value {
	once_Data_Map_Internal_eqMapIter.Do(func() {
		cache_Data_Map_Internal_eqMapIter = gopurs_runtime.Func2(func(dictEq_0_box gopurs_runtime.Value, dictEq1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Map_Internal_eqMapIter(dictEq_0_box, dictEq1_1_box)
})
	})
	return cache_Data_Map_Internal_eqMapIter
}

var cache_Data_Map_Internal_ordMapIter gopurs_runtime.Value
var once_Data_Map_Internal_ordMapIter sync.Once
func Get_Data_Map_Internal_ordMapIter() gopurs_runtime.Value {
	once_Data_Map_Internal_ordMapIter.Do(func() {
		cache_Data_Map_Internal_ordMapIter = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Map_Internal_ordMapIter(dictOrd_0_box)
})
	})
	return cache_Data_Map_Internal_ordMapIter
}

var cache_Data_Map_Internal_stepUnfoldr gopurs_runtime.Value
var once_Data_Map_Internal_stepUnfoldr sync.Once
func Get_Data_Map_Internal_stepUnfoldr() gopurs_runtime.Value {
	once_Data_Map_Internal_stepUnfoldr.Do(func() {
		cache_Data_Map_Internal_stepUnfoldr = Call_Data_Map_Internal_stepWith(Get_Data_Map_Internal_iterMapL(), gopurs_runtime.Func3(func(k_0 gopurs_runtime.Value, __local_var_1 gopurs_runtime.Value, __local_var_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_Map_Internal_2549197956_3094389156(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[*Constructor_Data_Tuple_Tuple[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value], gopurs_runtime.Value]]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer((&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer((&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, k_0, __local_var_1}))}, __local_var_2}))}, true}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())))}
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_Map_Internal_2549197956_3094389156(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[*Constructor_Data_Tuple_Tuple[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value], gopurs_runtime.Value]]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())))}
}))
	})
	return cache_Data_Map_Internal_stepUnfoldr
}

var cache_Data_Map_Internal_toUnfoldable gopurs_runtime.Value
var once_Data_Map_Internal_toUnfoldable sync.Once
func Get_Data_Map_Internal_toUnfoldable() gopurs_runtime.Value {
	once_Data_Map_Internal_toUnfoldable.Do(func() {
		cache_Data_Map_Internal_toUnfoldable = gopurs_runtime.Func(func(dictUnfoldable_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Map_Internal_toUnfoldable(gopurs_runtime.CoerceToStruct[Constructor_Data_Unfoldable_Unfoldable[gopurs_runtime.Value]](dictUnfoldable_0_box))
})
	})
	return cache_Data_Map_Internal_toUnfoldable
}

var cache_Data_Map_Internal_toUnfoldable__994050345 gopurs_runtime.Value
var once_Data_Map_Internal_toUnfoldable__994050345 sync.Once
func Get_Data_Map_Internal_toUnfoldable__994050345() gopurs_runtime.Value {
	once_Data_Map_Internal_toUnfoldable__994050345.Do(func() {
		cache_Data_Map_Internal_toUnfoldable__994050345 = gopurs_runtime.Func(func(__eta_norm_0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Rebox_Data_Map_Internal_2442833393_849153993(Call_Data_Map_Internal_toUnfoldable__994050345(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[uint32, float64]](__eta_norm_0_0_box))))}
})
	})
	return cache_Data_Map_Internal_toUnfoldable__994050345
}

var cache_Data_Map_Internal_showMap gopurs_runtime.Value
var once_Data_Map_Internal_showMap sync.Once
func Get_Data_Map_Internal_showMap() gopurs_runtime.Value {
	once_Data_Map_Internal_showMap.Do(func() {
		cache_Data_Map_Internal_showMap = gopurs_runtime.Func2(func(dictShow_0_box gopurs_runtime.Value, dictShow1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Map_Internal_showMap(dictShow_0_box, dictShow1_1_box)
})
	})
	return cache_Data_Map_Internal_showMap
}

var cache_Data_Map_Internal_isSubmap gopurs_runtime.Value
var once_Data_Map_Internal_isSubmap sync.Once
func Get_Data_Map_Internal_isSubmap() gopurs_runtime.Value {
	once_Data_Map_Internal_isSubmap.Do(func() {
		cache_Data_Map_Internal_isSubmap = gopurs_runtime.Func2(func(dictOrd_0_box gopurs_runtime.Value, dictEq_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Map_Internal_isSubmap(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](dictOrd_0_box), gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](dictEq_1_box))
})
	})
	return cache_Data_Map_Internal_isSubmap
}

var cache_Data_Map_Internal_isEmpty gopurs_runtime.Value
var once_Data_Map_Internal_isEmpty sync.Once
func Get_Data_Map_Internal_isEmpty() gopurs_runtime.Value {
	once_Data_Map_Internal_isEmpty.Do(func() {
		cache_Data_Map_Internal_isEmpty = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_Data_Map_Internal_isEmpty(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](v_0_box)))
})
	})
	return cache_Data_Map_Internal_isEmpty
}

var cache_Data_Map_Internal_intersectionWith gopurs_runtime.Value
var once_Data_Map_Internal_intersectionWith sync.Once
func Get_Data_Map_Internal_intersectionWith() gopurs_runtime.Value {
	once_Data_Map_Internal_intersectionWith.Do(func() {
		cache_Data_Map_Internal_intersectionWith = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Map_Internal_intersectionWith(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](dictOrd_0_box))
})
	})
	return cache_Data_Map_Internal_intersectionWith
}

var cache_Data_Map_Internal_intersection gopurs_runtime.Value
var once_Data_Map_Internal_intersection sync.Once
func Get_Data_Map_Internal_intersection() gopurs_runtime.Value {
	once_Data_Map_Internal_intersection.Do(func() {
		cache_Data_Map_Internal_intersection = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Map_Internal_intersection(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](dictOrd_0_box))
})
	})
	return cache_Data_Map_Internal_intersection
}

var cache_Data_Map_Internal_insertWith gopurs_runtime.Value
var once_Data_Map_Internal_insertWith sync.Once
func Get_Data_Map_Internal_insertWith() gopurs_runtime.Value {
	once_Data_Map_Internal_insertWith.Do(func() {
		cache_Data_Map_Internal_insertWith = gopurs_runtime.Func4(func(dictOrd_0_box gopurs_runtime.Value, app_1_box gopurs_runtime.Value, k_2_box gopurs_runtime.Value, v_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Map_Internal_insertWith(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](dictOrd_0_box), app_1_box, k_2_box, v_3_box)
})
	})
	return cache_Data_Map_Internal_insertWith
}

var cache_Data_Map_Internal_insert gopurs_runtime.Value
var once_Data_Map_Internal_insert sync.Once
func Get_Data_Map_Internal_insert() gopurs_runtime.Value {
	once_Data_Map_Internal_insert.Do(func() {
		cache_Data_Map_Internal_insert = gopurs_runtime.Func3(func(dictOrd_0_box gopurs_runtime.Value, k_1_box gopurs_runtime.Value, v_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Map_Internal_insert(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](dictOrd_0_box), k_1_box, v_2_box)
})
	})
	return cache_Data_Map_Internal_insert
}

var cache_Data_Map_Internal_functorMap gopurs_runtime.Value
var once_Data_Map_Internal_functorMap sync.Once
func Get_Data_Map_Internal_functorMap() gopurs_runtime.Value {
	once_Data_Map_Internal_functorMap.Do(func() {
		cache_Data_Map_Internal_functorMap = gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(Rebox_Data_Map_Internal_3281783655_2812149806((&Constructor_Data_Functor_Functor[*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_1_0_21 gopurs_runtime.Value
_ = go__go_1_0_21
// FALLBACK TCO: isLoop=false len=1
go__go_1_0_21 = gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t3 *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]
{
var __t_tag_1 *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](v_2)
if (__t_tag_1 == nil) {
__t3 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(nil))})
goto end_branch_3
} else {

}
}
{
var __t_tag_2 *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](v_2)
if (__t_tag_2 != nil) {
__t3 = (&Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V0, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V1, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V2, gopurs_runtime.Apply(f_0, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V3), gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(go__go_1_0_21, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V4)})), gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(go__go_1_0_21, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V5)}))})
goto end_branch_3
} else {

}
}
{
__t3 = func() *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value] { panic("Failed pattern match") }()
}
end_branch_3:
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(__t3)}
})
return go__go_1_0_21
})})))}
	})
	return cache_Data_Map_Internal_functorMap
}

var cache_Data_Map_Internal_functorWithIndexMap gopurs_runtime.Value
var once_Data_Map_Internal_functorWithIndexMap sync.Once
func Get_Data_Map_Internal_functorWithIndexMap() gopurs_runtime.Value {
	once_Data_Map_Internal_functorWithIndexMap.Do(func() {
		cache_Data_Map_Internal_functorWithIndexMap = gopurs_runtime.Value{Type: 9, IntVal: 4077743418, UnsafePtr: unsafe.Pointer(Rebox_Data_Map_Internal_1397444001_2412140840((&Constructor_Data_FunctorWithIndex_FunctorWithIndex[gopurs_runtime.Value, *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(Rebox_Data_Map_Internal_3281783655_2812149806(Rebox_Data_Map_Internal_2812149806_3281783655(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Data_Map_Internal_functorMap()))))}
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_1_0_22 gopurs_runtime.Value
_ = go__go_1_0_22
// FALLBACK TCO: isLoop=false len=1
go__go_1_0_22 = gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t3 *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]
{
var __t_tag_1 *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](v_2)
if (__t_tag_1 == nil) {
__t3 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(nil))})
goto end_branch_3
} else {

}
}
{
var __t_tag_2 *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](v_2)
if (__t_tag_2 != nil) {
__t3 = (&Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V0, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V1, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V2, gopurs_runtime.Apply2(f_0, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V3), gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(go__go_1_0_22, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V4)})), gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(go__go_1_0_22, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V5)}))})
goto end_branch_3
} else {

}
}
{
__t3 = func() *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value] { panic("Failed pattern match") }()
}
end_branch_3:
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(__t3)}
})
return go__go_1_0_22
})})))}
	})
	return cache_Data_Map_Internal_functorWithIndexMap
}

var cache_Data_Map_Internal_foldableMap gopurs_runtime.Value
var once_Data_Map_Internal_foldableMap sync.Once
func Get_Data_Map_Internal_foldableMap() gopurs_runtime.Value {
	once_Data_Map_Internal_foldableMap.Do(func() {
		cache_Data_Map_Internal_foldableMap = gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(Rebox_Data_Map_Internal_3596835815_1680800814((&Constructor_Data_Foldable_Foldable[*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Semigroup0_1_0 shape=App(Other) bindingType=(ADT ["Data","Semigroup","Semigroup"] [(TypeVar m)])
Semigroup0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_1_0
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_3_1_23 gopurs_runtime.Value
_ = go__go_3_1_23
// FALLBACK TCO: isLoop=false len=1
go__go_3_1_23 = gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t4 gopurs_runtime.Value
{
var __t_tag_2 *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](v_4)
if (__t_tag_2 == nil) {
__t4 = gopurs_runtime.RecordGet(dictMonoid_0, "mempty")
goto end_branch_4
} else {

}
}
{
var __t_tag_3 *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](v_4)
if (__t_tag_3 != nil) {
__t4 = gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_1_0.V0), gopurs_runtime.Apply(go__go_3_1_23, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V4)}), gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_1_0.V0), gopurs_runtime.Apply(f_2, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V3), gopurs_runtime.Apply(go__go_3_1_23, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V5)})))
goto end_branch_4
} else {

}
}
{
__t4 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_4:
return __t4
})
return go__go_3_1_23
})
}), gopurs_runtime.Func2(func(f_0 gopurs_runtime.Value, z_1 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_2_5_24 gopurs_runtime.Value
_ = go__go_2_5_24
// FALLBACK TCO: isLoop=false len=1
go__go_2_5_24 = gopurs_runtime.Func2(func(z_prime__3 gopurs_runtime.Value, __local_var_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t8 gopurs_runtime.Value
{
var __t_tag_6 *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__local_var_4)
if (__t_tag_6 == nil) {
__t8 = z_prime__3
goto end_branch_8
} else {

}
}
{
var __t_tag_7 *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__local_var_4)
if (__t_tag_7 != nil) {
__t8 = gopurs_runtime.UncurriedApp2(go__go_2_5_24, gopurs_runtime.Apply2(f_0, gopurs_runtime.UncurriedApp2(go__go_2_5_24, z_prime__3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_4.UnsafePtr).V4)}), (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_4.UnsafePtr).V3), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_4.UnsafePtr).V5)})
goto end_branch_8
} else {

}
}
{
__t8 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_8:
return __t8
})
return gopurs_runtime.Func(func(m_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.UncurriedApp2(go__go_2_5_24, z_1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](m_3))})
})
}), gopurs_runtime.Func2(func(f_0 gopurs_runtime.Value, z_1 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_2_9_25 gopurs_runtime.Value
_ = go__go_2_9_25
// FALLBACK TCO: isLoop=false len=1
go__go_2_9_25 = gopurs_runtime.Func2(func(m_prime__3 gopurs_runtime.Value, __local_var_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t12 gopurs_runtime.Value
{
var __t_tag_10 *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](m_prime__3)
if (__t_tag_10 == nil) {
__t12 = __local_var_4
goto end_branch_12
} else {

}
}
{
var __t_tag_11 *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](m_prime__3)
if (__t_tag_11 != nil) {
__t12 = gopurs_runtime.UncurriedApp2(go__go_2_9_25, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(m_prime__3.UnsafePtr).V4)}, gopurs_runtime.Apply2(f_0, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(m_prime__3.UnsafePtr).V3, gopurs_runtime.UncurriedApp2(go__go_2_9_25, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(m_prime__3.UnsafePtr).V5)}, __local_var_4)))
goto end_branch_12
} else {

}
}
{
__t12 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_12:
return __t12
})
return gopurs_runtime.Func(func(m_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.UncurriedApp2(go__go_2_9_25, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](m_3))}, z_1)
})
})})))}
	})
	return cache_Data_Map_Internal_foldableMap
}

var cache_Data_Map_Internal_foldableWithIndexMap gopurs_runtime.Value
var once_Data_Map_Internal_foldableWithIndexMap sync.Once
func Get_Data_Map_Internal_foldableWithIndexMap() gopurs_runtime.Value {
	once_Data_Map_Internal_foldableWithIndexMap.Do(func() {
		cache_Data_Map_Internal_foldableWithIndexMap = gopurs_runtime.Value{Type: 9, IntVal: 74250362, UnsafePtr: unsafe.Pointer(Rebox_Data_Map_Internal_1245395425_3725484264((&Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(Rebox_Data_Map_Internal_3596835815_1680800814(Rebox_Data_Map_Internal_1680800814_3596835815(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_Map_Internal_foldableMap()))))}
}), gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Semigroup0_1_0 shape=App(Other) bindingType=(ADT ["Data","Semigroup","Semigroup"] [(TypeVar m)])
Semigroup0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_1_0
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_3_1_26 gopurs_runtime.Value
_ = go__go_3_1_26
// FALLBACK TCO: isLoop=false len=1
go__go_3_1_26 = gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t4 gopurs_runtime.Value
{
var __t_tag_2 *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](v_4)
if (__t_tag_2 == nil) {
__t4 = gopurs_runtime.RecordGet(dictMonoid_0, "mempty")
goto end_branch_4
} else {

}
}
{
var __t_tag_3 *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](v_4)
if (__t_tag_3 != nil) {
__t4 = gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_1_0.V0), gopurs_runtime.Apply(go__go_3_1_26, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V4)}), gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_1_0.V0), gopurs_runtime.Apply2(f_2, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V3), gopurs_runtime.Apply(go__go_3_1_26, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V5)})))
goto end_branch_4
} else {

}
}
{
__t4 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_4:
return __t4
})
return go__go_3_1_26
})
}), gopurs_runtime.Func2(func(f_0 gopurs_runtime.Value, z_1 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_2_5_27 gopurs_runtime.Value
_ = go__go_2_5_27
// FALLBACK TCO: isLoop=false len=1
go__go_2_5_27 = gopurs_runtime.Func2(func(z_prime__3 gopurs_runtime.Value, __local_var_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t8 gopurs_runtime.Value
{
var __t_tag_6 *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__local_var_4)
if (__t_tag_6 == nil) {
__t8 = z_prime__3
goto end_branch_8
} else {

}
}
{
var __t_tag_7 *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__local_var_4)
if (__t_tag_7 != nil) {
__t8 = gopurs_runtime.UncurriedApp2(go__go_2_5_27, gopurs_runtime.Apply3(f_0, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_4.UnsafePtr).V2, gopurs_runtime.UncurriedApp2(go__go_2_5_27, z_prime__3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_4.UnsafePtr).V4)}), (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_4.UnsafePtr).V3), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_4.UnsafePtr).V5)})
goto end_branch_8
} else {

}
}
{
__t8 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_8:
return __t8
})
return gopurs_runtime.Func(func(m_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.UncurriedApp2(go__go_2_5_27, z_1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](m_3))})
})
}), gopurs_runtime.Func2(func(f_0 gopurs_runtime.Value, z_1 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_2_9_28 gopurs_runtime.Value
_ = go__go_2_9_28
// FALLBACK TCO: isLoop=false len=1
go__go_2_9_28 = gopurs_runtime.Func2(func(m_prime__3 gopurs_runtime.Value, __local_var_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t12 gopurs_runtime.Value
{
var __t_tag_10 *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](m_prime__3)
if (__t_tag_10 == nil) {
__t12 = __local_var_4
goto end_branch_12
} else {

}
}
{
var __t_tag_11 *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](m_prime__3)
if (__t_tag_11 != nil) {
__t12 = gopurs_runtime.UncurriedApp2(go__go_2_9_28, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(m_prime__3.UnsafePtr).V4)}, gopurs_runtime.Apply3(f_0, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(m_prime__3.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(m_prime__3.UnsafePtr).V3, gopurs_runtime.UncurriedApp2(go__go_2_9_28, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(m_prime__3.UnsafePtr).V5)}, __local_var_4)))
goto end_branch_12
} else {

}
}
{
__t12 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_12:
return __t12
})
return gopurs_runtime.Func(func(m_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.UncurriedApp2(go__go_2_9_28, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](m_3))}, z_1)
})
})})))}
	})
	return cache_Data_Map_Internal_foldableWithIndexMap
}

var cache_Data_Map_Internal_keys gopurs_runtime.Value
var once_Data_Map_Internal_keys sync.Once
func Get_Data_Map_Internal_keys() gopurs_runtime.Value {
	once_Data_Map_Internal_keys.Do(func() {
		cache_Data_Map_Internal_keys = func() gopurs_runtime.Value {
var go__go_0_0_29 gopurs_runtime.Value
_ = go__go_0_0_29
// FALLBACK TCO: isLoop=false len=1
go__go_0_0_29 = gopurs_runtime.Func2(func(m_prime__1 gopurs_runtime.Value, __local_var_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t3 gopurs_runtime.Value
{
var __t_tag_1 *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](m_prime__1)
if (__t_tag_1 == nil) {
__t3 = __local_var_2
goto end_branch_3
} else {

}
}
{
var __t_tag_2 *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](m_prime__1)
if (__t_tag_2 != nil) {
__t3 = gopurs_runtime.UncurriedApp2(go__go_0_0_29, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(m_prime__1.UnsafePtr).V4)}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(m_prime__1.UnsafePtr).V2, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.UncurriedApp2(go__go_0_0_29, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(m_prime__1.UnsafePtr).V5)}, __local_var_2))}))})
goto end_branch_3
} else {

}
}
{
__t3 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_3:
return __t3
})
return gopurs_runtime.Func(func(m_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.UncurriedApp2(go__go_0_0_29, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](m_1))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil))})
})
}()
	})
	return cache_Data_Map_Internal_keys
}

var cache_Data_Map_Internal_traversableMap gopurs_runtime.Value
var once_Data_Map_Internal_traversableMap sync.Once
func Get_Data_Map_Internal_traversableMap() gopurs_runtime.Value {
	once_Data_Map_Internal_traversableMap.Do(func() {
		cache_Data_Map_Internal_traversableMap = gopurs_runtime.Value{Type: 9, IntVal: 3941073978, UnsafePtr: unsafe.Pointer(Rebox_Data_Map_Internal_2141765991_3043886126((&Constructor_Data_Traversable_Traversable[*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(Rebox_Data_Map_Internal_3596835815_1680800814(Rebox_Data_Map_Internal_1680800814_3596835815(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_Map_Internal_foldableMap()))))}
}), gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(Rebox_Data_Map_Internal_3281783655_2812149806(Rebox_Data_Map_Internal_2812149806_3281783655(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Data_Map_Internal_functorMap()))))}
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_Map_Internal_traversableMap(), "traverse"), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](dictApplicative_0))}, gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Category_Category[gopurs_runtime.Value]](Get_Control_Category_categoryFn()).V1))
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Apply0_1_0 shape=App(Other) bindingType=(ADT ["Control","Apply","Apply"] [(TypeVar m)])
Apply0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}))
_ = Apply0_1_0
// TAST (Let): Functor0_2_1 shape=App(Other) bindingType=(ADT ["Data","Functor","Functor"] [(TypeVar m)])
Functor0_2_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_2_1
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_4_2_30 gopurs_runtime.Value
_ = go__go_4_2_30
// FALLBACK TCO: isLoop=false len=1
go__go_4_2_30 = gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
var __t8 gopurs_runtime.Value
{
var __t_tag_3 *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](v_5)
if (__t_tag_3 == nil) {
__t8 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "pure"), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(nil))}))})
goto end_branch_8
} else {

}
}
{
var __t_tag_4 *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](v_5)
if (__t_tag_4 != nil) {
// TAST (Let): __local_var_6_5 shape=Other bindingType=Any
__local_var_6_5 := (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_5.UnsafePtr).V0
_ = __local_var_6_5
// TAST (Let): __local_var_7_6 shape=Other bindingType=Any
__local_var_7_6 := (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_5.UnsafePtr).V2
_ = __local_var_7_6
// TAST (Let): __local_var_8_7 shape=Other bindingType=Any
__local_var_8_7 := (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_5.UnsafePtr).V1
_ = __local_var_8_7
__t8 = gopurs_runtime.Apply2(gopurs_runtime.Box(Apply0_1_0.V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Apply0_1_0.V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_2_1.V0), gopurs_runtime.Func3(func(l_prime__9 gopurs_runtime.Value, v_prime__10 gopurs_runtime.Value, r_prime__11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((&Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]{1, __local_var_6_5, __local_var_8_7, __local_var_7_6, v_prime__10, gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](l_prime__9), gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](r_prime__11)}))}
}), gopurs_runtime.Apply(go__go_4_2_30, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_5.UnsafePtr).V4)})), gopurs_runtime.Apply(f_3, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_5.UnsafePtr).V3)), gopurs_runtime.Apply(go__go_4_2_30, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_5.UnsafePtr).V5)}))
goto end_branch_8
} else {

}
}
{
__t8 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_8:
return __t8
})
return go__go_4_2_30
})
})})))}
	})
	return cache_Data_Map_Internal_traversableMap
}

var cache_Data_Map_Internal_traversableWithIndexMap gopurs_runtime.Value
var once_Data_Map_Internal_traversableWithIndexMap sync.Once
func Get_Data_Map_Internal_traversableWithIndexMap() gopurs_runtime.Value {
	once_Data_Map_Internal_traversableWithIndexMap.Do(func() {
		cache_Data_Map_Internal_traversableWithIndexMap = gopurs_runtime.Value{Type: 9, IntVal: 2078610234, UnsafePtr: unsafe.Pointer(Rebox_Data_Map_Internal_4018835873_1812164904((&Constructor_Data_TraversableWithIndex_TraversableWithIndex[gopurs_runtime.Value, *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 74250362, UnsafePtr: unsafe.Pointer(Rebox_Data_Map_Internal_1245395425_3725484264(Rebox_Data_Map_Internal_3725484264_1245395425(gopurs_runtime.CoerceToStruct[Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]](Get_Data_Map_Internal_foldableWithIndexMap()))))}
}), gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4077743418, UnsafePtr: unsafe.Pointer(Rebox_Data_Map_Internal_1397444001_2412140840(Rebox_Data_Map_Internal_2412140840_1397444001(gopurs_runtime.CoerceToStruct[Constructor_Data_FunctorWithIndex_FunctorWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]](Get_Data_Map_Internal_functorWithIndexMap()))))}
}), gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3941073978, UnsafePtr: unsafe.Pointer(Rebox_Data_Map_Internal_2141765991_3043886126(Rebox_Data_Map_Internal_3043886126_2141765991(gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]](Get_Data_Map_Internal_traversableMap()))))}
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Apply0_1_0 shape=App(Other) bindingType=(ADT ["Control","Apply","Apply"] [(TypeVar m)])
Apply0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}))
_ = Apply0_1_0
// TAST (Let): Functor0_2_1 shape=App(Other) bindingType=(ADT ["Data","Functor","Functor"] [(TypeVar m)])
Functor0_2_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_2_1
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_4_2_31 gopurs_runtime.Value
_ = go__go_4_2_31
// FALLBACK TCO: isLoop=false len=1
go__go_4_2_31 = gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
var __t8 gopurs_runtime.Value
{
var __t_tag_3 *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](v_5)
if (__t_tag_3 == nil) {
__t8 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "pure"), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(nil))}))})
goto end_branch_8
} else {

}
}
{
var __t_tag_4 *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](v_5)
if (__t_tag_4 != nil) {
// TAST (Let): __local_var_6_5 shape=Other bindingType=Any
__local_var_6_5 := (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_5.UnsafePtr).V0
_ = __local_var_6_5
// TAST (Let): __local_var_7_6 shape=Other bindingType=Any
__local_var_7_6 := (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_5.UnsafePtr).V2
_ = __local_var_7_6
// TAST (Let): __local_var_8_7 shape=Other bindingType=Any
__local_var_8_7 := (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_5.UnsafePtr).V1
_ = __local_var_8_7
__t8 = gopurs_runtime.Apply2(gopurs_runtime.Box(Apply0_1_0.V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Apply0_1_0.V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_2_1.V0), gopurs_runtime.Func3(func(l_prime__9 gopurs_runtime.Value, v_prime__10 gopurs_runtime.Value, r_prime__11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((&Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]{1, __local_var_6_5, __local_var_8_7, __local_var_7_6, v_prime__10, gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](l_prime__9), gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](r_prime__11)}))}
}), gopurs_runtime.Apply(go__go_4_2_31, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_5.UnsafePtr).V4)})), gopurs_runtime.Apply2(f_3, __local_var_7_6, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_5.UnsafePtr).V3)), gopurs_runtime.Apply(go__go_4_2_31, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_5.UnsafePtr).V5)}))
goto end_branch_8
} else {

}
}
{
__t8 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_8:
return __t8
})
return go__go_4_2_31
})
})})))}
	})
	return cache_Data_Map_Internal_traversableWithIndexMap
}

var cache_Data_Map_Internal_values gopurs_runtime.Value
var once_Data_Map_Internal_values sync.Once
func Get_Data_Map_Internal_values() gopurs_runtime.Value {
	once_Data_Map_Internal_values.Do(func() {
		cache_Data_Map_Internal_values = func() gopurs_runtime.Value {
var go__go_0_0_32 gopurs_runtime.Value
_ = go__go_0_0_32
// FALLBACK TCO: isLoop=false len=1
go__go_0_0_32 = gopurs_runtime.Func2(func(m_prime__1 gopurs_runtime.Value, __local_var_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t3 gopurs_runtime.Value
{
var __t_tag_1 *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](m_prime__1)
if (__t_tag_1 == nil) {
__t3 = __local_var_2
goto end_branch_3
} else {

}
}
{
var __t_tag_2 *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](m_prime__1)
if (__t_tag_2 != nil) {
__t3 = gopurs_runtime.UncurriedApp2(go__go_0_0_32, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(m_prime__1.UnsafePtr).V4)}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(m_prime__1.UnsafePtr).V3, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.UncurriedApp2(go__go_0_0_32, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(m_prime__1.UnsafePtr).V5)}, __local_var_2))}))})
goto end_branch_3
} else {

}
}
{
__t3 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_3:
return __t3
})
return gopurs_runtime.Func(func(m_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.UncurriedApp2(go__go_0_0_32, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](m_1))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil))})
})
}()
	})
	return cache_Data_Map_Internal_values
}

var cache_Data_Map_Internal_foldSubmapBy gopurs_runtime.Value
var once_Data_Map_Internal_foldSubmapBy sync.Once
func Get_Data_Map_Internal_foldSubmapBy() gopurs_runtime.Value {
	once_Data_Map_Internal_foldSubmapBy.Do(func() {
		cache_Data_Map_Internal_foldSubmapBy = gopurs_runtime.Func6(func(dictOrd_0_box gopurs_runtime.Value, appendFn_1_box gopurs_runtime.Value, memptyValue_2_box gopurs_runtime.Value, kmin_3_box gopurs_runtime.Value, kmax_4_box gopurs_runtime.Value, f_5_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Map_Internal_foldSubmapBy(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](dictOrd_0_box), appendFn_1_box, memptyValue_2_box, gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](kmin_3_box), gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](kmax_4_box), f_5_box)
})
	})
	return cache_Data_Map_Internal_foldSubmapBy
}

var cache_Data_Map_Internal_foldSubmap gopurs_runtime.Value
var once_Data_Map_Internal_foldSubmap sync.Once
func Get_Data_Map_Internal_foldSubmap() gopurs_runtime.Value {
	once_Data_Map_Internal_foldSubmap.Do(func() {
		cache_Data_Map_Internal_foldSubmap = gopurs_runtime.Func2(func(dictOrd_0_box gopurs_runtime.Value, dictMonoid_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Map_Internal_foldSubmap(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](dictOrd_0_box), gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]](dictMonoid_1_box))
})
	})
	return cache_Data_Map_Internal_foldSubmap
}

var cache_Data_Map_Internal_findMin gopurs_runtime.Value
var once_Data_Map_Internal_findMin sync.Once
func Get_Data_Map_Internal_findMin() gopurs_runtime.Value {
	once_Data_Map_Internal_findMin.Do(func() {
		cache_Data_Map_Internal_findMin = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				_v := Call_Data_Map_Internal_findMin(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](v_0_box))
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}()
})
	})
	return cache_Data_Map_Internal_findMin
}

var cache_Data_Map_Internal_lookupGT gopurs_runtime.Value
var once_Data_Map_Internal_lookupGT sync.Once
func Get_Data_Map_Internal_lookupGT() gopurs_runtime.Value {
	once_Data_Map_Internal_lookupGT.Do(func() {
		cache_Data_Map_Internal_lookupGT = gopurs_runtime.Func2(func(dictOrd_0_box gopurs_runtime.Value, k_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				_v := Call_Data_Map_Internal_lookupGT(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](dictOrd_0_box), k_1_box)
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}()
})
	})
	return cache_Data_Map_Internal_lookupGT
}

var cache_Data_Map_Internal_findMax gopurs_runtime.Value
var once_Data_Map_Internal_findMax sync.Once
func Get_Data_Map_Internal_findMax() gopurs_runtime.Value {
	once_Data_Map_Internal_findMax.Do(func() {
		cache_Data_Map_Internal_findMax = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				_v := Call_Data_Map_Internal_findMax(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](v_0_box))
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}()
})
	})
	return cache_Data_Map_Internal_findMax
}

var cache_Data_Map_Internal_lookupLT gopurs_runtime.Value
var once_Data_Map_Internal_lookupLT sync.Once
func Get_Data_Map_Internal_lookupLT() gopurs_runtime.Value {
	once_Data_Map_Internal_lookupLT.Do(func() {
		cache_Data_Map_Internal_lookupLT = gopurs_runtime.Func2(func(dictOrd_0_box gopurs_runtime.Value, k_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				_v := Call_Data_Map_Internal_lookupLT(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](dictOrd_0_box), k_1_box)
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}()
})
	})
	return cache_Data_Map_Internal_lookupLT
}

var cache_Data_Map_Internal_filterWithKey gopurs_runtime.Value
var once_Data_Map_Internal_filterWithKey sync.Once
func Get_Data_Map_Internal_filterWithKey() gopurs_runtime.Value {
	once_Data_Map_Internal_filterWithKey.Do(func() {
		cache_Data_Map_Internal_filterWithKey = gopurs_runtime.Func2(func(dictOrd_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Map_Internal_filterWithKey(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](dictOrd_0_box), f_1_box)
})
	})
	return cache_Data_Map_Internal_filterWithKey
}

var cache_Data_Map_Internal_filterKeys gopurs_runtime.Value
var once_Data_Map_Internal_filterKeys sync.Once
func Get_Data_Map_Internal_filterKeys() gopurs_runtime.Value {
	once_Data_Map_Internal_filterKeys.Do(func() {
		cache_Data_Map_Internal_filterKeys = gopurs_runtime.Func2(func(dictOrd_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Map_Internal_filterKeys(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](dictOrd_0_box), f_1_box)
})
	})
	return cache_Data_Map_Internal_filterKeys
}

var cache_Data_Map_Internal_filter gopurs_runtime.Value
var once_Data_Map_Internal_filter sync.Once
func Get_Data_Map_Internal_filter() gopurs_runtime.Value {
	once_Data_Map_Internal_filter.Do(func() {
		cache_Data_Map_Internal_filter = gopurs_runtime.Func2(func(dictOrd_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Map_Internal_filter(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](dictOrd_0_box), x_1_box)
})
	})
	return cache_Data_Map_Internal_filter
}

var cache_Data_Map_Internal_eqMap gopurs_runtime.Value
var once_Data_Map_Internal_eqMap sync.Once
func Get_Data_Map_Internal_eqMap() gopurs_runtime.Value {
	once_Data_Map_Internal_eqMap.Do(func() {
		cache_Data_Map_Internal_eqMap = gopurs_runtime.Func2(func(dictEq_0_box gopurs_runtime.Value, dictEq1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Map_Internal_eqMap(dictEq_0_box, dictEq1_1_box)
})
	})
	return cache_Data_Map_Internal_eqMap
}

var cache_Data_Map_Internal_ordMap gopurs_runtime.Value
var once_Data_Map_Internal_ordMap sync.Once
func Get_Data_Map_Internal_ordMap() gopurs_runtime.Value {
	once_Data_Map_Internal_ordMap.Do(func() {
		cache_Data_Map_Internal_ordMap = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Map_Internal_ordMap(dictOrd_0_box)
})
	})
	return cache_Data_Map_Internal_ordMap
}

var cache_Data_Map_Internal_eq1Map gopurs_runtime.Value
var once_Data_Map_Internal_eq1Map sync.Once
func Get_Data_Map_Internal_eq1Map() gopurs_runtime.Value {
	once_Data_Map_Internal_eq1Map.Do(func() {
		cache_Data_Map_Internal_eq1Map = gopurs_runtime.Func(func(dictEq_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Map_Internal_eq1Map(dictEq_0_box)
})
	})
	return cache_Data_Map_Internal_eq1Map
}

var cache_Data_Map_Internal_ord1Map gopurs_runtime.Value
var once_Data_Map_Internal_ord1Map sync.Once
func Get_Data_Map_Internal_ord1Map() gopurs_runtime.Value {
	once_Data_Map_Internal_ord1Map.Do(func() {
		cache_Data_Map_Internal_ord1Map = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Map_Internal_ord1Map(dictOrd_0_box)
})
	})
	return cache_Data_Map_Internal_ord1Map
}

var cache_Data_Map_Internal_empty gopurs_runtime.Value
var once_Data_Map_Internal_empty sync.Once
func Get_Data_Map_Internal_empty() gopurs_runtime.Value {
	once_Data_Map_Internal_empty.Do(func() {
		cache_Data_Map_Internal_empty = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(nil))}
	})
	return cache_Data_Map_Internal_empty
}

var cache_Data_Map_Internal_fromFoldable gopurs_runtime.Value
var once_Data_Map_Internal_fromFoldable sync.Once
func Get_Data_Map_Internal_fromFoldable() gopurs_runtime.Value {
	once_Data_Map_Internal_fromFoldable.Do(func() {
		cache_Data_Map_Internal_fromFoldable = gopurs_runtime.Func2(func(dictOrd_0_box gopurs_runtime.Value, dictFoldable_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Map_Internal_fromFoldable(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](dictOrd_0_box), gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](dictFoldable_1_box))
})
	})
	return cache_Data_Map_Internal_fromFoldable
}

var cache_Data_Map_Internal_fromFoldableWith gopurs_runtime.Value
var once_Data_Map_Internal_fromFoldableWith sync.Once
func Get_Data_Map_Internal_fromFoldableWith() gopurs_runtime.Value {
	once_Data_Map_Internal_fromFoldableWith.Do(func() {
		cache_Data_Map_Internal_fromFoldableWith = gopurs_runtime.Func3(func(dictOrd_0_box gopurs_runtime.Value, dictFoldable_1_box gopurs_runtime.Value, f_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Map_Internal_fromFoldableWith(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](dictOrd_0_box), gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](dictFoldable_1_box), f_2_box)
})
	})
	return cache_Data_Map_Internal_fromFoldableWith
}

var cache_Data_Map_Internal_fromFoldableWithIndex gopurs_runtime.Value
var once_Data_Map_Internal_fromFoldableWithIndex sync.Once
func Get_Data_Map_Internal_fromFoldableWithIndex() gopurs_runtime.Value {
	once_Data_Map_Internal_fromFoldableWithIndex.Do(func() {
		cache_Data_Map_Internal_fromFoldableWithIndex = gopurs_runtime.Func2(func(dictOrd_0_box gopurs_runtime.Value, dictFoldableWithIndex_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Map_Internal_fromFoldableWithIndex(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](dictOrd_0_box), gopurs_runtime.CoerceToStruct[Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]](dictFoldableWithIndex_1_box))
})
	})
	return cache_Data_Map_Internal_fromFoldableWithIndex
}

var cache_Data_Map_Internal_monoidSemigroupMap gopurs_runtime.Value
var once_Data_Map_Internal_monoidSemigroupMap sync.Once
func Get_Data_Map_Internal_monoidSemigroupMap() gopurs_runtime.Value {
	once_Data_Map_Internal_monoidSemigroupMap.Do(func() {
		cache_Data_Map_Internal_monoidSemigroupMap = gopurs_runtime.Func3(func(_dollar___unused_0_box gopurs_runtime.Value, dictOrd_1_box gopurs_runtime.Value, dictSemigroup_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Map_Internal_monoidSemigroupMap(_dollar___unused_0_box, dictOrd_1_box, dictSemigroup_2_box)
})
	})
	return cache_Data_Map_Internal_monoidSemigroupMap
}

var cache_Data_Map_Internal_submap gopurs_runtime.Value
var once_Data_Map_Internal_submap sync.Once
func Get_Data_Map_Internal_submap() gopurs_runtime.Value {
	once_Data_Map_Internal_submap.Do(func() {
		cache_Data_Map_Internal_submap = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Map_Internal_submap(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](dictOrd_0_box))
})
	})
	return cache_Data_Map_Internal_submap
}

var cache_Data_Map_Internal_unions gopurs_runtime.Value
var once_Data_Map_Internal_unions sync.Once
func Get_Data_Map_Internal_unions() gopurs_runtime.Value {
	once_Data_Map_Internal_unions.Do(func() {
		cache_Data_Map_Internal_unions = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Map_Internal_unions(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](dictOrd_0_box))
})
	})
	return cache_Data_Map_Internal_unions
}

var cache_Data_Map_Internal_difference gopurs_runtime.Value
var once_Data_Map_Internal_difference sync.Once
func Get_Data_Map_Internal_difference() gopurs_runtime.Value {
	once_Data_Map_Internal_difference.Do(func() {
		cache_Data_Map_Internal_difference = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Map_Internal_difference(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](dictOrd_0_box))
})
	})
	return cache_Data_Map_Internal_difference
}

var cache_Data_Map_Internal_go__delete gopurs_runtime.Value
var once_Data_Map_Internal_go__delete sync.Once
func Get_Data_Map_Internal_go__delete() gopurs_runtime.Value {
	once_Data_Map_Internal_go__delete.Do(func() {
		cache_Data_Map_Internal_go__delete = gopurs_runtime.Func2(func(dictOrd_0_box gopurs_runtime.Value, k_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Map_Internal_go__delete(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](dictOrd_0_box), k_1_box)
})
	})
	return cache_Data_Map_Internal_go__delete
}

var cache_Data_Map_Internal_checkValid gopurs_runtime.Value
var once_Data_Map_Internal_checkValid sync.Once
func Get_Data_Map_Internal_checkValid() gopurs_runtime.Value {
	once_Data_Map_Internal_checkValid.Do(func() {
		cache_Data_Map_Internal_checkValid = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Map_Internal_checkValid(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](dictOrd_0_box))
})
	})
	return cache_Data_Map_Internal_checkValid
}

var cache_Data_Map_Internal_catMaybes gopurs_runtime.Value
var once_Data_Map_Internal_catMaybes sync.Once
func Get_Data_Map_Internal_catMaybes() gopurs_runtime.Value {
	once_Data_Map_Internal_catMaybes.Do(func() {
		cache_Data_Map_Internal_catMaybes = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Map_Internal_catMaybes(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](dictOrd_0_box))
})
	})
	return cache_Data_Map_Internal_catMaybes
}

var cache_Data_Map_Internal_applyMap gopurs_runtime.Value
var once_Data_Map_Internal_applyMap sync.Once
func Get_Data_Map_Internal_applyMap() gopurs_runtime.Value {
	once_Data_Map_Internal_applyMap.Do(func() {
		cache_Data_Map_Internal_applyMap = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Map_Internal_applyMap(dictOrd_0_box)
})
	})
	return cache_Data_Map_Internal_applyMap
}

var cache_Data_Map_Internal_bindMap gopurs_runtime.Value
var once_Data_Map_Internal_bindMap sync.Once
func Get_Data_Map_Internal_bindMap() gopurs_runtime.Value {
	once_Data_Map_Internal_bindMap.Do(func() {
		cache_Data_Map_Internal_bindMap = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Map_Internal_bindMap(dictOrd_0_box)
})
	})
	return cache_Data_Map_Internal_bindMap
}

var cache_Data_Map_Internal_anyWithKey gopurs_runtime.Value
var once_Data_Map_Internal_anyWithKey sync.Once
func Get_Data_Map_Internal_anyWithKey() gopurs_runtime.Value {
	once_Data_Map_Internal_anyWithKey.Do(func() {
		cache_Data_Map_Internal_anyWithKey = gopurs_runtime.Func(func(predicate_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Map_Internal_anyWithKey(predicate_0_box)
})
	})
	return cache_Data_Map_Internal_anyWithKey
}

var cache_Data_Map_Internal_any gopurs_runtime.Value
var once_Data_Map_Internal_any sync.Once
func Get_Data_Map_Internal_any() gopurs_runtime.Value {
	once_Data_Map_Internal_any.Do(func() {
		cache_Data_Map_Internal_any = gopurs_runtime.Func(func(predicate_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Map_Internal_any(predicate_0_box)
})
	})
	return cache_Data_Map_Internal_any
}

var cache_Data_Map_Internal_alter gopurs_runtime.Value
var once_Data_Map_Internal_alter sync.Once
func Get_Data_Map_Internal_alter() gopurs_runtime.Value {
	once_Data_Map_Internal_alter.Do(func() {
		cache_Data_Map_Internal_alter = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Map_Internal_alter(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](dictOrd_0_box))
})
	})
	return cache_Data_Map_Internal_alter
}

var cache_Data_Map_Internal_altMap gopurs_runtime.Value
var once_Data_Map_Internal_altMap sync.Once
func Get_Data_Map_Internal_altMap() gopurs_runtime.Value {
	once_Data_Map_Internal_altMap.Do(func() {
		cache_Data_Map_Internal_altMap = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Map_Internal_altMap(dictOrd_0_box)
})
	})
	return cache_Data_Map_Internal_altMap
}

var cache_Data_Map_Internal_plusMap gopurs_runtime.Value
var once_Data_Map_Internal_plusMap sync.Once
func Get_Data_Map_Internal_plusMap() gopurs_runtime.Value {
	once_Data_Map_Internal_plusMap.Do(func() {
		cache_Data_Map_Internal_plusMap = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Map_Internal_plusMap(dictOrd_0_box)
})
	})
	return cache_Data_Map_Internal_plusMap
}

type Constructor_Data_Map_Internal_Leaf[T_k any, T_v any] struct {
	Rc uint32
}


type Constructor_Data_Map_Internal_Node[T_k any, T_v any] struct {
	Rc uint32
	V0 int64
	V1 int64
	V2 T_k
	V3 T_v
	V4 *Constructor_Data_Map_Internal_Node[T_k, T_v]
	V5 *Constructor_Data_Map_Internal_Node[T_k, T_v]
}


type Constructor_Data_Map_Internal_IterLeaf[T_k any, T_v any] struct {
	Rc uint32
}


type Constructor_Data_Map_Internal_IterEmit[T_k any, T_v any] struct {
	Rc uint32
	V0 T_k
	V1 T_v
	V2 gopurs_runtime.Value
}


type Constructor_Data_Map_Internal_IterNode[T_k any, T_v any] struct {
	Rc uint32
	V0 *Constructor_Data_Map_Internal_Node[T_k, T_v]
	V1 gopurs_runtime.Value
}


type Constructor_Data_Map_Internal_IterDone[T_k any, T_v any] struct {
	Rc uint32
}


type Constructor_Data_Map_Internal_IterNext[T_k any, T_v any] struct {
	Rc uint32
	V0 T_k
	V1 T_v
	V2 gopurs_runtime.Value
}


type Constructor_Data_Map_Internal_Split[T_k any, T_v any] struct {
	Rc uint32
	V0 *Constructor_Data_Maybe_Just[T_v]
	V1 *Constructor_Data_Map_Internal_Node[T_k, T_v]
	V2 *Constructor_Data_Map_Internal_Node[T_k, T_v]
}


type Constructor_Data_Map_Internal_SplitLast[T_k any, T_v any] struct {
	Rc uint32
	V0 T_k
	V1 T_v
	V2 *Constructor_Data_Map_Internal_Node[T_k, T_v]
}


func Call_Data_Map_Internal_Node__4238775810(__eta_norm_1_unused_0_loop int64, __eta_norm_0_unused_1_loop int64, __eta_norm_5_2_loop uint32, __eta_norm_4_3_loop float64, __eta_norm_3_4_loop *Constructor_Data_Map_Internal_Node[uint32, float64], __eta_norm_2_5_loop *Constructor_Data_Map_Internal_Node[uint32, float64]) *Constructor_Data_Map_Internal_Node[uint32, float64] {
Node__4238775810:
for {
if false { continue Node__4238775810 }
var __eta_norm_1_unused_0 int64 = __eta_norm_1_unused_0_loop
_ = __eta_norm_1_unused_0
var __eta_norm_0_unused_1 int64 = __eta_norm_0_unused_1_loop
_ = __eta_norm_0_unused_1
var __eta_norm_5_2 uint32 = __eta_norm_5_2_loop
_ = __eta_norm_5_2
var __eta_norm_4_3 float64 = __eta_norm_4_3_loop
_ = __eta_norm_4_3
var __eta_norm_3_4 *Constructor_Data_Map_Internal_Node[uint32, float64] = __eta_norm_3_4_loop
_ = __eta_norm_3_4
var __eta_norm_2_5 *Constructor_Data_Map_Internal_Node[uint32, float64] = __eta_norm_2_5_loop
_ = __eta_norm_2_5
return Rebox_Data_Map_Internal_2487766124_261879545(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply6(gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(value1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(value2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(value3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(value4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(value5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((&Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]{1, value0.IntVal, value1.IntVal, value2, value3, gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](value4), gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](value5)}))}
})
})
})
})
})
}), gopurs_runtime.Value{Type: 9, IntVal: int64(__eta_norm_5_2), UnsafePtr: nil}, gopurs_runtime.Float(__eta_norm_4_3), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(Rebox_Data_Map_Internal_261879545_2487766124(__eta_norm_3_4))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(Rebox_Data_Map_Internal_261879545_2487766124(__eta_norm_2_5))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(Rebox_Data_Map_Internal_261879545_2487766124((*Constructor_Data_Map_Internal_Node[uint32, float64])(nil)))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(Rebox_Data_Map_Internal_261879545_2487766124((*Constructor_Data_Map_Internal_Node[uint32, float64])(nil)))})))
}
}

func Call_Data_Map_Internal_unsafeNode(k_0_loop gopurs_runtime.Value, __local_var_1_loop gopurs_runtime.Value, __local_var_2_loop gopurs_runtime.Value, __local_var_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
var k_0 gopurs_runtime.Value = k_0_loop
_ = k_0
var __local_var_1 gopurs_runtime.Value = __local_var_1_loop
_ = __local_var_1
var __local_var_2 gopurs_runtime.Value = __local_var_2_loop
_ = __local_var_2
var __local_var_3 gopurs_runtime.Value = __local_var_3_loop
_ = __local_var_3
var __t9 *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]
{
var __t_tag_0 *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__local_var_2)
if (__t_tag_0 == nil) {
var __t3 *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]
{
var __t_tag_1 *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__local_var_3)
if (__t_tag_1 == nil) {
__t3 = (&Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]{1, int64(1), int64(1), k_0, __local_var_1, gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__local_var_2), gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__local_var_3)})
goto end_branch_3
} else {

}
}
{
var __t_tag_2 *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__local_var_3)
if (__t_tag_2 != nil) {
__t3 = (&Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]{1, (int64(1)) + ((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V0), (int64(1)) + ((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V1), k_0, __local_var_1, gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__local_var_2), gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__local_var_3)})
goto end_branch_3
} else {

}
}
{
__t3 = func() *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value] { panic("Failed pattern match") }()
}
end_branch_3:
__t9 = __t3
goto end_branch_9
} else {

}
}
{
var __t_tag_4 *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__local_var_2)
if (__t_tag_4 != nil) {
var __t8 *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]
{
var __t_tag_5 *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__local_var_3)
if (__t_tag_5 == nil) {
__t8 = (&Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]{1, (int64(1)) + ((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V0), (int64(1)) + ((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V1), k_0, __local_var_1, gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__local_var_2), gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__local_var_3)})
goto end_branch_8
} else {

}
}
{
var __t_tag_6 *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__local_var_3)
if (__t_tag_6 != nil) {
var __t7 int64
{
if ((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V0) > ((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V0) {
__t7 = (int64(1)) + ((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V0)
goto end_branch_7
} else {

}
}
{
__t7 = (int64(1)) + ((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V0)
}
end_branch_7:
__t8 = (&Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]{1, __t7, ((int64(1)) + ((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V1)) + ((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V1), k_0, __local_var_1, gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__local_var_2), gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__local_var_3)})
goto end_branch_8
} else {

}
}
{
__t8 = func() *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value] { panic("Failed pattern match") }()
}
end_branch_8:
__t9 = __t8
goto end_branch_9
} else {

}
}
{
__t9 = func() *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value] { panic("Failed pattern match") }()
}
end_branch_9:
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(__t9)}
}

func Call_Data_Map_Internal_toMapIter(a_0_loop *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var a_0 *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value] = a_0_loop
_ = a_0
return gopurs_runtime.Value{Type: 9, IntVal: 2861335956, UnsafePtr: unsafe.Pointer((&Constructor_Data_Map_Internal_IterNode[gopurs_runtime.Value, gopurs_runtime.Value]{1, a_0, gopurs_runtime.Value{Type: 9, IntVal: 2509360378, UnsafePtr: unsafe.Pointer(nil)}}))}
}

func Call_Data_Map_Internal_stepWith(f_0_loop gopurs_runtime.Value, next_1_loop gopurs_runtime.Value, done_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var next_1 gopurs_runtime.Value = next_1_loop
_ = next_1
var done_2 gopurs_runtime.Value = done_2_loop
_ = done_2
var Call_local_Data_Map_Internal_go__go_3_0_0 func(gopurs_runtime.Value) gopurs_runtime.Value
_ = Call_local_Data_Map_Internal_go__go_3_0_0
var go__go_3_0_0 gopurs_runtime.Value
_ = go__go_3_0_0
Call_local_Data_Map_Internal_go__go_3_0_0 = func(v_4_loop gopurs_runtime.Value) gopurs_runtime.Value {
go__go_3_0_0:
for {
if false { continue go__go_3_0_0 }
var v_4 gopurs_runtime.Value = v_4_loop
_ = v_4
var __t1 gopurs_runtime.Value
{
if (v_4.Type == 9 && v_4.IntVal == 2509360378) {
__t1 = gopurs_runtime.Apply(done_2, Get_Data_Unit_unit())
goto end_branch_1
} else {

}
}
{
if (v_4.Type == 9 && v_4.IntVal == 1343415489) {
__t1 = gopurs_runtime.UncurriedApp3(next_1, (*Constructor_Data_Map_Internal_IterEmit[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V0, (*Constructor_Data_Map_Internal_IterEmit[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V1, (*Constructor_Data_Map_Internal_IterEmit[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V2)
goto end_branch_1
} else {

}
}
{
if (v_4.Type == 9 && v_4.IntVal == 2861335956) {
v_4_loop = gopurs_runtime.Apply2(f_0, (*Constructor_Data_Map_Internal_IterNode[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_IterNode[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V0)})
continue go__go_3_0_0
__t1 = func() gopurs_runtime.Value { panic("unreachable") }()
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
}
go__go_3_0_0 = gopurs_runtime.Func(func(v_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return Call_local_Data_Map_Internal_go__go_3_0_0(v_4_loop_val)
})
return go__go_3_0_0
}

func Call_Data_Map_Internal_size(v_0_loop *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]) int64 {
var v_0 *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value] = v_0_loop
_ = v_0
var __t0 int64
{
if (v_0 == nil) {
__t0 = int64(0)
goto end_branch_0
} else {

}
}
{
if (v_0 != nil) {
__t0 = (v_0).V1
goto end_branch_0
} else {

}
}
{
__t0 = func() int64 { panic("Failed pattern match") }()
}
end_branch_0:
return __t0
}

func Call_Data_Map_Internal_size__1004260195(v_0_loop *Constructor_Data_Map_Internal_Node[uint32, float64]) int64 {
size__1004260195:
for {
if false { continue size__1004260195 }
var v_0 *Constructor_Data_Map_Internal_Node[uint32, float64] = v_0_loop
_ = v_0
var __t0 int64
{
if (v_0 == nil) {
__t0 = int64(0)
goto end_branch_0
} else {

}
}
{
if (v_0 != nil) {
__t0 = (v_0).V1
goto end_branch_0
} else {

}
}
{
__t0 = func() int64 { panic("Failed pattern match") }()
}
end_branch_0:
return __t0
}
}

func Call_Data_Map_Internal_singleton(k_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value] {
var k_0 gopurs_runtime.Value = k_0_loop
_ = k_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return (&Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]{1, int64(1), int64(1), k_0, v_1, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(nil), (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(nil)})
}

func Call_Data_Map_Internal_singleton__981367005(k_0_loop uint32, v_1_loop float64) *Constructor_Data_Map_Internal_Node[uint32, float64] {
singleton__981367005:
for {
if false { continue singleton__981367005 }
var k_0 uint32 = k_0_loop
_ = k_0
var v_1 float64 = v_1_loop
_ = v_1
return (&Constructor_Data_Map_Internal_Node[uint32, float64]{1, gopurs_runtime.Value{Type: 9, IntVal: int64(k_0), UnsafePtr: nil}.IntVal, gopurs_runtime.Float(v_1).IntVal, uint32(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(Rebox_Data_Map_Internal_261879545_2487766124((*Constructor_Data_Map_Internal_Node[uint32, float64])(nil)))}.IntVal), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(Rebox_Data_Map_Internal_261879545_2487766124((*Constructor_Data_Map_Internal_Node[uint32, float64])(nil)))}.FloatVal(), (*Constructor_Data_Map_Internal_Node[uint32, float64])(nil), (*Constructor_Data_Map_Internal_Node[uint32, float64])(nil)})
}
}

func Call_Data_Map_Internal_unsafeBalancedNode(k_0_loop gopurs_runtime.Value, __local_var_1_loop gopurs_runtime.Value, __local_var_2_loop gopurs_runtime.Value, __local_var_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
var k_0 gopurs_runtime.Value = k_0_loop
_ = k_0
var __local_var_1 gopurs_runtime.Value = __local_var_1_loop
_ = __local_var_1
var __local_var_2 gopurs_runtime.Value = __local_var_2_loop
_ = __local_var_2
var __local_var_3 gopurs_runtime.Value = __local_var_3_loop
_ = __local_var_3
var __t33 *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]
{
var __t_tag_0 *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__local_var_2)
if (__t_tag_0 == nil) {
var __t9 *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]
{
var __t_tag_1 *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__local_var_3)
if (__t_tag_1 == nil) {
__t9 = (&Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]{1, int64(1), int64(1), k_0, __local_var_1, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(nil), (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(nil)})
goto end_branch_9
} else {

}
}
{
var __t_tag_2 *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__local_var_3)
if ((__t_tag_2 != nil)) && (((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V0) > (int64(1))) {
var __t8 *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]
{
var __t_tag_3 *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value] = (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V4
var __t_and_7 bool = false
if (__t_tag_3 != nil) {

var __t6 bool
{
var __t_tag_4 *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value] = (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V5
if (__t_tag_4 == nil) {
__t6 = (((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V4).V0) > (int64(0))
goto end_branch_6
} else {

}
}
{
var __t_tag_5 *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value] = (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V5
if (__t_tag_5 != nil) {
__t6 = (((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V4).V0) > (((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V5).V0)
goto end_branch_6
} else {

}
}
{
__t6 = func() bool { panic("Failed pattern match") }()
}
end_branch_6:
__t_and_7 = __t6
}
if __t_and_7 {
__t8 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeNode(), ((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V4).V2, ((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V4).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeNode(), k_0, __local_var_1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__local_var_2))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V4).V4)})))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeNode(), (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V4).V5)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V5)})))}))
goto end_branch_8
} else {

}
}
{
__t8 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeNode(), (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeNode(), k_0, __local_var_1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__local_var_2))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V4)})))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V5)}))
}
end_branch_8:
__t9 = __t8
goto end_branch_9
} else {

}
}
{
__t9 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeNode(), k_0, __local_var_1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__local_var_2))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__local_var_3))}))
}
end_branch_9:
__t33 = __t9
goto end_branch_33
} else {

}
}
{
var __t_tag_10 *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__local_var_2)
if (__t_tag_10 != nil) {
var __t32 *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]
{
var __t_tag_11 *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__local_var_3)
if (__t_tag_11 != nil) {
var __t24 *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]
{
if ((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V0) > (((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V0) + (int64(1))) {
var __t17 *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]
{
var __t_tag_12 *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value] = (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V4
var __t_and_16 bool = false
if (__t_tag_12 != nil) {

var __t15 bool
{
var __t_tag_13 *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value] = (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V5
if (__t_tag_13 == nil) {
__t15 = (((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V4).V0) > (int64(0))
goto end_branch_15
} else {

}
}
{
var __t_tag_14 *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value] = (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V5
if (__t_tag_14 != nil) {
__t15 = (((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V4).V0) > (((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V5).V0)
goto end_branch_15
} else {

}
}
{
__t15 = func() bool { panic("Failed pattern match") }()
}
end_branch_15:
__t_and_16 = __t15
}
if __t_and_16 {
__t17 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeNode(), ((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V4).V2, ((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V4).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeNode(), k_0, __local_var_1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__local_var_2))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V4).V4)})))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeNode(), (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V4).V5)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V5)})))}))
goto end_branch_17
} else {

}
}
{
__t17 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeNode(), (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeNode(), k_0, __local_var_1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__local_var_2))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V4)})))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V5)}))
}
end_branch_17:
__t24 = __t17
goto end_branch_24
} else {

}
}
{
if ((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V0) > (((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V0) + (int64(1))) {
var __t23 *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]
{
var __t_tag_18 *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value] = (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V5
var __t_and_22 bool = false
if (__t_tag_18 != nil) {

var __t21 bool
{
var __t_tag_19 *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value] = (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V4
if (__t_tag_19 == nil) {
__t21 = (int64(0)) <= (((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V5).V0)
goto end_branch_21
} else {

}
}
{
var __t_tag_20 *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value] = (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V4
if (__t_tag_20 != nil) {
__t21 = (((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V4).V0) <= (((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V5).V0)
goto end_branch_21
} else {

}
}
{
__t21 = func() bool { panic("Failed pattern match") }()
}
end_branch_21:
__t_and_22 = __t21
}
if __t_and_22 {
__t23 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeNode(), ((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V5).V2, ((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V5).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeNode(), (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V4)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V5).V4)})))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeNode(), k_0, __local_var_1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V5).V5)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__local_var_3))})))}))
goto end_branch_23
} else {

}
}
{
__t23 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeNode(), (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V4)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeNode(), k_0, __local_var_1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V5)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__local_var_3))})))}))
}
end_branch_23:
__t24 = __t23
goto end_branch_24
} else {

}
}
{
__t24 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeNode(), k_0, __local_var_1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__local_var_2))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__local_var_3))}))
}
end_branch_24:
__t32 = __t24
goto end_branch_32
} else {

}
}
{
var __t_tag_25 *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__local_var_3)
if ((__t_tag_25 == nil)) && (((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V0) > (int64(1))) {
var __t31 *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]
{
var __t_tag_26 *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value] = (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V5
var __t_and_30 bool = false
if (__t_tag_26 != nil) {

var __t29 bool
{
var __t_tag_27 *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value] = (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V4
if (__t_tag_27 == nil) {
__t29 = (int64(0)) <= (((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V5).V0)
goto end_branch_29
} else {

}
}
{
var __t_tag_28 *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value] = (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V4
if (__t_tag_28 != nil) {
__t29 = (((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V4).V0) <= (((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V5).V0)
goto end_branch_29
} else {

}
}
{
__t29 = func() bool { panic("Failed pattern match") }()
}
end_branch_29:
__t_and_30 = __t29
}
if __t_and_30 {
__t31 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeNode(), ((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V5).V2, ((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V5).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeNode(), (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V4)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V5).V4)})))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeNode(), k_0, __local_var_1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V5).V5)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__local_var_3))})))}))
goto end_branch_31
} else {

}
}
{
__t31 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeNode(), (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V4)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeNode(), k_0, __local_var_1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V5)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__local_var_3))})))}))
}
end_branch_31:
__t32 = __t31
goto end_branch_32
} else {

}
}
{
__t32 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeNode(), k_0, __local_var_1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__local_var_2))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__local_var_3))}))
}
end_branch_32:
__t33 = __t32
goto end_branch_33
} else {

}
}
{
__t33 = func() *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value] { panic("Failed pattern match") }()
}
end_branch_33:
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(__t33)}
}

func Call_Data_Map_Internal_unsafeSplit(comp_0_loop gopurs_runtime.Value, __local_var_1_loop gopurs_runtime.Value, __local_var_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
unsafeSplit:
for {
if false { continue unsafeSplit }
var comp_0 gopurs_runtime.Value = comp_0_loop
_ = comp_0
var __local_var_1 gopurs_runtime.Value = __local_var_1_loop
_ = __local_var_1
var __local_var_2 gopurs_runtime.Value = __local_var_2_loop
_ = __local_var_2
var __t6 *Constructor_Data_Map_Internal_Split[gopurs_runtime.Value, gopurs_runtime.Value]
{
var __t_tag_0 *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__local_var_2)
if (__t_tag_0 == nil) {
__t6 = (&Constructor_Data_Map_Internal_Split[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil), (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(nil), (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(nil)})
goto end_branch_6
} else {

}
}
{
var __t_tag_1 *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__local_var_2)
if (__t_tag_1 != nil) {
// TAST (Let): v_3_2 shape=App(Other) bindingType=(ADT ["Data","Ordering","Ordering"] [])
v_3_2 := uint32(gopurs_runtime.Apply2(comp_0, __local_var_1, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V2).IntVal)
_ = v_3_2
var __t5 *Constructor_Data_Map_Internal_Split[gopurs_runtime.Value, gopurs_runtime.Value]
{
if (v_3_2 == 1527465420) {
// TAST (Let): v1_4_3 shape=UncurriedApp(Var) bindingType=(ADT ["Data","Map","Internal","Split"] [(TypeVar k), (TypeVar v)])
v1_4_3 := gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Split[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp3(Get_Data_Map_Internal_unsafeSplit(), comp_0, __local_var_1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V4)}))
_ = v1_4_3
__t5 = (&Constructor_Data_Map_Internal_Split[gopurs_runtime.Value, gopurs_runtime.Value]{1, (v1_4_3).V0, (v1_4_3).V1, gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeBalancedNode(), (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((v1_4_3).V2)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V5)}))})
goto end_branch_5
} else {

}
}
{
if (v_3_2 == 380165415) {
// TAST (Let): v1_4_4 shape=UncurriedApp(Var) bindingType=(ADT ["Data","Map","Internal","Split"] [(TypeVar k), (TypeVar v)])
v1_4_4 := gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Split[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp3(Get_Data_Map_Internal_unsafeSplit(), comp_0, __local_var_1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V5)}))
_ = v1_4_4
__t5 = (&Constructor_Data_Map_Internal_Split[gopurs_runtime.Value, gopurs_runtime.Value]{1, (v1_4_4).V0, gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeBalancedNode(), (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V4)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((v1_4_4).V1)})), (v1_4_4).V2})
goto end_branch_5
} else {

}
}
{
if (v_3_2 == 902936544) {
__t5 = (&Constructor_Data_Map_Internal_Split[gopurs_runtime.Value, gopurs_runtime.Value]{1, (&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{1, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V3}), (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V4, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V5})
goto end_branch_5
} else {

}
}
{
__t5 = func() *Constructor_Data_Map_Internal_Split[gopurs_runtime.Value, gopurs_runtime.Value] { panic("Failed pattern match") }()
}
end_branch_5:
__t6 = __t5
goto end_branch_6
} else {

}
}
{
__t6 = func() *Constructor_Data_Map_Internal_Split[gopurs_runtime.Value, gopurs_runtime.Value] { panic("Failed pattern match") }()
}
end_branch_6:
return gopurs_runtime.Value{Type: 9, IntVal: 3373277644, UnsafePtr: unsafe.Pointer(__t6)}
}
}

func Call_Data_Map_Internal_unsafeSplitLast(k_0_loop gopurs_runtime.Value, __local_var_1_loop gopurs_runtime.Value, __local_var_2_loop gopurs_runtime.Value, __local_var_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
unsafeSplitLast:
for {
if false { continue unsafeSplitLast }
var k_0 gopurs_runtime.Value = k_0_loop
_ = k_0
var __local_var_1 gopurs_runtime.Value = __local_var_1_loop
_ = __local_var_1
var __local_var_2 gopurs_runtime.Value = __local_var_2_loop
_ = __local_var_2
var __local_var_3 gopurs_runtime.Value = __local_var_3_loop
_ = __local_var_3
var __t3 *Constructor_Data_Map_Internal_SplitLast[gopurs_runtime.Value, gopurs_runtime.Value]
{
var __t_tag_0 *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__local_var_3)
if (__t_tag_0 == nil) {
__t3 = (&Constructor_Data_Map_Internal_SplitLast[gopurs_runtime.Value, gopurs_runtime.Value]{1, k_0, __local_var_1, gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__local_var_2)})
goto end_branch_3
} else {

}
}
{
var __t_tag_1 *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__local_var_3)
if (__t_tag_1 != nil) {
// TAST (Let): v1_4_2 shape=UncurriedApp(Var) bindingType=(ADT ["Data","Map","Internal","SplitLast"] [(TypeVar k), (TypeVar v)])
v1_4_2 := gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_SplitLast[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeSplitLast(), (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V4)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V5)}))
_ = v1_4_2
__t3 = (&Constructor_Data_Map_Internal_SplitLast[gopurs_runtime.Value, gopurs_runtime.Value]{1, (v1_4_2).V0, (v1_4_2).V1, gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeBalancedNode(), k_0, __local_var_1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__local_var_2))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((v1_4_2).V2)}))})
goto end_branch_3
} else {

}
}
{
__t3 = func() *Constructor_Data_Map_Internal_SplitLast[gopurs_runtime.Value, gopurs_runtime.Value] { panic("Failed pattern match") }()
}
end_branch_3:
return gopurs_runtime.Value{Type: 9, IntVal: 2668112006, UnsafePtr: unsafe.Pointer(__t3)}
}
}

func Call_Data_Map_Internal_unsafeJoinNodes(v_0_loop gopurs_runtime.Value, __local_var_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var __local_var_1 gopurs_runtime.Value = __local_var_1_loop
_ = __local_var_1
var __t3 *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]
{
var __t_tag_0 *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](v_0)
if (__t_tag_0 == nil) {
__t3 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__local_var_1)
goto end_branch_3
} else {

}
}
{
var __t_tag_1 *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](v_0)
if (__t_tag_1 != nil) {
// TAST (Let): v2_2_2 shape=UncurriedApp(Var) bindingType=(ADT ["Data","Map","Internal","SplitLast"] [(TypeVar k), (TypeVar v)])
v2_2_2 := gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_SplitLast[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeSplitLast(), (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_0.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_0.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_0.UnsafePtr).V4)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_0.UnsafePtr).V5)}))
_ = v2_2_2
__t3 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeBalancedNode(), (v2_2_2).V0, (v2_2_2).V1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((v2_2_2).V2)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__local_var_1))}))
goto end_branch_3
} else {

}
}
{
__t3 = func() *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value] { panic("Failed pattern match") }()
}
end_branch_3:
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(__t3)}
}

func Call_Data_Map_Internal_unsafeDifference(comp_0_loop gopurs_runtime.Value, __local_var_1_loop gopurs_runtime.Value, __local_var_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
unsafeDifference:
for {
if false { continue unsafeDifference }
var comp_0 gopurs_runtime.Value = comp_0_loop
_ = comp_0
var __local_var_1 gopurs_runtime.Value = __local_var_1_loop
_ = __local_var_1
var __local_var_2 gopurs_runtime.Value = __local_var_2_loop
_ = __local_var_2
var __t4 *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]
{
var __t_tag_0 *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__local_var_1)
if (__t_tag_0 == nil) {
__t4 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(nil))})
goto end_branch_4
} else {

}
}
{
var __t_tag_1 *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__local_var_2)
if (__t_tag_1 == nil) {
__t4 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__local_var_1)
goto end_branch_4
} else {

}
}
{
var __t_tag_2 *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__local_var_2)
if (__t_tag_2 != nil) {
// TAST (Let): v_3_3 shape=UncurriedApp(Var) bindingType=(ADT ["Data","Map","Internal","Split"] [(TypeVar k), (TypeVar v)])
v_3_3 := gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Split[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp3(Get_Data_Map_Internal_unsafeSplit(), comp_0, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V2, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__local_var_1))}))
_ = v_3_3
__t4 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp2(Get_Data_Map_Internal_unsafeJoinNodes(), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp3(Get_Data_Map_Internal_unsafeDifference(), comp_0, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((v_3_3).V1)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V4)})))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp3(Get_Data_Map_Internal_unsafeDifference(), comp_0, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((v_3_3).V2)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V5)})))}))
goto end_branch_4
} else {

}
}
{
__t4 = func() *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value] { panic("Failed pattern match") }()
}
end_branch_4:
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(__t4)}
}
}

func Call_Data_Map_Internal_unsafeIntersectionWith(comp_0_loop gopurs_runtime.Value, __local_var_1_loop gopurs_runtime.Value, __local_var_2_loop gopurs_runtime.Value, __local_var_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
unsafeIntersectionWith:
for {
if false { continue unsafeIntersectionWith }
var comp_0 gopurs_runtime.Value = comp_0_loop
_ = comp_0
var __local_var_1 gopurs_runtime.Value = __local_var_1_loop
_ = __local_var_1
var __local_var_2 gopurs_runtime.Value = __local_var_2_loop
_ = __local_var_2
var __local_var_3 gopurs_runtime.Value = __local_var_3_loop
_ = __local_var_3
var __t9 *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]
{
var __t_tag_0 *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__local_var_2)
if (__t_tag_0 == nil) {
__t9 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(nil))})
goto end_branch_9
} else {

}
}
{
var __t_tag_1 *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__local_var_3)
if (__t_tag_1 == nil) {
__t9 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(nil))})
goto end_branch_9
} else {

}
}
{
var __t_tag_2 *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__local_var_3)
if (__t_tag_2 != nil) {
// TAST (Let): v_4_3 shape=UncurriedApp(Var) bindingType=(ADT ["Data","Map","Internal","Split"] [(TypeVar k), (TypeVar a)])
v_4_3 := gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Split[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp3(Get_Data_Map_Internal_unsafeSplit(), comp_0, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V2, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__local_var_2))}))
_ = v_4_3
// TAST (Let): l_prime__5_4 shape=UncurriedApp(Var) bindingType=Any
l_prime__5_4 := gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeIntersectionWith(), comp_0, __local_var_1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((v_4_3).V1)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V4)})
_ = l_prime__5_4
// TAST (Let): r_prime__6_5 shape=UncurriedApp(Var) bindingType=Any
r_prime__6_5 := gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeIntersectionWith(), comp_0, __local_var_1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((v_4_3).V2)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V5)})
_ = r_prime__6_5
var __t8 *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]
{
var __t_tag_6 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = (v_4_3).V0
if (__t_tag_6 != nil) {
__t8 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeBalancedNode(), (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V2, gopurs_runtime.Apply2(__local_var_1, ((v_4_3).V0).V0, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V3), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](l_prime__5_4))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](r_prime__6_5))}))
goto end_branch_8
} else {

}
}
{
var __t_tag_7 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = (v_4_3).V0
if (__t_tag_7 == nil) {
__t8 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp2(Get_Data_Map_Internal_unsafeJoinNodes(), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](l_prime__5_4))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](r_prime__6_5))}))
goto end_branch_8
} else {

}
}
{
__t8 = func() *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value] { panic("Failed pattern match") }()
}
end_branch_8:
__t9 = __t8
goto end_branch_9
} else {

}
}
{
__t9 = func() *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value] { panic("Failed pattern match") }()
}
end_branch_9:
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(__t9)}
}
}

func Call_Data_Map_Internal_unsafeUnionWith(comp_0_loop gopurs_runtime.Value, __local_var_1_loop gopurs_runtime.Value, __local_var_2_loop gopurs_runtime.Value, __local_var_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
unsafeUnionWith:
for {
if false { continue unsafeUnionWith }
var comp_0 gopurs_runtime.Value = comp_0_loop
_ = comp_0
var __local_var_1 gopurs_runtime.Value = __local_var_1_loop
_ = __local_var_1
var __local_var_2 gopurs_runtime.Value = __local_var_2_loop
_ = __local_var_2
var __local_var_3 gopurs_runtime.Value = __local_var_3_loop
_ = __local_var_3
var __t9 *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]
{
var __t_tag_0 *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__local_var_2)
if (__t_tag_0 == nil) {
__t9 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__local_var_3)
goto end_branch_9
} else {

}
}
{
var __t_tag_1 *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__local_var_3)
if (__t_tag_1 == nil) {
__t9 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__local_var_2)
goto end_branch_9
} else {

}
}
{
var __t_tag_2 *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__local_var_3)
if (__t_tag_2 != nil) {
// TAST (Let): v_4_3 shape=UncurriedApp(Var) bindingType=(ADT ["Data","Map","Internal","Split"] [(TypeVar k), (TypeVar v)])
v_4_3 := gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Split[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp3(Get_Data_Map_Internal_unsafeSplit(), comp_0, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V2, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__local_var_2))}))
_ = v_4_3
// TAST (Let): l_prime__5_4 shape=UncurriedApp(Var) bindingType=Any
l_prime__5_4 := gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeUnionWith(), comp_0, __local_var_1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((v_4_3).V1)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V4)})
_ = l_prime__5_4
// TAST (Let): r_prime__6_5 shape=UncurriedApp(Var) bindingType=Any
r_prime__6_5 := gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeUnionWith(), comp_0, __local_var_1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((v_4_3).V2)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V5)})
_ = r_prime__6_5
var __t8 *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]
{
var __t_tag_6 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = (v_4_3).V0
if (__t_tag_6 != nil) {
__t8 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeBalancedNode(), (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V2, gopurs_runtime.Apply2(__local_var_1, ((v_4_3).V0).V0, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V3), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](l_prime__5_4))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](r_prime__6_5))}))
goto end_branch_8
} else {

}
}
{
var __t_tag_7 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = (v_4_3).V0
if (__t_tag_7 == nil) {
__t8 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeBalancedNode(), (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](l_prime__5_4))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](r_prime__6_5))}))
goto end_branch_8
} else {

}
}
{
__t8 = func() *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value] { panic("Failed pattern match") }()
}
end_branch_8:
__t9 = __t8
goto end_branch_9
} else {

}
}
{
__t9 = func() *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value] { panic("Failed pattern match") }()
}
end_branch_9:
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(__t9)}
}
}

func Call_Data_Map_Internal_unionWith(dictOrd_0_loop *Constructor_Data_Ord_Ord[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictOrd_0 *Constructor_Data_Ord_Ord[gopurs_runtime.Value] = dictOrd_0_loop
_ = dictOrd_0
// TAST (Let): compare_1_0 shape=Other bindingType=(Func [(TypeVar k), (TypeVar k)] (ADT ["Data","Ordering","Ordering"] []))
compare_1_0 := gopurs_runtime.Box(dictOrd_0.V1)
_ = compare_1_0
return gopurs_runtime.Func3(func(app_2 gopurs_runtime.Value, m1_3 gopurs_runtime.Value, m2_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeUnionWith(), compare_1_0, app_2, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](m1_3))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](m2_4))})))}
})
}

func Call_Data_Map_Internal_unionWith__2195946919(__eta_norm_2_unused_0_loop gopurs_runtime.Value, __eta_norm_1_1_loop gopurs_runtime.Value, __eta_norm_0_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
unionWith__2195946919:
for {
if false { continue unionWith__2195946919 }
var __eta_norm_2_unused_0 gopurs_runtime.Value = __eta_norm_2_unused_0_loop
_ = __eta_norm_2_unused_0
var __eta_norm_1_1 gopurs_runtime.Value = __eta_norm_1_1_loop
_ = __eta_norm_1_1
var __eta_norm_0_2 gopurs_runtime.Value = __eta_norm_0_2_loop
_ = __eta_norm_0_2
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(Rebox_Data_Map_Internal_261879545_2487766124(Rebox_Data_Map_Internal_2487766124_261879545(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeUnionWith(), gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[uint32]](Get_Data_Interval_Duration_ordDurationComponent()).V1), Get_Data_Semiring_numAdd(), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(Rebox_Data_Map_Internal_261879545_2487766124(Rebox_Data_Map_Internal_2487766124_261879545(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__eta_norm_1_1))))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(Rebox_Data_Map_Internal_261879545_2487766124(Rebox_Data_Map_Internal_2487766124_261879545(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__eta_norm_0_2))))})))))}
}
}

func Call_Data_Map_Internal_union(dictOrd_0_loop *Constructor_Data_Ord_Ord[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictOrd_0 *Constructor_Data_Ord_Ord[gopurs_runtime.Value] = dictOrd_0_loop
_ = dictOrd_0
// TAST (Let): compare_1_0 shape=Other bindingType=(Func [(TypeVar k), (TypeVar k)] (ADT ["Data","Ordering","Ordering"] []))
compare_1_0 := gopurs_runtime.Box(dictOrd_0.V1)
_ = compare_1_0
return gopurs_runtime.Func2(func(m1_2 gopurs_runtime.Value, m2_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeUnionWith(), compare_1_0, Get_Data_Function_go__const(), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](m1_2))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](m2_3))})))}
})
}

func Call_Data_Map_Internal_update(dictOrd_0_loop *Constructor_Data_Ord_Ord[gopurs_runtime.Value], f_1_loop gopurs_runtime.Value, k_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 *Constructor_Data_Ord_Ord[gopurs_runtime.Value] = dictOrd_0_loop
_ = dictOrd_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var k_2 gopurs_runtime.Value = k_2_loop
_ = k_2
var go__go_3_0_1 gopurs_runtime.Value
_ = go__go_3_0_1
// FALLBACK TCO: isLoop=false len=1
go__go_3_0_1 = gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t7 *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]
{
var __t_tag_1 *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](v_4)
if (__t_tag_1 == nil) {
__t7 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(nil))})
goto end_branch_7
} else {

}
}
{
var __t_tag_2 *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](v_4)
if (__t_tag_2 != nil) {
// TAST (Let): v1_5_3 shape=App(Other) bindingType=(ADT ["Data","Ordering","Ordering"] [])
v1_5_3 := uint32(gopurs_runtime.Apply2(gopurs_runtime.Box(dictOrd_0.V1), k_2, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V2).IntVal)
_ = v1_5_3
var __t6 *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]
{
if (v1_5_3 == 1527465420) {
__t6 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeBalancedNode(), (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(go__go_3_0_1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V4)})))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V5)}))
goto end_branch_6
} else {

}
}
{
if (v1_5_3 == 380165415) {
__t6 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeBalancedNode(), (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V4)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(go__go_3_0_1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V5)})))}))
goto end_branch_6
} else {

}
}
{
if (v1_5_3 == 902936544) {
// TAST (Let): v2_6_4 shape=App(Other) bindingType=(ADT ["Data","Maybe","Maybe"] [(TypeVar v)])
v2_6_4 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(f_1, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V3))
_ = v2_6_4
var __t5 *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]
{
if (v2_6_4 == nil) {
__t5 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp2(Get_Data_Map_Internal_unsafeJoinNodes(), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V4)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V5)}))
goto end_branch_5
} else {

}
}
{
if (v2_6_4 != nil) {
__t5 = (&Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V0, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V1, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V2, (v2_6_4).V0, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V4, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V5})
goto end_branch_5
} else {

}
}
{
__t5 = func() *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value] { panic("Failed pattern match") }()
}
end_branch_5:
__t6 = __t5
goto end_branch_6
} else {

}
}
{
__t6 = func() *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value] { panic("Failed pattern match") }()
}
end_branch_6:
__t7 = __t6
goto end_branch_7
} else {

}
}
{
__t7 = func() *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value] { panic("Failed pattern match") }()
}
end_branch_7:
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(__t7)}
})
return go__go_3_0_1
}

func Call_Data_Map_Internal_showTree(dictShow_0_loop *Constructor_Data_Show_Show[gopurs_runtime.Value], dictShow1_1_loop *Constructor_Data_Show_Show[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictShow_0 *Constructor_Data_Show_Show[gopurs_runtime.Value] = dictShow_0_loop
_ = dictShow_0
var dictShow1_1 *Constructor_Data_Show_Show[gopurs_runtime.Value] = dictShow1_1_loop
_ = dictShow1_1
var go__1866553730_2_0_2 gopurs_runtime.Value
_ = go__1866553730_2_0_2
// FALLBACK TCO: isLoop=false len=1
go__1866553730_2_0_2 = gopurs_runtime.Func2(func(ind_3 gopurs_runtime.Value, v_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t3 string
{
var __t_tag_1 *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](v_4)
if (__t_tag_1 == nil) {
__t3 = (ind_3.StrVal()) + ("Leaf")
goto end_branch_3
} else {

}
}
{
var __t_tag_2 *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](v_4)
if (__t_tag_2 != nil) {
__t3 = ((((((((((ind_3.StrVal()) + ("[")) + (gopurs_runtime.Apply(Get_Data_Show_showIntImpl(), gopurs_runtime.Int((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V0)).StrVal())) + ("] ")) + (gopurs_runtime.Apply(gopurs_runtime.Box(dictShow_0.V0), (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V2).StrVal())) + (" => ")) + (gopurs_runtime.Apply(gopurs_runtime.Box(dictShow1_1.V0), (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V3).StrVal())) + ("\x0a")) + (gopurs_runtime.Apply2(go__1866553730_2_0_2, gopurs_runtime.Str((ind_3.StrVal()) + ("    ")), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V4)}).StrVal())) + ("\x0a")) + (gopurs_runtime.Apply2(go__1866553730_2_0_2, gopurs_runtime.Str((ind_3.StrVal()) + ("    ")), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V5)}).StrVal())
goto end_branch_3
} else {

}
}
{
__t3 = func() string { panic("Failed pattern match") }()
}
end_branch_3:
return gopurs_runtime.Str(__t3)
})
var go__go_3_4_3 gopurs_runtime.Value
_ = go__go_3_4_3
// FALLBACK TCO: isLoop=false len=1
go__go_3_4_3 = gopurs_runtime.Func2(func(ind_4 gopurs_runtime.Value, v_5 gopurs_runtime.Value) gopurs_runtime.Value {
var __t7 string
{
var __t_tag_5 *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](v_5)
if (__t_tag_5 == nil) {
__t7 = (ind_4.StrVal()) + ("Leaf")
goto end_branch_7
} else {

}
}
{
var __t_tag_6 *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](v_5)
if (__t_tag_6 != nil) {
__t7 = ((((((((((ind_4.StrVal()) + ("[")) + (gopurs_runtime.Apply(Get_Data_Show_showIntImpl(), gopurs_runtime.Int((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_5.UnsafePtr).V0)).StrVal())) + ("] ")) + (gopurs_runtime.Apply(gopurs_runtime.Box(dictShow_0.V0), (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_5.UnsafePtr).V2).StrVal())) + (" => ")) + (gopurs_runtime.Apply(gopurs_runtime.Box(dictShow1_1.V0), (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_5.UnsafePtr).V3).StrVal())) + ("\x0a")) + (gopurs_runtime.Apply2(go__1866553730_2_0_2, gopurs_runtime.Str((ind_4.StrVal()) + ("    ")), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_5.UnsafePtr).V4)}).StrVal())) + ("\x0a")) + (gopurs_runtime.Apply2(go__1866553730_2_0_2, gopurs_runtime.Str((ind_4.StrVal()) + ("    ")), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_5.UnsafePtr).V5)}).StrVal())
goto end_branch_7
} else {

}
}
{
__t7 = func() string { panic("Failed pattern match") }()
}
end_branch_7:
return gopurs_runtime.Str(__t7)
})
return gopurs_runtime.Apply(go__1866553730_2_0_2, gopurs_runtime.Str(""))
}

func Call_Data_Map_Internal_semigroupMap(_dollar___unused_0_loop gopurs_runtime.Value, dictOrd_1_loop gopurs_runtime.Value, dictSemigroup_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var _dollar___unused_0 gopurs_runtime.Value = _dollar___unused_0_loop
_ = _dollar___unused_0
var dictOrd_1 gopurs_runtime.Value = dictOrd_1_loop
_ = dictOrd_1
var dictSemigroup_2 gopurs_runtime.Value = dictSemigroup_2_loop
_ = dictSemigroup_2
// TAST (Let): compare_3_0 shape=Other bindingType=(Func [(TypeVar k), (TypeVar k)] (ADT ["Data","Ordering","Ordering"] []))
compare_3_0 := gopurs_runtime.RecordGet(dictOrd_1, "compare")
_ = compare_3_0
// TAST (Let): __local_var_4_1 shape=Other bindingType=(Func [(TypeVar v), (TypeVar v)] (TypeVar v))
__local_var_4_1 := gopurs_runtime.RecordGet(dictSemigroup_2, "append")
_ = __local_var_4_1
return gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer(Rebox_Data_Map_Internal_3854229351_4179793454((&Constructor_Data_Semigroup_Semigroup[*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]]{1, gopurs_runtime.Func2(func(m1_5 gopurs_runtime.Value, m2_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeUnionWith(), compare_3_0, __local_var_4_1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](m1_5))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](m2_6))})))}
})})))}
}

func Call_Data_Map_Internal_semigroupMap1(dictOrd_0_loop gopurs_runtime.Value, dictSemigroup_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
var dictSemigroup_1 gopurs_runtime.Value = dictSemigroup_1_loop
_ = dictSemigroup_1
// TAST (Let): compare_2_0 shape=Other bindingType=(Func [(TypeVar k), (TypeVar k)] (ADT ["Data","Ordering","Ordering"] []))
compare_2_0 := gopurs_runtime.RecordGet(dictOrd_0, "compare")
_ = compare_2_0
// TAST (Let): __local_var_3_1 shape=Other bindingType=(Func [(TypeVar v), (TypeVar v)] (TypeVar v))
__local_var_3_1 := gopurs_runtime.RecordGet(dictSemigroup_1, "append")
_ = __local_var_3_1
return gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer(Rebox_Data_Map_Internal_3854229351_4179793454((&Constructor_Data_Semigroup_Semigroup[*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]]{1, gopurs_runtime.Func2(func(m1_4 gopurs_runtime.Value, m2_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeUnionWith(), compare_2_0, __local_var_3_1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](m1_4))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](m2_5))})))}
})})))}
}

func Call_Data_Map_Internal_pop(dictOrd_0_loop *Constructor_Data_Ord_Ord[gopurs_runtime.Value]) struct{V0 gopurs_runtime.Value; V1 bool} {
var dictOrd_0 *Constructor_Data_Ord_Ord[gopurs_runtime.Value] = dictOrd_0_loop
_ = dictOrd_0
// TAST (Let): compare_1_0 shape=Other bindingType=(Func [(TypeVar k), (TypeVar k)] (ADT ["Data","Ordering","Ordering"] []))
compare_1_0 := gopurs_runtime.Box(dictOrd_0.V1)
_ = compare_1_0
return func() struct{V0 gopurs_runtime.Value; V1 bool} {
				_v := gopurs_runtime.Func2(func(k_2 gopurs_runtime.Value, m_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): v_4_1 shape=UncurriedApp(Var) bindingType=(ADT ["Data","Map","Internal","Split"] [(TypeVar k), (TypeVar v)])
v_4_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Split[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp3(Get_Data_Map_Internal_unsafeSplit(), compare_1_0, k_2, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](m_3))}))
_ = v_4_1
var __t3 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
var __t_tag_2 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = (v_4_1).V0
if (__t_tag_2 != nil) {
__t3 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Rebox_Data_Map_Internal_3113727329_138441832(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value}{((v_4_1).V0).V0, gopurs_runtime.UncurriedApp2(Get_Data_Map_Internal_unsafeJoinNodes(), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((v_4_1).V1)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((v_4_1).V2)})}
				return gopurs_runtime.Box(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})
			}())))}, true}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())
goto end_branch_3
} else {

}
}
{
__t3 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())
}
end_branch_3:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t3)}
})
				if _v.Type == 9 && _v.IntVal == 3562950408 && _v.UnsafePtr != nil {
					return struct{V0 gopurs_runtime.Value; V1 bool}{V0: (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr).V0, V1: true}
				}
				return struct{V0 gopurs_runtime.Value; V1 bool}{V0: gopurs_runtime.Value{}, V1: false}
			}()
}

func Call_Data_Map_Internal_member(dictOrd_0_loop *Constructor_Data_Ord_Ord[gopurs_runtime.Value], k_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 *Constructor_Data_Ord_Ord[gopurs_runtime.Value] = dictOrd_0_loop
_ = dictOrd_0
var k_1 gopurs_runtime.Value = k_1_loop
_ = k_1
var Call_local_Data_Map_Internal_go__go_2_0_4 func(*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]) bool
_ = Call_local_Data_Map_Internal_go__go_2_0_4
var go__go_2_0_4 gopurs_runtime.Value
_ = go__go_2_0_4
Call_local_Data_Map_Internal_go__go_2_0_4 = func(v_3_loop *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]) bool {
go__go_2_0_4:
for {
if false { continue go__go_2_0_4 }
var v_3 *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value] = v_3_loop
_ = v_3
var __t3 bool
{
if (v_3 == nil) {
__t3 = false
goto end_branch_3
} else {

}
}
{
if (v_3 != nil) {
// TAST (Let): v1_4_1 shape=App(Other) bindingType=(ADT ["Data","Ordering","Ordering"] [])
v1_4_1 := uint32(gopurs_runtime.Apply2(gopurs_runtime.Box(dictOrd_0.V1), k_1, (v_3).V2).IntVal)
_ = v1_4_1
var __t2 bool
{
if (v1_4_1 == 1527465420) {
v_3_loop = (v_3).V4
continue go__go_2_0_4
__t2 = func() bool { panic("unreachable") }()
goto end_branch_2
} else {

}
}
{
if (v1_4_1 == 380165415) {
v_3_loop = (v_3).V5
continue go__go_2_0_4
__t2 = func() bool { panic("unreachable") }()
goto end_branch_2
} else {

}
}
{
if (v1_4_1 == 902936544) {
__t2 = true
goto end_branch_2
} else {

}
}
{
__t2 = func() bool { panic("Failed pattern match") }()
}
end_branch_2:
__t3 = __t2
goto end_branch_3
} else {

}
}
{
__t3 = func() bool { panic("Failed pattern match") }()
}
end_branch_3:
return __t3
}
}
go__go_2_0_4 = gopurs_runtime.Func(func(v_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_local_Data_Map_Internal_go__go_2_0_4(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](v_3_loop_val)))
})
return go__go_2_0_4
}

func Call_Data_Map_Internal_mapMaybeWithKey(dictOrd_0_loop *Constructor_Data_Ord_Ord[gopurs_runtime.Value], f_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 *Constructor_Data_Ord_Ord[gopurs_runtime.Value] = dictOrd_0_loop
_ = dictOrd_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var go__go_2_0_5 gopurs_runtime.Value
_ = go__go_2_0_5
// FALLBACK TCO: isLoop=false len=1
go__go_2_0_5 = gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t5 *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]
{
var __t_tag_1 *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](v_3)
if (__t_tag_1 == nil) {
__t5 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(nil))})
goto end_branch_5
} else {

}
}
{
var __t_tag_2 *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](v_3)
if (__t_tag_2 != nil) {
// TAST (Let): v2_4_3 shape=App(Other) bindingType=(ADT ["Data","Maybe","Maybe"] [(TypeVar b)])
v2_4_3 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply2(f_1, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V3))
_ = v2_4_3
var __t4 *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]
{
if (v2_4_3 != nil) {
__t4 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeBalancedNode(), (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V2, (v2_4_3).V0, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(go__go_2_0_5, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V4)})))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(go__go_2_0_5, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V5)})))}))
goto end_branch_4
} else {

}
}
{
if (v2_4_3 == nil) {
__t4 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp2(Get_Data_Map_Internal_unsafeJoinNodes(), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(go__go_2_0_5, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V4)})))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(go__go_2_0_5, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V5)})))}))
goto end_branch_4
} else {

}
}
{
__t4 = func() *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value] { panic("Failed pattern match") }()
}
end_branch_4:
__t5 = __t4
goto end_branch_5
} else {

}
}
{
__t5 = func() *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value] { panic("Failed pattern match") }()
}
end_branch_5:
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(__t5)}
})
return go__go_2_0_5
}

func Call_Data_Map_Internal_mapMaybe(dictOrd_0_loop *Constructor_Data_Ord_Ord[gopurs_runtime.Value], x_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 *Constructor_Data_Ord_Ord[gopurs_runtime.Value] = dictOrd_0_loop
_ = dictOrd_0
var x_1 gopurs_runtime.Value = x_1_loop
_ = x_1
return Call_Data_Map_Internal_mapMaybeWithKey(dictOrd_0, gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return x_1
}))
}

func Call_Data_Map_Internal_lookupLE(dictOrd_0_loop *Constructor_Data_Ord_Ord[gopurs_runtime.Value], k_1_loop gopurs_runtime.Value) struct{V0 gopurs_runtime.Value; V1 bool} {
var dictOrd_0 *Constructor_Data_Ord_Ord[gopurs_runtime.Value] = dictOrd_0_loop
_ = dictOrd_0
var k_1 gopurs_runtime.Value = k_1_loop
_ = k_1
var go__go_2_0_6 gopurs_runtime.Value
_ = go__go_2_0_6
// FALLBACK TCO: isLoop=false len=1
go__go_2_0_6 = gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t7 *Constructor_Data_Maybe_Just[struct{
	key gopurs_runtime.Value
	value gopurs_runtime.Value
}]
{
var __t_tag_1 *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](v_3)
if (__t_tag_1 == nil) {
__t7 = Rebox_Data_Map_Internal_3094389156_2758605161(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil))}))
goto end_branch_7
} else {

}
}
{
var __t_tag_2 *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](v_3)
if (__t_tag_2 != nil) {
// TAST (Let): v1_4_3 shape=App(Other) bindingType=(ADT ["Data","Ordering","Ordering"] [])
v1_4_3 := uint32(gopurs_runtime.Apply2(gopurs_runtime.Box(dictOrd_0.V1), k_1, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V2).IntVal)
_ = v1_4_3
var __t6 *Constructor_Data_Maybe_Just[struct{
	key gopurs_runtime.Value
	value gopurs_runtime.Value
}]
{
if (v1_4_3 == 1527465420) {
__t6 = Rebox_Data_Map_Internal_3094389156_2758605161(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(go__go_2_0_6, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V4)})))
goto end_branch_6
} else {

}
}
{
if (v1_4_3 == 380165415) {
// TAST (Let): v2_5_4 shape=App(Other) bindingType=(ADT ["Data","Maybe","Maybe"] [(Record (Row [key: (TypeVar k), value: (TypeVar v)] Any))])
v2_5_4 := Rebox_Data_Map_Internal_3094389156_2758605161(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(go__go_2_0_6, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V5)})))
_ = v2_5_4
var __t5 *Constructor_Data_Maybe_Just[struct{
	key gopurs_runtime.Value
	value gopurs_runtime.Value
}]
{
if (v2_5_4 == nil) {
__t5 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[struct{
	key gopurs_runtime.Value
	value gopurs_runtime.Value
}]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.RecordDict2("key", "value", (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V3), true}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())
goto end_branch_5
} else {

}
}
{
__t5 = v2_5_4
}
end_branch_5:
__t6 = __t5
goto end_branch_6
} else {

}
}
{
if (v1_4_3 == 902936544) {
__t6 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[struct{
	key gopurs_runtime.Value
	value gopurs_runtime.Value
}]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.RecordDict2("key", "value", (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V3), true}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())
goto end_branch_6
} else {

}
}
{
__t6 = func() *Constructor_Data_Maybe_Just[struct{
	key gopurs_runtime.Value
	value gopurs_runtime.Value
}] { panic("Failed pattern match") }()
}
end_branch_6:
__t7 = __t6
goto end_branch_7
} else {

}
}
{
__t7 = func() *Constructor_Data_Maybe_Just[struct{
	key gopurs_runtime.Value
	value gopurs_runtime.Value
}] { panic("Failed pattern match") }()
}
end_branch_7:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_Map_Internal_2758605161_3094389156(__t7))}
})
return func() struct{V0 gopurs_runtime.Value; V1 bool} {
				_v := go__go_2_0_6
				if _v.Type == 9 && _v.IntVal == 3562950408 && _v.UnsafePtr != nil {
					return struct{V0 gopurs_runtime.Value; V1 bool}{V0: (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr).V0, V1: true}
				}
				return struct{V0 gopurs_runtime.Value; V1 bool}{V0: gopurs_runtime.Value{}, V1: false}
			}()
}

func Call_Data_Map_Internal_lookupGE(dictOrd_0_loop *Constructor_Data_Ord_Ord[gopurs_runtime.Value], k_1_loop gopurs_runtime.Value) struct{V0 gopurs_runtime.Value; V1 bool} {
var dictOrd_0 *Constructor_Data_Ord_Ord[gopurs_runtime.Value] = dictOrd_0_loop
_ = dictOrd_0
var k_1 gopurs_runtime.Value = k_1_loop
_ = k_1
var go__go_2_0_7 gopurs_runtime.Value
_ = go__go_2_0_7
// FALLBACK TCO: isLoop=false len=1
go__go_2_0_7 = gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t7 *Constructor_Data_Maybe_Just[struct{
	key gopurs_runtime.Value
	value gopurs_runtime.Value
}]
{
var __t_tag_1 *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](v_3)
if (__t_tag_1 == nil) {
__t7 = Rebox_Data_Map_Internal_3094389156_2758605161(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil))}))
goto end_branch_7
} else {

}
}
{
var __t_tag_2 *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](v_3)
if (__t_tag_2 != nil) {
// TAST (Let): v1_4_3 shape=App(Other) bindingType=(ADT ["Data","Ordering","Ordering"] [])
v1_4_3 := uint32(gopurs_runtime.Apply2(gopurs_runtime.Box(dictOrd_0.V1), k_1, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V2).IntVal)
_ = v1_4_3
var __t6 *Constructor_Data_Maybe_Just[struct{
	key gopurs_runtime.Value
	value gopurs_runtime.Value
}]
{
if (v1_4_3 == 1527465420) {
// TAST (Let): v2_5_4 shape=App(Other) bindingType=(ADT ["Data","Maybe","Maybe"] [(Record (Row [key: (TypeVar k), value: (TypeVar v)] Any))])
v2_5_4 := Rebox_Data_Map_Internal_3094389156_2758605161(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(go__go_2_0_7, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V4)})))
_ = v2_5_4
var __t5 *Constructor_Data_Maybe_Just[struct{
	key gopurs_runtime.Value
	value gopurs_runtime.Value
}]
{
if (v2_5_4 == nil) {
__t5 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[struct{
	key gopurs_runtime.Value
	value gopurs_runtime.Value
}]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.RecordDict2("key", "value", (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V3), true}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())
goto end_branch_5
} else {

}
}
{
__t5 = v2_5_4
}
end_branch_5:
__t6 = __t5
goto end_branch_6
} else {

}
}
{
if (v1_4_3 == 380165415) {
__t6 = Rebox_Data_Map_Internal_3094389156_2758605161(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(go__go_2_0_7, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V5)})))
goto end_branch_6
} else {

}
}
{
if (v1_4_3 == 902936544) {
__t6 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[struct{
	key gopurs_runtime.Value
	value gopurs_runtime.Value
}]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.RecordDict2("key", "value", (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V3), true}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())
goto end_branch_6
} else {

}
}
{
__t6 = func() *Constructor_Data_Maybe_Just[struct{
	key gopurs_runtime.Value
	value gopurs_runtime.Value
}] { panic("Failed pattern match") }()
}
end_branch_6:
__t7 = __t6
goto end_branch_7
} else {

}
}
{
__t7 = func() *Constructor_Data_Maybe_Just[struct{
	key gopurs_runtime.Value
	value gopurs_runtime.Value
}] { panic("Failed pattern match") }()
}
end_branch_7:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_Map_Internal_2758605161_3094389156(__t7))}
})
return func() struct{V0 gopurs_runtime.Value; V1 bool} {
				_v := go__go_2_0_7
				if _v.Type == 9 && _v.IntVal == 3562950408 && _v.UnsafePtr != nil {
					return struct{V0 gopurs_runtime.Value; V1 bool}{V0: (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr).V0, V1: true}
				}
				return struct{V0 gopurs_runtime.Value; V1 bool}{V0: gopurs_runtime.Value{}, V1: false}
			}()
}

func Call_Data_Map_Internal_lookup(dictOrd_0_loop *Constructor_Data_Ord_Ord[gopurs_runtime.Value], k_1_loop gopurs_runtime.Value) struct{V0 gopurs_runtime.Value; V1 bool} {
var dictOrd_0 *Constructor_Data_Ord_Ord[gopurs_runtime.Value] = dictOrd_0_loop
_ = dictOrd_0
var k_1 gopurs_runtime.Value = k_1_loop
_ = k_1
var Call_local_Data_Map_Internal_go__go_2_0_8 func(*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]) *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
_ = Call_local_Data_Map_Internal_go__go_2_0_8
var go__go_2_0_8 gopurs_runtime.Value
_ = go__go_2_0_8
Call_local_Data_Map_Internal_go__go_2_0_8 = func(v_3_loop *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]) *Constructor_Data_Maybe_Just[gopurs_runtime.Value] {
go__go_2_0_8:
for {
if false { continue go__go_2_0_8 }
var v_3 *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value] = v_3_loop
_ = v_3
var __t3 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
if (v_3 == nil) {
__t3 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil))})
goto end_branch_3
} else {

}
}
{
if (v_3 != nil) {
// TAST (Let): v1_4_1 shape=App(Other) bindingType=(ADT ["Data","Ordering","Ordering"] [])
v1_4_1 := uint32(gopurs_runtime.Apply2(gopurs_runtime.Box(dictOrd_0.V1), k_1, (v_3).V2).IntVal)
_ = v1_4_1
var __t2 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
if (v1_4_1 == 1527465420) {
v_3_loop = (v_3).V4
continue go__go_2_0_8
__t2 = func() *Constructor_Data_Maybe_Just[gopurs_runtime.Value] { panic("unreachable") }()
goto end_branch_2
} else {

}
}
{
if (v1_4_1 == 380165415) {
v_3_loop = (v_3).V5
continue go__go_2_0_8
__t2 = func() *Constructor_Data_Maybe_Just[gopurs_runtime.Value] { panic("unreachable") }()
goto end_branch_2
} else {

}
}
{
if (v1_4_1 == 902936544) {
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{(v_3).V3, true}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())
goto end_branch_2
} else {

}
}
{
__t2 = func() *Constructor_Data_Maybe_Just[gopurs_runtime.Value] { panic("Failed pattern match") }()
}
end_branch_2:
__t3 = __t2
goto end_branch_3
} else {

}
}
{
__t3 = func() *Constructor_Data_Maybe_Just[gopurs_runtime.Value] { panic("Failed pattern match") }()
}
end_branch_3:
return __t3
}
}
go__go_2_0_8 = gopurs_runtime.Func(func(v_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_local_Data_Map_Internal_go__go_2_0_8(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](v_3_loop_val)))}
})
return func() struct{V0 gopurs_runtime.Value; V1 bool} {
				_v := go__go_2_0_8
				if _v.Type == 9 && _v.IntVal == 3562950408 && _v.UnsafePtr != nil {
					return struct{V0 gopurs_runtime.Value; V1 bool}{V0: (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr).V0, V1: true}
				}
				return struct{V0 gopurs_runtime.Value; V1 bool}{V0: gopurs_runtime.Value{}, V1: false}
			}()
}

func Call_Data_Map_Internal_lookup__634903124(k_unused_0_loop uint32, __eta_norm_0_1_loop *Constructor_Data_Map_Internal_Node[uint32, float64]) struct{V0 gopurs_runtime.Value; V1 bool} {
lookup__634903124:
for {
if false { continue lookup__634903124 }
var k_unused_0 uint32 = k_unused_0_loop
_ = k_unused_0
var __eta_norm_0_1 *Constructor_Data_Map_Internal_Node[uint32, float64] = __eta_norm_0_1_loop
_ = __eta_norm_0_1
var Call_local_Data_Map_Internal_go__go_2_0_9 func(*Constructor_Data_Map_Internal_Node[uint32, float64]) *Constructor_Data_Maybe_Just[float64]
_ = Call_local_Data_Map_Internal_go__go_2_0_9
var go__go_2_0_9 gopurs_runtime.Value
_ = go__go_2_0_9
Call_local_Data_Map_Internal_go__go_2_0_9 = func(v_3_loop *Constructor_Data_Map_Internal_Node[uint32, float64]) *Constructor_Data_Maybe_Just[float64] {
go__go_2_0_9:
for {
if false { continue go__go_2_0_9 }
var v_3 *Constructor_Data_Map_Internal_Node[uint32, float64] = v_3_loop
_ = v_3
var __t7 *Constructor_Data_Maybe_Just[float64]
{
if (v_3 == nil) {
__t7 = Rebox_Data_Map_Internal_3094389156_3240988860(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil))}))
goto end_branch_7
} else {

}
}
{
if (v_3 != nil) {
var __t6 *Constructor_Data_Maybe_Just[float64]
{
var __t_tag_1 uint32 = (v_3).V2
if (uint32(__t_tag_1) == 3908053364) {
v_3_loop = (v_3).V5
continue go__go_2_0_9
__t6 = func() *Constructor_Data_Maybe_Just[float64] { panic("unreachable") }()
goto end_branch_6
} else {

}
}
{
var __t_tag_2 uint32 = (v_3).V2
if (uint32(__t_tag_2) == 217821258) {
v_3_loop = (v_3).V5
continue go__go_2_0_9
__t6 = func() *Constructor_Data_Maybe_Just[float64] { panic("unreachable") }()
goto end_branch_6
} else {

}
}
{
var __t_tag_3 uint32 = (v_3).V2
if (uint32(__t_tag_3) == 1292308612) {
v_3_loop = (v_3).V5
continue go__go_2_0_9
__t6 = func() *Constructor_Data_Maybe_Just[float64] { panic("unreachable") }()
goto end_branch_6
} else {

}
}
{
var __t_tag_4 uint32 = (v_3).V2
if (uint32(__t_tag_4) == 2311060696) {
v_3_loop = (v_3).V5
continue go__go_2_0_9
__t6 = func() *Constructor_Data_Maybe_Just[float64] { panic("unreachable") }()
goto end_branch_6
} else {

}
}
{
var __t_tag_5 uint32 = (v_3).V2
if (uint32(__t_tag_5) == 401302776) {
__t6 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[float64]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Float((v_3).V3), true}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())
goto end_branch_6
} else {

}
}
{
v_3_loop = (v_3).V4
continue go__go_2_0_9
__t6 = func() *Constructor_Data_Maybe_Just[float64] { panic("unreachable") }()
}
end_branch_6:
__t7 = __t6
goto end_branch_7
} else {

}
}
{
__t7 = func() *Constructor_Data_Maybe_Just[float64] { panic("Failed pattern match") }()
}
end_branch_7:
return __t7
}
}
go__go_2_0_9 = gopurs_runtime.Func(func(v_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_Map_Internal_3240988860_3094389156(Call_local_Data_Map_Internal_go__go_2_0_9(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[uint32, float64]](v_3_loop_val))))}
})
return func() struct{V0 gopurs_runtime.Value; V1 bool} {
				_v := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_Map_Internal_3240988860_3094389156(Call_local_Data_Map_Internal_go__go_2_0_9(__eta_norm_0_1)))}
				if _v.Type == 9 && _v.IntVal == 3562950408 && _v.UnsafePtr != nil {
					return struct{V0 gopurs_runtime.Value; V1 bool}{V0: (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr).V0, V1: true}
				}
				return struct{V0 gopurs_runtime.Value; V1 bool}{V0: gopurs_runtime.Value{}, V1: false}
			}()
}
}

func Call_Data_Map_Internal_iterMapU(iter_0_loop gopurs_runtime.Value, v_1_loop *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var iter_0 gopurs_runtime.Value = iter_0_loop
_ = iter_0
var v_1 *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value] = v_1_loop
_ = v_1
var __t6 gopurs_runtime.Value
{
if (v_1 == nil) {
__t6 = iter_0
goto end_branch_6
} else {

}
}
{
if (v_1 != nil) {
var __t5 *Constructor_Data_Map_Internal_IterEmit[gopurs_runtime.Value, gopurs_runtime.Value]
{
var __t_tag_2 *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value] = (v_1).V4
if (__t_tag_2 == nil) {
var __t4 *Constructor_Data_Map_Internal_IterEmit[gopurs_runtime.Value, gopurs_runtime.Value]
{
var __t_tag_3 *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value] = (v_1).V5
if (__t_tag_3 == nil) {
__t4 = (&Constructor_Data_Map_Internal_IterEmit[gopurs_runtime.Value, gopurs_runtime.Value]{1, (v_1).V2, (v_1).V3, iter_0})
goto end_branch_4
} else {

}
}
{
__t4 = (&Constructor_Data_Map_Internal_IterEmit[gopurs_runtime.Value, gopurs_runtime.Value]{1, (v_1).V2, (v_1).V3, gopurs_runtime.Value{Type: 9, IntVal: 2861335956, UnsafePtr: unsafe.Pointer((&Constructor_Data_Map_Internal_IterNode[gopurs_runtime.Value, gopurs_runtime.Value]{1, (v_1).V5, iter_0}))}})
}
end_branch_4:
__t5 = __t4
goto end_branch_5
} else {

}
}
{
var __t1 *Constructor_Data_Map_Internal_IterEmit[gopurs_runtime.Value, gopurs_runtime.Value]
{
var __t_tag_0 *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value] = (v_1).V5
if (__t_tag_0 == nil) {
__t1 = (&Constructor_Data_Map_Internal_IterEmit[gopurs_runtime.Value, gopurs_runtime.Value]{1, (v_1).V2, (v_1).V3, gopurs_runtime.Value{Type: 9, IntVal: 2861335956, UnsafePtr: unsafe.Pointer((&Constructor_Data_Map_Internal_IterNode[gopurs_runtime.Value, gopurs_runtime.Value]{1, (v_1).V4, iter_0}))}})
goto end_branch_1
} else {

}
}
{
__t1 = (&Constructor_Data_Map_Internal_IterEmit[gopurs_runtime.Value, gopurs_runtime.Value]{1, (v_1).V2, (v_1).V3, gopurs_runtime.Value{Type: 9, IntVal: 2861335956, UnsafePtr: unsafe.Pointer((&Constructor_Data_Map_Internal_IterNode[gopurs_runtime.Value, gopurs_runtime.Value]{1, (v_1).V4, gopurs_runtime.Value{Type: 9, IntVal: 2861335956, UnsafePtr: unsafe.Pointer((&Constructor_Data_Map_Internal_IterNode[gopurs_runtime.Value, gopurs_runtime.Value]{1, (v_1).V5, iter_0}))}}))}})
}
end_branch_1:
__t5 = __t1
}
end_branch_5:
__t6 = gopurs_runtime.Value{Type: 9, IntVal: 1343415489, UnsafePtr: unsafe.Pointer(__t5)}
goto end_branch_6
} else {

}
}
{
__t6 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_6:
return __t6
}

func Call_Data_Map_Internal_toUnfoldableUnordered(dictUnfoldable_0_loop *Constructor_Data_Unfoldable_Unfoldable[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictUnfoldable_0 *Constructor_Data_Unfoldable_Unfoldable[gopurs_runtime.Value] = dictUnfoldable_0_loop
_ = dictUnfoldable_0
// TAST (Let): __local_var_1_0 shape=App(Other) bindingType=Any
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.Box(dictUnfoldable_0.V1), Get_Data_Map_Internal_stepUnfoldrUnordered())
_ = __local_var_1_0
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_1_0, gopurs_runtime.Value{Type: 9, IntVal: 2861335956, UnsafePtr: unsafe.Pointer((&Constructor_Data_Map_Internal_IterNode[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](x_2), gopurs_runtime.Value{Type: 9, IntVal: 2509360378, UnsafePtr: unsafe.Pointer(nil)}}))})
})
}

func Call_Data_Map_Internal_eqMapIter(dictEq_0_loop gopurs_runtime.Value, dictEq1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictEq_0 gopurs_runtime.Value = dictEq_0_loop
_ = dictEq_0
var dictEq1_1 gopurs_runtime.Value = dictEq1_1_loop
_ = dictEq1_1
var go__go_2_0_12 gopurs_runtime.Value
_ = go__go_2_0_12
// FALLBACK TCO: isLoop=false len=1
go__go_2_0_12 = gopurs_runtime.Func2(func(a_3 gopurs_runtime.Value, b_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): v_5_1 shape=App(Var) bindingType=(ADT ["Data","Map","Internal","MapIterStep"] [(TypeVar k), (TypeVar v)])
v_5_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_IterNext[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Map_Internal_stepAsc(), a_3))
_ = v_5_1
var __t3 bool
{
if (v_5_1 != nil) {
// TAST (Let): v2_6_2 shape=App(Var) bindingType=(ADT ["Data","Map","Internal","MapIterStep"] [(TypeVar k), (TypeVar v)])
v2_6_2 := gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_IterNext[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Map_Internal_stepAsc(), b_4))
_ = v2_6_2
__t3 = ((v2_6_2 != nil)) && ((((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq_0, "eq"), (v_5_1).V0, (v2_6_2).V0).IntVal) != (0)) && ((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq1_1, "eq"), (v_5_1).V1, (v2_6_2).V1).IntVal) != (0))) && ((gopurs_runtime.Apply2(go__go_2_0_12, (v_5_1).V2, (v2_6_2).V2).IntVal) != (0)))
goto end_branch_3
} else {

}
}
{
if (v_5_1 == nil) {
__t3 = true
goto end_branch_3
} else {

}
}
{
__t3 = func() bool { panic("Failed pattern match") }()
}
end_branch_3:
return gopurs_runtime.Bool(__t3)
})
return gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer((&Constructor_Data_Eq_Eq[gopurs_runtime.Value]{1, go__go_2_0_12}))}
}

func Call_Data_Map_Internal_ordMapIter(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
// TAST (Let): __local_var_1_1 shape=App(Other) bindingType=Any
__local_var_1_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd_0, "Eq0"), gopurs_runtime.Value{})
_ = __local_var_1_1
// TAST (Let): eqMapIter1__193435443_1_0 shape=Let(Abs(LitRecord)) bindingType=Any
eqMapIter1__193435443_1_0 := gopurs_runtime.Func(func(dictEq1_2 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_3_2_13 gopurs_runtime.Value
_ = go__go_3_2_13
// FALLBACK TCO: isLoop=false len=1
go__go_3_2_13 = gopurs_runtime.Func2(func(a_4 gopurs_runtime.Value, b_5 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): v_6_3 shape=App(Var) bindingType=(ADT ["Data","Map","Internal","MapIterStep"] [(TypeVar k), (TypeVar v)])
v_6_3 := gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_IterNext[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Map_Internal_stepAsc(), a_4))
_ = v_6_3
var __t5 bool
{
if (v_6_3 != nil) {
// TAST (Let): v2_7_4 shape=App(Var) bindingType=(ADT ["Data","Map","Internal","MapIterStep"] [(TypeVar k), (TypeVar v)])
v2_7_4 := gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_IterNext[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Map_Internal_stepAsc(), b_5))
_ = v2_7_4
__t5 = ((v2_7_4 != nil)) && ((((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_1, "eq"), (v_6_3).V0, (v2_7_4).V0).IntVal) != (0)) && ((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq1_2, "eq"), (v_6_3).V1, (v2_7_4).V1).IntVal) != (0))) && ((gopurs_runtime.Apply2(go__go_3_2_13, (v_6_3).V2, (v2_7_4).V2).IntVal) != (0)))
goto end_branch_5
} else {

}
}
{
if (v_6_3 == nil) {
__t5 = true
goto end_branch_5
} else {

}
}
{
__t5 = func() bool { panic("Failed pattern match") }()
}
end_branch_5:
return gopurs_runtime.Bool(__t5)
})
return gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer((&Constructor_Data_Eq_Eq[gopurs_runtime.Value]{1, go__go_3_2_13}))}
})
_ = eqMapIter1__193435443_1_0
return gopurs_runtime.Func(func(dictOrd1_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): eqMapIter2_3_6 shape=App(Other) bindingType=(ADT ["Data","Eq","Eq"] [(ADT ["Data","Map","Internal","MapIter"] [(TypeVar k), (TypeVar v)])])
eqMapIter2_3_6 := gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](gopurs_runtime.Apply(eqMapIter1__193435443_1_0, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd1_2, "Eq0"), gopurs_runtime.Value{})))
_ = eqMapIter2_3_6
var Call_local_Data_Map_Internal_go__go_4_7_14 func(gopurs_runtime.Value, gopurs_runtime.Value) uint32
_ = Call_local_Data_Map_Internal_go__go_4_7_14
var go__go_4_7_14 gopurs_runtime.Value
_ = go__go_4_7_14
Call_local_Data_Map_Internal_go__go_4_7_14 = func(a_5_loop gopurs_runtime.Value, b_6_loop gopurs_runtime.Value) uint32 {
go__go_4_7_14:
for {
if false { continue go__go_4_7_14 }
var a_5 gopurs_runtime.Value = a_5_loop
_ = a_5
var b_6 gopurs_runtime.Value = b_6_loop
_ = b_6
// TAST (Let): v_7_8 shape=App(Var) bindingType=(ADT ["Data","Map","Internal","MapIterStep"] [(TypeVar k), (TypeVar v)])
v_7_8 := gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_IterNext[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Map_Internal_stepAsc(), b_6))
_ = v_7_8
// TAST (Let): v1_8_9 shape=App(Var) bindingType=(ADT ["Data","Map","Internal","MapIterStep"] [(TypeVar k), (TypeVar v)])
v1_8_9 := gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_IterNext[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Map_Internal_stepAsc(), a_5))
_ = v1_8_9
var __t16 uint32
{
if (v1_8_9 != nil) {
var __t14 uint32
{
if (v_7_8 != nil) {
// TAST (Let): v3_9_10 shape=App(Other) bindingType=(ADT ["Data","Ordering","Ordering"] [])
v3_9_10 := uint32(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictOrd_0, "compare"), (v1_8_9).V0, (v_7_8).V0).IntVal)
_ = v3_9_10
var __t13 uint32
{
if (v3_9_10 == 902936544) {
// TAST (Let): v4_10_11 shape=App(Other) bindingType=(ADT ["Data","Ordering","Ordering"] [])
v4_10_11 := uint32(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictOrd1_2, "compare"), (v1_8_9).V1, (v_7_8).V1).IntVal)
_ = v4_10_11
var __t12 uint32
{
if (v4_10_11 == 902936544) {
a_5_loop = (v1_8_9).V2
b_6_loop = (v_7_8).V2
continue go__go_4_7_14
__t12 = func() uint32 { panic("unreachable") }()
goto end_branch_12
} else {

}
}
{
__t12 = v4_10_11
}
end_branch_12:
__t13 = __t12
goto end_branch_13
} else {

}
}
{
__t13 = v3_9_10
}
end_branch_13:
__t14 = __t13
goto end_branch_14
} else {

}
}
{
if (v_7_8 == nil) {
__t14 = 380165415
goto end_branch_14
} else {

}
}
{
__t14 = func() uint32 { panic("Failed pattern match") }()
}
end_branch_14:
__t16 = __t14
goto end_branch_16
} else {

}
}
{
if (v1_8_9 == nil) {
var __t15 uint32
{
if (v_7_8 == nil) {
__t15 = 902936544
goto end_branch_15
} else {

}
}
{
__t15 = 1527465420
}
end_branch_15:
__t16 = __t15
goto end_branch_16
} else {

}
}
{
if (v_7_8 == nil) {
__t16 = 380165415
goto end_branch_16
} else {

}
}
{
__t16 = func() uint32 { panic("Failed pattern match") }()
}
end_branch_16:
return __t16
}
}
go__go_4_7_14 = gopurs_runtime.Func(func(a_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: int64(Call_local_Data_Map_Internal_go__go_4_7_14(a_5_loop_val, b_6_loop_val)), UnsafePtr: nil}
})
})
return gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer((&Constructor_Data_Ord_Ord[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(eqMapIter2_3_6)}
}), go__go_4_7_14}))}
})
}

func Call_Data_Map_Internal_toUnfoldable(dictUnfoldable_0_loop *Constructor_Data_Unfoldable_Unfoldable[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictUnfoldable_0 *Constructor_Data_Unfoldable_Unfoldable[gopurs_runtime.Value] = dictUnfoldable_0_loop
_ = dictUnfoldable_0
// TAST (Let): __local_var_1_0 shape=App(Other) bindingType=Any
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.Box(dictUnfoldable_0.V1), Get_Data_Map_Internal_stepUnfoldr())
_ = __local_var_1_0
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_1_0, gopurs_runtime.Value{Type: 9, IntVal: 2861335956, UnsafePtr: unsafe.Pointer((&Constructor_Data_Map_Internal_IterNode[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](x_2), gopurs_runtime.Value{Type: 9, IntVal: 2509360378, UnsafePtr: unsafe.Pointer(nil)}}))})
})
}

func Call_Data_Map_Internal_toUnfoldable__994050345(__eta_norm_0_0_loop *Constructor_Data_Map_Internal_Node[uint32, float64]) *Constructor_Data_List_Types_Cons[*Constructor_Data_Tuple_Tuple[uint32, float64]] {
toUnfoldable__994050345:
for {
if false { continue toUnfoldable__994050345 }
var __eta_norm_0_0 *Constructor_Data_Map_Internal_Node[uint32, float64] = __eta_norm_0_0_loop
_ = __eta_norm_0_0
var Call_local_Data_Map_Internal_go__go_1_0_15 func(gopurs_runtime.Value, *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
_ = Call_local_Data_Map_Internal_go__go_1_0_15
var go__go_1_0_15 gopurs_runtime.Value
_ = go__go_1_0_15
Call_local_Data_Map_Internal_go__go_1_0_15 = func(source_2_loop gopurs_runtime.Value, memo_3_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] {
go__go_1_0_15:
for {
if false { continue go__go_1_0_15 }
var source_2 gopurs_runtime.Value = source_2_loop
_ = source_2
var memo_3 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = memo_3_loop
_ = memo_3
// TAST (Let): v_4_1 shape=App(Var) bindingType=(ADT ["Data","Maybe","Maybe"] [(ADT ["Data","Tuple","Tuple"] [(TypeVar a), (TypeVar b)])])
v_4_1 := Rebox_Data_Map_Internal_3094389156_4010058633(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Map_Internal_stepUnfoldr(), source_2)))
_ = v_4_1
var __t4 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
{
if (v_4_1 == nil) {
var Call_local_Data_Map_Internal_go__go_5_2_16 func(gopurs_runtime.Value, *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) gopurs_runtime.Value
_ = Call_local_Data_Map_Internal_go__go_5_2_16
var go__go_5_2_16 gopurs_runtime.Value
_ = go__go_5_2_16
Call_local_Data_Map_Internal_go__go_5_2_16 = func(b_6_loop gopurs_runtime.Value, v_7_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) gopurs_runtime.Value {
go__go_5_2_16:
for {
if false { continue go__go_5_2_16 }
var b_6 gopurs_runtime.Value = b_6_loop
_ = b_6
var v_7 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v_7_loop
_ = v_7
var __t3 gopurs_runtime.Value
{
if (v_7 == nil) {
__t3 = b_6
goto end_branch_3
} else {

}
}
{
if (v_7 != nil) {
b_6_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, (v_7).V0, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](b_6)}))}
v_7_loop = (v_7).V1
continue go__go_5_2_16
__t3 = func() gopurs_runtime.Value { panic("unreachable") }()
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
}
go__go_5_2_16 = gopurs_runtime.Func(func(b_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_7_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return Call_local_Data_Map_Internal_go__go_5_2_16(b_6_loop_val, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v_7_loop_val))
})
})
__t4 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](Call_local_Data_Map_Internal_go__go_5_2_16(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil))}, memo_3))
goto end_branch_4
} else {

}
}
{
if (v_4_1 != nil) {
source_2_loop = ((v_4_1).V0).V1
memo_3_loop = (&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, ((v_4_1).V0).V0, memo_3})
continue go__go_1_0_15
__t4 = func() *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] { panic("unreachable") }()
goto end_branch_4
} else {

}
}
{
__t4 = func() *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] { panic("Failed pattern match") }()
}
end_branch_4:
return __t4
}
}
go__go_1_0_15 = gopurs_runtime.Func(func(source_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(memo_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_local_Data_Map_Internal_go__go_1_0_15(source_2_loop_val, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](memo_3_loop_val)))}
})
})
return Rebox_Data_Map_Internal_849153993_2442833393(Call_local_Data_Map_Internal_go__go_1_0_15(gopurs_runtime.Value{Type: 9, IntVal: 2861335956, UnsafePtr: unsafe.Pointer((&Constructor_Data_Map_Internal_IterNode[gopurs_runtime.Value, gopurs_runtime.Value]{1, Rebox_Data_Map_Internal_261879545_2487766124(__eta_norm_0_0), gopurs_runtime.Value{Type: 9, IntVal: 2509360378, UnsafePtr: unsafe.Pointer(nil)}}))}, (*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil)))
}
}

func Call_Data_Map_Internal_showMap(dictShow_0_loop gopurs_runtime.Value, dictShow1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictShow_0 gopurs_runtime.Value = dictShow_0_loop
_ = dictShow_0
var dictShow1_1 gopurs_runtime.Value = dictShow1_1_loop
_ = dictShow1_1
// TAST (Let): showArray_2_0 shape=LitRecord bindingType=(ADT ["Data","Show","Show"] [(Array (ADT ["Data","Tuple","Tuple"] [(TypeVar k), (TypeVar v)]))])
showArray_2_0 := (&Constructor_Data_Show_Show[[]*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]{1, gopurs_runtime.Apply(Get_Data_Show_showArrayImpl(), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str((((("(Tuple ") + (gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictShow_0, "show"), (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V0).StrVal())) + (" ")) + (gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictShow1_1, "show"), (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V1).StrVal())) + (")"))
}))})
_ = showArray_2_0
return gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Data_Map_Internal_2638796135_1386611502((&Constructor_Data_Show_Show[*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(as_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str((("(fromFoldable ") + (gopurs_runtime.Apply(gopurs_runtime.Box(showArray_2_0.V0), func() gopurs_runtime.Value {
					arr := func() []*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value] {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Unfoldable_Unfoldable[gopurs_runtime.Value]](Get_Data_Unfoldable_unfoldableArray()).V1), Get_Data_Map_Internal_stepUnfoldr(), gopurs_runtime.Value{Type: 9, IntVal: 2861335956, UnsafePtr: unsafe.Pointer((&Constructor_Data_Map_Internal_IterNode[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](as_3), gopurs_runtime.Value{Type: 9, IntVal: 2509360378, UnsafePtr: unsafe.Pointer(nil)}}))}).UnsafePtr)
					unboxed := make([]*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value], len(arr))
					for i, v := range arr { unboxed[i] = gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](v) }
					return unboxed
				}()
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v)} }
					return gopurs_runtime.Array(boxed)
				}()).StrVal())) + (")"))
})})))}
}

func Call_Data_Map_Internal_isSubmap(dictOrd_0_loop *Constructor_Data_Ord_Ord[gopurs_runtime.Value], dictEq_1_loop *Constructor_Data_Eq_Eq[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictOrd_0 *Constructor_Data_Ord_Ord[gopurs_runtime.Value] = dictOrd_0_loop
_ = dictOrd_0
var dictEq_1 *Constructor_Data_Eq_Eq[gopurs_runtime.Value] = dictEq_1_loop
_ = dictEq_1
var go__go_2_0_17 gopurs_runtime.Value
_ = go__go_2_0_17
// FALLBACK TCO: isLoop=false len=1
go__go_2_0_17 = gopurs_runtime.Func2(func(m1_3 gopurs_runtime.Value, m2_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t9 bool
{
var __t_tag_1 *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](m1_3)
if (__t_tag_1 == nil) {
__t9 = true
goto end_branch_9
} else {

}
}
{
var __t_tag_2 *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](m1_3)
if (__t_tag_2 != nil) {
var Call_local_Data_Map_Internal_go__go_5_4_18 func(*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]) *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
_ = Call_local_Data_Map_Internal_go__go_5_4_18
var go__go_5_4_18 gopurs_runtime.Value
_ = go__go_5_4_18
Call_local_Data_Map_Internal_go__go_5_4_18 = func(v_6_loop *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]) *Constructor_Data_Maybe_Just[gopurs_runtime.Value] {
go__go_5_4_18:
for {
if false { continue go__go_5_4_18 }
var v_6 *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value] = v_6_loop
_ = v_6
var __t7 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
if (v_6 == nil) {
__t7 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil))})
goto end_branch_7
} else {

}
}
{
if (v_6 != nil) {
// TAST (Let): v1_7_5 shape=App(Other) bindingType=(ADT ["Data","Ordering","Ordering"] [])
v1_7_5 := uint32(gopurs_runtime.Apply2(gopurs_runtime.Box(dictOrd_0.V1), (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(m1_3.UnsafePtr).V2, (v_6).V2).IntVal)
_ = v1_7_5
var __t6 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
if (v1_7_5 == 1527465420) {
v_6_loop = (v_6).V4
continue go__go_5_4_18
__t6 = func() *Constructor_Data_Maybe_Just[gopurs_runtime.Value] { panic("unreachable") }()
goto end_branch_6
} else {

}
}
{
if (v1_7_5 == 380165415) {
v_6_loop = (v_6).V5
continue go__go_5_4_18
__t6 = func() *Constructor_Data_Maybe_Just[gopurs_runtime.Value] { panic("unreachable") }()
goto end_branch_6
} else {

}
}
{
if (v1_7_5 == 902936544) {
__t6 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{(v_6).V3, true}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())
goto end_branch_6
} else {

}
}
{
__t6 = func() *Constructor_Data_Maybe_Just[gopurs_runtime.Value] { panic("Failed pattern match") }()
}
end_branch_6:
__t7 = __t6
goto end_branch_7
} else {

}
}
{
__t7 = func() *Constructor_Data_Maybe_Just[gopurs_runtime.Value] { panic("Failed pattern match") }()
}
end_branch_7:
return __t7
}
}
go__go_5_4_18 = gopurs_runtime.Func(func(v_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_local_Data_Map_Internal_go__go_5_4_18(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](v_6_loop_val)))}
})
// TAST (Let): v1_5_3 shape=LetRec(App(Other)) bindingType=(ADT ["Data","Maybe","Maybe"] [(TypeVar v)])
v1_5_3 := Call_local_Data_Map_Internal_go__go_5_4_18(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](m2_4))
_ = v1_5_3
var __t8 bool
{
if (v1_5_3 == nil) {
__t8 = false
goto end_branch_8
} else {

}
}
{
if (v1_5_3 != nil) {
__t8 = ((gopurs_runtime.Apply2(gopurs_runtime.Box(dictEq_1.V0), (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(m1_3.UnsafePtr).V3, (v1_5_3).V0).IntVal) != (0)) && (((gopurs_runtime.Apply2(go__go_2_0_17, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(m1_3.UnsafePtr).V4)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](m2_4))}).IntVal) != (0)) && ((gopurs_runtime.Apply2(go__go_2_0_17, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(m1_3.UnsafePtr).V5)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](m2_4))}).IntVal) != (0)))
goto end_branch_8
} else {

}
}
{
__t8 = func() bool { panic("Failed pattern match") }()
}
end_branch_8:
__t9 = __t8
goto end_branch_9
} else {

}
}
{
__t9 = func() bool { panic("Failed pattern match") }()
}
end_branch_9:
return gopurs_runtime.Bool(__t9)
})
return go__go_2_0_17
}

func Call_Data_Map_Internal_isEmpty(v_0_loop *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]) bool {
var v_0 *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value] = v_0_loop
_ = v_0
return (v_0 == nil)
}

func Call_Data_Map_Internal_intersectionWith(dictOrd_0_loop *Constructor_Data_Ord_Ord[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictOrd_0 *Constructor_Data_Ord_Ord[gopurs_runtime.Value] = dictOrd_0_loop
_ = dictOrd_0
// TAST (Let): compare_1_0 shape=Other bindingType=(Func [(TypeVar k), (TypeVar k)] (ADT ["Data","Ordering","Ordering"] []))
compare_1_0 := gopurs_runtime.Box(dictOrd_0.V1)
_ = compare_1_0
return gopurs_runtime.Func3(func(app_2 gopurs_runtime.Value, m1_3 gopurs_runtime.Value, m2_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeIntersectionWith(), compare_1_0, app_2, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](m1_3))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](m2_4))})))}
})
}

func Call_Data_Map_Internal_intersection(dictOrd_0_loop *Constructor_Data_Ord_Ord[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictOrd_0 *Constructor_Data_Ord_Ord[gopurs_runtime.Value] = dictOrd_0_loop
_ = dictOrd_0
// TAST (Let): compare_1_0 shape=Other bindingType=(Func [(TypeVar k), (TypeVar k)] (ADT ["Data","Ordering","Ordering"] []))
compare_1_0 := gopurs_runtime.Box(dictOrd_0.V1)
_ = compare_1_0
return gopurs_runtime.Func2(func(m1_2 gopurs_runtime.Value, m2_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeIntersectionWith(), compare_1_0, Get_Data_Function_go__const(), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](m1_2))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](m2_3))})))}
})
}

func Call_Data_Map_Internal_insertWith(dictOrd_0_loop *Constructor_Data_Ord_Ord[gopurs_runtime.Value], app_1_loop gopurs_runtime.Value, k_2_loop gopurs_runtime.Value, v_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 *Constructor_Data_Ord_Ord[gopurs_runtime.Value] = dictOrd_0_loop
_ = dictOrd_0
var app_1 gopurs_runtime.Value = app_1_loop
_ = app_1
var k_2 gopurs_runtime.Value = k_2_loop
_ = k_2
var v_3 gopurs_runtime.Value = v_3_loop
_ = v_3
var go__go_4_0_19 gopurs_runtime.Value
_ = go__go_4_0_19
// FALLBACK TCO: isLoop=false len=1
go__go_4_0_19 = gopurs_runtime.Func(func(v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
var __t5 *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]
{
var __t_tag_1 *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](v1_5)
if (__t_tag_1 == nil) {
__t5 = (&Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]{1, int64(1), int64(1), k_2, v_3, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(nil), (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(nil)})
goto end_branch_5
} else {

}
}
{
var __t_tag_2 *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](v1_5)
if (__t_tag_2 != nil) {
// TAST (Let): v2_6_3 shape=App(Other) bindingType=(ADT ["Data","Ordering","Ordering"] [])
v2_6_3 := uint32(gopurs_runtime.Apply2(gopurs_runtime.Box(dictOrd_0.V1), k_2, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v1_5.UnsafePtr).V2).IntVal)
_ = v2_6_3
var __t4 *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]
{
if (v2_6_3 == 1527465420) {
__t4 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeBalancedNode(), (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v1_5.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v1_5.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(go__go_4_0_19, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v1_5.UnsafePtr).V4)})))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v1_5.UnsafePtr).V5)}))
goto end_branch_4
} else {

}
}
{
if (v2_6_3 == 380165415) {
__t4 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeBalancedNode(), (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v1_5.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v1_5.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v1_5.UnsafePtr).V4)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(go__go_4_0_19, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v1_5.UnsafePtr).V5)})))}))
goto end_branch_4
} else {

}
}
{
if (v2_6_3 == 902936544) {
__t4 = (&Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v1_5.UnsafePtr).V0, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v1_5.UnsafePtr).V1, k_2, gopurs_runtime.Apply2(app_1, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v1_5.UnsafePtr).V3, v_3), (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v1_5.UnsafePtr).V4, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v1_5.UnsafePtr).V5})
goto end_branch_4
} else {

}
}
{
__t4 = func() *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value] { panic("Failed pattern match") }()
}
end_branch_4:
__t5 = __t4
goto end_branch_5
} else {

}
}
{
__t5 = func() *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value] { panic("Failed pattern match") }()
}
end_branch_5:
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(__t5)}
})
return go__go_4_0_19
}

func Call_Data_Map_Internal_insert(dictOrd_0_loop *Constructor_Data_Ord_Ord[gopurs_runtime.Value], k_1_loop gopurs_runtime.Value, v_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 *Constructor_Data_Ord_Ord[gopurs_runtime.Value] = dictOrd_0_loop
_ = dictOrd_0
var k_1 gopurs_runtime.Value = k_1_loop
_ = k_1
var v_2 gopurs_runtime.Value = v_2_loop
_ = v_2
var go__go_3_0_20 gopurs_runtime.Value
_ = go__go_3_0_20
// FALLBACK TCO: isLoop=false len=1
go__go_3_0_20 = gopurs_runtime.Func(func(v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t5 *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]
{
var __t_tag_1 *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](v1_4)
if (__t_tag_1 == nil) {
__t5 = (&Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]{1, int64(1), int64(1), k_1, v_2, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(nil), (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(nil)})
goto end_branch_5
} else {

}
}
{
var __t_tag_2 *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](v1_4)
if (__t_tag_2 != nil) {
// TAST (Let): v2_5_3 shape=App(Other) bindingType=(ADT ["Data","Ordering","Ordering"] [])
v2_5_3 := uint32(gopurs_runtime.Apply2(gopurs_runtime.Box(dictOrd_0.V1), k_1, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v1_4.UnsafePtr).V2).IntVal)
_ = v2_5_3
var __t4 *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]
{
if (v2_5_3 == 1527465420) {
__t4 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeBalancedNode(), (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v1_4.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v1_4.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(go__go_3_0_20, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v1_4.UnsafePtr).V4)})))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v1_4.UnsafePtr).V5)}))
goto end_branch_4
} else {

}
}
{
if (v2_5_3 == 380165415) {
__t4 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeBalancedNode(), (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v1_4.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v1_4.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v1_4.UnsafePtr).V4)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(go__go_3_0_20, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v1_4.UnsafePtr).V5)})))}))
goto end_branch_4
} else {

}
}
{
if (v2_5_3 == 902936544) {
__t4 = (&Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v1_4.UnsafePtr).V0, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v1_4.UnsafePtr).V1, k_1, v_2, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v1_4.UnsafePtr).V4, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v1_4.UnsafePtr).V5})
goto end_branch_4
} else {

}
}
{
__t4 = func() *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value] { panic("Failed pattern match") }()
}
end_branch_4:
__t5 = __t4
goto end_branch_5
} else {

}
}
{
__t5 = func() *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value] { panic("Failed pattern match") }()
}
end_branch_5:
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(__t5)}
})
return go__go_3_0_20
}

func Call_Data_Map_Internal_foldSubmapBy(dictOrd_0_loop *Constructor_Data_Ord_Ord[gopurs_runtime.Value], appendFn_1_loop gopurs_runtime.Value, memptyValue_2_loop gopurs_runtime.Value, kmin_3_loop *Constructor_Data_Maybe_Just[gopurs_runtime.Value], kmax_4_loop *Constructor_Data_Maybe_Just[gopurs_runtime.Value], f_5_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 *Constructor_Data_Ord_Ord[gopurs_runtime.Value] = dictOrd_0_loop
_ = dictOrd_0
var appendFn_1 gopurs_runtime.Value = appendFn_1_loop
_ = appendFn_1
var memptyValue_2 gopurs_runtime.Value = memptyValue_2_loop
_ = memptyValue_2
var kmin_3 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = kmin_3_loop
_ = kmin_3
var kmax_4 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = kmax_4_loop
_ = kmax_4
var f_5 gopurs_runtime.Value = f_5_loop
_ = f_5
var __t3 gopurs_runtime.Value
{
if (kmin_3 != nil) {
// TAST (Let): __local_var_6_1 shape=Other bindingType=Any
__local_var_6_1 := (kmin_3).V0
_ = __local_var_6_1
__t3 = gopurs_runtime.Func(func(k_7 gopurs_runtime.Value) gopurs_runtime.Value {
var __t_tag_2 uint32 = uint32(gopurs_runtime.Apply2(gopurs_runtime.Box(dictOrd_0.V1), k_7, __local_var_6_1).IntVal)
return gopurs_runtime.Bool((uint32(__t_tag_2) == 1527465420))
})
goto end_branch_3
} else {

}
}
{
if (kmin_3 == nil) {
__t3 = gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(false)
})
goto end_branch_3
} else {

}
}
{
__t3 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_3:
// TAST (Let): tooSmall_6_0 shape=Branch(Let(Abs(Other)), Abs(LitBoolean), def=Other) bindingType=Any
tooSmall_6_0 := __t3
_ = tooSmall_6_0
var __t7 gopurs_runtime.Value
{
if (kmax_4 != nil) {
// TAST (Let): __local_var_7_5 shape=Other bindingType=Any
__local_var_7_5 := (kmax_4).V0
_ = __local_var_7_5
__t7 = gopurs_runtime.Func(func(k_8 gopurs_runtime.Value) gopurs_runtime.Value {
var __t_tag_6 uint32 = uint32(gopurs_runtime.Apply2(gopurs_runtime.Box(dictOrd_0.V1), k_8, __local_var_7_5).IntVal)
return gopurs_runtime.Bool((uint32(__t_tag_6) == 380165415))
})
goto end_branch_7
} else {

}
}
{
if (kmax_4 == nil) {
__t7 = gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(false)
})
goto end_branch_7
} else {

}
}
{
__t7 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_7:
// TAST (Let): tooLarge_7_4 shape=Branch(Let(Abs(Other)), Abs(LitBoolean), def=Other) bindingType=Any
tooLarge_7_4 := __t7
_ = tooLarge_7_4
var __t20 gopurs_runtime.Value
{
if (kmin_3 != nil) {
var __t16 gopurs_runtime.Value
{
if (kmax_4 != nil) {
// TAST (Let): __local_var_8_9 shape=Other bindingType=Any
__local_var_8_9 := (kmax_4).V0
_ = __local_var_8_9
// TAST (Let): __local_var_9_10 shape=Other bindingType=Any
__local_var_9_10 := (kmin_3).V0
_ = __local_var_9_10
__t16 = gopurs_runtime.Func(func(k_10 gopurs_runtime.Value) gopurs_runtime.Value {
var __t_tag_11 uint32 = uint32(gopurs_runtime.Apply2(gopurs_runtime.Box(dictOrd_0.V1), __local_var_9_10, k_10).IntVal)
var __t_and_13 bool = false
if ((uint32(__t_tag_11) == 380165415)) != (true) {

var __t_tag_12 uint32 = uint32(gopurs_runtime.Apply2(gopurs_runtime.Box(dictOrd_0.V1), k_10, __local_var_8_9).IntVal)
__t_and_13 = ((uint32(__t_tag_12) == 380165415)) != (true)
}
return gopurs_runtime.Bool(__t_and_13)
})
goto end_branch_16
} else {

}
}
{
if (kmax_4 == nil) {
// TAST (Let): __local_var_8_14 shape=Other bindingType=Any
__local_var_8_14 := (kmin_3).V0
_ = __local_var_8_14
__t16 = gopurs_runtime.Func(func(k_9 gopurs_runtime.Value) gopurs_runtime.Value {
var __t_tag_15 uint32 = uint32(gopurs_runtime.Apply2(gopurs_runtime.Box(dictOrd_0.V1), __local_var_8_14, k_9).IntVal)
return gopurs_runtime.Bool(((uint32(__t_tag_15) == 380165415)) != (true))
})
goto end_branch_16
} else {

}
}
{
__t16 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_16:
__t20 = __t16
goto end_branch_20
} else {

}
}
{
if (kmin_3 == nil) {
var __t19 gopurs_runtime.Value
{
if (kmax_4 != nil) {
// TAST (Let): __local_var_8_17 shape=Other bindingType=Any
__local_var_8_17 := (kmax_4).V0
_ = __local_var_8_17
__t19 = gopurs_runtime.Func(func(k_9 gopurs_runtime.Value) gopurs_runtime.Value {
var __t_tag_18 uint32 = uint32(gopurs_runtime.Apply2(gopurs_runtime.Box(dictOrd_0.V1), k_9, __local_var_8_17).IntVal)
return gopurs_runtime.Bool(((uint32(__t_tag_18) == 380165415)) != (true))
})
goto end_branch_19
} else {

}
}
{
if (kmax_4 == nil) {
__t19 = gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(true)
})
goto end_branch_19
} else {

}
}
{
__t19 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_19:
__t20 = __t19
goto end_branch_20
} else {

}
}
{
__t20 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_20:
// TAST (Let): inBounds_8_8 shape=Branch(Branch(Let(Let(Abs(Other))), Let(Abs(Other)), def=Other), Branch(Let(Abs(Other)), Abs(LitBoolean), def=Other), def=Other) bindingType=Any
inBounds_8_8 := __t20
_ = inBounds_8_8
var go__go_9_21_33 gopurs_runtime.Value
_ = go__go_9_21_33
// FALLBACK TCO: isLoop=false len=1
go__go_9_21_33 = gopurs_runtime.Func(func(v_10 gopurs_runtime.Value) gopurs_runtime.Value {
var __t27 gopurs_runtime.Value
{
var __t_tag_22 *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](v_10)
if (__t_tag_22 == nil) {
__t27 = memptyValue_2
goto end_branch_27
} else {

}
}
{
var __t_tag_23 *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](v_10)
if (__t_tag_23 != nil) {
var __t24 gopurs_runtime.Value
{
if (gopurs_runtime.Apply(tooSmall_6_0, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_10.UnsafePtr).V2).IntVal) != (0) {
__t24 = memptyValue_2
goto end_branch_24
} else {

}
}
{
__t24 = gopurs_runtime.Apply(go__go_9_21_33, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_10.UnsafePtr).V4)})
}
end_branch_24:
var __t25 gopurs_runtime.Value
{
if (gopurs_runtime.Apply(inBounds_8_8, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_10.UnsafePtr).V2).IntVal) != (0) {
__t25 = gopurs_runtime.Apply2(f_5, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_10.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_10.UnsafePtr).V3)
goto end_branch_25
} else {

}
}
{
__t25 = memptyValue_2
}
end_branch_25:
var __t26 gopurs_runtime.Value
{
if (gopurs_runtime.Apply(tooLarge_7_4, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_10.UnsafePtr).V2).IntVal) != (0) {
__t26 = memptyValue_2
goto end_branch_26
} else {

}
}
{
__t26 = gopurs_runtime.Apply(go__go_9_21_33, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_10.UnsafePtr).V5)})
}
end_branch_26:
__t27 = gopurs_runtime.Apply2(appendFn_1, gopurs_runtime.Apply2(appendFn_1, __t24, __t25), __t26)
goto end_branch_27
} else {

}
}
{
__t27 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_27:
return __t27
})
return go__go_9_21_33
}

func Call_Data_Map_Internal_foldSubmap(dictOrd_0_loop *Constructor_Data_Ord_Ord[gopurs_runtime.Value], dictMonoid_1_loop *Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictOrd_0 *Constructor_Data_Ord_Ord[gopurs_runtime.Value] = dictOrd_0_loop
_ = dictOrd_0
var dictMonoid_1 *Constructor_Data_Monoid_Monoid[gopurs_runtime.Value] = dictMonoid_1_loop
_ = dictMonoid_1
return gopurs_runtime.Apply3(Get_Data_Map_Internal_foldSubmapBy(), gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(dictOrd_0)}, gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.Box(dictMonoid_1.V0), gopurs_runtime.Value{}), "append"), gopurs_runtime.Box(dictMonoid_1.V1))
}

func Call_Data_Map_Internal_findMin(v_0_loop *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]) struct{V0 gopurs_runtime.Value; V1 bool} {
findMin:
for {
if false { continue findMin }
var v_0 *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value] = v_0_loop
_ = v_0
var __t2 *Constructor_Data_Maybe_Just[struct{
	key gopurs_runtime.Value
	value gopurs_runtime.Value
}]
{
if (v_0 == nil) {
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[struct{
	key gopurs_runtime.Value
	value gopurs_runtime.Value
}]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())
goto end_branch_2
} else {

}
}
{
if (v_0 != nil) {
var __t1 *Constructor_Data_Maybe_Just[struct{
	key gopurs_runtime.Value
	value gopurs_runtime.Value
}]
{
var __t_tag_0 *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value] = (v_0).V4
if (__t_tag_0 == nil) {
__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[struct{
	key gopurs_runtime.Value
	value gopurs_runtime.Value
}]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.RecordDict2("key", "value", (v_0).V2, (v_0).V3), true}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())
goto end_branch_1
} else {

}
}
{
v_0_loop = (v_0).V4
continue findMin
__t1 = func() *Constructor_Data_Maybe_Just[struct{
	key gopurs_runtime.Value
	value gopurs_runtime.Value
}] { panic("unreachable") }()
}
end_branch_1:
__t2 = __t1
goto end_branch_2
} else {

}
}
{
__t2 = func() *Constructor_Data_Maybe_Just[struct{
	key gopurs_runtime.Value
	value gopurs_runtime.Value
}] { panic("Failed pattern match") }()
}
end_branch_2:
return func() struct{V0 gopurs_runtime.Value; V1 bool} {
				_v := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_Map_Internal_2758605161_3094389156(__t2))}
				if _v.Type == 9 && _v.IntVal == 3562950408 && _v.UnsafePtr != nil {
					return struct{V0 gopurs_runtime.Value; V1 bool}{V0: (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr).V0, V1: true}
				}
				return struct{V0 gopurs_runtime.Value; V1 bool}{V0: gopurs_runtime.Value{}, V1: false}
			}()
}
}

func Call_Data_Map_Internal_lookupGT(dictOrd_0_loop *Constructor_Data_Ord_Ord[gopurs_runtime.Value], k_1_loop gopurs_runtime.Value) struct{V0 gopurs_runtime.Value; V1 bool} {
var dictOrd_0 *Constructor_Data_Ord_Ord[gopurs_runtime.Value] = dictOrd_0_loop
_ = dictOrd_0
var k_1 gopurs_runtime.Value = k_1_loop
_ = k_1
var go__go_2_0_34 gopurs_runtime.Value
_ = go__go_2_0_34
// FALLBACK TCO: isLoop=false len=1
go__go_2_0_34 = gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t7 *Constructor_Data_Maybe_Just[struct{
	key gopurs_runtime.Value
	value gopurs_runtime.Value
}]
{
var __t_tag_1 *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](v_3)
if (__t_tag_1 == nil) {
__t7 = Rebox_Data_Map_Internal_3094389156_2758605161(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil))}))
goto end_branch_7
} else {

}
}
{
var __t_tag_2 *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](v_3)
if (__t_tag_2 != nil) {
// TAST (Let): v1_4_3 shape=App(Other) bindingType=(ADT ["Data","Ordering","Ordering"] [])
v1_4_3 := uint32(gopurs_runtime.Apply2(gopurs_runtime.Box(dictOrd_0.V1), k_1, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V2).IntVal)
_ = v1_4_3
var __t6 *Constructor_Data_Maybe_Just[struct{
	key gopurs_runtime.Value
	value gopurs_runtime.Value
}]
{
if (v1_4_3 == 1527465420) {
// TAST (Let): v2_5_4 shape=App(Other) bindingType=(ADT ["Data","Maybe","Maybe"] [(Record (Row [key: (TypeVar k), value: (TypeVar v)] Any))])
v2_5_4 := Rebox_Data_Map_Internal_3094389156_2758605161(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(go__go_2_0_34, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V4)})))
_ = v2_5_4
var __t5 *Constructor_Data_Maybe_Just[struct{
	key gopurs_runtime.Value
	value gopurs_runtime.Value
}]
{
if (v2_5_4 == nil) {
__t5 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[struct{
	key gopurs_runtime.Value
	value gopurs_runtime.Value
}]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.RecordDict2("key", "value", (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V3), true}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())
goto end_branch_5
} else {

}
}
{
__t5 = v2_5_4
}
end_branch_5:
__t6 = __t5
goto end_branch_6
} else {

}
}
{
if (v1_4_3 == 380165415) {
__t6 = Rebox_Data_Map_Internal_3094389156_2758605161(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(go__go_2_0_34, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V5)})))
goto end_branch_6
} else {

}
}
{
if (v1_4_3 == 902936544) {
__t6 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[struct{
	key gopurs_runtime.Value
	value gopurs_runtime.Value
}]](func() gopurs_runtime.Value {
				_v := Call_Data_Map_Internal_findMin((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V5)
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())
goto end_branch_6
} else {

}
}
{
__t6 = func() *Constructor_Data_Maybe_Just[struct{
	key gopurs_runtime.Value
	value gopurs_runtime.Value
}] { panic("Failed pattern match") }()
}
end_branch_6:
__t7 = __t6
goto end_branch_7
} else {

}
}
{
__t7 = func() *Constructor_Data_Maybe_Just[struct{
	key gopurs_runtime.Value
	value gopurs_runtime.Value
}] { panic("Failed pattern match") }()
}
end_branch_7:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_Map_Internal_2758605161_3094389156(__t7))}
})
return func() struct{V0 gopurs_runtime.Value; V1 bool} {
				_v := go__go_2_0_34
				if _v.Type == 9 && _v.IntVal == 3562950408 && _v.UnsafePtr != nil {
					return struct{V0 gopurs_runtime.Value; V1 bool}{V0: (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr).V0, V1: true}
				}
				return struct{V0 gopurs_runtime.Value; V1 bool}{V0: gopurs_runtime.Value{}, V1: false}
			}()
}

func Call_Data_Map_Internal_findMax(v_0_loop *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]) struct{V0 gopurs_runtime.Value; V1 bool} {
findMax:
for {
if false { continue findMax }
var v_0 *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value] = v_0_loop
_ = v_0
var __t2 *Constructor_Data_Maybe_Just[struct{
	key gopurs_runtime.Value
	value gopurs_runtime.Value
}]
{
if (v_0 == nil) {
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[struct{
	key gopurs_runtime.Value
	value gopurs_runtime.Value
}]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())
goto end_branch_2
} else {

}
}
{
if (v_0 != nil) {
var __t1 *Constructor_Data_Maybe_Just[struct{
	key gopurs_runtime.Value
	value gopurs_runtime.Value
}]
{
var __t_tag_0 *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value] = (v_0).V5
if (__t_tag_0 == nil) {
__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[struct{
	key gopurs_runtime.Value
	value gopurs_runtime.Value
}]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.RecordDict2("key", "value", (v_0).V2, (v_0).V3), true}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())
goto end_branch_1
} else {

}
}
{
v_0_loop = (v_0).V5
continue findMax
__t1 = func() *Constructor_Data_Maybe_Just[struct{
	key gopurs_runtime.Value
	value gopurs_runtime.Value
}] { panic("unreachable") }()
}
end_branch_1:
__t2 = __t1
goto end_branch_2
} else {

}
}
{
__t2 = func() *Constructor_Data_Maybe_Just[struct{
	key gopurs_runtime.Value
	value gopurs_runtime.Value
}] { panic("Failed pattern match") }()
}
end_branch_2:
return func() struct{V0 gopurs_runtime.Value; V1 bool} {
				_v := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_Map_Internal_2758605161_3094389156(__t2))}
				if _v.Type == 9 && _v.IntVal == 3562950408 && _v.UnsafePtr != nil {
					return struct{V0 gopurs_runtime.Value; V1 bool}{V0: (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr).V0, V1: true}
				}
				return struct{V0 gopurs_runtime.Value; V1 bool}{V0: gopurs_runtime.Value{}, V1: false}
			}()
}
}

func Call_Data_Map_Internal_lookupLT(dictOrd_0_loop *Constructor_Data_Ord_Ord[gopurs_runtime.Value], k_1_loop gopurs_runtime.Value) struct{V0 gopurs_runtime.Value; V1 bool} {
var dictOrd_0 *Constructor_Data_Ord_Ord[gopurs_runtime.Value] = dictOrd_0_loop
_ = dictOrd_0
var k_1 gopurs_runtime.Value = k_1_loop
_ = k_1
var go__go_2_0_35 gopurs_runtime.Value
_ = go__go_2_0_35
// FALLBACK TCO: isLoop=false len=1
go__go_2_0_35 = gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t7 *Constructor_Data_Maybe_Just[struct{
	key gopurs_runtime.Value
	value gopurs_runtime.Value
}]
{
var __t_tag_1 *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](v_3)
if (__t_tag_1 == nil) {
__t7 = Rebox_Data_Map_Internal_3094389156_2758605161(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil))}))
goto end_branch_7
} else {

}
}
{
var __t_tag_2 *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](v_3)
if (__t_tag_2 != nil) {
// TAST (Let): v1_4_3 shape=App(Other) bindingType=(ADT ["Data","Ordering","Ordering"] [])
v1_4_3 := uint32(gopurs_runtime.Apply2(gopurs_runtime.Box(dictOrd_0.V1), k_1, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V2).IntVal)
_ = v1_4_3
var __t6 *Constructor_Data_Maybe_Just[struct{
	key gopurs_runtime.Value
	value gopurs_runtime.Value
}]
{
if (v1_4_3 == 1527465420) {
__t6 = Rebox_Data_Map_Internal_3094389156_2758605161(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(go__go_2_0_35, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V4)})))
goto end_branch_6
} else {

}
}
{
if (v1_4_3 == 380165415) {
// TAST (Let): v2_5_4 shape=App(Other) bindingType=(ADT ["Data","Maybe","Maybe"] [(Record (Row [key: (TypeVar k), value: (TypeVar v)] Any))])
v2_5_4 := Rebox_Data_Map_Internal_3094389156_2758605161(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(go__go_2_0_35, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V5)})))
_ = v2_5_4
var __t5 *Constructor_Data_Maybe_Just[struct{
	key gopurs_runtime.Value
	value gopurs_runtime.Value
}]
{
if (v2_5_4 == nil) {
__t5 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[struct{
	key gopurs_runtime.Value
	value gopurs_runtime.Value
}]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.RecordDict2("key", "value", (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V3), true}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())
goto end_branch_5
} else {

}
}
{
__t5 = v2_5_4
}
end_branch_5:
__t6 = __t5
goto end_branch_6
} else {

}
}
{
if (v1_4_3 == 902936544) {
__t6 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[struct{
	key gopurs_runtime.Value
	value gopurs_runtime.Value
}]](func() gopurs_runtime.Value {
				_v := Call_Data_Map_Internal_findMax((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V4)
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())
goto end_branch_6
} else {

}
}
{
__t6 = func() *Constructor_Data_Maybe_Just[struct{
	key gopurs_runtime.Value
	value gopurs_runtime.Value
}] { panic("Failed pattern match") }()
}
end_branch_6:
__t7 = __t6
goto end_branch_7
} else {

}
}
{
__t7 = func() *Constructor_Data_Maybe_Just[struct{
	key gopurs_runtime.Value
	value gopurs_runtime.Value
}] { panic("Failed pattern match") }()
}
end_branch_7:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_Map_Internal_2758605161_3094389156(__t7))}
})
return func() struct{V0 gopurs_runtime.Value; V1 bool} {
				_v := go__go_2_0_35
				if _v.Type == 9 && _v.IntVal == 3562950408 && _v.UnsafePtr != nil {
					return struct{V0 gopurs_runtime.Value; V1 bool}{V0: (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr).V0, V1: true}
				}
				return struct{V0 gopurs_runtime.Value; V1 bool}{V0: gopurs_runtime.Value{}, V1: false}
			}()
}

func Call_Data_Map_Internal_filterWithKey(dictOrd_0_loop *Constructor_Data_Ord_Ord[gopurs_runtime.Value], f_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 *Constructor_Data_Ord_Ord[gopurs_runtime.Value] = dictOrd_0_loop
_ = dictOrd_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var go__go_2_0_36 gopurs_runtime.Value
_ = go__go_2_0_36
// FALLBACK TCO: isLoop=false len=1
go__go_2_0_36 = gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t4 *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]
{
var __t_tag_1 *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](v_3)
if (__t_tag_1 == nil) {
__t4 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(nil))})
goto end_branch_4
} else {

}
}
{
var __t_tag_2 *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](v_3)
if (__t_tag_2 != nil) {
var __t3 *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]
{
if (gopurs_runtime.Apply2(f_1, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V3).IntVal) != (0) {
__t3 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeBalancedNode(), (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(go__go_2_0_36, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V4)})))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(go__go_2_0_36, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V5)})))}))
goto end_branch_3
} else {

}
}
{
__t3 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp2(Get_Data_Map_Internal_unsafeJoinNodes(), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(go__go_2_0_36, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V4)})))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(go__go_2_0_36, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V5)})))}))
}
end_branch_3:
__t4 = __t3
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
return go__go_2_0_36
}

func Call_Data_Map_Internal_filterKeys(dictOrd_0_loop *Constructor_Data_Ord_Ord[gopurs_runtime.Value], f_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 *Constructor_Data_Ord_Ord[gopurs_runtime.Value] = dictOrd_0_loop
_ = dictOrd_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var go__go_2_0_37 gopurs_runtime.Value
_ = go__go_2_0_37
// FALLBACK TCO: isLoop=false len=1
go__go_2_0_37 = gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t4 *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]
{
var __t_tag_1 *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](v_3)
if (__t_tag_1 == nil) {
__t4 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(nil))})
goto end_branch_4
} else {

}
}
{
var __t_tag_2 *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](v_3)
if (__t_tag_2 != nil) {
var __t3 *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]
{
if (gopurs_runtime.Apply(f_1, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V2).IntVal) != (0) {
__t3 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeBalancedNode(), (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(go__go_2_0_37, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V4)})))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(go__go_2_0_37, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V5)})))}))
goto end_branch_3
} else {

}
}
{
__t3 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp2(Get_Data_Map_Internal_unsafeJoinNodes(), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(go__go_2_0_37, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V4)})))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(go__go_2_0_37, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V5)})))}))
}
end_branch_3:
__t4 = __t3
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
return go__go_2_0_37
}

func Call_Data_Map_Internal_filter(dictOrd_0_loop *Constructor_Data_Ord_Ord[gopurs_runtime.Value], x_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 *Constructor_Data_Ord_Ord[gopurs_runtime.Value] = dictOrd_0_loop
_ = dictOrd_0
var x_1 gopurs_runtime.Value = x_1_loop
_ = x_1
return Call_Data_Map_Internal_filterWithKey(dictOrd_0, gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return x_1
}))
}

func Call_Data_Map_Internal_eqMap(dictEq_0_loop gopurs_runtime.Value, dictEq1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictEq_0 gopurs_runtime.Value = dictEq_0_loop
_ = dictEq_0
var dictEq1_1 gopurs_runtime.Value = dictEq1_1_loop
_ = dictEq1_1
var go__go_2_1_38 gopurs_runtime.Value
_ = go__go_2_1_38
// FALLBACK TCO: isLoop=false len=1
go__go_2_1_38 = gopurs_runtime.Func2(func(a_3 gopurs_runtime.Value, b_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): v_5_2 shape=App(Var) bindingType=(ADT ["Data","Map","Internal","MapIterStep"] [(TypeVar k), (TypeVar v)])
v_5_2 := gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_IterNext[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Map_Internal_stepAsc(), a_3))
_ = v_5_2
var __t4 bool
{
if (v_5_2 != nil) {
// TAST (Let): v2_6_3 shape=App(Var) bindingType=(ADT ["Data","Map","Internal","MapIterStep"] [(TypeVar k), (TypeVar v)])
v2_6_3 := gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_IterNext[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Map_Internal_stepAsc(), b_4))
_ = v2_6_3
__t4 = ((v2_6_3 != nil)) && ((((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq_0, "eq"), (v_5_2).V0, (v2_6_3).V0).IntVal) != (0)) && ((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq1_1, "eq"), (v_5_2).V1, (v2_6_3).V1).IntVal) != (0))) && ((gopurs_runtime.Apply2(go__go_2_1_38, (v_5_2).V2, (v2_6_3).V2).IntVal) != (0)))
goto end_branch_4
} else {

}
}
{
if (v_5_2 == nil) {
__t4 = true
goto end_branch_4
} else {

}
}
{
__t4 = func() bool { panic("Failed pattern match") }()
}
end_branch_4:
return gopurs_runtime.Bool(__t4)
})
// TAST (Let): eqMapIter2_2_0 shape=LitRecord bindingType=(ADT ["Data","Eq","Eq"] [(ADT ["Data","Map","Internal","MapIter"] [(TypeVar k), (TypeVar v)])])
eqMapIter2_2_0 := (&Constructor_Data_Eq_Eq[gopurs_runtime.Value]{1, go__go_2_1_38})
_ = eqMapIter2_2_0
return gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Data_Map_Internal_1444241223_3790796878((&Constructor_Data_Eq_Eq[*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]]{1, gopurs_runtime.Func2(func(xs_3 gopurs_runtime.Value, ys_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t9 bool
{
var __t_tag_5 *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](xs_3)
if (__t_tag_5 == nil) {
var __t_tag_6 *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](ys_4)
__t9 = (__t_tag_6 == nil)
goto end_branch_9
} else {

}
}
{
var __t_tag_7 *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](xs_3)
if (__t_tag_7 != nil) {
var __t_tag_8 *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](ys_4)
__t9 = ((__t_tag_8 != nil)) && ((((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(xs_3.UnsafePtr).V1) == ((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(ys_4.UnsafePtr).V1)) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(eqMapIter2_2_0.V0), gopurs_runtime.Value{Type: 9, IntVal: 2861335956, UnsafePtr: unsafe.Pointer((&Constructor_Data_Map_Internal_IterNode[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](xs_3), gopurs_runtime.Value{Type: 9, IntVal: 2509360378, UnsafePtr: unsafe.Pointer(nil)}}))}, gopurs_runtime.Value{Type: 9, IntVal: 2861335956, UnsafePtr: unsafe.Pointer((&Constructor_Data_Map_Internal_IterNode[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](ys_4), gopurs_runtime.Value{Type: 9, IntVal: 2509360378, UnsafePtr: unsafe.Pointer(nil)}}))}).IntVal) != (0)))
goto end_branch_9
} else {

}
}
{
__t9 = func() bool { panic("Failed pattern match") }()
}
end_branch_9:
return gopurs_runtime.Bool(__t9)
})})))}
}

func Call_Data_Map_Internal_ordMap(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
// TAST (Let): __local_var_1_2 shape=App(Other) bindingType=Any
__local_var_1_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd_0, "Eq0"), gopurs_runtime.Value{})
_ = __local_var_1_2
// TAST (Let): eqMapIter1__193435443_1_1 shape=Let(Abs(LitRecord)) bindingType=Any
eqMapIter1__193435443_1_1 := gopurs_runtime.Func(func(dictEq1_2 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_3_3_39 gopurs_runtime.Value
_ = go__go_3_3_39
// FALLBACK TCO: isLoop=false len=1
go__go_3_3_39 = gopurs_runtime.Func2(func(a_4 gopurs_runtime.Value, b_5 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): v_6_4 shape=App(Var) bindingType=(ADT ["Data","Map","Internal","MapIterStep"] [(TypeVar k), (TypeVar v)])
v_6_4 := gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_IterNext[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Map_Internal_stepAsc(), a_4))
_ = v_6_4
var __t6 bool
{
if (v_6_4 != nil) {
// TAST (Let): v2_7_5 shape=App(Var) bindingType=(ADT ["Data","Map","Internal","MapIterStep"] [(TypeVar k), (TypeVar v)])
v2_7_5 := gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_IterNext[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Map_Internal_stepAsc(), b_5))
_ = v2_7_5
__t6 = ((v2_7_5 != nil)) && ((((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_2, "eq"), (v_6_4).V0, (v2_7_5).V0).IntVal) != (0)) && ((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq1_2, "eq"), (v_6_4).V1, (v2_7_5).V1).IntVal) != (0))) && ((gopurs_runtime.Apply2(go__go_3_3_39, (v_6_4).V2, (v2_7_5).V2).IntVal) != (0)))
goto end_branch_6
} else {

}
}
{
if (v_6_4 == nil) {
__t6 = true
goto end_branch_6
} else {

}
}
{
__t6 = func() bool { panic("Failed pattern match") }()
}
end_branch_6:
return gopurs_runtime.Bool(__t6)
})
return gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer((&Constructor_Data_Eq_Eq[gopurs_runtime.Value]{1, go__go_3_3_39}))}
})
_ = eqMapIter1__193435443_1_1
// TAST (Let): ordMapIter1__193435443_1_0 shape=Let(Abs(Let(LitRecord))) bindingType=Any
ordMapIter1__193435443_1_0 := gopurs_runtime.Func(func(dictOrd1_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): eqMapIter2_3_7 shape=App(Other) bindingType=(ADT ["Data","Eq","Eq"] [(ADT ["Data","Map","Internal","MapIter"] [(TypeVar k), (TypeVar v)])])
eqMapIter2_3_7 := gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](gopurs_runtime.Apply(eqMapIter1__193435443_1_1, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd1_2, "Eq0"), gopurs_runtime.Value{})))
_ = eqMapIter2_3_7
var Call_local_Data_Map_Internal_go__go_4_8_40 func(gopurs_runtime.Value, gopurs_runtime.Value) uint32
_ = Call_local_Data_Map_Internal_go__go_4_8_40
var go__go_4_8_40 gopurs_runtime.Value
_ = go__go_4_8_40
Call_local_Data_Map_Internal_go__go_4_8_40 = func(a_5_loop gopurs_runtime.Value, b_6_loop gopurs_runtime.Value) uint32 {
go__go_4_8_40:
for {
if false { continue go__go_4_8_40 }
var a_5 gopurs_runtime.Value = a_5_loop
_ = a_5
var b_6 gopurs_runtime.Value = b_6_loop
_ = b_6
// TAST (Let): v_7_9 shape=App(Var) bindingType=(ADT ["Data","Map","Internal","MapIterStep"] [(TypeVar k), (TypeVar v)])
v_7_9 := gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_IterNext[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Map_Internal_stepAsc(), b_6))
_ = v_7_9
// TAST (Let): v1_8_10 shape=App(Var) bindingType=(ADT ["Data","Map","Internal","MapIterStep"] [(TypeVar k), (TypeVar v)])
v1_8_10 := gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_IterNext[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Map_Internal_stepAsc(), a_5))
_ = v1_8_10
var __t17 uint32
{
if (v1_8_10 != nil) {
var __t15 uint32
{
if (v_7_9 != nil) {
// TAST (Let): v3_9_11 shape=App(Other) bindingType=(ADT ["Data","Ordering","Ordering"] [])
v3_9_11 := uint32(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictOrd_0, "compare"), (v1_8_10).V0, (v_7_9).V0).IntVal)
_ = v3_9_11
var __t14 uint32
{
if (v3_9_11 == 902936544) {
// TAST (Let): v4_10_12 shape=App(Other) bindingType=(ADT ["Data","Ordering","Ordering"] [])
v4_10_12 := uint32(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictOrd1_2, "compare"), (v1_8_10).V1, (v_7_9).V1).IntVal)
_ = v4_10_12
var __t13 uint32
{
if (v4_10_12 == 902936544) {
a_5_loop = (v1_8_10).V2
b_6_loop = (v_7_9).V2
continue go__go_4_8_40
__t13 = func() uint32 { panic("unreachable") }()
goto end_branch_13
} else {

}
}
{
__t13 = v4_10_12
}
end_branch_13:
__t14 = __t13
goto end_branch_14
} else {

}
}
{
__t14 = v3_9_11
}
end_branch_14:
__t15 = __t14
goto end_branch_15
} else {

}
}
{
if (v_7_9 == nil) {
__t15 = 380165415
goto end_branch_15
} else {

}
}
{
__t15 = func() uint32 { panic("Failed pattern match") }()
}
end_branch_15:
__t17 = __t15
goto end_branch_17
} else {

}
}
{
if (v1_8_10 == nil) {
var __t16 uint32
{
if (v_7_9 == nil) {
__t16 = 902936544
goto end_branch_16
} else {

}
}
{
__t16 = 1527465420
}
end_branch_16:
__t17 = __t16
goto end_branch_17
} else {

}
}
{
if (v_7_9 == nil) {
__t17 = 380165415
goto end_branch_17
} else {

}
}
{
__t17 = func() uint32 { panic("Failed pattern match") }()
}
end_branch_17:
return __t17
}
}
go__go_4_8_40 = gopurs_runtime.Func(func(a_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: int64(Call_local_Data_Map_Internal_go__go_4_8_40(a_5_loop_val, b_6_loop_val)), UnsafePtr: nil}
})
})
return gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer((&Constructor_Data_Ord_Ord[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(eqMapIter2_3_7)}
}), go__go_4_8_40}))}
})
_ = ordMapIter1__193435443_1_0
// TAST (Let): __local_var_2_19 shape=App(Other) bindingType=Any
__local_var_2_19 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd_0, "Eq0"), gopurs_runtime.Value{})
_ = __local_var_2_19
// TAST (Let): eqMap1__193435443_2_18 shape=Let(Abs(Let(LitRecord))) bindingType=Any
eqMap1__193435443_2_18 := gopurs_runtime.Func(func(dictEq1_3 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_4_21_41 gopurs_runtime.Value
_ = go__go_4_21_41
// FALLBACK TCO: isLoop=false len=1
go__go_4_21_41 = gopurs_runtime.Func2(func(a_5 gopurs_runtime.Value, b_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): v_7_22 shape=App(Var) bindingType=(ADT ["Data","Map","Internal","MapIterStep"] [(TypeVar k), (TypeVar v)])
v_7_22 := gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_IterNext[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Map_Internal_stepAsc(), a_5))
_ = v_7_22
var __t24 bool
{
if (v_7_22 != nil) {
// TAST (Let): v2_8_23 shape=App(Var) bindingType=(ADT ["Data","Map","Internal","MapIterStep"] [(TypeVar k), (TypeVar v)])
v2_8_23 := gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_IterNext[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Map_Internal_stepAsc(), b_6))
_ = v2_8_23
__t24 = ((v2_8_23 != nil)) && ((((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_19, "eq"), (v_7_22).V0, (v2_8_23).V0).IntVal) != (0)) && ((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq1_3, "eq"), (v_7_22).V1, (v2_8_23).V1).IntVal) != (0))) && ((gopurs_runtime.Apply2(go__go_4_21_41, (v_7_22).V2, (v2_8_23).V2).IntVal) != (0)))
goto end_branch_24
} else {

}
}
{
if (v_7_22 == nil) {
__t24 = true
goto end_branch_24
} else {

}
}
{
__t24 = func() bool { panic("Failed pattern match") }()
}
end_branch_24:
return gopurs_runtime.Bool(__t24)
})
// TAST (Let): eqMapIter2_4_20 shape=LitRecord bindingType=(ADT ["Data","Eq","Eq"] [(ADT ["Data","Map","Internal","MapIter"] [(TypeVar k), (TypeVar v)])])
eqMapIter2_4_20 := (&Constructor_Data_Eq_Eq[gopurs_runtime.Value]{1, go__go_4_21_41})
_ = eqMapIter2_4_20
return gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Data_Map_Internal_1444241223_3790796878((&Constructor_Data_Eq_Eq[*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]]{1, gopurs_runtime.Func2(func(xs_5 gopurs_runtime.Value, ys_6 gopurs_runtime.Value) gopurs_runtime.Value {
var __t29 bool
{
var __t_tag_25 *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](xs_5)
if (__t_tag_25 == nil) {
var __t_tag_26 *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](ys_6)
__t29 = (__t_tag_26 == nil)
goto end_branch_29
} else {

}
}
{
var __t_tag_27 *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](xs_5)
if (__t_tag_27 != nil) {
var __t_tag_28 *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](ys_6)
__t29 = ((__t_tag_28 != nil)) && ((((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(xs_5.UnsafePtr).V1) == ((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(ys_6.UnsafePtr).V1)) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(eqMapIter2_4_20.V0), gopurs_runtime.Value{Type: 9, IntVal: 2861335956, UnsafePtr: unsafe.Pointer((&Constructor_Data_Map_Internal_IterNode[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](xs_5), gopurs_runtime.Value{Type: 9, IntVal: 2509360378, UnsafePtr: unsafe.Pointer(nil)}}))}, gopurs_runtime.Value{Type: 9, IntVal: 2861335956, UnsafePtr: unsafe.Pointer((&Constructor_Data_Map_Internal_IterNode[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](ys_6), gopurs_runtime.Value{Type: 9, IntVal: 2509360378, UnsafePtr: unsafe.Pointer(nil)}}))}).IntVal) != (0)))
goto end_branch_29
} else {

}
}
{
__t29 = func() bool { panic("Failed pattern match") }()
}
end_branch_29:
return gopurs_runtime.Bool(__t29)
})})))}
})
_ = eqMap1__193435443_2_18
return gopurs_runtime.Func(func(dictOrd1_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): ordMapIter2_4_30 shape=App(Other) bindingType=(ADT ["Data","Ord","Ord"] [(ADT ["Data","Map","Internal","MapIter"] [(TypeVar k), (TypeVar v)])])
ordMapIter2_4_30 := gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](gopurs_runtime.Apply(ordMapIter1__193435443_1_0, dictOrd1_3))
_ = ordMapIter2_4_30
// TAST (Let): eqMap2_5_31 shape=App(Other) bindingType=(ADT ["Data","Eq","Eq"] [(ADT ["Data","Map","Internal","Map"] [(TypeVar k), (TypeVar v)])])
eqMap2_5_31 := Rebox_Data_Map_Internal_3790796878_1444241223(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](gopurs_runtime.Apply(eqMap1__193435443_2_18, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd1_3, "Eq0"), gopurs_runtime.Value{}))))
_ = eqMap2_5_31
return gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(Rebox_Data_Map_Internal_1910448679_4177771502((&Constructor_Data_Ord_Ord[*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(_dollar___unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Data_Map_Internal_1444241223_3790796878(eqMap2_5_31))}
}), gopurs_runtime.Func2(func(xs_6 gopurs_runtime.Value, ys_7 gopurs_runtime.Value) gopurs_runtime.Value {
var __t37 uint32
{
var __t_tag_34 *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](xs_6)
if (__t_tag_34 == nil) {
var __t36 uint32
{
var __t_tag_35 *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](ys_7)
if (__t_tag_35 == nil) {
__t36 = 902936544
goto end_branch_36
} else {

}
}
{
__t36 = 1527465420
}
end_branch_36:
__t37 = __t36
goto end_branch_37
} else {

}
}
{
var __t33 uint32
{
var __t_tag_32 *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](ys_7)
if (__t_tag_32 == nil) {
__t33 = 380165415
goto end_branch_33
} else {

}
}
{
__t33 = uint32(gopurs_runtime.Apply2(gopurs_runtime.Box(ordMapIter2_4_30.V1), gopurs_runtime.Value{Type: 9, IntVal: 2861335956, UnsafePtr: unsafe.Pointer((&Constructor_Data_Map_Internal_IterNode[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](xs_6), gopurs_runtime.Value{Type: 9, IntVal: 2509360378, UnsafePtr: unsafe.Pointer(nil)}}))}, gopurs_runtime.Value{Type: 9, IntVal: 2861335956, UnsafePtr: unsafe.Pointer((&Constructor_Data_Map_Internal_IterNode[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](ys_7), gopurs_runtime.Value{Type: 9, IntVal: 2509360378, UnsafePtr: unsafe.Pointer(nil)}}))}).IntVal)
}
end_branch_33:
__t37 = __t33
}
end_branch_37:
return gopurs_runtime.Value{Type: 9, IntVal: int64(__t37), UnsafePtr: nil}
})})))}
})
}

func Call_Data_Map_Internal_eq1Map(dictEq_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictEq_0 gopurs_runtime.Value = dictEq_0_loop
_ = dictEq_0
return gopurs_runtime.Value{Type: 9, IntVal: 1715248107, UnsafePtr: unsafe.Pointer(Rebox_Data_Map_Internal_3691144502_1766074591((&Constructor_Data_Eq_Eq1[*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(dictEq1_1 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_2_1_42 gopurs_runtime.Value
_ = go__go_2_1_42
// FALLBACK TCO: isLoop=false len=1
go__go_2_1_42 = gopurs_runtime.Func2(func(a_3 gopurs_runtime.Value, b_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): v_5_2 shape=App(Var) bindingType=(ADT ["Data","Map","Internal","MapIterStep"] [(TypeVar k), (TypeVar v)])
v_5_2 := gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_IterNext[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Map_Internal_stepAsc(), a_3))
_ = v_5_2
var __t4 bool
{
if (v_5_2 != nil) {
// TAST (Let): v2_6_3 shape=App(Var) bindingType=(ADT ["Data","Map","Internal","MapIterStep"] [(TypeVar k), (TypeVar v)])
v2_6_3 := gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_IterNext[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Map_Internal_stepAsc(), b_4))
_ = v2_6_3
__t4 = ((v2_6_3 != nil)) && ((((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq_0, "eq"), (v_5_2).V0, (v2_6_3).V0).IntVal) != (0)) && ((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq1_1, "eq"), (v_5_2).V1, (v2_6_3).V1).IntVal) != (0))) && ((gopurs_runtime.Apply2(go__go_2_1_42, (v_5_2).V2, (v2_6_3).V2).IntVal) != (0)))
goto end_branch_4
} else {

}
}
{
if (v_5_2 == nil) {
__t4 = true
goto end_branch_4
} else {

}
}
{
__t4 = func() bool { panic("Failed pattern match") }()
}
end_branch_4:
return gopurs_runtime.Bool(__t4)
})
// TAST (Let): eqMapIter2_2_0 shape=LitRecord bindingType=(ADT ["Data","Eq","Eq"] [(ADT ["Data","Map","Internal","MapIter"] [(TypeVar k), (TypeVar v)])])
eqMapIter2_2_0 := (&Constructor_Data_Eq_Eq[gopurs_runtime.Value]{1, go__go_2_1_42})
_ = eqMapIter2_2_0
return gopurs_runtime.Func2(func(xs_3 gopurs_runtime.Value, ys_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t9 bool
{
var __t_tag_5 *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](xs_3)
if (__t_tag_5 == nil) {
var __t_tag_6 *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](ys_4)
__t9 = (__t_tag_6 == nil)
goto end_branch_9
} else {

}
}
{
var __t_tag_7 *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](xs_3)
if (__t_tag_7 != nil) {
var __t_tag_8 *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](ys_4)
__t9 = ((__t_tag_8 != nil)) && ((((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(xs_3.UnsafePtr).V1) == ((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(ys_4.UnsafePtr).V1)) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(eqMapIter2_2_0.V0), gopurs_runtime.Value{Type: 9, IntVal: 2861335956, UnsafePtr: unsafe.Pointer((&Constructor_Data_Map_Internal_IterNode[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](xs_3), gopurs_runtime.Value{Type: 9, IntVal: 2509360378, UnsafePtr: unsafe.Pointer(nil)}}))}, gopurs_runtime.Value{Type: 9, IntVal: 2861335956, UnsafePtr: unsafe.Pointer((&Constructor_Data_Map_Internal_IterNode[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](ys_4), gopurs_runtime.Value{Type: 9, IntVal: 2509360378, UnsafePtr: unsafe.Pointer(nil)}}))}).IntVal) != (0)))
goto end_branch_9
} else {

}
}
{
__t9 = func() bool { panic("Failed pattern match") }()
}
end_branch_9:
return gopurs_runtime.Bool(__t9)
})
})})))}
}

func Call_Data_Map_Internal_ord1Map(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
// TAST (Let): __local_var_1_3 shape=App(Other) bindingType=Any
__local_var_1_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd_0, "Eq0"), gopurs_runtime.Value{})
_ = __local_var_1_3
// TAST (Let): eqMapIter1__193435443_1_2 shape=Let(Abs(LitRecord)) bindingType=Any
eqMapIter1__193435443_1_2 := gopurs_runtime.Func(func(dictEq1_2 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_3_4_43 gopurs_runtime.Value
_ = go__go_3_4_43
// FALLBACK TCO: isLoop=false len=1
go__go_3_4_43 = gopurs_runtime.Func2(func(a_4 gopurs_runtime.Value, b_5 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): v_6_5 shape=App(Var) bindingType=(ADT ["Data","Map","Internal","MapIterStep"] [(TypeVar k), (TypeVar v)])
v_6_5 := gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_IterNext[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Map_Internal_stepAsc(), a_4))
_ = v_6_5
var __t7 bool
{
if (v_6_5 != nil) {
// TAST (Let): v2_7_6 shape=App(Var) bindingType=(ADT ["Data","Map","Internal","MapIterStep"] [(TypeVar k), (TypeVar v)])
v2_7_6 := gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_IterNext[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Map_Internal_stepAsc(), b_5))
_ = v2_7_6
__t7 = ((v2_7_6 != nil)) && ((((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_3, "eq"), (v_6_5).V0, (v2_7_6).V0).IntVal) != (0)) && ((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq1_2, "eq"), (v_6_5).V1, (v2_7_6).V1).IntVal) != (0))) && ((gopurs_runtime.Apply2(go__go_3_4_43, (v_6_5).V2, (v2_7_6).V2).IntVal) != (0)))
goto end_branch_7
} else {

}
}
{
if (v_6_5 == nil) {
__t7 = true
goto end_branch_7
} else {

}
}
{
__t7 = func() bool { panic("Failed pattern match") }()
}
end_branch_7:
return gopurs_runtime.Bool(__t7)
})
return gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer((&Constructor_Data_Eq_Eq[gopurs_runtime.Value]{1, go__go_3_4_43}))}
})
_ = eqMapIter1__193435443_1_2
// TAST (Let): ordMapIter1__193435443_1_1 shape=Let(Abs(Let(LitRecord))) bindingType=Any
ordMapIter1__193435443_1_1 := gopurs_runtime.Func(func(dictOrd1_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): eqMapIter2_3_8 shape=App(Other) bindingType=(ADT ["Data","Eq","Eq"] [(ADT ["Data","Map","Internal","MapIter"] [(TypeVar k), (TypeVar v)])])
eqMapIter2_3_8 := gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](gopurs_runtime.Apply(eqMapIter1__193435443_1_2, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd1_2, "Eq0"), gopurs_runtime.Value{})))
_ = eqMapIter2_3_8
var Call_local_Data_Map_Internal_go__go_4_9_44 func(gopurs_runtime.Value, gopurs_runtime.Value) uint32
_ = Call_local_Data_Map_Internal_go__go_4_9_44
var go__go_4_9_44 gopurs_runtime.Value
_ = go__go_4_9_44
Call_local_Data_Map_Internal_go__go_4_9_44 = func(a_5_loop gopurs_runtime.Value, b_6_loop gopurs_runtime.Value) uint32 {
go__go_4_9_44:
for {
if false { continue go__go_4_9_44 }
var a_5 gopurs_runtime.Value = a_5_loop
_ = a_5
var b_6 gopurs_runtime.Value = b_6_loop
_ = b_6
// TAST (Let): v_7_10 shape=App(Var) bindingType=(ADT ["Data","Map","Internal","MapIterStep"] [(TypeVar k), (TypeVar v)])
v_7_10 := gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_IterNext[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Map_Internal_stepAsc(), b_6))
_ = v_7_10
// TAST (Let): v1_8_11 shape=App(Var) bindingType=(ADT ["Data","Map","Internal","MapIterStep"] [(TypeVar k), (TypeVar v)])
v1_8_11 := gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_IterNext[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Map_Internal_stepAsc(), a_5))
_ = v1_8_11
var __t18 uint32
{
if (v1_8_11 != nil) {
var __t16 uint32
{
if (v_7_10 != nil) {
// TAST (Let): v3_9_12 shape=App(Other) bindingType=(ADT ["Data","Ordering","Ordering"] [])
v3_9_12 := uint32(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictOrd_0, "compare"), (v1_8_11).V0, (v_7_10).V0).IntVal)
_ = v3_9_12
var __t15 uint32
{
if (v3_9_12 == 902936544) {
// TAST (Let): v4_10_13 shape=App(Other) bindingType=(ADT ["Data","Ordering","Ordering"] [])
v4_10_13 := uint32(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictOrd1_2, "compare"), (v1_8_11).V1, (v_7_10).V1).IntVal)
_ = v4_10_13
var __t14 uint32
{
if (v4_10_13 == 902936544) {
a_5_loop = (v1_8_11).V2
b_6_loop = (v_7_10).V2
continue go__go_4_9_44
__t14 = func() uint32 { panic("unreachable") }()
goto end_branch_14
} else {

}
}
{
__t14 = v4_10_13
}
end_branch_14:
__t15 = __t14
goto end_branch_15
} else {

}
}
{
__t15 = v3_9_12
}
end_branch_15:
__t16 = __t15
goto end_branch_16
} else {

}
}
{
if (v_7_10 == nil) {
__t16 = 380165415
goto end_branch_16
} else {

}
}
{
__t16 = func() uint32 { panic("Failed pattern match") }()
}
end_branch_16:
__t18 = __t16
goto end_branch_18
} else {

}
}
{
if (v1_8_11 == nil) {
var __t17 uint32
{
if (v_7_10 == nil) {
__t17 = 902936544
goto end_branch_17
} else {

}
}
{
__t17 = 1527465420
}
end_branch_17:
__t18 = __t17
goto end_branch_18
} else {

}
}
{
if (v_7_10 == nil) {
__t18 = 380165415
goto end_branch_18
} else {

}
}
{
__t18 = func() uint32 { panic("Failed pattern match") }()
}
end_branch_18:
return __t18
}
}
go__go_4_9_44 = gopurs_runtime.Func(func(a_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: int64(Call_local_Data_Map_Internal_go__go_4_9_44(a_5_loop_val, b_6_loop_val)), UnsafePtr: nil}
})
})
return gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer((&Constructor_Data_Ord_Ord[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(eqMapIter2_3_8)}
}), go__go_4_9_44}))}
})
_ = ordMapIter1__193435443_1_1
// TAST (Let): __local_var_2_20 shape=App(Other) bindingType=Any
__local_var_2_20 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd_0, "Eq0"), gopurs_runtime.Value{})
_ = __local_var_2_20
// TAST (Let): eqMap1__193435443_2_19 shape=Let(Abs(Let(LitRecord))) bindingType=Any
eqMap1__193435443_2_19 := gopurs_runtime.Func(func(dictEq1_3 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_4_22_45 gopurs_runtime.Value
_ = go__go_4_22_45
// FALLBACK TCO: isLoop=false len=1
go__go_4_22_45 = gopurs_runtime.Func2(func(a_5 gopurs_runtime.Value, b_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): v_7_23 shape=App(Var) bindingType=(ADT ["Data","Map","Internal","MapIterStep"] [(TypeVar k), (TypeVar v)])
v_7_23 := gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_IterNext[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Map_Internal_stepAsc(), a_5))
_ = v_7_23
var __t25 bool
{
if (v_7_23 != nil) {
// TAST (Let): v2_8_24 shape=App(Var) bindingType=(ADT ["Data","Map","Internal","MapIterStep"] [(TypeVar k), (TypeVar v)])
v2_8_24 := gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_IterNext[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Map_Internal_stepAsc(), b_6))
_ = v2_8_24
__t25 = ((v2_8_24 != nil)) && ((((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_20, "eq"), (v_7_23).V0, (v2_8_24).V0).IntVal) != (0)) && ((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq1_3, "eq"), (v_7_23).V1, (v2_8_24).V1).IntVal) != (0))) && ((gopurs_runtime.Apply2(go__go_4_22_45, (v_7_23).V2, (v2_8_24).V2).IntVal) != (0)))
goto end_branch_25
} else {

}
}
{
if (v_7_23 == nil) {
__t25 = true
goto end_branch_25
} else {

}
}
{
__t25 = func() bool { panic("Failed pattern match") }()
}
end_branch_25:
return gopurs_runtime.Bool(__t25)
})
// TAST (Let): eqMapIter2_4_21 shape=LitRecord bindingType=(ADT ["Data","Eq","Eq"] [(ADT ["Data","Map","Internal","MapIter"] [(TypeVar k), (TypeVar v)])])
eqMapIter2_4_21 := (&Constructor_Data_Eq_Eq[gopurs_runtime.Value]{1, go__go_4_22_45})
_ = eqMapIter2_4_21
return gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Data_Map_Internal_1444241223_3790796878((&Constructor_Data_Eq_Eq[*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]]{1, gopurs_runtime.Func2(func(xs_5 gopurs_runtime.Value, ys_6 gopurs_runtime.Value) gopurs_runtime.Value {
var __t30 bool
{
var __t_tag_26 *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](xs_5)
if (__t_tag_26 == nil) {
var __t_tag_27 *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](ys_6)
__t30 = (__t_tag_27 == nil)
goto end_branch_30
} else {

}
}
{
var __t_tag_28 *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](xs_5)
if (__t_tag_28 != nil) {
var __t_tag_29 *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](ys_6)
__t30 = ((__t_tag_29 != nil)) && ((((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(xs_5.UnsafePtr).V1) == ((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(ys_6.UnsafePtr).V1)) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(eqMapIter2_4_21.V0), gopurs_runtime.Value{Type: 9, IntVal: 2861335956, UnsafePtr: unsafe.Pointer((&Constructor_Data_Map_Internal_IterNode[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](xs_5), gopurs_runtime.Value{Type: 9, IntVal: 2509360378, UnsafePtr: unsafe.Pointer(nil)}}))}, gopurs_runtime.Value{Type: 9, IntVal: 2861335956, UnsafePtr: unsafe.Pointer((&Constructor_Data_Map_Internal_IterNode[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](ys_6), gopurs_runtime.Value{Type: 9, IntVal: 2509360378, UnsafePtr: unsafe.Pointer(nil)}}))}).IntVal) != (0)))
goto end_branch_30
} else {

}
}
{
__t30 = func() bool { panic("Failed pattern match") }()
}
end_branch_30:
return gopurs_runtime.Bool(__t30)
})})))}
})
_ = eqMap1__193435443_2_19
// TAST (Let): ordMap1__193435443_1_0 shape=Let(Let(Abs(Let(Let(LitRecord))))) bindingType=Any
ordMap1__193435443_1_0 := gopurs_runtime.Func(func(dictOrd1_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): ordMapIter2_4_31 shape=App(Other) bindingType=(ADT ["Data","Ord","Ord"] [(ADT ["Data","Map","Internal","MapIter"] [(TypeVar k), (TypeVar v)])])
ordMapIter2_4_31 := gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](gopurs_runtime.Apply(ordMapIter1__193435443_1_1, dictOrd1_3))
_ = ordMapIter2_4_31
// TAST (Let): eqMap2_5_32 shape=App(Other) bindingType=(ADT ["Data","Eq","Eq"] [(ADT ["Data","Map","Internal","Map"] [(TypeVar k), (TypeVar v)])])
eqMap2_5_32 := Rebox_Data_Map_Internal_3790796878_1444241223(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](gopurs_runtime.Apply(eqMap1__193435443_2_19, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd1_3, "Eq0"), gopurs_runtime.Value{}))))
_ = eqMap2_5_32
return gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(Rebox_Data_Map_Internal_1910448679_4177771502((&Constructor_Data_Ord_Ord[*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(_dollar___unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Data_Map_Internal_1444241223_3790796878(eqMap2_5_32))}
}), gopurs_runtime.Func2(func(xs_6 gopurs_runtime.Value, ys_7 gopurs_runtime.Value) gopurs_runtime.Value {
var __t38 uint32
{
var __t_tag_35 *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](xs_6)
if (__t_tag_35 == nil) {
var __t37 uint32
{
var __t_tag_36 *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](ys_7)
if (__t_tag_36 == nil) {
__t37 = 902936544
goto end_branch_37
} else {

}
}
{
__t37 = 1527465420
}
end_branch_37:
__t38 = __t37
goto end_branch_38
} else {

}
}
{
var __t34 uint32
{
var __t_tag_33 *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](ys_7)
if (__t_tag_33 == nil) {
__t34 = 380165415
goto end_branch_34
} else {

}
}
{
__t34 = uint32(gopurs_runtime.Apply2(gopurs_runtime.Box(ordMapIter2_4_31.V1), gopurs_runtime.Value{Type: 9, IntVal: 2861335956, UnsafePtr: unsafe.Pointer((&Constructor_Data_Map_Internal_IterNode[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](xs_6), gopurs_runtime.Value{Type: 9, IntVal: 2509360378, UnsafePtr: unsafe.Pointer(nil)}}))}, gopurs_runtime.Value{Type: 9, IntVal: 2861335956, UnsafePtr: unsafe.Pointer((&Constructor_Data_Map_Internal_IterNode[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](ys_7), gopurs_runtime.Value{Type: 9, IntVal: 2509360378, UnsafePtr: unsafe.Pointer(nil)}}))}).IntVal)
}
end_branch_34:
__t38 = __t34
}
end_branch_38:
return gopurs_runtime.Value{Type: 9, IntVal: int64(__t38), UnsafePtr: nil}
})})))}
})
_ = ordMap1__193435443_1_0
// TAST (Let): __local_var_2_40 shape=App(Other) bindingType=Any
__local_var_2_40 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd_0, "Eq0"), gopurs_runtime.Value{})
_ = __local_var_2_40
// TAST (Let): eq1Map1_2_39 shape=Let(LitRecord) bindingType=(ADT ["Data","Eq","Eq1"] [(ADT ["Data","Map","Internal","Map"] [(TypeVar k)])])
eq1Map1_2_39 := (&Constructor_Data_Eq_Eq1[*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(dictEq1_3 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_4_42_46 gopurs_runtime.Value
_ = go__go_4_42_46
// FALLBACK TCO: isLoop=false len=1
go__go_4_42_46 = gopurs_runtime.Func2(func(a_5 gopurs_runtime.Value, b_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): v_7_43 shape=App(Var) bindingType=(ADT ["Data","Map","Internal","MapIterStep"] [(TypeVar k), (TypeVar v)])
v_7_43 := gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_IterNext[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Map_Internal_stepAsc(), a_5))
_ = v_7_43
var __t45 bool
{
if (v_7_43 != nil) {
// TAST (Let): v2_8_44 shape=App(Var) bindingType=(ADT ["Data","Map","Internal","MapIterStep"] [(TypeVar k), (TypeVar v)])
v2_8_44 := gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_IterNext[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Map_Internal_stepAsc(), b_6))
_ = v2_8_44
__t45 = ((v2_8_44 != nil)) && ((((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_40, "eq"), (v_7_43).V0, (v2_8_44).V0).IntVal) != (0)) && ((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq1_3, "eq"), (v_7_43).V1, (v2_8_44).V1).IntVal) != (0))) && ((gopurs_runtime.Apply2(go__go_4_42_46, (v_7_43).V2, (v2_8_44).V2).IntVal) != (0)))
goto end_branch_45
} else {

}
}
{
if (v_7_43 == nil) {
__t45 = true
goto end_branch_45
} else {

}
}
{
__t45 = func() bool { panic("Failed pattern match") }()
}
end_branch_45:
return gopurs_runtime.Bool(__t45)
})
// TAST (Let): eqMapIter2_4_41 shape=LitRecord bindingType=(ADT ["Data","Eq","Eq"] [(ADT ["Data","Map","Internal","MapIter"] [(TypeVar k), (TypeVar v)])])
eqMapIter2_4_41 := (&Constructor_Data_Eq_Eq[gopurs_runtime.Value]{1, go__go_4_42_46})
_ = eqMapIter2_4_41
return gopurs_runtime.Func2(func(xs_5 gopurs_runtime.Value, ys_6 gopurs_runtime.Value) gopurs_runtime.Value {
var __t50 bool
{
var __t_tag_46 *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](xs_5)
if (__t_tag_46 == nil) {
var __t_tag_47 *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](ys_6)
__t50 = (__t_tag_47 == nil)
goto end_branch_50
} else {

}
}
{
var __t_tag_48 *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](xs_5)
if (__t_tag_48 != nil) {
var __t_tag_49 *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](ys_6)
__t50 = ((__t_tag_49 != nil)) && ((((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(xs_5.UnsafePtr).V1) == ((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(ys_6.UnsafePtr).V1)) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(eqMapIter2_4_41.V0), gopurs_runtime.Value{Type: 9, IntVal: 2861335956, UnsafePtr: unsafe.Pointer((&Constructor_Data_Map_Internal_IterNode[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](xs_5), gopurs_runtime.Value{Type: 9, IntVal: 2509360378, UnsafePtr: unsafe.Pointer(nil)}}))}, gopurs_runtime.Value{Type: 9, IntVal: 2861335956, UnsafePtr: unsafe.Pointer((&Constructor_Data_Map_Internal_IterNode[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](ys_6), gopurs_runtime.Value{Type: 9, IntVal: 2509360378, UnsafePtr: unsafe.Pointer(nil)}}))}).IntVal) != (0)))
goto end_branch_50
} else {

}
}
{
__t50 = func() bool { panic("Failed pattern match") }()
}
end_branch_50:
return gopurs_runtime.Bool(__t50)
})
})})
_ = eq1Map1_2_39
return gopurs_runtime.Value{Type: 9, IntVal: 1632188299, UnsafePtr: unsafe.Pointer(Rebox_Data_Map_Internal_849406934_3985601471((&Constructor_Data_Ord_Ord1[*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(_dollar___unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1715248107, UnsafePtr: unsafe.Pointer(Rebox_Data_Map_Internal_3691144502_1766074591(eq1Map1_2_39))}
}), gopurs_runtime.Func(func(dictOrd1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(gopurs_runtime.Apply(ordMap1__193435443_1_0, dictOrd1_3), "compare")
})})))}
}

func Call_Data_Map_Internal_fromFoldable(dictOrd_0_loop *Constructor_Data_Ord_Ord[gopurs_runtime.Value], dictFoldable_1_loop *Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictOrd_0 *Constructor_Data_Ord_Ord[gopurs_runtime.Value] = dictOrd_0_loop
_ = dictOrd_0
var dictFoldable_1 *Constructor_Data_Foldable_Foldable[gopurs_runtime.Value] = dictFoldable_1_loop
_ = dictFoldable_1
return gopurs_runtime.Apply2(gopurs_runtime.Box(dictFoldable_1.V1), gopurs_runtime.Func2(func(m_2 gopurs_runtime.Value, v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(Call_Data_Map_Internal_insert(dictOrd_0, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V0, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V1), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](m_2))})))}
}), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(nil))})
}

func Call_Data_Map_Internal_fromFoldableWith(dictOrd_0_loop *Constructor_Data_Ord_Ord[gopurs_runtime.Value], dictFoldable_1_loop *Constructor_Data_Foldable_Foldable[gopurs_runtime.Value], f_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 *Constructor_Data_Ord_Ord[gopurs_runtime.Value] = dictOrd_0_loop
_ = dictOrd_0
var dictFoldable_1 *Constructor_Data_Foldable_Foldable[gopurs_runtime.Value] = dictFoldable_1_loop
_ = dictFoldable_1
var f_2 gopurs_runtime.Value = f_2_loop
_ = f_2
return gopurs_runtime.Apply2(gopurs_runtime.Box(dictFoldable_1.V1), gopurs_runtime.Func2(func(m_3 gopurs_runtime.Value, v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(Call_Data_Map_Internal_insertWith(dictOrd_0, gopurs_runtime.Func2(func(b_5 gopurs_runtime.Value, a_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_2, a_6, b_5)
}), (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V0, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V1), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](m_3))})))}
}), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(nil))})
}

func Call_Data_Map_Internal_fromFoldableWithIndex(dictOrd_0_loop *Constructor_Data_Ord_Ord[gopurs_runtime.Value], dictFoldableWithIndex_1_loop *Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var dictOrd_0 *Constructor_Data_Ord_Ord[gopurs_runtime.Value] = dictOrd_0_loop
_ = dictOrd_0
var dictFoldableWithIndex_1 *Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value] = dictFoldableWithIndex_1_loop
_ = dictFoldableWithIndex_1
return gopurs_runtime.Apply2(gopurs_runtime.Box(dictFoldableWithIndex_1.V2), gopurs_runtime.Func3(func(k_2 gopurs_runtime.Value, m_3 gopurs_runtime.Value, v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(Call_Data_Map_Internal_insert(dictOrd_0, k_2, v_4), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](m_3))})))}
}), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(nil))})
}

func Call_Data_Map_Internal_monoidSemigroupMap(_dollar___unused_0_loop gopurs_runtime.Value, dictOrd_1_loop gopurs_runtime.Value, dictSemigroup_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var _dollar___unused_0 gopurs_runtime.Value = _dollar___unused_0_loop
_ = _dollar___unused_0
var dictOrd_1 gopurs_runtime.Value = dictOrd_1_loop
_ = dictOrd_1
var dictSemigroup_2 gopurs_runtime.Value = dictSemigroup_2_loop
_ = dictSemigroup_2
// TAST (Let): compare_3_1 shape=Other bindingType=(Func [(TypeVar k), (TypeVar k)] (ADT ["Data","Ordering","Ordering"] []))
compare_3_1 := gopurs_runtime.RecordGet(dictOrd_1, "compare")
_ = compare_3_1
// TAST (Let): __local_var_4_2 shape=Other bindingType=(Func [(TypeVar v), (TypeVar v)] (TypeVar v))
__local_var_4_2 := gopurs_runtime.RecordGet(dictSemigroup_2, "append")
_ = __local_var_4_2
// TAST (Let): semigroupMap3_3_0 shape=LitRecord bindingType=(ADT ["Data","Semigroup","Semigroup"] [(ADT ["Data","Map","Internal","Map"] [(TypeVar k), (TypeVar v)])])
semigroupMap3_3_0 := (&Constructor_Data_Semigroup_Semigroup[*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]]{1, gopurs_runtime.Func2(func(m1_5 gopurs_runtime.Value, m2_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeUnionWith(), compare_3_1, __local_var_4_2, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](m1_5))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](m2_6))})))}
})})
_ = semigroupMap3_3_0
return gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(Rebox_Data_Map_Internal_1291127239_1201789390((&Constructor_Data_Monoid_Monoid[*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(_dollar___unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer(Rebox_Data_Map_Internal_3854229351_4179793454(semigroupMap3_3_0))}
}), gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(nil))})})))}
}

func Call_Data_Map_Internal_submap(dictOrd_0_loop *Constructor_Data_Ord_Ord[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictOrd_0 *Constructor_Data_Ord_Ord[gopurs_runtime.Value] = dictOrd_0_loop
_ = dictOrd_0
// TAST (Let): compare_1_1 shape=Other bindingType=(Func [(TypeVar k), (TypeVar k)] (ADT ["Data","Ordering","Ordering"] []))
compare_1_1 := gopurs_runtime.Box(dictOrd_0.V1)
_ = compare_1_1
// TAST (Let): union1_1_0 shape=Let(Abs(Abs(UncurriedApp(Var)))) bindingType=(Func [(ADT ["Data","Map","Internal","Map"] [(TypeVar k), (TypeVar v)]), (ADT ["Data","Map","Internal","Map"] [(TypeVar k), (TypeVar v)])] (ADT ["Data","Map","Internal","Map"] [(TypeVar k), (TypeVar v)]))
union1_1_0 := gopurs_runtime.Func2(func(m1_2 gopurs_runtime.Value, m2_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeUnionWith(), compare_1_1, Get_Data_Function_go__const(), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](m1_2))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](m2_3))})))}
})
_ = union1_1_0
return gopurs_runtime.Func2(func(kmin_2 gopurs_runtime.Value, kmax_3 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Map_Internal_foldSubmapBy(dictOrd_0, union1_1_0, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(nil))}, gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](kmin_2), gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](kmax_3), Get_Data_Map_Internal_singleton())
})
}

func Call_Data_Map_Internal_unions(dictOrd_0_loop *Constructor_Data_Ord_Ord[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictOrd_0 *Constructor_Data_Ord_Ord[gopurs_runtime.Value] = dictOrd_0_loop
_ = dictOrd_0
// TAST (Let): compare_1_1 shape=Other bindingType=(Func [(TypeVar k), (TypeVar k)] (ADT ["Data","Ordering","Ordering"] []))
compare_1_1 := gopurs_runtime.Box(dictOrd_0.V1)
_ = compare_1_1
// TAST (Let): union1_1_0 shape=Let(Abs(Abs(UncurriedApp(Var)))) bindingType=(Func [(ADT ["Data","Map","Internal","Map"] [(TypeVar k), (TypeVar v)]), (ADT ["Data","Map","Internal","Map"] [(TypeVar k), (TypeVar v)])] (ADT ["Data","Map","Internal","Map"] [(TypeVar k), (TypeVar v)]))
union1_1_0 := gopurs_runtime.Func2(func(m1_2 gopurs_runtime.Value, m2_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeUnionWith(), compare_1_1, Get_Data_Function_go__const(), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](m1_2))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](m2_3))})))}
})
_ = union1_1_0
return gopurs_runtime.Func(func(dictFoldable_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFoldable_2, "foldl"), union1_1_0, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(nil))})
})
}

func Call_Data_Map_Internal_difference(dictOrd_0_loop *Constructor_Data_Ord_Ord[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictOrd_0 *Constructor_Data_Ord_Ord[gopurs_runtime.Value] = dictOrd_0_loop
_ = dictOrd_0
// TAST (Let): compare_1_0 shape=Other bindingType=(Func [(TypeVar k), (TypeVar k)] (ADT ["Data","Ordering","Ordering"] []))
compare_1_0 := gopurs_runtime.Box(dictOrd_0.V1)
_ = compare_1_0
return gopurs_runtime.Func2(func(m1_2 gopurs_runtime.Value, m2_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp3(Get_Data_Map_Internal_unsafeDifference(), compare_1_0, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](m1_2))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](m2_3))})))}
})
}

func Call_Data_Map_Internal_go__delete(dictOrd_0_loop *Constructor_Data_Ord_Ord[gopurs_runtime.Value], k_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 *Constructor_Data_Ord_Ord[gopurs_runtime.Value] = dictOrd_0_loop
_ = dictOrd_0
var k_1 gopurs_runtime.Value = k_1_loop
_ = k_1
var go__go_2_0_47 gopurs_runtime.Value
_ = go__go_2_0_47
// FALLBACK TCO: isLoop=false len=1
go__go_2_0_47 = gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t5 *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]
{
var __t_tag_1 *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](v_3)
if (__t_tag_1 == nil) {
__t5 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(nil))})
goto end_branch_5
} else {

}
}
{
var __t_tag_2 *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](v_3)
if (__t_tag_2 != nil) {
// TAST (Let): v1_4_3 shape=App(Other) bindingType=(ADT ["Data","Ordering","Ordering"] [])
v1_4_3 := uint32(gopurs_runtime.Apply2(gopurs_runtime.Box(dictOrd_0.V1), k_1, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V2).IntVal)
_ = v1_4_3
var __t4 *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]
{
if (v1_4_3 == 1527465420) {
__t4 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeBalancedNode(), (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(go__go_2_0_47, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V4)})))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V5)}))
goto end_branch_4
} else {

}
}
{
if (v1_4_3 == 380165415) {
__t4 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeBalancedNode(), (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V4)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(go__go_2_0_47, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V5)})))}))
goto end_branch_4
} else {

}
}
{
if (v1_4_3 == 902936544) {
__t4 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp2(Get_Data_Map_Internal_unsafeJoinNodes(), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V4)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V5)}))
goto end_branch_4
} else {

}
}
{
__t4 = func() *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value] { panic("Failed pattern match") }()
}
end_branch_4:
__t5 = __t4
goto end_branch_5
} else {

}
}
{
__t5 = func() *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value] { panic("Failed pattern match") }()
}
end_branch_5:
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(__t5)}
})
return go__go_2_0_47
}

func Call_Data_Map_Internal_checkValid(dictOrd_0_loop *Constructor_Data_Ord_Ord[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictOrd_0 *Constructor_Data_Ord_Ord[gopurs_runtime.Value] = dictOrd_0_loop
_ = dictOrd_0
var go__go_1_0_48 gopurs_runtime.Value
_ = go__go_1_0_48
// FALLBACK TCO: isLoop=false len=1
go__go_1_0_48 = gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t25 bool
{
var __t_tag_1 *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](v_2)
if (__t_tag_1 == nil) {
__t25 = true
goto end_branch_25
} else {

}
}
{
var __t_tag_2 *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](v_2)
if (__t_tag_2 != nil) {
var __t24 bool
{
var __t_tag_3 *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value] = (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V4
if (__t_tag_3 == nil) {
var __t10 bool
{
var __t_tag_4 *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value] = (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V5
if (__t_tag_4 == nil) {
__t10 = true
goto end_branch_10
} else {

}
}
{
var __t_tag_5 *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value] = (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V5
if (__t_tag_5 != nil) {
var __t_and_9 bool = false
if ((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V0) == (int64(2)) {

var __t_and_8 bool = false
if (((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V5).V0) == (int64(1)) {

var __t_and_7 bool = false
if ((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V1) > (((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V5).V1) {

var __t_tag_6 uint32 = uint32(gopurs_runtime.Apply2(gopurs_runtime.Box(dictOrd_0.V1), ((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V5).V2, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V2).IntVal)
__t_and_7 = ((uint32(__t_tag_6) == 380165415)) && ((gopurs_runtime.Apply(go__go_1_0_48, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V5)}).IntVal) != (0))
}
__t_and_8 = __t_and_7
}
__t_and_9 = __t_and_8
}
__t10 = __t_and_9
goto end_branch_10
} else {

}
}
{
__t10 = func() bool { panic("Failed pattern match") }()
}
end_branch_10:
__t24 = __t10
goto end_branch_24
} else {

}
}
{
var __t_tag_11 *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value] = (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V4
if (__t_tag_11 != nil) {
var __t23 bool
{
var __t_tag_12 *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value] = (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V5
if (__t_tag_12 == nil) {
var __t_and_16 bool = false
if ((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V0) == (int64(2)) {

var __t_and_15 bool = false
if (((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V4).V0) == (int64(1)) {

var __t_and_14 bool = false
if ((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V1) > (((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V4).V1) {

var __t_tag_13 uint32 = uint32(gopurs_runtime.Apply2(gopurs_runtime.Box(dictOrd_0.V1), ((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V4).V2, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V2).IntVal)
__t_and_14 = ((uint32(__t_tag_13) == 1527465420)) && ((gopurs_runtime.Apply(go__go_1_0_48, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V4)}).IntVal) != (0))
}
__t_and_15 = __t_and_14
}
__t_and_16 = __t_and_15
}
__t23 = __t_and_16
goto end_branch_23
} else {

}
}
{
var __t_tag_17 *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value] = (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V5
if (__t_tag_17 != nil) {
var __t_and_22 bool = false
if ((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V0) > (((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V5).V0) {

var __t_tag_18 uint32 = uint32(gopurs_runtime.Apply2(gopurs_runtime.Box(dictOrd_0.V1), ((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V5).V2, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V2).IntVal)
var __t_and_21 bool = false
if (uint32(__t_tag_18) == 380165415) {

var __t_and_20 bool = false
if ((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V0) > (((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V4).V0) {

var __t_tag_19 uint32 = uint32(gopurs_runtime.Apply2(gopurs_runtime.Box(dictOrd_0.V1), ((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V4).V2, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V2).IntVal)
__t_and_20 = ((uint32(__t_tag_19) == 1527465420)) && ((((((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V5).V0) - (((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V4).V0)) < (int64(2))) && (((((((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V5).V1) + (((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V4).V1)) + (int64(1))) == ((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V1)) && (((gopurs_runtime.Apply(go__go_1_0_48, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V4)}).IntVal) != (0)) && ((gopurs_runtime.Apply(go__go_1_0_48, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V5)}).IntVal) != (0)))))
}
__t_and_21 = __t_and_20
}
__t_and_22 = __t_and_21
}
__t23 = __t_and_22
goto end_branch_23
} else {

}
}
{
__t23 = func() bool { panic("Failed pattern match") }()
}
end_branch_23:
__t24 = __t23
goto end_branch_24
} else {

}
}
{
__t24 = func() bool { panic("Failed pattern match") }()
}
end_branch_24:
__t25 = __t24
goto end_branch_25
} else {

}
}
{
__t25 = func() bool { panic("Failed pattern match") }()
}
end_branch_25:
return gopurs_runtime.Bool(__t25)
})
return go__go_1_0_48
}

func Call_Data_Map_Internal_catMaybes(dictOrd_0_loop *Constructor_Data_Ord_Ord[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictOrd_0 *Constructor_Data_Ord_Ord[gopurs_runtime.Value] = dictOrd_0_loop
_ = dictOrd_0
return Call_Data_Map_Internal_mapMaybeWithKey(dictOrd_0, gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Category_Category[gopurs_runtime.Value]](Get_Control_Category_categoryFn()).V1)
}))
}

func Call_Data_Map_Internal_applyMap(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
// TAST (Let): compare_1_0 shape=Other bindingType=(Func [(TypeVar k), (TypeVar k)] (ADT ["Data","Ordering","Ordering"] []))
compare_1_0 := gopurs_runtime.RecordGet(dictOrd_0, "compare")
_ = compare_1_0
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(Rebox_Data_Map_Internal_1628034448_3741347833((&Constructor_Control_Apply_Apply[*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(_dollar___unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(Rebox_Data_Map_Internal_3281783655_2812149806(Rebox_Data_Map_Internal_2812149806_3281783655(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Data_Map_Internal_functorMap()))))}
}), gopurs_runtime.Func2(func(m1_2 gopurs_runtime.Value, m2_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeIntersectionWith(), compare_1_0, gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Category_Category[gopurs_runtime.Value]](Get_Control_Category_categoryFn()).V1), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](m1_2))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](m2_3))})))}
})})))}
}

func Call_Data_Map_Internal_bindMap(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
// TAST (Let): compare_1_1 shape=Other bindingType=(Func [(TypeVar k), (TypeVar k)] (ADT ["Data","Ordering","Ordering"] []))
compare_1_1 := gopurs_runtime.RecordGet(dictOrd_0, "compare")
_ = compare_1_1
// TAST (Let): applyMap1_1_0 shape=LitRecord bindingType=(ADT ["Control","Apply","Apply"] [(ADT ["Data","Map","Internal","Map"] [(TypeVar k)])])
applyMap1_1_0 := (&Constructor_Control_Apply_Apply[*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(_dollar___unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(Rebox_Data_Map_Internal_3281783655_2812149806(Rebox_Data_Map_Internal_2812149806_3281783655(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Data_Map_Internal_functorMap()))))}
}), gopurs_runtime.Func2(func(m1_2 gopurs_runtime.Value, m2_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeIntersectionWith(), compare_1_1, gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Category_Category[gopurs_runtime.Value]](Get_Control_Category_categoryFn()).V1), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](m1_2))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](m2_3))})))}
})})
_ = applyMap1_1_0
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(Rebox_Data_Map_Internal_1860323088_2748095225((&Constructor_Control_Bind_Bind[*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(_dollar___unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(Rebox_Data_Map_Internal_1628034448_3741347833(applyMap1_1_0))}
}), gopurs_runtime.Func2(func(m_2 gopurs_runtime.Value, f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(Call_Data_Map_Internal_mapMaybeWithKey(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](dictOrd_0), gopurs_runtime.Func(func(k_4 gopurs_runtime.Value) gopurs_runtime.Value {
var Call_local_Data_Map_Internal_go__go_5_2_49 func(*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]) *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
_ = Call_local_Data_Map_Internal_go__go_5_2_49
var go__go_5_2_49 gopurs_runtime.Value
_ = go__go_5_2_49
Call_local_Data_Map_Internal_go__go_5_2_49 = func(v_6_loop *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]) *Constructor_Data_Maybe_Just[gopurs_runtime.Value] {
go__go_5_2_49:
for {
if false { continue go__go_5_2_49 }
var v_6 *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value] = v_6_loop
_ = v_6
var __t5 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
if (v_6 == nil) {
__t5 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil))})
goto end_branch_5
} else {

}
}
{
if (v_6 != nil) {
// TAST (Let): v1_7_3 shape=App(Other) bindingType=(ADT ["Data","Ordering","Ordering"] [])
v1_7_3 := uint32(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictOrd_0, "compare"), k_4, (v_6).V2).IntVal)
_ = v1_7_3
var __t4 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
if (v1_7_3 == 1527465420) {
v_6_loop = (v_6).V4
continue go__go_5_2_49
__t4 = func() *Constructor_Data_Maybe_Just[gopurs_runtime.Value] { panic("unreachable") }()
goto end_branch_4
} else {

}
}
{
if (v1_7_3 == 380165415) {
v_6_loop = (v_6).V5
continue go__go_5_2_49
__t4 = func() *Constructor_Data_Maybe_Just[gopurs_runtime.Value] { panic("unreachable") }()
goto end_branch_4
} else {

}
}
{
if (v1_7_3 == 902936544) {
__t4 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{(v_6).V3, true}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())
goto end_branch_4
} else {

}
}
{
__t4 = func() *Constructor_Data_Maybe_Just[gopurs_runtime.Value] { panic("Failed pattern match") }()
}
end_branch_4:
__t5 = __t4
goto end_branch_5
} else {

}
}
{
__t5 = func() *Constructor_Data_Maybe_Just[gopurs_runtime.Value] { panic("Failed pattern match") }()
}
end_branch_5:
return __t5
}
}
go__go_5_2_49 = gopurs_runtime.Func(func(v_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_local_Data_Map_Internal_go__go_5_2_49(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](v_6_loop_val)))}
})
return gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_local_Data_Map_Internal_go__go_5_2_49(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(f_3, x_6))))}
})
})), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](m_2))})))}
})})))}
}

func Call_Data_Map_Internal_anyWithKey(predicate_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var predicate_0 gopurs_runtime.Value = predicate_0_loop
_ = predicate_0
var go__go_1_0_50 gopurs_runtime.Value
_ = go__go_1_0_50
// FALLBACK TCO: isLoop=false len=1
go__go_1_0_50 = gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t3 bool
{
var __t_tag_1 *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](v_2)
if (__t_tag_1 == nil) {
__t3 = false
goto end_branch_3
} else {

}
}
{
var __t_tag_2 *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](v_2)
if (__t_tag_2 != nil) {
__t3 = ((gopurs_runtime.Apply2(predicate_0, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V3).IntVal) != (0)) || (((gopurs_runtime.Apply(go__go_1_0_50, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V4)}).IntVal) != (0)) || ((gopurs_runtime.Apply(go__go_1_0_50, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V5)}).IntVal) != (0)))
goto end_branch_3
} else {

}
}
{
__t3 = func() bool { panic("Failed pattern match") }()
}
end_branch_3:
return gopurs_runtime.Bool(__t3)
})
return go__go_1_0_50
}

func Call_Data_Map_Internal_any(predicate_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var predicate_0 gopurs_runtime.Value = predicate_0_loop
_ = predicate_0
var go__go_1_0_51 gopurs_runtime.Value
_ = go__go_1_0_51
// FALLBACK TCO: isLoop=false len=1
go__go_1_0_51 = gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t3 bool
{
var __t_tag_1 *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](v_2)
if (__t_tag_1 == nil) {
__t3 = false
goto end_branch_3
} else {

}
}
{
var __t_tag_2 *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](v_2)
if (__t_tag_2 != nil) {
__t3 = ((gopurs_runtime.Apply(predicate_0, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V3).IntVal) != (0)) || (((gopurs_runtime.Apply(go__go_1_0_51, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V4)}).IntVal) != (0)) || ((gopurs_runtime.Apply(go__go_1_0_51, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V5)}).IntVal) != (0)))
goto end_branch_3
} else {

}
}
{
__t3 = func() bool { panic("Failed pattern match") }()
}
end_branch_3:
return gopurs_runtime.Bool(__t3)
})
return go__go_1_0_51
}

func Call_Data_Map_Internal_alter(dictOrd_0_loop *Constructor_Data_Ord_Ord[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictOrd_0 *Constructor_Data_Ord_Ord[gopurs_runtime.Value] = dictOrd_0_loop
_ = dictOrd_0
// TAST (Let): compare_1_0 shape=Other bindingType=(Func [(TypeVar k), (TypeVar k)] (ADT ["Data","Ordering","Ordering"] []))
compare_1_0 := gopurs_runtime.Box(dictOrd_0.V1)
_ = compare_1_0
return gopurs_runtime.Func3(func(f_2 gopurs_runtime.Value, k_3 gopurs_runtime.Value, m_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): v_5_1 shape=UncurriedApp(Var) bindingType=(ADT ["Data","Map","Internal","Split"] [(TypeVar k), (TypeVar v)])
v_5_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Split[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp3(Get_Data_Map_Internal_unsafeSplit(), compare_1_0, k_3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](m_4))}))
_ = v_5_1
// TAST (Let): v2_6_2 shape=App(Other) bindingType=(ADT ["Data","Maybe","Maybe"] [(TypeVar v)])
v2_6_2 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(f_2, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((v_5_1).V0)}))
_ = v2_6_2
var __t3 *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]
{
if (v2_6_2 == nil) {
__t3 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp2(Get_Data_Map_Internal_unsafeJoinNodes(), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((v_5_1).V1)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((v_5_1).V2)}))
goto end_branch_3
} else {

}
}
{
if (v2_6_2 != nil) {
__t3 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeBalancedNode(), k_3, (v2_6_2).V0, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((v_5_1).V1)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((v_5_1).V2)}))
goto end_branch_3
} else {

}
}
{
__t3 = func() *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value] { panic("Failed pattern match") }()
}
end_branch_3:
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(__t3)}
})
}

func Call_Data_Map_Internal_altMap(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
// TAST (Let): compare_1_0 shape=Other bindingType=(Func [(TypeVar k), (TypeVar k)] (ADT ["Data","Ordering","Ordering"] []))
compare_1_0 := gopurs_runtime.RecordGet(dictOrd_0, "compare")
_ = compare_1_0
return gopurs_runtime.Value{Type: 9, IntVal: 4060500237, UnsafePtr: unsafe.Pointer(Rebox_Data_Map_Internal_2686810384_3421983481((&Constructor_Control_Alt_Alt[*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(_dollar___unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(Rebox_Data_Map_Internal_3281783655_2812149806(Rebox_Data_Map_Internal_2812149806_3281783655(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Data_Map_Internal_functorMap()))))}
}), gopurs_runtime.Func2(func(m1_2 gopurs_runtime.Value, m2_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeUnionWith(), compare_1_0, Get_Data_Function_go__const(), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](m1_2))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](m2_3))})))}
})})))}
}

func Call_Data_Map_Internal_plusMap(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
// TAST (Let): compare_1_1 shape=Other bindingType=(Func [(TypeVar k), (TypeVar k)] (ADT ["Data","Ordering","Ordering"] []))
compare_1_1 := gopurs_runtime.RecordGet(dictOrd_0, "compare")
_ = compare_1_1
// TAST (Let): altMap1_1_0 shape=LitRecord bindingType=(ADT ["Control","Alt","Alt"] [(ADT ["Data","Map","Internal","Map"] [(TypeVar k)])])
altMap1_1_0 := (&Constructor_Control_Alt_Alt[*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(_dollar___unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(Rebox_Data_Map_Internal_3281783655_2812149806(Rebox_Data_Map_Internal_2812149806_3281783655(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Data_Map_Internal_functorMap()))))}
}), gopurs_runtime.Func2(func(m1_2 gopurs_runtime.Value, m2_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeUnionWith(), compare_1_1, Get_Data_Function_go__const(), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](m1_2))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](m2_3))})))}
})})
_ = altMap1_1_0
return gopurs_runtime.Value{Type: 9, IntVal: 3709470893, UnsafePtr: unsafe.Pointer(Rebox_Data_Map_Internal_4212347312_3706288089((&Constructor_Control_Plus_Plus[*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(_dollar___unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4060500237, UnsafePtr: unsafe.Pointer(Rebox_Data_Map_Internal_2686810384_3421983481(altMap1_1_0))}
}), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(nil))}})))}
}

func Rebox_Data_Map_Internal_1245395425_3725484264(in *Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]]) *Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = in.V1
		out.V2 = in.V2
		out.V3 = in.V3
	return out
}

func Rebox_Data_Map_Internal_1291127239_1201789390(in *Constructor_Data_Monoid_Monoid[*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]]) *Constructor_Data_Monoid_Monoid[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(in.V1)}
	return out
}

func Rebox_Data_Map_Internal_138441832_3132786365(in *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]) *Constructor_Data_Tuple_Tuple[uint32, float64] {
	if in == nil { return nil }
	out := &Constructor_Data_Tuple_Tuple[uint32, float64]{}
		out.V0 = uint32(in.V0.IntVal)
		out.V1 = in.V1.FloatVal()
	return out
}

func Rebox_Data_Map_Internal_1397444001_2412140840(in *Constructor_Data_FunctorWithIndex_FunctorWithIndex[gopurs_runtime.Value, *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]]) *Constructor_Data_FunctorWithIndex_FunctorWithIndex[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_FunctorWithIndex_FunctorWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = in.V1
	return out
}

func Rebox_Data_Map_Internal_1444241223_3790796878(in *Constructor_Data_Eq_Eq[*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]]) *Constructor_Data_Eq_Eq[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Eq_Eq[gopurs_runtime.Value]{}
		out.V0 = in.V0
	return out
}

func Rebox_Data_Map_Internal_1628034448_3741347833(in *Constructor_Control_Apply_Apply[*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]]) *Constructor_Control_Apply_Apply[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Control_Apply_Apply[gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = in.V1
	return out
}

func Rebox_Data_Map_Internal_1680800814_3596835815(in *Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]) *Constructor_Data_Foldable_Foldable[*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]] {
	if in == nil { return nil }
	out := &Constructor_Data_Foldable_Foldable[*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]]{}
		out.V0 = in.V0
		out.V1 = in.V1
		out.V2 = in.V2
	return out
}

func Rebox_Data_Map_Internal_1785332133_138441832(in *Constructor_Data_Tuple_Tuple[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value], gopurs_runtime.Value]) *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{}
		out.V0 = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(in.V0)}
		out.V1 = in.V1
	return out
}

func Rebox_Data_Map_Internal_1860323088_2748095225(in *Constructor_Control_Bind_Bind[*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]]) *Constructor_Control_Bind_Bind[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Control_Bind_Bind[gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = in.V1
	return out
}

func Rebox_Data_Map_Internal_1910448679_4177771502(in *Constructor_Data_Ord_Ord[*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]]) *Constructor_Data_Ord_Ord[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Ord_Ord[gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = in.V1
	return out
}

func Rebox_Data_Map_Internal_2141765991_3043886126(in *Constructor_Data_Traversable_Traversable[*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]]) *Constructor_Data_Traversable_Traversable[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = in.V1
		out.V2 = in.V2
		out.V3 = in.V3
	return out
}

func Rebox_Data_Map_Internal_2412140840_1397444001(in *Constructor_Data_FunctorWithIndex_FunctorWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]) *Constructor_Data_FunctorWithIndex_FunctorWithIndex[gopurs_runtime.Value, *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]] {
	if in == nil { return nil }
	out := &Constructor_Data_FunctorWithIndex_FunctorWithIndex[gopurs_runtime.Value, *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]]{}
		out.V0 = in.V0
		out.V1 = in.V1
	return out
}

func Rebox_Data_Map_Internal_2442833393_849153993(in *Constructor_Data_List_Types_Cons[*Constructor_Data_Tuple_Tuple[uint32, float64]]) *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{}
		out.V0 = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Rebox_Data_Map_Internal_3132786365_138441832(in.V0))}
		out.V1 = Rebox_Data_Map_Internal_2442833393_849153993(in.V1)
	return out
}

func Rebox_Data_Map_Internal_2487766124_261879545(in *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]) *Constructor_Data_Map_Internal_Node[uint32, float64] {
	if in == nil { return nil }
	out := &Constructor_Data_Map_Internal_Node[uint32, float64]{}
		out.V0 = in.V0
		out.V1 = in.V1
		out.V2 = uint32(in.V2.IntVal)
		out.V3 = in.V3.FloatVal()
		out.V4 = Rebox_Data_Map_Internal_2487766124_261879545(in.V4)
		out.V5 = Rebox_Data_Map_Internal_2487766124_261879545(in.V5)
	return out
}

func Rebox_Data_Map_Internal_2549197956_3094389156(in *Constructor_Data_Maybe_Just[*Constructor_Data_Tuple_Tuple[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value], gopurs_runtime.Value]]) *Constructor_Data_Maybe_Just[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Maybe_Just[gopurs_runtime.Value]{}
		out.V0 = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Rebox_Data_Map_Internal_1785332133_138441832(in.V0))}
	return out
}

func Rebox_Data_Map_Internal_261879545_2487766124(in *Constructor_Data_Map_Internal_Node[uint32, float64]) *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = in.V1
		out.V2 = gopurs_runtime.Value{Type: 9, IntVal: int64(in.V2), UnsafePtr: nil}
		out.V3 = gopurs_runtime.Float(in.V3)
		out.V4 = Rebox_Data_Map_Internal_261879545_2487766124(in.V4)
		out.V5 = Rebox_Data_Map_Internal_261879545_2487766124(in.V5)
	return out
}

func Rebox_Data_Map_Internal_2638796135_1386611502(in *Constructor_Data_Show_Show[*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]]) *Constructor_Data_Show_Show[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Show_Show[gopurs_runtime.Value]{}
		out.V0 = in.V0
	return out
}

func Rebox_Data_Map_Internal_2686810384_3421983481(in *Constructor_Control_Alt_Alt[*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]]) *Constructor_Control_Alt_Alt[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Control_Alt_Alt[gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = in.V1
	return out
}

func Rebox_Data_Map_Internal_2758605161_3094389156(in *Constructor_Data_Maybe_Just[struct{
	key gopurs_runtime.Value
	value gopurs_runtime.Value
}]) *Constructor_Data_Maybe_Just[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Maybe_Just[gopurs_runtime.Value]{}
		out.V0 = func() gopurs_runtime.Value {
				orig := in.V0
				_ = orig
				return gopurs_runtime.RecordDict2("key", "value", orig.key, orig.value)
				}()
	return out
}

func Rebox_Data_Map_Internal_2812149806_3281783655(in *Constructor_Data_Functor_Functor[gopurs_runtime.Value]) *Constructor_Data_Functor_Functor[*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]] {
	if in == nil { return nil }
	out := &Constructor_Data_Functor_Functor[*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]]{}
		out.V0 = in.V0
	return out
}

func Rebox_Data_Map_Internal_3043886126_2141765991(in *Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]) *Constructor_Data_Traversable_Traversable[*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]] {
	if in == nil { return nil }
	out := &Constructor_Data_Traversable_Traversable[*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]]{}
		out.V0 = in.V0
		out.V1 = in.V1
		out.V2 = in.V2
		out.V3 = in.V3
	return out
}

func Rebox_Data_Map_Internal_3094389156_2758605161(in *Constructor_Data_Maybe_Just[gopurs_runtime.Value]) *Constructor_Data_Maybe_Just[struct{
	key gopurs_runtime.Value
	value gopurs_runtime.Value
}] {
	if in == nil { return nil }
	out := &Constructor_Data_Maybe_Just[struct{
	key gopurs_runtime.Value
	value gopurs_runtime.Value
}]{}
		out.V0 = func() struct{
	key gopurs_runtime.Value
	value gopurs_runtime.Value
} {
					orig := in.V0
					_ = orig
					clone := struct{
	key gopurs_runtime.Value
	value gopurs_runtime.Value
}{}
					clone.key = gopurs_runtime.RecordGet(orig, "key")
					clone.value = gopurs_runtime.RecordGet(orig, "value")
					return clone
				}()
	return out
}

func Rebox_Data_Map_Internal_3094389156_3240988860(in *Constructor_Data_Maybe_Just[gopurs_runtime.Value]) *Constructor_Data_Maybe_Just[float64] {
	if in == nil { return nil }
	out := &Constructor_Data_Maybe_Just[float64]{}
		out.V0 = in.V0.FloatVal()
	return out
}

func Rebox_Data_Map_Internal_3094389156_4010058633(in *Constructor_Data_Maybe_Just[gopurs_runtime.Value]) *Constructor_Data_Maybe_Just[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]] {
	if in == nil { return nil }
	out := &Constructor_Data_Maybe_Just[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]{}
		out.V0 = gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](in.V0)
	return out
}

func Rebox_Data_Map_Internal_3113727329_138441832(in *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]]) *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(in.V1)}
	return out
}

func Rebox_Data_Map_Internal_3132786365_138441832(in *Constructor_Data_Tuple_Tuple[uint32, float64]) *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{}
		out.V0 = gopurs_runtime.Value{Type: 9, IntVal: int64(in.V0), UnsafePtr: nil}
		out.V1 = gopurs_runtime.Float(in.V1)
	return out
}

func Rebox_Data_Map_Internal_3240988860_3094389156(in *Constructor_Data_Maybe_Just[float64]) *Constructor_Data_Maybe_Just[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Maybe_Just[gopurs_runtime.Value]{}
		out.V0 = gopurs_runtime.Float(in.V0)
	return out
}

func Rebox_Data_Map_Internal_3281783655_2812149806(in *Constructor_Data_Functor_Functor[*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]]) *Constructor_Data_Functor_Functor[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Functor_Functor[gopurs_runtime.Value]{}
		out.V0 = in.V0
	return out
}

func Rebox_Data_Map_Internal_3596835815_1680800814(in *Constructor_Data_Foldable_Foldable[*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]]) *Constructor_Data_Foldable_Foldable[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = in.V1
		out.V2 = in.V2
	return out
}

func Rebox_Data_Map_Internal_3691144502_1766074591(in *Constructor_Data_Eq_Eq1[*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]]) *Constructor_Data_Eq_Eq1[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Eq_Eq1[gopurs_runtime.Value]{}
		out.V0 = in.V0
	return out
}

func Rebox_Data_Map_Internal_3725484264_1245395425(in *Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]) *Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]] {
	if in == nil { return nil }
	out := &Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]]{}
		out.V0 = in.V0
		out.V1 = in.V1
		out.V2 = in.V2
		out.V3 = in.V3
	return out
}

func Rebox_Data_Map_Internal_3790796878_1444241223(in *Constructor_Data_Eq_Eq[gopurs_runtime.Value]) *Constructor_Data_Eq_Eq[*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]] {
	if in == nil { return nil }
	out := &Constructor_Data_Eq_Eq[*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]]{}
		out.V0 = in.V0
	return out
}

func Rebox_Data_Map_Internal_3854229351_4179793454(in *Constructor_Data_Semigroup_Semigroup[*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]]) *Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]{}
		out.V0 = in.V0
	return out
}

func Rebox_Data_Map_Internal_4018835873_1812164904(in *Constructor_Data_TraversableWithIndex_TraversableWithIndex[gopurs_runtime.Value, *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]]) *Constructor_Data_TraversableWithIndex_TraversableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_TraversableWithIndex_TraversableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = in.V1
		out.V2 = in.V2
		out.V3 = in.V3
	return out
}

func Rebox_Data_Map_Internal_4212347312_3706288089(in *Constructor_Control_Plus_Plus[*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]]) *Constructor_Control_Plus_Plus[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Control_Plus_Plus[gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = in.V1
	return out
}

func Rebox_Data_Map_Internal_849153993_2442833393(in *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_List_Types_Cons[*Constructor_Data_Tuple_Tuple[uint32, float64]] {
	if in == nil { return nil }
	out := &Constructor_Data_List_Types_Cons[*Constructor_Data_Tuple_Tuple[uint32, float64]]{}
		out.V0 = Rebox_Data_Map_Internal_138441832_3132786365(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](in.V0))
		out.V1 = Rebox_Data_Map_Internal_849153993_2442833393(in.V1)
	return out
}

func Rebox_Data_Map_Internal_849406934_3985601471(in *Constructor_Data_Ord_Ord1[*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]]) *Constructor_Data_Ord_Ord1[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Ord_Ord1[gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = in.V1
	return out
}


