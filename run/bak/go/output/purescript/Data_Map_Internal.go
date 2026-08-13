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
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_Data_Map_Internal_identity1(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](x_0_box)))}
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
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(&Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]{1, value0.IntVal, value1.IntVal, value2, value3, gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](value4), gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](value5)})}
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
return gopurs_runtime.Value{Type: 9, IntVal: 1343415489, UnsafePtr: unsafe.Pointer(&Constructor_Data_Map_Internal_IterEmit[gopurs_runtime.Value, gopurs_runtime.Value]{1, value0, value1, value2})}
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
return gopurs_runtime.Value{Type: 9, IntVal: 2861335956, UnsafePtr: unsafe.Pointer(&Constructor_Data_Map_Internal_IterNode[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](value0), value1})}
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
return gopurs_runtime.Value{Type: 9, IntVal: 953589075, UnsafePtr: unsafe.Pointer(&Constructor_Data_Map_Internal_IterNext[gopurs_runtime.Value, gopurs_runtime.Value]{1, value0, value1, value2})}
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
return gopurs_runtime.Value{Type: 9, IntVal: 3373277644, UnsafePtr: unsafe.Pointer(&Constructor_Data_Map_Internal_Split[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](value0), gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](value1), gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](value2)})}
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
return gopurs_runtime.Value{Type: 9, IntVal: 2668112006, UnsafePtr: unsafe.Pointer(&Constructor_Data_Map_Internal_SplitLast[gopurs_runtime.Value, gopurs_runtime.Value]{1, value0, value1, gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](value2)})}
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
return Call_Data_Map_Internal_unionWith(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](dictOrd_0_box))
})
	})
	return cache_Data_Map_Internal_unionWith
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
		cache_Data_Map_Internal_semigroupMap1 = gopurs_runtime.Apply(Get_Data_Map_Internal_semigroupMap(), gopurs_runtime.Value{})
	})
	return cache_Data_Map_Internal_semigroupMap1
}

var cache_Data_Map_Internal_pop gopurs_runtime.Value
var once_Data_Map_Internal_pop sync.Once
func Get_Data_Map_Internal_pop() gopurs_runtime.Value {
	once_Data_Map_Internal_pop.Do(func() {
		cache_Data_Map_Internal_pop = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Map_Internal_pop(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](dictOrd_0_box))
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
return Call_Data_Map_Internal_lookupLE(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](dictOrd_0_box), k_1_box)
})
	})
	return cache_Data_Map_Internal_lookupLE
}

var cache_Data_Map_Internal_lookupGE gopurs_runtime.Value
var once_Data_Map_Internal_lookupGE sync.Once
func Get_Data_Map_Internal_lookupGE() gopurs_runtime.Value {
	once_Data_Map_Internal_lookupGE.Do(func() {
		cache_Data_Map_Internal_lookupGE = gopurs_runtime.Func2(func(dictOrd_0_box gopurs_runtime.Value, k_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Map_Internal_lookupGE(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](dictOrd_0_box), k_1_box)
})
	})
	return cache_Data_Map_Internal_lookupGE
}

var cache_Data_Map_Internal_lookup gopurs_runtime.Value
var once_Data_Map_Internal_lookup sync.Once
func Get_Data_Map_Internal_lookup() gopurs_runtime.Value {
	once_Data_Map_Internal_lookup.Do(func() {
		cache_Data_Map_Internal_lookup = gopurs_runtime.Func2(func(dictOrd_0_box gopurs_runtime.Value, k_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Map_Internal_lookup(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](dictOrd_0_box), k_1_box)
})
	})
	return cache_Data_Map_Internal_lookup
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
		cache_Data_Map_Internal_stepUnfoldrUnordered = Call_Data_Map_Internal_stepWith(Get_Data_Map_Internal_iterMapU(), gopurs_runtime.Func3(func(__local_var_0 gopurs_runtime.Value, __local_var_1 gopurs_runtime.Value, __local_var_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, __local_var_0, __local_var_1})}, __local_var_2})}})}
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil))}
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
		cache_Data_Map_Internal_stepUnordered = Call_Data_Map_Internal_stepWith(Get_Data_Map_Internal_iterMapU(), gopurs_runtime.Func3(func(__local_var_0 gopurs_runtime.Value, __local_var_1 gopurs_runtime.Value, __local_var_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 953589075, UnsafePtr: unsafe.Pointer(&Constructor_Data_Map_Internal_IterNext[gopurs_runtime.Value, gopurs_runtime.Value]{1, __local_var_0, __local_var_1, __local_var_2})}
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
var go__go_0_0_8 gopurs_runtime.Value
go__go_0_0_8 = gopurs_runtime.Func(func(iter_1_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var iter_1_loop gopurs_runtime.Value = iter_1_loop_val
var v_2_loop gopurs_runtime.Value = v_2_loop_val
go__go_0_0_8:
for {
if false { continue go__go_0_0_8 }
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
var __t_tag_1 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V5)}
if (__t_tag_1.Type == 9 && __t_tag_1.IntVal == 324739070 && __t_tag_1.UnsafePtr == nil) {
iter_1_loop = gopurs_runtime.Value{Type: 9, IntVal: 1343415489, UnsafePtr: unsafe.Pointer(&Constructor_Data_Map_Internal_IterEmit[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V3, iter_1})}
v_2_loop = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V4)}
continue go__go_0_0_8
__t2 = gopurs_runtime.Value{}
goto end_branch_2
} else {

}
}
{
iter_1_loop = gopurs_runtime.Value{Type: 9, IntVal: 1343415489, UnsafePtr: unsafe.Pointer(&Constructor_Data_Map_Internal_IterEmit[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 2861335956, UnsafePtr: unsafe.Pointer(&Constructor_Data_Map_Internal_IterNode[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V4, iter_1})}})}
v_2_loop = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V5)}
continue go__go_0_0_8
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
return go__go_0_0_8
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
return gopurs_runtime.Value{Type: 9, IntVal: 953589075, UnsafePtr: unsafe.Pointer(&Constructor_Data_Map_Internal_IterNext[gopurs_runtime.Value, gopurs_runtime.Value]{1, __local_var_0, __local_var_1, __local_var_2})}
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
var __t_tag_1 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V5)}
if (__t_tag_1.Type == 9 && __t_tag_1.IntVal == 324739070 && __t_tag_1.UnsafePtr == nil) {
iter_1_loop = gopurs_runtime.Value{Type: 9, IntVal: 1343415489, UnsafePtr: unsafe.Pointer(&Constructor_Data_Map_Internal_IterEmit[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V3, iter_1})}
v_2_loop = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V4)}
continue go__go_0_0_9
__t2 = gopurs_runtime.Value{}
goto end_branch_2
} else {

}
}
{
iter_1_loop = gopurs_runtime.Value{Type: 9, IntVal: 1343415489, UnsafePtr: unsafe.Pointer(&Constructor_Data_Map_Internal_IterEmit[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 2861335956, UnsafePtr: unsafe.Pointer(&Constructor_Data_Map_Internal_IterNode[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V5, iter_1})}})}
v_2_loop = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V4)}
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
return gopurs_runtime.Value{Type: 9, IntVal: 953589075, UnsafePtr: unsafe.Pointer(&Constructor_Data_Map_Internal_IterNext[gopurs_runtime.Value, gopurs_runtime.Value]{1, __local_var_0, __local_var_1, __local_var_2})}
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
		cache_Data_Map_Internal_stepUnfoldr = Call_Data_Map_Internal_stepWith(Get_Data_Map_Internal_iterMapL(), gopurs_runtime.Func3(func(__local_var_0 gopurs_runtime.Value, __local_var_1 gopurs_runtime.Value, __local_var_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, __local_var_0, __local_var_1})}, __local_var_2})}})}
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil))}
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
		cache_Data_Map_Internal_functorMap = gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_1_0_16 gopurs_runtime.Value
_ = go__go_1_0_16
go__go_1_0_16 = gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if (v_2.Type == 9 && v_2.IntVal == 324739070 && v_2.UnsafePtr == nil) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(nil))}
goto end_branch_1
} else {

}
}
{
if (v_2.Type == 9 && v_2.IntVal == 324739070 && v_2.UnsafePtr != nil) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(&Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V0, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V1, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V2, gopurs_runtime.Apply(f_0, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V3), gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(go__go_1_0_16, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V4)})), gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(go__go_1_0_16, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V5)}))})}
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__t1))}
})
return go__go_1_0_16
}))
	})
	return cache_Data_Map_Internal_functorMap
}

var cache_Data_Map_Internal_functorWithIndexMap gopurs_runtime.Value
var once_Data_Map_Internal_functorWithIndexMap sync.Once
func Get_Data_Map_Internal_functorWithIndexMap() gopurs_runtime.Value {
	once_Data_Map_Internal_functorWithIndexMap.Do(func() {
		cache_Data_Map_Internal_functorWithIndexMap = gopurs_runtime.RecordDict2("Functor0", "mapWithIndex", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Map_Internal_functorMap()
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_1_0_17 gopurs_runtime.Value
_ = go__go_1_0_17
go__go_1_0_17 = gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if (v_2.Type == 9 && v_2.IntVal == 324739070 && v_2.UnsafePtr == nil) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(nil))}
goto end_branch_1
} else {

}
}
{
if (v_2.Type == 9 && v_2.IntVal == 324739070 && v_2.UnsafePtr != nil) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(&Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V0, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V1, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V2, gopurs_runtime.Apply2(f_0, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V3), gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(go__go_1_0_17, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V4)})), gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(go__go_1_0_17, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V5)}))})}
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__t1))}
})
return go__go_1_0_17
}))
	})
	return cache_Data_Map_Internal_functorWithIndexMap
}

var cache_Data_Map_Internal_foldableMap gopurs_runtime.Value
var once_Data_Map_Internal_foldableMap sync.Once
func Get_Data_Map_Internal_foldableMap() gopurs_runtime.Value {
	once_Data_Map_Internal_foldableMap.Do(func() {
		cache_Data_Map_Internal_foldableMap = gopurs_runtime.RecordDict3("foldMap", "foldl", "foldr", gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Semigroup0_1_0 -> *Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]
Semigroup0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_1_0
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_3_1_18 gopurs_runtime.Value
_ = go__go_3_1_18
go__go_3_1_18 = gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
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
__t2 = gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_1_0.V0), gopurs_runtime.Apply(go__go_3_1_18, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V4)}), gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_1_0.V0), gopurs_runtime.Apply(f_2, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V3), gopurs_runtime.Apply(go__go_3_1_18, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V5)})))
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
return go__go_3_1_18
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(z_1 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_2_3_19 gopurs_runtime.Value
go__go_2_3_19 = gopurs_runtime.Func(func(__local_var_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(__local_var_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var __local_var_3_loop gopurs_runtime.Value = __local_var_3_loop_val
var __local_var_4_loop gopurs_runtime.Value = __local_var_4_loop_val
go__go_2_3_19:
for {
if false { continue go__go_2_3_19 }
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
__t4 = gopurs_runtime.UncurriedApp2(go__go_2_3_19, gopurs_runtime.Apply2(f_0, gopurs_runtime.UncurriedApp2(go__go_2_3_19, __local_var_3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_4.UnsafePtr).V4)}), (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_4.UnsafePtr).V3), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_4.UnsafePtr).V5)})
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
return gopurs_runtime.UncurriedApp2(go__go_2_3_19, z_1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](m_3))})
})
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(z_1 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_2_5_20 gopurs_runtime.Value
_ = go__go_2_5_20
go__go_2_5_20 = gopurs_runtime.Func2(func(__local_var_3 gopurs_runtime.Value, __local_var_4 gopurs_runtime.Value) gopurs_runtime.Value {
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
__t6 = gopurs_runtime.UncurriedApp2(go__go_2_5_20, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V4)}, gopurs_runtime.Apply2(f_0, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V3, gopurs_runtime.UncurriedApp2(go__go_2_5_20, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V5)}, __local_var_4)))
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
return gopurs_runtime.UncurriedApp2(go__go_2_5_20, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](m_3))}, z_1)
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
return Get_Data_Map_Internal_foldableMap()
}), gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Semigroup0_1_0 -> *Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]
Semigroup0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{}))
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
__t2 = gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_1_0.V0), gopurs_runtime.Apply(go__go_3_1_21, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V4)}), gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_1_0.V0), gopurs_runtime.Apply2(f_2, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V3), gopurs_runtime.Apply(go__go_3_1_21, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V5)})))
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
__t4 = gopurs_runtime.UncurriedApp2(go__go_2_3_22, gopurs_runtime.Apply3(f_0, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_4.UnsafePtr).V2, gopurs_runtime.UncurriedApp2(go__go_2_3_22, __local_var_3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_4.UnsafePtr).V4)}), (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_4.UnsafePtr).V3), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_4.UnsafePtr).V5)})
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
return gopurs_runtime.UncurriedApp2(go__go_2_3_22, z_1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](m_3))})
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
__t6 = gopurs_runtime.UncurriedApp2(go__go_2_5_23, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V4)}, gopurs_runtime.Apply3(f_0, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V3, gopurs_runtime.UncurriedApp2(go__go_2_5_23, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V5)}, __local_var_4)))
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
return gopurs_runtime.UncurriedApp2(go__go_2_5_23, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](m_3))}, z_1)
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
		cache_Data_Map_Internal_keys = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_Map_Internal_foldableWithIndexMap(), "foldrWithIndex"), gopurs_runtime.Func(func(k_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(acc_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, k_0, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](acc_2)})}
})
})
}), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil))})
	})
	return cache_Data_Map_Internal_keys
}

var cache_Data_Map_Internal_traversableMap gopurs_runtime.Value
var once_Data_Map_Internal_traversableMap sync.Once
func Get_Data_Map_Internal_traversableMap() gopurs_runtime.Value {
	once_Data_Map_Internal_traversableMap.Do(func() {
		cache_Data_Map_Internal_traversableMap = gopurs_runtime.RecordDict4("Foldable1", "Functor0", "sequence", "traverse", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Map_Internal_foldableMap()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Map_Internal_functorMap()
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_Map_Internal_traversableMap(), "traverse"), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](dictApplicative_0))}, Get_Data_Map_Internal_identity())
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Apply0_1_0 -> *Constructor_Control_Apply_Apply[gopurs_runtime.Value]
Apply0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}))
_ = Apply0_1_0
// TAST (Let): Functor0_2_1 -> *Constructor_Data_Functor_Functor[gopurs_runtime.Value]
Functor0_2_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_2_1
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_4_2_24 gopurs_runtime.Value
_ = go__go_4_2_24
go__go_4_2_24 = gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
var __t6 gopurs_runtime.Value
{
if (v_5.Type == 9 && v_5.IntVal == 324739070 && v_5.UnsafePtr == nil) {
__t6 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "pure"), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(nil))})
goto end_branch_6
} else {

}
}
{
if (v_5.Type == 9 && v_5.IntVal == 324739070 && v_5.UnsafePtr != nil) {
// TAST (Let): __local_var_6_3 -> gopurs_runtime.Value
var __local_var_6_3 gopurs_runtime.Value = gopurs_runtime.Int((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_5.UnsafePtr).V0)
// TAST (Let): __local_var_7_4 -> gopurs_runtime.Value
__local_var_7_4 := (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_5.UnsafePtr).V2
_ = __local_var_7_4
// TAST (Let): __local_var_8_5 -> gopurs_runtime.Value
var __local_var_8_5 gopurs_runtime.Value = gopurs_runtime.Int((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_5.UnsafePtr).V1)
__t6 = gopurs_runtime.Apply2(gopurs_runtime.Box(Apply0_1_0.V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Apply0_1_0.V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_2_1.V0), gopurs_runtime.Func(func(l_prime_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_prime_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(r_prime_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(&Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]{1, __local_var_6_3.IntVal, __local_var_8_5.IntVal, __local_var_7_4, v_prime_10, gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](l_prime_9), gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](r_prime_11)})}
})
})
}), gopurs_runtime.Apply(go__go_4_2_24, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_5.UnsafePtr).V4)})), gopurs_runtime.Apply(f_3, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_5.UnsafePtr).V3)), gopurs_runtime.Apply(go__go_4_2_24, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_5.UnsafePtr).V5)}))
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
return go__go_4_2_24
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
return Get_Data_Map_Internal_foldableWithIndexMap()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Map_Internal_functorWithIndexMap()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Map_Internal_traversableMap()
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Apply0_1_0 -> *Constructor_Control_Apply_Apply[gopurs_runtime.Value]
Apply0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}))
_ = Apply0_1_0
// TAST (Let): Functor0_2_1 -> *Constructor_Data_Functor_Functor[gopurs_runtime.Value]
Functor0_2_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_2_1
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_4_2_25 gopurs_runtime.Value
_ = go__go_4_2_25
go__go_4_2_25 = gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
var __t6 gopurs_runtime.Value
{
if (v_5.Type == 9 && v_5.IntVal == 324739070 && v_5.UnsafePtr == nil) {
__t6 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "pure"), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(nil))})
goto end_branch_6
} else {

}
}
{
if (v_5.Type == 9 && v_5.IntVal == 324739070 && v_5.UnsafePtr != nil) {
// TAST (Let): __local_var_6_3 -> gopurs_runtime.Value
var __local_var_6_3 gopurs_runtime.Value = gopurs_runtime.Int((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_5.UnsafePtr).V0)
// TAST (Let): __local_var_7_4 -> gopurs_runtime.Value
__local_var_7_4 := (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_5.UnsafePtr).V2
_ = __local_var_7_4
// TAST (Let): __local_var_8_5 -> gopurs_runtime.Value
var __local_var_8_5 gopurs_runtime.Value = gopurs_runtime.Int((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_5.UnsafePtr).V1)
__t6 = gopurs_runtime.Apply2(gopurs_runtime.Box(Apply0_1_0.V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Apply0_1_0.V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_2_1.V0), gopurs_runtime.Func(func(l_prime_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_prime_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(r_prime_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(&Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]{1, __local_var_6_3.IntVal, __local_var_8_5.IntVal, __local_var_7_4, v_prime_10, gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](l_prime_9), gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](r_prime_11)})}
})
})
}), gopurs_runtime.Apply(go__go_4_2_25, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_5.UnsafePtr).V4)})), gopurs_runtime.Apply2(f_3, __local_var_7_4, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_5.UnsafePtr).V3)), gopurs_runtime.Apply(go__go_4_2_25, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_5.UnsafePtr).V5)}))
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
return go__go_4_2_25
})
}))
	})
	return cache_Data_Map_Internal_traversableWithIndexMap
}

var cache_Data_Map_Internal_values gopurs_runtime.Value
var once_Data_Map_Internal_values sync.Once
func Get_Data_Map_Internal_values() gopurs_runtime.Value {
	once_Data_Map_Internal_values.Do(func() {
		cache_Data_Map_Internal_values = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_Map_Internal_foldableMap(), "foldr"), Get_Data_List_Types_Cons(), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil))})
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
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_Data_Map_Internal_findMin(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](v_0_box)))}
})
	})
	return cache_Data_Map_Internal_findMin
}

var cache_Data_Map_Internal_lookupGT gopurs_runtime.Value
var once_Data_Map_Internal_lookupGT sync.Once
func Get_Data_Map_Internal_lookupGT() gopurs_runtime.Value {
	once_Data_Map_Internal_lookupGT.Do(func() {
		cache_Data_Map_Internal_lookupGT = gopurs_runtime.Func2(func(dictOrd_0_box gopurs_runtime.Value, k_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Map_Internal_lookupGT(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](dictOrd_0_box), k_1_box)
})
	})
	return cache_Data_Map_Internal_lookupGT
}

var cache_Data_Map_Internal_findMax gopurs_runtime.Value
var once_Data_Map_Internal_findMax sync.Once
func Get_Data_Map_Internal_findMax() gopurs_runtime.Value {
	once_Data_Map_Internal_findMax.Do(func() {
		cache_Data_Map_Internal_findMax = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_Data_Map_Internal_findMax(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](v_0_box)))}
})
	})
	return cache_Data_Map_Internal_findMax
}

var cache_Data_Map_Internal_lookupLT gopurs_runtime.Value
var once_Data_Map_Internal_lookupLT sync.Once
func Get_Data_Map_Internal_lookupLT() gopurs_runtime.Value {
	once_Data_Map_Internal_lookupLT.Do(func() {
		cache_Data_Map_Internal_lookupLT = gopurs_runtime.Func2(func(dictOrd_0_box gopurs_runtime.Value, k_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Map_Internal_lookupLT(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](dictOrd_0_box), k_1_box)
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

var cache_Data_Map_Internal_delete gopurs_runtime.Value
var once_Data_Map_Internal_delete sync.Once
func Get_Data_Map_Internal_delete() gopurs_runtime.Value {
	once_Data_Map_Internal_delete.Do(func() {
		cache_Data_Map_Internal_delete = gopurs_runtime.Func2(func(dictOrd_0_box gopurs_runtime.Value, k_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Map_Internal_delete(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](dictOrd_0_box), k_1_box)
})
	})
	return cache_Data_Map_Internal_delete
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

var cache_Data_Map_Internal_alter__2325420954 gopurs_runtime.Value
var once_Data_Map_Internal_alter__2325420954 sync.Once
func Get_Data_Map_Internal_alter__2325420954() gopurs_runtime.Value {
	once_Data_Map_Internal_alter__2325420954.Do(func() {
		cache_Data_Map_Internal_alter__2325420954 = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Map_Internal_alter__2325420954(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](dictOrd_0_box))
})
	})
	return cache_Data_Map_Internal_alter__2325420954
}

var cache_Data_Map_Internal_alter__1204655226 gopurs_runtime.Value
var once_Data_Map_Internal_alter__1204655226 sync.Once
func Get_Data_Map_Internal_alter__1204655226() gopurs_runtime.Value {
	once_Data_Map_Internal_alter__1204655226.Do(func() {
		cache_Data_Map_Internal_alter__1204655226 = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Map_Internal_alter__1204655226(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](dictOrd_0_box))
})
	})
	return cache_Data_Map_Internal_alter__1204655226
}

var cache_Data_Map_Internal_empty__2198260019 gopurs_runtime.Value
var once_Data_Map_Internal_empty__2198260019 sync.Once
func Get_Data_Map_Internal_empty__2198260019() gopurs_runtime.Value {
	once_Data_Map_Internal_empty__2198260019.Do(func() {
		cache_Data_Map_Internal_empty__2198260019 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(nil))}
	})
	return cache_Data_Map_Internal_empty__2198260019
}

var cache_Data_Map_Internal_empty__1818220131 gopurs_runtime.Value
var once_Data_Map_Internal_empty__1818220131 sync.Once
func Get_Data_Map_Internal_empty__1818220131() gopurs_runtime.Value {
	once_Data_Map_Internal_empty__1818220131.Do(func() {
		cache_Data_Map_Internal_empty__1818220131 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(nil))}
	})
	return cache_Data_Map_Internal_empty__1818220131
}

var cache_Data_Map_Internal_empty__1299254065 gopurs_runtime.Value
var once_Data_Map_Internal_empty__1299254065 sync.Once
func Get_Data_Map_Internal_empty__1299254065() gopurs_runtime.Value {
	once_Data_Map_Internal_empty__1299254065.Do(func() {
		cache_Data_Map_Internal_empty__1299254065 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(nil))}
	})
	return cache_Data_Map_Internal_empty__1299254065
}

var cache_Data_Map_Internal_empty__1794046843 gopurs_runtime.Value
var once_Data_Map_Internal_empty__1794046843 sync.Once
func Get_Data_Map_Internal_empty__1794046843() gopurs_runtime.Value {
	once_Data_Map_Internal_empty__1794046843.Do(func() {
		cache_Data_Map_Internal_empty__1794046843 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(nil))}
	})
	return cache_Data_Map_Internal_empty__1794046843
}

var cache_Data_Map_Internal_findMax__2266220649 gopurs_runtime.Value
var once_Data_Map_Internal_findMax__2266220649 sync.Once
func Get_Data_Map_Internal_findMax__2266220649() gopurs_runtime.Value {
	once_Data_Map_Internal_findMax__2266220649.Do(func() {
		cache_Data_Map_Internal_findMax__2266220649 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_Data_Map_Internal_findMax__2266220649(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](v_0_box)))}
})
	})
	return cache_Data_Map_Internal_findMax__2266220649
}

var cache_Data_Map_Internal_findMax__528468393 gopurs_runtime.Value
var once_Data_Map_Internal_findMax__528468393 sync.Once
func Get_Data_Map_Internal_findMax__528468393() gopurs_runtime.Value {
	once_Data_Map_Internal_findMax__528468393.Do(func() {
		cache_Data_Map_Internal_findMax__528468393 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_Data_Map_Internal_findMax__528468393(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](v_0_box)))}
})
	})
	return cache_Data_Map_Internal_findMax__528468393
}

var cache_Data_Map_Internal_findMin__2266220649 gopurs_runtime.Value
var once_Data_Map_Internal_findMin__2266220649 sync.Once
func Get_Data_Map_Internal_findMin__2266220649() gopurs_runtime.Value {
	once_Data_Map_Internal_findMin__2266220649.Do(func() {
		cache_Data_Map_Internal_findMin__2266220649 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_Data_Map_Internal_findMin__2266220649(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](v_0_box)))}
})
	})
	return cache_Data_Map_Internal_findMin__2266220649
}

var cache_Data_Map_Internal_findMin__528468393 gopurs_runtime.Value
var once_Data_Map_Internal_findMin__528468393 sync.Once
func Get_Data_Map_Internal_findMin__528468393() gopurs_runtime.Value {
	once_Data_Map_Internal_findMin__528468393.Do(func() {
		cache_Data_Map_Internal_findMin__528468393 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_Data_Map_Internal_findMin__528468393(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](v_0_box)))}
})
	})
	return cache_Data_Map_Internal_findMin__528468393
}

var cache_Data_Map_Internal_foldSubmapBy__3050108409 gopurs_runtime.Value
var once_Data_Map_Internal_foldSubmapBy__3050108409 sync.Once
func Get_Data_Map_Internal_foldSubmapBy__3050108409() gopurs_runtime.Value {
	once_Data_Map_Internal_foldSubmapBy__3050108409.Do(func() {
		cache_Data_Map_Internal_foldSubmapBy__3050108409 = gopurs_runtime.Func6(func(dictOrd_0_box gopurs_runtime.Value, appendFn_1_box gopurs_runtime.Value, memptyValue_2_box gopurs_runtime.Value, kmin_3_box gopurs_runtime.Value, kmax_4_box gopurs_runtime.Value, f_5_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Map_Internal_foldSubmapBy__3050108409(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](dictOrd_0_box), appendFn_1_box, memptyValue_2_box, gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](kmin_3_box), gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](kmax_4_box), f_5_box)
})
	})
	return cache_Data_Map_Internal_foldSubmapBy__3050108409
}

var cache_Data_Map_Internal_foldSubmapBy__3128450809 gopurs_runtime.Value
var once_Data_Map_Internal_foldSubmapBy__3128450809 sync.Once
func Get_Data_Map_Internal_foldSubmapBy__3128450809() gopurs_runtime.Value {
	once_Data_Map_Internal_foldSubmapBy__3128450809.Do(func() {
		cache_Data_Map_Internal_foldSubmapBy__3128450809 = gopurs_runtime.Func6(func(dictOrd_0_box gopurs_runtime.Value, appendFn_1_box gopurs_runtime.Value, memptyValue_2_box gopurs_runtime.Value, kmin_3_box gopurs_runtime.Value, kmax_4_box gopurs_runtime.Value, f_5_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Map_Internal_foldSubmapBy__3128450809(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](dictOrd_0_box), appendFn_1_box, gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](memptyValue_2_box), gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](kmin_3_box), gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](kmax_4_box), f_5_box)
})
	})
	return cache_Data_Map_Internal_foldSubmapBy__3128450809
}

var cache_Data_Map_Internal_foldableMap__767959947 gopurs_runtime.Value
var once_Data_Map_Internal_foldableMap__767959947 sync.Once
func Get_Data_Map_Internal_foldableMap__767959947() gopurs_runtime.Value {
	once_Data_Map_Internal_foldableMap__767959947.Do(func() {
		cache_Data_Map_Internal_foldableMap__767959947 = gopurs_runtime.RecordDict3("foldMap", "foldl", "foldr", gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Semigroup0_1_0 -> *Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]
Semigroup0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_1_0
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_3_1_38 gopurs_runtime.Value
_ = go__go_3_1_38
go__go_3_1_38 = gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
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
__t2 = gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_1_0.V0), gopurs_runtime.Apply(go__go_3_1_38, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V4)}), gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_1_0.V0), gopurs_runtime.Apply(f_2, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V3), gopurs_runtime.Apply(go__go_3_1_38, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V5)})))
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
return go__go_3_1_38
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(z_1 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_2_3_39 gopurs_runtime.Value
go__go_2_3_39 = gopurs_runtime.Func(func(__local_var_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(__local_var_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var __local_var_3_loop gopurs_runtime.Value = __local_var_3_loop_val
var __local_var_4_loop gopurs_runtime.Value = __local_var_4_loop_val
go__go_2_3_39:
for {
if false { continue go__go_2_3_39 }
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
__t4 = gopurs_runtime.UncurriedApp2(go__go_2_3_39, gopurs_runtime.Apply2(f_0, gopurs_runtime.UncurriedApp2(go__go_2_3_39, __local_var_3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_4.UnsafePtr).V4)}), (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_4.UnsafePtr).V3), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_4.UnsafePtr).V5)})
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
return gopurs_runtime.UncurriedApp2(go__go_2_3_39, z_1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](m_3))})
})
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(z_1 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_2_5_40 gopurs_runtime.Value
_ = go__go_2_5_40
go__go_2_5_40 = gopurs_runtime.Func2(func(__local_var_3 gopurs_runtime.Value, __local_var_4 gopurs_runtime.Value) gopurs_runtime.Value {
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
__t6 = gopurs_runtime.UncurriedApp2(go__go_2_5_40, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V4)}, gopurs_runtime.Apply2(f_0, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V3, gopurs_runtime.UncurriedApp2(go__go_2_5_40, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V5)}, __local_var_4)))
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
return gopurs_runtime.UncurriedApp2(go__go_2_5_40, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](m_3))}, z_1)
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
// TAST (Let): Semigroup0_1_0 -> *Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]
Semigroup0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_1_0
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_3_1_41 gopurs_runtime.Value
_ = go__go_3_1_41
go__go_3_1_41 = gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
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
__t2 = gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_1_0.V0), gopurs_runtime.Apply(go__go_3_1_41, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V4)}), gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_1_0.V0), gopurs_runtime.Apply(f_2, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V3), gopurs_runtime.Apply(go__go_3_1_41, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V5)})))
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
return go__go_3_1_41
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(z_1 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_2_3_42 gopurs_runtime.Value
go__go_2_3_42 = gopurs_runtime.Func(func(__local_var_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(__local_var_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var __local_var_3_loop gopurs_runtime.Value = __local_var_3_loop_val
var __local_var_4_loop gopurs_runtime.Value = __local_var_4_loop_val
go__go_2_3_42:
for {
if false { continue go__go_2_3_42 }
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
__t4 = gopurs_runtime.UncurriedApp2(go__go_2_3_42, gopurs_runtime.Apply2(f_0, gopurs_runtime.UncurriedApp2(go__go_2_3_42, __local_var_3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_4.UnsafePtr).V4)}), (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_4.UnsafePtr).V3), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_4.UnsafePtr).V5)})
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
return gopurs_runtime.UncurriedApp2(go__go_2_3_42, z_1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](m_3))})
})
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(z_1 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_2_5_43 gopurs_runtime.Value
_ = go__go_2_5_43
go__go_2_5_43 = gopurs_runtime.Func2(func(__local_var_3 gopurs_runtime.Value, __local_var_4 gopurs_runtime.Value) gopurs_runtime.Value {
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
__t6 = gopurs_runtime.UncurriedApp2(go__go_2_5_43, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V4)}, gopurs_runtime.Apply2(f_0, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V3, gopurs_runtime.UncurriedApp2(go__go_2_5_43, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V5)}, __local_var_4)))
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
return gopurs_runtime.UncurriedApp2(go__go_2_5_43, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](m_3))}, z_1)
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
return Get_Data_Map_Internal_foldableMap()
}), gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Semigroup0_1_0 -> *Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]
Semigroup0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_1_0
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_3_1_44 gopurs_runtime.Value
_ = go__go_3_1_44
go__go_3_1_44 = gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
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
__t2 = gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_1_0.V0), gopurs_runtime.Apply(go__go_3_1_44, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V4)}), gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_1_0.V0), gopurs_runtime.Apply2(f_2, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V3), gopurs_runtime.Apply(go__go_3_1_44, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V5)})))
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
return go__go_3_1_44
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(z_1 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_2_3_45 gopurs_runtime.Value
go__go_2_3_45 = gopurs_runtime.Func(func(__local_var_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(__local_var_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var __local_var_3_loop gopurs_runtime.Value = __local_var_3_loop_val
var __local_var_4_loop gopurs_runtime.Value = __local_var_4_loop_val
go__go_2_3_45:
for {
if false { continue go__go_2_3_45 }
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
__t4 = gopurs_runtime.UncurriedApp2(go__go_2_3_45, gopurs_runtime.Apply3(f_0, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_4.UnsafePtr).V2, gopurs_runtime.UncurriedApp2(go__go_2_3_45, __local_var_3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_4.UnsafePtr).V4)}), (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_4.UnsafePtr).V3), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_4.UnsafePtr).V5)})
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
return gopurs_runtime.UncurriedApp2(go__go_2_3_45, z_1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](m_3))})
})
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(z_1 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_2_5_46 gopurs_runtime.Value
_ = go__go_2_5_46
go__go_2_5_46 = gopurs_runtime.Func2(func(__local_var_3 gopurs_runtime.Value, __local_var_4 gopurs_runtime.Value) gopurs_runtime.Value {
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
__t6 = gopurs_runtime.UncurriedApp2(go__go_2_5_46, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V4)}, gopurs_runtime.Apply3(f_0, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V3, gopurs_runtime.UncurriedApp2(go__go_2_5_46, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V5)}, __local_var_4)))
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
return gopurs_runtime.UncurriedApp2(go__go_2_5_46, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](m_3))}, z_1)
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
return Get_Data_Map_Internal_foldableMap()
}), gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Semigroup0_1_0 -> *Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]
Semigroup0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_1_0
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_3_1_47 gopurs_runtime.Value
_ = go__go_3_1_47
go__go_3_1_47 = gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
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
__t2 = gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_1_0.V0), gopurs_runtime.Apply(go__go_3_1_47, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V4)}), gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_1_0.V0), gopurs_runtime.Apply2(f_2, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V3), gopurs_runtime.Apply(go__go_3_1_47, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V5)})))
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
return go__go_3_1_47
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(z_1 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_2_3_48 gopurs_runtime.Value
go__go_2_3_48 = gopurs_runtime.Func(func(__local_var_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(__local_var_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var __local_var_3_loop gopurs_runtime.Value = __local_var_3_loop_val
var __local_var_4_loop gopurs_runtime.Value = __local_var_4_loop_val
go__go_2_3_48:
for {
if false { continue go__go_2_3_48 }
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
__t4 = gopurs_runtime.UncurriedApp2(go__go_2_3_48, gopurs_runtime.Apply3(f_0, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_4.UnsafePtr).V2, gopurs_runtime.UncurriedApp2(go__go_2_3_48, __local_var_3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_4.UnsafePtr).V4)}), (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_4.UnsafePtr).V3), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_4.UnsafePtr).V5)})
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
return gopurs_runtime.UncurriedApp2(go__go_2_3_48, z_1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](m_3))})
})
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(z_1 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_2_5_49 gopurs_runtime.Value
_ = go__go_2_5_49
go__go_2_5_49 = gopurs_runtime.Func2(func(__local_var_3 gopurs_runtime.Value, __local_var_4 gopurs_runtime.Value) gopurs_runtime.Value {
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
__t6 = gopurs_runtime.UncurriedApp2(go__go_2_5_49, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V4)}, gopurs_runtime.Apply3(f_0, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V3, gopurs_runtime.UncurriedApp2(go__go_2_5_49, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V5)}, __local_var_4)))
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
return gopurs_runtime.UncurriedApp2(go__go_2_5_49, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](m_3))}, z_1)
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
return Get_Data_Map_Internal_foldableMap()
}), gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Semigroup0_1_0 -> *Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]
Semigroup0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_1_0
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_3_1_50 gopurs_runtime.Value
_ = go__go_3_1_50
go__go_3_1_50 = gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
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
__t2 = gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_1_0.V0), gopurs_runtime.Apply(go__go_3_1_50, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V4)}), gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_1_0.V0), gopurs_runtime.Apply2(f_2, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V3), gopurs_runtime.Apply(go__go_3_1_50, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V5)})))
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
return go__go_3_1_50
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(z_1 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_2_3_51 gopurs_runtime.Value
go__go_2_3_51 = gopurs_runtime.Func(func(__local_var_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(__local_var_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var __local_var_3_loop gopurs_runtime.Value = __local_var_3_loop_val
var __local_var_4_loop gopurs_runtime.Value = __local_var_4_loop_val
go__go_2_3_51:
for {
if false { continue go__go_2_3_51 }
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
__t4 = gopurs_runtime.UncurriedApp2(go__go_2_3_51, gopurs_runtime.Apply3(f_0, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_4.UnsafePtr).V2, gopurs_runtime.UncurriedApp2(go__go_2_3_51, __local_var_3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_4.UnsafePtr).V4)}), (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_4.UnsafePtr).V3), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_4.UnsafePtr).V5)})
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
return gopurs_runtime.UncurriedApp2(go__go_2_3_51, z_1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](m_3))})
})
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(z_1 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_2_5_52 gopurs_runtime.Value
_ = go__go_2_5_52
go__go_2_5_52 = gopurs_runtime.Func2(func(__local_var_3 gopurs_runtime.Value, __local_var_4 gopurs_runtime.Value) gopurs_runtime.Value {
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
__t6 = gopurs_runtime.UncurriedApp2(go__go_2_5_52, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V4)}, gopurs_runtime.Apply3(f_0, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V3, gopurs_runtime.UncurriedApp2(go__go_2_5_52, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V5)}, __local_var_4)))
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
return gopurs_runtime.UncurriedApp2(go__go_2_5_52, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](m_3))}, z_1)
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
var go__go_1_0_53 gopurs_runtime.Value
_ = go__go_1_0_53
go__go_1_0_53 = gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if (v_2.Type == 9 && v_2.IntVal == 324739070 && v_2.UnsafePtr == nil) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(nil))}
goto end_branch_1
} else {

}
}
{
if (v_2.Type == 9 && v_2.IntVal == 324739070 && v_2.UnsafePtr != nil) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(&Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V0, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V1, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V2, gopurs_runtime.Apply(f_0, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V3), gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(go__go_1_0_53, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V4)})), gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(go__go_1_0_53, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V5)}))})}
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__t1))}
})
return go__go_1_0_53
}))
	})
	return cache_Data_Map_Internal_functorMap__2501170929
}

var cache_Data_Map_Internal_functorWithIndexMap__3138419015 gopurs_runtime.Value
var once_Data_Map_Internal_functorWithIndexMap__3138419015 sync.Once
func Get_Data_Map_Internal_functorWithIndexMap__3138419015() gopurs_runtime.Value {
	once_Data_Map_Internal_functorWithIndexMap__3138419015.Do(func() {
		cache_Data_Map_Internal_functorWithIndexMap__3138419015 = gopurs_runtime.RecordDict2("Functor0", "mapWithIndex", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Map_Internal_functorMap()
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_1_0_54 gopurs_runtime.Value
_ = go__go_1_0_54
go__go_1_0_54 = gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if (v_2.Type == 9 && v_2.IntVal == 324739070 && v_2.UnsafePtr == nil) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(nil))}
goto end_branch_1
} else {

}
}
{
if (v_2.Type == 9 && v_2.IntVal == 324739070 && v_2.UnsafePtr != nil) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(&Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V0, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V1, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V2, gopurs_runtime.Apply2(f_0, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V3), gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(go__go_1_0_54, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V4)})), gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(go__go_1_0_54, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V5)}))})}
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__t1))}
})
return go__go_1_0_54
}))
	})
	return cache_Data_Map_Internal_functorWithIndexMap__3138419015
}

var cache_Data_Map_Internal_insert__3204212386 gopurs_runtime.Value
var once_Data_Map_Internal_insert__3204212386 sync.Once
func Get_Data_Map_Internal_insert__3204212386() gopurs_runtime.Value {
	once_Data_Map_Internal_insert__3204212386.Do(func() {
		cache_Data_Map_Internal_insert__3204212386 = gopurs_runtime.Func3(func(dictOrd_0_box gopurs_runtime.Value, k_1_box gopurs_runtime.Value, v_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Map_Internal_insert__3204212386(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](dictOrd_0_box), k_1_box, v_2_box)
})
	})
	return cache_Data_Map_Internal_insert__3204212386
}

var cache_Data_Map_Internal_insert__4289641298 gopurs_runtime.Value
var once_Data_Map_Internal_insert__4289641298 sync.Once
func Get_Data_Map_Internal_insert__4289641298() gopurs_runtime.Value {
	once_Data_Map_Internal_insert__4289641298.Do(func() {
		cache_Data_Map_Internal_insert__4289641298 = gopurs_runtime.Func3(func(dictOrd_0_box gopurs_runtime.Value, k_1_box gopurs_runtime.Value, v_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Map_Internal_insert__4289641298(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](dictOrd_0_box), k_1_box, v_2_box)
})
	})
	return cache_Data_Map_Internal_insert__4289641298
}

var cache_Data_Map_Internal_insert__2073142786 gopurs_runtime.Value
var once_Data_Map_Internal_insert__2073142786 sync.Once
func Get_Data_Map_Internal_insert__2073142786() gopurs_runtime.Value {
	once_Data_Map_Internal_insert__2073142786.Do(func() {
		cache_Data_Map_Internal_insert__2073142786 = gopurs_runtime.Func3(func(dictOrd_0_box gopurs_runtime.Value, k_1_box gopurs_runtime.Value, v_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Map_Internal_insert__2073142786(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]]](dictOrd_0_box), gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](k_1_box), v_2_box)
})
	})
	return cache_Data_Map_Internal_insert__2073142786
}

var cache_Data_Map_Internal_insertWith__118979962 gopurs_runtime.Value
var once_Data_Map_Internal_insertWith__118979962 sync.Once
func Get_Data_Map_Internal_insertWith__118979962() gopurs_runtime.Value {
	once_Data_Map_Internal_insertWith__118979962.Do(func() {
		cache_Data_Map_Internal_insertWith__118979962 = gopurs_runtime.Func4(func(dictOrd_0_box gopurs_runtime.Value, app_1_box gopurs_runtime.Value, k_2_box gopurs_runtime.Value, v_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Map_Internal_insertWith__118979962(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](dictOrd_0_box), app_1_box, k_2_box, v_3_box)
})
	})
	return cache_Data_Map_Internal_insertWith__118979962
}

var cache_Data_Map_Internal_intersectionWith__3717755541 gopurs_runtime.Value
var once_Data_Map_Internal_intersectionWith__3717755541 sync.Once
func Get_Data_Map_Internal_intersectionWith__3717755541() gopurs_runtime.Value {
	once_Data_Map_Internal_intersectionWith__3717755541.Do(func() {
		cache_Data_Map_Internal_intersectionWith__3717755541 = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Map_Internal_intersectionWith__3717755541(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](dictOrd_0_box))
})
	})
	return cache_Data_Map_Internal_intersectionWith__3717755541
}

var cache_Data_Map_Internal_intersectionWith__4144106805 gopurs_runtime.Value
var once_Data_Map_Internal_intersectionWith__4144106805 sync.Once
func Get_Data_Map_Internal_intersectionWith__4144106805() gopurs_runtime.Value {
	once_Data_Map_Internal_intersectionWith__4144106805.Do(func() {
		cache_Data_Map_Internal_intersectionWith__4144106805 = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Map_Internal_intersectionWith__4144106805(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](dictOrd_0_box))
})
	})
	return cache_Data_Map_Internal_intersectionWith__4144106805
}

var cache_Data_Map_Internal_isEmpty__1620059593 gopurs_runtime.Value
var once_Data_Map_Internal_isEmpty__1620059593 sync.Once
func Get_Data_Map_Internal_isEmpty__1620059593() gopurs_runtime.Value {
	once_Data_Map_Internal_isEmpty__1620059593.Do(func() {
		cache_Data_Map_Internal_isEmpty__1620059593 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_Data_Map_Internal_isEmpty__1620059593(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](v_0_box)))
})
	})
	return cache_Data_Map_Internal_isEmpty__1620059593
}

var cache_Data_Map_Internal_iterMapL__3394814354 gopurs_runtime.Value
var once_Data_Map_Internal_iterMapL__3394814354 sync.Once
func Get_Data_Map_Internal_iterMapL__3394814354() gopurs_runtime.Value {
	once_Data_Map_Internal_iterMapL__3394814354.Do(func() {
		cache_Data_Map_Internal_iterMapL__3394814354 = func() gopurs_runtime.Value {
var go__go_0_0_59 gopurs_runtime.Value
go__go_0_0_59 = gopurs_runtime.Func(func(iter_1_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var iter_1_loop gopurs_runtime.Value = iter_1_loop_val
var v_2_loop gopurs_runtime.Value = v_2_loop_val
go__go_0_0_59:
for {
if false { continue go__go_0_0_59 }
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
var __t_tag_1 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V5)}
if (__t_tag_1.Type == 9 && __t_tag_1.IntVal == 324739070 && __t_tag_1.UnsafePtr == nil) {
iter_1_loop = gopurs_runtime.Value{Type: 9, IntVal: 1343415489, UnsafePtr: unsafe.Pointer(&Constructor_Data_Map_Internal_IterEmit[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V3, iter_1})}
v_2_loop = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V4)}
continue go__go_0_0_59
__t2 = gopurs_runtime.Value{}
goto end_branch_2
} else {

}
}
{
iter_1_loop = gopurs_runtime.Value{Type: 9, IntVal: 1343415489, UnsafePtr: unsafe.Pointer(&Constructor_Data_Map_Internal_IterEmit[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 2861335956, UnsafePtr: unsafe.Pointer(&Constructor_Data_Map_Internal_IterNode[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V5, iter_1})}})}
v_2_loop = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V4)}
continue go__go_0_0_59
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
return go__go_0_0_59
}()
	})
	return cache_Data_Map_Internal_iterMapL__3394814354
}

var cache_Data_Map_Internal_iterMapL__878452066 gopurs_runtime.Value
var once_Data_Map_Internal_iterMapL__878452066 sync.Once
func Get_Data_Map_Internal_iterMapL__878452066() gopurs_runtime.Value {
	once_Data_Map_Internal_iterMapL__878452066.Do(func() {
		cache_Data_Map_Internal_iterMapL__878452066 = func() gopurs_runtime.Value {
var go__go_0_0_60 gopurs_runtime.Value
go__go_0_0_60 = gopurs_runtime.Func(func(iter_1_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var iter_1_loop gopurs_runtime.Value = iter_1_loop_val
var v_2_loop gopurs_runtime.Value = v_2_loop_val
go__go_0_0_60:
for {
if false { continue go__go_0_0_60 }
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
var __t_tag_1 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V5)}
if (__t_tag_1.Type == 9 && __t_tag_1.IntVal == 324739070 && __t_tag_1.UnsafePtr == nil) {
iter_1_loop = gopurs_runtime.Value{Type: 9, IntVal: 1343415489, UnsafePtr: unsafe.Pointer(&Constructor_Data_Map_Internal_IterEmit[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V3, iter_1})}
v_2_loop = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V4)}
continue go__go_0_0_60
__t2 = gopurs_runtime.Value{}
goto end_branch_2
} else {

}
}
{
iter_1_loop = gopurs_runtime.Value{Type: 9, IntVal: 1343415489, UnsafePtr: unsafe.Pointer(&Constructor_Data_Map_Internal_IterEmit[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 2861335956, UnsafePtr: unsafe.Pointer(&Constructor_Data_Map_Internal_IterNode[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V5, iter_1})}})}
v_2_loop = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V4)}
continue go__go_0_0_60
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
return go__go_0_0_60
}()
	})
	return cache_Data_Map_Internal_iterMapL__878452066
}

var cache_Data_Map_Internal_iterMapL__1101342704 gopurs_runtime.Value
var once_Data_Map_Internal_iterMapL__1101342704 sync.Once
func Get_Data_Map_Internal_iterMapL__1101342704() gopurs_runtime.Value {
	once_Data_Map_Internal_iterMapL__1101342704.Do(func() {
		cache_Data_Map_Internal_iterMapL__1101342704 = func() gopurs_runtime.Value {
var go__go_0_0_61 gopurs_runtime.Value
go__go_0_0_61 = gopurs_runtime.Func(func(iter_1_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var iter_1_loop gopurs_runtime.Value = iter_1_loop_val
var v_2_loop gopurs_runtime.Value = v_2_loop_val
go__go_0_0_61:
for {
if false { continue go__go_0_0_61 }
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
var __t_tag_1 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V5)}
if (__t_tag_1.Type == 9 && __t_tag_1.IntVal == 324739070 && __t_tag_1.UnsafePtr == nil) {
iter_1_loop = gopurs_runtime.Value{Type: 9, IntVal: 1343415489, UnsafePtr: unsafe.Pointer(&Constructor_Data_Map_Internal_IterEmit[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: int64(uint32((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V2.IntVal)), UnsafePtr: nil}, gopurs_runtime.Float((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V3.FloatVal()), iter_1})}
v_2_loop = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V4)}
continue go__go_0_0_61
__t2 = gopurs_runtime.Value{}
goto end_branch_2
} else {

}
}
{
iter_1_loop = gopurs_runtime.Value{Type: 9, IntVal: 1343415489, UnsafePtr: unsafe.Pointer(&Constructor_Data_Map_Internal_IterEmit[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: int64(uint32((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V2.IntVal)), UnsafePtr: nil}, gopurs_runtime.Float((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V3.FloatVal()), gopurs_runtime.Value{Type: 9, IntVal: 2861335956, UnsafePtr: unsafe.Pointer(&Constructor_Data_Map_Internal_IterNode[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V5, iter_1})}})}
v_2_loop = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V4)}
continue go__go_0_0_61
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
return go__go_0_0_61
}()
	})
	return cache_Data_Map_Internal_iterMapL__1101342704
}

var cache_Data_Map_Internal_iterMapR__878452066 gopurs_runtime.Value
var once_Data_Map_Internal_iterMapR__878452066 sync.Once
func Get_Data_Map_Internal_iterMapR__878452066() gopurs_runtime.Value {
	once_Data_Map_Internal_iterMapR__878452066.Do(func() {
		cache_Data_Map_Internal_iterMapR__878452066 = func() gopurs_runtime.Value {
var go__go_0_0_62 gopurs_runtime.Value
go__go_0_0_62 = gopurs_runtime.Func(func(iter_1_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var iter_1_loop gopurs_runtime.Value = iter_1_loop_val
var v_2_loop gopurs_runtime.Value = v_2_loop_val
go__go_0_0_62:
for {
if false { continue go__go_0_0_62 }
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
var __t_tag_1 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V5)}
if (__t_tag_1.Type == 9 && __t_tag_1.IntVal == 324739070 && __t_tag_1.UnsafePtr == nil) {
iter_1_loop = gopurs_runtime.Value{Type: 9, IntVal: 1343415489, UnsafePtr: unsafe.Pointer(&Constructor_Data_Map_Internal_IterEmit[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V3, iter_1})}
v_2_loop = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V4)}
continue go__go_0_0_62
__t2 = gopurs_runtime.Value{}
goto end_branch_2
} else {

}
}
{
iter_1_loop = gopurs_runtime.Value{Type: 9, IntVal: 1343415489, UnsafePtr: unsafe.Pointer(&Constructor_Data_Map_Internal_IterEmit[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 2861335956, UnsafePtr: unsafe.Pointer(&Constructor_Data_Map_Internal_IterNode[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V4, iter_1})}})}
v_2_loop = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V5)}
continue go__go_0_0_62
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
return go__go_0_0_62
}()
	})
	return cache_Data_Map_Internal_iterMapR__878452066
}

var cache_Data_Map_Internal_iterMapU__878452066 gopurs_runtime.Value
var once_Data_Map_Internal_iterMapU__878452066 sync.Once
func Get_Data_Map_Internal_iterMapU__878452066() gopurs_runtime.Value {
	once_Data_Map_Internal_iterMapU__878452066.Do(func() {
		cache_Data_Map_Internal_iterMapU__878452066 = gopurs_runtime.Func2(func(iter_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Map_Internal_iterMapU__878452066(iter_0_box, gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](v_1_box))
})
	})
	return cache_Data_Map_Internal_iterMapU__878452066
}

var cache_Data_Map_Internal_keys__3504999702 gopurs_runtime.Value
var once_Data_Map_Internal_keys__3504999702 sync.Once
func Get_Data_Map_Internal_keys__3504999702() gopurs_runtime.Value {
	once_Data_Map_Internal_keys__3504999702.Do(func() {
		cache_Data_Map_Internal_keys__3504999702 = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_Map_Internal_foldableWithIndexMap(), "foldrWithIndex"), gopurs_runtime.Func(func(k_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(acc_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, k_0, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](acc_2)})}
})
})
}), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil))})
	})
	return cache_Data_Map_Internal_keys__3504999702
}

var cache_Data_Map_Internal_keys__2406038214 gopurs_runtime.Value
var once_Data_Map_Internal_keys__2406038214 sync.Once
func Get_Data_Map_Internal_keys__2406038214() gopurs_runtime.Value {
	once_Data_Map_Internal_keys__2406038214.Do(func() {
		cache_Data_Map_Internal_keys__2406038214 = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_Map_Internal_foldableWithIndexMap(), "foldrWithIndex"), gopurs_runtime.Func(func(k_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(acc_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, k_0, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](acc_2)})}
})
})
}), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil))})
	})
	return cache_Data_Map_Internal_keys__2406038214
}

var cache_Data_Map_Internal_keys__2813649686 gopurs_runtime.Value
var once_Data_Map_Internal_keys__2813649686 sync.Once
func Get_Data_Map_Internal_keys__2813649686() gopurs_runtime.Value {
	once_Data_Map_Internal_keys__2813649686.Do(func() {
		cache_Data_Map_Internal_keys__2813649686 = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_Map_Internal_foldableWithIndexMap(), "foldrWithIndex"), gopurs_runtime.Func(func(k_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(acc_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](k_0))}, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]]](acc_2))})})}
})
})
}), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil))})
	})
	return cache_Data_Map_Internal_keys__2813649686
}

var cache_Data_Map_Internal_lookup__3378638282 gopurs_runtime.Value
var once_Data_Map_Internal_lookup__3378638282 sync.Once
func Get_Data_Map_Internal_lookup__3378638282() gopurs_runtime.Value {
	once_Data_Map_Internal_lookup__3378638282.Do(func() {
		cache_Data_Map_Internal_lookup__3378638282 = gopurs_runtime.Func2(func(dictOrd_0_box gopurs_runtime.Value, k_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Map_Internal_lookup__3378638282(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](dictOrd_0_box), k_1_box)
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
return Call_Data_Map_Internal_mapMaybe__3426301240(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](dictOrd_0_box), x_1_box)
})
	})
	return cache_Data_Map_Internal_mapMaybe__3426301240
}

var cache_Data_Map_Internal_mapMaybe__1970555288 gopurs_runtime.Value
var once_Data_Map_Internal_mapMaybe__1970555288 sync.Once
func Get_Data_Map_Internal_mapMaybe__1970555288() gopurs_runtime.Value {
	once_Data_Map_Internal_mapMaybe__1970555288.Do(func() {
		cache_Data_Map_Internal_mapMaybe__1970555288 = gopurs_runtime.Func2(func(dictOrd_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Map_Internal_mapMaybe__1970555288(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](dictOrd_0_box), x_1_box)
})
	})
	return cache_Data_Map_Internal_mapMaybe__1970555288
}

var cache_Data_Map_Internal_mapMaybeWithKey__817660689 gopurs_runtime.Value
var once_Data_Map_Internal_mapMaybeWithKey__817660689 sync.Once
func Get_Data_Map_Internal_mapMaybeWithKey__817660689() gopurs_runtime.Value {
	once_Data_Map_Internal_mapMaybeWithKey__817660689.Do(func() {
		cache_Data_Map_Internal_mapMaybeWithKey__817660689 = gopurs_runtime.Func2(func(dictOrd_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Map_Internal_mapMaybeWithKey__817660689(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](dictOrd_0_box), f_1_box)
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
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(Call_Data_Map_Internal_singleton__1300483034(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](k_0_box), v_1_box))}
})
	})
	return cache_Data_Map_Internal_singleton__1300483034
}

var cache_Data_Map_Internal_size__909390430 gopurs_runtime.Value
var once_Data_Map_Internal_size__909390430 sync.Once
func Get_Data_Map_Internal_size__909390430() gopurs_runtime.Value {
	once_Data_Map_Internal_size__909390430.Do(func() {
		cache_Data_Map_Internal_size__909390430 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int(Call_Data_Map_Internal_size__909390430(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](v_0_box)))
})
	})
	return cache_Data_Map_Internal_size__909390430
}

var cache_Data_Map_Internal_size__1374028086 gopurs_runtime.Value
var once_Data_Map_Internal_size__1374028086 sync.Once
func Get_Data_Map_Internal_size__1374028086() gopurs_runtime.Value {
	once_Data_Map_Internal_size__1374028086.Do(func() {
		cache_Data_Map_Internal_size__1374028086 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int(Call_Data_Map_Internal_size__1374028086(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](v_0_box)))
})
	})
	return cache_Data_Map_Internal_size__1374028086
}

var cache_Data_Map_Internal_size__2382154916 gopurs_runtime.Value
var once_Data_Map_Internal_size__2382154916 sync.Once
func Get_Data_Map_Internal_size__2382154916() gopurs_runtime.Value {
	once_Data_Map_Internal_size__2382154916.Do(func() {
		cache_Data_Map_Internal_size__2382154916 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int(Call_Data_Map_Internal_size__2382154916(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[uint32, float64]](v_0_box)))
})
	})
	return cache_Data_Map_Internal_size__2382154916
}

var cache_Data_Map_Internal_stepAsc__2098920977 gopurs_runtime.Value
var once_Data_Map_Internal_stepAsc__2098920977 sync.Once
func Get_Data_Map_Internal_stepAsc__2098920977() gopurs_runtime.Value {
	once_Data_Map_Internal_stepAsc__2098920977.Do(func() {
		cache_Data_Map_Internal_stepAsc__2098920977 = Call_Data_Map_Internal_stepWith(Get_Data_Map_Internal_iterMapL(), gopurs_runtime.Func3(func(__local_var_0 gopurs_runtime.Value, __local_var_1 gopurs_runtime.Value, __local_var_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 953589075, UnsafePtr: unsafe.Pointer(&Constructor_Data_Map_Internal_IterNext[gopurs_runtime.Value, gopurs_runtime.Value]{1, __local_var_0, __local_var_1, __local_var_2})}
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 953589075, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_IterNext[gopurs_runtime.Value, gopurs_runtime.Value])(nil))}
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
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, __local_var_0, __local_var_1})}, __local_var_2})}})}
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil))}
}))
	})
	return cache_Data_Map_Internal_stepUnfoldr__966001626
}

var cache_Data_Map_Internal_stepUnfoldr__575593864 gopurs_runtime.Value
var once_Data_Map_Internal_stepUnfoldr__575593864 sync.Once
func Get_Data_Map_Internal_stepUnfoldr__575593864() gopurs_runtime.Value {
	once_Data_Map_Internal_stepUnfoldr__575593864.Do(func() {
		cache_Data_Map_Internal_stepUnfoldr__575593864 = Call_Data_Map_Internal_stepWith(Get_Data_Map_Internal_iterMapL(), gopurs_runtime.Func3(func(__local_var_0 gopurs_runtime.Value, __local_var_1 gopurs_runtime.Value, __local_var_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(__local_var_0.IntVal)), UnsafePtr: nil}, gopurs_runtime.Float(__local_var_1.FloatVal())})}, __local_var_2})}})}
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil))}
}))
	})
	return cache_Data_Map_Internal_stepUnfoldr__575593864
}

var cache_Data_Map_Internal_stepUnfoldrUnordered__966001626 gopurs_runtime.Value
var once_Data_Map_Internal_stepUnfoldrUnordered__966001626 sync.Once
func Get_Data_Map_Internal_stepUnfoldrUnordered__966001626() gopurs_runtime.Value {
	once_Data_Map_Internal_stepUnfoldrUnordered__966001626.Do(func() {
		cache_Data_Map_Internal_stepUnfoldrUnordered__966001626 = Call_Data_Map_Internal_stepWith(Get_Data_Map_Internal_iterMapU(), gopurs_runtime.Func3(func(__local_var_0 gopurs_runtime.Value, __local_var_1 gopurs_runtime.Value, __local_var_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, __local_var_0, __local_var_1})}, __local_var_2})}})}
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil))}
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
return Call_Data_Map_Internal_toMapIter__1799172593(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](a_0_box))
})
	})
	return cache_Data_Map_Internal_toMapIter__1799172593
}

var cache_Data_Map_Internal_toMapIter__2014410513 gopurs_runtime.Value
var once_Data_Map_Internal_toMapIter__2014410513 sync.Once
func Get_Data_Map_Internal_toMapIter__2014410513() gopurs_runtime.Value {
	once_Data_Map_Internal_toMapIter__2014410513.Do(func() {
		cache_Data_Map_Internal_toMapIter__2014410513 = gopurs_runtime.Func(func(a_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Map_Internal_toMapIter__2014410513(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](a_0_box))
})
	})
	return cache_Data_Map_Internal_toMapIter__2014410513
}

var cache_Data_Map_Internal_toMapIter__772765521 gopurs_runtime.Value
var once_Data_Map_Internal_toMapIter__772765521 sync.Once
func Get_Data_Map_Internal_toMapIter__772765521() gopurs_runtime.Value {
	once_Data_Map_Internal_toMapIter__772765521.Do(func() {
		cache_Data_Map_Internal_toMapIter__772765521 = gopurs_runtime.Func(func(a_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Map_Internal_toMapIter__772765521(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[uint32, float64]](a_0_box))
})
	})
	return cache_Data_Map_Internal_toMapIter__772765521
}

var cache_Data_Map_Internal_toMapIter__1738891721 gopurs_runtime.Value
var once_Data_Map_Internal_toMapIter__1738891721 sync.Once
func Get_Data_Map_Internal_toMapIter__1738891721() gopurs_runtime.Value {
	once_Data_Map_Internal_toMapIter__1738891721.Do(func() {
		cache_Data_Map_Internal_toMapIter__1738891721 = gopurs_runtime.Func(func(a_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Map_Internal_toMapIter__1738891721(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](a_0_box))
})
	})
	return cache_Data_Map_Internal_toMapIter__1738891721
}

var cache_Data_Map_Internal_toUnfoldable__2183602684 gopurs_runtime.Value
var once_Data_Map_Internal_toUnfoldable__2183602684 sync.Once
func Get_Data_Map_Internal_toUnfoldable__2183602684() gopurs_runtime.Value {
	once_Data_Map_Internal_toUnfoldable__2183602684.Do(func() {
		cache_Data_Map_Internal_toUnfoldable__2183602684 = gopurs_runtime.Func(func(dictUnfoldable_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Map_Internal_toUnfoldable__2183602684(gopurs_runtime.CoerceToStruct[Constructor_Data_Unfoldable_Unfoldable[gopurs_runtime.Value]](dictUnfoldable_0_box))
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
return Get_Data_Map_Internal_foldableMap()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Map_Internal_functorMap()
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_Map_Internal_traversableMap(), "traverse"), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](dictApplicative_0))}, Get_Data_Map_Internal_identity())
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Apply0_1_0 -> *Constructor_Control_Apply_Apply[gopurs_runtime.Value]
Apply0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}))
_ = Apply0_1_0
// TAST (Let): Functor0_2_1 -> *Constructor_Data_Functor_Functor[gopurs_runtime.Value]
Functor0_2_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_2_1
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_4_2_73 gopurs_runtime.Value
_ = go__go_4_2_73
go__go_4_2_73 = gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
var __t6 gopurs_runtime.Value
{
if (v_5.Type == 9 && v_5.IntVal == 324739070 && v_5.UnsafePtr == nil) {
__t6 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "pure"), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(nil))})
goto end_branch_6
} else {

}
}
{
if (v_5.Type == 9 && v_5.IntVal == 324739070 && v_5.UnsafePtr != nil) {
// TAST (Let): __local_var_6_3 -> gopurs_runtime.Value
var __local_var_6_3 gopurs_runtime.Value = gopurs_runtime.Int((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_5.UnsafePtr).V0)
// TAST (Let): __local_var_7_4 -> gopurs_runtime.Value
__local_var_7_4 := (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_5.UnsafePtr).V2
_ = __local_var_7_4
// TAST (Let): __local_var_8_5 -> gopurs_runtime.Value
var __local_var_8_5 gopurs_runtime.Value = gopurs_runtime.Int((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_5.UnsafePtr).V1)
__t6 = gopurs_runtime.Apply2(gopurs_runtime.Box(Apply0_1_0.V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Apply0_1_0.V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_2_1.V0), gopurs_runtime.Func(func(l_prime_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_prime_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(r_prime_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(&Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]{1, __local_var_6_3.IntVal, __local_var_8_5.IntVal, __local_var_7_4, v_prime_10, gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](l_prime_9), gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](r_prime_11)})}
})
})
}), gopurs_runtime.Apply(go__go_4_2_73, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_5.UnsafePtr).V4)})), gopurs_runtime.Apply(f_3, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_5.UnsafePtr).V3)), gopurs_runtime.Apply(go__go_4_2_73, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_5.UnsafePtr).V5)}))
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
return go__go_4_2_73
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
return Get_Data_Map_Internal_foldableMap()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Map_Internal_functorMap()
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_Map_Internal_traversableMap(), "traverse"), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](dictApplicative_0))}, Get_Data_Map_Internal_identity())
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Apply0_1_0 -> *Constructor_Control_Apply_Apply[gopurs_runtime.Value]
Apply0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}))
_ = Apply0_1_0
// TAST (Let): Functor0_2_1 -> *Constructor_Data_Functor_Functor[gopurs_runtime.Value]
Functor0_2_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_2_1
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_4_2_74 gopurs_runtime.Value
_ = go__go_4_2_74
go__go_4_2_74 = gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
var __t6 gopurs_runtime.Value
{
if (v_5.Type == 9 && v_5.IntVal == 324739070 && v_5.UnsafePtr == nil) {
__t6 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "pure"), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(nil))})
goto end_branch_6
} else {

}
}
{
if (v_5.Type == 9 && v_5.IntVal == 324739070 && v_5.UnsafePtr != nil) {
// TAST (Let): __local_var_6_3 -> gopurs_runtime.Value
var __local_var_6_3 gopurs_runtime.Value = gopurs_runtime.Int((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_5.UnsafePtr).V0)
// TAST (Let): __local_var_7_4 -> gopurs_runtime.Value
__local_var_7_4 := (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_5.UnsafePtr).V2
_ = __local_var_7_4
// TAST (Let): __local_var_8_5 -> gopurs_runtime.Value
var __local_var_8_5 gopurs_runtime.Value = gopurs_runtime.Int((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_5.UnsafePtr).V1)
__t6 = gopurs_runtime.Apply2(gopurs_runtime.Box(Apply0_1_0.V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Apply0_1_0.V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_2_1.V0), gopurs_runtime.Func(func(l_prime_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_prime_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(r_prime_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(&Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]{1, __local_var_6_3.IntVal, __local_var_8_5.IntVal, __local_var_7_4, v_prime_10, gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](l_prime_9), gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](r_prime_11)})}
})
})
}), gopurs_runtime.Apply(go__go_4_2_74, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_5.UnsafePtr).V4)})), gopurs_runtime.Apply(f_3, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_5.UnsafePtr).V3)), gopurs_runtime.Apply(go__go_4_2_74, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_5.UnsafePtr).V5)}))
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
return go__go_4_2_74
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
return Get_Data_Map_Internal_foldableWithIndexMap()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Map_Internal_functorWithIndexMap()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Map_Internal_traversableMap()
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Apply0_1_0 -> *Constructor_Control_Apply_Apply[gopurs_runtime.Value]
Apply0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}))
_ = Apply0_1_0
// TAST (Let): Functor0_2_1 -> *Constructor_Data_Functor_Functor[gopurs_runtime.Value]
Functor0_2_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_2_1
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_4_2_75 gopurs_runtime.Value
_ = go__go_4_2_75
go__go_4_2_75 = gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
var __t6 gopurs_runtime.Value
{
if (v_5.Type == 9 && v_5.IntVal == 324739070 && v_5.UnsafePtr == nil) {
__t6 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "pure"), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(nil))})
goto end_branch_6
} else {

}
}
{
if (v_5.Type == 9 && v_5.IntVal == 324739070 && v_5.UnsafePtr != nil) {
// TAST (Let): __local_var_6_3 -> gopurs_runtime.Value
var __local_var_6_3 gopurs_runtime.Value = gopurs_runtime.Int((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_5.UnsafePtr).V0)
// TAST (Let): __local_var_7_4 -> gopurs_runtime.Value
__local_var_7_4 := (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_5.UnsafePtr).V2
_ = __local_var_7_4
// TAST (Let): __local_var_8_5 -> gopurs_runtime.Value
var __local_var_8_5 gopurs_runtime.Value = gopurs_runtime.Int((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_5.UnsafePtr).V1)
__t6 = gopurs_runtime.Apply2(gopurs_runtime.Box(Apply0_1_0.V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Apply0_1_0.V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_2_1.V0), gopurs_runtime.Func(func(l_prime_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_prime_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(r_prime_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(&Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]{1, __local_var_6_3.IntVal, __local_var_8_5.IntVal, __local_var_7_4, v_prime_10, gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](l_prime_9), gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](r_prime_11)})}
})
})
}), gopurs_runtime.Apply(go__go_4_2_75, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_5.UnsafePtr).V4)})), gopurs_runtime.Apply2(f_3, __local_var_7_4, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_5.UnsafePtr).V3)), gopurs_runtime.Apply(go__go_4_2_75, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_5.UnsafePtr).V5)}))
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
return go__go_4_2_75
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
return Call_Data_Map_Internal_unionWith__2507192643(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](dictOrd_0_box))
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

type Constructor_Data_Map_Internal_Leaf[T_k any, T_v any] struct {
	Rc uint32
}


type Constructor_Data_Map_Internal_Node[T_k any, T_v any] struct {
	Rc uint32
	V0 int64
	V1 int64
	V2 T_k
	V3 T_v
	V4 *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]
	V5 *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]
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
	V0 *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]
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
	V0 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
	V1 *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]
	V2 *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]
}


type Constructor_Data_Map_Internal_SplitLast[T_k any, T_v any] struct {
	Rc uint32
	V0 T_k
	V1 T_v
	V2 *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]
}


func Call_Data_Map_Internal_identity(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_Data_Map_Internal_identity1(x_0_loop *Constructor_Data_Maybe_Just[gopurs_runtime.Value]) *Constructor_Data_Maybe_Just[gopurs_runtime.Value] {
var x_0 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = x_0_loop
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
var __t4 gopurs_runtime.Value
{
if (__local_var_2.Type == 9 && __local_var_2.IntVal == 324739070 && __local_var_2.UnsafePtr == nil) {
var __t0 gopurs_runtime.Value
{
if (__local_var_3.Type == 9 && __local_var_3.IntVal == 324739070 && __local_var_3.UnsafePtr == nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(&Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]{1, 1, 1, __local_var_0, __local_var_1, gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__local_var_2), gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__local_var_3)})}
goto end_branch_0
} else {

}
}
{
if (__local_var_3.Type == 9 && __local_var_3.IntVal == 324739070 && __local_var_3.UnsafePtr != nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(&Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]{1, (1) + ((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V0), (1) + ((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V1), __local_var_0, __local_var_1, gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__local_var_2), gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__local_var_3)})}
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
__t4 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__t0))}
goto end_branch_4
} else {

}
}
{
if (__local_var_2.Type == 9 && __local_var_2.IntVal == 324739070 && __local_var_2.UnsafePtr != nil) {
var __t3 gopurs_runtime.Value
{
if (__local_var_3.Type == 9 && __local_var_3.IntVal == 324739070 && __local_var_3.UnsafePtr == nil) {
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(&Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]{1, (1) + ((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V0), (1) + ((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V1), __local_var_0, __local_var_1, gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__local_var_2), gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__local_var_3)})}
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
if ((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V0) > ((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V0) {
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
__t2 = (1) + ((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V0)
goto end_branch_2
} else {

}
}
{
__t2 = (1) + ((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V0)
}
end_branch_2:
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(&Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]{1, __t2, ((1) + ((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V1)) + ((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V1), __local_var_0, __local_var_1, gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__local_var_2), gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__local_var_3)})}
goto end_branch_3
} else {

}
}
{
__t3 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_3:
__t4 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__t3))}
goto end_branch_4
} else {

}
}
{
__t4 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_4:
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__t4))}
}

func Call_Data_Map_Internal_toMapIter(a_0_loop *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var a_0 *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value] = a_0_loop
_ = a_0
return gopurs_runtime.Value{Type: 9, IntVal: 2861335956, UnsafePtr: unsafe.Pointer(&Constructor_Data_Map_Internal_IterNode[gopurs_runtime.Value, gopurs_runtime.Value]{1, a_0, gopurs_runtime.Value{Type: 9, IntVal: 2509360378, UnsafePtr: unsafe.Pointer(nil)}})}
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
__t1 = gopurs_runtime.UncurriedApp3(next_1, (*Constructor_Data_Map_Internal_IterEmit[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V0, (*Constructor_Data_Map_Internal_IterEmit[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V1, (*Constructor_Data_Map_Internal_IterEmit[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V2)
goto end_branch_1
} else {

}
}
{
if (v_4.Type == 9 && v_4.IntVal == 2861335956) {
v_4_loop = gopurs_runtime.Apply2(f_0, (*Constructor_Data_Map_Internal_IterNode[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_IterNode[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V0)})
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

func Call_Data_Map_Internal_size(v_0_loop *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]) int64 {
var v_0 *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value] = v_0_loop
_ = v_0
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_0)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_0)}.IntVal == 324739070 && gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr == nil) {
__t0 = gopurs_runtime.Int(0)
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_0)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_0)}.IntVal == 324739070 && gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr != nil) {
__t0 = gopurs_runtime.Int((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V1)
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return __t0.IntVal
}

func Call_Data_Map_Internal_singleton(k_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value] {
var k_0 gopurs_runtime.Value = k_0_loop
_ = k_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(&Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]{1, 1, 1, k_0, v_1, gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(nil))}), gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(nil))})})})
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
var __t37 gopurs_runtime.Value
{
if (__local_var_2.Type == 9 && __local_var_2.IntVal == 324739070 && __local_var_2.UnsafePtr == nil) {
var __t9 gopurs_runtime.Value
{
if (__local_var_3.Type == 9 && __local_var_3.IntVal == 324739070 && __local_var_3.UnsafePtr == nil) {
__t9 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(&Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]{1, 1, 1, __local_var_0, __local_var_1, gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(nil))}), gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(nil))})})}
goto end_branch_9
} else {

}
}
{
var __t_and_1 bool = false
if (__local_var_3.Type == 9 && __local_var_3.IntVal == 324739070 && __local_var_3.UnsafePtr != nil) {

var __t0 bool
{
if ((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V0) > (1) {
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
var __t8 *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]
{
var __t_tag_2 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V4)}
var __t_and_7 bool = false
if (__t_tag_2.Type == 9 && __t_tag_2.IntVal == 324739070 && __t_tag_2.UnsafePtr != nil) {

var __t6 bool
{
var __t5 gopurs_runtime.Value
{
var __t_tag_3 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V5)}
if (__t_tag_3.Type == 9 && __t_tag_3.IntVal == 324739070 && __t_tag_3.UnsafePtr == nil) {
__t5 = gopurs_runtime.Int(0)
goto end_branch_5
} else {

}
}
{
var __t_tag_4 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V5)}
if (__t_tag_4.Type == 9 && __t_tag_4.IntVal == 324739070 && __t_tag_4.UnsafePtr != nil) {
__t5 = gopurs_runtime.Int((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V5)}.UnsafePtr).V0)
goto end_branch_5
} else {

}
}
{
__t5 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_5:
if ((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V4)}.UnsafePtr).V0) > (__t5.IntVal) {
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
__t8 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeNode(), (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V4)}.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V4)}.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeNode(), __local_var_0, __local_var_1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__local_var_2))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V4)}.UnsafePtr).V4)})))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeNode(), (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V4)}.UnsafePtr).V5)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V5)})))}))
goto end_branch_8
} else {

}
}
{
__t8 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeNode(), (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeNode(), __local_var_0, __local_var_1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__local_var_2))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V4)})))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V5)}))
}
end_branch_8:
__t9 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(__t8)}
goto end_branch_9
} else {

}
}
{
__t9 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeNode(), __local_var_0, __local_var_1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__local_var_2))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__local_var_3))})))}
}
end_branch_9:
__t37 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__t9))}
goto end_branch_37
} else {

}
}
{
if (__local_var_2.Type == 9 && __local_var_2.IntVal == 324739070 && __local_var_2.UnsafePtr != nil) {
var __t36 *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]
{
if (__local_var_3.Type == 9 && __local_var_3.IntVal == 324739070 && __local_var_3.UnsafePtr != nil) {
var __t26 *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]
{
var __t10 bool
{
if ((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V0) > (((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V0) + (1)) {
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
var __t17 *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]
{
var __t_tag_11 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V4)}
var __t_and_16 bool = false
if (__t_tag_11.Type == 9 && __t_tag_11.IntVal == 324739070 && __t_tag_11.UnsafePtr != nil) {

var __t15 bool
{
var __t14 gopurs_runtime.Value
{
var __t_tag_12 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V5)}
if (__t_tag_12.Type == 9 && __t_tag_12.IntVal == 324739070 && __t_tag_12.UnsafePtr == nil) {
__t14 = gopurs_runtime.Int(0)
goto end_branch_14
} else {

}
}
{
var __t_tag_13 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V5)}
if (__t_tag_13.Type == 9 && __t_tag_13.IntVal == 324739070 && __t_tag_13.UnsafePtr != nil) {
__t14 = gopurs_runtime.Int((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V5)}.UnsafePtr).V0)
goto end_branch_14
} else {

}
}
{
__t14 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_14:
if ((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V4)}.UnsafePtr).V0) > (__t14.IntVal) {
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
__t17 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeNode(), (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V4)}.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V4)}.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeNode(), __local_var_0, __local_var_1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__local_var_2))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V4)}.UnsafePtr).V4)})))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeNode(), (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V4)}.UnsafePtr).V5)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V5)})))}))
goto end_branch_17
} else {

}
}
{
__t17 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeNode(), (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeNode(), __local_var_0, __local_var_1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__local_var_2))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V4)})))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V5)}))
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
if ((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V0) > (((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V0) + (1)) {
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
var __t25 *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]
{
var __t_tag_19 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V5)}
var __t_and_24 bool = false
if (__t_tag_19.Type == 9 && __t_tag_19.IntVal == 324739070 && __t_tag_19.UnsafePtr != nil) {

var __t23 bool
{
var __t22 gopurs_runtime.Value
{
var __t_tag_20 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V4)}
if (__t_tag_20.Type == 9 && __t_tag_20.IntVal == 324739070 && __t_tag_20.UnsafePtr == nil) {
__t22 = gopurs_runtime.Int(0)
goto end_branch_22
} else {

}
}
{
var __t_tag_21 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V4)}
if (__t_tag_21.Type == 9 && __t_tag_21.IntVal == 324739070 && __t_tag_21.UnsafePtr != nil) {
__t22 = gopurs_runtime.Int((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V4)}.UnsafePtr).V0)
goto end_branch_22
} else {

}
}
{
__t22 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_22:
if (__t22.IntVal) > ((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V5)}.UnsafePtr).V0) {
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
__t25 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeNode(), (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V5)}.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V5)}.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeNode(), (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V4)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V5)}.UnsafePtr).V4)})))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeNode(), __local_var_0, __local_var_1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V5)}.UnsafePtr).V5)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__local_var_3))})))}))
goto end_branch_25
} else {

}
}
{
__t25 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeNode(), (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V4)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeNode(), __local_var_0, __local_var_1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V5)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__local_var_3))})))}))
}
end_branch_25:
__t26 = __t25
goto end_branch_26
} else {

}
}
{
__t26 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeNode(), __local_var_0, __local_var_1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__local_var_2))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__local_var_3))}))
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
if ((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V0) > (1) {
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
var __t35 *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]
{
var __t_tag_29 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V5)}
var __t_and_34 bool = false
if (__t_tag_29.Type == 9 && __t_tag_29.IntVal == 324739070 && __t_tag_29.UnsafePtr != nil) {

var __t33 bool
{
var __t32 gopurs_runtime.Value
{
var __t_tag_30 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V4)}
if (__t_tag_30.Type == 9 && __t_tag_30.IntVal == 324739070 && __t_tag_30.UnsafePtr == nil) {
__t32 = gopurs_runtime.Int(0)
goto end_branch_32
} else {

}
}
{
var __t_tag_31 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V4)}
if (__t_tag_31.Type == 9 && __t_tag_31.IntVal == 324739070 && __t_tag_31.UnsafePtr != nil) {
__t32 = gopurs_runtime.Int((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V4)}.UnsafePtr).V0)
goto end_branch_32
} else {

}
}
{
__t32 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_32:
if (__t32.IntVal) > ((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V5)}.UnsafePtr).V0) {
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
__t35 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeNode(), (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V5)}.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V5)}.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeNode(), (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V4)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V5)}.UnsafePtr).V4)})))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeNode(), __local_var_0, __local_var_1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V5)}.UnsafePtr).V5)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__local_var_3))})))}))
goto end_branch_35
} else {

}
}
{
__t35 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeNode(), (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V4)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeNode(), __local_var_0, __local_var_1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V5)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__local_var_3))})))}))
}
end_branch_35:
__t36 = __t35
goto end_branch_36
} else {

}
}
{
__t36 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeNode(), __local_var_0, __local_var_1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__local_var_2))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__local_var_3))}))
}
end_branch_36:
__t37 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(__t36)}
goto end_branch_37
} else {

}
}
{
__t37 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_37:
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__t37))}
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
var __t4 gopurs_runtime.Value
{
if (__local_var_2.Type == 9 && __local_var_2.IntVal == 324739070 && __local_var_2.UnsafePtr == nil) {
__t4 = gopurs_runtime.Value{Type: 9, IntVal: 3373277644, UnsafePtr: unsafe.Pointer(&Constructor_Data_Map_Internal_Split[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil))}), gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(nil))}), gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(nil))})})}
goto end_branch_4
} else {

}
}
{
if (__local_var_2.Type == 9 && __local_var_2.IntVal == 324739070 && __local_var_2.UnsafePtr != nil) {
// TAST (Let): v_3_0 -> gopurs_runtime.Value
v_3_0 := gopurs_runtime.Apply2(__local_var_0, __local_var_1, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V2)
_ = v_3_0
var __t3 gopurs_runtime.Value
{
if (uint32(v_3_0.IntVal) == 1527465420) {
// TAST (Let): v1_4_1 -> gopurs_runtime.Value
v1_4_1 := gopurs_runtime.UncurriedApp3(Get_Data_Map_Internal_unsafeSplit(), __local_var_0, __local_var_1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V4)})
_ = v1_4_1
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 3373277644, UnsafePtr: unsafe.Pointer(&Constructor_Data_Map_Internal_Split[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Data_Map_Internal_Split[gopurs_runtime.Value, gopurs_runtime.Value])(v1_4_1.UnsafePtr).V0, (*Constructor_Data_Map_Internal_Split[gopurs_runtime.Value, gopurs_runtime.Value])(v1_4_1.UnsafePtr).V1, gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeBalancedNode(), (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Split[gopurs_runtime.Value, gopurs_runtime.Value])(v1_4_1.UnsafePtr).V2)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V5)}))})}
goto end_branch_3
} else {

}
}
{
if (uint32(v_3_0.IntVal) == 380165415) {
// TAST (Let): v1_4_2 -> gopurs_runtime.Value
v1_4_2 := gopurs_runtime.UncurriedApp3(Get_Data_Map_Internal_unsafeSplit(), __local_var_0, __local_var_1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V5)})
_ = v1_4_2
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 3373277644, UnsafePtr: unsafe.Pointer(&Constructor_Data_Map_Internal_Split[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Data_Map_Internal_Split[gopurs_runtime.Value, gopurs_runtime.Value])(v1_4_2.UnsafePtr).V0, gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeBalancedNode(), (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V4)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Split[gopurs_runtime.Value, gopurs_runtime.Value])(v1_4_2.UnsafePtr).V1)})), (*Constructor_Data_Map_Internal_Split[gopurs_runtime.Value, gopurs_runtime.Value])(v1_4_2.UnsafePtr).V2})}
goto end_branch_3
} else {

}
}
{
if (uint32(v_3_0.IntVal) == 902936544) {
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 3373277644, UnsafePtr: unsafe.Pointer(&Constructor_Data_Map_Internal_Split[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{1, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V3})}), (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V4, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V5})}
goto end_branch_3
} else {

}
}
{
__t3 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_3:
__t4 = gopurs_runtime.Value{Type: 9, IntVal: 3373277644, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Split[gopurs_runtime.Value, gopurs_runtime.Value]](__t3))}
goto end_branch_4
} else {

}
}
{
__t4 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_4:
return gopurs_runtime.Value{Type: 9, IntVal: 3373277644, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Split[gopurs_runtime.Value, gopurs_runtime.Value]](__t4))}
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
var __t1 gopurs_runtime.Value
{
if (__local_var_3.Type == 9 && __local_var_3.IntVal == 324739070 && __local_var_3.UnsafePtr == nil) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 2668112006, UnsafePtr: unsafe.Pointer(&Constructor_Data_Map_Internal_SplitLast[gopurs_runtime.Value, gopurs_runtime.Value]{1, __local_var_0, __local_var_1, gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__local_var_2)})}
goto end_branch_1
} else {

}
}
{
if (__local_var_3.Type == 9 && __local_var_3.IntVal == 324739070 && __local_var_3.UnsafePtr != nil) {
// TAST (Let): v1_4_0 -> gopurs_runtime.Value
v1_4_0 := gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeSplitLast(), (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V4)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V5)})
_ = v1_4_0
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 2668112006, UnsafePtr: unsafe.Pointer(&Constructor_Data_Map_Internal_SplitLast[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Data_Map_Internal_SplitLast[gopurs_runtime.Value, gopurs_runtime.Value])(v1_4_0.UnsafePtr).V0, (*Constructor_Data_Map_Internal_SplitLast[gopurs_runtime.Value, gopurs_runtime.Value])(v1_4_0.UnsafePtr).V1, gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeBalancedNode(), __local_var_0, __local_var_1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__local_var_2))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_SplitLast[gopurs_runtime.Value, gopurs_runtime.Value])(v1_4_0.UnsafePtr).V2)}))})}
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
return gopurs_runtime.Value{Type: 9, IntVal: 2668112006, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_SplitLast[gopurs_runtime.Value, gopurs_runtime.Value]](__t1))}
}
}

func Call_Data_Map_Internal_unsafeJoinNodes(__local_var_0_loop gopurs_runtime.Value, __local_var_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __local_var_0 gopurs_runtime.Value = __local_var_0_loop
_ = __local_var_0
var __local_var_1 gopurs_runtime.Value = __local_var_1_loop
_ = __local_var_1
var __t1 gopurs_runtime.Value
{
if (__local_var_0.Type == 9 && __local_var_0.IntVal == 324739070 && __local_var_0.UnsafePtr == nil) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__local_var_1))}
goto end_branch_1
} else {

}
}
{
if (__local_var_0.Type == 9 && __local_var_0.IntVal == 324739070 && __local_var_0.UnsafePtr != nil) {
// TAST (Let): v2_2_0 -> gopurs_runtime.Value
v2_2_0 := gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeSplitLast(), (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_0.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_0.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_0.UnsafePtr).V4)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_0.UnsafePtr).V5)})
_ = v2_2_0
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeBalancedNode(), (*Constructor_Data_Map_Internal_SplitLast[gopurs_runtime.Value, gopurs_runtime.Value])(v2_2_0.UnsafePtr).V0, (*Constructor_Data_Map_Internal_SplitLast[gopurs_runtime.Value, gopurs_runtime.Value])(v2_2_0.UnsafePtr).V1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_SplitLast[gopurs_runtime.Value, gopurs_runtime.Value])(v2_2_0.UnsafePtr).V2)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__local_var_1))})))}
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__t1))}
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
var __t1 gopurs_runtime.Value
{
if (__local_var_1.Type == 9 && __local_var_1.IntVal == 324739070 && __local_var_1.UnsafePtr == nil) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(nil))}
goto end_branch_1
} else {

}
}
{
if (__local_var_2.Type == 9 && __local_var_2.IntVal == 324739070 && __local_var_2.UnsafePtr == nil) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__local_var_1))}
goto end_branch_1
} else {

}
}
{
if (__local_var_2.Type == 9 && __local_var_2.IntVal == 324739070 && __local_var_2.UnsafePtr != nil) {
// TAST (Let): v_3_0 -> gopurs_runtime.Value
v_3_0 := gopurs_runtime.UncurriedApp3(Get_Data_Map_Internal_unsafeSplit(), __local_var_0, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V2, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__local_var_1))})
_ = v_3_0
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp2(Get_Data_Map_Internal_unsafeJoinNodes(), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp3(Get_Data_Map_Internal_unsafeDifference(), __local_var_0, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Split[gopurs_runtime.Value, gopurs_runtime.Value])(v_3_0.UnsafePtr).V1)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V4)})))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp3(Get_Data_Map_Internal_unsafeDifference(), __local_var_0, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Split[gopurs_runtime.Value, gopurs_runtime.Value])(v_3_0.UnsafePtr).V2)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V5)})))})))}
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__t1))}
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
var __t6 gopurs_runtime.Value
{
if (__local_var_2.Type == 9 && __local_var_2.IntVal == 324739070 && __local_var_2.UnsafePtr == nil) {
__t6 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(nil))}
goto end_branch_6
} else {

}
}
{
if (__local_var_3.Type == 9 && __local_var_3.IntVal == 324739070 && __local_var_3.UnsafePtr == nil) {
__t6 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(nil))}
goto end_branch_6
} else {

}
}
{
if (__local_var_3.Type == 9 && __local_var_3.IntVal == 324739070 && __local_var_3.UnsafePtr != nil) {
// TAST (Let): v_4_0 -> gopurs_runtime.Value
v_4_0 := gopurs_runtime.UncurriedApp3(Get_Data_Map_Internal_unsafeSplit(), __local_var_0, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V2, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__local_var_2))})
_ = v_4_0
// TAST (Let): l_prime_5_1 -> gopurs_runtime.Value
l_prime_5_1 := gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeIntersectionWith(), __local_var_0, __local_var_1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Split[gopurs_runtime.Value, gopurs_runtime.Value])(v_4_0.UnsafePtr).V1)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V4)})
_ = l_prime_5_1
// TAST (Let): r_prime_6_2 -> gopurs_runtime.Value
r_prime_6_2 := gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeIntersectionWith(), __local_var_0, __local_var_1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Split[gopurs_runtime.Value, gopurs_runtime.Value])(v_4_0.UnsafePtr).V2)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V5)})
_ = r_prime_6_2
var __t5 gopurs_runtime.Value
{
var __t_tag_3 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Split[gopurs_runtime.Value, gopurs_runtime.Value])(v_4_0.UnsafePtr).V0)}
if (__t_tag_3.Type == 9 && __t_tag_3.IntVal == 930809136 && __t_tag_3.UnsafePtr != nil) {
__t5 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeBalancedNode(), (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V2, gopurs_runtime.Apply2(__local_var_1, (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Split[gopurs_runtime.Value, gopurs_runtime.Value])(v_4_0.UnsafePtr).V0)}.UnsafePtr).V0, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V3), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](l_prime_5_1))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](r_prime_6_2))})))}
goto end_branch_5
} else {

}
}
{
var __t_tag_4 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Split[gopurs_runtime.Value, gopurs_runtime.Value])(v_4_0.UnsafePtr).V0)}
if (__t_tag_4.Type == 9 && __t_tag_4.IntVal == 930809136 && __t_tag_4.UnsafePtr == nil) {
__t5 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp2(Get_Data_Map_Internal_unsafeJoinNodes(), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](l_prime_5_1))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](r_prime_6_2))})))}
goto end_branch_5
} else {

}
}
{
__t5 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_5:
__t6 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__t5))}
goto end_branch_6
} else {

}
}
{
__t6 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_6:
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__t6))}
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
var __t6 gopurs_runtime.Value
{
if (__local_var_2.Type == 9 && __local_var_2.IntVal == 324739070 && __local_var_2.UnsafePtr == nil) {
__t6 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__local_var_3))}
goto end_branch_6
} else {

}
}
{
if (__local_var_3.Type == 9 && __local_var_3.IntVal == 324739070 && __local_var_3.UnsafePtr == nil) {
__t6 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__local_var_2))}
goto end_branch_6
} else {

}
}
{
if (__local_var_3.Type == 9 && __local_var_3.IntVal == 324739070 && __local_var_3.UnsafePtr != nil) {
// TAST (Let): v_4_0 -> gopurs_runtime.Value
v_4_0 := gopurs_runtime.UncurriedApp3(Get_Data_Map_Internal_unsafeSplit(), __local_var_0, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V2, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__local_var_2))})
_ = v_4_0
// TAST (Let): l_prime_5_1 -> gopurs_runtime.Value
l_prime_5_1 := gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeUnionWith(), __local_var_0, __local_var_1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Split[gopurs_runtime.Value, gopurs_runtime.Value])(v_4_0.UnsafePtr).V1)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V4)})
_ = l_prime_5_1
// TAST (Let): r_prime_6_2 -> gopurs_runtime.Value
r_prime_6_2 := gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeUnionWith(), __local_var_0, __local_var_1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Split[gopurs_runtime.Value, gopurs_runtime.Value])(v_4_0.UnsafePtr).V2)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V5)})
_ = r_prime_6_2
var __t5 gopurs_runtime.Value
{
var __t_tag_3 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Split[gopurs_runtime.Value, gopurs_runtime.Value])(v_4_0.UnsafePtr).V0)}
if (__t_tag_3.Type == 9 && __t_tag_3.IntVal == 930809136 && __t_tag_3.UnsafePtr != nil) {
__t5 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeBalancedNode(), (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V2, gopurs_runtime.Apply2(__local_var_1, (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Split[gopurs_runtime.Value, gopurs_runtime.Value])(v_4_0.UnsafePtr).V0)}.UnsafePtr).V0, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V3), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](l_prime_5_1))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](r_prime_6_2))})))}
goto end_branch_5
} else {

}
}
{
var __t_tag_4 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Split[gopurs_runtime.Value, gopurs_runtime.Value])(v_4_0.UnsafePtr).V0)}
if (__t_tag_4.Type == 9 && __t_tag_4.IntVal == 930809136 && __t_tag_4.UnsafePtr == nil) {
__t5 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeBalancedNode(), (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](l_prime_5_1))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](r_prime_6_2))})))}
goto end_branch_5
} else {

}
}
{
__t5 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_5:
__t6 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__t5))}
goto end_branch_6
} else {

}
}
{
__t6 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_6:
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__t6))}
}
}

func Call_Data_Map_Internal_unionWith(dictOrd_0_loop *Constructor_Data_Ord_Ord[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictOrd_0 *Constructor_Data_Ord_Ord[gopurs_runtime.Value] = dictOrd_0_loop
_ = dictOrd_0
// TAST (Let): compare_1_0 -> gopurs_runtime.Value
compare_1_0 := gopurs_runtime.Box(dictOrd_0.V1)
_ = compare_1_0
return gopurs_runtime.Func(func(app_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m2_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeUnionWith(), compare_1_0, app_2, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](m1_3))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](m2_4))})))}
})
})
})
}

func Call_Data_Map_Internal_union(dictOrd_0_loop *Constructor_Data_Ord_Ord[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictOrd_0 *Constructor_Data_Ord_Ord[gopurs_runtime.Value] = dictOrd_0_loop
_ = dictOrd_0
// TAST (Let): compare_1_0 -> gopurs_runtime.Value
compare_1_0 := gopurs_runtime.Box(dictOrd_0.V1)
_ = compare_1_0
return gopurs_runtime.Func(func(m1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m2_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeUnionWith(), compare_1_0, Get_Data_Function_go__const(), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](m1_2))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](m2_3))})))}
})
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
go__go_3_0_1 = gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t5 gopurs_runtime.Value
{
if (v_4.Type == 9 && v_4.IntVal == 324739070 && v_4.UnsafePtr == nil) {
__t5 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(nil))}
goto end_branch_5
} else {

}
}
{
if (v_4.Type == 9 && v_4.IntVal == 324739070 && v_4.UnsafePtr != nil) {
// TAST (Let): v1_5_1 -> gopurs_runtime.Value
v1_5_1 := gopurs_runtime.Apply2(gopurs_runtime.Box(dictOrd_0.V1), k_2, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V2)
_ = v1_5_1
var __t4 gopurs_runtime.Value
{
if (uint32(v1_5_1.IntVal) == 1527465420) {
__t4 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeBalancedNode(), (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(go__go_3_0_1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V4)})))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V5)})))}
goto end_branch_4
} else {

}
}
{
if (uint32(v1_5_1.IntVal) == 380165415) {
__t4 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeBalancedNode(), (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V4)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(go__go_3_0_1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V5)})))})))}
goto end_branch_4
} else {

}
}
{
if (uint32(v1_5_1.IntVal) == 902936544) {
// TAST (Let): v2_6_2 -> *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
v2_6_2 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(f_1, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V3))
_ = v2_6_2
var __t3 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_6_2)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_6_2)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_6_2)}.UnsafePtr == nil) {
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp2(Get_Data_Map_Internal_unsafeJoinNodes(), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V4)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V5)})))}
goto end_branch_3
} else {

}
}
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_6_2)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_6_2)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_6_2)}.UnsafePtr != nil) {
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(&Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V0, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V1, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V2, (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_6_2)}.UnsafePtr).V0, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V4, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V5})}
goto end_branch_3
} else {

}
}
{
__t3 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_3:
__t4 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__t3))}
goto end_branch_4
} else {

}
}
{
__t4 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_4:
__t5 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__t4))}
goto end_branch_5
} else {

}
}
{
__t5 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_5:
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__t5))}
})
return go__go_3_0_1
}

func Call_Data_Map_Internal_showTree(dictShow_0_loop *Constructor_Data_Show_Show[gopurs_runtime.Value], dictShow1_1_loop *Constructor_Data_Show_Show[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictShow_0 *Constructor_Data_Show_Show[gopurs_runtime.Value] = dictShow_0_loop
_ = dictShow_0
var dictShow1_1 *Constructor_Data_Show_Show[gopurs_runtime.Value] = dictShow1_1_loop
_ = dictShow1_1
var go__go_2_0_2 gopurs_runtime.Value
_ = go__go_2_0_2
go__go_2_0_2 = gopurs_runtime.Func(func(ind_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if (v_4.Type == 9 && v_4.IntVal == 324739070 && v_4.UnsafePtr == nil) {
__t1 = gopurs_runtime.Str((ind_3.StrVal()) + ("Leaf"))
goto end_branch_1
} else {

}
}
{
if (v_4.Type == 9 && v_4.IntVal == 324739070 && v_4.UnsafePtr != nil) {
__t1 = gopurs_runtime.Str(((((((((((ind_3.StrVal()) + ("[")) + (gopurs_runtime.Apply(Get_Data_Show_showIntImpl(), gopurs_runtime.Int((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V0)).StrVal())) + ("] ")) + (gopurs_runtime.Apply(gopurs_runtime.Box(dictShow_0.V0), (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V2).StrVal())) + (" => ")) + (gopurs_runtime.Apply(gopurs_runtime.Box(dictShow1_1.V0), (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V3).StrVal())) + ("\x0a")) + (gopurs_runtime.Apply2(go__go_2_0_2, gopurs_runtime.Str((ind_3.StrVal()) + ("    ")), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V4)}).StrVal())) + ("\x0a")) + (gopurs_runtime.Apply2(go__go_2_0_2, gopurs_runtime.Str((ind_3.StrVal()) + ("    ")), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V5)}).StrVal()))
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
return gopurs_runtime.Str(__t1.StrVal())
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
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeUnionWith(), compare_3_0, __local_var_4_1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](m1_5))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](m2_6))})))}
})
}))
}

func Call_Data_Map_Internal_pop(dictOrd_0_loop *Constructor_Data_Ord_Ord[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictOrd_0 *Constructor_Data_Ord_Ord[gopurs_runtime.Value] = dictOrd_0_loop
_ = dictOrd_0
// TAST (Let): compare_1_0 -> gopurs_runtime.Value
compare_1_0 := gopurs_runtime.Box(dictOrd_0.V1)
_ = compare_1_0
return gopurs_runtime.Func(func(k_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): v_4_1 -> gopurs_runtime.Value
v_4_1 := gopurs_runtime.UncurriedApp3(Get_Data_Map_Internal_unsafeSplit(), compare_1_0, k_2, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](m_3))})
_ = v_4_1
// TAST (Let): __local_var_5_2 -> gopurs_runtime.Value
var __local_var_5_2 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Split[gopurs_runtime.Value, gopurs_runtime.Value])(v_4_1.UnsafePtr).V1)}
// TAST (Let): __local_var_6_3 -> gopurs_runtime.Value
var __local_var_6_3 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Split[gopurs_runtime.Value, gopurs_runtime.Value])(v_4_1.UnsafePtr).V2)}
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]]]](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_Maybe_functorMaybe(), "map"), gopurs_runtime.Func(func(a_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, a_7, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp2(Get_Data_Map_Internal_unsafeJoinNodes(), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__local_var_5_2))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__local_var_6_3))})))}})}
}), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Split[gopurs_runtime.Value, gopurs_runtime.Value])(v_4_1.UnsafePtr).V0)})))}
})
})
}

func Call_Data_Map_Internal_member(dictOrd_0_loop *Constructor_Data_Ord_Ord[gopurs_runtime.Value], k_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 *Constructor_Data_Ord_Ord[gopurs_runtime.Value] = dictOrd_0_loop
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
var __t3 gopurs_runtime.Value
{
if (v_3.Type == 9 && v_3.IntVal == 324739070 && v_3.UnsafePtr == nil) {
__t3 = gopurs_runtime.Bool(false)
goto end_branch_3
} else {

}
}
{
if (v_3.Type == 9 && v_3.IntVal == 324739070 && v_3.UnsafePtr != nil) {
// TAST (Let): v1_4_1 -> gopurs_runtime.Value
v1_4_1 := gopurs_runtime.Apply2(gopurs_runtime.Box(dictOrd_0.V1), k_1, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V2)
_ = v1_4_1
var __t2 gopurs_runtime.Value
{
if (uint32(v1_4_1.IntVal) == 1527465420) {
v_3_loop = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V4)}
continue go__go_2_0_3
__t2 = gopurs_runtime.Bool((gopurs_runtime.Value{}.IntVal) != (0))
goto end_branch_2
} else {

}
}
{
if (uint32(v1_4_1.IntVal) == 380165415) {
v_3_loop = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V5)}
continue go__go_2_0_3
__t2 = gopurs_runtime.Bool((gopurs_runtime.Value{}.IntVal) != (0))
goto end_branch_2
} else {

}
}
{
if (uint32(v1_4_1.IntVal) == 902936544) {
__t2 = gopurs_runtime.Bool(true)
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_2:
__t3 = gopurs_runtime.Bool((__t2.IntVal) != (0))
goto end_branch_3
} else {

}
}
{
__t3 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_3:
return gopurs_runtime.Bool((__t3.IntVal) != (0))
}
}()
})
return go__go_2_0_3
}

func Call_Data_Map_Internal_mapMaybeWithKey(dictOrd_0_loop *Constructor_Data_Ord_Ord[gopurs_runtime.Value], f_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 *Constructor_Data_Ord_Ord[gopurs_runtime.Value] = dictOrd_0_loop
_ = dictOrd_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var go__go_2_0_4 gopurs_runtime.Value
_ = go__go_2_0_4
go__go_2_0_4 = gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t3 gopurs_runtime.Value
{
if (v_3.Type == 9 && v_3.IntVal == 324739070 && v_3.UnsafePtr == nil) {
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(nil))}
goto end_branch_3
} else {

}
}
{
if (v_3.Type == 9 && v_3.IntVal == 324739070 && v_3.UnsafePtr != nil) {
// TAST (Let): v2_4_1 -> gopurs_runtime.Value
v2_4_1 := gopurs_runtime.Apply2(f_1, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V3)
_ = v2_4_1
var __t2 gopurs_runtime.Value
{
if (v2_4_1.Type == 9 && v2_4_1.IntVal == 930809136 && v2_4_1.UnsafePtr != nil) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeBalancedNode(), (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V2, (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(v2_4_1.UnsafePtr).V0, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(go__go_2_0_4, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V4)})))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(go__go_2_0_4, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V5)})))})))}
goto end_branch_2
} else {

}
}
{
if (v2_4_1.Type == 9 && v2_4_1.IntVal == 930809136 && v2_4_1.UnsafePtr == nil) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp2(Get_Data_Map_Internal_unsafeJoinNodes(), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(go__go_2_0_4, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V4)})))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(go__go_2_0_4, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V5)})))})))}
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_2:
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__t2))}
goto end_branch_3
} else {

}
}
{
__t3 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_3:
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__t3))}
})
return go__go_2_0_4
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

func Call_Data_Map_Internal_lookupLE(dictOrd_0_loop *Constructor_Data_Ord_Ord[gopurs_runtime.Value], k_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 *Constructor_Data_Ord_Ord[gopurs_runtime.Value] = dictOrd_0_loop
_ = dictOrd_0
var k_1 gopurs_runtime.Value = k_1_loop
_ = k_1
var go__go_2_0_5 gopurs_runtime.Value
_ = go__go_2_0_5
go__go_2_0_5 = gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t5 gopurs_runtime.Value
{
if (v_3.Type == 9 && v_3.IntVal == 324739070 && v_3.UnsafePtr == nil) {
__t5 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil))}
goto end_branch_5
} else {

}
}
{
if (v_3.Type == 9 && v_3.IntVal == 324739070 && v_3.UnsafePtr != nil) {
// TAST (Let): v1_4_1 -> gopurs_runtime.Value
v1_4_1 := gopurs_runtime.Apply2(gopurs_runtime.Box(dictOrd_0.V1), k_1, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V2)
_ = v1_4_1
var __t4 gopurs_runtime.Value
{
if (uint32(v1_4_1.IntVal) == 1527465420) {
__t4 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(go__go_2_0_5, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V4)})))}
goto end_branch_4
} else {

}
}
{
if (uint32(v1_4_1.IntVal) == 380165415) {
// TAST (Let): v2_5_2 -> *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
v2_5_2 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(go__go_2_0_5, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V5)}))
_ = v2_5_2
var __t3 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_5_2)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_5_2)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_5_2)}.UnsafePtr == nil) {
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{1, gopurs_runtime.RecordDict2("key", "value", (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V3)})}
goto end_branch_3
} else {

}
}
{
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_5_2)}
}
end_branch_3:
__t4 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](__t3))}
goto end_branch_4
} else {

}
}
{
if (uint32(v1_4_1.IntVal) == 902936544) {
__t4 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{1, gopurs_runtime.RecordDict2("key", "value", (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V3)})}
goto end_branch_4
} else {

}
}
{
__t4 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_4:
__t5 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](__t4))}
goto end_branch_5
} else {

}
}
{
__t5 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_5:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](__t5))}
})
return go__go_2_0_5
}

func Call_Data_Map_Internal_lookupGE(dictOrd_0_loop *Constructor_Data_Ord_Ord[gopurs_runtime.Value], k_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 *Constructor_Data_Ord_Ord[gopurs_runtime.Value] = dictOrd_0_loop
_ = dictOrd_0
var k_1 gopurs_runtime.Value = k_1_loop
_ = k_1
var go__go_2_0_6 gopurs_runtime.Value
_ = go__go_2_0_6
go__go_2_0_6 = gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t5 gopurs_runtime.Value
{
if (v_3.Type == 9 && v_3.IntVal == 324739070 && v_3.UnsafePtr == nil) {
__t5 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil))}
goto end_branch_5
} else {

}
}
{
if (v_3.Type == 9 && v_3.IntVal == 324739070 && v_3.UnsafePtr != nil) {
// TAST (Let): v1_4_1 -> gopurs_runtime.Value
v1_4_1 := gopurs_runtime.Apply2(gopurs_runtime.Box(dictOrd_0.V1), k_1, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V2)
_ = v1_4_1
var __t4 gopurs_runtime.Value
{
if (uint32(v1_4_1.IntVal) == 1527465420) {
// TAST (Let): v2_5_2 -> *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
v2_5_2 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(go__go_2_0_6, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V4)}))
_ = v2_5_2
var __t3 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_5_2)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_5_2)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_5_2)}.UnsafePtr == nil) {
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{1, gopurs_runtime.RecordDict2("key", "value", (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V3)})}
goto end_branch_3
} else {

}
}
{
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_5_2)}
}
end_branch_3:
__t4 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](__t3))}
goto end_branch_4
} else {

}
}
{
if (uint32(v1_4_1.IntVal) == 380165415) {
__t4 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(go__go_2_0_6, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V5)})))}
goto end_branch_4
} else {

}
}
{
if (uint32(v1_4_1.IntVal) == 902936544) {
__t4 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{1, gopurs_runtime.RecordDict2("key", "value", (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V3)})}
goto end_branch_4
} else {

}
}
{
__t4 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_4:
__t5 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](__t4))}
goto end_branch_5
} else {

}
}
{
__t5 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_5:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](__t5))}
})
return go__go_2_0_6
}

func Call_Data_Map_Internal_lookup(dictOrd_0_loop *Constructor_Data_Ord_Ord[gopurs_runtime.Value], k_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 *Constructor_Data_Ord_Ord[gopurs_runtime.Value] = dictOrd_0_loop
_ = dictOrd_0
var k_1 gopurs_runtime.Value = k_1_loop
_ = k_1
var go__go_2_0_7 gopurs_runtime.Value
go__go_2_0_7 = gopurs_runtime.Func(func(v_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_3_loop gopurs_runtime.Value = v_3_loop_val
go__go_2_0_7:
for {
if false { continue go__go_2_0_7 }
var v_3 gopurs_runtime.Value = v_3_loop
_ = v_3
var __t3 gopurs_runtime.Value
{
if (v_3.Type == 9 && v_3.IntVal == 324739070 && v_3.UnsafePtr == nil) {
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil))}
goto end_branch_3
} else {

}
}
{
if (v_3.Type == 9 && v_3.IntVal == 324739070 && v_3.UnsafePtr != nil) {
// TAST (Let): v1_4_1 -> gopurs_runtime.Value
v1_4_1 := gopurs_runtime.Apply2(gopurs_runtime.Box(dictOrd_0.V1), k_1, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V2)
_ = v1_4_1
var __t2 gopurs_runtime.Value
{
if (uint32(v1_4_1.IntVal) == 1527465420) {
v_3_loop = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V4)}
continue go__go_2_0_7
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Value{}))}
goto end_branch_2
} else {

}
}
{
if (uint32(v1_4_1.IntVal) == 380165415) {
v_3_loop = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V5)}
continue go__go_2_0_7
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Value{}))}
goto end_branch_2
} else {

}
}
{
if (uint32(v1_4_1.IntVal) == 902936544) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{1, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V3})}
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_2:
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](__t2))}
goto end_branch_3
} else {

}
}
{
__t3 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_3:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](__t3))}
}
}()
})
return go__go_2_0_7
}

func Call_Data_Map_Internal_iterMapU(iter_0_loop gopurs_runtime.Value, v_1_loop *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var iter_0 gopurs_runtime.Value = iter_0_loop
_ = iter_0
var v_1 *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value] = v_1_loop
_ = v_1
var __t6 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_1)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_1)}.IntVal == 324739070 && gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_1)}.UnsafePtr == nil) {
__t6 = iter_0
goto end_branch_6
} else {

}
}
{
if (gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_1)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_1)}.IntVal == 324739070 && gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_1)}.UnsafePtr != nil) {
var __t5 *Constructor_Data_Map_Internal_IterEmit[gopurs_runtime.Value, gopurs_runtime.Value]
{
var __t_tag_2 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_1)}.UnsafePtr).V4)}
if (__t_tag_2.Type == 9 && __t_tag_2.IntVal == 324739070 && __t_tag_2.UnsafePtr == nil) {
var __t4 *Constructor_Data_Map_Internal_IterEmit[gopurs_runtime.Value, gopurs_runtime.Value]
{
var __t_tag_3 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_1)}.UnsafePtr).V5)}
if (__t_tag_3.Type == 9 && __t_tag_3.IntVal == 324739070 && __t_tag_3.UnsafePtr == nil) {
__t4 = &Constructor_Data_Map_Internal_IterEmit[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_1)}.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_1)}.UnsafePtr).V3, iter_0}
goto end_branch_4
} else {

}
}
{
__t4 = &Constructor_Data_Map_Internal_IterEmit[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_1)}.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_1)}.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 2861335956, UnsafePtr: unsafe.Pointer(&Constructor_Data_Map_Internal_IterNode[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_1)}.UnsafePtr).V5, iter_0})}}
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
var __t_tag_0 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_1)}.UnsafePtr).V5)}
if (__t_tag_0.Type == 9 && __t_tag_0.IntVal == 324739070 && __t_tag_0.UnsafePtr == nil) {
__t1 = &Constructor_Data_Map_Internal_IterEmit[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_1)}.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_1)}.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 2861335956, UnsafePtr: unsafe.Pointer(&Constructor_Data_Map_Internal_IterNode[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_1)}.UnsafePtr).V4, iter_0})}}
goto end_branch_1
} else {

}
}
{
__t1 = &Constructor_Data_Map_Internal_IterEmit[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_1)}.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_1)}.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 2861335956, UnsafePtr: unsafe.Pointer(&Constructor_Data_Map_Internal_IterNode[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_1)}.UnsafePtr).V4, gopurs_runtime.Value{Type: 9, IntVal: 2861335956, UnsafePtr: unsafe.Pointer(&Constructor_Data_Map_Internal_IterNode[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_1)}.UnsafePtr).V5, iter_0})}})}}
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
// TAST (Let): __local_var_1_0 -> gopurs_runtime.Value
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.Box(dictUnfoldable_0.V1), Get_Data_Map_Internal_stepUnfoldrUnordered())
_ = __local_var_1_0
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_1_0, gopurs_runtime.Value{Type: 9, IntVal: 2861335956, UnsafePtr: unsafe.Pointer(&Constructor_Data_Map_Internal_IterNode[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](x_2), gopurs_runtime.Value{Type: 9, IntVal: 2509360378, UnsafePtr: unsafe.Pointer(nil)}})})
})
}

func Call_Data_Map_Internal_eqMapIter(dictEq_0_loop gopurs_runtime.Value, dictEq1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictEq_0 gopurs_runtime.Value = dictEq_0_loop
_ = dictEq_0
var dictEq1_1 gopurs_runtime.Value = dictEq1_1_loop
_ = dictEq1_1
var go__go_2_0_10 gopurs_runtime.Value
go__go_2_0_10 = gopurs_runtime.Func(func(a_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var a_3_loop gopurs_runtime.Value = a_3_loop_val
var b_4_loop gopurs_runtime.Value = b_4_loop_val
go__go_2_0_10:
for {
if false { continue go__go_2_0_10 }
var a_3 gopurs_runtime.Value = a_3_loop
_ = a_3
var b_4 gopurs_runtime.Value = b_4_loop
_ = b_4
// TAST (Let): v_5_1 -> *Constructor_Data_Map_Internal_IterNext[gopurs_runtime.Value, gopurs_runtime.Value]
v_5_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_IterNext[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Map_Internal_stepAsc(), a_3))
_ = v_5_1
var __t4 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 953589075, UnsafePtr: unsafe.Pointer(v_5_1)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 953589075, UnsafePtr: unsafe.Pointer(v_5_1)}.IntVal == 953589075 && gopurs_runtime.Value{Type: 9, IntVal: 953589075, UnsafePtr: unsafe.Pointer(v_5_1)}.UnsafePtr != nil) {
// TAST (Let): v2_6_2 -> *Constructor_Data_Map_Internal_IterNext[gopurs_runtime.Value, gopurs_runtime.Value]
v2_6_2 := gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_IterNext[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Map_Internal_stepAsc(), b_4))
_ = v2_6_2
var __t3 bool
{
if ((gopurs_runtime.Value{Type: 9, IntVal: 953589075, UnsafePtr: unsafe.Pointer(v2_6_2)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 953589075, UnsafePtr: unsafe.Pointer(v2_6_2)}.IntVal == 953589075 && gopurs_runtime.Value{Type: 9, IntVal: 953589075, UnsafePtr: unsafe.Pointer(v2_6_2)}.UnsafePtr != nil)) && (((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq_0, "eq"), (*Constructor_Data_Map_Internal_IterNext[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 953589075, UnsafePtr: unsafe.Pointer(v_5_1)}.UnsafePtr).V0, (*Constructor_Data_Map_Internal_IterNext[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 953589075, UnsafePtr: unsafe.Pointer(v2_6_2)}.UnsafePtr).V0).IntVal) != (0)) && ((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq1_1, "eq"), (*Constructor_Data_Map_Internal_IterNext[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 953589075, UnsafePtr: unsafe.Pointer(v_5_1)}.UnsafePtr).V1, (*Constructor_Data_Map_Internal_IterNext[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 953589075, UnsafePtr: unsafe.Pointer(v2_6_2)}.UnsafePtr).V1).IntVal) != (0))) {
a_3_loop = (*Constructor_Data_Map_Internal_IterNext[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 953589075, UnsafePtr: unsafe.Pointer(v_5_1)}.UnsafePtr).V2
b_4_loop = (*Constructor_Data_Map_Internal_IterNext[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 953589075, UnsafePtr: unsafe.Pointer(v2_6_2)}.UnsafePtr).V2
continue go__go_2_0_10
__t3 = (gopurs_runtime.Value{}.IntVal) != (0)
goto end_branch_3
} else {

}
}
{
__t3 = false
}
end_branch_3:
__t4 = gopurs_runtime.Bool(__t3)
goto end_branch_4
} else {

}
}
{
if (gopurs_runtime.Value{Type: 9, IntVal: 953589075, UnsafePtr: unsafe.Pointer(v_5_1)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 953589075, UnsafePtr: unsafe.Pointer(v_5_1)}.IntVal == 953589075 && gopurs_runtime.Value{Type: 9, IntVal: 953589075, UnsafePtr: unsafe.Pointer(v_5_1)}.UnsafePtr == nil) {
__t4 = gopurs_runtime.Bool(true)
goto end_branch_4
} else {

}
}
{
__t4 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_4:
return gopurs_runtime.Bool((__t4.IntVal) != (0))
}
}()
})
})
return gopurs_runtime.RecordDict1("eq", go__go_2_0_10)
}

func Call_Data_Map_Internal_ordMapIter(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
// TAST (Let): eqMapIter1_1_0 -> gopurs_runtime.Value
eqMapIter1_1_0 := gopurs_runtime.Apply(Get_Data_Map_Internal_eqMapIter(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd_0, "Eq0"), gopurs_runtime.Value{}))
_ = eqMapIter1_1_0
return gopurs_runtime.Func(func(dictOrd1_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): eqMapIter2_3_1 -> gopurs_runtime.Value
eqMapIter2_3_1 := gopurs_runtime.Apply(eqMapIter1_1_0, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd1_2, "Eq0"), gopurs_runtime.Value{}))
_ = eqMapIter2_3_1
var go__go_4_2_11 gopurs_runtime.Value
go__go_4_2_11 = gopurs_runtime.Func(func(a_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var a_5_loop gopurs_runtime.Value = a_5_loop_val
var b_6_loop gopurs_runtime.Value = b_6_loop_val
go__go_4_2_11:
for {
if false { continue go__go_4_2_11 }
var a_5 gopurs_runtime.Value = a_5_loop
_ = a_5
var b_6 gopurs_runtime.Value = b_6_loop
_ = b_6
// TAST (Let): v_7_3 -> *Constructor_Data_Map_Internal_IterNext[gopurs_runtime.Value, gopurs_runtime.Value]
v_7_3 := gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_IterNext[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Map_Internal_stepAsc(), b_6))
_ = v_7_3
// TAST (Let): v1_8_4 -> *Constructor_Data_Map_Internal_IterNext[gopurs_runtime.Value, gopurs_runtime.Value]
v1_8_4 := gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_IterNext[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Map_Internal_stepAsc(), a_5))
_ = v1_8_4
var __t11 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 953589075, UnsafePtr: unsafe.Pointer(v1_8_4)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 953589075, UnsafePtr: unsafe.Pointer(v1_8_4)}.IntVal == 953589075 && gopurs_runtime.Value{Type: 9, IntVal: 953589075, UnsafePtr: unsafe.Pointer(v1_8_4)}.UnsafePtr != nil) {
var __t9 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 953589075, UnsafePtr: unsafe.Pointer(v_7_3)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 953589075, UnsafePtr: unsafe.Pointer(v_7_3)}.IntVal == 953589075 && gopurs_runtime.Value{Type: 9, IntVal: 953589075, UnsafePtr: unsafe.Pointer(v_7_3)}.UnsafePtr != nil) {
// TAST (Let): v3_9_5 -> gopurs_runtime.Value
v3_9_5 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictOrd_0, "compare"), (*Constructor_Data_Map_Internal_IterNext[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 953589075, UnsafePtr: unsafe.Pointer(v1_8_4)}.UnsafePtr).V0, (*Constructor_Data_Map_Internal_IterNext[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 953589075, UnsafePtr: unsafe.Pointer(v_7_3)}.UnsafePtr).V0)
_ = v3_9_5
var __t8 uint32
{
if (uint32(v3_9_5.IntVal) == 902936544) {
// TAST (Let): v4_10_6 -> gopurs_runtime.Value
v4_10_6 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictOrd1_2, "compare"), (*Constructor_Data_Map_Internal_IterNext[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 953589075, UnsafePtr: unsafe.Pointer(v1_8_4)}.UnsafePtr).V1, (*Constructor_Data_Map_Internal_IterNext[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 953589075, UnsafePtr: unsafe.Pointer(v_7_3)}.UnsafePtr).V1)
_ = v4_10_6
var __t7 uint32
{
if (uint32(v4_10_6.IntVal) == 902936544) {
a_5_loop = (*Constructor_Data_Map_Internal_IterNext[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 953589075, UnsafePtr: unsafe.Pointer(v1_8_4)}.UnsafePtr).V2
b_6_loop = (*Constructor_Data_Map_Internal_IterNext[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 953589075, UnsafePtr: unsafe.Pointer(v_7_3)}.UnsafePtr).V2
continue go__go_4_2_11
__t7 = uint32(gopurs_runtime.Value{}.IntVal)
goto end_branch_7
} else {

}
}
{
__t7 = uint32(v4_10_6.IntVal)
}
end_branch_7:
__t8 = __t7
goto end_branch_8
} else {

}
}
{
__t8 = uint32(v3_9_5.IntVal)
}
end_branch_8:
__t9 = gopurs_runtime.Value{Type: 9, IntVal: int64(__t8), UnsafePtr: nil}
goto end_branch_9
} else {

}
}
{
if (gopurs_runtime.Value{Type: 9, IntVal: 953589075, UnsafePtr: unsafe.Pointer(v_7_3)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 953589075, UnsafePtr: unsafe.Pointer(v_7_3)}.IntVal == 953589075 && gopurs_runtime.Value{Type: 9, IntVal: 953589075, UnsafePtr: unsafe.Pointer(v_7_3)}.UnsafePtr == nil) {
__t9 = gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}
goto end_branch_9
} else {

}
}
{
__t9 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_9:
__t11 = __t9
goto end_branch_11
} else {

}
}
{
if (gopurs_runtime.Value{Type: 9, IntVal: 953589075, UnsafePtr: unsafe.Pointer(v1_8_4)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 953589075, UnsafePtr: unsafe.Pointer(v1_8_4)}.IntVal == 953589075 && gopurs_runtime.Value{Type: 9, IntVal: 953589075, UnsafePtr: unsafe.Pointer(v1_8_4)}.UnsafePtr == nil) {
var __t10 uint32
{
if (gopurs_runtime.Value{Type: 9, IntVal: 953589075, UnsafePtr: unsafe.Pointer(v_7_3)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 953589075, UnsafePtr: unsafe.Pointer(v_7_3)}.IntVal == 953589075 && gopurs_runtime.Value{Type: 9, IntVal: 953589075, UnsafePtr: unsafe.Pointer(v_7_3)}.UnsafePtr == nil) {
__t10 = 902936544
goto end_branch_10
} else {

}
}
{
__t10 = 1527465420
}
end_branch_10:
__t11 = gopurs_runtime.Value{Type: 9, IntVal: int64(__t10), UnsafePtr: nil}
goto end_branch_11
} else {

}
}
{
if (gopurs_runtime.Value{Type: 9, IntVal: 953589075, UnsafePtr: unsafe.Pointer(v_7_3)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 953589075, UnsafePtr: unsafe.Pointer(v_7_3)}.IntVal == 953589075 && gopurs_runtime.Value{Type: 9, IntVal: 953589075, UnsafePtr: unsafe.Pointer(v_7_3)}.UnsafePtr == nil) {
__t11 = gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}
goto end_branch_11
} else {

}
}
{
__t11 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_11:
return gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(__t11.IntVal)), UnsafePtr: nil}
}
}()
})
})
return gopurs_runtime.RecordDict2("Eq0", "compare", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return eqMapIter2_3_1
}), go__go_4_2_11)
})
}

func Call_Data_Map_Internal_toUnfoldable(dictUnfoldable_0_loop *Constructor_Data_Unfoldable_Unfoldable[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictUnfoldable_0 *Constructor_Data_Unfoldable_Unfoldable[gopurs_runtime.Value] = dictUnfoldable_0_loop
_ = dictUnfoldable_0
// TAST (Let): __local_var_1_0 -> gopurs_runtime.Value
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.Box(dictUnfoldable_0.V1), Get_Data_Map_Internal_stepUnfoldr())
_ = __local_var_1_0
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_1_0, gopurs_runtime.Value{Type: 9, IntVal: 2861335956, UnsafePtr: unsafe.Pointer(&Constructor_Data_Map_Internal_IterNode[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](x_2), gopurs_runtime.Value{Type: 9, IntVal: 2509360378, UnsafePtr: unsafe.Pointer(nil)}})})
})
}

func Call_Data_Map_Internal_showMap(dictShow_0_loop gopurs_runtime.Value, dictShow1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictShow_0 gopurs_runtime.Value = dictShow_0_loop
_ = dictShow_0
var dictShow1_1 gopurs_runtime.Value = dictShow1_1_loop
_ = dictShow1_1
// TAST (Let): showArray_2_0 -> *Constructor_Data_Show_Show[[]*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]
showArray_2_0 := &Constructor_Data_Show_Show[[]*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]{1, gopurs_runtime.Apply(Get_Data_Show_showArrayImpl(), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str((((("(Tuple ") + (gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictShow_0, "show"), (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V0).StrVal())) + (" ")) + (gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictShow1_1, "show"), (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V1).StrVal())) + (")"))
}))}
_ = showArray_2_0
return gopurs_runtime.RecordDict1("show", gopurs_runtime.Func(func(as_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str((("(fromFoldable ") + (gopurs_runtime.Apply(gopurs_runtime.Box(showArray_2_0.V0), func() gopurs_runtime.Value {
					arr := func() []*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value] {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_Unfoldable_unfoldableArray(), "unfoldr"), Get_Data_Map_Internal_stepUnfoldr(), gopurs_runtime.Value{Type: 9, IntVal: 2861335956, UnsafePtr: unsafe.Pointer(&Constructor_Data_Map_Internal_IterNode[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](as_3), gopurs_runtime.Value{Type: 9, IntVal: 2509360378, UnsafePtr: unsafe.Pointer(nil)}})}).UnsafePtr)
					unboxed := make([]*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value], len(arr))
					for i, v := range arr { unboxed[i] = gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](v) }
					return unboxed
				}()
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v)} }
					return gopurs_runtime.Array(boxed)
				}()).StrVal())) + (")"))
}))
}

func Call_Data_Map_Internal_isSubmap(dictOrd_0_loop *Constructor_Data_Ord_Ord[gopurs_runtime.Value], dictEq_1_loop *Constructor_Data_Eq_Eq[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictOrd_0 *Constructor_Data_Ord_Ord[gopurs_runtime.Value] = dictOrd_0_loop
_ = dictOrd_0
var dictEq_1 *Constructor_Data_Eq_Eq[gopurs_runtime.Value] = dictEq_1_loop
_ = dictEq_1
var go__go_2_0_12 gopurs_runtime.Value
_ = go__go_2_0_12
go__go_2_0_12 = gopurs_runtime.Func(func(m1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m2_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t8 gopurs_runtime.Value
{
if (m1_3.Type == 9 && m1_3.IntVal == 324739070 && m1_3.UnsafePtr == nil) {
__t8 = gopurs_runtime.Bool(true)
goto end_branch_8
} else {

}
}
{
if (m1_3.Type == 9 && m1_3.IntVal == 324739070 && m1_3.UnsafePtr != nil) {
// TAST (Let): __local_var_5_1 -> gopurs_runtime.Value
__local_var_5_1 := (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(m1_3.UnsafePtr).V2
_ = __local_var_5_1
var go__go_6_2_13 gopurs_runtime.Value
go__go_6_2_13 = gopurs_runtime.Func(func(v_7_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_7_loop gopurs_runtime.Value = v_7_loop_val
go__go_6_2_13:
for {
if false { continue go__go_6_2_13 }
var v_7 gopurs_runtime.Value = v_7_loop
_ = v_7
var __t5 gopurs_runtime.Value
{
if (v_7.Type == 9 && v_7.IntVal == 324739070 && v_7.UnsafePtr == nil) {
__t5 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil))}
goto end_branch_5
} else {

}
}
{
if (v_7.Type == 9 && v_7.IntVal == 324739070 && v_7.UnsafePtr != nil) {
// TAST (Let): v1_8_3 -> gopurs_runtime.Value
v1_8_3 := gopurs_runtime.Apply2(gopurs_runtime.Box(dictOrd_0.V1), __local_var_5_1, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_7.UnsafePtr).V2)
_ = v1_8_3
var __t4 gopurs_runtime.Value
{
if (uint32(v1_8_3.IntVal) == 1527465420) {
v_7_loop = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_7.UnsafePtr).V4)}
continue go__go_6_2_13
__t4 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Value{}))}
goto end_branch_4
} else {

}
}
{
if (uint32(v1_8_3.IntVal) == 380165415) {
v_7_loop = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_7.UnsafePtr).V5)}
continue go__go_6_2_13
__t4 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Value{}))}
goto end_branch_4
} else {

}
}
{
if (uint32(v1_8_3.IntVal) == 902936544) {
__t4 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{1, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_7.UnsafePtr).V3})}
goto end_branch_4
} else {

}
}
{
__t4 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_4:
__t5 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](__t4))}
goto end_branch_5
} else {

}
}
{
__t5 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_5:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](__t5))}
}
}()
})
// TAST (Let): v1_7_6 -> gopurs_runtime.Value
v1_7_6 := gopurs_runtime.Apply(go__go_6_2_13, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](m2_4))})
_ = v1_7_6
var __t7 gopurs_runtime.Value
{
if (v1_7_6.Type == 9 && v1_7_6.IntVal == 930809136 && v1_7_6.UnsafePtr == nil) {
__t7 = gopurs_runtime.Bool(false)
goto end_branch_7
} else {

}
}
{
if (v1_7_6.Type == 9 && v1_7_6.IntVal == 930809136 && v1_7_6.UnsafePtr != nil) {
__t7 = gopurs_runtime.Bool(((gopurs_runtime.Apply2(gopurs_runtime.Box(dictEq_1.V0), (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(m1_3.UnsafePtr).V3, (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(v1_7_6.UnsafePtr).V0).IntVal) != (0)) && (((gopurs_runtime.Apply2(go__go_2_0_12, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(m1_3.UnsafePtr).V4)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](m2_4))}).IntVal) != (0)) && ((gopurs_runtime.Apply2(go__go_2_0_12, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(m1_3.UnsafePtr).V5)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](m2_4))}).IntVal) != (0))))
goto end_branch_7
} else {

}
}
{
__t7 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_7:
__t8 = gopurs_runtime.Bool((__t7.IntVal) != (0))
goto end_branch_8
} else {

}
}
{
__t8 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_8:
return gopurs_runtime.Bool((__t8.IntVal) != (0))
})
})
return go__go_2_0_12
}

func Call_Data_Map_Internal_isEmpty(v_0_loop *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]) bool {
var v_0 *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value] = v_0_loop
_ = v_0
var __t0 bool
{
if (gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_0)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_0)}.IntVal == 324739070 && gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr == nil) {
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

func Call_Data_Map_Internal_intersectionWith(dictOrd_0_loop *Constructor_Data_Ord_Ord[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictOrd_0 *Constructor_Data_Ord_Ord[gopurs_runtime.Value] = dictOrd_0_loop
_ = dictOrd_0
// TAST (Let): compare_1_0 -> gopurs_runtime.Value
compare_1_0 := gopurs_runtime.Box(dictOrd_0.V1)
_ = compare_1_0
return gopurs_runtime.Func(func(app_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m2_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeIntersectionWith(), compare_1_0, app_2, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](m1_3))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](m2_4))})))}
})
})
})
}

func Call_Data_Map_Internal_intersection(dictOrd_0_loop *Constructor_Data_Ord_Ord[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictOrd_0 *Constructor_Data_Ord_Ord[gopurs_runtime.Value] = dictOrd_0_loop
_ = dictOrd_0
// TAST (Let): compare_1_0 -> gopurs_runtime.Value
compare_1_0 := gopurs_runtime.Box(dictOrd_0.V1)
_ = compare_1_0
return gopurs_runtime.Func(func(m1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m2_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeIntersectionWith(), compare_1_0, Get_Data_Function_go__const(), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](m1_2))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](m2_3))})))}
})
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
var go__go_4_0_14 gopurs_runtime.Value
_ = go__go_4_0_14
go__go_4_0_14 = gopurs_runtime.Func(func(v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
var __t3 gopurs_runtime.Value
{
if (v1_5.Type == 9 && v1_5.IntVal == 324739070 && v1_5.UnsafePtr == nil) {
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(&Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]{1, 1, 1, k_2, v_3, gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(nil))}), gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(nil))})})}
goto end_branch_3
} else {

}
}
{
if (v1_5.Type == 9 && v1_5.IntVal == 324739070 && v1_5.UnsafePtr != nil) {
// TAST (Let): v2_6_1 -> gopurs_runtime.Value
v2_6_1 := gopurs_runtime.Apply2(gopurs_runtime.Box(dictOrd_0.V1), k_2, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v1_5.UnsafePtr).V2)
_ = v2_6_1
var __t2 gopurs_runtime.Value
{
if (uint32(v2_6_1.IntVal) == 1527465420) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeBalancedNode(), (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v1_5.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v1_5.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(go__go_4_0_14, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v1_5.UnsafePtr).V4)})))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v1_5.UnsafePtr).V5)})))}
goto end_branch_2
} else {

}
}
{
if (uint32(v2_6_1.IntVal) == 380165415) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeBalancedNode(), (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v1_5.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v1_5.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v1_5.UnsafePtr).V4)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(go__go_4_0_14, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v1_5.UnsafePtr).V5)})))})))}
goto end_branch_2
} else {

}
}
{
if (uint32(v2_6_1.IntVal) == 902936544) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(&Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v1_5.UnsafePtr).V0, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v1_5.UnsafePtr).V1, k_2, gopurs_runtime.Apply2(app_1, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v1_5.UnsafePtr).V3, v_3), (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v1_5.UnsafePtr).V4, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v1_5.UnsafePtr).V5})}
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_2:
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__t2))}
goto end_branch_3
} else {

}
}
{
__t3 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_3:
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__t3))}
})
return go__go_4_0_14
}

func Call_Data_Map_Internal_insert(dictOrd_0_loop *Constructor_Data_Ord_Ord[gopurs_runtime.Value], k_1_loop gopurs_runtime.Value, v_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 *Constructor_Data_Ord_Ord[gopurs_runtime.Value] = dictOrd_0_loop
_ = dictOrd_0
var k_1 gopurs_runtime.Value = k_1_loop
_ = k_1
var v_2 gopurs_runtime.Value = v_2_loop
_ = v_2
var go__go_3_0_15 gopurs_runtime.Value
_ = go__go_3_0_15
go__go_3_0_15 = gopurs_runtime.Func(func(v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t3 gopurs_runtime.Value
{
if (v1_4.Type == 9 && v1_4.IntVal == 324739070 && v1_4.UnsafePtr == nil) {
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(&Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]{1, 1, 1, k_1, v_2, gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(nil))}), gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(nil))})})}
goto end_branch_3
} else {

}
}
{
if (v1_4.Type == 9 && v1_4.IntVal == 324739070 && v1_4.UnsafePtr != nil) {
// TAST (Let): v2_5_1 -> gopurs_runtime.Value
v2_5_1 := gopurs_runtime.Apply2(gopurs_runtime.Box(dictOrd_0.V1), k_1, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v1_4.UnsafePtr).V2)
_ = v2_5_1
var __t2 gopurs_runtime.Value
{
if (uint32(v2_5_1.IntVal) == 1527465420) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeBalancedNode(), (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v1_4.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v1_4.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(go__go_3_0_15, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v1_4.UnsafePtr).V4)})))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v1_4.UnsafePtr).V5)})))}
goto end_branch_2
} else {

}
}
{
if (uint32(v2_5_1.IntVal) == 380165415) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeBalancedNode(), (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v1_4.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v1_4.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v1_4.UnsafePtr).V4)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(go__go_3_0_15, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v1_4.UnsafePtr).V5)})))})))}
goto end_branch_2
} else {

}
}
{
if (uint32(v2_5_1.IntVal) == 902936544) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(&Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v1_4.UnsafePtr).V0, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v1_4.UnsafePtr).V1, k_1, v_2, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v1_4.UnsafePtr).V4, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v1_4.UnsafePtr).V5})}
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_2:
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__t2))}
goto end_branch_3
} else {

}
}
{
__t3 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_3:
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__t3))}
})
return go__go_3_0_15
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
var __t4 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(kmin_3)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(kmin_3)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(kmin_3)}.UnsafePtr != nil) {
// TAST (Let): __local_var_6_1 -> gopurs_runtime.Value
__local_var_6_1 := (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(kmin_3)}.UnsafePtr).V0
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
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(kmin_3)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(kmin_3)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(kmin_3)}.UnsafePtr == nil) {
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
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(kmax_4)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(kmax_4)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(kmax_4)}.UnsafePtr != nil) {
// TAST (Let): __local_var_7_6 -> gopurs_runtime.Value
__local_var_7_6 := (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(kmax_4)}.UnsafePtr).V0
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
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(kmax_4)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(kmax_4)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(kmax_4)}.UnsafePtr == nil) {
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
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(kmin_3)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(kmin_3)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(kmin_3)}.UnsafePtr != nil) {
var __t21 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(kmax_4)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(kmax_4)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(kmax_4)}.UnsafePtr != nil) {
// TAST (Let): __local_var_8_11 -> gopurs_runtime.Value
__local_var_8_11 := (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(kmax_4)}.UnsafePtr).V0
_ = __local_var_8_11
// TAST (Let): __local_var_9_12 -> gopurs_runtime.Value
__local_var_9_12 := (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(kmin_3)}.UnsafePtr).V0
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
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(kmax_4)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(kmax_4)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(kmax_4)}.UnsafePtr == nil) {
// TAST (Let): __local_var_8_18 -> gopurs_runtime.Value
__local_var_8_18 := (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(kmin_3)}.UnsafePtr).V0
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
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(kmin_3)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(kmin_3)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(kmin_3)}.UnsafePtr == nil) {
var __t25 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(kmax_4)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(kmax_4)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(kmax_4)}.UnsafePtr != nil) {
// TAST (Let): __local_var_8_22 -> gopurs_runtime.Value
__local_var_8_22 := (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(kmax_4)}.UnsafePtr).V0
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
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(kmax_4)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(kmax_4)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(kmax_4)}.UnsafePtr == nil) {
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
var go__go_9_27_26 gopurs_runtime.Value
_ = go__go_9_27_26
go__go_9_27_26 = gopurs_runtime.Func(func(v_10 gopurs_runtime.Value) gopurs_runtime.Value {
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
if (gopurs_runtime.Apply(tooSmall_6_0, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_10.UnsafePtr).V2).IntVal) != (0) {
__t28 = memptyValue_2
goto end_branch_28
} else {

}
}
{
__t28 = gopurs_runtime.Apply(go__go_9_27_26, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_10.UnsafePtr).V4)})
}
end_branch_28:
var __t29 gopurs_runtime.Value
{
if (gopurs_runtime.Apply(inBounds_8_10, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_10.UnsafePtr).V2).IntVal) != (0) {
__t29 = gopurs_runtime.Apply2(f_5, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_10.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_10.UnsafePtr).V3)
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
if (gopurs_runtime.Apply(tooLarge_7_5, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_10.UnsafePtr).V2).IntVal) != (0) {
__t30 = memptyValue_2
goto end_branch_30
} else {

}
}
{
__t30 = gopurs_runtime.Apply(go__go_9_27_26, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_10.UnsafePtr).V5)})
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
return go__go_9_27_26
}

func Call_Data_Map_Internal_foldSubmap(dictOrd_0_loop *Constructor_Data_Ord_Ord[gopurs_runtime.Value], dictMonoid_1_loop *Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictOrd_0 *Constructor_Data_Ord_Ord[gopurs_runtime.Value] = dictOrd_0_loop
_ = dictOrd_0
var dictMonoid_1 *Constructor_Data_Monoid_Monoid[gopurs_runtime.Value] = dictMonoid_1_loop
_ = dictMonoid_1
return gopurs_runtime.Apply3(Get_Data_Map_Internal_foldSubmapBy(), gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(dictOrd_0)}, gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.Box(dictMonoid_1.V0), gopurs_runtime.Value{}), "append"), gopurs_runtime.Box(dictMonoid_1.V1))
}

func Call_Data_Map_Internal_findMin(v_0_loop *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]) *Constructor_Data_Maybe_Just[gopurs_runtime.Value] {
findMin:
for {
if false { continue findMin }
var v_0 *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value] = v_0_loop
_ = v_0
var __t2 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_0)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_0)}.IntVal == 324739070 && gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr == nil) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil))}
goto end_branch_2
} else {

}
}
{
if (gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_0)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_0)}.IntVal == 324739070 && gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr != nil) {
var __t1 gopurs_runtime.Value
{
var __t_tag_0 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V4)}
if (__t_tag_0.Type == 9 && __t_tag_0.IntVal == 324739070 && __t_tag_0.UnsafePtr == nil) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{1, gopurs_runtime.RecordDict2("key", "value", (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V3)})}
goto end_branch_1
} else {

}
}
{
v_0_loop = (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V4
continue findMin
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Value{}))}
}
end_branch_1:
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](__t1))}
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_2:
return gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](__t2)
}
}

func Call_Data_Map_Internal_lookupGT(dictOrd_0_loop *Constructor_Data_Ord_Ord[gopurs_runtime.Value], k_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 *Constructor_Data_Ord_Ord[gopurs_runtime.Value] = dictOrd_0_loop
_ = dictOrd_0
var k_1 gopurs_runtime.Value = k_1_loop
_ = k_1
var go__go_2_0_27 gopurs_runtime.Value
_ = go__go_2_0_27
go__go_2_0_27 = gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t5 gopurs_runtime.Value
{
if (v_3.Type == 9 && v_3.IntVal == 324739070 && v_3.UnsafePtr == nil) {
__t5 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil))}
goto end_branch_5
} else {

}
}
{
if (v_3.Type == 9 && v_3.IntVal == 324739070 && v_3.UnsafePtr != nil) {
// TAST (Let): v1_4_1 -> gopurs_runtime.Value
v1_4_1 := gopurs_runtime.Apply2(gopurs_runtime.Box(dictOrd_0.V1), k_1, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V2)
_ = v1_4_1
var __t4 gopurs_runtime.Value
{
if (uint32(v1_4_1.IntVal) == 1527465420) {
// TAST (Let): v2_5_2 -> *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
v2_5_2 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(go__go_2_0_27, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V4)}))
_ = v2_5_2
var __t3 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_5_2)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_5_2)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_5_2)}.UnsafePtr == nil) {
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{1, gopurs_runtime.RecordDict2("key", "value", (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V3)})}
goto end_branch_3
} else {

}
}
{
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_5_2)}
}
end_branch_3:
__t4 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](__t3))}
goto end_branch_4
} else {

}
}
{
if (uint32(v1_4_1.IntVal) == 380165415) {
__t4 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(go__go_2_0_27, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V5)})))}
goto end_branch_4
} else {

}
}
{
if (uint32(v1_4_1.IntVal) == 902936544) {
__t4 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_Data_Map_Internal_findMin((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V5))}))}
goto end_branch_4
} else {

}
}
{
__t4 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_4:
__t5 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](__t4))}
goto end_branch_5
} else {

}
}
{
__t5 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_5:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](__t5))}
})
return go__go_2_0_27
}

func Call_Data_Map_Internal_findMax(v_0_loop *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]) *Constructor_Data_Maybe_Just[gopurs_runtime.Value] {
findMax:
for {
if false { continue findMax }
var v_0 *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value] = v_0_loop
_ = v_0
var __t2 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_0)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_0)}.IntVal == 324739070 && gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr == nil) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil))}
goto end_branch_2
} else {

}
}
{
if (gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_0)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_0)}.IntVal == 324739070 && gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr != nil) {
var __t1 gopurs_runtime.Value
{
var __t_tag_0 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V5)}
if (__t_tag_0.Type == 9 && __t_tag_0.IntVal == 324739070 && __t_tag_0.UnsafePtr == nil) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{1, gopurs_runtime.RecordDict2("key", "value", (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V3)})}
goto end_branch_1
} else {

}
}
{
v_0_loop = (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V5
continue findMax
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Value{}))}
}
end_branch_1:
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](__t1))}
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_2:
return gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](__t2)
}
}

func Call_Data_Map_Internal_lookupLT(dictOrd_0_loop *Constructor_Data_Ord_Ord[gopurs_runtime.Value], k_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 *Constructor_Data_Ord_Ord[gopurs_runtime.Value] = dictOrd_0_loop
_ = dictOrd_0
var k_1 gopurs_runtime.Value = k_1_loop
_ = k_1
var go__go_2_0_28 gopurs_runtime.Value
_ = go__go_2_0_28
go__go_2_0_28 = gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t5 gopurs_runtime.Value
{
if (v_3.Type == 9 && v_3.IntVal == 324739070 && v_3.UnsafePtr == nil) {
__t5 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil))}
goto end_branch_5
} else {

}
}
{
if (v_3.Type == 9 && v_3.IntVal == 324739070 && v_3.UnsafePtr != nil) {
// TAST (Let): v1_4_1 -> gopurs_runtime.Value
v1_4_1 := gopurs_runtime.Apply2(gopurs_runtime.Box(dictOrd_0.V1), k_1, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V2)
_ = v1_4_1
var __t4 gopurs_runtime.Value
{
if (uint32(v1_4_1.IntVal) == 1527465420) {
__t4 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(go__go_2_0_28, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V4)})))}
goto end_branch_4
} else {

}
}
{
if (uint32(v1_4_1.IntVal) == 380165415) {
// TAST (Let): v2_5_2 -> *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
v2_5_2 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(go__go_2_0_28, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V5)}))
_ = v2_5_2
var __t3 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_5_2)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_5_2)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_5_2)}.UnsafePtr == nil) {
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{1, gopurs_runtime.RecordDict2("key", "value", (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V3)})}
goto end_branch_3
} else {

}
}
{
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_5_2)}
}
end_branch_3:
__t4 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](__t3))}
goto end_branch_4
} else {

}
}
{
if (uint32(v1_4_1.IntVal) == 902936544) {
__t4 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_Data_Map_Internal_findMax((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V4))}))}
goto end_branch_4
} else {

}
}
{
__t4 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_4:
__t5 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](__t4))}
goto end_branch_5
} else {

}
}
{
__t5 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_5:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](__t5))}
})
return go__go_2_0_28
}

func Call_Data_Map_Internal_filterWithKey(dictOrd_0_loop *Constructor_Data_Ord_Ord[gopurs_runtime.Value], f_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 *Constructor_Data_Ord_Ord[gopurs_runtime.Value] = dictOrd_0_loop
_ = dictOrd_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var go__go_2_0_29 gopurs_runtime.Value
_ = go__go_2_0_29
go__go_2_0_29 = gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 gopurs_runtime.Value
{
if (v_3.Type == 9 && v_3.IntVal == 324739070 && v_3.UnsafePtr == nil) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(nil))}
goto end_branch_2
} else {

}
}
{
if (v_3.Type == 9 && v_3.IntVal == 324739070 && v_3.UnsafePtr != nil) {
var __t1 *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]
{
if (gopurs_runtime.Apply2(f_1, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V3).IntVal) != (0) {
__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeBalancedNode(), (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(go__go_2_0_29, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V4)})))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(go__go_2_0_29, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V5)})))}))
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp2(Get_Data_Map_Internal_unsafeJoinNodes(), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(go__go_2_0_29, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V4)})))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(go__go_2_0_29, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V5)})))}))
}
end_branch_1:
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(__t1)}
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_2:
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__t2))}
})
return go__go_2_0_29
}

func Call_Data_Map_Internal_filterKeys(dictOrd_0_loop *Constructor_Data_Ord_Ord[gopurs_runtime.Value], f_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 *Constructor_Data_Ord_Ord[gopurs_runtime.Value] = dictOrd_0_loop
_ = dictOrd_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var go__go_2_0_30 gopurs_runtime.Value
_ = go__go_2_0_30
go__go_2_0_30 = gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 gopurs_runtime.Value
{
if (v_3.Type == 9 && v_3.IntVal == 324739070 && v_3.UnsafePtr == nil) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(nil))}
goto end_branch_2
} else {

}
}
{
if (v_3.Type == 9 && v_3.IntVal == 324739070 && v_3.UnsafePtr != nil) {
var __t1 *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]
{
if (gopurs_runtime.Apply(f_1, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V2).IntVal) != (0) {
__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeBalancedNode(), (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(go__go_2_0_30, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V4)})))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(go__go_2_0_30, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V5)})))}))
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp2(Get_Data_Map_Internal_unsafeJoinNodes(), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(go__go_2_0_30, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V4)})))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(go__go_2_0_30, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V5)})))}))
}
end_branch_1:
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(__t1)}
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_2:
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__t2))}
})
return go__go_2_0_30
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
// TAST (Let): eqMapIter2_2_0 -> *Constructor_Data_Eq_Eq[gopurs_runtime.Value]
eqMapIter2_2_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Call_Data_Map_Internal_eqMapIter(dictEq_0, dictEq1_1))
_ = eqMapIter2_2_0
return gopurs_runtime.RecordDict1("eq", gopurs_runtime.Func(func(xs_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(ys_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t3 gopurs_runtime.Value
{
if (xs_3.Type == 9 && xs_3.IntVal == 324739070 && xs_3.UnsafePtr == nil) {
var __t1 bool
{
if (ys_4.Type == 9 && ys_4.IntVal == 324739070 && ys_4.UnsafePtr == nil) {
__t1 = true
goto end_branch_1
} else {

}
}
{
__t1 = false
}
end_branch_1:
__t3 = gopurs_runtime.Bool(__t1)
goto end_branch_3
} else {

}
}
{
if (xs_3.Type == 9 && xs_3.IntVal == 324739070 && xs_3.UnsafePtr != nil) {
var __t2 bool
{
if ((ys_4.Type == 9 && ys_4.IntVal == 324739070 && ys_4.UnsafePtr != nil)) && (((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(xs_3.UnsafePtr).V1) == ((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(ys_4.UnsafePtr).V1)) {
__t2 = (gopurs_runtime.Apply2(gopurs_runtime.Box(eqMapIter2_2_0.V0), gopurs_runtime.Value{Type: 9, IntVal: 2861335956, UnsafePtr: unsafe.Pointer(&Constructor_Data_Map_Internal_IterNode[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](xs_3), gopurs_runtime.Value{Type: 9, IntVal: 2509360378, UnsafePtr: unsafe.Pointer(nil)}})}, gopurs_runtime.Value{Type: 9, IntVal: 2861335956, UnsafePtr: unsafe.Pointer(&Constructor_Data_Map_Internal_IterNode[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](ys_4), gopurs_runtime.Value{Type: 9, IntVal: 2509360378, UnsafePtr: unsafe.Pointer(nil)}})}).IntVal) != (0)
goto end_branch_2
} else {

}
}
{
__t2 = false
}
end_branch_2:
__t3 = gopurs_runtime.Bool(__t2)
goto end_branch_3
} else {

}
}
{
__t3 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_3:
return gopurs_runtime.Bool((__t3.IntVal) != (0))
})
}))
}

func Call_Data_Map_Internal_ordMap(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
// TAST (Let): ordMapIter1_1_0 -> gopurs_runtime.Value
ordMapIter1_1_0 := Call_Data_Map_Internal_ordMapIter(dictOrd_0)
_ = ordMapIter1_1_0
// TAST (Let): eqMap1_2_1 -> gopurs_runtime.Value
eqMap1_2_1 := gopurs_runtime.Apply(Get_Data_Map_Internal_eqMap(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd_0, "Eq0"), gopurs_runtime.Value{}))
_ = eqMap1_2_1
return gopurs_runtime.Func(func(dictOrd1_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): ordMapIter2_4_2 -> *Constructor_Data_Ord_Ord[gopurs_runtime.Value]
ordMapIter2_4_2 := gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](gopurs_runtime.Apply(ordMapIter1_1_0, dictOrd1_3))
_ = ordMapIter2_4_2
// TAST (Let): eqMap2_5_3 -> gopurs_runtime.Value
eqMap2_5_3 := gopurs_runtime.Apply(eqMap1_2_1, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd1_3, "Eq0"), gopurs_runtime.Value{}))
_ = eqMap2_5_3
return gopurs_runtime.RecordDict2("Eq0", "compare", gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return eqMap2_5_3
}), gopurs_runtime.Func(func(xs_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(ys_7 gopurs_runtime.Value) gopurs_runtime.Value {
var __t6 uint32
{
if (xs_6.Type == 9 && xs_6.IntVal == 324739070 && xs_6.UnsafePtr == nil) {
var __t5 uint32
{
if (ys_7.Type == 9 && ys_7.IntVal == 324739070 && ys_7.UnsafePtr == nil) {
__t5 = 902936544
goto end_branch_5
} else {

}
}
{
__t5 = 1527465420
}
end_branch_5:
__t6 = __t5
goto end_branch_6
} else {

}
}
{
var __t4 uint32
{
if (ys_7.Type == 9 && ys_7.IntVal == 324739070 && ys_7.UnsafePtr == nil) {
__t4 = 380165415
goto end_branch_4
} else {

}
}
{
__t4 = uint32(gopurs_runtime.Apply2(gopurs_runtime.Box(ordMapIter2_4_2.V1), gopurs_runtime.Value{Type: 9, IntVal: 2861335956, UnsafePtr: unsafe.Pointer(&Constructor_Data_Map_Internal_IterNode[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](xs_6), gopurs_runtime.Value{Type: 9, IntVal: 2509360378, UnsafePtr: unsafe.Pointer(nil)}})}, gopurs_runtime.Value{Type: 9, IntVal: 2861335956, UnsafePtr: unsafe.Pointer(&Constructor_Data_Map_Internal_IterNode[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](ys_7), gopurs_runtime.Value{Type: 9, IntVal: 2509360378, UnsafePtr: unsafe.Pointer(nil)}})}).IntVal)
}
end_branch_4:
__t6 = __t4
}
end_branch_6:
return gopurs_runtime.Value{Type: 9, IntVal: int64(__t6), UnsafePtr: nil}
})
}))
})
}

func Call_Data_Map_Internal_eq1Map(dictEq_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictEq_0 gopurs_runtime.Value = dictEq_0_loop
_ = dictEq_0
return gopurs_runtime.RecordDict1("eq1", gopurs_runtime.Func(func(dictEq1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(Call_Data_Map_Internal_eqMap(dictEq_0, dictEq1_1), "eq")
}))
}

func Call_Data_Map_Internal_ord1Map(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
// TAST (Let): ordMap1_1_0 -> gopurs_runtime.Value
ordMap1_1_0 := Call_Data_Map_Internal_ordMap(dictOrd_0)
_ = ordMap1_1_0
// TAST (Let): __local_var_2_2 -> gopurs_runtime.Value
__local_var_2_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd_0, "Eq0"), gopurs_runtime.Value{})
_ = __local_var_2_2
// TAST (Let): eq1Map1_2_1 -> gopurs_runtime.Value
eq1Map1_2_1 := gopurs_runtime.RecordDict1("eq1", gopurs_runtime.Func(func(dictEq1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(Call_Data_Map_Internal_eqMap(__local_var_2_2, dictEq1_3), "eq")
}))
_ = eq1Map1_2_1
return gopurs_runtime.RecordDict2("Eq10", "compare1", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return eq1Map1_2_1
}), gopurs_runtime.Func(func(dictOrd1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(gopurs_runtime.Apply(ordMap1_1_0, dictOrd1_3), "compare")
}))
}

func Call_Data_Map_Internal_fromFoldable(dictOrd_0_loop *Constructor_Data_Ord_Ord[gopurs_runtime.Value], dictFoldable_1_loop *Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictOrd_0 *Constructor_Data_Ord_Ord[gopurs_runtime.Value] = dictOrd_0_loop
_ = dictOrd_0
var dictFoldable_1 *Constructor_Data_Foldable_Foldable[gopurs_runtime.Value] = dictFoldable_1_loop
_ = dictFoldable_1
return gopurs_runtime.Apply2(gopurs_runtime.Box(dictFoldable_1.V1), gopurs_runtime.Func(func(m_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(Call_Data_Map_Internal_insert(dictOrd_0, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V0, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V1), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](m_2))})))}
})
}), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(nil))})
}

func Call_Data_Map_Internal_fromFoldableWith(dictOrd_0_loop *Constructor_Data_Ord_Ord[gopurs_runtime.Value], dictFoldable_1_loop *Constructor_Data_Foldable_Foldable[gopurs_runtime.Value], f_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 *Constructor_Data_Ord_Ord[gopurs_runtime.Value] = dictOrd_0_loop
_ = dictOrd_0
var dictFoldable_1 *Constructor_Data_Foldable_Foldable[gopurs_runtime.Value] = dictFoldable_1_loop
_ = dictFoldable_1
var f_2 gopurs_runtime.Value = f_2_loop
_ = f_2
return gopurs_runtime.Apply2(gopurs_runtime.Box(dictFoldable_1.V1), gopurs_runtime.Func(func(m_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(Call_Data_Map_Internal_insertWith(dictOrd_0, gopurs_runtime.Func(func(b_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_2, a_6, b_5)
})
}), (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V0, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V1), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](m_3))})))}
})
}), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(nil))})
}

func Call_Data_Map_Internal_fromFoldableWithIndex(dictOrd_0_loop *Constructor_Data_Ord_Ord[gopurs_runtime.Value], dictFoldableWithIndex_1_loop *Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var dictOrd_0 *Constructor_Data_Ord_Ord[gopurs_runtime.Value] = dictOrd_0_loop
_ = dictOrd_0
var dictFoldableWithIndex_1 *Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value] = dictFoldableWithIndex_1_loop
_ = dictFoldableWithIndex_1
return gopurs_runtime.Apply2(gopurs_runtime.Box(dictFoldableWithIndex_1.V2), gopurs_runtime.Func(func(k_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(Call_Data_Map_Internal_insert(dictOrd_0, k_2, v_4), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](m_3))})))}
})
})
}), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(nil))})
}

func Call_Data_Map_Internal_monoidSemigroupMap(_dollar__unused_0_loop gopurs_runtime.Value, dictOrd_1_loop gopurs_runtime.Value, dictSemigroup_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var _dollar__unused_0 gopurs_runtime.Value = _dollar__unused_0_loop
_ = _dollar__unused_0
var dictOrd_1 gopurs_runtime.Value = dictOrd_1_loop
_ = dictOrd_1
var dictSemigroup_2 gopurs_runtime.Value = dictSemigroup_2_loop
_ = dictSemigroup_2
// TAST (Let): semigroupMap3_3_0 -> gopurs_runtime.Value
semigroupMap3_3_0 := Call_Data_Map_Internal_semigroupMap(gopurs_runtime.Value{}, dictOrd_1, dictSemigroup_2)
_ = semigroupMap3_3_0
return gopurs_runtime.RecordDict2("Semigroup0", "mempty", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return semigroupMap3_3_0
}), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(nil))})
}

func Call_Data_Map_Internal_submap(dictOrd_0_loop *Constructor_Data_Ord_Ord[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictOrd_0 *Constructor_Data_Ord_Ord[gopurs_runtime.Value] = dictOrd_0_loop
_ = dictOrd_0
// TAST (Let): compare_1_1 -> gopurs_runtime.Value
compare_1_1 := gopurs_runtime.Box(dictOrd_0.V1)
_ = compare_1_1
// TAST (Let): union1_1_0 -> gopurs_runtime.Value
union1_1_0 := gopurs_runtime.Func(func(m1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m2_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeUnionWith(), compare_1_1, Get_Data_Function_go__const(), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](m1_2))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](m2_3))})))}
})
})
_ = union1_1_0
return gopurs_runtime.Func(func(kmin_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(kmax_3 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Map_Internal_foldSubmapBy(dictOrd_0, union1_1_0, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(nil))}, gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](kmin_2), gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](kmax_3), Get_Data_Map_Internal_singleton())
})
})
}

func Call_Data_Map_Internal_unions(dictOrd_0_loop *Constructor_Data_Ord_Ord[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictOrd_0 *Constructor_Data_Ord_Ord[gopurs_runtime.Value] = dictOrd_0_loop
_ = dictOrd_0
// TAST (Let): compare_1_1 -> gopurs_runtime.Value
compare_1_1 := gopurs_runtime.Box(dictOrd_0.V1)
_ = compare_1_1
// TAST (Let): union1_1_0 -> gopurs_runtime.Value
union1_1_0 := gopurs_runtime.Func(func(m1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m2_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeUnionWith(), compare_1_1, Get_Data_Function_go__const(), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](m1_2))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](m2_3))})))}
})
})
_ = union1_1_0
return gopurs_runtime.Func(func(dictFoldable_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFoldable_2, "foldl"), union1_1_0, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(nil))})
})
}

func Call_Data_Map_Internal_difference(dictOrd_0_loop *Constructor_Data_Ord_Ord[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictOrd_0 *Constructor_Data_Ord_Ord[gopurs_runtime.Value] = dictOrd_0_loop
_ = dictOrd_0
// TAST (Let): compare_1_0 -> gopurs_runtime.Value
compare_1_0 := gopurs_runtime.Box(dictOrd_0.V1)
_ = compare_1_0
return gopurs_runtime.Func(func(m1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m2_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp3(Get_Data_Map_Internal_unsafeDifference(), compare_1_0, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](m1_2))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](m2_3))})))}
})
})
}

func Call_Data_Map_Internal_delete(dictOrd_0_loop *Constructor_Data_Ord_Ord[gopurs_runtime.Value], k_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 *Constructor_Data_Ord_Ord[gopurs_runtime.Value] = dictOrd_0_loop
_ = dictOrd_0
var k_1 gopurs_runtime.Value = k_1_loop
_ = k_1
var go__go_2_0_31 gopurs_runtime.Value
_ = go__go_2_0_31
go__go_2_0_31 = gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t3 gopurs_runtime.Value
{
if (v_3.Type == 9 && v_3.IntVal == 324739070 && v_3.UnsafePtr == nil) {
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(nil))}
goto end_branch_3
} else {

}
}
{
if (v_3.Type == 9 && v_3.IntVal == 324739070 && v_3.UnsafePtr != nil) {
// TAST (Let): v1_4_1 -> gopurs_runtime.Value
v1_4_1 := gopurs_runtime.Apply2(gopurs_runtime.Box(dictOrd_0.V1), k_1, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V2)
_ = v1_4_1
var __t2 gopurs_runtime.Value
{
if (uint32(v1_4_1.IntVal) == 1527465420) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeBalancedNode(), (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(go__go_2_0_31, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V4)})))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V5)})))}
goto end_branch_2
} else {

}
}
{
if (uint32(v1_4_1.IntVal) == 380165415) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeBalancedNode(), (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V4)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(go__go_2_0_31, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V5)})))})))}
goto end_branch_2
} else {

}
}
{
if (uint32(v1_4_1.IntVal) == 902936544) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp2(Get_Data_Map_Internal_unsafeJoinNodes(), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V4)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V5)})))}
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_2:
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__t2))}
goto end_branch_3
} else {

}
}
{
__t3 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_3:
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__t3))}
})
return go__go_2_0_31
}

func Call_Data_Map_Internal_checkValid(dictOrd_0_loop *Constructor_Data_Ord_Ord[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictOrd_0 *Constructor_Data_Ord_Ord[gopurs_runtime.Value] = dictOrd_0_loop
_ = dictOrd_0
var go__go_1_0_32 gopurs_runtime.Value
_ = go__go_1_0_32
go__go_1_0_32 = gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t33 gopurs_runtime.Value
{
if (v_2.Type == 9 && v_2.IntVal == 324739070 && v_2.UnsafePtr == nil) {
__t33 = gopurs_runtime.Bool(true)
goto end_branch_33
} else {

}
}
{
if (v_2.Type == 9 && v_2.IntVal == 324739070 && v_2.UnsafePtr != nil) {
var __t32 gopurs_runtime.Value
{
var __t_tag_1 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V4)}
if (__t_tag_1.Type == 9 && __t_tag_1.IntVal == 324739070 && __t_tag_1.UnsafePtr == nil) {
var __t10 gopurs_runtime.Value
{
var __t_tag_2 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V5)}
if (__t_tag_2.Type == 9 && __t_tag_2.IntVal == 324739070 && __t_tag_2.UnsafePtr == nil) {
__t10 = gopurs_runtime.Bool(true)
goto end_branch_10
} else {

}
}
{
var __t_tag_3 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V5)}
if (__t_tag_3.Type == 9 && __t_tag_3.IntVal == 324739070 && __t_tag_3.UnsafePtr != nil) {
var __t_and_9 bool = false
if ((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V0) == (2) {

var __t_and_8 bool = false
if ((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V5)}.UnsafePtr).V0) == (1) {

var __t4 bool
{
if ((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V1) > ((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V5)}.UnsafePtr).V1) {
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
var __t_tag_5 gopurs_runtime.Value = gopurs_runtime.Apply2(gopurs_runtime.Box(dictOrd_0.V1), (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V5)}.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V2)
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
__t_and_7 = (__t6) && ((gopurs_runtime.Apply(go__go_1_0_32, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V5)}).IntVal) != (0))
}
__t_and_8 = __t_and_7
}
__t_and_9 = __t_and_8
}
__t10 = gopurs_runtime.Bool(__t_and_9)
goto end_branch_10
} else {

}
}
{
__t10 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_10:
__t32 = gopurs_runtime.Bool((__t10.IntVal) != (0))
goto end_branch_32
} else {

}
}
{
var __t_tag_11 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V4)}
if (__t_tag_11.Type == 9 && __t_tag_11.IntVal == 324739070 && __t_tag_11.UnsafePtr != nil) {
var __t31 gopurs_runtime.Value
{
var __t_tag_12 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V5)}
if (__t_tag_12.Type == 9 && __t_tag_12.IntVal == 324739070 && __t_tag_12.UnsafePtr == nil) {
var __t_and_18 bool = false
if ((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V0) == (2) {

var __t_and_17 bool = false
if ((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V4)}.UnsafePtr).V0) == (1) {

var __t13 bool
{
if ((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V1) > ((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V4)}.UnsafePtr).V1) {
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
var __t_tag_14 gopurs_runtime.Value = gopurs_runtime.Apply2(gopurs_runtime.Box(dictOrd_0.V1), (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V4)}.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V2)
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
__t_and_16 = (__t15) && ((gopurs_runtime.Apply(go__go_1_0_32, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V4)}).IntVal) != (0))
}
__t_and_17 = __t_and_16
}
__t_and_18 = __t_and_17
}
__t31 = gopurs_runtime.Bool(__t_and_18)
goto end_branch_31
} else {

}
}
{
var __t_tag_19 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V5)}
if (__t_tag_19.Type == 9 && __t_tag_19.IntVal == 324739070 && __t_tag_19.UnsafePtr != nil) {
var __t20 bool
{
if ((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V0) > ((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V5)}.UnsafePtr).V0) {
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
var __t_tag_21 gopurs_runtime.Value = gopurs_runtime.Apply2(gopurs_runtime.Box(dictOrd_0.V1), (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V5)}.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V2)
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
if ((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V0) > ((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V4)}.UnsafePtr).V0) {
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
var __t_tag_24 gopurs_runtime.Value = gopurs_runtime.Apply2(gopurs_runtime.Box(dictOrd_0.V1), (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V4)}.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V2)
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
if (gopurs_runtime.Apply(Get_Data_Ord_abs__1599282999(), gopurs_runtime.Int(((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V5)}.UnsafePtr).V0) - ((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V4)}.UnsafePtr).V0))).IntVal) < (2) {
__t26 = true
goto end_branch_26
} else {

}
}
{
__t26 = false
}
end_branch_26:
__t_and_27 = (__t26) && ((((((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V5)}.UnsafePtr).V1) + ((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V4)}.UnsafePtr).V1)) + (1)) == ((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V1)) && (((gopurs_runtime.Apply(go__go_1_0_32, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V4)}).IntVal) != (0)) && ((gopurs_runtime.Apply(go__go_1_0_32, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V5)}).IntVal) != (0))))
}
__t_and_28 = __t_and_27
}
__t_and_29 = __t_and_28
}
__t_and_30 = __t_and_29
}
__t31 = gopurs_runtime.Bool(__t_and_30)
goto end_branch_31
} else {

}
}
{
__t31 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_31:
__t32 = gopurs_runtime.Bool((__t31.IntVal) != (0))
goto end_branch_32
} else {

}
}
{
__t32 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_32:
__t33 = gopurs_runtime.Bool((__t32.IntVal) != (0))
goto end_branch_33
} else {

}
}
{
__t33 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_33:
return gopurs_runtime.Bool((__t33.IntVal) != (0))
})
return go__go_1_0_32
}

func Call_Data_Map_Internal_catMaybes(dictOrd_0_loop *Constructor_Data_Ord_Ord[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictOrd_0 *Constructor_Data_Ord_Ord[gopurs_runtime.Value] = dictOrd_0_loop
_ = dictOrd_0
return Call_Data_Map_Internal_mapMaybeWithKey(dictOrd_0, gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Map_Internal_identity1()
}))
}

func Call_Data_Map_Internal_applyMap(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
// TAST (Let): compare_1_0 -> gopurs_runtime.Value
compare_1_0 := gopurs_runtime.RecordGet(dictOrd_0, "compare")
_ = compare_1_0
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Map_Internal_functorMap()
}), gopurs_runtime.Func(func(m1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m2_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeIntersectionWith(), compare_1_0, Get_Data_Map_Internal_identity2(), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](m1_2))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](m2_3))})))}
})
}))
}

func Call_Data_Map_Internal_bindMap(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
// TAST (Let): compare_1_1 -> gopurs_runtime.Value
compare_1_1 := gopurs_runtime.RecordGet(dictOrd_0, "compare")
_ = compare_1_1
// TAST (Let): applyMap1_1_0 -> gopurs_runtime.Value
applyMap1_1_0 := gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Map_Internal_functorMap()
}), gopurs_runtime.Func(func(m1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m2_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeIntersectionWith(), compare_1_1, Get_Data_Map_Internal_identity2(), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](m1_2))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](m2_3))})))}
})
}))
_ = applyMap1_1_0
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return applyMap1_1_0
}), gopurs_runtime.Func(func(m_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(Call_Data_Map_Internal_mapMaybeWithKey(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](dictOrd_0), gopurs_runtime.Func(func(k_4 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_5_3_33 gopurs_runtime.Value
go__go_5_3_33 = gopurs_runtime.Func(func(v_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_6_loop gopurs_runtime.Value = v_6_loop_val
go__go_5_3_33:
for {
if false { continue go__go_5_3_33 }
var v_6 gopurs_runtime.Value = v_6_loop
_ = v_6
var __t6 gopurs_runtime.Value
{
if (v_6.Type == 9 && v_6.IntVal == 324739070 && v_6.UnsafePtr == nil) {
__t6 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil))}
goto end_branch_6
} else {

}
}
{
if (v_6.Type == 9 && v_6.IntVal == 324739070 && v_6.UnsafePtr != nil) {
// TAST (Let): v1_7_4 -> gopurs_runtime.Value
v1_7_4 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictOrd_0, "compare"), k_4, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_6.UnsafePtr).V2)
_ = v1_7_4
var __t5 gopurs_runtime.Value
{
if (uint32(v1_7_4.IntVal) == 1527465420) {
v_6_loop = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_6.UnsafePtr).V4)}
continue go__go_5_3_33
__t5 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Value{}))}
goto end_branch_5
} else {

}
}
{
if (uint32(v1_7_4.IntVal) == 380165415) {
v_6_loop = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_6.UnsafePtr).V5)}
continue go__go_5_3_33
__t5 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Value{}))}
goto end_branch_5
} else {

}
}
{
if (uint32(v1_7_4.IntVal) == 902936544) {
__t5 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{1, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_6.UnsafePtr).V3})}
goto end_branch_5
} else {

}
}
{
__t5 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_5:
__t6 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](__t5))}
goto end_branch_6
} else {

}
}
{
__t6 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_6:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](__t6))}
}
}()
})
// TAST (Let): __local_var_5_2 -> gopurs_runtime.Value
__local_var_5_2 := go__go_5_3_33
_ = __local_var_5_2
return gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_5_2, gopurs_runtime.Apply(f_3, x_6))
})
})), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](m_2))})))}
})
}))
}

func Call_Data_Map_Internal_anyWithKey(predicate_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var predicate_0 gopurs_runtime.Value = predicate_0_loop
_ = predicate_0
var go__go_1_0_34 gopurs_runtime.Value
_ = go__go_1_0_34
go__go_1_0_34 = gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if (v_2.Type == 9 && v_2.IntVal == 324739070 && v_2.UnsafePtr == nil) {
__t1 = gopurs_runtime.Bool(false)
goto end_branch_1
} else {

}
}
{
if (v_2.Type == 9 && v_2.IntVal == 324739070 && v_2.UnsafePtr != nil) {
__t1 = gopurs_runtime.Bool(((gopurs_runtime.Apply2(predicate_0, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V3).IntVal) != (0)) || (((gopurs_runtime.Apply(go__go_1_0_34, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V4)}).IntVal) != (0)) || ((gopurs_runtime.Apply(go__go_1_0_34, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V5)}).IntVal) != (0))))
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
return gopurs_runtime.Bool((__t1.IntVal) != (0))
})
return go__go_1_0_34
}

func Call_Data_Map_Internal_any(predicate_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var predicate_0 gopurs_runtime.Value = predicate_0_loop
_ = predicate_0
var go__go_1_0_35 gopurs_runtime.Value
_ = go__go_1_0_35
go__go_1_0_35 = gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if (v_2.Type == 9 && v_2.IntVal == 324739070 && v_2.UnsafePtr == nil) {
__t1 = gopurs_runtime.Bool(false)
goto end_branch_1
} else {

}
}
{
if (v_2.Type == 9 && v_2.IntVal == 324739070 && v_2.UnsafePtr != nil) {
__t1 = gopurs_runtime.Bool(((gopurs_runtime.Apply(predicate_0, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V3).IntVal) != (0)) || (((gopurs_runtime.Apply(go__go_1_0_35, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V4)}).IntVal) != (0)) || ((gopurs_runtime.Apply(go__go_1_0_35, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V5)}).IntVal) != (0))))
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
return gopurs_runtime.Bool((__t1.IntVal) != (0))
})
return go__go_1_0_35
}

func Call_Data_Map_Internal_alter(dictOrd_0_loop *Constructor_Data_Ord_Ord[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictOrd_0 *Constructor_Data_Ord_Ord[gopurs_runtime.Value] = dictOrd_0_loop
_ = dictOrd_0
// TAST (Let): compare_1_0 -> gopurs_runtime.Value
compare_1_0 := gopurs_runtime.Box(dictOrd_0.V1)
_ = compare_1_0
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): v_5_1 -> gopurs_runtime.Value
v_5_1 := gopurs_runtime.UncurriedApp3(Get_Data_Map_Internal_unsafeSplit(), compare_1_0, k_3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](m_4))})
_ = v_5_1
// TAST (Let): v2_6_2 -> *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
v2_6_2 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(f_2, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Split[gopurs_runtime.Value, gopurs_runtime.Value])(v_5_1.UnsafePtr).V0)}))
_ = v2_6_2
var __t3 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_6_2)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_6_2)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_6_2)}.UnsafePtr == nil) {
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp2(Get_Data_Map_Internal_unsafeJoinNodes(), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Split[gopurs_runtime.Value, gopurs_runtime.Value])(v_5_1.UnsafePtr).V1)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Split[gopurs_runtime.Value, gopurs_runtime.Value])(v_5_1.UnsafePtr).V2)})))}
goto end_branch_3
} else {

}
}
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_6_2)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_6_2)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_6_2)}.UnsafePtr != nil) {
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeBalancedNode(), k_3, (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_6_2)}.UnsafePtr).V0, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Split[gopurs_runtime.Value, gopurs_runtime.Value])(v_5_1.UnsafePtr).V1)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Split[gopurs_runtime.Value, gopurs_runtime.Value])(v_5_1.UnsafePtr).V2)})))}
goto end_branch_3
} else {

}
}
{
__t3 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_3:
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__t3))}
})
})
})
}

func Call_Data_Map_Internal_altMap(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
// TAST (Let): compare_1_0 -> gopurs_runtime.Value
compare_1_0 := gopurs_runtime.RecordGet(dictOrd_0, "compare")
_ = compare_1_0
return gopurs_runtime.RecordDict2("Functor0", "alt", gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Map_Internal_functorMap()
}), gopurs_runtime.Func(func(m1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m2_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeUnionWith(), compare_1_0, Get_Data_Function_go__const(), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](m1_2))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](m2_3))})))}
})
}))
}

func Call_Data_Map_Internal_plusMap(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
// TAST (Let): compare_1_1 -> gopurs_runtime.Value
compare_1_1 := gopurs_runtime.RecordGet(dictOrd_0, "compare")
_ = compare_1_1
// TAST (Let): altMap1_1_0 -> gopurs_runtime.Value
altMap1_1_0 := gopurs_runtime.RecordDict2("Functor0", "alt", gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Map_Internal_functorMap()
}), gopurs_runtime.Func(func(m1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m2_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeUnionWith(), compare_1_1, Get_Data_Function_go__const(), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](m1_2))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](m2_3))})))}
})
}))
_ = altMap1_1_0
return gopurs_runtime.RecordDict2("Alt0", "empty", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return altMap1_1_0
}), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(nil))})
}

func Call_Data_Map_Internal_alter__2325420954(dictOrd_0_loop *Constructor_Data_Ord_Ord[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictOrd_0 *Constructor_Data_Ord_Ord[gopurs_runtime.Value] = dictOrd_0_loop
_ = dictOrd_0
// TAST (Let): compare_1_0 -> gopurs_runtime.Value
compare_1_0 := gopurs_runtime.Box(dictOrd_0.V1)
_ = compare_1_0
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): v_5_1 -> gopurs_runtime.Value
v_5_1 := gopurs_runtime.UncurriedApp3(Get_Data_Map_Internal_unsafeSplit(), compare_1_0, k_3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](m_4))})
_ = v_5_1
// TAST (Let): v2_6_2 -> *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
v2_6_2 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(f_2, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Split[gopurs_runtime.Value, gopurs_runtime.Value])(v_5_1.UnsafePtr).V0)}))
_ = v2_6_2
var __t3 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_6_2)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_6_2)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_6_2)}.UnsafePtr == nil) {
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp2(Get_Data_Map_Internal_unsafeJoinNodes(), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Split[gopurs_runtime.Value, gopurs_runtime.Value])(v_5_1.UnsafePtr).V1)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Split[gopurs_runtime.Value, gopurs_runtime.Value])(v_5_1.UnsafePtr).V2)})))}
goto end_branch_3
} else {

}
}
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_6_2)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_6_2)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_6_2)}.UnsafePtr != nil) {
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeBalancedNode(), k_3, (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_6_2)}.UnsafePtr).V0, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Split[gopurs_runtime.Value, gopurs_runtime.Value])(v_5_1.UnsafePtr).V1)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Split[gopurs_runtime.Value, gopurs_runtime.Value])(v_5_1.UnsafePtr).V2)})))}
goto end_branch_3
} else {

}
}
{
__t3 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_3:
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__t3))}
})
})
})
}

func Call_Data_Map_Internal_alter__1204655226(dictOrd_0_loop *Constructor_Data_Ord_Ord[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictOrd_0 *Constructor_Data_Ord_Ord[gopurs_runtime.Value] = dictOrd_0_loop
_ = dictOrd_0
// TAST (Let): compare_1_0 -> gopurs_runtime.Value
compare_1_0 := gopurs_runtime.Box(dictOrd_0.V1)
_ = compare_1_0
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): v_5_1 -> gopurs_runtime.Value
v_5_1 := gopurs_runtime.UncurriedApp3(Get_Data_Map_Internal_unsafeSplit(), compare_1_0, k_3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](m_4))})
_ = v_5_1
// TAST (Let): v2_6_2 -> *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
v2_6_2 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(f_2, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Split[gopurs_runtime.Value, gopurs_runtime.Value])(v_5_1.UnsafePtr).V0)}))
_ = v2_6_2
var __t3 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_6_2)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_6_2)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_6_2)}.UnsafePtr == nil) {
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp2(Get_Data_Map_Internal_unsafeJoinNodes(), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Split[gopurs_runtime.Value, gopurs_runtime.Value])(v_5_1.UnsafePtr).V1)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Split[gopurs_runtime.Value, gopurs_runtime.Value])(v_5_1.UnsafePtr).V2)})))}
goto end_branch_3
} else {

}
}
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_6_2)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_6_2)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_6_2)}.UnsafePtr != nil) {
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeBalancedNode(), k_3, (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_6_2)}.UnsafePtr).V0, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Split[gopurs_runtime.Value, gopurs_runtime.Value])(v_5_1.UnsafePtr).V1)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Split[gopurs_runtime.Value, gopurs_runtime.Value])(v_5_1.UnsafePtr).V2)})))}
goto end_branch_3
} else {

}
}
{
__t3 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_3:
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__t3))}
})
})
})
}

func Call_Data_Map_Internal_findMax__2266220649(v_0_loop *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]) *Constructor_Data_Maybe_Just[gopurs_runtime.Value] {
var v_0 *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value] = v_0_loop
_ = v_0
var __t2 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_0)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_0)}.IntVal == 324739070 && gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr == nil) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil))}
goto end_branch_2
} else {

}
}
{
if (gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_0)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_0)}.IntVal == 324739070 && gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr != nil) {
var __t1 gopurs_runtime.Value
{
var __t_tag_0 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V5)}
if (__t_tag_0.Type == 9 && __t_tag_0.IntVal == 324739070 && __t_tag_0.UnsafePtr == nil) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{1, gopurs_runtime.RecordDict2("key", "value", (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V3)})}
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_Data_Map_Internal_findMax((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V5))}))}
}
end_branch_1:
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](__t1))}
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_2:
return gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](__t2)
}

func Call_Data_Map_Internal_findMax__528468393(v_0_loop *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]) *Constructor_Data_Maybe_Just[gopurs_runtime.Value] {
var v_0 *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value] = v_0_loop
_ = v_0
var __t2 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_0)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_0)}.IntVal == 324739070 && gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr == nil) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil))}
goto end_branch_2
} else {

}
}
{
if (gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_0)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_0)}.IntVal == 324739070 && gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr != nil) {
var __t1 gopurs_runtime.Value
{
var __t_tag_0 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V5)}
if (__t_tag_0.Type == 9 && __t_tag_0.IntVal == 324739070 && __t_tag_0.UnsafePtr == nil) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{1, gopurs_runtime.RecordDict2("key", "value", (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V3)})}
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_Data_Map_Internal_findMax((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V5))}))}
}
end_branch_1:
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](__t1))}
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_2:
return gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](__t2)
}

func Call_Data_Map_Internal_findMin__2266220649(v_0_loop *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]) *Constructor_Data_Maybe_Just[gopurs_runtime.Value] {
var v_0 *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value] = v_0_loop
_ = v_0
var __t2 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_0)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_0)}.IntVal == 324739070 && gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr == nil) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil))}
goto end_branch_2
} else {

}
}
{
if (gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_0)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_0)}.IntVal == 324739070 && gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr != nil) {
var __t1 gopurs_runtime.Value
{
var __t_tag_0 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V4)}
if (__t_tag_0.Type == 9 && __t_tag_0.IntVal == 324739070 && __t_tag_0.UnsafePtr == nil) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{1, gopurs_runtime.RecordDict2("key", "value", (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V3)})}
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_Data_Map_Internal_findMin((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V4))}))}
}
end_branch_1:
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](__t1))}
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_2:
return gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](__t2)
}

func Call_Data_Map_Internal_findMin__528468393(v_0_loop *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]) *Constructor_Data_Maybe_Just[gopurs_runtime.Value] {
var v_0 *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value] = v_0_loop
_ = v_0
var __t2 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_0)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_0)}.IntVal == 324739070 && gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr == nil) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil))}
goto end_branch_2
} else {

}
}
{
if (gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_0)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_0)}.IntVal == 324739070 && gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr != nil) {
var __t1 gopurs_runtime.Value
{
var __t_tag_0 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V4)}
if (__t_tag_0.Type == 9 && __t_tag_0.IntVal == 324739070 && __t_tag_0.UnsafePtr == nil) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{1, gopurs_runtime.RecordDict2("key", "value", (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V3)})}
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_Data_Map_Internal_findMin((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V4))}))}
}
end_branch_1:
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](__t1))}
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_2:
return gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](__t2)
}

func Call_Data_Map_Internal_foldSubmapBy__3050108409(dictOrd_0_loop *Constructor_Data_Ord_Ord[gopurs_runtime.Value], appendFn_1_loop gopurs_runtime.Value, memptyValue_2_loop gopurs_runtime.Value, kmin_3_loop *Constructor_Data_Maybe_Just[gopurs_runtime.Value], kmax_4_loop *Constructor_Data_Maybe_Just[gopurs_runtime.Value], f_5_loop gopurs_runtime.Value) gopurs_runtime.Value {
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
var __t4 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(kmin_3)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(kmin_3)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(kmin_3)}.UnsafePtr != nil) {
// TAST (Let): __local_var_6_1 -> gopurs_runtime.Value
__local_var_6_1 := (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(kmin_3)}.UnsafePtr).V0
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
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(kmin_3)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(kmin_3)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(kmin_3)}.UnsafePtr == nil) {
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
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(kmax_4)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(kmax_4)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(kmax_4)}.UnsafePtr != nil) {
// TAST (Let): __local_var_7_6 -> gopurs_runtime.Value
__local_var_7_6 := (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(kmax_4)}.UnsafePtr).V0
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
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(kmax_4)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(kmax_4)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(kmax_4)}.UnsafePtr == nil) {
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
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(kmin_3)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(kmin_3)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(kmin_3)}.UnsafePtr != nil) {
var __t21 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(kmax_4)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(kmax_4)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(kmax_4)}.UnsafePtr != nil) {
// TAST (Let): __local_var_8_11 -> gopurs_runtime.Value
__local_var_8_11 := (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(kmax_4)}.UnsafePtr).V0
_ = __local_var_8_11
// TAST (Let): __local_var_9_12 -> gopurs_runtime.Value
__local_var_9_12 := (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(kmin_3)}.UnsafePtr).V0
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
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(kmax_4)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(kmax_4)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(kmax_4)}.UnsafePtr == nil) {
// TAST (Let): __local_var_8_18 -> gopurs_runtime.Value
__local_var_8_18 := (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(kmin_3)}.UnsafePtr).V0
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
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(kmin_3)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(kmin_3)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(kmin_3)}.UnsafePtr == nil) {
var __t25 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(kmax_4)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(kmax_4)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(kmax_4)}.UnsafePtr != nil) {
// TAST (Let): __local_var_8_22 -> gopurs_runtime.Value
__local_var_8_22 := (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(kmax_4)}.UnsafePtr).V0
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
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(kmax_4)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(kmax_4)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(kmax_4)}.UnsafePtr == nil) {
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
var go__go_9_27_36 gopurs_runtime.Value
_ = go__go_9_27_36
go__go_9_27_36 = gopurs_runtime.Func(func(v_10 gopurs_runtime.Value) gopurs_runtime.Value {
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
if (gopurs_runtime.Apply(tooSmall_6_0, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_10.UnsafePtr).V2).IntVal) != (0) {
__t28 = memptyValue_2
goto end_branch_28
} else {

}
}
{
__t28 = gopurs_runtime.Apply(go__go_9_27_36, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_10.UnsafePtr).V4)})
}
end_branch_28:
var __t29 gopurs_runtime.Value
{
if (gopurs_runtime.Apply(inBounds_8_10, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_10.UnsafePtr).V2).IntVal) != (0) {
__t29 = gopurs_runtime.Apply2(f_5, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_10.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_10.UnsafePtr).V3)
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
if (gopurs_runtime.Apply(tooLarge_7_5, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_10.UnsafePtr).V2).IntVal) != (0) {
__t30 = memptyValue_2
goto end_branch_30
} else {

}
}
{
__t30 = gopurs_runtime.Apply(go__go_9_27_36, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_10.UnsafePtr).V5)})
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
return go__go_9_27_36
}

func Call_Data_Map_Internal_foldSubmapBy__3128450809(dictOrd_0_loop *Constructor_Data_Ord_Ord[gopurs_runtime.Value], appendFn_1_loop gopurs_runtime.Value, memptyValue_2_loop *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value], kmin_3_loop *Constructor_Data_Maybe_Just[gopurs_runtime.Value], kmax_4_loop *Constructor_Data_Maybe_Just[gopurs_runtime.Value], f_5_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 *Constructor_Data_Ord_Ord[gopurs_runtime.Value] = dictOrd_0_loop
_ = dictOrd_0
var appendFn_1 gopurs_runtime.Value = appendFn_1_loop
_ = appendFn_1
var memptyValue_2 *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value] = memptyValue_2_loop
_ = memptyValue_2
var kmin_3 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = kmin_3_loop
_ = kmin_3
var kmax_4 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = kmax_4_loop
_ = kmax_4
var f_5 gopurs_runtime.Value = f_5_loop
_ = f_5
var __t4 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(kmin_3)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(kmin_3)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(kmin_3)}.UnsafePtr != nil) {
// TAST (Let): __local_var_6_1 -> gopurs_runtime.Value
__local_var_6_1 := (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(kmin_3)}.UnsafePtr).V0
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
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(kmin_3)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(kmin_3)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(kmin_3)}.UnsafePtr == nil) {
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
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(kmax_4)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(kmax_4)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(kmax_4)}.UnsafePtr != nil) {
// TAST (Let): __local_var_7_6 -> gopurs_runtime.Value
__local_var_7_6 := (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(kmax_4)}.UnsafePtr).V0
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
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(kmax_4)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(kmax_4)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(kmax_4)}.UnsafePtr == nil) {
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
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(kmin_3)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(kmin_3)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(kmin_3)}.UnsafePtr != nil) {
var __t21 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(kmax_4)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(kmax_4)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(kmax_4)}.UnsafePtr != nil) {
// TAST (Let): __local_var_8_11 -> gopurs_runtime.Value
__local_var_8_11 := (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(kmax_4)}.UnsafePtr).V0
_ = __local_var_8_11
// TAST (Let): __local_var_9_12 -> gopurs_runtime.Value
__local_var_9_12 := (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(kmin_3)}.UnsafePtr).V0
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
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(kmax_4)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(kmax_4)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(kmax_4)}.UnsafePtr == nil) {
// TAST (Let): __local_var_8_18 -> gopurs_runtime.Value
__local_var_8_18 := (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(kmin_3)}.UnsafePtr).V0
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
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(kmin_3)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(kmin_3)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(kmin_3)}.UnsafePtr == nil) {
var __t25 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(kmax_4)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(kmax_4)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(kmax_4)}.UnsafePtr != nil) {
// TAST (Let): __local_var_8_22 -> gopurs_runtime.Value
__local_var_8_22 := (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(kmax_4)}.UnsafePtr).V0
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
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(kmax_4)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(kmax_4)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(kmax_4)}.UnsafePtr == nil) {
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
var go__go_9_27_37 gopurs_runtime.Value
_ = go__go_9_27_37
go__go_9_27_37 = gopurs_runtime.Func(func(v_10 gopurs_runtime.Value) gopurs_runtime.Value {
var __t31 gopurs_runtime.Value
{
if (v_10.Type == 9 && v_10.IntVal == 324739070 && v_10.UnsafePtr == nil) {
__t31 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(memptyValue_2)}
goto end_branch_31
} else {

}
}
{
if (v_10.Type == 9 && v_10.IntVal == 324739070 && v_10.UnsafePtr != nil) {
var __t28 *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]
{
if (gopurs_runtime.Apply(tooSmall_6_0, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_10.UnsafePtr).V2).IntVal) != (0) {
__t28 = memptyValue_2
goto end_branch_28
} else {

}
}
{
__t28 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(go__go_9_27_37, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_10.UnsafePtr).V4)}))
}
end_branch_28:
var __t29 *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]
{
if (gopurs_runtime.Apply(inBounds_8_10, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_10.UnsafePtr).V2).IntVal) != (0) {
__t29 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply2(f_5, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_10.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_10.UnsafePtr).V3))
goto end_branch_29
} else {

}
}
{
__t29 = memptyValue_2
}
end_branch_29:
var __t30 *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]
{
if (gopurs_runtime.Apply(tooLarge_7_5, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_10.UnsafePtr).V2).IntVal) != (0) {
__t30 = memptyValue_2
goto end_branch_30
} else {

}
}
{
__t30 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(go__go_9_27_37, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_10.UnsafePtr).V5)}))
}
end_branch_30:
__t31 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply2(appendFn_1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply2(appendFn_1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(__t28)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(__t29)})))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(__t30)})))}
goto end_branch_31
} else {

}
}
{
__t31 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_31:
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__t31))}
})
return go__go_9_27_37
}

func Call_Data_Map_Internal_insert__3204212386(dictOrd_0_loop *Constructor_Data_Ord_Ord[gopurs_runtime.Value], k_1_loop gopurs_runtime.Value, v_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 *Constructor_Data_Ord_Ord[gopurs_runtime.Value] = dictOrd_0_loop
_ = dictOrd_0
var k_1 gopurs_runtime.Value = k_1_loop
_ = k_1
var v_2 gopurs_runtime.Value = v_2_loop
_ = v_2
var go__go_3_0_55 gopurs_runtime.Value
_ = go__go_3_0_55
go__go_3_0_55 = gopurs_runtime.Func(func(v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t3 gopurs_runtime.Value
{
if (v1_4.Type == 9 && v1_4.IntVal == 324739070 && v1_4.UnsafePtr == nil) {
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(&Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]{1, 1, 1, k_1, v_2, gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(nil))}), gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(nil))})})}
goto end_branch_3
} else {

}
}
{
if (v1_4.Type == 9 && v1_4.IntVal == 324739070 && v1_4.UnsafePtr != nil) {
// TAST (Let): v2_5_1 -> gopurs_runtime.Value
v2_5_1 := gopurs_runtime.Apply2(gopurs_runtime.Box(dictOrd_0.V1), k_1, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v1_4.UnsafePtr).V2)
_ = v2_5_1
var __t2 gopurs_runtime.Value
{
if (uint32(v2_5_1.IntVal) == 1527465420) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeBalancedNode(), (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v1_4.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v1_4.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(go__go_3_0_55, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v1_4.UnsafePtr).V4)})))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v1_4.UnsafePtr).V5)})))}
goto end_branch_2
} else {

}
}
{
if (uint32(v2_5_1.IntVal) == 380165415) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeBalancedNode(), (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v1_4.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v1_4.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v1_4.UnsafePtr).V4)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(go__go_3_0_55, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v1_4.UnsafePtr).V5)})))})))}
goto end_branch_2
} else {

}
}
{
if (uint32(v2_5_1.IntVal) == 902936544) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(&Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v1_4.UnsafePtr).V0, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v1_4.UnsafePtr).V1, k_1, v_2, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v1_4.UnsafePtr).V4, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v1_4.UnsafePtr).V5})}
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_2:
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__t2))}
goto end_branch_3
} else {

}
}
{
__t3 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_3:
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__t3))}
})
return go__go_3_0_55
}

func Call_Data_Map_Internal_insert__4289641298(dictOrd_0_loop *Constructor_Data_Ord_Ord[gopurs_runtime.Value], k_1_loop gopurs_runtime.Value, v_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 *Constructor_Data_Ord_Ord[gopurs_runtime.Value] = dictOrd_0_loop
_ = dictOrd_0
var k_1 gopurs_runtime.Value = k_1_loop
_ = k_1
var v_2 gopurs_runtime.Value = v_2_loop
_ = v_2
var go__go_3_0_56 gopurs_runtime.Value
_ = go__go_3_0_56
go__go_3_0_56 = gopurs_runtime.Func(func(v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t3 gopurs_runtime.Value
{
if (v1_4.Type == 9 && v1_4.IntVal == 324739070 && v1_4.UnsafePtr == nil) {
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(&Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]{1, 1, 1, k_1, v_2, gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(nil))}), gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(nil))})})}
goto end_branch_3
} else {

}
}
{
if (v1_4.Type == 9 && v1_4.IntVal == 324739070 && v1_4.UnsafePtr != nil) {
// TAST (Let): v2_5_1 -> gopurs_runtime.Value
v2_5_1 := gopurs_runtime.Apply2(gopurs_runtime.Box(dictOrd_0.V1), k_1, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v1_4.UnsafePtr).V2)
_ = v2_5_1
var __t2 gopurs_runtime.Value
{
if (uint32(v2_5_1.IntVal) == 1527465420) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeBalancedNode(), (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v1_4.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v1_4.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(go__go_3_0_56, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v1_4.UnsafePtr).V4)})))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v1_4.UnsafePtr).V5)})))}
goto end_branch_2
} else {

}
}
{
if (uint32(v2_5_1.IntVal) == 380165415) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeBalancedNode(), (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v1_4.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v1_4.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v1_4.UnsafePtr).V4)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(go__go_3_0_56, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v1_4.UnsafePtr).V5)})))})))}
goto end_branch_2
} else {

}
}
{
if (uint32(v2_5_1.IntVal) == 902936544) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(&Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v1_4.UnsafePtr).V0, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v1_4.UnsafePtr).V1, k_1, v_2, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v1_4.UnsafePtr).V4, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v1_4.UnsafePtr).V5})}
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_2:
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__t2))}
goto end_branch_3
} else {

}
}
{
__t3 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_3:
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__t3))}
})
return go__go_3_0_56
}

func Call_Data_Map_Internal_insert__2073142786(dictOrd_0_loop *Constructor_Data_Ord_Ord[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]], k_1_loop *Constructor_Data_Maybe_Just[gopurs_runtime.Value], v_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 *Constructor_Data_Ord_Ord[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]] = dictOrd_0_loop
_ = dictOrd_0
var k_1 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = k_1_loop
_ = k_1
var v_2 gopurs_runtime.Value = v_2_loop
_ = v_2
var go__go_3_0_57 gopurs_runtime.Value
_ = go__go_3_0_57
go__go_3_0_57 = gopurs_runtime.Func(func(v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t3 gopurs_runtime.Value
{
if (v1_4.Type == 9 && v1_4.IntVal == 324739070 && v1_4.UnsafePtr == nil) {
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(&Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]{1, 1, 1, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(k_1)}, v_2, gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(nil))}), gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(nil))})})}
goto end_branch_3
} else {

}
}
{
if (v1_4.Type == 9 && v1_4.IntVal == 324739070 && v1_4.UnsafePtr != nil) {
// TAST (Let): v2_5_1 -> gopurs_runtime.Value
v2_5_1 := gopurs_runtime.Apply2(gopurs_runtime.Box(dictOrd_0.V1), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(k_1)}, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]]((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v1_4.UnsafePtr).V2))})
_ = v2_5_1
var __t2 gopurs_runtime.Value
{
if (uint32(v2_5_1.IntVal) == 1527465420) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[*Constructor_Data_Maybe_Just[gopurs_runtime.Value], gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeBalancedNode(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]]((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v1_4.UnsafePtr).V2))}, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v1_4.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[*Constructor_Data_Maybe_Just[gopurs_runtime.Value], gopurs_runtime.Value]](gopurs_runtime.Apply(go__go_3_0_57, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v1_4.UnsafePtr).V4)})))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v1_4.UnsafePtr).V5)})))}
goto end_branch_2
} else {

}
}
{
if (uint32(v2_5_1.IntVal) == 380165415) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[*Constructor_Data_Maybe_Just[gopurs_runtime.Value], gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeBalancedNode(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]]((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v1_4.UnsafePtr).V2))}, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v1_4.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v1_4.UnsafePtr).V4)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[*Constructor_Data_Maybe_Just[gopurs_runtime.Value], gopurs_runtime.Value]](gopurs_runtime.Apply(go__go_3_0_57, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v1_4.UnsafePtr).V5)})))})))}
goto end_branch_2
} else {

}
}
{
if (uint32(v2_5_1.IntVal) == 902936544) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(&Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v1_4.UnsafePtr).V0, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v1_4.UnsafePtr).V1, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(k_1)}, v_2, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v1_4.UnsafePtr).V4, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v1_4.UnsafePtr).V5})}
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_2:
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[*Constructor_Data_Maybe_Just[gopurs_runtime.Value], gopurs_runtime.Value]](__t2))}
goto end_branch_3
} else {

}
}
{
__t3 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_3:
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[*Constructor_Data_Maybe_Just[gopurs_runtime.Value], gopurs_runtime.Value]](__t3))}
})
return go__go_3_0_57
}

func Call_Data_Map_Internal_insertWith__118979962(dictOrd_0_loop *Constructor_Data_Ord_Ord[gopurs_runtime.Value], app_1_loop gopurs_runtime.Value, k_2_loop gopurs_runtime.Value, v_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 *Constructor_Data_Ord_Ord[gopurs_runtime.Value] = dictOrd_0_loop
_ = dictOrd_0
var app_1 gopurs_runtime.Value = app_1_loop
_ = app_1
var k_2 gopurs_runtime.Value = k_2_loop
_ = k_2
var v_3 gopurs_runtime.Value = v_3_loop
_ = v_3
var go__go_4_0_58 gopurs_runtime.Value
_ = go__go_4_0_58
go__go_4_0_58 = gopurs_runtime.Func(func(v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
var __t3 gopurs_runtime.Value
{
if (v1_5.Type == 9 && v1_5.IntVal == 324739070 && v1_5.UnsafePtr == nil) {
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(&Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]{1, 1, 1, k_2, v_3, gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(nil))}), gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(nil))})})}
goto end_branch_3
} else {

}
}
{
if (v1_5.Type == 9 && v1_5.IntVal == 324739070 && v1_5.UnsafePtr != nil) {
// TAST (Let): v2_6_1 -> gopurs_runtime.Value
v2_6_1 := gopurs_runtime.Apply2(gopurs_runtime.Box(dictOrd_0.V1), k_2, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v1_5.UnsafePtr).V2)
_ = v2_6_1
var __t2 gopurs_runtime.Value
{
if (uint32(v2_6_1.IntVal) == 1527465420) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeBalancedNode(), (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v1_5.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v1_5.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(go__go_4_0_58, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v1_5.UnsafePtr).V4)})))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v1_5.UnsafePtr).V5)})))}
goto end_branch_2
} else {

}
}
{
if (uint32(v2_6_1.IntVal) == 380165415) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeBalancedNode(), (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v1_5.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v1_5.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v1_5.UnsafePtr).V4)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(go__go_4_0_58, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v1_5.UnsafePtr).V5)})))})))}
goto end_branch_2
} else {

}
}
{
if (uint32(v2_6_1.IntVal) == 902936544) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(&Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v1_5.UnsafePtr).V0, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v1_5.UnsafePtr).V1, k_2, gopurs_runtime.Apply2(app_1, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v1_5.UnsafePtr).V3, v_3), (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v1_5.UnsafePtr).V4, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v1_5.UnsafePtr).V5})}
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_2:
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__t2))}
goto end_branch_3
} else {

}
}
{
__t3 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_3:
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__t3))}
})
return go__go_4_0_58
}

func Call_Data_Map_Internal_intersectionWith__3717755541(dictOrd_0_loop *Constructor_Data_Ord_Ord[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictOrd_0 *Constructor_Data_Ord_Ord[gopurs_runtime.Value] = dictOrd_0_loop
_ = dictOrd_0
// TAST (Let): compare_1_0 -> gopurs_runtime.Value
compare_1_0 := gopurs_runtime.Box(dictOrd_0.V1)
_ = compare_1_0
return gopurs_runtime.Func(func(app_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m2_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeIntersectionWith(), compare_1_0, app_2, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](m1_3))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](m2_4))})))}
})
})
})
}

func Call_Data_Map_Internal_intersectionWith__4144106805(dictOrd_0_loop *Constructor_Data_Ord_Ord[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictOrd_0 *Constructor_Data_Ord_Ord[gopurs_runtime.Value] = dictOrd_0_loop
_ = dictOrd_0
// TAST (Let): compare_1_0 -> gopurs_runtime.Value
compare_1_0 := gopurs_runtime.Box(dictOrd_0.V1)
_ = compare_1_0
return gopurs_runtime.Func(func(app_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m2_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeIntersectionWith(), compare_1_0, app_2, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](m1_3))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](m2_4))})))}
})
})
})
}

func Call_Data_Map_Internal_isEmpty__1620059593(v_0_loop *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]) bool {
var v_0 *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value] = v_0_loop
_ = v_0
var __t0 bool
{
if (gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_0)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_0)}.IntVal == 324739070 && gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr == nil) {
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

func Call_Data_Map_Internal_iterMapU__878452066(iter_0_loop gopurs_runtime.Value, v_1_loop *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var iter_0 gopurs_runtime.Value = iter_0_loop
_ = iter_0
var v_1 *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value] = v_1_loop
_ = v_1
var __t6 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_1)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_1)}.IntVal == 324739070 && gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_1)}.UnsafePtr == nil) {
__t6 = iter_0
goto end_branch_6
} else {

}
}
{
if (gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_1)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_1)}.IntVal == 324739070 && gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_1)}.UnsafePtr != nil) {
var __t5 *Constructor_Data_Map_Internal_IterEmit[gopurs_runtime.Value, gopurs_runtime.Value]
{
var __t_tag_2 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_1)}.UnsafePtr).V4)}
if (__t_tag_2.Type == 9 && __t_tag_2.IntVal == 324739070 && __t_tag_2.UnsafePtr == nil) {
var __t4 *Constructor_Data_Map_Internal_IterEmit[gopurs_runtime.Value, gopurs_runtime.Value]
{
var __t_tag_3 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_1)}.UnsafePtr).V5)}
if (__t_tag_3.Type == 9 && __t_tag_3.IntVal == 324739070 && __t_tag_3.UnsafePtr == nil) {
__t4 = &Constructor_Data_Map_Internal_IterEmit[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_1)}.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_1)}.UnsafePtr).V3, iter_0}
goto end_branch_4
} else {

}
}
{
__t4 = &Constructor_Data_Map_Internal_IterEmit[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_1)}.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_1)}.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 2861335956, UnsafePtr: unsafe.Pointer(&Constructor_Data_Map_Internal_IterNode[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_1)}.UnsafePtr).V5, iter_0})}}
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
var __t_tag_0 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_1)}.UnsafePtr).V5)}
if (__t_tag_0.Type == 9 && __t_tag_0.IntVal == 324739070 && __t_tag_0.UnsafePtr == nil) {
__t1 = &Constructor_Data_Map_Internal_IterEmit[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_1)}.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_1)}.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 2861335956, UnsafePtr: unsafe.Pointer(&Constructor_Data_Map_Internal_IterNode[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_1)}.UnsafePtr).V4, iter_0})}}
goto end_branch_1
} else {

}
}
{
__t1 = &Constructor_Data_Map_Internal_IterEmit[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_1)}.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_1)}.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 2861335956, UnsafePtr: unsafe.Pointer(&Constructor_Data_Map_Internal_IterNode[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_1)}.UnsafePtr).V4, gopurs_runtime.Value{Type: 9, IntVal: 2861335956, UnsafePtr: unsafe.Pointer(&Constructor_Data_Map_Internal_IterNode[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_1)}.UnsafePtr).V5, iter_0})}})}}
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

func Call_Data_Map_Internal_lookup__3378638282(dictOrd_0_loop *Constructor_Data_Ord_Ord[gopurs_runtime.Value], k_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 *Constructor_Data_Ord_Ord[gopurs_runtime.Value] = dictOrd_0_loop
_ = dictOrd_0
var k_1 gopurs_runtime.Value = k_1_loop
_ = k_1
var go__go_2_0_63 gopurs_runtime.Value
go__go_2_0_63 = gopurs_runtime.Func(func(v_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_3_loop gopurs_runtime.Value = v_3_loop_val
go__go_2_0_63:
for {
if false { continue go__go_2_0_63 }
var v_3 gopurs_runtime.Value = v_3_loop
_ = v_3
var __t3 gopurs_runtime.Value
{
if (v_3.Type == 9 && v_3.IntVal == 324739070 && v_3.UnsafePtr == nil) {
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil))}
goto end_branch_3
} else {

}
}
{
if (v_3.Type == 9 && v_3.IntVal == 324739070 && v_3.UnsafePtr != nil) {
// TAST (Let): v1_4_1 -> gopurs_runtime.Value
v1_4_1 := gopurs_runtime.Apply2(gopurs_runtime.Box(dictOrd_0.V1), k_1, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V2)
_ = v1_4_1
var __t2 gopurs_runtime.Value
{
if (uint32(v1_4_1.IntVal) == 1527465420) {
v_3_loop = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V4)}
continue go__go_2_0_63
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Value{}))}
goto end_branch_2
} else {

}
}
{
if (uint32(v1_4_1.IntVal) == 380165415) {
v_3_loop = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V5)}
continue go__go_2_0_63
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Value{}))}
goto end_branch_2
} else {

}
}
{
if (uint32(v1_4_1.IntVal) == 902936544) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{1, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V3})}
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_2:
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](__t2))}
goto end_branch_3
} else {

}
}
{
__t3 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_3:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](__t3))}
}
}()
})
return go__go_2_0_63
}

func Call_Data_Map_Internal_lookup__1040249709(k_0_loop uint32) gopurs_runtime.Value {
var k_0 uint32 = k_0_loop
_ = k_0
var go__go_1_0_64 gopurs_runtime.Value
go__go_1_0_64 = gopurs_runtime.Func(func(v_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_2_loop gopurs_runtime.Value = v_2_loop_val
go__go_1_0_64:
for {
if false { continue go__go_1_0_64 }
var v_2 gopurs_runtime.Value = v_2_loop
_ = v_2
var __t3 gopurs_runtime.Value
{
if (v_2.Type == 9 && v_2.IntVal == 324739070 && v_2.UnsafePtr == nil) {
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil))}
goto end_branch_3
} else {

}
}
{
if (v_2.Type == 9 && v_2.IntVal == 324739070 && v_2.UnsafePtr != nil) {
// TAST (Let): v1_3_1 -> gopurs_runtime.Value
v1_3_1 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_Interval_Duration_ordDurationComponent(), "compare"), gopurs_runtime.Value{Type: 9, IntVal: int64(k_0), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(uint32((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V2.IntVal)), UnsafePtr: nil})
_ = v1_3_1
var __t2 gopurs_runtime.Value
{
if (uint32(v1_3_1.IntVal) == 1527465420) {
v_2_loop = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V4)}
continue go__go_1_0_64
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[float64]](gopurs_runtime.Value{}))}
goto end_branch_2
} else {

}
}
{
if (uint32(v1_3_1.IntVal) == 380165415) {
v_2_loop = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V5)}
continue go__go_1_0_64
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[float64]](gopurs_runtime.Value{}))}
goto end_branch_2
} else {

}
}
{
if (uint32(v1_3_1.IntVal) == 902936544) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{1, gopurs_runtime.Float((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V3.FloatVal())})}
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_2:
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[float64]](__t2))}
goto end_branch_3
} else {

}
}
{
__t3 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_3:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[float64]](__t3))}
}
}()
})
return go__go_1_0_64
}

func Call_Data_Map_Internal_mapMaybe__3426301240(dictOrd_0_loop *Constructor_Data_Ord_Ord[gopurs_runtime.Value], x_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 *Constructor_Data_Ord_Ord[gopurs_runtime.Value] = dictOrd_0_loop
_ = dictOrd_0
var x_1 gopurs_runtime.Value = x_1_loop
_ = x_1
return Call_Data_Map_Internal_mapMaybeWithKey(dictOrd_0, gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return x_1
}))
}

func Call_Data_Map_Internal_mapMaybe__1970555288(dictOrd_0_loop *Constructor_Data_Ord_Ord[gopurs_runtime.Value], x_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 *Constructor_Data_Ord_Ord[gopurs_runtime.Value] = dictOrd_0_loop
_ = dictOrd_0
var x_1 gopurs_runtime.Value = x_1_loop
_ = x_1
return Call_Data_Map_Internal_mapMaybeWithKey(dictOrd_0, gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return x_1
}))
}

func Call_Data_Map_Internal_mapMaybeWithKey__817660689(dictOrd_0_loop *Constructor_Data_Ord_Ord[gopurs_runtime.Value], f_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 *Constructor_Data_Ord_Ord[gopurs_runtime.Value] = dictOrd_0_loop
_ = dictOrd_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var go__go_2_0_65 gopurs_runtime.Value
_ = go__go_2_0_65
go__go_2_0_65 = gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t3 gopurs_runtime.Value
{
if (v_3.Type == 9 && v_3.IntVal == 324739070 && v_3.UnsafePtr == nil) {
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(nil))}
goto end_branch_3
} else {

}
}
{
if (v_3.Type == 9 && v_3.IntVal == 324739070 && v_3.UnsafePtr != nil) {
// TAST (Let): v2_4_1 -> gopurs_runtime.Value
v2_4_1 := gopurs_runtime.Apply2(f_1, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V3)
_ = v2_4_1
var __t2 gopurs_runtime.Value
{
if (v2_4_1.Type == 9 && v2_4_1.IntVal == 930809136 && v2_4_1.UnsafePtr != nil) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeBalancedNode(), (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V2, (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(v2_4_1.UnsafePtr).V0, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(go__go_2_0_65, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V4)})))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(go__go_2_0_65, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V5)})))})))}
goto end_branch_2
} else {

}
}
{
if (v2_4_1.Type == 9 && v2_4_1.IntVal == 930809136 && v2_4_1.UnsafePtr == nil) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp2(Get_Data_Map_Internal_unsafeJoinNodes(), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(go__go_2_0_65, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V4)})))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(go__go_2_0_65, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V5)})))})))}
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_2:
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__t2))}
goto end_branch_3
} else {

}
}
{
__t3 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_3:
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__t3))}
})
return go__go_2_0_65
}

func Call_Data_Map_Internal_singleton__3511563426(k_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value] {
var k_0 gopurs_runtime.Value = k_0_loop
_ = k_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(&Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]{1, 1, 1, k_0, v_1, gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(nil))}), gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(nil))})})})
}

func Call_Data_Map_Internal_singleton__943571066(k_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value] {
var k_0 gopurs_runtime.Value = k_0_loop
_ = k_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(&Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]{1, 1, 1, k_0, v_1, gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(nil))}), gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(nil))})})})
}

func Call_Data_Map_Internal_singleton__2450056090(k_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value] {
var k_0 gopurs_runtime.Value = k_0_loop
_ = k_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(&Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]{1, 1, 1, k_0, v_1, gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(nil))}), gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(nil))})})})
}

func Call_Data_Map_Internal_singleton__3707014010(k_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value] {
var k_0 gopurs_runtime.Value = k_0_loop
_ = k_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(&Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]{1, 1, 1, k_0, v_1, gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(nil))}), gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(nil))})})})
}

func Call_Data_Map_Internal_singleton__1518627866(k_0_loop uint32, v_1_loop float64) *Constructor_Data_Map_Internal_Node[uint32, float64] {
var k_0 uint32 = k_0_loop
_ = k_0
var v_1 float64 = v_1_loop
_ = v_1
return gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[uint32, float64]](gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(&Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]{1, 1, 1, gopurs_runtime.Value{Type: 9, IntVal: int64(k_0), UnsafePtr: nil}, gopurs_runtime.Float(v_1), gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(nil))}), gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(nil))})})})
}

func Call_Data_Map_Internal_singleton__1300483034(k_0_loop *Constructor_Data_Maybe_Just[gopurs_runtime.Value], v_1_loop gopurs_runtime.Value) *Constructor_Data_Map_Internal_Node[*Constructor_Data_Maybe_Just[gopurs_runtime.Value], gopurs_runtime.Value] {
var k_0 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = k_0_loop
_ = k_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[*Constructor_Data_Maybe_Just[gopurs_runtime.Value], gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(&Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]{1, 1, 1, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(k_0)}, v_1, gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(nil))}), gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(nil))})})})
}

func Call_Data_Map_Internal_size__909390430(v_0_loop *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]) int64 {
var v_0 *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value] = v_0_loop
_ = v_0
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_0)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_0)}.IntVal == 324739070 && gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr == nil) {
__t0 = gopurs_runtime.Int(0)
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_0)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_0)}.IntVal == 324739070 && gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr != nil) {
__t0 = gopurs_runtime.Int((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V1)
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return __t0.IntVal
}

func Call_Data_Map_Internal_size__1374028086(v_0_loop *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]) int64 {
var v_0 *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value] = v_0_loop
_ = v_0
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_0)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_0)}.IntVal == 324739070 && gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr == nil) {
__t0 = gopurs_runtime.Int(0)
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_0)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_0)}.IntVal == 324739070 && gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr != nil) {
__t0 = gopurs_runtime.Int((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V1)
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return __t0.IntVal
}

func Call_Data_Map_Internal_size__2382154916(v_0_loop *Constructor_Data_Map_Internal_Node[uint32, float64]) int64 {
var v_0 *Constructor_Data_Map_Internal_Node[uint32, float64] = v_0_loop
_ = v_0
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_0)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_0)}.IntVal == 324739070 && gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr == nil) {
__t0 = gopurs_runtime.Int(0)
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_0)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_0)}.IntVal == 324739070 && gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr != nil) {
__t0 = gopurs_runtime.Int((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V1)
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return __t0.IntVal
}

func Call_Data_Map_Internal_stepWith__2632420966(f_0_loop gopurs_runtime.Value, next_1_loop gopurs_runtime.Value, done_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var next_1 gopurs_runtime.Value = next_1_loop
_ = next_1
var done_2 gopurs_runtime.Value = done_2_loop
_ = done_2
var go__go_3_0_66 gopurs_runtime.Value
go__go_3_0_66 = gopurs_runtime.Func(func(v_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_4_loop gopurs_runtime.Value = v_4_loop_val
go__go_3_0_66:
for {
if false { continue go__go_3_0_66 }
var v_4 gopurs_runtime.Value = v_4_loop
_ = v_4
var __t1 gopurs_runtime.Value
{
if (v_4.Type == 9 && v_4.IntVal == 2509360378) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]](gopurs_runtime.Apply(done_2, Get_Data_Unit_unit())))}
goto end_branch_1
} else {

}
}
{
if (v_4.Type == 9 && v_4.IntVal == 1343415489) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]](gopurs_runtime.UncurriedApp3(next_1, (*Constructor_Data_Map_Internal_IterEmit[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V0, (*Constructor_Data_Map_Internal_IterEmit[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V1, (*Constructor_Data_Map_Internal_IterEmit[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V2)))}
goto end_branch_1
} else {

}
}
{
if (v_4.Type == 9 && v_4.IntVal == 2861335956) {
v_4_loop = gopurs_runtime.Apply2(f_0, (*Constructor_Data_Map_Internal_IterNode[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_IterNode[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V0)})
continue go__go_3_0_66
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]](gopurs_runtime.Value{}))}
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]](__t1))}
}
}()
})
return go__go_3_0_66
}

func Call_Data_Map_Internal_stepWith__603436967(f_0_loop gopurs_runtime.Value, next_1_loop gopurs_runtime.Value, done_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var next_1 gopurs_runtime.Value = next_1_loop
_ = next_1
var done_2 gopurs_runtime.Value = done_2_loop
_ = done_2
var go__go_3_0_67 gopurs_runtime.Value
go__go_3_0_67 = gopurs_runtime.Func(func(v_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_4_loop gopurs_runtime.Value = v_4_loop_val
go__go_3_0_67:
for {
if false { continue go__go_3_0_67 }
var v_4 gopurs_runtime.Value = v_4_loop
_ = v_4
var __t1 gopurs_runtime.Value
{
if (v_4.Type == 9 && v_4.IntVal == 2509360378) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(done_2, Get_Data_Unit_unit())))}
goto end_branch_1
} else {

}
}
{
if (v_4.Type == 9 && v_4.IntVal == 1343415489) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp3(next_1, (*Constructor_Data_Map_Internal_IterEmit[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V0, (*Constructor_Data_Map_Internal_IterEmit[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V1, (*Constructor_Data_Map_Internal_IterEmit[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V2)))}
goto end_branch_1
} else {

}
}
{
if (v_4.Type == 9 && v_4.IntVal == 2861335956) {
v_4_loop = gopurs_runtime.Apply2(f_0, (*Constructor_Data_Map_Internal_IterNode[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_IterNode[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V0)})
continue go__go_3_0_67
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Value{}))}
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](__t1))}
}
}()
})
return go__go_3_0_67
}

func Call_Data_Map_Internal_stepWith__3186376421(f_0_loop gopurs_runtime.Value, next_1_loop gopurs_runtime.Value, done_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var next_1 gopurs_runtime.Value = next_1_loop
_ = next_1
var done_2 gopurs_runtime.Value = done_2_loop
_ = done_2
var go__go_3_0_68 gopurs_runtime.Value
go__go_3_0_68 = gopurs_runtime.Func(func(v_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_4_loop gopurs_runtime.Value = v_4_loop_val
go__go_3_0_68:
for {
if false { continue go__go_3_0_68 }
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
continue go__go_3_0_68
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
return go__go_3_0_68
}

func Call_Data_Map_Internal_stepWith__2866328237(f_0_loop gopurs_runtime.Value, next_1_loop gopurs_runtime.Value, done_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var next_1 gopurs_runtime.Value = next_1_loop
_ = next_1
var done_2 gopurs_runtime.Value = done_2_loop
_ = done_2
var go__go_3_0_69 gopurs_runtime.Value
go__go_3_0_69 = gopurs_runtime.Func(func(v_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_4_loop gopurs_runtime.Value = v_4_loop_val
go__go_3_0_69:
for {
if false { continue go__go_3_0_69 }
var v_4 gopurs_runtime.Value = v_4_loop
_ = v_4
var __t1 gopurs_runtime.Value
{
if (v_4.Type == 9 && v_4.IntVal == 2509360378) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 953589075, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_IterNext[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(done_2, Get_Data_Unit_unit())))}
goto end_branch_1
} else {

}
}
{
if (v_4.Type == 9 && v_4.IntVal == 1343415489) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 953589075, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_IterNext[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp3(next_1, (*Constructor_Data_Map_Internal_IterEmit[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V0, (*Constructor_Data_Map_Internal_IterEmit[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V1, (*Constructor_Data_Map_Internal_IterEmit[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V2)))}
goto end_branch_1
} else {

}
}
{
if (v_4.Type == 9 && v_4.IntVal == 2861335956) {
v_4_loop = gopurs_runtime.Apply2(f_0, (*Constructor_Data_Map_Internal_IterNode[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_IterNode[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V0)})
continue go__go_3_0_69
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 953589075, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_IterNext[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Value{}))}
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
return gopurs_runtime.Value{Type: 9, IntVal: 953589075, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_IterNext[gopurs_runtime.Value, gopurs_runtime.Value]](__t1))}
}
}()
})
return go__go_3_0_69
}

func Call_Data_Map_Internal_stepWith__280335550(f_0_loop gopurs_runtime.Value, next_1_loop gopurs_runtime.Value, done_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var next_1 gopurs_runtime.Value = next_1_loop
_ = next_1
var done_2 gopurs_runtime.Value = done_2_loop
_ = done_2
var go__go_3_0_70 gopurs_runtime.Value
go__go_3_0_70 = gopurs_runtime.Func(func(v_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_4_loop gopurs_runtime.Value = v_4_loop_val
go__go_3_0_70:
for {
if false { continue go__go_3_0_70 }
var v_4 gopurs_runtime.Value = v_4_loop
_ = v_4
var __t1 gopurs_runtime.Value
{
if (v_4.Type == 9 && v_4.IntVal == 2509360378) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[*Constructor_Data_Tuple_Tuple[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value], gopurs_runtime.Value]]](gopurs_runtime.Apply(done_2, Get_Data_Unit_unit())))}
goto end_branch_1
} else {

}
}
{
if (v_4.Type == 9 && v_4.IntVal == 1343415489) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[*Constructor_Data_Tuple_Tuple[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value], gopurs_runtime.Value]]](gopurs_runtime.UncurriedApp3(next_1, (*Constructor_Data_Map_Internal_IterEmit[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V0, (*Constructor_Data_Map_Internal_IterEmit[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V1, (*Constructor_Data_Map_Internal_IterEmit[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V2)))}
goto end_branch_1
} else {

}
}
{
if (v_4.Type == 9 && v_4.IntVal == 2861335956) {
v_4_loop = gopurs_runtime.Apply2(f_0, (*Constructor_Data_Map_Internal_IterNode[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_IterNode[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V0)})
continue go__go_3_0_70
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[*Constructor_Data_Tuple_Tuple[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value], gopurs_runtime.Value]]](gopurs_runtime.Value{}))}
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[*Constructor_Data_Tuple_Tuple[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value], gopurs_runtime.Value]]](__t1))}
}
}()
})
return go__go_3_0_70
}

func Call_Data_Map_Internal_stepWith__2834533669(f_0_loop gopurs_runtime.Value, next_1_loop gopurs_runtime.Value, done_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var next_1 gopurs_runtime.Value = next_1_loop
_ = next_1
var done_2 gopurs_runtime.Value = done_2_loop
_ = done_2
var go__go_3_0_71 gopurs_runtime.Value
go__go_3_0_71 = gopurs_runtime.Func(func(v_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_4_loop gopurs_runtime.Value = v_4_loop_val
go__go_3_0_71:
for {
if false { continue go__go_3_0_71 }
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
__t1 = gopurs_runtime.UncurriedApp3(next_1, gopurs_runtime.Value{Type: 9, IntVal: int64(uint32((*Constructor_Data_Map_Internal_IterEmit[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V0.IntVal)), UnsafePtr: nil}, gopurs_runtime.Float((*Constructor_Data_Map_Internal_IterEmit[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V1.FloatVal()), (*Constructor_Data_Map_Internal_IterEmit[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V2)
goto end_branch_1
} else {

}
}
{
if (v_4.Type == 9 && v_4.IntVal == 2861335956) {
v_4_loop = gopurs_runtime.Apply2(f_0, (*Constructor_Data_Map_Internal_IterNode[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_IterNode[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V0)})
continue go__go_3_0_71
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
return go__go_3_0_71
}

func Call_Data_Map_Internal_stepWith__1463181374(f_0_loop gopurs_runtime.Value, next_1_loop gopurs_runtime.Value, done_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var next_1 gopurs_runtime.Value = next_1_loop
_ = next_1
var done_2 gopurs_runtime.Value = done_2_loop
_ = done_2
var go__go_3_0_72 gopurs_runtime.Value
go__go_3_0_72 = gopurs_runtime.Func(func(v_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_4_loop gopurs_runtime.Value = v_4_loop_val
go__go_3_0_72:
for {
if false { continue go__go_3_0_72 }
var v_4 gopurs_runtime.Value = v_4_loop
_ = v_4
var __t1 gopurs_runtime.Value
{
if (v_4.Type == 9 && v_4.IntVal == 2509360378) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[*Constructor_Data_Tuple_Tuple[*Constructor_Data_Tuple_Tuple[uint32, float64], gopurs_runtime.Value]]](gopurs_runtime.Apply(done_2, Get_Data_Unit_unit())))}
goto end_branch_1
} else {

}
}
{
if (v_4.Type == 9 && v_4.IntVal == 1343415489) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[*Constructor_Data_Tuple_Tuple[*Constructor_Data_Tuple_Tuple[uint32, float64], gopurs_runtime.Value]]](gopurs_runtime.UncurriedApp3(next_1, gopurs_runtime.Value{Type: 9, IntVal: int64(uint32((*Constructor_Data_Map_Internal_IterEmit[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V0.IntVal)), UnsafePtr: nil}, gopurs_runtime.Float((*Constructor_Data_Map_Internal_IterEmit[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V1.FloatVal()), (*Constructor_Data_Map_Internal_IterEmit[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V2)))}
goto end_branch_1
} else {

}
}
{
if (v_4.Type == 9 && v_4.IntVal == 2861335956) {
v_4_loop = gopurs_runtime.Apply2(f_0, (*Constructor_Data_Map_Internal_IterNode[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_IterNode[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V0)})
continue go__go_3_0_72
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[*Constructor_Data_Tuple_Tuple[*Constructor_Data_Tuple_Tuple[uint32, float64], gopurs_runtime.Value]]](gopurs_runtime.Value{}))}
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[*Constructor_Data_Tuple_Tuple[*Constructor_Data_Tuple_Tuple[uint32, float64], gopurs_runtime.Value]]](__t1))}
}
}()
})
return go__go_3_0_72
}

func Call_Data_Map_Internal_toMapIter__1799172593(a_0_loop *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var a_0 *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value] = a_0_loop
_ = a_0
return gopurs_runtime.Value{Type: 9, IntVal: 2861335956, UnsafePtr: unsafe.Pointer(&Constructor_Data_Map_Internal_IterNode[gopurs_runtime.Value, gopurs_runtime.Value]{1, a_0, gopurs_runtime.Value{Type: 9, IntVal: 2509360378, UnsafePtr: unsafe.Pointer(nil)}})}
}

func Call_Data_Map_Internal_toMapIter__2014410513(a_0_loop *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var a_0 *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value] = a_0_loop
_ = a_0
return gopurs_runtime.Value{Type: 9, IntVal: 2861335956, UnsafePtr: unsafe.Pointer(&Constructor_Data_Map_Internal_IterNode[gopurs_runtime.Value, gopurs_runtime.Value]{1, a_0, gopurs_runtime.Value{Type: 9, IntVal: 2509360378, UnsafePtr: unsafe.Pointer(nil)}})}
}

func Call_Data_Map_Internal_toMapIter__772765521(a_0_loop *Constructor_Data_Map_Internal_Node[uint32, float64]) gopurs_runtime.Value {
var a_0 *Constructor_Data_Map_Internal_Node[uint32, float64] = a_0_loop
_ = a_0
return gopurs_runtime.Value{Type: 9, IntVal: 2861335956, UnsafePtr: unsafe.Pointer(&Constructor_Data_Map_Internal_IterNode[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(a_0)}), gopurs_runtime.Value{Type: 9, IntVal: 2509360378, UnsafePtr: unsafe.Pointer(nil)}})}
}

func Call_Data_Map_Internal_toMapIter__1738891721(a_0_loop *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var a_0 *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value] = a_0_loop
_ = a_0
return gopurs_runtime.Value{Type: 9, IntVal: 2861335956, UnsafePtr: unsafe.Pointer(&Constructor_Data_Map_Internal_IterNode[gopurs_runtime.Value, gopurs_runtime.Value]{1, a_0, gopurs_runtime.Value{Type: 9, IntVal: 2509360378, UnsafePtr: unsafe.Pointer(nil)}})}
}

func Call_Data_Map_Internal_toUnfoldable__2183602684(dictUnfoldable_0_loop *Constructor_Data_Unfoldable_Unfoldable[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictUnfoldable_0 *Constructor_Data_Unfoldable_Unfoldable[gopurs_runtime.Value] = dictUnfoldable_0_loop
_ = dictUnfoldable_0
// TAST (Let): __local_var_1_0 -> gopurs_runtime.Value
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.Box(dictUnfoldable_0.V1), Get_Data_Map_Internal_stepUnfoldr())
_ = __local_var_1_0
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_1_0, gopurs_runtime.Value{Type: 9, IntVal: 2861335956, UnsafePtr: unsafe.Pointer(&Constructor_Data_Map_Internal_IterNode[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](x_2), gopurs_runtime.Value{Type: 9, IntVal: 2509360378, UnsafePtr: unsafe.Pointer(nil)}})})
})
}

func Call_Data_Map_Internal_toUnfoldable__2567957978(__eta0_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[*Constructor_Data_Tuple_Tuple[uint32, float64]]](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_List_Types_unfoldableList(), "unfoldr"), Get_Data_Map_Internal_stepUnfoldr(), gopurs_runtime.Value{Type: 9, IntVal: 2861335956, UnsafePtr: unsafe.Pointer(&Constructor_Data_Map_Internal_IterNode[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__eta0_0), gopurs_runtime.Value{Type: 9, IntVal: 2509360378, UnsafePtr: unsafe.Pointer(nil)}})})))}
}

func Call_Data_Map_Internal_unionWith__2507192643(dictOrd_0_loop *Constructor_Data_Ord_Ord[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictOrd_0 *Constructor_Data_Ord_Ord[gopurs_runtime.Value] = dictOrd_0_loop
_ = dictOrd_0
// TAST (Let): compare_1_0 -> gopurs_runtime.Value
compare_1_0 := gopurs_runtime.Box(dictOrd_0.V1)
_ = compare_1_0
return gopurs_runtime.Func(func(app_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m2_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeUnionWith(), compare_1_0, app_2, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](m1_3))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](m2_4))})))}
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
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[uint32, float64]](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeUnionWith(), gopurs_runtime.RecordGet(Get_Data_Interval_Duration_ordDurationComponent(), "compare"), __eta0_0, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[uint32, float64]](__eta1_1))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[uint32, float64]](__eta2_2))})))}
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
var __t37 gopurs_runtime.Value
{
if (__local_var_2.Type == 9 && __local_var_2.IntVal == 324739070 && __local_var_2.UnsafePtr == nil) {
var __t9 gopurs_runtime.Value
{
if (__local_var_3.Type == 9 && __local_var_3.IntVal == 324739070 && __local_var_3.UnsafePtr == nil) {
__t9 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(&Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]{1, 1, 1, __local_var_0, __local_var_1, gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(nil))}), gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(nil))})})}
goto end_branch_9
} else {

}
}
{
var __t_and_1 bool = false
if (__local_var_3.Type == 9 && __local_var_3.IntVal == 324739070 && __local_var_3.UnsafePtr != nil) {

var __t0 bool
{
if ((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V0) > (1) {
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
var __t8 *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]
{
var __t_tag_2 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V4)}
var __t_and_7 bool = false
if (__t_tag_2.Type == 9 && __t_tag_2.IntVal == 324739070 && __t_tag_2.UnsafePtr != nil) {

var __t6 bool
{
var __t5 gopurs_runtime.Value
{
var __t_tag_3 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V5)}
if (__t_tag_3.Type == 9 && __t_tag_3.IntVal == 324739070 && __t_tag_3.UnsafePtr == nil) {
__t5 = gopurs_runtime.Int(0)
goto end_branch_5
} else {

}
}
{
var __t_tag_4 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V5)}
if (__t_tag_4.Type == 9 && __t_tag_4.IntVal == 324739070 && __t_tag_4.UnsafePtr != nil) {
__t5 = gopurs_runtime.Int((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V5)}.UnsafePtr).V0)
goto end_branch_5
} else {

}
}
{
__t5 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_5:
if ((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V4)}.UnsafePtr).V0) > (__t5.IntVal) {
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
__t8 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeNode(), (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V4)}.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V4)}.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeNode(), __local_var_0, __local_var_1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__local_var_2))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V4)}.UnsafePtr).V4)})))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeNode(), (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V4)}.UnsafePtr).V5)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V5)})))}))
goto end_branch_8
} else {

}
}
{
__t8 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeNode(), (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeNode(), __local_var_0, __local_var_1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__local_var_2))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V4)})))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V5)}))
}
end_branch_8:
__t9 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(__t8)}
goto end_branch_9
} else {

}
}
{
__t9 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeNode(), __local_var_0, __local_var_1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__local_var_2))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__local_var_3))})))}
}
end_branch_9:
__t37 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__t9))}
goto end_branch_37
} else {

}
}
{
if (__local_var_2.Type == 9 && __local_var_2.IntVal == 324739070 && __local_var_2.UnsafePtr != nil) {
var __t36 *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]
{
if (__local_var_3.Type == 9 && __local_var_3.IntVal == 324739070 && __local_var_3.UnsafePtr != nil) {
var __t26 *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]
{
var __t10 bool
{
if ((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V0) > (((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V0) + (1)) {
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
var __t17 *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]
{
var __t_tag_11 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V4)}
var __t_and_16 bool = false
if (__t_tag_11.Type == 9 && __t_tag_11.IntVal == 324739070 && __t_tag_11.UnsafePtr != nil) {

var __t15 bool
{
var __t14 gopurs_runtime.Value
{
var __t_tag_12 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V5)}
if (__t_tag_12.Type == 9 && __t_tag_12.IntVal == 324739070 && __t_tag_12.UnsafePtr == nil) {
__t14 = gopurs_runtime.Int(0)
goto end_branch_14
} else {

}
}
{
var __t_tag_13 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V5)}
if (__t_tag_13.Type == 9 && __t_tag_13.IntVal == 324739070 && __t_tag_13.UnsafePtr != nil) {
__t14 = gopurs_runtime.Int((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V5)}.UnsafePtr).V0)
goto end_branch_14
} else {

}
}
{
__t14 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_14:
if ((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V4)}.UnsafePtr).V0) > (__t14.IntVal) {
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
__t17 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeNode(), (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V4)}.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V4)}.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeNode(), __local_var_0, __local_var_1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__local_var_2))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V4)}.UnsafePtr).V4)})))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeNode(), (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V4)}.UnsafePtr).V5)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V5)})))}))
goto end_branch_17
} else {

}
}
{
__t17 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeNode(), (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeNode(), __local_var_0, __local_var_1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__local_var_2))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V4)})))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V5)}))
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
if ((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V0) > (((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V0) + (1)) {
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
var __t25 *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]
{
var __t_tag_19 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V5)}
var __t_and_24 bool = false
if (__t_tag_19.Type == 9 && __t_tag_19.IntVal == 324739070 && __t_tag_19.UnsafePtr != nil) {

var __t23 bool
{
var __t22 gopurs_runtime.Value
{
var __t_tag_20 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V4)}
if (__t_tag_20.Type == 9 && __t_tag_20.IntVal == 324739070 && __t_tag_20.UnsafePtr == nil) {
__t22 = gopurs_runtime.Int(0)
goto end_branch_22
} else {

}
}
{
var __t_tag_21 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V4)}
if (__t_tag_21.Type == 9 && __t_tag_21.IntVal == 324739070 && __t_tag_21.UnsafePtr != nil) {
__t22 = gopurs_runtime.Int((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V4)}.UnsafePtr).V0)
goto end_branch_22
} else {

}
}
{
__t22 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_22:
if (__t22.IntVal) > ((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V5)}.UnsafePtr).V0) {
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
__t25 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeNode(), (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V5)}.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V5)}.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeNode(), (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V4)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V5)}.UnsafePtr).V4)})))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeNode(), __local_var_0, __local_var_1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V5)}.UnsafePtr).V5)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__local_var_3))})))}))
goto end_branch_25
} else {

}
}
{
__t25 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeNode(), (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V4)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeNode(), __local_var_0, __local_var_1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V5)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__local_var_3))})))}))
}
end_branch_25:
__t26 = __t25
goto end_branch_26
} else {

}
}
{
__t26 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeNode(), __local_var_0, __local_var_1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__local_var_2))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__local_var_3))}))
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
if ((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V0) > (1) {
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
var __t35 *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]
{
var __t_tag_29 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V5)}
var __t_and_34 bool = false
if (__t_tag_29.Type == 9 && __t_tag_29.IntVal == 324739070 && __t_tag_29.UnsafePtr != nil) {

var __t33 bool
{
var __t32 gopurs_runtime.Value
{
var __t_tag_30 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V4)}
if (__t_tag_30.Type == 9 && __t_tag_30.IntVal == 324739070 && __t_tag_30.UnsafePtr == nil) {
__t32 = gopurs_runtime.Int(0)
goto end_branch_32
} else {

}
}
{
var __t_tag_31 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V4)}
if (__t_tag_31.Type == 9 && __t_tag_31.IntVal == 324739070 && __t_tag_31.UnsafePtr != nil) {
__t32 = gopurs_runtime.Int((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V4)}.UnsafePtr).V0)
goto end_branch_32
} else {

}
}
{
__t32 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_32:
if (__t32.IntVal) > ((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V5)}.UnsafePtr).V0) {
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
__t35 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeNode(), (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V5)}.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V5)}.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeNode(), (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V4)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V5)}.UnsafePtr).V4)})))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeNode(), __local_var_0, __local_var_1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V5)}.UnsafePtr).V5)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__local_var_3))})))}))
goto end_branch_35
} else {

}
}
{
__t35 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeNode(), (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V4)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeNode(), __local_var_0, __local_var_1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V5)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__local_var_3))})))}))
}
end_branch_35:
__t36 = __t35
goto end_branch_36
} else {

}
}
{
__t36 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeNode(), __local_var_0, __local_var_1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__local_var_2))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__local_var_3))}))
}
end_branch_36:
__t37 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(__t36)}
goto end_branch_37
} else {

}
}
{
__t37 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_37:
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__t37))}
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
var __t37 gopurs_runtime.Value
{
if (__local_var_2.Type == 9 && __local_var_2.IntVal == 324739070 && __local_var_2.UnsafePtr == nil) {
var __t9 gopurs_runtime.Value
{
if (__local_var_3.Type == 9 && __local_var_3.IntVal == 324739070 && __local_var_3.UnsafePtr == nil) {
__t9 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(&Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]{1, 1, 1, __local_var_0, __local_var_1, gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(nil))}), gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(nil))})})}
goto end_branch_9
} else {

}
}
{
var __t_and_1 bool = false
if (__local_var_3.Type == 9 && __local_var_3.IntVal == 324739070 && __local_var_3.UnsafePtr != nil) {

var __t0 bool
{
if ((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V0) > (1) {
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
var __t8 *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]
{
var __t_tag_2 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V4)}
var __t_and_7 bool = false
if (__t_tag_2.Type == 9 && __t_tag_2.IntVal == 324739070 && __t_tag_2.UnsafePtr != nil) {

var __t6 bool
{
var __t5 gopurs_runtime.Value
{
var __t_tag_3 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V5)}
if (__t_tag_3.Type == 9 && __t_tag_3.IntVal == 324739070 && __t_tag_3.UnsafePtr == nil) {
__t5 = gopurs_runtime.Int(0)
goto end_branch_5
} else {

}
}
{
var __t_tag_4 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V5)}
if (__t_tag_4.Type == 9 && __t_tag_4.IntVal == 324739070 && __t_tag_4.UnsafePtr != nil) {
__t5 = gopurs_runtime.Int((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V5)}.UnsafePtr).V0)
goto end_branch_5
} else {

}
}
{
__t5 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_5:
if ((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V4)}.UnsafePtr).V0) > (__t5.IntVal) {
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
__t8 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeNode(), (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V4)}.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V4)}.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeNode(), __local_var_0, __local_var_1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__local_var_2))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V4)}.UnsafePtr).V4)})))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeNode(), (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V4)}.UnsafePtr).V5)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V5)})))}))
goto end_branch_8
} else {

}
}
{
__t8 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeNode(), (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeNode(), __local_var_0, __local_var_1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__local_var_2))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V4)})))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V5)}))
}
end_branch_8:
__t9 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(__t8)}
goto end_branch_9
} else {

}
}
{
__t9 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeNode(), __local_var_0, __local_var_1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__local_var_2))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__local_var_3))})))}
}
end_branch_9:
__t37 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__t9))}
goto end_branch_37
} else {

}
}
{
if (__local_var_2.Type == 9 && __local_var_2.IntVal == 324739070 && __local_var_2.UnsafePtr != nil) {
var __t36 *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]
{
if (__local_var_3.Type == 9 && __local_var_3.IntVal == 324739070 && __local_var_3.UnsafePtr != nil) {
var __t26 *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]
{
var __t10 bool
{
if ((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V0) > (((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V0) + (1)) {
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
var __t17 *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]
{
var __t_tag_11 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V4)}
var __t_and_16 bool = false
if (__t_tag_11.Type == 9 && __t_tag_11.IntVal == 324739070 && __t_tag_11.UnsafePtr != nil) {

var __t15 bool
{
var __t14 gopurs_runtime.Value
{
var __t_tag_12 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V5)}
if (__t_tag_12.Type == 9 && __t_tag_12.IntVal == 324739070 && __t_tag_12.UnsafePtr == nil) {
__t14 = gopurs_runtime.Int(0)
goto end_branch_14
} else {

}
}
{
var __t_tag_13 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V5)}
if (__t_tag_13.Type == 9 && __t_tag_13.IntVal == 324739070 && __t_tag_13.UnsafePtr != nil) {
__t14 = gopurs_runtime.Int((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V5)}.UnsafePtr).V0)
goto end_branch_14
} else {

}
}
{
__t14 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_14:
if ((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V4)}.UnsafePtr).V0) > (__t14.IntVal) {
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
__t17 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeNode(), (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V4)}.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V4)}.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeNode(), __local_var_0, __local_var_1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__local_var_2))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V4)}.UnsafePtr).V4)})))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeNode(), (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V4)}.UnsafePtr).V5)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V5)})))}))
goto end_branch_17
} else {

}
}
{
__t17 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeNode(), (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeNode(), __local_var_0, __local_var_1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__local_var_2))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V4)})))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V5)}))
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
if ((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V0) > (((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V0) + (1)) {
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
var __t25 *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]
{
var __t_tag_19 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V5)}
var __t_and_24 bool = false
if (__t_tag_19.Type == 9 && __t_tag_19.IntVal == 324739070 && __t_tag_19.UnsafePtr != nil) {

var __t23 bool
{
var __t22 gopurs_runtime.Value
{
var __t_tag_20 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V4)}
if (__t_tag_20.Type == 9 && __t_tag_20.IntVal == 324739070 && __t_tag_20.UnsafePtr == nil) {
__t22 = gopurs_runtime.Int(0)
goto end_branch_22
} else {

}
}
{
var __t_tag_21 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V4)}
if (__t_tag_21.Type == 9 && __t_tag_21.IntVal == 324739070 && __t_tag_21.UnsafePtr != nil) {
__t22 = gopurs_runtime.Int((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V4)}.UnsafePtr).V0)
goto end_branch_22
} else {

}
}
{
__t22 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_22:
if (__t22.IntVal) > ((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V5)}.UnsafePtr).V0) {
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
__t25 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeNode(), (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V5)}.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V5)}.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeNode(), (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V4)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V5)}.UnsafePtr).V4)})))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeNode(), __local_var_0, __local_var_1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V5)}.UnsafePtr).V5)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__local_var_3))})))}))
goto end_branch_25
} else {

}
}
{
__t25 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeNode(), (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V4)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeNode(), __local_var_0, __local_var_1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V5)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__local_var_3))})))}))
}
end_branch_25:
__t26 = __t25
goto end_branch_26
} else {

}
}
{
__t26 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeNode(), __local_var_0, __local_var_1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__local_var_2))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__local_var_3))}))
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
if ((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V0) > (1) {
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
var __t35 *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]
{
var __t_tag_29 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V5)}
var __t_and_34 bool = false
if (__t_tag_29.Type == 9 && __t_tag_29.IntVal == 324739070 && __t_tag_29.UnsafePtr != nil) {

var __t33 bool
{
var __t32 gopurs_runtime.Value
{
var __t_tag_30 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V4)}
if (__t_tag_30.Type == 9 && __t_tag_30.IntVal == 324739070 && __t_tag_30.UnsafePtr == nil) {
__t32 = gopurs_runtime.Int(0)
goto end_branch_32
} else {

}
}
{
var __t_tag_31 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V4)}
if (__t_tag_31.Type == 9 && __t_tag_31.IntVal == 324739070 && __t_tag_31.UnsafePtr != nil) {
__t32 = gopurs_runtime.Int((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V4)}.UnsafePtr).V0)
goto end_branch_32
} else {

}
}
{
__t32 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_32:
if (__t32.IntVal) > ((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V5)}.UnsafePtr).V0) {
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
__t35 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeNode(), (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V5)}.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V5)}.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeNode(), (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V4)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V5)}.UnsafePtr).V4)})))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeNode(), __local_var_0, __local_var_1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V5)}.UnsafePtr).V5)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__local_var_3))})))}))
goto end_branch_35
} else {

}
}
{
__t35 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeNode(), (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V4)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeNode(), __local_var_0, __local_var_1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V5)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__local_var_3))})))}))
}
end_branch_35:
__t36 = __t35
goto end_branch_36
} else {

}
}
{
__t36 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeNode(), __local_var_0, __local_var_1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__local_var_2))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__local_var_3))}))
}
end_branch_36:
__t37 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(__t36)}
goto end_branch_37
} else {

}
}
{
__t37 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_37:
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__t37))}
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
var __t37 gopurs_runtime.Value
{
if (__local_var_2.Type == 9 && __local_var_2.IntVal == 324739070 && __local_var_2.UnsafePtr == nil) {
var __t9 gopurs_runtime.Value
{
if (__local_var_3.Type == 9 && __local_var_3.IntVal == 324739070 && __local_var_3.UnsafePtr == nil) {
__t9 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(&Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]{1, 1, 1, __local_var_0, __local_var_1, gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(nil))}), gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(nil))})})}
goto end_branch_9
} else {

}
}
{
var __t_and_1 bool = false
if (__local_var_3.Type == 9 && __local_var_3.IntVal == 324739070 && __local_var_3.UnsafePtr != nil) {

var __t0 bool
{
if ((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V0) > (1) {
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
var __t8 *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]
{
var __t_tag_2 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V4)}
var __t_and_7 bool = false
if (__t_tag_2.Type == 9 && __t_tag_2.IntVal == 324739070 && __t_tag_2.UnsafePtr != nil) {

var __t6 bool
{
var __t5 gopurs_runtime.Value
{
var __t_tag_3 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V5)}
if (__t_tag_3.Type == 9 && __t_tag_3.IntVal == 324739070 && __t_tag_3.UnsafePtr == nil) {
__t5 = gopurs_runtime.Int(0)
goto end_branch_5
} else {

}
}
{
var __t_tag_4 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V5)}
if (__t_tag_4.Type == 9 && __t_tag_4.IntVal == 324739070 && __t_tag_4.UnsafePtr != nil) {
__t5 = gopurs_runtime.Int((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V5)}.UnsafePtr).V0)
goto end_branch_5
} else {

}
}
{
__t5 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_5:
if ((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V4)}.UnsafePtr).V0) > (__t5.IntVal) {
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
__t8 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeNode(), (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V4)}.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V4)}.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeNode(), __local_var_0, __local_var_1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__local_var_2))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V4)}.UnsafePtr).V4)})))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeNode(), (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V4)}.UnsafePtr).V5)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V5)})))}))
goto end_branch_8
} else {

}
}
{
__t8 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeNode(), (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeNode(), __local_var_0, __local_var_1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__local_var_2))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V4)})))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V5)}))
}
end_branch_8:
__t9 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(__t8)}
goto end_branch_9
} else {

}
}
{
__t9 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeNode(), __local_var_0, __local_var_1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__local_var_2))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__local_var_3))})))}
}
end_branch_9:
__t37 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__t9))}
goto end_branch_37
} else {

}
}
{
if (__local_var_2.Type == 9 && __local_var_2.IntVal == 324739070 && __local_var_2.UnsafePtr != nil) {
var __t36 *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]
{
if (__local_var_3.Type == 9 && __local_var_3.IntVal == 324739070 && __local_var_3.UnsafePtr != nil) {
var __t26 *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]
{
var __t10 bool
{
if ((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V0) > (((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V0) + (1)) {
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
var __t17 *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]
{
var __t_tag_11 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V4)}
var __t_and_16 bool = false
if (__t_tag_11.Type == 9 && __t_tag_11.IntVal == 324739070 && __t_tag_11.UnsafePtr != nil) {

var __t15 bool
{
var __t14 gopurs_runtime.Value
{
var __t_tag_12 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V5)}
if (__t_tag_12.Type == 9 && __t_tag_12.IntVal == 324739070 && __t_tag_12.UnsafePtr == nil) {
__t14 = gopurs_runtime.Int(0)
goto end_branch_14
} else {

}
}
{
var __t_tag_13 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V5)}
if (__t_tag_13.Type == 9 && __t_tag_13.IntVal == 324739070 && __t_tag_13.UnsafePtr != nil) {
__t14 = gopurs_runtime.Int((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V5)}.UnsafePtr).V0)
goto end_branch_14
} else {

}
}
{
__t14 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_14:
if ((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V4)}.UnsafePtr).V0) > (__t14.IntVal) {
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
__t17 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeNode(), (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V4)}.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V4)}.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeNode(), __local_var_0, __local_var_1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__local_var_2))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V4)}.UnsafePtr).V4)})))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeNode(), (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V4)}.UnsafePtr).V5)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V5)})))}))
goto end_branch_17
} else {

}
}
{
__t17 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeNode(), (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeNode(), __local_var_0, __local_var_1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__local_var_2))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V4)})))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V5)}))
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
if ((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V0) > (((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V0) + (1)) {
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
var __t25 *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]
{
var __t_tag_19 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V5)}
var __t_and_24 bool = false
if (__t_tag_19.Type == 9 && __t_tag_19.IntVal == 324739070 && __t_tag_19.UnsafePtr != nil) {

var __t23 bool
{
var __t22 gopurs_runtime.Value
{
var __t_tag_20 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V4)}
if (__t_tag_20.Type == 9 && __t_tag_20.IntVal == 324739070 && __t_tag_20.UnsafePtr == nil) {
__t22 = gopurs_runtime.Int(0)
goto end_branch_22
} else {

}
}
{
var __t_tag_21 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V4)}
if (__t_tag_21.Type == 9 && __t_tag_21.IntVal == 324739070 && __t_tag_21.UnsafePtr != nil) {
__t22 = gopurs_runtime.Int((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V4)}.UnsafePtr).V0)
goto end_branch_22
} else {

}
}
{
__t22 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_22:
if (__t22.IntVal) > ((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V5)}.UnsafePtr).V0) {
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
__t25 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeNode(), (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V5)}.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V5)}.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeNode(), (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V4)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V5)}.UnsafePtr).V4)})))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeNode(), __local_var_0, __local_var_1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V5)}.UnsafePtr).V5)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__local_var_3))})))}))
goto end_branch_25
} else {

}
}
{
__t25 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeNode(), (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V4)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeNode(), __local_var_0, __local_var_1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V5)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__local_var_3))})))}))
}
end_branch_25:
__t26 = __t25
goto end_branch_26
} else {

}
}
{
__t26 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeNode(), __local_var_0, __local_var_1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__local_var_2))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__local_var_3))}))
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
if ((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V0) > (1) {
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
var __t35 *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]
{
var __t_tag_29 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V5)}
var __t_and_34 bool = false
if (__t_tag_29.Type == 9 && __t_tag_29.IntVal == 324739070 && __t_tag_29.UnsafePtr != nil) {

var __t33 bool
{
var __t32 gopurs_runtime.Value
{
var __t_tag_30 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V4)}
if (__t_tag_30.Type == 9 && __t_tag_30.IntVal == 324739070 && __t_tag_30.UnsafePtr == nil) {
__t32 = gopurs_runtime.Int(0)
goto end_branch_32
} else {

}
}
{
var __t_tag_31 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V4)}
if (__t_tag_31.Type == 9 && __t_tag_31.IntVal == 324739070 && __t_tag_31.UnsafePtr != nil) {
__t32 = gopurs_runtime.Int((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V4)}.UnsafePtr).V0)
goto end_branch_32
} else {

}
}
{
__t32 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_32:
if (__t32.IntVal) > ((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V5)}.UnsafePtr).V0) {
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
__t35 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeNode(), (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V5)}.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V5)}.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeNode(), (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V4)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V5)}.UnsafePtr).V4)})))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeNode(), __local_var_0, __local_var_1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V5)}.UnsafePtr).V5)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__local_var_3))})))}))
goto end_branch_35
} else {

}
}
{
__t35 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeNode(), (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V4)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeNode(), __local_var_0, __local_var_1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V5)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__local_var_3))})))}))
}
end_branch_35:
__t36 = __t35
goto end_branch_36
} else {

}
}
{
__t36 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeNode(), __local_var_0, __local_var_1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__local_var_2))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__local_var_3))}))
}
end_branch_36:
__t37 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(__t36)}
goto end_branch_37
} else {

}
}
{
__t37 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_37:
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__t37))}
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
var __t37 gopurs_runtime.Value
{
if (__local_var_2.Type == 9 && __local_var_2.IntVal == 324739070 && __local_var_2.UnsafePtr == nil) {
var __t9 gopurs_runtime.Value
{
if (__local_var_3.Type == 9 && __local_var_3.IntVal == 324739070 && __local_var_3.UnsafePtr == nil) {
__t9 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(&Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]{1, 1, 1, gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(__local_var_0.IntVal)), UnsafePtr: nil}, gopurs_runtime.Float(__local_var_1.FloatVal()), gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(nil))}), gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(nil))})})}
goto end_branch_9
} else {

}
}
{
var __t_and_1 bool = false
if (__local_var_3.Type == 9 && __local_var_3.IntVal == 324739070 && __local_var_3.UnsafePtr != nil) {

var __t0 bool
{
if ((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V0) > (1) {
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
var __t8 *Constructor_Data_Map_Internal_Node[uint32, float64]
{
var __t_tag_2 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V4)}
var __t_and_7 bool = false
if (__t_tag_2.Type == 9 && __t_tag_2.IntVal == 324739070 && __t_tag_2.UnsafePtr != nil) {

var __t6 bool
{
var __t5 gopurs_runtime.Value
{
var __t_tag_3 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V5)}
if (__t_tag_3.Type == 9 && __t_tag_3.IntVal == 324739070 && __t_tag_3.UnsafePtr == nil) {
__t5 = gopurs_runtime.Int(0)
goto end_branch_5
} else {

}
}
{
var __t_tag_4 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V5)}
if (__t_tag_4.Type == 9 && __t_tag_4.IntVal == 324739070 && __t_tag_4.UnsafePtr != nil) {
__t5 = gopurs_runtime.Int((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V5)}.UnsafePtr).V0)
goto end_branch_5
} else {

}
}
{
__t5 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_5:
if ((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V4)}.UnsafePtr).V0) > (__t5.IntVal) {
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
__t8 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[uint32, float64]](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeNode(), gopurs_runtime.Value{Type: 9, IntVal: int64(uint32((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V4)}.UnsafePtr).V2.IntVal)), UnsafePtr: nil}, gopurs_runtime.Float((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V4)}.UnsafePtr).V3.FloatVal()), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[uint32, float64]](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeNode(), gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(__local_var_0.IntVal)), UnsafePtr: nil}, gopurs_runtime.Float(__local_var_1.FloatVal()), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[uint32, float64]](__local_var_2))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V4)}.UnsafePtr).V4)})))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[uint32, float64]](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeNode(), gopurs_runtime.Value{Type: 9, IntVal: int64(uint32((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V2.IntVal)), UnsafePtr: nil}, gopurs_runtime.Float((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V3.FloatVal()), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V4)}.UnsafePtr).V5)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V5)})))}))
goto end_branch_8
} else {

}
}
{
__t8 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[uint32, float64]](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeNode(), gopurs_runtime.Value{Type: 9, IntVal: int64(uint32((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V2.IntVal)), UnsafePtr: nil}, gopurs_runtime.Float((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V3.FloatVal()), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[uint32, float64]](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeNode(), gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(__local_var_0.IntVal)), UnsafePtr: nil}, gopurs_runtime.Float(__local_var_1.FloatVal()), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[uint32, float64]](__local_var_2))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V4)})))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V5)}))
}
end_branch_8:
__t9 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(__t8)}
goto end_branch_9
} else {

}
}
{
__t9 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[uint32, float64]](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeNode(), gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(__local_var_0.IntVal)), UnsafePtr: nil}, gopurs_runtime.Float(__local_var_1.FloatVal()), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[uint32, float64]](__local_var_2))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[uint32, float64]](__local_var_3))})))}
}
end_branch_9:
__t37 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[uint32, float64]](__t9))}
goto end_branch_37
} else {

}
}
{
if (__local_var_2.Type == 9 && __local_var_2.IntVal == 324739070 && __local_var_2.UnsafePtr != nil) {
var __t36 *Constructor_Data_Map_Internal_Node[uint32, float64]
{
if (__local_var_3.Type == 9 && __local_var_3.IntVal == 324739070 && __local_var_3.UnsafePtr != nil) {
var __t26 *Constructor_Data_Map_Internal_Node[uint32, float64]
{
var __t10 bool
{
if ((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V0) > (((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V0) + (1)) {
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
var __t17 *Constructor_Data_Map_Internal_Node[uint32, float64]
{
var __t_tag_11 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V4)}
var __t_and_16 bool = false
if (__t_tag_11.Type == 9 && __t_tag_11.IntVal == 324739070 && __t_tag_11.UnsafePtr != nil) {

var __t15 bool
{
var __t14 gopurs_runtime.Value
{
var __t_tag_12 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V5)}
if (__t_tag_12.Type == 9 && __t_tag_12.IntVal == 324739070 && __t_tag_12.UnsafePtr == nil) {
__t14 = gopurs_runtime.Int(0)
goto end_branch_14
} else {

}
}
{
var __t_tag_13 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V5)}
if (__t_tag_13.Type == 9 && __t_tag_13.IntVal == 324739070 && __t_tag_13.UnsafePtr != nil) {
__t14 = gopurs_runtime.Int((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V5)}.UnsafePtr).V0)
goto end_branch_14
} else {

}
}
{
__t14 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_14:
if ((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V4)}.UnsafePtr).V0) > (__t14.IntVal) {
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
__t17 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[uint32, float64]](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeNode(), gopurs_runtime.Value{Type: 9, IntVal: int64(uint32((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V4)}.UnsafePtr).V2.IntVal)), UnsafePtr: nil}, gopurs_runtime.Float((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V4)}.UnsafePtr).V3.FloatVal()), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[uint32, float64]](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeNode(), gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(__local_var_0.IntVal)), UnsafePtr: nil}, gopurs_runtime.Float(__local_var_1.FloatVal()), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[uint32, float64]](__local_var_2))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V4)}.UnsafePtr).V4)})))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[uint32, float64]](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeNode(), gopurs_runtime.Value{Type: 9, IntVal: int64(uint32((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V2.IntVal)), UnsafePtr: nil}, gopurs_runtime.Float((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V3.FloatVal()), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V4)}.UnsafePtr).V5)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V5)})))}))
goto end_branch_17
} else {

}
}
{
__t17 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[uint32, float64]](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeNode(), gopurs_runtime.Value{Type: 9, IntVal: int64(uint32((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V2.IntVal)), UnsafePtr: nil}, gopurs_runtime.Float((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V3.FloatVal()), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[uint32, float64]](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeNode(), gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(__local_var_0.IntVal)), UnsafePtr: nil}, gopurs_runtime.Float(__local_var_1.FloatVal()), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[uint32, float64]](__local_var_2))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V4)})))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V5)}))
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
if ((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V0) > (((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V0) + (1)) {
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
var __t25 *Constructor_Data_Map_Internal_Node[uint32, float64]
{
var __t_tag_19 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V5)}
var __t_and_24 bool = false
if (__t_tag_19.Type == 9 && __t_tag_19.IntVal == 324739070 && __t_tag_19.UnsafePtr != nil) {

var __t23 bool
{
var __t22 gopurs_runtime.Value
{
var __t_tag_20 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V4)}
if (__t_tag_20.Type == 9 && __t_tag_20.IntVal == 324739070 && __t_tag_20.UnsafePtr == nil) {
__t22 = gopurs_runtime.Int(0)
goto end_branch_22
} else {

}
}
{
var __t_tag_21 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V4)}
if (__t_tag_21.Type == 9 && __t_tag_21.IntVal == 324739070 && __t_tag_21.UnsafePtr != nil) {
__t22 = gopurs_runtime.Int((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V4)}.UnsafePtr).V0)
goto end_branch_22
} else {

}
}
{
__t22 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_22:
if (__t22.IntVal) > ((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V5)}.UnsafePtr).V0) {
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
__t25 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[uint32, float64]](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeNode(), gopurs_runtime.Value{Type: 9, IntVal: int64(uint32((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V5)}.UnsafePtr).V2.IntVal)), UnsafePtr: nil}, gopurs_runtime.Float((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V5)}.UnsafePtr).V3.FloatVal()), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[uint32, float64]](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeNode(), gopurs_runtime.Value{Type: 9, IntVal: int64(uint32((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V2.IntVal)), UnsafePtr: nil}, gopurs_runtime.Float((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V3.FloatVal()), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V4)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V5)}.UnsafePtr).V4)})))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[uint32, float64]](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeNode(), gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(__local_var_0.IntVal)), UnsafePtr: nil}, gopurs_runtime.Float(__local_var_1.FloatVal()), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V5)}.UnsafePtr).V5)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[uint32, float64]](__local_var_3))})))}))
goto end_branch_25
} else {

}
}
{
__t25 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[uint32, float64]](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeNode(), gopurs_runtime.Value{Type: 9, IntVal: int64(uint32((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V2.IntVal)), UnsafePtr: nil}, gopurs_runtime.Float((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V3.FloatVal()), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V4)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[uint32, float64]](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeNode(), gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(__local_var_0.IntVal)), UnsafePtr: nil}, gopurs_runtime.Float(__local_var_1.FloatVal()), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V5)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[uint32, float64]](__local_var_3))})))}))
}
end_branch_25:
__t26 = __t25
goto end_branch_26
} else {

}
}
{
__t26 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[uint32, float64]](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeNode(), gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(__local_var_0.IntVal)), UnsafePtr: nil}, gopurs_runtime.Float(__local_var_1.FloatVal()), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[uint32, float64]](__local_var_2))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[uint32, float64]](__local_var_3))}))
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
if ((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V0) > (1) {
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
var __t35 *Constructor_Data_Map_Internal_Node[uint32, float64]
{
var __t_tag_29 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V5)}
var __t_and_34 bool = false
if (__t_tag_29.Type == 9 && __t_tag_29.IntVal == 324739070 && __t_tag_29.UnsafePtr != nil) {

var __t33 bool
{
var __t32 gopurs_runtime.Value
{
var __t_tag_30 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V4)}
if (__t_tag_30.Type == 9 && __t_tag_30.IntVal == 324739070 && __t_tag_30.UnsafePtr == nil) {
__t32 = gopurs_runtime.Int(0)
goto end_branch_32
} else {

}
}
{
var __t_tag_31 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V4)}
if (__t_tag_31.Type == 9 && __t_tag_31.IntVal == 324739070 && __t_tag_31.UnsafePtr != nil) {
__t32 = gopurs_runtime.Int((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V4)}.UnsafePtr).V0)
goto end_branch_32
} else {

}
}
{
__t32 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_32:
if (__t32.IntVal) > ((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V5)}.UnsafePtr).V0) {
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
__t35 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[uint32, float64]](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeNode(), gopurs_runtime.Value{Type: 9, IntVal: int64(uint32((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V5)}.UnsafePtr).V2.IntVal)), UnsafePtr: nil}, gopurs_runtime.Float((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V5)}.UnsafePtr).V3.FloatVal()), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[uint32, float64]](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeNode(), gopurs_runtime.Value{Type: 9, IntVal: int64(uint32((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V2.IntVal)), UnsafePtr: nil}, gopurs_runtime.Float((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V3.FloatVal()), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V4)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V5)}.UnsafePtr).V4)})))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[uint32, float64]](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeNode(), gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(__local_var_0.IntVal)), UnsafePtr: nil}, gopurs_runtime.Float(__local_var_1.FloatVal()), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V5)}.UnsafePtr).V5)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[uint32, float64]](__local_var_3))})))}))
goto end_branch_35
} else {

}
}
{
__t35 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[uint32, float64]](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeNode(), gopurs_runtime.Value{Type: 9, IntVal: int64(uint32((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V2.IntVal)), UnsafePtr: nil}, gopurs_runtime.Float((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V3.FloatVal()), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V4)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[uint32, float64]](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeNode(), gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(__local_var_0.IntVal)), UnsafePtr: nil}, gopurs_runtime.Float(__local_var_1.FloatVal()), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V5)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[uint32, float64]](__local_var_3))})))}))
}
end_branch_35:
__t36 = __t35
goto end_branch_36
} else {

}
}
{
__t36 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[uint32, float64]](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeNode(), gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(__local_var_0.IntVal)), UnsafePtr: nil}, gopurs_runtime.Float(__local_var_1.FloatVal()), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[uint32, float64]](__local_var_2))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[uint32, float64]](__local_var_3))}))
}
end_branch_36:
__t37 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(__t36)}
goto end_branch_37
} else {

}
}
{
__t37 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_37:
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[uint32, float64]](__t37))}
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
var __t37 gopurs_runtime.Value
{
if (__local_var_2.Type == 9 && __local_var_2.IntVal == 324739070 && __local_var_2.UnsafePtr == nil) {
var __t9 gopurs_runtime.Value
{
if (__local_var_3.Type == 9 && __local_var_3.IntVal == 324739070 && __local_var_3.UnsafePtr == nil) {
__t9 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(&Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]{1, 1, 1, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](__local_var_0))}, __local_var_1, gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(nil))}), gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(nil))})})}
goto end_branch_9
} else {

}
}
{
var __t_and_1 bool = false
if (__local_var_3.Type == 9 && __local_var_3.IntVal == 324739070 && __local_var_3.UnsafePtr != nil) {

var __t0 bool
{
if ((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V0) > (1) {
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
var __t8 *Constructor_Data_Map_Internal_Node[*Constructor_Data_Maybe_Just[gopurs_runtime.Value], gopurs_runtime.Value]
{
var __t_tag_2 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V4)}
var __t_and_7 bool = false
if (__t_tag_2.Type == 9 && __t_tag_2.IntVal == 324739070 && __t_tag_2.UnsafePtr != nil) {

var __t6 bool
{
var __t5 gopurs_runtime.Value
{
var __t_tag_3 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V5)}
if (__t_tag_3.Type == 9 && __t_tag_3.IntVal == 324739070 && __t_tag_3.UnsafePtr == nil) {
__t5 = gopurs_runtime.Int(0)
goto end_branch_5
} else {

}
}
{
var __t_tag_4 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V5)}
if (__t_tag_4.Type == 9 && __t_tag_4.IntVal == 324739070 && __t_tag_4.UnsafePtr != nil) {
__t5 = gopurs_runtime.Int((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V5)}.UnsafePtr).V0)
goto end_branch_5
} else {

}
}
{
__t5 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_5:
if ((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V4)}.UnsafePtr).V0) > (__t5.IntVal) {
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
__t8 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[*Constructor_Data_Maybe_Just[gopurs_runtime.Value], gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeNode(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]]((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V4)}.UnsafePtr).V2))}, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V4)}.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[*Constructor_Data_Maybe_Just[gopurs_runtime.Value], gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeNode(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](__local_var_0))}, __local_var_1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[*Constructor_Data_Maybe_Just[gopurs_runtime.Value], gopurs_runtime.Value]](__local_var_2))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V4)}.UnsafePtr).V4)})))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[*Constructor_Data_Maybe_Just[gopurs_runtime.Value], gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeNode(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]]((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V2))}, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V4)}.UnsafePtr).V5)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V5)})))}))
goto end_branch_8
} else {

}
}
{
__t8 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[*Constructor_Data_Maybe_Just[gopurs_runtime.Value], gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeNode(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]]((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V2))}, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[*Constructor_Data_Maybe_Just[gopurs_runtime.Value], gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeNode(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](__local_var_0))}, __local_var_1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[*Constructor_Data_Maybe_Just[gopurs_runtime.Value], gopurs_runtime.Value]](__local_var_2))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V4)})))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V5)}))
}
end_branch_8:
__t9 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(__t8)}
goto end_branch_9
} else {

}
}
{
__t9 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[*Constructor_Data_Maybe_Just[gopurs_runtime.Value], gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeNode(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](__local_var_0))}, __local_var_1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[*Constructor_Data_Maybe_Just[gopurs_runtime.Value], gopurs_runtime.Value]](__local_var_2))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[*Constructor_Data_Maybe_Just[gopurs_runtime.Value], gopurs_runtime.Value]](__local_var_3))})))}
}
end_branch_9:
__t37 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[*Constructor_Data_Maybe_Just[gopurs_runtime.Value], gopurs_runtime.Value]](__t9))}
goto end_branch_37
} else {

}
}
{
if (__local_var_2.Type == 9 && __local_var_2.IntVal == 324739070 && __local_var_2.UnsafePtr != nil) {
var __t36 *Constructor_Data_Map_Internal_Node[*Constructor_Data_Maybe_Just[gopurs_runtime.Value], gopurs_runtime.Value]
{
if (__local_var_3.Type == 9 && __local_var_3.IntVal == 324739070 && __local_var_3.UnsafePtr != nil) {
var __t26 *Constructor_Data_Map_Internal_Node[*Constructor_Data_Maybe_Just[gopurs_runtime.Value], gopurs_runtime.Value]
{
var __t10 bool
{
if ((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V0) > (((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V0) + (1)) {
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
var __t17 *Constructor_Data_Map_Internal_Node[*Constructor_Data_Maybe_Just[gopurs_runtime.Value], gopurs_runtime.Value]
{
var __t_tag_11 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V4)}
var __t_and_16 bool = false
if (__t_tag_11.Type == 9 && __t_tag_11.IntVal == 324739070 && __t_tag_11.UnsafePtr != nil) {

var __t15 bool
{
var __t14 gopurs_runtime.Value
{
var __t_tag_12 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V5)}
if (__t_tag_12.Type == 9 && __t_tag_12.IntVal == 324739070 && __t_tag_12.UnsafePtr == nil) {
__t14 = gopurs_runtime.Int(0)
goto end_branch_14
} else {

}
}
{
var __t_tag_13 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V5)}
if (__t_tag_13.Type == 9 && __t_tag_13.IntVal == 324739070 && __t_tag_13.UnsafePtr != nil) {
__t14 = gopurs_runtime.Int((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V5)}.UnsafePtr).V0)
goto end_branch_14
} else {

}
}
{
__t14 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_14:
if ((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V4)}.UnsafePtr).V0) > (__t14.IntVal) {
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
__t17 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[*Constructor_Data_Maybe_Just[gopurs_runtime.Value], gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeNode(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]]((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V4)}.UnsafePtr).V2))}, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V4)}.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[*Constructor_Data_Maybe_Just[gopurs_runtime.Value], gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeNode(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](__local_var_0))}, __local_var_1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[*Constructor_Data_Maybe_Just[gopurs_runtime.Value], gopurs_runtime.Value]](__local_var_2))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V4)}.UnsafePtr).V4)})))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[*Constructor_Data_Maybe_Just[gopurs_runtime.Value], gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeNode(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]]((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V2))}, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V4)}.UnsafePtr).V5)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V5)})))}))
goto end_branch_17
} else {

}
}
{
__t17 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[*Constructor_Data_Maybe_Just[gopurs_runtime.Value], gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeNode(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]]((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V2))}, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[*Constructor_Data_Maybe_Just[gopurs_runtime.Value], gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeNode(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](__local_var_0))}, __local_var_1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[*Constructor_Data_Maybe_Just[gopurs_runtime.Value], gopurs_runtime.Value]](__local_var_2))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V4)})))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V5)}))
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
if ((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V0) > (((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V0) + (1)) {
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
var __t25 *Constructor_Data_Map_Internal_Node[*Constructor_Data_Maybe_Just[gopurs_runtime.Value], gopurs_runtime.Value]
{
var __t_tag_19 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V5)}
var __t_and_24 bool = false
if (__t_tag_19.Type == 9 && __t_tag_19.IntVal == 324739070 && __t_tag_19.UnsafePtr != nil) {

var __t23 bool
{
var __t22 gopurs_runtime.Value
{
var __t_tag_20 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V4)}
if (__t_tag_20.Type == 9 && __t_tag_20.IntVal == 324739070 && __t_tag_20.UnsafePtr == nil) {
__t22 = gopurs_runtime.Int(0)
goto end_branch_22
} else {

}
}
{
var __t_tag_21 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V4)}
if (__t_tag_21.Type == 9 && __t_tag_21.IntVal == 324739070 && __t_tag_21.UnsafePtr != nil) {
__t22 = gopurs_runtime.Int((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V4)}.UnsafePtr).V0)
goto end_branch_22
} else {

}
}
{
__t22 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_22:
if (__t22.IntVal) > ((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V5)}.UnsafePtr).V0) {
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
__t25 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[*Constructor_Data_Maybe_Just[gopurs_runtime.Value], gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeNode(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]]((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V5)}.UnsafePtr).V2))}, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V5)}.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[*Constructor_Data_Maybe_Just[gopurs_runtime.Value], gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeNode(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]]((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V2))}, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V4)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V5)}.UnsafePtr).V4)})))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[*Constructor_Data_Maybe_Just[gopurs_runtime.Value], gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeNode(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](__local_var_0))}, __local_var_1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V5)}.UnsafePtr).V5)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[*Constructor_Data_Maybe_Just[gopurs_runtime.Value], gopurs_runtime.Value]](__local_var_3))})))}))
goto end_branch_25
} else {

}
}
{
__t25 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[*Constructor_Data_Maybe_Just[gopurs_runtime.Value], gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeNode(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]]((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V2))}, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V4)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[*Constructor_Data_Maybe_Just[gopurs_runtime.Value], gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeNode(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](__local_var_0))}, __local_var_1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V5)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[*Constructor_Data_Maybe_Just[gopurs_runtime.Value], gopurs_runtime.Value]](__local_var_3))})))}))
}
end_branch_25:
__t26 = __t25
goto end_branch_26
} else {

}
}
{
__t26 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[*Constructor_Data_Maybe_Just[gopurs_runtime.Value], gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeNode(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](__local_var_0))}, __local_var_1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[*Constructor_Data_Maybe_Just[gopurs_runtime.Value], gopurs_runtime.Value]](__local_var_2))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[*Constructor_Data_Maybe_Just[gopurs_runtime.Value], gopurs_runtime.Value]](__local_var_3))}))
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
if ((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V0) > (1) {
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
var __t35 *Constructor_Data_Map_Internal_Node[*Constructor_Data_Maybe_Just[gopurs_runtime.Value], gopurs_runtime.Value]
{
var __t_tag_29 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V5)}
var __t_and_34 bool = false
if (__t_tag_29.Type == 9 && __t_tag_29.IntVal == 324739070 && __t_tag_29.UnsafePtr != nil) {

var __t33 bool
{
var __t32 gopurs_runtime.Value
{
var __t_tag_30 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V4)}
if (__t_tag_30.Type == 9 && __t_tag_30.IntVal == 324739070 && __t_tag_30.UnsafePtr == nil) {
__t32 = gopurs_runtime.Int(0)
goto end_branch_32
} else {

}
}
{
var __t_tag_31 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V4)}
if (__t_tag_31.Type == 9 && __t_tag_31.IntVal == 324739070 && __t_tag_31.UnsafePtr != nil) {
__t32 = gopurs_runtime.Int((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V4)}.UnsafePtr).V0)
goto end_branch_32
} else {

}
}
{
__t32 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_32:
if (__t32.IntVal) > ((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V5)}.UnsafePtr).V0) {
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
__t35 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[*Constructor_Data_Maybe_Just[gopurs_runtime.Value], gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeNode(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]]((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V5)}.UnsafePtr).V2))}, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V5)}.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[*Constructor_Data_Maybe_Just[gopurs_runtime.Value], gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeNode(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]]((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V2))}, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V4)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V5)}.UnsafePtr).V4)})))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[*Constructor_Data_Maybe_Just[gopurs_runtime.Value], gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeNode(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](__local_var_0))}, __local_var_1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V5)}.UnsafePtr).V5)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[*Constructor_Data_Maybe_Just[gopurs_runtime.Value], gopurs_runtime.Value]](__local_var_3))})))}))
goto end_branch_35
} else {

}
}
{
__t35 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[*Constructor_Data_Maybe_Just[gopurs_runtime.Value], gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeNode(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]]((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V2))}, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V4)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[*Constructor_Data_Maybe_Just[gopurs_runtime.Value], gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeNode(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](__local_var_0))}, __local_var_1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V5)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[*Constructor_Data_Maybe_Just[gopurs_runtime.Value], gopurs_runtime.Value]](__local_var_3))})))}))
}
end_branch_35:
__t36 = __t35
goto end_branch_36
} else {

}
}
{
__t36 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[*Constructor_Data_Maybe_Just[gopurs_runtime.Value], gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeNode(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](__local_var_0))}, __local_var_1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[*Constructor_Data_Maybe_Just[gopurs_runtime.Value], gopurs_runtime.Value]](__local_var_2))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[*Constructor_Data_Maybe_Just[gopurs_runtime.Value], gopurs_runtime.Value]](__local_var_3))}))
}
end_branch_36:
__t37 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(__t36)}
goto end_branch_37
} else {

}
}
{
__t37 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_37:
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[*Constructor_Data_Maybe_Just[gopurs_runtime.Value], gopurs_runtime.Value]](__t37))}
}

func Call_Data_Map_Internal_unsafeDifference__4097927905(__local_var_0_loop gopurs_runtime.Value, __local_var_1_loop gopurs_runtime.Value, __local_var_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __local_var_0 gopurs_runtime.Value = __local_var_0_loop
_ = __local_var_0
var __local_var_1 gopurs_runtime.Value = __local_var_1_loop
_ = __local_var_1
var __local_var_2 gopurs_runtime.Value = __local_var_2_loop
_ = __local_var_2
var __t1 gopurs_runtime.Value
{
if (__local_var_1.Type == 9 && __local_var_1.IntVal == 324739070 && __local_var_1.UnsafePtr == nil) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(nil))}
goto end_branch_1
} else {

}
}
{
if (__local_var_2.Type == 9 && __local_var_2.IntVal == 324739070 && __local_var_2.UnsafePtr == nil) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__local_var_1))}
goto end_branch_1
} else {

}
}
{
if (__local_var_2.Type == 9 && __local_var_2.IntVal == 324739070 && __local_var_2.UnsafePtr != nil) {
// TAST (Let): v_3_0 -> gopurs_runtime.Value
v_3_0 := gopurs_runtime.UncurriedApp3(Get_Data_Map_Internal_unsafeSplit(), __local_var_0, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V2, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__local_var_1))})
_ = v_3_0
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp2(Get_Data_Map_Internal_unsafeJoinNodes(), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp3(Get_Data_Map_Internal_unsafeDifference(), __local_var_0, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Split[gopurs_runtime.Value, gopurs_runtime.Value])(v_3_0.UnsafePtr).V1)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V4)})))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp3(Get_Data_Map_Internal_unsafeDifference(), __local_var_0, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Split[gopurs_runtime.Value, gopurs_runtime.Value])(v_3_0.UnsafePtr).V2)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V5)})))})))}
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__t1))}
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
var __t6 gopurs_runtime.Value
{
if (__local_var_2.Type == 9 && __local_var_2.IntVal == 324739070 && __local_var_2.UnsafePtr == nil) {
__t6 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(nil))}
goto end_branch_6
} else {

}
}
{
if (__local_var_3.Type == 9 && __local_var_3.IntVal == 324739070 && __local_var_3.UnsafePtr == nil) {
__t6 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(nil))}
goto end_branch_6
} else {

}
}
{
if (__local_var_3.Type == 9 && __local_var_3.IntVal == 324739070 && __local_var_3.UnsafePtr != nil) {
// TAST (Let): v_4_0 -> gopurs_runtime.Value
v_4_0 := gopurs_runtime.UncurriedApp3(Get_Data_Map_Internal_unsafeSplit(), __local_var_0, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V2, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__local_var_2))})
_ = v_4_0
// TAST (Let): l_prime_5_1 -> gopurs_runtime.Value
l_prime_5_1 := gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeIntersectionWith(), __local_var_0, __local_var_1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Split[gopurs_runtime.Value, gopurs_runtime.Value])(v_4_0.UnsafePtr).V1)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V4)})
_ = l_prime_5_1
// TAST (Let): r_prime_6_2 -> gopurs_runtime.Value
r_prime_6_2 := gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeIntersectionWith(), __local_var_0, __local_var_1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Split[gopurs_runtime.Value, gopurs_runtime.Value])(v_4_0.UnsafePtr).V2)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V5)})
_ = r_prime_6_2
var __t5 gopurs_runtime.Value
{
var __t_tag_3 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Split[gopurs_runtime.Value, gopurs_runtime.Value])(v_4_0.UnsafePtr).V0)}
if (__t_tag_3.Type == 9 && __t_tag_3.IntVal == 930809136 && __t_tag_3.UnsafePtr != nil) {
__t5 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeBalancedNode(), (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V2, gopurs_runtime.Apply2(__local_var_1, (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Split[gopurs_runtime.Value, gopurs_runtime.Value])(v_4_0.UnsafePtr).V0)}.UnsafePtr).V0, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V3), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](l_prime_5_1))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](r_prime_6_2))})))}
goto end_branch_5
} else {

}
}
{
var __t_tag_4 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Split[gopurs_runtime.Value, gopurs_runtime.Value])(v_4_0.UnsafePtr).V0)}
if (__t_tag_4.Type == 9 && __t_tag_4.IntVal == 930809136 && __t_tag_4.UnsafePtr == nil) {
__t5 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp2(Get_Data_Map_Internal_unsafeJoinNodes(), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](l_prime_5_1))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](r_prime_6_2))})))}
goto end_branch_5
} else {

}
}
{
__t5 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_5:
__t6 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__t5))}
goto end_branch_6
} else {

}
}
{
__t6 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_6:
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__t6))}
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
var __t6 gopurs_runtime.Value
{
if (__local_var_2.Type == 9 && __local_var_2.IntVal == 324739070 && __local_var_2.UnsafePtr == nil) {
__t6 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(nil))}
goto end_branch_6
} else {

}
}
{
if (__local_var_3.Type == 9 && __local_var_3.IntVal == 324739070 && __local_var_3.UnsafePtr == nil) {
__t6 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(nil))}
goto end_branch_6
} else {

}
}
{
if (__local_var_3.Type == 9 && __local_var_3.IntVal == 324739070 && __local_var_3.UnsafePtr != nil) {
// TAST (Let): v_4_0 -> gopurs_runtime.Value
v_4_0 := gopurs_runtime.UncurriedApp3(Get_Data_Map_Internal_unsafeSplit(), __local_var_0, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V2, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__local_var_2))})
_ = v_4_0
// TAST (Let): l_prime_5_1 -> gopurs_runtime.Value
l_prime_5_1 := gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeIntersectionWith(), __local_var_0, __local_var_1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Split[gopurs_runtime.Value, gopurs_runtime.Value])(v_4_0.UnsafePtr).V1)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V4)})
_ = l_prime_5_1
// TAST (Let): r_prime_6_2 -> gopurs_runtime.Value
r_prime_6_2 := gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeIntersectionWith(), __local_var_0, __local_var_1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Split[gopurs_runtime.Value, gopurs_runtime.Value])(v_4_0.UnsafePtr).V2)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V5)})
_ = r_prime_6_2
var __t5 gopurs_runtime.Value
{
var __t_tag_3 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Split[gopurs_runtime.Value, gopurs_runtime.Value])(v_4_0.UnsafePtr).V0)}
if (__t_tag_3.Type == 9 && __t_tag_3.IntVal == 930809136 && __t_tag_3.UnsafePtr != nil) {
__t5 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeBalancedNode(), (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V2, gopurs_runtime.Apply2(__local_var_1, (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Split[gopurs_runtime.Value, gopurs_runtime.Value])(v_4_0.UnsafePtr).V0)}.UnsafePtr).V0, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V3), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](l_prime_5_1))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](r_prime_6_2))})))}
goto end_branch_5
} else {

}
}
{
var __t_tag_4 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Split[gopurs_runtime.Value, gopurs_runtime.Value])(v_4_0.UnsafePtr).V0)}
if (__t_tag_4.Type == 9 && __t_tag_4.IntVal == 930809136 && __t_tag_4.UnsafePtr == nil) {
__t5 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp2(Get_Data_Map_Internal_unsafeJoinNodes(), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](l_prime_5_1))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](r_prime_6_2))})))}
goto end_branch_5
} else {

}
}
{
__t5 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_5:
__t6 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__t5))}
goto end_branch_6
} else {

}
}
{
__t6 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_6:
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__t6))}
}

func Call_Data_Map_Internal_unsafeJoinNodes__2531831408(__local_var_0_loop gopurs_runtime.Value, __local_var_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __local_var_0 gopurs_runtime.Value = __local_var_0_loop
_ = __local_var_0
var __local_var_1 gopurs_runtime.Value = __local_var_1_loop
_ = __local_var_1
var __t1 gopurs_runtime.Value
{
if (__local_var_0.Type == 9 && __local_var_0.IntVal == 324739070 && __local_var_0.UnsafePtr == nil) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__local_var_1))}
goto end_branch_1
} else {

}
}
{
if (__local_var_0.Type == 9 && __local_var_0.IntVal == 324739070 && __local_var_0.UnsafePtr != nil) {
// TAST (Let): v2_2_0 -> gopurs_runtime.Value
v2_2_0 := gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeSplitLast(), (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_0.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_0.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_0.UnsafePtr).V4)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_0.UnsafePtr).V5)})
_ = v2_2_0
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeBalancedNode(), (*Constructor_Data_Map_Internal_SplitLast[gopurs_runtime.Value, gopurs_runtime.Value])(v2_2_0.UnsafePtr).V0, (*Constructor_Data_Map_Internal_SplitLast[gopurs_runtime.Value, gopurs_runtime.Value])(v2_2_0.UnsafePtr).V1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_SplitLast[gopurs_runtime.Value, gopurs_runtime.Value])(v2_2_0.UnsafePtr).V2)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__local_var_1))})))}
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__t1))}
}

func Call_Data_Map_Internal_unsafeJoinNodes__3967876672(__local_var_0_loop gopurs_runtime.Value, __local_var_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __local_var_0 gopurs_runtime.Value = __local_var_0_loop
_ = __local_var_0
var __local_var_1 gopurs_runtime.Value = __local_var_1_loop
_ = __local_var_1
var __t1 gopurs_runtime.Value
{
if (__local_var_0.Type == 9 && __local_var_0.IntVal == 324739070 && __local_var_0.UnsafePtr == nil) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__local_var_1))}
goto end_branch_1
} else {

}
}
{
if (__local_var_0.Type == 9 && __local_var_0.IntVal == 324739070 && __local_var_0.UnsafePtr != nil) {
// TAST (Let): v2_2_0 -> gopurs_runtime.Value
v2_2_0 := gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeSplitLast(), (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_0.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_0.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_0.UnsafePtr).V4)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_0.UnsafePtr).V5)})
_ = v2_2_0
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeBalancedNode(), (*Constructor_Data_Map_Internal_SplitLast[gopurs_runtime.Value, gopurs_runtime.Value])(v2_2_0.UnsafePtr).V0, (*Constructor_Data_Map_Internal_SplitLast[gopurs_runtime.Value, gopurs_runtime.Value])(v2_2_0.UnsafePtr).V1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_SplitLast[gopurs_runtime.Value, gopurs_runtime.Value])(v2_2_0.UnsafePtr).V2)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__local_var_1))})))}
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__t1))}
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
var __t4 gopurs_runtime.Value
{
if (__local_var_2.Type == 9 && __local_var_2.IntVal == 324739070 && __local_var_2.UnsafePtr == nil) {
var __t0 gopurs_runtime.Value
{
if (__local_var_3.Type == 9 && __local_var_3.IntVal == 324739070 && __local_var_3.UnsafePtr == nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(&Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]{1, 1, 1, __local_var_0, __local_var_1, gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__local_var_2), gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__local_var_3)})}
goto end_branch_0
} else {

}
}
{
if (__local_var_3.Type == 9 && __local_var_3.IntVal == 324739070 && __local_var_3.UnsafePtr != nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(&Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]{1, (1) + ((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V0), (1) + ((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V1), __local_var_0, __local_var_1, gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__local_var_2), gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__local_var_3)})}
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
__t4 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__t0))}
goto end_branch_4
} else {

}
}
{
if (__local_var_2.Type == 9 && __local_var_2.IntVal == 324739070 && __local_var_2.UnsafePtr != nil) {
var __t3 gopurs_runtime.Value
{
if (__local_var_3.Type == 9 && __local_var_3.IntVal == 324739070 && __local_var_3.UnsafePtr == nil) {
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(&Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]{1, (1) + ((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V0), (1) + ((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V1), __local_var_0, __local_var_1, gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__local_var_2), gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__local_var_3)})}
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
if ((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V0) > ((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V0) {
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
__t2 = (1) + ((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V0)
goto end_branch_2
} else {

}
}
{
__t2 = (1) + ((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V0)
}
end_branch_2:
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(&Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]{1, __t2, ((1) + ((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V1)) + ((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V1), __local_var_0, __local_var_1, gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__local_var_2), gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__local_var_3)})}
goto end_branch_3
} else {

}
}
{
__t3 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_3:
__t4 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__t3))}
goto end_branch_4
} else {

}
}
{
__t4 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_4:
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__t4))}
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
var __t4 gopurs_runtime.Value
{
if (__local_var_2.Type == 9 && __local_var_2.IntVal == 324739070 && __local_var_2.UnsafePtr == nil) {
var __t0 gopurs_runtime.Value
{
if (__local_var_3.Type == 9 && __local_var_3.IntVal == 324739070 && __local_var_3.UnsafePtr == nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(&Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]{1, 1, 1, __local_var_0, __local_var_1, gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__local_var_2), gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__local_var_3)})}
goto end_branch_0
} else {

}
}
{
if (__local_var_3.Type == 9 && __local_var_3.IntVal == 324739070 && __local_var_3.UnsafePtr != nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(&Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]{1, (1) + ((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V0), (1) + ((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V1), __local_var_0, __local_var_1, gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__local_var_2), gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__local_var_3)})}
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
__t4 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__t0))}
goto end_branch_4
} else {

}
}
{
if (__local_var_2.Type == 9 && __local_var_2.IntVal == 324739070 && __local_var_2.UnsafePtr != nil) {
var __t3 gopurs_runtime.Value
{
if (__local_var_3.Type == 9 && __local_var_3.IntVal == 324739070 && __local_var_3.UnsafePtr == nil) {
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(&Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]{1, (1) + ((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V0), (1) + ((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V1), __local_var_0, __local_var_1, gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__local_var_2), gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__local_var_3)})}
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
if ((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V0) > ((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V0) {
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
__t2 = (1) + ((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V0)
goto end_branch_2
} else {

}
}
{
__t2 = (1) + ((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V0)
}
end_branch_2:
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(&Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]{1, __t2, ((1) + ((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V1)) + ((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V1), __local_var_0, __local_var_1, gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__local_var_2), gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__local_var_3)})}
goto end_branch_3
} else {

}
}
{
__t3 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_3:
__t4 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__t3))}
goto end_branch_4
} else {

}
}
{
__t4 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_4:
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__t4))}
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
var __t4 gopurs_runtime.Value
{
if (__local_var_2.Type == 9 && __local_var_2.IntVal == 324739070 && __local_var_2.UnsafePtr == nil) {
var __t0 gopurs_runtime.Value
{
if (__local_var_3.Type == 9 && __local_var_3.IntVal == 324739070 && __local_var_3.UnsafePtr == nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(&Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]{1, 1, 1, __local_var_0, __local_var_1, gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__local_var_2), gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__local_var_3)})}
goto end_branch_0
} else {

}
}
{
if (__local_var_3.Type == 9 && __local_var_3.IntVal == 324739070 && __local_var_3.UnsafePtr != nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(&Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]{1, (1) + ((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V0), (1) + ((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V1), __local_var_0, __local_var_1, gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__local_var_2), gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__local_var_3)})}
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
__t4 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__t0))}
goto end_branch_4
} else {

}
}
{
if (__local_var_2.Type == 9 && __local_var_2.IntVal == 324739070 && __local_var_2.UnsafePtr != nil) {
var __t3 gopurs_runtime.Value
{
if (__local_var_3.Type == 9 && __local_var_3.IntVal == 324739070 && __local_var_3.UnsafePtr == nil) {
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(&Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]{1, (1) + ((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V0), (1) + ((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V1), __local_var_0, __local_var_1, gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__local_var_2), gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__local_var_3)})}
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
if ((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V0) > ((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V0) {
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
__t2 = (1) + ((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V0)
goto end_branch_2
} else {

}
}
{
__t2 = (1) + ((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V0)
}
end_branch_2:
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(&Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]{1, __t2, ((1) + ((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V1)) + ((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V1), __local_var_0, __local_var_1, gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__local_var_2), gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__local_var_3)})}
goto end_branch_3
} else {

}
}
{
__t3 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_3:
__t4 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__t3))}
goto end_branch_4
} else {

}
}
{
__t4 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_4:
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__t4))}
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
var __t4 gopurs_runtime.Value
{
if (__local_var_2.Type == 9 && __local_var_2.IntVal == 324739070 && __local_var_2.UnsafePtr == nil) {
var __t0 gopurs_runtime.Value
{
if (__local_var_3.Type == 9 && __local_var_3.IntVal == 324739070 && __local_var_3.UnsafePtr == nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(&Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]{1, 1, 1, gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(__local_var_0.IntVal)), UnsafePtr: nil}, gopurs_runtime.Float(__local_var_1.FloatVal()), gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[uint32, float64]](__local_var_2))}), gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[uint32, float64]](__local_var_3))})})}
goto end_branch_0
} else {

}
}
{
if (__local_var_3.Type == 9 && __local_var_3.IntVal == 324739070 && __local_var_3.UnsafePtr != nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(&Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]{1, (1) + ((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V0), (1) + ((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V1), gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(__local_var_0.IntVal)), UnsafePtr: nil}, gopurs_runtime.Float(__local_var_1.FloatVal()), gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[uint32, float64]](__local_var_2))}), gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[uint32, float64]](__local_var_3))})})}
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
__t4 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[uint32, float64]](__t0))}
goto end_branch_4
} else {

}
}
{
if (__local_var_2.Type == 9 && __local_var_2.IntVal == 324739070 && __local_var_2.UnsafePtr != nil) {
var __t3 gopurs_runtime.Value
{
if (__local_var_3.Type == 9 && __local_var_3.IntVal == 324739070 && __local_var_3.UnsafePtr == nil) {
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(&Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]{1, (1) + ((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V0), (1) + ((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V1), gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(__local_var_0.IntVal)), UnsafePtr: nil}, gopurs_runtime.Float(__local_var_1.FloatVal()), gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[uint32, float64]](__local_var_2))}), gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[uint32, float64]](__local_var_3))})})}
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
if ((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V0) > ((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V0) {
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
__t2 = (1) + ((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V0)
goto end_branch_2
} else {

}
}
{
__t2 = (1) + ((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V0)
}
end_branch_2:
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(&Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]{1, __t2, ((1) + ((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V1)) + ((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V1), gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(__local_var_0.IntVal)), UnsafePtr: nil}, gopurs_runtime.Float(__local_var_1.FloatVal()), gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[uint32, float64]](__local_var_2))}), gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[uint32, float64]](__local_var_3))})})}
goto end_branch_3
} else {

}
}
{
__t3 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_3:
__t4 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[uint32, float64]](__t3))}
goto end_branch_4
} else {

}
}
{
__t4 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_4:
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[uint32, float64]](__t4))}
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
var __t4 gopurs_runtime.Value
{
if (__local_var_2.Type == 9 && __local_var_2.IntVal == 324739070 && __local_var_2.UnsafePtr == nil) {
var __t0 gopurs_runtime.Value
{
if (__local_var_3.Type == 9 && __local_var_3.IntVal == 324739070 && __local_var_3.UnsafePtr == nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(&Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]{1, 1, 1, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](__local_var_0))}, __local_var_1, gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[*Constructor_Data_Maybe_Just[gopurs_runtime.Value], gopurs_runtime.Value]](__local_var_2))}), gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[*Constructor_Data_Maybe_Just[gopurs_runtime.Value], gopurs_runtime.Value]](__local_var_3))})})}
goto end_branch_0
} else {

}
}
{
if (__local_var_3.Type == 9 && __local_var_3.IntVal == 324739070 && __local_var_3.UnsafePtr != nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(&Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]{1, (1) + ((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V0), (1) + ((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V1), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](__local_var_0))}, __local_var_1, gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[*Constructor_Data_Maybe_Just[gopurs_runtime.Value], gopurs_runtime.Value]](__local_var_2))}), gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[*Constructor_Data_Maybe_Just[gopurs_runtime.Value], gopurs_runtime.Value]](__local_var_3))})})}
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
__t4 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[*Constructor_Data_Maybe_Just[gopurs_runtime.Value], gopurs_runtime.Value]](__t0))}
goto end_branch_4
} else {

}
}
{
if (__local_var_2.Type == 9 && __local_var_2.IntVal == 324739070 && __local_var_2.UnsafePtr != nil) {
var __t3 gopurs_runtime.Value
{
if (__local_var_3.Type == 9 && __local_var_3.IntVal == 324739070 && __local_var_3.UnsafePtr == nil) {
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(&Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]{1, (1) + ((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V0), (1) + ((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V1), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](__local_var_0))}, __local_var_1, gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[*Constructor_Data_Maybe_Just[gopurs_runtime.Value], gopurs_runtime.Value]](__local_var_2))}), gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[*Constructor_Data_Maybe_Just[gopurs_runtime.Value], gopurs_runtime.Value]](__local_var_3))})})}
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
if ((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V0) > ((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V0) {
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
__t2 = (1) + ((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V0)
goto end_branch_2
} else {

}
}
{
__t2 = (1) + ((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V0)
}
end_branch_2:
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(&Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]{1, __t2, ((1) + ((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V1)) + ((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V1), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](__local_var_0))}, __local_var_1, gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[*Constructor_Data_Maybe_Just[gopurs_runtime.Value], gopurs_runtime.Value]](__local_var_2))}), gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[*Constructor_Data_Maybe_Just[gopurs_runtime.Value], gopurs_runtime.Value]](__local_var_3))})})}
goto end_branch_3
} else {

}
}
{
__t3 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_3:
__t4 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[*Constructor_Data_Maybe_Just[gopurs_runtime.Value], gopurs_runtime.Value]](__t3))}
goto end_branch_4
} else {

}
}
{
__t4 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_4:
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[*Constructor_Data_Maybe_Just[gopurs_runtime.Value], gopurs_runtime.Value]](__t4))}
}

func Call_Data_Map_Internal_unsafeSplit__1094566431(__local_var_0_loop gopurs_runtime.Value, __local_var_1_loop gopurs_runtime.Value, __local_var_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __local_var_0 gopurs_runtime.Value = __local_var_0_loop
_ = __local_var_0
var __local_var_1 gopurs_runtime.Value = __local_var_1_loop
_ = __local_var_1
var __local_var_2 gopurs_runtime.Value = __local_var_2_loop
_ = __local_var_2
var __t4 gopurs_runtime.Value
{
if (__local_var_2.Type == 9 && __local_var_2.IntVal == 324739070 && __local_var_2.UnsafePtr == nil) {
__t4 = gopurs_runtime.Value{Type: 9, IntVal: 3373277644, UnsafePtr: unsafe.Pointer(&Constructor_Data_Map_Internal_Split[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil))}), gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(nil))}), gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(nil))})})}
goto end_branch_4
} else {

}
}
{
if (__local_var_2.Type == 9 && __local_var_2.IntVal == 324739070 && __local_var_2.UnsafePtr != nil) {
// TAST (Let): v_3_0 -> gopurs_runtime.Value
v_3_0 := gopurs_runtime.Apply2(__local_var_0, __local_var_1, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V2)
_ = v_3_0
var __t3 gopurs_runtime.Value
{
if (uint32(v_3_0.IntVal) == 1527465420) {
// TAST (Let): v1_4_1 -> gopurs_runtime.Value
v1_4_1 := gopurs_runtime.UncurriedApp3(Get_Data_Map_Internal_unsafeSplit(), __local_var_0, __local_var_1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V4)})
_ = v1_4_1
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 3373277644, UnsafePtr: unsafe.Pointer(&Constructor_Data_Map_Internal_Split[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Data_Map_Internal_Split[gopurs_runtime.Value, gopurs_runtime.Value])(v1_4_1.UnsafePtr).V0, (*Constructor_Data_Map_Internal_Split[gopurs_runtime.Value, gopurs_runtime.Value])(v1_4_1.UnsafePtr).V1, gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeBalancedNode(), (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Split[gopurs_runtime.Value, gopurs_runtime.Value])(v1_4_1.UnsafePtr).V2)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V5)}))})}
goto end_branch_3
} else {

}
}
{
if (uint32(v_3_0.IntVal) == 380165415) {
// TAST (Let): v1_4_2 -> gopurs_runtime.Value
v1_4_2 := gopurs_runtime.UncurriedApp3(Get_Data_Map_Internal_unsafeSplit(), __local_var_0, __local_var_1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V5)})
_ = v1_4_2
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 3373277644, UnsafePtr: unsafe.Pointer(&Constructor_Data_Map_Internal_Split[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Data_Map_Internal_Split[gopurs_runtime.Value, gopurs_runtime.Value])(v1_4_2.UnsafePtr).V0, gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeBalancedNode(), (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V4)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Split[gopurs_runtime.Value, gopurs_runtime.Value])(v1_4_2.UnsafePtr).V1)})), (*Constructor_Data_Map_Internal_Split[gopurs_runtime.Value, gopurs_runtime.Value])(v1_4_2.UnsafePtr).V2})}
goto end_branch_3
} else {

}
}
{
if (uint32(v_3_0.IntVal) == 902936544) {
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 3373277644, UnsafePtr: unsafe.Pointer(&Constructor_Data_Map_Internal_Split[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{1, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V3})}), (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V4, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V5})}
goto end_branch_3
} else {

}
}
{
__t3 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_3:
__t4 = gopurs_runtime.Value{Type: 9, IntVal: 3373277644, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Split[gopurs_runtime.Value, gopurs_runtime.Value]](__t3))}
goto end_branch_4
} else {

}
}
{
__t4 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_4:
return gopurs_runtime.Value{Type: 9, IntVal: 3373277644, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Split[gopurs_runtime.Value, gopurs_runtime.Value]](__t4))}
}

func Call_Data_Map_Internal_unsafeSplit__4154869695(__local_var_0_loop gopurs_runtime.Value, __local_var_1_loop gopurs_runtime.Value, __local_var_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __local_var_0 gopurs_runtime.Value = __local_var_0_loop
_ = __local_var_0
var __local_var_1 gopurs_runtime.Value = __local_var_1_loop
_ = __local_var_1
var __local_var_2 gopurs_runtime.Value = __local_var_2_loop
_ = __local_var_2
var __t4 gopurs_runtime.Value
{
if (__local_var_2.Type == 9 && __local_var_2.IntVal == 324739070 && __local_var_2.UnsafePtr == nil) {
__t4 = gopurs_runtime.Value{Type: 9, IntVal: 3373277644, UnsafePtr: unsafe.Pointer(&Constructor_Data_Map_Internal_Split[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil))}), gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(nil))}), gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(nil))})})}
goto end_branch_4
} else {

}
}
{
if (__local_var_2.Type == 9 && __local_var_2.IntVal == 324739070 && __local_var_2.UnsafePtr != nil) {
// TAST (Let): v_3_0 -> gopurs_runtime.Value
v_3_0 := gopurs_runtime.Apply2(__local_var_0, __local_var_1, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V2)
_ = v_3_0
var __t3 gopurs_runtime.Value
{
if (uint32(v_3_0.IntVal) == 1527465420) {
// TAST (Let): v1_4_1 -> gopurs_runtime.Value
v1_4_1 := gopurs_runtime.UncurriedApp3(Get_Data_Map_Internal_unsafeSplit(), __local_var_0, __local_var_1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V4)})
_ = v1_4_1
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 3373277644, UnsafePtr: unsafe.Pointer(&Constructor_Data_Map_Internal_Split[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Data_Map_Internal_Split[gopurs_runtime.Value, gopurs_runtime.Value])(v1_4_1.UnsafePtr).V0, (*Constructor_Data_Map_Internal_Split[gopurs_runtime.Value, gopurs_runtime.Value])(v1_4_1.UnsafePtr).V1, gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeBalancedNode(), (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Split[gopurs_runtime.Value, gopurs_runtime.Value])(v1_4_1.UnsafePtr).V2)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V5)}))})}
goto end_branch_3
} else {

}
}
{
if (uint32(v_3_0.IntVal) == 380165415) {
// TAST (Let): v1_4_2 -> gopurs_runtime.Value
v1_4_2 := gopurs_runtime.UncurriedApp3(Get_Data_Map_Internal_unsafeSplit(), __local_var_0, __local_var_1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V5)})
_ = v1_4_2
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 3373277644, UnsafePtr: unsafe.Pointer(&Constructor_Data_Map_Internal_Split[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Data_Map_Internal_Split[gopurs_runtime.Value, gopurs_runtime.Value])(v1_4_2.UnsafePtr).V0, gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeBalancedNode(), (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V4)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Split[gopurs_runtime.Value, gopurs_runtime.Value])(v1_4_2.UnsafePtr).V1)})), (*Constructor_Data_Map_Internal_Split[gopurs_runtime.Value, gopurs_runtime.Value])(v1_4_2.UnsafePtr).V2})}
goto end_branch_3
} else {

}
}
{
if (uint32(v_3_0.IntVal) == 902936544) {
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 3373277644, UnsafePtr: unsafe.Pointer(&Constructor_Data_Map_Internal_Split[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{1, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V3})}), (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V4, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V5})}
goto end_branch_3
} else {

}
}
{
__t3 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_3:
__t4 = gopurs_runtime.Value{Type: 9, IntVal: 3373277644, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Split[gopurs_runtime.Value, gopurs_runtime.Value]](__t3))}
goto end_branch_4
} else {

}
}
{
__t4 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_4:
return gopurs_runtime.Value{Type: 9, IntVal: 3373277644, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Split[gopurs_runtime.Value, gopurs_runtime.Value]](__t4))}
}

func Call_Data_Map_Internal_unsafeSplit__1308258847(__local_var_0_loop gopurs_runtime.Value, __local_var_1_loop gopurs_runtime.Value, __local_var_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __local_var_0 gopurs_runtime.Value = __local_var_0_loop
_ = __local_var_0
var __local_var_1 gopurs_runtime.Value = __local_var_1_loop
_ = __local_var_1
var __local_var_2 gopurs_runtime.Value = __local_var_2_loop
_ = __local_var_2
var __t4 gopurs_runtime.Value
{
if (__local_var_2.Type == 9 && __local_var_2.IntVal == 324739070 && __local_var_2.UnsafePtr == nil) {
__t4 = gopurs_runtime.Value{Type: 9, IntVal: 3373277644, UnsafePtr: unsafe.Pointer(&Constructor_Data_Map_Internal_Split[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil))}), gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(nil))}), gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(nil))})})}
goto end_branch_4
} else {

}
}
{
if (__local_var_2.Type == 9 && __local_var_2.IntVal == 324739070 && __local_var_2.UnsafePtr != nil) {
// TAST (Let): v_3_0 -> gopurs_runtime.Value
v_3_0 := gopurs_runtime.Apply2(__local_var_0, __local_var_1, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V2)
_ = v_3_0
var __t3 gopurs_runtime.Value
{
if (uint32(v_3_0.IntVal) == 1527465420) {
// TAST (Let): v1_4_1 -> gopurs_runtime.Value
v1_4_1 := gopurs_runtime.UncurriedApp3(Get_Data_Map_Internal_unsafeSplit(), __local_var_0, __local_var_1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V4)})
_ = v1_4_1
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 3373277644, UnsafePtr: unsafe.Pointer(&Constructor_Data_Map_Internal_Split[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Data_Map_Internal_Split[gopurs_runtime.Value, gopurs_runtime.Value])(v1_4_1.UnsafePtr).V0, (*Constructor_Data_Map_Internal_Split[gopurs_runtime.Value, gopurs_runtime.Value])(v1_4_1.UnsafePtr).V1, gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeBalancedNode(), (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Split[gopurs_runtime.Value, gopurs_runtime.Value])(v1_4_1.UnsafePtr).V2)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V5)}))})}
goto end_branch_3
} else {

}
}
{
if (uint32(v_3_0.IntVal) == 380165415) {
// TAST (Let): v1_4_2 -> gopurs_runtime.Value
v1_4_2 := gopurs_runtime.UncurriedApp3(Get_Data_Map_Internal_unsafeSplit(), __local_var_0, __local_var_1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V5)})
_ = v1_4_2
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 3373277644, UnsafePtr: unsafe.Pointer(&Constructor_Data_Map_Internal_Split[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Data_Map_Internal_Split[gopurs_runtime.Value, gopurs_runtime.Value])(v1_4_2.UnsafePtr).V0, gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeBalancedNode(), (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V4)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Split[gopurs_runtime.Value, gopurs_runtime.Value])(v1_4_2.UnsafePtr).V1)})), (*Constructor_Data_Map_Internal_Split[gopurs_runtime.Value, gopurs_runtime.Value])(v1_4_2.UnsafePtr).V2})}
goto end_branch_3
} else {

}
}
{
if (uint32(v_3_0.IntVal) == 902936544) {
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 3373277644, UnsafePtr: unsafe.Pointer(&Constructor_Data_Map_Internal_Split[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{1, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V3})}), (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V4, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V5})}
goto end_branch_3
} else {

}
}
{
__t3 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_3:
__t4 = gopurs_runtime.Value{Type: 9, IntVal: 3373277644, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Split[gopurs_runtime.Value, gopurs_runtime.Value]](__t3))}
goto end_branch_4
} else {

}
}
{
__t4 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_4:
return gopurs_runtime.Value{Type: 9, IntVal: 3373277644, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Split[gopurs_runtime.Value, gopurs_runtime.Value]](__t4))}
}

func Call_Data_Map_Internal_unsafeSplit__1115245464(__local_var_0_loop gopurs_runtime.Value, __local_var_1_loop gopurs_runtime.Value, __local_var_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __local_var_0 gopurs_runtime.Value = __local_var_0_loop
_ = __local_var_0
var __local_var_1 gopurs_runtime.Value = __local_var_1_loop
_ = __local_var_1
var __local_var_2 gopurs_runtime.Value = __local_var_2_loop
_ = __local_var_2
var __t4 gopurs_runtime.Value
{
if (__local_var_2.Type == 9 && __local_var_2.IntVal == 324739070 && __local_var_2.UnsafePtr == nil) {
__t4 = gopurs_runtime.Value{Type: 9, IntVal: 3373277644, UnsafePtr: unsafe.Pointer(&Constructor_Data_Map_Internal_Split[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil))}), gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(nil))}), gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(nil))})})}
goto end_branch_4
} else {

}
}
{
if (__local_var_2.Type == 9 && __local_var_2.IntVal == 324739070 && __local_var_2.UnsafePtr != nil) {
// TAST (Let): v_3_0 -> gopurs_runtime.Value
v_3_0 := gopurs_runtime.Apply2(__local_var_0, gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(__local_var_1.IntVal)), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(uint32((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V2.IntVal)), UnsafePtr: nil})
_ = v_3_0
var __t3 gopurs_runtime.Value
{
if (uint32(v_3_0.IntVal) == 1527465420) {
// TAST (Let): v1_4_1 -> gopurs_runtime.Value
v1_4_1 := gopurs_runtime.UncurriedApp3(Get_Data_Map_Internal_unsafeSplit(), __local_var_0, gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(__local_var_1.IntVal)), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V4)})
_ = v1_4_1
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 3373277644, UnsafePtr: unsafe.Pointer(&Constructor_Data_Map_Internal_Split[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Data_Map_Internal_Split[gopurs_runtime.Value, gopurs_runtime.Value])(v1_4_1.UnsafePtr).V0, (*Constructor_Data_Map_Internal_Split[gopurs_runtime.Value, gopurs_runtime.Value])(v1_4_1.UnsafePtr).V1, gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[uint32, float64]](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeBalancedNode(), gopurs_runtime.Value{Type: 9, IntVal: int64(uint32((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V2.IntVal)), UnsafePtr: nil}, gopurs_runtime.Float((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V3.FloatVal()), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Split[gopurs_runtime.Value, gopurs_runtime.Value])(v1_4_1.UnsafePtr).V2)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V5)})))})})}
goto end_branch_3
} else {

}
}
{
if (uint32(v_3_0.IntVal) == 380165415) {
// TAST (Let): v1_4_2 -> gopurs_runtime.Value
v1_4_2 := gopurs_runtime.UncurriedApp3(Get_Data_Map_Internal_unsafeSplit(), __local_var_0, gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(__local_var_1.IntVal)), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V5)})
_ = v1_4_2
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 3373277644, UnsafePtr: unsafe.Pointer(&Constructor_Data_Map_Internal_Split[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Data_Map_Internal_Split[gopurs_runtime.Value, gopurs_runtime.Value])(v1_4_2.UnsafePtr).V0, gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[uint32, float64]](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeBalancedNode(), gopurs_runtime.Value{Type: 9, IntVal: int64(uint32((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V2.IntVal)), UnsafePtr: nil}, gopurs_runtime.Float((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V3.FloatVal()), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V4)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Split[gopurs_runtime.Value, gopurs_runtime.Value])(v1_4_2.UnsafePtr).V1)})))}), (*Constructor_Data_Map_Internal_Split[gopurs_runtime.Value, gopurs_runtime.Value])(v1_4_2.UnsafePtr).V2})}
goto end_branch_3
} else {

}
}
{
if (uint32(v_3_0.IntVal) == 902936544) {
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 3373277644, UnsafePtr: unsafe.Pointer(&Constructor_Data_Map_Internal_Split[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{1, gopurs_runtime.Float((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V3.FloatVal())})}), (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V4, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V5})}
goto end_branch_3
} else {

}
}
{
__t3 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_3:
__t4 = gopurs_runtime.Value{Type: 9, IntVal: 3373277644, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Split[uint32, float64]](__t3))}
goto end_branch_4
} else {

}
}
{
__t4 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_4:
return gopurs_runtime.Value{Type: 9, IntVal: 3373277644, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Split[uint32, float64]](__t4))}
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
var __t1 gopurs_runtime.Value
{
if (__local_var_3.Type == 9 && __local_var_3.IntVal == 324739070 && __local_var_3.UnsafePtr == nil) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 2668112006, UnsafePtr: unsafe.Pointer(&Constructor_Data_Map_Internal_SplitLast[gopurs_runtime.Value, gopurs_runtime.Value]{1, __local_var_0, __local_var_1, gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__local_var_2)})}
goto end_branch_1
} else {

}
}
{
if (__local_var_3.Type == 9 && __local_var_3.IntVal == 324739070 && __local_var_3.UnsafePtr != nil) {
// TAST (Let): v1_4_0 -> gopurs_runtime.Value
v1_4_0 := gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeSplitLast(), (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V4)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V5)})
_ = v1_4_0
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 2668112006, UnsafePtr: unsafe.Pointer(&Constructor_Data_Map_Internal_SplitLast[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Data_Map_Internal_SplitLast[gopurs_runtime.Value, gopurs_runtime.Value])(v1_4_0.UnsafePtr).V0, (*Constructor_Data_Map_Internal_SplitLast[gopurs_runtime.Value, gopurs_runtime.Value])(v1_4_0.UnsafePtr).V1, gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeBalancedNode(), __local_var_0, __local_var_1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__local_var_2))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_SplitLast[gopurs_runtime.Value, gopurs_runtime.Value])(v1_4_0.UnsafePtr).V2)}))})}
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
return gopurs_runtime.Value{Type: 9, IntVal: 2668112006, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_SplitLast[gopurs_runtime.Value, gopurs_runtime.Value]](__t1))}
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
var __t1 gopurs_runtime.Value
{
if (__local_var_3.Type == 9 && __local_var_3.IntVal == 324739070 && __local_var_3.UnsafePtr == nil) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 2668112006, UnsafePtr: unsafe.Pointer(&Constructor_Data_Map_Internal_SplitLast[gopurs_runtime.Value, gopurs_runtime.Value]{1, __local_var_0, __local_var_1, gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__local_var_2)})}
goto end_branch_1
} else {

}
}
{
if (__local_var_3.Type == 9 && __local_var_3.IntVal == 324739070 && __local_var_3.UnsafePtr != nil) {
// TAST (Let): v1_4_0 -> gopurs_runtime.Value
v1_4_0 := gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeSplitLast(), (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V4)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V5)})
_ = v1_4_0
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 2668112006, UnsafePtr: unsafe.Pointer(&Constructor_Data_Map_Internal_SplitLast[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Data_Map_Internal_SplitLast[gopurs_runtime.Value, gopurs_runtime.Value])(v1_4_0.UnsafePtr).V0, (*Constructor_Data_Map_Internal_SplitLast[gopurs_runtime.Value, gopurs_runtime.Value])(v1_4_0.UnsafePtr).V1, gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeBalancedNode(), __local_var_0, __local_var_1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__local_var_2))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_SplitLast[gopurs_runtime.Value, gopurs_runtime.Value])(v1_4_0.UnsafePtr).V2)}))})}
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
return gopurs_runtime.Value{Type: 9, IntVal: 2668112006, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_SplitLast[gopurs_runtime.Value, gopurs_runtime.Value]](__t1))}
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
var __t6 gopurs_runtime.Value
{
if (__local_var_2.Type == 9 && __local_var_2.IntVal == 324739070 && __local_var_2.UnsafePtr == nil) {
__t6 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__local_var_3))}
goto end_branch_6
} else {

}
}
{
if (__local_var_3.Type == 9 && __local_var_3.IntVal == 324739070 && __local_var_3.UnsafePtr == nil) {
__t6 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__local_var_2))}
goto end_branch_6
} else {

}
}
{
if (__local_var_3.Type == 9 && __local_var_3.IntVal == 324739070 && __local_var_3.UnsafePtr != nil) {
// TAST (Let): v_4_0 -> gopurs_runtime.Value
v_4_0 := gopurs_runtime.UncurriedApp3(Get_Data_Map_Internal_unsafeSplit(), __local_var_0, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V2, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__local_var_2))})
_ = v_4_0
// TAST (Let): l_prime_5_1 -> gopurs_runtime.Value
l_prime_5_1 := gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeUnionWith(), __local_var_0, __local_var_1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Split[gopurs_runtime.Value, gopurs_runtime.Value])(v_4_0.UnsafePtr).V1)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V4)})
_ = l_prime_5_1
// TAST (Let): r_prime_6_2 -> gopurs_runtime.Value
r_prime_6_2 := gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeUnionWith(), __local_var_0, __local_var_1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Split[gopurs_runtime.Value, gopurs_runtime.Value])(v_4_0.UnsafePtr).V2)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V5)})
_ = r_prime_6_2
var __t5 gopurs_runtime.Value
{
var __t_tag_3 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Split[gopurs_runtime.Value, gopurs_runtime.Value])(v_4_0.UnsafePtr).V0)}
if (__t_tag_3.Type == 9 && __t_tag_3.IntVal == 930809136 && __t_tag_3.UnsafePtr != nil) {
__t5 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeBalancedNode(), (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V2, gopurs_runtime.Apply2(__local_var_1, (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Split[gopurs_runtime.Value, gopurs_runtime.Value])(v_4_0.UnsafePtr).V0)}.UnsafePtr).V0, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V3), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](l_prime_5_1))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](r_prime_6_2))})))}
goto end_branch_5
} else {

}
}
{
var __t_tag_4 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Split[gopurs_runtime.Value, gopurs_runtime.Value])(v_4_0.UnsafePtr).V0)}
if (__t_tag_4.Type == 9 && __t_tag_4.IntVal == 930809136 && __t_tag_4.UnsafePtr == nil) {
__t5 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeBalancedNode(), (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](l_prime_5_1))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](r_prime_6_2))})))}
goto end_branch_5
} else {

}
}
{
__t5 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_5:
__t6 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__t5))}
goto end_branch_6
} else {

}
}
{
__t6 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_6:
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__t6))}
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
var __t6 gopurs_runtime.Value
{
if (__local_var_2.Type == 9 && __local_var_2.IntVal == 324739070 && __local_var_2.UnsafePtr == nil) {
__t6 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[uint32, float64]](__local_var_3))}
goto end_branch_6
} else {

}
}
{
if (__local_var_3.Type == 9 && __local_var_3.IntVal == 324739070 && __local_var_3.UnsafePtr == nil) {
__t6 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[uint32, float64]](__local_var_2))}
goto end_branch_6
} else {

}
}
{
if (__local_var_3.Type == 9 && __local_var_3.IntVal == 324739070 && __local_var_3.UnsafePtr != nil) {
// TAST (Let): v_4_0 -> gopurs_runtime.Value
v_4_0 := gopurs_runtime.UncurriedApp3(Get_Data_Map_Internal_unsafeSplit(), __local_var_0, gopurs_runtime.Value{Type: 9, IntVal: int64(uint32((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V2.IntVal)), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[uint32, float64]](__local_var_2))})
_ = v_4_0
// TAST (Let): l_prime_5_1 -> gopurs_runtime.Value
l_prime_5_1 := gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeUnionWith(), __local_var_0, __local_var_1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Split[gopurs_runtime.Value, gopurs_runtime.Value])(v_4_0.UnsafePtr).V1)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V4)})
_ = l_prime_5_1
// TAST (Let): r_prime_6_2 -> gopurs_runtime.Value
r_prime_6_2 := gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeUnionWith(), __local_var_0, __local_var_1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Split[gopurs_runtime.Value, gopurs_runtime.Value])(v_4_0.UnsafePtr).V2)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V5)})
_ = r_prime_6_2
var __t5 gopurs_runtime.Value
{
var __t_tag_3 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Split[gopurs_runtime.Value, gopurs_runtime.Value])(v_4_0.UnsafePtr).V0)}
if (__t_tag_3.Type == 9 && __t_tag_3.IntVal == 930809136 && __t_tag_3.UnsafePtr != nil) {
__t5 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[uint32, float64]](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeBalancedNode(), gopurs_runtime.Value{Type: 9, IntVal: int64(uint32((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V2.IntVal)), UnsafePtr: nil}, gopurs_runtime.Float(gopurs_runtime.Apply2(__local_var_1, gopurs_runtime.Float((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Split[gopurs_runtime.Value, gopurs_runtime.Value])(v_4_0.UnsafePtr).V0)}.UnsafePtr).V0.FloatVal()), gopurs_runtime.Float((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V3.FloatVal())).FloatVal()), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[uint32, float64]](l_prime_5_1))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[uint32, float64]](r_prime_6_2))})))}
goto end_branch_5
} else {

}
}
{
var __t_tag_4 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Split[gopurs_runtime.Value, gopurs_runtime.Value])(v_4_0.UnsafePtr).V0)}
if (__t_tag_4.Type == 9 && __t_tag_4.IntVal == 930809136 && __t_tag_4.UnsafePtr == nil) {
__t5 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[uint32, float64]](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeBalancedNode(), gopurs_runtime.Value{Type: 9, IntVal: int64(uint32((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V2.IntVal)), UnsafePtr: nil}, gopurs_runtime.Float((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V3.FloatVal()), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[uint32, float64]](l_prime_5_1))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[uint32, float64]](r_prime_6_2))})))}
goto end_branch_5
} else {

}
}
{
__t5 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_5:
__t6 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[uint32, float64]](__t5))}
goto end_branch_6
} else {

}
}
{
__t6 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_6:
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[uint32, float64]](__t6))}
}


