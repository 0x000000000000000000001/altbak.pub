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
		cache_Data_Map_Internal_identity = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Map_Internal_identity(x_0_box)
})
	})
	return cache_Data_Map_Internal_identity
}

var cache_Data_Map_Internal_identity1 gopurs_runtime.Value
var once_Data_Map_Internal_identity1 sync.Once
func Get_Data_Map_Internal_identity1() gopurs_runtime.Value {
	once_Data_Map_Internal_identity1.Do(func() {
		cache_Data_Map_Internal_identity1 = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_Data_Map_Internal_identity1(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](x_0_box)))}
})
	})
	return cache_Data_Map_Internal_identity1
}

var cache_Data_Map_Internal_identity2 gopurs_runtime.Value
var once_Data_Map_Internal_identity2 sync.Once
func Get_Data_Map_Internal_identity2() gopurs_runtime.Value {
	once_Data_Map_Internal_identity2.Do(func() {
		cache_Data_Map_Internal_identity2 = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Map_Internal_identity2(x_0_box)
})
	})
	return cache_Data_Map_Internal_identity2
}

var cache_Data_Map_Internal_Leaf gopurs_runtime.Value
var once_Data_Map_Internal_Leaf sync.Once
func Get_Data_Map_Internal_Leaf() gopurs_runtime.Value {
	once_Data_Map_Internal_Leaf.Do(func() {
		cache_Data_Map_Internal_Leaf = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(nil))}
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
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(&Constructor_Data_Map_Internal_Node{1, value0.IntVal, value1.IntVal, value2, value3, gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](value4), gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](value5)})}
})
})
})
})
})
})
	})
	return cache_Data_Map_Internal_Node
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
return gopurs_runtime.Value{Type: 9, IntVal: 1343415489, UnsafePtr: unsafe.Pointer(&Constructor_Data_Map_Internal_IterEmit{1, value0, value1, value2})}
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
return gopurs_runtime.Value{Type: 9, IntVal: 2861335956, UnsafePtr: unsafe.Pointer(&Constructor_Data_Map_Internal_IterNode{1, gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](value0), value1})}
})
})
	})
	return cache_Data_Map_Internal_IterNode
}

var cache_Data_Map_Internal_IterDone gopurs_runtime.Value
var once_Data_Map_Internal_IterDone sync.Once
func Get_Data_Map_Internal_IterDone() gopurs_runtime.Value {
	once_Data_Map_Internal_IterDone.Do(func() {
		cache_Data_Map_Internal_IterDone = gopurs_runtime.Value{Type: 9, IntVal: 953589075, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_IterNext)(nil))}
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
return gopurs_runtime.Value{Type: 9, IntVal: 953589075, UnsafePtr: unsafe.Pointer(&Constructor_Data_Map_Internal_IterNext{1, value0, value1, value2})}
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
return gopurs_runtime.Value{Type: 9, IntVal: 3373277644, UnsafePtr: unsafe.Pointer(&Constructor_Data_Map_Internal_Split{1, gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](value0), gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](value1), gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](value2)})}
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
return gopurs_runtime.Value{Type: 9, IntVal: 2668112006, UnsafePtr: unsafe.Pointer(&Constructor_Data_Map_Internal_SplitLast{1, value0, value1, gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](value2)})}
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
		cache_Data_Map_Internal_unsafeNode = gopurs_runtime.Func4(func(__local_var_0_box gopurs_runtime.Value, __local_var_1_box gopurs_runtime.Value, __local_var_2_box gopurs_runtime.Value, __local_var_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Map_Internal_unsafeNode(__local_var_0_box, __local_var_1_box, __local_var_2_box, __local_var_3_box)
})
	})
	return cache_Data_Map_Internal_unsafeNode
}

var cache_Data_Map_Internal_toMapIter gopurs_runtime.Value
var once_Data_Map_Internal_toMapIter sync.Once
func Get_Data_Map_Internal_toMapIter() gopurs_runtime.Value {
	once_Data_Map_Internal_toMapIter.Do(func() {
		cache_Data_Map_Internal_toMapIter = gopurs_runtime.Func(func(a_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Map_Internal_toMapIter(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](a_0_box))
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
return gopurs_runtime.Int(Call_Data_Map_Internal_size(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](v_0_box)))
})
	})
	return cache_Data_Map_Internal_size
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

var cache_Data_Map_Internal_unsafeBalancedNode gopurs_runtime.Value
var once_Data_Map_Internal_unsafeBalancedNode sync.Once
func Get_Data_Map_Internal_unsafeBalancedNode() gopurs_runtime.Value {
	once_Data_Map_Internal_unsafeBalancedNode.Do(func() {
		cache_Data_Map_Internal_unsafeBalancedNode = gopurs_runtime.Func4(func(__local_var_0_box gopurs_runtime.Value, __local_var_1_box gopurs_runtime.Value, __local_var_2_box gopurs_runtime.Value, __local_var_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Map_Internal_unsafeBalancedNode(__local_var_0_box, __local_var_1_box, __local_var_2_box, __local_var_3_box)
})
	})
	return cache_Data_Map_Internal_unsafeBalancedNode
}

var cache_Data_Map_Internal_unsafeSplit gopurs_runtime.Value
var once_Data_Map_Internal_unsafeSplit sync.Once
func Get_Data_Map_Internal_unsafeSplit() gopurs_runtime.Value {
	once_Data_Map_Internal_unsafeSplit.Do(func() {
		cache_Data_Map_Internal_unsafeSplit = gopurs_runtime.Func3(func(__local_var_0_box gopurs_runtime.Value, __local_var_1_box gopurs_runtime.Value, __local_var_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Map_Internal_unsafeSplit(__local_var_0_box, __local_var_1_box, __local_var_2_box)
})
	})
	return cache_Data_Map_Internal_unsafeSplit
}

var cache_Data_Map_Internal_unsafeSplitLast gopurs_runtime.Value
var once_Data_Map_Internal_unsafeSplitLast sync.Once
func Get_Data_Map_Internal_unsafeSplitLast() gopurs_runtime.Value {
	once_Data_Map_Internal_unsafeSplitLast.Do(func() {
		cache_Data_Map_Internal_unsafeSplitLast = gopurs_runtime.Func4(func(__local_var_0_box gopurs_runtime.Value, __local_var_1_box gopurs_runtime.Value, __local_var_2_box gopurs_runtime.Value, __local_var_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Map_Internal_unsafeSplitLast(__local_var_0_box, __local_var_1_box, __local_var_2_box, __local_var_3_box)
})
	})
	return cache_Data_Map_Internal_unsafeSplitLast
}

var cache_Data_Map_Internal_unsafeJoinNodes gopurs_runtime.Value
var once_Data_Map_Internal_unsafeJoinNodes sync.Once
func Get_Data_Map_Internal_unsafeJoinNodes() gopurs_runtime.Value {
	once_Data_Map_Internal_unsafeJoinNodes.Do(func() {
		cache_Data_Map_Internal_unsafeJoinNodes = gopurs_runtime.Func2(func(__local_var_0_box gopurs_runtime.Value, __local_var_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Map_Internal_unsafeJoinNodes(__local_var_0_box, __local_var_1_box)
})
	})
	return cache_Data_Map_Internal_unsafeJoinNodes
}

var cache_Data_Map_Internal_unsafeDifference gopurs_runtime.Value
var once_Data_Map_Internal_unsafeDifference sync.Once
func Get_Data_Map_Internal_unsafeDifference() gopurs_runtime.Value {
	once_Data_Map_Internal_unsafeDifference.Do(func() {
		cache_Data_Map_Internal_unsafeDifference = gopurs_runtime.Func3(func(__local_var_0_box gopurs_runtime.Value, __local_var_1_box gopurs_runtime.Value, __local_var_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Map_Internal_unsafeDifference(__local_var_0_box, __local_var_1_box, __local_var_2_box)
})
	})
	return cache_Data_Map_Internal_unsafeDifference
}

var cache_Data_Map_Internal_unsafeIntersectionWith gopurs_runtime.Value
var once_Data_Map_Internal_unsafeIntersectionWith sync.Once
func Get_Data_Map_Internal_unsafeIntersectionWith() gopurs_runtime.Value {
	once_Data_Map_Internal_unsafeIntersectionWith.Do(func() {
		cache_Data_Map_Internal_unsafeIntersectionWith = gopurs_runtime.Func4(func(__local_var_0_box gopurs_runtime.Value, __local_var_1_box gopurs_runtime.Value, __local_var_2_box gopurs_runtime.Value, __local_var_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Map_Internal_unsafeIntersectionWith(__local_var_0_box, __local_var_1_box, __local_var_2_box, __local_var_3_box)
})
	})
	return cache_Data_Map_Internal_unsafeIntersectionWith
}

var cache_Data_Map_Internal_unsafeUnionWith gopurs_runtime.Value
var once_Data_Map_Internal_unsafeUnionWith sync.Once
func Get_Data_Map_Internal_unsafeUnionWith() gopurs_runtime.Value {
	once_Data_Map_Internal_unsafeUnionWith.Do(func() {
		cache_Data_Map_Internal_unsafeUnionWith = gopurs_runtime.Func4(func(__local_var_0_box gopurs_runtime.Value, __local_var_1_box gopurs_runtime.Value, __local_var_2_box gopurs_runtime.Value, __local_var_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Map_Internal_unsafeUnionWith(__local_var_0_box, __local_var_1_box, __local_var_2_box, __local_var_3_box)
})
	})
	return cache_Data_Map_Internal_unsafeUnionWith
}

var cache_Data_Map_Internal_unionWith gopurs_runtime.Value
var once_Data_Map_Internal_unionWith sync.Once
func Get_Data_Map_Internal_unionWith() gopurs_runtime.Value {
	once_Data_Map_Internal_unionWith.Do(func() {
		cache_Data_Map_Internal_unionWith = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Map_Internal_unionWith(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord](dictOrd_0_box))
})
	})
	return cache_Data_Map_Internal_unionWith
}

var cache_Data_Map_Internal_union gopurs_runtime.Value
var once_Data_Map_Internal_union sync.Once
func Get_Data_Map_Internal_union() gopurs_runtime.Value {
	once_Data_Map_Internal_union.Do(func() {
		cache_Data_Map_Internal_union = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Map_Internal_union(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord](dictOrd_0_box))
})
	})
	return cache_Data_Map_Internal_union
}

var cache_Data_Map_Internal_update gopurs_runtime.Value
var once_Data_Map_Internal_update sync.Once
func Get_Data_Map_Internal_update() gopurs_runtime.Value {
	once_Data_Map_Internal_update.Do(func() {
		cache_Data_Map_Internal_update = gopurs_runtime.Func3(func(dictOrd_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, k_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Map_Internal_update(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord](dictOrd_0_box), f_1_box, k_2_box)
})
	})
	return cache_Data_Map_Internal_update
}

var cache_Data_Map_Internal_showTree gopurs_runtime.Value
var once_Data_Map_Internal_showTree sync.Once
func Get_Data_Map_Internal_showTree() gopurs_runtime.Value {
	once_Data_Map_Internal_showTree.Do(func() {
		cache_Data_Map_Internal_showTree = gopurs_runtime.Func2(func(dictShow_0_box gopurs_runtime.Value, dictShow1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Map_Internal_showTree(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show](dictShow_0_box), gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show](dictShow1_1_box))
})
	})
	return cache_Data_Map_Internal_showTree
}

var cache_Data_Map_Internal_semigroupMap gopurs_runtime.Value
var once_Data_Map_Internal_semigroupMap sync.Once
func Get_Data_Map_Internal_semigroupMap() gopurs_runtime.Value {
	once_Data_Map_Internal_semigroupMap.Do(func() {
		cache_Data_Map_Internal_semigroupMap = gopurs_runtime.Func3(func(_dollar__unused_0_box gopurs_runtime.Value, dictOrd_1_box gopurs_runtime.Value, dictSemigroup_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Map_Internal_semigroupMap(_dollar__unused_0_box, dictOrd_1_box, dictSemigroup_2_box)
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
return Call_Data_Map_Internal_pop(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord](dictOrd_0_box))
})
	})
	return cache_Data_Map_Internal_pop
}

var cache_Data_Map_Internal_member gopurs_runtime.Value
var once_Data_Map_Internal_member sync.Once
func Get_Data_Map_Internal_member() gopurs_runtime.Value {
	once_Data_Map_Internal_member.Do(func() {
		cache_Data_Map_Internal_member = gopurs_runtime.Func2(func(dictOrd_0_box gopurs_runtime.Value, k_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Map_Internal_member(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord](dictOrd_0_box), k_1_box)
})
	})
	return cache_Data_Map_Internal_member
}

var cache_Data_Map_Internal_mapMaybeWithKey gopurs_runtime.Value
var once_Data_Map_Internal_mapMaybeWithKey sync.Once
func Get_Data_Map_Internal_mapMaybeWithKey() gopurs_runtime.Value {
	once_Data_Map_Internal_mapMaybeWithKey.Do(func() {
		cache_Data_Map_Internal_mapMaybeWithKey = gopurs_runtime.Func2(func(dictOrd_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Map_Internal_mapMaybeWithKey(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord](dictOrd_0_box), f_1_box)
})
	})
	return cache_Data_Map_Internal_mapMaybeWithKey
}

var cache_Data_Map_Internal_mapMaybe gopurs_runtime.Value
var once_Data_Map_Internal_mapMaybe sync.Once
func Get_Data_Map_Internal_mapMaybe() gopurs_runtime.Value {
	once_Data_Map_Internal_mapMaybe.Do(func() {
		cache_Data_Map_Internal_mapMaybe = gopurs_runtime.Func2(func(dictOrd_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Map_Internal_mapMaybe(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord](dictOrd_0_box), x_1_box)
})
	})
	return cache_Data_Map_Internal_mapMaybe
}

var cache_Data_Map_Internal_lookupLE gopurs_runtime.Value
var once_Data_Map_Internal_lookupLE sync.Once
func Get_Data_Map_Internal_lookupLE() gopurs_runtime.Value {
	once_Data_Map_Internal_lookupLE.Do(func() {
		cache_Data_Map_Internal_lookupLE = gopurs_runtime.Func2(func(dictOrd_0_box gopurs_runtime.Value, k_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Map_Internal_lookupLE(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord](dictOrd_0_box), k_1_box)
})
	})
	return cache_Data_Map_Internal_lookupLE
}

var cache_Data_Map_Internal_lookupGE gopurs_runtime.Value
var once_Data_Map_Internal_lookupGE sync.Once
func Get_Data_Map_Internal_lookupGE() gopurs_runtime.Value {
	once_Data_Map_Internal_lookupGE.Do(func() {
		cache_Data_Map_Internal_lookupGE = gopurs_runtime.Func2(func(dictOrd_0_box gopurs_runtime.Value, k_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Map_Internal_lookupGE(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord](dictOrd_0_box), k_1_box)
})
	})
	return cache_Data_Map_Internal_lookupGE
}

var cache_Data_Map_Internal_lookup gopurs_runtime.Value
var once_Data_Map_Internal_lookup sync.Once
func Get_Data_Map_Internal_lookup() gopurs_runtime.Value {
	once_Data_Map_Internal_lookup.Do(func() {
		cache_Data_Map_Internal_lookup = gopurs_runtime.Func2(func(dictOrd_0_box gopurs_runtime.Value, k_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Map_Internal_lookup(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord](dictOrd_0_box), k_1_box)
})
	})
	return cache_Data_Map_Internal_lookup
}

var cache_Data_Map_Internal_iterMapU gopurs_runtime.Value
var once_Data_Map_Internal_iterMapU sync.Once
func Get_Data_Map_Internal_iterMapU() gopurs_runtime.Value {
	once_Data_Map_Internal_iterMapU.Do(func() {
		cache_Data_Map_Internal_iterMapU = gopurs_runtime.Func2(func(iter_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Map_Internal_iterMapU(iter_0_box, gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](v_1_box))
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
		cache_Data_Map_Internal_stepUnfoldrUnordered = Call_Data_Map_Internal_stepWith(Get_Data_Map_Internal_iterMapU(), gopurs_runtime.Func3(func(__local_var_0 gopurs_runtime.Value, __local_var_1 gopurs_runtime.Value, __local_var_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, __local_var_0, __local_var_1})}, __local_var_2})}})}
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))}
}))
	})
	return cache_Data_Map_Internal_stepUnfoldrUnordered
}

var cache_Data_Map_Internal_toUnfoldableUnordered gopurs_runtime.Value
var once_Data_Map_Internal_toUnfoldableUnordered sync.Once
func Get_Data_Map_Internal_toUnfoldableUnordered() gopurs_runtime.Value {
	once_Data_Map_Internal_toUnfoldableUnordered.Do(func() {
		cache_Data_Map_Internal_toUnfoldableUnordered = gopurs_runtime.Func(func(dictUnfoldable_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Map_Internal_toUnfoldableUnordered(gopurs_runtime.CoerceToStruct[Constructor_Data_Unfoldable_Unfoldable](dictUnfoldable_0_box))
})
	})
	return cache_Data_Map_Internal_toUnfoldableUnordered
}

var cache_Data_Map_Internal_stepUnordered gopurs_runtime.Value
var once_Data_Map_Internal_stepUnordered sync.Once
func Get_Data_Map_Internal_stepUnordered() gopurs_runtime.Value {
	once_Data_Map_Internal_stepUnordered.Do(func() {
		cache_Data_Map_Internal_stepUnordered = Call_Data_Map_Internal_stepWith(Get_Data_Map_Internal_iterMapU(), gopurs_runtime.Func3(func(__local_var_0 gopurs_runtime.Value, __local_var_1 gopurs_runtime.Value, __local_var_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 953589075, UnsafePtr: unsafe.Pointer(&Constructor_Data_Map_Internal_IterNext{1, __local_var_0, __local_var_1, __local_var_2})}
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 953589075, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_IterNext)(nil))}
}))
	})
	return cache_Data_Map_Internal_stepUnordered
}

var cache_Data_Map_Internal_iterMapR gopurs_runtime.Value
var once_Data_Map_Internal_iterMapR sync.Once
func Get_Data_Map_Internal_iterMapR() gopurs_runtime.Value {
	once_Data_Map_Internal_iterMapR.Do(func() {
		cache_Data_Map_Internal_iterMapR = func() gopurs_runtime.Value {
var go__go_0_0_9 gopurs_runtime.Value
go__go_0_0_9 = gopurs_runtime.Func(func(iter_1_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var iter_1_loop gopurs_runtime.Value = iter_1_loop_val
var v_2_loop gopurs_runtime.Value = v_2_loop_val
go__go_0_0_9:
for {
if false { continue go__go_0_0_9 }
var iter_1 gopurs_runtime.Value = iter_1_loop
_ = iter_1
var v_2 gopurs_runtime.Value = v_2_loop
_ = v_2
var __t3 gopurs_runtime.Value
{
if (v_2.Type == 9 && v_2.IntVal == 324739070 && v_2.UnsafePtr == nil) {
__t3 = iter_1
goto end_branch_3
} else {

}
}
{
if (v_2.Type == 9 && v_2.IntVal == 324739070 && v_2.UnsafePtr != nil) {
var __t2 gopurs_runtime.Value
{
var __t_tag_1 *Constructor_Data_Map_Internal_Node = (*Constructor_Data_Map_Internal_Node)(v_2.UnsafePtr).V5
if (__t_tag_1 == nil) {
iter_1_loop = gopurs_runtime.Value{Type: 9, IntVal: 1343415489, UnsafePtr: unsafe.Pointer(&Constructor_Data_Map_Internal_IterEmit{1, (*Constructor_Data_Map_Internal_Node)(v_2.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node)(v_2.UnsafePtr).V3, iter_1})}
v_2_loop = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v_2.UnsafePtr).V4)}
continue go__go_0_0_9
__t2 = gopurs_runtime.Value{}
goto end_branch_2
} else {

}
}
{
iter_1_loop = gopurs_runtime.Value{Type: 9, IntVal: 1343415489, UnsafePtr: unsafe.Pointer(&Constructor_Data_Map_Internal_IterEmit{1, (*Constructor_Data_Map_Internal_Node)(v_2.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node)(v_2.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 2861335956, UnsafePtr: unsafe.Pointer(&Constructor_Data_Map_Internal_IterNode{1, (*Constructor_Data_Map_Internal_Node)(v_2.UnsafePtr).V4, iter_1})}})}
v_2_loop = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v_2.UnsafePtr).V5)}
continue go__go_0_0_9
__t2 = gopurs_runtime.Value{}
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
}()
})
})
return go__go_0_0_9
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
		cache_Data_Map_Internal_stepDesc = Call_Data_Map_Internal_stepWith(Get_Data_Map_Internal_iterMapR(), gopurs_runtime.Func3(func(__local_var_0 gopurs_runtime.Value, __local_var_1 gopurs_runtime.Value, __local_var_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 953589075, UnsafePtr: unsafe.Pointer(&Constructor_Data_Map_Internal_IterNext{1, __local_var_0, __local_var_1, __local_var_2})}
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 953589075, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_IterNext)(nil))}
}))
	})
	return cache_Data_Map_Internal_stepDesc
}

var cache_Data_Map_Internal_iterMapL gopurs_runtime.Value
var once_Data_Map_Internal_iterMapL sync.Once
func Get_Data_Map_Internal_iterMapL() gopurs_runtime.Value {
	once_Data_Map_Internal_iterMapL.Do(func() {
		cache_Data_Map_Internal_iterMapL = func() gopurs_runtime.Value {
var go__go_0_0_10 gopurs_runtime.Value
go__go_0_0_10 = gopurs_runtime.Func(func(iter_1_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var iter_1_loop gopurs_runtime.Value = iter_1_loop_val
var v_2_loop gopurs_runtime.Value = v_2_loop_val
go__go_0_0_10:
for {
if false { continue go__go_0_0_10 }
var iter_1 gopurs_runtime.Value = iter_1_loop
_ = iter_1
var v_2 gopurs_runtime.Value = v_2_loop
_ = v_2
var __t3 gopurs_runtime.Value
{
if (v_2.Type == 9 && v_2.IntVal == 324739070 && v_2.UnsafePtr == nil) {
__t3 = iter_1
goto end_branch_3
} else {

}
}
{
if (v_2.Type == 9 && v_2.IntVal == 324739070 && v_2.UnsafePtr != nil) {
var __t2 gopurs_runtime.Value
{
var __t_tag_1 *Constructor_Data_Map_Internal_Node = (*Constructor_Data_Map_Internal_Node)(v_2.UnsafePtr).V5
if (__t_tag_1 == nil) {
iter_1_loop = gopurs_runtime.Value{Type: 9, IntVal: 1343415489, UnsafePtr: unsafe.Pointer(&Constructor_Data_Map_Internal_IterEmit{1, (*Constructor_Data_Map_Internal_Node)(v_2.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node)(v_2.UnsafePtr).V3, iter_1})}
v_2_loop = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v_2.UnsafePtr).V4)}
continue go__go_0_0_10
__t2 = gopurs_runtime.Value{}
goto end_branch_2
} else {

}
}
{
iter_1_loop = gopurs_runtime.Value{Type: 9, IntVal: 1343415489, UnsafePtr: unsafe.Pointer(&Constructor_Data_Map_Internal_IterEmit{1, (*Constructor_Data_Map_Internal_Node)(v_2.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node)(v_2.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 2861335956, UnsafePtr: unsafe.Pointer(&Constructor_Data_Map_Internal_IterNode{1, (*Constructor_Data_Map_Internal_Node)(v_2.UnsafePtr).V5, iter_1})}})}
v_2_loop = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v_2.UnsafePtr).V4)}
continue go__go_0_0_10
__t2 = gopurs_runtime.Value{}
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
}()
})
})
return go__go_0_0_10
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
		cache_Data_Map_Internal_stepAsc = Call_Data_Map_Internal_stepWith(Get_Data_Map_Internal_iterMapL(), gopurs_runtime.Func3(func(__local_var_0 gopurs_runtime.Value, __local_var_1 gopurs_runtime.Value, __local_var_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 953589075, UnsafePtr: unsafe.Pointer(&Constructor_Data_Map_Internal_IterNext{1, __local_var_0, __local_var_1, __local_var_2})}
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 953589075, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_IterNext)(nil))}
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
		cache_Data_Map_Internal_stepUnfoldr = Call_Data_Map_Internal_stepWith(Get_Data_Map_Internal_iterMapL(), gopurs_runtime.Func3(func(__local_var_0 gopurs_runtime.Value, __local_var_1 gopurs_runtime.Value, __local_var_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, __local_var_0, __local_var_1})}, __local_var_2})}})}
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))}
}))
	})
	return cache_Data_Map_Internal_stepUnfoldr
}

var cache_Data_Map_Internal_toUnfoldable gopurs_runtime.Value
var once_Data_Map_Internal_toUnfoldable sync.Once
func Get_Data_Map_Internal_toUnfoldable() gopurs_runtime.Value {
	once_Data_Map_Internal_toUnfoldable.Do(func() {
		cache_Data_Map_Internal_toUnfoldable = gopurs_runtime.Func(func(dictUnfoldable_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Map_Internal_toUnfoldable(gopurs_runtime.CoerceToStruct[Constructor_Data_Unfoldable_Unfoldable](dictUnfoldable_0_box))
})
	})
	return cache_Data_Map_Internal_toUnfoldable
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
return Call_Data_Map_Internal_isSubmap(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord](dictOrd_0_box), gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq](dictEq_1_box))
})
	})
	return cache_Data_Map_Internal_isSubmap
}

var cache_Data_Map_Internal_isEmpty gopurs_runtime.Value
var once_Data_Map_Internal_isEmpty sync.Once
func Get_Data_Map_Internal_isEmpty() gopurs_runtime.Value {
	once_Data_Map_Internal_isEmpty.Do(func() {
		cache_Data_Map_Internal_isEmpty = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_Data_Map_Internal_isEmpty(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](v_0_box)))
})
	})
	return cache_Data_Map_Internal_isEmpty
}

var cache_Data_Map_Internal_intersectionWith gopurs_runtime.Value
var once_Data_Map_Internal_intersectionWith sync.Once
func Get_Data_Map_Internal_intersectionWith() gopurs_runtime.Value {
	once_Data_Map_Internal_intersectionWith.Do(func() {
		cache_Data_Map_Internal_intersectionWith = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Map_Internal_intersectionWith(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord](dictOrd_0_box))
})
	})
	return cache_Data_Map_Internal_intersectionWith
}

var cache_Data_Map_Internal_intersection gopurs_runtime.Value
var once_Data_Map_Internal_intersection sync.Once
func Get_Data_Map_Internal_intersection() gopurs_runtime.Value {
	once_Data_Map_Internal_intersection.Do(func() {
		cache_Data_Map_Internal_intersection = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Map_Internal_intersection(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord](dictOrd_0_box))
})
	})
	return cache_Data_Map_Internal_intersection
}

var cache_Data_Map_Internal_insertWith gopurs_runtime.Value
var once_Data_Map_Internal_insertWith sync.Once
func Get_Data_Map_Internal_insertWith() gopurs_runtime.Value {
	once_Data_Map_Internal_insertWith.Do(func() {
		cache_Data_Map_Internal_insertWith = gopurs_runtime.Func4(func(dictOrd_0_box gopurs_runtime.Value, app_1_box gopurs_runtime.Value, k_2_box gopurs_runtime.Value, v_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Map_Internal_insertWith(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord](dictOrd_0_box), app_1_box, k_2_box, v_3_box)
})
	})
	return cache_Data_Map_Internal_insertWith
}

var cache_Data_Map_Internal_insert gopurs_runtime.Value
var once_Data_Map_Internal_insert sync.Once
func Get_Data_Map_Internal_insert() gopurs_runtime.Value {
	once_Data_Map_Internal_insert.Do(func() {
		cache_Data_Map_Internal_insert = gopurs_runtime.Func3(func(dictOrd_0_box gopurs_runtime.Value, k_1_box gopurs_runtime.Value, v_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Map_Internal_insert(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord](dictOrd_0_box), k_1_box, v_2_box)
})
	})
	return cache_Data_Map_Internal_insert
}

var cache_Data_Map_Internal_functorMap gopurs_runtime.Value
var once_Data_Map_Internal_functorMap sync.Once
func Get_Data_Map_Internal_functorMap() gopurs_runtime.Value {
	once_Data_Map_Internal_functorMap.Do(func() {
		cache_Data_Map_Internal_functorMap = gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_1_0_18 gopurs_runtime.Value
_ = go__go_1_0_18
go__go_1_0_18 = gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 *Constructor_Data_Map_Internal_Node
{
if (v_2.Type == 9 && v_2.IntVal == 324739070 && v_2.UnsafePtr == nil) {
__t1 = (*Constructor_Data_Map_Internal_Node)(nil)
goto end_branch_1
} else {

}
}
{
if (v_2.Type == 9 && v_2.IntVal == 324739070 && v_2.UnsafePtr != nil) {
__t1 = &Constructor_Data_Map_Internal_Node{1, (*Constructor_Data_Map_Internal_Node)(v_2.UnsafePtr).V0, (*Constructor_Data_Map_Internal_Node)(v_2.UnsafePtr).V1, (*Constructor_Data_Map_Internal_Node)(v_2.UnsafePtr).V2, gopurs_runtime.Apply(f_0, (*Constructor_Data_Map_Internal_Node)(v_2.UnsafePtr).V3), gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.Apply(go__go_1_0_18, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v_2.UnsafePtr).V4)})), gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.Apply(go__go_1_0_18, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v_2.UnsafePtr).V5)}))}
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_1:
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(__t1)}
})
return go__go_1_0_18
}))
	})
	return cache_Data_Map_Internal_functorMap
}

var cache_Data_Map_Internal_functorWithIndexMap gopurs_runtime.Value
var once_Data_Map_Internal_functorWithIndexMap sync.Once
func Get_Data_Map_Internal_functorWithIndexMap() gopurs_runtime.Value {
	once_Data_Map_Internal_functorWithIndexMap.Do(func() {
		cache_Data_Map_Internal_functorWithIndexMap = gopurs_runtime.RecordDict2("Functor0", "mapWithIndex", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_2_0_19 gopurs_runtime.Value
_ = go__go_2_0_19
go__go_2_0_19 = gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 *Constructor_Data_Map_Internal_Node
{
if (v_3.Type == 9 && v_3.IntVal == 324739070 && v_3.UnsafePtr == nil) {
__t1 = (*Constructor_Data_Map_Internal_Node)(nil)
goto end_branch_1
} else {

}
}
{
if (v_3.Type == 9 && v_3.IntVal == 324739070 && v_3.UnsafePtr != nil) {
__t1 = &Constructor_Data_Map_Internal_Node{1, (*Constructor_Data_Map_Internal_Node)(v_3.UnsafePtr).V0, (*Constructor_Data_Map_Internal_Node)(v_3.UnsafePtr).V1, (*Constructor_Data_Map_Internal_Node)(v_3.UnsafePtr).V2, gopurs_runtime.Apply(f_1, (*Constructor_Data_Map_Internal_Node)(v_3.UnsafePtr).V3), gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.Apply(go__go_2_0_19, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v_3.UnsafePtr).V4)})), gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.Apply(go__go_2_0_19, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v_3.UnsafePtr).V5)}))}
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_1:
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(__t1)}
})
return go__go_2_0_19
}))
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_1_2_20 gopurs_runtime.Value
_ = go__go_1_2_20
go__go_1_2_20 = gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t3 *Constructor_Data_Map_Internal_Node
{
if (v_2.Type == 9 && v_2.IntVal == 324739070 && v_2.UnsafePtr == nil) {
__t3 = (*Constructor_Data_Map_Internal_Node)(nil)
goto end_branch_3
} else {

}
}
{
if (v_2.Type == 9 && v_2.IntVal == 324739070 && v_2.UnsafePtr != nil) {
__t3 = &Constructor_Data_Map_Internal_Node{1, (*Constructor_Data_Map_Internal_Node)(v_2.UnsafePtr).V0, (*Constructor_Data_Map_Internal_Node)(v_2.UnsafePtr).V1, (*Constructor_Data_Map_Internal_Node)(v_2.UnsafePtr).V2, gopurs_runtime.Apply2(f_0, (*Constructor_Data_Map_Internal_Node)(v_2.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node)(v_2.UnsafePtr).V3), gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.Apply(go__go_1_2_20, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v_2.UnsafePtr).V4)})), gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.Apply(go__go_1_2_20, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v_2.UnsafePtr).V5)}))}
goto end_branch_3
} else {

}
}
{
__t3 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_3:
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(__t3)}
})
return go__go_1_2_20
}))
	})
	return cache_Data_Map_Internal_functorWithIndexMap
}

var cache_Data_Map_Internal_foldableMap gopurs_runtime.Value
var once_Data_Map_Internal_foldableMap sync.Once
func Get_Data_Map_Internal_foldableMap() gopurs_runtime.Value {
	once_Data_Map_Internal_foldableMap.Do(func() {
		cache_Data_Map_Internal_foldableMap = gopurs_runtime.RecordDict3("foldMap", "foldl", "foldr", gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Semigroup0_1_0 -> *Constructor_Data_Semigroup_Semigroup
Semigroup0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_1_0
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_3_1_21 gopurs_runtime.Value
_ = go__go_3_1_21
go__go_3_1_21 = gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
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
__t2 = gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_1_0.V0), gopurs_runtime.Apply(go__go_3_1_21, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v_4.UnsafePtr).V4)}), gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_1_0.V0), gopurs_runtime.Apply(f_2, (*Constructor_Data_Map_Internal_Node)(v_4.UnsafePtr).V3), gopurs_runtime.Apply(go__go_3_1_21, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v_4.UnsafePtr).V5)})))
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
return go__go_3_1_21
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(z_1 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_2_3_22 gopurs_runtime.Value
go__go_2_3_22 = gopurs_runtime.Func(func(__local_var_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(__local_var_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var __local_var_3_loop gopurs_runtime.Value = __local_var_3_loop_val
var __local_var_4_loop gopurs_runtime.Value = __local_var_4_loop_val
go__go_2_3_22:
for {
if false { continue go__go_2_3_22 }
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
__t4 = gopurs_runtime.UncurriedApp2(go__go_2_3_22, gopurs_runtime.Apply2(f_0, gopurs_runtime.UncurriedApp2(go__go_2_3_22, __local_var_3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_4.UnsafePtr).V4)}), (*Constructor_Data_Map_Internal_Node)(__local_var_4.UnsafePtr).V3), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_4.UnsafePtr).V5)})
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
return gopurs_runtime.UncurriedApp2(go__go_2_3_22, z_1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](m_3))})
})
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(z_1 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_2_5_23 gopurs_runtime.Value
_ = go__go_2_5_23
go__go_2_5_23 = gopurs_runtime.Func2(func(__local_var_3 gopurs_runtime.Value, __local_var_4 gopurs_runtime.Value) gopurs_runtime.Value {
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
__t6 = gopurs_runtime.UncurriedApp2(go__go_2_5_23, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_3.UnsafePtr).V4)}, gopurs_runtime.Apply2(f_0, (*Constructor_Data_Map_Internal_Node)(__local_var_3.UnsafePtr).V3, gopurs_runtime.UncurriedApp2(go__go_2_5_23, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_3.UnsafePtr).V5)}, __local_var_4)))
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
return gopurs_runtime.UncurriedApp2(go__go_2_5_23, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](m_3))}, z_1)
})
})
}))
	})
	return cache_Data_Map_Internal_foldableMap
}

var cache_Data_Map_Internal_foldableWithIndexMap gopurs_runtime.Value
var once_Data_Map_Internal_foldableWithIndexMap sync.Once
func Get_Data_Map_Internal_foldableWithIndexMap() gopurs_runtime.Value {
	once_Data_Map_Internal_foldableWithIndexMap.Do(func() {
		cache_Data_Map_Internal_foldableWithIndexMap = gopurs_runtime.RecordDict4("Foldable0", "foldMapWithIndex", "foldlWithIndex", "foldrWithIndex", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict3("foldMap", "foldl", "foldr", gopurs_runtime.Func(func(dictMonoid_1 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Semigroup0_2_0 -> *Constructor_Data_Semigroup_Semigroup
Semigroup0_2_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_1, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_2_0
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_4_1_24 gopurs_runtime.Value
_ = go__go_4_1_24
go__go_4_1_24 = gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 gopurs_runtime.Value
{
if (v_5.Type == 9 && v_5.IntVal == 324739070 && v_5.UnsafePtr == nil) {
__t2 = gopurs_runtime.RecordGet(dictMonoid_1, "mempty")
goto end_branch_2
} else {

}
}
{
if (v_5.Type == 9 && v_5.IntVal == 324739070 && v_5.UnsafePtr != nil) {
__t2 = gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_2_0.V0), gopurs_runtime.Apply(go__go_4_1_24, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v_5.UnsafePtr).V4)}), gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_2_0.V0), gopurs_runtime.Apply(f_3, (*Constructor_Data_Map_Internal_Node)(v_5.UnsafePtr).V3), gopurs_runtime.Apply(go__go_4_1_24, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v_5.UnsafePtr).V5)})))
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
return go__go_4_1_24
})
}), gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(z_2 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_3_3_25 gopurs_runtime.Value
go__go_3_3_25 = gopurs_runtime.Func(func(__local_var_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(__local_var_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var __local_var_4_loop gopurs_runtime.Value = __local_var_4_loop_val
var __local_var_5_loop gopurs_runtime.Value = __local_var_5_loop_val
go__go_3_3_25:
for {
if false { continue go__go_3_3_25 }
var __local_var_4 gopurs_runtime.Value = __local_var_4_loop
_ = __local_var_4
var __local_var_5 gopurs_runtime.Value = __local_var_5_loop
_ = __local_var_5
var __t4 gopurs_runtime.Value
{
if (__local_var_5.Type == 9 && __local_var_5.IntVal == 324739070 && __local_var_5.UnsafePtr == nil) {
__t4 = __local_var_4
goto end_branch_4
} else {

}
}
{
if (__local_var_5.Type == 9 && __local_var_5.IntVal == 324739070 && __local_var_5.UnsafePtr != nil) {
__t4 = gopurs_runtime.UncurriedApp2(go__go_3_3_25, gopurs_runtime.Apply2(f_1, gopurs_runtime.UncurriedApp2(go__go_3_3_25, __local_var_4, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_5.UnsafePtr).V4)}), (*Constructor_Data_Map_Internal_Node)(__local_var_5.UnsafePtr).V3), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_5.UnsafePtr).V5)})
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
return gopurs_runtime.Func(func(m_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.UncurriedApp2(go__go_3_3_25, z_2, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](m_4))})
})
})
}), gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(z_2 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_3_5_26 gopurs_runtime.Value
_ = go__go_3_5_26
go__go_3_5_26 = gopurs_runtime.Func2(func(__local_var_4 gopurs_runtime.Value, __local_var_5 gopurs_runtime.Value) gopurs_runtime.Value {
var __t6 gopurs_runtime.Value
{
if (__local_var_4.Type == 9 && __local_var_4.IntVal == 324739070 && __local_var_4.UnsafePtr == nil) {
__t6 = __local_var_5
goto end_branch_6
} else {

}
}
{
if (__local_var_4.Type == 9 && __local_var_4.IntVal == 324739070 && __local_var_4.UnsafePtr != nil) {
__t6 = gopurs_runtime.UncurriedApp2(go__go_3_5_26, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_4.UnsafePtr).V4)}, gopurs_runtime.Apply2(f_1, (*Constructor_Data_Map_Internal_Node)(__local_var_4.UnsafePtr).V3, gopurs_runtime.UncurriedApp2(go__go_3_5_26, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_4.UnsafePtr).V5)}, __local_var_5)))
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
return gopurs_runtime.Func(func(m_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.UncurriedApp2(go__go_3_5_26, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](m_4))}, z_2)
})
})
}))
}), gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Semigroup0_1_7 -> *Constructor_Data_Semigroup_Semigroup
Semigroup0_1_7 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_1_7
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_3_8_27 gopurs_runtime.Value
_ = go__go_3_8_27
go__go_3_8_27 = gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t9 gopurs_runtime.Value
{
if (v_4.Type == 9 && v_4.IntVal == 324739070 && v_4.UnsafePtr == nil) {
__t9 = gopurs_runtime.RecordGet(dictMonoid_0, "mempty")
goto end_branch_9
} else {

}
}
{
if (v_4.Type == 9 && v_4.IntVal == 324739070 && v_4.UnsafePtr != nil) {
__t9 = gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_1_7.V0), gopurs_runtime.Apply(go__go_3_8_27, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v_4.UnsafePtr).V4)}), gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_1_7.V0), gopurs_runtime.Apply2(f_2, (*Constructor_Data_Map_Internal_Node)(v_4.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node)(v_4.UnsafePtr).V3), gopurs_runtime.Apply(go__go_3_8_27, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v_4.UnsafePtr).V5)})))
goto end_branch_9
} else {

}
}
{
__t9 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_9:
return __t9
})
return go__go_3_8_27
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(z_1 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_2_10_28 gopurs_runtime.Value
go__go_2_10_28 = gopurs_runtime.Func(func(__local_var_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(__local_var_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var __local_var_3_loop gopurs_runtime.Value = __local_var_3_loop_val
var __local_var_4_loop gopurs_runtime.Value = __local_var_4_loop_val
go__go_2_10_28:
for {
if false { continue go__go_2_10_28 }
var __local_var_3 gopurs_runtime.Value = __local_var_3_loop
_ = __local_var_3
var __local_var_4 gopurs_runtime.Value = __local_var_4_loop
_ = __local_var_4
var __t11 gopurs_runtime.Value
{
if (__local_var_4.Type == 9 && __local_var_4.IntVal == 324739070 && __local_var_4.UnsafePtr == nil) {
__t11 = __local_var_3
goto end_branch_11
} else {

}
}
{
if (__local_var_4.Type == 9 && __local_var_4.IntVal == 324739070 && __local_var_4.UnsafePtr != nil) {
__t11 = gopurs_runtime.UncurriedApp2(go__go_2_10_28, gopurs_runtime.Apply3(f_0, (*Constructor_Data_Map_Internal_Node)(__local_var_4.UnsafePtr).V2, gopurs_runtime.UncurriedApp2(go__go_2_10_28, __local_var_3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_4.UnsafePtr).V4)}), (*Constructor_Data_Map_Internal_Node)(__local_var_4.UnsafePtr).V3), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_4.UnsafePtr).V5)})
goto end_branch_11
} else {

}
}
{
__t11 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_11:
return __t11
}
}()
})
})
return gopurs_runtime.Func(func(m_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.UncurriedApp2(go__go_2_10_28, z_1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](m_3))})
})
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(z_1 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_2_12_29 gopurs_runtime.Value
_ = go__go_2_12_29
go__go_2_12_29 = gopurs_runtime.Func2(func(__local_var_3 gopurs_runtime.Value, __local_var_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t13 gopurs_runtime.Value
{
if (__local_var_3.Type == 9 && __local_var_3.IntVal == 324739070 && __local_var_3.UnsafePtr == nil) {
__t13 = __local_var_4
goto end_branch_13
} else {

}
}
{
if (__local_var_3.Type == 9 && __local_var_3.IntVal == 324739070 && __local_var_3.UnsafePtr != nil) {
__t13 = gopurs_runtime.UncurriedApp2(go__go_2_12_29, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_3.UnsafePtr).V4)}, gopurs_runtime.Apply3(f_0, (*Constructor_Data_Map_Internal_Node)(__local_var_3.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node)(__local_var_3.UnsafePtr).V3, gopurs_runtime.UncurriedApp2(go__go_2_12_29, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_3.UnsafePtr).V5)}, __local_var_4)))
goto end_branch_13
} else {

}
}
{
__t13 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_13:
return __t13
})
return gopurs_runtime.Func(func(m_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.UncurriedApp2(go__go_2_12_29, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](m_3))}, z_1)
})
})
}))
	})
	return cache_Data_Map_Internal_foldableWithIndexMap
}

var cache_Data_Map_Internal_keys gopurs_runtime.Value
var once_Data_Map_Internal_keys sync.Once
func Get_Data_Map_Internal_keys() gopurs_runtime.Value {
	once_Data_Map_Internal_keys.Do(func() {
		cache_Data_Map_Internal_keys = func() gopurs_runtime.Value {
var go__go_0_0_30 gopurs_runtime.Value
_ = go__go_0_0_30
go__go_0_0_30 = gopurs_runtime.Func2(func(__local_var_1 gopurs_runtime.Value, __local_var_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if (__local_var_1.Type == 9 && __local_var_1.IntVal == 324739070 && __local_var_1.UnsafePtr == nil) {
__t1 = __local_var_2
goto end_branch_1
} else {

}
}
{
if (__local_var_1.Type == 9 && __local_var_1.IntVal == 324739070 && __local_var_1.UnsafePtr != nil) {
__t1 = gopurs_runtime.UncurriedApp2(go__go_0_0_30, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_1.UnsafePtr).V4)}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, (*Constructor_Data_Map_Internal_Node)(__local_var_1.UnsafePtr).V2, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.UncurriedApp2(go__go_0_0_30, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_1.UnsafePtr).V5)}, __local_var_2))})})
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
return gopurs_runtime.Func(func(m_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.UncurriedApp2(go__go_0_0_30, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](m_1))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(nil))})
})
}()
	})
	return cache_Data_Map_Internal_keys
}

var cache_Data_Map_Internal_traversableMap gopurs_runtime.Value
var once_Data_Map_Internal_traversableMap sync.Once
func Get_Data_Map_Internal_traversableMap() gopurs_runtime.Value {
	once_Data_Map_Internal_traversableMap.Do(func() {
		cache_Data_Map_Internal_traversableMap = gopurs_runtime.RecordDict4("Foldable1", "Functor0", "sequence", "traverse", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict3("foldMap", "foldl", "foldr", gopurs_runtime.Func(func(dictMonoid_1 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Semigroup0_2_0 -> *Constructor_Data_Semigroup_Semigroup
Semigroup0_2_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_1, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_2_0
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_4_1_31 gopurs_runtime.Value
_ = go__go_4_1_31
go__go_4_1_31 = gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 gopurs_runtime.Value
{
if (v_5.Type == 9 && v_5.IntVal == 324739070 && v_5.UnsafePtr == nil) {
__t2 = gopurs_runtime.RecordGet(dictMonoid_1, "mempty")
goto end_branch_2
} else {

}
}
{
if (v_5.Type == 9 && v_5.IntVal == 324739070 && v_5.UnsafePtr != nil) {
__t2 = gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_2_0.V0), gopurs_runtime.Apply(go__go_4_1_31, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v_5.UnsafePtr).V4)}), gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_2_0.V0), gopurs_runtime.Apply(f_3, (*Constructor_Data_Map_Internal_Node)(v_5.UnsafePtr).V3), gopurs_runtime.Apply(go__go_4_1_31, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v_5.UnsafePtr).V5)})))
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
return go__go_4_1_31
})
}), gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(z_2 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_3_3_32 gopurs_runtime.Value
go__go_3_3_32 = gopurs_runtime.Func(func(__local_var_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(__local_var_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var __local_var_4_loop gopurs_runtime.Value = __local_var_4_loop_val
var __local_var_5_loop gopurs_runtime.Value = __local_var_5_loop_val
go__go_3_3_32:
for {
if false { continue go__go_3_3_32 }
var __local_var_4 gopurs_runtime.Value = __local_var_4_loop
_ = __local_var_4
var __local_var_5 gopurs_runtime.Value = __local_var_5_loop
_ = __local_var_5
var __t4 gopurs_runtime.Value
{
if (__local_var_5.Type == 9 && __local_var_5.IntVal == 324739070 && __local_var_5.UnsafePtr == nil) {
__t4 = __local_var_4
goto end_branch_4
} else {

}
}
{
if (__local_var_5.Type == 9 && __local_var_5.IntVal == 324739070 && __local_var_5.UnsafePtr != nil) {
__t4 = gopurs_runtime.UncurriedApp2(go__go_3_3_32, gopurs_runtime.Apply2(f_1, gopurs_runtime.UncurriedApp2(go__go_3_3_32, __local_var_4, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_5.UnsafePtr).V4)}), (*Constructor_Data_Map_Internal_Node)(__local_var_5.UnsafePtr).V3), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_5.UnsafePtr).V5)})
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
return gopurs_runtime.Func(func(m_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.UncurriedApp2(go__go_3_3_32, z_2, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](m_4))})
})
})
}), gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(z_2 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_3_5_33 gopurs_runtime.Value
_ = go__go_3_5_33
go__go_3_5_33 = gopurs_runtime.Func2(func(__local_var_4 gopurs_runtime.Value, __local_var_5 gopurs_runtime.Value) gopurs_runtime.Value {
var __t6 gopurs_runtime.Value
{
if (__local_var_4.Type == 9 && __local_var_4.IntVal == 324739070 && __local_var_4.UnsafePtr == nil) {
__t6 = __local_var_5
goto end_branch_6
} else {

}
}
{
if (__local_var_4.Type == 9 && __local_var_4.IntVal == 324739070 && __local_var_4.UnsafePtr != nil) {
__t6 = gopurs_runtime.UncurriedApp2(go__go_3_5_33, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_4.UnsafePtr).V4)}, gopurs_runtime.Apply2(f_1, (*Constructor_Data_Map_Internal_Node)(__local_var_4.UnsafePtr).V3, gopurs_runtime.UncurriedApp2(go__go_3_5_33, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_4.UnsafePtr).V5)}, __local_var_5)))
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
return gopurs_runtime.Func(func(m_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.UncurriedApp2(go__go_3_5_33, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](m_4))}, z_2)
})
})
}))
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_2_7_34 gopurs_runtime.Value
_ = go__go_2_7_34
go__go_2_7_34 = gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t8 *Constructor_Data_Map_Internal_Node
{
if (v_3.Type == 9 && v_3.IntVal == 324739070 && v_3.UnsafePtr == nil) {
__t8 = (*Constructor_Data_Map_Internal_Node)(nil)
goto end_branch_8
} else {

}
}
{
if (v_3.Type == 9 && v_3.IntVal == 324739070 && v_3.UnsafePtr != nil) {
__t8 = &Constructor_Data_Map_Internal_Node{1, (*Constructor_Data_Map_Internal_Node)(v_3.UnsafePtr).V0, (*Constructor_Data_Map_Internal_Node)(v_3.UnsafePtr).V1, (*Constructor_Data_Map_Internal_Node)(v_3.UnsafePtr).V2, gopurs_runtime.Apply(f_1, (*Constructor_Data_Map_Internal_Node)(v_3.UnsafePtr).V3), gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.Apply(go__go_2_7_34, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v_3.UnsafePtr).V4)})), gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.Apply(go__go_2_7_34, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v_3.UnsafePtr).V5)}))}
goto end_branch_8
} else {

}
}
{
__t8 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_8:
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(__t8)}
})
return go__go_2_7_34
}))
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_Map_Internal_traversableMap(), "traverse"), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](dictApplicative_0))}, gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return x_1
}))
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Apply0_1_9 -> *Constructor_Control_Apply_Apply
Apply0_1_9 := gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}))
_ = Apply0_1_9
// TAST (Let): Functor0_2_10 -> *Constructor_Data_Functor_Functor
Functor0_2_10 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_2_10
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_4_11_35 gopurs_runtime.Value
_ = go__go_4_11_35
go__go_4_11_35 = gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
var __t15 gopurs_runtime.Value
{
if (v_5.Type == 9 && v_5.IntVal == 324739070 && v_5.UnsafePtr == nil) {
__t15 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "pure"), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(nil))})
goto end_branch_15
} else {

}
}
{
if (v_5.Type == 9 && v_5.IntVal == 324739070 && v_5.UnsafePtr != nil) {
// TAST (Let): __local_var_6_12 -> gopurs_runtime.Value
var __local_var_6_12 gopurs_runtime.Value = gopurs_runtime.Int((*Constructor_Data_Map_Internal_Node)(v_5.UnsafePtr).V0)
// TAST (Let): __local_var_7_13 -> gopurs_runtime.Value
__local_var_7_13 := (*Constructor_Data_Map_Internal_Node)(v_5.UnsafePtr).V2
_ = __local_var_7_13
// TAST (Let): __local_var_8_14 -> gopurs_runtime.Value
var __local_var_8_14 gopurs_runtime.Value = gopurs_runtime.Int((*Constructor_Data_Map_Internal_Node)(v_5.UnsafePtr).V1)
__t15 = gopurs_runtime.Apply2(gopurs_runtime.Box(Apply0_1_9.V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Apply0_1_9.V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_2_10.V0), gopurs_runtime.Func(func(l_prime_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_prime_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(r_prime_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(&Constructor_Data_Map_Internal_Node{1, __local_var_6_12.IntVal, __local_var_8_14.IntVal, __local_var_7_13, v_prime_10, gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](l_prime_9), gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](r_prime_11)})}
})
})
}), gopurs_runtime.Apply(go__go_4_11_35, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v_5.UnsafePtr).V4)})), gopurs_runtime.Apply(f_3, (*Constructor_Data_Map_Internal_Node)(v_5.UnsafePtr).V3)), gopurs_runtime.Apply(go__go_4_11_35, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v_5.UnsafePtr).V5)}))
goto end_branch_15
} else {

}
}
{
__t15 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_15:
return __t15
})
return go__go_4_11_35
})
}))
	})
	return cache_Data_Map_Internal_traversableMap
}

var cache_Data_Map_Internal_traversableWithIndexMap gopurs_runtime.Value
var once_Data_Map_Internal_traversableWithIndexMap sync.Once
func Get_Data_Map_Internal_traversableWithIndexMap() gopurs_runtime.Value {
	once_Data_Map_Internal_traversableWithIndexMap.Do(func() {
		cache_Data_Map_Internal_traversableWithIndexMap = gopurs_runtime.RecordDict4("FoldableWithIndex1", "FunctorWithIndex0", "Traversable2", "traverseWithIndex", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict4("Foldable0", "foldMapWithIndex", "foldlWithIndex", "foldrWithIndex", gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict3("foldMap", "foldl", "foldr", gopurs_runtime.Func(func(dictMonoid_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Semigroup0_3_0 -> *Constructor_Data_Semigroup_Semigroup
Semigroup0_3_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_2, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_3_0
return gopurs_runtime.Func(func(f_4 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_5_1_36 gopurs_runtime.Value
_ = go__go_5_1_36
go__go_5_1_36 = gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 gopurs_runtime.Value
{
if (v_6.Type == 9 && v_6.IntVal == 324739070 && v_6.UnsafePtr == nil) {
__t2 = gopurs_runtime.RecordGet(dictMonoid_2, "mempty")
goto end_branch_2
} else {

}
}
{
if (v_6.Type == 9 && v_6.IntVal == 324739070 && v_6.UnsafePtr != nil) {
__t2 = gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_3_0.V0), gopurs_runtime.Apply(go__go_5_1_36, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v_6.UnsafePtr).V4)}), gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_3_0.V0), gopurs_runtime.Apply(f_4, (*Constructor_Data_Map_Internal_Node)(v_6.UnsafePtr).V3), gopurs_runtime.Apply(go__go_5_1_36, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v_6.UnsafePtr).V5)})))
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
return go__go_5_1_36
})
}), gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(z_3 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_4_3_37 gopurs_runtime.Value
go__go_4_3_37 = gopurs_runtime.Func(func(__local_var_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(__local_var_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var __local_var_5_loop gopurs_runtime.Value = __local_var_5_loop_val
var __local_var_6_loop gopurs_runtime.Value = __local_var_6_loop_val
go__go_4_3_37:
for {
if false { continue go__go_4_3_37 }
var __local_var_5 gopurs_runtime.Value = __local_var_5_loop
_ = __local_var_5
var __local_var_6 gopurs_runtime.Value = __local_var_6_loop
_ = __local_var_6
var __t4 gopurs_runtime.Value
{
if (__local_var_6.Type == 9 && __local_var_6.IntVal == 324739070 && __local_var_6.UnsafePtr == nil) {
__t4 = __local_var_5
goto end_branch_4
} else {

}
}
{
if (__local_var_6.Type == 9 && __local_var_6.IntVal == 324739070 && __local_var_6.UnsafePtr != nil) {
__t4 = gopurs_runtime.UncurriedApp2(go__go_4_3_37, gopurs_runtime.Apply2(f_2, gopurs_runtime.UncurriedApp2(go__go_4_3_37, __local_var_5, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_6.UnsafePtr).V4)}), (*Constructor_Data_Map_Internal_Node)(__local_var_6.UnsafePtr).V3), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_6.UnsafePtr).V5)})
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
return gopurs_runtime.Func(func(m_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.UncurriedApp2(go__go_4_3_37, z_3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](m_5))})
})
})
}), gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(z_3 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_4_5_38 gopurs_runtime.Value
_ = go__go_4_5_38
go__go_4_5_38 = gopurs_runtime.Func2(func(__local_var_5 gopurs_runtime.Value, __local_var_6 gopurs_runtime.Value) gopurs_runtime.Value {
var __t6 gopurs_runtime.Value
{
if (__local_var_5.Type == 9 && __local_var_5.IntVal == 324739070 && __local_var_5.UnsafePtr == nil) {
__t6 = __local_var_6
goto end_branch_6
} else {

}
}
{
if (__local_var_5.Type == 9 && __local_var_5.IntVal == 324739070 && __local_var_5.UnsafePtr != nil) {
__t6 = gopurs_runtime.UncurriedApp2(go__go_4_5_38, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_5.UnsafePtr).V4)}, gopurs_runtime.Apply2(f_2, (*Constructor_Data_Map_Internal_Node)(__local_var_5.UnsafePtr).V3, gopurs_runtime.UncurriedApp2(go__go_4_5_38, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_5.UnsafePtr).V5)}, __local_var_6)))
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
return gopurs_runtime.Func(func(m_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.UncurriedApp2(go__go_4_5_38, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](m_5))}, z_3)
})
})
}))
}), gopurs_runtime.Func(func(dictMonoid_1 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Semigroup0_2_7 -> *Constructor_Data_Semigroup_Semigroup
Semigroup0_2_7 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_1, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_2_7
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_4_8_39 gopurs_runtime.Value
_ = go__go_4_8_39
go__go_4_8_39 = gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
var __t9 gopurs_runtime.Value
{
if (v_5.Type == 9 && v_5.IntVal == 324739070 && v_5.UnsafePtr == nil) {
__t9 = gopurs_runtime.RecordGet(dictMonoid_1, "mempty")
goto end_branch_9
} else {

}
}
{
if (v_5.Type == 9 && v_5.IntVal == 324739070 && v_5.UnsafePtr != nil) {
__t9 = gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_2_7.V0), gopurs_runtime.Apply(go__go_4_8_39, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v_5.UnsafePtr).V4)}), gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_2_7.V0), gopurs_runtime.Apply2(f_3, (*Constructor_Data_Map_Internal_Node)(v_5.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node)(v_5.UnsafePtr).V3), gopurs_runtime.Apply(go__go_4_8_39, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v_5.UnsafePtr).V5)})))
goto end_branch_9
} else {

}
}
{
__t9 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_9:
return __t9
})
return go__go_4_8_39
})
}), gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(z_2 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_3_10_40 gopurs_runtime.Value
go__go_3_10_40 = gopurs_runtime.Func(func(__local_var_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(__local_var_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var __local_var_4_loop gopurs_runtime.Value = __local_var_4_loop_val
var __local_var_5_loop gopurs_runtime.Value = __local_var_5_loop_val
go__go_3_10_40:
for {
if false { continue go__go_3_10_40 }
var __local_var_4 gopurs_runtime.Value = __local_var_4_loop
_ = __local_var_4
var __local_var_5 gopurs_runtime.Value = __local_var_5_loop
_ = __local_var_5
var __t11 gopurs_runtime.Value
{
if (__local_var_5.Type == 9 && __local_var_5.IntVal == 324739070 && __local_var_5.UnsafePtr == nil) {
__t11 = __local_var_4
goto end_branch_11
} else {

}
}
{
if (__local_var_5.Type == 9 && __local_var_5.IntVal == 324739070 && __local_var_5.UnsafePtr != nil) {
__t11 = gopurs_runtime.UncurriedApp2(go__go_3_10_40, gopurs_runtime.Apply3(f_1, (*Constructor_Data_Map_Internal_Node)(__local_var_5.UnsafePtr).V2, gopurs_runtime.UncurriedApp2(go__go_3_10_40, __local_var_4, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_5.UnsafePtr).V4)}), (*Constructor_Data_Map_Internal_Node)(__local_var_5.UnsafePtr).V3), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_5.UnsafePtr).V5)})
goto end_branch_11
} else {

}
}
{
__t11 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_11:
return __t11
}
}()
})
})
return gopurs_runtime.Func(func(m_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.UncurriedApp2(go__go_3_10_40, z_2, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](m_4))})
})
})
}), gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(z_2 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_3_12_41 gopurs_runtime.Value
_ = go__go_3_12_41
go__go_3_12_41 = gopurs_runtime.Func2(func(__local_var_4 gopurs_runtime.Value, __local_var_5 gopurs_runtime.Value) gopurs_runtime.Value {
var __t13 gopurs_runtime.Value
{
if (__local_var_4.Type == 9 && __local_var_4.IntVal == 324739070 && __local_var_4.UnsafePtr == nil) {
__t13 = __local_var_5
goto end_branch_13
} else {

}
}
{
if (__local_var_4.Type == 9 && __local_var_4.IntVal == 324739070 && __local_var_4.UnsafePtr != nil) {
__t13 = gopurs_runtime.UncurriedApp2(go__go_3_12_41, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_4.UnsafePtr).V4)}, gopurs_runtime.Apply3(f_1, (*Constructor_Data_Map_Internal_Node)(__local_var_4.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node)(__local_var_4.UnsafePtr).V3, gopurs_runtime.UncurriedApp2(go__go_3_12_41, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_4.UnsafePtr).V5)}, __local_var_5)))
goto end_branch_13
} else {

}
}
{
__t13 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_13:
return __t13
})
return gopurs_runtime.Func(func(m_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.UncurriedApp2(go__go_3_12_41, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](m_4))}, z_2)
})
})
}))
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Functor0", "mapWithIndex", gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_3_14_42 gopurs_runtime.Value
_ = go__go_3_14_42
go__go_3_14_42 = gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t15 *Constructor_Data_Map_Internal_Node
{
if (v_4.Type == 9 && v_4.IntVal == 324739070 && v_4.UnsafePtr == nil) {
__t15 = (*Constructor_Data_Map_Internal_Node)(nil)
goto end_branch_15
} else {

}
}
{
if (v_4.Type == 9 && v_4.IntVal == 324739070 && v_4.UnsafePtr != nil) {
__t15 = &Constructor_Data_Map_Internal_Node{1, (*Constructor_Data_Map_Internal_Node)(v_4.UnsafePtr).V0, (*Constructor_Data_Map_Internal_Node)(v_4.UnsafePtr).V1, (*Constructor_Data_Map_Internal_Node)(v_4.UnsafePtr).V2, gopurs_runtime.Apply(f_2, (*Constructor_Data_Map_Internal_Node)(v_4.UnsafePtr).V3), gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.Apply(go__go_3_14_42, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v_4.UnsafePtr).V4)})), gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.Apply(go__go_3_14_42, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v_4.UnsafePtr).V5)}))}
goto end_branch_15
} else {

}
}
{
__t15 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_15:
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(__t15)}
})
return go__go_3_14_42
}))
}), gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_2_16_43 gopurs_runtime.Value
_ = go__go_2_16_43
go__go_2_16_43 = gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t17 *Constructor_Data_Map_Internal_Node
{
if (v_3.Type == 9 && v_3.IntVal == 324739070 && v_3.UnsafePtr == nil) {
__t17 = (*Constructor_Data_Map_Internal_Node)(nil)
goto end_branch_17
} else {

}
}
{
if (v_3.Type == 9 && v_3.IntVal == 324739070 && v_3.UnsafePtr != nil) {
__t17 = &Constructor_Data_Map_Internal_Node{1, (*Constructor_Data_Map_Internal_Node)(v_3.UnsafePtr).V0, (*Constructor_Data_Map_Internal_Node)(v_3.UnsafePtr).V1, (*Constructor_Data_Map_Internal_Node)(v_3.UnsafePtr).V2, gopurs_runtime.Apply2(f_1, (*Constructor_Data_Map_Internal_Node)(v_3.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node)(v_3.UnsafePtr).V3), gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.Apply(go__go_2_16_43, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v_3.UnsafePtr).V4)})), gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.Apply(go__go_2_16_43, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v_3.UnsafePtr).V5)}))}
goto end_branch_17
} else {

}
}
{
__t17 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_17:
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(__t17)}
})
return go__go_2_16_43
}))
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict4("Foldable1", "Functor0", "sequence", "traverse", gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict3("foldMap", "foldl", "foldr", gopurs_runtime.Func(func(dictMonoid_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Semigroup0_3_18 -> *Constructor_Data_Semigroup_Semigroup
Semigroup0_3_18 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_2, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_3_18
return gopurs_runtime.Func(func(f_4 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_5_19_44 gopurs_runtime.Value
_ = go__go_5_19_44
go__go_5_19_44 = gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
var __t20 gopurs_runtime.Value
{
if (v_6.Type == 9 && v_6.IntVal == 324739070 && v_6.UnsafePtr == nil) {
__t20 = gopurs_runtime.RecordGet(dictMonoid_2, "mempty")
goto end_branch_20
} else {

}
}
{
if (v_6.Type == 9 && v_6.IntVal == 324739070 && v_6.UnsafePtr != nil) {
__t20 = gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_3_18.V0), gopurs_runtime.Apply(go__go_5_19_44, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v_6.UnsafePtr).V4)}), gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_3_18.V0), gopurs_runtime.Apply(f_4, (*Constructor_Data_Map_Internal_Node)(v_6.UnsafePtr).V3), gopurs_runtime.Apply(go__go_5_19_44, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v_6.UnsafePtr).V5)})))
goto end_branch_20
} else {

}
}
{
__t20 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_20:
return __t20
})
return go__go_5_19_44
})
}), gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(z_3 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_4_21_45 gopurs_runtime.Value
go__go_4_21_45 = gopurs_runtime.Func(func(__local_var_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(__local_var_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var __local_var_5_loop gopurs_runtime.Value = __local_var_5_loop_val
var __local_var_6_loop gopurs_runtime.Value = __local_var_6_loop_val
go__go_4_21_45:
for {
if false { continue go__go_4_21_45 }
var __local_var_5 gopurs_runtime.Value = __local_var_5_loop
_ = __local_var_5
var __local_var_6 gopurs_runtime.Value = __local_var_6_loop
_ = __local_var_6
var __t22 gopurs_runtime.Value
{
if (__local_var_6.Type == 9 && __local_var_6.IntVal == 324739070 && __local_var_6.UnsafePtr == nil) {
__t22 = __local_var_5
goto end_branch_22
} else {

}
}
{
if (__local_var_6.Type == 9 && __local_var_6.IntVal == 324739070 && __local_var_6.UnsafePtr != nil) {
__t22 = gopurs_runtime.UncurriedApp2(go__go_4_21_45, gopurs_runtime.Apply2(f_2, gopurs_runtime.UncurriedApp2(go__go_4_21_45, __local_var_5, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_6.UnsafePtr).V4)}), (*Constructor_Data_Map_Internal_Node)(__local_var_6.UnsafePtr).V3), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_6.UnsafePtr).V5)})
goto end_branch_22
} else {

}
}
{
__t22 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_22:
return __t22
}
}()
})
})
return gopurs_runtime.Func(func(m_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.UncurriedApp2(go__go_4_21_45, z_3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](m_5))})
})
})
}), gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(z_3 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_4_23_46 gopurs_runtime.Value
_ = go__go_4_23_46
go__go_4_23_46 = gopurs_runtime.Func2(func(__local_var_5 gopurs_runtime.Value, __local_var_6 gopurs_runtime.Value) gopurs_runtime.Value {
var __t24 gopurs_runtime.Value
{
if (__local_var_5.Type == 9 && __local_var_5.IntVal == 324739070 && __local_var_5.UnsafePtr == nil) {
__t24 = __local_var_6
goto end_branch_24
} else {

}
}
{
if (__local_var_5.Type == 9 && __local_var_5.IntVal == 324739070 && __local_var_5.UnsafePtr != nil) {
__t24 = gopurs_runtime.UncurriedApp2(go__go_4_23_46, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_5.UnsafePtr).V4)}, gopurs_runtime.Apply2(f_2, (*Constructor_Data_Map_Internal_Node)(__local_var_5.UnsafePtr).V3, gopurs_runtime.UncurriedApp2(go__go_4_23_46, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_5.UnsafePtr).V5)}, __local_var_6)))
goto end_branch_24
} else {

}
}
{
__t24 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_24:
return __t24
})
return gopurs_runtime.Func(func(m_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.UncurriedApp2(go__go_4_23_46, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](m_5))}, z_3)
})
})
}))
}), gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_3_25_47 gopurs_runtime.Value
_ = go__go_3_25_47
go__go_3_25_47 = gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t26 *Constructor_Data_Map_Internal_Node
{
if (v_4.Type == 9 && v_4.IntVal == 324739070 && v_4.UnsafePtr == nil) {
__t26 = (*Constructor_Data_Map_Internal_Node)(nil)
goto end_branch_26
} else {

}
}
{
if (v_4.Type == 9 && v_4.IntVal == 324739070 && v_4.UnsafePtr != nil) {
__t26 = &Constructor_Data_Map_Internal_Node{1, (*Constructor_Data_Map_Internal_Node)(v_4.UnsafePtr).V0, (*Constructor_Data_Map_Internal_Node)(v_4.UnsafePtr).V1, (*Constructor_Data_Map_Internal_Node)(v_4.UnsafePtr).V2, gopurs_runtime.Apply(f_2, (*Constructor_Data_Map_Internal_Node)(v_4.UnsafePtr).V3), gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.Apply(go__go_3_25_47, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v_4.UnsafePtr).V4)})), gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.Apply(go__go_3_25_47, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v_4.UnsafePtr).V5)}))}
goto end_branch_26
} else {

}
}
{
__t26 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_26:
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(__t26)}
})
return go__go_3_25_47
}))
}), gopurs_runtime.Func(func(dictApplicative_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_Map_Internal_traversableMap(), "traverse"), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](dictApplicative_1))}, gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return x_2
}))
}), gopurs_runtime.Func(func(dictApplicative_1 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Apply0_2_27 -> *Constructor_Control_Apply_Apply
Apply0_2_27 := gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_1, "Apply0"), gopurs_runtime.Value{}))
_ = Apply0_2_27
// TAST (Let): Functor0_3_28 -> *Constructor_Data_Functor_Functor
Functor0_3_28 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_1, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_3_28
return gopurs_runtime.Func(func(f_4 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_5_29_48 gopurs_runtime.Value
_ = go__go_5_29_48
go__go_5_29_48 = gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
var __t33 gopurs_runtime.Value
{
if (v_6.Type == 9 && v_6.IntVal == 324739070 && v_6.UnsafePtr == nil) {
__t33 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_1, "pure"), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(nil))})
goto end_branch_33
} else {

}
}
{
if (v_6.Type == 9 && v_6.IntVal == 324739070 && v_6.UnsafePtr != nil) {
// TAST (Let): __local_var_7_30 -> gopurs_runtime.Value
var __local_var_7_30 gopurs_runtime.Value = gopurs_runtime.Int((*Constructor_Data_Map_Internal_Node)(v_6.UnsafePtr).V0)
// TAST (Let): __local_var_8_31 -> gopurs_runtime.Value
__local_var_8_31 := (*Constructor_Data_Map_Internal_Node)(v_6.UnsafePtr).V2
_ = __local_var_8_31
// TAST (Let): __local_var_9_32 -> gopurs_runtime.Value
var __local_var_9_32 gopurs_runtime.Value = gopurs_runtime.Int((*Constructor_Data_Map_Internal_Node)(v_6.UnsafePtr).V1)
__t33 = gopurs_runtime.Apply2(gopurs_runtime.Box(Apply0_2_27.V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Apply0_2_27.V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_3_28.V0), gopurs_runtime.Func(func(l_prime_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_prime_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(r_prime_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(&Constructor_Data_Map_Internal_Node{1, __local_var_7_30.IntVal, __local_var_9_32.IntVal, __local_var_8_31, v_prime_11, gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](l_prime_10), gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](r_prime_12)})}
})
})
}), gopurs_runtime.Apply(go__go_5_29_48, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v_6.UnsafePtr).V4)})), gopurs_runtime.Apply(f_4, (*Constructor_Data_Map_Internal_Node)(v_6.UnsafePtr).V3)), gopurs_runtime.Apply(go__go_5_29_48, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v_6.UnsafePtr).V5)}))
goto end_branch_33
} else {

}
}
{
__t33 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_33:
return __t33
})
return go__go_5_29_48
})
}))
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Apply0_1_34 -> *Constructor_Control_Apply_Apply
Apply0_1_34 := gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}))
_ = Apply0_1_34
// TAST (Let): Functor0_2_35 -> *Constructor_Data_Functor_Functor
Functor0_2_35 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_2_35
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_4_36_49 gopurs_runtime.Value
_ = go__go_4_36_49
go__go_4_36_49 = gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
var __t40 gopurs_runtime.Value
{
if (v_5.Type == 9 && v_5.IntVal == 324739070 && v_5.UnsafePtr == nil) {
__t40 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "pure"), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(nil))})
goto end_branch_40
} else {

}
}
{
if (v_5.Type == 9 && v_5.IntVal == 324739070 && v_5.UnsafePtr != nil) {
// TAST (Let): __local_var_6_37 -> gopurs_runtime.Value
var __local_var_6_37 gopurs_runtime.Value = gopurs_runtime.Int((*Constructor_Data_Map_Internal_Node)(v_5.UnsafePtr).V0)
// TAST (Let): __local_var_7_38 -> gopurs_runtime.Value
__local_var_7_38 := (*Constructor_Data_Map_Internal_Node)(v_5.UnsafePtr).V2
_ = __local_var_7_38
// TAST (Let): __local_var_8_39 -> gopurs_runtime.Value
var __local_var_8_39 gopurs_runtime.Value = gopurs_runtime.Int((*Constructor_Data_Map_Internal_Node)(v_5.UnsafePtr).V1)
__t40 = gopurs_runtime.Apply2(gopurs_runtime.Box(Apply0_1_34.V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Apply0_1_34.V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_2_35.V0), gopurs_runtime.Func(func(l_prime_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_prime_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(r_prime_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(&Constructor_Data_Map_Internal_Node{1, __local_var_6_37.IntVal, __local_var_8_39.IntVal, __local_var_7_38, v_prime_10, gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](l_prime_9), gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](r_prime_11)})}
})
})
}), gopurs_runtime.Apply(go__go_4_36_49, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v_5.UnsafePtr).V4)})), gopurs_runtime.Apply2(f_3, __local_var_7_38, (*Constructor_Data_Map_Internal_Node)(v_5.UnsafePtr).V3)), gopurs_runtime.Apply(go__go_4_36_49, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v_5.UnsafePtr).V5)}))
goto end_branch_40
} else {

}
}
{
__t40 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_40:
return __t40
})
return go__go_4_36_49
})
}))
	})
	return cache_Data_Map_Internal_traversableWithIndexMap
}

var cache_Data_Map_Internal_values gopurs_runtime.Value
var once_Data_Map_Internal_values sync.Once
func Get_Data_Map_Internal_values() gopurs_runtime.Value {
	once_Data_Map_Internal_values.Do(func() {
		cache_Data_Map_Internal_values = func() gopurs_runtime.Value {
var go__go_0_0_50 gopurs_runtime.Value
_ = go__go_0_0_50
go__go_0_0_50 = gopurs_runtime.Func2(func(__local_var_1 gopurs_runtime.Value, __local_var_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if (__local_var_1.Type == 9 && __local_var_1.IntVal == 324739070 && __local_var_1.UnsafePtr == nil) {
__t1 = __local_var_2
goto end_branch_1
} else {

}
}
{
if (__local_var_1.Type == 9 && __local_var_1.IntVal == 324739070 && __local_var_1.UnsafePtr != nil) {
__t1 = gopurs_runtime.UncurriedApp2(go__go_0_0_50, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_1.UnsafePtr).V4)}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, (*Constructor_Data_Map_Internal_Node)(__local_var_1.UnsafePtr).V3, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.UncurriedApp2(go__go_0_0_50, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_1.UnsafePtr).V5)}, __local_var_2))})})
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
return gopurs_runtime.Func(func(m_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.UncurriedApp2(go__go_0_0_50, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](m_1))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(nil))})
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
return Call_Data_Map_Internal_foldSubmapBy(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord](dictOrd_0_box), appendFn_1_box, memptyValue_2_box, gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](kmin_3_box), gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](kmax_4_box), f_5_box)
})
	})
	return cache_Data_Map_Internal_foldSubmapBy
}

var cache_Data_Map_Internal_foldSubmap gopurs_runtime.Value
var once_Data_Map_Internal_foldSubmap sync.Once
func Get_Data_Map_Internal_foldSubmap() gopurs_runtime.Value {
	once_Data_Map_Internal_foldSubmap.Do(func() {
		cache_Data_Map_Internal_foldSubmap = gopurs_runtime.Func2(func(dictOrd_0_box gopurs_runtime.Value, dictMonoid_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Map_Internal_foldSubmap(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord](dictOrd_0_box), gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid](dictMonoid_1_box))
})
	})
	return cache_Data_Map_Internal_foldSubmap
}

var cache_Data_Map_Internal_findMin gopurs_runtime.Value
var once_Data_Map_Internal_findMin sync.Once
func Get_Data_Map_Internal_findMin() gopurs_runtime.Value {
	once_Data_Map_Internal_findMin.Do(func() {
		cache_Data_Map_Internal_findMin = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_Data_Map_Internal_findMin(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](v_0_box)))}
})
	})
	return cache_Data_Map_Internal_findMin
}

var cache_Data_Map_Internal_lookupGT gopurs_runtime.Value
var once_Data_Map_Internal_lookupGT sync.Once
func Get_Data_Map_Internal_lookupGT() gopurs_runtime.Value {
	once_Data_Map_Internal_lookupGT.Do(func() {
		cache_Data_Map_Internal_lookupGT = gopurs_runtime.Func2(func(dictOrd_0_box gopurs_runtime.Value, k_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Map_Internal_lookupGT(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord](dictOrd_0_box), k_1_box)
})
	})
	return cache_Data_Map_Internal_lookupGT
}

var cache_Data_Map_Internal_findMax gopurs_runtime.Value
var once_Data_Map_Internal_findMax sync.Once
func Get_Data_Map_Internal_findMax() gopurs_runtime.Value {
	once_Data_Map_Internal_findMax.Do(func() {
		cache_Data_Map_Internal_findMax = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_Data_Map_Internal_findMax(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](v_0_box)))}
})
	})
	return cache_Data_Map_Internal_findMax
}

var cache_Data_Map_Internal_lookupLT gopurs_runtime.Value
var once_Data_Map_Internal_lookupLT sync.Once
func Get_Data_Map_Internal_lookupLT() gopurs_runtime.Value {
	once_Data_Map_Internal_lookupLT.Do(func() {
		cache_Data_Map_Internal_lookupLT = gopurs_runtime.Func2(func(dictOrd_0_box gopurs_runtime.Value, k_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Map_Internal_lookupLT(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord](dictOrd_0_box), k_1_box)
})
	})
	return cache_Data_Map_Internal_lookupLT
}

var cache_Data_Map_Internal_filterWithKey gopurs_runtime.Value
var once_Data_Map_Internal_filterWithKey sync.Once
func Get_Data_Map_Internal_filterWithKey() gopurs_runtime.Value {
	once_Data_Map_Internal_filterWithKey.Do(func() {
		cache_Data_Map_Internal_filterWithKey = gopurs_runtime.Func2(func(dictOrd_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Map_Internal_filterWithKey(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord](dictOrd_0_box), f_1_box)
})
	})
	return cache_Data_Map_Internal_filterWithKey
}

var cache_Data_Map_Internal_filterKeys gopurs_runtime.Value
var once_Data_Map_Internal_filterKeys sync.Once
func Get_Data_Map_Internal_filterKeys() gopurs_runtime.Value {
	once_Data_Map_Internal_filterKeys.Do(func() {
		cache_Data_Map_Internal_filterKeys = gopurs_runtime.Func2(func(dictOrd_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Map_Internal_filterKeys(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord](dictOrd_0_box), f_1_box)
})
	})
	return cache_Data_Map_Internal_filterKeys
}

var cache_Data_Map_Internal_filter gopurs_runtime.Value
var once_Data_Map_Internal_filter sync.Once
func Get_Data_Map_Internal_filter() gopurs_runtime.Value {
	once_Data_Map_Internal_filter.Do(func() {
		cache_Data_Map_Internal_filter = gopurs_runtime.Func2(func(dictOrd_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Map_Internal_filter(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord](dictOrd_0_box), x_1_box)
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
		cache_Data_Map_Internal_empty = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(nil))}
	})
	return cache_Data_Map_Internal_empty
}

var cache_Data_Map_Internal_fromFoldable gopurs_runtime.Value
var once_Data_Map_Internal_fromFoldable sync.Once
func Get_Data_Map_Internal_fromFoldable() gopurs_runtime.Value {
	once_Data_Map_Internal_fromFoldable.Do(func() {
		cache_Data_Map_Internal_fromFoldable = gopurs_runtime.Func2(func(dictOrd_0_box gopurs_runtime.Value, dictFoldable_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Map_Internal_fromFoldable(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord](dictOrd_0_box), gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable](dictFoldable_1_box))
})
	})
	return cache_Data_Map_Internal_fromFoldable
}

var cache_Data_Map_Internal_fromFoldableWith gopurs_runtime.Value
var once_Data_Map_Internal_fromFoldableWith sync.Once
func Get_Data_Map_Internal_fromFoldableWith() gopurs_runtime.Value {
	once_Data_Map_Internal_fromFoldableWith.Do(func() {
		cache_Data_Map_Internal_fromFoldableWith = gopurs_runtime.Func3(func(dictOrd_0_box gopurs_runtime.Value, dictFoldable_1_box gopurs_runtime.Value, f_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Map_Internal_fromFoldableWith(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord](dictOrd_0_box), gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable](dictFoldable_1_box), f_2_box)
})
	})
	return cache_Data_Map_Internal_fromFoldableWith
}

var cache_Data_Map_Internal_fromFoldableWithIndex gopurs_runtime.Value
var once_Data_Map_Internal_fromFoldableWithIndex sync.Once
func Get_Data_Map_Internal_fromFoldableWithIndex() gopurs_runtime.Value {
	once_Data_Map_Internal_fromFoldableWithIndex.Do(func() {
		cache_Data_Map_Internal_fromFoldableWithIndex = gopurs_runtime.Func2(func(dictOrd_0_box gopurs_runtime.Value, dictFoldableWithIndex_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Map_Internal_fromFoldableWithIndex(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord](dictOrd_0_box), gopurs_runtime.CoerceToStruct[Constructor_Data_FoldableWithIndex_FoldableWithIndex](dictFoldableWithIndex_1_box))
})
	})
	return cache_Data_Map_Internal_fromFoldableWithIndex
}

var cache_Data_Map_Internal_monoidSemigroupMap gopurs_runtime.Value
var once_Data_Map_Internal_monoidSemigroupMap sync.Once
func Get_Data_Map_Internal_monoidSemigroupMap() gopurs_runtime.Value {
	once_Data_Map_Internal_monoidSemigroupMap.Do(func() {
		cache_Data_Map_Internal_monoidSemigroupMap = gopurs_runtime.Func3(func(_dollar__unused_0_box gopurs_runtime.Value, dictOrd_1_box gopurs_runtime.Value, dictSemigroup_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Map_Internal_monoidSemigroupMap(_dollar__unused_0_box, dictOrd_1_box, dictSemigroup_2_box)
})
	})
	return cache_Data_Map_Internal_monoidSemigroupMap
}

var cache_Data_Map_Internal_submap gopurs_runtime.Value
var once_Data_Map_Internal_submap sync.Once
func Get_Data_Map_Internal_submap() gopurs_runtime.Value {
	once_Data_Map_Internal_submap.Do(func() {
		cache_Data_Map_Internal_submap = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Map_Internal_submap(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord](dictOrd_0_box))
})
	})
	return cache_Data_Map_Internal_submap
}

var cache_Data_Map_Internal_unions gopurs_runtime.Value
var once_Data_Map_Internal_unions sync.Once
func Get_Data_Map_Internal_unions() gopurs_runtime.Value {
	once_Data_Map_Internal_unions.Do(func() {
		cache_Data_Map_Internal_unions = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Map_Internal_unions(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord](dictOrd_0_box))
})
	})
	return cache_Data_Map_Internal_unions
}

var cache_Data_Map_Internal_difference gopurs_runtime.Value
var once_Data_Map_Internal_difference sync.Once
func Get_Data_Map_Internal_difference() gopurs_runtime.Value {
	once_Data_Map_Internal_difference.Do(func() {
		cache_Data_Map_Internal_difference = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Map_Internal_difference(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord](dictOrd_0_box))
})
	})
	return cache_Data_Map_Internal_difference
}

var cache_Data_Map_Internal_delete gopurs_runtime.Value
var once_Data_Map_Internal_delete sync.Once
func Get_Data_Map_Internal_delete() gopurs_runtime.Value {
	once_Data_Map_Internal_delete.Do(func() {
		cache_Data_Map_Internal_delete = gopurs_runtime.Func2(func(dictOrd_0_box gopurs_runtime.Value, k_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Map_Internal_delete(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord](dictOrd_0_box), k_1_box)
})
	})
	return cache_Data_Map_Internal_delete
}

var cache_Data_Map_Internal_checkValid gopurs_runtime.Value
var once_Data_Map_Internal_checkValid sync.Once
func Get_Data_Map_Internal_checkValid() gopurs_runtime.Value {
	once_Data_Map_Internal_checkValid.Do(func() {
		cache_Data_Map_Internal_checkValid = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Map_Internal_checkValid(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord](dictOrd_0_box))
})
	})
	return cache_Data_Map_Internal_checkValid
}

var cache_Data_Map_Internal_catMaybes gopurs_runtime.Value
var once_Data_Map_Internal_catMaybes sync.Once
func Get_Data_Map_Internal_catMaybes() gopurs_runtime.Value {
	once_Data_Map_Internal_catMaybes.Do(func() {
		cache_Data_Map_Internal_catMaybes = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Map_Internal_catMaybes(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord](dictOrd_0_box))
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
return Call_Data_Map_Internal_alter(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord](dictOrd_0_box))
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

var cache_Data_Map_Internal_alter__2325420954 gopurs_runtime.Value
var once_Data_Map_Internal_alter__2325420954 sync.Once
func Get_Data_Map_Internal_alter__2325420954() gopurs_runtime.Value {
	once_Data_Map_Internal_alter__2325420954.Do(func() {
		cache_Data_Map_Internal_alter__2325420954 = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Map_Internal_alter__2325420954(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord](dictOrd_0_box))
})
	})
	return cache_Data_Map_Internal_alter__2325420954
}

var cache_Data_Map_Internal_alter__1204655226 gopurs_runtime.Value
var once_Data_Map_Internal_alter__1204655226 sync.Once
func Get_Data_Map_Internal_alter__1204655226() gopurs_runtime.Value {
	once_Data_Map_Internal_alter__1204655226.Do(func() {
		cache_Data_Map_Internal_alter__1204655226 = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Map_Internal_alter__1204655226(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord](dictOrd_0_box))
})
	})
	return cache_Data_Map_Internal_alter__1204655226
}

var cache_Data_Map_Internal_empty__2198260019 gopurs_runtime.Value
var once_Data_Map_Internal_empty__2198260019 sync.Once
func Get_Data_Map_Internal_empty__2198260019() gopurs_runtime.Value {
	once_Data_Map_Internal_empty__2198260019.Do(func() {
		cache_Data_Map_Internal_empty__2198260019 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(nil))}
	})
	return cache_Data_Map_Internal_empty__2198260019
}

var cache_Data_Map_Internal_empty__1818220131 gopurs_runtime.Value
var once_Data_Map_Internal_empty__1818220131 sync.Once
func Get_Data_Map_Internal_empty__1818220131() gopurs_runtime.Value {
	once_Data_Map_Internal_empty__1818220131.Do(func() {
		cache_Data_Map_Internal_empty__1818220131 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(nil))}
	})
	return cache_Data_Map_Internal_empty__1818220131
}

var cache_Data_Map_Internal_empty__1299254065 gopurs_runtime.Value
var once_Data_Map_Internal_empty__1299254065 sync.Once
func Get_Data_Map_Internal_empty__1299254065() gopurs_runtime.Value {
	once_Data_Map_Internal_empty__1299254065.Do(func() {
		cache_Data_Map_Internal_empty__1299254065 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(nil))}
	})
	return cache_Data_Map_Internal_empty__1299254065
}

var cache_Data_Map_Internal_empty__1794046843 gopurs_runtime.Value
var once_Data_Map_Internal_empty__1794046843 sync.Once
func Get_Data_Map_Internal_empty__1794046843() gopurs_runtime.Value {
	once_Data_Map_Internal_empty__1794046843.Do(func() {
		cache_Data_Map_Internal_empty__1794046843 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(nil))}
	})
	return cache_Data_Map_Internal_empty__1794046843
}

var cache_Data_Map_Internal_findMax__2266220649 gopurs_runtime.Value
var once_Data_Map_Internal_findMax__2266220649 sync.Once
func Get_Data_Map_Internal_findMax__2266220649() gopurs_runtime.Value {
	once_Data_Map_Internal_findMax__2266220649.Do(func() {
		cache_Data_Map_Internal_findMax__2266220649 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_Data_Map_Internal_findMax__2266220649(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](v_0_box)))}
})
	})
	return cache_Data_Map_Internal_findMax__2266220649
}

var cache_Data_Map_Internal_findMax__528468393 gopurs_runtime.Value
var once_Data_Map_Internal_findMax__528468393 sync.Once
func Get_Data_Map_Internal_findMax__528468393() gopurs_runtime.Value {
	once_Data_Map_Internal_findMax__528468393.Do(func() {
		cache_Data_Map_Internal_findMax__528468393 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_Data_Map_Internal_findMax__528468393(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](v_0_box)))}
})
	})
	return cache_Data_Map_Internal_findMax__528468393
}

var cache_Data_Map_Internal_findMin__2266220649 gopurs_runtime.Value
var once_Data_Map_Internal_findMin__2266220649 sync.Once
func Get_Data_Map_Internal_findMin__2266220649() gopurs_runtime.Value {
	once_Data_Map_Internal_findMin__2266220649.Do(func() {
		cache_Data_Map_Internal_findMin__2266220649 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_Data_Map_Internal_findMin__2266220649(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](v_0_box)))}
})
	})
	return cache_Data_Map_Internal_findMin__2266220649
}

var cache_Data_Map_Internal_findMin__528468393 gopurs_runtime.Value
var once_Data_Map_Internal_findMin__528468393 sync.Once
func Get_Data_Map_Internal_findMin__528468393() gopurs_runtime.Value {
	once_Data_Map_Internal_findMin__528468393.Do(func() {
		cache_Data_Map_Internal_findMin__528468393 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_Data_Map_Internal_findMin__528468393(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](v_0_box)))}
})
	})
	return cache_Data_Map_Internal_findMin__528468393
}

var cache_Data_Map_Internal_foldSubmapBy__3050108409 gopurs_runtime.Value
var once_Data_Map_Internal_foldSubmapBy__3050108409 sync.Once
func Get_Data_Map_Internal_foldSubmapBy__3050108409() gopurs_runtime.Value {
	once_Data_Map_Internal_foldSubmapBy__3050108409.Do(func() {
		cache_Data_Map_Internal_foldSubmapBy__3050108409 = gopurs_runtime.Func6(func(dictOrd_0_box gopurs_runtime.Value, appendFn_1_box gopurs_runtime.Value, memptyValue_2_box gopurs_runtime.Value, kmin_3_box gopurs_runtime.Value, kmax_4_box gopurs_runtime.Value, f_5_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Map_Internal_foldSubmapBy__3050108409(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord](dictOrd_0_box), appendFn_1_box, memptyValue_2_box, gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](kmin_3_box), gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](kmax_4_box), f_5_box)
})
	})
	return cache_Data_Map_Internal_foldSubmapBy__3050108409
}

var cache_Data_Map_Internal_foldSubmapBy__3128450809 gopurs_runtime.Value
var once_Data_Map_Internal_foldSubmapBy__3128450809 sync.Once
func Get_Data_Map_Internal_foldSubmapBy__3128450809() gopurs_runtime.Value {
	once_Data_Map_Internal_foldSubmapBy__3128450809.Do(func() {
		cache_Data_Map_Internal_foldSubmapBy__3128450809 = gopurs_runtime.Func6(func(dictOrd_0_box gopurs_runtime.Value, appendFn_1_box gopurs_runtime.Value, memptyValue_2_box gopurs_runtime.Value, kmin_3_box gopurs_runtime.Value, kmax_4_box gopurs_runtime.Value, f_5_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Map_Internal_foldSubmapBy__3128450809(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord](dictOrd_0_box), appendFn_1_box, gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](memptyValue_2_box), gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](kmin_3_box), gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](kmax_4_box), f_5_box)
})
	})
	return cache_Data_Map_Internal_foldSubmapBy__3128450809
}

var cache_Data_Map_Internal_foldableMap__767959947 gopurs_runtime.Value
var once_Data_Map_Internal_foldableMap__767959947 sync.Once
func Get_Data_Map_Internal_foldableMap__767959947() gopurs_runtime.Value {
	once_Data_Map_Internal_foldableMap__767959947.Do(func() {
		cache_Data_Map_Internal_foldableMap__767959947 = gopurs_runtime.RecordDict3("foldMap", "foldl", "foldr", gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Semigroup0_1_0 -> *Constructor_Data_Semigroup_Semigroup
Semigroup0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_1_0
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_3_1_83 gopurs_runtime.Value
_ = go__go_3_1_83
go__go_3_1_83 = gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
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
__t2 = gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_1_0.V0), gopurs_runtime.Apply(go__go_3_1_83, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v_4.UnsafePtr).V4)}), gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_1_0.V0), gopurs_runtime.Apply(f_2, (*Constructor_Data_Map_Internal_Node)(v_4.UnsafePtr).V3), gopurs_runtime.Apply(go__go_3_1_83, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v_4.UnsafePtr).V5)})))
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
return go__go_3_1_83
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(z_1 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_2_3_84 gopurs_runtime.Value
go__go_2_3_84 = gopurs_runtime.Func(func(__local_var_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(__local_var_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var __local_var_3_loop gopurs_runtime.Value = __local_var_3_loop_val
var __local_var_4_loop gopurs_runtime.Value = __local_var_4_loop_val
go__go_2_3_84:
for {
if false { continue go__go_2_3_84 }
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
__t4 = gopurs_runtime.UncurriedApp2(go__go_2_3_84, gopurs_runtime.Apply2(f_0, gopurs_runtime.UncurriedApp2(go__go_2_3_84, __local_var_3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_4.UnsafePtr).V4)}), (*Constructor_Data_Map_Internal_Node)(__local_var_4.UnsafePtr).V3), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_4.UnsafePtr).V5)})
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
return gopurs_runtime.UncurriedApp2(go__go_2_3_84, z_1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](m_3))})
})
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(z_1 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_2_5_85 gopurs_runtime.Value
_ = go__go_2_5_85
go__go_2_5_85 = gopurs_runtime.Func2(func(__local_var_3 gopurs_runtime.Value, __local_var_4 gopurs_runtime.Value) gopurs_runtime.Value {
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
__t6 = gopurs_runtime.UncurriedApp2(go__go_2_5_85, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_3.UnsafePtr).V4)}, gopurs_runtime.Apply2(f_0, (*Constructor_Data_Map_Internal_Node)(__local_var_3.UnsafePtr).V3, gopurs_runtime.UncurriedApp2(go__go_2_5_85, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_3.UnsafePtr).V5)}, __local_var_4)))
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
return gopurs_runtime.UncurriedApp2(go__go_2_5_85, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](m_3))}, z_1)
})
})
}))
	})
	return cache_Data_Map_Internal_foldableMap__767959947
}

var cache_Data_Map_Internal_foldableMap__373570208 gopurs_runtime.Value
var once_Data_Map_Internal_foldableMap__373570208 sync.Once
func Get_Data_Map_Internal_foldableMap__373570208() gopurs_runtime.Value {
	once_Data_Map_Internal_foldableMap__373570208.Do(func() {
		cache_Data_Map_Internal_foldableMap__373570208 = gopurs_runtime.RecordDict3("foldMap", "foldl", "foldr", gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Semigroup0_1_0 -> *Constructor_Data_Semigroup_Semigroup
Semigroup0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_1_0
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_3_1_86 gopurs_runtime.Value
_ = go__go_3_1_86
go__go_3_1_86 = gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
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
__t2 = gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_1_0.V0), gopurs_runtime.Apply(go__go_3_1_86, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v_4.UnsafePtr).V4)}), gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_1_0.V0), gopurs_runtime.Apply(f_2, (*Constructor_Data_Map_Internal_Node)(v_4.UnsafePtr).V3), gopurs_runtime.Apply(go__go_3_1_86, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v_4.UnsafePtr).V5)})))
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
return go__go_3_1_86
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(z_1 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_2_3_87 gopurs_runtime.Value
go__go_2_3_87 = gopurs_runtime.Func(func(__local_var_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(__local_var_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var __local_var_3_loop gopurs_runtime.Value = __local_var_3_loop_val
var __local_var_4_loop gopurs_runtime.Value = __local_var_4_loop_val
go__go_2_3_87:
for {
if false { continue go__go_2_3_87 }
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
__t4 = gopurs_runtime.UncurriedApp2(go__go_2_3_87, gopurs_runtime.Apply2(f_0, gopurs_runtime.UncurriedApp2(go__go_2_3_87, __local_var_3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_4.UnsafePtr).V4)}), (*Constructor_Data_Map_Internal_Node)(__local_var_4.UnsafePtr).V3), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_4.UnsafePtr).V5)})
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
return gopurs_runtime.UncurriedApp2(go__go_2_3_87, z_1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](m_3))})
})
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(z_1 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_2_5_88 gopurs_runtime.Value
_ = go__go_2_5_88
go__go_2_5_88 = gopurs_runtime.Func2(func(__local_var_3 gopurs_runtime.Value, __local_var_4 gopurs_runtime.Value) gopurs_runtime.Value {
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
__t6 = gopurs_runtime.UncurriedApp2(go__go_2_5_88, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_3.UnsafePtr).V4)}, gopurs_runtime.Apply2(f_0, (*Constructor_Data_Map_Internal_Node)(__local_var_3.UnsafePtr).V3, gopurs_runtime.UncurriedApp2(go__go_2_5_88, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_3.UnsafePtr).V5)}, __local_var_4)))
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
return gopurs_runtime.UncurriedApp2(go__go_2_5_88, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](m_3))}, z_1)
})
})
}))
	})
	return cache_Data_Map_Internal_foldableMap__373570208
}

var cache_Data_Map_Internal_foldableWithIndexMap__1634502082 gopurs_runtime.Value
var once_Data_Map_Internal_foldableWithIndexMap__1634502082 sync.Once
func Get_Data_Map_Internal_foldableWithIndexMap__1634502082() gopurs_runtime.Value {
	once_Data_Map_Internal_foldableWithIndexMap__1634502082.Do(func() {
		cache_Data_Map_Internal_foldableWithIndexMap__1634502082 = gopurs_runtime.RecordDict4("Foldable0", "foldMapWithIndex", "foldlWithIndex", "foldrWithIndex", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict3("foldMap", "foldl", "foldr", gopurs_runtime.Func(func(dictMonoid_1 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Semigroup0_2_0 -> *Constructor_Data_Semigroup_Semigroup
Semigroup0_2_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_1, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_2_0
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_4_1_89 gopurs_runtime.Value
_ = go__go_4_1_89
go__go_4_1_89 = gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 gopurs_runtime.Value
{
if (v_5.Type == 9 && v_5.IntVal == 324739070 && v_5.UnsafePtr == nil) {
__t2 = gopurs_runtime.RecordGet(dictMonoid_1, "mempty")
goto end_branch_2
} else {

}
}
{
if (v_5.Type == 9 && v_5.IntVal == 324739070 && v_5.UnsafePtr != nil) {
__t2 = gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_2_0.V0), gopurs_runtime.Apply(go__go_4_1_89, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v_5.UnsafePtr).V4)}), gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_2_0.V0), gopurs_runtime.Apply(f_3, (*Constructor_Data_Map_Internal_Node)(v_5.UnsafePtr).V3), gopurs_runtime.Apply(go__go_4_1_89, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v_5.UnsafePtr).V5)})))
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
return go__go_4_1_89
})
}), gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(z_2 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_3_3_90 gopurs_runtime.Value
go__go_3_3_90 = gopurs_runtime.Func(func(__local_var_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(__local_var_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var __local_var_4_loop gopurs_runtime.Value = __local_var_4_loop_val
var __local_var_5_loop gopurs_runtime.Value = __local_var_5_loop_val
go__go_3_3_90:
for {
if false { continue go__go_3_3_90 }
var __local_var_4 gopurs_runtime.Value = __local_var_4_loop
_ = __local_var_4
var __local_var_5 gopurs_runtime.Value = __local_var_5_loop
_ = __local_var_5
var __t4 gopurs_runtime.Value
{
if (__local_var_5.Type == 9 && __local_var_5.IntVal == 324739070 && __local_var_5.UnsafePtr == nil) {
__t4 = __local_var_4
goto end_branch_4
} else {

}
}
{
if (__local_var_5.Type == 9 && __local_var_5.IntVal == 324739070 && __local_var_5.UnsafePtr != nil) {
__t4 = gopurs_runtime.UncurriedApp2(go__go_3_3_90, gopurs_runtime.Apply2(f_1, gopurs_runtime.UncurriedApp2(go__go_3_3_90, __local_var_4, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_5.UnsafePtr).V4)}), (*Constructor_Data_Map_Internal_Node)(__local_var_5.UnsafePtr).V3), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_5.UnsafePtr).V5)})
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
return gopurs_runtime.Func(func(m_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.UncurriedApp2(go__go_3_3_90, z_2, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](m_4))})
})
})
}), gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(z_2 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_3_5_91 gopurs_runtime.Value
_ = go__go_3_5_91
go__go_3_5_91 = gopurs_runtime.Func2(func(__local_var_4 gopurs_runtime.Value, __local_var_5 gopurs_runtime.Value) gopurs_runtime.Value {
var __t6 gopurs_runtime.Value
{
if (__local_var_4.Type == 9 && __local_var_4.IntVal == 324739070 && __local_var_4.UnsafePtr == nil) {
__t6 = __local_var_5
goto end_branch_6
} else {

}
}
{
if (__local_var_4.Type == 9 && __local_var_4.IntVal == 324739070 && __local_var_4.UnsafePtr != nil) {
__t6 = gopurs_runtime.UncurriedApp2(go__go_3_5_91, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_4.UnsafePtr).V4)}, gopurs_runtime.Apply2(f_1, (*Constructor_Data_Map_Internal_Node)(__local_var_4.UnsafePtr).V3, gopurs_runtime.UncurriedApp2(go__go_3_5_91, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_4.UnsafePtr).V5)}, __local_var_5)))
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
return gopurs_runtime.Func(func(m_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.UncurriedApp2(go__go_3_5_91, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](m_4))}, z_2)
})
})
}))
}), gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Semigroup0_1_7 -> *Constructor_Data_Semigroup_Semigroup
Semigroup0_1_7 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_1_7
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_3_8_92 gopurs_runtime.Value
_ = go__go_3_8_92
go__go_3_8_92 = gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t9 gopurs_runtime.Value
{
if (v_4.Type == 9 && v_4.IntVal == 324739070 && v_4.UnsafePtr == nil) {
__t9 = gopurs_runtime.RecordGet(dictMonoid_0, "mempty")
goto end_branch_9
} else {

}
}
{
if (v_4.Type == 9 && v_4.IntVal == 324739070 && v_4.UnsafePtr != nil) {
__t9 = gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_1_7.V0), gopurs_runtime.Apply(go__go_3_8_92, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v_4.UnsafePtr).V4)}), gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_1_7.V0), gopurs_runtime.Apply2(f_2, (*Constructor_Data_Map_Internal_Node)(v_4.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node)(v_4.UnsafePtr).V3), gopurs_runtime.Apply(go__go_3_8_92, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v_4.UnsafePtr).V5)})))
goto end_branch_9
} else {

}
}
{
__t9 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_9:
return __t9
})
return go__go_3_8_92
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(z_1 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_2_10_93 gopurs_runtime.Value
go__go_2_10_93 = gopurs_runtime.Func(func(__local_var_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(__local_var_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var __local_var_3_loop gopurs_runtime.Value = __local_var_3_loop_val
var __local_var_4_loop gopurs_runtime.Value = __local_var_4_loop_val
go__go_2_10_93:
for {
if false { continue go__go_2_10_93 }
var __local_var_3 gopurs_runtime.Value = __local_var_3_loop
_ = __local_var_3
var __local_var_4 gopurs_runtime.Value = __local_var_4_loop
_ = __local_var_4
var __t11 gopurs_runtime.Value
{
if (__local_var_4.Type == 9 && __local_var_4.IntVal == 324739070 && __local_var_4.UnsafePtr == nil) {
__t11 = __local_var_3
goto end_branch_11
} else {

}
}
{
if (__local_var_4.Type == 9 && __local_var_4.IntVal == 324739070 && __local_var_4.UnsafePtr != nil) {
__t11 = gopurs_runtime.UncurriedApp2(go__go_2_10_93, gopurs_runtime.Apply3(f_0, (*Constructor_Data_Map_Internal_Node)(__local_var_4.UnsafePtr).V2, gopurs_runtime.UncurriedApp2(go__go_2_10_93, __local_var_3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_4.UnsafePtr).V4)}), (*Constructor_Data_Map_Internal_Node)(__local_var_4.UnsafePtr).V3), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_4.UnsafePtr).V5)})
goto end_branch_11
} else {

}
}
{
__t11 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_11:
return __t11
}
}()
})
})
return gopurs_runtime.Func(func(m_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.UncurriedApp2(go__go_2_10_93, z_1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](m_3))})
})
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(z_1 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_2_12_94 gopurs_runtime.Value
_ = go__go_2_12_94
go__go_2_12_94 = gopurs_runtime.Func2(func(__local_var_3 gopurs_runtime.Value, __local_var_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t13 gopurs_runtime.Value
{
if (__local_var_3.Type == 9 && __local_var_3.IntVal == 324739070 && __local_var_3.UnsafePtr == nil) {
__t13 = __local_var_4
goto end_branch_13
} else {

}
}
{
if (__local_var_3.Type == 9 && __local_var_3.IntVal == 324739070 && __local_var_3.UnsafePtr != nil) {
__t13 = gopurs_runtime.UncurriedApp2(go__go_2_12_94, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_3.UnsafePtr).V4)}, gopurs_runtime.Apply3(f_0, (*Constructor_Data_Map_Internal_Node)(__local_var_3.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node)(__local_var_3.UnsafePtr).V3, gopurs_runtime.UncurriedApp2(go__go_2_12_94, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_3.UnsafePtr).V5)}, __local_var_4)))
goto end_branch_13
} else {

}
}
{
__t13 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_13:
return __t13
})
return gopurs_runtime.Func(func(m_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.UncurriedApp2(go__go_2_12_94, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](m_3))}, z_1)
})
})
}))
	})
	return cache_Data_Map_Internal_foldableWithIndexMap__1634502082
}

var cache_Data_Map_Internal_foldableWithIndexMap__1035756962 gopurs_runtime.Value
var once_Data_Map_Internal_foldableWithIndexMap__1035756962 sync.Once
func Get_Data_Map_Internal_foldableWithIndexMap__1035756962() gopurs_runtime.Value {
	once_Data_Map_Internal_foldableWithIndexMap__1035756962.Do(func() {
		cache_Data_Map_Internal_foldableWithIndexMap__1035756962 = gopurs_runtime.RecordDict4("Foldable0", "foldMapWithIndex", "foldlWithIndex", "foldrWithIndex", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict3("foldMap", "foldl", "foldr", gopurs_runtime.Func(func(dictMonoid_1 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Semigroup0_2_0 -> *Constructor_Data_Semigroup_Semigroup
Semigroup0_2_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_1, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_2_0
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_4_1_95 gopurs_runtime.Value
_ = go__go_4_1_95
go__go_4_1_95 = gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 gopurs_runtime.Value
{
if (v_5.Type == 9 && v_5.IntVal == 324739070 && v_5.UnsafePtr == nil) {
__t2 = gopurs_runtime.RecordGet(dictMonoid_1, "mempty")
goto end_branch_2
} else {

}
}
{
if (v_5.Type == 9 && v_5.IntVal == 324739070 && v_5.UnsafePtr != nil) {
__t2 = gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_2_0.V0), gopurs_runtime.Apply(go__go_4_1_95, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v_5.UnsafePtr).V4)}), gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_2_0.V0), gopurs_runtime.Apply(f_3, (*Constructor_Data_Map_Internal_Node)(v_5.UnsafePtr).V3), gopurs_runtime.Apply(go__go_4_1_95, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v_5.UnsafePtr).V5)})))
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
return go__go_4_1_95
})
}), gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(z_2 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_3_3_96 gopurs_runtime.Value
go__go_3_3_96 = gopurs_runtime.Func(func(__local_var_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(__local_var_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var __local_var_4_loop gopurs_runtime.Value = __local_var_4_loop_val
var __local_var_5_loop gopurs_runtime.Value = __local_var_5_loop_val
go__go_3_3_96:
for {
if false { continue go__go_3_3_96 }
var __local_var_4 gopurs_runtime.Value = __local_var_4_loop
_ = __local_var_4
var __local_var_5 gopurs_runtime.Value = __local_var_5_loop
_ = __local_var_5
var __t4 gopurs_runtime.Value
{
if (__local_var_5.Type == 9 && __local_var_5.IntVal == 324739070 && __local_var_5.UnsafePtr == nil) {
__t4 = __local_var_4
goto end_branch_4
} else {

}
}
{
if (__local_var_5.Type == 9 && __local_var_5.IntVal == 324739070 && __local_var_5.UnsafePtr != nil) {
__t4 = gopurs_runtime.UncurriedApp2(go__go_3_3_96, gopurs_runtime.Apply2(f_1, gopurs_runtime.UncurriedApp2(go__go_3_3_96, __local_var_4, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_5.UnsafePtr).V4)}), (*Constructor_Data_Map_Internal_Node)(__local_var_5.UnsafePtr).V3), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_5.UnsafePtr).V5)})
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
return gopurs_runtime.Func(func(m_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.UncurriedApp2(go__go_3_3_96, z_2, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](m_4))})
})
})
}), gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(z_2 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_3_5_97 gopurs_runtime.Value
_ = go__go_3_5_97
go__go_3_5_97 = gopurs_runtime.Func2(func(__local_var_4 gopurs_runtime.Value, __local_var_5 gopurs_runtime.Value) gopurs_runtime.Value {
var __t6 gopurs_runtime.Value
{
if (__local_var_4.Type == 9 && __local_var_4.IntVal == 324739070 && __local_var_4.UnsafePtr == nil) {
__t6 = __local_var_5
goto end_branch_6
} else {

}
}
{
if (__local_var_4.Type == 9 && __local_var_4.IntVal == 324739070 && __local_var_4.UnsafePtr != nil) {
__t6 = gopurs_runtime.UncurriedApp2(go__go_3_5_97, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_4.UnsafePtr).V4)}, gopurs_runtime.Apply2(f_1, (*Constructor_Data_Map_Internal_Node)(__local_var_4.UnsafePtr).V3, gopurs_runtime.UncurriedApp2(go__go_3_5_97, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_4.UnsafePtr).V5)}, __local_var_5)))
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
return gopurs_runtime.Func(func(m_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.UncurriedApp2(go__go_3_5_97, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](m_4))}, z_2)
})
})
}))
}), gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Semigroup0_1_7 -> *Constructor_Data_Semigroup_Semigroup
Semigroup0_1_7 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_1_7
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_3_8_98 gopurs_runtime.Value
_ = go__go_3_8_98
go__go_3_8_98 = gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t9 gopurs_runtime.Value
{
if (v_4.Type == 9 && v_4.IntVal == 324739070 && v_4.UnsafePtr == nil) {
__t9 = gopurs_runtime.RecordGet(dictMonoid_0, "mempty")
goto end_branch_9
} else {

}
}
{
if (v_4.Type == 9 && v_4.IntVal == 324739070 && v_4.UnsafePtr != nil) {
__t9 = gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_1_7.V0), gopurs_runtime.Apply(go__go_3_8_98, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v_4.UnsafePtr).V4)}), gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_1_7.V0), gopurs_runtime.Apply2(f_2, (*Constructor_Data_Map_Internal_Node)(v_4.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node)(v_4.UnsafePtr).V3), gopurs_runtime.Apply(go__go_3_8_98, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v_4.UnsafePtr).V5)})))
goto end_branch_9
} else {

}
}
{
__t9 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_9:
return __t9
})
return go__go_3_8_98
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(z_1 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_2_10_99 gopurs_runtime.Value
go__go_2_10_99 = gopurs_runtime.Func(func(__local_var_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(__local_var_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var __local_var_3_loop gopurs_runtime.Value = __local_var_3_loop_val
var __local_var_4_loop gopurs_runtime.Value = __local_var_4_loop_val
go__go_2_10_99:
for {
if false { continue go__go_2_10_99 }
var __local_var_3 gopurs_runtime.Value = __local_var_3_loop
_ = __local_var_3
var __local_var_4 gopurs_runtime.Value = __local_var_4_loop
_ = __local_var_4
var __t11 gopurs_runtime.Value
{
if (__local_var_4.Type == 9 && __local_var_4.IntVal == 324739070 && __local_var_4.UnsafePtr == nil) {
__t11 = __local_var_3
goto end_branch_11
} else {

}
}
{
if (__local_var_4.Type == 9 && __local_var_4.IntVal == 324739070 && __local_var_4.UnsafePtr != nil) {
__t11 = gopurs_runtime.UncurriedApp2(go__go_2_10_99, gopurs_runtime.Apply3(f_0, (*Constructor_Data_Map_Internal_Node)(__local_var_4.UnsafePtr).V2, gopurs_runtime.UncurriedApp2(go__go_2_10_99, __local_var_3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_4.UnsafePtr).V4)}), (*Constructor_Data_Map_Internal_Node)(__local_var_4.UnsafePtr).V3), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_4.UnsafePtr).V5)})
goto end_branch_11
} else {

}
}
{
__t11 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_11:
return __t11
}
}()
})
})
return gopurs_runtime.Func(func(m_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.UncurriedApp2(go__go_2_10_99, z_1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](m_3))})
})
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(z_1 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_2_12_100 gopurs_runtime.Value
_ = go__go_2_12_100
go__go_2_12_100 = gopurs_runtime.Func2(func(__local_var_3 gopurs_runtime.Value, __local_var_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t13 gopurs_runtime.Value
{
if (__local_var_3.Type == 9 && __local_var_3.IntVal == 324739070 && __local_var_3.UnsafePtr == nil) {
__t13 = __local_var_4
goto end_branch_13
} else {

}
}
{
if (__local_var_3.Type == 9 && __local_var_3.IntVal == 324739070 && __local_var_3.UnsafePtr != nil) {
__t13 = gopurs_runtime.UncurriedApp2(go__go_2_12_100, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_3.UnsafePtr).V4)}, gopurs_runtime.Apply3(f_0, (*Constructor_Data_Map_Internal_Node)(__local_var_3.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node)(__local_var_3.UnsafePtr).V3, gopurs_runtime.UncurriedApp2(go__go_2_12_100, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_3.UnsafePtr).V5)}, __local_var_4)))
goto end_branch_13
} else {

}
}
{
__t13 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_13:
return __t13
})
return gopurs_runtime.Func(func(m_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.UncurriedApp2(go__go_2_12_100, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](m_3))}, z_1)
})
})
}))
	})
	return cache_Data_Map_Internal_foldableWithIndexMap__1035756962
}

var cache_Data_Map_Internal_foldableWithIndexMap__1966365627 gopurs_runtime.Value
var once_Data_Map_Internal_foldableWithIndexMap__1966365627 sync.Once
func Get_Data_Map_Internal_foldableWithIndexMap__1966365627() gopurs_runtime.Value {
	once_Data_Map_Internal_foldableWithIndexMap__1966365627.Do(func() {
		cache_Data_Map_Internal_foldableWithIndexMap__1966365627 = gopurs_runtime.RecordDict4("Foldable0", "foldMapWithIndex", "foldlWithIndex", "foldrWithIndex", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict3("foldMap", "foldl", "foldr", gopurs_runtime.Func(func(dictMonoid_1 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Semigroup0_2_0 -> *Constructor_Data_Semigroup_Semigroup
Semigroup0_2_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_1, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_2_0
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_4_1_101 gopurs_runtime.Value
_ = go__go_4_1_101
go__go_4_1_101 = gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 gopurs_runtime.Value
{
if (v_5.Type == 9 && v_5.IntVal == 324739070 && v_5.UnsafePtr == nil) {
__t2 = gopurs_runtime.RecordGet(dictMonoid_1, "mempty")
goto end_branch_2
} else {

}
}
{
if (v_5.Type == 9 && v_5.IntVal == 324739070 && v_5.UnsafePtr != nil) {
__t2 = gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_2_0.V0), gopurs_runtime.Apply(go__go_4_1_101, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v_5.UnsafePtr).V4)}), gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_2_0.V0), gopurs_runtime.Apply(f_3, (*Constructor_Data_Map_Internal_Node)(v_5.UnsafePtr).V3), gopurs_runtime.Apply(go__go_4_1_101, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v_5.UnsafePtr).V5)})))
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
return go__go_4_1_101
})
}), gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(z_2 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_3_3_102 gopurs_runtime.Value
go__go_3_3_102 = gopurs_runtime.Func(func(__local_var_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(__local_var_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var __local_var_4_loop gopurs_runtime.Value = __local_var_4_loop_val
var __local_var_5_loop gopurs_runtime.Value = __local_var_5_loop_val
go__go_3_3_102:
for {
if false { continue go__go_3_3_102 }
var __local_var_4 gopurs_runtime.Value = __local_var_4_loop
_ = __local_var_4
var __local_var_5 gopurs_runtime.Value = __local_var_5_loop
_ = __local_var_5
var __t4 gopurs_runtime.Value
{
if (__local_var_5.Type == 9 && __local_var_5.IntVal == 324739070 && __local_var_5.UnsafePtr == nil) {
__t4 = __local_var_4
goto end_branch_4
} else {

}
}
{
if (__local_var_5.Type == 9 && __local_var_5.IntVal == 324739070 && __local_var_5.UnsafePtr != nil) {
__t4 = gopurs_runtime.UncurriedApp2(go__go_3_3_102, gopurs_runtime.Apply2(f_1, gopurs_runtime.UncurriedApp2(go__go_3_3_102, __local_var_4, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_5.UnsafePtr).V4)}), (*Constructor_Data_Map_Internal_Node)(__local_var_5.UnsafePtr).V3), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_5.UnsafePtr).V5)})
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
return gopurs_runtime.Func(func(m_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.UncurriedApp2(go__go_3_3_102, z_2, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](m_4))})
})
})
}), gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(z_2 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_3_5_103 gopurs_runtime.Value
_ = go__go_3_5_103
go__go_3_5_103 = gopurs_runtime.Func2(func(__local_var_4 gopurs_runtime.Value, __local_var_5 gopurs_runtime.Value) gopurs_runtime.Value {
var __t6 gopurs_runtime.Value
{
if (__local_var_4.Type == 9 && __local_var_4.IntVal == 324739070 && __local_var_4.UnsafePtr == nil) {
__t6 = __local_var_5
goto end_branch_6
} else {

}
}
{
if (__local_var_4.Type == 9 && __local_var_4.IntVal == 324739070 && __local_var_4.UnsafePtr != nil) {
__t6 = gopurs_runtime.UncurriedApp2(go__go_3_5_103, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_4.UnsafePtr).V4)}, gopurs_runtime.Apply2(f_1, (*Constructor_Data_Map_Internal_Node)(__local_var_4.UnsafePtr).V3, gopurs_runtime.UncurriedApp2(go__go_3_5_103, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_4.UnsafePtr).V5)}, __local_var_5)))
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
return gopurs_runtime.Func(func(m_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.UncurriedApp2(go__go_3_5_103, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](m_4))}, z_2)
})
})
}))
}), gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Semigroup0_1_7 -> *Constructor_Data_Semigroup_Semigroup
Semigroup0_1_7 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_1_7
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_3_8_104 gopurs_runtime.Value
_ = go__go_3_8_104
go__go_3_8_104 = gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t9 gopurs_runtime.Value
{
if (v_4.Type == 9 && v_4.IntVal == 324739070 && v_4.UnsafePtr == nil) {
__t9 = gopurs_runtime.RecordGet(dictMonoid_0, "mempty")
goto end_branch_9
} else {

}
}
{
if (v_4.Type == 9 && v_4.IntVal == 324739070 && v_4.UnsafePtr != nil) {
__t9 = gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_1_7.V0), gopurs_runtime.Apply(go__go_3_8_104, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v_4.UnsafePtr).V4)}), gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_1_7.V0), gopurs_runtime.Apply2(f_2, (*Constructor_Data_Map_Internal_Node)(v_4.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node)(v_4.UnsafePtr).V3), gopurs_runtime.Apply(go__go_3_8_104, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v_4.UnsafePtr).V5)})))
goto end_branch_9
} else {

}
}
{
__t9 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_9:
return __t9
})
return go__go_3_8_104
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(z_1 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_2_10_105 gopurs_runtime.Value
go__go_2_10_105 = gopurs_runtime.Func(func(__local_var_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(__local_var_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var __local_var_3_loop gopurs_runtime.Value = __local_var_3_loop_val
var __local_var_4_loop gopurs_runtime.Value = __local_var_4_loop_val
go__go_2_10_105:
for {
if false { continue go__go_2_10_105 }
var __local_var_3 gopurs_runtime.Value = __local_var_3_loop
_ = __local_var_3
var __local_var_4 gopurs_runtime.Value = __local_var_4_loop
_ = __local_var_4
var __t11 gopurs_runtime.Value
{
if (__local_var_4.Type == 9 && __local_var_4.IntVal == 324739070 && __local_var_4.UnsafePtr == nil) {
__t11 = __local_var_3
goto end_branch_11
} else {

}
}
{
if (__local_var_4.Type == 9 && __local_var_4.IntVal == 324739070 && __local_var_4.UnsafePtr != nil) {
__t11 = gopurs_runtime.UncurriedApp2(go__go_2_10_105, gopurs_runtime.Apply3(f_0, (*Constructor_Data_Map_Internal_Node)(__local_var_4.UnsafePtr).V2, gopurs_runtime.UncurriedApp2(go__go_2_10_105, __local_var_3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_4.UnsafePtr).V4)}), (*Constructor_Data_Map_Internal_Node)(__local_var_4.UnsafePtr).V3), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_4.UnsafePtr).V5)})
goto end_branch_11
} else {

}
}
{
__t11 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_11:
return __t11
}
}()
})
})
return gopurs_runtime.Func(func(m_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.UncurriedApp2(go__go_2_10_105, z_1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](m_3))})
})
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(z_1 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_2_12_106 gopurs_runtime.Value
_ = go__go_2_12_106
go__go_2_12_106 = gopurs_runtime.Func2(func(__local_var_3 gopurs_runtime.Value, __local_var_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t13 gopurs_runtime.Value
{
if (__local_var_3.Type == 9 && __local_var_3.IntVal == 324739070 && __local_var_3.UnsafePtr == nil) {
__t13 = __local_var_4
goto end_branch_13
} else {

}
}
{
if (__local_var_3.Type == 9 && __local_var_3.IntVal == 324739070 && __local_var_3.UnsafePtr != nil) {
__t13 = gopurs_runtime.UncurriedApp2(go__go_2_12_106, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_3.UnsafePtr).V4)}, gopurs_runtime.Apply3(f_0, (*Constructor_Data_Map_Internal_Node)(__local_var_3.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node)(__local_var_3.UnsafePtr).V3, gopurs_runtime.UncurriedApp2(go__go_2_12_106, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_3.UnsafePtr).V5)}, __local_var_4)))
goto end_branch_13
} else {

}
}
{
__t13 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_13:
return __t13
})
return gopurs_runtime.Func(func(m_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.UncurriedApp2(go__go_2_12_106, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](m_3))}, z_1)
})
})
}))
	})
	return cache_Data_Map_Internal_foldableWithIndexMap__1966365627
}

var cache_Data_Map_Internal_functorMap__2501170929 gopurs_runtime.Value
var once_Data_Map_Internal_functorMap__2501170929 sync.Once
func Get_Data_Map_Internal_functorMap__2501170929() gopurs_runtime.Value {
	once_Data_Map_Internal_functorMap__2501170929.Do(func() {
		cache_Data_Map_Internal_functorMap__2501170929 = gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_1_0_107 gopurs_runtime.Value
_ = go__go_1_0_107
go__go_1_0_107 = gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 *Constructor_Data_Map_Internal_Node
{
if (v_2.Type == 9 && v_2.IntVal == 324739070 && v_2.UnsafePtr == nil) {
__t1 = (*Constructor_Data_Map_Internal_Node)(nil)
goto end_branch_1
} else {

}
}
{
if (v_2.Type == 9 && v_2.IntVal == 324739070 && v_2.UnsafePtr != nil) {
__t1 = &Constructor_Data_Map_Internal_Node{1, (*Constructor_Data_Map_Internal_Node)(v_2.UnsafePtr).V0, (*Constructor_Data_Map_Internal_Node)(v_2.UnsafePtr).V1, (*Constructor_Data_Map_Internal_Node)(v_2.UnsafePtr).V2, gopurs_runtime.Apply(f_0, (*Constructor_Data_Map_Internal_Node)(v_2.UnsafePtr).V3), gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.Apply(go__go_1_0_107, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v_2.UnsafePtr).V4)})), gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.Apply(go__go_1_0_107, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v_2.UnsafePtr).V5)}))}
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_1:
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(__t1)}
})
return go__go_1_0_107
}))
	})
	return cache_Data_Map_Internal_functorMap__2501170929
}

var cache_Data_Map_Internal_functorWithIndexMap__3138419015 gopurs_runtime.Value
var once_Data_Map_Internal_functorWithIndexMap__3138419015 sync.Once
func Get_Data_Map_Internal_functorWithIndexMap__3138419015() gopurs_runtime.Value {
	once_Data_Map_Internal_functorWithIndexMap__3138419015.Do(func() {
		cache_Data_Map_Internal_functorWithIndexMap__3138419015 = gopurs_runtime.RecordDict2("Functor0", "mapWithIndex", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_2_0_108 gopurs_runtime.Value
_ = go__go_2_0_108
go__go_2_0_108 = gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 *Constructor_Data_Map_Internal_Node
{
if (v_3.Type == 9 && v_3.IntVal == 324739070 && v_3.UnsafePtr == nil) {
__t1 = (*Constructor_Data_Map_Internal_Node)(nil)
goto end_branch_1
} else {

}
}
{
if (v_3.Type == 9 && v_3.IntVal == 324739070 && v_3.UnsafePtr != nil) {
__t1 = &Constructor_Data_Map_Internal_Node{1, (*Constructor_Data_Map_Internal_Node)(v_3.UnsafePtr).V0, (*Constructor_Data_Map_Internal_Node)(v_3.UnsafePtr).V1, (*Constructor_Data_Map_Internal_Node)(v_3.UnsafePtr).V2, gopurs_runtime.Apply(f_1, (*Constructor_Data_Map_Internal_Node)(v_3.UnsafePtr).V3), gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.Apply(go__go_2_0_108, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v_3.UnsafePtr).V4)})), gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.Apply(go__go_2_0_108, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v_3.UnsafePtr).V5)}))}
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_1:
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(__t1)}
})
return go__go_2_0_108
}))
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_1_2_109 gopurs_runtime.Value
_ = go__go_1_2_109
go__go_1_2_109 = gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t3 *Constructor_Data_Map_Internal_Node
{
if (v_2.Type == 9 && v_2.IntVal == 324739070 && v_2.UnsafePtr == nil) {
__t3 = (*Constructor_Data_Map_Internal_Node)(nil)
goto end_branch_3
} else {

}
}
{
if (v_2.Type == 9 && v_2.IntVal == 324739070 && v_2.UnsafePtr != nil) {
__t3 = &Constructor_Data_Map_Internal_Node{1, (*Constructor_Data_Map_Internal_Node)(v_2.UnsafePtr).V0, (*Constructor_Data_Map_Internal_Node)(v_2.UnsafePtr).V1, (*Constructor_Data_Map_Internal_Node)(v_2.UnsafePtr).V2, gopurs_runtime.Apply2(f_0, (*Constructor_Data_Map_Internal_Node)(v_2.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node)(v_2.UnsafePtr).V3), gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.Apply(go__go_1_2_109, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v_2.UnsafePtr).V4)})), gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.Apply(go__go_1_2_109, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v_2.UnsafePtr).V5)}))}
goto end_branch_3
} else {

}
}
{
__t3 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_3:
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(__t3)}
})
return go__go_1_2_109
}))
	})
	return cache_Data_Map_Internal_functorWithIndexMap__3138419015
}

var cache_Data_Map_Internal_insert__3204212386 gopurs_runtime.Value
var once_Data_Map_Internal_insert__3204212386 sync.Once
func Get_Data_Map_Internal_insert__3204212386() gopurs_runtime.Value {
	once_Data_Map_Internal_insert__3204212386.Do(func() {
		cache_Data_Map_Internal_insert__3204212386 = gopurs_runtime.Func3(func(dictOrd_0_box gopurs_runtime.Value, k_1_box gopurs_runtime.Value, v_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Map_Internal_insert__3204212386(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord](dictOrd_0_box), k_1_box, v_2_box)
})
	})
	return cache_Data_Map_Internal_insert__3204212386
}

var cache_Data_Map_Internal_insert__4289641298 gopurs_runtime.Value
var once_Data_Map_Internal_insert__4289641298 sync.Once
func Get_Data_Map_Internal_insert__4289641298() gopurs_runtime.Value {
	once_Data_Map_Internal_insert__4289641298.Do(func() {
		cache_Data_Map_Internal_insert__4289641298 = gopurs_runtime.Func3(func(dictOrd_0_box gopurs_runtime.Value, k_1_box gopurs_runtime.Value, v_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Map_Internal_insert__4289641298(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord](dictOrd_0_box), k_1_box, v_2_box)
})
	})
	return cache_Data_Map_Internal_insert__4289641298
}

var cache_Data_Map_Internal_insert__2073142786 gopurs_runtime.Value
var once_Data_Map_Internal_insert__2073142786 sync.Once
func Get_Data_Map_Internal_insert__2073142786() gopurs_runtime.Value {
	once_Data_Map_Internal_insert__2073142786.Do(func() {
		cache_Data_Map_Internal_insert__2073142786 = gopurs_runtime.Func3(func(dictOrd_0_box gopurs_runtime.Value, k_1_box gopurs_runtime.Value, v_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Map_Internal_insert__2073142786(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord](dictOrd_0_box), gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](k_1_box), v_2_box)
})
	})
	return cache_Data_Map_Internal_insert__2073142786
}

var cache_Data_Map_Internal_insertWith__118979962 gopurs_runtime.Value
var once_Data_Map_Internal_insertWith__118979962 sync.Once
func Get_Data_Map_Internal_insertWith__118979962() gopurs_runtime.Value {
	once_Data_Map_Internal_insertWith__118979962.Do(func() {
		cache_Data_Map_Internal_insertWith__118979962 = gopurs_runtime.Func4(func(dictOrd_0_box gopurs_runtime.Value, app_1_box gopurs_runtime.Value, k_2_box gopurs_runtime.Value, v_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Map_Internal_insertWith__118979962(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord](dictOrd_0_box), app_1_box, k_2_box, v_3_box)
})
	})
	return cache_Data_Map_Internal_insertWith__118979962
}

var cache_Data_Map_Internal_intersectionWith__3717755541 gopurs_runtime.Value
var once_Data_Map_Internal_intersectionWith__3717755541 sync.Once
func Get_Data_Map_Internal_intersectionWith__3717755541() gopurs_runtime.Value {
	once_Data_Map_Internal_intersectionWith__3717755541.Do(func() {
		cache_Data_Map_Internal_intersectionWith__3717755541 = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Map_Internal_intersectionWith__3717755541(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord](dictOrd_0_box))
})
	})
	return cache_Data_Map_Internal_intersectionWith__3717755541
}

var cache_Data_Map_Internal_intersectionWith__4144106805 gopurs_runtime.Value
var once_Data_Map_Internal_intersectionWith__4144106805 sync.Once
func Get_Data_Map_Internal_intersectionWith__4144106805() gopurs_runtime.Value {
	once_Data_Map_Internal_intersectionWith__4144106805.Do(func() {
		cache_Data_Map_Internal_intersectionWith__4144106805 = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Map_Internal_intersectionWith__4144106805(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord](dictOrd_0_box))
})
	})
	return cache_Data_Map_Internal_intersectionWith__4144106805
}

var cache_Data_Map_Internal_isEmpty__1620059593 gopurs_runtime.Value
var once_Data_Map_Internal_isEmpty__1620059593 sync.Once
func Get_Data_Map_Internal_isEmpty__1620059593() gopurs_runtime.Value {
	once_Data_Map_Internal_isEmpty__1620059593.Do(func() {
		cache_Data_Map_Internal_isEmpty__1620059593 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_Data_Map_Internal_isEmpty__1620059593(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](v_0_box)))
})
	})
	return cache_Data_Map_Internal_isEmpty__1620059593
}

var cache_Data_Map_Internal_iterMapL__3394814354 gopurs_runtime.Value
var once_Data_Map_Internal_iterMapL__3394814354 sync.Once
func Get_Data_Map_Internal_iterMapL__3394814354() gopurs_runtime.Value {
	once_Data_Map_Internal_iterMapL__3394814354.Do(func() {
		cache_Data_Map_Internal_iterMapL__3394814354 = func() gopurs_runtime.Value {
var go__go_0_0_114 gopurs_runtime.Value
go__go_0_0_114 = gopurs_runtime.Func(func(iter_1_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var iter_1_loop gopurs_runtime.Value = iter_1_loop_val
var v_2_loop gopurs_runtime.Value = v_2_loop_val
go__go_0_0_114:
for {
if false { continue go__go_0_0_114 }
var iter_1 gopurs_runtime.Value = iter_1_loop
_ = iter_1
var v_2 gopurs_runtime.Value = v_2_loop
_ = v_2
var __t3 gopurs_runtime.Value
{
if (v_2.Type == 9 && v_2.IntVal == 324739070 && v_2.UnsafePtr == nil) {
__t3 = iter_1
goto end_branch_3
} else {

}
}
{
if (v_2.Type == 9 && v_2.IntVal == 324739070 && v_2.UnsafePtr != nil) {
var __t2 gopurs_runtime.Value
{
var __t_tag_1 *Constructor_Data_Map_Internal_Node = (*Constructor_Data_Map_Internal_Node)(v_2.UnsafePtr).V5
if (__t_tag_1 == nil) {
iter_1_loop = gopurs_runtime.Value{Type: 9, IntVal: 1343415489, UnsafePtr: unsafe.Pointer(&Constructor_Data_Map_Internal_IterEmit{1, (*Constructor_Data_Map_Internal_Node)(v_2.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node)(v_2.UnsafePtr).V3, iter_1})}
v_2_loop = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v_2.UnsafePtr).V4)}
continue go__go_0_0_114
__t2 = gopurs_runtime.Value{}
goto end_branch_2
} else {

}
}
{
iter_1_loop = gopurs_runtime.Value{Type: 9, IntVal: 1343415489, UnsafePtr: unsafe.Pointer(&Constructor_Data_Map_Internal_IterEmit{1, (*Constructor_Data_Map_Internal_Node)(v_2.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node)(v_2.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 2861335956, UnsafePtr: unsafe.Pointer(&Constructor_Data_Map_Internal_IterNode{1, (*Constructor_Data_Map_Internal_Node)(v_2.UnsafePtr).V5, iter_1})}})}
v_2_loop = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v_2.UnsafePtr).V4)}
continue go__go_0_0_114
__t2 = gopurs_runtime.Value{}
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
}()
})
})
return go__go_0_0_114
}()
	})
	return cache_Data_Map_Internal_iterMapL__3394814354
}

var cache_Data_Map_Internal_iterMapL__878452066 gopurs_runtime.Value
var once_Data_Map_Internal_iterMapL__878452066 sync.Once
func Get_Data_Map_Internal_iterMapL__878452066() gopurs_runtime.Value {
	once_Data_Map_Internal_iterMapL__878452066.Do(func() {
		cache_Data_Map_Internal_iterMapL__878452066 = func() gopurs_runtime.Value {
var go__go_0_0_115 gopurs_runtime.Value
go__go_0_0_115 = gopurs_runtime.Func(func(iter_1_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var iter_1_loop gopurs_runtime.Value = iter_1_loop_val
var v_2_loop gopurs_runtime.Value = v_2_loop_val
go__go_0_0_115:
for {
if false { continue go__go_0_0_115 }
var iter_1 gopurs_runtime.Value = iter_1_loop
_ = iter_1
var v_2 gopurs_runtime.Value = v_2_loop
_ = v_2
var __t3 gopurs_runtime.Value
{
if (v_2.Type == 9 && v_2.IntVal == 324739070 && v_2.UnsafePtr == nil) {
__t3 = iter_1
goto end_branch_3
} else {

}
}
{
if (v_2.Type == 9 && v_2.IntVal == 324739070 && v_2.UnsafePtr != nil) {
var __t2 gopurs_runtime.Value
{
var __t_tag_1 *Constructor_Data_Map_Internal_Node = (*Constructor_Data_Map_Internal_Node)(v_2.UnsafePtr).V5
if (__t_tag_1 == nil) {
iter_1_loop = gopurs_runtime.Value{Type: 9, IntVal: 1343415489, UnsafePtr: unsafe.Pointer(&Constructor_Data_Map_Internal_IterEmit{1, (*Constructor_Data_Map_Internal_Node)(v_2.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node)(v_2.UnsafePtr).V3, iter_1})}
v_2_loop = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v_2.UnsafePtr).V4)}
continue go__go_0_0_115
__t2 = gopurs_runtime.Value{}
goto end_branch_2
} else {

}
}
{
iter_1_loop = gopurs_runtime.Value{Type: 9, IntVal: 1343415489, UnsafePtr: unsafe.Pointer(&Constructor_Data_Map_Internal_IterEmit{1, (*Constructor_Data_Map_Internal_Node)(v_2.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node)(v_2.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 2861335956, UnsafePtr: unsafe.Pointer(&Constructor_Data_Map_Internal_IterNode{1, (*Constructor_Data_Map_Internal_Node)(v_2.UnsafePtr).V5, iter_1})}})}
v_2_loop = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v_2.UnsafePtr).V4)}
continue go__go_0_0_115
__t2 = gopurs_runtime.Value{}
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
}()
})
})
return go__go_0_0_115
}()
	})
	return cache_Data_Map_Internal_iterMapL__878452066
}

var cache_Data_Map_Internal_iterMapL__1101342704 gopurs_runtime.Value
var once_Data_Map_Internal_iterMapL__1101342704 sync.Once
func Get_Data_Map_Internal_iterMapL__1101342704() gopurs_runtime.Value {
	once_Data_Map_Internal_iterMapL__1101342704.Do(func() {
		cache_Data_Map_Internal_iterMapL__1101342704 = func() gopurs_runtime.Value {
var go__go_0_0_116 gopurs_runtime.Value
go__go_0_0_116 = gopurs_runtime.Func(func(iter_1_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var iter_1_loop gopurs_runtime.Value = iter_1_loop_val
var v_2_loop gopurs_runtime.Value = v_2_loop_val
go__go_0_0_116:
for {
if false { continue go__go_0_0_116 }
var iter_1 gopurs_runtime.Value = iter_1_loop
_ = iter_1
var v_2 gopurs_runtime.Value = v_2_loop
_ = v_2
var __t3 gopurs_runtime.Value
{
if (v_2.Type == 9 && v_2.IntVal == 324739070 && v_2.UnsafePtr == nil) {
__t3 = iter_1
goto end_branch_3
} else {

}
}
{
if (v_2.Type == 9 && v_2.IntVal == 324739070 && v_2.UnsafePtr != nil) {
var __t2 gopurs_runtime.Value
{
var __t_tag_1 *Constructor_Data_Map_Internal_Node = (*Constructor_Data_Map_Internal_Node)(v_2.UnsafePtr).V5
if (__t_tag_1 == nil) {
iter_1_loop = gopurs_runtime.Value{Type: 9, IntVal: 1343415489, UnsafePtr: unsafe.Pointer(&Constructor_Data_Map_Internal_IterEmit{1, gopurs_runtime.Value{Type: 9, IntVal: int64(uint32((*Constructor_Data_Map_Internal_Node)(v_2.UnsafePtr).V2.IntVal)), UnsafePtr: nil}, gopurs_runtime.Float((*Constructor_Data_Map_Internal_Node)(v_2.UnsafePtr).V3.FloatVal()), iter_1})}
v_2_loop = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v_2.UnsafePtr).V4)}
continue go__go_0_0_116
__t2 = gopurs_runtime.Value{}
goto end_branch_2
} else {

}
}
{
iter_1_loop = gopurs_runtime.Value{Type: 9, IntVal: 1343415489, UnsafePtr: unsafe.Pointer(&Constructor_Data_Map_Internal_IterEmit{1, gopurs_runtime.Value{Type: 9, IntVal: int64(uint32((*Constructor_Data_Map_Internal_Node)(v_2.UnsafePtr).V2.IntVal)), UnsafePtr: nil}, gopurs_runtime.Float((*Constructor_Data_Map_Internal_Node)(v_2.UnsafePtr).V3.FloatVal()), gopurs_runtime.Value{Type: 9, IntVal: 2861335956, UnsafePtr: unsafe.Pointer(&Constructor_Data_Map_Internal_IterNode{1, (*Constructor_Data_Map_Internal_Node)(v_2.UnsafePtr).V5, iter_1})}})}
v_2_loop = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v_2.UnsafePtr).V4)}
continue go__go_0_0_116
__t2 = gopurs_runtime.Value{}
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
}()
})
})
return go__go_0_0_116
}()
	})
	return cache_Data_Map_Internal_iterMapL__1101342704
}

var cache_Data_Map_Internal_iterMapR__878452066 gopurs_runtime.Value
var once_Data_Map_Internal_iterMapR__878452066 sync.Once
func Get_Data_Map_Internal_iterMapR__878452066() gopurs_runtime.Value {
	once_Data_Map_Internal_iterMapR__878452066.Do(func() {
		cache_Data_Map_Internal_iterMapR__878452066 = func() gopurs_runtime.Value {
var go__go_0_0_117 gopurs_runtime.Value
go__go_0_0_117 = gopurs_runtime.Func(func(iter_1_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var iter_1_loop gopurs_runtime.Value = iter_1_loop_val
var v_2_loop gopurs_runtime.Value = v_2_loop_val
go__go_0_0_117:
for {
if false { continue go__go_0_0_117 }
var iter_1 gopurs_runtime.Value = iter_1_loop
_ = iter_1
var v_2 gopurs_runtime.Value = v_2_loop
_ = v_2
var __t3 gopurs_runtime.Value
{
if (v_2.Type == 9 && v_2.IntVal == 324739070 && v_2.UnsafePtr == nil) {
__t3 = iter_1
goto end_branch_3
} else {

}
}
{
if (v_2.Type == 9 && v_2.IntVal == 324739070 && v_2.UnsafePtr != nil) {
var __t2 gopurs_runtime.Value
{
var __t_tag_1 *Constructor_Data_Map_Internal_Node = (*Constructor_Data_Map_Internal_Node)(v_2.UnsafePtr).V5
if (__t_tag_1 == nil) {
iter_1_loop = gopurs_runtime.Value{Type: 9, IntVal: 1343415489, UnsafePtr: unsafe.Pointer(&Constructor_Data_Map_Internal_IterEmit{1, (*Constructor_Data_Map_Internal_Node)(v_2.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node)(v_2.UnsafePtr).V3, iter_1})}
v_2_loop = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v_2.UnsafePtr).V4)}
continue go__go_0_0_117
__t2 = gopurs_runtime.Value{}
goto end_branch_2
} else {

}
}
{
iter_1_loop = gopurs_runtime.Value{Type: 9, IntVal: 1343415489, UnsafePtr: unsafe.Pointer(&Constructor_Data_Map_Internal_IterEmit{1, (*Constructor_Data_Map_Internal_Node)(v_2.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node)(v_2.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 2861335956, UnsafePtr: unsafe.Pointer(&Constructor_Data_Map_Internal_IterNode{1, (*Constructor_Data_Map_Internal_Node)(v_2.UnsafePtr).V4, iter_1})}})}
v_2_loop = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v_2.UnsafePtr).V5)}
continue go__go_0_0_117
__t2 = gopurs_runtime.Value{}
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
}()
})
})
return go__go_0_0_117
}()
	})
	return cache_Data_Map_Internal_iterMapR__878452066
}

var cache_Data_Map_Internal_iterMapU__878452066 gopurs_runtime.Value
var once_Data_Map_Internal_iterMapU__878452066 sync.Once
func Get_Data_Map_Internal_iterMapU__878452066() gopurs_runtime.Value {
	once_Data_Map_Internal_iterMapU__878452066.Do(func() {
		cache_Data_Map_Internal_iterMapU__878452066 = gopurs_runtime.Func2(func(iter_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Map_Internal_iterMapU__878452066(iter_0_box, gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](v_1_box))
})
	})
	return cache_Data_Map_Internal_iterMapU__878452066
}

var cache_Data_Map_Internal_keys__3504999702 gopurs_runtime.Value
var once_Data_Map_Internal_keys__3504999702 sync.Once
func Get_Data_Map_Internal_keys__3504999702() gopurs_runtime.Value {
	once_Data_Map_Internal_keys__3504999702.Do(func() {
		cache_Data_Map_Internal_keys__3504999702 = func() gopurs_runtime.Value {
var go__go_0_0_118 gopurs_runtime.Value
_ = go__go_0_0_118
go__go_0_0_118 = gopurs_runtime.Func2(func(__local_var_1 gopurs_runtime.Value, __local_var_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if (__local_var_1.Type == 9 && __local_var_1.IntVal == 324739070 && __local_var_1.UnsafePtr == nil) {
__t1 = __local_var_2
goto end_branch_1
} else {

}
}
{
if (__local_var_1.Type == 9 && __local_var_1.IntVal == 324739070 && __local_var_1.UnsafePtr != nil) {
__t1 = gopurs_runtime.UncurriedApp2(go__go_0_0_118, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_1.UnsafePtr).V4)}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, (*Constructor_Data_Map_Internal_Node)(__local_var_1.UnsafePtr).V2, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.UncurriedApp2(go__go_0_0_118, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_1.UnsafePtr).V5)}, __local_var_2))})})
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
return gopurs_runtime.Func(func(m_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.UncurriedApp2(go__go_0_0_118, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](m_1))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(nil))})
})
}()
	})
	return cache_Data_Map_Internal_keys__3504999702
}

var cache_Data_Map_Internal_keys__2406038214 gopurs_runtime.Value
var once_Data_Map_Internal_keys__2406038214 sync.Once
func Get_Data_Map_Internal_keys__2406038214() gopurs_runtime.Value {
	once_Data_Map_Internal_keys__2406038214.Do(func() {
		cache_Data_Map_Internal_keys__2406038214 = func() gopurs_runtime.Value {
var go__go_0_0_119 gopurs_runtime.Value
_ = go__go_0_0_119
go__go_0_0_119 = gopurs_runtime.Func2(func(__local_var_1 gopurs_runtime.Value, __local_var_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if (__local_var_1.Type == 9 && __local_var_1.IntVal == 324739070 && __local_var_1.UnsafePtr == nil) {
__t1 = __local_var_2
goto end_branch_1
} else {

}
}
{
if (__local_var_1.Type == 9 && __local_var_1.IntVal == 324739070 && __local_var_1.UnsafePtr != nil) {
__t1 = gopurs_runtime.UncurriedApp2(go__go_0_0_119, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_1.UnsafePtr).V4)}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, (*Constructor_Data_Map_Internal_Node)(__local_var_1.UnsafePtr).V2, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.UncurriedApp2(go__go_0_0_119, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_1.UnsafePtr).V5)}, __local_var_2))})})
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
return gopurs_runtime.Func(func(m_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.UncurriedApp2(go__go_0_0_119, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](m_1))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(nil))})
})
}()
	})
	return cache_Data_Map_Internal_keys__2406038214
}

var cache_Data_Map_Internal_keys__2813649686 gopurs_runtime.Value
var once_Data_Map_Internal_keys__2813649686 sync.Once
func Get_Data_Map_Internal_keys__2813649686() gopurs_runtime.Value {
	once_Data_Map_Internal_keys__2813649686.Do(func() {
		cache_Data_Map_Internal_keys__2813649686 = func() gopurs_runtime.Value {
var go__go_0_0_120 gopurs_runtime.Value
_ = go__go_0_0_120
go__go_0_0_120 = gopurs_runtime.Func2(func(__local_var_1 gopurs_runtime.Value, __local_var_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if (__local_var_1.Type == 9 && __local_var_1.IntVal == 324739070 && __local_var_1.UnsafePtr == nil) {
__t1 = __local_var_2
goto end_branch_1
} else {

}
}
{
if (__local_var_1.Type == 9 && __local_var_1.IntVal == 324739070 && __local_var_1.UnsafePtr != nil) {
__t1 = gopurs_runtime.UncurriedApp2(go__go_0_0_120, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_1.UnsafePtr).V4)}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just]((*Constructor_Data_Map_Internal_Node)(__local_var_1.UnsafePtr).V2))}, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.UncurriedApp2(go__go_0_0_120, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_1.UnsafePtr).V5)}, __local_var_2)))})})})
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
return gopurs_runtime.Func(func(m_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.UncurriedApp2(go__go_0_0_120, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](m_1))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(nil))})
})
}()
	})
	return cache_Data_Map_Internal_keys__2813649686
}

var cache_Data_Map_Internal_lookup__3378638282 gopurs_runtime.Value
var once_Data_Map_Internal_lookup__3378638282 sync.Once
func Get_Data_Map_Internal_lookup__3378638282() gopurs_runtime.Value {
	once_Data_Map_Internal_lookup__3378638282.Do(func() {
		cache_Data_Map_Internal_lookup__3378638282 = gopurs_runtime.Func2(func(dictOrd_0_box gopurs_runtime.Value, k_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Map_Internal_lookup__3378638282(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord](dictOrd_0_box), k_1_box)
})
	})
	return cache_Data_Map_Internal_lookup__3378638282
}

var cache_Data_Map_Internal_lookup__1040249709 gopurs_runtime.Value
var once_Data_Map_Internal_lookup__1040249709 sync.Once
func Get_Data_Map_Internal_lookup__1040249709() gopurs_runtime.Value {
	once_Data_Map_Internal_lookup__1040249709.Do(func() {
		cache_Data_Map_Internal_lookup__1040249709 = gopurs_runtime.Func(func(k_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Map_Internal_lookup__1040249709(uint32(k_0_box.IntVal))
})
	})
	return cache_Data_Map_Internal_lookup__1040249709
}

var cache_Data_Map_Internal_mapMaybe__3426301240 gopurs_runtime.Value
var once_Data_Map_Internal_mapMaybe__3426301240 sync.Once
func Get_Data_Map_Internal_mapMaybe__3426301240() gopurs_runtime.Value {
	once_Data_Map_Internal_mapMaybe__3426301240.Do(func() {
		cache_Data_Map_Internal_mapMaybe__3426301240 = gopurs_runtime.Func2(func(dictOrd_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Map_Internal_mapMaybe__3426301240(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord](dictOrd_0_box), x_1_box)
})
	})
	return cache_Data_Map_Internal_mapMaybe__3426301240
}

var cache_Data_Map_Internal_mapMaybe__1970555288 gopurs_runtime.Value
var once_Data_Map_Internal_mapMaybe__1970555288 sync.Once
func Get_Data_Map_Internal_mapMaybe__1970555288() gopurs_runtime.Value {
	once_Data_Map_Internal_mapMaybe__1970555288.Do(func() {
		cache_Data_Map_Internal_mapMaybe__1970555288 = gopurs_runtime.Func2(func(dictOrd_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Map_Internal_mapMaybe__1970555288(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord](dictOrd_0_box), x_1_box)
})
	})
	return cache_Data_Map_Internal_mapMaybe__1970555288
}

var cache_Data_Map_Internal_mapMaybeWithKey__817660689 gopurs_runtime.Value
var once_Data_Map_Internal_mapMaybeWithKey__817660689 sync.Once
func Get_Data_Map_Internal_mapMaybeWithKey__817660689() gopurs_runtime.Value {
	once_Data_Map_Internal_mapMaybeWithKey__817660689.Do(func() {
		cache_Data_Map_Internal_mapMaybeWithKey__817660689 = gopurs_runtime.Func2(func(dictOrd_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Map_Internal_mapMaybeWithKey__817660689(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord](dictOrd_0_box), f_1_box)
})
	})
	return cache_Data_Map_Internal_mapMaybeWithKey__817660689
}

var cache_Data_Map_Internal_singleton__3511563426 gopurs_runtime.Value
var once_Data_Map_Internal_singleton__3511563426 sync.Once
func Get_Data_Map_Internal_singleton__3511563426() gopurs_runtime.Value {
	once_Data_Map_Internal_singleton__3511563426.Do(func() {
		cache_Data_Map_Internal_singleton__3511563426 = gopurs_runtime.Func2(func(k_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(Call_Data_Map_Internal_singleton__3511563426(k_0_box, v_1_box))}
})
	})
	return cache_Data_Map_Internal_singleton__3511563426
}

var cache_Data_Map_Internal_singleton__943571066 gopurs_runtime.Value
var once_Data_Map_Internal_singleton__943571066 sync.Once
func Get_Data_Map_Internal_singleton__943571066() gopurs_runtime.Value {
	once_Data_Map_Internal_singleton__943571066.Do(func() {
		cache_Data_Map_Internal_singleton__943571066 = gopurs_runtime.Func2(func(k_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(Call_Data_Map_Internal_singleton__943571066(k_0_box, v_1_box))}
})
	})
	return cache_Data_Map_Internal_singleton__943571066
}

var cache_Data_Map_Internal_singleton__2450056090 gopurs_runtime.Value
var once_Data_Map_Internal_singleton__2450056090 sync.Once
func Get_Data_Map_Internal_singleton__2450056090() gopurs_runtime.Value {
	once_Data_Map_Internal_singleton__2450056090.Do(func() {
		cache_Data_Map_Internal_singleton__2450056090 = gopurs_runtime.Func2(func(k_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(Call_Data_Map_Internal_singleton__2450056090(k_0_box, v_1_box))}
})
	})
	return cache_Data_Map_Internal_singleton__2450056090
}

var cache_Data_Map_Internal_singleton__3707014010 gopurs_runtime.Value
var once_Data_Map_Internal_singleton__3707014010 sync.Once
func Get_Data_Map_Internal_singleton__3707014010() gopurs_runtime.Value {
	once_Data_Map_Internal_singleton__3707014010.Do(func() {
		cache_Data_Map_Internal_singleton__3707014010 = gopurs_runtime.Func2(func(k_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(Call_Data_Map_Internal_singleton__3707014010(k_0_box, v_1_box))}
})
	})
	return cache_Data_Map_Internal_singleton__3707014010
}

var cache_Data_Map_Internal_singleton__1518627866 gopurs_runtime.Value
var once_Data_Map_Internal_singleton__1518627866 sync.Once
func Get_Data_Map_Internal_singleton__1518627866() gopurs_runtime.Value {
	once_Data_Map_Internal_singleton__1518627866.Do(func() {
		cache_Data_Map_Internal_singleton__1518627866 = gopurs_runtime.Func2(func(k_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(Call_Data_Map_Internal_singleton__1518627866(uint32(k_0_box.IntVal), v_1_box.FloatVal()))}
})
	})
	return cache_Data_Map_Internal_singleton__1518627866
}

var cache_Data_Map_Internal_singleton__1300483034 gopurs_runtime.Value
var once_Data_Map_Internal_singleton__1300483034 sync.Once
func Get_Data_Map_Internal_singleton__1300483034() gopurs_runtime.Value {
	once_Data_Map_Internal_singleton__1300483034.Do(func() {
		cache_Data_Map_Internal_singleton__1300483034 = gopurs_runtime.Func2(func(k_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(Call_Data_Map_Internal_singleton__1300483034(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](k_0_box), v_1_box))}
})
	})
	return cache_Data_Map_Internal_singleton__1300483034
}

var cache_Data_Map_Internal_size__909390430 gopurs_runtime.Value
var once_Data_Map_Internal_size__909390430 sync.Once
func Get_Data_Map_Internal_size__909390430() gopurs_runtime.Value {
	once_Data_Map_Internal_size__909390430.Do(func() {
		cache_Data_Map_Internal_size__909390430 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int(Call_Data_Map_Internal_size__909390430(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](v_0_box)))
})
	})
	return cache_Data_Map_Internal_size__909390430
}

var cache_Data_Map_Internal_size__1374028086 gopurs_runtime.Value
var once_Data_Map_Internal_size__1374028086 sync.Once
func Get_Data_Map_Internal_size__1374028086() gopurs_runtime.Value {
	once_Data_Map_Internal_size__1374028086.Do(func() {
		cache_Data_Map_Internal_size__1374028086 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int(Call_Data_Map_Internal_size__1374028086(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](v_0_box)))
})
	})
	return cache_Data_Map_Internal_size__1374028086
}

var cache_Data_Map_Internal_size__2382154916 gopurs_runtime.Value
var once_Data_Map_Internal_size__2382154916 sync.Once
func Get_Data_Map_Internal_size__2382154916() gopurs_runtime.Value {
	once_Data_Map_Internal_size__2382154916.Do(func() {
		cache_Data_Map_Internal_size__2382154916 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int(Call_Data_Map_Internal_size__2382154916(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](v_0_box)))
})
	})
	return cache_Data_Map_Internal_size__2382154916
}

var cache_Data_Map_Internal_stepAsc__2098920977 gopurs_runtime.Value
var once_Data_Map_Internal_stepAsc__2098920977 sync.Once
func Get_Data_Map_Internal_stepAsc__2098920977() gopurs_runtime.Value {
	once_Data_Map_Internal_stepAsc__2098920977.Do(func() {
		cache_Data_Map_Internal_stepAsc__2098920977 = Call_Data_Map_Internal_stepWith(Get_Data_Map_Internal_iterMapL(), gopurs_runtime.Func3(func(__local_var_0 gopurs_runtime.Value, __local_var_1 gopurs_runtime.Value, __local_var_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 953589075, UnsafePtr: unsafe.Pointer(&Constructor_Data_Map_Internal_IterNext{1, __local_var_0, __local_var_1, __local_var_2})}
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 953589075, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_IterNext)(nil))}
}))
	})
	return cache_Data_Map_Internal_stepAsc__2098920977
}

var cache_Data_Map_Internal_stepAscCps__4029696238 gopurs_runtime.Value
var once_Data_Map_Internal_stepAscCps__4029696238 sync.Once
func Get_Data_Map_Internal_stepAscCps__4029696238() gopurs_runtime.Value {
	once_Data_Map_Internal_stepAscCps__4029696238.Do(func() {
		cache_Data_Map_Internal_stepAscCps__4029696238 = gopurs_runtime.Apply(Get_Data_Map_Internal_stepWith(), Get_Data_Map_Internal_iterMapL())
	})
	return cache_Data_Map_Internal_stepAscCps__4029696238
}

var cache_Data_Map_Internal_stepAscCps__3405418095 gopurs_runtime.Value
var once_Data_Map_Internal_stepAscCps__3405418095 sync.Once
func Get_Data_Map_Internal_stepAscCps__3405418095() gopurs_runtime.Value {
	once_Data_Map_Internal_stepAscCps__3405418095.Do(func() {
		cache_Data_Map_Internal_stepAscCps__3405418095 = gopurs_runtime.Apply(Get_Data_Map_Internal_stepWith(), Get_Data_Map_Internal_iterMapL())
	})
	return cache_Data_Map_Internal_stepAscCps__3405418095
}

var cache_Data_Map_Internal_stepAscCps__3090303421 gopurs_runtime.Value
var once_Data_Map_Internal_stepAscCps__3090303421 sync.Once
func Get_Data_Map_Internal_stepAscCps__3090303421() gopurs_runtime.Value {
	once_Data_Map_Internal_stepAscCps__3090303421.Do(func() {
		cache_Data_Map_Internal_stepAscCps__3090303421 = gopurs_runtime.Apply(Get_Data_Map_Internal_stepWith(), Get_Data_Map_Internal_iterMapL())
	})
	return cache_Data_Map_Internal_stepAscCps__3090303421
}

var cache_Data_Map_Internal_stepAscCps__2463496949 gopurs_runtime.Value
var once_Data_Map_Internal_stepAscCps__2463496949 sync.Once
func Get_Data_Map_Internal_stepAscCps__2463496949() gopurs_runtime.Value {
	once_Data_Map_Internal_stepAscCps__2463496949.Do(func() {
		cache_Data_Map_Internal_stepAscCps__2463496949 = gopurs_runtime.Apply(Get_Data_Map_Internal_stepWith(), Get_Data_Map_Internal_iterMapL())
	})
	return cache_Data_Map_Internal_stepAscCps__2463496949
}

var cache_Data_Map_Internal_stepAscCps__1323290822 gopurs_runtime.Value
var once_Data_Map_Internal_stepAscCps__1323290822 sync.Once
func Get_Data_Map_Internal_stepAscCps__1323290822() gopurs_runtime.Value {
	once_Data_Map_Internal_stepAscCps__1323290822.Do(func() {
		cache_Data_Map_Internal_stepAscCps__1323290822 = gopurs_runtime.Apply(Get_Data_Map_Internal_stepWith(), Get_Data_Map_Internal_iterMapL())
	})
	return cache_Data_Map_Internal_stepAscCps__1323290822
}

var cache_Data_Map_Internal_stepAscCps__895533487 gopurs_runtime.Value
var once_Data_Map_Internal_stepAscCps__895533487 sync.Once
func Get_Data_Map_Internal_stepAscCps__895533487() gopurs_runtime.Value {
	once_Data_Map_Internal_stepAscCps__895533487.Do(func() {
		cache_Data_Map_Internal_stepAscCps__895533487 = gopurs_runtime.Apply(Get_Data_Map_Internal_stepWith(), Get_Data_Map_Internal_iterMapL())
	})
	return cache_Data_Map_Internal_stepAscCps__895533487
}

var cache_Data_Map_Internal_stepAscCps__212019860 gopurs_runtime.Value
var once_Data_Map_Internal_stepAscCps__212019860 sync.Once
func Get_Data_Map_Internal_stepAscCps__212019860() gopurs_runtime.Value {
	once_Data_Map_Internal_stepAscCps__212019860.Do(func() {
		cache_Data_Map_Internal_stepAscCps__212019860 = gopurs_runtime.Apply(Get_Data_Map_Internal_stepWith(), Get_Data_Map_Internal_iterMapL())
	})
	return cache_Data_Map_Internal_stepAscCps__212019860
}

var cache_Data_Map_Internal_stepDescCps__3090303421 gopurs_runtime.Value
var once_Data_Map_Internal_stepDescCps__3090303421 sync.Once
func Get_Data_Map_Internal_stepDescCps__3090303421() gopurs_runtime.Value {
	once_Data_Map_Internal_stepDescCps__3090303421.Do(func() {
		cache_Data_Map_Internal_stepDescCps__3090303421 = gopurs_runtime.Apply(Get_Data_Map_Internal_stepWith(), Get_Data_Map_Internal_iterMapR())
	})
	return cache_Data_Map_Internal_stepDescCps__3090303421
}

var cache_Data_Map_Internal_stepDescCps__2463496949 gopurs_runtime.Value
var once_Data_Map_Internal_stepDescCps__2463496949 sync.Once
func Get_Data_Map_Internal_stepDescCps__2463496949() gopurs_runtime.Value {
	once_Data_Map_Internal_stepDescCps__2463496949.Do(func() {
		cache_Data_Map_Internal_stepDescCps__2463496949 = gopurs_runtime.Apply(Get_Data_Map_Internal_stepWith(), Get_Data_Map_Internal_iterMapR())
	})
	return cache_Data_Map_Internal_stepDescCps__2463496949
}

var cache_Data_Map_Internal_stepUnfoldr__966001626 gopurs_runtime.Value
var once_Data_Map_Internal_stepUnfoldr__966001626 sync.Once
func Get_Data_Map_Internal_stepUnfoldr__966001626() gopurs_runtime.Value {
	once_Data_Map_Internal_stepUnfoldr__966001626.Do(func() {
		cache_Data_Map_Internal_stepUnfoldr__966001626 = Call_Data_Map_Internal_stepWith(Get_Data_Map_Internal_iterMapL(), gopurs_runtime.Func3(func(__local_var_0 gopurs_runtime.Value, __local_var_1 gopurs_runtime.Value, __local_var_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, __local_var_0, __local_var_1})}, __local_var_2})}})}
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))}
}))
	})
	return cache_Data_Map_Internal_stepUnfoldr__966001626
}

var cache_Data_Map_Internal_stepUnfoldr__575593864 gopurs_runtime.Value
var once_Data_Map_Internal_stepUnfoldr__575593864 sync.Once
func Get_Data_Map_Internal_stepUnfoldr__575593864() gopurs_runtime.Value {
	once_Data_Map_Internal_stepUnfoldr__575593864.Do(func() {
		cache_Data_Map_Internal_stepUnfoldr__575593864 = Call_Data_Map_Internal_stepWith(Get_Data_Map_Internal_iterMapL(), gopurs_runtime.Func3(func(__local_var_0 gopurs_runtime.Value, __local_var_1 gopurs_runtime.Value, __local_var_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(__local_var_0.IntVal)), UnsafePtr: nil}, gopurs_runtime.Float(__local_var_1.FloatVal())})}, __local_var_2})}})}
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))}
}))
	})
	return cache_Data_Map_Internal_stepUnfoldr__575593864
}

var cache_Data_Map_Internal_stepUnfoldrUnordered__966001626 gopurs_runtime.Value
var once_Data_Map_Internal_stepUnfoldrUnordered__966001626 sync.Once
func Get_Data_Map_Internal_stepUnfoldrUnordered__966001626() gopurs_runtime.Value {
	once_Data_Map_Internal_stepUnfoldrUnordered__966001626.Do(func() {
		cache_Data_Map_Internal_stepUnfoldrUnordered__966001626 = Call_Data_Map_Internal_stepWith(Get_Data_Map_Internal_iterMapU(), gopurs_runtime.Func3(func(__local_var_0 gopurs_runtime.Value, __local_var_1 gopurs_runtime.Value, __local_var_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, __local_var_0, __local_var_1})}, __local_var_2})}})}
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))}
}))
	})
	return cache_Data_Map_Internal_stepUnfoldrUnordered__966001626
}

var cache_Data_Map_Internal_stepUnorderedCps__3090303421 gopurs_runtime.Value
var once_Data_Map_Internal_stepUnorderedCps__3090303421 sync.Once
func Get_Data_Map_Internal_stepUnorderedCps__3090303421() gopurs_runtime.Value {
	once_Data_Map_Internal_stepUnorderedCps__3090303421.Do(func() {
		cache_Data_Map_Internal_stepUnorderedCps__3090303421 = gopurs_runtime.Apply(Get_Data_Map_Internal_stepWith(), Get_Data_Map_Internal_iterMapU())
	})
	return cache_Data_Map_Internal_stepUnorderedCps__3090303421
}

var cache_Data_Map_Internal_stepUnorderedCps__2463496949 gopurs_runtime.Value
var once_Data_Map_Internal_stepUnorderedCps__2463496949 sync.Once
func Get_Data_Map_Internal_stepUnorderedCps__2463496949() gopurs_runtime.Value {
	once_Data_Map_Internal_stepUnorderedCps__2463496949.Do(func() {
		cache_Data_Map_Internal_stepUnorderedCps__2463496949 = gopurs_runtime.Apply(Get_Data_Map_Internal_stepWith(), Get_Data_Map_Internal_iterMapU())
	})
	return cache_Data_Map_Internal_stepUnorderedCps__2463496949
}

var cache_Data_Map_Internal_stepUnorderedCps__1323290822 gopurs_runtime.Value
var once_Data_Map_Internal_stepUnorderedCps__1323290822 sync.Once
func Get_Data_Map_Internal_stepUnorderedCps__1323290822() gopurs_runtime.Value {
	once_Data_Map_Internal_stepUnorderedCps__1323290822.Do(func() {
		cache_Data_Map_Internal_stepUnorderedCps__1323290822 = gopurs_runtime.Apply(Get_Data_Map_Internal_stepWith(), Get_Data_Map_Internal_iterMapU())
	})
	return cache_Data_Map_Internal_stepUnorderedCps__1323290822
}

var cache_Data_Map_Internal_stepWith__2632420966 gopurs_runtime.Value
var once_Data_Map_Internal_stepWith__2632420966 sync.Once
func Get_Data_Map_Internal_stepWith__2632420966() gopurs_runtime.Value {
	once_Data_Map_Internal_stepWith__2632420966.Do(func() {
		cache_Data_Map_Internal_stepWith__2632420966 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, next_1_box gopurs_runtime.Value, done_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Map_Internal_stepWith__2632420966(f_0_box, next_1_box, done_2_box)
})
	})
	return cache_Data_Map_Internal_stepWith__2632420966
}

var cache_Data_Map_Internal_stepWith__603436967 gopurs_runtime.Value
var once_Data_Map_Internal_stepWith__603436967 sync.Once
func Get_Data_Map_Internal_stepWith__603436967() gopurs_runtime.Value {
	once_Data_Map_Internal_stepWith__603436967.Do(func() {
		cache_Data_Map_Internal_stepWith__603436967 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, next_1_box gopurs_runtime.Value, done_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Map_Internal_stepWith__603436967(f_0_box, next_1_box, done_2_box)
})
	})
	return cache_Data_Map_Internal_stepWith__603436967
}

var cache_Data_Map_Internal_stepWith__3186376421 gopurs_runtime.Value
var once_Data_Map_Internal_stepWith__3186376421 sync.Once
func Get_Data_Map_Internal_stepWith__3186376421() gopurs_runtime.Value {
	once_Data_Map_Internal_stepWith__3186376421.Do(func() {
		cache_Data_Map_Internal_stepWith__3186376421 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, next_1_box gopurs_runtime.Value, done_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Map_Internal_stepWith__3186376421(f_0_box, next_1_box, done_2_box)
})
	})
	return cache_Data_Map_Internal_stepWith__3186376421
}

var cache_Data_Map_Internal_stepWith__2866328237 gopurs_runtime.Value
var once_Data_Map_Internal_stepWith__2866328237 sync.Once
func Get_Data_Map_Internal_stepWith__2866328237() gopurs_runtime.Value {
	once_Data_Map_Internal_stepWith__2866328237.Do(func() {
		cache_Data_Map_Internal_stepWith__2866328237 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, next_1_box gopurs_runtime.Value, done_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Map_Internal_stepWith__2866328237(f_0_box, next_1_box, done_2_box)
})
	})
	return cache_Data_Map_Internal_stepWith__2866328237
}

var cache_Data_Map_Internal_stepWith__280335550 gopurs_runtime.Value
var once_Data_Map_Internal_stepWith__280335550 sync.Once
func Get_Data_Map_Internal_stepWith__280335550() gopurs_runtime.Value {
	once_Data_Map_Internal_stepWith__280335550.Do(func() {
		cache_Data_Map_Internal_stepWith__280335550 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, next_1_box gopurs_runtime.Value, done_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Map_Internal_stepWith__280335550(f_0_box, next_1_box, done_2_box)
})
	})
	return cache_Data_Map_Internal_stepWith__280335550
}

var cache_Data_Map_Internal_stepWith__2834533669 gopurs_runtime.Value
var once_Data_Map_Internal_stepWith__2834533669 sync.Once
func Get_Data_Map_Internal_stepWith__2834533669() gopurs_runtime.Value {
	once_Data_Map_Internal_stepWith__2834533669.Do(func() {
		cache_Data_Map_Internal_stepWith__2834533669 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, next_1_box gopurs_runtime.Value, done_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Map_Internal_stepWith__2834533669(f_0_box, next_1_box, done_2_box)
})
	})
	return cache_Data_Map_Internal_stepWith__2834533669
}

var cache_Data_Map_Internal_stepWith__1463181374 gopurs_runtime.Value
var once_Data_Map_Internal_stepWith__1463181374 sync.Once
func Get_Data_Map_Internal_stepWith__1463181374() gopurs_runtime.Value {
	once_Data_Map_Internal_stepWith__1463181374.Do(func() {
		cache_Data_Map_Internal_stepWith__1463181374 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, next_1_box gopurs_runtime.Value, done_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Map_Internal_stepWith__1463181374(f_0_box, next_1_box, done_2_box)
})
	})
	return cache_Data_Map_Internal_stepWith__1463181374
}

var cache_Data_Map_Internal_toMapIter__1799172593 gopurs_runtime.Value
var once_Data_Map_Internal_toMapIter__1799172593 sync.Once
func Get_Data_Map_Internal_toMapIter__1799172593() gopurs_runtime.Value {
	once_Data_Map_Internal_toMapIter__1799172593.Do(func() {
		cache_Data_Map_Internal_toMapIter__1799172593 = gopurs_runtime.Func(func(a_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Map_Internal_toMapIter__1799172593(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](a_0_box))
})
	})
	return cache_Data_Map_Internal_toMapIter__1799172593
}

var cache_Data_Map_Internal_toMapIter__2014410513 gopurs_runtime.Value
var once_Data_Map_Internal_toMapIter__2014410513 sync.Once
func Get_Data_Map_Internal_toMapIter__2014410513() gopurs_runtime.Value {
	once_Data_Map_Internal_toMapIter__2014410513.Do(func() {
		cache_Data_Map_Internal_toMapIter__2014410513 = gopurs_runtime.Func(func(a_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Map_Internal_toMapIter__2014410513(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](a_0_box))
})
	})
	return cache_Data_Map_Internal_toMapIter__2014410513
}

var cache_Data_Map_Internal_toMapIter__772765521 gopurs_runtime.Value
var once_Data_Map_Internal_toMapIter__772765521 sync.Once
func Get_Data_Map_Internal_toMapIter__772765521() gopurs_runtime.Value {
	once_Data_Map_Internal_toMapIter__772765521.Do(func() {
		cache_Data_Map_Internal_toMapIter__772765521 = gopurs_runtime.Func(func(a_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Map_Internal_toMapIter__772765521(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](a_0_box))
})
	})
	return cache_Data_Map_Internal_toMapIter__772765521
}

var cache_Data_Map_Internal_toMapIter__1738891721 gopurs_runtime.Value
var once_Data_Map_Internal_toMapIter__1738891721 sync.Once
func Get_Data_Map_Internal_toMapIter__1738891721() gopurs_runtime.Value {
	once_Data_Map_Internal_toMapIter__1738891721.Do(func() {
		cache_Data_Map_Internal_toMapIter__1738891721 = gopurs_runtime.Func(func(a_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Map_Internal_toMapIter__1738891721(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](a_0_box))
})
	})
	return cache_Data_Map_Internal_toMapIter__1738891721
}

var cache_Data_Map_Internal_toUnfoldable__2183602684 gopurs_runtime.Value
var once_Data_Map_Internal_toUnfoldable__2183602684 sync.Once
func Get_Data_Map_Internal_toUnfoldable__2183602684() gopurs_runtime.Value {
	once_Data_Map_Internal_toUnfoldable__2183602684.Do(func() {
		cache_Data_Map_Internal_toUnfoldable__2183602684 = gopurs_runtime.Func(func(dictUnfoldable_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Map_Internal_toUnfoldable__2183602684(gopurs_runtime.CoerceToStruct[Constructor_Data_Unfoldable_Unfoldable](dictUnfoldable_0_box))
})
	})
	return cache_Data_Map_Internal_toUnfoldable__2183602684
}

var cache_Data_Map_Internal_toUnfoldable__2567957978 gopurs_runtime.Value
var once_Data_Map_Internal_toUnfoldable__2567957978 sync.Once
func Get_Data_Map_Internal_toUnfoldable__2567957978() gopurs_runtime.Value {
	once_Data_Map_Internal_toUnfoldable__2567957978.Do(func() {
		cache_Data_Map_Internal_toUnfoldable__2567957978 = gopurs_runtime.Func(func(__eta0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Map_Internal_toUnfoldable__2567957978(__eta0_0_box)
})
	})
	return cache_Data_Map_Internal_toUnfoldable__2567957978
}

var cache_Data_Map_Internal_traversableMap__1002539403 gopurs_runtime.Value
var once_Data_Map_Internal_traversableMap__1002539403 sync.Once
func Get_Data_Map_Internal_traversableMap__1002539403() gopurs_runtime.Value {
	once_Data_Map_Internal_traversableMap__1002539403.Do(func() {
		cache_Data_Map_Internal_traversableMap__1002539403 = gopurs_runtime.RecordDict4("Foldable1", "Functor0", "sequence", "traverse", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict3("foldMap", "foldl", "foldr", gopurs_runtime.Func(func(dictMonoid_1 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Semigroup0_2_0 -> *Constructor_Data_Semigroup_Semigroup
Semigroup0_2_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_1, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_2_0
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_4_1_133 gopurs_runtime.Value
_ = go__go_4_1_133
go__go_4_1_133 = gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 gopurs_runtime.Value
{
if (v_5.Type == 9 && v_5.IntVal == 324739070 && v_5.UnsafePtr == nil) {
__t2 = gopurs_runtime.RecordGet(dictMonoid_1, "mempty")
goto end_branch_2
} else {

}
}
{
if (v_5.Type == 9 && v_5.IntVal == 324739070 && v_5.UnsafePtr != nil) {
__t2 = gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_2_0.V0), gopurs_runtime.Apply(go__go_4_1_133, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v_5.UnsafePtr).V4)}), gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_2_0.V0), gopurs_runtime.Apply(f_3, (*Constructor_Data_Map_Internal_Node)(v_5.UnsafePtr).V3), gopurs_runtime.Apply(go__go_4_1_133, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v_5.UnsafePtr).V5)})))
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
return go__go_4_1_133
})
}), gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(z_2 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_3_3_134 gopurs_runtime.Value
go__go_3_3_134 = gopurs_runtime.Func(func(__local_var_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(__local_var_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var __local_var_4_loop gopurs_runtime.Value = __local_var_4_loop_val
var __local_var_5_loop gopurs_runtime.Value = __local_var_5_loop_val
go__go_3_3_134:
for {
if false { continue go__go_3_3_134 }
var __local_var_4 gopurs_runtime.Value = __local_var_4_loop
_ = __local_var_4
var __local_var_5 gopurs_runtime.Value = __local_var_5_loop
_ = __local_var_5
var __t4 gopurs_runtime.Value
{
if (__local_var_5.Type == 9 && __local_var_5.IntVal == 324739070 && __local_var_5.UnsafePtr == nil) {
__t4 = __local_var_4
goto end_branch_4
} else {

}
}
{
if (__local_var_5.Type == 9 && __local_var_5.IntVal == 324739070 && __local_var_5.UnsafePtr != nil) {
__t4 = gopurs_runtime.UncurriedApp2(go__go_3_3_134, gopurs_runtime.Apply2(f_1, gopurs_runtime.UncurriedApp2(go__go_3_3_134, __local_var_4, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_5.UnsafePtr).V4)}), (*Constructor_Data_Map_Internal_Node)(__local_var_5.UnsafePtr).V3), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_5.UnsafePtr).V5)})
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
return gopurs_runtime.Func(func(m_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.UncurriedApp2(go__go_3_3_134, z_2, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](m_4))})
})
})
}), gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(z_2 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_3_5_135 gopurs_runtime.Value
_ = go__go_3_5_135
go__go_3_5_135 = gopurs_runtime.Func2(func(__local_var_4 gopurs_runtime.Value, __local_var_5 gopurs_runtime.Value) gopurs_runtime.Value {
var __t6 gopurs_runtime.Value
{
if (__local_var_4.Type == 9 && __local_var_4.IntVal == 324739070 && __local_var_4.UnsafePtr == nil) {
__t6 = __local_var_5
goto end_branch_6
} else {

}
}
{
if (__local_var_4.Type == 9 && __local_var_4.IntVal == 324739070 && __local_var_4.UnsafePtr != nil) {
__t6 = gopurs_runtime.UncurriedApp2(go__go_3_5_135, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_4.UnsafePtr).V4)}, gopurs_runtime.Apply2(f_1, (*Constructor_Data_Map_Internal_Node)(__local_var_4.UnsafePtr).V3, gopurs_runtime.UncurriedApp2(go__go_3_5_135, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_4.UnsafePtr).V5)}, __local_var_5)))
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
return gopurs_runtime.Func(func(m_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.UncurriedApp2(go__go_3_5_135, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](m_4))}, z_2)
})
})
}))
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_2_7_136 gopurs_runtime.Value
_ = go__go_2_7_136
go__go_2_7_136 = gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t8 *Constructor_Data_Map_Internal_Node
{
if (v_3.Type == 9 && v_3.IntVal == 324739070 && v_3.UnsafePtr == nil) {
__t8 = (*Constructor_Data_Map_Internal_Node)(nil)
goto end_branch_8
} else {

}
}
{
if (v_3.Type == 9 && v_3.IntVal == 324739070 && v_3.UnsafePtr != nil) {
__t8 = &Constructor_Data_Map_Internal_Node{1, (*Constructor_Data_Map_Internal_Node)(v_3.UnsafePtr).V0, (*Constructor_Data_Map_Internal_Node)(v_3.UnsafePtr).V1, (*Constructor_Data_Map_Internal_Node)(v_3.UnsafePtr).V2, gopurs_runtime.Apply(f_1, (*Constructor_Data_Map_Internal_Node)(v_3.UnsafePtr).V3), gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.Apply(go__go_2_7_136, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v_3.UnsafePtr).V4)})), gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.Apply(go__go_2_7_136, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v_3.UnsafePtr).V5)}))}
goto end_branch_8
} else {

}
}
{
__t8 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_8:
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(__t8)}
})
return go__go_2_7_136
}))
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Apply0_1_9 -> *Constructor_Control_Apply_Apply
Apply0_1_9 := gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}))
_ = Apply0_1_9
// TAST (Let): Functor0_2_10 -> *Constructor_Data_Functor_Functor
Functor0_2_10 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_2_10
var go__go_3_11_137 gopurs_runtime.Value
_ = go__go_3_11_137
go__go_3_11_137 = gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t15 gopurs_runtime.Value
{
if (v_4.Type == 9 && v_4.IntVal == 324739070 && v_4.UnsafePtr == nil) {
__t15 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "pure"), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(nil))})
goto end_branch_15
} else {

}
}
{
if (v_4.Type == 9 && v_4.IntVal == 324739070 && v_4.UnsafePtr != nil) {
// TAST (Let): __local_var_5_12 -> gopurs_runtime.Value
var __local_var_5_12 gopurs_runtime.Value = gopurs_runtime.Int((*Constructor_Data_Map_Internal_Node)(v_4.UnsafePtr).V0)
// TAST (Let): __local_var_6_13 -> gopurs_runtime.Value
__local_var_6_13 := (*Constructor_Data_Map_Internal_Node)(v_4.UnsafePtr).V2
_ = __local_var_6_13
// TAST (Let): __local_var_7_14 -> gopurs_runtime.Value
var __local_var_7_14 gopurs_runtime.Value = gopurs_runtime.Int((*Constructor_Data_Map_Internal_Node)(v_4.UnsafePtr).V1)
__t15 = gopurs_runtime.Apply2(gopurs_runtime.Box(Apply0_1_9.V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Apply0_1_9.V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_2_10.V0), gopurs_runtime.Func(func(l_prime_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_prime_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(r_prime_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(&Constructor_Data_Map_Internal_Node{1, __local_var_5_12.IntVal, __local_var_7_14.IntVal, __local_var_6_13, v_prime_9, gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](l_prime_8), gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](r_prime_10)})}
})
})
}), gopurs_runtime.Apply(go__go_3_11_137, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v_4.UnsafePtr).V4)})), (*Constructor_Data_Map_Internal_Node)(v_4.UnsafePtr).V3), gopurs_runtime.Apply(go__go_3_11_137, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v_4.UnsafePtr).V5)}))
goto end_branch_15
} else {

}
}
{
__t15 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_15:
return __t15
})
return go__go_3_11_137
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Apply0_1_16 -> *Constructor_Control_Apply_Apply
Apply0_1_16 := gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}))
_ = Apply0_1_16
// TAST (Let): Functor0_2_17 -> *Constructor_Data_Functor_Functor
Functor0_2_17 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_2_17
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_4_18_138 gopurs_runtime.Value
_ = go__go_4_18_138
go__go_4_18_138 = gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
var __t22 gopurs_runtime.Value
{
if (v_5.Type == 9 && v_5.IntVal == 324739070 && v_5.UnsafePtr == nil) {
__t22 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "pure"), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(nil))})
goto end_branch_22
} else {

}
}
{
if (v_5.Type == 9 && v_5.IntVal == 324739070 && v_5.UnsafePtr != nil) {
// TAST (Let): __local_var_6_19 -> gopurs_runtime.Value
var __local_var_6_19 gopurs_runtime.Value = gopurs_runtime.Int((*Constructor_Data_Map_Internal_Node)(v_5.UnsafePtr).V0)
// TAST (Let): __local_var_7_20 -> gopurs_runtime.Value
__local_var_7_20 := (*Constructor_Data_Map_Internal_Node)(v_5.UnsafePtr).V2
_ = __local_var_7_20
// TAST (Let): __local_var_8_21 -> gopurs_runtime.Value
var __local_var_8_21 gopurs_runtime.Value = gopurs_runtime.Int((*Constructor_Data_Map_Internal_Node)(v_5.UnsafePtr).V1)
__t22 = gopurs_runtime.Apply2(gopurs_runtime.Box(Apply0_1_16.V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Apply0_1_16.V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_2_17.V0), gopurs_runtime.Func(func(l_prime_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_prime_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(r_prime_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(&Constructor_Data_Map_Internal_Node{1, __local_var_6_19.IntVal, __local_var_8_21.IntVal, __local_var_7_20, v_prime_10, gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](l_prime_9), gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](r_prime_11)})}
})
})
}), gopurs_runtime.Apply(go__go_4_18_138, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v_5.UnsafePtr).V4)})), gopurs_runtime.Apply(f_3, (*Constructor_Data_Map_Internal_Node)(v_5.UnsafePtr).V3)), gopurs_runtime.Apply(go__go_4_18_138, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v_5.UnsafePtr).V5)}))
goto end_branch_22
} else {

}
}
{
__t22 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_22:
return __t22
})
return go__go_4_18_138
})
}))
	})
	return cache_Data_Map_Internal_traversableMap__1002539403
}

var cache_Data_Map_Internal_traversableMap__2256206635 gopurs_runtime.Value
var once_Data_Map_Internal_traversableMap__2256206635 sync.Once
func Get_Data_Map_Internal_traversableMap__2256206635() gopurs_runtime.Value {
	once_Data_Map_Internal_traversableMap__2256206635.Do(func() {
		cache_Data_Map_Internal_traversableMap__2256206635 = gopurs_runtime.RecordDict4("Foldable1", "Functor0", "sequence", "traverse", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict3("foldMap", "foldl", "foldr", gopurs_runtime.Func(func(dictMonoid_1 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Semigroup0_2_0 -> *Constructor_Data_Semigroup_Semigroup
Semigroup0_2_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_1, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_2_0
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_4_1_139 gopurs_runtime.Value
_ = go__go_4_1_139
go__go_4_1_139 = gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 gopurs_runtime.Value
{
if (v_5.Type == 9 && v_5.IntVal == 324739070 && v_5.UnsafePtr == nil) {
__t2 = gopurs_runtime.RecordGet(dictMonoid_1, "mempty")
goto end_branch_2
} else {

}
}
{
if (v_5.Type == 9 && v_5.IntVal == 324739070 && v_5.UnsafePtr != nil) {
__t2 = gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_2_0.V0), gopurs_runtime.Apply(go__go_4_1_139, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v_5.UnsafePtr).V4)}), gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_2_0.V0), gopurs_runtime.Apply(f_3, (*Constructor_Data_Map_Internal_Node)(v_5.UnsafePtr).V3), gopurs_runtime.Apply(go__go_4_1_139, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v_5.UnsafePtr).V5)})))
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
return go__go_4_1_139
})
}), gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(z_2 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_3_3_140 gopurs_runtime.Value
go__go_3_3_140 = gopurs_runtime.Func(func(__local_var_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(__local_var_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var __local_var_4_loop gopurs_runtime.Value = __local_var_4_loop_val
var __local_var_5_loop gopurs_runtime.Value = __local_var_5_loop_val
go__go_3_3_140:
for {
if false { continue go__go_3_3_140 }
var __local_var_4 gopurs_runtime.Value = __local_var_4_loop
_ = __local_var_4
var __local_var_5 gopurs_runtime.Value = __local_var_5_loop
_ = __local_var_5
var __t4 gopurs_runtime.Value
{
if (__local_var_5.Type == 9 && __local_var_5.IntVal == 324739070 && __local_var_5.UnsafePtr == nil) {
__t4 = __local_var_4
goto end_branch_4
} else {

}
}
{
if (__local_var_5.Type == 9 && __local_var_5.IntVal == 324739070 && __local_var_5.UnsafePtr != nil) {
__t4 = gopurs_runtime.UncurriedApp2(go__go_3_3_140, gopurs_runtime.Apply2(f_1, gopurs_runtime.UncurriedApp2(go__go_3_3_140, __local_var_4, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_5.UnsafePtr).V4)}), (*Constructor_Data_Map_Internal_Node)(__local_var_5.UnsafePtr).V3), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_5.UnsafePtr).V5)})
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
return gopurs_runtime.Func(func(m_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.UncurriedApp2(go__go_3_3_140, z_2, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](m_4))})
})
})
}), gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(z_2 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_3_5_141 gopurs_runtime.Value
_ = go__go_3_5_141
go__go_3_5_141 = gopurs_runtime.Func2(func(__local_var_4 gopurs_runtime.Value, __local_var_5 gopurs_runtime.Value) gopurs_runtime.Value {
var __t6 gopurs_runtime.Value
{
if (__local_var_4.Type == 9 && __local_var_4.IntVal == 324739070 && __local_var_4.UnsafePtr == nil) {
__t6 = __local_var_5
goto end_branch_6
} else {

}
}
{
if (__local_var_4.Type == 9 && __local_var_4.IntVal == 324739070 && __local_var_4.UnsafePtr != nil) {
__t6 = gopurs_runtime.UncurriedApp2(go__go_3_5_141, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_4.UnsafePtr).V4)}, gopurs_runtime.Apply2(f_1, (*Constructor_Data_Map_Internal_Node)(__local_var_4.UnsafePtr).V3, gopurs_runtime.UncurriedApp2(go__go_3_5_141, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_4.UnsafePtr).V5)}, __local_var_5)))
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
return gopurs_runtime.Func(func(m_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.UncurriedApp2(go__go_3_5_141, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](m_4))}, z_2)
})
})
}))
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_2_7_142 gopurs_runtime.Value
_ = go__go_2_7_142
go__go_2_7_142 = gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t8 *Constructor_Data_Map_Internal_Node
{
if (v_3.Type == 9 && v_3.IntVal == 324739070 && v_3.UnsafePtr == nil) {
__t8 = (*Constructor_Data_Map_Internal_Node)(nil)
goto end_branch_8
} else {

}
}
{
if (v_3.Type == 9 && v_3.IntVal == 324739070 && v_3.UnsafePtr != nil) {
__t8 = &Constructor_Data_Map_Internal_Node{1, (*Constructor_Data_Map_Internal_Node)(v_3.UnsafePtr).V0, (*Constructor_Data_Map_Internal_Node)(v_3.UnsafePtr).V1, (*Constructor_Data_Map_Internal_Node)(v_3.UnsafePtr).V2, gopurs_runtime.Apply(f_1, (*Constructor_Data_Map_Internal_Node)(v_3.UnsafePtr).V3), gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.Apply(go__go_2_7_142, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v_3.UnsafePtr).V4)})), gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.Apply(go__go_2_7_142, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v_3.UnsafePtr).V5)}))}
goto end_branch_8
} else {

}
}
{
__t8 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_8:
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(__t8)}
})
return go__go_2_7_142
}))
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Apply0_1_9 -> *Constructor_Control_Apply_Apply
Apply0_1_9 := gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}))
_ = Apply0_1_9
// TAST (Let): Functor0_2_10 -> *Constructor_Data_Functor_Functor
Functor0_2_10 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_2_10
var go__go_3_11_143 gopurs_runtime.Value
_ = go__go_3_11_143
go__go_3_11_143 = gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t15 gopurs_runtime.Value
{
if (v_4.Type == 9 && v_4.IntVal == 324739070 && v_4.UnsafePtr == nil) {
__t15 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "pure"), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(nil))})
goto end_branch_15
} else {

}
}
{
if (v_4.Type == 9 && v_4.IntVal == 324739070 && v_4.UnsafePtr != nil) {
// TAST (Let): __local_var_5_12 -> gopurs_runtime.Value
var __local_var_5_12 gopurs_runtime.Value = gopurs_runtime.Int((*Constructor_Data_Map_Internal_Node)(v_4.UnsafePtr).V0)
// TAST (Let): __local_var_6_13 -> gopurs_runtime.Value
__local_var_6_13 := (*Constructor_Data_Map_Internal_Node)(v_4.UnsafePtr).V2
_ = __local_var_6_13
// TAST (Let): __local_var_7_14 -> gopurs_runtime.Value
var __local_var_7_14 gopurs_runtime.Value = gopurs_runtime.Int((*Constructor_Data_Map_Internal_Node)(v_4.UnsafePtr).V1)
__t15 = gopurs_runtime.Apply2(gopurs_runtime.Box(Apply0_1_9.V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Apply0_1_9.V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_2_10.V0), gopurs_runtime.Func(func(l_prime_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_prime_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(r_prime_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(&Constructor_Data_Map_Internal_Node{1, __local_var_5_12.IntVal, __local_var_7_14.IntVal, __local_var_6_13, v_prime_9, gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](l_prime_8), gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](r_prime_10)})}
})
})
}), gopurs_runtime.Apply(go__go_3_11_143, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v_4.UnsafePtr).V4)})), (*Constructor_Data_Map_Internal_Node)(v_4.UnsafePtr).V3), gopurs_runtime.Apply(go__go_3_11_143, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v_4.UnsafePtr).V5)}))
goto end_branch_15
} else {

}
}
{
__t15 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_15:
return __t15
})
return go__go_3_11_143
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Apply0_1_16 -> *Constructor_Control_Apply_Apply
Apply0_1_16 := gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}))
_ = Apply0_1_16
// TAST (Let): Functor0_2_17 -> *Constructor_Data_Functor_Functor
Functor0_2_17 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_2_17
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_4_18_144 gopurs_runtime.Value
_ = go__go_4_18_144
go__go_4_18_144 = gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
var __t22 gopurs_runtime.Value
{
if (v_5.Type == 9 && v_5.IntVal == 324739070 && v_5.UnsafePtr == nil) {
__t22 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "pure"), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(nil))})
goto end_branch_22
} else {

}
}
{
if (v_5.Type == 9 && v_5.IntVal == 324739070 && v_5.UnsafePtr != nil) {
// TAST (Let): __local_var_6_19 -> gopurs_runtime.Value
var __local_var_6_19 gopurs_runtime.Value = gopurs_runtime.Int((*Constructor_Data_Map_Internal_Node)(v_5.UnsafePtr).V0)
// TAST (Let): __local_var_7_20 -> gopurs_runtime.Value
__local_var_7_20 := (*Constructor_Data_Map_Internal_Node)(v_5.UnsafePtr).V2
_ = __local_var_7_20
// TAST (Let): __local_var_8_21 -> gopurs_runtime.Value
var __local_var_8_21 gopurs_runtime.Value = gopurs_runtime.Int((*Constructor_Data_Map_Internal_Node)(v_5.UnsafePtr).V1)
__t22 = gopurs_runtime.Apply2(gopurs_runtime.Box(Apply0_1_16.V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Apply0_1_16.V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_2_17.V0), gopurs_runtime.Func(func(l_prime_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_prime_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(r_prime_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(&Constructor_Data_Map_Internal_Node{1, __local_var_6_19.IntVal, __local_var_8_21.IntVal, __local_var_7_20, v_prime_10, gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](l_prime_9), gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](r_prime_11)})}
})
})
}), gopurs_runtime.Apply(go__go_4_18_144, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v_5.UnsafePtr).V4)})), gopurs_runtime.Apply(f_3, (*Constructor_Data_Map_Internal_Node)(v_5.UnsafePtr).V3)), gopurs_runtime.Apply(go__go_4_18_144, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v_5.UnsafePtr).V5)}))
goto end_branch_22
} else {

}
}
{
__t22 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_22:
return __t22
})
return go__go_4_18_144
})
}))
	})
	return cache_Data_Map_Internal_traversableMap__2256206635
}

var cache_Data_Map_Internal_traversableWithIndexMap__3269014446 gopurs_runtime.Value
var once_Data_Map_Internal_traversableWithIndexMap__3269014446 sync.Once
func Get_Data_Map_Internal_traversableWithIndexMap__3269014446() gopurs_runtime.Value {
	once_Data_Map_Internal_traversableWithIndexMap__3269014446.Do(func() {
		cache_Data_Map_Internal_traversableWithIndexMap__3269014446 = gopurs_runtime.RecordDict4("FoldableWithIndex1", "FunctorWithIndex0", "Traversable2", "traverseWithIndex", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict4("Foldable0", "foldMapWithIndex", "foldlWithIndex", "foldrWithIndex", gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict3("foldMap", "foldl", "foldr", gopurs_runtime.Func(func(dictMonoid_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Semigroup0_3_0 -> *Constructor_Data_Semigroup_Semigroup
Semigroup0_3_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_2, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_3_0
return gopurs_runtime.Func(func(f_4 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_5_1_145 gopurs_runtime.Value
_ = go__go_5_1_145
go__go_5_1_145 = gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 gopurs_runtime.Value
{
if (v_6.Type == 9 && v_6.IntVal == 324739070 && v_6.UnsafePtr == nil) {
__t2 = gopurs_runtime.RecordGet(dictMonoid_2, "mempty")
goto end_branch_2
} else {

}
}
{
if (v_6.Type == 9 && v_6.IntVal == 324739070 && v_6.UnsafePtr != nil) {
__t2 = gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_3_0.V0), gopurs_runtime.Apply(go__go_5_1_145, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v_6.UnsafePtr).V4)}), gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_3_0.V0), gopurs_runtime.Apply(f_4, (*Constructor_Data_Map_Internal_Node)(v_6.UnsafePtr).V3), gopurs_runtime.Apply(go__go_5_1_145, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v_6.UnsafePtr).V5)})))
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
return go__go_5_1_145
})
}), gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(z_3 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_4_3_146 gopurs_runtime.Value
go__go_4_3_146 = gopurs_runtime.Func(func(__local_var_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(__local_var_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var __local_var_5_loop gopurs_runtime.Value = __local_var_5_loop_val
var __local_var_6_loop gopurs_runtime.Value = __local_var_6_loop_val
go__go_4_3_146:
for {
if false { continue go__go_4_3_146 }
var __local_var_5 gopurs_runtime.Value = __local_var_5_loop
_ = __local_var_5
var __local_var_6 gopurs_runtime.Value = __local_var_6_loop
_ = __local_var_6
var __t4 gopurs_runtime.Value
{
if (__local_var_6.Type == 9 && __local_var_6.IntVal == 324739070 && __local_var_6.UnsafePtr == nil) {
__t4 = __local_var_5
goto end_branch_4
} else {

}
}
{
if (__local_var_6.Type == 9 && __local_var_6.IntVal == 324739070 && __local_var_6.UnsafePtr != nil) {
__t4 = gopurs_runtime.UncurriedApp2(go__go_4_3_146, gopurs_runtime.Apply2(f_2, gopurs_runtime.UncurriedApp2(go__go_4_3_146, __local_var_5, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_6.UnsafePtr).V4)}), (*Constructor_Data_Map_Internal_Node)(__local_var_6.UnsafePtr).V3), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_6.UnsafePtr).V5)})
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
return gopurs_runtime.Func(func(m_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.UncurriedApp2(go__go_4_3_146, z_3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](m_5))})
})
})
}), gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(z_3 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_4_5_147 gopurs_runtime.Value
_ = go__go_4_5_147
go__go_4_5_147 = gopurs_runtime.Func2(func(__local_var_5 gopurs_runtime.Value, __local_var_6 gopurs_runtime.Value) gopurs_runtime.Value {
var __t6 gopurs_runtime.Value
{
if (__local_var_5.Type == 9 && __local_var_5.IntVal == 324739070 && __local_var_5.UnsafePtr == nil) {
__t6 = __local_var_6
goto end_branch_6
} else {

}
}
{
if (__local_var_5.Type == 9 && __local_var_5.IntVal == 324739070 && __local_var_5.UnsafePtr != nil) {
__t6 = gopurs_runtime.UncurriedApp2(go__go_4_5_147, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_5.UnsafePtr).V4)}, gopurs_runtime.Apply2(f_2, (*Constructor_Data_Map_Internal_Node)(__local_var_5.UnsafePtr).V3, gopurs_runtime.UncurriedApp2(go__go_4_5_147, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_5.UnsafePtr).V5)}, __local_var_6)))
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
return gopurs_runtime.Func(func(m_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.UncurriedApp2(go__go_4_5_147, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](m_5))}, z_3)
})
})
}))
}), gopurs_runtime.Func(func(dictMonoid_1 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Semigroup0_2_7 -> *Constructor_Data_Semigroup_Semigroup
Semigroup0_2_7 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_1, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_2_7
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_4_8_148 gopurs_runtime.Value
_ = go__go_4_8_148
go__go_4_8_148 = gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
var __t9 gopurs_runtime.Value
{
if (v_5.Type == 9 && v_5.IntVal == 324739070 && v_5.UnsafePtr == nil) {
__t9 = gopurs_runtime.RecordGet(dictMonoid_1, "mempty")
goto end_branch_9
} else {

}
}
{
if (v_5.Type == 9 && v_5.IntVal == 324739070 && v_5.UnsafePtr != nil) {
__t9 = gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_2_7.V0), gopurs_runtime.Apply(go__go_4_8_148, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v_5.UnsafePtr).V4)}), gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_2_7.V0), gopurs_runtime.Apply2(f_3, (*Constructor_Data_Map_Internal_Node)(v_5.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node)(v_5.UnsafePtr).V3), gopurs_runtime.Apply(go__go_4_8_148, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v_5.UnsafePtr).V5)})))
goto end_branch_9
} else {

}
}
{
__t9 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_9:
return __t9
})
return go__go_4_8_148
})
}), gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(z_2 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_3_10_149 gopurs_runtime.Value
go__go_3_10_149 = gopurs_runtime.Func(func(__local_var_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(__local_var_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var __local_var_4_loop gopurs_runtime.Value = __local_var_4_loop_val
var __local_var_5_loop gopurs_runtime.Value = __local_var_5_loop_val
go__go_3_10_149:
for {
if false { continue go__go_3_10_149 }
var __local_var_4 gopurs_runtime.Value = __local_var_4_loop
_ = __local_var_4
var __local_var_5 gopurs_runtime.Value = __local_var_5_loop
_ = __local_var_5
var __t11 gopurs_runtime.Value
{
if (__local_var_5.Type == 9 && __local_var_5.IntVal == 324739070 && __local_var_5.UnsafePtr == nil) {
__t11 = __local_var_4
goto end_branch_11
} else {

}
}
{
if (__local_var_5.Type == 9 && __local_var_5.IntVal == 324739070 && __local_var_5.UnsafePtr != nil) {
__t11 = gopurs_runtime.UncurriedApp2(go__go_3_10_149, gopurs_runtime.Apply3(f_1, (*Constructor_Data_Map_Internal_Node)(__local_var_5.UnsafePtr).V2, gopurs_runtime.UncurriedApp2(go__go_3_10_149, __local_var_4, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_5.UnsafePtr).V4)}), (*Constructor_Data_Map_Internal_Node)(__local_var_5.UnsafePtr).V3), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_5.UnsafePtr).V5)})
goto end_branch_11
} else {

}
}
{
__t11 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_11:
return __t11
}
}()
})
})
return gopurs_runtime.Func(func(m_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.UncurriedApp2(go__go_3_10_149, z_2, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](m_4))})
})
})
}), gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(z_2 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_3_12_150 gopurs_runtime.Value
_ = go__go_3_12_150
go__go_3_12_150 = gopurs_runtime.Func2(func(__local_var_4 gopurs_runtime.Value, __local_var_5 gopurs_runtime.Value) gopurs_runtime.Value {
var __t13 gopurs_runtime.Value
{
if (__local_var_4.Type == 9 && __local_var_4.IntVal == 324739070 && __local_var_4.UnsafePtr == nil) {
__t13 = __local_var_5
goto end_branch_13
} else {

}
}
{
if (__local_var_4.Type == 9 && __local_var_4.IntVal == 324739070 && __local_var_4.UnsafePtr != nil) {
__t13 = gopurs_runtime.UncurriedApp2(go__go_3_12_150, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_4.UnsafePtr).V4)}, gopurs_runtime.Apply3(f_1, (*Constructor_Data_Map_Internal_Node)(__local_var_4.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node)(__local_var_4.UnsafePtr).V3, gopurs_runtime.UncurriedApp2(go__go_3_12_150, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_4.UnsafePtr).V5)}, __local_var_5)))
goto end_branch_13
} else {

}
}
{
__t13 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_13:
return __t13
})
return gopurs_runtime.Func(func(m_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.UncurriedApp2(go__go_3_12_150, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](m_4))}, z_2)
})
})
}))
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Functor0", "mapWithIndex", gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_3_14_151 gopurs_runtime.Value
_ = go__go_3_14_151
go__go_3_14_151 = gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t15 *Constructor_Data_Map_Internal_Node
{
if (v_4.Type == 9 && v_4.IntVal == 324739070 && v_4.UnsafePtr == nil) {
__t15 = (*Constructor_Data_Map_Internal_Node)(nil)
goto end_branch_15
} else {

}
}
{
if (v_4.Type == 9 && v_4.IntVal == 324739070 && v_4.UnsafePtr != nil) {
__t15 = &Constructor_Data_Map_Internal_Node{1, (*Constructor_Data_Map_Internal_Node)(v_4.UnsafePtr).V0, (*Constructor_Data_Map_Internal_Node)(v_4.UnsafePtr).V1, (*Constructor_Data_Map_Internal_Node)(v_4.UnsafePtr).V2, gopurs_runtime.Apply(f_2, (*Constructor_Data_Map_Internal_Node)(v_4.UnsafePtr).V3), gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.Apply(go__go_3_14_151, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v_4.UnsafePtr).V4)})), gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.Apply(go__go_3_14_151, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v_4.UnsafePtr).V5)}))}
goto end_branch_15
} else {

}
}
{
__t15 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_15:
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(__t15)}
})
return go__go_3_14_151
}))
}), gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_2_16_152 gopurs_runtime.Value
_ = go__go_2_16_152
go__go_2_16_152 = gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t17 *Constructor_Data_Map_Internal_Node
{
if (v_3.Type == 9 && v_3.IntVal == 324739070 && v_3.UnsafePtr == nil) {
__t17 = (*Constructor_Data_Map_Internal_Node)(nil)
goto end_branch_17
} else {

}
}
{
if (v_3.Type == 9 && v_3.IntVal == 324739070 && v_3.UnsafePtr != nil) {
__t17 = &Constructor_Data_Map_Internal_Node{1, (*Constructor_Data_Map_Internal_Node)(v_3.UnsafePtr).V0, (*Constructor_Data_Map_Internal_Node)(v_3.UnsafePtr).V1, (*Constructor_Data_Map_Internal_Node)(v_3.UnsafePtr).V2, gopurs_runtime.Apply2(f_1, (*Constructor_Data_Map_Internal_Node)(v_3.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node)(v_3.UnsafePtr).V3), gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.Apply(go__go_2_16_152, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v_3.UnsafePtr).V4)})), gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.Apply(go__go_2_16_152, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v_3.UnsafePtr).V5)}))}
goto end_branch_17
} else {

}
}
{
__t17 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_17:
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(__t17)}
})
return go__go_2_16_152
}))
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict4("Foldable1", "Functor0", "sequence", "traverse", gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict3("foldMap", "foldl", "foldr", gopurs_runtime.Func(func(dictMonoid_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Semigroup0_3_18 -> *Constructor_Data_Semigroup_Semigroup
Semigroup0_3_18 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_2, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_3_18
return gopurs_runtime.Func(func(f_4 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_5_19_153 gopurs_runtime.Value
_ = go__go_5_19_153
go__go_5_19_153 = gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
var __t20 gopurs_runtime.Value
{
if (v_6.Type == 9 && v_6.IntVal == 324739070 && v_6.UnsafePtr == nil) {
__t20 = gopurs_runtime.RecordGet(dictMonoid_2, "mempty")
goto end_branch_20
} else {

}
}
{
if (v_6.Type == 9 && v_6.IntVal == 324739070 && v_6.UnsafePtr != nil) {
__t20 = gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_3_18.V0), gopurs_runtime.Apply(go__go_5_19_153, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v_6.UnsafePtr).V4)}), gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_3_18.V0), gopurs_runtime.Apply(f_4, (*Constructor_Data_Map_Internal_Node)(v_6.UnsafePtr).V3), gopurs_runtime.Apply(go__go_5_19_153, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v_6.UnsafePtr).V5)})))
goto end_branch_20
} else {

}
}
{
__t20 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_20:
return __t20
})
return go__go_5_19_153
})
}), gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(z_3 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_4_21_154 gopurs_runtime.Value
go__go_4_21_154 = gopurs_runtime.Func(func(__local_var_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(__local_var_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var __local_var_5_loop gopurs_runtime.Value = __local_var_5_loop_val
var __local_var_6_loop gopurs_runtime.Value = __local_var_6_loop_val
go__go_4_21_154:
for {
if false { continue go__go_4_21_154 }
var __local_var_5 gopurs_runtime.Value = __local_var_5_loop
_ = __local_var_5
var __local_var_6 gopurs_runtime.Value = __local_var_6_loop
_ = __local_var_6
var __t22 gopurs_runtime.Value
{
if (__local_var_6.Type == 9 && __local_var_6.IntVal == 324739070 && __local_var_6.UnsafePtr == nil) {
__t22 = __local_var_5
goto end_branch_22
} else {

}
}
{
if (__local_var_6.Type == 9 && __local_var_6.IntVal == 324739070 && __local_var_6.UnsafePtr != nil) {
__t22 = gopurs_runtime.UncurriedApp2(go__go_4_21_154, gopurs_runtime.Apply2(f_2, gopurs_runtime.UncurriedApp2(go__go_4_21_154, __local_var_5, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_6.UnsafePtr).V4)}), (*Constructor_Data_Map_Internal_Node)(__local_var_6.UnsafePtr).V3), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_6.UnsafePtr).V5)})
goto end_branch_22
} else {

}
}
{
__t22 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_22:
return __t22
}
}()
})
})
return gopurs_runtime.Func(func(m_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.UncurriedApp2(go__go_4_21_154, z_3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](m_5))})
})
})
}), gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(z_3 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_4_23_155 gopurs_runtime.Value
_ = go__go_4_23_155
go__go_4_23_155 = gopurs_runtime.Func2(func(__local_var_5 gopurs_runtime.Value, __local_var_6 gopurs_runtime.Value) gopurs_runtime.Value {
var __t24 gopurs_runtime.Value
{
if (__local_var_5.Type == 9 && __local_var_5.IntVal == 324739070 && __local_var_5.UnsafePtr == nil) {
__t24 = __local_var_6
goto end_branch_24
} else {

}
}
{
if (__local_var_5.Type == 9 && __local_var_5.IntVal == 324739070 && __local_var_5.UnsafePtr != nil) {
__t24 = gopurs_runtime.UncurriedApp2(go__go_4_23_155, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_5.UnsafePtr).V4)}, gopurs_runtime.Apply2(f_2, (*Constructor_Data_Map_Internal_Node)(__local_var_5.UnsafePtr).V3, gopurs_runtime.UncurriedApp2(go__go_4_23_155, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_5.UnsafePtr).V5)}, __local_var_6)))
goto end_branch_24
} else {

}
}
{
__t24 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_24:
return __t24
})
return gopurs_runtime.Func(func(m_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.UncurriedApp2(go__go_4_23_155, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](m_5))}, z_3)
})
})
}))
}), gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_3_25_156 gopurs_runtime.Value
_ = go__go_3_25_156
go__go_3_25_156 = gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t26 *Constructor_Data_Map_Internal_Node
{
if (v_4.Type == 9 && v_4.IntVal == 324739070 && v_4.UnsafePtr == nil) {
__t26 = (*Constructor_Data_Map_Internal_Node)(nil)
goto end_branch_26
} else {

}
}
{
if (v_4.Type == 9 && v_4.IntVal == 324739070 && v_4.UnsafePtr != nil) {
__t26 = &Constructor_Data_Map_Internal_Node{1, (*Constructor_Data_Map_Internal_Node)(v_4.UnsafePtr).V0, (*Constructor_Data_Map_Internal_Node)(v_4.UnsafePtr).V1, (*Constructor_Data_Map_Internal_Node)(v_4.UnsafePtr).V2, gopurs_runtime.Apply(f_2, (*Constructor_Data_Map_Internal_Node)(v_4.UnsafePtr).V3), gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.Apply(go__go_3_25_156, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v_4.UnsafePtr).V4)})), gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.Apply(go__go_3_25_156, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v_4.UnsafePtr).V5)}))}
goto end_branch_26
} else {

}
}
{
__t26 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_26:
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(__t26)}
})
return go__go_3_25_156
}))
}), gopurs_runtime.Func(func(dictApplicative_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_Map_Internal_traversableMap(), "traverse"), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](dictApplicative_1))}, gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return x_2
}))
}), gopurs_runtime.Func(func(dictApplicative_1 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Apply0_2_27 -> *Constructor_Control_Apply_Apply
Apply0_2_27 := gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_1, "Apply0"), gopurs_runtime.Value{}))
_ = Apply0_2_27
// TAST (Let): Functor0_3_28 -> *Constructor_Data_Functor_Functor
Functor0_3_28 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_1, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_3_28
return gopurs_runtime.Func(func(f_4 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_5_29_157 gopurs_runtime.Value
_ = go__go_5_29_157
go__go_5_29_157 = gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
var __t33 gopurs_runtime.Value
{
if (v_6.Type == 9 && v_6.IntVal == 324739070 && v_6.UnsafePtr == nil) {
__t33 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_1, "pure"), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(nil))})
goto end_branch_33
} else {

}
}
{
if (v_6.Type == 9 && v_6.IntVal == 324739070 && v_6.UnsafePtr != nil) {
// TAST (Let): __local_var_7_30 -> gopurs_runtime.Value
var __local_var_7_30 gopurs_runtime.Value = gopurs_runtime.Int((*Constructor_Data_Map_Internal_Node)(v_6.UnsafePtr).V0)
// TAST (Let): __local_var_8_31 -> gopurs_runtime.Value
__local_var_8_31 := (*Constructor_Data_Map_Internal_Node)(v_6.UnsafePtr).V2
_ = __local_var_8_31
// TAST (Let): __local_var_9_32 -> gopurs_runtime.Value
var __local_var_9_32 gopurs_runtime.Value = gopurs_runtime.Int((*Constructor_Data_Map_Internal_Node)(v_6.UnsafePtr).V1)
__t33 = gopurs_runtime.Apply2(gopurs_runtime.Box(Apply0_2_27.V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Apply0_2_27.V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_3_28.V0), gopurs_runtime.Func(func(l_prime_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_prime_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(r_prime_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(&Constructor_Data_Map_Internal_Node{1, __local_var_7_30.IntVal, __local_var_9_32.IntVal, __local_var_8_31, v_prime_11, gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](l_prime_10), gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](r_prime_12)})}
})
})
}), gopurs_runtime.Apply(go__go_5_29_157, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v_6.UnsafePtr).V4)})), gopurs_runtime.Apply(f_4, (*Constructor_Data_Map_Internal_Node)(v_6.UnsafePtr).V3)), gopurs_runtime.Apply(go__go_5_29_157, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v_6.UnsafePtr).V5)}))
goto end_branch_33
} else {

}
}
{
__t33 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_33:
return __t33
})
return go__go_5_29_157
})
}))
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Apply0_1_34 -> *Constructor_Control_Apply_Apply
Apply0_1_34 := gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}))
_ = Apply0_1_34
// TAST (Let): Functor0_2_35 -> *Constructor_Data_Functor_Functor
Functor0_2_35 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_2_35
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_4_36_158 gopurs_runtime.Value
_ = go__go_4_36_158
go__go_4_36_158 = gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
var __t40 gopurs_runtime.Value
{
if (v_5.Type == 9 && v_5.IntVal == 324739070 && v_5.UnsafePtr == nil) {
__t40 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "pure"), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(nil))})
goto end_branch_40
} else {

}
}
{
if (v_5.Type == 9 && v_5.IntVal == 324739070 && v_5.UnsafePtr != nil) {
// TAST (Let): __local_var_6_37 -> gopurs_runtime.Value
var __local_var_6_37 gopurs_runtime.Value = gopurs_runtime.Int((*Constructor_Data_Map_Internal_Node)(v_5.UnsafePtr).V0)
// TAST (Let): __local_var_7_38 -> gopurs_runtime.Value
__local_var_7_38 := (*Constructor_Data_Map_Internal_Node)(v_5.UnsafePtr).V2
_ = __local_var_7_38
// TAST (Let): __local_var_8_39 -> gopurs_runtime.Value
var __local_var_8_39 gopurs_runtime.Value = gopurs_runtime.Int((*Constructor_Data_Map_Internal_Node)(v_5.UnsafePtr).V1)
__t40 = gopurs_runtime.Apply2(gopurs_runtime.Box(Apply0_1_34.V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Apply0_1_34.V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_2_35.V0), gopurs_runtime.Func(func(l_prime_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_prime_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(r_prime_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(&Constructor_Data_Map_Internal_Node{1, __local_var_6_37.IntVal, __local_var_8_39.IntVal, __local_var_7_38, v_prime_10, gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](l_prime_9), gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](r_prime_11)})}
})
})
}), gopurs_runtime.Apply(go__go_4_36_158, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v_5.UnsafePtr).V4)})), gopurs_runtime.Apply2(f_3, __local_var_7_38, (*Constructor_Data_Map_Internal_Node)(v_5.UnsafePtr).V3)), gopurs_runtime.Apply(go__go_4_36_158, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v_5.UnsafePtr).V5)}))
goto end_branch_40
} else {

}
}
{
__t40 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_40:
return __t40
})
return go__go_4_36_158
})
}))
	})
	return cache_Data_Map_Internal_traversableWithIndexMap__3269014446
}

var cache_Data_Map_Internal_unionWith__2507192643 gopurs_runtime.Value
var once_Data_Map_Internal_unionWith__2507192643 sync.Once
func Get_Data_Map_Internal_unionWith__2507192643() gopurs_runtime.Value {
	once_Data_Map_Internal_unionWith__2507192643.Do(func() {
		cache_Data_Map_Internal_unionWith__2507192643 = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Map_Internal_unionWith__2507192643(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord](dictOrd_0_box))
})
	})
	return cache_Data_Map_Internal_unionWith__2507192643
}

var cache_Data_Map_Internal_unionWith__952079555 gopurs_runtime.Value
var once_Data_Map_Internal_unionWith__952079555 sync.Once
func Get_Data_Map_Internal_unionWith__952079555() gopurs_runtime.Value {
	once_Data_Map_Internal_unionWith__952079555.Do(func() {
		cache_Data_Map_Internal_unionWith__952079555 = gopurs_runtime.Func3(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value, __eta2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Map_Internal_unionWith__952079555(__eta0_0_box, __eta1_1_box, __eta2_2_box)
})
	})
	return cache_Data_Map_Internal_unionWith__952079555
}

var cache_Data_Map_Internal_unsafeBalancedNode__1259503046 gopurs_runtime.Value
var once_Data_Map_Internal_unsafeBalancedNode__1259503046 sync.Once
func Get_Data_Map_Internal_unsafeBalancedNode__1259503046() gopurs_runtime.Value {
	once_Data_Map_Internal_unsafeBalancedNode__1259503046.Do(func() {
		cache_Data_Map_Internal_unsafeBalancedNode__1259503046 = gopurs_runtime.Func4(func(__local_var_0_box gopurs_runtime.Value, __local_var_1_box gopurs_runtime.Value, __local_var_2_box gopurs_runtime.Value, __local_var_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Map_Internal_unsafeBalancedNode__1259503046(__local_var_0_box, __local_var_1_box, __local_var_2_box, __local_var_3_box)
})
	})
	return cache_Data_Map_Internal_unsafeBalancedNode__1259503046
}

var cache_Data_Map_Internal_unsafeBalancedNode__1305301638 gopurs_runtime.Value
var once_Data_Map_Internal_unsafeBalancedNode__1305301638 sync.Once
func Get_Data_Map_Internal_unsafeBalancedNode__1305301638() gopurs_runtime.Value {
	once_Data_Map_Internal_unsafeBalancedNode__1305301638.Do(func() {
		cache_Data_Map_Internal_unsafeBalancedNode__1305301638 = gopurs_runtime.Func4(func(__local_var_0_box gopurs_runtime.Value, __local_var_1_box gopurs_runtime.Value, __local_var_2_box gopurs_runtime.Value, __local_var_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Map_Internal_unsafeBalancedNode__1305301638(__local_var_0_box, __local_var_1_box, __local_var_2_box, __local_var_3_box)
})
	})
	return cache_Data_Map_Internal_unsafeBalancedNode__1305301638
}

var cache_Data_Map_Internal_unsafeBalancedNode__954819782 gopurs_runtime.Value
var once_Data_Map_Internal_unsafeBalancedNode__954819782 sync.Once
func Get_Data_Map_Internal_unsafeBalancedNode__954819782() gopurs_runtime.Value {
	once_Data_Map_Internal_unsafeBalancedNode__954819782.Do(func() {
		cache_Data_Map_Internal_unsafeBalancedNode__954819782 = gopurs_runtime.Func4(func(__local_var_0_box gopurs_runtime.Value, __local_var_1_box gopurs_runtime.Value, __local_var_2_box gopurs_runtime.Value, __local_var_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Map_Internal_unsafeBalancedNode__954819782(__local_var_0_box, __local_var_1_box, __local_var_2_box, __local_var_3_box)
})
	})
	return cache_Data_Map_Internal_unsafeBalancedNode__954819782
}

var cache_Data_Map_Internal_unsafeBalancedNode__1776657286 gopurs_runtime.Value
var once_Data_Map_Internal_unsafeBalancedNode__1776657286 sync.Once
func Get_Data_Map_Internal_unsafeBalancedNode__1776657286() gopurs_runtime.Value {
	once_Data_Map_Internal_unsafeBalancedNode__1776657286.Do(func() {
		cache_Data_Map_Internal_unsafeBalancedNode__1776657286 = gopurs_runtime.Func4(func(__local_var_0_box gopurs_runtime.Value, __local_var_1_box gopurs_runtime.Value, __local_var_2_box gopurs_runtime.Value, __local_var_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Map_Internal_unsafeBalancedNode__1776657286(__local_var_0_box, __local_var_1_box, __local_var_2_box, __local_var_3_box)
})
	})
	return cache_Data_Map_Internal_unsafeBalancedNode__1776657286
}

var cache_Data_Map_Internal_unsafeBalancedNode__1902536198 gopurs_runtime.Value
var once_Data_Map_Internal_unsafeBalancedNode__1902536198 sync.Once
func Get_Data_Map_Internal_unsafeBalancedNode__1902536198() gopurs_runtime.Value {
	once_Data_Map_Internal_unsafeBalancedNode__1902536198.Do(func() {
		cache_Data_Map_Internal_unsafeBalancedNode__1902536198 = gopurs_runtime.Func4(func(__local_var_0_box gopurs_runtime.Value, __local_var_1_box gopurs_runtime.Value, __local_var_2_box gopurs_runtime.Value, __local_var_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Map_Internal_unsafeBalancedNode__1902536198(__local_var_0_box, __local_var_1_box, __local_var_2_box, __local_var_3_box)
})
	})
	return cache_Data_Map_Internal_unsafeBalancedNode__1902536198
}

var cache_Data_Map_Internal_unsafeDifference__4097927905 gopurs_runtime.Value
var once_Data_Map_Internal_unsafeDifference__4097927905 sync.Once
func Get_Data_Map_Internal_unsafeDifference__4097927905() gopurs_runtime.Value {
	once_Data_Map_Internal_unsafeDifference__4097927905.Do(func() {
		cache_Data_Map_Internal_unsafeDifference__4097927905 = gopurs_runtime.Func3(func(__local_var_0_box gopurs_runtime.Value, __local_var_1_box gopurs_runtime.Value, __local_var_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Map_Internal_unsafeDifference__4097927905(__local_var_0_box, __local_var_1_box, __local_var_2_box)
})
	})
	return cache_Data_Map_Internal_unsafeDifference__4097927905
}

var cache_Data_Map_Internal_unsafeIntersectionWith__4109280494 gopurs_runtime.Value
var once_Data_Map_Internal_unsafeIntersectionWith__4109280494 sync.Once
func Get_Data_Map_Internal_unsafeIntersectionWith__4109280494() gopurs_runtime.Value {
	once_Data_Map_Internal_unsafeIntersectionWith__4109280494.Do(func() {
		cache_Data_Map_Internal_unsafeIntersectionWith__4109280494 = gopurs_runtime.Func4(func(__local_var_0_box gopurs_runtime.Value, __local_var_1_box gopurs_runtime.Value, __local_var_2_box gopurs_runtime.Value, __local_var_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Map_Internal_unsafeIntersectionWith__4109280494(__local_var_0_box, __local_var_1_box, __local_var_2_box, __local_var_3_box)
})
	})
	return cache_Data_Map_Internal_unsafeIntersectionWith__4109280494
}

var cache_Data_Map_Internal_unsafeIntersectionWith__2517966 gopurs_runtime.Value
var once_Data_Map_Internal_unsafeIntersectionWith__2517966 sync.Once
func Get_Data_Map_Internal_unsafeIntersectionWith__2517966() gopurs_runtime.Value {
	once_Data_Map_Internal_unsafeIntersectionWith__2517966.Do(func() {
		cache_Data_Map_Internal_unsafeIntersectionWith__2517966 = gopurs_runtime.Func4(func(__local_var_0_box gopurs_runtime.Value, __local_var_1_box gopurs_runtime.Value, __local_var_2_box gopurs_runtime.Value, __local_var_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Map_Internal_unsafeIntersectionWith__2517966(__local_var_0_box, __local_var_1_box, __local_var_2_box, __local_var_3_box)
})
	})
	return cache_Data_Map_Internal_unsafeIntersectionWith__2517966
}

var cache_Data_Map_Internal_unsafeJoinNodes__2531831408 gopurs_runtime.Value
var once_Data_Map_Internal_unsafeJoinNodes__2531831408 sync.Once
func Get_Data_Map_Internal_unsafeJoinNodes__2531831408() gopurs_runtime.Value {
	once_Data_Map_Internal_unsafeJoinNodes__2531831408.Do(func() {
		cache_Data_Map_Internal_unsafeJoinNodes__2531831408 = gopurs_runtime.Func2(func(__local_var_0_box gopurs_runtime.Value, __local_var_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Map_Internal_unsafeJoinNodes__2531831408(__local_var_0_box, __local_var_1_box)
})
	})
	return cache_Data_Map_Internal_unsafeJoinNodes__2531831408
}

var cache_Data_Map_Internal_unsafeJoinNodes__3967876672 gopurs_runtime.Value
var once_Data_Map_Internal_unsafeJoinNodes__3967876672 sync.Once
func Get_Data_Map_Internal_unsafeJoinNodes__3967876672() gopurs_runtime.Value {
	once_Data_Map_Internal_unsafeJoinNodes__3967876672.Do(func() {
		cache_Data_Map_Internal_unsafeJoinNodes__3967876672 = gopurs_runtime.Func2(func(__local_var_0_box gopurs_runtime.Value, __local_var_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Map_Internal_unsafeJoinNodes__3967876672(__local_var_0_box, __local_var_1_box)
})
	})
	return cache_Data_Map_Internal_unsafeJoinNodes__3967876672
}

var cache_Data_Map_Internal_unsafeNode__1259503046 gopurs_runtime.Value
var once_Data_Map_Internal_unsafeNode__1259503046 sync.Once
func Get_Data_Map_Internal_unsafeNode__1259503046() gopurs_runtime.Value {
	once_Data_Map_Internal_unsafeNode__1259503046.Do(func() {
		cache_Data_Map_Internal_unsafeNode__1259503046 = gopurs_runtime.Func4(func(__local_var_0_box gopurs_runtime.Value, __local_var_1_box gopurs_runtime.Value, __local_var_2_box gopurs_runtime.Value, __local_var_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Map_Internal_unsafeNode__1259503046(__local_var_0_box, __local_var_1_box, __local_var_2_box, __local_var_3_box)
})
	})
	return cache_Data_Map_Internal_unsafeNode__1259503046
}

var cache_Data_Map_Internal_unsafeNode__1305301638 gopurs_runtime.Value
var once_Data_Map_Internal_unsafeNode__1305301638 sync.Once
func Get_Data_Map_Internal_unsafeNode__1305301638() gopurs_runtime.Value {
	once_Data_Map_Internal_unsafeNode__1305301638.Do(func() {
		cache_Data_Map_Internal_unsafeNode__1305301638 = gopurs_runtime.Func4(func(__local_var_0_box gopurs_runtime.Value, __local_var_1_box gopurs_runtime.Value, __local_var_2_box gopurs_runtime.Value, __local_var_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Map_Internal_unsafeNode__1305301638(__local_var_0_box, __local_var_1_box, __local_var_2_box, __local_var_3_box)
})
	})
	return cache_Data_Map_Internal_unsafeNode__1305301638
}

var cache_Data_Map_Internal_unsafeNode__954819782 gopurs_runtime.Value
var once_Data_Map_Internal_unsafeNode__954819782 sync.Once
func Get_Data_Map_Internal_unsafeNode__954819782() gopurs_runtime.Value {
	once_Data_Map_Internal_unsafeNode__954819782.Do(func() {
		cache_Data_Map_Internal_unsafeNode__954819782 = gopurs_runtime.Func4(func(__local_var_0_box gopurs_runtime.Value, __local_var_1_box gopurs_runtime.Value, __local_var_2_box gopurs_runtime.Value, __local_var_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Map_Internal_unsafeNode__954819782(__local_var_0_box, __local_var_1_box, __local_var_2_box, __local_var_3_box)
})
	})
	return cache_Data_Map_Internal_unsafeNode__954819782
}

var cache_Data_Map_Internal_unsafeNode__1776657286 gopurs_runtime.Value
var once_Data_Map_Internal_unsafeNode__1776657286 sync.Once
func Get_Data_Map_Internal_unsafeNode__1776657286() gopurs_runtime.Value {
	once_Data_Map_Internal_unsafeNode__1776657286.Do(func() {
		cache_Data_Map_Internal_unsafeNode__1776657286 = gopurs_runtime.Func4(func(__local_var_0_box gopurs_runtime.Value, __local_var_1_box gopurs_runtime.Value, __local_var_2_box gopurs_runtime.Value, __local_var_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Map_Internal_unsafeNode__1776657286(__local_var_0_box, __local_var_1_box, __local_var_2_box, __local_var_3_box)
})
	})
	return cache_Data_Map_Internal_unsafeNode__1776657286
}

var cache_Data_Map_Internal_unsafeNode__1902536198 gopurs_runtime.Value
var once_Data_Map_Internal_unsafeNode__1902536198 sync.Once
func Get_Data_Map_Internal_unsafeNode__1902536198() gopurs_runtime.Value {
	once_Data_Map_Internal_unsafeNode__1902536198.Do(func() {
		cache_Data_Map_Internal_unsafeNode__1902536198 = gopurs_runtime.Func4(func(__local_var_0_box gopurs_runtime.Value, __local_var_1_box gopurs_runtime.Value, __local_var_2_box gopurs_runtime.Value, __local_var_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Map_Internal_unsafeNode__1902536198(__local_var_0_box, __local_var_1_box, __local_var_2_box, __local_var_3_box)
})
	})
	return cache_Data_Map_Internal_unsafeNode__1902536198
}

var cache_Data_Map_Internal_unsafeSplit__1094566431 gopurs_runtime.Value
var once_Data_Map_Internal_unsafeSplit__1094566431 sync.Once
func Get_Data_Map_Internal_unsafeSplit__1094566431() gopurs_runtime.Value {
	once_Data_Map_Internal_unsafeSplit__1094566431.Do(func() {
		cache_Data_Map_Internal_unsafeSplit__1094566431 = gopurs_runtime.Func3(func(__local_var_0_box gopurs_runtime.Value, __local_var_1_box gopurs_runtime.Value, __local_var_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Map_Internal_unsafeSplit__1094566431(__local_var_0_box, __local_var_1_box, __local_var_2_box)
})
	})
	return cache_Data_Map_Internal_unsafeSplit__1094566431
}

var cache_Data_Map_Internal_unsafeSplit__4154869695 gopurs_runtime.Value
var once_Data_Map_Internal_unsafeSplit__4154869695 sync.Once
func Get_Data_Map_Internal_unsafeSplit__4154869695() gopurs_runtime.Value {
	once_Data_Map_Internal_unsafeSplit__4154869695.Do(func() {
		cache_Data_Map_Internal_unsafeSplit__4154869695 = gopurs_runtime.Func3(func(__local_var_0_box gopurs_runtime.Value, __local_var_1_box gopurs_runtime.Value, __local_var_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Map_Internal_unsafeSplit__4154869695(__local_var_0_box, __local_var_1_box, __local_var_2_box)
})
	})
	return cache_Data_Map_Internal_unsafeSplit__4154869695
}

var cache_Data_Map_Internal_unsafeSplit__1308258847 gopurs_runtime.Value
var once_Data_Map_Internal_unsafeSplit__1308258847 sync.Once
func Get_Data_Map_Internal_unsafeSplit__1308258847() gopurs_runtime.Value {
	once_Data_Map_Internal_unsafeSplit__1308258847.Do(func() {
		cache_Data_Map_Internal_unsafeSplit__1308258847 = gopurs_runtime.Func3(func(__local_var_0_box gopurs_runtime.Value, __local_var_1_box gopurs_runtime.Value, __local_var_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Map_Internal_unsafeSplit__1308258847(__local_var_0_box, __local_var_1_box, __local_var_2_box)
})
	})
	return cache_Data_Map_Internal_unsafeSplit__1308258847
}

var cache_Data_Map_Internal_unsafeSplit__1115245464 gopurs_runtime.Value
var once_Data_Map_Internal_unsafeSplit__1115245464 sync.Once
func Get_Data_Map_Internal_unsafeSplit__1115245464() gopurs_runtime.Value {
	once_Data_Map_Internal_unsafeSplit__1115245464.Do(func() {
		cache_Data_Map_Internal_unsafeSplit__1115245464 = gopurs_runtime.Func3(func(__local_var_0_box gopurs_runtime.Value, __local_var_1_box gopurs_runtime.Value, __local_var_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Map_Internal_unsafeSplit__1115245464(__local_var_0_box, __local_var_1_box, __local_var_2_box)
})
	})
	return cache_Data_Map_Internal_unsafeSplit__1115245464
}

var cache_Data_Map_Internal_unsafeSplitLast__1494186946 gopurs_runtime.Value
var once_Data_Map_Internal_unsafeSplitLast__1494186946 sync.Once
func Get_Data_Map_Internal_unsafeSplitLast__1494186946() gopurs_runtime.Value {
	once_Data_Map_Internal_unsafeSplitLast__1494186946.Do(func() {
		cache_Data_Map_Internal_unsafeSplitLast__1494186946 = gopurs_runtime.Func4(func(__local_var_0_box gopurs_runtime.Value, __local_var_1_box gopurs_runtime.Value, __local_var_2_box gopurs_runtime.Value, __local_var_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Map_Internal_unsafeSplitLast__1494186946(__local_var_0_box, __local_var_1_box, __local_var_2_box, __local_var_3_box)
})
	})
	return cache_Data_Map_Internal_unsafeSplitLast__1494186946
}

var cache_Data_Map_Internal_unsafeSplitLast__224676098 gopurs_runtime.Value
var once_Data_Map_Internal_unsafeSplitLast__224676098 sync.Once
func Get_Data_Map_Internal_unsafeSplitLast__224676098() gopurs_runtime.Value {
	once_Data_Map_Internal_unsafeSplitLast__224676098.Do(func() {
		cache_Data_Map_Internal_unsafeSplitLast__224676098 = gopurs_runtime.Func4(func(__local_var_0_box gopurs_runtime.Value, __local_var_1_box gopurs_runtime.Value, __local_var_2_box gopurs_runtime.Value, __local_var_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Map_Internal_unsafeSplitLast__224676098(__local_var_0_box, __local_var_1_box, __local_var_2_box, __local_var_3_box)
})
	})
	return cache_Data_Map_Internal_unsafeSplitLast__224676098
}

var cache_Data_Map_Internal_unsafeUnionWith__4109280494 gopurs_runtime.Value
var once_Data_Map_Internal_unsafeUnionWith__4109280494 sync.Once
func Get_Data_Map_Internal_unsafeUnionWith__4109280494() gopurs_runtime.Value {
	once_Data_Map_Internal_unsafeUnionWith__4109280494.Do(func() {
		cache_Data_Map_Internal_unsafeUnionWith__4109280494 = gopurs_runtime.Func4(func(__local_var_0_box gopurs_runtime.Value, __local_var_1_box gopurs_runtime.Value, __local_var_2_box gopurs_runtime.Value, __local_var_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Map_Internal_unsafeUnionWith__4109280494(__local_var_0_box, __local_var_1_box, __local_var_2_box, __local_var_3_box)
})
	})
	return cache_Data_Map_Internal_unsafeUnionWith__4109280494
}

var cache_Data_Map_Internal_unsafeUnionWith__3421363785 gopurs_runtime.Value
var once_Data_Map_Internal_unsafeUnionWith__3421363785 sync.Once
func Get_Data_Map_Internal_unsafeUnionWith__3421363785() gopurs_runtime.Value {
	once_Data_Map_Internal_unsafeUnionWith__3421363785.Do(func() {
		cache_Data_Map_Internal_unsafeUnionWith__3421363785 = gopurs_runtime.Func4(func(__local_var_0_box gopurs_runtime.Value, __local_var_1_box gopurs_runtime.Value, __local_var_2_box gopurs_runtime.Value, __local_var_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Map_Internal_unsafeUnionWith__3421363785(__local_var_0_box, __local_var_1_box, __local_var_2_box, __local_var_3_box)
})
	})
	return cache_Data_Map_Internal_unsafeUnionWith__3421363785
}

type Constructor_Data_Map_Internal_Leaf struct {
	Rc uint32
}


type Constructor_Data_Map_Internal_Node struct {
	Rc uint32
	V0 int64
	V1 int64
	V2 gopurs_runtime.Value
	V3 gopurs_runtime.Value
	V4 *Constructor_Data_Map_Internal_Node
	V5 *Constructor_Data_Map_Internal_Node
}


type Constructor_Data_Map_Internal_IterLeaf struct {
	Rc uint32
}


type Constructor_Data_Map_Internal_IterEmit struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
	V2 gopurs_runtime.Value
}


type Constructor_Data_Map_Internal_IterNode struct {
	Rc uint32
	V0 *Constructor_Data_Map_Internal_Node
	V1 gopurs_runtime.Value
}


type Constructor_Data_Map_Internal_IterDone struct {
	Rc uint32
}


type Constructor_Data_Map_Internal_IterNext struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
	V2 gopurs_runtime.Value
}


type Constructor_Data_Map_Internal_Split struct {
	Rc uint32
	V0 *Constructor_Data_Maybe_Just
	V1 *Constructor_Data_Map_Internal_Node
	V2 *Constructor_Data_Map_Internal_Node
}


type Constructor_Data_Map_Internal_SplitLast struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
	V2 *Constructor_Data_Map_Internal_Node
}


func Call_Data_Map_Internal_identity(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_Data_Map_Internal_identity1(x_0_loop *Constructor_Data_Maybe_Just) *Constructor_Data_Maybe_Just {
var x_0 *Constructor_Data_Maybe_Just = x_0_loop
_ = x_0
return x_0
}

func Call_Data_Map_Internal_identity2(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_Data_Map_Internal_unsafeNode(__local_var_0_loop gopurs_runtime.Value, __local_var_1_loop gopurs_runtime.Value, __local_var_2_loop gopurs_runtime.Value, __local_var_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __local_var_0 gopurs_runtime.Value = __local_var_0_loop
_ = __local_var_0
var __local_var_1 gopurs_runtime.Value = __local_var_1_loop
_ = __local_var_1
var __local_var_2 gopurs_runtime.Value = __local_var_2_loop
_ = __local_var_2
var __local_var_3 gopurs_runtime.Value = __local_var_3_loop
_ = __local_var_3
var __t4 *Constructor_Data_Map_Internal_Node
{
if (__local_var_2.Type == 9 && __local_var_2.IntVal == 324739070 && __local_var_2.UnsafePtr == nil) {
var __t0 *Constructor_Data_Map_Internal_Node
{
if (__local_var_3.Type == 9 && __local_var_3.IntVal == 324739070 && __local_var_3.UnsafePtr == nil) {
__t0 = &Constructor_Data_Map_Internal_Node{1, 1, 1, __local_var_0, __local_var_1, gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](__local_var_2), gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](__local_var_3)}
goto end_branch_0
} else {

}
}
{
if (__local_var_3.Type == 9 && __local_var_3.IntVal == 324739070 && __local_var_3.UnsafePtr != nil) {
__t0 = &Constructor_Data_Map_Internal_Node{1, (1) + ((*Constructor_Data_Map_Internal_Node)(__local_var_3.UnsafePtr).V0), (1) + ((*Constructor_Data_Map_Internal_Node)(__local_var_3.UnsafePtr).V1), __local_var_0, __local_var_1, gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](__local_var_2), gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](__local_var_3)}
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_0:
__t4 = __t0
goto end_branch_4
} else {

}
}
{
if (__local_var_2.Type == 9 && __local_var_2.IntVal == 324739070 && __local_var_2.UnsafePtr != nil) {
var __t3 *Constructor_Data_Map_Internal_Node
{
if (__local_var_3.Type == 9 && __local_var_3.IntVal == 324739070 && __local_var_3.UnsafePtr == nil) {
__t3 = &Constructor_Data_Map_Internal_Node{1, (1) + ((*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V0), (1) + ((*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V1), __local_var_0, __local_var_1, gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](__local_var_2), gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](__local_var_3)}
goto end_branch_3
} else {

}
}
{
if (__local_var_3.Type == 9 && __local_var_3.IntVal == 324739070 && __local_var_3.UnsafePtr != nil) {
var __t2 int64
{
var __t1 bool
{
if ((*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V0) > ((*Constructor_Data_Map_Internal_Node)(__local_var_3.UnsafePtr).V0) {
__t1 = true
goto end_branch_1
} else {

}
}
{
__t1 = false
}
end_branch_1:
if __t1 {
__t2 = (1) + ((*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V0)
goto end_branch_2
} else {

}
}
{
__t2 = (1) + ((*Constructor_Data_Map_Internal_Node)(__local_var_3.UnsafePtr).V0)
}
end_branch_2:
__t3 = &Constructor_Data_Map_Internal_Node{1, __t2, ((1) + ((*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V1)) + ((*Constructor_Data_Map_Internal_Node)(__local_var_3.UnsafePtr).V1), __local_var_0, __local_var_1, gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](__local_var_2), gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](__local_var_3)}
goto end_branch_3
} else {

}
}
{
__t3 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_3:
__t4 = __t3
goto end_branch_4
} else {

}
}
{
__t4 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_4:
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(__t4)}
}

func Call_Data_Map_Internal_toMapIter(a_0_loop *Constructor_Data_Map_Internal_Node) gopurs_runtime.Value {
var a_0 *Constructor_Data_Map_Internal_Node = a_0_loop
_ = a_0
return gopurs_runtime.Value{Type: 9, IntVal: 2861335956, UnsafePtr: unsafe.Pointer(&Constructor_Data_Map_Internal_IterNode{1, a_0, gopurs_runtime.Value{Type: 9, IntVal: 2509360378, UnsafePtr: unsafe.Pointer(nil)}})}
}

func Call_Data_Map_Internal_stepWith(f_0_loop gopurs_runtime.Value, next_1_loop gopurs_runtime.Value, done_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var next_1 gopurs_runtime.Value = next_1_loop
_ = next_1
var done_2 gopurs_runtime.Value = done_2_loop
_ = done_2
var go__go_3_0_0 gopurs_runtime.Value
go__go_3_0_0 = gopurs_runtime.Func(func(v_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_4_loop gopurs_runtime.Value = v_4_loop_val
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
__t1 = gopurs_runtime.UncurriedApp3(next_1, (*Constructor_Data_Map_Internal_IterEmit)(v_4.UnsafePtr).V0, (*Constructor_Data_Map_Internal_IterEmit)(v_4.UnsafePtr).V1, (*Constructor_Data_Map_Internal_IterEmit)(v_4.UnsafePtr).V2)
goto end_branch_1
} else {

}
}
{
if (v_4.Type == 9 && v_4.IntVal == 2861335956) {
v_4_loop = gopurs_runtime.Apply2(f_0, (*Constructor_Data_Map_Internal_IterNode)(v_4.UnsafePtr).V1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_IterNode)(v_4.UnsafePtr).V0)})
continue go__go_3_0_0
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
return go__go_3_0_0
}

func Call_Data_Map_Internal_size(v_0_loop *Constructor_Data_Map_Internal_Node) int64 {
var v_0 *Constructor_Data_Map_Internal_Node = v_0_loop
_ = v_0
var __t0 int64
{
if (v_0 == nil) {
__t0 = 0
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
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal
}
end_branch_0:
return __t0
}

func Call_Data_Map_Internal_singleton(k_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) *Constructor_Data_Map_Internal_Node {
var k_0 gopurs_runtime.Value = k_0_loop
_ = k_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return &Constructor_Data_Map_Internal_Node{1, 1, 1, k_0, v_1, (*Constructor_Data_Map_Internal_Node)(nil), (*Constructor_Data_Map_Internal_Node)(nil)}
}

func Call_Data_Map_Internal_unsafeBalancedNode(__local_var_0_loop gopurs_runtime.Value, __local_var_1_loop gopurs_runtime.Value, __local_var_2_loop gopurs_runtime.Value, __local_var_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __local_var_0 gopurs_runtime.Value = __local_var_0_loop
_ = __local_var_0
var __local_var_1 gopurs_runtime.Value = __local_var_1_loop
_ = __local_var_1
var __local_var_2 gopurs_runtime.Value = __local_var_2_loop
_ = __local_var_2
var __local_var_3 gopurs_runtime.Value = __local_var_3_loop
_ = __local_var_3
var __t37 *Constructor_Data_Map_Internal_Node
{
if (__local_var_2.Type == 9 && __local_var_2.IntVal == 324739070 && __local_var_2.UnsafePtr == nil) {
var __t9 *Constructor_Data_Map_Internal_Node
{
if (__local_var_3.Type == 9 && __local_var_3.IntVal == 324739070 && __local_var_3.UnsafePtr == nil) {
__t9 = &Constructor_Data_Map_Internal_Node{1, 1, 1, __local_var_0, __local_var_1, (*Constructor_Data_Map_Internal_Node)(nil), (*Constructor_Data_Map_Internal_Node)(nil)}
goto end_branch_9
} else {

}
}
{
var __t_and_1 bool = false
if (__local_var_3.Type == 9 && __local_var_3.IntVal == 324739070 && __local_var_3.UnsafePtr != nil) {

var __t0 bool
{
if ((*Constructor_Data_Map_Internal_Node)(__local_var_3.UnsafePtr).V0) > (1) {
__t0 = true
goto end_branch_0
} else {

}
}
{
__t0 = false
}
end_branch_0:
__t_and_1 = __t0
}
if __t_and_1 {
var __t8 *Constructor_Data_Map_Internal_Node
{
var __t_tag_2 *Constructor_Data_Map_Internal_Node = (*Constructor_Data_Map_Internal_Node)(__local_var_3.UnsafePtr).V4
var __t_and_7 bool = false
if (__t_tag_2 != nil) {

var __t6 bool
{
var __t5 int64
{
var __t_tag_3 *Constructor_Data_Map_Internal_Node = (*Constructor_Data_Map_Internal_Node)(__local_var_3.UnsafePtr).V5
if (__t_tag_3 == nil) {
__t5 = 0
goto end_branch_5
} else {

}
}
{
var __t_tag_4 *Constructor_Data_Map_Internal_Node = (*Constructor_Data_Map_Internal_Node)(__local_var_3.UnsafePtr).V5
if (__t_tag_4 != nil) {
__t5 = ((*Constructor_Data_Map_Internal_Node)(__local_var_3.UnsafePtr).V5).V0
goto end_branch_5
} else {

}
}
{
__t5 = func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal
}
end_branch_5:
if (((*Constructor_Data_Map_Internal_Node)(__local_var_3.UnsafePtr).V4).V0) > (__t5) {
__t6 = true
goto end_branch_6
} else {

}
}
{
__t6 = false
}
end_branch_6:
__t_and_7 = __t6
}
if __t_and_7 {
__t8 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeNode(), ((*Constructor_Data_Map_Internal_Node)(__local_var_3.UnsafePtr).V4).V2, ((*Constructor_Data_Map_Internal_Node)(__local_var_3.UnsafePtr).V4).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeNode(), __local_var_0, __local_var_1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](__local_var_2))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(((*Constructor_Data_Map_Internal_Node)(__local_var_3.UnsafePtr).V4).V4)})))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeNode(), (*Constructor_Data_Map_Internal_Node)(__local_var_3.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node)(__local_var_3.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(((*Constructor_Data_Map_Internal_Node)(__local_var_3.UnsafePtr).V4).V5)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_3.UnsafePtr).V5)})))}))
goto end_branch_8
} else {

}
}
{
__t8 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeNode(), (*Constructor_Data_Map_Internal_Node)(__local_var_3.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node)(__local_var_3.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeNode(), __local_var_0, __local_var_1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](__local_var_2))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_3.UnsafePtr).V4)})))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_3.UnsafePtr).V5)}))
}
end_branch_8:
__t9 = __t8
goto end_branch_9
} else {

}
}
{
__t9 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeNode(), __local_var_0, __local_var_1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](__local_var_2))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](__local_var_3))}))
}
end_branch_9:
__t37 = __t9
goto end_branch_37
} else {

}
}
{
if (__local_var_2.Type == 9 && __local_var_2.IntVal == 324739070 && __local_var_2.UnsafePtr != nil) {
var __t36 *Constructor_Data_Map_Internal_Node
{
if (__local_var_3.Type == 9 && __local_var_3.IntVal == 324739070 && __local_var_3.UnsafePtr != nil) {
var __t26 *Constructor_Data_Map_Internal_Node
{
var __t10 bool
{
if ((*Constructor_Data_Map_Internal_Node)(__local_var_3.UnsafePtr).V0) > (((*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V0) + (1)) {
__t10 = true
goto end_branch_10
} else {

}
}
{
__t10 = false
}
end_branch_10:
if __t10 {
var __t17 *Constructor_Data_Map_Internal_Node
{
var __t_tag_11 *Constructor_Data_Map_Internal_Node = (*Constructor_Data_Map_Internal_Node)(__local_var_3.UnsafePtr).V4
var __t_and_16 bool = false
if (__t_tag_11 != nil) {

var __t15 bool
{
var __t14 int64
{
var __t_tag_12 *Constructor_Data_Map_Internal_Node = (*Constructor_Data_Map_Internal_Node)(__local_var_3.UnsafePtr).V5
if (__t_tag_12 == nil) {
__t14 = 0
goto end_branch_14
} else {

}
}
{
var __t_tag_13 *Constructor_Data_Map_Internal_Node = (*Constructor_Data_Map_Internal_Node)(__local_var_3.UnsafePtr).V5
if (__t_tag_13 != nil) {
__t14 = ((*Constructor_Data_Map_Internal_Node)(__local_var_3.UnsafePtr).V5).V0
goto end_branch_14
} else {

}
}
{
__t14 = func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal
}
end_branch_14:
if (((*Constructor_Data_Map_Internal_Node)(__local_var_3.UnsafePtr).V4).V0) > (__t14) {
__t15 = true
goto end_branch_15
} else {

}
}
{
__t15 = false
}
end_branch_15:
__t_and_16 = __t15
}
if __t_and_16 {
__t17 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeNode(), ((*Constructor_Data_Map_Internal_Node)(__local_var_3.UnsafePtr).V4).V2, ((*Constructor_Data_Map_Internal_Node)(__local_var_3.UnsafePtr).V4).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeNode(), __local_var_0, __local_var_1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](__local_var_2))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(((*Constructor_Data_Map_Internal_Node)(__local_var_3.UnsafePtr).V4).V4)})))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeNode(), (*Constructor_Data_Map_Internal_Node)(__local_var_3.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node)(__local_var_3.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(((*Constructor_Data_Map_Internal_Node)(__local_var_3.UnsafePtr).V4).V5)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_3.UnsafePtr).V5)})))}))
goto end_branch_17
} else {

}
}
{
__t17 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeNode(), (*Constructor_Data_Map_Internal_Node)(__local_var_3.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node)(__local_var_3.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeNode(), __local_var_0, __local_var_1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](__local_var_2))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_3.UnsafePtr).V4)})))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_3.UnsafePtr).V5)}))
}
end_branch_17:
__t26 = __t17
goto end_branch_26
} else {

}
}
{
var __t18 bool
{
if ((*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V0) > (((*Constructor_Data_Map_Internal_Node)(__local_var_3.UnsafePtr).V0) + (1)) {
__t18 = true
goto end_branch_18
} else {

}
}
{
__t18 = false
}
end_branch_18:
if __t18 {
var __t25 *Constructor_Data_Map_Internal_Node
{
var __t_tag_19 *Constructor_Data_Map_Internal_Node = (*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V5
var __t_and_24 bool = false
if (__t_tag_19 != nil) {

var __t23 bool
{
var __t22 int64
{
var __t_tag_20 *Constructor_Data_Map_Internal_Node = (*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V4
if (__t_tag_20 == nil) {
__t22 = 0
goto end_branch_22
} else {

}
}
{
var __t_tag_21 *Constructor_Data_Map_Internal_Node = (*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V4
if (__t_tag_21 != nil) {
__t22 = ((*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V4).V0
goto end_branch_22
} else {

}
}
{
__t22 = func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal
}
end_branch_22:
if (__t22) > (((*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V5).V0) {
__t23 = false
goto end_branch_23
} else {

}
}
{
__t23 = true
}
end_branch_23:
__t_and_24 = __t23
}
if __t_and_24 {
__t25 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeNode(), ((*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V5).V2, ((*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V5).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeNode(), (*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V4)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(((*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V5).V4)})))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeNode(), __local_var_0, __local_var_1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(((*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V5).V5)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](__local_var_3))})))}))
goto end_branch_25
} else {

}
}
{
__t25 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeNode(), (*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V4)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeNode(), __local_var_0, __local_var_1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V5)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](__local_var_3))})))}))
}
end_branch_25:
__t26 = __t25
goto end_branch_26
} else {

}
}
{
__t26 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeNode(), __local_var_0, __local_var_1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](__local_var_2))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](__local_var_3))}))
}
end_branch_26:
__t36 = __t26
goto end_branch_36
} else {

}
}
{
var __t_and_28 bool = false
if (__local_var_3.Type == 9 && __local_var_3.IntVal == 324739070 && __local_var_3.UnsafePtr == nil) {

var __t27 bool
{
if ((*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V0) > (1) {
__t27 = true
goto end_branch_27
} else {

}
}
{
__t27 = false
}
end_branch_27:
__t_and_28 = __t27
}
if __t_and_28 {
var __t35 *Constructor_Data_Map_Internal_Node
{
var __t_tag_29 *Constructor_Data_Map_Internal_Node = (*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V5
var __t_and_34 bool = false
if (__t_tag_29 != nil) {

var __t33 bool
{
var __t32 int64
{
var __t_tag_30 *Constructor_Data_Map_Internal_Node = (*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V4
if (__t_tag_30 == nil) {
__t32 = 0
goto end_branch_32
} else {

}
}
{
var __t_tag_31 *Constructor_Data_Map_Internal_Node = (*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V4
if (__t_tag_31 != nil) {
__t32 = ((*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V4).V0
goto end_branch_32
} else {

}
}
{
__t32 = func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal
}
end_branch_32:
if (__t32) > (((*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V5).V0) {
__t33 = false
goto end_branch_33
} else {

}
}
{
__t33 = true
}
end_branch_33:
__t_and_34 = __t33
}
if __t_and_34 {
__t35 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeNode(), ((*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V5).V2, ((*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V5).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeNode(), (*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V4)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(((*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V5).V4)})))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeNode(), __local_var_0, __local_var_1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(((*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V5).V5)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](__local_var_3))})))}))
goto end_branch_35
} else {

}
}
{
__t35 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeNode(), (*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V4)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeNode(), __local_var_0, __local_var_1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V5)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](__local_var_3))})))}))
}
end_branch_35:
__t36 = __t35
goto end_branch_36
} else {

}
}
{
__t36 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeNode(), __local_var_0, __local_var_1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](__local_var_2))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](__local_var_3))}))
}
end_branch_36:
__t37 = __t36
goto end_branch_37
} else {

}
}
{
__t37 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_37:
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(__t37)}
}

func Call_Data_Map_Internal_unsafeSplit(__local_var_0_loop gopurs_runtime.Value, __local_var_1_loop gopurs_runtime.Value, __local_var_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
unsafeSplit:
for {
if false { continue unsafeSplit }
var __local_var_0 gopurs_runtime.Value = __local_var_0_loop
_ = __local_var_0
var __local_var_1 gopurs_runtime.Value = __local_var_1_loop
_ = __local_var_1
var __local_var_2 gopurs_runtime.Value = __local_var_2_loop
_ = __local_var_2
var __t4 *Constructor_Data_Map_Internal_Split
{
if (__local_var_2.Type == 9 && __local_var_2.IntVal == 324739070 && __local_var_2.UnsafePtr == nil) {
__t4 = &Constructor_Data_Map_Internal_Split{1, (*Constructor_Data_Maybe_Just)(nil), (*Constructor_Data_Map_Internal_Node)(nil), (*Constructor_Data_Map_Internal_Node)(nil)}
goto end_branch_4
} else {

}
}
{
if (__local_var_2.Type == 9 && __local_var_2.IntVal == 324739070 && __local_var_2.UnsafePtr != nil) {
// TAST (Let): v_3_0 -> gopurs_runtime.Value
v_3_0 := gopurs_runtime.Apply2(__local_var_0, __local_var_1, (*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V2)
_ = v_3_0
var __t3 *Constructor_Data_Map_Internal_Split
{
if (uint32(v_3_0.IntVal) == 1527465420) {
// TAST (Let): v1_4_1 -> gopurs_runtime.Value
v1_4_1 := gopurs_runtime.UncurriedApp3(Get_Data_Map_Internal_unsafeSplit(), __local_var_0, __local_var_1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V4)})
_ = v1_4_1
__t3 = &Constructor_Data_Map_Internal_Split{1, (*Constructor_Data_Map_Internal_Split)(v1_4_1.UnsafePtr).V0, (*Constructor_Data_Map_Internal_Split)(v1_4_1.UnsafePtr).V1, gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeBalancedNode(), (*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Split)(v1_4_1.UnsafePtr).V2)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V5)}))}
goto end_branch_3
} else {

}
}
{
if (uint32(v_3_0.IntVal) == 380165415) {
// TAST (Let): v1_4_2 -> gopurs_runtime.Value
v1_4_2 := gopurs_runtime.UncurriedApp3(Get_Data_Map_Internal_unsafeSplit(), __local_var_0, __local_var_1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V5)})
_ = v1_4_2
__t3 = &Constructor_Data_Map_Internal_Split{1, (*Constructor_Data_Map_Internal_Split)(v1_4_2.UnsafePtr).V0, gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeBalancedNode(), (*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V4)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Split)(v1_4_2.UnsafePtr).V1)})), (*Constructor_Data_Map_Internal_Split)(v1_4_2.UnsafePtr).V2}
goto end_branch_3
} else {

}
}
{
if (uint32(v_3_0.IntVal) == 902936544) {
__t3 = &Constructor_Data_Map_Internal_Split{1, &Constructor_Data_Maybe_Just{1, (*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V3}, (*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V4, (*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V5}
goto end_branch_3
} else {

}
}
{
__t3 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Split](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_3:
__t4 = __t3
goto end_branch_4
} else {

}
}
{
__t4 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Split](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_4:
return gopurs_runtime.Value{Type: 9, IntVal: 3373277644, UnsafePtr: unsafe.Pointer(__t4)}
}
}

func Call_Data_Map_Internal_unsafeSplitLast(__local_var_0_loop gopurs_runtime.Value, __local_var_1_loop gopurs_runtime.Value, __local_var_2_loop gopurs_runtime.Value, __local_var_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
unsafeSplitLast:
for {
if false { continue unsafeSplitLast }
var __local_var_0 gopurs_runtime.Value = __local_var_0_loop
_ = __local_var_0
var __local_var_1 gopurs_runtime.Value = __local_var_1_loop
_ = __local_var_1
var __local_var_2 gopurs_runtime.Value = __local_var_2_loop
_ = __local_var_2
var __local_var_3 gopurs_runtime.Value = __local_var_3_loop
_ = __local_var_3
var __t1 *Constructor_Data_Map_Internal_SplitLast
{
if (__local_var_3.Type == 9 && __local_var_3.IntVal == 324739070 && __local_var_3.UnsafePtr == nil) {
__t1 = &Constructor_Data_Map_Internal_SplitLast{1, __local_var_0, __local_var_1, gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](__local_var_2)}
goto end_branch_1
} else {

}
}
{
if (__local_var_3.Type == 9 && __local_var_3.IntVal == 324739070 && __local_var_3.UnsafePtr != nil) {
// TAST (Let): v1_4_0 -> gopurs_runtime.Value
v1_4_0 := gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeSplitLast(), (*Constructor_Data_Map_Internal_Node)(__local_var_3.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node)(__local_var_3.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_3.UnsafePtr).V4)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_3.UnsafePtr).V5)})
_ = v1_4_0
__t1 = &Constructor_Data_Map_Internal_SplitLast{1, (*Constructor_Data_Map_Internal_SplitLast)(v1_4_0.UnsafePtr).V0, (*Constructor_Data_Map_Internal_SplitLast)(v1_4_0.UnsafePtr).V1, gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeBalancedNode(), __local_var_0, __local_var_1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](__local_var_2))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_SplitLast)(v1_4_0.UnsafePtr).V2)}))}
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_SplitLast](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_1:
return gopurs_runtime.Value{Type: 9, IntVal: 2668112006, UnsafePtr: unsafe.Pointer(__t1)}
}
}

func Call_Data_Map_Internal_unsafeJoinNodes(__local_var_0_loop gopurs_runtime.Value, __local_var_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __local_var_0 gopurs_runtime.Value = __local_var_0_loop
_ = __local_var_0
var __local_var_1 gopurs_runtime.Value = __local_var_1_loop
_ = __local_var_1
var __t1 *Constructor_Data_Map_Internal_Node
{
if (__local_var_0.Type == 9 && __local_var_0.IntVal == 324739070 && __local_var_0.UnsafePtr == nil) {
__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](__local_var_1)
goto end_branch_1
} else {

}
}
{
if (__local_var_0.Type == 9 && __local_var_0.IntVal == 324739070 && __local_var_0.UnsafePtr != nil) {
// TAST (Let): v2_2_0 -> gopurs_runtime.Value
v2_2_0 := gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeSplitLast(), (*Constructor_Data_Map_Internal_Node)(__local_var_0.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node)(__local_var_0.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_0.UnsafePtr).V4)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_0.UnsafePtr).V5)})
_ = v2_2_0
__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeBalancedNode(), (*Constructor_Data_Map_Internal_SplitLast)(v2_2_0.UnsafePtr).V0, (*Constructor_Data_Map_Internal_SplitLast)(v2_2_0.UnsafePtr).V1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_SplitLast)(v2_2_0.UnsafePtr).V2)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](__local_var_1))}))
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_1:
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(__t1)}
}

func Call_Data_Map_Internal_unsafeDifference(__local_var_0_loop gopurs_runtime.Value, __local_var_1_loop gopurs_runtime.Value, __local_var_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
unsafeDifference:
for {
if false { continue unsafeDifference }
var __local_var_0 gopurs_runtime.Value = __local_var_0_loop
_ = __local_var_0
var __local_var_1 gopurs_runtime.Value = __local_var_1_loop
_ = __local_var_1
var __local_var_2 gopurs_runtime.Value = __local_var_2_loop
_ = __local_var_2
var __t1 *Constructor_Data_Map_Internal_Node
{
if (__local_var_1.Type == 9 && __local_var_1.IntVal == 324739070 && __local_var_1.UnsafePtr == nil) {
__t1 = (*Constructor_Data_Map_Internal_Node)(nil)
goto end_branch_1
} else {

}
}
{
if (__local_var_2.Type == 9 && __local_var_2.IntVal == 324739070 && __local_var_2.UnsafePtr == nil) {
__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](__local_var_1)
goto end_branch_1
} else {

}
}
{
if (__local_var_2.Type == 9 && __local_var_2.IntVal == 324739070 && __local_var_2.UnsafePtr != nil) {
// TAST (Let): v_3_0 -> gopurs_runtime.Value
v_3_0 := gopurs_runtime.UncurriedApp3(Get_Data_Map_Internal_unsafeSplit(), __local_var_0, (*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V2, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](__local_var_1))})
_ = v_3_0
__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.UncurriedApp2(Get_Data_Map_Internal_unsafeJoinNodes(), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.UncurriedApp3(Get_Data_Map_Internal_unsafeDifference(), __local_var_0, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Split)(v_3_0.UnsafePtr).V1)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V4)})))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.UncurriedApp3(Get_Data_Map_Internal_unsafeDifference(), __local_var_0, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Split)(v_3_0.UnsafePtr).V2)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V5)})))}))
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_1:
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(__t1)}
}
}

func Call_Data_Map_Internal_unsafeIntersectionWith(__local_var_0_loop gopurs_runtime.Value, __local_var_1_loop gopurs_runtime.Value, __local_var_2_loop gopurs_runtime.Value, __local_var_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
unsafeIntersectionWith:
for {
if false { continue unsafeIntersectionWith }
var __local_var_0 gopurs_runtime.Value = __local_var_0_loop
_ = __local_var_0
var __local_var_1 gopurs_runtime.Value = __local_var_1_loop
_ = __local_var_1
var __local_var_2 gopurs_runtime.Value = __local_var_2_loop
_ = __local_var_2
var __local_var_3 gopurs_runtime.Value = __local_var_3_loop
_ = __local_var_3
var __t6 *Constructor_Data_Map_Internal_Node
{
if (__local_var_2.Type == 9 && __local_var_2.IntVal == 324739070 && __local_var_2.UnsafePtr == nil) {
__t6 = (*Constructor_Data_Map_Internal_Node)(nil)
goto end_branch_6
} else {

}
}
{
if (__local_var_3.Type == 9 && __local_var_3.IntVal == 324739070 && __local_var_3.UnsafePtr == nil) {
__t6 = (*Constructor_Data_Map_Internal_Node)(nil)
goto end_branch_6
} else {

}
}
{
if (__local_var_3.Type == 9 && __local_var_3.IntVal == 324739070 && __local_var_3.UnsafePtr != nil) {
// TAST (Let): v_4_0 -> gopurs_runtime.Value
v_4_0 := gopurs_runtime.UncurriedApp3(Get_Data_Map_Internal_unsafeSplit(), __local_var_0, (*Constructor_Data_Map_Internal_Node)(__local_var_3.UnsafePtr).V2, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](__local_var_2))})
_ = v_4_0
// TAST (Let): l_prime_5_1 -> gopurs_runtime.Value
l_prime_5_1 := gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeIntersectionWith(), __local_var_0, __local_var_1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Split)(v_4_0.UnsafePtr).V1)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_3.UnsafePtr).V4)})
_ = l_prime_5_1
// TAST (Let): r_prime_6_2 -> gopurs_runtime.Value
r_prime_6_2 := gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeIntersectionWith(), __local_var_0, __local_var_1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Split)(v_4_0.UnsafePtr).V2)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_3.UnsafePtr).V5)})
_ = r_prime_6_2
var __t5 *Constructor_Data_Map_Internal_Node
{
var __t_tag_3 *Constructor_Data_Maybe_Just = (*Constructor_Data_Map_Internal_Split)(v_4_0.UnsafePtr).V0
if (__t_tag_3 != nil) {
__t5 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeBalancedNode(), (*Constructor_Data_Map_Internal_Node)(__local_var_3.UnsafePtr).V2, gopurs_runtime.Apply2(__local_var_1, ((*Constructor_Data_Map_Internal_Split)(v_4_0.UnsafePtr).V0).V0, (*Constructor_Data_Map_Internal_Node)(__local_var_3.UnsafePtr).V3), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](l_prime_5_1))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](r_prime_6_2))}))
goto end_branch_5
} else {

}
}
{
var __t_tag_4 *Constructor_Data_Maybe_Just = (*Constructor_Data_Map_Internal_Split)(v_4_0.UnsafePtr).V0
if (__t_tag_4 == nil) {
__t5 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.UncurriedApp2(Get_Data_Map_Internal_unsafeJoinNodes(), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](l_prime_5_1))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](r_prime_6_2))}))
goto end_branch_5
} else {

}
}
{
__t5 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_5:
__t6 = __t5
goto end_branch_6
} else {

}
}
{
__t6 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_6:
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(__t6)}
}
}

func Call_Data_Map_Internal_unsafeUnionWith(__local_var_0_loop gopurs_runtime.Value, __local_var_1_loop gopurs_runtime.Value, __local_var_2_loop gopurs_runtime.Value, __local_var_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
unsafeUnionWith:
for {
if false { continue unsafeUnionWith }
var __local_var_0 gopurs_runtime.Value = __local_var_0_loop
_ = __local_var_0
var __local_var_1 gopurs_runtime.Value = __local_var_1_loop
_ = __local_var_1
var __local_var_2 gopurs_runtime.Value = __local_var_2_loop
_ = __local_var_2
var __local_var_3 gopurs_runtime.Value = __local_var_3_loop
_ = __local_var_3
var __t6 *Constructor_Data_Map_Internal_Node
{
if (__local_var_2.Type == 9 && __local_var_2.IntVal == 324739070 && __local_var_2.UnsafePtr == nil) {
__t6 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](__local_var_3)
goto end_branch_6
} else {

}
}
{
if (__local_var_3.Type == 9 && __local_var_3.IntVal == 324739070 && __local_var_3.UnsafePtr == nil) {
__t6 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](__local_var_2)
goto end_branch_6
} else {

}
}
{
if (__local_var_3.Type == 9 && __local_var_3.IntVal == 324739070 && __local_var_3.UnsafePtr != nil) {
// TAST (Let): v_4_0 -> gopurs_runtime.Value
v_4_0 := gopurs_runtime.UncurriedApp3(Get_Data_Map_Internal_unsafeSplit(), __local_var_0, (*Constructor_Data_Map_Internal_Node)(__local_var_3.UnsafePtr).V2, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](__local_var_2))})
_ = v_4_0
// TAST (Let): l_prime_5_1 -> gopurs_runtime.Value
l_prime_5_1 := gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeUnionWith(), __local_var_0, __local_var_1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Split)(v_4_0.UnsafePtr).V1)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_3.UnsafePtr).V4)})
_ = l_prime_5_1
// TAST (Let): r_prime_6_2 -> gopurs_runtime.Value
r_prime_6_2 := gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeUnionWith(), __local_var_0, __local_var_1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Split)(v_4_0.UnsafePtr).V2)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_3.UnsafePtr).V5)})
_ = r_prime_6_2
var __t5 *Constructor_Data_Map_Internal_Node
{
var __t_tag_3 *Constructor_Data_Maybe_Just = (*Constructor_Data_Map_Internal_Split)(v_4_0.UnsafePtr).V0
if (__t_tag_3 != nil) {
__t5 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeBalancedNode(), (*Constructor_Data_Map_Internal_Node)(__local_var_3.UnsafePtr).V2, gopurs_runtime.Apply2(__local_var_1, ((*Constructor_Data_Map_Internal_Split)(v_4_0.UnsafePtr).V0).V0, (*Constructor_Data_Map_Internal_Node)(__local_var_3.UnsafePtr).V3), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](l_prime_5_1))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](r_prime_6_2))}))
goto end_branch_5
} else {

}
}
{
var __t_tag_4 *Constructor_Data_Maybe_Just = (*Constructor_Data_Map_Internal_Split)(v_4_0.UnsafePtr).V0
if (__t_tag_4 == nil) {
__t5 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeBalancedNode(), (*Constructor_Data_Map_Internal_Node)(__local_var_3.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node)(__local_var_3.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](l_prime_5_1))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](r_prime_6_2))}))
goto end_branch_5
} else {

}
}
{
__t5 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_5:
__t6 = __t5
goto end_branch_6
} else {

}
}
{
__t6 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_6:
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(__t6)}
}
}

func Call_Data_Map_Internal_unionWith(dictOrd_0_loop *Constructor_Data_Ord_Ord) gopurs_runtime.Value {
var dictOrd_0 *Constructor_Data_Ord_Ord = dictOrd_0_loop
_ = dictOrd_0
// TAST (Let): compare_1_0 -> gopurs_runtime.Value
compare_1_0 := gopurs_runtime.Box(dictOrd_0.V1)
_ = compare_1_0
return gopurs_runtime.Func(func(app_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m2_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeUnionWith(), compare_1_0, app_2, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](m1_3))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](m2_4))})))}
})
})
})
}

func Call_Data_Map_Internal_union(dictOrd_0_loop *Constructor_Data_Ord_Ord) gopurs_runtime.Value {
var dictOrd_0 *Constructor_Data_Ord_Ord = dictOrd_0_loop
_ = dictOrd_0
// TAST (Let): compare_1_0 -> gopurs_runtime.Value
compare_1_0 := gopurs_runtime.Box(dictOrd_0.V1)
_ = compare_1_0
return gopurs_runtime.Func(func(m1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m2_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeUnionWith(), compare_1_0, Get_Data_Function_go__const(), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](m1_2))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](m2_3))})))}
})
})
}

func Call_Data_Map_Internal_update(dictOrd_0_loop *Constructor_Data_Ord_Ord, f_1_loop gopurs_runtime.Value, k_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 *Constructor_Data_Ord_Ord = dictOrd_0_loop
_ = dictOrd_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var k_2 gopurs_runtime.Value = k_2_loop
_ = k_2
var go__go_3_0_1 gopurs_runtime.Value
_ = go__go_3_0_1
go__go_3_0_1 = gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t5 *Constructor_Data_Map_Internal_Node
{
if (v_4.Type == 9 && v_4.IntVal == 324739070 && v_4.UnsafePtr == nil) {
__t5 = (*Constructor_Data_Map_Internal_Node)(nil)
goto end_branch_5
} else {

}
}
{
if (v_4.Type == 9 && v_4.IntVal == 324739070 && v_4.UnsafePtr != nil) {
// TAST (Let): v1_5_1 -> gopurs_runtime.Value
v1_5_1 := gopurs_runtime.Apply2(gopurs_runtime.Box(dictOrd_0.V1), k_2, (*Constructor_Data_Map_Internal_Node)(v_4.UnsafePtr).V2)
_ = v1_5_1
var __t4 *Constructor_Data_Map_Internal_Node
{
if (uint32(v1_5_1.IntVal) == 1527465420) {
__t4 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeBalancedNode(), (*Constructor_Data_Map_Internal_Node)(v_4.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node)(v_4.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.Apply(go__go_3_0_1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v_4.UnsafePtr).V4)})))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v_4.UnsafePtr).V5)}))
goto end_branch_4
} else {

}
}
{
if (uint32(v1_5_1.IntVal) == 380165415) {
__t4 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeBalancedNode(), (*Constructor_Data_Map_Internal_Node)(v_4.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node)(v_4.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v_4.UnsafePtr).V4)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.Apply(go__go_3_0_1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v_4.UnsafePtr).V5)})))}))
goto end_branch_4
} else {

}
}
{
if (uint32(v1_5_1.IntVal) == 902936544) {
// TAST (Let): v2_6_2 -> *Constructor_Data_Maybe_Just
v2_6_2 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply(f_1, (*Constructor_Data_Map_Internal_Node)(v_4.UnsafePtr).V3))
_ = v2_6_2
var __t3 *Constructor_Data_Map_Internal_Node
{
if (v2_6_2 == nil) {
__t3 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.UncurriedApp2(Get_Data_Map_Internal_unsafeJoinNodes(), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v_4.UnsafePtr).V4)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v_4.UnsafePtr).V5)}))
goto end_branch_3
} else {

}
}
{
if (v2_6_2 != nil) {
__t3 = &Constructor_Data_Map_Internal_Node{1, (*Constructor_Data_Map_Internal_Node)(v_4.UnsafePtr).V0, (*Constructor_Data_Map_Internal_Node)(v_4.UnsafePtr).V1, (*Constructor_Data_Map_Internal_Node)(v_4.UnsafePtr).V2, (v2_6_2).V0, (*Constructor_Data_Map_Internal_Node)(v_4.UnsafePtr).V4, (*Constructor_Data_Map_Internal_Node)(v_4.UnsafePtr).V5}
goto end_branch_3
} else {

}
}
{
__t3 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_3:
__t4 = __t3
goto end_branch_4
} else {

}
}
{
__t4 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_4:
__t5 = __t4
goto end_branch_5
} else {

}
}
{
__t5 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_5:
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(__t5)}
})
return go__go_3_0_1
}

func Call_Data_Map_Internal_showTree(dictShow_0_loop *Constructor_Data_Show_Show, dictShow1_1_loop *Constructor_Data_Show_Show) gopurs_runtime.Value {
var dictShow_0 *Constructor_Data_Show_Show = dictShow_0_loop
_ = dictShow_0
var dictShow1_1 *Constructor_Data_Show_Show = dictShow1_1_loop
_ = dictShow1_1
var go__go_2_0_2 gopurs_runtime.Value
_ = go__go_2_0_2
go__go_2_0_2 = gopurs_runtime.Func(func(ind_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 string
{
if (v_4.Type == 9 && v_4.IntVal == 324739070 && v_4.UnsafePtr == nil) {
__t1 = (ind_3.StrVal()) + ("Leaf")
goto end_branch_1
} else {

}
}
{
if (v_4.Type == 9 && v_4.IntVal == 324739070 && v_4.UnsafePtr != nil) {
__t1 = ((((((((((ind_3.StrVal()) + ("[")) + (gopurs_runtime.Apply(Get_Data_Show_showIntImpl(), gopurs_runtime.Int((*Constructor_Data_Map_Internal_Node)(v_4.UnsafePtr).V0)).StrVal())) + ("] ")) + (gopurs_runtime.Apply(gopurs_runtime.Box(dictShow_0.V0), (*Constructor_Data_Map_Internal_Node)(v_4.UnsafePtr).V2).StrVal())) + (" => ")) + (gopurs_runtime.Apply(gopurs_runtime.Box(dictShow1_1.V0), (*Constructor_Data_Map_Internal_Node)(v_4.UnsafePtr).V3).StrVal())) + ("\x0a")) + (gopurs_runtime.Apply2(go__go_2_0_2, gopurs_runtime.Str((ind_3.StrVal()) + ("    ")), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v_4.UnsafePtr).V4)}).StrVal())) + ("\x0a")) + (gopurs_runtime.Apply2(go__go_2_0_2, gopurs_runtime.Str((ind_3.StrVal()) + ("    ")), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v_4.UnsafePtr).V5)}).StrVal())
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }().StrVal()
}
end_branch_1:
return gopurs_runtime.Str(__t1)
})
})
return gopurs_runtime.Apply(go__go_2_0_2, gopurs_runtime.Str(""))
}

func Call_Data_Map_Internal_semigroupMap(_dollar__unused_0_loop gopurs_runtime.Value, dictOrd_1_loop gopurs_runtime.Value, dictSemigroup_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var _dollar__unused_0 gopurs_runtime.Value = _dollar__unused_0_loop
_ = _dollar__unused_0
var dictOrd_1 gopurs_runtime.Value = dictOrd_1_loop
_ = dictOrd_1
var dictSemigroup_2 gopurs_runtime.Value = dictSemigroup_2_loop
_ = dictSemigroup_2
// TAST (Let): compare_3_0 -> gopurs_runtime.Value
compare_3_0 := gopurs_runtime.RecordGet(dictOrd_1, "compare")
_ = compare_3_0
// TAST (Let): __local_var_4_1 -> gopurs_runtime.Value
__local_var_4_1 := gopurs_runtime.RecordGet(dictSemigroup_2, "append")
_ = __local_var_4_1
return gopurs_runtime.RecordDict1("append", gopurs_runtime.Func(func(m1_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m2_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeUnionWith(), compare_3_0, __local_var_4_1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](m1_5))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](m2_6))})))}
})
}))
}

func Call_Data_Map_Internal_semigroupMap1(dictOrd_0_loop gopurs_runtime.Value, dictSemigroup_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
var dictSemigroup_1 gopurs_runtime.Value = dictSemigroup_1_loop
_ = dictSemigroup_1
// TAST (Let): compare_2_0 -> gopurs_runtime.Value
compare_2_0 := gopurs_runtime.RecordGet(dictOrd_0, "compare")
_ = compare_2_0
// TAST (Let): __local_var_3_1 -> gopurs_runtime.Value
__local_var_3_1 := gopurs_runtime.RecordGet(dictSemigroup_1, "append")
_ = __local_var_3_1
return gopurs_runtime.RecordDict1("append", gopurs_runtime.Func(func(m1_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m2_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeUnionWith(), compare_2_0, __local_var_3_1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](m1_4))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](m2_5))})))}
})
}))
}

func Call_Data_Map_Internal_pop(dictOrd_0_loop *Constructor_Data_Ord_Ord) gopurs_runtime.Value {
var dictOrd_0 *Constructor_Data_Ord_Ord = dictOrd_0_loop
_ = dictOrd_0
// TAST (Let): compare_1_0 -> gopurs_runtime.Value
compare_1_0 := gopurs_runtime.Box(dictOrd_0.V1)
_ = compare_1_0
return gopurs_runtime.Func(func(k_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): v_4_1 -> gopurs_runtime.Value
v_4_1 := gopurs_runtime.UncurriedApp3(Get_Data_Map_Internal_unsafeSplit(), compare_1_0, k_2, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](m_3))})
_ = v_4_1
// TAST (Let): __local_var_5_2 -> gopurs_runtime.Value
var __local_var_5_2 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Split)(v_4_1.UnsafePtr).V1)}
// TAST (Let): __local_var_6_3 -> gopurs_runtime.Value
var __local_var_6_3 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Split)(v_4_1.UnsafePtr).V2)}
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_Maybe_functorMaybe(), "map"), gopurs_runtime.Func(func(a_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, a_7, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.UncurriedApp2(Get_Data_Map_Internal_unsafeJoinNodes(), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](__local_var_5_2))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](__local_var_6_3))})))}})}
}), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Split)(v_4_1.UnsafePtr).V0)})))}
})
})
}

func Call_Data_Map_Internal_member(dictOrd_0_loop *Constructor_Data_Ord_Ord, k_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 *Constructor_Data_Ord_Ord = dictOrd_0_loop
_ = dictOrd_0
var k_1 gopurs_runtime.Value = k_1_loop
_ = k_1
var go__go_2_0_3 gopurs_runtime.Value
go__go_2_0_3 = gopurs_runtime.Func(func(v_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_3_loop gopurs_runtime.Value = v_3_loop_val
go__go_2_0_3:
for {
if false { continue go__go_2_0_3 }
var v_3 gopurs_runtime.Value = v_3_loop
_ = v_3
var __t3 bool
{
if (v_3.Type == 9 && v_3.IntVal == 324739070 && v_3.UnsafePtr == nil) {
__t3 = false
goto end_branch_3
} else {

}
}
{
if (v_3.Type == 9 && v_3.IntVal == 324739070 && v_3.UnsafePtr != nil) {
// TAST (Let): v1_4_1 -> gopurs_runtime.Value
v1_4_1 := gopurs_runtime.Apply2(gopurs_runtime.Box(dictOrd_0.V1), k_1, (*Constructor_Data_Map_Internal_Node)(v_3.UnsafePtr).V2)
_ = v1_4_1
var __t2 bool
{
if (uint32(v1_4_1.IntVal) == 1527465420) {
v_3_loop = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v_3.UnsafePtr).V4)}
continue go__go_2_0_3
__t2 = (gopurs_runtime.Value{}.IntVal) != (0)
goto end_branch_2
} else {

}
}
{
if (uint32(v1_4_1.IntVal) == 380165415) {
v_3_loop = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v_3.UnsafePtr).V5)}
continue go__go_2_0_3
__t2 = (gopurs_runtime.Value{}.IntVal) != (0)
goto end_branch_2
} else {

}
}
{
if (uint32(v1_4_1.IntVal) == 902936544) {
__t2 = true
goto end_branch_2
} else {

}
}
{
__t2 = (func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal) != (0)
}
end_branch_2:
__t3 = __t2
goto end_branch_3
} else {

}
}
{
__t3 = (func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal) != (0)
}
end_branch_3:
return gopurs_runtime.Bool(__t3)
}
}()
})
return go__go_2_0_3
}

func Call_Data_Map_Internal_mapMaybeWithKey(dictOrd_0_loop *Constructor_Data_Ord_Ord, f_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 *Constructor_Data_Ord_Ord = dictOrd_0_loop
_ = dictOrd_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var go__go_2_0_4 gopurs_runtime.Value
_ = go__go_2_0_4
go__go_2_0_4 = gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t3 *Constructor_Data_Map_Internal_Node
{
if (v_3.Type == 9 && v_3.IntVal == 324739070 && v_3.UnsafePtr == nil) {
__t3 = (*Constructor_Data_Map_Internal_Node)(nil)
goto end_branch_3
} else {

}
}
{
if (v_3.Type == 9 && v_3.IntVal == 324739070 && v_3.UnsafePtr != nil) {
// TAST (Let): v2_4_1 -> gopurs_runtime.Value
v2_4_1 := gopurs_runtime.Apply2(f_1, (*Constructor_Data_Map_Internal_Node)(v_3.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node)(v_3.UnsafePtr).V3)
_ = v2_4_1
var __t2 *Constructor_Data_Map_Internal_Node
{
if (v2_4_1.Type == 9 && v2_4_1.IntVal == 930809136 && v2_4_1.UnsafePtr != nil) {
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeBalancedNode(), (*Constructor_Data_Map_Internal_Node)(v_3.UnsafePtr).V2, (*Constructor_Data_Maybe_Just)(v2_4_1.UnsafePtr).V0, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.Apply(go__go_2_0_4, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v_3.UnsafePtr).V4)})))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.Apply(go__go_2_0_4, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v_3.UnsafePtr).V5)})))}))
goto end_branch_2
} else {

}
}
{
if (v2_4_1.Type == 9 && v2_4_1.IntVal == 930809136 && v2_4_1.UnsafePtr == nil) {
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.UncurriedApp2(Get_Data_Map_Internal_unsafeJoinNodes(), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.Apply(go__go_2_0_4, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v_3.UnsafePtr).V4)})))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.Apply(go__go_2_0_4, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v_3.UnsafePtr).V5)})))}))
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_2:
__t3 = __t2
goto end_branch_3
} else {

}
}
{
__t3 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_3:
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(__t3)}
})
return go__go_2_0_4
}

func Call_Data_Map_Internal_mapMaybe(dictOrd_0_loop *Constructor_Data_Ord_Ord, x_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 *Constructor_Data_Ord_Ord = dictOrd_0_loop
_ = dictOrd_0
var x_1 gopurs_runtime.Value = x_1_loop
_ = x_1
var go__go_2_0_5 gopurs_runtime.Value
_ = go__go_2_0_5
go__go_2_0_5 = gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t3 *Constructor_Data_Map_Internal_Node
{
if (v_3.Type == 9 && v_3.IntVal == 324739070 && v_3.UnsafePtr == nil) {
__t3 = (*Constructor_Data_Map_Internal_Node)(nil)
goto end_branch_3
} else {

}
}
{
if (v_3.Type == 9 && v_3.IntVal == 324739070 && v_3.UnsafePtr != nil) {
// TAST (Let): v2_4_1 -> gopurs_runtime.Value
v2_4_1 := gopurs_runtime.Apply(x_1, (*Constructor_Data_Map_Internal_Node)(v_3.UnsafePtr).V3)
_ = v2_4_1
var __t2 *Constructor_Data_Map_Internal_Node
{
if (v2_4_1.Type == 9 && v2_4_1.IntVal == 930809136 && v2_4_1.UnsafePtr != nil) {
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeBalancedNode(), (*Constructor_Data_Map_Internal_Node)(v_3.UnsafePtr).V2, (*Constructor_Data_Maybe_Just)(v2_4_1.UnsafePtr).V0, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.Apply(go__go_2_0_5, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v_3.UnsafePtr).V4)})))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.Apply(go__go_2_0_5, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v_3.UnsafePtr).V5)})))}))
goto end_branch_2
} else {

}
}
{
if (v2_4_1.Type == 9 && v2_4_1.IntVal == 930809136 && v2_4_1.UnsafePtr == nil) {
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.UncurriedApp2(Get_Data_Map_Internal_unsafeJoinNodes(), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.Apply(go__go_2_0_5, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v_3.UnsafePtr).V4)})))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.Apply(go__go_2_0_5, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v_3.UnsafePtr).V5)})))}))
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_2:
__t3 = __t2
goto end_branch_3
} else {

}
}
{
__t3 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_3:
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(__t3)}
})
return go__go_2_0_5
}

func Call_Data_Map_Internal_lookupLE(dictOrd_0_loop *Constructor_Data_Ord_Ord, k_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 *Constructor_Data_Ord_Ord = dictOrd_0_loop
_ = dictOrd_0
var k_1 gopurs_runtime.Value = k_1_loop
_ = k_1
var go__go_2_0_6 gopurs_runtime.Value
_ = go__go_2_0_6
go__go_2_0_6 = gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t5 *Constructor_Data_Maybe_Just
{
if (v_3.Type == 9 && v_3.IntVal == 324739070 && v_3.UnsafePtr == nil) {
__t5 = (*Constructor_Data_Maybe_Just)(nil)
goto end_branch_5
} else {

}
}
{
if (v_3.Type == 9 && v_3.IntVal == 324739070 && v_3.UnsafePtr != nil) {
// TAST (Let): v1_4_1 -> gopurs_runtime.Value
v1_4_1 := gopurs_runtime.Apply2(gopurs_runtime.Box(dictOrd_0.V1), k_1, (*Constructor_Data_Map_Internal_Node)(v_3.UnsafePtr).V2)
_ = v1_4_1
var __t4 *Constructor_Data_Maybe_Just
{
if (uint32(v1_4_1.IntVal) == 1527465420) {
__t4 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply(go__go_2_0_6, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v_3.UnsafePtr).V4)}))
goto end_branch_4
} else {

}
}
{
if (uint32(v1_4_1.IntVal) == 380165415) {
// TAST (Let): v2_5_2 -> *Constructor_Data_Maybe_Just
v2_5_2 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply(go__go_2_0_6, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v_3.UnsafePtr).V5)}))
_ = v2_5_2
var __t3 *Constructor_Data_Maybe_Just
{
if (v2_5_2 == nil) {
__t3 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, gopurs_runtime.RecordDict2("key", "value", (*Constructor_Data_Map_Internal_Node)(v_3.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node)(v_3.UnsafePtr).V3)})})
goto end_branch_3
} else {

}
}
{
__t3 = v2_5_2
}
end_branch_3:
__t4 = __t3
goto end_branch_4
} else {

}
}
{
if (uint32(v1_4_1.IntVal) == 902936544) {
__t4 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, gopurs_runtime.RecordDict2("key", "value", (*Constructor_Data_Map_Internal_Node)(v_3.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node)(v_3.UnsafePtr).V3)})})
goto end_branch_4
} else {

}
}
{
__t4 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_4:
__t5 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t4)})
goto end_branch_5
} else {

}
}
{
__t5 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_5:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t5)}
})
return go__go_2_0_6
}

func Call_Data_Map_Internal_lookupGE(dictOrd_0_loop *Constructor_Data_Ord_Ord, k_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 *Constructor_Data_Ord_Ord = dictOrd_0_loop
_ = dictOrd_0
var k_1 gopurs_runtime.Value = k_1_loop
_ = k_1
var go__go_2_0_7 gopurs_runtime.Value
_ = go__go_2_0_7
go__go_2_0_7 = gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t5 *Constructor_Data_Maybe_Just
{
if (v_3.Type == 9 && v_3.IntVal == 324739070 && v_3.UnsafePtr == nil) {
__t5 = (*Constructor_Data_Maybe_Just)(nil)
goto end_branch_5
} else {

}
}
{
if (v_3.Type == 9 && v_3.IntVal == 324739070 && v_3.UnsafePtr != nil) {
// TAST (Let): v1_4_1 -> gopurs_runtime.Value
v1_4_1 := gopurs_runtime.Apply2(gopurs_runtime.Box(dictOrd_0.V1), k_1, (*Constructor_Data_Map_Internal_Node)(v_3.UnsafePtr).V2)
_ = v1_4_1
var __t4 *Constructor_Data_Maybe_Just
{
if (uint32(v1_4_1.IntVal) == 1527465420) {
// TAST (Let): v2_5_2 -> *Constructor_Data_Maybe_Just
v2_5_2 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply(go__go_2_0_7, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v_3.UnsafePtr).V4)}))
_ = v2_5_2
var __t3 *Constructor_Data_Maybe_Just
{
if (v2_5_2 == nil) {
__t3 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, gopurs_runtime.RecordDict2("key", "value", (*Constructor_Data_Map_Internal_Node)(v_3.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node)(v_3.UnsafePtr).V3)})})
goto end_branch_3
} else {

}
}
{
__t3 = v2_5_2
}
end_branch_3:
__t4 = __t3
goto end_branch_4
} else {

}
}
{
if (uint32(v1_4_1.IntVal) == 380165415) {
__t4 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply(go__go_2_0_7, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v_3.UnsafePtr).V5)}))
goto end_branch_4
} else {

}
}
{
if (uint32(v1_4_1.IntVal) == 902936544) {
__t4 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, gopurs_runtime.RecordDict2("key", "value", (*Constructor_Data_Map_Internal_Node)(v_3.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node)(v_3.UnsafePtr).V3)})})
goto end_branch_4
} else {

}
}
{
__t4 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_4:
__t5 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t4)})
goto end_branch_5
} else {

}
}
{
__t5 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_5:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t5)}
})
return go__go_2_0_7
}

func Call_Data_Map_Internal_lookup(dictOrd_0_loop *Constructor_Data_Ord_Ord, k_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 *Constructor_Data_Ord_Ord = dictOrd_0_loop
_ = dictOrd_0
var k_1 gopurs_runtime.Value = k_1_loop
_ = k_1
var go__go_2_0_8 gopurs_runtime.Value
go__go_2_0_8 = gopurs_runtime.Func(func(v_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_3_loop gopurs_runtime.Value = v_3_loop_val
go__go_2_0_8:
for {
if false { continue go__go_2_0_8 }
var v_3 gopurs_runtime.Value = v_3_loop
_ = v_3
var __t3 *Constructor_Data_Maybe_Just
{
if (v_3.Type == 9 && v_3.IntVal == 324739070 && v_3.UnsafePtr == nil) {
__t3 = (*Constructor_Data_Maybe_Just)(nil)
goto end_branch_3
} else {

}
}
{
if (v_3.Type == 9 && v_3.IntVal == 324739070 && v_3.UnsafePtr != nil) {
// TAST (Let): v1_4_1 -> gopurs_runtime.Value
v1_4_1 := gopurs_runtime.Apply2(gopurs_runtime.Box(dictOrd_0.V1), k_1, (*Constructor_Data_Map_Internal_Node)(v_3.UnsafePtr).V2)
_ = v1_4_1
var __t2 *Constructor_Data_Maybe_Just
{
if (uint32(v1_4_1.IntVal) == 1527465420) {
v_3_loop = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v_3.UnsafePtr).V4)}
continue go__go_2_0_8
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Value{})
goto end_branch_2
} else {

}
}
{
if (uint32(v1_4_1.IntVal) == 380165415) {
v_3_loop = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v_3.UnsafePtr).V5)}
continue go__go_2_0_8
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Value{})
goto end_branch_2
} else {

}
}
{
if (uint32(v1_4_1.IntVal) == 902936544) {
__t2 = &Constructor_Data_Maybe_Just{1, (*Constructor_Data_Map_Internal_Node)(v_3.UnsafePtr).V3}
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_2:
__t3 = __t2
goto end_branch_3
} else {

}
}
{
__t3 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_3:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t3)}
}
}()
})
return go__go_2_0_8
}

func Call_Data_Map_Internal_iterMapU(iter_0_loop gopurs_runtime.Value, v_1_loop *Constructor_Data_Map_Internal_Node) gopurs_runtime.Value {
var iter_0 gopurs_runtime.Value = iter_0_loop
_ = iter_0
var v_1 *Constructor_Data_Map_Internal_Node = v_1_loop
_ = v_1
var __t6 *Constructor_Data_Map_Internal_IterEmit
{
if (v_1 == nil) {
__t6 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_IterEmit](iter_0)
goto end_branch_6
} else {

}
}
{
if (v_1 != nil) {
var __t5 *Constructor_Data_Map_Internal_IterEmit
{
var __t_tag_2 *Constructor_Data_Map_Internal_Node = (v_1).V4
if (__t_tag_2 == nil) {
var __t4 *Constructor_Data_Map_Internal_IterEmit
{
var __t_tag_3 *Constructor_Data_Map_Internal_Node = (v_1).V5
if (__t_tag_3 == nil) {
__t4 = &Constructor_Data_Map_Internal_IterEmit{1, (v_1).V2, (v_1).V3, iter_0}
goto end_branch_4
} else {

}
}
{
__t4 = &Constructor_Data_Map_Internal_IterEmit{1, (v_1).V2, (v_1).V3, gopurs_runtime.Value{Type: 9, IntVal: 2861335956, UnsafePtr: unsafe.Pointer(&Constructor_Data_Map_Internal_IterNode{1, (v_1).V5, iter_0})}}
}
end_branch_4:
__t5 = __t4
goto end_branch_5
} else {

}
}
{
var __t1 *Constructor_Data_Map_Internal_IterEmit
{
var __t_tag_0 *Constructor_Data_Map_Internal_Node = (v_1).V5
if (__t_tag_0 == nil) {
__t1 = &Constructor_Data_Map_Internal_IterEmit{1, (v_1).V2, (v_1).V3, gopurs_runtime.Value{Type: 9, IntVal: 2861335956, UnsafePtr: unsafe.Pointer(&Constructor_Data_Map_Internal_IterNode{1, (v_1).V4, iter_0})}}
goto end_branch_1
} else {

}
}
{
__t1 = &Constructor_Data_Map_Internal_IterEmit{1, (v_1).V2, (v_1).V3, gopurs_runtime.Value{Type: 9, IntVal: 2861335956, UnsafePtr: unsafe.Pointer(&Constructor_Data_Map_Internal_IterNode{1, (v_1).V4, gopurs_runtime.Value{Type: 9, IntVal: 2861335956, UnsafePtr: unsafe.Pointer(&Constructor_Data_Map_Internal_IterNode{1, (v_1).V5, iter_0})}})}}
}
end_branch_1:
__t5 = __t1
}
end_branch_5:
__t6 = __t5
goto end_branch_6
} else {

}
}
{
__t6 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_IterEmit](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_6:
return gopurs_runtime.Value{Type: 9, IntVal: 1343415489, UnsafePtr: unsafe.Pointer(__t6)}
}

func Call_Data_Map_Internal_toUnfoldableUnordered(dictUnfoldable_0_loop *Constructor_Data_Unfoldable_Unfoldable) gopurs_runtime.Value {
var dictUnfoldable_0 *Constructor_Data_Unfoldable_Unfoldable = dictUnfoldable_0_loop
_ = dictUnfoldable_0
// TAST (Let): __local_var_1_0 -> gopurs_runtime.Value
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.Box(dictUnfoldable_0.V1), Get_Data_Map_Internal_stepUnfoldrUnordered())
_ = __local_var_1_0
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_1_0, gopurs_runtime.Value{Type: 9, IntVal: 2861335956, UnsafePtr: unsafe.Pointer(&Constructor_Data_Map_Internal_IterNode{1, gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](x_2), gopurs_runtime.Value{Type: 9, IntVal: 2509360378, UnsafePtr: unsafe.Pointer(nil)}})})
})
}

func Call_Data_Map_Internal_eqMapIter(dictEq_0_loop gopurs_runtime.Value, dictEq1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictEq_0 gopurs_runtime.Value = dictEq_0_loop
_ = dictEq_0
var dictEq1_1 gopurs_runtime.Value = dictEq1_1_loop
_ = dictEq1_1
var go__go_2_0_11 gopurs_runtime.Value
go__go_2_0_11 = gopurs_runtime.Func(func(a_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var a_3_loop gopurs_runtime.Value = a_3_loop_val
var b_4_loop gopurs_runtime.Value = b_4_loop_val
go__go_2_0_11:
for {
if false { continue go__go_2_0_11 }
var a_3 gopurs_runtime.Value = a_3_loop
_ = a_3
var b_4 gopurs_runtime.Value = b_4_loop
_ = b_4
// TAST (Let): v_5_1 -> *Constructor_Data_Map_Internal_IterNext
v_5_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_IterNext](gopurs_runtime.Apply(Get_Data_Map_Internal_stepAsc(), a_3))
_ = v_5_1
var __t4 bool
{
if (v_5_1 != nil) {
// TAST (Let): v2_6_2 -> *Constructor_Data_Map_Internal_IterNext
v2_6_2 := gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_IterNext](gopurs_runtime.Apply(Get_Data_Map_Internal_stepAsc(), b_4))
_ = v2_6_2
var __t3 bool
{
if ((v2_6_2 != nil)) && (((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq_0, "eq"), (v_5_1).V0, (v2_6_2).V0).IntVal) != (0)) && ((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq1_1, "eq"), (v_5_1).V1, (v2_6_2).V1).IntVal) != (0))) {
a_3_loop = (v_5_1).V2
b_4_loop = (v2_6_2).V2
continue go__go_2_0_11
__t3 = (gopurs_runtime.Value{}.IntVal) != (0)
goto end_branch_3
} else {

}
}
{
__t3 = false
}
end_branch_3:
__t4 = __t3
goto end_branch_4
} else {

}
}
{
if (v_5_1 == nil) {
__t4 = true
goto end_branch_4
} else {

}
}
{
__t4 = (func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal) != (0)
}
end_branch_4:
return gopurs_runtime.Bool(__t4)
}
}()
})
})
return gopurs_runtime.RecordDict1("eq", go__go_2_0_11)
}

func Call_Data_Map_Internal_ordMapIter(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
// TAST (Let): __local_var_1_0 -> gopurs_runtime.Value
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd_0, "Eq0"), gopurs_runtime.Value{})
_ = __local_var_1_0
return gopurs_runtime.Func(func(dictOrd1_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_3_2 -> gopurs_runtime.Value
__local_var_3_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd1_2, "Eq0"), gopurs_runtime.Value{})
_ = __local_var_3_2
var go__go_4_3_12 gopurs_runtime.Value
go__go_4_3_12 = gopurs_runtime.Func(func(a_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var a_5_loop gopurs_runtime.Value = a_5_loop_val
var b_6_loop gopurs_runtime.Value = b_6_loop_val
go__go_4_3_12:
for {
if false { continue go__go_4_3_12 }
var a_5 gopurs_runtime.Value = a_5_loop
_ = a_5
var b_6 gopurs_runtime.Value = b_6_loop
_ = b_6
// TAST (Let): v_7_4 -> *Constructor_Data_Map_Internal_IterNext
v_7_4 := gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_IterNext](gopurs_runtime.Apply(Get_Data_Map_Internal_stepAsc(), a_5))
_ = v_7_4
var __t7 bool
{
if (v_7_4 != nil) {
// TAST (Let): v2_8_5 -> *Constructor_Data_Map_Internal_IterNext
v2_8_5 := gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_IterNext](gopurs_runtime.Apply(Get_Data_Map_Internal_stepAsc(), b_6))
_ = v2_8_5
var __t6 bool
{
if ((v2_8_5 != nil)) && (((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_0, "eq"), (v_7_4).V0, (v2_8_5).V0).IntVal) != (0)) && ((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_3_2, "eq"), (v_7_4).V1, (v2_8_5).V1).IntVal) != (0))) {
a_5_loop = (v_7_4).V2
b_6_loop = (v2_8_5).V2
continue go__go_4_3_12
__t6 = (gopurs_runtime.Value{}.IntVal) != (0)
goto end_branch_6
} else {

}
}
{
__t6 = false
}
end_branch_6:
__t7 = __t6
goto end_branch_7
} else {

}
}
{
if (v_7_4 == nil) {
__t7 = true
goto end_branch_7
} else {

}
}
{
__t7 = (func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal) != (0)
}
end_branch_7:
return gopurs_runtime.Bool(__t7)
}
}()
})
})
// TAST (Let): eqMapIter2_3_1 -> gopurs_runtime.Value
eqMapIter2_3_1 := gopurs_runtime.RecordDict1("eq", go__go_4_3_12)
_ = eqMapIter2_3_1
var go__go_4_8_13 gopurs_runtime.Value
go__go_4_8_13 = gopurs_runtime.Func(func(a_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var a_5_loop gopurs_runtime.Value = a_5_loop_val
var b_6_loop gopurs_runtime.Value = b_6_loop_val
go__go_4_8_13:
for {
if false { continue go__go_4_8_13 }
var a_5 gopurs_runtime.Value = a_5_loop
_ = a_5
var b_6 gopurs_runtime.Value = b_6_loop
_ = b_6
// TAST (Let): v_7_9 -> *Constructor_Data_Map_Internal_IterNext
v_7_9 := gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_IterNext](gopurs_runtime.Apply(Get_Data_Map_Internal_stepAsc(), b_6))
_ = v_7_9
// TAST (Let): v1_8_10 -> *Constructor_Data_Map_Internal_IterNext
v1_8_10 := gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_IterNext](gopurs_runtime.Apply(Get_Data_Map_Internal_stepAsc(), a_5))
_ = v1_8_10
var __t17 uint32
{
if (v1_8_10 != nil) {
var __t15 uint32
{
if (v_7_9 != nil) {
// TAST (Let): v3_9_11 -> gopurs_runtime.Value
v3_9_11 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictOrd_0, "compare"), (v1_8_10).V0, (v_7_9).V0)
_ = v3_9_11
var __t14 uint32
{
if (uint32(v3_9_11.IntVal) == 902936544) {
// TAST (Let): v4_10_12 -> gopurs_runtime.Value
v4_10_12 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictOrd1_2, "compare"), (v1_8_10).V1, (v_7_9).V1)
_ = v4_10_12
var __t13 uint32
{
if (uint32(v4_10_12.IntVal) == 902936544) {
a_5_loop = (v1_8_10).V2
b_6_loop = (v_7_9).V2
continue go__go_4_8_13
__t13 = uint32(gopurs_runtime.Value{}.IntVal)
goto end_branch_13
} else {

}
}
{
__t13 = uint32(v4_10_12.IntVal)
}
end_branch_13:
__t14 = __t13
goto end_branch_14
} else {

}
}
{
__t14 = uint32(v3_9_11.IntVal)
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
__t15 = uint32(func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal)
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
__t17 = uint32(func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal)
}
end_branch_17:
return gopurs_runtime.Value{Type: 9, IntVal: int64(__t17), UnsafePtr: nil}
}
}()
})
})
return gopurs_runtime.RecordDict2("Eq0", "compare", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return eqMapIter2_3_1
}), go__go_4_8_13)
})
}

func Call_Data_Map_Internal_toUnfoldable(dictUnfoldable_0_loop *Constructor_Data_Unfoldable_Unfoldable) gopurs_runtime.Value {
var dictUnfoldable_0 *Constructor_Data_Unfoldable_Unfoldable = dictUnfoldable_0_loop
_ = dictUnfoldable_0
// TAST (Let): __local_var_1_0 -> gopurs_runtime.Value
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.Box(dictUnfoldable_0.V1), Get_Data_Map_Internal_stepUnfoldr())
_ = __local_var_1_0
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_1_0, gopurs_runtime.Value{Type: 9, IntVal: 2861335956, UnsafePtr: unsafe.Pointer(&Constructor_Data_Map_Internal_IterNode{1, gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](x_2), gopurs_runtime.Value{Type: 9, IntVal: 2509360378, UnsafePtr: unsafe.Pointer(nil)}})})
})
}

func Call_Data_Map_Internal_showMap(dictShow_0_loop gopurs_runtime.Value, dictShow1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictShow_0 gopurs_runtime.Value = dictShow_0_loop
_ = dictShow_0
var dictShow1_1 gopurs_runtime.Value = dictShow1_1_loop
_ = dictShow1_1
// TAST (Let): showArray_2_0 -> *Constructor_Data_Show_Show
showArray_2_0 := &Constructor_Data_Show_Show{1, gopurs_runtime.Apply(Get_Data_Show_showArrayImpl(), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str((((("(Tuple ") + (gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictShow_0, "show"), (*Constructor_Data_Tuple_Tuple)(v_2.UnsafePtr).V0).StrVal())) + (" ")) + (gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictShow1_1, "show"), (*Constructor_Data_Tuple_Tuple)(v_2.UnsafePtr).V1).StrVal())) + (")"))
}))}
_ = showArray_2_0
return gopurs_runtime.RecordDict1("show", gopurs_runtime.Func(func(as_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str((("(fromFoldable ") + (gopurs_runtime.Apply(gopurs_runtime.Box(showArray_2_0.V0), func() gopurs_runtime.Value {
					arr := func() []*Constructor_Data_Tuple_Tuple {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_Unfoldable_unfoldableArray(), "unfoldr"), Get_Data_Map_Internal_stepUnfoldr(), gopurs_runtime.Value{Type: 9, IntVal: 2861335956, UnsafePtr: unsafe.Pointer(&Constructor_Data_Map_Internal_IterNode{1, gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](as_3), gopurs_runtime.Value{Type: 9, IntVal: 2509360378, UnsafePtr: unsafe.Pointer(nil)}})}).UnsafePtr)
					unboxed := make([]*Constructor_Data_Tuple_Tuple, len(arr))
					for i, v := range arr { unboxed[i] = gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple](v) }
					return unboxed
				}()
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v)} }
					return gopurs_runtime.Array(boxed)
				}()).StrVal())) + (")"))
}))
}

func Call_Data_Map_Internal_isSubmap(dictOrd_0_loop *Constructor_Data_Ord_Ord, dictEq_1_loop *Constructor_Data_Eq_Eq) gopurs_runtime.Value {
var dictOrd_0 *Constructor_Data_Ord_Ord = dictOrd_0_loop
_ = dictOrd_0
var dictEq_1 *Constructor_Data_Eq_Eq = dictEq_1_loop
_ = dictEq_1
var go__go_2_0_14 gopurs_runtime.Value
_ = go__go_2_0_14
go__go_2_0_14 = gopurs_runtime.Func(func(m1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m2_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t8 bool
{
if (m1_3.Type == 9 && m1_3.IntVal == 324739070 && m1_3.UnsafePtr == nil) {
__t8 = true
goto end_branch_8
} else {

}
}
{
if (m1_3.Type == 9 && m1_3.IntVal == 324739070 && m1_3.UnsafePtr != nil) {
// TAST (Let): __local_var_5_1 -> gopurs_runtime.Value
__local_var_5_1 := (*Constructor_Data_Map_Internal_Node)(m1_3.UnsafePtr).V2
_ = __local_var_5_1
var go__go_6_2_15 gopurs_runtime.Value
go__go_6_2_15 = gopurs_runtime.Func(func(v_7_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_7_loop gopurs_runtime.Value = v_7_loop_val
go__go_6_2_15:
for {
if false { continue go__go_6_2_15 }
var v_7 gopurs_runtime.Value = v_7_loop
_ = v_7
var __t5 *Constructor_Data_Maybe_Just
{
if (v_7.Type == 9 && v_7.IntVal == 324739070 && v_7.UnsafePtr == nil) {
__t5 = (*Constructor_Data_Maybe_Just)(nil)
goto end_branch_5
} else {

}
}
{
if (v_7.Type == 9 && v_7.IntVal == 324739070 && v_7.UnsafePtr != nil) {
// TAST (Let): v1_8_3 -> gopurs_runtime.Value
v1_8_3 := gopurs_runtime.Apply2(gopurs_runtime.Box(dictOrd_0.V1), __local_var_5_1, (*Constructor_Data_Map_Internal_Node)(v_7.UnsafePtr).V2)
_ = v1_8_3
var __t4 *Constructor_Data_Maybe_Just
{
if (uint32(v1_8_3.IntVal) == 1527465420) {
v_7_loop = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v_7.UnsafePtr).V4)}
continue go__go_6_2_15
__t4 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Value{})
goto end_branch_4
} else {

}
}
{
if (uint32(v1_8_3.IntVal) == 380165415) {
v_7_loop = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v_7.UnsafePtr).V5)}
continue go__go_6_2_15
__t4 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Value{})
goto end_branch_4
} else {

}
}
{
if (uint32(v1_8_3.IntVal) == 902936544) {
__t4 = &Constructor_Data_Maybe_Just{1, (*Constructor_Data_Map_Internal_Node)(v_7.UnsafePtr).V3}
goto end_branch_4
} else {

}
}
{
__t4 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_4:
__t5 = __t4
goto end_branch_5
} else {

}
}
{
__t5 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_5:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t5)}
}
}()
})
// TAST (Let): v1_7_6 -> gopurs_runtime.Value
v1_7_6 := gopurs_runtime.Apply(go__go_6_2_15, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](m2_4))})
_ = v1_7_6
var __t7 bool
{
if (v1_7_6.Type == 9 && v1_7_6.IntVal == 930809136 && v1_7_6.UnsafePtr == nil) {
__t7 = false
goto end_branch_7
} else {

}
}
{
if (v1_7_6.Type == 9 && v1_7_6.IntVal == 930809136 && v1_7_6.UnsafePtr != nil) {
__t7 = ((gopurs_runtime.Apply2(gopurs_runtime.Box(dictEq_1.V0), (*Constructor_Data_Map_Internal_Node)(m1_3.UnsafePtr).V3, (*Constructor_Data_Maybe_Just)(v1_7_6.UnsafePtr).V0).IntVal) != (0)) && (((gopurs_runtime.Apply2(go__go_2_0_14, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(m1_3.UnsafePtr).V4)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](m2_4))}).IntVal) != (0)) && ((gopurs_runtime.Apply2(go__go_2_0_14, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(m1_3.UnsafePtr).V5)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](m2_4))}).IntVal) != (0)))
goto end_branch_7
} else {

}
}
{
__t7 = (func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal) != (0)
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
return go__go_2_0_14
}

func Call_Data_Map_Internal_isEmpty(v_0_loop *Constructor_Data_Map_Internal_Node) bool {
var v_0 *Constructor_Data_Map_Internal_Node = v_0_loop
_ = v_0
var __t0 bool
{
if (v_0 == nil) {
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

func Call_Data_Map_Internal_intersectionWith(dictOrd_0_loop *Constructor_Data_Ord_Ord) gopurs_runtime.Value {
var dictOrd_0 *Constructor_Data_Ord_Ord = dictOrd_0_loop
_ = dictOrd_0
// TAST (Let): compare_1_0 -> gopurs_runtime.Value
compare_1_0 := gopurs_runtime.Box(dictOrd_0.V1)
_ = compare_1_0
return gopurs_runtime.Func(func(app_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m2_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeIntersectionWith(), compare_1_0, app_2, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](m1_3))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](m2_4))})))}
})
})
})
}

func Call_Data_Map_Internal_intersection(dictOrd_0_loop *Constructor_Data_Ord_Ord) gopurs_runtime.Value {
var dictOrd_0 *Constructor_Data_Ord_Ord = dictOrd_0_loop
_ = dictOrd_0
// TAST (Let): compare_1_0 -> gopurs_runtime.Value
compare_1_0 := gopurs_runtime.Box(dictOrd_0.V1)
_ = compare_1_0
return gopurs_runtime.Func(func(m1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m2_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeIntersectionWith(), compare_1_0, Get_Data_Function_go__const(), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](m1_2))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](m2_3))})))}
})
})
}

func Call_Data_Map_Internal_insertWith(dictOrd_0_loop *Constructor_Data_Ord_Ord, app_1_loop gopurs_runtime.Value, k_2_loop gopurs_runtime.Value, v_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 *Constructor_Data_Ord_Ord = dictOrd_0_loop
_ = dictOrd_0
var app_1 gopurs_runtime.Value = app_1_loop
_ = app_1
var k_2 gopurs_runtime.Value = k_2_loop
_ = k_2
var v_3 gopurs_runtime.Value = v_3_loop
_ = v_3
var go__go_4_0_16 gopurs_runtime.Value
_ = go__go_4_0_16
go__go_4_0_16 = gopurs_runtime.Func(func(v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
var __t3 *Constructor_Data_Map_Internal_Node
{
if (v1_5.Type == 9 && v1_5.IntVal == 324739070 && v1_5.UnsafePtr == nil) {
__t3 = &Constructor_Data_Map_Internal_Node{1, 1, 1, k_2, v_3, (*Constructor_Data_Map_Internal_Node)(nil), (*Constructor_Data_Map_Internal_Node)(nil)}
goto end_branch_3
} else {

}
}
{
if (v1_5.Type == 9 && v1_5.IntVal == 324739070 && v1_5.UnsafePtr != nil) {
// TAST (Let): v2_6_1 -> gopurs_runtime.Value
v2_6_1 := gopurs_runtime.Apply2(gopurs_runtime.Box(dictOrd_0.V1), k_2, (*Constructor_Data_Map_Internal_Node)(v1_5.UnsafePtr).V2)
_ = v2_6_1
var __t2 *Constructor_Data_Map_Internal_Node
{
if (uint32(v2_6_1.IntVal) == 1527465420) {
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeBalancedNode(), (*Constructor_Data_Map_Internal_Node)(v1_5.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node)(v1_5.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.Apply(go__go_4_0_16, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v1_5.UnsafePtr).V4)})))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v1_5.UnsafePtr).V5)}))
goto end_branch_2
} else {

}
}
{
if (uint32(v2_6_1.IntVal) == 380165415) {
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeBalancedNode(), (*Constructor_Data_Map_Internal_Node)(v1_5.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node)(v1_5.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v1_5.UnsafePtr).V4)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.Apply(go__go_4_0_16, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v1_5.UnsafePtr).V5)})))}))
goto end_branch_2
} else {

}
}
{
if (uint32(v2_6_1.IntVal) == 902936544) {
__t2 = &Constructor_Data_Map_Internal_Node{1, (*Constructor_Data_Map_Internal_Node)(v1_5.UnsafePtr).V0, (*Constructor_Data_Map_Internal_Node)(v1_5.UnsafePtr).V1, k_2, gopurs_runtime.Apply2(app_1, (*Constructor_Data_Map_Internal_Node)(v1_5.UnsafePtr).V3, v_3), (*Constructor_Data_Map_Internal_Node)(v1_5.UnsafePtr).V4, (*Constructor_Data_Map_Internal_Node)(v1_5.UnsafePtr).V5}
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_2:
__t3 = __t2
goto end_branch_3
} else {

}
}
{
__t3 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_3:
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(__t3)}
})
return go__go_4_0_16
}

func Call_Data_Map_Internal_insert(dictOrd_0_loop *Constructor_Data_Ord_Ord, k_1_loop gopurs_runtime.Value, v_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 *Constructor_Data_Ord_Ord = dictOrd_0_loop
_ = dictOrd_0
var k_1 gopurs_runtime.Value = k_1_loop
_ = k_1
var v_2 gopurs_runtime.Value = v_2_loop
_ = v_2
var go__go_3_0_17 gopurs_runtime.Value
_ = go__go_3_0_17
go__go_3_0_17 = gopurs_runtime.Func(func(v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t3 *Constructor_Data_Map_Internal_Node
{
if (v1_4.Type == 9 && v1_4.IntVal == 324739070 && v1_4.UnsafePtr == nil) {
__t3 = &Constructor_Data_Map_Internal_Node{1, 1, 1, k_1, v_2, (*Constructor_Data_Map_Internal_Node)(nil), (*Constructor_Data_Map_Internal_Node)(nil)}
goto end_branch_3
} else {

}
}
{
if (v1_4.Type == 9 && v1_4.IntVal == 324739070 && v1_4.UnsafePtr != nil) {
// TAST (Let): v2_5_1 -> gopurs_runtime.Value
v2_5_1 := gopurs_runtime.Apply2(gopurs_runtime.Box(dictOrd_0.V1), k_1, (*Constructor_Data_Map_Internal_Node)(v1_4.UnsafePtr).V2)
_ = v2_5_1
var __t2 *Constructor_Data_Map_Internal_Node
{
if (uint32(v2_5_1.IntVal) == 1527465420) {
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeBalancedNode(), (*Constructor_Data_Map_Internal_Node)(v1_4.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node)(v1_4.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.Apply(go__go_3_0_17, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v1_4.UnsafePtr).V4)})))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v1_4.UnsafePtr).V5)}))
goto end_branch_2
} else {

}
}
{
if (uint32(v2_5_1.IntVal) == 380165415) {
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeBalancedNode(), (*Constructor_Data_Map_Internal_Node)(v1_4.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node)(v1_4.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v1_4.UnsafePtr).V4)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.Apply(go__go_3_0_17, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v1_4.UnsafePtr).V5)})))}))
goto end_branch_2
} else {

}
}
{
if (uint32(v2_5_1.IntVal) == 902936544) {
__t2 = &Constructor_Data_Map_Internal_Node{1, (*Constructor_Data_Map_Internal_Node)(v1_4.UnsafePtr).V0, (*Constructor_Data_Map_Internal_Node)(v1_4.UnsafePtr).V1, k_1, v_2, (*Constructor_Data_Map_Internal_Node)(v1_4.UnsafePtr).V4, (*Constructor_Data_Map_Internal_Node)(v1_4.UnsafePtr).V5}
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_2:
__t3 = __t2
goto end_branch_3
} else {

}
}
{
__t3 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_3:
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(__t3)}
})
return go__go_3_0_17
}

func Call_Data_Map_Internal_foldSubmapBy(dictOrd_0_loop *Constructor_Data_Ord_Ord, appendFn_1_loop gopurs_runtime.Value, memptyValue_2_loop gopurs_runtime.Value, kmin_3_loop *Constructor_Data_Maybe_Just, kmax_4_loop *Constructor_Data_Maybe_Just, f_5_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 *Constructor_Data_Ord_Ord = dictOrd_0_loop
_ = dictOrd_0
var appendFn_1 gopurs_runtime.Value = appendFn_1_loop
_ = appendFn_1
var memptyValue_2 gopurs_runtime.Value = memptyValue_2_loop
_ = memptyValue_2
var kmin_3 *Constructor_Data_Maybe_Just = kmin_3_loop
_ = kmin_3
var kmax_4 *Constructor_Data_Maybe_Just = kmax_4_loop
_ = kmax_4
var f_5 gopurs_runtime.Value = f_5_loop
_ = f_5
var __t4 gopurs_runtime.Value
{
if (kmin_3 != nil) {
// TAST (Let): __local_var_6_1 -> gopurs_runtime.Value
__local_var_6_1 := (kmin_3).V0
_ = __local_var_6_1
__t4 = gopurs_runtime.Func(func(k_7 gopurs_runtime.Value) gopurs_runtime.Value {
var __t3 bool
{
var __t_tag_2 gopurs_runtime.Value = gopurs_runtime.Apply2(gopurs_runtime.Box(dictOrd_0.V1), k_7, __local_var_6_1)
if (uint32(__t_tag_2.IntVal) == 1527465420) {
__t3 = true
goto end_branch_3
} else {

}
}
{
__t3 = false
}
end_branch_3:
return gopurs_runtime.Bool(__t3)
})
goto end_branch_4
} else {

}
}
{
if (kmin_3 == nil) {
__t4 = gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(false)
})
goto end_branch_4
} else {

}
}
{
__t4 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_4:
// TAST (Let): tooSmall_6_0 -> gopurs_runtime.Value
tooSmall_6_0 := __t4
_ = tooSmall_6_0
var __t9 gopurs_runtime.Value
{
if (kmax_4 != nil) {
// TAST (Let): __local_var_7_6 -> gopurs_runtime.Value
__local_var_7_6 := (kmax_4).V0
_ = __local_var_7_6
__t9 = gopurs_runtime.Func(func(k_8 gopurs_runtime.Value) gopurs_runtime.Value {
var __t8 bool
{
var __t_tag_7 gopurs_runtime.Value = gopurs_runtime.Apply2(gopurs_runtime.Box(dictOrd_0.V1), k_8, __local_var_7_6)
if (uint32(__t_tag_7.IntVal) == 380165415) {
__t8 = true
goto end_branch_8
} else {

}
}
{
__t8 = false
}
end_branch_8:
return gopurs_runtime.Bool(__t8)
})
goto end_branch_9
} else {

}
}
{
if (kmax_4 == nil) {
__t9 = gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(false)
})
goto end_branch_9
} else {

}
}
{
__t9 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_9:
// TAST (Let): tooLarge_7_5 -> gopurs_runtime.Value
tooLarge_7_5 := __t9
_ = tooLarge_7_5
var __t26 gopurs_runtime.Value
{
if (kmin_3 != nil) {
var __t21 gopurs_runtime.Value
{
if (kmax_4 != nil) {
// TAST (Let): __local_var_8_11 -> gopurs_runtime.Value
__local_var_8_11 := (kmax_4).V0
_ = __local_var_8_11
// TAST (Let): __local_var_9_12 -> gopurs_runtime.Value
__local_var_9_12 := (kmin_3).V0
_ = __local_var_9_12
__t21 = gopurs_runtime.Func(func(k_10 gopurs_runtime.Value) gopurs_runtime.Value {
var __t14 bool
{
var __t_tag_13 gopurs_runtime.Value = gopurs_runtime.Apply2(gopurs_runtime.Box(dictOrd_0.V1), __local_var_9_12, k_10)
if (uint32(__t_tag_13.IntVal) == 380165415) {
__t14 = false
goto end_branch_14
} else {

}
}
{
__t14 = true
}
end_branch_14:
var __t_and_17 bool = false
if __t14 {

var __t16 bool
{
var __t_tag_15 gopurs_runtime.Value = gopurs_runtime.Apply2(gopurs_runtime.Box(dictOrd_0.V1), k_10, __local_var_8_11)
if (uint32(__t_tag_15.IntVal) == 380165415) {
__t16 = false
goto end_branch_16
} else {

}
}
{
__t16 = true
}
end_branch_16:
__t_and_17 = __t16
}
return gopurs_runtime.Bool(__t_and_17)
})
goto end_branch_21
} else {

}
}
{
if (kmax_4 == nil) {
// TAST (Let): __local_var_8_18 -> gopurs_runtime.Value
__local_var_8_18 := (kmin_3).V0
_ = __local_var_8_18
__t21 = gopurs_runtime.Func(func(k_9 gopurs_runtime.Value) gopurs_runtime.Value {
var __t20 bool
{
var __t_tag_19 gopurs_runtime.Value = gopurs_runtime.Apply2(gopurs_runtime.Box(dictOrd_0.V1), __local_var_8_18, k_9)
if (uint32(__t_tag_19.IntVal) == 380165415) {
__t20 = false
goto end_branch_20
} else {

}
}
{
__t20 = true
}
end_branch_20:
return gopurs_runtime.Bool(__t20)
})
goto end_branch_21
} else {

}
}
{
__t21 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_21:
__t26 = __t21
goto end_branch_26
} else {

}
}
{
if (kmin_3 == nil) {
var __t25 gopurs_runtime.Value
{
if (kmax_4 != nil) {
// TAST (Let): __local_var_8_22 -> gopurs_runtime.Value
__local_var_8_22 := (kmax_4).V0
_ = __local_var_8_22
__t25 = gopurs_runtime.Func(func(k_9 gopurs_runtime.Value) gopurs_runtime.Value {
var __t24 bool
{
var __t_tag_23 gopurs_runtime.Value = gopurs_runtime.Apply2(gopurs_runtime.Box(dictOrd_0.V1), k_9, __local_var_8_22)
if (uint32(__t_tag_23.IntVal) == 380165415) {
__t24 = false
goto end_branch_24
} else {

}
}
{
__t24 = true
}
end_branch_24:
return gopurs_runtime.Bool(__t24)
})
goto end_branch_25
} else {

}
}
{
if (kmax_4 == nil) {
__t25 = gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(true)
})
goto end_branch_25
} else {

}
}
{
__t25 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_25:
__t26 = __t25
goto end_branch_26
} else {

}
}
{
__t26 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_26:
// TAST (Let): inBounds_8_10 -> gopurs_runtime.Value
inBounds_8_10 := __t26
_ = inBounds_8_10
var go__go_9_27_51 gopurs_runtime.Value
_ = go__go_9_27_51
go__go_9_27_51 = gopurs_runtime.Func(func(v_10 gopurs_runtime.Value) gopurs_runtime.Value {
var __t31 gopurs_runtime.Value
{
if (v_10.Type == 9 && v_10.IntVal == 324739070 && v_10.UnsafePtr == nil) {
__t31 = memptyValue_2
goto end_branch_31
} else {

}
}
{
if (v_10.Type == 9 && v_10.IntVal == 324739070 && v_10.UnsafePtr != nil) {
var __t28 gopurs_runtime.Value
{
if (gopurs_runtime.Apply(tooSmall_6_0, (*Constructor_Data_Map_Internal_Node)(v_10.UnsafePtr).V2).IntVal) != (0) {
__t28 = memptyValue_2
goto end_branch_28
} else {

}
}
{
__t28 = gopurs_runtime.Apply(go__go_9_27_51, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v_10.UnsafePtr).V4)})
}
end_branch_28:
var __t29 gopurs_runtime.Value
{
if (gopurs_runtime.Apply(inBounds_8_10, (*Constructor_Data_Map_Internal_Node)(v_10.UnsafePtr).V2).IntVal) != (0) {
__t29 = gopurs_runtime.Apply2(f_5, (*Constructor_Data_Map_Internal_Node)(v_10.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node)(v_10.UnsafePtr).V3)
goto end_branch_29
} else {

}
}
{
__t29 = memptyValue_2
}
end_branch_29:
var __t30 gopurs_runtime.Value
{
if (gopurs_runtime.Apply(tooLarge_7_5, (*Constructor_Data_Map_Internal_Node)(v_10.UnsafePtr).V2).IntVal) != (0) {
__t30 = memptyValue_2
goto end_branch_30
} else {

}
}
{
__t30 = gopurs_runtime.Apply(go__go_9_27_51, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v_10.UnsafePtr).V5)})
}
end_branch_30:
__t31 = gopurs_runtime.Apply2(appendFn_1, gopurs_runtime.Apply2(appendFn_1, __t28, __t29), __t30)
goto end_branch_31
} else {

}
}
{
__t31 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_31:
return __t31
})
return go__go_9_27_51
}

func Call_Data_Map_Internal_foldSubmap(dictOrd_0_loop *Constructor_Data_Ord_Ord, dictMonoid_1_loop *Constructor_Data_Monoid_Monoid) gopurs_runtime.Value {
var dictOrd_0 *Constructor_Data_Ord_Ord = dictOrd_0_loop
_ = dictOrd_0
var dictMonoid_1 *Constructor_Data_Monoid_Monoid = dictMonoid_1_loop
_ = dictMonoid_1
// TAST (Let): __local_var_2_0 -> gopurs_runtime.Value
__local_var_2_0 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.Box(dictMonoid_1.V0), gopurs_runtime.Value{}), "append")
_ = __local_var_2_0
// TAST (Let): __local_var_3_1 -> gopurs_runtime.Value
__local_var_3_1 := gopurs_runtime.Box(dictMonoid_1.V1)
_ = __local_var_3_1
return gopurs_runtime.Func(func(kmin_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(kmax_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_6 gopurs_runtime.Value) gopurs_runtime.Value {
var __t6 gopurs_runtime.Value
{
if (kmin_4.Type == 9 && kmin_4.IntVal == 930809136 && kmin_4.UnsafePtr != nil) {
// TAST (Let): __local_var_7_3 -> gopurs_runtime.Value
__local_var_7_3 := (*Constructor_Data_Maybe_Just)(kmin_4.UnsafePtr).V0
_ = __local_var_7_3
__t6 = gopurs_runtime.Func(func(k_8 gopurs_runtime.Value) gopurs_runtime.Value {
var __t5 bool
{
var __t_tag_4 gopurs_runtime.Value = gopurs_runtime.Apply2(gopurs_runtime.Box(dictOrd_0.V1), k_8, __local_var_7_3)
if (uint32(__t_tag_4.IntVal) == 1527465420) {
__t5 = true
goto end_branch_5
} else {

}
}
{
__t5 = false
}
end_branch_5:
return gopurs_runtime.Bool(__t5)
})
goto end_branch_6
} else {

}
}
{
if (kmin_4.Type == 9 && kmin_4.IntVal == 930809136 && kmin_4.UnsafePtr == nil) {
__t6 = gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(false)
})
goto end_branch_6
} else {

}
}
{
__t6 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_6:
// TAST (Let): tooSmall_7_2 -> gopurs_runtime.Value
tooSmall_7_2 := __t6
_ = tooSmall_7_2
var __t11 gopurs_runtime.Value
{
if (kmax_5.Type == 9 && kmax_5.IntVal == 930809136 && kmax_5.UnsafePtr != nil) {
// TAST (Let): __local_var_8_8 -> gopurs_runtime.Value
__local_var_8_8 := (*Constructor_Data_Maybe_Just)(kmax_5.UnsafePtr).V0
_ = __local_var_8_8
__t11 = gopurs_runtime.Func(func(k_9 gopurs_runtime.Value) gopurs_runtime.Value {
var __t10 bool
{
var __t_tag_9 gopurs_runtime.Value = gopurs_runtime.Apply2(gopurs_runtime.Box(dictOrd_0.V1), k_9, __local_var_8_8)
if (uint32(__t_tag_9.IntVal) == 380165415) {
__t10 = true
goto end_branch_10
} else {

}
}
{
__t10 = false
}
end_branch_10:
return gopurs_runtime.Bool(__t10)
})
goto end_branch_11
} else {

}
}
{
if (kmax_5.Type == 9 && kmax_5.IntVal == 930809136 && kmax_5.UnsafePtr == nil) {
__t11 = gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(false)
})
goto end_branch_11
} else {

}
}
{
__t11 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_11:
// TAST (Let): tooLarge_8_7 -> gopurs_runtime.Value
tooLarge_8_7 := __t11
_ = tooLarge_8_7
var __t28 gopurs_runtime.Value
{
if (kmin_4.Type == 9 && kmin_4.IntVal == 930809136 && kmin_4.UnsafePtr != nil) {
var __t23 gopurs_runtime.Value
{
if (kmax_5.Type == 9 && kmax_5.IntVal == 930809136 && kmax_5.UnsafePtr != nil) {
// TAST (Let): __local_var_9_13 -> gopurs_runtime.Value
__local_var_9_13 := (*Constructor_Data_Maybe_Just)(kmax_5.UnsafePtr).V0
_ = __local_var_9_13
// TAST (Let): __local_var_10_14 -> gopurs_runtime.Value
__local_var_10_14 := (*Constructor_Data_Maybe_Just)(kmin_4.UnsafePtr).V0
_ = __local_var_10_14
__t23 = gopurs_runtime.Func(func(k_11 gopurs_runtime.Value) gopurs_runtime.Value {
var __t16 bool
{
var __t_tag_15 gopurs_runtime.Value = gopurs_runtime.Apply2(gopurs_runtime.Box(dictOrd_0.V1), __local_var_10_14, k_11)
if (uint32(__t_tag_15.IntVal) == 380165415) {
__t16 = false
goto end_branch_16
} else {

}
}
{
__t16 = true
}
end_branch_16:
var __t_and_19 bool = false
if __t16 {

var __t18 bool
{
var __t_tag_17 gopurs_runtime.Value = gopurs_runtime.Apply2(gopurs_runtime.Box(dictOrd_0.V1), k_11, __local_var_9_13)
if (uint32(__t_tag_17.IntVal) == 380165415) {
__t18 = false
goto end_branch_18
} else {

}
}
{
__t18 = true
}
end_branch_18:
__t_and_19 = __t18
}
return gopurs_runtime.Bool(__t_and_19)
})
goto end_branch_23
} else {

}
}
{
if (kmax_5.Type == 9 && kmax_5.IntVal == 930809136 && kmax_5.UnsafePtr == nil) {
// TAST (Let): __local_var_9_20 -> gopurs_runtime.Value
__local_var_9_20 := (*Constructor_Data_Maybe_Just)(kmin_4.UnsafePtr).V0
_ = __local_var_9_20
__t23 = gopurs_runtime.Func(func(k_10 gopurs_runtime.Value) gopurs_runtime.Value {
var __t22 bool
{
var __t_tag_21 gopurs_runtime.Value = gopurs_runtime.Apply2(gopurs_runtime.Box(dictOrd_0.V1), __local_var_9_20, k_10)
if (uint32(__t_tag_21.IntVal) == 380165415) {
__t22 = false
goto end_branch_22
} else {

}
}
{
__t22 = true
}
end_branch_22:
return gopurs_runtime.Bool(__t22)
})
goto end_branch_23
} else {

}
}
{
__t23 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_23:
__t28 = __t23
goto end_branch_28
} else {

}
}
{
if (kmin_4.Type == 9 && kmin_4.IntVal == 930809136 && kmin_4.UnsafePtr == nil) {
var __t27 gopurs_runtime.Value
{
if (kmax_5.Type == 9 && kmax_5.IntVal == 930809136 && kmax_5.UnsafePtr != nil) {
// TAST (Let): __local_var_9_24 -> gopurs_runtime.Value
__local_var_9_24 := (*Constructor_Data_Maybe_Just)(kmax_5.UnsafePtr).V0
_ = __local_var_9_24
__t27 = gopurs_runtime.Func(func(k_10 gopurs_runtime.Value) gopurs_runtime.Value {
var __t26 bool
{
var __t_tag_25 gopurs_runtime.Value = gopurs_runtime.Apply2(gopurs_runtime.Box(dictOrd_0.V1), k_10, __local_var_9_24)
if (uint32(__t_tag_25.IntVal) == 380165415) {
__t26 = false
goto end_branch_26
} else {

}
}
{
__t26 = true
}
end_branch_26:
return gopurs_runtime.Bool(__t26)
})
goto end_branch_27
} else {

}
}
{
if (kmax_5.Type == 9 && kmax_5.IntVal == 930809136 && kmax_5.UnsafePtr == nil) {
__t27 = gopurs_runtime.Func(func(v_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(true)
})
goto end_branch_27
} else {

}
}
{
__t27 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_27:
__t28 = __t27
goto end_branch_28
} else {

}
}
{
__t28 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_28:
// TAST (Let): inBounds_9_12 -> gopurs_runtime.Value
inBounds_9_12 := __t28
_ = inBounds_9_12
var go__go_10_29_52 gopurs_runtime.Value
_ = go__go_10_29_52
go__go_10_29_52 = gopurs_runtime.Func(func(v_11 gopurs_runtime.Value) gopurs_runtime.Value {
var __t33 gopurs_runtime.Value
{
if (v_11.Type == 9 && v_11.IntVal == 324739070 && v_11.UnsafePtr == nil) {
__t33 = __local_var_3_1
goto end_branch_33
} else {

}
}
{
if (v_11.Type == 9 && v_11.IntVal == 324739070 && v_11.UnsafePtr != nil) {
var __t30 gopurs_runtime.Value
{
if (gopurs_runtime.Apply(tooSmall_7_2, (*Constructor_Data_Map_Internal_Node)(v_11.UnsafePtr).V2).IntVal) != (0) {
__t30 = __local_var_3_1
goto end_branch_30
} else {

}
}
{
__t30 = gopurs_runtime.Apply(go__go_10_29_52, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v_11.UnsafePtr).V4)})
}
end_branch_30:
var __t31 gopurs_runtime.Value
{
if (gopurs_runtime.Apply(inBounds_9_12, (*Constructor_Data_Map_Internal_Node)(v_11.UnsafePtr).V2).IntVal) != (0) {
__t31 = gopurs_runtime.Apply2(f_6, (*Constructor_Data_Map_Internal_Node)(v_11.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node)(v_11.UnsafePtr).V3)
goto end_branch_31
} else {

}
}
{
__t31 = __local_var_3_1
}
end_branch_31:
var __t32 gopurs_runtime.Value
{
if (gopurs_runtime.Apply(tooLarge_8_7, (*Constructor_Data_Map_Internal_Node)(v_11.UnsafePtr).V2).IntVal) != (0) {
__t32 = __local_var_3_1
goto end_branch_32
} else {

}
}
{
__t32 = gopurs_runtime.Apply(go__go_10_29_52, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v_11.UnsafePtr).V5)})
}
end_branch_32:
__t33 = gopurs_runtime.Apply2(__local_var_2_0, gopurs_runtime.Apply2(__local_var_2_0, __t30, __t31), __t32)
goto end_branch_33
} else {

}
}
{
__t33 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_33:
return __t33
})
return go__go_10_29_52
})
})
})
}

func Call_Data_Map_Internal_findMin(v_0_loop *Constructor_Data_Map_Internal_Node) *Constructor_Data_Maybe_Just {
findMin:
for {
if false { continue findMin }
var v_0 *Constructor_Data_Map_Internal_Node = v_0_loop
_ = v_0
var __t2 *Constructor_Data_Maybe_Just
{
if (v_0 == nil) {
__t2 = (*Constructor_Data_Maybe_Just)(nil)
goto end_branch_2
} else {

}
}
{
if (v_0 != nil) {
var __t1 *Constructor_Data_Maybe_Just
{
var __t_tag_0 *Constructor_Data_Map_Internal_Node = (v_0).V4
if (__t_tag_0 == nil) {
__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, gopurs_runtime.RecordDict2("key", "value", (v_0).V2, (v_0).V3)})})
goto end_branch_1
} else {

}
}
{
v_0_loop = (v_0).V4
continue findMin
__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Value{})
}
end_branch_1:
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t1)})
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_2:
return gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t2)})
}
}

func Call_Data_Map_Internal_lookupGT(dictOrd_0_loop *Constructor_Data_Ord_Ord, k_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 *Constructor_Data_Ord_Ord = dictOrd_0_loop
_ = dictOrd_0
var k_1 gopurs_runtime.Value = k_1_loop
_ = k_1
var go__go_2_0_53 gopurs_runtime.Value
_ = go__go_2_0_53
go__go_2_0_53 = gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t5 *Constructor_Data_Maybe_Just
{
if (v_3.Type == 9 && v_3.IntVal == 324739070 && v_3.UnsafePtr == nil) {
__t5 = (*Constructor_Data_Maybe_Just)(nil)
goto end_branch_5
} else {

}
}
{
if (v_3.Type == 9 && v_3.IntVal == 324739070 && v_3.UnsafePtr != nil) {
// TAST (Let): v1_4_1 -> gopurs_runtime.Value
v1_4_1 := gopurs_runtime.Apply2(gopurs_runtime.Box(dictOrd_0.V1), k_1, (*Constructor_Data_Map_Internal_Node)(v_3.UnsafePtr).V2)
_ = v1_4_1
var __t4 *Constructor_Data_Maybe_Just
{
if (uint32(v1_4_1.IntVal) == 1527465420) {
// TAST (Let): v2_5_2 -> *Constructor_Data_Maybe_Just
v2_5_2 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply(go__go_2_0_53, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v_3.UnsafePtr).V4)}))
_ = v2_5_2
var __t3 *Constructor_Data_Maybe_Just
{
if (v2_5_2 == nil) {
__t3 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, gopurs_runtime.RecordDict2("key", "value", (*Constructor_Data_Map_Internal_Node)(v_3.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node)(v_3.UnsafePtr).V3)})})
goto end_branch_3
} else {

}
}
{
__t3 = v2_5_2
}
end_branch_3:
__t4 = __t3
goto end_branch_4
} else {

}
}
{
if (uint32(v1_4_1.IntVal) == 380165415) {
__t4 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply(go__go_2_0_53, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v_3.UnsafePtr).V5)}))
goto end_branch_4
} else {

}
}
{
if (uint32(v1_4_1.IntVal) == 902936544) {
__t4 = Call_Data_Map_Internal_findMin((*Constructor_Data_Map_Internal_Node)(v_3.UnsafePtr).V5)
goto end_branch_4
} else {

}
}
{
__t4 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_4:
__t5 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t4)})
goto end_branch_5
} else {

}
}
{
__t5 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_5:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t5)}
})
return go__go_2_0_53
}

func Call_Data_Map_Internal_findMax(v_0_loop *Constructor_Data_Map_Internal_Node) *Constructor_Data_Maybe_Just {
findMax:
for {
if false { continue findMax }
var v_0 *Constructor_Data_Map_Internal_Node = v_0_loop
_ = v_0
var __t2 *Constructor_Data_Maybe_Just
{
if (v_0 == nil) {
__t2 = (*Constructor_Data_Maybe_Just)(nil)
goto end_branch_2
} else {

}
}
{
if (v_0 != nil) {
var __t1 *Constructor_Data_Maybe_Just
{
var __t_tag_0 *Constructor_Data_Map_Internal_Node = (v_0).V5
if (__t_tag_0 == nil) {
__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, gopurs_runtime.RecordDict2("key", "value", (v_0).V2, (v_0).V3)})})
goto end_branch_1
} else {

}
}
{
v_0_loop = (v_0).V5
continue findMax
__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Value{})
}
end_branch_1:
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t1)})
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_2:
return gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t2)})
}
}

func Call_Data_Map_Internal_lookupLT(dictOrd_0_loop *Constructor_Data_Ord_Ord, k_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 *Constructor_Data_Ord_Ord = dictOrd_0_loop
_ = dictOrd_0
var k_1 gopurs_runtime.Value = k_1_loop
_ = k_1
var go__go_2_0_54 gopurs_runtime.Value
_ = go__go_2_0_54
go__go_2_0_54 = gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t5 *Constructor_Data_Maybe_Just
{
if (v_3.Type == 9 && v_3.IntVal == 324739070 && v_3.UnsafePtr == nil) {
__t5 = (*Constructor_Data_Maybe_Just)(nil)
goto end_branch_5
} else {

}
}
{
if (v_3.Type == 9 && v_3.IntVal == 324739070 && v_3.UnsafePtr != nil) {
// TAST (Let): v1_4_1 -> gopurs_runtime.Value
v1_4_1 := gopurs_runtime.Apply2(gopurs_runtime.Box(dictOrd_0.V1), k_1, (*Constructor_Data_Map_Internal_Node)(v_3.UnsafePtr).V2)
_ = v1_4_1
var __t4 *Constructor_Data_Maybe_Just
{
if (uint32(v1_4_1.IntVal) == 1527465420) {
__t4 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply(go__go_2_0_54, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v_3.UnsafePtr).V4)}))
goto end_branch_4
} else {

}
}
{
if (uint32(v1_4_1.IntVal) == 380165415) {
// TAST (Let): v2_5_2 -> *Constructor_Data_Maybe_Just
v2_5_2 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply(go__go_2_0_54, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v_3.UnsafePtr).V5)}))
_ = v2_5_2
var __t3 *Constructor_Data_Maybe_Just
{
if (v2_5_2 == nil) {
__t3 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, gopurs_runtime.RecordDict2("key", "value", (*Constructor_Data_Map_Internal_Node)(v_3.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node)(v_3.UnsafePtr).V3)})})
goto end_branch_3
} else {

}
}
{
__t3 = v2_5_2
}
end_branch_3:
__t4 = __t3
goto end_branch_4
} else {

}
}
{
if (uint32(v1_4_1.IntVal) == 902936544) {
__t4 = Call_Data_Map_Internal_findMax((*Constructor_Data_Map_Internal_Node)(v_3.UnsafePtr).V4)
goto end_branch_4
} else {

}
}
{
__t4 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_4:
__t5 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t4)})
goto end_branch_5
} else {

}
}
{
__t5 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_5:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t5)}
})
return go__go_2_0_54
}

func Call_Data_Map_Internal_filterWithKey(dictOrd_0_loop *Constructor_Data_Ord_Ord, f_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 *Constructor_Data_Ord_Ord = dictOrd_0_loop
_ = dictOrd_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var go__go_2_0_55 gopurs_runtime.Value
_ = go__go_2_0_55
go__go_2_0_55 = gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 *Constructor_Data_Map_Internal_Node
{
if (v_3.Type == 9 && v_3.IntVal == 324739070 && v_3.UnsafePtr == nil) {
__t2 = (*Constructor_Data_Map_Internal_Node)(nil)
goto end_branch_2
} else {

}
}
{
if (v_3.Type == 9 && v_3.IntVal == 324739070 && v_3.UnsafePtr != nil) {
var __t1 *Constructor_Data_Map_Internal_Node
{
if (gopurs_runtime.Apply2(f_1, (*Constructor_Data_Map_Internal_Node)(v_3.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node)(v_3.UnsafePtr).V3).IntVal) != (0) {
__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeBalancedNode(), (*Constructor_Data_Map_Internal_Node)(v_3.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node)(v_3.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.Apply(go__go_2_0_55, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v_3.UnsafePtr).V4)})))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.Apply(go__go_2_0_55, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v_3.UnsafePtr).V5)})))}))
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.UncurriedApp2(Get_Data_Map_Internal_unsafeJoinNodes(), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.Apply(go__go_2_0_55, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v_3.UnsafePtr).V4)})))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.Apply(go__go_2_0_55, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v_3.UnsafePtr).V5)})))}))
}
end_branch_1:
__t2 = __t1
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
return go__go_2_0_55
}

func Call_Data_Map_Internal_filterKeys(dictOrd_0_loop *Constructor_Data_Ord_Ord, f_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 *Constructor_Data_Ord_Ord = dictOrd_0_loop
_ = dictOrd_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var go__go_2_0_56 gopurs_runtime.Value
_ = go__go_2_0_56
go__go_2_0_56 = gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 *Constructor_Data_Map_Internal_Node
{
if (v_3.Type == 9 && v_3.IntVal == 324739070 && v_3.UnsafePtr == nil) {
__t2 = (*Constructor_Data_Map_Internal_Node)(nil)
goto end_branch_2
} else {

}
}
{
if (v_3.Type == 9 && v_3.IntVal == 324739070 && v_3.UnsafePtr != nil) {
var __t1 *Constructor_Data_Map_Internal_Node
{
if (gopurs_runtime.Apply(f_1, (*Constructor_Data_Map_Internal_Node)(v_3.UnsafePtr).V2).IntVal) != (0) {
__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeBalancedNode(), (*Constructor_Data_Map_Internal_Node)(v_3.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node)(v_3.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.Apply(go__go_2_0_56, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v_3.UnsafePtr).V4)})))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.Apply(go__go_2_0_56, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v_3.UnsafePtr).V5)})))}))
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.UncurriedApp2(Get_Data_Map_Internal_unsafeJoinNodes(), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.Apply(go__go_2_0_56, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v_3.UnsafePtr).V4)})))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.Apply(go__go_2_0_56, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v_3.UnsafePtr).V5)})))}))
}
end_branch_1:
__t2 = __t1
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
return go__go_2_0_56
}

func Call_Data_Map_Internal_filter(dictOrd_0_loop *Constructor_Data_Ord_Ord, x_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 *Constructor_Data_Ord_Ord = dictOrd_0_loop
_ = dictOrd_0
var x_1 gopurs_runtime.Value = x_1_loop
_ = x_1
var go__go_2_0_57 gopurs_runtime.Value
_ = go__go_2_0_57
go__go_2_0_57 = gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 *Constructor_Data_Map_Internal_Node
{
if (v_3.Type == 9 && v_3.IntVal == 324739070 && v_3.UnsafePtr == nil) {
__t2 = (*Constructor_Data_Map_Internal_Node)(nil)
goto end_branch_2
} else {

}
}
{
if (v_3.Type == 9 && v_3.IntVal == 324739070 && v_3.UnsafePtr != nil) {
var __t1 *Constructor_Data_Map_Internal_Node
{
if (gopurs_runtime.Apply(x_1, (*Constructor_Data_Map_Internal_Node)(v_3.UnsafePtr).V3).IntVal) != (0) {
__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeBalancedNode(), (*Constructor_Data_Map_Internal_Node)(v_3.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node)(v_3.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.Apply(go__go_2_0_57, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v_3.UnsafePtr).V4)})))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.Apply(go__go_2_0_57, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v_3.UnsafePtr).V5)})))}))
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.UncurriedApp2(Get_Data_Map_Internal_unsafeJoinNodes(), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.Apply(go__go_2_0_57, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v_3.UnsafePtr).V4)})))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.Apply(go__go_2_0_57, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v_3.UnsafePtr).V5)})))}))
}
end_branch_1:
__t2 = __t1
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
return go__go_2_0_57
}

func Call_Data_Map_Internal_eqMap(dictEq_0_loop gopurs_runtime.Value, dictEq1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictEq_0 gopurs_runtime.Value = dictEq_0_loop
_ = dictEq_0
var dictEq1_1 gopurs_runtime.Value = dictEq1_1_loop
_ = dictEq1_1
var go__go_2_1_58 gopurs_runtime.Value
go__go_2_1_58 = gopurs_runtime.Func(func(a_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var a_3_loop gopurs_runtime.Value = a_3_loop_val
var b_4_loop gopurs_runtime.Value = b_4_loop_val
go__go_2_1_58:
for {
if false { continue go__go_2_1_58 }
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
continue go__go_2_1_58
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
eqMapIter2_2_0 := &Constructor_Data_Eq_Eq{1, go__go_2_1_58}
_ = eqMapIter2_2_0
return gopurs_runtime.RecordDict1("eq", gopurs_runtime.Func(func(xs_3 gopurs_runtime.Value) gopurs_runtime.Value {
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
}))
}

func Call_Data_Map_Internal_ordMap(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
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
var go__go_5_5_59 gopurs_runtime.Value
go__go_5_5_59 = gopurs_runtime.Func(func(a_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_7_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var a_6_loop gopurs_runtime.Value = a_6_loop_val
var b_7_loop gopurs_runtime.Value = b_7_loop_val
go__go_5_5_59:
for {
if false { continue go__go_5_5_59 }
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
continue go__go_5_5_59
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
// TAST (Let): eqMapIter2_4_3 -> gopurs_runtime.Value
eqMapIter2_4_3 := gopurs_runtime.RecordDict1("eq", go__go_5_5_59)
_ = eqMapIter2_4_3
var go__go_5_10_60 gopurs_runtime.Value
go__go_5_10_60 = gopurs_runtime.Func(func(a_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_7_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var a_6_loop gopurs_runtime.Value = a_6_loop_val
var b_7_loop gopurs_runtime.Value = b_7_loop_val
go__go_5_10_60:
for {
if false { continue go__go_5_10_60 }
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
continue go__go_5_10_60
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
ordMapIter2_4_2 := gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord](gopurs_runtime.RecordDict2("Eq0", "compare", gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return eqMapIter2_4_3
}), go__go_5_10_60))
_ = ordMapIter2_4_2
// TAST (Let): __local_var_5_21 -> gopurs_runtime.Value
__local_var_5_21 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd1_3, "Eq0"), gopurs_runtime.Value{})
_ = __local_var_5_21
var go__go_6_23_61 gopurs_runtime.Value
go__go_6_23_61 = gopurs_runtime.Func(func(a_7_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_8_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var a_7_loop gopurs_runtime.Value = a_7_loop_val
var b_8_loop gopurs_runtime.Value = b_8_loop_val
go__go_6_23_61:
for {
if false { continue go__go_6_23_61 }
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
continue go__go_6_23_61
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
eqMapIter2_6_22 := &Constructor_Data_Eq_Eq{1, go__go_6_23_61}
_ = eqMapIter2_6_22
// TAST (Let): eqMap2_5_20 -> gopurs_runtime.Value
eqMap2_5_20 := gopurs_runtime.RecordDict1("eq", gopurs_runtime.Func(func(xs_7 gopurs_runtime.Value) gopurs_runtime.Value {
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
}))
_ = eqMap2_5_20
return gopurs_runtime.RecordDict2("Eq0", "compare", gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return eqMap2_5_20
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
}))
})
}

func Call_Data_Map_Internal_eq1Map(dictEq_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictEq_0 gopurs_runtime.Value = dictEq_0_loop
_ = dictEq_0
return gopurs_runtime.RecordDict1("eq1", gopurs_runtime.Func(func(dictEq1_1 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_2_1_62 gopurs_runtime.Value
go__go_2_1_62 = gopurs_runtime.Func(func(a_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var a_3_loop gopurs_runtime.Value = a_3_loop_val
var b_4_loop gopurs_runtime.Value = b_4_loop_val
go__go_2_1_62:
for {
if false { continue go__go_2_1_62 }
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
continue go__go_2_1_62
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
eqMapIter2_2_0 := &Constructor_Data_Eq_Eq{1, go__go_2_1_62}
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
}))
}

func Call_Data_Map_Internal_ord1Map(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
// TAST (Let): __local_var_1_0 -> gopurs_runtime.Value
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd_0, "Eq0"), gopurs_runtime.Value{})
_ = __local_var_1_0
// TAST (Let): __local_var_2_2 -> gopurs_runtime.Value
__local_var_2_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd_0, "Eq0"), gopurs_runtime.Value{})
_ = __local_var_2_2
// TAST (Let): eq1Map1_2_1 -> gopurs_runtime.Value
eq1Map1_2_1 := gopurs_runtime.RecordDict1("eq1", gopurs_runtime.Func(func(dictEq1_3 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_4_4_63 gopurs_runtime.Value
go__go_4_4_63 = gopurs_runtime.Func(func(a_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var a_5_loop gopurs_runtime.Value = a_5_loop_val
var b_6_loop gopurs_runtime.Value = b_6_loop_val
go__go_4_4_63:
for {
if false { continue go__go_4_4_63 }
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
continue go__go_4_4_63
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
eqMapIter2_4_3 := &Constructor_Data_Eq_Eq{1, go__go_4_4_63}
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
}))
_ = eq1Map1_2_1
return gopurs_runtime.RecordDict2("Eq10", "compare1", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return eq1Map1_2_1
}), gopurs_runtime.Func(func(dictOrd1_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_4_14 -> gopurs_runtime.Value
__local_var_4_14 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd1_3, "Eq0"), gopurs_runtime.Value{})
_ = __local_var_4_14
var go__go_5_15_64 gopurs_runtime.Value
go__go_5_15_64 = gopurs_runtime.Func(func(a_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_7_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var a_6_loop gopurs_runtime.Value = a_6_loop_val
var b_7_loop gopurs_runtime.Value = b_7_loop_val
go__go_5_15_64:
for {
if false { continue go__go_5_15_64 }
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
continue go__go_5_15_64
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
// TAST (Let): eqMapIter2_4_13 -> gopurs_runtime.Value
eqMapIter2_4_13 := gopurs_runtime.RecordDict1("eq", go__go_5_15_64)
_ = eqMapIter2_4_13
var go__go_5_20_65 gopurs_runtime.Value
go__go_5_20_65 = gopurs_runtime.Func(func(a_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_7_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var a_6_loop gopurs_runtime.Value = a_6_loop_val
var b_7_loop gopurs_runtime.Value = b_7_loop_val
go__go_5_20_65:
for {
if false { continue go__go_5_20_65 }
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
continue go__go_5_20_65
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
ordMapIter2_4_12 := gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord](gopurs_runtime.RecordDict2("Eq0", "compare", gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return eqMapIter2_4_13
}), go__go_5_20_65))
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
}))
}

func Call_Data_Map_Internal_fromFoldable(dictOrd_0_loop *Constructor_Data_Ord_Ord, dictFoldable_1_loop *Constructor_Data_Foldable_Foldable) gopurs_runtime.Value {
var dictOrd_0 *Constructor_Data_Ord_Ord = dictOrd_0_loop
_ = dictOrd_0
var dictFoldable_1 *Constructor_Data_Foldable_Foldable = dictFoldable_1_loop
_ = dictFoldable_1
return gopurs_runtime.Apply2(gopurs_runtime.Box(dictFoldable_1.V1), gopurs_runtime.Func(func(m_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_4_0 -> gopurs_runtime.Value
__local_var_4_0 := (*Constructor_Data_Tuple_Tuple)(v_3.UnsafePtr).V0
_ = __local_var_4_0
// TAST (Let): __local_var_5_1 -> gopurs_runtime.Value
__local_var_5_1 := (*Constructor_Data_Tuple_Tuple)(v_3.UnsafePtr).V1
_ = __local_var_5_1
var go__go_6_2_66 gopurs_runtime.Value
_ = go__go_6_2_66
go__go_6_2_66 = gopurs_runtime.Func(func(v1_7 gopurs_runtime.Value) gopurs_runtime.Value {
var __t5 *Constructor_Data_Map_Internal_Node
{
if (v1_7.Type == 9 && v1_7.IntVal == 324739070 && v1_7.UnsafePtr == nil) {
__t5 = &Constructor_Data_Map_Internal_Node{1, 1, 1, __local_var_4_0, __local_var_5_1, (*Constructor_Data_Map_Internal_Node)(nil), (*Constructor_Data_Map_Internal_Node)(nil)}
goto end_branch_5
} else {

}
}
{
if (v1_7.Type == 9 && v1_7.IntVal == 324739070 && v1_7.UnsafePtr != nil) {
// TAST (Let): v2_8_3 -> gopurs_runtime.Value
v2_8_3 := gopurs_runtime.Apply2(gopurs_runtime.Box(dictOrd_0.V1), __local_var_4_0, (*Constructor_Data_Map_Internal_Node)(v1_7.UnsafePtr).V2)
_ = v2_8_3
var __t4 *Constructor_Data_Map_Internal_Node
{
if (uint32(v2_8_3.IntVal) == 1527465420) {
__t4 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeBalancedNode(), (*Constructor_Data_Map_Internal_Node)(v1_7.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node)(v1_7.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.Apply(go__go_6_2_66, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v1_7.UnsafePtr).V4)})))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v1_7.UnsafePtr).V5)}))
goto end_branch_4
} else {

}
}
{
if (uint32(v2_8_3.IntVal) == 380165415) {
__t4 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeBalancedNode(), (*Constructor_Data_Map_Internal_Node)(v1_7.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node)(v1_7.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v1_7.UnsafePtr).V4)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.Apply(go__go_6_2_66, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v1_7.UnsafePtr).V5)})))}))
goto end_branch_4
} else {

}
}
{
if (uint32(v2_8_3.IntVal) == 902936544) {
__t4 = &Constructor_Data_Map_Internal_Node{1, (*Constructor_Data_Map_Internal_Node)(v1_7.UnsafePtr).V0, (*Constructor_Data_Map_Internal_Node)(v1_7.UnsafePtr).V1, __local_var_4_0, __local_var_5_1, (*Constructor_Data_Map_Internal_Node)(v1_7.UnsafePtr).V4, (*Constructor_Data_Map_Internal_Node)(v1_7.UnsafePtr).V5}
goto end_branch_4
} else {

}
}
{
__t4 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_4:
__t5 = __t4
goto end_branch_5
} else {

}
}
{
__t5 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_5:
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(__t5)}
})
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.Apply(go__go_6_2_66, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](m_2))})))}
})
}), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(nil))})
}

func Call_Data_Map_Internal_fromFoldableWith(dictOrd_0_loop *Constructor_Data_Ord_Ord, dictFoldable_1_loop *Constructor_Data_Foldable_Foldable, f_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 *Constructor_Data_Ord_Ord = dictOrd_0_loop
_ = dictOrd_0
var dictFoldable_1 *Constructor_Data_Foldable_Foldable = dictFoldable_1_loop
_ = dictFoldable_1
var f_2 gopurs_runtime.Value = f_2_loop
_ = f_2
return gopurs_runtime.Apply2(gopurs_runtime.Box(dictFoldable_1.V1), gopurs_runtime.Func(func(m_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_5_0 -> gopurs_runtime.Value
__local_var_5_0 := (*Constructor_Data_Tuple_Tuple)(v_4.UnsafePtr).V0
_ = __local_var_5_0
// TAST (Let): __local_var_6_1 -> gopurs_runtime.Value
__local_var_6_1 := (*Constructor_Data_Tuple_Tuple)(v_4.UnsafePtr).V1
_ = __local_var_6_1
var go__go_7_2_67 gopurs_runtime.Value
_ = go__go_7_2_67
go__go_7_2_67 = gopurs_runtime.Func(func(v1_8 gopurs_runtime.Value) gopurs_runtime.Value {
var __t5 *Constructor_Data_Map_Internal_Node
{
if (v1_8.Type == 9 && v1_8.IntVal == 324739070 && v1_8.UnsafePtr == nil) {
__t5 = &Constructor_Data_Map_Internal_Node{1, 1, 1, __local_var_5_0, __local_var_6_1, (*Constructor_Data_Map_Internal_Node)(nil), (*Constructor_Data_Map_Internal_Node)(nil)}
goto end_branch_5
} else {

}
}
{
if (v1_8.Type == 9 && v1_8.IntVal == 324739070 && v1_8.UnsafePtr != nil) {
// TAST (Let): v2_9_3 -> gopurs_runtime.Value
v2_9_3 := gopurs_runtime.Apply2(gopurs_runtime.Box(dictOrd_0.V1), __local_var_5_0, (*Constructor_Data_Map_Internal_Node)(v1_8.UnsafePtr).V2)
_ = v2_9_3
var __t4 *Constructor_Data_Map_Internal_Node
{
if (uint32(v2_9_3.IntVal) == 1527465420) {
__t4 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeBalancedNode(), (*Constructor_Data_Map_Internal_Node)(v1_8.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node)(v1_8.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.Apply(go__go_7_2_67, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v1_8.UnsafePtr).V4)})))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v1_8.UnsafePtr).V5)}))
goto end_branch_4
} else {

}
}
{
if (uint32(v2_9_3.IntVal) == 380165415) {
__t4 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeBalancedNode(), (*Constructor_Data_Map_Internal_Node)(v1_8.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node)(v1_8.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v1_8.UnsafePtr).V4)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.Apply(go__go_7_2_67, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v1_8.UnsafePtr).V5)})))}))
goto end_branch_4
} else {

}
}
{
if (uint32(v2_9_3.IntVal) == 902936544) {
__t4 = &Constructor_Data_Map_Internal_Node{1, (*Constructor_Data_Map_Internal_Node)(v1_8.UnsafePtr).V0, (*Constructor_Data_Map_Internal_Node)(v1_8.UnsafePtr).V1, __local_var_5_0, gopurs_runtime.Apply2(f_2, __local_var_6_1, (*Constructor_Data_Map_Internal_Node)(v1_8.UnsafePtr).V3), (*Constructor_Data_Map_Internal_Node)(v1_8.UnsafePtr).V4, (*Constructor_Data_Map_Internal_Node)(v1_8.UnsafePtr).V5}
goto end_branch_4
} else {

}
}
{
__t4 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_4:
__t5 = __t4
goto end_branch_5
} else {

}
}
{
__t5 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_5:
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(__t5)}
})
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.Apply(go__go_7_2_67, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](m_3))})))}
})
}), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(nil))})
}

func Call_Data_Map_Internal_fromFoldableWithIndex(dictOrd_0_loop *Constructor_Data_Ord_Ord, dictFoldableWithIndex_1_loop *Constructor_Data_FoldableWithIndex_FoldableWithIndex) gopurs_runtime.Value {
var dictOrd_0 *Constructor_Data_Ord_Ord = dictOrd_0_loop
_ = dictOrd_0
var dictFoldableWithIndex_1 *Constructor_Data_FoldableWithIndex_FoldableWithIndex = dictFoldableWithIndex_1_loop
_ = dictFoldableWithIndex_1
return gopurs_runtime.Apply2(gopurs_runtime.Box(dictFoldableWithIndex_1.V2), gopurs_runtime.Func(func(k_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_5_0_68 gopurs_runtime.Value
_ = go__go_5_0_68
go__go_5_0_68 = gopurs_runtime.Func(func(v1_6 gopurs_runtime.Value) gopurs_runtime.Value {
var __t3 *Constructor_Data_Map_Internal_Node
{
if (v1_6.Type == 9 && v1_6.IntVal == 324739070 && v1_6.UnsafePtr == nil) {
__t3 = &Constructor_Data_Map_Internal_Node{1, 1, 1, k_2, v_4, (*Constructor_Data_Map_Internal_Node)(nil), (*Constructor_Data_Map_Internal_Node)(nil)}
goto end_branch_3
} else {

}
}
{
if (v1_6.Type == 9 && v1_6.IntVal == 324739070 && v1_6.UnsafePtr != nil) {
// TAST (Let): v2_7_1 -> gopurs_runtime.Value
v2_7_1 := gopurs_runtime.Apply2(gopurs_runtime.Box(dictOrd_0.V1), k_2, (*Constructor_Data_Map_Internal_Node)(v1_6.UnsafePtr).V2)
_ = v2_7_1
var __t2 *Constructor_Data_Map_Internal_Node
{
if (uint32(v2_7_1.IntVal) == 1527465420) {
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeBalancedNode(), (*Constructor_Data_Map_Internal_Node)(v1_6.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node)(v1_6.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.Apply(go__go_5_0_68, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v1_6.UnsafePtr).V4)})))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v1_6.UnsafePtr).V5)}))
goto end_branch_2
} else {

}
}
{
if (uint32(v2_7_1.IntVal) == 380165415) {
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeBalancedNode(), (*Constructor_Data_Map_Internal_Node)(v1_6.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node)(v1_6.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v1_6.UnsafePtr).V4)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.Apply(go__go_5_0_68, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v1_6.UnsafePtr).V5)})))}))
goto end_branch_2
} else {

}
}
{
if (uint32(v2_7_1.IntVal) == 902936544) {
__t2 = &Constructor_Data_Map_Internal_Node{1, (*Constructor_Data_Map_Internal_Node)(v1_6.UnsafePtr).V0, (*Constructor_Data_Map_Internal_Node)(v1_6.UnsafePtr).V1, k_2, v_4, (*Constructor_Data_Map_Internal_Node)(v1_6.UnsafePtr).V4, (*Constructor_Data_Map_Internal_Node)(v1_6.UnsafePtr).V5}
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_2:
__t3 = __t2
goto end_branch_3
} else {

}
}
{
__t3 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_3:
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(__t3)}
})
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.Apply(go__go_5_0_68, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](m_3))})))}
})
})
}), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(nil))})
}

func Call_Data_Map_Internal_monoidSemigroupMap(_dollar__unused_0_loop gopurs_runtime.Value, dictOrd_1_loop gopurs_runtime.Value, dictSemigroup_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var _dollar__unused_0 gopurs_runtime.Value = _dollar__unused_0_loop
_ = _dollar__unused_0
var dictOrd_1 gopurs_runtime.Value = dictOrd_1_loop
_ = dictOrd_1
var dictSemigroup_2 gopurs_runtime.Value = dictSemigroup_2_loop
_ = dictSemigroup_2
// TAST (Let): compare_3_1 -> gopurs_runtime.Value
compare_3_1 := gopurs_runtime.RecordGet(dictOrd_1, "compare")
_ = compare_3_1
// TAST (Let): __local_var_4_2 -> gopurs_runtime.Value
__local_var_4_2 := gopurs_runtime.RecordGet(dictSemigroup_2, "append")
_ = __local_var_4_2
// TAST (Let): semigroupMap3_3_0 -> gopurs_runtime.Value
semigroupMap3_3_0 := gopurs_runtime.RecordDict1("append", gopurs_runtime.Func(func(m1_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m2_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeUnionWith(), compare_3_1, __local_var_4_2, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](m1_5))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](m2_6))})))}
})
}))
_ = semigroupMap3_3_0
return gopurs_runtime.RecordDict2("Semigroup0", "mempty", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return semigroupMap3_3_0
}), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(nil))})
}

func Call_Data_Map_Internal_submap(dictOrd_0_loop *Constructor_Data_Ord_Ord) gopurs_runtime.Value {
var dictOrd_0 *Constructor_Data_Ord_Ord = dictOrd_0_loop
_ = dictOrd_0
// TAST (Let): compare_1_1 -> gopurs_runtime.Value
compare_1_1 := gopurs_runtime.Box(dictOrd_0.V1)
_ = compare_1_1
// TAST (Let): union1_1_0 -> gopurs_runtime.Value
union1_1_0 := gopurs_runtime.Func(func(m1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m2_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeUnionWith(), compare_1_1, Get_Data_Function_go__const(), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](m1_2))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](m2_3))})))}
})
})
_ = union1_1_0
return gopurs_runtime.Func(func(kmin_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(kmax_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t8 gopurs_runtime.Value
{
var __t_tag_3 *Constructor_Data_Maybe_Just = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](kmin_2)
if (__t_tag_3 != nil) {
// TAST (Let): __local_var_4_4 -> gopurs_runtime.Value
__local_var_4_4 := (*Constructor_Data_Maybe_Just)(kmin_2.UnsafePtr).V0
_ = __local_var_4_4
__t8 = gopurs_runtime.Func(func(k_5 gopurs_runtime.Value) gopurs_runtime.Value {
var __t6 bool
{
var __t_tag_5 gopurs_runtime.Value = gopurs_runtime.Apply2(gopurs_runtime.Box(dictOrd_0.V1), k_5, __local_var_4_4)
if (uint32(__t_tag_5.IntVal) == 1527465420) {
__t6 = true
goto end_branch_6
} else {

}
}
{
__t6 = false
}
end_branch_6:
return gopurs_runtime.Bool(__t6)
})
goto end_branch_8
} else {

}
}
{
var __t_tag_7 *Constructor_Data_Maybe_Just = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](kmin_2)
if (__t_tag_7 == nil) {
__t8 = gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(false)
})
goto end_branch_8
} else {

}
}
{
__t8 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_8:
// TAST (Let): tooSmall_4_2 -> gopurs_runtime.Value
tooSmall_4_2 := __t8
_ = tooSmall_4_2
var __t15 gopurs_runtime.Value
{
var __t_tag_10 *Constructor_Data_Maybe_Just = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](kmax_3)
if (__t_tag_10 != nil) {
// TAST (Let): __local_var_5_11 -> gopurs_runtime.Value
__local_var_5_11 := (*Constructor_Data_Maybe_Just)(kmax_3.UnsafePtr).V0
_ = __local_var_5_11
__t15 = gopurs_runtime.Func(func(k_6 gopurs_runtime.Value) gopurs_runtime.Value {
var __t13 bool
{
var __t_tag_12 gopurs_runtime.Value = gopurs_runtime.Apply2(gopurs_runtime.Box(dictOrd_0.V1), k_6, __local_var_5_11)
if (uint32(__t_tag_12.IntVal) == 380165415) {
__t13 = true
goto end_branch_13
} else {

}
}
{
__t13 = false
}
end_branch_13:
return gopurs_runtime.Bool(__t13)
})
goto end_branch_15
} else {

}
}
{
var __t_tag_14 *Constructor_Data_Maybe_Just = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](kmax_3)
if (__t_tag_14 == nil) {
__t15 = gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(false)
})
goto end_branch_15
} else {

}
}
{
__t15 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_15:
// TAST (Let): tooLarge_5_9 -> gopurs_runtime.Value
tooLarge_5_9 := __t15
_ = tooLarge_5_9
var __t38 gopurs_runtime.Value
{
var __t_tag_17 *Constructor_Data_Maybe_Just = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](kmin_2)
if (__t_tag_17 != nil) {
var __t30 gopurs_runtime.Value
{
var __t_tag_18 *Constructor_Data_Maybe_Just = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](kmax_3)
if (__t_tag_18 != nil) {
// TAST (Let): __local_var_6_19 -> gopurs_runtime.Value
__local_var_6_19 := (*Constructor_Data_Maybe_Just)(kmax_3.UnsafePtr).V0
_ = __local_var_6_19
// TAST (Let): __local_var_7_20 -> gopurs_runtime.Value
__local_var_7_20 := (*Constructor_Data_Maybe_Just)(kmin_2.UnsafePtr).V0
_ = __local_var_7_20
__t30 = gopurs_runtime.Func(func(k_8 gopurs_runtime.Value) gopurs_runtime.Value {
var __t22 bool
{
var __t_tag_21 gopurs_runtime.Value = gopurs_runtime.Apply2(gopurs_runtime.Box(dictOrd_0.V1), __local_var_7_20, k_8)
if (uint32(__t_tag_21.IntVal) == 380165415) {
__t22 = false
goto end_branch_22
} else {

}
}
{
__t22 = true
}
end_branch_22:
var __t_and_25 bool = false
if __t22 {

var __t24 bool
{
var __t_tag_23 gopurs_runtime.Value = gopurs_runtime.Apply2(gopurs_runtime.Box(dictOrd_0.V1), k_8, __local_var_6_19)
if (uint32(__t_tag_23.IntVal) == 380165415) {
__t24 = false
goto end_branch_24
} else {

}
}
{
__t24 = true
}
end_branch_24:
__t_and_25 = __t24
}
return gopurs_runtime.Bool(__t_and_25)
})
goto end_branch_30
} else {

}
}
{
var __t_tag_26 *Constructor_Data_Maybe_Just = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](kmax_3)
if (__t_tag_26 == nil) {
// TAST (Let): __local_var_6_27 -> gopurs_runtime.Value
__local_var_6_27 := (*Constructor_Data_Maybe_Just)(kmin_2.UnsafePtr).V0
_ = __local_var_6_27
__t30 = gopurs_runtime.Func(func(k_7 gopurs_runtime.Value) gopurs_runtime.Value {
var __t29 bool
{
var __t_tag_28 gopurs_runtime.Value = gopurs_runtime.Apply2(gopurs_runtime.Box(dictOrd_0.V1), __local_var_6_27, k_7)
if (uint32(__t_tag_28.IntVal) == 380165415) {
__t29 = false
goto end_branch_29
} else {

}
}
{
__t29 = true
}
end_branch_29:
return gopurs_runtime.Bool(__t29)
})
goto end_branch_30
} else {

}
}
{
__t30 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_30:
__t38 = __t30
goto end_branch_38
} else {

}
}
{
var __t_tag_31 *Constructor_Data_Maybe_Just = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](kmin_2)
if (__t_tag_31 == nil) {
var __t37 gopurs_runtime.Value
{
var __t_tag_32 *Constructor_Data_Maybe_Just = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](kmax_3)
if (__t_tag_32 != nil) {
// TAST (Let): __local_var_6_33 -> gopurs_runtime.Value
__local_var_6_33 := (*Constructor_Data_Maybe_Just)(kmax_3.UnsafePtr).V0
_ = __local_var_6_33
__t37 = gopurs_runtime.Func(func(k_7 gopurs_runtime.Value) gopurs_runtime.Value {
var __t35 bool
{
var __t_tag_34 gopurs_runtime.Value = gopurs_runtime.Apply2(gopurs_runtime.Box(dictOrd_0.V1), k_7, __local_var_6_33)
if (uint32(__t_tag_34.IntVal) == 380165415) {
__t35 = false
goto end_branch_35
} else {

}
}
{
__t35 = true
}
end_branch_35:
return gopurs_runtime.Bool(__t35)
})
goto end_branch_37
} else {

}
}
{
var __t_tag_36 *Constructor_Data_Maybe_Just = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](kmax_3)
if (__t_tag_36 == nil) {
__t37 = gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(true)
})
goto end_branch_37
} else {

}
}
{
__t37 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_37:
__t38 = __t37
goto end_branch_38
} else {

}
}
{
__t38 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_38:
// TAST (Let): inBounds_6_16 -> gopurs_runtime.Value
inBounds_6_16 := __t38
_ = inBounds_6_16
var go__go_7_39_69 gopurs_runtime.Value
_ = go__go_7_39_69
go__go_7_39_69 = gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
var __t43 *Constructor_Data_Map_Internal_Node
{
if (v_8.Type == 9 && v_8.IntVal == 324739070 && v_8.UnsafePtr == nil) {
__t43 = (*Constructor_Data_Map_Internal_Node)(nil)
goto end_branch_43
} else {

}
}
{
if (v_8.Type == 9 && v_8.IntVal == 324739070 && v_8.UnsafePtr != nil) {
var __t40 *Constructor_Data_Map_Internal_Node
{
if (gopurs_runtime.Apply(tooSmall_4_2, (*Constructor_Data_Map_Internal_Node)(v_8.UnsafePtr).V2).IntVal) != (0) {
__t40 = (*Constructor_Data_Map_Internal_Node)(nil)
goto end_branch_40
} else {

}
}
{
__t40 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.Apply(go__go_7_39_69, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v_8.UnsafePtr).V4)}))
}
end_branch_40:
var __t41 *Constructor_Data_Map_Internal_Node
{
if (gopurs_runtime.Apply(inBounds_6_16, (*Constructor_Data_Map_Internal_Node)(v_8.UnsafePtr).V2).IntVal) != (0) {
__t41 = &Constructor_Data_Map_Internal_Node{1, 1, 1, (*Constructor_Data_Map_Internal_Node)(v_8.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node)(v_8.UnsafePtr).V3, (*Constructor_Data_Map_Internal_Node)(nil), (*Constructor_Data_Map_Internal_Node)(nil)}
goto end_branch_41
} else {

}
}
{
__t41 = (*Constructor_Data_Map_Internal_Node)(nil)
}
end_branch_41:
var __t42 *Constructor_Data_Map_Internal_Node
{
if (gopurs_runtime.Apply(tooLarge_5_9, (*Constructor_Data_Map_Internal_Node)(v_8.UnsafePtr).V2).IntVal) != (0) {
__t42 = (*Constructor_Data_Map_Internal_Node)(nil)
goto end_branch_42
} else {

}
}
{
__t42 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.Apply(go__go_7_39_69, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v_8.UnsafePtr).V5)}))
}
end_branch_42:
__t43 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.Apply2(union1_1_0, gopurs_runtime.Apply2(union1_1_0, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(__t40)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(__t41)}), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(__t42)}))
goto end_branch_43
} else {

}
}
{
__t43 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_43:
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(__t43)}
})
return go__go_7_39_69
})
})
}

func Call_Data_Map_Internal_unions(dictOrd_0_loop *Constructor_Data_Ord_Ord) gopurs_runtime.Value {
var dictOrd_0 *Constructor_Data_Ord_Ord = dictOrd_0_loop
_ = dictOrd_0
// TAST (Let): compare_1_1 -> gopurs_runtime.Value
compare_1_1 := gopurs_runtime.Box(dictOrd_0.V1)
_ = compare_1_1
// TAST (Let): union1_1_0 -> gopurs_runtime.Value
union1_1_0 := gopurs_runtime.Func(func(m1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m2_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeUnionWith(), compare_1_1, Get_Data_Function_go__const(), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](m1_2))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](m2_3))})))}
})
})
_ = union1_1_0
return gopurs_runtime.Func(func(dictFoldable_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFoldable_2, "foldl"), union1_1_0, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(nil))})
})
}

func Call_Data_Map_Internal_difference(dictOrd_0_loop *Constructor_Data_Ord_Ord) gopurs_runtime.Value {
var dictOrd_0 *Constructor_Data_Ord_Ord = dictOrd_0_loop
_ = dictOrd_0
// TAST (Let): compare_1_0 -> gopurs_runtime.Value
compare_1_0 := gopurs_runtime.Box(dictOrd_0.V1)
_ = compare_1_0
return gopurs_runtime.Func(func(m1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m2_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.UncurriedApp3(Get_Data_Map_Internal_unsafeDifference(), compare_1_0, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](m1_2))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](m2_3))})))}
})
})
}

func Call_Data_Map_Internal_delete(dictOrd_0_loop *Constructor_Data_Ord_Ord, k_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 *Constructor_Data_Ord_Ord = dictOrd_0_loop
_ = dictOrd_0
var k_1 gopurs_runtime.Value = k_1_loop
_ = k_1
var go__go_2_0_70 gopurs_runtime.Value
_ = go__go_2_0_70
go__go_2_0_70 = gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t3 *Constructor_Data_Map_Internal_Node
{
if (v_3.Type == 9 && v_3.IntVal == 324739070 && v_3.UnsafePtr == nil) {
__t3 = (*Constructor_Data_Map_Internal_Node)(nil)
goto end_branch_3
} else {

}
}
{
if (v_3.Type == 9 && v_3.IntVal == 324739070 && v_3.UnsafePtr != nil) {
// TAST (Let): v1_4_1 -> gopurs_runtime.Value
v1_4_1 := gopurs_runtime.Apply2(gopurs_runtime.Box(dictOrd_0.V1), k_1, (*Constructor_Data_Map_Internal_Node)(v_3.UnsafePtr).V2)
_ = v1_4_1
var __t2 *Constructor_Data_Map_Internal_Node
{
if (uint32(v1_4_1.IntVal) == 1527465420) {
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeBalancedNode(), (*Constructor_Data_Map_Internal_Node)(v_3.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node)(v_3.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.Apply(go__go_2_0_70, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v_3.UnsafePtr).V4)})))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v_3.UnsafePtr).V5)}))
goto end_branch_2
} else {

}
}
{
if (uint32(v1_4_1.IntVal) == 380165415) {
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeBalancedNode(), (*Constructor_Data_Map_Internal_Node)(v_3.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node)(v_3.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v_3.UnsafePtr).V4)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.Apply(go__go_2_0_70, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v_3.UnsafePtr).V5)})))}))
goto end_branch_2
} else {

}
}
{
if (uint32(v1_4_1.IntVal) == 902936544) {
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.UncurriedApp2(Get_Data_Map_Internal_unsafeJoinNodes(), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v_3.UnsafePtr).V4)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v_3.UnsafePtr).V5)}))
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_2:
__t3 = __t2
goto end_branch_3
} else {

}
}
{
__t3 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_3:
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(__t3)}
})
return go__go_2_0_70
}

func Call_Data_Map_Internal_checkValid(dictOrd_0_loop *Constructor_Data_Ord_Ord) gopurs_runtime.Value {
var dictOrd_0 *Constructor_Data_Ord_Ord = dictOrd_0_loop
_ = dictOrd_0
var go__go_1_0_71 gopurs_runtime.Value
_ = go__go_1_0_71
go__go_1_0_71 = gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t33 bool
{
if (v_2.Type == 9 && v_2.IntVal == 324739070 && v_2.UnsafePtr == nil) {
__t33 = true
goto end_branch_33
} else {

}
}
{
if (v_2.Type == 9 && v_2.IntVal == 324739070 && v_2.UnsafePtr != nil) {
var __t32 bool
{
var __t_tag_1 *Constructor_Data_Map_Internal_Node = (*Constructor_Data_Map_Internal_Node)(v_2.UnsafePtr).V4
if (__t_tag_1 == nil) {
var __t10 bool
{
var __t_tag_2 *Constructor_Data_Map_Internal_Node = (*Constructor_Data_Map_Internal_Node)(v_2.UnsafePtr).V5
if (__t_tag_2 == nil) {
__t10 = true
goto end_branch_10
} else {

}
}
{
var __t_tag_3 *Constructor_Data_Map_Internal_Node = (*Constructor_Data_Map_Internal_Node)(v_2.UnsafePtr).V5
if (__t_tag_3 != nil) {
var __t_and_9 bool = false
if ((*Constructor_Data_Map_Internal_Node)(v_2.UnsafePtr).V0) == (2) {

var __t_and_8 bool = false
if (((*Constructor_Data_Map_Internal_Node)(v_2.UnsafePtr).V5).V0) == (1) {

var __t4 bool
{
if ((*Constructor_Data_Map_Internal_Node)(v_2.UnsafePtr).V1) > (((*Constructor_Data_Map_Internal_Node)(v_2.UnsafePtr).V5).V1) {
__t4 = true
goto end_branch_4
} else {

}
}
{
__t4 = false
}
end_branch_4:
var __t_and_7 bool = false
if __t4 {

var __t6 bool
{
var __t_tag_5 gopurs_runtime.Value = gopurs_runtime.Apply2(gopurs_runtime.Box(dictOrd_0.V1), ((*Constructor_Data_Map_Internal_Node)(v_2.UnsafePtr).V5).V2, (*Constructor_Data_Map_Internal_Node)(v_2.UnsafePtr).V2)
if (uint32(__t_tag_5.IntVal) == 380165415) {
__t6 = true
goto end_branch_6
} else {

}
}
{
__t6 = false
}
end_branch_6:
__t_and_7 = (__t6) && ((gopurs_runtime.Apply(go__go_1_0_71, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v_2.UnsafePtr).V5)}).IntVal) != (0))
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
__t10 = (func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal) != (0)
}
end_branch_10:
__t32 = __t10
goto end_branch_32
} else {

}
}
{
var __t_tag_11 *Constructor_Data_Map_Internal_Node = (*Constructor_Data_Map_Internal_Node)(v_2.UnsafePtr).V4
if (__t_tag_11 != nil) {
var __t31 bool
{
var __t_tag_12 *Constructor_Data_Map_Internal_Node = (*Constructor_Data_Map_Internal_Node)(v_2.UnsafePtr).V5
if (__t_tag_12 == nil) {
var __t_and_18 bool = false
if ((*Constructor_Data_Map_Internal_Node)(v_2.UnsafePtr).V0) == (2) {

var __t_and_17 bool = false
if (((*Constructor_Data_Map_Internal_Node)(v_2.UnsafePtr).V4).V0) == (1) {

var __t13 bool
{
if ((*Constructor_Data_Map_Internal_Node)(v_2.UnsafePtr).V1) > (((*Constructor_Data_Map_Internal_Node)(v_2.UnsafePtr).V4).V1) {
__t13 = true
goto end_branch_13
} else {

}
}
{
__t13 = false
}
end_branch_13:
var __t_and_16 bool = false
if __t13 {

var __t15 bool
{
var __t_tag_14 gopurs_runtime.Value = gopurs_runtime.Apply2(gopurs_runtime.Box(dictOrd_0.V1), ((*Constructor_Data_Map_Internal_Node)(v_2.UnsafePtr).V4).V2, (*Constructor_Data_Map_Internal_Node)(v_2.UnsafePtr).V2)
if (uint32(__t_tag_14.IntVal) == 1527465420) {
__t15 = true
goto end_branch_15
} else {

}
}
{
__t15 = false
}
end_branch_15:
__t_and_16 = (__t15) && ((gopurs_runtime.Apply(go__go_1_0_71, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v_2.UnsafePtr).V4)}).IntVal) != (0))
}
__t_and_17 = __t_and_16
}
__t_and_18 = __t_and_17
}
__t31 = __t_and_18
goto end_branch_31
} else {

}
}
{
var __t_tag_19 *Constructor_Data_Map_Internal_Node = (*Constructor_Data_Map_Internal_Node)(v_2.UnsafePtr).V5
if (__t_tag_19 != nil) {
var __t20 bool
{
if ((*Constructor_Data_Map_Internal_Node)(v_2.UnsafePtr).V0) > (((*Constructor_Data_Map_Internal_Node)(v_2.UnsafePtr).V5).V0) {
__t20 = true
goto end_branch_20
} else {

}
}
{
__t20 = false
}
end_branch_20:
var __t_and_30 bool = false
if __t20 {

var __t22 bool
{
var __t_tag_21 gopurs_runtime.Value = gopurs_runtime.Apply2(gopurs_runtime.Box(dictOrd_0.V1), ((*Constructor_Data_Map_Internal_Node)(v_2.UnsafePtr).V5).V2, (*Constructor_Data_Map_Internal_Node)(v_2.UnsafePtr).V2)
if (uint32(__t_tag_21.IntVal) == 380165415) {
__t22 = true
goto end_branch_22
} else {

}
}
{
__t22 = false
}
end_branch_22:
var __t_and_29 bool = false
if __t22 {

var __t23 bool
{
if ((*Constructor_Data_Map_Internal_Node)(v_2.UnsafePtr).V0) > (((*Constructor_Data_Map_Internal_Node)(v_2.UnsafePtr).V4).V0) {
__t23 = true
goto end_branch_23
} else {

}
}
{
__t23 = false
}
end_branch_23:
var __t_and_28 bool = false
if __t23 {

var __t25 bool
{
var __t_tag_24 gopurs_runtime.Value = gopurs_runtime.Apply2(gopurs_runtime.Box(dictOrd_0.V1), ((*Constructor_Data_Map_Internal_Node)(v_2.UnsafePtr).V4).V2, (*Constructor_Data_Map_Internal_Node)(v_2.UnsafePtr).V2)
if (uint32(__t_tag_24.IntVal) == 1527465420) {
__t25 = true
goto end_branch_25
} else {

}
}
{
__t25 = false
}
end_branch_25:
var __t_and_27 bool = false
if __t25 {

var __t26 bool
{
if (gopurs_runtime.Apply(Get_Data_Ord_abs__1599282999(), gopurs_runtime.Int((((*Constructor_Data_Map_Internal_Node)(v_2.UnsafePtr).V5).V0) - (((*Constructor_Data_Map_Internal_Node)(v_2.UnsafePtr).V4).V0))).IntVal) < (2) {
__t26 = true
goto end_branch_26
} else {

}
}
{
__t26 = false
}
end_branch_26:
__t_and_27 = (__t26) && (((((((*Constructor_Data_Map_Internal_Node)(v_2.UnsafePtr).V5).V1) + (((*Constructor_Data_Map_Internal_Node)(v_2.UnsafePtr).V4).V1)) + (1)) == ((*Constructor_Data_Map_Internal_Node)(v_2.UnsafePtr).V1)) && (((gopurs_runtime.Apply(go__go_1_0_71, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v_2.UnsafePtr).V4)}).IntVal) != (0)) && ((gopurs_runtime.Apply(go__go_1_0_71, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v_2.UnsafePtr).V5)}).IntVal) != (0))))
}
__t_and_28 = __t_and_27
}
__t_and_29 = __t_and_28
}
__t_and_30 = __t_and_29
}
__t31 = __t_and_30
goto end_branch_31
} else {

}
}
{
__t31 = (func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal) != (0)
}
end_branch_31:
__t32 = __t31
goto end_branch_32
} else {

}
}
{
__t32 = (func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal) != (0)
}
end_branch_32:
__t33 = __t32
goto end_branch_33
} else {

}
}
{
__t33 = (func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal) != (0)
}
end_branch_33:
return gopurs_runtime.Bool(__t33)
})
return go__go_1_0_71
}

func Call_Data_Map_Internal_catMaybes(dictOrd_0_loop *Constructor_Data_Ord_Ord) gopurs_runtime.Value {
var dictOrd_0 *Constructor_Data_Ord_Ord = dictOrd_0_loop
_ = dictOrd_0
var go__go_1_0_72 gopurs_runtime.Value
_ = go__go_1_0_72
go__go_1_0_72 = gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t4 *Constructor_Data_Map_Internal_Node
{
if (v_2.Type == 9 && v_2.IntVal == 324739070 && v_2.UnsafePtr == nil) {
__t4 = (*Constructor_Data_Map_Internal_Node)(nil)
goto end_branch_4
} else {

}
}
{
if (v_2.Type == 9 && v_2.IntVal == 324739070 && v_2.UnsafePtr != nil) {
var __t3 *Constructor_Data_Map_Internal_Node
{
var __t_tag_1 gopurs_runtime.Value = (*Constructor_Data_Map_Internal_Node)(v_2.UnsafePtr).V3
if (__t_tag_1.Type == 9 && __t_tag_1.IntVal == 930809136 && __t_tag_1.UnsafePtr != nil) {
__t3 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeBalancedNode(), (*Constructor_Data_Map_Internal_Node)(v_2.UnsafePtr).V2, (*Constructor_Data_Maybe_Just)((*Constructor_Data_Map_Internal_Node)(v_2.UnsafePtr).V3.UnsafePtr).V0, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.Apply(go__go_1_0_72, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v_2.UnsafePtr).V4)})))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.Apply(go__go_1_0_72, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v_2.UnsafePtr).V5)})))}))
goto end_branch_3
} else {

}
}
{
var __t_tag_2 gopurs_runtime.Value = (*Constructor_Data_Map_Internal_Node)(v_2.UnsafePtr).V3
if (__t_tag_2.Type == 9 && __t_tag_2.IntVal == 930809136 && __t_tag_2.UnsafePtr == nil) {
__t3 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.UncurriedApp2(Get_Data_Map_Internal_unsafeJoinNodes(), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.Apply(go__go_1_0_72, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v_2.UnsafePtr).V4)})))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.Apply(go__go_1_0_72, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v_2.UnsafePtr).V5)})))}))
goto end_branch_3
} else {

}
}
{
__t3 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_3:
__t4 = __t3
goto end_branch_4
} else {

}
}
{
__t4 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_4:
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(__t4)}
})
return go__go_1_0_72
}

func Call_Data_Map_Internal_applyMap(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
// TAST (Let): compare_1_2 -> gopurs_runtime.Value
compare_1_2 := gopurs_runtime.RecordGet(dictOrd_0, "compare")
_ = compare_1_2
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_3_0_73 gopurs_runtime.Value
_ = go__go_3_0_73
go__go_3_0_73 = gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 *Constructor_Data_Map_Internal_Node
{
if (v_4.Type == 9 && v_4.IntVal == 324739070 && v_4.UnsafePtr == nil) {
__t1 = (*Constructor_Data_Map_Internal_Node)(nil)
goto end_branch_1
} else {

}
}
{
if (v_4.Type == 9 && v_4.IntVal == 324739070 && v_4.UnsafePtr != nil) {
__t1 = &Constructor_Data_Map_Internal_Node{1, (*Constructor_Data_Map_Internal_Node)(v_4.UnsafePtr).V0, (*Constructor_Data_Map_Internal_Node)(v_4.UnsafePtr).V1, (*Constructor_Data_Map_Internal_Node)(v_4.UnsafePtr).V2, gopurs_runtime.Apply(f_2, (*Constructor_Data_Map_Internal_Node)(v_4.UnsafePtr).V3), gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.Apply(go__go_3_0_73, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v_4.UnsafePtr).V4)})), gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.Apply(go__go_3_0_73, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v_4.UnsafePtr).V5)}))}
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_1:
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(__t1)}
})
return go__go_3_0_73
}))
}), gopurs_runtime.Func(func(m1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m2_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeIntersectionWith(), compare_1_2, gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return x_4
}), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](m1_2))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](m2_3))})))}
})
}))
}

func Call_Data_Map_Internal_bindMap(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
// TAST (Let): compare_1_3 -> gopurs_runtime.Value
compare_1_3 := gopurs_runtime.RecordGet(dictOrd_0, "compare")
_ = compare_1_3
// TAST (Let): applyMap1_1_0 -> gopurs_runtime.Value
applyMap1_1_0 := gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_3_1_74 gopurs_runtime.Value
_ = go__go_3_1_74
go__go_3_1_74 = gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 *Constructor_Data_Map_Internal_Node
{
if (v_4.Type == 9 && v_4.IntVal == 324739070 && v_4.UnsafePtr == nil) {
__t2 = (*Constructor_Data_Map_Internal_Node)(nil)
goto end_branch_2
} else {

}
}
{
if (v_4.Type == 9 && v_4.IntVal == 324739070 && v_4.UnsafePtr != nil) {
__t2 = &Constructor_Data_Map_Internal_Node{1, (*Constructor_Data_Map_Internal_Node)(v_4.UnsafePtr).V0, (*Constructor_Data_Map_Internal_Node)(v_4.UnsafePtr).V1, (*Constructor_Data_Map_Internal_Node)(v_4.UnsafePtr).V2, gopurs_runtime.Apply(f_2, (*Constructor_Data_Map_Internal_Node)(v_4.UnsafePtr).V3), gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.Apply(go__go_3_1_74, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v_4.UnsafePtr).V4)})), gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.Apply(go__go_3_1_74, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v_4.UnsafePtr).V5)}))}
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
return go__go_3_1_74
}))
}), gopurs_runtime.Func(func(m1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m2_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeIntersectionWith(), compare_1_3, gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return x_4
}), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](m1_2))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](m2_3))})))}
})
}))
_ = applyMap1_1_0
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return applyMap1_1_0
}), gopurs_runtime.Func(func(m_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_4_4_75 gopurs_runtime.Value
_ = go__go_4_4_75
go__go_4_4_75 = gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
var __t12 *Constructor_Data_Map_Internal_Node
{
if (v_5.Type == 9 && v_5.IntVal == 324739070 && v_5.UnsafePtr == nil) {
__t12 = (*Constructor_Data_Map_Internal_Node)(nil)
goto end_branch_12
} else {

}
}
{
if (v_5.Type == 9 && v_5.IntVal == 324739070 && v_5.UnsafePtr != nil) {
// TAST (Let): __local_var_6_5 -> gopurs_runtime.Value
__local_var_6_5 := (*Constructor_Data_Map_Internal_Node)(v_5.UnsafePtr).V2
_ = __local_var_6_5
var go__go_7_7_76 gopurs_runtime.Value
go__go_7_7_76 = gopurs_runtime.Func(func(v_8_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_8_loop gopurs_runtime.Value = v_8_loop_val
go__go_7_7_76:
for {
if false { continue go__go_7_7_76 }
var v_8 gopurs_runtime.Value = v_8_loop
_ = v_8
var __t10 *Constructor_Data_Maybe_Just
{
if (v_8.Type == 9 && v_8.IntVal == 324739070 && v_8.UnsafePtr == nil) {
__t10 = (*Constructor_Data_Maybe_Just)(nil)
goto end_branch_10
} else {

}
}
{
if (v_8.Type == 9 && v_8.IntVal == 324739070 && v_8.UnsafePtr != nil) {
// TAST (Let): v1_9_8 -> gopurs_runtime.Value
v1_9_8 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictOrd_0, "compare"), __local_var_6_5, (*Constructor_Data_Map_Internal_Node)(v_8.UnsafePtr).V2)
_ = v1_9_8
var __t9 *Constructor_Data_Maybe_Just
{
if (uint32(v1_9_8.IntVal) == 1527465420) {
v_8_loop = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v_8.UnsafePtr).V4)}
continue go__go_7_7_76
__t9 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Value{})
goto end_branch_9
} else {

}
}
{
if (uint32(v1_9_8.IntVal) == 380165415) {
v_8_loop = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v_8.UnsafePtr).V5)}
continue go__go_7_7_76
__t9 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Value{})
goto end_branch_9
} else {

}
}
{
if (uint32(v1_9_8.IntVal) == 902936544) {
__t9 = &Constructor_Data_Maybe_Just{1, (*Constructor_Data_Map_Internal_Node)(v_8.UnsafePtr).V3}
goto end_branch_9
} else {

}
}
{
__t9 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_9:
__t10 = __t9
goto end_branch_10
} else {

}
}
{
__t10 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_10:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t10)}
}
}()
})
// TAST (Let): v2_7_6 -> gopurs_runtime.Value
v2_7_6 := gopurs_runtime.Apply(go__go_7_7_76, gopurs_runtime.Apply(f_3, (*Constructor_Data_Map_Internal_Node)(v_5.UnsafePtr).V3))
_ = v2_7_6
var __t11 *Constructor_Data_Map_Internal_Node
{
if (v2_7_6.Type == 9 && v2_7_6.IntVal == 930809136 && v2_7_6.UnsafePtr != nil) {
__t11 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeBalancedNode(), (*Constructor_Data_Map_Internal_Node)(v_5.UnsafePtr).V2, (*Constructor_Data_Maybe_Just)(v2_7_6.UnsafePtr).V0, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.Apply(go__go_4_4_75, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v_5.UnsafePtr).V4)})))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.Apply(go__go_4_4_75, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v_5.UnsafePtr).V5)})))}))
goto end_branch_11
} else {

}
}
{
if (v2_7_6.Type == 9 && v2_7_6.IntVal == 930809136 && v2_7_6.UnsafePtr == nil) {
__t11 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.UncurriedApp2(Get_Data_Map_Internal_unsafeJoinNodes(), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.Apply(go__go_4_4_75, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v_5.UnsafePtr).V4)})))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.Apply(go__go_4_4_75, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v_5.UnsafePtr).V5)})))}))
goto end_branch_11
} else {

}
}
{
__t11 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_11:
__t12 = __t11
goto end_branch_12
} else {

}
}
{
__t12 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_12:
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(__t12)}
})
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.Apply(go__go_4_4_75, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](m_2))})))}
})
}))
}

func Call_Data_Map_Internal_anyWithKey(predicate_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var predicate_0 gopurs_runtime.Value = predicate_0_loop
_ = predicate_0
var go__go_1_0_77 gopurs_runtime.Value
_ = go__go_1_0_77
go__go_1_0_77 = gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 bool
{
if (v_2.Type == 9 && v_2.IntVal == 324739070 && v_2.UnsafePtr == nil) {
__t1 = false
goto end_branch_1
} else {

}
}
{
if (v_2.Type == 9 && v_2.IntVal == 324739070 && v_2.UnsafePtr != nil) {
__t1 = ((gopurs_runtime.Apply2(predicate_0, (*Constructor_Data_Map_Internal_Node)(v_2.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node)(v_2.UnsafePtr).V3).IntVal) != (0)) || (((gopurs_runtime.Apply(go__go_1_0_77, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v_2.UnsafePtr).V4)}).IntVal) != (0)) || ((gopurs_runtime.Apply(go__go_1_0_77, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v_2.UnsafePtr).V5)}).IntVal) != (0)))
goto end_branch_1
} else {

}
}
{
__t1 = (func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal) != (0)
}
end_branch_1:
return gopurs_runtime.Bool(__t1)
})
return go__go_1_0_77
}

func Call_Data_Map_Internal_any(predicate_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var predicate_0 gopurs_runtime.Value = predicate_0_loop
_ = predicate_0
var go__go_1_0_78 gopurs_runtime.Value
_ = go__go_1_0_78
go__go_1_0_78 = gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 bool
{
if (v_2.Type == 9 && v_2.IntVal == 324739070 && v_2.UnsafePtr == nil) {
__t1 = false
goto end_branch_1
} else {

}
}
{
if (v_2.Type == 9 && v_2.IntVal == 324739070 && v_2.UnsafePtr != nil) {
__t1 = ((gopurs_runtime.Apply(predicate_0, (*Constructor_Data_Map_Internal_Node)(v_2.UnsafePtr).V3).IntVal) != (0)) || (((gopurs_runtime.Apply(go__go_1_0_78, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v_2.UnsafePtr).V4)}).IntVal) != (0)) || ((gopurs_runtime.Apply(go__go_1_0_78, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v_2.UnsafePtr).V5)}).IntVal) != (0)))
goto end_branch_1
} else {

}
}
{
__t1 = (func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal) != (0)
}
end_branch_1:
return gopurs_runtime.Bool(__t1)
})
return go__go_1_0_78
}

func Call_Data_Map_Internal_alter(dictOrd_0_loop *Constructor_Data_Ord_Ord) gopurs_runtime.Value {
var dictOrd_0 *Constructor_Data_Ord_Ord = dictOrd_0_loop
_ = dictOrd_0
// TAST (Let): compare_1_0 -> gopurs_runtime.Value
compare_1_0 := gopurs_runtime.Box(dictOrd_0.V1)
_ = compare_1_0
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): v_5_1 -> gopurs_runtime.Value
v_5_1 := gopurs_runtime.UncurriedApp3(Get_Data_Map_Internal_unsafeSplit(), compare_1_0, k_3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](m_4))})
_ = v_5_1
// TAST (Let): v2_6_2 -> *Constructor_Data_Maybe_Just
v2_6_2 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply(f_2, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Split)(v_5_1.UnsafePtr).V0)}))
_ = v2_6_2
var __t3 *Constructor_Data_Map_Internal_Node
{
if (v2_6_2 == nil) {
__t3 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.UncurriedApp2(Get_Data_Map_Internal_unsafeJoinNodes(), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Split)(v_5_1.UnsafePtr).V1)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Split)(v_5_1.UnsafePtr).V2)}))
goto end_branch_3
} else {

}
}
{
if (v2_6_2 != nil) {
__t3 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeBalancedNode(), k_3, (v2_6_2).V0, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Split)(v_5_1.UnsafePtr).V1)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Split)(v_5_1.UnsafePtr).V2)}))
goto end_branch_3
} else {

}
}
{
__t3 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_3:
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(__t3)}
})
})
})
}

func Call_Data_Map_Internal_altMap(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
// TAST (Let): compare_1_2 -> gopurs_runtime.Value
compare_1_2 := gopurs_runtime.RecordGet(dictOrd_0, "compare")
_ = compare_1_2
return gopurs_runtime.RecordDict2("Functor0", "alt", gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_3_0_79 gopurs_runtime.Value
_ = go__go_3_0_79
go__go_3_0_79 = gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 *Constructor_Data_Map_Internal_Node
{
if (v_4.Type == 9 && v_4.IntVal == 324739070 && v_4.UnsafePtr == nil) {
__t1 = (*Constructor_Data_Map_Internal_Node)(nil)
goto end_branch_1
} else {

}
}
{
if (v_4.Type == 9 && v_4.IntVal == 324739070 && v_4.UnsafePtr != nil) {
__t1 = &Constructor_Data_Map_Internal_Node{1, (*Constructor_Data_Map_Internal_Node)(v_4.UnsafePtr).V0, (*Constructor_Data_Map_Internal_Node)(v_4.UnsafePtr).V1, (*Constructor_Data_Map_Internal_Node)(v_4.UnsafePtr).V2, gopurs_runtime.Apply(f_2, (*Constructor_Data_Map_Internal_Node)(v_4.UnsafePtr).V3), gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.Apply(go__go_3_0_79, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v_4.UnsafePtr).V4)})), gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.Apply(go__go_3_0_79, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v_4.UnsafePtr).V5)}))}
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_1:
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(__t1)}
})
return go__go_3_0_79
}))
}), gopurs_runtime.Func(func(m1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m2_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeUnionWith(), compare_1_2, Get_Data_Function_go__const(), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](m1_2))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](m2_3))})))}
})
}))
}

func Call_Data_Map_Internal_plusMap(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
// TAST (Let): compare_1_3 -> gopurs_runtime.Value
compare_1_3 := gopurs_runtime.RecordGet(dictOrd_0, "compare")
_ = compare_1_3
// TAST (Let): altMap1_1_0 -> gopurs_runtime.Value
altMap1_1_0 := gopurs_runtime.RecordDict2("Functor0", "alt", gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_3_1_80 gopurs_runtime.Value
_ = go__go_3_1_80
go__go_3_1_80 = gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 *Constructor_Data_Map_Internal_Node
{
if (v_4.Type == 9 && v_4.IntVal == 324739070 && v_4.UnsafePtr == nil) {
__t2 = (*Constructor_Data_Map_Internal_Node)(nil)
goto end_branch_2
} else {

}
}
{
if (v_4.Type == 9 && v_4.IntVal == 324739070 && v_4.UnsafePtr != nil) {
__t2 = &Constructor_Data_Map_Internal_Node{1, (*Constructor_Data_Map_Internal_Node)(v_4.UnsafePtr).V0, (*Constructor_Data_Map_Internal_Node)(v_4.UnsafePtr).V1, (*Constructor_Data_Map_Internal_Node)(v_4.UnsafePtr).V2, gopurs_runtime.Apply(f_2, (*Constructor_Data_Map_Internal_Node)(v_4.UnsafePtr).V3), gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.Apply(go__go_3_1_80, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v_4.UnsafePtr).V4)})), gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.Apply(go__go_3_1_80, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v_4.UnsafePtr).V5)}))}
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
return go__go_3_1_80
}))
}), gopurs_runtime.Func(func(m1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m2_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeUnionWith(), compare_1_3, Get_Data_Function_go__const(), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](m1_2))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](m2_3))})))}
})
}))
_ = altMap1_1_0
return gopurs_runtime.RecordDict2("Alt0", "empty", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return altMap1_1_0
}), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(nil))})
}

func Call_Data_Map_Internal_alter__2325420954(dictOrd_0_loop *Constructor_Data_Ord_Ord) gopurs_runtime.Value {
var dictOrd_0 *Constructor_Data_Ord_Ord = dictOrd_0_loop
_ = dictOrd_0
// TAST (Let): compare_1_0 -> gopurs_runtime.Value
compare_1_0 := gopurs_runtime.Box(dictOrd_0.V1)
_ = compare_1_0
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): v_5_1 -> gopurs_runtime.Value
v_5_1 := gopurs_runtime.UncurriedApp3(Get_Data_Map_Internal_unsafeSplit(), compare_1_0, k_3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](m_4))})
_ = v_5_1
// TAST (Let): v2_6_2 -> *Constructor_Data_Maybe_Just
v2_6_2 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply(f_2, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Split)(v_5_1.UnsafePtr).V0)}))
_ = v2_6_2
var __t3 *Constructor_Data_Map_Internal_Node
{
if (v2_6_2 == nil) {
__t3 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.UncurriedApp2(Get_Data_Map_Internal_unsafeJoinNodes(), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Split)(v_5_1.UnsafePtr).V1)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Split)(v_5_1.UnsafePtr).V2)}))
goto end_branch_3
} else {

}
}
{
if (v2_6_2 != nil) {
__t3 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeBalancedNode(), k_3, (v2_6_2).V0, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Split)(v_5_1.UnsafePtr).V1)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Split)(v_5_1.UnsafePtr).V2)}))
goto end_branch_3
} else {

}
}
{
__t3 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_3:
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(__t3)}
})
})
})
}

func Call_Data_Map_Internal_alter__1204655226(dictOrd_0_loop *Constructor_Data_Ord_Ord) gopurs_runtime.Value {
var dictOrd_0 *Constructor_Data_Ord_Ord = dictOrd_0_loop
_ = dictOrd_0
// TAST (Let): compare_1_0 -> gopurs_runtime.Value
compare_1_0 := gopurs_runtime.Box(dictOrd_0.V1)
_ = compare_1_0
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): v_5_1 -> gopurs_runtime.Value
v_5_1 := gopurs_runtime.UncurriedApp3(Get_Data_Map_Internal_unsafeSplit(), compare_1_0, k_3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](m_4))})
_ = v_5_1
// TAST (Let): v2_6_2 -> *Constructor_Data_Maybe_Just
v2_6_2 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply(f_2, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Split)(v_5_1.UnsafePtr).V0)}))
_ = v2_6_2
var __t3 *Constructor_Data_Map_Internal_Node
{
if (v2_6_2 == nil) {
__t3 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.UncurriedApp2(Get_Data_Map_Internal_unsafeJoinNodes(), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Split)(v_5_1.UnsafePtr).V1)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Split)(v_5_1.UnsafePtr).V2)}))
goto end_branch_3
} else {

}
}
{
if (v2_6_2 != nil) {
__t3 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeBalancedNode(), k_3, (v2_6_2).V0, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Split)(v_5_1.UnsafePtr).V1)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Split)(v_5_1.UnsafePtr).V2)}))
goto end_branch_3
} else {

}
}
{
__t3 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_3:
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(__t3)}
})
})
})
}

func Call_Data_Map_Internal_findMax__2266220649(v_0_loop *Constructor_Data_Map_Internal_Node) *Constructor_Data_Maybe_Just {
var v_0 *Constructor_Data_Map_Internal_Node = v_0_loop
_ = v_0
var __t2 *Constructor_Data_Maybe_Just
{
if (v_0 == nil) {
__t2 = (*Constructor_Data_Maybe_Just)(nil)
goto end_branch_2
} else {

}
}
{
if (v_0 != nil) {
var __t1 *Constructor_Data_Maybe_Just
{
var __t_tag_0 *Constructor_Data_Map_Internal_Node = (v_0).V5
if (__t_tag_0 == nil) {
__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, gopurs_runtime.RecordDict2("key", "value", (v_0).V2, (v_0).V3)})})
goto end_branch_1
} else {

}
}
{
__t1 = Call_Data_Map_Internal_findMax((v_0).V5)
}
end_branch_1:
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t1)})
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_2:
return gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t2)})
}

func Call_Data_Map_Internal_findMax__528468393(v_0_loop *Constructor_Data_Map_Internal_Node) *Constructor_Data_Maybe_Just {
var v_0 *Constructor_Data_Map_Internal_Node = v_0_loop
_ = v_0
var __t2 *Constructor_Data_Maybe_Just
{
if (v_0 == nil) {
__t2 = (*Constructor_Data_Maybe_Just)(nil)
goto end_branch_2
} else {

}
}
{
if (v_0 != nil) {
var __t1 *Constructor_Data_Maybe_Just
{
var __t_tag_0 *Constructor_Data_Map_Internal_Node = (v_0).V5
if (__t_tag_0 == nil) {
__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, gopurs_runtime.RecordDict2("key", "value", (v_0).V2, (v_0).V3)})})
goto end_branch_1
} else {

}
}
{
__t1 = Call_Data_Map_Internal_findMax((v_0).V5)
}
end_branch_1:
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t1)})
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_2:
return gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t2)})
}

func Call_Data_Map_Internal_findMin__2266220649(v_0_loop *Constructor_Data_Map_Internal_Node) *Constructor_Data_Maybe_Just {
var v_0 *Constructor_Data_Map_Internal_Node = v_0_loop
_ = v_0
var __t2 *Constructor_Data_Maybe_Just
{
if (v_0 == nil) {
__t2 = (*Constructor_Data_Maybe_Just)(nil)
goto end_branch_2
} else {

}
}
{
if (v_0 != nil) {
var __t1 *Constructor_Data_Maybe_Just
{
var __t_tag_0 *Constructor_Data_Map_Internal_Node = (v_0).V4
if (__t_tag_0 == nil) {
__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, gopurs_runtime.RecordDict2("key", "value", (v_0).V2, (v_0).V3)})})
goto end_branch_1
} else {

}
}
{
__t1 = Call_Data_Map_Internal_findMin((v_0).V4)
}
end_branch_1:
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t1)})
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_2:
return gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t2)})
}

func Call_Data_Map_Internal_findMin__528468393(v_0_loop *Constructor_Data_Map_Internal_Node) *Constructor_Data_Maybe_Just {
var v_0 *Constructor_Data_Map_Internal_Node = v_0_loop
_ = v_0
var __t2 *Constructor_Data_Maybe_Just
{
if (v_0 == nil) {
__t2 = (*Constructor_Data_Maybe_Just)(nil)
goto end_branch_2
} else {

}
}
{
if (v_0 != nil) {
var __t1 *Constructor_Data_Maybe_Just
{
var __t_tag_0 *Constructor_Data_Map_Internal_Node = (v_0).V4
if (__t_tag_0 == nil) {
__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, gopurs_runtime.RecordDict2("key", "value", (v_0).V2, (v_0).V3)})})
goto end_branch_1
} else {

}
}
{
__t1 = Call_Data_Map_Internal_findMin((v_0).V4)
}
end_branch_1:
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t1)})
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_2:
return gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t2)})
}

func Call_Data_Map_Internal_foldSubmapBy__3050108409(dictOrd_0_loop *Constructor_Data_Ord_Ord, appendFn_1_loop gopurs_runtime.Value, memptyValue_2_loop gopurs_runtime.Value, kmin_3_loop *Constructor_Data_Maybe_Just, kmax_4_loop *Constructor_Data_Maybe_Just, f_5_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 *Constructor_Data_Ord_Ord = dictOrd_0_loop
_ = dictOrd_0
var appendFn_1 gopurs_runtime.Value = appendFn_1_loop
_ = appendFn_1
var memptyValue_2 gopurs_runtime.Value = memptyValue_2_loop
_ = memptyValue_2
var kmin_3 *Constructor_Data_Maybe_Just = kmin_3_loop
_ = kmin_3
var kmax_4 *Constructor_Data_Maybe_Just = kmax_4_loop
_ = kmax_4
var f_5 gopurs_runtime.Value = f_5_loop
_ = f_5
var __t4 gopurs_runtime.Value
{
if (kmin_3 != nil) {
// TAST (Let): __local_var_6_1 -> gopurs_runtime.Value
__local_var_6_1 := (kmin_3).V0
_ = __local_var_6_1
__t4 = gopurs_runtime.Func(func(k_7 gopurs_runtime.Value) gopurs_runtime.Value {
var __t3 bool
{
var __t_tag_2 gopurs_runtime.Value = gopurs_runtime.Apply2(gopurs_runtime.Box(dictOrd_0.V1), k_7, __local_var_6_1)
if (uint32(__t_tag_2.IntVal) == 1527465420) {
__t3 = true
goto end_branch_3
} else {

}
}
{
__t3 = false
}
end_branch_3:
return gopurs_runtime.Bool(__t3)
})
goto end_branch_4
} else {

}
}
{
if (kmin_3 == nil) {
__t4 = gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(false)
})
goto end_branch_4
} else {

}
}
{
__t4 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_4:
// TAST (Let): tooSmall_6_0 -> gopurs_runtime.Value
tooSmall_6_0 := __t4
_ = tooSmall_6_0
var __t9 gopurs_runtime.Value
{
if (kmax_4 != nil) {
// TAST (Let): __local_var_7_6 -> gopurs_runtime.Value
__local_var_7_6 := (kmax_4).V0
_ = __local_var_7_6
__t9 = gopurs_runtime.Func(func(k_8 gopurs_runtime.Value) gopurs_runtime.Value {
var __t8 bool
{
var __t_tag_7 gopurs_runtime.Value = gopurs_runtime.Apply2(gopurs_runtime.Box(dictOrd_0.V1), k_8, __local_var_7_6)
if (uint32(__t_tag_7.IntVal) == 380165415) {
__t8 = true
goto end_branch_8
} else {

}
}
{
__t8 = false
}
end_branch_8:
return gopurs_runtime.Bool(__t8)
})
goto end_branch_9
} else {

}
}
{
if (kmax_4 == nil) {
__t9 = gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(false)
})
goto end_branch_9
} else {

}
}
{
__t9 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_9:
// TAST (Let): tooLarge_7_5 -> gopurs_runtime.Value
tooLarge_7_5 := __t9
_ = tooLarge_7_5
var __t26 gopurs_runtime.Value
{
if (kmin_3 != nil) {
var __t21 gopurs_runtime.Value
{
if (kmax_4 != nil) {
// TAST (Let): __local_var_8_11 -> gopurs_runtime.Value
__local_var_8_11 := (kmax_4).V0
_ = __local_var_8_11
// TAST (Let): __local_var_9_12 -> gopurs_runtime.Value
__local_var_9_12 := (kmin_3).V0
_ = __local_var_9_12
__t21 = gopurs_runtime.Func(func(k_10 gopurs_runtime.Value) gopurs_runtime.Value {
var __t14 bool
{
var __t_tag_13 gopurs_runtime.Value = gopurs_runtime.Apply2(gopurs_runtime.Box(dictOrd_0.V1), __local_var_9_12, k_10)
if (uint32(__t_tag_13.IntVal) == 380165415) {
__t14 = false
goto end_branch_14
} else {

}
}
{
__t14 = true
}
end_branch_14:
var __t_and_17 bool = false
if __t14 {

var __t16 bool
{
var __t_tag_15 gopurs_runtime.Value = gopurs_runtime.Apply2(gopurs_runtime.Box(dictOrd_0.V1), k_10, __local_var_8_11)
if (uint32(__t_tag_15.IntVal) == 380165415) {
__t16 = false
goto end_branch_16
} else {

}
}
{
__t16 = true
}
end_branch_16:
__t_and_17 = __t16
}
return gopurs_runtime.Bool(__t_and_17)
})
goto end_branch_21
} else {

}
}
{
if (kmax_4 == nil) {
// TAST (Let): __local_var_8_18 -> gopurs_runtime.Value
__local_var_8_18 := (kmin_3).V0
_ = __local_var_8_18
__t21 = gopurs_runtime.Func(func(k_9 gopurs_runtime.Value) gopurs_runtime.Value {
var __t20 bool
{
var __t_tag_19 gopurs_runtime.Value = gopurs_runtime.Apply2(gopurs_runtime.Box(dictOrd_0.V1), __local_var_8_18, k_9)
if (uint32(__t_tag_19.IntVal) == 380165415) {
__t20 = false
goto end_branch_20
} else {

}
}
{
__t20 = true
}
end_branch_20:
return gopurs_runtime.Bool(__t20)
})
goto end_branch_21
} else {

}
}
{
__t21 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_21:
__t26 = __t21
goto end_branch_26
} else {

}
}
{
if (kmin_3 == nil) {
var __t25 gopurs_runtime.Value
{
if (kmax_4 != nil) {
// TAST (Let): __local_var_8_22 -> gopurs_runtime.Value
__local_var_8_22 := (kmax_4).V0
_ = __local_var_8_22
__t25 = gopurs_runtime.Func(func(k_9 gopurs_runtime.Value) gopurs_runtime.Value {
var __t24 bool
{
var __t_tag_23 gopurs_runtime.Value = gopurs_runtime.Apply2(gopurs_runtime.Box(dictOrd_0.V1), k_9, __local_var_8_22)
if (uint32(__t_tag_23.IntVal) == 380165415) {
__t24 = false
goto end_branch_24
} else {

}
}
{
__t24 = true
}
end_branch_24:
return gopurs_runtime.Bool(__t24)
})
goto end_branch_25
} else {

}
}
{
if (kmax_4 == nil) {
__t25 = gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(true)
})
goto end_branch_25
} else {

}
}
{
__t25 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_25:
__t26 = __t25
goto end_branch_26
} else {

}
}
{
__t26 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_26:
// TAST (Let): inBounds_8_10 -> gopurs_runtime.Value
inBounds_8_10 := __t26
_ = inBounds_8_10
var go__go_9_27_81 gopurs_runtime.Value
_ = go__go_9_27_81
go__go_9_27_81 = gopurs_runtime.Func(func(v_10 gopurs_runtime.Value) gopurs_runtime.Value {
var __t31 gopurs_runtime.Value
{
if (v_10.Type == 9 && v_10.IntVal == 324739070 && v_10.UnsafePtr == nil) {
__t31 = memptyValue_2
goto end_branch_31
} else {

}
}
{
if (v_10.Type == 9 && v_10.IntVal == 324739070 && v_10.UnsafePtr != nil) {
var __t28 gopurs_runtime.Value
{
if (gopurs_runtime.Apply(tooSmall_6_0, (*Constructor_Data_Map_Internal_Node)(v_10.UnsafePtr).V2).IntVal) != (0) {
__t28 = memptyValue_2
goto end_branch_28
} else {

}
}
{
__t28 = gopurs_runtime.Apply(go__go_9_27_81, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v_10.UnsafePtr).V4)})
}
end_branch_28:
var __t29 gopurs_runtime.Value
{
if (gopurs_runtime.Apply(inBounds_8_10, (*Constructor_Data_Map_Internal_Node)(v_10.UnsafePtr).V2).IntVal) != (0) {
__t29 = gopurs_runtime.Apply2(f_5, (*Constructor_Data_Map_Internal_Node)(v_10.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node)(v_10.UnsafePtr).V3)
goto end_branch_29
} else {

}
}
{
__t29 = memptyValue_2
}
end_branch_29:
var __t30 gopurs_runtime.Value
{
if (gopurs_runtime.Apply(tooLarge_7_5, (*Constructor_Data_Map_Internal_Node)(v_10.UnsafePtr).V2).IntVal) != (0) {
__t30 = memptyValue_2
goto end_branch_30
} else {

}
}
{
__t30 = gopurs_runtime.Apply(go__go_9_27_81, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v_10.UnsafePtr).V5)})
}
end_branch_30:
__t31 = gopurs_runtime.Apply2(appendFn_1, gopurs_runtime.Apply2(appendFn_1, __t28, __t29), __t30)
goto end_branch_31
} else {

}
}
{
__t31 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_31:
return __t31
})
return go__go_9_27_81
}

func Call_Data_Map_Internal_foldSubmapBy__3128450809(dictOrd_0_loop *Constructor_Data_Ord_Ord, appendFn_1_loop gopurs_runtime.Value, memptyValue_2_loop *Constructor_Data_Map_Internal_Node, kmin_3_loop *Constructor_Data_Maybe_Just, kmax_4_loop *Constructor_Data_Maybe_Just, f_5_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 *Constructor_Data_Ord_Ord = dictOrd_0_loop
_ = dictOrd_0
var appendFn_1 gopurs_runtime.Value = appendFn_1_loop
_ = appendFn_1
var memptyValue_2 *Constructor_Data_Map_Internal_Node = memptyValue_2_loop
_ = memptyValue_2
var kmin_3 *Constructor_Data_Maybe_Just = kmin_3_loop
_ = kmin_3
var kmax_4 *Constructor_Data_Maybe_Just = kmax_4_loop
_ = kmax_4
var f_5 gopurs_runtime.Value = f_5_loop
_ = f_5
var __t4 gopurs_runtime.Value
{
if (kmin_3 != nil) {
// TAST (Let): __local_var_6_1 -> gopurs_runtime.Value
__local_var_6_1 := (kmin_3).V0
_ = __local_var_6_1
__t4 = gopurs_runtime.Func(func(k_7 gopurs_runtime.Value) gopurs_runtime.Value {
var __t3 bool
{
var __t_tag_2 gopurs_runtime.Value = gopurs_runtime.Apply2(gopurs_runtime.Box(dictOrd_0.V1), k_7, __local_var_6_1)
if (uint32(__t_tag_2.IntVal) == 1527465420) {
__t3 = true
goto end_branch_3
} else {

}
}
{
__t3 = false
}
end_branch_3:
return gopurs_runtime.Bool(__t3)
})
goto end_branch_4
} else {

}
}
{
if (kmin_3 == nil) {
__t4 = gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(false)
})
goto end_branch_4
} else {

}
}
{
__t4 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_4:
// TAST (Let): tooSmall_6_0 -> gopurs_runtime.Value
tooSmall_6_0 := __t4
_ = tooSmall_6_0
var __t9 gopurs_runtime.Value
{
if (kmax_4 != nil) {
// TAST (Let): __local_var_7_6 -> gopurs_runtime.Value
__local_var_7_6 := (kmax_4).V0
_ = __local_var_7_6
__t9 = gopurs_runtime.Func(func(k_8 gopurs_runtime.Value) gopurs_runtime.Value {
var __t8 bool
{
var __t_tag_7 gopurs_runtime.Value = gopurs_runtime.Apply2(gopurs_runtime.Box(dictOrd_0.V1), k_8, __local_var_7_6)
if (uint32(__t_tag_7.IntVal) == 380165415) {
__t8 = true
goto end_branch_8
} else {

}
}
{
__t8 = false
}
end_branch_8:
return gopurs_runtime.Bool(__t8)
})
goto end_branch_9
} else {

}
}
{
if (kmax_4 == nil) {
__t9 = gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(false)
})
goto end_branch_9
} else {

}
}
{
__t9 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_9:
// TAST (Let): tooLarge_7_5 -> gopurs_runtime.Value
tooLarge_7_5 := __t9
_ = tooLarge_7_5
var __t26 gopurs_runtime.Value
{
if (kmin_3 != nil) {
var __t21 gopurs_runtime.Value
{
if (kmax_4 != nil) {
// TAST (Let): __local_var_8_11 -> gopurs_runtime.Value
__local_var_8_11 := (kmax_4).V0
_ = __local_var_8_11
// TAST (Let): __local_var_9_12 -> gopurs_runtime.Value
__local_var_9_12 := (kmin_3).V0
_ = __local_var_9_12
__t21 = gopurs_runtime.Func(func(k_10 gopurs_runtime.Value) gopurs_runtime.Value {
var __t14 bool
{
var __t_tag_13 gopurs_runtime.Value = gopurs_runtime.Apply2(gopurs_runtime.Box(dictOrd_0.V1), __local_var_9_12, k_10)
if (uint32(__t_tag_13.IntVal) == 380165415) {
__t14 = false
goto end_branch_14
} else {

}
}
{
__t14 = true
}
end_branch_14:
var __t_and_17 bool = false
if __t14 {

var __t16 bool
{
var __t_tag_15 gopurs_runtime.Value = gopurs_runtime.Apply2(gopurs_runtime.Box(dictOrd_0.V1), k_10, __local_var_8_11)
if (uint32(__t_tag_15.IntVal) == 380165415) {
__t16 = false
goto end_branch_16
} else {

}
}
{
__t16 = true
}
end_branch_16:
__t_and_17 = __t16
}
return gopurs_runtime.Bool(__t_and_17)
})
goto end_branch_21
} else {

}
}
{
if (kmax_4 == nil) {
// TAST (Let): __local_var_8_18 -> gopurs_runtime.Value
__local_var_8_18 := (kmin_3).V0
_ = __local_var_8_18
__t21 = gopurs_runtime.Func(func(k_9 gopurs_runtime.Value) gopurs_runtime.Value {
var __t20 bool
{
var __t_tag_19 gopurs_runtime.Value = gopurs_runtime.Apply2(gopurs_runtime.Box(dictOrd_0.V1), __local_var_8_18, k_9)
if (uint32(__t_tag_19.IntVal) == 380165415) {
__t20 = false
goto end_branch_20
} else {

}
}
{
__t20 = true
}
end_branch_20:
return gopurs_runtime.Bool(__t20)
})
goto end_branch_21
} else {

}
}
{
__t21 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_21:
__t26 = __t21
goto end_branch_26
} else {

}
}
{
if (kmin_3 == nil) {
var __t25 gopurs_runtime.Value
{
if (kmax_4 != nil) {
// TAST (Let): __local_var_8_22 -> gopurs_runtime.Value
__local_var_8_22 := (kmax_4).V0
_ = __local_var_8_22
__t25 = gopurs_runtime.Func(func(k_9 gopurs_runtime.Value) gopurs_runtime.Value {
var __t24 bool
{
var __t_tag_23 gopurs_runtime.Value = gopurs_runtime.Apply2(gopurs_runtime.Box(dictOrd_0.V1), k_9, __local_var_8_22)
if (uint32(__t_tag_23.IntVal) == 380165415) {
__t24 = false
goto end_branch_24
} else {

}
}
{
__t24 = true
}
end_branch_24:
return gopurs_runtime.Bool(__t24)
})
goto end_branch_25
} else {

}
}
{
if (kmax_4 == nil) {
__t25 = gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(true)
})
goto end_branch_25
} else {

}
}
{
__t25 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_25:
__t26 = __t25
goto end_branch_26
} else {

}
}
{
__t26 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_26:
// TAST (Let): inBounds_8_10 -> gopurs_runtime.Value
inBounds_8_10 := __t26
_ = inBounds_8_10
var go__go_9_27_82 gopurs_runtime.Value
_ = go__go_9_27_82
go__go_9_27_82 = gopurs_runtime.Func(func(v_10 gopurs_runtime.Value) gopurs_runtime.Value {
var __t31 *Constructor_Data_Map_Internal_Node
{
if (v_10.Type == 9 && v_10.IntVal == 324739070 && v_10.UnsafePtr == nil) {
__t31 = memptyValue_2
goto end_branch_31
} else {

}
}
{
if (v_10.Type == 9 && v_10.IntVal == 324739070 && v_10.UnsafePtr != nil) {
var __t28 *Constructor_Data_Map_Internal_Node
{
if (gopurs_runtime.Apply(tooSmall_6_0, (*Constructor_Data_Map_Internal_Node)(v_10.UnsafePtr).V2).IntVal) != (0) {
__t28 = memptyValue_2
goto end_branch_28
} else {

}
}
{
__t28 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.Apply(go__go_9_27_82, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v_10.UnsafePtr).V4)}))
}
end_branch_28:
var __t29 *Constructor_Data_Map_Internal_Node
{
if (gopurs_runtime.Apply(inBounds_8_10, (*Constructor_Data_Map_Internal_Node)(v_10.UnsafePtr).V2).IntVal) != (0) {
__t29 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.Apply2(f_5, (*Constructor_Data_Map_Internal_Node)(v_10.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node)(v_10.UnsafePtr).V3))
goto end_branch_29
} else {

}
}
{
__t29 = memptyValue_2
}
end_branch_29:
var __t30 *Constructor_Data_Map_Internal_Node
{
if (gopurs_runtime.Apply(tooLarge_7_5, (*Constructor_Data_Map_Internal_Node)(v_10.UnsafePtr).V2).IntVal) != (0) {
__t30 = memptyValue_2
goto end_branch_30
} else {

}
}
{
__t30 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.Apply(go__go_9_27_82, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v_10.UnsafePtr).V5)}))
}
end_branch_30:
__t31 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.Apply2(appendFn_1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.Apply2(appendFn_1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(__t28)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(__t29)})))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(__t30)}))
goto end_branch_31
} else {

}
}
{
__t31 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_31:
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(__t31)}
})
return go__go_9_27_82
}

func Call_Data_Map_Internal_insert__3204212386(dictOrd_0_loop *Constructor_Data_Ord_Ord, k_1_loop gopurs_runtime.Value, v_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 *Constructor_Data_Ord_Ord = dictOrd_0_loop
_ = dictOrd_0
var k_1 gopurs_runtime.Value = k_1_loop
_ = k_1
var v_2 gopurs_runtime.Value = v_2_loop
_ = v_2
var go__go_3_0_110 gopurs_runtime.Value
_ = go__go_3_0_110
go__go_3_0_110 = gopurs_runtime.Func(func(v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t3 *Constructor_Data_Map_Internal_Node
{
if (v1_4.Type == 9 && v1_4.IntVal == 324739070 && v1_4.UnsafePtr == nil) {
__t3 = &Constructor_Data_Map_Internal_Node{1, 1, 1, k_1, v_2, (*Constructor_Data_Map_Internal_Node)(nil), (*Constructor_Data_Map_Internal_Node)(nil)}
goto end_branch_3
} else {

}
}
{
if (v1_4.Type == 9 && v1_4.IntVal == 324739070 && v1_4.UnsafePtr != nil) {
// TAST (Let): v2_5_1 -> gopurs_runtime.Value
v2_5_1 := gopurs_runtime.Apply2(gopurs_runtime.Box(dictOrd_0.V1), k_1, (*Constructor_Data_Map_Internal_Node)(v1_4.UnsafePtr).V2)
_ = v2_5_1
var __t2 *Constructor_Data_Map_Internal_Node
{
if (uint32(v2_5_1.IntVal) == 1527465420) {
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeBalancedNode(), (*Constructor_Data_Map_Internal_Node)(v1_4.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node)(v1_4.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.Apply(go__go_3_0_110, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v1_4.UnsafePtr).V4)})))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v1_4.UnsafePtr).V5)}))
goto end_branch_2
} else {

}
}
{
if (uint32(v2_5_1.IntVal) == 380165415) {
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeBalancedNode(), (*Constructor_Data_Map_Internal_Node)(v1_4.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node)(v1_4.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v1_4.UnsafePtr).V4)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.Apply(go__go_3_0_110, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v1_4.UnsafePtr).V5)})))}))
goto end_branch_2
} else {

}
}
{
if (uint32(v2_5_1.IntVal) == 902936544) {
__t2 = &Constructor_Data_Map_Internal_Node{1, (*Constructor_Data_Map_Internal_Node)(v1_4.UnsafePtr).V0, (*Constructor_Data_Map_Internal_Node)(v1_4.UnsafePtr).V1, k_1, v_2, (*Constructor_Data_Map_Internal_Node)(v1_4.UnsafePtr).V4, (*Constructor_Data_Map_Internal_Node)(v1_4.UnsafePtr).V5}
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_2:
__t3 = __t2
goto end_branch_3
} else {

}
}
{
__t3 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_3:
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(__t3)}
})
return go__go_3_0_110
}

func Call_Data_Map_Internal_insert__4289641298(dictOrd_0_loop *Constructor_Data_Ord_Ord, k_1_loop gopurs_runtime.Value, v_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 *Constructor_Data_Ord_Ord = dictOrd_0_loop
_ = dictOrd_0
var k_1 gopurs_runtime.Value = k_1_loop
_ = k_1
var v_2 gopurs_runtime.Value = v_2_loop
_ = v_2
var go__go_3_0_111 gopurs_runtime.Value
_ = go__go_3_0_111
go__go_3_0_111 = gopurs_runtime.Func(func(v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t3 *Constructor_Data_Map_Internal_Node
{
if (v1_4.Type == 9 && v1_4.IntVal == 324739070 && v1_4.UnsafePtr == nil) {
__t3 = &Constructor_Data_Map_Internal_Node{1, 1, 1, k_1, v_2, (*Constructor_Data_Map_Internal_Node)(nil), (*Constructor_Data_Map_Internal_Node)(nil)}
goto end_branch_3
} else {

}
}
{
if (v1_4.Type == 9 && v1_4.IntVal == 324739070 && v1_4.UnsafePtr != nil) {
// TAST (Let): v2_5_1 -> gopurs_runtime.Value
v2_5_1 := gopurs_runtime.Apply2(gopurs_runtime.Box(dictOrd_0.V1), k_1, (*Constructor_Data_Map_Internal_Node)(v1_4.UnsafePtr).V2)
_ = v2_5_1
var __t2 *Constructor_Data_Map_Internal_Node
{
if (uint32(v2_5_1.IntVal) == 1527465420) {
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeBalancedNode(), (*Constructor_Data_Map_Internal_Node)(v1_4.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node)(v1_4.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.Apply(go__go_3_0_111, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v1_4.UnsafePtr).V4)})))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v1_4.UnsafePtr).V5)}))
goto end_branch_2
} else {

}
}
{
if (uint32(v2_5_1.IntVal) == 380165415) {
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeBalancedNode(), (*Constructor_Data_Map_Internal_Node)(v1_4.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node)(v1_4.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v1_4.UnsafePtr).V4)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.Apply(go__go_3_0_111, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v1_4.UnsafePtr).V5)})))}))
goto end_branch_2
} else {

}
}
{
if (uint32(v2_5_1.IntVal) == 902936544) {
__t2 = &Constructor_Data_Map_Internal_Node{1, (*Constructor_Data_Map_Internal_Node)(v1_4.UnsafePtr).V0, (*Constructor_Data_Map_Internal_Node)(v1_4.UnsafePtr).V1, k_1, v_2, (*Constructor_Data_Map_Internal_Node)(v1_4.UnsafePtr).V4, (*Constructor_Data_Map_Internal_Node)(v1_4.UnsafePtr).V5}
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_2:
__t3 = __t2
goto end_branch_3
} else {

}
}
{
__t3 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_3:
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(__t3)}
})
return go__go_3_0_111
}

func Call_Data_Map_Internal_insert__2073142786(dictOrd_0_loop *Constructor_Data_Ord_Ord, k_1_loop *Constructor_Data_Maybe_Just, v_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 *Constructor_Data_Ord_Ord = dictOrd_0_loop
_ = dictOrd_0
var k_1 *Constructor_Data_Maybe_Just = k_1_loop
_ = k_1
var v_2 gopurs_runtime.Value = v_2_loop
_ = v_2
var go__go_3_0_112 gopurs_runtime.Value
_ = go__go_3_0_112
go__go_3_0_112 = gopurs_runtime.Func(func(v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t3 *Constructor_Data_Map_Internal_Node
{
if (v1_4.Type == 9 && v1_4.IntVal == 324739070 && v1_4.UnsafePtr == nil) {
__t3 = &Constructor_Data_Map_Internal_Node{1, 1, 1, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(k_1)}, v_2, (*Constructor_Data_Map_Internal_Node)(nil), (*Constructor_Data_Map_Internal_Node)(nil)}
goto end_branch_3
} else {

}
}
{
if (v1_4.Type == 9 && v1_4.IntVal == 324739070 && v1_4.UnsafePtr != nil) {
// TAST (Let): v2_5_1 -> gopurs_runtime.Value
v2_5_1 := gopurs_runtime.Apply2(gopurs_runtime.Box(dictOrd_0.V1), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(k_1)}, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just]((*Constructor_Data_Map_Internal_Node)(v1_4.UnsafePtr).V2))})
_ = v2_5_1
var __t2 *Constructor_Data_Map_Internal_Node
{
if (uint32(v2_5_1.IntVal) == 1527465420) {
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeBalancedNode(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just]((*Constructor_Data_Map_Internal_Node)(v1_4.UnsafePtr).V2))}, (*Constructor_Data_Map_Internal_Node)(v1_4.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.Apply(go__go_3_0_112, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v1_4.UnsafePtr).V4)})))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v1_4.UnsafePtr).V5)}))
goto end_branch_2
} else {

}
}
{
if (uint32(v2_5_1.IntVal) == 380165415) {
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeBalancedNode(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just]((*Constructor_Data_Map_Internal_Node)(v1_4.UnsafePtr).V2))}, (*Constructor_Data_Map_Internal_Node)(v1_4.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v1_4.UnsafePtr).V4)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.Apply(go__go_3_0_112, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v1_4.UnsafePtr).V5)})))}))
goto end_branch_2
} else {

}
}
{
if (uint32(v2_5_1.IntVal) == 902936544) {
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(&Constructor_Data_Map_Internal_Node{1, (*Constructor_Data_Map_Internal_Node)(v1_4.UnsafePtr).V0, (*Constructor_Data_Map_Internal_Node)(v1_4.UnsafePtr).V1, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(k_1)}, v_2, (*Constructor_Data_Map_Internal_Node)(v1_4.UnsafePtr).V4, (*Constructor_Data_Map_Internal_Node)(v1_4.UnsafePtr).V5})})
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_2:
__t3 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(__t2)})
goto end_branch_3
} else {

}
}
{
__t3 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_3:
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(__t3)}
})
return go__go_3_0_112
}

func Call_Data_Map_Internal_insertWith__118979962(dictOrd_0_loop *Constructor_Data_Ord_Ord, app_1_loop gopurs_runtime.Value, k_2_loop gopurs_runtime.Value, v_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 *Constructor_Data_Ord_Ord = dictOrd_0_loop
_ = dictOrd_0
var app_1 gopurs_runtime.Value = app_1_loop
_ = app_1
var k_2 gopurs_runtime.Value = k_2_loop
_ = k_2
var v_3 gopurs_runtime.Value = v_3_loop
_ = v_3
var go__go_4_0_113 gopurs_runtime.Value
_ = go__go_4_0_113
go__go_4_0_113 = gopurs_runtime.Func(func(v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
var __t3 *Constructor_Data_Map_Internal_Node
{
if (v1_5.Type == 9 && v1_5.IntVal == 324739070 && v1_5.UnsafePtr == nil) {
__t3 = &Constructor_Data_Map_Internal_Node{1, 1, 1, k_2, v_3, (*Constructor_Data_Map_Internal_Node)(nil), (*Constructor_Data_Map_Internal_Node)(nil)}
goto end_branch_3
} else {

}
}
{
if (v1_5.Type == 9 && v1_5.IntVal == 324739070 && v1_5.UnsafePtr != nil) {
// TAST (Let): v2_6_1 -> gopurs_runtime.Value
v2_6_1 := gopurs_runtime.Apply2(gopurs_runtime.Box(dictOrd_0.V1), k_2, (*Constructor_Data_Map_Internal_Node)(v1_5.UnsafePtr).V2)
_ = v2_6_1
var __t2 *Constructor_Data_Map_Internal_Node
{
if (uint32(v2_6_1.IntVal) == 1527465420) {
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeBalancedNode(), (*Constructor_Data_Map_Internal_Node)(v1_5.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node)(v1_5.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.Apply(go__go_4_0_113, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v1_5.UnsafePtr).V4)})))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v1_5.UnsafePtr).V5)}))
goto end_branch_2
} else {

}
}
{
if (uint32(v2_6_1.IntVal) == 380165415) {
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeBalancedNode(), (*Constructor_Data_Map_Internal_Node)(v1_5.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node)(v1_5.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v1_5.UnsafePtr).V4)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.Apply(go__go_4_0_113, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v1_5.UnsafePtr).V5)})))}))
goto end_branch_2
} else {

}
}
{
if (uint32(v2_6_1.IntVal) == 902936544) {
__t2 = &Constructor_Data_Map_Internal_Node{1, (*Constructor_Data_Map_Internal_Node)(v1_5.UnsafePtr).V0, (*Constructor_Data_Map_Internal_Node)(v1_5.UnsafePtr).V1, k_2, gopurs_runtime.Apply2(app_1, (*Constructor_Data_Map_Internal_Node)(v1_5.UnsafePtr).V3, v_3), (*Constructor_Data_Map_Internal_Node)(v1_5.UnsafePtr).V4, (*Constructor_Data_Map_Internal_Node)(v1_5.UnsafePtr).V5}
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_2:
__t3 = __t2
goto end_branch_3
} else {

}
}
{
__t3 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_3:
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(__t3)}
})
return go__go_4_0_113
}

func Call_Data_Map_Internal_intersectionWith__3717755541(dictOrd_0_loop *Constructor_Data_Ord_Ord) gopurs_runtime.Value {
var dictOrd_0 *Constructor_Data_Ord_Ord = dictOrd_0_loop
_ = dictOrd_0
// TAST (Let): compare_1_0 -> gopurs_runtime.Value
compare_1_0 := gopurs_runtime.Box(dictOrd_0.V1)
_ = compare_1_0
return gopurs_runtime.Func(func(app_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m2_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeIntersectionWith(), compare_1_0, app_2, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](m1_3))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](m2_4))})))}
})
})
})
}

func Call_Data_Map_Internal_intersectionWith__4144106805(dictOrd_0_loop *Constructor_Data_Ord_Ord) gopurs_runtime.Value {
var dictOrd_0 *Constructor_Data_Ord_Ord = dictOrd_0_loop
_ = dictOrd_0
// TAST (Let): compare_1_0 -> gopurs_runtime.Value
compare_1_0 := gopurs_runtime.Box(dictOrd_0.V1)
_ = compare_1_0
return gopurs_runtime.Func(func(app_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m2_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeIntersectionWith(), compare_1_0, app_2, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](m1_3))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](m2_4))})))}
})
})
})
}

func Call_Data_Map_Internal_isEmpty__1620059593(v_0_loop *Constructor_Data_Map_Internal_Node) bool {
var v_0 *Constructor_Data_Map_Internal_Node = v_0_loop
_ = v_0
var __t0 bool
{
if (v_0 == nil) {
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

func Call_Data_Map_Internal_iterMapU__878452066(iter_0_loop gopurs_runtime.Value, v_1_loop *Constructor_Data_Map_Internal_Node) gopurs_runtime.Value {
var iter_0 gopurs_runtime.Value = iter_0_loop
_ = iter_0
var v_1 *Constructor_Data_Map_Internal_Node = v_1_loop
_ = v_1
var __t6 *Constructor_Data_Map_Internal_IterEmit
{
if (v_1 == nil) {
__t6 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_IterEmit](iter_0)
goto end_branch_6
} else {

}
}
{
if (v_1 != nil) {
var __t5 *Constructor_Data_Map_Internal_IterEmit
{
var __t_tag_2 *Constructor_Data_Map_Internal_Node = (v_1).V4
if (__t_tag_2 == nil) {
var __t4 *Constructor_Data_Map_Internal_IterEmit
{
var __t_tag_3 *Constructor_Data_Map_Internal_Node = (v_1).V5
if (__t_tag_3 == nil) {
__t4 = &Constructor_Data_Map_Internal_IterEmit{1, (v_1).V2, (v_1).V3, iter_0}
goto end_branch_4
} else {

}
}
{
__t4 = &Constructor_Data_Map_Internal_IterEmit{1, (v_1).V2, (v_1).V3, gopurs_runtime.Value{Type: 9, IntVal: 2861335956, UnsafePtr: unsafe.Pointer(&Constructor_Data_Map_Internal_IterNode{1, (v_1).V5, iter_0})}}
}
end_branch_4:
__t5 = __t4
goto end_branch_5
} else {

}
}
{
var __t1 *Constructor_Data_Map_Internal_IterEmit
{
var __t_tag_0 *Constructor_Data_Map_Internal_Node = (v_1).V5
if (__t_tag_0 == nil) {
__t1 = &Constructor_Data_Map_Internal_IterEmit{1, (v_1).V2, (v_1).V3, gopurs_runtime.Value{Type: 9, IntVal: 2861335956, UnsafePtr: unsafe.Pointer(&Constructor_Data_Map_Internal_IterNode{1, (v_1).V4, iter_0})}}
goto end_branch_1
} else {

}
}
{
__t1 = &Constructor_Data_Map_Internal_IterEmit{1, (v_1).V2, (v_1).V3, gopurs_runtime.Value{Type: 9, IntVal: 2861335956, UnsafePtr: unsafe.Pointer(&Constructor_Data_Map_Internal_IterNode{1, (v_1).V4, gopurs_runtime.Value{Type: 9, IntVal: 2861335956, UnsafePtr: unsafe.Pointer(&Constructor_Data_Map_Internal_IterNode{1, (v_1).V5, iter_0})}})}}
}
end_branch_1:
__t5 = __t1
}
end_branch_5:
__t6 = __t5
goto end_branch_6
} else {

}
}
{
__t6 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_IterEmit](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_6:
return gopurs_runtime.Value{Type: 9, IntVal: 1343415489, UnsafePtr: unsafe.Pointer(__t6)}
}

func Call_Data_Map_Internal_lookup__3378638282(dictOrd_0_loop *Constructor_Data_Ord_Ord, k_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 *Constructor_Data_Ord_Ord = dictOrd_0_loop
_ = dictOrd_0
var k_1 gopurs_runtime.Value = k_1_loop
_ = k_1
var go__go_2_0_121 gopurs_runtime.Value
go__go_2_0_121 = gopurs_runtime.Func(func(v_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_3_loop gopurs_runtime.Value = v_3_loop_val
go__go_2_0_121:
for {
if false { continue go__go_2_0_121 }
var v_3 gopurs_runtime.Value = v_3_loop
_ = v_3
var __t3 *Constructor_Data_Maybe_Just
{
if (v_3.Type == 9 && v_3.IntVal == 324739070 && v_3.UnsafePtr == nil) {
__t3 = (*Constructor_Data_Maybe_Just)(nil)
goto end_branch_3
} else {

}
}
{
if (v_3.Type == 9 && v_3.IntVal == 324739070 && v_3.UnsafePtr != nil) {
// TAST (Let): v1_4_1 -> gopurs_runtime.Value
v1_4_1 := gopurs_runtime.Apply2(gopurs_runtime.Box(dictOrd_0.V1), k_1, (*Constructor_Data_Map_Internal_Node)(v_3.UnsafePtr).V2)
_ = v1_4_1
var __t2 *Constructor_Data_Maybe_Just
{
if (uint32(v1_4_1.IntVal) == 1527465420) {
v_3_loop = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v_3.UnsafePtr).V4)}
continue go__go_2_0_121
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Value{})
goto end_branch_2
} else {

}
}
{
if (uint32(v1_4_1.IntVal) == 380165415) {
v_3_loop = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v_3.UnsafePtr).V5)}
continue go__go_2_0_121
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Value{})
goto end_branch_2
} else {

}
}
{
if (uint32(v1_4_1.IntVal) == 902936544) {
__t2 = &Constructor_Data_Maybe_Just{1, (*Constructor_Data_Map_Internal_Node)(v_3.UnsafePtr).V3}
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_2:
__t3 = __t2
goto end_branch_3
} else {

}
}
{
__t3 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_3:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t3)}
}
}()
})
return go__go_2_0_121
}

func Call_Data_Map_Internal_lookup__1040249709(k_0_loop uint32) gopurs_runtime.Value {
var k_0 uint32 = k_0_loop
_ = k_0
var go__go_1_0_122 gopurs_runtime.Value
go__go_1_0_122 = gopurs_runtime.Func(func(v_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_2_loop gopurs_runtime.Value = v_2_loop_val
go__go_1_0_122:
for {
if false { continue go__go_1_0_122 }
var v_2 gopurs_runtime.Value = v_2_loop
_ = v_2
var __t3 *Constructor_Data_Maybe_Just
{
if (v_2.Type == 9 && v_2.IntVal == 324739070 && v_2.UnsafePtr == nil) {
__t3 = (*Constructor_Data_Maybe_Just)(nil)
goto end_branch_3
} else {

}
}
{
if (v_2.Type == 9 && v_2.IntVal == 324739070 && v_2.UnsafePtr != nil) {
// TAST (Let): v1_3_1 -> gopurs_runtime.Value
v1_3_1 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_Interval_Duration_ordDurationComponent(), "compare"), gopurs_runtime.Value{Type: 9, IntVal: int64(k_0), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(uint32((*Constructor_Data_Map_Internal_Node)(v_2.UnsafePtr).V2.IntVal)), UnsafePtr: nil})
_ = v1_3_1
var __t2 *Constructor_Data_Maybe_Just
{
if (uint32(v1_3_1.IntVal) == 1527465420) {
v_2_loop = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v_2.UnsafePtr).V4)}
continue go__go_1_0_122
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Value{})
goto end_branch_2
} else {

}
}
{
if (uint32(v1_3_1.IntVal) == 380165415) {
v_2_loop = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v_2.UnsafePtr).V5)}
continue go__go_1_0_122
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Value{})
goto end_branch_2
} else {

}
}
{
if (uint32(v1_3_1.IntVal) == 902936544) {
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, gopurs_runtime.Float((*Constructor_Data_Map_Internal_Node)(v_2.UnsafePtr).V3.FloatVal())})})
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_2:
__t3 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t2)})
goto end_branch_3
} else {

}
}
{
__t3 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_3:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t3)}
}
}()
})
return go__go_1_0_122
}

func Call_Data_Map_Internal_mapMaybe__3426301240(dictOrd_0_loop *Constructor_Data_Ord_Ord, x_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 *Constructor_Data_Ord_Ord = dictOrd_0_loop
_ = dictOrd_0
var x_1 gopurs_runtime.Value = x_1_loop
_ = x_1
var go__go_2_0_123 gopurs_runtime.Value
_ = go__go_2_0_123
go__go_2_0_123 = gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t3 *Constructor_Data_Map_Internal_Node
{
if (v_3.Type == 9 && v_3.IntVal == 324739070 && v_3.UnsafePtr == nil) {
__t3 = (*Constructor_Data_Map_Internal_Node)(nil)
goto end_branch_3
} else {

}
}
{
if (v_3.Type == 9 && v_3.IntVal == 324739070 && v_3.UnsafePtr != nil) {
// TAST (Let): v2_4_1 -> gopurs_runtime.Value
v2_4_1 := gopurs_runtime.Apply(x_1, (*Constructor_Data_Map_Internal_Node)(v_3.UnsafePtr).V3)
_ = v2_4_1
var __t2 *Constructor_Data_Map_Internal_Node
{
if (v2_4_1.Type == 9 && v2_4_1.IntVal == 930809136 && v2_4_1.UnsafePtr != nil) {
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeBalancedNode(), (*Constructor_Data_Map_Internal_Node)(v_3.UnsafePtr).V2, (*Constructor_Data_Maybe_Just)(v2_4_1.UnsafePtr).V0, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.Apply(go__go_2_0_123, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v_3.UnsafePtr).V4)})))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.Apply(go__go_2_0_123, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v_3.UnsafePtr).V5)})))}))
goto end_branch_2
} else {

}
}
{
if (v2_4_1.Type == 9 && v2_4_1.IntVal == 930809136 && v2_4_1.UnsafePtr == nil) {
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.UncurriedApp2(Get_Data_Map_Internal_unsafeJoinNodes(), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.Apply(go__go_2_0_123, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v_3.UnsafePtr).V4)})))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.Apply(go__go_2_0_123, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v_3.UnsafePtr).V5)})))}))
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_2:
__t3 = __t2
goto end_branch_3
} else {

}
}
{
__t3 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_3:
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(__t3)}
})
return go__go_2_0_123
}

func Call_Data_Map_Internal_mapMaybe__1970555288(dictOrd_0_loop *Constructor_Data_Ord_Ord, x_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 *Constructor_Data_Ord_Ord = dictOrd_0_loop
_ = dictOrd_0
var x_1 gopurs_runtime.Value = x_1_loop
_ = x_1
var go__go_2_0_124 gopurs_runtime.Value
_ = go__go_2_0_124
go__go_2_0_124 = gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t3 *Constructor_Data_Map_Internal_Node
{
if (v_3.Type == 9 && v_3.IntVal == 324739070 && v_3.UnsafePtr == nil) {
__t3 = (*Constructor_Data_Map_Internal_Node)(nil)
goto end_branch_3
} else {

}
}
{
if (v_3.Type == 9 && v_3.IntVal == 324739070 && v_3.UnsafePtr != nil) {
// TAST (Let): v2_4_1 -> gopurs_runtime.Value
v2_4_1 := gopurs_runtime.Apply(x_1, (*Constructor_Data_Map_Internal_Node)(v_3.UnsafePtr).V3)
_ = v2_4_1
var __t2 *Constructor_Data_Map_Internal_Node
{
if (v2_4_1.Type == 9 && v2_4_1.IntVal == 930809136 && v2_4_1.UnsafePtr != nil) {
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeBalancedNode(), (*Constructor_Data_Map_Internal_Node)(v_3.UnsafePtr).V2, (*Constructor_Data_Maybe_Just)(v2_4_1.UnsafePtr).V0, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.Apply(go__go_2_0_124, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v_3.UnsafePtr).V4)})))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.Apply(go__go_2_0_124, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v_3.UnsafePtr).V5)})))}))
goto end_branch_2
} else {

}
}
{
if (v2_4_1.Type == 9 && v2_4_1.IntVal == 930809136 && v2_4_1.UnsafePtr == nil) {
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.UncurriedApp2(Get_Data_Map_Internal_unsafeJoinNodes(), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.Apply(go__go_2_0_124, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v_3.UnsafePtr).V4)})))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.Apply(go__go_2_0_124, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v_3.UnsafePtr).V5)})))}))
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_2:
__t3 = __t2
goto end_branch_3
} else {

}
}
{
__t3 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_3:
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(__t3)}
})
return go__go_2_0_124
}

func Call_Data_Map_Internal_mapMaybeWithKey__817660689(dictOrd_0_loop *Constructor_Data_Ord_Ord, f_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 *Constructor_Data_Ord_Ord = dictOrd_0_loop
_ = dictOrd_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var go__go_2_0_125 gopurs_runtime.Value
_ = go__go_2_0_125
go__go_2_0_125 = gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t3 *Constructor_Data_Map_Internal_Node
{
if (v_3.Type == 9 && v_3.IntVal == 324739070 && v_3.UnsafePtr == nil) {
__t3 = (*Constructor_Data_Map_Internal_Node)(nil)
goto end_branch_3
} else {

}
}
{
if (v_3.Type == 9 && v_3.IntVal == 324739070 && v_3.UnsafePtr != nil) {
// TAST (Let): v2_4_1 -> gopurs_runtime.Value
v2_4_1 := gopurs_runtime.Apply2(f_1, (*Constructor_Data_Map_Internal_Node)(v_3.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node)(v_3.UnsafePtr).V3)
_ = v2_4_1
var __t2 *Constructor_Data_Map_Internal_Node
{
if (v2_4_1.Type == 9 && v2_4_1.IntVal == 930809136 && v2_4_1.UnsafePtr != nil) {
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeBalancedNode(), (*Constructor_Data_Map_Internal_Node)(v_3.UnsafePtr).V2, (*Constructor_Data_Maybe_Just)(v2_4_1.UnsafePtr).V0, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.Apply(go__go_2_0_125, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v_3.UnsafePtr).V4)})))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.Apply(go__go_2_0_125, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v_3.UnsafePtr).V5)})))}))
goto end_branch_2
} else {

}
}
{
if (v2_4_1.Type == 9 && v2_4_1.IntVal == 930809136 && v2_4_1.UnsafePtr == nil) {
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.UncurriedApp2(Get_Data_Map_Internal_unsafeJoinNodes(), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.Apply(go__go_2_0_125, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v_3.UnsafePtr).V4)})))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.Apply(go__go_2_0_125, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v_3.UnsafePtr).V5)})))}))
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_2:
__t3 = __t2
goto end_branch_3
} else {

}
}
{
__t3 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_3:
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(__t3)}
})
return go__go_2_0_125
}

func Call_Data_Map_Internal_singleton__3511563426(k_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) *Constructor_Data_Map_Internal_Node {
var k_0 gopurs_runtime.Value = k_0_loop
_ = k_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return &Constructor_Data_Map_Internal_Node{1, 1, 1, k_0, v_1, (*Constructor_Data_Map_Internal_Node)(nil), (*Constructor_Data_Map_Internal_Node)(nil)}
}

func Call_Data_Map_Internal_singleton__943571066(k_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) *Constructor_Data_Map_Internal_Node {
var k_0 gopurs_runtime.Value = k_0_loop
_ = k_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return &Constructor_Data_Map_Internal_Node{1, 1, 1, k_0, v_1, (*Constructor_Data_Map_Internal_Node)(nil), (*Constructor_Data_Map_Internal_Node)(nil)}
}

func Call_Data_Map_Internal_singleton__2450056090(k_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) *Constructor_Data_Map_Internal_Node {
var k_0 gopurs_runtime.Value = k_0_loop
_ = k_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return &Constructor_Data_Map_Internal_Node{1, 1, 1, k_0, v_1, (*Constructor_Data_Map_Internal_Node)(nil), (*Constructor_Data_Map_Internal_Node)(nil)}
}

func Call_Data_Map_Internal_singleton__3707014010(k_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) *Constructor_Data_Map_Internal_Node {
var k_0 gopurs_runtime.Value = k_0_loop
_ = k_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return &Constructor_Data_Map_Internal_Node{1, 1, 1, k_0, v_1, (*Constructor_Data_Map_Internal_Node)(nil), (*Constructor_Data_Map_Internal_Node)(nil)}
}

func Call_Data_Map_Internal_singleton__1518627866(k_0_loop uint32, v_1_loop float64) *Constructor_Data_Map_Internal_Node {
var k_0 uint32 = k_0_loop
_ = k_0
var v_1 float64 = v_1_loop
_ = v_1
return gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(&Constructor_Data_Map_Internal_Node{1, 1, 1, gopurs_runtime.Value{Type: 9, IntVal: int64(k_0), UnsafePtr: nil}, gopurs_runtime.Float(v_1), (*Constructor_Data_Map_Internal_Node)(nil), (*Constructor_Data_Map_Internal_Node)(nil)})})
}

func Call_Data_Map_Internal_singleton__1300483034(k_0_loop *Constructor_Data_Maybe_Just, v_1_loop gopurs_runtime.Value) *Constructor_Data_Map_Internal_Node {
var k_0 *Constructor_Data_Maybe_Just = k_0_loop
_ = k_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(&Constructor_Data_Map_Internal_Node{1, 1, 1, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(k_0)}, v_1, (*Constructor_Data_Map_Internal_Node)(nil), (*Constructor_Data_Map_Internal_Node)(nil)})})
}

func Call_Data_Map_Internal_size__909390430(v_0_loop *Constructor_Data_Map_Internal_Node) int64 {
var v_0 *Constructor_Data_Map_Internal_Node = v_0_loop
_ = v_0
var __t0 int64
{
if (v_0 == nil) {
__t0 = 0
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
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal
}
end_branch_0:
return __t0
}

func Call_Data_Map_Internal_size__1374028086(v_0_loop *Constructor_Data_Map_Internal_Node) int64 {
var v_0 *Constructor_Data_Map_Internal_Node = v_0_loop
_ = v_0
var __t0 int64
{
if (v_0 == nil) {
__t0 = 0
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
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal
}
end_branch_0:
return __t0
}

func Call_Data_Map_Internal_size__2382154916(v_0_loop *Constructor_Data_Map_Internal_Node) int64 {
var v_0 *Constructor_Data_Map_Internal_Node = v_0_loop
_ = v_0
var __t0 int64
{
if (v_0 == nil) {
__t0 = 0
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
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal
}
end_branch_0:
return __t0
}

func Call_Data_Map_Internal_stepWith__2632420966(f_0_loop gopurs_runtime.Value, next_1_loop gopurs_runtime.Value, done_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var next_1 gopurs_runtime.Value = next_1_loop
_ = next_1
var done_2 gopurs_runtime.Value = done_2_loop
_ = done_2
var go__go_3_0_126 gopurs_runtime.Value
go__go_3_0_126 = gopurs_runtime.Func(func(v_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_4_loop gopurs_runtime.Value = v_4_loop_val
go__go_3_0_126:
for {
if false { continue go__go_3_0_126 }
var v_4 gopurs_runtime.Value = v_4_loop
_ = v_4
var __t1 *Constructor_Data_Maybe_Just
{
if (v_4.Type == 9 && v_4.IntVal == 2509360378) {
__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply(done_2, Get_Data_Unit_unit()))
goto end_branch_1
} else {

}
}
{
if (v_4.Type == 9 && v_4.IntVal == 1343415489) {
__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.UncurriedApp3(next_1, (*Constructor_Data_Map_Internal_IterEmit)(v_4.UnsafePtr).V0, (*Constructor_Data_Map_Internal_IterEmit)(v_4.UnsafePtr).V1, (*Constructor_Data_Map_Internal_IterEmit)(v_4.UnsafePtr).V2))
goto end_branch_1
} else {

}
}
{
if (v_4.Type == 9 && v_4.IntVal == 2861335956) {
v_4_loop = gopurs_runtime.Apply2(f_0, (*Constructor_Data_Map_Internal_IterNode)(v_4.UnsafePtr).V1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_IterNode)(v_4.UnsafePtr).V0)})
continue go__go_3_0_126
__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Value{})
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_1:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t1)}
}
}()
})
return go__go_3_0_126
}

func Call_Data_Map_Internal_stepWith__603436967(f_0_loop gopurs_runtime.Value, next_1_loop gopurs_runtime.Value, done_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var next_1 gopurs_runtime.Value = next_1_loop
_ = next_1
var done_2 gopurs_runtime.Value = done_2_loop
_ = done_2
var go__go_3_0_127 gopurs_runtime.Value
go__go_3_0_127 = gopurs_runtime.Func(func(v_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_4_loop gopurs_runtime.Value = v_4_loop_val
go__go_3_0_127:
for {
if false { continue go__go_3_0_127 }
var v_4 gopurs_runtime.Value = v_4_loop
_ = v_4
var __t1 *Constructor_Data_Tuple_Tuple
{
if (v_4.Type == 9 && v_4.IntVal == 2509360378) {
__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple](gopurs_runtime.Apply(done_2, Get_Data_Unit_unit()))
goto end_branch_1
} else {

}
}
{
if (v_4.Type == 9 && v_4.IntVal == 1343415489) {
__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple](gopurs_runtime.UncurriedApp3(next_1, (*Constructor_Data_Map_Internal_IterEmit)(v_4.UnsafePtr).V0, (*Constructor_Data_Map_Internal_IterEmit)(v_4.UnsafePtr).V1, (*Constructor_Data_Map_Internal_IterEmit)(v_4.UnsafePtr).V2))
goto end_branch_1
} else {

}
}
{
if (v_4.Type == 9 && v_4.IntVal == 2861335956) {
v_4_loop = gopurs_runtime.Apply2(f_0, (*Constructor_Data_Map_Internal_IterNode)(v_4.UnsafePtr).V1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_IterNode)(v_4.UnsafePtr).V0)})
continue go__go_3_0_127
__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple](gopurs_runtime.Value{})
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_1:
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(__t1)}
}
}()
})
return go__go_3_0_127
}

func Call_Data_Map_Internal_stepWith__3186376421(f_0_loop gopurs_runtime.Value, next_1_loop gopurs_runtime.Value, done_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var next_1 gopurs_runtime.Value = next_1_loop
_ = next_1
var done_2 gopurs_runtime.Value = done_2_loop
_ = done_2
var go__go_3_0_128 gopurs_runtime.Value
go__go_3_0_128 = gopurs_runtime.Func(func(v_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_4_loop gopurs_runtime.Value = v_4_loop_val
go__go_3_0_128:
for {
if false { continue go__go_3_0_128 }
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
__t1 = gopurs_runtime.UncurriedApp3(next_1, (*Constructor_Data_Map_Internal_IterEmit)(v_4.UnsafePtr).V0, (*Constructor_Data_Map_Internal_IterEmit)(v_4.UnsafePtr).V1, (*Constructor_Data_Map_Internal_IterEmit)(v_4.UnsafePtr).V2)
goto end_branch_1
} else {

}
}
{
if (v_4.Type == 9 && v_4.IntVal == 2861335956) {
v_4_loop = gopurs_runtime.Apply2(f_0, (*Constructor_Data_Map_Internal_IterNode)(v_4.UnsafePtr).V1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_IterNode)(v_4.UnsafePtr).V0)})
continue go__go_3_0_128
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
return go__go_3_0_128
}

func Call_Data_Map_Internal_stepWith__2866328237(f_0_loop gopurs_runtime.Value, next_1_loop gopurs_runtime.Value, done_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var next_1 gopurs_runtime.Value = next_1_loop
_ = next_1
var done_2 gopurs_runtime.Value = done_2_loop
_ = done_2
var go__go_3_0_129 gopurs_runtime.Value
go__go_3_0_129 = gopurs_runtime.Func(func(v_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_4_loop gopurs_runtime.Value = v_4_loop_val
go__go_3_0_129:
for {
if false { continue go__go_3_0_129 }
var v_4 gopurs_runtime.Value = v_4_loop
_ = v_4
var __t1 *Constructor_Data_Map_Internal_IterNext
{
if (v_4.Type == 9 && v_4.IntVal == 2509360378) {
__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_IterNext](gopurs_runtime.Apply(done_2, Get_Data_Unit_unit()))
goto end_branch_1
} else {

}
}
{
if (v_4.Type == 9 && v_4.IntVal == 1343415489) {
__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_IterNext](gopurs_runtime.UncurriedApp3(next_1, (*Constructor_Data_Map_Internal_IterEmit)(v_4.UnsafePtr).V0, (*Constructor_Data_Map_Internal_IterEmit)(v_4.UnsafePtr).V1, (*Constructor_Data_Map_Internal_IterEmit)(v_4.UnsafePtr).V2))
goto end_branch_1
} else {

}
}
{
if (v_4.Type == 9 && v_4.IntVal == 2861335956) {
v_4_loop = gopurs_runtime.Apply2(f_0, (*Constructor_Data_Map_Internal_IterNode)(v_4.UnsafePtr).V1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_IterNode)(v_4.UnsafePtr).V0)})
continue go__go_3_0_129
__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_IterNext](gopurs_runtime.Value{})
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_IterNext](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_1:
return gopurs_runtime.Value{Type: 9, IntVal: 953589075, UnsafePtr: unsafe.Pointer(__t1)}
}
}()
})
return go__go_3_0_129
}

func Call_Data_Map_Internal_stepWith__280335550(f_0_loop gopurs_runtime.Value, next_1_loop gopurs_runtime.Value, done_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var next_1 gopurs_runtime.Value = next_1_loop
_ = next_1
var done_2 gopurs_runtime.Value = done_2_loop
_ = done_2
var go__go_3_0_130 gopurs_runtime.Value
go__go_3_0_130 = gopurs_runtime.Func(func(v_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_4_loop gopurs_runtime.Value = v_4_loop_val
go__go_3_0_130:
for {
if false { continue go__go_3_0_130 }
var v_4 gopurs_runtime.Value = v_4_loop
_ = v_4
var __t1 *Constructor_Data_Maybe_Just
{
if (v_4.Type == 9 && v_4.IntVal == 2509360378) {
__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply(done_2, Get_Data_Unit_unit()))
goto end_branch_1
} else {

}
}
{
if (v_4.Type == 9 && v_4.IntVal == 1343415489) {
__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.UncurriedApp3(next_1, (*Constructor_Data_Map_Internal_IterEmit)(v_4.UnsafePtr).V0, (*Constructor_Data_Map_Internal_IterEmit)(v_4.UnsafePtr).V1, (*Constructor_Data_Map_Internal_IterEmit)(v_4.UnsafePtr).V2))
goto end_branch_1
} else {

}
}
{
if (v_4.Type == 9 && v_4.IntVal == 2861335956) {
v_4_loop = gopurs_runtime.Apply2(f_0, (*Constructor_Data_Map_Internal_IterNode)(v_4.UnsafePtr).V1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_IterNode)(v_4.UnsafePtr).V0)})
continue go__go_3_0_130
__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Value{})
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_1:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t1)}
}
}()
})
return go__go_3_0_130
}

func Call_Data_Map_Internal_stepWith__2834533669(f_0_loop gopurs_runtime.Value, next_1_loop gopurs_runtime.Value, done_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var next_1 gopurs_runtime.Value = next_1_loop
_ = next_1
var done_2 gopurs_runtime.Value = done_2_loop
_ = done_2
var go__go_3_0_131 gopurs_runtime.Value
go__go_3_0_131 = gopurs_runtime.Func(func(v_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_4_loop gopurs_runtime.Value = v_4_loop_val
go__go_3_0_131:
for {
if false { continue go__go_3_0_131 }
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
__t1 = gopurs_runtime.UncurriedApp3(next_1, gopurs_runtime.Value{Type: 9, IntVal: int64(uint32((*Constructor_Data_Map_Internal_IterEmit)(v_4.UnsafePtr).V0.IntVal)), UnsafePtr: nil}, gopurs_runtime.Float((*Constructor_Data_Map_Internal_IterEmit)(v_4.UnsafePtr).V1.FloatVal()), (*Constructor_Data_Map_Internal_IterEmit)(v_4.UnsafePtr).V2)
goto end_branch_1
} else {

}
}
{
if (v_4.Type == 9 && v_4.IntVal == 2861335956) {
v_4_loop = gopurs_runtime.Apply2(f_0, (*Constructor_Data_Map_Internal_IterNode)(v_4.UnsafePtr).V1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_IterNode)(v_4.UnsafePtr).V0)})
continue go__go_3_0_131
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
return go__go_3_0_131
}

func Call_Data_Map_Internal_stepWith__1463181374(f_0_loop gopurs_runtime.Value, next_1_loop gopurs_runtime.Value, done_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var next_1 gopurs_runtime.Value = next_1_loop
_ = next_1
var done_2 gopurs_runtime.Value = done_2_loop
_ = done_2
var go__go_3_0_132 gopurs_runtime.Value
go__go_3_0_132 = gopurs_runtime.Func(func(v_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_4_loop gopurs_runtime.Value = v_4_loop_val
go__go_3_0_132:
for {
if false { continue go__go_3_0_132 }
var v_4 gopurs_runtime.Value = v_4_loop
_ = v_4
var __t1 *Constructor_Data_Maybe_Just
{
if (v_4.Type == 9 && v_4.IntVal == 2509360378) {
__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply(done_2, Get_Data_Unit_unit()))
goto end_branch_1
} else {

}
}
{
if (v_4.Type == 9 && v_4.IntVal == 1343415489) {
__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.UncurriedApp3(next_1, gopurs_runtime.Value{Type: 9, IntVal: int64(uint32((*Constructor_Data_Map_Internal_IterEmit)(v_4.UnsafePtr).V0.IntVal)), UnsafePtr: nil}, gopurs_runtime.Float((*Constructor_Data_Map_Internal_IterEmit)(v_4.UnsafePtr).V1.FloatVal()), (*Constructor_Data_Map_Internal_IterEmit)(v_4.UnsafePtr).V2))
goto end_branch_1
} else {

}
}
{
if (v_4.Type == 9 && v_4.IntVal == 2861335956) {
v_4_loop = gopurs_runtime.Apply2(f_0, (*Constructor_Data_Map_Internal_IterNode)(v_4.UnsafePtr).V1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_IterNode)(v_4.UnsafePtr).V0)})
continue go__go_3_0_132
__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Value{})
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_1:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t1)}
}
}()
})
return go__go_3_0_132
}

func Call_Data_Map_Internal_toMapIter__1799172593(a_0_loop *Constructor_Data_Map_Internal_Node) gopurs_runtime.Value {
var a_0 *Constructor_Data_Map_Internal_Node = a_0_loop
_ = a_0
return gopurs_runtime.Value{Type: 9, IntVal: 2861335956, UnsafePtr: unsafe.Pointer(&Constructor_Data_Map_Internal_IterNode{1, a_0, gopurs_runtime.Value{Type: 9, IntVal: 2509360378, UnsafePtr: unsafe.Pointer(nil)}})}
}

func Call_Data_Map_Internal_toMapIter__2014410513(a_0_loop *Constructor_Data_Map_Internal_Node) gopurs_runtime.Value {
var a_0 *Constructor_Data_Map_Internal_Node = a_0_loop
_ = a_0
return gopurs_runtime.Value{Type: 9, IntVal: 2861335956, UnsafePtr: unsafe.Pointer(&Constructor_Data_Map_Internal_IterNode{1, a_0, gopurs_runtime.Value{Type: 9, IntVal: 2509360378, UnsafePtr: unsafe.Pointer(nil)}})}
}

func Call_Data_Map_Internal_toMapIter__772765521(a_0_loop *Constructor_Data_Map_Internal_Node) gopurs_runtime.Value {
var a_0 *Constructor_Data_Map_Internal_Node = a_0_loop
_ = a_0
return gopurs_runtime.Value{Type: 9, IntVal: 2861335956, UnsafePtr: unsafe.Pointer(&Constructor_Data_Map_Internal_IterNode{1, gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(a_0)}), gopurs_runtime.Value{Type: 9, IntVal: 2509360378, UnsafePtr: unsafe.Pointer(nil)}})}
}

func Call_Data_Map_Internal_toMapIter__1738891721(a_0_loop *Constructor_Data_Map_Internal_Node) gopurs_runtime.Value {
var a_0 *Constructor_Data_Map_Internal_Node = a_0_loop
_ = a_0
return gopurs_runtime.Value{Type: 9, IntVal: 2861335956, UnsafePtr: unsafe.Pointer(&Constructor_Data_Map_Internal_IterNode{1, a_0, gopurs_runtime.Value{Type: 9, IntVal: 2509360378, UnsafePtr: unsafe.Pointer(nil)}})}
}

func Call_Data_Map_Internal_toUnfoldable__2183602684(dictUnfoldable_0_loop *Constructor_Data_Unfoldable_Unfoldable) gopurs_runtime.Value {
var dictUnfoldable_0 *Constructor_Data_Unfoldable_Unfoldable = dictUnfoldable_0_loop
_ = dictUnfoldable_0
// TAST (Let): __local_var_1_0 -> gopurs_runtime.Value
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.Box(dictUnfoldable_0.V1), Get_Data_Map_Internal_stepUnfoldr())
_ = __local_var_1_0
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_1_0, gopurs_runtime.Value{Type: 9, IntVal: 2861335956, UnsafePtr: unsafe.Pointer(&Constructor_Data_Map_Internal_IterNode{1, gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](x_2), gopurs_runtime.Value{Type: 9, IntVal: 2509360378, UnsafePtr: unsafe.Pointer(nil)}})})
})
}

func Call_Data_Map_Internal_toUnfoldable__2567957978(__eta0_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_List_Types_unfoldableList(), "unfoldr"), Get_Data_Map_Internal_stepUnfoldr(), gopurs_runtime.Value{Type: 9, IntVal: 2861335956, UnsafePtr: unsafe.Pointer(&Constructor_Data_Map_Internal_IterNode{1, gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](__eta0_0), gopurs_runtime.Value{Type: 9, IntVal: 2509360378, UnsafePtr: unsafe.Pointer(nil)}})})))}
}

func Call_Data_Map_Internal_unionWith__2507192643(dictOrd_0_loop *Constructor_Data_Ord_Ord) gopurs_runtime.Value {
var dictOrd_0 *Constructor_Data_Ord_Ord = dictOrd_0_loop
_ = dictOrd_0
// TAST (Let): compare_1_0 -> gopurs_runtime.Value
compare_1_0 := gopurs_runtime.Box(dictOrd_0.V1)
_ = compare_1_0
return gopurs_runtime.Func(func(app_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m2_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeUnionWith(), compare_1_0, app_2, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](m1_3))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](m2_4))})))}
})
})
})
}

func Call_Data_Map_Internal_unionWith__952079555(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value, __eta2_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
_ = __eta1_1
var __eta2_2 gopurs_runtime.Value = __eta2_2_loop
_ = __eta2_2
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeUnionWith(), gopurs_runtime.RecordGet(Get_Data_Interval_Duration_ordDurationComponent(), "compare"), __eta0_0, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](__eta1_1))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](__eta2_2))})))}
}

func Call_Data_Map_Internal_unsafeBalancedNode__1259503046(__local_var_0_loop gopurs_runtime.Value, __local_var_1_loop gopurs_runtime.Value, __local_var_2_loop gopurs_runtime.Value, __local_var_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __local_var_0 gopurs_runtime.Value = __local_var_0_loop
_ = __local_var_0
var __local_var_1 gopurs_runtime.Value = __local_var_1_loop
_ = __local_var_1
var __local_var_2 gopurs_runtime.Value = __local_var_2_loop
_ = __local_var_2
var __local_var_3 gopurs_runtime.Value = __local_var_3_loop
_ = __local_var_3
var __t37 *Constructor_Data_Map_Internal_Node
{
if (__local_var_2.Type == 9 && __local_var_2.IntVal == 324739070 && __local_var_2.UnsafePtr == nil) {
var __t9 *Constructor_Data_Map_Internal_Node
{
if (__local_var_3.Type == 9 && __local_var_3.IntVal == 324739070 && __local_var_3.UnsafePtr == nil) {
__t9 = &Constructor_Data_Map_Internal_Node{1, 1, 1, __local_var_0, __local_var_1, (*Constructor_Data_Map_Internal_Node)(nil), (*Constructor_Data_Map_Internal_Node)(nil)}
goto end_branch_9
} else {

}
}
{
var __t_and_1 bool = false
if (__local_var_3.Type == 9 && __local_var_3.IntVal == 324739070 && __local_var_3.UnsafePtr != nil) {

var __t0 bool
{
if ((*Constructor_Data_Map_Internal_Node)(__local_var_3.UnsafePtr).V0) > (1) {
__t0 = true
goto end_branch_0
} else {

}
}
{
__t0 = false
}
end_branch_0:
__t_and_1 = __t0
}
if __t_and_1 {
var __t8 *Constructor_Data_Map_Internal_Node
{
var __t_tag_2 *Constructor_Data_Map_Internal_Node = (*Constructor_Data_Map_Internal_Node)(__local_var_3.UnsafePtr).V4
var __t_and_7 bool = false
if (__t_tag_2 != nil) {

var __t6 bool
{
var __t5 int64
{
var __t_tag_3 *Constructor_Data_Map_Internal_Node = (*Constructor_Data_Map_Internal_Node)(__local_var_3.UnsafePtr).V5
if (__t_tag_3 == nil) {
__t5 = 0
goto end_branch_5
} else {

}
}
{
var __t_tag_4 *Constructor_Data_Map_Internal_Node = (*Constructor_Data_Map_Internal_Node)(__local_var_3.UnsafePtr).V5
if (__t_tag_4 != nil) {
__t5 = ((*Constructor_Data_Map_Internal_Node)(__local_var_3.UnsafePtr).V5).V0
goto end_branch_5
} else {

}
}
{
__t5 = func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal
}
end_branch_5:
if (((*Constructor_Data_Map_Internal_Node)(__local_var_3.UnsafePtr).V4).V0) > (__t5) {
__t6 = true
goto end_branch_6
} else {

}
}
{
__t6 = false
}
end_branch_6:
__t_and_7 = __t6
}
if __t_and_7 {
__t8 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeNode(), ((*Constructor_Data_Map_Internal_Node)(__local_var_3.UnsafePtr).V4).V2, ((*Constructor_Data_Map_Internal_Node)(__local_var_3.UnsafePtr).V4).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeNode(), __local_var_0, __local_var_1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](__local_var_2))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(((*Constructor_Data_Map_Internal_Node)(__local_var_3.UnsafePtr).V4).V4)})))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeNode(), (*Constructor_Data_Map_Internal_Node)(__local_var_3.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node)(__local_var_3.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(((*Constructor_Data_Map_Internal_Node)(__local_var_3.UnsafePtr).V4).V5)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_3.UnsafePtr).V5)})))}))
goto end_branch_8
} else {

}
}
{
__t8 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeNode(), (*Constructor_Data_Map_Internal_Node)(__local_var_3.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node)(__local_var_3.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeNode(), __local_var_0, __local_var_1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](__local_var_2))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_3.UnsafePtr).V4)})))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_3.UnsafePtr).V5)}))
}
end_branch_8:
__t9 = __t8
goto end_branch_9
} else {

}
}
{
__t9 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeNode(), __local_var_0, __local_var_1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](__local_var_2))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](__local_var_3))}))
}
end_branch_9:
__t37 = __t9
goto end_branch_37
} else {

}
}
{
if (__local_var_2.Type == 9 && __local_var_2.IntVal == 324739070 && __local_var_2.UnsafePtr != nil) {
var __t36 *Constructor_Data_Map_Internal_Node
{
if (__local_var_3.Type == 9 && __local_var_3.IntVal == 324739070 && __local_var_3.UnsafePtr != nil) {
var __t26 *Constructor_Data_Map_Internal_Node
{
var __t10 bool
{
if ((*Constructor_Data_Map_Internal_Node)(__local_var_3.UnsafePtr).V0) > (((*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V0) + (1)) {
__t10 = true
goto end_branch_10
} else {

}
}
{
__t10 = false
}
end_branch_10:
if __t10 {
var __t17 *Constructor_Data_Map_Internal_Node
{
var __t_tag_11 *Constructor_Data_Map_Internal_Node = (*Constructor_Data_Map_Internal_Node)(__local_var_3.UnsafePtr).V4
var __t_and_16 bool = false
if (__t_tag_11 != nil) {

var __t15 bool
{
var __t14 int64
{
var __t_tag_12 *Constructor_Data_Map_Internal_Node = (*Constructor_Data_Map_Internal_Node)(__local_var_3.UnsafePtr).V5
if (__t_tag_12 == nil) {
__t14 = 0
goto end_branch_14
} else {

}
}
{
var __t_tag_13 *Constructor_Data_Map_Internal_Node = (*Constructor_Data_Map_Internal_Node)(__local_var_3.UnsafePtr).V5
if (__t_tag_13 != nil) {
__t14 = ((*Constructor_Data_Map_Internal_Node)(__local_var_3.UnsafePtr).V5).V0
goto end_branch_14
} else {

}
}
{
__t14 = func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal
}
end_branch_14:
if (((*Constructor_Data_Map_Internal_Node)(__local_var_3.UnsafePtr).V4).V0) > (__t14) {
__t15 = true
goto end_branch_15
} else {

}
}
{
__t15 = false
}
end_branch_15:
__t_and_16 = __t15
}
if __t_and_16 {
__t17 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeNode(), ((*Constructor_Data_Map_Internal_Node)(__local_var_3.UnsafePtr).V4).V2, ((*Constructor_Data_Map_Internal_Node)(__local_var_3.UnsafePtr).V4).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeNode(), __local_var_0, __local_var_1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](__local_var_2))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(((*Constructor_Data_Map_Internal_Node)(__local_var_3.UnsafePtr).V4).V4)})))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeNode(), (*Constructor_Data_Map_Internal_Node)(__local_var_3.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node)(__local_var_3.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(((*Constructor_Data_Map_Internal_Node)(__local_var_3.UnsafePtr).V4).V5)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_3.UnsafePtr).V5)})))}))
goto end_branch_17
} else {

}
}
{
__t17 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeNode(), (*Constructor_Data_Map_Internal_Node)(__local_var_3.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node)(__local_var_3.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeNode(), __local_var_0, __local_var_1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](__local_var_2))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_3.UnsafePtr).V4)})))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_3.UnsafePtr).V5)}))
}
end_branch_17:
__t26 = __t17
goto end_branch_26
} else {

}
}
{
var __t18 bool
{
if ((*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V0) > (((*Constructor_Data_Map_Internal_Node)(__local_var_3.UnsafePtr).V0) + (1)) {
__t18 = true
goto end_branch_18
} else {

}
}
{
__t18 = false
}
end_branch_18:
if __t18 {
var __t25 *Constructor_Data_Map_Internal_Node
{
var __t_tag_19 *Constructor_Data_Map_Internal_Node = (*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V5
var __t_and_24 bool = false
if (__t_tag_19 != nil) {

var __t23 bool
{
var __t22 int64
{
var __t_tag_20 *Constructor_Data_Map_Internal_Node = (*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V4
if (__t_tag_20 == nil) {
__t22 = 0
goto end_branch_22
} else {

}
}
{
var __t_tag_21 *Constructor_Data_Map_Internal_Node = (*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V4
if (__t_tag_21 != nil) {
__t22 = ((*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V4).V0
goto end_branch_22
} else {

}
}
{
__t22 = func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal
}
end_branch_22:
if (__t22) > (((*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V5).V0) {
__t23 = false
goto end_branch_23
} else {

}
}
{
__t23 = true
}
end_branch_23:
__t_and_24 = __t23
}
if __t_and_24 {
__t25 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeNode(), ((*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V5).V2, ((*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V5).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeNode(), (*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V4)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(((*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V5).V4)})))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeNode(), __local_var_0, __local_var_1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(((*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V5).V5)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](__local_var_3))})))}))
goto end_branch_25
} else {

}
}
{
__t25 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeNode(), (*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V4)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeNode(), __local_var_0, __local_var_1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V5)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](__local_var_3))})))}))
}
end_branch_25:
__t26 = __t25
goto end_branch_26
} else {

}
}
{
__t26 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeNode(), __local_var_0, __local_var_1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](__local_var_2))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](__local_var_3))}))
}
end_branch_26:
__t36 = __t26
goto end_branch_36
} else {

}
}
{
var __t_and_28 bool = false
if (__local_var_3.Type == 9 && __local_var_3.IntVal == 324739070 && __local_var_3.UnsafePtr == nil) {

var __t27 bool
{
if ((*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V0) > (1) {
__t27 = true
goto end_branch_27
} else {

}
}
{
__t27 = false
}
end_branch_27:
__t_and_28 = __t27
}
if __t_and_28 {
var __t35 *Constructor_Data_Map_Internal_Node
{
var __t_tag_29 *Constructor_Data_Map_Internal_Node = (*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V5
var __t_and_34 bool = false
if (__t_tag_29 != nil) {

var __t33 bool
{
var __t32 int64
{
var __t_tag_30 *Constructor_Data_Map_Internal_Node = (*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V4
if (__t_tag_30 == nil) {
__t32 = 0
goto end_branch_32
} else {

}
}
{
var __t_tag_31 *Constructor_Data_Map_Internal_Node = (*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V4
if (__t_tag_31 != nil) {
__t32 = ((*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V4).V0
goto end_branch_32
} else {

}
}
{
__t32 = func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal
}
end_branch_32:
if (__t32) > (((*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V5).V0) {
__t33 = false
goto end_branch_33
} else {

}
}
{
__t33 = true
}
end_branch_33:
__t_and_34 = __t33
}
if __t_and_34 {
__t35 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeNode(), ((*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V5).V2, ((*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V5).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeNode(), (*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V4)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(((*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V5).V4)})))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeNode(), __local_var_0, __local_var_1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(((*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V5).V5)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](__local_var_3))})))}))
goto end_branch_35
} else {

}
}
{
__t35 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeNode(), (*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V4)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeNode(), __local_var_0, __local_var_1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V5)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](__local_var_3))})))}))
}
end_branch_35:
__t36 = __t35
goto end_branch_36
} else {

}
}
{
__t36 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeNode(), __local_var_0, __local_var_1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](__local_var_2))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](__local_var_3))}))
}
end_branch_36:
__t37 = __t36
goto end_branch_37
} else {

}
}
{
__t37 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_37:
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(__t37)}
}

func Call_Data_Map_Internal_unsafeBalancedNode__1305301638(__local_var_0_loop gopurs_runtime.Value, __local_var_1_loop gopurs_runtime.Value, __local_var_2_loop gopurs_runtime.Value, __local_var_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __local_var_0 gopurs_runtime.Value = __local_var_0_loop
_ = __local_var_0
var __local_var_1 gopurs_runtime.Value = __local_var_1_loop
_ = __local_var_1
var __local_var_2 gopurs_runtime.Value = __local_var_2_loop
_ = __local_var_2
var __local_var_3 gopurs_runtime.Value = __local_var_3_loop
_ = __local_var_3
var __t37 *Constructor_Data_Map_Internal_Node
{
if (__local_var_2.Type == 9 && __local_var_2.IntVal == 324739070 && __local_var_2.UnsafePtr == nil) {
var __t9 *Constructor_Data_Map_Internal_Node
{
if (__local_var_3.Type == 9 && __local_var_3.IntVal == 324739070 && __local_var_3.UnsafePtr == nil) {
__t9 = &Constructor_Data_Map_Internal_Node{1, 1, 1, __local_var_0, __local_var_1, (*Constructor_Data_Map_Internal_Node)(nil), (*Constructor_Data_Map_Internal_Node)(nil)}
goto end_branch_9
} else {

}
}
{
var __t_and_1 bool = false
if (__local_var_3.Type == 9 && __local_var_3.IntVal == 324739070 && __local_var_3.UnsafePtr != nil) {

var __t0 bool
{
if ((*Constructor_Data_Map_Internal_Node)(__local_var_3.UnsafePtr).V0) > (1) {
__t0 = true
goto end_branch_0
} else {

}
}
{
__t0 = false
}
end_branch_0:
__t_and_1 = __t0
}
if __t_and_1 {
var __t8 *Constructor_Data_Map_Internal_Node
{
var __t_tag_2 *Constructor_Data_Map_Internal_Node = (*Constructor_Data_Map_Internal_Node)(__local_var_3.UnsafePtr).V4
var __t_and_7 bool = false
if (__t_tag_2 != nil) {

var __t6 bool
{
var __t5 int64
{
var __t_tag_3 *Constructor_Data_Map_Internal_Node = (*Constructor_Data_Map_Internal_Node)(__local_var_3.UnsafePtr).V5
if (__t_tag_3 == nil) {
__t5 = 0
goto end_branch_5
} else {

}
}
{
var __t_tag_4 *Constructor_Data_Map_Internal_Node = (*Constructor_Data_Map_Internal_Node)(__local_var_3.UnsafePtr).V5
if (__t_tag_4 != nil) {
__t5 = ((*Constructor_Data_Map_Internal_Node)(__local_var_3.UnsafePtr).V5).V0
goto end_branch_5
} else {

}
}
{
__t5 = func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal
}
end_branch_5:
if (((*Constructor_Data_Map_Internal_Node)(__local_var_3.UnsafePtr).V4).V0) > (__t5) {
__t6 = true
goto end_branch_6
} else {

}
}
{
__t6 = false
}
end_branch_6:
__t_and_7 = __t6
}
if __t_and_7 {
__t8 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeNode(), ((*Constructor_Data_Map_Internal_Node)(__local_var_3.UnsafePtr).V4).V2, ((*Constructor_Data_Map_Internal_Node)(__local_var_3.UnsafePtr).V4).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeNode(), __local_var_0, __local_var_1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](__local_var_2))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(((*Constructor_Data_Map_Internal_Node)(__local_var_3.UnsafePtr).V4).V4)})))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeNode(), (*Constructor_Data_Map_Internal_Node)(__local_var_3.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node)(__local_var_3.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(((*Constructor_Data_Map_Internal_Node)(__local_var_3.UnsafePtr).V4).V5)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_3.UnsafePtr).V5)})))}))
goto end_branch_8
} else {

}
}
{
__t8 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeNode(), (*Constructor_Data_Map_Internal_Node)(__local_var_3.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node)(__local_var_3.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeNode(), __local_var_0, __local_var_1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](__local_var_2))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_3.UnsafePtr).V4)})))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_3.UnsafePtr).V5)}))
}
end_branch_8:
__t9 = __t8
goto end_branch_9
} else {

}
}
{
__t9 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeNode(), __local_var_0, __local_var_1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](__local_var_2))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](__local_var_3))}))
}
end_branch_9:
__t37 = __t9
goto end_branch_37
} else {

}
}
{
if (__local_var_2.Type == 9 && __local_var_2.IntVal == 324739070 && __local_var_2.UnsafePtr != nil) {
var __t36 *Constructor_Data_Map_Internal_Node
{
if (__local_var_3.Type == 9 && __local_var_3.IntVal == 324739070 && __local_var_3.UnsafePtr != nil) {
var __t26 *Constructor_Data_Map_Internal_Node
{
var __t10 bool
{
if ((*Constructor_Data_Map_Internal_Node)(__local_var_3.UnsafePtr).V0) > (((*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V0) + (1)) {
__t10 = true
goto end_branch_10
} else {

}
}
{
__t10 = false
}
end_branch_10:
if __t10 {
var __t17 *Constructor_Data_Map_Internal_Node
{
var __t_tag_11 *Constructor_Data_Map_Internal_Node = (*Constructor_Data_Map_Internal_Node)(__local_var_3.UnsafePtr).V4
var __t_and_16 bool = false
if (__t_tag_11 != nil) {

var __t15 bool
{
var __t14 int64
{
var __t_tag_12 *Constructor_Data_Map_Internal_Node = (*Constructor_Data_Map_Internal_Node)(__local_var_3.UnsafePtr).V5
if (__t_tag_12 == nil) {
__t14 = 0
goto end_branch_14
} else {

}
}
{
var __t_tag_13 *Constructor_Data_Map_Internal_Node = (*Constructor_Data_Map_Internal_Node)(__local_var_3.UnsafePtr).V5
if (__t_tag_13 != nil) {
__t14 = ((*Constructor_Data_Map_Internal_Node)(__local_var_3.UnsafePtr).V5).V0
goto end_branch_14
} else {

}
}
{
__t14 = func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal
}
end_branch_14:
if (((*Constructor_Data_Map_Internal_Node)(__local_var_3.UnsafePtr).V4).V0) > (__t14) {
__t15 = true
goto end_branch_15
} else {

}
}
{
__t15 = false
}
end_branch_15:
__t_and_16 = __t15
}
if __t_and_16 {
__t17 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeNode(), ((*Constructor_Data_Map_Internal_Node)(__local_var_3.UnsafePtr).V4).V2, ((*Constructor_Data_Map_Internal_Node)(__local_var_3.UnsafePtr).V4).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeNode(), __local_var_0, __local_var_1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](__local_var_2))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(((*Constructor_Data_Map_Internal_Node)(__local_var_3.UnsafePtr).V4).V4)})))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeNode(), (*Constructor_Data_Map_Internal_Node)(__local_var_3.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node)(__local_var_3.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(((*Constructor_Data_Map_Internal_Node)(__local_var_3.UnsafePtr).V4).V5)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_3.UnsafePtr).V5)})))}))
goto end_branch_17
} else {

}
}
{
__t17 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeNode(), (*Constructor_Data_Map_Internal_Node)(__local_var_3.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node)(__local_var_3.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeNode(), __local_var_0, __local_var_1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](__local_var_2))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_3.UnsafePtr).V4)})))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_3.UnsafePtr).V5)}))
}
end_branch_17:
__t26 = __t17
goto end_branch_26
} else {

}
}
{
var __t18 bool
{
if ((*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V0) > (((*Constructor_Data_Map_Internal_Node)(__local_var_3.UnsafePtr).V0) + (1)) {
__t18 = true
goto end_branch_18
} else {

}
}
{
__t18 = false
}
end_branch_18:
if __t18 {
var __t25 *Constructor_Data_Map_Internal_Node
{
var __t_tag_19 *Constructor_Data_Map_Internal_Node = (*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V5
var __t_and_24 bool = false
if (__t_tag_19 != nil) {

var __t23 bool
{
var __t22 int64
{
var __t_tag_20 *Constructor_Data_Map_Internal_Node = (*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V4
if (__t_tag_20 == nil) {
__t22 = 0
goto end_branch_22
} else {

}
}
{
var __t_tag_21 *Constructor_Data_Map_Internal_Node = (*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V4
if (__t_tag_21 != nil) {
__t22 = ((*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V4).V0
goto end_branch_22
} else {

}
}
{
__t22 = func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal
}
end_branch_22:
if (__t22) > (((*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V5).V0) {
__t23 = false
goto end_branch_23
} else {

}
}
{
__t23 = true
}
end_branch_23:
__t_and_24 = __t23
}
if __t_and_24 {
__t25 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeNode(), ((*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V5).V2, ((*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V5).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeNode(), (*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V4)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(((*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V5).V4)})))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeNode(), __local_var_0, __local_var_1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(((*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V5).V5)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](__local_var_3))})))}))
goto end_branch_25
} else {

}
}
{
__t25 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeNode(), (*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V4)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeNode(), __local_var_0, __local_var_1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V5)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](__local_var_3))})))}))
}
end_branch_25:
__t26 = __t25
goto end_branch_26
} else {

}
}
{
__t26 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeNode(), __local_var_0, __local_var_1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](__local_var_2))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](__local_var_3))}))
}
end_branch_26:
__t36 = __t26
goto end_branch_36
} else {

}
}
{
var __t_and_28 bool = false
if (__local_var_3.Type == 9 && __local_var_3.IntVal == 324739070 && __local_var_3.UnsafePtr == nil) {

var __t27 bool
{
if ((*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V0) > (1) {
__t27 = true
goto end_branch_27
} else {

}
}
{
__t27 = false
}
end_branch_27:
__t_and_28 = __t27
}
if __t_and_28 {
var __t35 *Constructor_Data_Map_Internal_Node
{
var __t_tag_29 *Constructor_Data_Map_Internal_Node = (*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V5
var __t_and_34 bool = false
if (__t_tag_29 != nil) {

var __t33 bool
{
var __t32 int64
{
var __t_tag_30 *Constructor_Data_Map_Internal_Node = (*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V4
if (__t_tag_30 == nil) {
__t32 = 0
goto end_branch_32
} else {

}
}
{
var __t_tag_31 *Constructor_Data_Map_Internal_Node = (*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V4
if (__t_tag_31 != nil) {
__t32 = ((*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V4).V0
goto end_branch_32
} else {

}
}
{
__t32 = func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal
}
end_branch_32:
if (__t32) > (((*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V5).V0) {
__t33 = false
goto end_branch_33
} else {

}
}
{
__t33 = true
}
end_branch_33:
__t_and_34 = __t33
}
if __t_and_34 {
__t35 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeNode(), ((*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V5).V2, ((*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V5).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeNode(), (*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V4)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(((*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V5).V4)})))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeNode(), __local_var_0, __local_var_1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(((*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V5).V5)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](__local_var_3))})))}))
goto end_branch_35
} else {

}
}
{
__t35 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeNode(), (*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V4)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeNode(), __local_var_0, __local_var_1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V5)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](__local_var_3))})))}))
}
end_branch_35:
__t36 = __t35
goto end_branch_36
} else {

}
}
{
__t36 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeNode(), __local_var_0, __local_var_1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](__local_var_2))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](__local_var_3))}))
}
end_branch_36:
__t37 = __t36
goto end_branch_37
} else {

}
}
{
__t37 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_37:
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(__t37)}
}

func Call_Data_Map_Internal_unsafeBalancedNode__954819782(__local_var_0_loop gopurs_runtime.Value, __local_var_1_loop gopurs_runtime.Value, __local_var_2_loop gopurs_runtime.Value, __local_var_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __local_var_0 gopurs_runtime.Value = __local_var_0_loop
_ = __local_var_0
var __local_var_1 gopurs_runtime.Value = __local_var_1_loop
_ = __local_var_1
var __local_var_2 gopurs_runtime.Value = __local_var_2_loop
_ = __local_var_2
var __local_var_3 gopurs_runtime.Value = __local_var_3_loop
_ = __local_var_3
var __t37 *Constructor_Data_Map_Internal_Node
{
if (__local_var_2.Type == 9 && __local_var_2.IntVal == 324739070 && __local_var_2.UnsafePtr == nil) {
var __t9 *Constructor_Data_Map_Internal_Node
{
if (__local_var_3.Type == 9 && __local_var_3.IntVal == 324739070 && __local_var_3.UnsafePtr == nil) {
__t9 = &Constructor_Data_Map_Internal_Node{1, 1, 1, __local_var_0, __local_var_1, (*Constructor_Data_Map_Internal_Node)(nil), (*Constructor_Data_Map_Internal_Node)(nil)}
goto end_branch_9
} else {

}
}
{
var __t_and_1 bool = false
if (__local_var_3.Type == 9 && __local_var_3.IntVal == 324739070 && __local_var_3.UnsafePtr != nil) {

var __t0 bool
{
if ((*Constructor_Data_Map_Internal_Node)(__local_var_3.UnsafePtr).V0) > (1) {
__t0 = true
goto end_branch_0
} else {

}
}
{
__t0 = false
}
end_branch_0:
__t_and_1 = __t0
}
if __t_and_1 {
var __t8 *Constructor_Data_Map_Internal_Node
{
var __t_tag_2 *Constructor_Data_Map_Internal_Node = (*Constructor_Data_Map_Internal_Node)(__local_var_3.UnsafePtr).V4
var __t_and_7 bool = false
if (__t_tag_2 != nil) {

var __t6 bool
{
var __t5 int64
{
var __t_tag_3 *Constructor_Data_Map_Internal_Node = (*Constructor_Data_Map_Internal_Node)(__local_var_3.UnsafePtr).V5
if (__t_tag_3 == nil) {
__t5 = 0
goto end_branch_5
} else {

}
}
{
var __t_tag_4 *Constructor_Data_Map_Internal_Node = (*Constructor_Data_Map_Internal_Node)(__local_var_3.UnsafePtr).V5
if (__t_tag_4 != nil) {
__t5 = ((*Constructor_Data_Map_Internal_Node)(__local_var_3.UnsafePtr).V5).V0
goto end_branch_5
} else {

}
}
{
__t5 = func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal
}
end_branch_5:
if (((*Constructor_Data_Map_Internal_Node)(__local_var_3.UnsafePtr).V4).V0) > (__t5) {
__t6 = true
goto end_branch_6
} else {

}
}
{
__t6 = false
}
end_branch_6:
__t_and_7 = __t6
}
if __t_and_7 {
__t8 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeNode(), ((*Constructor_Data_Map_Internal_Node)(__local_var_3.UnsafePtr).V4).V2, ((*Constructor_Data_Map_Internal_Node)(__local_var_3.UnsafePtr).V4).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeNode(), __local_var_0, __local_var_1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](__local_var_2))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(((*Constructor_Data_Map_Internal_Node)(__local_var_3.UnsafePtr).V4).V4)})))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeNode(), (*Constructor_Data_Map_Internal_Node)(__local_var_3.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node)(__local_var_3.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(((*Constructor_Data_Map_Internal_Node)(__local_var_3.UnsafePtr).V4).V5)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_3.UnsafePtr).V5)})))}))
goto end_branch_8
} else {

}
}
{
__t8 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeNode(), (*Constructor_Data_Map_Internal_Node)(__local_var_3.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node)(__local_var_3.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeNode(), __local_var_0, __local_var_1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](__local_var_2))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_3.UnsafePtr).V4)})))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_3.UnsafePtr).V5)}))
}
end_branch_8:
__t9 = __t8
goto end_branch_9
} else {

}
}
{
__t9 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeNode(), __local_var_0, __local_var_1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](__local_var_2))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](__local_var_3))}))
}
end_branch_9:
__t37 = __t9
goto end_branch_37
} else {

}
}
{
if (__local_var_2.Type == 9 && __local_var_2.IntVal == 324739070 && __local_var_2.UnsafePtr != nil) {
var __t36 *Constructor_Data_Map_Internal_Node
{
if (__local_var_3.Type == 9 && __local_var_3.IntVal == 324739070 && __local_var_3.UnsafePtr != nil) {
var __t26 *Constructor_Data_Map_Internal_Node
{
var __t10 bool
{
if ((*Constructor_Data_Map_Internal_Node)(__local_var_3.UnsafePtr).V0) > (((*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V0) + (1)) {
__t10 = true
goto end_branch_10
} else {

}
}
{
__t10 = false
}
end_branch_10:
if __t10 {
var __t17 *Constructor_Data_Map_Internal_Node
{
var __t_tag_11 *Constructor_Data_Map_Internal_Node = (*Constructor_Data_Map_Internal_Node)(__local_var_3.UnsafePtr).V4
var __t_and_16 bool = false
if (__t_tag_11 != nil) {

var __t15 bool
{
var __t14 int64
{
var __t_tag_12 *Constructor_Data_Map_Internal_Node = (*Constructor_Data_Map_Internal_Node)(__local_var_3.UnsafePtr).V5
if (__t_tag_12 == nil) {
__t14 = 0
goto end_branch_14
} else {

}
}
{
var __t_tag_13 *Constructor_Data_Map_Internal_Node = (*Constructor_Data_Map_Internal_Node)(__local_var_3.UnsafePtr).V5
if (__t_tag_13 != nil) {
__t14 = ((*Constructor_Data_Map_Internal_Node)(__local_var_3.UnsafePtr).V5).V0
goto end_branch_14
} else {

}
}
{
__t14 = func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal
}
end_branch_14:
if (((*Constructor_Data_Map_Internal_Node)(__local_var_3.UnsafePtr).V4).V0) > (__t14) {
__t15 = true
goto end_branch_15
} else {

}
}
{
__t15 = false
}
end_branch_15:
__t_and_16 = __t15
}
if __t_and_16 {
__t17 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeNode(), ((*Constructor_Data_Map_Internal_Node)(__local_var_3.UnsafePtr).V4).V2, ((*Constructor_Data_Map_Internal_Node)(__local_var_3.UnsafePtr).V4).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeNode(), __local_var_0, __local_var_1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](__local_var_2))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(((*Constructor_Data_Map_Internal_Node)(__local_var_3.UnsafePtr).V4).V4)})))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeNode(), (*Constructor_Data_Map_Internal_Node)(__local_var_3.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node)(__local_var_3.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(((*Constructor_Data_Map_Internal_Node)(__local_var_3.UnsafePtr).V4).V5)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_3.UnsafePtr).V5)})))}))
goto end_branch_17
} else {

}
}
{
__t17 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeNode(), (*Constructor_Data_Map_Internal_Node)(__local_var_3.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node)(__local_var_3.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeNode(), __local_var_0, __local_var_1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](__local_var_2))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_3.UnsafePtr).V4)})))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_3.UnsafePtr).V5)}))
}
end_branch_17:
__t26 = __t17
goto end_branch_26
} else {

}
}
{
var __t18 bool
{
if ((*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V0) > (((*Constructor_Data_Map_Internal_Node)(__local_var_3.UnsafePtr).V0) + (1)) {
__t18 = true
goto end_branch_18
} else {

}
}
{
__t18 = false
}
end_branch_18:
if __t18 {
var __t25 *Constructor_Data_Map_Internal_Node
{
var __t_tag_19 *Constructor_Data_Map_Internal_Node = (*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V5
var __t_and_24 bool = false
if (__t_tag_19 != nil) {

var __t23 bool
{
var __t22 int64
{
var __t_tag_20 *Constructor_Data_Map_Internal_Node = (*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V4
if (__t_tag_20 == nil) {
__t22 = 0
goto end_branch_22
} else {

}
}
{
var __t_tag_21 *Constructor_Data_Map_Internal_Node = (*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V4
if (__t_tag_21 != nil) {
__t22 = ((*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V4).V0
goto end_branch_22
} else {

}
}
{
__t22 = func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal
}
end_branch_22:
if (__t22) > (((*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V5).V0) {
__t23 = false
goto end_branch_23
} else {

}
}
{
__t23 = true
}
end_branch_23:
__t_and_24 = __t23
}
if __t_and_24 {
__t25 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeNode(), ((*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V5).V2, ((*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V5).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeNode(), (*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V4)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(((*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V5).V4)})))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeNode(), __local_var_0, __local_var_1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(((*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V5).V5)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](__local_var_3))})))}))
goto end_branch_25
} else {

}
}
{
__t25 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeNode(), (*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V4)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeNode(), __local_var_0, __local_var_1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V5)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](__local_var_3))})))}))
}
end_branch_25:
__t26 = __t25
goto end_branch_26
} else {

}
}
{
__t26 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeNode(), __local_var_0, __local_var_1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](__local_var_2))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](__local_var_3))}))
}
end_branch_26:
__t36 = __t26
goto end_branch_36
} else {

}
}
{
var __t_and_28 bool = false
if (__local_var_3.Type == 9 && __local_var_3.IntVal == 324739070 && __local_var_3.UnsafePtr == nil) {

var __t27 bool
{
if ((*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V0) > (1) {
__t27 = true
goto end_branch_27
} else {

}
}
{
__t27 = false
}
end_branch_27:
__t_and_28 = __t27
}
if __t_and_28 {
var __t35 *Constructor_Data_Map_Internal_Node
{
var __t_tag_29 *Constructor_Data_Map_Internal_Node = (*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V5
var __t_and_34 bool = false
if (__t_tag_29 != nil) {

var __t33 bool
{
var __t32 int64
{
var __t_tag_30 *Constructor_Data_Map_Internal_Node = (*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V4
if (__t_tag_30 == nil) {
__t32 = 0
goto end_branch_32
} else {

}
}
{
var __t_tag_31 *Constructor_Data_Map_Internal_Node = (*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V4
if (__t_tag_31 != nil) {
__t32 = ((*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V4).V0
goto end_branch_32
} else {

}
}
{
__t32 = func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal
}
end_branch_32:
if (__t32) > (((*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V5).V0) {
__t33 = false
goto end_branch_33
} else {

}
}
{
__t33 = true
}
end_branch_33:
__t_and_34 = __t33
}
if __t_and_34 {
__t35 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeNode(), ((*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V5).V2, ((*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V5).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeNode(), (*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V4)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(((*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V5).V4)})))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeNode(), __local_var_0, __local_var_1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(((*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V5).V5)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](__local_var_3))})))}))
goto end_branch_35
} else {

}
}
{
__t35 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeNode(), (*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V4)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeNode(), __local_var_0, __local_var_1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V5)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](__local_var_3))})))}))
}
end_branch_35:
__t36 = __t35
goto end_branch_36
} else {

}
}
{
__t36 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeNode(), __local_var_0, __local_var_1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](__local_var_2))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](__local_var_3))}))
}
end_branch_36:
__t37 = __t36
goto end_branch_37
} else {

}
}
{
__t37 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_37:
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(__t37)}
}

func Call_Data_Map_Internal_unsafeBalancedNode__1776657286(__local_var_0_loop gopurs_runtime.Value, __local_var_1_loop gopurs_runtime.Value, __local_var_2_loop gopurs_runtime.Value, __local_var_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __local_var_0 gopurs_runtime.Value = __local_var_0_loop
_ = __local_var_0
var __local_var_1 gopurs_runtime.Value = __local_var_1_loop
_ = __local_var_1
var __local_var_2 gopurs_runtime.Value = __local_var_2_loop
_ = __local_var_2
var __local_var_3 gopurs_runtime.Value = __local_var_3_loop
_ = __local_var_3
var __t37 *Constructor_Data_Map_Internal_Node
{
if (__local_var_2.Type == 9 && __local_var_2.IntVal == 324739070 && __local_var_2.UnsafePtr == nil) {
var __t9 *Constructor_Data_Map_Internal_Node
{
if (__local_var_3.Type == 9 && __local_var_3.IntVal == 324739070 && __local_var_3.UnsafePtr == nil) {
__t9 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(&Constructor_Data_Map_Internal_Node{1, 1, 1, gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(__local_var_0.IntVal)), UnsafePtr: nil}, gopurs_runtime.Float(__local_var_1.FloatVal()), (*Constructor_Data_Map_Internal_Node)(nil), (*Constructor_Data_Map_Internal_Node)(nil)})})
goto end_branch_9
} else {

}
}
{
var __t_and_1 bool = false
if (__local_var_3.Type == 9 && __local_var_3.IntVal == 324739070 && __local_var_3.UnsafePtr != nil) {

var __t0 bool
{
if ((*Constructor_Data_Map_Internal_Node)(__local_var_3.UnsafePtr).V0) > (1) {
__t0 = true
goto end_branch_0
} else {

}
}
{
__t0 = false
}
end_branch_0:
__t_and_1 = __t0
}
if __t_and_1 {
var __t8 *Constructor_Data_Map_Internal_Node
{
var __t_tag_2 *Constructor_Data_Map_Internal_Node = (*Constructor_Data_Map_Internal_Node)(__local_var_3.UnsafePtr).V4
var __t_and_7 bool = false
if (__t_tag_2 != nil) {

var __t6 bool
{
var __t5 int64
{
var __t_tag_3 *Constructor_Data_Map_Internal_Node = (*Constructor_Data_Map_Internal_Node)(__local_var_3.UnsafePtr).V5
if (__t_tag_3 == nil) {
__t5 = 0
goto end_branch_5
} else {

}
}
{
var __t_tag_4 *Constructor_Data_Map_Internal_Node = (*Constructor_Data_Map_Internal_Node)(__local_var_3.UnsafePtr).V5
if (__t_tag_4 != nil) {
__t5 = ((*Constructor_Data_Map_Internal_Node)(__local_var_3.UnsafePtr).V5).V0
goto end_branch_5
} else {

}
}
{
__t5 = func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal
}
end_branch_5:
if (((*Constructor_Data_Map_Internal_Node)(__local_var_3.UnsafePtr).V4).V0) > (__t5) {
__t6 = true
goto end_branch_6
} else {

}
}
{
__t6 = false
}
end_branch_6:
__t_and_7 = __t6
}
if __t_and_7 {
__t8 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeNode(), gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(((*Constructor_Data_Map_Internal_Node)(__local_var_3.UnsafePtr).V4).V2.IntVal)), UnsafePtr: nil}, gopurs_runtime.Float(((*Constructor_Data_Map_Internal_Node)(__local_var_3.UnsafePtr).V4).V3.FloatVal()), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeNode(), gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(__local_var_0.IntVal)), UnsafePtr: nil}, gopurs_runtime.Float(__local_var_1.FloatVal()), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](__local_var_2))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(((*Constructor_Data_Map_Internal_Node)(__local_var_3.UnsafePtr).V4).V4)})))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeNode(), gopurs_runtime.Value{Type: 9, IntVal: int64(uint32((*Constructor_Data_Map_Internal_Node)(__local_var_3.UnsafePtr).V2.IntVal)), UnsafePtr: nil}, gopurs_runtime.Float((*Constructor_Data_Map_Internal_Node)(__local_var_3.UnsafePtr).V3.FloatVal()), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(((*Constructor_Data_Map_Internal_Node)(__local_var_3.UnsafePtr).V4).V5)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_3.UnsafePtr).V5)})))}))
goto end_branch_8
} else {

}
}
{
__t8 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeNode(), gopurs_runtime.Value{Type: 9, IntVal: int64(uint32((*Constructor_Data_Map_Internal_Node)(__local_var_3.UnsafePtr).V2.IntVal)), UnsafePtr: nil}, gopurs_runtime.Float((*Constructor_Data_Map_Internal_Node)(__local_var_3.UnsafePtr).V3.FloatVal()), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeNode(), gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(__local_var_0.IntVal)), UnsafePtr: nil}, gopurs_runtime.Float(__local_var_1.FloatVal()), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](__local_var_2))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_3.UnsafePtr).V4)})))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_3.UnsafePtr).V5)}))
}
end_branch_8:
__t9 = __t8
goto end_branch_9
} else {

}
}
{
__t9 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeNode(), gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(__local_var_0.IntVal)), UnsafePtr: nil}, gopurs_runtime.Float(__local_var_1.FloatVal()), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](__local_var_2))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](__local_var_3))}))
}
end_branch_9:
__t37 = __t9
goto end_branch_37
} else {

}
}
{
if (__local_var_2.Type == 9 && __local_var_2.IntVal == 324739070 && __local_var_2.UnsafePtr != nil) {
var __t36 *Constructor_Data_Map_Internal_Node
{
if (__local_var_3.Type == 9 && __local_var_3.IntVal == 324739070 && __local_var_3.UnsafePtr != nil) {
var __t26 *Constructor_Data_Map_Internal_Node
{
var __t10 bool
{
if ((*Constructor_Data_Map_Internal_Node)(__local_var_3.UnsafePtr).V0) > (((*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V0) + (1)) {
__t10 = true
goto end_branch_10
} else {

}
}
{
__t10 = false
}
end_branch_10:
if __t10 {
var __t17 *Constructor_Data_Map_Internal_Node
{
var __t_tag_11 *Constructor_Data_Map_Internal_Node = (*Constructor_Data_Map_Internal_Node)(__local_var_3.UnsafePtr).V4
var __t_and_16 bool = false
if (__t_tag_11 != nil) {

var __t15 bool
{
var __t14 int64
{
var __t_tag_12 *Constructor_Data_Map_Internal_Node = (*Constructor_Data_Map_Internal_Node)(__local_var_3.UnsafePtr).V5
if (__t_tag_12 == nil) {
__t14 = 0
goto end_branch_14
} else {

}
}
{
var __t_tag_13 *Constructor_Data_Map_Internal_Node = (*Constructor_Data_Map_Internal_Node)(__local_var_3.UnsafePtr).V5
if (__t_tag_13 != nil) {
__t14 = ((*Constructor_Data_Map_Internal_Node)(__local_var_3.UnsafePtr).V5).V0
goto end_branch_14
} else {

}
}
{
__t14 = func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal
}
end_branch_14:
if (((*Constructor_Data_Map_Internal_Node)(__local_var_3.UnsafePtr).V4).V0) > (__t14) {
__t15 = true
goto end_branch_15
} else {

}
}
{
__t15 = false
}
end_branch_15:
__t_and_16 = __t15
}
if __t_and_16 {
__t17 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeNode(), gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(((*Constructor_Data_Map_Internal_Node)(__local_var_3.UnsafePtr).V4).V2.IntVal)), UnsafePtr: nil}, gopurs_runtime.Float(((*Constructor_Data_Map_Internal_Node)(__local_var_3.UnsafePtr).V4).V3.FloatVal()), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeNode(), gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(__local_var_0.IntVal)), UnsafePtr: nil}, gopurs_runtime.Float(__local_var_1.FloatVal()), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](__local_var_2))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(((*Constructor_Data_Map_Internal_Node)(__local_var_3.UnsafePtr).V4).V4)})))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeNode(), gopurs_runtime.Value{Type: 9, IntVal: int64(uint32((*Constructor_Data_Map_Internal_Node)(__local_var_3.UnsafePtr).V2.IntVal)), UnsafePtr: nil}, gopurs_runtime.Float((*Constructor_Data_Map_Internal_Node)(__local_var_3.UnsafePtr).V3.FloatVal()), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(((*Constructor_Data_Map_Internal_Node)(__local_var_3.UnsafePtr).V4).V5)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_3.UnsafePtr).V5)})))}))
goto end_branch_17
} else {

}
}
{
__t17 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeNode(), gopurs_runtime.Value{Type: 9, IntVal: int64(uint32((*Constructor_Data_Map_Internal_Node)(__local_var_3.UnsafePtr).V2.IntVal)), UnsafePtr: nil}, gopurs_runtime.Float((*Constructor_Data_Map_Internal_Node)(__local_var_3.UnsafePtr).V3.FloatVal()), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeNode(), gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(__local_var_0.IntVal)), UnsafePtr: nil}, gopurs_runtime.Float(__local_var_1.FloatVal()), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](__local_var_2))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_3.UnsafePtr).V4)})))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_3.UnsafePtr).V5)}))
}
end_branch_17:
__t26 = __t17
goto end_branch_26
} else {

}
}
{
var __t18 bool
{
if ((*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V0) > (((*Constructor_Data_Map_Internal_Node)(__local_var_3.UnsafePtr).V0) + (1)) {
__t18 = true
goto end_branch_18
} else {

}
}
{
__t18 = false
}
end_branch_18:
if __t18 {
var __t25 *Constructor_Data_Map_Internal_Node
{
var __t_tag_19 *Constructor_Data_Map_Internal_Node = (*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V5
var __t_and_24 bool = false
if (__t_tag_19 != nil) {

var __t23 bool
{
var __t22 int64
{
var __t_tag_20 *Constructor_Data_Map_Internal_Node = (*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V4
if (__t_tag_20 == nil) {
__t22 = 0
goto end_branch_22
} else {

}
}
{
var __t_tag_21 *Constructor_Data_Map_Internal_Node = (*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V4
if (__t_tag_21 != nil) {
__t22 = ((*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V4).V0
goto end_branch_22
} else {

}
}
{
__t22 = func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal
}
end_branch_22:
if (__t22) > (((*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V5).V0) {
__t23 = false
goto end_branch_23
} else {

}
}
{
__t23 = true
}
end_branch_23:
__t_and_24 = __t23
}
if __t_and_24 {
__t25 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeNode(), gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(((*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V5).V2.IntVal)), UnsafePtr: nil}, gopurs_runtime.Float(((*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V5).V3.FloatVal()), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeNode(), gopurs_runtime.Value{Type: 9, IntVal: int64(uint32((*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V2.IntVal)), UnsafePtr: nil}, gopurs_runtime.Float((*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V3.FloatVal()), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V4)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(((*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V5).V4)})))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeNode(), gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(__local_var_0.IntVal)), UnsafePtr: nil}, gopurs_runtime.Float(__local_var_1.FloatVal()), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(((*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V5).V5)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](__local_var_3))})))}))
goto end_branch_25
} else {

}
}
{
__t25 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeNode(), gopurs_runtime.Value{Type: 9, IntVal: int64(uint32((*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V2.IntVal)), UnsafePtr: nil}, gopurs_runtime.Float((*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V3.FloatVal()), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V4)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeNode(), gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(__local_var_0.IntVal)), UnsafePtr: nil}, gopurs_runtime.Float(__local_var_1.FloatVal()), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V5)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](__local_var_3))})))}))
}
end_branch_25:
__t26 = __t25
goto end_branch_26
} else {

}
}
{
__t26 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeNode(), gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(__local_var_0.IntVal)), UnsafePtr: nil}, gopurs_runtime.Float(__local_var_1.FloatVal()), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](__local_var_2))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](__local_var_3))}))
}
end_branch_26:
__t36 = __t26
goto end_branch_36
} else {

}
}
{
var __t_and_28 bool = false
if (__local_var_3.Type == 9 && __local_var_3.IntVal == 324739070 && __local_var_3.UnsafePtr == nil) {

var __t27 bool
{
if ((*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V0) > (1) {
__t27 = true
goto end_branch_27
} else {

}
}
{
__t27 = false
}
end_branch_27:
__t_and_28 = __t27
}
if __t_and_28 {
var __t35 *Constructor_Data_Map_Internal_Node
{
var __t_tag_29 *Constructor_Data_Map_Internal_Node = (*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V5
var __t_and_34 bool = false
if (__t_tag_29 != nil) {

var __t33 bool
{
var __t32 int64
{
var __t_tag_30 *Constructor_Data_Map_Internal_Node = (*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V4
if (__t_tag_30 == nil) {
__t32 = 0
goto end_branch_32
} else {

}
}
{
var __t_tag_31 *Constructor_Data_Map_Internal_Node = (*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V4
if (__t_tag_31 != nil) {
__t32 = ((*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V4).V0
goto end_branch_32
} else {

}
}
{
__t32 = func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal
}
end_branch_32:
if (__t32) > (((*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V5).V0) {
__t33 = false
goto end_branch_33
} else {

}
}
{
__t33 = true
}
end_branch_33:
__t_and_34 = __t33
}
if __t_and_34 {
__t35 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeNode(), gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(((*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V5).V2.IntVal)), UnsafePtr: nil}, gopurs_runtime.Float(((*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V5).V3.FloatVal()), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeNode(), gopurs_runtime.Value{Type: 9, IntVal: int64(uint32((*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V2.IntVal)), UnsafePtr: nil}, gopurs_runtime.Float((*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V3.FloatVal()), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V4)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(((*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V5).V4)})))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeNode(), gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(__local_var_0.IntVal)), UnsafePtr: nil}, gopurs_runtime.Float(__local_var_1.FloatVal()), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(((*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V5).V5)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](__local_var_3))})))}))
goto end_branch_35
} else {

}
}
{
__t35 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeNode(), gopurs_runtime.Value{Type: 9, IntVal: int64(uint32((*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V2.IntVal)), UnsafePtr: nil}, gopurs_runtime.Float((*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V3.FloatVal()), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V4)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeNode(), gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(__local_var_0.IntVal)), UnsafePtr: nil}, gopurs_runtime.Float(__local_var_1.FloatVal()), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V5)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](__local_var_3))})))}))
}
end_branch_35:
__t36 = __t35
goto end_branch_36
} else {

}
}
{
__t36 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeNode(), gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(__local_var_0.IntVal)), UnsafePtr: nil}, gopurs_runtime.Float(__local_var_1.FloatVal()), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](__local_var_2))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](__local_var_3))}))
}
end_branch_36:
__t37 = __t36
goto end_branch_37
} else {

}
}
{
__t37 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_37:
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(__t37)}
}

func Call_Data_Map_Internal_unsafeBalancedNode__1902536198(__local_var_0_loop gopurs_runtime.Value, __local_var_1_loop gopurs_runtime.Value, __local_var_2_loop gopurs_runtime.Value, __local_var_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __local_var_0 gopurs_runtime.Value = __local_var_0_loop
_ = __local_var_0
var __local_var_1 gopurs_runtime.Value = __local_var_1_loop
_ = __local_var_1
var __local_var_2 gopurs_runtime.Value = __local_var_2_loop
_ = __local_var_2
var __local_var_3 gopurs_runtime.Value = __local_var_3_loop
_ = __local_var_3
var __t37 *Constructor_Data_Map_Internal_Node
{
if (__local_var_2.Type == 9 && __local_var_2.IntVal == 324739070 && __local_var_2.UnsafePtr == nil) {
var __t9 *Constructor_Data_Map_Internal_Node
{
if (__local_var_3.Type == 9 && __local_var_3.IntVal == 324739070 && __local_var_3.UnsafePtr == nil) {
__t9 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(&Constructor_Data_Map_Internal_Node{1, 1, 1, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](__local_var_0))}, __local_var_1, (*Constructor_Data_Map_Internal_Node)(nil), (*Constructor_Data_Map_Internal_Node)(nil)})})
goto end_branch_9
} else {

}
}
{
var __t_and_1 bool = false
if (__local_var_3.Type == 9 && __local_var_3.IntVal == 324739070 && __local_var_3.UnsafePtr != nil) {

var __t0 bool
{
if ((*Constructor_Data_Map_Internal_Node)(__local_var_3.UnsafePtr).V0) > (1) {
__t0 = true
goto end_branch_0
} else {

}
}
{
__t0 = false
}
end_branch_0:
__t_and_1 = __t0
}
if __t_and_1 {
var __t8 *Constructor_Data_Map_Internal_Node
{
var __t_tag_2 *Constructor_Data_Map_Internal_Node = (*Constructor_Data_Map_Internal_Node)(__local_var_3.UnsafePtr).V4
var __t_and_7 bool = false
if (__t_tag_2 != nil) {

var __t6 bool
{
var __t5 int64
{
var __t_tag_3 *Constructor_Data_Map_Internal_Node = (*Constructor_Data_Map_Internal_Node)(__local_var_3.UnsafePtr).V5
if (__t_tag_3 == nil) {
__t5 = 0
goto end_branch_5
} else {

}
}
{
var __t_tag_4 *Constructor_Data_Map_Internal_Node = (*Constructor_Data_Map_Internal_Node)(__local_var_3.UnsafePtr).V5
if (__t_tag_4 != nil) {
__t5 = ((*Constructor_Data_Map_Internal_Node)(__local_var_3.UnsafePtr).V5).V0
goto end_branch_5
} else {

}
}
{
__t5 = func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal
}
end_branch_5:
if (((*Constructor_Data_Map_Internal_Node)(__local_var_3.UnsafePtr).V4).V0) > (__t5) {
__t6 = true
goto end_branch_6
} else {

}
}
{
__t6 = false
}
end_branch_6:
__t_and_7 = __t6
}
if __t_and_7 {
__t8 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeNode(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](((*Constructor_Data_Map_Internal_Node)(__local_var_3.UnsafePtr).V4).V2))}, ((*Constructor_Data_Map_Internal_Node)(__local_var_3.UnsafePtr).V4).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeNode(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](__local_var_0))}, __local_var_1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](__local_var_2))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(((*Constructor_Data_Map_Internal_Node)(__local_var_3.UnsafePtr).V4).V4)})))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeNode(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just]((*Constructor_Data_Map_Internal_Node)(__local_var_3.UnsafePtr).V2))}, (*Constructor_Data_Map_Internal_Node)(__local_var_3.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(((*Constructor_Data_Map_Internal_Node)(__local_var_3.UnsafePtr).V4).V5)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_3.UnsafePtr).V5)})))}))
goto end_branch_8
} else {

}
}
{
__t8 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeNode(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just]((*Constructor_Data_Map_Internal_Node)(__local_var_3.UnsafePtr).V2))}, (*Constructor_Data_Map_Internal_Node)(__local_var_3.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeNode(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](__local_var_0))}, __local_var_1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](__local_var_2))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_3.UnsafePtr).V4)})))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_3.UnsafePtr).V5)}))
}
end_branch_8:
__t9 = __t8
goto end_branch_9
} else {

}
}
{
__t9 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeNode(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](__local_var_0))}, __local_var_1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](__local_var_2))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](__local_var_3))}))
}
end_branch_9:
__t37 = __t9
goto end_branch_37
} else {

}
}
{
if (__local_var_2.Type == 9 && __local_var_2.IntVal == 324739070 && __local_var_2.UnsafePtr != nil) {
var __t36 *Constructor_Data_Map_Internal_Node
{
if (__local_var_3.Type == 9 && __local_var_3.IntVal == 324739070 && __local_var_3.UnsafePtr != nil) {
var __t26 *Constructor_Data_Map_Internal_Node
{
var __t10 bool
{
if ((*Constructor_Data_Map_Internal_Node)(__local_var_3.UnsafePtr).V0) > (((*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V0) + (1)) {
__t10 = true
goto end_branch_10
} else {

}
}
{
__t10 = false
}
end_branch_10:
if __t10 {
var __t17 *Constructor_Data_Map_Internal_Node
{
var __t_tag_11 *Constructor_Data_Map_Internal_Node = (*Constructor_Data_Map_Internal_Node)(__local_var_3.UnsafePtr).V4
var __t_and_16 bool = false
if (__t_tag_11 != nil) {

var __t15 bool
{
var __t14 int64
{
var __t_tag_12 *Constructor_Data_Map_Internal_Node = (*Constructor_Data_Map_Internal_Node)(__local_var_3.UnsafePtr).V5
if (__t_tag_12 == nil) {
__t14 = 0
goto end_branch_14
} else {

}
}
{
var __t_tag_13 *Constructor_Data_Map_Internal_Node = (*Constructor_Data_Map_Internal_Node)(__local_var_3.UnsafePtr).V5
if (__t_tag_13 != nil) {
__t14 = ((*Constructor_Data_Map_Internal_Node)(__local_var_3.UnsafePtr).V5).V0
goto end_branch_14
} else {

}
}
{
__t14 = func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal
}
end_branch_14:
if (((*Constructor_Data_Map_Internal_Node)(__local_var_3.UnsafePtr).V4).V0) > (__t14) {
__t15 = true
goto end_branch_15
} else {

}
}
{
__t15 = false
}
end_branch_15:
__t_and_16 = __t15
}
if __t_and_16 {
__t17 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeNode(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](((*Constructor_Data_Map_Internal_Node)(__local_var_3.UnsafePtr).V4).V2))}, ((*Constructor_Data_Map_Internal_Node)(__local_var_3.UnsafePtr).V4).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeNode(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](__local_var_0))}, __local_var_1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](__local_var_2))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(((*Constructor_Data_Map_Internal_Node)(__local_var_3.UnsafePtr).V4).V4)})))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeNode(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just]((*Constructor_Data_Map_Internal_Node)(__local_var_3.UnsafePtr).V2))}, (*Constructor_Data_Map_Internal_Node)(__local_var_3.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(((*Constructor_Data_Map_Internal_Node)(__local_var_3.UnsafePtr).V4).V5)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_3.UnsafePtr).V5)})))}))
goto end_branch_17
} else {

}
}
{
__t17 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeNode(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just]((*Constructor_Data_Map_Internal_Node)(__local_var_3.UnsafePtr).V2))}, (*Constructor_Data_Map_Internal_Node)(__local_var_3.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeNode(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](__local_var_0))}, __local_var_1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](__local_var_2))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_3.UnsafePtr).V4)})))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_3.UnsafePtr).V5)}))
}
end_branch_17:
__t26 = __t17
goto end_branch_26
} else {

}
}
{
var __t18 bool
{
if ((*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V0) > (((*Constructor_Data_Map_Internal_Node)(__local_var_3.UnsafePtr).V0) + (1)) {
__t18 = true
goto end_branch_18
} else {

}
}
{
__t18 = false
}
end_branch_18:
if __t18 {
var __t25 *Constructor_Data_Map_Internal_Node
{
var __t_tag_19 *Constructor_Data_Map_Internal_Node = (*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V5
var __t_and_24 bool = false
if (__t_tag_19 != nil) {

var __t23 bool
{
var __t22 int64
{
var __t_tag_20 *Constructor_Data_Map_Internal_Node = (*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V4
if (__t_tag_20 == nil) {
__t22 = 0
goto end_branch_22
} else {

}
}
{
var __t_tag_21 *Constructor_Data_Map_Internal_Node = (*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V4
if (__t_tag_21 != nil) {
__t22 = ((*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V4).V0
goto end_branch_22
} else {

}
}
{
__t22 = func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal
}
end_branch_22:
if (__t22) > (((*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V5).V0) {
__t23 = false
goto end_branch_23
} else {

}
}
{
__t23 = true
}
end_branch_23:
__t_and_24 = __t23
}
if __t_and_24 {
__t25 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeNode(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](((*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V5).V2))}, ((*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V5).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeNode(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just]((*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V2))}, (*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V4)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(((*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V5).V4)})))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeNode(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](__local_var_0))}, __local_var_1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(((*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V5).V5)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](__local_var_3))})))}))
goto end_branch_25
} else {

}
}
{
__t25 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeNode(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just]((*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V2))}, (*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V4)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeNode(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](__local_var_0))}, __local_var_1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V5)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](__local_var_3))})))}))
}
end_branch_25:
__t26 = __t25
goto end_branch_26
} else {

}
}
{
__t26 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeNode(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](__local_var_0))}, __local_var_1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](__local_var_2))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](__local_var_3))}))
}
end_branch_26:
__t36 = __t26
goto end_branch_36
} else {

}
}
{
var __t_and_28 bool = false
if (__local_var_3.Type == 9 && __local_var_3.IntVal == 324739070 && __local_var_3.UnsafePtr == nil) {

var __t27 bool
{
if ((*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V0) > (1) {
__t27 = true
goto end_branch_27
} else {

}
}
{
__t27 = false
}
end_branch_27:
__t_and_28 = __t27
}
if __t_and_28 {
var __t35 *Constructor_Data_Map_Internal_Node
{
var __t_tag_29 *Constructor_Data_Map_Internal_Node = (*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V5
var __t_and_34 bool = false
if (__t_tag_29 != nil) {

var __t33 bool
{
var __t32 int64
{
var __t_tag_30 *Constructor_Data_Map_Internal_Node = (*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V4
if (__t_tag_30 == nil) {
__t32 = 0
goto end_branch_32
} else {

}
}
{
var __t_tag_31 *Constructor_Data_Map_Internal_Node = (*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V4
if (__t_tag_31 != nil) {
__t32 = ((*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V4).V0
goto end_branch_32
} else {

}
}
{
__t32 = func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal
}
end_branch_32:
if (__t32) > (((*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V5).V0) {
__t33 = false
goto end_branch_33
} else {

}
}
{
__t33 = true
}
end_branch_33:
__t_and_34 = __t33
}
if __t_and_34 {
__t35 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeNode(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](((*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V5).V2))}, ((*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V5).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeNode(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just]((*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V2))}, (*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V4)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(((*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V5).V4)})))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeNode(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](__local_var_0))}, __local_var_1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(((*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V5).V5)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](__local_var_3))})))}))
goto end_branch_35
} else {

}
}
{
__t35 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeNode(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just]((*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V2))}, (*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V4)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeNode(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](__local_var_0))}, __local_var_1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V5)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](__local_var_3))})))}))
}
end_branch_35:
__t36 = __t35
goto end_branch_36
} else {

}
}
{
__t36 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeNode(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](__local_var_0))}, __local_var_1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](__local_var_2))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](__local_var_3))}))
}
end_branch_36:
__t37 = __t36
goto end_branch_37
} else {

}
}
{
__t37 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_37:
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(__t37)}
}

func Call_Data_Map_Internal_unsafeDifference__4097927905(__local_var_0_loop gopurs_runtime.Value, __local_var_1_loop gopurs_runtime.Value, __local_var_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __local_var_0 gopurs_runtime.Value = __local_var_0_loop
_ = __local_var_0
var __local_var_1 gopurs_runtime.Value = __local_var_1_loop
_ = __local_var_1
var __local_var_2 gopurs_runtime.Value = __local_var_2_loop
_ = __local_var_2
var __t1 *Constructor_Data_Map_Internal_Node
{
if (__local_var_1.Type == 9 && __local_var_1.IntVal == 324739070 && __local_var_1.UnsafePtr == nil) {
__t1 = (*Constructor_Data_Map_Internal_Node)(nil)
goto end_branch_1
} else {

}
}
{
if (__local_var_2.Type == 9 && __local_var_2.IntVal == 324739070 && __local_var_2.UnsafePtr == nil) {
__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](__local_var_1)
goto end_branch_1
} else {

}
}
{
if (__local_var_2.Type == 9 && __local_var_2.IntVal == 324739070 && __local_var_2.UnsafePtr != nil) {
// TAST (Let): v_3_0 -> gopurs_runtime.Value
v_3_0 := gopurs_runtime.UncurriedApp3(Get_Data_Map_Internal_unsafeSplit(), __local_var_0, (*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V2, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](__local_var_1))})
_ = v_3_0
__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.UncurriedApp2(Get_Data_Map_Internal_unsafeJoinNodes(), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.UncurriedApp3(Get_Data_Map_Internal_unsafeDifference(), __local_var_0, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Split)(v_3_0.UnsafePtr).V1)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V4)})))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.UncurriedApp3(Get_Data_Map_Internal_unsafeDifference(), __local_var_0, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Split)(v_3_0.UnsafePtr).V2)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V5)})))}))
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_1:
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(__t1)}
}

func Call_Data_Map_Internal_unsafeIntersectionWith__4109280494(__local_var_0_loop gopurs_runtime.Value, __local_var_1_loop gopurs_runtime.Value, __local_var_2_loop gopurs_runtime.Value, __local_var_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __local_var_0 gopurs_runtime.Value = __local_var_0_loop
_ = __local_var_0
var __local_var_1 gopurs_runtime.Value = __local_var_1_loop
_ = __local_var_1
var __local_var_2 gopurs_runtime.Value = __local_var_2_loop
_ = __local_var_2
var __local_var_3 gopurs_runtime.Value = __local_var_3_loop
_ = __local_var_3
var __t6 *Constructor_Data_Map_Internal_Node
{
if (__local_var_2.Type == 9 && __local_var_2.IntVal == 324739070 && __local_var_2.UnsafePtr == nil) {
__t6 = (*Constructor_Data_Map_Internal_Node)(nil)
goto end_branch_6
} else {

}
}
{
if (__local_var_3.Type == 9 && __local_var_3.IntVal == 324739070 && __local_var_3.UnsafePtr == nil) {
__t6 = (*Constructor_Data_Map_Internal_Node)(nil)
goto end_branch_6
} else {

}
}
{
if (__local_var_3.Type == 9 && __local_var_3.IntVal == 324739070 && __local_var_3.UnsafePtr != nil) {
// TAST (Let): v_4_0 -> gopurs_runtime.Value
v_4_0 := gopurs_runtime.UncurriedApp3(Get_Data_Map_Internal_unsafeSplit(), __local_var_0, (*Constructor_Data_Map_Internal_Node)(__local_var_3.UnsafePtr).V2, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](__local_var_2))})
_ = v_4_0
// TAST (Let): l_prime_5_1 -> gopurs_runtime.Value
l_prime_5_1 := gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeIntersectionWith(), __local_var_0, __local_var_1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Split)(v_4_0.UnsafePtr).V1)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_3.UnsafePtr).V4)})
_ = l_prime_5_1
// TAST (Let): r_prime_6_2 -> gopurs_runtime.Value
r_prime_6_2 := gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeIntersectionWith(), __local_var_0, __local_var_1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Split)(v_4_0.UnsafePtr).V2)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_3.UnsafePtr).V5)})
_ = r_prime_6_2
var __t5 *Constructor_Data_Map_Internal_Node
{
var __t_tag_3 *Constructor_Data_Maybe_Just = (*Constructor_Data_Map_Internal_Split)(v_4_0.UnsafePtr).V0
if (__t_tag_3 != nil) {
__t5 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeBalancedNode(), (*Constructor_Data_Map_Internal_Node)(__local_var_3.UnsafePtr).V2, gopurs_runtime.Apply2(__local_var_1, ((*Constructor_Data_Map_Internal_Split)(v_4_0.UnsafePtr).V0).V0, (*Constructor_Data_Map_Internal_Node)(__local_var_3.UnsafePtr).V3), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](l_prime_5_1))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](r_prime_6_2))}))
goto end_branch_5
} else {

}
}
{
var __t_tag_4 *Constructor_Data_Maybe_Just = (*Constructor_Data_Map_Internal_Split)(v_4_0.UnsafePtr).V0
if (__t_tag_4 == nil) {
__t5 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.UncurriedApp2(Get_Data_Map_Internal_unsafeJoinNodes(), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](l_prime_5_1))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](r_prime_6_2))}))
goto end_branch_5
} else {

}
}
{
__t5 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_5:
__t6 = __t5
goto end_branch_6
} else {

}
}
{
__t6 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_6:
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(__t6)}
}

func Call_Data_Map_Internal_unsafeIntersectionWith__2517966(__local_var_0_loop gopurs_runtime.Value, __local_var_1_loop gopurs_runtime.Value, __local_var_2_loop gopurs_runtime.Value, __local_var_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __local_var_0 gopurs_runtime.Value = __local_var_0_loop
_ = __local_var_0
var __local_var_1 gopurs_runtime.Value = __local_var_1_loop
_ = __local_var_1
var __local_var_2 gopurs_runtime.Value = __local_var_2_loop
_ = __local_var_2
var __local_var_3 gopurs_runtime.Value = __local_var_3_loop
_ = __local_var_3
var __t6 *Constructor_Data_Map_Internal_Node
{
if (__local_var_2.Type == 9 && __local_var_2.IntVal == 324739070 && __local_var_2.UnsafePtr == nil) {
__t6 = (*Constructor_Data_Map_Internal_Node)(nil)
goto end_branch_6
} else {

}
}
{
if (__local_var_3.Type == 9 && __local_var_3.IntVal == 324739070 && __local_var_3.UnsafePtr == nil) {
__t6 = (*Constructor_Data_Map_Internal_Node)(nil)
goto end_branch_6
} else {

}
}
{
if (__local_var_3.Type == 9 && __local_var_3.IntVal == 324739070 && __local_var_3.UnsafePtr != nil) {
// TAST (Let): v_4_0 -> gopurs_runtime.Value
v_4_0 := gopurs_runtime.UncurriedApp3(Get_Data_Map_Internal_unsafeSplit(), __local_var_0, (*Constructor_Data_Map_Internal_Node)(__local_var_3.UnsafePtr).V2, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](__local_var_2))})
_ = v_4_0
// TAST (Let): l_prime_5_1 -> gopurs_runtime.Value
l_prime_5_1 := gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeIntersectionWith(), __local_var_0, __local_var_1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Split)(v_4_0.UnsafePtr).V1)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_3.UnsafePtr).V4)})
_ = l_prime_5_1
// TAST (Let): r_prime_6_2 -> gopurs_runtime.Value
r_prime_6_2 := gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeIntersectionWith(), __local_var_0, __local_var_1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Split)(v_4_0.UnsafePtr).V2)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_3.UnsafePtr).V5)})
_ = r_prime_6_2
var __t5 *Constructor_Data_Map_Internal_Node
{
var __t_tag_3 *Constructor_Data_Maybe_Just = (*Constructor_Data_Map_Internal_Split)(v_4_0.UnsafePtr).V0
if (__t_tag_3 != nil) {
__t5 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeBalancedNode(), (*Constructor_Data_Map_Internal_Node)(__local_var_3.UnsafePtr).V2, gopurs_runtime.Apply2(__local_var_1, ((*Constructor_Data_Map_Internal_Split)(v_4_0.UnsafePtr).V0).V0, (*Constructor_Data_Map_Internal_Node)(__local_var_3.UnsafePtr).V3), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](l_prime_5_1))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](r_prime_6_2))}))
goto end_branch_5
} else {

}
}
{
var __t_tag_4 *Constructor_Data_Maybe_Just = (*Constructor_Data_Map_Internal_Split)(v_4_0.UnsafePtr).V0
if (__t_tag_4 == nil) {
__t5 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.UncurriedApp2(Get_Data_Map_Internal_unsafeJoinNodes(), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](l_prime_5_1))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](r_prime_6_2))}))
goto end_branch_5
} else {

}
}
{
__t5 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_5:
__t6 = __t5
goto end_branch_6
} else {

}
}
{
__t6 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_6:
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(__t6)}
}

func Call_Data_Map_Internal_unsafeJoinNodes__2531831408(__local_var_0_loop gopurs_runtime.Value, __local_var_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __local_var_0 gopurs_runtime.Value = __local_var_0_loop
_ = __local_var_0
var __local_var_1 gopurs_runtime.Value = __local_var_1_loop
_ = __local_var_1
var __t1 *Constructor_Data_Map_Internal_Node
{
if (__local_var_0.Type == 9 && __local_var_0.IntVal == 324739070 && __local_var_0.UnsafePtr == nil) {
__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](__local_var_1)
goto end_branch_1
} else {

}
}
{
if (__local_var_0.Type == 9 && __local_var_0.IntVal == 324739070 && __local_var_0.UnsafePtr != nil) {
// TAST (Let): v2_2_0 -> gopurs_runtime.Value
v2_2_0 := gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeSplitLast(), (*Constructor_Data_Map_Internal_Node)(__local_var_0.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node)(__local_var_0.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_0.UnsafePtr).V4)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_0.UnsafePtr).V5)})
_ = v2_2_0
__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeBalancedNode(), (*Constructor_Data_Map_Internal_SplitLast)(v2_2_0.UnsafePtr).V0, (*Constructor_Data_Map_Internal_SplitLast)(v2_2_0.UnsafePtr).V1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_SplitLast)(v2_2_0.UnsafePtr).V2)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](__local_var_1))}))
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_1:
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(__t1)}
}

func Call_Data_Map_Internal_unsafeJoinNodes__3967876672(__local_var_0_loop gopurs_runtime.Value, __local_var_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __local_var_0 gopurs_runtime.Value = __local_var_0_loop
_ = __local_var_0
var __local_var_1 gopurs_runtime.Value = __local_var_1_loop
_ = __local_var_1
var __t1 *Constructor_Data_Map_Internal_Node
{
if (__local_var_0.Type == 9 && __local_var_0.IntVal == 324739070 && __local_var_0.UnsafePtr == nil) {
__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](__local_var_1)
goto end_branch_1
} else {

}
}
{
if (__local_var_0.Type == 9 && __local_var_0.IntVal == 324739070 && __local_var_0.UnsafePtr != nil) {
// TAST (Let): v2_2_0 -> gopurs_runtime.Value
v2_2_0 := gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeSplitLast(), (*Constructor_Data_Map_Internal_Node)(__local_var_0.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node)(__local_var_0.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_0.UnsafePtr).V4)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_0.UnsafePtr).V5)})
_ = v2_2_0
__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeBalancedNode(), (*Constructor_Data_Map_Internal_SplitLast)(v2_2_0.UnsafePtr).V0, (*Constructor_Data_Map_Internal_SplitLast)(v2_2_0.UnsafePtr).V1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_SplitLast)(v2_2_0.UnsafePtr).V2)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](__local_var_1))}))
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_1:
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(__t1)}
}

func Call_Data_Map_Internal_unsafeNode__1259503046(__local_var_0_loop gopurs_runtime.Value, __local_var_1_loop gopurs_runtime.Value, __local_var_2_loop gopurs_runtime.Value, __local_var_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __local_var_0 gopurs_runtime.Value = __local_var_0_loop
_ = __local_var_0
var __local_var_1 gopurs_runtime.Value = __local_var_1_loop
_ = __local_var_1
var __local_var_2 gopurs_runtime.Value = __local_var_2_loop
_ = __local_var_2
var __local_var_3 gopurs_runtime.Value = __local_var_3_loop
_ = __local_var_3
var __t4 *Constructor_Data_Map_Internal_Node
{
if (__local_var_2.Type == 9 && __local_var_2.IntVal == 324739070 && __local_var_2.UnsafePtr == nil) {
var __t0 *Constructor_Data_Map_Internal_Node
{
if (__local_var_3.Type == 9 && __local_var_3.IntVal == 324739070 && __local_var_3.UnsafePtr == nil) {
__t0 = &Constructor_Data_Map_Internal_Node{1, 1, 1, __local_var_0, __local_var_1, gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](__local_var_2), gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](__local_var_3)}
goto end_branch_0
} else {

}
}
{
if (__local_var_3.Type == 9 && __local_var_3.IntVal == 324739070 && __local_var_3.UnsafePtr != nil) {
__t0 = &Constructor_Data_Map_Internal_Node{1, (1) + ((*Constructor_Data_Map_Internal_Node)(__local_var_3.UnsafePtr).V0), (1) + ((*Constructor_Data_Map_Internal_Node)(__local_var_3.UnsafePtr).V1), __local_var_0, __local_var_1, gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](__local_var_2), gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](__local_var_3)}
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_0:
__t4 = __t0
goto end_branch_4
} else {

}
}
{
if (__local_var_2.Type == 9 && __local_var_2.IntVal == 324739070 && __local_var_2.UnsafePtr != nil) {
var __t3 *Constructor_Data_Map_Internal_Node
{
if (__local_var_3.Type == 9 && __local_var_3.IntVal == 324739070 && __local_var_3.UnsafePtr == nil) {
__t3 = &Constructor_Data_Map_Internal_Node{1, (1) + ((*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V0), (1) + ((*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V1), __local_var_0, __local_var_1, gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](__local_var_2), gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](__local_var_3)}
goto end_branch_3
} else {

}
}
{
if (__local_var_3.Type == 9 && __local_var_3.IntVal == 324739070 && __local_var_3.UnsafePtr != nil) {
var __t2 int64
{
var __t1 bool
{
if ((*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V0) > ((*Constructor_Data_Map_Internal_Node)(__local_var_3.UnsafePtr).V0) {
__t1 = true
goto end_branch_1
} else {

}
}
{
__t1 = false
}
end_branch_1:
if __t1 {
__t2 = (1) + ((*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V0)
goto end_branch_2
} else {

}
}
{
__t2 = (1) + ((*Constructor_Data_Map_Internal_Node)(__local_var_3.UnsafePtr).V0)
}
end_branch_2:
__t3 = &Constructor_Data_Map_Internal_Node{1, __t2, ((1) + ((*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V1)) + ((*Constructor_Data_Map_Internal_Node)(__local_var_3.UnsafePtr).V1), __local_var_0, __local_var_1, gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](__local_var_2), gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](__local_var_3)}
goto end_branch_3
} else {

}
}
{
__t3 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_3:
__t4 = __t3
goto end_branch_4
} else {

}
}
{
__t4 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_4:
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(__t4)}
}

func Call_Data_Map_Internal_unsafeNode__1305301638(__local_var_0_loop gopurs_runtime.Value, __local_var_1_loop gopurs_runtime.Value, __local_var_2_loop gopurs_runtime.Value, __local_var_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __local_var_0 gopurs_runtime.Value = __local_var_0_loop
_ = __local_var_0
var __local_var_1 gopurs_runtime.Value = __local_var_1_loop
_ = __local_var_1
var __local_var_2 gopurs_runtime.Value = __local_var_2_loop
_ = __local_var_2
var __local_var_3 gopurs_runtime.Value = __local_var_3_loop
_ = __local_var_3
var __t4 *Constructor_Data_Map_Internal_Node
{
if (__local_var_2.Type == 9 && __local_var_2.IntVal == 324739070 && __local_var_2.UnsafePtr == nil) {
var __t0 *Constructor_Data_Map_Internal_Node
{
if (__local_var_3.Type == 9 && __local_var_3.IntVal == 324739070 && __local_var_3.UnsafePtr == nil) {
__t0 = &Constructor_Data_Map_Internal_Node{1, 1, 1, __local_var_0, __local_var_1, gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](__local_var_2), gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](__local_var_3)}
goto end_branch_0
} else {

}
}
{
if (__local_var_3.Type == 9 && __local_var_3.IntVal == 324739070 && __local_var_3.UnsafePtr != nil) {
__t0 = &Constructor_Data_Map_Internal_Node{1, (1) + ((*Constructor_Data_Map_Internal_Node)(__local_var_3.UnsafePtr).V0), (1) + ((*Constructor_Data_Map_Internal_Node)(__local_var_3.UnsafePtr).V1), __local_var_0, __local_var_1, gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](__local_var_2), gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](__local_var_3)}
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_0:
__t4 = __t0
goto end_branch_4
} else {

}
}
{
if (__local_var_2.Type == 9 && __local_var_2.IntVal == 324739070 && __local_var_2.UnsafePtr != nil) {
var __t3 *Constructor_Data_Map_Internal_Node
{
if (__local_var_3.Type == 9 && __local_var_3.IntVal == 324739070 && __local_var_3.UnsafePtr == nil) {
__t3 = &Constructor_Data_Map_Internal_Node{1, (1) + ((*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V0), (1) + ((*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V1), __local_var_0, __local_var_1, gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](__local_var_2), gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](__local_var_3)}
goto end_branch_3
} else {

}
}
{
if (__local_var_3.Type == 9 && __local_var_3.IntVal == 324739070 && __local_var_3.UnsafePtr != nil) {
var __t2 int64
{
var __t1 bool
{
if ((*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V0) > ((*Constructor_Data_Map_Internal_Node)(__local_var_3.UnsafePtr).V0) {
__t1 = true
goto end_branch_1
} else {

}
}
{
__t1 = false
}
end_branch_1:
if __t1 {
__t2 = (1) + ((*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V0)
goto end_branch_2
} else {

}
}
{
__t2 = (1) + ((*Constructor_Data_Map_Internal_Node)(__local_var_3.UnsafePtr).V0)
}
end_branch_2:
__t3 = &Constructor_Data_Map_Internal_Node{1, __t2, ((1) + ((*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V1)) + ((*Constructor_Data_Map_Internal_Node)(__local_var_3.UnsafePtr).V1), __local_var_0, __local_var_1, gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](__local_var_2), gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](__local_var_3)}
goto end_branch_3
} else {

}
}
{
__t3 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_3:
__t4 = __t3
goto end_branch_4
} else {

}
}
{
__t4 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_4:
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(__t4)}
}

func Call_Data_Map_Internal_unsafeNode__954819782(__local_var_0_loop gopurs_runtime.Value, __local_var_1_loop gopurs_runtime.Value, __local_var_2_loop gopurs_runtime.Value, __local_var_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __local_var_0 gopurs_runtime.Value = __local_var_0_loop
_ = __local_var_0
var __local_var_1 gopurs_runtime.Value = __local_var_1_loop
_ = __local_var_1
var __local_var_2 gopurs_runtime.Value = __local_var_2_loop
_ = __local_var_2
var __local_var_3 gopurs_runtime.Value = __local_var_3_loop
_ = __local_var_3
var __t4 *Constructor_Data_Map_Internal_Node
{
if (__local_var_2.Type == 9 && __local_var_2.IntVal == 324739070 && __local_var_2.UnsafePtr == nil) {
var __t0 *Constructor_Data_Map_Internal_Node
{
if (__local_var_3.Type == 9 && __local_var_3.IntVal == 324739070 && __local_var_3.UnsafePtr == nil) {
__t0 = &Constructor_Data_Map_Internal_Node{1, 1, 1, __local_var_0, __local_var_1, gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](__local_var_2), gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](__local_var_3)}
goto end_branch_0
} else {

}
}
{
if (__local_var_3.Type == 9 && __local_var_3.IntVal == 324739070 && __local_var_3.UnsafePtr != nil) {
__t0 = &Constructor_Data_Map_Internal_Node{1, (1) + ((*Constructor_Data_Map_Internal_Node)(__local_var_3.UnsafePtr).V0), (1) + ((*Constructor_Data_Map_Internal_Node)(__local_var_3.UnsafePtr).V1), __local_var_0, __local_var_1, gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](__local_var_2), gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](__local_var_3)}
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_0:
__t4 = __t0
goto end_branch_4
} else {

}
}
{
if (__local_var_2.Type == 9 && __local_var_2.IntVal == 324739070 && __local_var_2.UnsafePtr != nil) {
var __t3 *Constructor_Data_Map_Internal_Node
{
if (__local_var_3.Type == 9 && __local_var_3.IntVal == 324739070 && __local_var_3.UnsafePtr == nil) {
__t3 = &Constructor_Data_Map_Internal_Node{1, (1) + ((*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V0), (1) + ((*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V1), __local_var_0, __local_var_1, gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](__local_var_2), gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](__local_var_3)}
goto end_branch_3
} else {

}
}
{
if (__local_var_3.Type == 9 && __local_var_3.IntVal == 324739070 && __local_var_3.UnsafePtr != nil) {
var __t2 int64
{
var __t1 bool
{
if ((*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V0) > ((*Constructor_Data_Map_Internal_Node)(__local_var_3.UnsafePtr).V0) {
__t1 = true
goto end_branch_1
} else {

}
}
{
__t1 = false
}
end_branch_1:
if __t1 {
__t2 = (1) + ((*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V0)
goto end_branch_2
} else {

}
}
{
__t2 = (1) + ((*Constructor_Data_Map_Internal_Node)(__local_var_3.UnsafePtr).V0)
}
end_branch_2:
__t3 = &Constructor_Data_Map_Internal_Node{1, __t2, ((1) + ((*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V1)) + ((*Constructor_Data_Map_Internal_Node)(__local_var_3.UnsafePtr).V1), __local_var_0, __local_var_1, gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](__local_var_2), gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](__local_var_3)}
goto end_branch_3
} else {

}
}
{
__t3 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_3:
__t4 = __t3
goto end_branch_4
} else {

}
}
{
__t4 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_4:
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(__t4)}
}

func Call_Data_Map_Internal_unsafeNode__1776657286(__local_var_0_loop gopurs_runtime.Value, __local_var_1_loop gopurs_runtime.Value, __local_var_2_loop gopurs_runtime.Value, __local_var_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __local_var_0 gopurs_runtime.Value = __local_var_0_loop
_ = __local_var_0
var __local_var_1 gopurs_runtime.Value = __local_var_1_loop
_ = __local_var_1
var __local_var_2 gopurs_runtime.Value = __local_var_2_loop
_ = __local_var_2
var __local_var_3 gopurs_runtime.Value = __local_var_3_loop
_ = __local_var_3
var __t4 *Constructor_Data_Map_Internal_Node
{
if (__local_var_2.Type == 9 && __local_var_2.IntVal == 324739070 && __local_var_2.UnsafePtr == nil) {
var __t0 *Constructor_Data_Map_Internal_Node
{
if (__local_var_3.Type == 9 && __local_var_3.IntVal == 324739070 && __local_var_3.UnsafePtr == nil) {
__t0 = &Constructor_Data_Map_Internal_Node{1, 1, 1, gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(__local_var_0.IntVal)), UnsafePtr: nil}, gopurs_runtime.Float(__local_var_1.FloatVal()), gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](__local_var_2))}), gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](__local_var_3))})}
goto end_branch_0
} else {

}
}
{
if (__local_var_3.Type == 9 && __local_var_3.IntVal == 324739070 && __local_var_3.UnsafePtr != nil) {
__t0 = &Constructor_Data_Map_Internal_Node{1, (1) + ((*Constructor_Data_Map_Internal_Node)(__local_var_3.UnsafePtr).V0), (1) + ((*Constructor_Data_Map_Internal_Node)(__local_var_3.UnsafePtr).V1), gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(__local_var_0.IntVal)), UnsafePtr: nil}, gopurs_runtime.Float(__local_var_1.FloatVal()), gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](__local_var_2))}), gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](__local_var_3))})}
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_0:
__t4 = __t0
goto end_branch_4
} else {

}
}
{
if (__local_var_2.Type == 9 && __local_var_2.IntVal == 324739070 && __local_var_2.UnsafePtr != nil) {
var __t3 *Constructor_Data_Map_Internal_Node
{
if (__local_var_3.Type == 9 && __local_var_3.IntVal == 324739070 && __local_var_3.UnsafePtr == nil) {
__t3 = &Constructor_Data_Map_Internal_Node{1, (1) + ((*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V0), (1) + ((*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V1), gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(__local_var_0.IntVal)), UnsafePtr: nil}, gopurs_runtime.Float(__local_var_1.FloatVal()), gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](__local_var_2))}), gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](__local_var_3))})}
goto end_branch_3
} else {

}
}
{
if (__local_var_3.Type == 9 && __local_var_3.IntVal == 324739070 && __local_var_3.UnsafePtr != nil) {
var __t2 int64
{
var __t1 bool
{
if ((*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V0) > ((*Constructor_Data_Map_Internal_Node)(__local_var_3.UnsafePtr).V0) {
__t1 = true
goto end_branch_1
} else {

}
}
{
__t1 = false
}
end_branch_1:
if __t1 {
__t2 = (1) + ((*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V0)
goto end_branch_2
} else {

}
}
{
__t2 = (1) + ((*Constructor_Data_Map_Internal_Node)(__local_var_3.UnsafePtr).V0)
}
end_branch_2:
__t3 = &Constructor_Data_Map_Internal_Node{1, __t2, ((1) + ((*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V1)) + ((*Constructor_Data_Map_Internal_Node)(__local_var_3.UnsafePtr).V1), gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(__local_var_0.IntVal)), UnsafePtr: nil}, gopurs_runtime.Float(__local_var_1.FloatVal()), gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](__local_var_2))}), gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](__local_var_3))})}
goto end_branch_3
} else {

}
}
{
__t3 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_3:
__t4 = __t3
goto end_branch_4
} else {

}
}
{
__t4 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_4:
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(__t4)}
}

func Call_Data_Map_Internal_unsafeNode__1902536198(__local_var_0_loop gopurs_runtime.Value, __local_var_1_loop gopurs_runtime.Value, __local_var_2_loop gopurs_runtime.Value, __local_var_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __local_var_0 gopurs_runtime.Value = __local_var_0_loop
_ = __local_var_0
var __local_var_1 gopurs_runtime.Value = __local_var_1_loop
_ = __local_var_1
var __local_var_2 gopurs_runtime.Value = __local_var_2_loop
_ = __local_var_2
var __local_var_3 gopurs_runtime.Value = __local_var_3_loop
_ = __local_var_3
var __t4 *Constructor_Data_Map_Internal_Node
{
if (__local_var_2.Type == 9 && __local_var_2.IntVal == 324739070 && __local_var_2.UnsafePtr == nil) {
var __t0 *Constructor_Data_Map_Internal_Node
{
if (__local_var_3.Type == 9 && __local_var_3.IntVal == 324739070 && __local_var_3.UnsafePtr == nil) {
__t0 = &Constructor_Data_Map_Internal_Node{1, 1, 1, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](__local_var_0))}, __local_var_1, gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](__local_var_2))}), gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](__local_var_3))})}
goto end_branch_0
} else {

}
}
{
if (__local_var_3.Type == 9 && __local_var_3.IntVal == 324739070 && __local_var_3.UnsafePtr != nil) {
__t0 = &Constructor_Data_Map_Internal_Node{1, (1) + ((*Constructor_Data_Map_Internal_Node)(__local_var_3.UnsafePtr).V0), (1) + ((*Constructor_Data_Map_Internal_Node)(__local_var_3.UnsafePtr).V1), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](__local_var_0))}, __local_var_1, gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](__local_var_2))}), gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](__local_var_3))})}
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_0:
__t4 = __t0
goto end_branch_4
} else {

}
}
{
if (__local_var_2.Type == 9 && __local_var_2.IntVal == 324739070 && __local_var_2.UnsafePtr != nil) {
var __t3 *Constructor_Data_Map_Internal_Node
{
if (__local_var_3.Type == 9 && __local_var_3.IntVal == 324739070 && __local_var_3.UnsafePtr == nil) {
__t3 = &Constructor_Data_Map_Internal_Node{1, (1) + ((*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V0), (1) + ((*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V1), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](__local_var_0))}, __local_var_1, gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](__local_var_2))}), gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](__local_var_3))})}
goto end_branch_3
} else {

}
}
{
if (__local_var_3.Type == 9 && __local_var_3.IntVal == 324739070 && __local_var_3.UnsafePtr != nil) {
var __t2 int64
{
var __t1 bool
{
if ((*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V0) > ((*Constructor_Data_Map_Internal_Node)(__local_var_3.UnsafePtr).V0) {
__t1 = true
goto end_branch_1
} else {

}
}
{
__t1 = false
}
end_branch_1:
if __t1 {
__t2 = (1) + ((*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V0)
goto end_branch_2
} else {

}
}
{
__t2 = (1) + ((*Constructor_Data_Map_Internal_Node)(__local_var_3.UnsafePtr).V0)
}
end_branch_2:
__t3 = &Constructor_Data_Map_Internal_Node{1, __t2, ((1) + ((*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V1)) + ((*Constructor_Data_Map_Internal_Node)(__local_var_3.UnsafePtr).V1), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](__local_var_0))}, __local_var_1, gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](__local_var_2))}), gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](__local_var_3))})}
goto end_branch_3
} else {

}
}
{
__t3 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_3:
__t4 = __t3
goto end_branch_4
} else {

}
}
{
__t4 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_4:
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(__t4)}
}

func Call_Data_Map_Internal_unsafeSplit__1094566431(__local_var_0_loop gopurs_runtime.Value, __local_var_1_loop gopurs_runtime.Value, __local_var_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __local_var_0 gopurs_runtime.Value = __local_var_0_loop
_ = __local_var_0
var __local_var_1 gopurs_runtime.Value = __local_var_1_loop
_ = __local_var_1
var __local_var_2 gopurs_runtime.Value = __local_var_2_loop
_ = __local_var_2
var __t4 *Constructor_Data_Map_Internal_Split
{
if (__local_var_2.Type == 9 && __local_var_2.IntVal == 324739070 && __local_var_2.UnsafePtr == nil) {
__t4 = &Constructor_Data_Map_Internal_Split{1, (*Constructor_Data_Maybe_Just)(nil), (*Constructor_Data_Map_Internal_Node)(nil), (*Constructor_Data_Map_Internal_Node)(nil)}
goto end_branch_4
} else {

}
}
{
if (__local_var_2.Type == 9 && __local_var_2.IntVal == 324739070 && __local_var_2.UnsafePtr != nil) {
// TAST (Let): v_3_0 -> gopurs_runtime.Value
v_3_0 := gopurs_runtime.Apply2(__local_var_0, __local_var_1, (*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V2)
_ = v_3_0
var __t3 *Constructor_Data_Map_Internal_Split
{
if (uint32(v_3_0.IntVal) == 1527465420) {
// TAST (Let): v1_4_1 -> gopurs_runtime.Value
v1_4_1 := gopurs_runtime.UncurriedApp3(Get_Data_Map_Internal_unsafeSplit(), __local_var_0, __local_var_1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V4)})
_ = v1_4_1
__t3 = &Constructor_Data_Map_Internal_Split{1, (*Constructor_Data_Map_Internal_Split)(v1_4_1.UnsafePtr).V0, (*Constructor_Data_Map_Internal_Split)(v1_4_1.UnsafePtr).V1, gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeBalancedNode(), (*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Split)(v1_4_1.UnsafePtr).V2)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V5)}))}
goto end_branch_3
} else {

}
}
{
if (uint32(v_3_0.IntVal) == 380165415) {
// TAST (Let): v1_4_2 -> gopurs_runtime.Value
v1_4_2 := gopurs_runtime.UncurriedApp3(Get_Data_Map_Internal_unsafeSplit(), __local_var_0, __local_var_1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V5)})
_ = v1_4_2
__t3 = &Constructor_Data_Map_Internal_Split{1, (*Constructor_Data_Map_Internal_Split)(v1_4_2.UnsafePtr).V0, gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeBalancedNode(), (*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V4)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Split)(v1_4_2.UnsafePtr).V1)})), (*Constructor_Data_Map_Internal_Split)(v1_4_2.UnsafePtr).V2}
goto end_branch_3
} else {

}
}
{
if (uint32(v_3_0.IntVal) == 902936544) {
__t3 = &Constructor_Data_Map_Internal_Split{1, &Constructor_Data_Maybe_Just{1, (*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V3}, (*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V4, (*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V5}
goto end_branch_3
} else {

}
}
{
__t3 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Split](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_3:
__t4 = __t3
goto end_branch_4
} else {

}
}
{
__t4 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Split](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_4:
return gopurs_runtime.Value{Type: 9, IntVal: 3373277644, UnsafePtr: unsafe.Pointer(__t4)}
}

func Call_Data_Map_Internal_unsafeSplit__4154869695(__local_var_0_loop gopurs_runtime.Value, __local_var_1_loop gopurs_runtime.Value, __local_var_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __local_var_0 gopurs_runtime.Value = __local_var_0_loop
_ = __local_var_0
var __local_var_1 gopurs_runtime.Value = __local_var_1_loop
_ = __local_var_1
var __local_var_2 gopurs_runtime.Value = __local_var_2_loop
_ = __local_var_2
var __t4 *Constructor_Data_Map_Internal_Split
{
if (__local_var_2.Type == 9 && __local_var_2.IntVal == 324739070 && __local_var_2.UnsafePtr == nil) {
__t4 = &Constructor_Data_Map_Internal_Split{1, (*Constructor_Data_Maybe_Just)(nil), (*Constructor_Data_Map_Internal_Node)(nil), (*Constructor_Data_Map_Internal_Node)(nil)}
goto end_branch_4
} else {

}
}
{
if (__local_var_2.Type == 9 && __local_var_2.IntVal == 324739070 && __local_var_2.UnsafePtr != nil) {
// TAST (Let): v_3_0 -> gopurs_runtime.Value
v_3_0 := gopurs_runtime.Apply2(__local_var_0, __local_var_1, (*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V2)
_ = v_3_0
var __t3 *Constructor_Data_Map_Internal_Split
{
if (uint32(v_3_0.IntVal) == 1527465420) {
// TAST (Let): v1_4_1 -> gopurs_runtime.Value
v1_4_1 := gopurs_runtime.UncurriedApp3(Get_Data_Map_Internal_unsafeSplit(), __local_var_0, __local_var_1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V4)})
_ = v1_4_1
__t3 = &Constructor_Data_Map_Internal_Split{1, (*Constructor_Data_Map_Internal_Split)(v1_4_1.UnsafePtr).V0, (*Constructor_Data_Map_Internal_Split)(v1_4_1.UnsafePtr).V1, gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeBalancedNode(), (*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Split)(v1_4_1.UnsafePtr).V2)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V5)}))}
goto end_branch_3
} else {

}
}
{
if (uint32(v_3_0.IntVal) == 380165415) {
// TAST (Let): v1_4_2 -> gopurs_runtime.Value
v1_4_2 := gopurs_runtime.UncurriedApp3(Get_Data_Map_Internal_unsafeSplit(), __local_var_0, __local_var_1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V5)})
_ = v1_4_2
__t3 = &Constructor_Data_Map_Internal_Split{1, (*Constructor_Data_Map_Internal_Split)(v1_4_2.UnsafePtr).V0, gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeBalancedNode(), (*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V4)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Split)(v1_4_2.UnsafePtr).V1)})), (*Constructor_Data_Map_Internal_Split)(v1_4_2.UnsafePtr).V2}
goto end_branch_3
} else {

}
}
{
if (uint32(v_3_0.IntVal) == 902936544) {
__t3 = &Constructor_Data_Map_Internal_Split{1, &Constructor_Data_Maybe_Just{1, (*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V3}, (*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V4, (*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V5}
goto end_branch_3
} else {

}
}
{
__t3 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Split](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_3:
__t4 = __t3
goto end_branch_4
} else {

}
}
{
__t4 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Split](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_4:
return gopurs_runtime.Value{Type: 9, IntVal: 3373277644, UnsafePtr: unsafe.Pointer(__t4)}
}

func Call_Data_Map_Internal_unsafeSplit__1308258847(__local_var_0_loop gopurs_runtime.Value, __local_var_1_loop gopurs_runtime.Value, __local_var_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __local_var_0 gopurs_runtime.Value = __local_var_0_loop
_ = __local_var_0
var __local_var_1 gopurs_runtime.Value = __local_var_1_loop
_ = __local_var_1
var __local_var_2 gopurs_runtime.Value = __local_var_2_loop
_ = __local_var_2
var __t4 *Constructor_Data_Map_Internal_Split
{
if (__local_var_2.Type == 9 && __local_var_2.IntVal == 324739070 && __local_var_2.UnsafePtr == nil) {
__t4 = &Constructor_Data_Map_Internal_Split{1, (*Constructor_Data_Maybe_Just)(nil), (*Constructor_Data_Map_Internal_Node)(nil), (*Constructor_Data_Map_Internal_Node)(nil)}
goto end_branch_4
} else {

}
}
{
if (__local_var_2.Type == 9 && __local_var_2.IntVal == 324739070 && __local_var_2.UnsafePtr != nil) {
// TAST (Let): v_3_0 -> gopurs_runtime.Value
v_3_0 := gopurs_runtime.Apply2(__local_var_0, __local_var_1, (*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V2)
_ = v_3_0
var __t3 *Constructor_Data_Map_Internal_Split
{
if (uint32(v_3_0.IntVal) == 1527465420) {
// TAST (Let): v1_4_1 -> gopurs_runtime.Value
v1_4_1 := gopurs_runtime.UncurriedApp3(Get_Data_Map_Internal_unsafeSplit(), __local_var_0, __local_var_1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V4)})
_ = v1_4_1
__t3 = &Constructor_Data_Map_Internal_Split{1, (*Constructor_Data_Map_Internal_Split)(v1_4_1.UnsafePtr).V0, (*Constructor_Data_Map_Internal_Split)(v1_4_1.UnsafePtr).V1, gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeBalancedNode(), (*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Split)(v1_4_1.UnsafePtr).V2)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V5)}))}
goto end_branch_3
} else {

}
}
{
if (uint32(v_3_0.IntVal) == 380165415) {
// TAST (Let): v1_4_2 -> gopurs_runtime.Value
v1_4_2 := gopurs_runtime.UncurriedApp3(Get_Data_Map_Internal_unsafeSplit(), __local_var_0, __local_var_1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V5)})
_ = v1_4_2
__t3 = &Constructor_Data_Map_Internal_Split{1, (*Constructor_Data_Map_Internal_Split)(v1_4_2.UnsafePtr).V0, gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeBalancedNode(), (*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V4)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Split)(v1_4_2.UnsafePtr).V1)})), (*Constructor_Data_Map_Internal_Split)(v1_4_2.UnsafePtr).V2}
goto end_branch_3
} else {

}
}
{
if (uint32(v_3_0.IntVal) == 902936544) {
__t3 = &Constructor_Data_Map_Internal_Split{1, &Constructor_Data_Maybe_Just{1, (*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V3}, (*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V4, (*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V5}
goto end_branch_3
} else {

}
}
{
__t3 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Split](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_3:
__t4 = __t3
goto end_branch_4
} else {

}
}
{
__t4 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Split](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_4:
return gopurs_runtime.Value{Type: 9, IntVal: 3373277644, UnsafePtr: unsafe.Pointer(__t4)}
}

func Call_Data_Map_Internal_unsafeSplit__1115245464(__local_var_0_loop gopurs_runtime.Value, __local_var_1_loop gopurs_runtime.Value, __local_var_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __local_var_0 gopurs_runtime.Value = __local_var_0_loop
_ = __local_var_0
var __local_var_1 gopurs_runtime.Value = __local_var_1_loop
_ = __local_var_1
var __local_var_2 gopurs_runtime.Value = __local_var_2_loop
_ = __local_var_2
var __t4 *Constructor_Data_Map_Internal_Split
{
if (__local_var_2.Type == 9 && __local_var_2.IntVal == 324739070 && __local_var_2.UnsafePtr == nil) {
__t4 = &Constructor_Data_Map_Internal_Split{1, (*Constructor_Data_Maybe_Just)(nil), (*Constructor_Data_Map_Internal_Node)(nil), (*Constructor_Data_Map_Internal_Node)(nil)}
goto end_branch_4
} else {

}
}
{
if (__local_var_2.Type == 9 && __local_var_2.IntVal == 324739070 && __local_var_2.UnsafePtr != nil) {
// TAST (Let): v_3_0 -> gopurs_runtime.Value
v_3_0 := gopurs_runtime.Apply2(__local_var_0, gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(__local_var_1.IntVal)), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(uint32((*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V2.IntVal)), UnsafePtr: nil})
_ = v_3_0
var __t3 *Constructor_Data_Map_Internal_Split
{
if (uint32(v_3_0.IntVal) == 1527465420) {
// TAST (Let): v1_4_1 -> gopurs_runtime.Value
v1_4_1 := gopurs_runtime.UncurriedApp3(Get_Data_Map_Internal_unsafeSplit(), __local_var_0, gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(__local_var_1.IntVal)), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V4)})
_ = v1_4_1
__t3 = &Constructor_Data_Map_Internal_Split{1, (*Constructor_Data_Map_Internal_Split)(v1_4_1.UnsafePtr).V0, (*Constructor_Data_Map_Internal_Split)(v1_4_1.UnsafePtr).V1, gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeBalancedNode(), gopurs_runtime.Value{Type: 9, IntVal: int64(uint32((*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V2.IntVal)), UnsafePtr: nil}, gopurs_runtime.Float((*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V3.FloatVal()), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Split)(v1_4_1.UnsafePtr).V2)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V5)})))})}
goto end_branch_3
} else {

}
}
{
if (uint32(v_3_0.IntVal) == 380165415) {
// TAST (Let): v1_4_2 -> gopurs_runtime.Value
v1_4_2 := gopurs_runtime.UncurriedApp3(Get_Data_Map_Internal_unsafeSplit(), __local_var_0, gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(__local_var_1.IntVal)), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V5)})
_ = v1_4_2
__t3 = &Constructor_Data_Map_Internal_Split{1, (*Constructor_Data_Map_Internal_Split)(v1_4_2.UnsafePtr).V0, gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeBalancedNode(), gopurs_runtime.Value{Type: 9, IntVal: int64(uint32((*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V2.IntVal)), UnsafePtr: nil}, gopurs_runtime.Float((*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V3.FloatVal()), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V4)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Split)(v1_4_2.UnsafePtr).V1)})))}), (*Constructor_Data_Map_Internal_Split)(v1_4_2.UnsafePtr).V2}
goto end_branch_3
} else {

}
}
{
if (uint32(v_3_0.IntVal) == 902936544) {
__t3 = &Constructor_Data_Map_Internal_Split{1, &Constructor_Data_Maybe_Just{1, gopurs_runtime.Float((*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V3.FloatVal())}, (*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V4, (*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V5}
goto end_branch_3
} else {

}
}
{
__t3 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Split](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_3:
__t4 = __t3
goto end_branch_4
} else {

}
}
{
__t4 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Split](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_4:
return gopurs_runtime.Value{Type: 9, IntVal: 3373277644, UnsafePtr: unsafe.Pointer(__t4)}
}

func Call_Data_Map_Internal_unsafeSplitLast__1494186946(__local_var_0_loop gopurs_runtime.Value, __local_var_1_loop gopurs_runtime.Value, __local_var_2_loop gopurs_runtime.Value, __local_var_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __local_var_0 gopurs_runtime.Value = __local_var_0_loop
_ = __local_var_0
var __local_var_1 gopurs_runtime.Value = __local_var_1_loop
_ = __local_var_1
var __local_var_2 gopurs_runtime.Value = __local_var_2_loop
_ = __local_var_2
var __local_var_3 gopurs_runtime.Value = __local_var_3_loop
_ = __local_var_3
var __t1 *Constructor_Data_Map_Internal_SplitLast
{
if (__local_var_3.Type == 9 && __local_var_3.IntVal == 324739070 && __local_var_3.UnsafePtr == nil) {
__t1 = &Constructor_Data_Map_Internal_SplitLast{1, __local_var_0, __local_var_1, gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](__local_var_2)}
goto end_branch_1
} else {

}
}
{
if (__local_var_3.Type == 9 && __local_var_3.IntVal == 324739070 && __local_var_3.UnsafePtr != nil) {
// TAST (Let): v1_4_0 -> gopurs_runtime.Value
v1_4_0 := gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeSplitLast(), (*Constructor_Data_Map_Internal_Node)(__local_var_3.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node)(__local_var_3.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_3.UnsafePtr).V4)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_3.UnsafePtr).V5)})
_ = v1_4_0
__t1 = &Constructor_Data_Map_Internal_SplitLast{1, (*Constructor_Data_Map_Internal_SplitLast)(v1_4_0.UnsafePtr).V0, (*Constructor_Data_Map_Internal_SplitLast)(v1_4_0.UnsafePtr).V1, gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeBalancedNode(), __local_var_0, __local_var_1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](__local_var_2))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_SplitLast)(v1_4_0.UnsafePtr).V2)}))}
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_SplitLast](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_1:
return gopurs_runtime.Value{Type: 9, IntVal: 2668112006, UnsafePtr: unsafe.Pointer(__t1)}
}

func Call_Data_Map_Internal_unsafeSplitLast__224676098(__local_var_0_loop gopurs_runtime.Value, __local_var_1_loop gopurs_runtime.Value, __local_var_2_loop gopurs_runtime.Value, __local_var_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __local_var_0 gopurs_runtime.Value = __local_var_0_loop
_ = __local_var_0
var __local_var_1 gopurs_runtime.Value = __local_var_1_loop
_ = __local_var_1
var __local_var_2 gopurs_runtime.Value = __local_var_2_loop
_ = __local_var_2
var __local_var_3 gopurs_runtime.Value = __local_var_3_loop
_ = __local_var_3
var __t1 *Constructor_Data_Map_Internal_SplitLast
{
if (__local_var_3.Type == 9 && __local_var_3.IntVal == 324739070 && __local_var_3.UnsafePtr == nil) {
__t1 = &Constructor_Data_Map_Internal_SplitLast{1, __local_var_0, __local_var_1, gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](__local_var_2)}
goto end_branch_1
} else {

}
}
{
if (__local_var_3.Type == 9 && __local_var_3.IntVal == 324739070 && __local_var_3.UnsafePtr != nil) {
// TAST (Let): v1_4_0 -> gopurs_runtime.Value
v1_4_0 := gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeSplitLast(), (*Constructor_Data_Map_Internal_Node)(__local_var_3.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node)(__local_var_3.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_3.UnsafePtr).V4)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_3.UnsafePtr).V5)})
_ = v1_4_0
__t1 = &Constructor_Data_Map_Internal_SplitLast{1, (*Constructor_Data_Map_Internal_SplitLast)(v1_4_0.UnsafePtr).V0, (*Constructor_Data_Map_Internal_SplitLast)(v1_4_0.UnsafePtr).V1, gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeBalancedNode(), __local_var_0, __local_var_1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](__local_var_2))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_SplitLast)(v1_4_0.UnsafePtr).V2)}))}
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_SplitLast](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_1:
return gopurs_runtime.Value{Type: 9, IntVal: 2668112006, UnsafePtr: unsafe.Pointer(__t1)}
}

func Call_Data_Map_Internal_unsafeUnionWith__4109280494(__local_var_0_loop gopurs_runtime.Value, __local_var_1_loop gopurs_runtime.Value, __local_var_2_loop gopurs_runtime.Value, __local_var_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __local_var_0 gopurs_runtime.Value = __local_var_0_loop
_ = __local_var_0
var __local_var_1 gopurs_runtime.Value = __local_var_1_loop
_ = __local_var_1
var __local_var_2 gopurs_runtime.Value = __local_var_2_loop
_ = __local_var_2
var __local_var_3 gopurs_runtime.Value = __local_var_3_loop
_ = __local_var_3
var __t6 *Constructor_Data_Map_Internal_Node
{
if (__local_var_2.Type == 9 && __local_var_2.IntVal == 324739070 && __local_var_2.UnsafePtr == nil) {
__t6 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](__local_var_3)
goto end_branch_6
} else {

}
}
{
if (__local_var_3.Type == 9 && __local_var_3.IntVal == 324739070 && __local_var_3.UnsafePtr == nil) {
__t6 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](__local_var_2)
goto end_branch_6
} else {

}
}
{
if (__local_var_3.Type == 9 && __local_var_3.IntVal == 324739070 && __local_var_3.UnsafePtr != nil) {
// TAST (Let): v_4_0 -> gopurs_runtime.Value
v_4_0 := gopurs_runtime.UncurriedApp3(Get_Data_Map_Internal_unsafeSplit(), __local_var_0, (*Constructor_Data_Map_Internal_Node)(__local_var_3.UnsafePtr).V2, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](__local_var_2))})
_ = v_4_0
// TAST (Let): l_prime_5_1 -> gopurs_runtime.Value
l_prime_5_1 := gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeUnionWith(), __local_var_0, __local_var_1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Split)(v_4_0.UnsafePtr).V1)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_3.UnsafePtr).V4)})
_ = l_prime_5_1
// TAST (Let): r_prime_6_2 -> gopurs_runtime.Value
r_prime_6_2 := gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeUnionWith(), __local_var_0, __local_var_1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Split)(v_4_0.UnsafePtr).V2)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_3.UnsafePtr).V5)})
_ = r_prime_6_2
var __t5 *Constructor_Data_Map_Internal_Node
{
var __t_tag_3 *Constructor_Data_Maybe_Just = (*Constructor_Data_Map_Internal_Split)(v_4_0.UnsafePtr).V0
if (__t_tag_3 != nil) {
__t5 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeBalancedNode(), (*Constructor_Data_Map_Internal_Node)(__local_var_3.UnsafePtr).V2, gopurs_runtime.Apply2(__local_var_1, ((*Constructor_Data_Map_Internal_Split)(v_4_0.UnsafePtr).V0).V0, (*Constructor_Data_Map_Internal_Node)(__local_var_3.UnsafePtr).V3), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](l_prime_5_1))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](r_prime_6_2))}))
goto end_branch_5
} else {

}
}
{
var __t_tag_4 *Constructor_Data_Maybe_Just = (*Constructor_Data_Map_Internal_Split)(v_4_0.UnsafePtr).V0
if (__t_tag_4 == nil) {
__t5 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeBalancedNode(), (*Constructor_Data_Map_Internal_Node)(__local_var_3.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node)(__local_var_3.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](l_prime_5_1))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](r_prime_6_2))}))
goto end_branch_5
} else {

}
}
{
__t5 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_5:
__t6 = __t5
goto end_branch_6
} else {

}
}
{
__t6 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_6:
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(__t6)}
}

func Call_Data_Map_Internal_unsafeUnionWith__3421363785(__local_var_0_loop gopurs_runtime.Value, __local_var_1_loop gopurs_runtime.Value, __local_var_2_loop gopurs_runtime.Value, __local_var_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __local_var_0 gopurs_runtime.Value = __local_var_0_loop
_ = __local_var_0
var __local_var_1 gopurs_runtime.Value = __local_var_1_loop
_ = __local_var_1
var __local_var_2 gopurs_runtime.Value = __local_var_2_loop
_ = __local_var_2
var __local_var_3 gopurs_runtime.Value = __local_var_3_loop
_ = __local_var_3
var __t6 *Constructor_Data_Map_Internal_Node
{
if (__local_var_2.Type == 9 && __local_var_2.IntVal == 324739070 && __local_var_2.UnsafePtr == nil) {
__t6 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](__local_var_3)
goto end_branch_6
} else {

}
}
{
if (__local_var_3.Type == 9 && __local_var_3.IntVal == 324739070 && __local_var_3.UnsafePtr == nil) {
__t6 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](__local_var_2)
goto end_branch_6
} else {

}
}
{
if (__local_var_3.Type == 9 && __local_var_3.IntVal == 324739070 && __local_var_3.UnsafePtr != nil) {
// TAST (Let): v_4_0 -> gopurs_runtime.Value
v_4_0 := gopurs_runtime.UncurriedApp3(Get_Data_Map_Internal_unsafeSplit(), __local_var_0, gopurs_runtime.Value{Type: 9, IntVal: int64(uint32((*Constructor_Data_Map_Internal_Node)(__local_var_3.UnsafePtr).V2.IntVal)), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](__local_var_2))})
_ = v_4_0
// TAST (Let): l_prime_5_1 -> gopurs_runtime.Value
l_prime_5_1 := gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeUnionWith(), __local_var_0, __local_var_1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Split)(v_4_0.UnsafePtr).V1)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_3.UnsafePtr).V4)})
_ = l_prime_5_1
// TAST (Let): r_prime_6_2 -> gopurs_runtime.Value
r_prime_6_2 := gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeUnionWith(), __local_var_0, __local_var_1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Split)(v_4_0.UnsafePtr).V2)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_3.UnsafePtr).V5)})
_ = r_prime_6_2
var __t5 *Constructor_Data_Map_Internal_Node
{
var __t_tag_3 *Constructor_Data_Maybe_Just = (*Constructor_Data_Map_Internal_Split)(v_4_0.UnsafePtr).V0
if (__t_tag_3 != nil) {
__t5 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeBalancedNode(), gopurs_runtime.Value{Type: 9, IntVal: int64(uint32((*Constructor_Data_Map_Internal_Node)(__local_var_3.UnsafePtr).V2.IntVal)), UnsafePtr: nil}, gopurs_runtime.Float(gopurs_runtime.Apply2(__local_var_1, gopurs_runtime.Float(((*Constructor_Data_Map_Internal_Split)(v_4_0.UnsafePtr).V0).V0.FloatVal()), gopurs_runtime.Float((*Constructor_Data_Map_Internal_Node)(__local_var_3.UnsafePtr).V3.FloatVal())).FloatVal()), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](l_prime_5_1))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](r_prime_6_2))}))
goto end_branch_5
} else {

}
}
{
var __t_tag_4 *Constructor_Data_Maybe_Just = (*Constructor_Data_Map_Internal_Split)(v_4_0.UnsafePtr).V0
if (__t_tag_4 == nil) {
__t5 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeBalancedNode(), gopurs_runtime.Value{Type: 9, IntVal: int64(uint32((*Constructor_Data_Map_Internal_Node)(__local_var_3.UnsafePtr).V2.IntVal)), UnsafePtr: nil}, gopurs_runtime.Float((*Constructor_Data_Map_Internal_Node)(__local_var_3.UnsafePtr).V3.FloatVal()), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](l_prime_5_1))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](r_prime_6_2))}))
goto end_branch_5
} else {

}
}
{
__t5 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_5:
__t6 = __t5
goto end_branch_6
} else {

}
}
{
__t6 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_6:
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(__t6)}
}


