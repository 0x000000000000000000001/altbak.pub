package Data_Map_Internal

import (
	pkg_Control_Applicative "gopurs/output/Control.Applicative"
	pkg_Control_Apply "gopurs/output/Control.Apply"
	pkg_Control_Semigroupoid "gopurs/output/Control.Semigroupoid"
	pkg_Data_Eq "gopurs/output/Data.Eq"
	pkg_Data_Foldable "gopurs/output/Data.Foldable"
	pkg_Data_FoldableWithIndex "gopurs/output/Data.FoldableWithIndex"
	pkg_Data_Function "gopurs/output/Data.Function"
	pkg_Data_Functor "gopurs/output/Data.Functor"
	pkg_Data_HeytingAlgebra "gopurs/output/Data.HeytingAlgebra"
	pkg_Data_List_Types "gopurs/output/Data.List.Types"
	pkg_Data_Maybe "gopurs/output/Data.Maybe"
	pkg_Data_Monoid "gopurs/output/Data.Monoid"
	pkg_Data_Ord "gopurs/output/Data.Ord"
	pkg_Data_Ring "gopurs/output/Data.Ring"
	pkg_Data_Semigroup "gopurs/output/Data.Semigroup"
	pkg_Data_Semiring "gopurs/output/Data.Semiring"
	pkg_Data_Show "gopurs/output/Data.Show"
	pkg_Data_Traversable "gopurs/output/Data.Traversable"
	pkg_Data_Tuple "gopurs/output/Data.Tuple"
	pkg_Data_Unfoldable "gopurs/output/Data.Unfoldable"
	pkg_Data_Unfoldable1 "gopurs/output/Data.Unfoldable1"
	pkg_Data_Unit "gopurs/output/Data.Unit"
	pkg_Partial_Unsafe "gopurs/output/Partial.Unsafe"
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_identity gopurs_runtime.Value
var once_identity sync.Once
func Get_identity() gopurs_runtime.Value {
	once_identity.Do(func() {
		cache_identity = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_identity(x_0_box)
})
	})
	return cache_identity
}

var cache_identity1 gopurs_runtime.Value
var once_identity1 sync.Once
func Get_identity1() gopurs_runtime.Value {
	once_identity1.Do(func() {
		cache_identity1 = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_identity1(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](x_0_box)))}
})
	})
	return cache_identity1
}

var cache_identity2 gopurs_runtime.Value
var once_identity2 sync.Once
func Get_identity2() gopurs_runtime.Value {
	once_identity2.Do(func() {
		cache_identity2 = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_identity2(x_0_box)
})
	})
	return cache_identity2
}

var cache_Leaf gopurs_runtime.Value
var once_Leaf sync.Once
func Get_Leaf() gopurs_runtime.Value {
	once_Leaf.Do(func() {
		cache_Leaf = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(nil))}
	})
	return cache_Leaf
}

var cache_Node gopurs_runtime.Value
var once_Node sync.Once
func Get_Node() gopurs_runtime.Value {
	once_Node.Do(func() {
		cache_Node = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(value1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(value2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(value3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(value4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(value5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(&Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]{1, value0.IntVal, value1.IntVal, value2, value3, gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](value4), gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](value5)})}
})
})
})
})
})
})
	})
	return cache_Node
}

var cache_IterLeaf gopurs_runtime.Value
var once_IterLeaf sync.Once
func Get_IterLeaf() gopurs_runtime.Value {
	once_IterLeaf.Do(func() {
		cache_IterLeaf = gopurs_runtime.Value{Type: 9, IntVal: 2509360378, UnsafePtr: unsafe.Pointer(nil)}
	})
	return cache_IterLeaf
}

var cache_IterEmit gopurs_runtime.Value
var once_IterEmit sync.Once
func Get_IterEmit() gopurs_runtime.Value {
	once_IterEmit.Do(func() {
		cache_IterEmit = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(value1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(value2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1343415489, UnsafePtr: unsafe.Pointer(&Constructor_IterEmit[gopurs_runtime.Value, gopurs_runtime.Value]{1, value0, value1, value2})}
})
})
})
	})
	return cache_IterEmit
}

var cache_IterNode gopurs_runtime.Value
var once_IterNode sync.Once
func Get_IterNode() gopurs_runtime.Value {
	once_IterNode.Do(func() {
		cache_IterNode = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(value1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2861335956, UnsafePtr: unsafe.Pointer(&Constructor_IterNode[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](value0), value1})}
})
})
	})
	return cache_IterNode
}

var cache_IterDone gopurs_runtime.Value
var once_IterDone sync.Once
func Get_IterDone() gopurs_runtime.Value {
	once_IterDone.Do(func() {
		cache_IterDone = gopurs_runtime.Value{Type: 9, IntVal: 953589075, UnsafePtr: unsafe.Pointer((*Constructor_IterNext[gopurs_runtime.Value, gopurs_runtime.Value])(nil))}
	})
	return cache_IterDone
}

var cache_IterNext gopurs_runtime.Value
var once_IterNext sync.Once
func Get_IterNext() gopurs_runtime.Value {
	once_IterNext.Do(func() {
		cache_IterNext = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(value1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(value2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 953589075, UnsafePtr: unsafe.Pointer(&Constructor_IterNext[gopurs_runtime.Value, gopurs_runtime.Value]{1, value0, value1, value2})}
})
})
})
	})
	return cache_IterNext
}

var cache_Split gopurs_runtime.Value
var once_Split sync.Once
func Get_Split() gopurs_runtime.Value {
	once_Split.Do(func() {
		cache_Split = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(value1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(value2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3373277644, UnsafePtr: unsafe.Pointer(&Constructor_Split[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](value0), gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](value1), gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](value2)})}
})
})
})
	})
	return cache_Split
}

var cache_SplitLast gopurs_runtime.Value
var once_SplitLast sync.Once
func Get_SplitLast() gopurs_runtime.Value {
	once_SplitLast.Do(func() {
		cache_SplitLast = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(value1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(value2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2668112006, UnsafePtr: unsafe.Pointer(&Constructor_SplitLast[gopurs_runtime.Value, gopurs_runtime.Value]{1, value0, value1, gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](value2)})}
})
})
})
	})
	return cache_SplitLast
}

var cache_unsafeNode gopurs_runtime.Value
var once_unsafeNode sync.Once
func Get_unsafeNode() gopurs_runtime.Value {
	once_unsafeNode.Do(func() {
		cache_unsafeNode = gopurs_runtime.Func4(func(__local_var_0_box gopurs_runtime.Value, __local_var_1_box gopurs_runtime.Value, __local_var_2_box gopurs_runtime.Value, __local_var_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_unsafeNode(__local_var_0_box, __local_var_1_box, __local_var_2_box, __local_var_3_box)
})
	})
	return cache_unsafeNode
}

var cache_toMapIter gopurs_runtime.Value
var once_toMapIter sync.Once
func Get_toMapIter() gopurs_runtime.Value {
	once_toMapIter.Do(func() {
		cache_toMapIter = gopurs_runtime.Func(func(a_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_toMapIter(gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](a_0_box))
})
	})
	return cache_toMapIter
}

var cache_stepWith gopurs_runtime.Value
var once_stepWith sync.Once
func Get_stepWith() gopurs_runtime.Value {
	once_stepWith.Do(func() {
		cache_stepWith = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, next_1_box gopurs_runtime.Value, done_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_stepWith(f_0_box, next_1_box, done_2_box)
})
	})
	return cache_stepWith
}

var cache_size gopurs_runtime.Value
var once_size sync.Once
func Get_size() gopurs_runtime.Value {
	once_size.Do(func() {
		cache_size = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int(Call_size(gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](v_0_box)))
})
	})
	return cache_size
}

var cache_singleton gopurs_runtime.Value
var once_singleton sync.Once
func Get_singleton() gopurs_runtime.Value {
	once_singleton.Do(func() {
		cache_singleton = gopurs_runtime.Func2(func(k_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(Call_singleton(k_0_box, v_1_box))}
})
	})
	return cache_singleton
}

var cache_unsafeBalancedNode gopurs_runtime.Value
var once_unsafeBalancedNode sync.Once
func Get_unsafeBalancedNode() gopurs_runtime.Value {
	once_unsafeBalancedNode.Do(func() {
		cache_unsafeBalancedNode = gopurs_runtime.Func4(func(__local_var_0_box gopurs_runtime.Value, __local_var_1_box gopurs_runtime.Value, __local_var_2_box gopurs_runtime.Value, __local_var_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_unsafeBalancedNode(__local_var_0_box, __local_var_1_box, __local_var_2_box, __local_var_3_box)
})
	})
	return cache_unsafeBalancedNode
}

var cache_unsafeSplit gopurs_runtime.Value
var once_unsafeSplit sync.Once
func Get_unsafeSplit() gopurs_runtime.Value {
	once_unsafeSplit.Do(func() {
		cache_unsafeSplit = gopurs_runtime.Func3(func(__local_var_0_box gopurs_runtime.Value, __local_var_1_box gopurs_runtime.Value, __local_var_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_unsafeSplit(__local_var_0_box, __local_var_1_box, __local_var_2_box)
})
	})
	return cache_unsafeSplit
}

var cache_unsafeSplitLast gopurs_runtime.Value
var once_unsafeSplitLast sync.Once
func Get_unsafeSplitLast() gopurs_runtime.Value {
	once_unsafeSplitLast.Do(func() {
		cache_unsafeSplitLast = gopurs_runtime.Func4(func(__local_var_0_box gopurs_runtime.Value, __local_var_1_box gopurs_runtime.Value, __local_var_2_box gopurs_runtime.Value, __local_var_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_unsafeSplitLast(__local_var_0_box, __local_var_1_box, __local_var_2_box, __local_var_3_box)
})
	})
	return cache_unsafeSplitLast
}

var cache_unsafeJoinNodes gopurs_runtime.Value
var once_unsafeJoinNodes sync.Once
func Get_unsafeJoinNodes() gopurs_runtime.Value {
	once_unsafeJoinNodes.Do(func() {
		cache_unsafeJoinNodes = gopurs_runtime.Func2(func(__local_var_0_box gopurs_runtime.Value, __local_var_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_unsafeJoinNodes(__local_var_0_box, __local_var_1_box)
})
	})
	return cache_unsafeJoinNodes
}

var cache_unsafeDifference gopurs_runtime.Value
var once_unsafeDifference sync.Once
func Get_unsafeDifference() gopurs_runtime.Value {
	once_unsafeDifference.Do(func() {
		cache_unsafeDifference = gopurs_runtime.Func3(func(__local_var_0_box gopurs_runtime.Value, __local_var_1_box gopurs_runtime.Value, __local_var_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_unsafeDifference(__local_var_0_box, __local_var_1_box, __local_var_2_box)
})
	})
	return cache_unsafeDifference
}

var cache_unsafeIntersectionWith gopurs_runtime.Value
var once_unsafeIntersectionWith sync.Once
func Get_unsafeIntersectionWith() gopurs_runtime.Value {
	once_unsafeIntersectionWith.Do(func() {
		cache_unsafeIntersectionWith = gopurs_runtime.Func4(func(__local_var_0_box gopurs_runtime.Value, __local_var_1_box gopurs_runtime.Value, __local_var_2_box gopurs_runtime.Value, __local_var_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_unsafeIntersectionWith(__local_var_0_box, __local_var_1_box, __local_var_2_box, __local_var_3_box)
})
	})
	return cache_unsafeIntersectionWith
}

var cache_unsafeUnionWith gopurs_runtime.Value
var once_unsafeUnionWith sync.Once
func Get_unsafeUnionWith() gopurs_runtime.Value {
	once_unsafeUnionWith.Do(func() {
		cache_unsafeUnionWith = gopurs_runtime.Func4(func(__local_var_0_box gopurs_runtime.Value, __local_var_1_box gopurs_runtime.Value, __local_var_2_box gopurs_runtime.Value, __local_var_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_unsafeUnionWith(__local_var_0_box, __local_var_1_box, __local_var_2_box, __local_var_3_box)
})
	})
	return cache_unsafeUnionWith
}

var cache_unionWith gopurs_runtime.Value
var once_unionWith sync.Once
func Get_unionWith() gopurs_runtime.Value {
	once_unionWith.Do(func() {
		cache_unionWith = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_unionWith(gopurs_runtime.CoerceToStruct[pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value]](dictOrd_0_box))
})
	})
	return cache_unionWith
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

var cache_update gopurs_runtime.Value
var once_update sync.Once
func Get_update() gopurs_runtime.Value {
	once_update.Do(func() {
		cache_update = gopurs_runtime.Func3(func(dictOrd_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, k_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_update(gopurs_runtime.CoerceToStruct[pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value]](dictOrd_0_box), f_1_box, k_2_box)
})
	})
	return cache_update
}

var cache_showTree gopurs_runtime.Value
var once_showTree sync.Once
func Get_showTree() gopurs_runtime.Value {
	once_showTree.Do(func() {
		cache_showTree = gopurs_runtime.Func2(func(dictShow_0_box gopurs_runtime.Value, dictShow1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_showTree(gopurs_runtime.CoerceToStruct[pkg_Data_Show.Constructor_Show[gopurs_runtime.Value]](dictShow_0_box), gopurs_runtime.CoerceToStruct[pkg_Data_Show.Constructor_Show[gopurs_runtime.Value]](dictShow1_1_box))
})
	})
	return cache_showTree
}

var cache_semigroupMap gopurs_runtime.Value
var once_semigroupMap sync.Once
func Get_semigroupMap() gopurs_runtime.Value {
	once_semigroupMap.Do(func() {
		cache_semigroupMap = gopurs_runtime.Func3(func(_dollar__unused_0_box gopurs_runtime.Value, dictOrd_1_box gopurs_runtime.Value, dictSemigroup_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_semigroupMap(_dollar__unused_0_box, dictOrd_1_box, dictSemigroup_2_box)
})
	})
	return cache_semigroupMap
}

var cache_pop gopurs_runtime.Value
var once_pop sync.Once
func Get_pop() gopurs_runtime.Value {
	once_pop.Do(func() {
		cache_pop = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_pop(gopurs_runtime.CoerceToStruct[pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value]](dictOrd_0_box))
})
	})
	return cache_pop
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

var cache_mapMaybeWithKey gopurs_runtime.Value
var once_mapMaybeWithKey sync.Once
func Get_mapMaybeWithKey() gopurs_runtime.Value {
	once_mapMaybeWithKey.Do(func() {
		cache_mapMaybeWithKey = gopurs_runtime.Func2(func(dictOrd_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_mapMaybeWithKey(gopurs_runtime.CoerceToStruct[pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value]](dictOrd_0_box), f_1_box)
})
	})
	return cache_mapMaybeWithKey
}

var cache_mapMaybe gopurs_runtime.Value
var once_mapMaybe sync.Once
func Get_mapMaybe() gopurs_runtime.Value {
	once_mapMaybe.Do(func() {
		cache_mapMaybe = gopurs_runtime.Func2(func(dictOrd_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_mapMaybe(gopurs_runtime.CoerceToStruct[pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value]](dictOrd_0_box), x_1_box)
})
	})
	return cache_mapMaybe
}

var cache_lookupLE gopurs_runtime.Value
var once_lookupLE sync.Once
func Get_lookupLE() gopurs_runtime.Value {
	once_lookupLE.Do(func() {
		cache_lookupLE = gopurs_runtime.Func2(func(dictOrd_0_box gopurs_runtime.Value, k_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_lookupLE(gopurs_runtime.CoerceToStruct[pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value]](dictOrd_0_box), k_1_box)
})
	})
	return cache_lookupLE
}

var cache_lookupGE gopurs_runtime.Value
var once_lookupGE sync.Once
func Get_lookupGE() gopurs_runtime.Value {
	once_lookupGE.Do(func() {
		cache_lookupGE = gopurs_runtime.Func2(func(dictOrd_0_box gopurs_runtime.Value, k_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_lookupGE(gopurs_runtime.CoerceToStruct[pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value]](dictOrd_0_box), k_1_box)
})
	})
	return cache_lookupGE
}

var cache_lookup gopurs_runtime.Value
var once_lookup sync.Once
func Get_lookup() gopurs_runtime.Value {
	once_lookup.Do(func() {
		cache_lookup = gopurs_runtime.Func2(func(dictOrd_0_box gopurs_runtime.Value, k_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_lookup(gopurs_runtime.CoerceToStruct[pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value]](dictOrd_0_box), k_1_box)
})
	})
	return cache_lookup
}

var cache_iterMapU gopurs_runtime.Value
var once_iterMapU sync.Once
func Get_iterMapU() gopurs_runtime.Value {
	once_iterMapU.Do(func() {
		cache_iterMapU = gopurs_runtime.Func2(func(iter_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_iterMapU(iter_0_box, gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](v_1_box))
})
	})
	return cache_iterMapU
}

var cache_stepUnorderedCps gopurs_runtime.Value
var once_stepUnorderedCps sync.Once
func Get_stepUnorderedCps() gopurs_runtime.Value {
	once_stepUnorderedCps.Do(func() {
		cache_stepUnorderedCps = gopurs_runtime.Apply(Get_stepWith(), Get_iterMapU())
	})
	return cache_stepUnorderedCps
}

var cache_stepUnfoldrUnordered gopurs_runtime.Value
var once_stepUnfoldrUnordered sync.Once
func Get_stepUnfoldrUnordered() gopurs_runtime.Value {
	once_stepUnfoldrUnordered.Do(func() {
		cache_stepUnfoldrUnordered = Call_stepWith(Get_iterMapU(), gopurs_runtime.Func3(func(__local_var_0 gopurs_runtime.Value, __local_var_1 gopurs_runtime.Value, __local_var_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, __local_var_0, __local_var_1})}, __local_var_2})}})}
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(nil))}
}))
	})
	return cache_stepUnfoldrUnordered
}

var cache_toUnfoldableUnordered gopurs_runtime.Value
var once_toUnfoldableUnordered sync.Once
func Get_toUnfoldableUnordered() gopurs_runtime.Value {
	once_toUnfoldableUnordered.Do(func() {
		cache_toUnfoldableUnordered = gopurs_runtime.Func(func(dictUnfoldable_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_toUnfoldableUnordered(gopurs_runtime.CoerceToStruct[pkg_Data_Unfoldable.Constructor_Unfoldable[gopurs_runtime.Value]](dictUnfoldable_0_box))
})
	})
	return cache_toUnfoldableUnordered
}

var cache_stepUnordered gopurs_runtime.Value
var once_stepUnordered sync.Once
func Get_stepUnordered() gopurs_runtime.Value {
	once_stepUnordered.Do(func() {
		cache_stepUnordered = Call_stepWith(Get_iterMapU(), gopurs_runtime.Func3(func(__local_var_0 gopurs_runtime.Value, __local_var_1 gopurs_runtime.Value, __local_var_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 953589075, UnsafePtr: unsafe.Pointer(&Constructor_IterNext[gopurs_runtime.Value, gopurs_runtime.Value]{1, __local_var_0, __local_var_1, __local_var_2})}
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 953589075, UnsafePtr: unsafe.Pointer((*Constructor_IterNext[gopurs_runtime.Value, gopurs_runtime.Value])(nil))}
}))
	})
	return cache_stepUnordered
}

var cache_iterMapR gopurs_runtime.Value
var once_iterMapR sync.Once
func Get_iterMapR() gopurs_runtime.Value {
	once_iterMapR.Do(func() {
		cache_iterMapR = func() gopurs_runtime.Value {
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
var __t_tag_1 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V5)}
if (__t_tag_1.Type == 9 && __t_tag_1.IntVal == 324739070 && __t_tag_1.UnsafePtr == nil) {
iter_1_loop = gopurs_runtime.Value{Type: 9, IntVal: 1343415489, UnsafePtr: unsafe.Pointer(&Constructor_IterEmit[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V2, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V3, iter_1})}
v_2_loop = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V4)}
continue go__go_0_0_8
__t2 = gopurs_runtime.Value{}
goto end_branch_2
} else {

}
}
{
iter_1_loop = gopurs_runtime.Value{Type: 9, IntVal: 1343415489, UnsafePtr: unsafe.Pointer(&Constructor_IterEmit[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V2, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 2861335956, UnsafePtr: unsafe.Pointer(&Constructor_IterNode[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V4, iter_1})}})}
v_2_loop = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V5)}
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
	return cache_iterMapR
}

var cache_stepDescCps gopurs_runtime.Value
var once_stepDescCps sync.Once
func Get_stepDescCps() gopurs_runtime.Value {
	once_stepDescCps.Do(func() {
		cache_stepDescCps = gopurs_runtime.Apply(Get_stepWith(), Get_iterMapR())
	})
	return cache_stepDescCps
}

var cache_stepDesc gopurs_runtime.Value
var once_stepDesc sync.Once
func Get_stepDesc() gopurs_runtime.Value {
	once_stepDesc.Do(func() {
		cache_stepDesc = Call_stepWith(Get_iterMapR(), gopurs_runtime.Func3(func(__local_var_0 gopurs_runtime.Value, __local_var_1 gopurs_runtime.Value, __local_var_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 953589075, UnsafePtr: unsafe.Pointer(&Constructor_IterNext[gopurs_runtime.Value, gopurs_runtime.Value]{1, __local_var_0, __local_var_1, __local_var_2})}
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 953589075, UnsafePtr: unsafe.Pointer((*Constructor_IterNext[gopurs_runtime.Value, gopurs_runtime.Value])(nil))}
}))
	})
	return cache_stepDesc
}

var cache_iterMapL gopurs_runtime.Value
var once_iterMapL sync.Once
func Get_iterMapL() gopurs_runtime.Value {
	once_iterMapL.Do(func() {
		cache_iterMapL = func() gopurs_runtime.Value {
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
var __t_tag_1 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V5)}
if (__t_tag_1.Type == 9 && __t_tag_1.IntVal == 324739070 && __t_tag_1.UnsafePtr == nil) {
iter_1_loop = gopurs_runtime.Value{Type: 9, IntVal: 1343415489, UnsafePtr: unsafe.Pointer(&Constructor_IterEmit[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V2, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V3, iter_1})}
v_2_loop = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V4)}
continue go__go_0_0_9
__t2 = gopurs_runtime.Value{}
goto end_branch_2
} else {

}
}
{
iter_1_loop = gopurs_runtime.Value{Type: 9, IntVal: 1343415489, UnsafePtr: unsafe.Pointer(&Constructor_IterEmit[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V2, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 2861335956, UnsafePtr: unsafe.Pointer(&Constructor_IterNode[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V5, iter_1})}})}
v_2_loop = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V4)}
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
	return cache_iterMapL
}

var cache_stepAscCps gopurs_runtime.Value
var once_stepAscCps sync.Once
func Get_stepAscCps() gopurs_runtime.Value {
	once_stepAscCps.Do(func() {
		cache_stepAscCps = gopurs_runtime.Apply(Get_stepWith(), Get_iterMapL())
	})
	return cache_stepAscCps
}

var cache_stepAsc gopurs_runtime.Value
var once_stepAsc sync.Once
func Get_stepAsc() gopurs_runtime.Value {
	once_stepAsc.Do(func() {
		cache_stepAsc = Call_stepWith(Get_iterMapL(), gopurs_runtime.Func3(func(__local_var_0 gopurs_runtime.Value, __local_var_1 gopurs_runtime.Value, __local_var_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 953589075, UnsafePtr: unsafe.Pointer(&Constructor_IterNext[gopurs_runtime.Value, gopurs_runtime.Value]{1, __local_var_0, __local_var_1, __local_var_2})}
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 953589075, UnsafePtr: unsafe.Pointer((*Constructor_IterNext[gopurs_runtime.Value, gopurs_runtime.Value])(nil))}
}))
	})
	return cache_stepAsc
}

var cache_eqMapIter gopurs_runtime.Value
var once_eqMapIter sync.Once
func Get_eqMapIter() gopurs_runtime.Value {
	once_eqMapIter.Do(func() {
		cache_eqMapIter = gopurs_runtime.Func2(func(dictEq_0_box gopurs_runtime.Value, dictEq1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_eqMapIter(dictEq_0_box, dictEq1_1_box)
})
	})
	return cache_eqMapIter
}

var cache_ordMapIter gopurs_runtime.Value
var once_ordMapIter sync.Once
func Get_ordMapIter() gopurs_runtime.Value {
	once_ordMapIter.Do(func() {
		cache_ordMapIter = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_ordMapIter(dictOrd_0_box)
})
	})
	return cache_ordMapIter
}

var cache_stepUnfoldr gopurs_runtime.Value
var once_stepUnfoldr sync.Once
func Get_stepUnfoldr() gopurs_runtime.Value {
	once_stepUnfoldr.Do(func() {
		cache_stepUnfoldr = Call_stepWith(Get_iterMapL(), gopurs_runtime.Func3(func(__local_var_0 gopurs_runtime.Value, __local_var_1 gopurs_runtime.Value, __local_var_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, __local_var_0, __local_var_1})}, __local_var_2})}})}
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(nil))}
}))
	})
	return cache_stepUnfoldr
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

var cache_showMap gopurs_runtime.Value
var once_showMap sync.Once
func Get_showMap() gopurs_runtime.Value {
	once_showMap.Do(func() {
		cache_showMap = gopurs_runtime.Func2(func(dictShow_0_box gopurs_runtime.Value, dictShow1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_showMap(dictShow_0_box, dictShow1_1_box)
})
	})
	return cache_showMap
}

var cache_isSubmap gopurs_runtime.Value
var once_isSubmap sync.Once
func Get_isSubmap() gopurs_runtime.Value {
	once_isSubmap.Do(func() {
		cache_isSubmap = gopurs_runtime.Func2(func(dictOrd_0_box gopurs_runtime.Value, dictEq_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_isSubmap(gopurs_runtime.CoerceToStruct[pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value]](dictOrd_0_box), gopurs_runtime.CoerceToStruct[pkg_Data_Eq.Constructor_Eq[gopurs_runtime.Value]](dictEq_1_box))
})
	})
	return cache_isSubmap
}

var cache_isEmpty gopurs_runtime.Value
var once_isEmpty sync.Once
func Get_isEmpty() gopurs_runtime.Value {
	once_isEmpty.Do(func() {
		cache_isEmpty = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_isEmpty(gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](v_0_box)))
})
	})
	return cache_isEmpty
}

var cache_intersectionWith gopurs_runtime.Value
var once_intersectionWith sync.Once
func Get_intersectionWith() gopurs_runtime.Value {
	once_intersectionWith.Do(func() {
		cache_intersectionWith = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_intersectionWith(gopurs_runtime.CoerceToStruct[pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value]](dictOrd_0_box))
})
	})
	return cache_intersectionWith
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

var cache_insertWith gopurs_runtime.Value
var once_insertWith sync.Once
func Get_insertWith() gopurs_runtime.Value {
	once_insertWith.Do(func() {
		cache_insertWith = gopurs_runtime.Func4(func(dictOrd_0_box gopurs_runtime.Value, app_1_box gopurs_runtime.Value, k_2_box gopurs_runtime.Value, v_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_insertWith(gopurs_runtime.CoerceToStruct[pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value]](dictOrd_0_box), app_1_box, k_2_box, v_3_box)
})
	})
	return cache_insertWith
}

var cache_insert gopurs_runtime.Value
var once_insert sync.Once
func Get_insert() gopurs_runtime.Value {
	once_insert.Do(func() {
		cache_insert = gopurs_runtime.Func3(func(dictOrd_0_box gopurs_runtime.Value, k_1_box gopurs_runtime.Value, v_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_insert(gopurs_runtime.CoerceToStruct[pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value]](dictOrd_0_box), k_1_box, v_2_box)
})
	})
	return cache_insert
}

var cache_functorMap gopurs_runtime.Value
var once_functorMap sync.Once
func Get_functorMap() gopurs_runtime.Value {
	once_functorMap.Do(func() {
		cache_functorMap = gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_1_0_16 gopurs_runtime.Value
_ = go__go_1_0_16
go__go_1_0_16 = gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if (v_2.Type == 9 && v_2.IntVal == 324739070 && v_2.UnsafePtr == nil) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(nil))}
goto end_branch_1
} else {

}
}
{
if (v_2.Type == 9 && v_2.IntVal == 324739070 && v_2.UnsafePtr != nil) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(&Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V0, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V1, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V2, gopurs_runtime.Apply(f_0, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V3), gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(go__go_1_0_16, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V4)})), gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(go__go_1_0_16, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V5)}))})}
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__t1))}
})
return go__go_1_0_16
}))
	})
	return cache_functorMap
}

var cache_functorWithIndexMap gopurs_runtime.Value
var once_functorWithIndexMap sync.Once
func Get_functorWithIndexMap() gopurs_runtime.Value {
	once_functorWithIndexMap.Do(func() {
		cache_functorWithIndexMap = gopurs_runtime.RecordDict2("Functor0", "mapWithIndex", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_functorMap()
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_1_0_17 gopurs_runtime.Value
_ = go__go_1_0_17
go__go_1_0_17 = gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if (v_2.Type == 9 && v_2.IntVal == 324739070 && v_2.UnsafePtr == nil) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(nil))}
goto end_branch_1
} else {

}
}
{
if (v_2.Type == 9 && v_2.IntVal == 324739070 && v_2.UnsafePtr != nil) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(&Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V0, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V1, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V2, gopurs_runtime.Apply2(f_0, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V2, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V3), gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(go__go_1_0_17, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V4)})), gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(go__go_1_0_17, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V5)}))})}
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__t1))}
})
return go__go_1_0_17
}))
	})
	return cache_functorWithIndexMap
}

var cache_foldableMap gopurs_runtime.Value
var once_foldableMap sync.Once
func Get_foldableMap() gopurs_runtime.Value {
	once_foldableMap.Do(func() {
		cache_foldableMap = gopurs_runtime.RecordDict3("foldMap", "foldl", "foldr", gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
Semigroup0_1_0 := gopurs_runtime.CoerceToStruct[pkg_Data_Semigroup.Constructor_Semigroup[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{}))
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
__t2 = gopurs_runtime.Apply2(Semigroup0_1_0.V0, gopurs_runtime.Apply(go__go_3_1_18, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V4)}), gopurs_runtime.Apply2(Semigroup0_1_0.V0, gopurs_runtime.Apply(f_2, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V3), gopurs_runtime.Apply(go__go_3_1_18, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V5)})))
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
__t4 = gopurs_runtime.UncurriedApp2(go__go_2_3_19, gopurs_runtime.Apply2(f_0, gopurs_runtime.UncurriedApp2(go__go_2_3_19, __local_var_3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_4.UnsafePtr).V4)}), (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_4.UnsafePtr).V3), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_4.UnsafePtr).V5)})
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
return gopurs_runtime.UncurriedApp2(go__go_2_3_19, z_1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](m_3))})
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
__t6 = gopurs_runtime.UncurriedApp2(go__go_2_5_20, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V4)}, gopurs_runtime.Apply2(f_0, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V3, gopurs_runtime.UncurriedApp2(go__go_2_5_20, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V5)}, __local_var_4)))
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
return gopurs_runtime.UncurriedApp2(go__go_2_5_20, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](m_3))}, z_1)
})
})
}))
	})
	return cache_foldableMap
}

var cache_foldableWithIndexMap gopurs_runtime.Value
var once_foldableWithIndexMap sync.Once
func Get_foldableWithIndexMap() gopurs_runtime.Value {
	once_foldableWithIndexMap.Do(func() {
		cache_foldableWithIndexMap = gopurs_runtime.RecordDict4("Foldable0", "foldMapWithIndex", "foldlWithIndex", "foldrWithIndex", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_foldableMap()
}), gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
Semigroup0_1_0 := gopurs_runtime.CoerceToStruct[pkg_Data_Semigroup.Constructor_Semigroup[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{}))
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
__t2 = gopurs_runtime.Apply2(Semigroup0_1_0.V0, gopurs_runtime.Apply(go__go_3_1_21, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V4)}), gopurs_runtime.Apply2(Semigroup0_1_0.V0, gopurs_runtime.Apply2(f_2, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V2, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V3), gopurs_runtime.Apply(go__go_3_1_21, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V5)})))
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
__t4 = gopurs_runtime.UncurriedApp2(go__go_2_3_22, gopurs_runtime.Apply3(f_0, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_4.UnsafePtr).V2, gopurs_runtime.UncurriedApp2(go__go_2_3_22, __local_var_3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_4.UnsafePtr).V4)}), (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_4.UnsafePtr).V3), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_4.UnsafePtr).V5)})
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
return gopurs_runtime.UncurriedApp2(go__go_2_3_22, z_1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](m_3))})
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
__t6 = gopurs_runtime.UncurriedApp2(go__go_2_5_23, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V4)}, gopurs_runtime.Apply3(f_0, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V2, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V3, gopurs_runtime.UncurriedApp2(go__go_2_5_23, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V5)}, __local_var_4)))
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
return gopurs_runtime.UncurriedApp2(go__go_2_5_23, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](m_3))}, z_1)
})
})
}))
	})
	return cache_foldableWithIndexMap
}

var cache_keys gopurs_runtime.Value
var once_keys sync.Once
func Get_keys() gopurs_runtime.Value {
	once_keys.Do(func() {
		cache_keys = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_foldableWithIndexMap(), "foldrWithIndex"), gopurs_runtime.Func(func(k_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(acc_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]{1, k_0, gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](acc_2)})}
})
})
}), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(nil))})
	})
	return cache_keys
}

var cache_traversableMap gopurs_runtime.Value
var once_traversableMap sync.Once
func Get_traversableMap() gopurs_runtime.Value {
	once_traversableMap.Do(func() {
		cache_traversableMap = gopurs_runtime.RecordDict4("Foldable1", "Functor0", "sequence", "traverse", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_foldableMap()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_functorMap()
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_traversableMap(), "traverse"), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]](dictApplicative_0))}, Get_identity())
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
Apply0_1_0 := gopurs_runtime.CoerceToStruct[pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}))
_ = Apply0_1_0
Functor0_2_1 := gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_2_1
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_4_2_24 gopurs_runtime.Value
_ = go__go_4_2_24
go__go_4_2_24 = gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
var __t6 gopurs_runtime.Value
{
if (v_5.Type == 9 && v_5.IntVal == 324739070 && v_5.UnsafePtr == nil) {
__t6 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "pure"), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(nil))})
goto end_branch_6
} else {

}
}
{
if (v_5.Type == 9 && v_5.IntVal == 324739070 && v_5.UnsafePtr != nil) {
__local_var_6_3 := (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_5.UnsafePtr).V0
_ = __local_var_6_3
__local_var_7_4 := (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_5.UnsafePtr).V2
_ = __local_var_7_4
__local_var_8_5 := (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_5.UnsafePtr).V1
_ = __local_var_8_5
__t6 = gopurs_runtime.Apply2(Apply0_1_0.V1, gopurs_runtime.Apply2(Apply0_1_0.V1, gopurs_runtime.Apply2(Functor0_2_1.V0, gopurs_runtime.Func(func(l_prime_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_prime_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(r_prime_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(&Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]{1, __local_var_6_3, __local_var_8_5, __local_var_7_4, v_prime_10, gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](l_prime_9), gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](r_prime_11)})}
})
})
}), gopurs_runtime.Apply(go__go_4_2_24, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_5.UnsafePtr).V4)})), gopurs_runtime.Apply(f_3, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_5.UnsafePtr).V3)), gopurs_runtime.Apply(go__go_4_2_24, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_5.UnsafePtr).V5)}))
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
	return cache_traversableMap
}

var cache_traversableWithIndexMap gopurs_runtime.Value
var once_traversableWithIndexMap sync.Once
func Get_traversableWithIndexMap() gopurs_runtime.Value {
	once_traversableWithIndexMap.Do(func() {
		cache_traversableWithIndexMap = gopurs_runtime.RecordDict4("FoldableWithIndex1", "FunctorWithIndex0", "Traversable2", "traverseWithIndex", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_foldableWithIndexMap()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_functorWithIndexMap()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_traversableMap()
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
Apply0_1_0 := gopurs_runtime.CoerceToStruct[pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}))
_ = Apply0_1_0
Functor0_2_1 := gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_2_1
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_4_2_25 gopurs_runtime.Value
_ = go__go_4_2_25
go__go_4_2_25 = gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
var __t6 gopurs_runtime.Value
{
if (v_5.Type == 9 && v_5.IntVal == 324739070 && v_5.UnsafePtr == nil) {
__t6 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "pure"), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(nil))})
goto end_branch_6
} else {

}
}
{
if (v_5.Type == 9 && v_5.IntVal == 324739070 && v_5.UnsafePtr != nil) {
__local_var_6_3 := (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_5.UnsafePtr).V0
_ = __local_var_6_3
__local_var_7_4 := (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_5.UnsafePtr).V2
_ = __local_var_7_4
__local_var_8_5 := (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_5.UnsafePtr).V1
_ = __local_var_8_5
__t6 = gopurs_runtime.Apply2(Apply0_1_0.V1, gopurs_runtime.Apply2(Apply0_1_0.V1, gopurs_runtime.Apply2(Functor0_2_1.V0, gopurs_runtime.Func(func(l_prime_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_prime_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(r_prime_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(&Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]{1, __local_var_6_3, __local_var_8_5, __local_var_7_4, v_prime_10, gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](l_prime_9), gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](r_prime_11)})}
})
})
}), gopurs_runtime.Apply(go__go_4_2_25, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_5.UnsafePtr).V4)})), gopurs_runtime.Apply2(f_3, __local_var_7_4, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_5.UnsafePtr).V3)), gopurs_runtime.Apply(go__go_4_2_25, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_5.UnsafePtr).V5)}))
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
	return cache_traversableWithIndexMap
}

var cache_values gopurs_runtime.Value
var once_values sync.Once
func Get_values() gopurs_runtime.Value {
	once_values.Do(func() {
		cache_values = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_foldableMap(), "foldr"), pkg_Data_List_Types.Get_Cons(), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(nil))})
	})
	return cache_values
}

var cache_foldSubmapBy gopurs_runtime.Value
var once_foldSubmapBy sync.Once
func Get_foldSubmapBy() gopurs_runtime.Value {
	once_foldSubmapBy.Do(func() {
		cache_foldSubmapBy = gopurs_runtime.Func6(func(dictOrd_0_box gopurs_runtime.Value, appendFn_1_box gopurs_runtime.Value, memptyValue_2_box gopurs_runtime.Value, kmin_3_box gopurs_runtime.Value, kmax_4_box gopurs_runtime.Value, f_5_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldSubmapBy(gopurs_runtime.CoerceToStruct[pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value]](dictOrd_0_box), appendFn_1_box, memptyValue_2_box, gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](kmin_3_box), gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](kmax_4_box), f_5_box)
})
	})
	return cache_foldSubmapBy
}

var cache_foldSubmap gopurs_runtime.Value
var once_foldSubmap sync.Once
func Get_foldSubmap() gopurs_runtime.Value {
	once_foldSubmap.Do(func() {
		cache_foldSubmap = gopurs_runtime.Func2(func(dictOrd_0_box gopurs_runtime.Value, dictMonoid_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldSubmap(gopurs_runtime.CoerceToStruct[pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value]](dictOrd_0_box), gopurs_runtime.CoerceToStruct[pkg_Data_Monoid.Constructor_Monoid[gopurs_runtime.Value]](dictMonoid_1_box))
})
	})
	return cache_foldSubmap
}

var cache_findMin gopurs_runtime.Value
var once_findMin sync.Once
func Get_findMin() gopurs_runtime.Value {
	once_findMin.Do(func() {
		cache_findMin = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_findMin(gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](v_0_box)))}
})
	})
	return cache_findMin
}

var cache_lookupGT gopurs_runtime.Value
var once_lookupGT sync.Once
func Get_lookupGT() gopurs_runtime.Value {
	once_lookupGT.Do(func() {
		cache_lookupGT = gopurs_runtime.Func2(func(dictOrd_0_box gopurs_runtime.Value, k_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_lookupGT(gopurs_runtime.CoerceToStruct[pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value]](dictOrd_0_box), k_1_box)
})
	})
	return cache_lookupGT
}

var cache_findMax gopurs_runtime.Value
var once_findMax sync.Once
func Get_findMax() gopurs_runtime.Value {
	once_findMax.Do(func() {
		cache_findMax = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_findMax(gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](v_0_box)))}
})
	})
	return cache_findMax
}

var cache_lookupLT gopurs_runtime.Value
var once_lookupLT sync.Once
func Get_lookupLT() gopurs_runtime.Value {
	once_lookupLT.Do(func() {
		cache_lookupLT = gopurs_runtime.Func2(func(dictOrd_0_box gopurs_runtime.Value, k_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_lookupLT(gopurs_runtime.CoerceToStruct[pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value]](dictOrd_0_box), k_1_box)
})
	})
	return cache_lookupLT
}

var cache_filterWithKey gopurs_runtime.Value
var once_filterWithKey sync.Once
func Get_filterWithKey() gopurs_runtime.Value {
	once_filterWithKey.Do(func() {
		cache_filterWithKey = gopurs_runtime.Func2(func(dictOrd_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_filterWithKey(gopurs_runtime.CoerceToStruct[pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value]](dictOrd_0_box), f_1_box)
})
	})
	return cache_filterWithKey
}

var cache_filterKeys gopurs_runtime.Value
var once_filterKeys sync.Once
func Get_filterKeys() gopurs_runtime.Value {
	once_filterKeys.Do(func() {
		cache_filterKeys = gopurs_runtime.Func2(func(dictOrd_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_filterKeys(gopurs_runtime.CoerceToStruct[pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value]](dictOrd_0_box), f_1_box)
})
	})
	return cache_filterKeys
}

var cache_filter gopurs_runtime.Value
var once_filter sync.Once
func Get_filter() gopurs_runtime.Value {
	once_filter.Do(func() {
		cache_filter = gopurs_runtime.Func2(func(dictOrd_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_filter(gopurs_runtime.CoerceToStruct[pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value]](dictOrd_0_box), x_1_box)
})
	})
	return cache_filter
}

var cache_eqMap gopurs_runtime.Value
var once_eqMap sync.Once
func Get_eqMap() gopurs_runtime.Value {
	once_eqMap.Do(func() {
		cache_eqMap = gopurs_runtime.Func2(func(dictEq_0_box gopurs_runtime.Value, dictEq1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_eqMap(dictEq_0_box, dictEq1_1_box)
})
	})
	return cache_eqMap
}

var cache_ordMap gopurs_runtime.Value
var once_ordMap sync.Once
func Get_ordMap() gopurs_runtime.Value {
	once_ordMap.Do(func() {
		cache_ordMap = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_ordMap(dictOrd_0_box)
})
	})
	return cache_ordMap
}

var cache_eq1Map gopurs_runtime.Value
var once_eq1Map sync.Once
func Get_eq1Map() gopurs_runtime.Value {
	once_eq1Map.Do(func() {
		cache_eq1Map = gopurs_runtime.Func(func(dictEq_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_eq1Map(dictEq_0_box)
})
	})
	return cache_eq1Map
}

var cache_ord1Map gopurs_runtime.Value
var once_ord1Map sync.Once
func Get_ord1Map() gopurs_runtime.Value {
	once_ord1Map.Do(func() {
		cache_ord1Map = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_ord1Map(dictOrd_0_box)
})
	})
	return cache_ord1Map
}

var cache_empty gopurs_runtime.Value
var once_empty sync.Once
func Get_empty() gopurs_runtime.Value {
	once_empty.Do(func() {
		cache_empty = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(nil))}
	})
	return cache_empty
}

var cache_fromFoldable gopurs_runtime.Value
var once_fromFoldable sync.Once
func Get_fromFoldable() gopurs_runtime.Value {
	once_fromFoldable.Do(func() {
		cache_fromFoldable = gopurs_runtime.Func2(func(dictOrd_0_box gopurs_runtime.Value, dictFoldable_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_fromFoldable(gopurs_runtime.CoerceToStruct[pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value]](dictOrd_0_box), gopurs_runtime.CoerceToStruct[pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value]](dictFoldable_1_box))
})
	})
	return cache_fromFoldable
}

var cache_fromFoldableWith gopurs_runtime.Value
var once_fromFoldableWith sync.Once
func Get_fromFoldableWith() gopurs_runtime.Value {
	once_fromFoldableWith.Do(func() {
		cache_fromFoldableWith = gopurs_runtime.Func3(func(dictOrd_0_box gopurs_runtime.Value, dictFoldable_1_box gopurs_runtime.Value, f_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_fromFoldableWith(gopurs_runtime.CoerceToStruct[pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value]](dictOrd_0_box), gopurs_runtime.CoerceToStruct[pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value]](dictFoldable_1_box), f_2_box)
})
	})
	return cache_fromFoldableWith
}

var cache_fromFoldableWithIndex gopurs_runtime.Value
var once_fromFoldableWithIndex sync.Once
func Get_fromFoldableWithIndex() gopurs_runtime.Value {
	once_fromFoldableWithIndex.Do(func() {
		cache_fromFoldableWithIndex = gopurs_runtime.Func2(func(dictOrd_0_box gopurs_runtime.Value, dictFoldableWithIndex_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_fromFoldableWithIndex(gopurs_runtime.CoerceToStruct[pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value]](dictOrd_0_box), gopurs_runtime.CoerceToStruct[pkg_Data_FoldableWithIndex.Constructor_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]](dictFoldableWithIndex_1_box))
})
	})
	return cache_fromFoldableWithIndex
}

var cache_monoidSemigroupMap gopurs_runtime.Value
var once_monoidSemigroupMap sync.Once
func Get_monoidSemigroupMap() gopurs_runtime.Value {
	once_monoidSemigroupMap.Do(func() {
		cache_monoidSemigroupMap = gopurs_runtime.Func3(func(_dollar__unused_0_box gopurs_runtime.Value, dictOrd_1_box gopurs_runtime.Value, dictSemigroup_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_monoidSemigroupMap(_dollar__unused_0_box, dictOrd_1_box, dictSemigroup_2_box)
})
	})
	return cache_monoidSemigroupMap
}

var cache_submap gopurs_runtime.Value
var once_submap sync.Once
func Get_submap() gopurs_runtime.Value {
	once_submap.Do(func() {
		cache_submap = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_submap(gopurs_runtime.CoerceToStruct[pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value]](dictOrd_0_box))
})
	})
	return cache_submap
}

var cache_unions gopurs_runtime.Value
var once_unions sync.Once
func Get_unions() gopurs_runtime.Value {
	once_unions.Do(func() {
		cache_unions = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_unions(gopurs_runtime.CoerceToStruct[pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value]](dictOrd_0_box))
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

var cache_delete gopurs_runtime.Value
var once_delete sync.Once
func Get_delete() gopurs_runtime.Value {
	once_delete.Do(func() {
		cache_delete = gopurs_runtime.Func2(func(dictOrd_0_box gopurs_runtime.Value, k_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_delete(gopurs_runtime.CoerceToStruct[pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value]](dictOrd_0_box), k_1_box)
})
	})
	return cache_delete
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

var cache_applyMap gopurs_runtime.Value
var once_applyMap sync.Once
func Get_applyMap() gopurs_runtime.Value {
	once_applyMap.Do(func() {
		cache_applyMap = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_applyMap(dictOrd_0_box)
})
	})
	return cache_applyMap
}

var cache_bindMap gopurs_runtime.Value
var once_bindMap sync.Once
func Get_bindMap() gopurs_runtime.Value {
	once_bindMap.Do(func() {
		cache_bindMap = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_bindMap(dictOrd_0_box)
})
	})
	return cache_bindMap
}

var cache_anyWithKey gopurs_runtime.Value
var once_anyWithKey sync.Once
func Get_anyWithKey() gopurs_runtime.Value {
	once_anyWithKey.Do(func() {
		cache_anyWithKey = gopurs_runtime.Func(func(predicate_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_anyWithKey(predicate_0_box)
})
	})
	return cache_anyWithKey
}

var cache_any gopurs_runtime.Value
var once_any sync.Once
func Get_any() gopurs_runtime.Value {
	once_any.Do(func() {
		cache_any = gopurs_runtime.Func(func(predicate_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_any(predicate_0_box)
})
	})
	return cache_any
}

var cache_alter gopurs_runtime.Value
var once_alter sync.Once
func Get_alter() gopurs_runtime.Value {
	once_alter.Do(func() {
		cache_alter = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_alter(gopurs_runtime.CoerceToStruct[pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value]](dictOrd_0_box))
})
	})
	return cache_alter
}

var cache_altMap gopurs_runtime.Value
var once_altMap sync.Once
func Get_altMap() gopurs_runtime.Value {
	once_altMap.Do(func() {
		cache_altMap = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_altMap(dictOrd_0_box)
})
	})
	return cache_altMap
}

var cache_plusMap gopurs_runtime.Value
var once_plusMap sync.Once
func Get_plusMap() gopurs_runtime.Value {
	once_plusMap.Do(func() {
		cache_plusMap = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_plusMap(dictOrd_0_box)
})
	})
	return cache_plusMap
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

var cache_eq__2843686287 gopurs_runtime.Value
var once_eq__2843686287 sync.Once
func Get_eq__2843686287() gopurs_runtime.Value {
	once_eq__2843686287.Do(func() {
		cache_eq__2843686287 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_eq__2843686287(__eta0_0_box, __eta1_1_box)
})
	})
	return cache_eq__2843686287
}

var cache_eq__2384498378 gopurs_runtime.Value
var once_eq__2384498378 sync.Once
func Get_eq__2384498378() gopurs_runtime.Value {
	once_eq__2384498378.Do(func() {
		cache_eq__2384498378 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_eq__2384498378(gopurs_runtime.CoerceToStruct[pkg_Data_Eq.Constructor_Eq[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_eq__2384498378
}

var cache_eq__2541178864 gopurs_runtime.Value
var once_eq__2541178864 sync.Once
func Get_eq__2541178864() gopurs_runtime.Value {
	once_eq__2541178864.Do(func() {
		cache_eq__2541178864 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_eq__2541178864(gopurs_runtime.CoerceToStruct[pkg_Data_Eq.Constructor_Eq[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_eq__2541178864
}

var cache_foldl__2151204251 gopurs_runtime.Value
var once_foldl__2151204251 sync.Once
func Get_foldl__2151204251() gopurs_runtime.Value {
	once_foldl__2151204251.Do(func() {
		cache_foldl__2151204251 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldl__2151204251(gopurs_runtime.CoerceToStruct[pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_foldl__2151204251
}

var cache_foldl__53736539 gopurs_runtime.Value
var once_foldl__53736539 sync.Once
func Get_foldl__53736539() gopurs_runtime.Value {
	once_foldl__53736539.Do(func() {
		cache_foldl__53736539 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldl__53736539(gopurs_runtime.CoerceToStruct[pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_foldl__53736539
}

var cache_foldl__22573083 gopurs_runtime.Value
var once_foldl__22573083 sync.Once
func Get_foldl__22573083() gopurs_runtime.Value {
	once_foldl__22573083.Do(func() {
		cache_foldl__22573083 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldl__22573083(gopurs_runtime.CoerceToStruct[pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_foldl__22573083
}

var cache_foldr__2151204251 gopurs_runtime.Value
var once_foldr__2151204251 sync.Once
func Get_foldr__2151204251() gopurs_runtime.Value {
	once_foldr__2151204251.Do(func() {
		cache_foldr__2151204251 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldr__2151204251(gopurs_runtime.CoerceToStruct[pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_foldr__2151204251
}

var cache_foldr__530094749 gopurs_runtime.Value
var once_foldr__530094749 sync.Once
func Get_foldr__530094749() gopurs_runtime.Value {
	once_foldr__530094749.Do(func() {
		cache_foldr__530094749 = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, z_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldr__530094749(f_0_box, gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](z_1_box))
})
	})
	return cache_foldr__530094749
}

var cache_foldlWithIndex__2986161357 gopurs_runtime.Value
var once_foldlWithIndex__2986161357 sync.Once
func Get_foldlWithIndex__2986161357() gopurs_runtime.Value {
	once_foldlWithIndex__2986161357.Do(func() {
		cache_foldlWithIndex__2986161357 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldlWithIndex__2986161357(gopurs_runtime.CoerceToStruct[pkg_Data_FoldableWithIndex.Constructor_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_foldlWithIndex__2986161357
}

var cache_foldlWithIndex__2499716749 gopurs_runtime.Value
var once_foldlWithIndex__2499716749 sync.Once
func Get_foldlWithIndex__2499716749() gopurs_runtime.Value {
	once_foldlWithIndex__2499716749.Do(func() {
		cache_foldlWithIndex__2499716749 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldlWithIndex__2499716749(gopurs_runtime.CoerceToStruct[pkg_Data_FoldableWithIndex.Constructor_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_foldlWithIndex__2499716749
}

var cache_foldrWithIndex__2986161357 gopurs_runtime.Value
var once_foldrWithIndex__2986161357 sync.Once
func Get_foldrWithIndex__2986161357() gopurs_runtime.Value {
	once_foldrWithIndex__2986161357.Do(func() {
		cache_foldrWithIndex__2986161357 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldrWithIndex__2986161357(gopurs_runtime.CoerceToStruct[pkg_Data_FoldableWithIndex.Constructor_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_foldrWithIndex__2986161357
}

var cache_foldrWithIndex__3511467915 gopurs_runtime.Value
var once_foldrWithIndex__3511467915 sync.Once
func Get_foldrWithIndex__3511467915() gopurs_runtime.Value {
	once_foldrWithIndex__3511467915.Do(func() {
		cache_foldrWithIndex__3511467915 = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, z_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldrWithIndex__3511467915(f_0_box, gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](z_1_box))
})
	})
	return cache_foldrWithIndex__3511467915
}

var cache_const__1496134642 gopurs_runtime.Value
var once_const__1496134642 sync.Once
func Get_const__1496134642() gopurs_runtime.Value {
	once_const__1496134642.Do(func() {
		cache_const__1496134642 = gopurs_runtime.Func2(func(a_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_const__1496134642(a_0_box, v_1_box)
})
	})
	return cache_const__1496134642
}

var cache_const__2390202835 gopurs_runtime.Value
var once_const__2390202835 sync.Once
func Get_const__2390202835() gopurs_runtime.Value {
	once_const__2390202835.Do(func() {
		cache_const__2390202835 = gopurs_runtime.Func2(func(a_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_const__2390202835((a_0_box.IntVal) != (0), v_1_box))
})
	})
	return cache_const__2390202835
}

var cache_const__702735379 gopurs_runtime.Value
var once_const__702735379 sync.Once
func Get_const__702735379() gopurs_runtime.Value {
	once_const__702735379.Do(func() {
		cache_const__702735379 = gopurs_runtime.Func2(func(a_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_const__702735379(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](a_0_box), v_1_box))}
})
	})
	return cache_const__702735379
}

var cache_const__220790420 gopurs_runtime.Value
var once_const__220790420 sync.Once
func Get_const__220790420() gopurs_runtime.Value {
	once_const__220790420.Do(func() {
		cache_const__220790420 = gopurs_runtime.Func2(func(a_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_const__220790420(a_0_box, v_1_box)
})
	})
	return cache_const__220790420
}

var cache_const__641934996 gopurs_runtime.Value
var once_const__641934996 sync.Once
func Get_const__641934996() gopurs_runtime.Value {
	once_const__641934996.Do(func() {
		cache_const__641934996 = gopurs_runtime.Func2(func(a_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_const__641934996(a_0_box, v_1_box)
})
	})
	return cache_const__641934996
}

var cache_const__1628252324 gopurs_runtime.Value
var once_const__1628252324 sync.Once
func Get_const__1628252324() gopurs_runtime.Value {
	once_const__1628252324.Do(func() {
		cache_const__1628252324 = gopurs_runtime.Func2(func(a_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_const__1628252324(a_0_box, v_1_box)
})
	})
	return cache_const__1628252324
}

var cache_flip__3525081280 gopurs_runtime.Value
var once_flip__3525081280 sync.Once
func Get_flip__3525081280() gopurs_runtime.Value {
	once_flip__3525081280.Do(func() {
		cache_flip__3525081280 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_flip__3525081280(f_0_box, b_1_box, a_2_box)
})
	})
	return cache_flip__3525081280
}

var cache_flip__1172723328 gopurs_runtime.Value
var once_flip__1172723328 sync.Once
func Get_flip__1172723328() gopurs_runtime.Value {
	once_flip__1172723328.Do(func() {
		cache_flip__1172723328 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_flip__1172723328(f_0_box, b_1_box, a_2_box)
})
	})
	return cache_flip__1172723328
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

var cache_map__4039429788 gopurs_runtime.Value
var once_map__4039429788 sync.Once
func Get_map__4039429788() gopurs_runtime.Value {
	once_map__4039429788.Do(func() {
		cache_map__4039429788 = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_map__4039429788(v_0_box, gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](v1_1_box)))}
})
	})
	return cache_map__4039429788
}

var cache_conj__3676519832 gopurs_runtime.Value
var once_conj__3676519832 sync.Once
func Get_conj__3676519832() gopurs_runtime.Value {
	once_conj__3676519832.Do(func() {
		cache_conj__3676519832 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_conj__3676519832(__eta0_0_box, __eta1_1_box)
})
	})
	return cache_conj__3676519832
}

var cache_conj__3472268504 gopurs_runtime.Value
var once_conj__3472268504 sync.Once
func Get_conj__3472268504() gopurs_runtime.Value {
	once_conj__3472268504.Do(func() {
		cache_conj__3472268504 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_conj__3472268504(gopurs_runtime.CoerceToStruct[pkg_Data_HeytingAlgebra.Constructor_HeytingAlgebra[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_conj__3472268504
}

var cache_disj__3676519832 gopurs_runtime.Value
var once_disj__3676519832 sync.Once
func Get_disj__3676519832() gopurs_runtime.Value {
	once_disj__3676519832.Do(func() {
		cache_disj__3676519832 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_disj__3676519832(__eta0_0_box, __eta1_1_box)
})
	})
	return cache_disj__3676519832
}

var cache_disj__3472268504 gopurs_runtime.Value
var once_disj__3472268504 sync.Once
func Get_disj__3472268504() gopurs_runtime.Value {
	once_disj__3472268504.Do(func() {
		cache_disj__3472268504 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_disj__3472268504(gopurs_runtime.CoerceToStruct[pkg_Data_HeytingAlgebra.Constructor_HeytingAlgebra[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_disj__3472268504
}

var cache_not__3201284355 gopurs_runtime.Value
var once_not__3201284355 sync.Once
func Get_not__3201284355() gopurs_runtime.Value {
	once_not__3201284355.Do(func() {
		cache_not__3201284355 = gopurs_runtime.Func(func(__eta0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_not__3201284355(__eta0_0_box)
})
	})
	return cache_not__3201284355
}

var cache_not__1505204753 gopurs_runtime.Value
var once_not__1505204753 sync.Once
func Get_not__1505204753() gopurs_runtime.Value {
	once_not__1505204753.Do(func() {
		cache_not__1505204753 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_not__1505204753(gopurs_runtime.CoerceToStruct[pkg_Data_HeytingAlgebra.Constructor_HeytingAlgebra[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_not__1505204753
}

var cache_empty__1818220131 gopurs_runtime.Value
var once_empty__1818220131 sync.Once
func Get_empty__1818220131() gopurs_runtime.Value {
	once_empty__1818220131.Do(func() {
		cache_empty__1818220131 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(nil))}
	})
	return cache_empty__1818220131
}

var cache_empty__1794046843 gopurs_runtime.Value
var once_empty__1794046843 sync.Once
func Get_empty__1794046843() gopurs_runtime.Value {
	once_empty__1794046843.Do(func() {
		cache_empty__1794046843 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(nil))}
	})
	return cache_empty__1794046843
}

var cache_findMax__528468393 gopurs_runtime.Value
var once_findMax__528468393 sync.Once
func Get_findMax__528468393() gopurs_runtime.Value {
	once_findMax__528468393.Do(func() {
		cache_findMax__528468393 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_findMax__528468393(gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](v_0_box)))}
})
	})
	return cache_findMax__528468393
}

var cache_findMin__528468393 gopurs_runtime.Value
var once_findMin__528468393 sync.Once
func Get_findMin__528468393() gopurs_runtime.Value {
	once_findMin__528468393.Do(func() {
		cache_findMin__528468393 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_findMin__528468393(gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](v_0_box)))}
})
	})
	return cache_findMin__528468393
}

var cache_foldSubmapBy__3050108409 gopurs_runtime.Value
var once_foldSubmapBy__3050108409 sync.Once
func Get_foldSubmapBy__3050108409() gopurs_runtime.Value {
	once_foldSubmapBy__3050108409.Do(func() {
		cache_foldSubmapBy__3050108409 = gopurs_runtime.Func6(func(dictOrd_0_box gopurs_runtime.Value, appendFn_1_box gopurs_runtime.Value, memptyValue_2_box gopurs_runtime.Value, kmin_3_box gopurs_runtime.Value, kmax_4_box gopurs_runtime.Value, f_5_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldSubmapBy__3050108409(gopurs_runtime.CoerceToStruct[pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value]](dictOrd_0_box), appendFn_1_box, memptyValue_2_box, gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](kmin_3_box), gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](kmax_4_box), f_5_box)
})
	})
	return cache_foldSubmapBy__3050108409
}

var cache_foldSubmapBy__3128450809 gopurs_runtime.Value
var once_foldSubmapBy__3128450809 sync.Once
func Get_foldSubmapBy__3128450809() gopurs_runtime.Value {
	once_foldSubmapBy__3128450809.Do(func() {
		cache_foldSubmapBy__3128450809 = gopurs_runtime.Func6(func(dictOrd_0_box gopurs_runtime.Value, appendFn_1_box gopurs_runtime.Value, memptyValue_2_box gopurs_runtime.Value, kmin_3_box gopurs_runtime.Value, kmax_4_box gopurs_runtime.Value, f_5_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldSubmapBy__3128450809(gopurs_runtime.CoerceToStruct[pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value]](dictOrd_0_box), appendFn_1_box, memptyValue_2_box, gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](kmin_3_box), gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](kmax_4_box), f_5_box)
})
	})
	return cache_foldSubmapBy__3128450809
}

var cache_foldableMap__767959947 gopurs_runtime.Value
var once_foldableMap__767959947 sync.Once
func Get_foldableMap__767959947() gopurs_runtime.Value {
	once_foldableMap__767959947.Do(func() {
		cache_foldableMap__767959947 = gopurs_runtime.RecordDict3("foldMap", "foldl", "foldr", gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
Semigroup0_1_0 := gopurs_runtime.CoerceToStruct[pkg_Data_Semigroup.Constructor_Semigroup[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_1_0
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_3_1_40 gopurs_runtime.Value
_ = go__go_3_1_40
go__go_3_1_40 = gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
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
__t2 = gopurs_runtime.Apply2(Semigroup0_1_0.V0, gopurs_runtime.Apply(go__go_3_1_40, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V4)}), gopurs_runtime.Apply2(Semigroup0_1_0.V0, gopurs_runtime.Apply(f_2, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V3), gopurs_runtime.Apply(go__go_3_1_40, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V5)})))
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
return go__go_3_1_40
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(z_1 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_2_3_41 gopurs_runtime.Value
go__go_2_3_41 = gopurs_runtime.Func(func(__local_var_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(__local_var_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var __local_var_3_loop gopurs_runtime.Value = __local_var_3_loop_val
var __local_var_4_loop gopurs_runtime.Value = __local_var_4_loop_val
go__go_2_3_41:
for {
if false { continue go__go_2_3_41 }
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
__t4 = gopurs_runtime.UncurriedApp2(go__go_2_3_41, gopurs_runtime.Apply2(f_0, gopurs_runtime.UncurriedApp2(go__go_2_3_41, __local_var_3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_4.UnsafePtr).V4)}), (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_4.UnsafePtr).V3), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_4.UnsafePtr).V5)})
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
return gopurs_runtime.UncurriedApp2(go__go_2_3_41, z_1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](m_3))})
})
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(z_1 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_2_5_42 gopurs_runtime.Value
_ = go__go_2_5_42
go__go_2_5_42 = gopurs_runtime.Func2(func(__local_var_3 gopurs_runtime.Value, __local_var_4 gopurs_runtime.Value) gopurs_runtime.Value {
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
__t6 = gopurs_runtime.UncurriedApp2(go__go_2_5_42, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V4)}, gopurs_runtime.Apply2(f_0, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V3, gopurs_runtime.UncurriedApp2(go__go_2_5_42, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V5)}, __local_var_4)))
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
return gopurs_runtime.UncurriedApp2(go__go_2_5_42, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](m_3))}, z_1)
})
})
}))
	})
	return cache_foldableMap__767959947
}

var cache_foldableMap__373570208 gopurs_runtime.Value
var once_foldableMap__373570208 sync.Once
func Get_foldableMap__373570208() gopurs_runtime.Value {
	once_foldableMap__373570208.Do(func() {
		cache_foldableMap__373570208 = gopurs_runtime.RecordDict3("foldMap", "foldl", "foldr", gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
Semigroup0_1_0 := gopurs_runtime.CoerceToStruct[pkg_Data_Semigroup.Constructor_Semigroup[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_1_0
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_3_1_43 gopurs_runtime.Value
_ = go__go_3_1_43
go__go_3_1_43 = gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
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
__t2 = gopurs_runtime.Apply2(Semigroup0_1_0.V0, gopurs_runtime.Apply(go__go_3_1_43, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V4)}), gopurs_runtime.Apply2(Semigroup0_1_0.V0, gopurs_runtime.Apply(f_2, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V3), gopurs_runtime.Apply(go__go_3_1_43, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V5)})))
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
return go__go_3_1_43
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(z_1 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_2_3_44 gopurs_runtime.Value
go__go_2_3_44 = gopurs_runtime.Func(func(__local_var_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(__local_var_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var __local_var_3_loop gopurs_runtime.Value = __local_var_3_loop_val
var __local_var_4_loop gopurs_runtime.Value = __local_var_4_loop_val
go__go_2_3_44:
for {
if false { continue go__go_2_3_44 }
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
__t4 = gopurs_runtime.UncurriedApp2(go__go_2_3_44, gopurs_runtime.Apply2(f_0, gopurs_runtime.UncurriedApp2(go__go_2_3_44, __local_var_3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_4.UnsafePtr).V4)}), (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_4.UnsafePtr).V3), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_4.UnsafePtr).V5)})
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
return gopurs_runtime.UncurriedApp2(go__go_2_3_44, z_1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](m_3))})
})
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(z_1 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_2_5_45 gopurs_runtime.Value
_ = go__go_2_5_45
go__go_2_5_45 = gopurs_runtime.Func2(func(__local_var_3 gopurs_runtime.Value, __local_var_4 gopurs_runtime.Value) gopurs_runtime.Value {
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
__t6 = gopurs_runtime.UncurriedApp2(go__go_2_5_45, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V4)}, gopurs_runtime.Apply2(f_0, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V3, gopurs_runtime.UncurriedApp2(go__go_2_5_45, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V5)}, __local_var_4)))
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
return gopurs_runtime.UncurriedApp2(go__go_2_5_45, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](m_3))}, z_1)
})
})
}))
	})
	return cache_foldableMap__373570208
}

var cache_foldableWithIndexMap__1634502082 gopurs_runtime.Value
var once_foldableWithIndexMap__1634502082 sync.Once
func Get_foldableWithIndexMap__1634502082() gopurs_runtime.Value {
	once_foldableWithIndexMap__1634502082.Do(func() {
		cache_foldableWithIndexMap__1634502082 = gopurs_runtime.RecordDict4("Foldable0", "foldMapWithIndex", "foldlWithIndex", "foldrWithIndex", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_foldableMap()
}), gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
Semigroup0_1_0 := gopurs_runtime.CoerceToStruct[pkg_Data_Semigroup.Constructor_Semigroup[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_1_0
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_3_1_46 gopurs_runtime.Value
_ = go__go_3_1_46
go__go_3_1_46 = gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
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
__t2 = gopurs_runtime.Apply2(Semigroup0_1_0.V0, gopurs_runtime.Apply(go__go_3_1_46, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V4)}), gopurs_runtime.Apply2(Semigroup0_1_0.V0, gopurs_runtime.Apply2(f_2, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V2, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V3), gopurs_runtime.Apply(go__go_3_1_46, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V5)})))
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
return go__go_3_1_46
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(z_1 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_2_3_47 gopurs_runtime.Value
go__go_2_3_47 = gopurs_runtime.Func(func(__local_var_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(__local_var_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var __local_var_3_loop gopurs_runtime.Value = __local_var_3_loop_val
var __local_var_4_loop gopurs_runtime.Value = __local_var_4_loop_val
go__go_2_3_47:
for {
if false { continue go__go_2_3_47 }
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
__t4 = gopurs_runtime.UncurriedApp2(go__go_2_3_47, gopurs_runtime.Apply3(f_0, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_4.UnsafePtr).V2, gopurs_runtime.UncurriedApp2(go__go_2_3_47, __local_var_3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_4.UnsafePtr).V4)}), (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_4.UnsafePtr).V3), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_4.UnsafePtr).V5)})
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
return gopurs_runtime.UncurriedApp2(go__go_2_3_47, z_1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](m_3))})
})
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(z_1 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_2_5_48 gopurs_runtime.Value
_ = go__go_2_5_48
go__go_2_5_48 = gopurs_runtime.Func2(func(__local_var_3 gopurs_runtime.Value, __local_var_4 gopurs_runtime.Value) gopurs_runtime.Value {
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
__t6 = gopurs_runtime.UncurriedApp2(go__go_2_5_48, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V4)}, gopurs_runtime.Apply3(f_0, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V2, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V3, gopurs_runtime.UncurriedApp2(go__go_2_5_48, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V5)}, __local_var_4)))
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
return gopurs_runtime.UncurriedApp2(go__go_2_5_48, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](m_3))}, z_1)
})
})
}))
	})
	return cache_foldableWithIndexMap__1634502082
}

var cache_foldableWithIndexMap__1966365627 gopurs_runtime.Value
var once_foldableWithIndexMap__1966365627 sync.Once
func Get_foldableWithIndexMap__1966365627() gopurs_runtime.Value {
	once_foldableWithIndexMap__1966365627.Do(func() {
		cache_foldableWithIndexMap__1966365627 = gopurs_runtime.RecordDict4("Foldable0", "foldMapWithIndex", "foldlWithIndex", "foldrWithIndex", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_foldableMap()
}), gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
Semigroup0_1_0 := gopurs_runtime.CoerceToStruct[pkg_Data_Semigroup.Constructor_Semigroup[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_1_0
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_3_1_49 gopurs_runtime.Value
_ = go__go_3_1_49
go__go_3_1_49 = gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
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
__t2 = gopurs_runtime.Apply2(Semigroup0_1_0.V0, gopurs_runtime.Apply(go__go_3_1_49, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V4)}), gopurs_runtime.Apply2(Semigroup0_1_0.V0, gopurs_runtime.Apply2(f_2, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V2, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V3), gopurs_runtime.Apply(go__go_3_1_49, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V5)})))
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
return go__go_3_1_49
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(z_1 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_2_3_50 gopurs_runtime.Value
go__go_2_3_50 = gopurs_runtime.Func(func(__local_var_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(__local_var_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var __local_var_3_loop gopurs_runtime.Value = __local_var_3_loop_val
var __local_var_4_loop gopurs_runtime.Value = __local_var_4_loop_val
go__go_2_3_50:
for {
if false { continue go__go_2_3_50 }
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
__t4 = gopurs_runtime.UncurriedApp2(go__go_2_3_50, gopurs_runtime.Apply3(f_0, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_4.UnsafePtr).V2, gopurs_runtime.UncurriedApp2(go__go_2_3_50, __local_var_3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_4.UnsafePtr).V4)}), (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_4.UnsafePtr).V3), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_4.UnsafePtr).V5)})
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
return gopurs_runtime.UncurriedApp2(go__go_2_3_50, z_1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](m_3))})
})
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(z_1 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_2_5_51 gopurs_runtime.Value
_ = go__go_2_5_51
go__go_2_5_51 = gopurs_runtime.Func2(func(__local_var_3 gopurs_runtime.Value, __local_var_4 gopurs_runtime.Value) gopurs_runtime.Value {
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
__t6 = gopurs_runtime.UncurriedApp2(go__go_2_5_51, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V4)}, gopurs_runtime.Apply3(f_0, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V2, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V3, gopurs_runtime.UncurriedApp2(go__go_2_5_51, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V5)}, __local_var_4)))
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
return gopurs_runtime.UncurriedApp2(go__go_2_5_51, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](m_3))}, z_1)
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
var go__go_1_0_52 gopurs_runtime.Value
_ = go__go_1_0_52
go__go_1_0_52 = gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if (v_2.Type == 9 && v_2.IntVal == 324739070 && v_2.UnsafePtr == nil) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(nil))}
goto end_branch_1
} else {

}
}
{
if (v_2.Type == 9 && v_2.IntVal == 324739070 && v_2.UnsafePtr != nil) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(&Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V0, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V1, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V2, gopurs_runtime.Apply(f_0, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V3), gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(go__go_1_0_52, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V4)})), gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(go__go_1_0_52, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V5)}))})}
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__t1))}
})
return go__go_1_0_52
}))
	})
	return cache_functorMap__2501170929
}

var cache_functorWithIndexMap__3138419015 gopurs_runtime.Value
var once_functorWithIndexMap__3138419015 sync.Once
func Get_functorWithIndexMap__3138419015() gopurs_runtime.Value {
	once_functorWithIndexMap__3138419015.Do(func() {
		cache_functorWithIndexMap__3138419015 = gopurs_runtime.RecordDict2("Functor0", "mapWithIndex", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_functorMap()
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_1_0_53 gopurs_runtime.Value
_ = go__go_1_0_53
go__go_1_0_53 = gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if (v_2.Type == 9 && v_2.IntVal == 324739070 && v_2.UnsafePtr == nil) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(nil))}
goto end_branch_1
} else {

}
}
{
if (v_2.Type == 9 && v_2.IntVal == 324739070 && v_2.UnsafePtr != nil) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(&Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V0, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V1, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V2, gopurs_runtime.Apply2(f_0, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V2, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V3), gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(go__go_1_0_53, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V4)})), gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(go__go_1_0_53, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V5)}))})}
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__t1))}
})
return go__go_1_0_53
}))
	})
	return cache_functorWithIndexMap__3138419015
}

var cache_insert__4289641298 gopurs_runtime.Value
var once_insert__4289641298 sync.Once
func Get_insert__4289641298() gopurs_runtime.Value {
	once_insert__4289641298.Do(func() {
		cache_insert__4289641298 = gopurs_runtime.Func3(func(dictOrd_0_box gopurs_runtime.Value, k_1_box gopurs_runtime.Value, v_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_insert__4289641298(gopurs_runtime.CoerceToStruct[pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value]](dictOrd_0_box), k_1_box, v_2_box)
})
	})
	return cache_insert__4289641298
}

var cache_insertWith__118979962 gopurs_runtime.Value
var once_insertWith__118979962 sync.Once
func Get_insertWith__118979962() gopurs_runtime.Value {
	once_insertWith__118979962.Do(func() {
		cache_insertWith__118979962 = gopurs_runtime.Func4(func(dictOrd_0_box gopurs_runtime.Value, app_1_box gopurs_runtime.Value, k_2_box gopurs_runtime.Value, v_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_insertWith__118979962(gopurs_runtime.CoerceToStruct[pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value]](dictOrd_0_box), app_1_box, k_2_box, v_3_box)
})
	})
	return cache_insertWith__118979962
}

var cache_intersectionWith__3717755541 gopurs_runtime.Value
var once_intersectionWith__3717755541 sync.Once
func Get_intersectionWith__3717755541() gopurs_runtime.Value {
	once_intersectionWith__3717755541.Do(func() {
		cache_intersectionWith__3717755541 = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_intersectionWith__3717755541(gopurs_runtime.CoerceToStruct[pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value]](dictOrd_0_box))
})
	})
	return cache_intersectionWith__3717755541
}

var cache_intersectionWith__4144106805 gopurs_runtime.Value
var once_intersectionWith__4144106805 sync.Once
func Get_intersectionWith__4144106805() gopurs_runtime.Value {
	once_intersectionWith__4144106805.Do(func() {
		cache_intersectionWith__4144106805 = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_intersectionWith__4144106805(gopurs_runtime.CoerceToStruct[pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value]](dictOrd_0_box))
})
	})
	return cache_intersectionWith__4144106805
}

var cache_iterMapL__878452066 gopurs_runtime.Value
var once_iterMapL__878452066 sync.Once
func Get_iterMapL__878452066() gopurs_runtime.Value {
	once_iterMapL__878452066.Do(func() {
		cache_iterMapL__878452066 = func() gopurs_runtime.Value {
var go__go_0_0_56 gopurs_runtime.Value
go__go_0_0_56 = gopurs_runtime.Func(func(iter_1_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var iter_1_loop gopurs_runtime.Value = iter_1_loop_val
var v_2_loop gopurs_runtime.Value = v_2_loop_val
go__go_0_0_56:
for {
if false { continue go__go_0_0_56 }
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
var __t_tag_1 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V5)}
if (__t_tag_1.Type == 9 && __t_tag_1.IntVal == 324739070 && __t_tag_1.UnsafePtr == nil) {
iter_1_loop = gopurs_runtime.Value{Type: 9, IntVal: 1343415489, UnsafePtr: unsafe.Pointer(&Constructor_IterEmit[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V2, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V3, iter_1})}
v_2_loop = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V4)}
continue go__go_0_0_56
__t2 = gopurs_runtime.Value{}
goto end_branch_2
} else {

}
}
{
iter_1_loop = gopurs_runtime.Value{Type: 9, IntVal: 1343415489, UnsafePtr: unsafe.Pointer(&Constructor_IterEmit[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V2, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 2861335956, UnsafePtr: unsafe.Pointer(&Constructor_IterNode[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V5, iter_1})}})}
v_2_loop = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V4)}
continue go__go_0_0_56
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
return go__go_0_0_56
}()
	})
	return cache_iterMapL__878452066
}

var cache_iterMapR__878452066 gopurs_runtime.Value
var once_iterMapR__878452066 sync.Once
func Get_iterMapR__878452066() gopurs_runtime.Value {
	once_iterMapR__878452066.Do(func() {
		cache_iterMapR__878452066 = func() gopurs_runtime.Value {
var go__go_0_0_57 gopurs_runtime.Value
go__go_0_0_57 = gopurs_runtime.Func(func(iter_1_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var iter_1_loop gopurs_runtime.Value = iter_1_loop_val
var v_2_loop gopurs_runtime.Value = v_2_loop_val
go__go_0_0_57:
for {
if false { continue go__go_0_0_57 }
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
var __t_tag_1 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V5)}
if (__t_tag_1.Type == 9 && __t_tag_1.IntVal == 324739070 && __t_tag_1.UnsafePtr == nil) {
iter_1_loop = gopurs_runtime.Value{Type: 9, IntVal: 1343415489, UnsafePtr: unsafe.Pointer(&Constructor_IterEmit[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V2, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V3, iter_1})}
v_2_loop = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V4)}
continue go__go_0_0_57
__t2 = gopurs_runtime.Value{}
goto end_branch_2
} else {

}
}
{
iter_1_loop = gopurs_runtime.Value{Type: 9, IntVal: 1343415489, UnsafePtr: unsafe.Pointer(&Constructor_IterEmit[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V2, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 2861335956, UnsafePtr: unsafe.Pointer(&Constructor_IterNode[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V4, iter_1})}})}
v_2_loop = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V5)}
continue go__go_0_0_57
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
return go__go_0_0_57
}()
	})
	return cache_iterMapR__878452066
}

var cache_iterMapU__878452066 gopurs_runtime.Value
var once_iterMapU__878452066 sync.Once
func Get_iterMapU__878452066() gopurs_runtime.Value {
	once_iterMapU__878452066.Do(func() {
		cache_iterMapU__878452066 = gopurs_runtime.Func2(func(iter_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_iterMapU__878452066(iter_0_box, gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](v_1_box))
})
	})
	return cache_iterMapU__878452066
}

var cache_lookup__3378638282 gopurs_runtime.Value
var once_lookup__3378638282 sync.Once
func Get_lookup__3378638282() gopurs_runtime.Value {
	once_lookup__3378638282.Do(func() {
		cache_lookup__3378638282 = gopurs_runtime.Func2(func(dictOrd_0_box gopurs_runtime.Value, k_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_lookup__3378638282(gopurs_runtime.CoerceToStruct[pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value]](dictOrd_0_box), k_1_box)
})
	})
	return cache_lookup__3378638282
}

var cache_mapMaybe__3426301240 gopurs_runtime.Value
var once_mapMaybe__3426301240 sync.Once
func Get_mapMaybe__3426301240() gopurs_runtime.Value {
	once_mapMaybe__3426301240.Do(func() {
		cache_mapMaybe__3426301240 = gopurs_runtime.Func2(func(dictOrd_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_mapMaybe__3426301240(gopurs_runtime.CoerceToStruct[pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value]](dictOrd_0_box), x_1_box)
})
	})
	return cache_mapMaybe__3426301240
}

var cache_mapMaybe__1970555288 gopurs_runtime.Value
var once_mapMaybe__1970555288 sync.Once
func Get_mapMaybe__1970555288() gopurs_runtime.Value {
	once_mapMaybe__1970555288.Do(func() {
		cache_mapMaybe__1970555288 = gopurs_runtime.Func2(func(dictOrd_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_mapMaybe__1970555288(gopurs_runtime.CoerceToStruct[pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value]](dictOrd_0_box), x_1_box)
})
	})
	return cache_mapMaybe__1970555288
}

var cache_mapMaybeWithKey__817660689 gopurs_runtime.Value
var once_mapMaybeWithKey__817660689 sync.Once
func Get_mapMaybeWithKey__817660689() gopurs_runtime.Value {
	once_mapMaybeWithKey__817660689.Do(func() {
		cache_mapMaybeWithKey__817660689 = gopurs_runtime.Func2(func(dictOrd_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_mapMaybeWithKey__817660689(gopurs_runtime.CoerceToStruct[pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value]](dictOrd_0_box), f_1_box)
})
	})
	return cache_mapMaybeWithKey__817660689
}

var cache_singleton__3511563426 gopurs_runtime.Value
var once_singleton__3511563426 sync.Once
func Get_singleton__3511563426() gopurs_runtime.Value {
	once_singleton__3511563426.Do(func() {
		cache_singleton__3511563426 = gopurs_runtime.Func2(func(k_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(Call_singleton__3511563426(k_0_box, v_1_box))}
})
	})
	return cache_singleton__3511563426
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

var cache_stepAsc__2098920977 gopurs_runtime.Value
var once_stepAsc__2098920977 sync.Once
func Get_stepAsc__2098920977() gopurs_runtime.Value {
	once_stepAsc__2098920977.Do(func() {
		cache_stepAsc__2098920977 = Call_stepWith(Get_iterMapL(), gopurs_runtime.Func3(func(__local_var_0 gopurs_runtime.Value, __local_var_1 gopurs_runtime.Value, __local_var_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 953589075, UnsafePtr: unsafe.Pointer(&Constructor_IterNext[gopurs_runtime.Value, gopurs_runtime.Value]{1, __local_var_0, __local_var_1, __local_var_2})}
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 953589075, UnsafePtr: unsafe.Pointer((*Constructor_IterNext[gopurs_runtime.Value, gopurs_runtime.Value])(nil))}
}))
	})
	return cache_stepAsc__2098920977
}

var cache_stepAscCps__3090303421 gopurs_runtime.Value
var once_stepAscCps__3090303421 sync.Once
func Get_stepAscCps__3090303421() gopurs_runtime.Value {
	once_stepAscCps__3090303421.Do(func() {
		cache_stepAscCps__3090303421 = gopurs_runtime.Apply(Get_stepWith(), Get_iterMapL())
	})
	return cache_stepAscCps__3090303421
}

var cache_stepAscCps__2463496949 gopurs_runtime.Value
var once_stepAscCps__2463496949 sync.Once
func Get_stepAscCps__2463496949() gopurs_runtime.Value {
	once_stepAscCps__2463496949.Do(func() {
		cache_stepAscCps__2463496949 = gopurs_runtime.Apply(Get_stepWith(), Get_iterMapL())
	})
	return cache_stepAscCps__2463496949
}

var cache_stepAscCps__1323290822 gopurs_runtime.Value
var once_stepAscCps__1323290822 sync.Once
func Get_stepAscCps__1323290822() gopurs_runtime.Value {
	once_stepAscCps__1323290822.Do(func() {
		cache_stepAscCps__1323290822 = gopurs_runtime.Apply(Get_stepWith(), Get_iterMapL())
	})
	return cache_stepAscCps__1323290822
}

var cache_stepDescCps__3090303421 gopurs_runtime.Value
var once_stepDescCps__3090303421 sync.Once
func Get_stepDescCps__3090303421() gopurs_runtime.Value {
	once_stepDescCps__3090303421.Do(func() {
		cache_stepDescCps__3090303421 = gopurs_runtime.Apply(Get_stepWith(), Get_iterMapR())
	})
	return cache_stepDescCps__3090303421
}

var cache_stepDescCps__2463496949 gopurs_runtime.Value
var once_stepDescCps__2463496949 sync.Once
func Get_stepDescCps__2463496949() gopurs_runtime.Value {
	once_stepDescCps__2463496949.Do(func() {
		cache_stepDescCps__2463496949 = gopurs_runtime.Apply(Get_stepWith(), Get_iterMapR())
	})
	return cache_stepDescCps__2463496949
}

var cache_stepUnfoldr__966001626 gopurs_runtime.Value
var once_stepUnfoldr__966001626 sync.Once
func Get_stepUnfoldr__966001626() gopurs_runtime.Value {
	once_stepUnfoldr__966001626.Do(func() {
		cache_stepUnfoldr__966001626 = Call_stepWith(Get_iterMapL(), gopurs_runtime.Func3(func(__local_var_0 gopurs_runtime.Value, __local_var_1 gopurs_runtime.Value, __local_var_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, __local_var_0, __local_var_1})}, __local_var_2})}})}
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(nil))}
}))
	})
	return cache_stepUnfoldr__966001626
}

var cache_stepUnfoldrUnordered__966001626 gopurs_runtime.Value
var once_stepUnfoldrUnordered__966001626 sync.Once
func Get_stepUnfoldrUnordered__966001626() gopurs_runtime.Value {
	once_stepUnfoldrUnordered__966001626.Do(func() {
		cache_stepUnfoldrUnordered__966001626 = Call_stepWith(Get_iterMapU(), gopurs_runtime.Func3(func(__local_var_0 gopurs_runtime.Value, __local_var_1 gopurs_runtime.Value, __local_var_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, __local_var_0, __local_var_1})}, __local_var_2})}})}
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(nil))}
}))
	})
	return cache_stepUnfoldrUnordered__966001626
}

var cache_stepUnorderedCps__3090303421 gopurs_runtime.Value
var once_stepUnorderedCps__3090303421 sync.Once
func Get_stepUnorderedCps__3090303421() gopurs_runtime.Value {
	once_stepUnorderedCps__3090303421.Do(func() {
		cache_stepUnorderedCps__3090303421 = gopurs_runtime.Apply(Get_stepWith(), Get_iterMapU())
	})
	return cache_stepUnorderedCps__3090303421
}

var cache_stepUnorderedCps__2463496949 gopurs_runtime.Value
var once_stepUnorderedCps__2463496949 sync.Once
func Get_stepUnorderedCps__2463496949() gopurs_runtime.Value {
	once_stepUnorderedCps__2463496949.Do(func() {
		cache_stepUnorderedCps__2463496949 = gopurs_runtime.Apply(Get_stepWith(), Get_iterMapU())
	})
	return cache_stepUnorderedCps__2463496949
}

var cache_stepUnorderedCps__1323290822 gopurs_runtime.Value
var once_stepUnorderedCps__1323290822 sync.Once
func Get_stepUnorderedCps__1323290822() gopurs_runtime.Value {
	once_stepUnorderedCps__1323290822.Do(func() {
		cache_stepUnorderedCps__1323290822 = gopurs_runtime.Apply(Get_stepWith(), Get_iterMapU())
	})
	return cache_stepUnorderedCps__1323290822
}

var cache_stepWith__3186376421 gopurs_runtime.Value
var once_stepWith__3186376421 sync.Once
func Get_stepWith__3186376421() gopurs_runtime.Value {
	once_stepWith__3186376421.Do(func() {
		cache_stepWith__3186376421 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, next_1_box gopurs_runtime.Value, done_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_stepWith__3186376421(f_0_box, next_1_box, done_2_box)
})
	})
	return cache_stepWith__3186376421
}

var cache_toMapIter__2014410513 gopurs_runtime.Value
var once_toMapIter__2014410513 sync.Once
func Get_toMapIter__2014410513() gopurs_runtime.Value {
	once_toMapIter__2014410513.Do(func() {
		cache_toMapIter__2014410513 = gopurs_runtime.Func(func(a_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_toMapIter__2014410513(gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](a_0_box))
})
	})
	return cache_toMapIter__2014410513
}

var cache_toMapIter__1738891721 gopurs_runtime.Value
var once_toMapIter__1738891721 sync.Once
func Get_toMapIter__1738891721() gopurs_runtime.Value {
	once_toMapIter__1738891721.Do(func() {
		cache_toMapIter__1738891721 = gopurs_runtime.Func(func(a_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_toMapIter__1738891721(gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](a_0_box))
})
	})
	return cache_toMapIter__1738891721
}

var cache_toUnfoldable__2183602684 gopurs_runtime.Value
var once_toUnfoldable__2183602684 sync.Once
func Get_toUnfoldable__2183602684() gopurs_runtime.Value {
	once_toUnfoldable__2183602684.Do(func() {
		cache_toUnfoldable__2183602684 = gopurs_runtime.Func(func(dictUnfoldable_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_toUnfoldable__2183602684(gopurs_runtime.CoerceToStruct[pkg_Data_Unfoldable.Constructor_Unfoldable[gopurs_runtime.Value]](dictUnfoldable_0_box))
})
	})
	return cache_toUnfoldable__2183602684
}

var cache_traversableMap__1002539403 gopurs_runtime.Value
var once_traversableMap__1002539403 sync.Once
func Get_traversableMap__1002539403() gopurs_runtime.Value {
	once_traversableMap__1002539403.Do(func() {
		cache_traversableMap__1002539403 = gopurs_runtime.RecordDict4("Foldable1", "Functor0", "sequence", "traverse", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_foldableMap()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_functorMap()
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_traversableMap(), "traverse"), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]](dictApplicative_0))}, Get_identity())
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
Apply0_1_0 := gopurs_runtime.CoerceToStruct[pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}))
_ = Apply0_1_0
Functor0_2_1 := gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_2_1
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_4_2_61 gopurs_runtime.Value
_ = go__go_4_2_61
go__go_4_2_61 = gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
var __t6 gopurs_runtime.Value
{
if (v_5.Type == 9 && v_5.IntVal == 324739070 && v_5.UnsafePtr == nil) {
__t6 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "pure"), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(nil))})
goto end_branch_6
} else {

}
}
{
if (v_5.Type == 9 && v_5.IntVal == 324739070 && v_5.UnsafePtr != nil) {
__local_var_6_3 := (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_5.UnsafePtr).V0
_ = __local_var_6_3
__local_var_7_4 := (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_5.UnsafePtr).V2
_ = __local_var_7_4
__local_var_8_5 := (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_5.UnsafePtr).V1
_ = __local_var_8_5
__t6 = gopurs_runtime.Apply2(Apply0_1_0.V1, gopurs_runtime.Apply2(Apply0_1_0.V1, gopurs_runtime.Apply2(Functor0_2_1.V0, gopurs_runtime.Func(func(l_prime_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_prime_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(r_prime_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(&Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]{1, __local_var_6_3, __local_var_8_5, __local_var_7_4, v_prime_10, gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](l_prime_9), gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](r_prime_11)})}
})
})
}), gopurs_runtime.Apply(go__go_4_2_61, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_5.UnsafePtr).V4)})), gopurs_runtime.Apply(f_3, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_5.UnsafePtr).V3)), gopurs_runtime.Apply(go__go_4_2_61, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_5.UnsafePtr).V5)}))
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
return go__go_4_2_61
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
return Get_foldableMap()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_functorMap()
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_traversableMap(), "traverse"), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]](dictApplicative_0))}, Get_identity())
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
Apply0_1_0 := gopurs_runtime.CoerceToStruct[pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}))
_ = Apply0_1_0
Functor0_2_1 := gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_2_1
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_4_2_62 gopurs_runtime.Value
_ = go__go_4_2_62
go__go_4_2_62 = gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
var __t6 gopurs_runtime.Value
{
if (v_5.Type == 9 && v_5.IntVal == 324739070 && v_5.UnsafePtr == nil) {
__t6 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "pure"), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(nil))})
goto end_branch_6
} else {

}
}
{
if (v_5.Type == 9 && v_5.IntVal == 324739070 && v_5.UnsafePtr != nil) {
__local_var_6_3 := (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_5.UnsafePtr).V0
_ = __local_var_6_3
__local_var_7_4 := (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_5.UnsafePtr).V2
_ = __local_var_7_4
__local_var_8_5 := (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_5.UnsafePtr).V1
_ = __local_var_8_5
__t6 = gopurs_runtime.Apply2(Apply0_1_0.V1, gopurs_runtime.Apply2(Apply0_1_0.V1, gopurs_runtime.Apply2(Functor0_2_1.V0, gopurs_runtime.Func(func(l_prime_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_prime_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(r_prime_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(&Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]{1, __local_var_6_3, __local_var_8_5, __local_var_7_4, v_prime_10, gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](l_prime_9), gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](r_prime_11)})}
})
})
}), gopurs_runtime.Apply(go__go_4_2_62, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_5.UnsafePtr).V4)})), gopurs_runtime.Apply(f_3, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_5.UnsafePtr).V3)), gopurs_runtime.Apply(go__go_4_2_62, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_5.UnsafePtr).V5)}))
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
return go__go_4_2_62
})
}))
	})
	return cache_traversableMap__2256206635
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

var cache_unsafeDifference__4097927905 gopurs_runtime.Value
var once_unsafeDifference__4097927905 sync.Once
func Get_unsafeDifference__4097927905() gopurs_runtime.Value {
	once_unsafeDifference__4097927905.Do(func() {
		cache_unsafeDifference__4097927905 = gopurs_runtime.Func3(func(__local_var_0_box gopurs_runtime.Value, __local_var_1_box gopurs_runtime.Value, __local_var_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_unsafeDifference__4097927905(__local_var_0_box, __local_var_1_box, __local_var_2_box)
})
	})
	return cache_unsafeDifference__4097927905
}

var cache_unsafeIntersectionWith__4109280494 gopurs_runtime.Value
var once_unsafeIntersectionWith__4109280494 sync.Once
func Get_unsafeIntersectionWith__4109280494() gopurs_runtime.Value {
	once_unsafeIntersectionWith__4109280494.Do(func() {
		cache_unsafeIntersectionWith__4109280494 = gopurs_runtime.Func4(func(__local_var_0_box gopurs_runtime.Value, __local_var_1_box gopurs_runtime.Value, __local_var_2_box gopurs_runtime.Value, __local_var_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_unsafeIntersectionWith__4109280494(__local_var_0_box, __local_var_1_box, __local_var_2_box, __local_var_3_box)
})
	})
	return cache_unsafeIntersectionWith__4109280494
}

var cache_unsafeJoinNodes__3967876672 gopurs_runtime.Value
var once_unsafeJoinNodes__3967876672 sync.Once
func Get_unsafeJoinNodes__3967876672() gopurs_runtime.Value {
	once_unsafeJoinNodes__3967876672.Do(func() {
		cache_unsafeJoinNodes__3967876672 = gopurs_runtime.Func2(func(__local_var_0_box gopurs_runtime.Value, __local_var_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_unsafeJoinNodes__3967876672(__local_var_0_box, __local_var_1_box)
})
	})
	return cache_unsafeJoinNodes__3967876672
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

var cache_unsafeSplitLast__224676098 gopurs_runtime.Value
var once_unsafeSplitLast__224676098 sync.Once
func Get_unsafeSplitLast__224676098() gopurs_runtime.Value {
	once_unsafeSplitLast__224676098.Do(func() {
		cache_unsafeSplitLast__224676098 = gopurs_runtime.Func4(func(__local_var_0_box gopurs_runtime.Value, __local_var_1_box gopurs_runtime.Value, __local_var_2_box gopurs_runtime.Value, __local_var_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_unsafeSplitLast__224676098(__local_var_0_box, __local_var_1_box, __local_var_2_box, __local_var_3_box)
})
	})
	return cache_unsafeSplitLast__224676098
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

var cache_fromJust__1791383420 gopurs_runtime.Value
var once_fromJust__1791383420 sync.Once
func Get_fromJust__1791383420() gopurs_runtime.Value {
	once_fromJust__1791383420.Do(func() {
		cache_fromJust__1791383420 = gopurs_runtime.Func2(func(_dollar__unused_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_fromJust__1791383420(_dollar__unused_0_box, gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](v_1_box))
})
	})
	return cache_fromJust__1791383420
}

var cache_functorMaybe__2569569018 gopurs_runtime.Value
var once_functorMaybe__2569569018 sync.Once
func Get_functorMaybe__2569569018() gopurs_runtime.Value {
	once_functorMaybe__2569569018.Do(func() {
		cache_functorMaybe__2569569018 = gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (v1_1.Type == 9 && v1_1.IntVal == 930809136 && v1_1.UnsafePtr != nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, gopurs_runtime.Apply(v_0, (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(v1_1.UnsafePtr).V0)})}
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(nil))}
}
end_branch_0:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](__t0))}
})
}))
	})
	return cache_functorMaybe__2569569018
}

var cache_isNothing__2591355336 gopurs_runtime.Value
var once_isNothing__2591355336 sync.Once
func Get_isNothing__2591355336() gopurs_runtime.Value {
	once_isNothing__2591355336.Do(func() {
		cache_isNothing__2591355336 = gopurs_runtime.Func(func(v2_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_isNothing__2591355336(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](v2_0_box)))
})
	})
	return cache_isNothing__2591355336
}

var cache_maybe__1594528518 gopurs_runtime.Value
var once_maybe__1594528518 sync.Once
func Get_maybe__1594528518() gopurs_runtime.Value {
	once_maybe__1594528518.Do(func() {
		cache_maybe__1594528518 = gopurs_runtime.Func3(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value, v2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_maybe__1594528518(v_0_box, v1_1_box, gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](v2_2_box))
})
	})
	return cache_maybe__1594528518
}

var cache_maybe__3658316244 gopurs_runtime.Value
var once_maybe__3658316244 sync.Once
func Get_maybe__3658316244() gopurs_runtime.Value {
	once_maybe__3658316244.Do(func() {
		cache_maybe__3658316244 = gopurs_runtime.Func3(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value, v2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_maybe__3658316244(v_0_box, v1_1_box, gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](v2_2_box))
})
	})
	return cache_maybe__3658316244
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

var cache_abs__1599282999 gopurs_runtime.Value
var once_abs__1599282999 sync.Once
func Get_abs__1599282999() gopurs_runtime.Value {
	once_abs__1599282999.Do(func() {
		cache_abs__1599282999 = gopurs_runtime.Func(func(__eta0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_abs__1599282999(__eta0_0_box)
})
	})
	return cache_abs__1599282999
}

var cache_abs__2515802711 gopurs_runtime.Value
var once_abs__2515802711 sync.Once
func Get_abs__2515802711() gopurs_runtime.Value {
	once_abs__2515802711.Do(func() {
		cache_abs__2515802711 = gopurs_runtime.Func2(func(dictOrd_0_box gopurs_runtime.Value, dictRing_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_abs__2515802711(gopurs_runtime.CoerceToStruct[pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value]](dictOrd_0_box), gopurs_runtime.CoerceToStruct[pkg_Data_Ring.Constructor_Ring[gopurs_runtime.Value]](dictRing_1_box))
})
	})
	return cache_abs__2515802711
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

var cache_compare__2802126154 gopurs_runtime.Value
var once_compare__2802126154 sync.Once
func Get_compare__2802126154() gopurs_runtime.Value {
	once_compare__2802126154.Do(func() {
		cache_compare__2802126154 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_compare__2802126154(gopurs_runtime.CoerceToStruct[pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_compare__2802126154
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

var cache_greaterThanOrEq__1697837627 gopurs_runtime.Value
var once_greaterThanOrEq__1697837627 sync.Once
func Get_greaterThanOrEq__1697837627() gopurs_runtime.Value {
	once_greaterThanOrEq__1697837627.Do(func() {
		cache_greaterThanOrEq__1697837627 = gopurs_runtime.Func3(func(dictOrd_0_box gopurs_runtime.Value, a1_1_box gopurs_runtime.Value, a2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_greaterThanOrEq__1697837627(gopurs_runtime.CoerceToStruct[pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value]](dictOrd_0_box), a1_1_box, a2_2_box))
})
	})
	return cache_greaterThanOrEq__1697837627
}

var cache_greaterThanOrEq__1409282474 gopurs_runtime.Value
var once_greaterThanOrEq__1409282474 sync.Once
func Get_greaterThanOrEq__1409282474() gopurs_runtime.Value {
	once_greaterThanOrEq__1409282474.Do(func() {
		cache_greaterThanOrEq__1409282474 = gopurs_runtime.Func3(func(dictOrd_0_box gopurs_runtime.Value, a1_1_box gopurs_runtime.Value, a2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_greaterThanOrEq__1409282474(gopurs_runtime.CoerceToStruct[pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value]](dictOrd_0_box), a1_1_box, a2_2_box))
})
	})
	return cache_greaterThanOrEq__1409282474
}

var cache_lessThan__4087042607 gopurs_runtime.Value
var once_lessThan__4087042607 sync.Once
func Get_lessThan__4087042607() gopurs_runtime.Value {
	once_lessThan__4087042607.Do(func() {
		cache_lessThan__4087042607 = gopurs_runtime.Func2(func(a1_0_box gopurs_runtime.Value, a2_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_lessThan__4087042607(a1_0_box, a2_1_box))
})
	})
	return cache_lessThan__4087042607
}

var cache_lessThan__1409282474 gopurs_runtime.Value
var once_lessThan__1409282474 sync.Once
func Get_lessThan__1409282474() gopurs_runtime.Value {
	once_lessThan__1409282474.Do(func() {
		cache_lessThan__1409282474 = gopurs_runtime.Func3(func(dictOrd_0_box gopurs_runtime.Value, a1_1_box gopurs_runtime.Value, a2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_lessThan__1409282474(gopurs_runtime.CoerceToStruct[pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value]](dictOrd_0_box), a1_1_box, a2_2_box))
})
	})
	return cache_lessThan__1409282474
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

var cache_negate__2674436160 gopurs_runtime.Value
var once_negate__2674436160 sync.Once
func Get_negate__2674436160() gopurs_runtime.Value {
	once_negate__2674436160.Do(func() {
		cache_negate__2674436160 = gopurs_runtime.Func(func(dictRing_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_negate__2674436160(gopurs_runtime.CoerceToStruct[pkg_Data_Ring.Constructor_Ring[gopurs_runtime.Value]](dictRing_0_box))
})
	})
	return cache_negate__2674436160
}

var cache_negate__1364373265 gopurs_runtime.Value
var once_negate__1364373265 sync.Once
func Get_negate__1364373265() gopurs_runtime.Value {
	once_negate__1364373265.Do(func() {
		cache_negate__1364373265 = gopurs_runtime.Func(func(dictRing_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_negate__1364373265(gopurs_runtime.CoerceToStruct[pkg_Data_Ring.Constructor_Ring[gopurs_runtime.Value]](dictRing_0_box))
})
	})
	return cache_negate__1364373265
}

var cache_sub__1043827704 gopurs_runtime.Value
var once_sub__1043827704 sync.Once
func Get_sub__1043827704() gopurs_runtime.Value {
	once_sub__1043827704.Do(func() {
		cache_sub__1043827704 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_sub__1043827704(__eta0_0_box, __eta1_1_box)
})
	})
	return cache_sub__1043827704
}

var cache_sub__3675938712 gopurs_runtime.Value
var once_sub__3675938712 sync.Once
func Get_sub__3675938712() gopurs_runtime.Value {
	once_sub__3675938712.Do(func() {
		cache_sub__3675938712 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_sub__3675938712(gopurs_runtime.CoerceToStruct[pkg_Data_Ring.Constructor_Ring[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_sub__3675938712
}

var cache_append__493084344 gopurs_runtime.Value
var once_append__493084344 sync.Once
func Get_append__493084344() gopurs_runtime.Value {
	once_append__493084344.Do(func() {
		cache_append__493084344 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_append__493084344(__eta0_0_box, __eta1_1_box)
})
	})
	return cache_append__493084344
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

var cache_zero__1204848985 gopurs_runtime.Value
var once_zero__1204848985 sync.Once
func Get_zero__1204848985() gopurs_runtime.Value {
	once_zero__1204848985.Do(func() {
		cache_zero__1204848985 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_zero__1204848985(dict_0_box)
})
	})
	return cache_zero__1204848985
}

var cache_show__1488465650 gopurs_runtime.Value
var once_show__1488465650 sync.Once
func Get_show__1488465650() gopurs_runtime.Value {
	once_show__1488465650.Do(func() {
		cache_show__1488465650 = gopurs_runtime.Func(func(__eta0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_show__1488465650(__eta0_0_box)
})
	})
	return cache_show__1488465650
}

var cache_show__2742601362 gopurs_runtime.Value
var once_show__2742601362 sync.Once
func Get_show__2742601362() gopurs_runtime.Value {
	once_show__2742601362.Do(func() {
		cache_show__2742601362 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_show__2742601362(gopurs_runtime.CoerceToStruct[pkg_Data_Show.Constructor_Show[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_show__2742601362
}

var cache_show__2696961938 gopurs_runtime.Value
var once_show__2696961938 sync.Once
func Get_show__2696961938() gopurs_runtime.Value {
	once_show__2696961938.Do(func() {
		cache_show__2696961938 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_show__2696961938(gopurs_runtime.CoerceToStruct[pkg_Data_Show.Constructor_Show[[]*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]](dict_0_box))
})
	})
	return cache_show__2696961938
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

var cache_fst__20422131 gopurs_runtime.Value
var once_fst__20422131 sync.Once
func Get_fst__20422131() gopurs_runtime.Value {
	once_fst__20422131.Do(func() {
		cache_fst__20422131 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_fst__20422131(gopurs_runtime.CoerceToStruct[pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](v_0_box))
})
	})
	return cache_fst__20422131
}

var cache_snd__20422131 gopurs_runtime.Value
var once_snd__20422131 sync.Once
func Get_snd__20422131() gopurs_runtime.Value {
	once_snd__20422131.Do(func() {
		cache_snd__20422131 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_snd__20422131(gopurs_runtime.CoerceToStruct[pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](v_0_box))
})
	})
	return cache_snd__20422131
}

var cache_unfoldableArray__644327338 gopurs_runtime.Value
var once_unfoldableArray__644327338 sync.Once
func Get_unfoldableArray__644327338() gopurs_runtime.Value {
	once_unfoldableArray__644327338.Do(func() {
		cache_unfoldableArray__644327338 = gopurs_runtime.RecordDict2("Unfoldable10", "unfoldr", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Unfoldable1.Get_unfoldable1Array()
}), gopurs_runtime.Apply4(pkg_Data_Unfoldable.Get_unfoldrArrayImpl(), pkg_Data_Maybe.Get_isNothing(), gopurs_runtime.Apply(pkg_Partial_Unsafe.Get__unsafePartial(), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (v_1.Type == 9 && v_1.IntVal == 930809136 && v_1.UnsafePtr != nil) {
__t0 = (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(v_1.UnsafePtr).V0
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return __t0
})
})), pkg_Data_Tuple.Get_fst(), pkg_Data_Tuple.Get_snd()))
	})
	return cache_unfoldableArray__644327338
}

var cache_unfoldr__1128708256 gopurs_runtime.Value
var once_unfoldr__1128708256 sync.Once
func Get_unfoldr__1128708256() gopurs_runtime.Value {
	once_unfoldr__1128708256.Do(func() {
		cache_unfoldr__1128708256 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_unfoldr__1128708256(gopurs_runtime.CoerceToStruct[pkg_Data_Unfoldable.Constructor_Unfoldable[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_unfoldr__1128708256
}

var cache_unfoldr__1519733018 gopurs_runtime.Value
var once_unfoldr__1519733018 sync.Once
func Get_unfoldr__1519733018() gopurs_runtime.Value {
	once_unfoldr__1519733018.Do(func() {
		cache_unfoldr__1519733018 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_unfoldr__1519733018(gopurs_runtime.CoerceToStruct[pkg_Data_Unfoldable.Constructor_Unfoldable[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_unfoldr__1519733018
}

var cache_unfoldable1Array__4196906331 gopurs_runtime.Value
var once_unfoldable1Array__4196906331 sync.Once
func Get_unfoldable1Array__4196906331() gopurs_runtime.Value {
	once_unfoldable1Array__4196906331.Do(func() {
		cache_unfoldable1Array__4196906331 = gopurs_runtime.RecordDict1("unfoldr1", gopurs_runtime.Apply4(pkg_Data_Unfoldable1.Get_unfoldr1ArrayImpl(), pkg_Data_Maybe.Get_isNothing(), gopurs_runtime.Apply(pkg_Partial_Unsafe.Get__unsafePartial(), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (v_1.Type == 9 && v_1.IntVal == 930809136 && v_1.UnsafePtr != nil) {
__t0 = (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(v_1.UnsafePtr).V0
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return __t0
})
})), pkg_Data_Tuple.Get_fst(), pkg_Data_Tuple.Get_snd()))
	})
	return cache_unfoldable1Array__4196906331
}

var cache_unsafePartial__1306634845 gopurs_runtime.Value
var once_unsafePartial__1306634845 sync.Once
func Get_unsafePartial__1306634845() gopurs_runtime.Value {
	once_unsafePartial__1306634845.Do(func() {
		cache_unsafePartial__1306634845 = pkg_Partial_Unsafe.Get__unsafePartial()
	})
	return cache_unsafePartial__1306634845
}

var cache_unsafePartial__1090765981 gopurs_runtime.Value
var once_unsafePartial__1090765981 sync.Once
func Get_unsafePartial__1090765981() gopurs_runtime.Value {
	once_unsafePartial__1090765981.Do(func() {
		cache_unsafePartial__1090765981 = pkg_Partial_Unsafe.Get__unsafePartial()
	})
	return cache_unsafePartial__1090765981
}

type Constructor_Leaf[T_k any, T_v any] struct {
	Rc uint32
}


type Constructor_Node[T_k any, T_v any] struct {
	Rc uint32
	V0 int64
	V1 int64
	V2 T_k
	V3 T_v
	V4 *Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]
	V5 *Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]
}


type Constructor_IterLeaf[T_k any, T_v any] struct {
	Rc uint32
}


type Constructor_IterEmit[T_k any, T_v any] struct {
	Rc uint32
	V0 T_k
	V1 T_v
	V2 gopurs_runtime.Value
}


type Constructor_IterNode[T_k any, T_v any] struct {
	Rc uint32
	V0 *Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]
	V1 gopurs_runtime.Value
}


type Constructor_IterDone[T_k any, T_v any] struct {
	Rc uint32
}


type Constructor_IterNext[T_k any, T_v any] struct {
	Rc uint32
	V0 T_k
	V1 T_v
	V2 gopurs_runtime.Value
}


type Constructor_Split[T_k any, T_v any] struct {
	Rc uint32
	V0 *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]
	V1 *Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]
	V2 *Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]
}


type Constructor_SplitLast[T_k any, T_v any] struct {
	Rc uint32
	V0 T_k
	V1 T_v
	V2 *Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]
}


func Call_identity(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_identity1(x_0_loop *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]) *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value] {
var x_0 *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value] = x_0_loop
_ = x_0
return x_0
}

func Call_identity2(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_unsafeNode(__local_var_0_loop gopurs_runtime.Value, __local_var_1_loop gopurs_runtime.Value, __local_var_2_loop gopurs_runtime.Value, __local_var_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
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
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(&Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]{1, 1, 1, __local_var_0, __local_var_1, gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__local_var_2), gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__local_var_3)})}
goto end_branch_0
} else {

}
}
{
if (__local_var_3.Type == 9 && __local_var_3.IntVal == 324739070 && __local_var_3.UnsafePtr != nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(&Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]{1, Call_add__560788792(gopurs_runtime.Int(1), gopurs_runtime.Int((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V0)).IntVal, Call_add__560788792(gopurs_runtime.Int(1), gopurs_runtime.Int((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V1)).IntVal, __local_var_0, __local_var_1, gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__local_var_2), gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__local_var_3)})}
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__t0))}
goto end_branch_3
} else {

}
}
{
if (__local_var_2.Type == 9 && __local_var_2.IntVal == 324739070 && __local_var_2.UnsafePtr != nil) {
var __t2 gopurs_runtime.Value
{
if (__local_var_3.Type == 9 && __local_var_3.IntVal == 324739070 && __local_var_3.UnsafePtr == nil) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(&Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]{1, Call_add__560788792(gopurs_runtime.Int(1), gopurs_runtime.Int((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V0)).IntVal, Call_add__560788792(gopurs_runtime.Int(1), gopurs_runtime.Int((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V1)).IntVal, __local_var_0, __local_var_1, gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__local_var_2), gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__local_var_3)})}
goto end_branch_2
} else {

}
}
{
if (__local_var_3.Type == 9 && __local_var_3.IntVal == 324739070 && __local_var_3.UnsafePtr != nil) {
var __t1 int64
{
if (gopurs_runtime.Bool(Call_greaterThan__4087042607(gopurs_runtime.Int((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V0), gopurs_runtime.Int((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V0))).IntVal) != (0) {
__t1 = (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V0
goto end_branch_1
} else {

}
}
{
__t1 = (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V0
}
end_branch_1:
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(&Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]{1, Call_add__560788792(gopurs_runtime.Int(1), gopurs_runtime.Int(__t1)).IntVal, Call_add__560788792(gopurs_runtime.Int(Call_add__560788792(gopurs_runtime.Int(1), gopurs_runtime.Int((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V1)).IntVal), gopurs_runtime.Int((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V1)).IntVal, __local_var_0, __local_var_1, gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__local_var_2), gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__local_var_3)})}
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_2:
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__t2))}
goto end_branch_3
} else {

}
}
{
__t3 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_3:
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__t3))}
}

func Call_toMapIter(a_0_loop *Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var a_0 *Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value] = a_0_loop
_ = a_0
return gopurs_runtime.Value{Type: 9, IntVal: 2861335956, UnsafePtr: unsafe.Pointer(&Constructor_IterNode[gopurs_runtime.Value, gopurs_runtime.Value]{1, a_0, gopurs_runtime.Value{Type: 9, IntVal: 2509360378, UnsafePtr: unsafe.Pointer(nil)}})}
}

func Call_stepWith(f_0_loop gopurs_runtime.Value, next_1_loop gopurs_runtime.Value, done_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
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
__t1 = gopurs_runtime.Apply(done_2, pkg_Data_Unit.Get_unit())
goto end_branch_1
} else {

}
}
{
if (v_4.Type == 9 && v_4.IntVal == 1343415489) {
__t1 = gopurs_runtime.UncurriedApp3(next_1, (*Constructor_IterEmit[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V0, (*Constructor_IterEmit[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V1, (*Constructor_IterEmit[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V2)
goto end_branch_1
} else {

}
}
{
if (v_4.Type == 9 && v_4.IntVal == 2861335956) {
v_4_loop = gopurs_runtime.Apply2(f_0, (*Constructor_IterNode[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_IterNode[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V0)})
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

func Call_size(v_0_loop *Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]) int64 {
var v_0 *Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value] = v_0_loop
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
__t0 = gopurs_runtime.Int((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V1)
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

func Call_singleton(k_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) *Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value] {
var k_0 gopurs_runtime.Value = k_0_loop
_ = k_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(&Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]{1, 1, 1, k_0, v_1, gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(nil))}), gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(nil))})})})
}

func Call_unsafeBalancedNode(__local_var_0_loop gopurs_runtime.Value, __local_var_1_loop gopurs_runtime.Value, __local_var_2_loop gopurs_runtime.Value, __local_var_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
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
__t6 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(&Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]{1, 1, 1, __local_var_0, __local_var_1, gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(nil))}), gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(nil))})})}
goto end_branch_6
} else {

}
}
{
if ((__local_var_3.Type == 9 && __local_var_3.IntVal == 324739070 && __local_var_3.UnsafePtr != nil)) && ((gopurs_runtime.Bool(Call_greaterThan__4087042607(gopurs_runtime.Int((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V0), gopurs_runtime.Int(1))).IntVal) != (0)) {
var __t5 *Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]
{
var __t_tag_0 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V4)}
var __t_and_4 bool = false
if (__t_tag_0.Type == 9 && __t_tag_0.IntVal == 324739070 && __t_tag_0.UnsafePtr != nil) {

var __t3 gopurs_runtime.Value
{
var __t_tag_1 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V5)}
if (__t_tag_1.Type == 9 && __t_tag_1.IntVal == 324739070 && __t_tag_1.UnsafePtr == nil) {
__t3 = gopurs_runtime.Int(0)
goto end_branch_3
} else {

}
}
{
var __t_tag_2 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V5)}
if (__t_tag_2.Type == 9 && __t_tag_2.IntVal == 324739070 && __t_tag_2.UnsafePtr != nil) {
__t3 = gopurs_runtime.Int((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V5)}.UnsafePtr).V0)
goto end_branch_3
} else {

}
}
{
__t3 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_3:
__t_and_4 = (gopurs_runtime.Bool(Call_greaterThan__4087042607(gopurs_runtime.Int((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V4)}.UnsafePtr).V0), gopurs_runtime.Int(__t3.IntVal))).IntVal) != (0)
}
if __t_and_4 {
__t5 = gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_unsafeNode(), (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V4)}.UnsafePtr).V2, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V4)}.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_unsafeNode(), __local_var_0, __local_var_1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__local_var_2))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V4)}.UnsafePtr).V4)})))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_unsafeNode(), (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V2, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V4)}.UnsafePtr).V5)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V5)})))}))
goto end_branch_5
} else {

}
}
{
__t5 = gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_unsafeNode(), (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V2, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_unsafeNode(), __local_var_0, __local_var_1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__local_var_2))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V4)})))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V5)}))
}
end_branch_5:
__t6 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(__t5)}
goto end_branch_6
} else {

}
}
{
__t6 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_unsafeNode(), __local_var_0, __local_var_1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__local_var_2))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__local_var_3))})))}
}
end_branch_6:
__t27 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__t6))}
goto end_branch_27
} else {

}
}
{
if (__local_var_2.Type == 9 && __local_var_2.IntVal == 324739070 && __local_var_2.UnsafePtr != nil) {
var __t26 *Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]
{
if (__local_var_3.Type == 9 && __local_var_3.IntVal == 324739070 && __local_var_3.UnsafePtr != nil) {
var __t19 *Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]
{
if (gopurs_runtime.Bool(Call_greaterThan__4087042607(gopurs_runtime.Int((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V0), gopurs_runtime.Int(Call_add__560788792(gopurs_runtime.Int((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V0), gopurs_runtime.Int(1)).IntVal))).IntVal) != (0) {
var __t12 *Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]
{
var __t_tag_7 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V4)}
var __t_and_11 bool = false
if (__t_tag_7.Type == 9 && __t_tag_7.IntVal == 324739070 && __t_tag_7.UnsafePtr != nil) {

var __t10 gopurs_runtime.Value
{
var __t_tag_8 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V5)}
if (__t_tag_8.Type == 9 && __t_tag_8.IntVal == 324739070 && __t_tag_8.UnsafePtr == nil) {
__t10 = gopurs_runtime.Int(0)
goto end_branch_10
} else {

}
}
{
var __t_tag_9 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V5)}
if (__t_tag_9.Type == 9 && __t_tag_9.IntVal == 324739070 && __t_tag_9.UnsafePtr != nil) {
__t10 = gopurs_runtime.Int((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V5)}.UnsafePtr).V0)
goto end_branch_10
} else {

}
}
{
__t10 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_10:
__t_and_11 = (gopurs_runtime.Bool(Call_greaterThan__4087042607(gopurs_runtime.Int((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V4)}.UnsafePtr).V0), gopurs_runtime.Int(__t10.IntVal))).IntVal) != (0)
}
if __t_and_11 {
__t12 = gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_unsafeNode(), (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V4)}.UnsafePtr).V2, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V4)}.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_unsafeNode(), __local_var_0, __local_var_1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__local_var_2))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V4)}.UnsafePtr).V4)})))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_unsafeNode(), (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V2, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V4)}.UnsafePtr).V5)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V5)})))}))
goto end_branch_12
} else {

}
}
{
__t12 = gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_unsafeNode(), (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V2, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_unsafeNode(), __local_var_0, __local_var_1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__local_var_2))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V4)})))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V5)}))
}
end_branch_12:
__t19 = __t12
goto end_branch_19
} else {

}
}
{
if (gopurs_runtime.Bool(Call_greaterThan__4087042607(gopurs_runtime.Int((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V0), gopurs_runtime.Int(Call_add__560788792(gopurs_runtime.Int((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V0), gopurs_runtime.Int(1)).IntVal))).IntVal) != (0) {
var __t18 *Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]
{
var __t_tag_13 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V5)}
var __t_and_17 bool = false
if (__t_tag_13.Type == 9 && __t_tag_13.IntVal == 324739070 && __t_tag_13.UnsafePtr != nil) {

var __t16 gopurs_runtime.Value
{
var __t_tag_14 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V4)}
if (__t_tag_14.Type == 9 && __t_tag_14.IntVal == 324739070 && __t_tag_14.UnsafePtr == nil) {
__t16 = gopurs_runtime.Int(0)
goto end_branch_16
} else {

}
}
{
var __t_tag_15 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V4)}
if (__t_tag_15.Type == 9 && __t_tag_15.IntVal == 324739070 && __t_tag_15.UnsafePtr != nil) {
__t16 = gopurs_runtime.Int((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V4)}.UnsafePtr).V0)
goto end_branch_16
} else {

}
}
{
__t16 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_16:
__t_and_17 = (gopurs_runtime.Bool(Call_lessThanOrEq__4087042607(gopurs_runtime.Int(__t16.IntVal), gopurs_runtime.Int((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V5)}.UnsafePtr).V0))).IntVal) != (0)
}
if __t_and_17 {
__t18 = gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_unsafeNode(), (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V5)}.UnsafePtr).V2, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V5)}.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_unsafeNode(), (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V2, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V4)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V5)}.UnsafePtr).V4)})))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_unsafeNode(), __local_var_0, __local_var_1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V5)}.UnsafePtr).V5)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__local_var_3))})))}))
goto end_branch_18
} else {

}
}
{
__t18 = gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_unsafeNode(), (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V2, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V4)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_unsafeNode(), __local_var_0, __local_var_1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V5)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__local_var_3))})))}))
}
end_branch_18:
__t19 = __t18
goto end_branch_19
} else {

}
}
{
__t19 = gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_unsafeNode(), __local_var_0, __local_var_1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__local_var_2))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__local_var_3))}))
}
end_branch_19:
__t26 = __t19
goto end_branch_26
} else {

}
}
{
if ((__local_var_3.Type == 9 && __local_var_3.IntVal == 324739070 && __local_var_3.UnsafePtr == nil)) && ((gopurs_runtime.Bool(Call_greaterThan__4087042607(gopurs_runtime.Int((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V0), gopurs_runtime.Int(1))).IntVal) != (0)) {
var __t25 *Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]
{
var __t_tag_20 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V5)}
var __t_and_24 bool = false
if (__t_tag_20.Type == 9 && __t_tag_20.IntVal == 324739070 && __t_tag_20.UnsafePtr != nil) {

var __t23 gopurs_runtime.Value
{
var __t_tag_21 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V4)}
if (__t_tag_21.Type == 9 && __t_tag_21.IntVal == 324739070 && __t_tag_21.UnsafePtr == nil) {
__t23 = gopurs_runtime.Int(0)
goto end_branch_23
} else {

}
}
{
var __t_tag_22 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V4)}
if (__t_tag_22.Type == 9 && __t_tag_22.IntVal == 324739070 && __t_tag_22.UnsafePtr != nil) {
__t23 = gopurs_runtime.Int((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V4)}.UnsafePtr).V0)
goto end_branch_23
} else {

}
}
{
__t23 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_23:
__t_and_24 = (gopurs_runtime.Bool(Call_lessThanOrEq__4087042607(gopurs_runtime.Int(__t23.IntVal), gopurs_runtime.Int((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V5)}.UnsafePtr).V0))).IntVal) != (0)
}
if __t_and_24 {
__t25 = gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_unsafeNode(), (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V5)}.UnsafePtr).V2, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V5)}.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_unsafeNode(), (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V2, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V4)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V5)}.UnsafePtr).V4)})))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_unsafeNode(), __local_var_0, __local_var_1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V5)}.UnsafePtr).V5)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__local_var_3))})))}))
goto end_branch_25
} else {

}
}
{
__t25 = gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_unsafeNode(), (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V2, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V4)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_unsafeNode(), __local_var_0, __local_var_1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V5)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__local_var_3))})))}))
}
end_branch_25:
__t26 = __t25
goto end_branch_26
} else {

}
}
{
__t26 = gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_unsafeNode(), __local_var_0, __local_var_1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__local_var_2))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__local_var_3))}))
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
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__t27))}
}

func Call_unsafeSplit(__local_var_0_loop gopurs_runtime.Value, __local_var_1_loop gopurs_runtime.Value, __local_var_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
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
__t4 = gopurs_runtime.Value{Type: 9, IntVal: 3373277644, UnsafePtr: unsafe.Pointer(&Constructor_Split[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(nil))}), gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(nil))}), gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(nil))})})}
goto end_branch_4
} else {

}
}
{
if (__local_var_2.Type == 9 && __local_var_2.IntVal == 324739070 && __local_var_2.UnsafePtr != nil) {
v_3_0 := gopurs_runtime.Apply2(__local_var_0, __local_var_1, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V2)
_ = v_3_0
var __t3 gopurs_runtime.Value
{
if (uint32(v_3_0.IntVal) == 1527465420) {
v1_4_1 := gopurs_runtime.UncurriedApp3(Get_unsafeSplit(), __local_var_0, __local_var_1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V4)})
_ = v1_4_1
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 3373277644, UnsafePtr: unsafe.Pointer(&Constructor_Split[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Split[gopurs_runtime.Value, gopurs_runtime.Value])(v1_4_1.UnsafePtr).V0, (*Constructor_Split[gopurs_runtime.Value, gopurs_runtime.Value])(v1_4_1.UnsafePtr).V1, gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_unsafeBalancedNode(), (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V2, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Split[gopurs_runtime.Value, gopurs_runtime.Value])(v1_4_1.UnsafePtr).V2)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V5)}))})}
goto end_branch_3
} else {

}
}
{
if (uint32(v_3_0.IntVal) == 380165415) {
v1_4_2 := gopurs_runtime.UncurriedApp3(Get_unsafeSplit(), __local_var_0, __local_var_1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V5)})
_ = v1_4_2
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 3373277644, UnsafePtr: unsafe.Pointer(&Constructor_Split[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Split[gopurs_runtime.Value, gopurs_runtime.Value])(v1_4_2.UnsafePtr).V0, gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_unsafeBalancedNode(), (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V2, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V4)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Split[gopurs_runtime.Value, gopurs_runtime.Value])(v1_4_2.UnsafePtr).V1)})), (*Constructor_Split[gopurs_runtime.Value, gopurs_runtime.Value])(v1_4_2.UnsafePtr).V2})}
goto end_branch_3
} else {

}
}
{
if (uint32(v_3_0.IntVal) == 902936544) {
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 3373277644, UnsafePtr: unsafe.Pointer(&Constructor_Split[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V3})}), (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V4, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V5})}
goto end_branch_3
} else {

}
}
{
__t3 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_3:
__t4 = gopurs_runtime.Value{Type: 9, IntVal: 3373277644, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Split[gopurs_runtime.Value, gopurs_runtime.Value]](__t3))}
goto end_branch_4
} else {

}
}
{
__t4 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_4:
return gopurs_runtime.Value{Type: 9, IntVal: 3373277644, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Split[gopurs_runtime.Value, gopurs_runtime.Value]](__t4))}
}
}

func Call_unsafeSplitLast(__local_var_0_loop gopurs_runtime.Value, __local_var_1_loop gopurs_runtime.Value, __local_var_2_loop gopurs_runtime.Value, __local_var_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
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
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 2668112006, UnsafePtr: unsafe.Pointer(&Constructor_SplitLast[gopurs_runtime.Value, gopurs_runtime.Value]{1, __local_var_0, __local_var_1, gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__local_var_2)})}
goto end_branch_1
} else {

}
}
{
if (__local_var_3.Type == 9 && __local_var_3.IntVal == 324739070 && __local_var_3.UnsafePtr != nil) {
v1_4_0 := gopurs_runtime.UncurriedApp4(Get_unsafeSplitLast(), (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V2, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V4)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V5)})
_ = v1_4_0
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 2668112006, UnsafePtr: unsafe.Pointer(&Constructor_SplitLast[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_SplitLast[gopurs_runtime.Value, gopurs_runtime.Value])(v1_4_0.UnsafePtr).V0, (*Constructor_SplitLast[gopurs_runtime.Value, gopurs_runtime.Value])(v1_4_0.UnsafePtr).V1, gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_unsafeBalancedNode(), __local_var_0, __local_var_1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__local_var_2))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_SplitLast[gopurs_runtime.Value, gopurs_runtime.Value])(v1_4_0.UnsafePtr).V2)}))})}
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
return gopurs_runtime.Value{Type: 9, IntVal: 2668112006, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_SplitLast[gopurs_runtime.Value, gopurs_runtime.Value]](__t1))}
}
}

func Call_unsafeJoinNodes(__local_var_0_loop gopurs_runtime.Value, __local_var_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __local_var_0 gopurs_runtime.Value = __local_var_0_loop
_ = __local_var_0
var __local_var_1 gopurs_runtime.Value = __local_var_1_loop
_ = __local_var_1
var __t1 gopurs_runtime.Value
{
if (__local_var_0.Type == 9 && __local_var_0.IntVal == 324739070 && __local_var_0.UnsafePtr == nil) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__local_var_1))}
goto end_branch_1
} else {

}
}
{
if (__local_var_0.Type == 9 && __local_var_0.IntVal == 324739070 && __local_var_0.UnsafePtr != nil) {
v2_2_0 := gopurs_runtime.UncurriedApp4(Get_unsafeSplitLast(), (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_0.UnsafePtr).V2, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_0.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_0.UnsafePtr).V4)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_0.UnsafePtr).V5)})
_ = v2_2_0
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_unsafeBalancedNode(), (*Constructor_SplitLast[gopurs_runtime.Value, gopurs_runtime.Value])(v2_2_0.UnsafePtr).V0, (*Constructor_SplitLast[gopurs_runtime.Value, gopurs_runtime.Value])(v2_2_0.UnsafePtr).V1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_SplitLast[gopurs_runtime.Value, gopurs_runtime.Value])(v2_2_0.UnsafePtr).V2)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__local_var_1))})))}
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__t1))}
}

func Call_unsafeDifference(__local_var_0_loop gopurs_runtime.Value, __local_var_1_loop gopurs_runtime.Value, __local_var_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
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
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(nil))}
goto end_branch_1
} else {

}
}
{
if (__local_var_2.Type == 9 && __local_var_2.IntVal == 324739070 && __local_var_2.UnsafePtr == nil) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__local_var_1))}
goto end_branch_1
} else {

}
}
{
if (__local_var_2.Type == 9 && __local_var_2.IntVal == 324739070 && __local_var_2.UnsafePtr != nil) {
v_3_0 := gopurs_runtime.UncurriedApp3(Get_unsafeSplit(), __local_var_0, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V2, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__local_var_1))})
_ = v_3_0
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp2(Get_unsafeJoinNodes(), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp3(Get_unsafeDifference(), __local_var_0, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Split[gopurs_runtime.Value, gopurs_runtime.Value])(v_3_0.UnsafePtr).V1)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V4)})))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp3(Get_unsafeDifference(), __local_var_0, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Split[gopurs_runtime.Value, gopurs_runtime.Value])(v_3_0.UnsafePtr).V2)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V5)})))})))}
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__t1))}
}
}

func Call_unsafeIntersectionWith(__local_var_0_loop gopurs_runtime.Value, __local_var_1_loop gopurs_runtime.Value, __local_var_2_loop gopurs_runtime.Value, __local_var_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
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
__t6 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(nil))}
goto end_branch_6
} else {

}
}
{
if (__local_var_3.Type == 9 && __local_var_3.IntVal == 324739070 && __local_var_3.UnsafePtr == nil) {
__t6 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(nil))}
goto end_branch_6
} else {

}
}
{
if (__local_var_3.Type == 9 && __local_var_3.IntVal == 324739070 && __local_var_3.UnsafePtr != nil) {
v_4_0 := gopurs_runtime.UncurriedApp3(Get_unsafeSplit(), __local_var_0, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V2, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__local_var_2))})
_ = v_4_0
l_prime_5_1 := gopurs_runtime.UncurriedApp4(Get_unsafeIntersectionWith(), __local_var_0, __local_var_1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Split[gopurs_runtime.Value, gopurs_runtime.Value])(v_4_0.UnsafePtr).V1)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V4)})
_ = l_prime_5_1
r_prime_6_2 := gopurs_runtime.UncurriedApp4(Get_unsafeIntersectionWith(), __local_var_0, __local_var_1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Split[gopurs_runtime.Value, gopurs_runtime.Value])(v_4_0.UnsafePtr).V2)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V5)})
_ = r_prime_6_2
var __t5 gopurs_runtime.Value
{
var __t_tag_3 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Split[gopurs_runtime.Value, gopurs_runtime.Value])(v_4_0.UnsafePtr).V0)}
if (__t_tag_3.Type == 9 && __t_tag_3.IntVal == 930809136 && __t_tag_3.UnsafePtr != nil) {
__t5 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_unsafeBalancedNode(), (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V2, gopurs_runtime.Apply2(__local_var_1, (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Split[gopurs_runtime.Value, gopurs_runtime.Value])(v_4_0.UnsafePtr).V0)}.UnsafePtr).V0, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V3), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](l_prime_5_1))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](r_prime_6_2))})))}
goto end_branch_5
} else {

}
}
{
var __t_tag_4 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Split[gopurs_runtime.Value, gopurs_runtime.Value])(v_4_0.UnsafePtr).V0)}
if (__t_tag_4.Type == 9 && __t_tag_4.IntVal == 930809136 && __t_tag_4.UnsafePtr == nil) {
__t5 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp2(Get_unsafeJoinNodes(), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](l_prime_5_1))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](r_prime_6_2))})))}
goto end_branch_5
} else {

}
}
{
__t5 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_5:
__t6 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__t5))}
goto end_branch_6
} else {

}
}
{
__t6 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_6:
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__t6))}
}
}

func Call_unsafeUnionWith(__local_var_0_loop gopurs_runtime.Value, __local_var_1_loop gopurs_runtime.Value, __local_var_2_loop gopurs_runtime.Value, __local_var_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
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
__t6 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__local_var_3))}
goto end_branch_6
} else {

}
}
{
if (__local_var_3.Type == 9 && __local_var_3.IntVal == 324739070 && __local_var_3.UnsafePtr == nil) {
__t6 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__local_var_2))}
goto end_branch_6
} else {

}
}
{
if (__local_var_3.Type == 9 && __local_var_3.IntVal == 324739070 && __local_var_3.UnsafePtr != nil) {
v_4_0 := gopurs_runtime.UncurriedApp3(Get_unsafeSplit(), __local_var_0, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V2, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__local_var_2))})
_ = v_4_0
l_prime_5_1 := gopurs_runtime.UncurriedApp4(Get_unsafeUnionWith(), __local_var_0, __local_var_1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Split[gopurs_runtime.Value, gopurs_runtime.Value])(v_4_0.UnsafePtr).V1)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V4)})
_ = l_prime_5_1
r_prime_6_2 := gopurs_runtime.UncurriedApp4(Get_unsafeUnionWith(), __local_var_0, __local_var_1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Split[gopurs_runtime.Value, gopurs_runtime.Value])(v_4_0.UnsafePtr).V2)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V5)})
_ = r_prime_6_2
var __t5 gopurs_runtime.Value
{
var __t_tag_3 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Split[gopurs_runtime.Value, gopurs_runtime.Value])(v_4_0.UnsafePtr).V0)}
if (__t_tag_3.Type == 9 && __t_tag_3.IntVal == 930809136 && __t_tag_3.UnsafePtr != nil) {
__t5 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_unsafeBalancedNode(), (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V2, gopurs_runtime.Apply2(__local_var_1, (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Split[gopurs_runtime.Value, gopurs_runtime.Value])(v_4_0.UnsafePtr).V0)}.UnsafePtr).V0, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V3), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](l_prime_5_1))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](r_prime_6_2))})))}
goto end_branch_5
} else {

}
}
{
var __t_tag_4 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Split[gopurs_runtime.Value, gopurs_runtime.Value])(v_4_0.UnsafePtr).V0)}
if (__t_tag_4.Type == 9 && __t_tag_4.IntVal == 930809136 && __t_tag_4.UnsafePtr == nil) {
__t5 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_unsafeBalancedNode(), (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V2, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](l_prime_5_1))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](r_prime_6_2))})))}
goto end_branch_5
} else {

}
}
{
__t5 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_5:
__t6 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__t5))}
goto end_branch_6
} else {

}
}
{
__t6 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_6:
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__t6))}
}
}

func Call_unionWith(dictOrd_0_loop *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictOrd_0 *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value] = dictOrd_0_loop
_ = dictOrd_0
compare_1_0 := dictOrd_0.V1
_ = compare_1_0
return gopurs_runtime.Func(func(app_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m2_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_unsafeUnionWith(), compare_1_0, app_2, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](m1_3))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](m2_4))})))}
})
})
})
}

func Call_union(dictOrd_0_loop *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictOrd_0 *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value] = dictOrd_0_loop
_ = dictOrd_0
compare_1_0 := dictOrd_0.V1
_ = compare_1_0
return gopurs_runtime.Func(func(m1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m2_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_unsafeUnionWith(), compare_1_0, pkg_Data_Function.Get_go__const(), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](m1_2))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](m2_3))})))}
})
})
}

func Call_update(dictOrd_0_loop *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value], f_1_loop gopurs_runtime.Value, k_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value] = dictOrd_0_loop
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
__t5 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(nil))}
goto end_branch_5
} else {

}
}
{
if (v_4.Type == 9 && v_4.IntVal == 324739070 && v_4.UnsafePtr != nil) {
v1_5_1 := gopurs_runtime.Apply2(dictOrd_0.V1, k_2, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V2)
_ = v1_5_1
var __t4 gopurs_runtime.Value
{
if (uint32(v1_5_1.IntVal) == 1527465420) {
__t4 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_unsafeBalancedNode(), (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V2, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(go__go_3_0_1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V4)})))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V5)})))}
goto end_branch_4
} else {

}
}
{
if (uint32(v1_5_1.IntVal) == 380165415) {
__t4 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_unsafeBalancedNode(), (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V2, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V4)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(go__go_3_0_1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V5)})))})))}
goto end_branch_4
} else {

}
}
{
if (uint32(v1_5_1.IntVal) == 902936544) {
v2_6_2 := gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(f_1, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V3))
_ = v2_6_2
var __t3 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_6_2)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_6_2)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_6_2)}.UnsafePtr == nil) {
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp2(Get_unsafeJoinNodes(), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V4)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V5)})))}
goto end_branch_3
} else {

}
}
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_6_2)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_6_2)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_6_2)}.UnsafePtr != nil) {
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(&Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V0, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V1, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V2, (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_6_2)}.UnsafePtr).V0, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V4, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V5})}
goto end_branch_3
} else {

}
}
{
__t3 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_3:
__t4 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__t3))}
goto end_branch_4
} else {

}
}
{
__t4 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_4:
__t5 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__t4))}
goto end_branch_5
} else {

}
}
{
__t5 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_5:
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__t5))}
})
return go__go_3_0_1
}

func Call_showTree(dictShow_0_loop *pkg_Data_Show.Constructor_Show[gopurs_runtime.Value], dictShow1_1_loop *pkg_Data_Show.Constructor_Show[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictShow_0 *pkg_Data_Show.Constructor_Show[gopurs_runtime.Value] = dictShow_0_loop
_ = dictShow_0
var dictShow1_1 *pkg_Data_Show.Constructor_Show[gopurs_runtime.Value] = dictShow1_1_loop
_ = dictShow1_1
var go__go_2_0_2 gopurs_runtime.Value
_ = go__go_2_0_2
go__go_2_0_2 = gopurs_runtime.Func(func(ind_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if (v_4.Type == 9 && v_4.IntVal == 324739070 && v_4.UnsafePtr == nil) {
__t1 = gopurs_runtime.Str(Call_append__493084344(gopurs_runtime.Str(ind_3.StrVal()), gopurs_runtime.Str("Leaf")).StrVal())
goto end_branch_1
} else {

}
}
{
if (v_4.Type == 9 && v_4.IntVal == 324739070 && v_4.UnsafePtr != nil) {
__t1 = gopurs_runtime.Str(Call_append__493084344(gopurs_runtime.Str(Call_append__493084344(gopurs_runtime.Str(ind_3.StrVal()), gopurs_runtime.Str(Call_append__493084344(gopurs_runtime.Str("["), gopurs_runtime.Str(Call_append__493084344(gopurs_runtime.Str(Call_show__1488465650(gopurs_runtime.Int((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V0)).StrVal()), gopurs_runtime.Str(Call_append__493084344(gopurs_runtime.Str("] "), gopurs_runtime.Str(Call_append__493084344(gopurs_runtime.Str(gopurs_runtime.Apply(dictShow_0.V0, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V2).StrVal()), gopurs_runtime.Str(Call_append__493084344(gopurs_runtime.Str(" => "), gopurs_runtime.Str(Call_append__493084344(gopurs_runtime.Str(gopurs_runtime.Apply(dictShow1_1.V0, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V3).StrVal()), gopurs_runtime.Str("\x0a")).StrVal())).StrVal())).StrVal())).StrVal())).StrVal())).StrVal())).StrVal()), gopurs_runtime.Str(Call_append__493084344(gopurs_runtime.Str(Call_append__493084344(gopurs_runtime.Str(gopurs_runtime.Apply2(go__go_2_0_2, gopurs_runtime.Str(Call_append__493084344(gopurs_runtime.Str(ind_3.StrVal()), gopurs_runtime.Str("    ")).StrVal()), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V4)}).StrVal()), gopurs_runtime.Str("\x0a")).StrVal()), gopurs_runtime.Str(gopurs_runtime.Apply2(go__go_2_0_2, gopurs_runtime.Str(Call_append__493084344(gopurs_runtime.Str(ind_3.StrVal()), gopurs_runtime.Str("    ")).StrVal()), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V5)}).StrVal())).StrVal())).StrVal())
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

func Call_semigroupMap(_dollar__unused_0_loop gopurs_runtime.Value, dictOrd_1_loop gopurs_runtime.Value, dictSemigroup_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var _dollar__unused_0 gopurs_runtime.Value = _dollar__unused_0_loop
_ = _dollar__unused_0
var dictOrd_1 gopurs_runtime.Value = dictOrd_1_loop
_ = dictOrd_1
var dictSemigroup_2 gopurs_runtime.Value = dictSemigroup_2_loop
_ = dictSemigroup_2
compare_3_0 := gopurs_runtime.RecordGet(dictOrd_1, "compare")
_ = compare_3_0
__local_var_4_1 := gopurs_runtime.RecordGet(dictSemigroup_2, "append")
_ = __local_var_4_1
return gopurs_runtime.RecordDict1("append", gopurs_runtime.Func(func(m1_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m2_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_unsafeUnionWith(), compare_3_0, __local_var_4_1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](m1_5))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](m2_6))})))}
})
}))
}

func Call_pop(dictOrd_0_loop *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictOrd_0 *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value] = dictOrd_0_loop
_ = dictOrd_0
compare_1_0 := dictOrd_0.V1
_ = compare_1_0
return gopurs_runtime.Func(func(k_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m_3 gopurs_runtime.Value) gopurs_runtime.Value {
v_4_1 := gopurs_runtime.UncurriedApp3(Get_unsafeSplit(), compare_1_0, k_2, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](m_3))})
_ = v_4_1
__local_var_5_2 := (*Constructor_Split[gopurs_runtime.Value, gopurs_runtime.Value])(v_4_1.UnsafePtr).V1
_ = __local_var_5_2
__local_var_6_3 := (*Constructor_Split[gopurs_runtime.Value, gopurs_runtime.Value])(v_4_1.UnsafePtr).V2
_ = __local_var_6_3
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, *Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]]]](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_functorMaybe(), "map"), gopurs_runtime.Func(func(a_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, a_7, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp2(Get_unsafeJoinNodes(), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(__local_var_5_2)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(__local_var_6_3)})))}})}
}), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Split[gopurs_runtime.Value, gopurs_runtime.Value])(v_4_1.UnsafePtr).V0)})))}
})
})
}

func Call_member(dictOrd_0_loop *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value], k_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value] = dictOrd_0_loop
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
v1_4_1 := gopurs_runtime.Apply2(dictOrd_0.V1, k_1, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V2)
_ = v1_4_1
var __t2 gopurs_runtime.Value
{
if (uint32(v1_4_1.IntVal) == 1527465420) {
v_3_loop = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V4)}
continue go__go_2_0_3
__t2 = gopurs_runtime.Bool((gopurs_runtime.Value{}.IntVal) != (0))
goto end_branch_2
} else {

}
}
{
if (uint32(v1_4_1.IntVal) == 380165415) {
v_3_loop = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V5)}
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

func Call_mapMaybeWithKey(dictOrd_0_loop *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value], f_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value] = dictOrd_0_loop
_ = dictOrd_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var go__go_2_0_4 gopurs_runtime.Value
_ = go__go_2_0_4
go__go_2_0_4 = gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t3 gopurs_runtime.Value
{
if (v_3.Type == 9 && v_3.IntVal == 324739070 && v_3.UnsafePtr == nil) {
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(nil))}
goto end_branch_3
} else {

}
}
{
if (v_3.Type == 9 && v_3.IntVal == 324739070 && v_3.UnsafePtr != nil) {
v2_4_1 := gopurs_runtime.Apply2(f_1, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V2, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V3)
_ = v2_4_1
var __t2 gopurs_runtime.Value
{
if (v2_4_1.Type == 9 && v2_4_1.IntVal == 930809136 && v2_4_1.UnsafePtr != nil) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_unsafeBalancedNode(), (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V2, (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(v2_4_1.UnsafePtr).V0, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(go__go_2_0_4, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V4)})))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(go__go_2_0_4, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V5)})))})))}
goto end_branch_2
} else {

}
}
{
if (v2_4_1.Type == 9 && v2_4_1.IntVal == 930809136 && v2_4_1.UnsafePtr == nil) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp2(Get_unsafeJoinNodes(), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(go__go_2_0_4, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V4)})))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(go__go_2_0_4, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V5)})))})))}
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_2:
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__t2))}
goto end_branch_3
} else {

}
}
{
__t3 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_3:
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__t3))}
})
return go__go_2_0_4
}

func Call_mapMaybe(dictOrd_0_loop *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value], x_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value] = dictOrd_0_loop
_ = dictOrd_0
var x_1 gopurs_runtime.Value = x_1_loop
_ = x_1
return Call_mapMaybeWithKey(dictOrd_0, gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return x_1
}))
}

func Call_lookupLE(dictOrd_0_loop *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value], k_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value] = dictOrd_0_loop
_ = dictOrd_0
var k_1 gopurs_runtime.Value = k_1_loop
_ = k_1
var go__go_2_0_5 gopurs_runtime.Value
_ = go__go_2_0_5
go__go_2_0_5 = gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t5 gopurs_runtime.Value
{
if (v_3.Type == 9 && v_3.IntVal == 324739070 && v_3.UnsafePtr == nil) {
__t5 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(nil))}
goto end_branch_5
} else {

}
}
{
if (v_3.Type == 9 && v_3.IntVal == 324739070 && v_3.UnsafePtr != nil) {
v1_4_1 := gopurs_runtime.Apply2(dictOrd_0.V1, k_1, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V2)
_ = v1_4_1
var __t4 gopurs_runtime.Value
{
if (uint32(v1_4_1.IntVal) == 1527465420) {
__t4 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(go__go_2_0_5, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V4)})))}
goto end_branch_4
} else {

}
}
{
if (uint32(v1_4_1.IntVal) == 380165415) {
v2_5_2 := gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(go__go_2_0_5, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V5)}))
_ = v2_5_2
var __t3 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_5_2)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_5_2)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_5_2)}.UnsafePtr == nil) {
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, gopurs_runtime.RecordDict2("key", "value", (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V2, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V3)})}
goto end_branch_3
} else {

}
}
{
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_5_2)}
}
end_branch_3:
__t4 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](__t3))}
goto end_branch_4
} else {

}
}
{
if (uint32(v1_4_1.IntVal) == 902936544) {
__t4 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, gopurs_runtime.RecordDict2("key", "value", (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V2, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V3)})}
goto end_branch_4
} else {

}
}
{
__t4 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_4:
__t5 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](__t4))}
goto end_branch_5
} else {

}
}
{
__t5 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_5:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](__t5))}
})
return go__go_2_0_5
}

func Call_lookupGE(dictOrd_0_loop *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value], k_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value] = dictOrd_0_loop
_ = dictOrd_0
var k_1 gopurs_runtime.Value = k_1_loop
_ = k_1
var go__go_2_0_6 gopurs_runtime.Value
_ = go__go_2_0_6
go__go_2_0_6 = gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t5 gopurs_runtime.Value
{
if (v_3.Type == 9 && v_3.IntVal == 324739070 && v_3.UnsafePtr == nil) {
__t5 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(nil))}
goto end_branch_5
} else {

}
}
{
if (v_3.Type == 9 && v_3.IntVal == 324739070 && v_3.UnsafePtr != nil) {
v1_4_1 := gopurs_runtime.Apply2(dictOrd_0.V1, k_1, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V2)
_ = v1_4_1
var __t4 gopurs_runtime.Value
{
if (uint32(v1_4_1.IntVal) == 1527465420) {
v2_5_2 := gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(go__go_2_0_6, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V4)}))
_ = v2_5_2
var __t3 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_5_2)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_5_2)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_5_2)}.UnsafePtr == nil) {
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, gopurs_runtime.RecordDict2("key", "value", (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V2, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V3)})}
goto end_branch_3
} else {

}
}
{
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_5_2)}
}
end_branch_3:
__t4 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](__t3))}
goto end_branch_4
} else {

}
}
{
if (uint32(v1_4_1.IntVal) == 380165415) {
__t4 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(go__go_2_0_6, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V5)})))}
goto end_branch_4
} else {

}
}
{
if (uint32(v1_4_1.IntVal) == 902936544) {
__t4 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, gopurs_runtime.RecordDict2("key", "value", (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V2, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V3)})}
goto end_branch_4
} else {

}
}
{
__t4 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_4:
__t5 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](__t4))}
goto end_branch_5
} else {

}
}
{
__t5 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_5:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](__t5))}
})
return go__go_2_0_6
}

func Call_lookup(dictOrd_0_loop *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value], k_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value] = dictOrd_0_loop
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
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(nil))}
goto end_branch_3
} else {

}
}
{
if (v_3.Type == 9 && v_3.IntVal == 324739070 && v_3.UnsafePtr != nil) {
v1_4_1 := gopurs_runtime.Apply2(dictOrd_0.V1, k_1, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V2)
_ = v1_4_1
var __t2 gopurs_runtime.Value
{
if (uint32(v1_4_1.IntVal) == 1527465420) {
v_3_loop = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V4)}
continue go__go_2_0_7
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](gopurs_runtime.Value{}))}
goto end_branch_2
} else {

}
}
{
if (uint32(v1_4_1.IntVal) == 380165415) {
v_3_loop = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V5)}
continue go__go_2_0_7
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](gopurs_runtime.Value{}))}
goto end_branch_2
} else {

}
}
{
if (uint32(v1_4_1.IntVal) == 902936544) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V3})}
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_2:
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](__t2))}
goto end_branch_3
} else {

}
}
{
__t3 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_3:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](__t3))}
}
}()
})
return go__go_2_0_7
}

func Call_iterMapU(iter_0_loop gopurs_runtime.Value, v_1_loop *Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var iter_0 gopurs_runtime.Value = iter_0_loop
_ = iter_0
var v_1 *Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value] = v_1_loop
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
var __t5 *Constructor_IterEmit[gopurs_runtime.Value, gopurs_runtime.Value]
{
var __t_tag_2 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_1)}.UnsafePtr).V4)}
if (__t_tag_2.Type == 9 && __t_tag_2.IntVal == 324739070 && __t_tag_2.UnsafePtr == nil) {
var __t4 *Constructor_IterEmit[gopurs_runtime.Value, gopurs_runtime.Value]
{
var __t_tag_3 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_1)}.UnsafePtr).V5)}
if (__t_tag_3.Type == 9 && __t_tag_3.IntVal == 324739070 && __t_tag_3.UnsafePtr == nil) {
__t4 = &Constructor_IterEmit[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_1)}.UnsafePtr).V2, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_1)}.UnsafePtr).V3, iter_0}
goto end_branch_4
} else {

}
}
{
__t4 = &Constructor_IterEmit[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_1)}.UnsafePtr).V2, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_1)}.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 2861335956, UnsafePtr: unsafe.Pointer(&Constructor_IterNode[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_1)}.UnsafePtr).V5, iter_0})}}
}
end_branch_4:
__t5 = __t4
goto end_branch_5
} else {

}
}
{
var __t1 *Constructor_IterEmit[gopurs_runtime.Value, gopurs_runtime.Value]
{
var __t_tag_0 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_1)}.UnsafePtr).V5)}
if (__t_tag_0.Type == 9 && __t_tag_0.IntVal == 324739070 && __t_tag_0.UnsafePtr == nil) {
__t1 = &Constructor_IterEmit[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_1)}.UnsafePtr).V2, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_1)}.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 2861335956, UnsafePtr: unsafe.Pointer(&Constructor_IterNode[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_1)}.UnsafePtr).V4, iter_0})}}
goto end_branch_1
} else {

}
}
{
__t1 = &Constructor_IterEmit[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_1)}.UnsafePtr).V2, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_1)}.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 2861335956, UnsafePtr: unsafe.Pointer(&Constructor_IterNode[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_1)}.UnsafePtr).V4, gopurs_runtime.Value{Type: 9, IntVal: 2861335956, UnsafePtr: unsafe.Pointer(&Constructor_IterNode[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_1)}.UnsafePtr).V5, iter_0})}})}}
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

func Call_toUnfoldableUnordered(dictUnfoldable_0_loop *pkg_Data_Unfoldable.Constructor_Unfoldable[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictUnfoldable_0 *pkg_Data_Unfoldable.Constructor_Unfoldable[gopurs_runtime.Value] = dictUnfoldable_0_loop
_ = dictUnfoldable_0
__local_var_1_0 := gopurs_runtime.Apply(dictUnfoldable_0.V1, Get_stepUnfoldrUnordered())
_ = __local_var_1_0
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_1_0, gopurs_runtime.Value{Type: 9, IntVal: 2861335956, UnsafePtr: unsafe.Pointer(&Constructor_IterNode[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](x_2), gopurs_runtime.Value{Type: 9, IntVal: 2509360378, UnsafePtr: unsafe.Pointer(nil)}})})
})
}

func Call_eqMapIter(dictEq_0_loop gopurs_runtime.Value, dictEq1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
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
v_5_1 := gopurs_runtime.CoerceToStruct[Constructor_IterNext[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(Get_stepAsc(), a_3))
_ = v_5_1
var __t4 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 953589075, UnsafePtr: unsafe.Pointer(v_5_1)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 953589075, UnsafePtr: unsafe.Pointer(v_5_1)}.IntVal == 953589075 && gopurs_runtime.Value{Type: 9, IntVal: 953589075, UnsafePtr: unsafe.Pointer(v_5_1)}.UnsafePtr != nil) {
v2_6_2 := gopurs_runtime.CoerceToStruct[Constructor_IterNext[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(Get_stepAsc(), b_4))
_ = v2_6_2
var __t3 bool
{
if ((gopurs_runtime.Value{Type: 9, IntVal: 953589075, UnsafePtr: unsafe.Pointer(v2_6_2)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 953589075, UnsafePtr: unsafe.Pointer(v2_6_2)}.IntVal == 953589075 && gopurs_runtime.Value{Type: 9, IntVal: 953589075, UnsafePtr: unsafe.Pointer(v2_6_2)}.UnsafePtr != nil)) && ((Call_conj__3676519832(gopurs_runtime.Bool((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq_0, "eq"), (*Constructor_IterNext[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 953589075, UnsafePtr: unsafe.Pointer(v_5_1)}.UnsafePtr).V0, (*Constructor_IterNext[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 953589075, UnsafePtr: unsafe.Pointer(v2_6_2)}.UnsafePtr).V0).IntVal) != (0)), gopurs_runtime.Bool((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq1_1, "eq"), (*Constructor_IterNext[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 953589075, UnsafePtr: unsafe.Pointer(v_5_1)}.UnsafePtr).V1, (*Constructor_IterNext[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 953589075, UnsafePtr: unsafe.Pointer(v2_6_2)}.UnsafePtr).V1).IntVal) != (0))).IntVal) != (0)) {
a_3_loop = (*Constructor_IterNext[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 953589075, UnsafePtr: unsafe.Pointer(v_5_1)}.UnsafePtr).V2
b_4_loop = (*Constructor_IterNext[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 953589075, UnsafePtr: unsafe.Pointer(v2_6_2)}.UnsafePtr).V2
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

func Call_ordMapIter(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
eqMapIter1_1_0 := gopurs_runtime.Apply(Get_eqMapIter(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd_0, "Eq0"), gopurs_runtime.Value{}))
_ = eqMapIter1_1_0
return gopurs_runtime.Func(func(dictOrd1_2 gopurs_runtime.Value) gopurs_runtime.Value {
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
v_7_3 := gopurs_runtime.CoerceToStruct[Constructor_IterNext[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(Get_stepAsc(), b_6))
_ = v_7_3
v1_8_4 := gopurs_runtime.CoerceToStruct[Constructor_IterNext[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(Get_stepAsc(), a_5))
_ = v1_8_4
var __t11 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 953589075, UnsafePtr: unsafe.Pointer(v1_8_4)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 953589075, UnsafePtr: unsafe.Pointer(v1_8_4)}.IntVal == 953589075 && gopurs_runtime.Value{Type: 9, IntVal: 953589075, UnsafePtr: unsafe.Pointer(v1_8_4)}.UnsafePtr != nil) {
var __t9 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 953589075, UnsafePtr: unsafe.Pointer(v_7_3)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 953589075, UnsafePtr: unsafe.Pointer(v_7_3)}.IntVal == 953589075 && gopurs_runtime.Value{Type: 9, IntVal: 953589075, UnsafePtr: unsafe.Pointer(v_7_3)}.UnsafePtr != nil) {
v3_9_5 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictOrd_0, "compare"), (*Constructor_IterNext[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 953589075, UnsafePtr: unsafe.Pointer(v1_8_4)}.UnsafePtr).V0, (*Constructor_IterNext[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 953589075, UnsafePtr: unsafe.Pointer(v_7_3)}.UnsafePtr).V0)
_ = v3_9_5
var __t8 uint32
{
if (uint32(v3_9_5.IntVal) == 902936544) {
v4_10_6 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictOrd1_2, "compare"), (*Constructor_IterNext[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 953589075, UnsafePtr: unsafe.Pointer(v1_8_4)}.UnsafePtr).V1, (*Constructor_IterNext[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 953589075, UnsafePtr: unsafe.Pointer(v_7_3)}.UnsafePtr).V1)
_ = v4_10_6
var __t7 uint32
{
if (uint32(v4_10_6.IntVal) == 902936544) {
a_5_loop = (*Constructor_IterNext[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 953589075, UnsafePtr: unsafe.Pointer(v1_8_4)}.UnsafePtr).V2
b_6_loop = (*Constructor_IterNext[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 953589075, UnsafePtr: unsafe.Pointer(v_7_3)}.UnsafePtr).V2
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

func Call_toUnfoldable(dictUnfoldable_0_loop *pkg_Data_Unfoldable.Constructor_Unfoldable[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictUnfoldable_0 *pkg_Data_Unfoldable.Constructor_Unfoldable[gopurs_runtime.Value] = dictUnfoldable_0_loop
_ = dictUnfoldable_0
__local_var_1_0 := gopurs_runtime.Apply(dictUnfoldable_0.V1, Get_stepUnfoldr())
_ = __local_var_1_0
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_1_0, gopurs_runtime.Value{Type: 9, IntVal: 2861335956, UnsafePtr: unsafe.Pointer(&Constructor_IterNode[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](x_2), gopurs_runtime.Value{Type: 9, IntVal: 2509360378, UnsafePtr: unsafe.Pointer(nil)}})})
})
}

func Call_showMap(dictShow_0_loop gopurs_runtime.Value, dictShow1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictShow_0 gopurs_runtime.Value = dictShow_0_loop
_ = dictShow_0
var dictShow1_1 gopurs_runtime.Value = dictShow1_1_loop
_ = dictShow1_1
showArray_2_0 := &pkg_Data_Show.Constructor_Show[[]*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]{1, gopurs_runtime.Apply(pkg_Data_Show.Get_showArrayImpl(), gopurs_runtime.RecordGet(gopurs_runtime.Apply2(pkg_Data_Tuple.Get_showTuple(), dictShow_0, dictShow1_1), "show"))}
_ = showArray_2_0
return gopurs_runtime.RecordDict1("show", gopurs_runtime.Func(func(as_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str(Call_append__493084344(gopurs_runtime.Str("(fromFoldable "), gopurs_runtime.Str(Call_append__493084344(gopurs_runtime.Str(gopurs_runtime.Apply(showArray_2_0.V0, func() gopurs_runtime.Value {
					arr := func() []*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value] {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Unfoldable.Get_unfoldableArray(), "unfoldr"), Get_stepUnfoldr(), gopurs_runtime.Value{Type: 9, IntVal: 2861335956, UnsafePtr: unsafe.Pointer(&Constructor_IterNode[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](as_3), gopurs_runtime.Value{Type: 9, IntVal: 2509360378, UnsafePtr: unsafe.Pointer(nil)}})}).UnsafePtr)
					unboxed := make([]*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value], len(arr))
					for i, v := range arr { unboxed[i] = gopurs_runtime.CoerceToStruct[pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](v) }
					return unboxed
				}()
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v)} }
					return gopurs_runtime.Array(boxed)
				}()).StrVal()), gopurs_runtime.Str(")")).StrVal())).StrVal())
}))
}

func Call_isSubmap(dictOrd_0_loop *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value], dictEq_1_loop *pkg_Data_Eq.Constructor_Eq[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictOrd_0 *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value] = dictOrd_0_loop
_ = dictOrd_0
var dictEq_1 *pkg_Data_Eq.Constructor_Eq[gopurs_runtime.Value] = dictEq_1_loop
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
__local_var_5_1 := (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(m1_3.UnsafePtr).V2
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
__t5 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(nil))}
goto end_branch_5
} else {

}
}
{
if (v_7.Type == 9 && v_7.IntVal == 324739070 && v_7.UnsafePtr != nil) {
v1_8_3 := gopurs_runtime.Apply2(dictOrd_0.V1, __local_var_5_1, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_7.UnsafePtr).V2)
_ = v1_8_3
var __t4 gopurs_runtime.Value
{
if (uint32(v1_8_3.IntVal) == 1527465420) {
v_7_loop = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_7.UnsafePtr).V4)}
continue go__go_6_2_13
__t4 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](gopurs_runtime.Value{}))}
goto end_branch_4
} else {

}
}
{
if (uint32(v1_8_3.IntVal) == 380165415) {
v_7_loop = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_7.UnsafePtr).V5)}
continue go__go_6_2_13
__t4 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](gopurs_runtime.Value{}))}
goto end_branch_4
} else {

}
}
{
if (uint32(v1_8_3.IntVal) == 902936544) {
__t4 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_7.UnsafePtr).V3})}
goto end_branch_4
} else {

}
}
{
__t4 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_4:
__t5 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](__t4))}
goto end_branch_5
} else {

}
}
{
__t5 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_5:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](__t5))}
}
}()
})
v1_7_6 := gopurs_runtime.Apply(go__go_6_2_13, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](m2_4))})
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
__t7 = gopurs_runtime.Bool((Call_conj__3676519832(gopurs_runtime.Bool((gopurs_runtime.Apply2(dictEq_1.V0, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(m1_3.UnsafePtr).V3, (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(v1_7_6.UnsafePtr).V0).IntVal) != (0)), gopurs_runtime.Bool((Call_conj__3676519832(gopurs_runtime.Bool((gopurs_runtime.Apply2(go__go_2_0_12, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(m1_3.UnsafePtr).V4)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](m2_4))}).IntVal) != (0)), gopurs_runtime.Bool((gopurs_runtime.Apply2(go__go_2_0_12, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(m1_3.UnsafePtr).V5)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](m2_4))}).IntVal) != (0))).IntVal) != (0))).IntVal) != (0))
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

func Call_isEmpty(v_0_loop *Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]) bool {
var v_0 *Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value] = v_0_loop
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

func Call_intersectionWith(dictOrd_0_loop *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictOrd_0 *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value] = dictOrd_0_loop
_ = dictOrd_0
compare_1_0 := dictOrd_0.V1
_ = compare_1_0
return gopurs_runtime.Func(func(app_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m2_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_unsafeIntersectionWith(), compare_1_0, app_2, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](m1_3))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](m2_4))})))}
})
})
})
}

func Call_intersection(dictOrd_0_loop *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictOrd_0 *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value] = dictOrd_0_loop
_ = dictOrd_0
compare_1_0 := dictOrd_0.V1
_ = compare_1_0
return gopurs_runtime.Func(func(m1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m2_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_unsafeIntersectionWith(), compare_1_0, pkg_Data_Function.Get_go__const(), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](m1_2))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](m2_3))})))}
})
})
}

func Call_insertWith(dictOrd_0_loop *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value], app_1_loop gopurs_runtime.Value, k_2_loop gopurs_runtime.Value, v_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value] = dictOrd_0_loop
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
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(&Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]{1, 1, 1, k_2, v_3, gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(nil))}), gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(nil))})})}
goto end_branch_3
} else {

}
}
{
if (v1_5.Type == 9 && v1_5.IntVal == 324739070 && v1_5.UnsafePtr != nil) {
v2_6_1 := gopurs_runtime.Apply2(dictOrd_0.V1, k_2, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v1_5.UnsafePtr).V2)
_ = v2_6_1
var __t2 gopurs_runtime.Value
{
if (uint32(v2_6_1.IntVal) == 1527465420) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_unsafeBalancedNode(), (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v1_5.UnsafePtr).V2, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v1_5.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(go__go_4_0_14, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v1_5.UnsafePtr).V4)})))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v1_5.UnsafePtr).V5)})))}
goto end_branch_2
} else {

}
}
{
if (uint32(v2_6_1.IntVal) == 380165415) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_unsafeBalancedNode(), (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v1_5.UnsafePtr).V2, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v1_5.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v1_5.UnsafePtr).V4)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(go__go_4_0_14, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v1_5.UnsafePtr).V5)})))})))}
goto end_branch_2
} else {

}
}
{
if (uint32(v2_6_1.IntVal) == 902936544) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(&Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v1_5.UnsafePtr).V0, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v1_5.UnsafePtr).V1, k_2, gopurs_runtime.Apply2(app_1, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v1_5.UnsafePtr).V3, v_3), (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v1_5.UnsafePtr).V4, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v1_5.UnsafePtr).V5})}
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_2:
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__t2))}
goto end_branch_3
} else {

}
}
{
__t3 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_3:
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__t3))}
})
return go__go_4_0_14
}

func Call_insert(dictOrd_0_loop *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value], k_1_loop gopurs_runtime.Value, v_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value] = dictOrd_0_loop
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
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(&Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]{1, 1, 1, k_1, v_2, gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(nil))}), gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(nil))})})}
goto end_branch_3
} else {

}
}
{
if (v1_4.Type == 9 && v1_4.IntVal == 324739070 && v1_4.UnsafePtr != nil) {
v2_5_1 := gopurs_runtime.Apply2(dictOrd_0.V1, k_1, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v1_4.UnsafePtr).V2)
_ = v2_5_1
var __t2 gopurs_runtime.Value
{
if (uint32(v2_5_1.IntVal) == 1527465420) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_unsafeBalancedNode(), (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v1_4.UnsafePtr).V2, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v1_4.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(go__go_3_0_15, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v1_4.UnsafePtr).V4)})))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v1_4.UnsafePtr).V5)})))}
goto end_branch_2
} else {

}
}
{
if (uint32(v2_5_1.IntVal) == 380165415) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_unsafeBalancedNode(), (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v1_4.UnsafePtr).V2, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v1_4.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v1_4.UnsafePtr).V4)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(go__go_3_0_15, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v1_4.UnsafePtr).V5)})))})))}
goto end_branch_2
} else {

}
}
{
if (uint32(v2_5_1.IntVal) == 902936544) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(&Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v1_4.UnsafePtr).V0, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v1_4.UnsafePtr).V1, k_1, v_2, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v1_4.UnsafePtr).V4, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v1_4.UnsafePtr).V5})}
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_2:
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__t2))}
goto end_branch_3
} else {

}
}
{
__t3 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_3:
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__t3))}
})
return go__go_3_0_15
}

func Call_foldSubmapBy(dictOrd_0_loop *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value], appendFn_1_loop gopurs_runtime.Value, memptyValue_2_loop gopurs_runtime.Value, kmin_3_loop *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value], kmax_4_loop *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value], f_5_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value] = dictOrd_0_loop
_ = dictOrd_0
var appendFn_1 gopurs_runtime.Value = appendFn_1_loop
_ = appendFn_1
var memptyValue_2 gopurs_runtime.Value = memptyValue_2_loop
_ = memptyValue_2
var kmin_3 *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value] = kmin_3_loop
_ = kmin_3
var kmax_4 *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value] = kmax_4_loop
_ = kmax_4
var f_5 gopurs_runtime.Value = f_5_loop
_ = f_5
var __t4 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(kmin_3)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(kmin_3)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(kmin_3)}.UnsafePtr != nil) {
__local_var_6_1 := (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(kmin_3)}.UnsafePtr).V0
_ = __local_var_6_1
__t4 = gopurs_runtime.Func(func(k_7 gopurs_runtime.Value) gopurs_runtime.Value {
var __t3 bool
{
var __t_tag_2 gopurs_runtime.Value = gopurs_runtime.Apply2(dictOrd_0.V1, k_7, __local_var_6_1)
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
tooSmall_6_0 := __t4
_ = tooSmall_6_0
var __t9 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(kmax_4)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(kmax_4)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(kmax_4)}.UnsafePtr != nil) {
__local_var_7_6 := (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(kmax_4)}.UnsafePtr).V0
_ = __local_var_7_6
__t9 = gopurs_runtime.Func(func(k_8 gopurs_runtime.Value) gopurs_runtime.Value {
var __t8 bool
{
var __t_tag_7 gopurs_runtime.Value = gopurs_runtime.Apply2(dictOrd_0.V1, k_8, __local_var_7_6)
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
tooLarge_7_5 := __t9
_ = tooLarge_7_5
var __t25 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(kmin_3)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(kmin_3)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(kmin_3)}.UnsafePtr != nil) {
var __t20 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(kmax_4)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(kmax_4)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(kmax_4)}.UnsafePtr != nil) {
__local_var_8_11 := (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(kmax_4)}.UnsafePtr).V0
_ = __local_var_8_11
__local_var_9_12 := (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(kmin_3)}.UnsafePtr).V0
_ = __local_var_9_12
__t20 = gopurs_runtime.Func(func(k_10 gopurs_runtime.Value) gopurs_runtime.Value {
var __t14 bool
{
var __t_tag_13 gopurs_runtime.Value = gopurs_runtime.Apply2(dictOrd_0.V1, __local_var_9_12, k_10)
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
var __t16 bool
{
var __t_tag_15 gopurs_runtime.Value = gopurs_runtime.Apply2(dictOrd_0.V1, k_10, __local_var_8_11)
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
return gopurs_runtime.Bool((Call_conj__3676519832(gopurs_runtime.Bool(__t14), gopurs_runtime.Bool(__t16)).IntVal) != (0))
})
goto end_branch_20
} else {

}
}
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(kmax_4)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(kmax_4)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(kmax_4)}.UnsafePtr == nil) {
__local_var_8_17 := (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(kmin_3)}.UnsafePtr).V0
_ = __local_var_8_17
__t20 = gopurs_runtime.Func(func(k_9 gopurs_runtime.Value) gopurs_runtime.Value {
var __t19 bool
{
var __t_tag_18 gopurs_runtime.Value = gopurs_runtime.Apply2(dictOrd_0.V1, __local_var_8_17, k_9)
if (uint32(__t_tag_18.IntVal) == 380165415) {
__t19 = false
goto end_branch_19
} else {

}
}
{
__t19 = true
}
end_branch_19:
return gopurs_runtime.Bool(__t19)
})
goto end_branch_20
} else {

}
}
{
__t20 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_20:
__t25 = __t20
goto end_branch_25
} else {

}
}
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(kmin_3)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(kmin_3)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(kmin_3)}.UnsafePtr == nil) {
var __t24 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(kmax_4)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(kmax_4)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(kmax_4)}.UnsafePtr != nil) {
__local_var_8_21 := (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(kmax_4)}.UnsafePtr).V0
_ = __local_var_8_21
__t24 = gopurs_runtime.Func(func(k_9 gopurs_runtime.Value) gopurs_runtime.Value {
var __t23 bool
{
var __t_tag_22 gopurs_runtime.Value = gopurs_runtime.Apply2(dictOrd_0.V1, k_9, __local_var_8_21)
if (uint32(__t_tag_22.IntVal) == 380165415) {
__t23 = false
goto end_branch_23
} else {

}
}
{
__t23 = true
}
end_branch_23:
return gopurs_runtime.Bool(__t23)
})
goto end_branch_24
} else {

}
}
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(kmax_4)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(kmax_4)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(kmax_4)}.UnsafePtr == nil) {
__t24 = gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(true)
})
goto end_branch_24
} else {

}
}
{
__t24 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_24:
__t25 = __t24
goto end_branch_25
} else {

}
}
{
__t25 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_25:
inBounds_8_10 := __t25
_ = inBounds_8_10
var go__go_9_26_26 gopurs_runtime.Value
_ = go__go_9_26_26
go__go_9_26_26 = gopurs_runtime.Func(func(v_10 gopurs_runtime.Value) gopurs_runtime.Value {
var __t30 gopurs_runtime.Value
{
if (v_10.Type == 9 && v_10.IntVal == 324739070 && v_10.UnsafePtr == nil) {
__t30 = memptyValue_2
goto end_branch_30
} else {

}
}
{
if (v_10.Type == 9 && v_10.IntVal == 324739070 && v_10.UnsafePtr != nil) {
var __t27 gopurs_runtime.Value
{
if (gopurs_runtime.Apply(tooSmall_6_0, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_10.UnsafePtr).V2).IntVal) != (0) {
__t27 = memptyValue_2
goto end_branch_27
} else {

}
}
{
__t27 = gopurs_runtime.Apply(go__go_9_26_26, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_10.UnsafePtr).V4)})
}
end_branch_27:
var __t28 gopurs_runtime.Value
{
if (gopurs_runtime.Apply(inBounds_8_10, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_10.UnsafePtr).V2).IntVal) != (0) {
__t28 = gopurs_runtime.Apply2(f_5, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_10.UnsafePtr).V2, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_10.UnsafePtr).V3)
goto end_branch_28
} else {

}
}
{
__t28 = memptyValue_2
}
end_branch_28:
var __t29 gopurs_runtime.Value
{
if (gopurs_runtime.Apply(tooLarge_7_5, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_10.UnsafePtr).V2).IntVal) != (0) {
__t29 = memptyValue_2
goto end_branch_29
} else {

}
}
{
__t29 = gopurs_runtime.Apply(go__go_9_26_26, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_10.UnsafePtr).V5)})
}
end_branch_29:
__t30 = gopurs_runtime.Apply2(appendFn_1, gopurs_runtime.Apply2(appendFn_1, __t27, __t28), __t29)
goto end_branch_30
} else {

}
}
{
__t30 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_30:
return __t30
})
return go__go_9_26_26
}

func Call_foldSubmap(dictOrd_0_loop *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value], dictMonoid_1_loop *pkg_Data_Monoid.Constructor_Monoid[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictOrd_0 *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value] = dictOrd_0_loop
_ = dictOrd_0
var dictMonoid_1 *pkg_Data_Monoid.Constructor_Monoid[gopurs_runtime.Value] = dictMonoid_1_loop
_ = dictMonoid_1
return gopurs_runtime.Apply3(Get_foldSubmapBy(), gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(dictOrd_0)}, gopurs_runtime.RecordGet(gopurs_runtime.Apply(dictMonoid_1.V0, gopurs_runtime.Value{}), "append"), dictMonoid_1.V1)
}

func Call_findMin(v_0_loop *Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]) *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value] {
findMin:
for {
if false { continue findMin }
var v_0 *Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value] = v_0_loop
_ = v_0
var __t2 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_0)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_0)}.IntVal == 324739070 && gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr == nil) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(nil))}
goto end_branch_2
} else {

}
}
{
if (gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_0)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_0)}.IntVal == 324739070 && gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr != nil) {
var __t1 gopurs_runtime.Value
{
var __t_tag_0 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V4)}
if (__t_tag_0.Type == 9 && __t_tag_0.IntVal == 324739070 && __t_tag_0.UnsafePtr == nil) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, gopurs_runtime.RecordDict2("key", "value", (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V2, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V3)})}
goto end_branch_1
} else {

}
}
{
v_0_loop = (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V4
continue findMin
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](gopurs_runtime.Value{}))}
}
end_branch_1:
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](__t1))}
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_2:
return gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](__t2)
}
}

func Call_lookupGT(dictOrd_0_loop *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value], k_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value] = dictOrd_0_loop
_ = dictOrd_0
var k_1 gopurs_runtime.Value = k_1_loop
_ = k_1
var go__go_2_0_27 gopurs_runtime.Value
_ = go__go_2_0_27
go__go_2_0_27 = gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t5 gopurs_runtime.Value
{
if (v_3.Type == 9 && v_3.IntVal == 324739070 && v_3.UnsafePtr == nil) {
__t5 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(nil))}
goto end_branch_5
} else {

}
}
{
if (v_3.Type == 9 && v_3.IntVal == 324739070 && v_3.UnsafePtr != nil) {
v1_4_1 := gopurs_runtime.Apply2(dictOrd_0.V1, k_1, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V2)
_ = v1_4_1
var __t4 gopurs_runtime.Value
{
if (uint32(v1_4_1.IntVal) == 1527465420) {
v2_5_2 := gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(go__go_2_0_27, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V4)}))
_ = v2_5_2
var __t3 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_5_2)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_5_2)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_5_2)}.UnsafePtr == nil) {
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, gopurs_runtime.RecordDict2("key", "value", (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V2, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V3)})}
goto end_branch_3
} else {

}
}
{
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_5_2)}
}
end_branch_3:
__t4 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](__t3))}
goto end_branch_4
} else {

}
}
{
if (uint32(v1_4_1.IntVal) == 380165415) {
__t4 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(go__go_2_0_27, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V5)})))}
goto end_branch_4
} else {

}
}
{
if (uint32(v1_4_1.IntVal) == 902936544) {
__t4 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_findMin((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V5))}))}
goto end_branch_4
} else {

}
}
{
__t4 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_4:
__t5 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](__t4))}
goto end_branch_5
} else {

}
}
{
__t5 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_5:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](__t5))}
})
return go__go_2_0_27
}

func Call_findMax(v_0_loop *Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]) *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value] {
findMax:
for {
if false { continue findMax }
var v_0 *Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value] = v_0_loop
_ = v_0
var __t2 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_0)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_0)}.IntVal == 324739070 && gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr == nil) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(nil))}
goto end_branch_2
} else {

}
}
{
if (gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_0)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_0)}.IntVal == 324739070 && gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr != nil) {
var __t1 gopurs_runtime.Value
{
var __t_tag_0 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V5)}
if (__t_tag_0.Type == 9 && __t_tag_0.IntVal == 324739070 && __t_tag_0.UnsafePtr == nil) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, gopurs_runtime.RecordDict2("key", "value", (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V2, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V3)})}
goto end_branch_1
} else {

}
}
{
v_0_loop = (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V5
continue findMax
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](gopurs_runtime.Value{}))}
}
end_branch_1:
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](__t1))}
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_2:
return gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](__t2)
}
}

func Call_lookupLT(dictOrd_0_loop *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value], k_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value] = dictOrd_0_loop
_ = dictOrd_0
var k_1 gopurs_runtime.Value = k_1_loop
_ = k_1
var go__go_2_0_28 gopurs_runtime.Value
_ = go__go_2_0_28
go__go_2_0_28 = gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t5 gopurs_runtime.Value
{
if (v_3.Type == 9 && v_3.IntVal == 324739070 && v_3.UnsafePtr == nil) {
__t5 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(nil))}
goto end_branch_5
} else {

}
}
{
if (v_3.Type == 9 && v_3.IntVal == 324739070 && v_3.UnsafePtr != nil) {
v1_4_1 := gopurs_runtime.Apply2(dictOrd_0.V1, k_1, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V2)
_ = v1_4_1
var __t4 gopurs_runtime.Value
{
if (uint32(v1_4_1.IntVal) == 1527465420) {
__t4 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(go__go_2_0_28, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V4)})))}
goto end_branch_4
} else {

}
}
{
if (uint32(v1_4_1.IntVal) == 380165415) {
v2_5_2 := gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(go__go_2_0_28, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V5)}))
_ = v2_5_2
var __t3 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_5_2)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_5_2)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_5_2)}.UnsafePtr == nil) {
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, gopurs_runtime.RecordDict2("key", "value", (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V2, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V3)})}
goto end_branch_3
} else {

}
}
{
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_5_2)}
}
end_branch_3:
__t4 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](__t3))}
goto end_branch_4
} else {

}
}
{
if (uint32(v1_4_1.IntVal) == 902936544) {
__t4 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_findMax((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V4))}))}
goto end_branch_4
} else {

}
}
{
__t4 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_4:
__t5 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](__t4))}
goto end_branch_5
} else {

}
}
{
__t5 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_5:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](__t5))}
})
return go__go_2_0_28
}

func Call_filterWithKey(dictOrd_0_loop *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value], f_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value] = dictOrd_0_loop
_ = dictOrd_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var go__go_2_0_29 gopurs_runtime.Value
_ = go__go_2_0_29
go__go_2_0_29 = gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 gopurs_runtime.Value
{
if (v_3.Type == 9 && v_3.IntVal == 324739070 && v_3.UnsafePtr == nil) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(nil))}
goto end_branch_2
} else {

}
}
{
if (v_3.Type == 9 && v_3.IntVal == 324739070 && v_3.UnsafePtr != nil) {
var __t1 *Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]
{
if (gopurs_runtime.Apply2(f_1, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V2, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V3).IntVal) != (0) {
__t1 = gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_unsafeBalancedNode(), (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V2, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(go__go_2_0_29, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V4)})))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(go__go_2_0_29, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V5)})))}))
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp2(Get_unsafeJoinNodes(), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(go__go_2_0_29, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V4)})))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(go__go_2_0_29, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V5)})))}))
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
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__t2))}
})
return go__go_2_0_29
}

func Call_filterKeys(dictOrd_0_loop *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value], f_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value] = dictOrd_0_loop
_ = dictOrd_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var go__go_2_0_30 gopurs_runtime.Value
_ = go__go_2_0_30
go__go_2_0_30 = gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 gopurs_runtime.Value
{
if (v_3.Type == 9 && v_3.IntVal == 324739070 && v_3.UnsafePtr == nil) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(nil))}
goto end_branch_2
} else {

}
}
{
if (v_3.Type == 9 && v_3.IntVal == 324739070 && v_3.UnsafePtr != nil) {
var __t1 *Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]
{
if (gopurs_runtime.Apply(f_1, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V2).IntVal) != (0) {
__t1 = gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_unsafeBalancedNode(), (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V2, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(go__go_2_0_30, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V4)})))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(go__go_2_0_30, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V5)})))}))
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp2(Get_unsafeJoinNodes(), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(go__go_2_0_30, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V4)})))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(go__go_2_0_30, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V5)})))}))
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
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__t2))}
})
return go__go_2_0_30
}

func Call_filter(dictOrd_0_loop *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value], x_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value] = dictOrd_0_loop
_ = dictOrd_0
var x_1 gopurs_runtime.Value = x_1_loop
_ = x_1
return Call_filterWithKey(dictOrd_0, gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return x_1
}))
}

func Call_eqMap(dictEq_0_loop gopurs_runtime.Value, dictEq1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictEq_0 gopurs_runtime.Value = dictEq_0_loop
_ = dictEq_0
var dictEq1_1 gopurs_runtime.Value = dictEq1_1_loop
_ = dictEq1_1
eqMapIter2_2_0 := gopurs_runtime.CoerceToStruct[pkg_Data_Eq.Constructor_Eq[gopurs_runtime.Value]](Call_eqMapIter(dictEq_0, dictEq1_1))
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
if ((ys_4.Type == 9 && ys_4.IntVal == 324739070 && ys_4.UnsafePtr != nil)) && ((Call_eq__2843686287(gopurs_runtime.Int((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(xs_3.UnsafePtr).V1), gopurs_runtime.Int((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(ys_4.UnsafePtr).V1)).IntVal) != (0)) {
__t2 = (gopurs_runtime.Apply2(eqMapIter2_2_0.V0, gopurs_runtime.Value{Type: 9, IntVal: 2861335956, UnsafePtr: unsafe.Pointer(&Constructor_IterNode[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](xs_3), gopurs_runtime.Value{Type: 9, IntVal: 2509360378, UnsafePtr: unsafe.Pointer(nil)}})}, gopurs_runtime.Value{Type: 9, IntVal: 2861335956, UnsafePtr: unsafe.Pointer(&Constructor_IterNode[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](ys_4), gopurs_runtime.Value{Type: 9, IntVal: 2509360378, UnsafePtr: unsafe.Pointer(nil)}})}).IntVal) != (0)
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

func Call_ordMap(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
ordMapIter1_1_0 := Call_ordMapIter(dictOrd_0)
_ = ordMapIter1_1_0
eqMap1_2_1 := gopurs_runtime.Apply(Get_eqMap(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd_0, "Eq0"), gopurs_runtime.Value{}))
_ = eqMap1_2_1
return gopurs_runtime.Func(func(dictOrd1_3 gopurs_runtime.Value) gopurs_runtime.Value {
ordMapIter2_4_2 := gopurs_runtime.CoerceToStruct[pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value]](gopurs_runtime.Apply(ordMapIter1_1_0, dictOrd1_3))
_ = ordMapIter2_4_2
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
__t4 = uint32(gopurs_runtime.Apply2(ordMapIter2_4_2.V1, gopurs_runtime.Value{Type: 9, IntVal: 2861335956, UnsafePtr: unsafe.Pointer(&Constructor_IterNode[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](xs_6), gopurs_runtime.Value{Type: 9, IntVal: 2509360378, UnsafePtr: unsafe.Pointer(nil)}})}, gopurs_runtime.Value{Type: 9, IntVal: 2861335956, UnsafePtr: unsafe.Pointer(&Constructor_IterNode[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](ys_7), gopurs_runtime.Value{Type: 9, IntVal: 2509360378, UnsafePtr: unsafe.Pointer(nil)}})}).IntVal)
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

func Call_eq1Map(dictEq_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictEq_0 gopurs_runtime.Value = dictEq_0_loop
_ = dictEq_0
return gopurs_runtime.RecordDict1("eq1", gopurs_runtime.Func(func(dictEq1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(Call_eqMap(dictEq_0, dictEq1_1), "eq")
}))
}

func Call_ord1Map(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
ordMap1_1_0 := Call_ordMap(dictOrd_0)
_ = ordMap1_1_0
__local_var_2_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd_0, "Eq0"), gopurs_runtime.Value{})
_ = __local_var_2_2
eq1Map1_2_1 := gopurs_runtime.RecordDict1("eq1", gopurs_runtime.Func(func(dictEq1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(Call_eqMap(__local_var_2_2, dictEq1_3), "eq")
}))
_ = eq1Map1_2_1
return gopurs_runtime.RecordDict2("Eq10", "compare1", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return eq1Map1_2_1
}), gopurs_runtime.Func(func(dictOrd1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(gopurs_runtime.Apply(ordMap1_1_0, dictOrd1_3), "compare")
}))
}

func Call_fromFoldable(dictOrd_0_loop *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value], dictFoldable_1_loop *pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictOrd_0 *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value] = dictOrd_0_loop
_ = dictOrd_0
var dictFoldable_1 *pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value] = dictFoldable_1_loop
_ = dictFoldable_1
return gopurs_runtime.Apply2(dictFoldable_1.V1, gopurs_runtime.Func(func(m_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(Call_insert(dictOrd_0, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V0, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V1), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](m_2))})))}
})
}), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(nil))})
}

func Call_fromFoldableWith(dictOrd_0_loop *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value], dictFoldable_1_loop *pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value], f_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value] = dictOrd_0_loop
_ = dictOrd_0
var dictFoldable_1 *pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value] = dictFoldable_1_loop
_ = dictFoldable_1
var f_2 gopurs_runtime.Value = f_2_loop
_ = f_2
return gopurs_runtime.Apply2(dictFoldable_1.V1, gopurs_runtime.Func(func(m_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(Call_insertWith(dictOrd_0, gopurs_runtime.Func(func(b_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_2, a_6, b_5)
})
}), (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V0, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V1), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](m_3))})))}
})
}), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(nil))})
}

func Call_fromFoldableWithIndex(dictOrd_0_loop *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value], dictFoldableWithIndex_1_loop *pkg_Data_FoldableWithIndex.Constructor_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var dictOrd_0 *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value] = dictOrd_0_loop
_ = dictOrd_0
var dictFoldableWithIndex_1 *pkg_Data_FoldableWithIndex.Constructor_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value] = dictFoldableWithIndex_1_loop
_ = dictFoldableWithIndex_1
return gopurs_runtime.Apply2(dictFoldableWithIndex_1.V2, gopurs_runtime.Func(func(k_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(Call_insert(dictOrd_0, k_2, v_4), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](m_3))})))}
})
})
}), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(nil))})
}

func Call_monoidSemigroupMap(_dollar__unused_0_loop gopurs_runtime.Value, dictOrd_1_loop gopurs_runtime.Value, dictSemigroup_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var _dollar__unused_0 gopurs_runtime.Value = _dollar__unused_0_loop
_ = _dollar__unused_0
var dictOrd_1 gopurs_runtime.Value = dictOrd_1_loop
_ = dictOrd_1
var dictSemigroup_2 gopurs_runtime.Value = dictSemigroup_2_loop
_ = dictSemigroup_2
semigroupMap3_3_0 := Call_semigroupMap(gopurs_runtime.Value{}, dictOrd_1, dictSemigroup_2)
_ = semigroupMap3_3_0
return gopurs_runtime.RecordDict2("Semigroup0", "mempty", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return semigroupMap3_3_0
}), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(nil))})
}

func Call_submap(dictOrd_0_loop *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictOrd_0 *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value] = dictOrd_0_loop
_ = dictOrd_0
compare_1_1 := dictOrd_0.V1
_ = compare_1_1
union1_1_0 := gopurs_runtime.Func(func(m1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m2_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_unsafeUnionWith(), compare_1_1, pkg_Data_Function.Get_go__const(), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](m1_2))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](m2_3))})))}
})
})
_ = union1_1_0
return gopurs_runtime.Func(func(kmin_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(kmax_3 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldSubmapBy(dictOrd_0, union1_1_0, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(nil))}, gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](kmin_2), gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](kmax_3), Get_singleton())
})
})
}

func Call_unions(dictOrd_0_loop *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictOrd_0 *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value] = dictOrd_0_loop
_ = dictOrd_0
compare_1_1 := dictOrd_0.V1
_ = compare_1_1
union1_1_0 := gopurs_runtime.Func(func(m1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m2_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_unsafeUnionWith(), compare_1_1, pkg_Data_Function.Get_go__const(), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](m1_2))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](m2_3))})))}
})
})
_ = union1_1_0
return gopurs_runtime.Func(func(dictFoldable_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFoldable_2, "foldl"), union1_1_0, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(nil))})
})
}

func Call_difference(dictOrd_0_loop *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictOrd_0 *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value] = dictOrd_0_loop
_ = dictOrd_0
compare_1_0 := dictOrd_0.V1
_ = compare_1_0
return gopurs_runtime.Func(func(m1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m2_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp3(Get_unsafeDifference(), compare_1_0, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](m1_2))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](m2_3))})))}
})
})
}

func Call_delete(dictOrd_0_loop *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value], k_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value] = dictOrd_0_loop
_ = dictOrd_0
var k_1 gopurs_runtime.Value = k_1_loop
_ = k_1
var go__go_2_0_31 gopurs_runtime.Value
_ = go__go_2_0_31
go__go_2_0_31 = gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t3 gopurs_runtime.Value
{
if (v_3.Type == 9 && v_3.IntVal == 324739070 && v_3.UnsafePtr == nil) {
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(nil))}
goto end_branch_3
} else {

}
}
{
if (v_3.Type == 9 && v_3.IntVal == 324739070 && v_3.UnsafePtr != nil) {
v1_4_1 := gopurs_runtime.Apply2(dictOrd_0.V1, k_1, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V2)
_ = v1_4_1
var __t2 gopurs_runtime.Value
{
if (uint32(v1_4_1.IntVal) == 1527465420) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_unsafeBalancedNode(), (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V2, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(go__go_2_0_31, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V4)})))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V5)})))}
goto end_branch_2
} else {

}
}
{
if (uint32(v1_4_1.IntVal) == 380165415) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_unsafeBalancedNode(), (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V2, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V4)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(go__go_2_0_31, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V5)})))})))}
goto end_branch_2
} else {

}
}
{
if (uint32(v1_4_1.IntVal) == 902936544) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp2(Get_unsafeJoinNodes(), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V4)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V5)})))}
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_2:
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__t2))}
goto end_branch_3
} else {

}
}
{
__t3 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_3:
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__t3))}
})
return go__go_2_0_31
}

func Call_checkValid(dictOrd_0_loop *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictOrd_0 *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value] = dictOrd_0_loop
_ = dictOrd_0
var go__go_1_0_32 gopurs_runtime.Value
_ = go__go_1_0_32
go__go_1_0_32 = gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t18 gopurs_runtime.Value
{
if (v_2.Type == 9 && v_2.IntVal == 324739070 && v_2.UnsafePtr == nil) {
__t18 = gopurs_runtime.Bool(true)
goto end_branch_18
} else {

}
}
{
if (v_2.Type == 9 && v_2.IntVal == 324739070 && v_2.UnsafePtr != nil) {
var __t17 gopurs_runtime.Value
{
var __t_tag_1 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V4)}
if (__t_tag_1.Type == 9 && __t_tag_1.IntVal == 324739070 && __t_tag_1.UnsafePtr == nil) {
var __t6 gopurs_runtime.Value
{
var __t_tag_2 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V5)}
if (__t_tag_2.Type == 9 && __t_tag_2.IntVal == 324739070 && __t_tag_2.UnsafePtr == nil) {
__t6 = gopurs_runtime.Bool(true)
goto end_branch_6
} else {

}
}
{
var __t_tag_3 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V5)}
if (__t_tag_3.Type == 9 && __t_tag_3.IntVal == 324739070 && __t_tag_3.UnsafePtr != nil) {
var __t5 bool
{
var __t_tag_4 gopurs_runtime.Value = gopurs_runtime.Apply2(dictOrd_0.V1, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V5)}.UnsafePtr).V2, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V2)
if (uint32(__t_tag_4.IntVal) == 380165415) {
__t5 = true
goto end_branch_5
} else {

}
}
{
__t5 = false
}
end_branch_5:
__t6 = gopurs_runtime.Bool((Call_conj__3676519832(gopurs_runtime.Bool((Call_eq__2843686287(gopurs_runtime.Int((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V0), gopurs_runtime.Int(2)).IntVal) != (0)), gopurs_runtime.Bool((Call_conj__3676519832(gopurs_runtime.Bool((Call_eq__2843686287(gopurs_runtime.Int((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V5)}.UnsafePtr).V0), gopurs_runtime.Int(1)).IntVal) != (0)), gopurs_runtime.Bool((Call_conj__3676519832(gopurs_runtime.Bool((gopurs_runtime.Bool(Call_greaterThan__4087042607(gopurs_runtime.Int((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V1), gopurs_runtime.Int((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V5)}.UnsafePtr).V1))).IntVal) != (0)), gopurs_runtime.Bool((Call_conj__3676519832(gopurs_runtime.Bool(__t5), gopurs_runtime.Bool((gopurs_runtime.Apply(go__go_1_0_32, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V5)}).IntVal) != (0))).IntVal) != (0))).IntVal) != (0))).IntVal) != (0))).IntVal) != (0))
goto end_branch_6
} else {

}
}
{
__t6 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_6:
__t17 = gopurs_runtime.Bool((__t6.IntVal) != (0))
goto end_branch_17
} else {

}
}
{
var __t_tag_7 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V4)}
if (__t_tag_7.Type == 9 && __t_tag_7.IntVal == 324739070 && __t_tag_7.UnsafePtr != nil) {
var __t16 gopurs_runtime.Value
{
var __t_tag_8 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V5)}
if (__t_tag_8.Type == 9 && __t_tag_8.IntVal == 324739070 && __t_tag_8.UnsafePtr == nil) {
var __t10 bool
{
var __t_tag_9 gopurs_runtime.Value = gopurs_runtime.Apply2(dictOrd_0.V1, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V4)}.UnsafePtr).V2, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V2)
if (uint32(__t_tag_9.IntVal) == 1527465420) {
__t10 = true
goto end_branch_10
} else {

}
}
{
__t10 = false
}
end_branch_10:
__t16 = gopurs_runtime.Bool((Call_conj__3676519832(gopurs_runtime.Bool((Call_eq__2843686287(gopurs_runtime.Int((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V0), gopurs_runtime.Int(2)).IntVal) != (0)), gopurs_runtime.Bool((Call_conj__3676519832(gopurs_runtime.Bool((Call_eq__2843686287(gopurs_runtime.Int((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V4)}.UnsafePtr).V0), gopurs_runtime.Int(1)).IntVal) != (0)), gopurs_runtime.Bool((Call_conj__3676519832(gopurs_runtime.Bool((gopurs_runtime.Bool(Call_greaterThan__4087042607(gopurs_runtime.Int((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V1), gopurs_runtime.Int((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V4)}.UnsafePtr).V1))).IntVal) != (0)), gopurs_runtime.Bool((Call_conj__3676519832(gopurs_runtime.Bool(__t10), gopurs_runtime.Bool((gopurs_runtime.Apply(go__go_1_0_32, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V4)}).IntVal) != (0))).IntVal) != (0))).IntVal) != (0))).IntVal) != (0))).IntVal) != (0))
goto end_branch_16
} else {

}
}
{
var __t_tag_11 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V5)}
if (__t_tag_11.Type == 9 && __t_tag_11.IntVal == 324739070 && __t_tag_11.UnsafePtr != nil) {
var __t13 bool
{
var __t_tag_12 gopurs_runtime.Value = gopurs_runtime.Apply2(dictOrd_0.V1, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V5)}.UnsafePtr).V2, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V2)
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
var __t15 bool
{
var __t_tag_14 gopurs_runtime.Value = gopurs_runtime.Apply2(dictOrd_0.V1, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V4)}.UnsafePtr).V2, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V2)
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
__t16 = gopurs_runtime.Bool((Call_conj__3676519832(gopurs_runtime.Bool((gopurs_runtime.Bool(Call_greaterThan__4087042607(gopurs_runtime.Int((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V0), gopurs_runtime.Int((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V5)}.UnsafePtr).V0))).IntVal) != (0)), gopurs_runtime.Bool((Call_conj__3676519832(gopurs_runtime.Bool(__t13), gopurs_runtime.Bool((Call_conj__3676519832(gopurs_runtime.Bool((gopurs_runtime.Bool(Call_greaterThan__4087042607(gopurs_runtime.Int((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V0), gopurs_runtime.Int((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V4)}.UnsafePtr).V0))).IntVal) != (0)), gopurs_runtime.Bool((Call_conj__3676519832(gopurs_runtime.Bool(__t15), gopurs_runtime.Bool((Call_conj__3676519832(gopurs_runtime.Bool((gopurs_runtime.Bool(Call_lessThan__4087042607(gopurs_runtime.Int(Call_abs__1599282999(gopurs_runtime.Int(Call_sub__1043827704(gopurs_runtime.Int((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V5)}.UnsafePtr).V0), gopurs_runtime.Int((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V4)}.UnsafePtr).V0)).IntVal)).IntVal), gopurs_runtime.Int(2))).IntVal) != (0)), gopurs_runtime.Bool((Call_conj__3676519832(gopurs_runtime.Bool((Call_eq__2843686287(gopurs_runtime.Int(Call_add__560788792(gopurs_runtime.Int(Call_add__560788792(gopurs_runtime.Int((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V5)}.UnsafePtr).V1), gopurs_runtime.Int((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V4)}.UnsafePtr).V1)).IntVal), gopurs_runtime.Int(1)).IntVal), gopurs_runtime.Int((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V1)).IntVal) != (0)), gopurs_runtime.Bool((Call_conj__3676519832(gopurs_runtime.Bool((gopurs_runtime.Apply(go__go_1_0_32, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V4)}).IntVal) != (0)), gopurs_runtime.Bool((gopurs_runtime.Apply(go__go_1_0_32, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V5)}).IntVal) != (0))).IntVal) != (0))).IntVal) != (0))).IntVal) != (0))).IntVal) != (0))).IntVal) != (0))).IntVal) != (0))).IntVal) != (0))
goto end_branch_16
} else {

}
}
{
__t16 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_16:
__t17 = gopurs_runtime.Bool((__t16.IntVal) != (0))
goto end_branch_17
} else {

}
}
{
__t17 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_17:
__t18 = gopurs_runtime.Bool((__t17.IntVal) != (0))
goto end_branch_18
} else {

}
}
{
__t18 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_18:
return gopurs_runtime.Bool((__t18.IntVal) != (0))
})
return go__go_1_0_32
}

func Call_catMaybes(dictOrd_0_loop *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictOrd_0 *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value] = dictOrd_0_loop
_ = dictOrd_0
return Call_mapMaybeWithKey(dictOrd_0, gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_identity1()
}))
}

func Call_applyMap(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
compare_1_0 := gopurs_runtime.RecordGet(dictOrd_0, "compare")
_ = compare_1_0
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_functorMap()
}), gopurs_runtime.Func(func(m1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m2_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_unsafeIntersectionWith(), compare_1_0, Get_identity2(), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](m1_2))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](m2_3))})))}
})
}))
}

func Call_bindMap(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
compare_1_1 := gopurs_runtime.RecordGet(dictOrd_0, "compare")
_ = compare_1_1
applyMap1_1_0 := gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_functorMap()
}), gopurs_runtime.Func(func(m1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m2_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_unsafeIntersectionWith(), compare_1_1, Get_identity2(), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](m1_2))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](m2_3))})))}
})
}))
_ = applyMap1_1_0
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return applyMap1_1_0
}), gopurs_runtime.Func(func(m_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(Call_mapMaybeWithKey(gopurs_runtime.CoerceToStruct[pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value]](dictOrd_0), gopurs_runtime.Func(func(k_4 gopurs_runtime.Value) gopurs_runtime.Value {
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
__t6 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(nil))}
goto end_branch_6
} else {

}
}
{
if (v_6.Type == 9 && v_6.IntVal == 324739070 && v_6.UnsafePtr != nil) {
v1_7_4 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictOrd_0, "compare"), k_4, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_6.UnsafePtr).V2)
_ = v1_7_4
var __t5 gopurs_runtime.Value
{
if (uint32(v1_7_4.IntVal) == 1527465420) {
v_6_loop = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_6.UnsafePtr).V4)}
continue go__go_5_3_33
__t5 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](gopurs_runtime.Value{}))}
goto end_branch_5
} else {

}
}
{
if (uint32(v1_7_4.IntVal) == 380165415) {
v_6_loop = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_6.UnsafePtr).V5)}
continue go__go_5_3_33
__t5 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](gopurs_runtime.Value{}))}
goto end_branch_5
} else {

}
}
{
if (uint32(v1_7_4.IntVal) == 902936544) {
__t5 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_6.UnsafePtr).V3})}
goto end_branch_5
} else {

}
}
{
__t5 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_5:
__t6 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](__t5))}
goto end_branch_6
} else {

}
}
{
__t6 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_6:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](__t6))}
}
}()
})
__local_var_5_2 := go__go_5_3_33
_ = __local_var_5_2
return gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_5_2, gopurs_runtime.Apply(f_3, x_6))
})
})), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](m_2))})))}
})
}))
}

func Call_anyWithKey(predicate_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
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
__t1 = gopurs_runtime.Bool((Call_disj__3676519832(gopurs_runtime.Bool((gopurs_runtime.Apply2(predicate_0, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V2, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V3).IntVal) != (0)), gopurs_runtime.Bool((Call_disj__3676519832(gopurs_runtime.Bool((gopurs_runtime.Apply(go__go_1_0_34, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V4)}).IntVal) != (0)), gopurs_runtime.Bool((gopurs_runtime.Apply(go__go_1_0_34, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V5)}).IntVal) != (0))).IntVal) != (0))).IntVal) != (0))
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

func Call_any(predicate_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
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
__t1 = gopurs_runtime.Bool((Call_disj__3676519832(gopurs_runtime.Bool((gopurs_runtime.Apply(predicate_0, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V3).IntVal) != (0)), gopurs_runtime.Bool((Call_disj__3676519832(gopurs_runtime.Bool((gopurs_runtime.Apply(go__go_1_0_35, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V4)}).IntVal) != (0)), gopurs_runtime.Bool((gopurs_runtime.Apply(go__go_1_0_35, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V5)}).IntVal) != (0))).IntVal) != (0))).IntVal) != (0))
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

func Call_alter(dictOrd_0_loop *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictOrd_0 *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value] = dictOrd_0_loop
_ = dictOrd_0
compare_1_0 := dictOrd_0.V1
_ = compare_1_0
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m_4 gopurs_runtime.Value) gopurs_runtime.Value {
v_5_1 := gopurs_runtime.UncurriedApp3(Get_unsafeSplit(), compare_1_0, k_3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](m_4))})
_ = v_5_1
v2_6_2 := gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(f_2, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Split[gopurs_runtime.Value, gopurs_runtime.Value])(v_5_1.UnsafePtr).V0)}))
_ = v2_6_2
var __t3 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_6_2)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_6_2)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_6_2)}.UnsafePtr == nil) {
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp2(Get_unsafeJoinNodes(), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Split[gopurs_runtime.Value, gopurs_runtime.Value])(v_5_1.UnsafePtr).V1)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Split[gopurs_runtime.Value, gopurs_runtime.Value])(v_5_1.UnsafePtr).V2)})))}
goto end_branch_3
} else {

}
}
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_6_2)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_6_2)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_6_2)}.UnsafePtr != nil) {
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_unsafeBalancedNode(), k_3, (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_6_2)}.UnsafePtr).V0, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Split[gopurs_runtime.Value, gopurs_runtime.Value])(v_5_1.UnsafePtr).V1)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Split[gopurs_runtime.Value, gopurs_runtime.Value])(v_5_1.UnsafePtr).V2)})))}
goto end_branch_3
} else {

}
}
{
__t3 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_3:
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__t3))}
})
})
})
}

func Call_altMap(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
compare_1_0 := gopurs_runtime.RecordGet(dictOrd_0, "compare")
_ = compare_1_0
return gopurs_runtime.RecordDict2("Functor0", "alt", gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_functorMap()
}), gopurs_runtime.Func(func(m1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m2_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_unsafeUnionWith(), compare_1_0, pkg_Data_Function.Get_go__const(), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](m1_2))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](m2_3))})))}
})
}))
}

func Call_plusMap(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
compare_1_1 := gopurs_runtime.RecordGet(dictOrd_0, "compare")
_ = compare_1_1
altMap1_1_0 := gopurs_runtime.RecordDict2("Functor0", "alt", gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_functorMap()
}), gopurs_runtime.Func(func(m1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m2_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_unsafeUnionWith(), compare_1_1, pkg_Data_Function.Get_go__const(), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](m1_2))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](m2_3))})))}
})
}))
_ = altMap1_1_0
return gopurs_runtime.RecordDict2("Alt0", "empty", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return altMap1_1_0
}), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(nil))})
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

func Call_eq__2843686287(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
_ = __eta1_1
return gopurs_runtime.Bool((__eta0_0.IntVal) == (__eta1_1.IntVal))
}

func Call_eq__2384498378(dict_0_loop *pkg_Data_Eq.Constructor_Eq[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Eq.Constructor_Eq[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_eq__2541178864(dict_0_loop *pkg_Data_Eq.Constructor_Eq[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Eq.Constructor_Eq[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_foldl__2151204251(dict_0_loop *pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_foldl__53736539(dict_0_loop *pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_foldl__22573083(dict_0_loop *pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_foldr__2151204251(dict_0_loop *pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V2
}

func Call_foldr__530094749(f_0_loop gopurs_runtime.Value, z_1_loop *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var z_1 *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = z_1_loop
_ = z_1
var go__go_2_0_36 gopurs_runtime.Value
_ = go__go_2_0_36
go__go_2_0_36 = gopurs_runtime.Func2(func(__local_var_3 gopurs_runtime.Value, __local_var_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if (__local_var_3.Type == 9 && __local_var_3.IntVal == 324739070 && __local_var_3.UnsafePtr == nil) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](__local_var_4))}
goto end_branch_1
} else {

}
}
{
if (__local_var_3.Type == 9 && __local_var_3.IntVal == 324739070 && __local_var_3.UnsafePtr != nil) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.UncurriedApp2(go__go_2_0_36, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V4)}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply2(f_0, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.UncurriedApp2(go__go_2_0_36, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V5)}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](__local_var_4))})))})))})))}
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](__t1))}
})
return gopurs_runtime.Func(func(m_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.UncurriedApp2(go__go_2_0_36, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](m_3))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(z_1)})))}
})
}

func Call_foldlWithIndex__2986161357(dict_0_loop *pkg_Data_FoldableWithIndex.Constructor_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_FoldableWithIndex.Constructor_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V2
}

func Call_foldlWithIndex__2499716749(dict_0_loop *pkg_Data_FoldableWithIndex.Constructor_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_FoldableWithIndex.Constructor_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V2
}

func Call_foldrWithIndex__2986161357(dict_0_loop *pkg_Data_FoldableWithIndex.Constructor_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_FoldableWithIndex.Constructor_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V3
}

func Call_foldrWithIndex__3511467915(f_0_loop gopurs_runtime.Value, z_1_loop *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var z_1 *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = z_1_loop
_ = z_1
var go__go_2_0_37 gopurs_runtime.Value
_ = go__go_2_0_37
go__go_2_0_37 = gopurs_runtime.Func2(func(__local_var_3 gopurs_runtime.Value, __local_var_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if (__local_var_3.Type == 9 && __local_var_3.IntVal == 324739070 && __local_var_3.UnsafePtr == nil) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](__local_var_4))}
goto end_branch_1
} else {

}
}
{
if (__local_var_3.Type == 9 && __local_var_3.IntVal == 324739070 && __local_var_3.UnsafePtr != nil) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.UncurriedApp2(go__go_2_0_37, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V4)}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply3(f_0, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V2, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.UncurriedApp2(go__go_2_0_37, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V5)}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](__local_var_4))})))})))})))}
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](__t1))}
})
return gopurs_runtime.Func(func(m_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.UncurriedApp2(go__go_2_0_37, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](m_3))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(z_1)})))}
})
}

func Call_const__1496134642(a_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var a_0 gopurs_runtime.Value = a_0_loop
_ = a_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return a_0
}

func Call_const__2390202835(a_0_loop bool, v_1_loop gopurs_runtime.Value) bool {
var a_0 bool = a_0_loop
_ = a_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return a_0
}

func Call_const__702735379(a_0_loop *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value], v_1_loop gopurs_runtime.Value) *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value] {
var a_0 *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value] = a_0_loop
_ = a_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return a_0
}

func Call_const__220790420(a_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var a_0 gopurs_runtime.Value = a_0_loop
_ = a_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return a_0
}

func Call_const__641934996(a_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var a_0 gopurs_runtime.Value = a_0_loop
_ = a_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return a_0
}

func Call_const__1628252324(a_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var a_0 gopurs_runtime.Value = a_0_loop
_ = a_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return a_0
}

func Call_flip__3525081280(f_0_loop gopurs_runtime.Value, b_1_loop gopurs_runtime.Value, a_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var b_1 gopurs_runtime.Value = b_1_loop
_ = b_1
var a_2 gopurs_runtime.Value = a_2_loop
_ = a_2
return gopurs_runtime.Apply2(f_0, a_2, b_1)
}

func Call_flip__1172723328(f_0_loop gopurs_runtime.Value, b_1_loop gopurs_runtime.Value, a_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var b_1 gopurs_runtime.Value = b_1_loop
_ = b_1
var a_2 gopurs_runtime.Value = a_2_loop
_ = a_2
return gopurs_runtime.Apply2(f_0, a_2, b_1)
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

func Call_map__4039429788(v_0_loop gopurs_runtime.Value, v1_1_loop *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]) *pkg_Data_Maybe.Constructor_Just[*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, *Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]]] {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var v1_1 *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value] = v1_1_loop
_ = v1_1
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v1_1)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v1_1)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v1_1)}.UnsafePtr != nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, *Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]]](gopurs_runtime.Apply(v_0, (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v1_1)}.UnsafePtr).V0)))}})}
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(nil))}
}
end_branch_0:
return gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, *Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]]]](__t0)
}

func Call_conj__3676519832(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
_ = __eta1_1
return gopurs_runtime.Bool(((__eta0_0.IntVal) != (0)) && ((__eta1_1.IntVal) != (0)))
}

func Call_conj__3472268504(dict_0_loop *pkg_Data_HeytingAlgebra.Constructor_HeytingAlgebra[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_HeytingAlgebra.Constructor_HeytingAlgebra[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_disj__3676519832(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
_ = __eta1_1
return gopurs_runtime.Bool(((__eta0_0.IntVal) != (0)) || ((__eta1_1.IntVal) != (0)))
}

func Call_disj__3472268504(dict_0_loop *pkg_Data_HeytingAlgebra.Constructor_HeytingAlgebra[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_HeytingAlgebra.Constructor_HeytingAlgebra[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_not__3201284355(__eta0_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
return gopurs_runtime.Bool(((__eta0_0.IntVal) != (0)) != (true))
}

func Call_not__1505204753(dict_0_loop *pkg_Data_HeytingAlgebra.Constructor_HeytingAlgebra[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_HeytingAlgebra.Constructor_HeytingAlgebra[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V4
}

func Call_findMax__528468393(v_0_loop *Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]) *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value] {
var v_0 *Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value] = v_0_loop
_ = v_0
var __t2 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_0)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_0)}.IntVal == 324739070 && gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr == nil) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(nil))}
goto end_branch_2
} else {

}
}
{
if (gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_0)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_0)}.IntVal == 324739070 && gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr != nil) {
var __t1 gopurs_runtime.Value
{
var __t_tag_0 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V5)}
if (__t_tag_0.Type == 9 && __t_tag_0.IntVal == 324739070 && __t_tag_0.UnsafePtr == nil) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, gopurs_runtime.RecordDict2("key", "value", (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V2, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V3)})}
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_findMax((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V5))}))}
}
end_branch_1:
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](__t1))}
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_2:
return gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](__t2)
}

func Call_findMin__528468393(v_0_loop *Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]) *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value] {
var v_0 *Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value] = v_0_loop
_ = v_0
var __t2 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_0)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_0)}.IntVal == 324739070 && gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr == nil) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(nil))}
goto end_branch_2
} else {

}
}
{
if (gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_0)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_0)}.IntVal == 324739070 && gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr != nil) {
var __t1 gopurs_runtime.Value
{
var __t_tag_0 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V4)}
if (__t_tag_0.Type == 9 && __t_tag_0.IntVal == 324739070 && __t_tag_0.UnsafePtr == nil) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, gopurs_runtime.RecordDict2("key", "value", (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V2, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V3)})}
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_findMin((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V4))}))}
}
end_branch_1:
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](__t1))}
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_2:
return gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](__t2)
}

func Call_foldSubmapBy__3050108409(dictOrd_0_loop *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value], appendFn_1_loop gopurs_runtime.Value, memptyValue_2_loop gopurs_runtime.Value, kmin_3_loop *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value], kmax_4_loop *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value], f_5_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value] = dictOrd_0_loop
_ = dictOrd_0
var appendFn_1 gopurs_runtime.Value = appendFn_1_loop
_ = appendFn_1
var memptyValue_2 gopurs_runtime.Value = memptyValue_2_loop
_ = memptyValue_2
var kmin_3 *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value] = kmin_3_loop
_ = kmin_3
var kmax_4 *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value] = kmax_4_loop
_ = kmax_4
var f_5 gopurs_runtime.Value = f_5_loop
_ = f_5
var __t4 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(kmin_3)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(kmin_3)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(kmin_3)}.UnsafePtr != nil) {
__local_var_6_1 := (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(kmin_3)}.UnsafePtr).V0
_ = __local_var_6_1
__t4 = gopurs_runtime.Func(func(k_7 gopurs_runtime.Value) gopurs_runtime.Value {
var __t3 bool
{
var __t_tag_2 gopurs_runtime.Value = gopurs_runtime.Apply2(dictOrd_0.V1, k_7, __local_var_6_1)
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
tooSmall_6_0 := __t4
_ = tooSmall_6_0
var __t9 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(kmax_4)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(kmax_4)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(kmax_4)}.UnsafePtr != nil) {
__local_var_7_6 := (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(kmax_4)}.UnsafePtr).V0
_ = __local_var_7_6
__t9 = gopurs_runtime.Func(func(k_8 gopurs_runtime.Value) gopurs_runtime.Value {
var __t8 bool
{
var __t_tag_7 gopurs_runtime.Value = gopurs_runtime.Apply2(dictOrd_0.V1, k_8, __local_var_7_6)
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
tooLarge_7_5 := __t9
_ = tooLarge_7_5
var __t26 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(kmin_3)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(kmin_3)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(kmin_3)}.UnsafePtr != nil) {
var __t21 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(kmax_4)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(kmax_4)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(kmax_4)}.UnsafePtr != nil) {
__local_var_8_11 := (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(kmax_4)}.UnsafePtr).V0
_ = __local_var_8_11
__local_var_9_12 := (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(kmin_3)}.UnsafePtr).V0
_ = __local_var_9_12
__t21 = gopurs_runtime.Func(func(k_10 gopurs_runtime.Value) gopurs_runtime.Value {
var __t14 bool
{
var __t_tag_13 gopurs_runtime.Value = gopurs_runtime.Apply2(dictOrd_0.V1, __local_var_9_12, k_10)
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
var __t_tag_15 gopurs_runtime.Value = gopurs_runtime.Apply2(dictOrd_0.V1, k_10, __local_var_8_11)
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
__local_var_8_18 := (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(kmin_3)}.UnsafePtr).V0
_ = __local_var_8_18
__t21 = gopurs_runtime.Func(func(k_9 gopurs_runtime.Value) gopurs_runtime.Value {
var __t20 bool
{
var __t_tag_19 gopurs_runtime.Value = gopurs_runtime.Apply2(dictOrd_0.V1, __local_var_8_18, k_9)
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
__local_var_8_22 := (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(kmax_4)}.UnsafePtr).V0
_ = __local_var_8_22
__t25 = gopurs_runtime.Func(func(k_9 gopurs_runtime.Value) gopurs_runtime.Value {
var __t24 bool
{
var __t_tag_23 gopurs_runtime.Value = gopurs_runtime.Apply2(dictOrd_0.V1, k_9, __local_var_8_22)
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
inBounds_8_10 := __t26
_ = inBounds_8_10
var go__go_9_27_38 gopurs_runtime.Value
_ = go__go_9_27_38
go__go_9_27_38 = gopurs_runtime.Func(func(v_10 gopurs_runtime.Value) gopurs_runtime.Value {
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
if (gopurs_runtime.Apply(tooSmall_6_0, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_10.UnsafePtr).V2).IntVal) != (0) {
__t28 = memptyValue_2
goto end_branch_28
} else {

}
}
{
__t28 = gopurs_runtime.Apply(go__go_9_27_38, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_10.UnsafePtr).V4)})
}
end_branch_28:
var __t29 gopurs_runtime.Value
{
if (gopurs_runtime.Apply(inBounds_8_10, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_10.UnsafePtr).V2).IntVal) != (0) {
__t29 = gopurs_runtime.Apply2(f_5, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_10.UnsafePtr).V2, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_10.UnsafePtr).V3)
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
if (gopurs_runtime.Apply(tooLarge_7_5, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_10.UnsafePtr).V2).IntVal) != (0) {
__t30 = memptyValue_2
goto end_branch_30
} else {

}
}
{
__t30 = gopurs_runtime.Apply(go__go_9_27_38, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_10.UnsafePtr).V5)})
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
return go__go_9_27_38
}

func Call_foldSubmapBy__3128450809(dictOrd_0_loop *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value], appendFn_1_loop gopurs_runtime.Value, memptyValue_2_loop gopurs_runtime.Value, kmin_3_loop *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value], kmax_4_loop *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value], f_5_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value] = dictOrd_0_loop
_ = dictOrd_0
var appendFn_1 gopurs_runtime.Value = appendFn_1_loop
_ = appendFn_1
var memptyValue_2 gopurs_runtime.Value = memptyValue_2_loop
_ = memptyValue_2
var kmin_3 *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value] = kmin_3_loop
_ = kmin_3
var kmax_4 *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value] = kmax_4_loop
_ = kmax_4
var f_5 gopurs_runtime.Value = f_5_loop
_ = f_5
var __t4 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(kmin_3)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(kmin_3)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(kmin_3)}.UnsafePtr != nil) {
__local_var_6_1 := (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(kmin_3)}.UnsafePtr).V0
_ = __local_var_6_1
__t4 = gopurs_runtime.Func(func(k_7 gopurs_runtime.Value) gopurs_runtime.Value {
var __t3 bool
{
var __t_tag_2 gopurs_runtime.Value = gopurs_runtime.Apply2(dictOrd_0.V1, k_7, __local_var_6_1)
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
tooSmall_6_0 := __t4
_ = tooSmall_6_0
var __t9 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(kmax_4)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(kmax_4)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(kmax_4)}.UnsafePtr != nil) {
__local_var_7_6 := (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(kmax_4)}.UnsafePtr).V0
_ = __local_var_7_6
__t9 = gopurs_runtime.Func(func(k_8 gopurs_runtime.Value) gopurs_runtime.Value {
var __t8 bool
{
var __t_tag_7 gopurs_runtime.Value = gopurs_runtime.Apply2(dictOrd_0.V1, k_8, __local_var_7_6)
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
tooLarge_7_5 := __t9
_ = tooLarge_7_5
var __t26 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(kmin_3)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(kmin_3)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(kmin_3)}.UnsafePtr != nil) {
var __t21 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(kmax_4)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(kmax_4)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(kmax_4)}.UnsafePtr != nil) {
__local_var_8_11 := (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(kmax_4)}.UnsafePtr).V0
_ = __local_var_8_11
__local_var_9_12 := (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(kmin_3)}.UnsafePtr).V0
_ = __local_var_9_12
__t21 = gopurs_runtime.Func(func(k_10 gopurs_runtime.Value) gopurs_runtime.Value {
var __t14 bool
{
var __t_tag_13 gopurs_runtime.Value = gopurs_runtime.Apply2(dictOrd_0.V1, __local_var_9_12, k_10)
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
var __t_tag_15 gopurs_runtime.Value = gopurs_runtime.Apply2(dictOrd_0.V1, k_10, __local_var_8_11)
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
__local_var_8_18 := (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(kmin_3)}.UnsafePtr).V0
_ = __local_var_8_18
__t21 = gopurs_runtime.Func(func(k_9 gopurs_runtime.Value) gopurs_runtime.Value {
var __t20 bool
{
var __t_tag_19 gopurs_runtime.Value = gopurs_runtime.Apply2(dictOrd_0.V1, __local_var_8_18, k_9)
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
__local_var_8_22 := (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(kmax_4)}.UnsafePtr).V0
_ = __local_var_8_22
__t25 = gopurs_runtime.Func(func(k_9 gopurs_runtime.Value) gopurs_runtime.Value {
var __t24 bool
{
var __t_tag_23 gopurs_runtime.Value = gopurs_runtime.Apply2(dictOrd_0.V1, k_9, __local_var_8_22)
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
inBounds_8_10 := __t26
_ = inBounds_8_10
var go__go_9_27_39 gopurs_runtime.Value
_ = go__go_9_27_39
go__go_9_27_39 = gopurs_runtime.Func(func(v_10 gopurs_runtime.Value) gopurs_runtime.Value {
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
if (gopurs_runtime.Apply(tooSmall_6_0, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_10.UnsafePtr).V2).IntVal) != (0) {
__t28 = memptyValue_2
goto end_branch_28
} else {

}
}
{
__t28 = gopurs_runtime.Apply(go__go_9_27_39, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_10.UnsafePtr).V4)})
}
end_branch_28:
var __t29 gopurs_runtime.Value
{
if (gopurs_runtime.Apply(inBounds_8_10, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_10.UnsafePtr).V2).IntVal) != (0) {
__t29 = gopurs_runtime.Apply2(f_5, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_10.UnsafePtr).V2, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_10.UnsafePtr).V3)
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
if (gopurs_runtime.Apply(tooLarge_7_5, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_10.UnsafePtr).V2).IntVal) != (0) {
__t30 = memptyValue_2
goto end_branch_30
} else {

}
}
{
__t30 = gopurs_runtime.Apply(go__go_9_27_39, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_10.UnsafePtr).V5)})
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
return go__go_9_27_39
}

func Call_insert__4289641298(dictOrd_0_loop *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value], k_1_loop gopurs_runtime.Value, v_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value] = dictOrd_0_loop
_ = dictOrd_0
var k_1 gopurs_runtime.Value = k_1_loop
_ = k_1
var v_2 gopurs_runtime.Value = v_2_loop
_ = v_2
var go__go_3_0_54 gopurs_runtime.Value
_ = go__go_3_0_54
go__go_3_0_54 = gopurs_runtime.Func(func(v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t3 gopurs_runtime.Value
{
if (v1_4.Type == 9 && v1_4.IntVal == 324739070 && v1_4.UnsafePtr == nil) {
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(&Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]{1, 1, 1, k_1, v_2, gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(nil))}), gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(nil))})})}
goto end_branch_3
} else {

}
}
{
if (v1_4.Type == 9 && v1_4.IntVal == 324739070 && v1_4.UnsafePtr != nil) {
v2_5_1 := gopurs_runtime.Apply2(dictOrd_0.V1, k_1, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v1_4.UnsafePtr).V2)
_ = v2_5_1
var __t2 gopurs_runtime.Value
{
if (uint32(v2_5_1.IntVal) == 1527465420) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_unsafeBalancedNode(), (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v1_4.UnsafePtr).V2, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v1_4.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(go__go_3_0_54, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v1_4.UnsafePtr).V4)})))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v1_4.UnsafePtr).V5)})))}
goto end_branch_2
} else {

}
}
{
if (uint32(v2_5_1.IntVal) == 380165415) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_unsafeBalancedNode(), (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v1_4.UnsafePtr).V2, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v1_4.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v1_4.UnsafePtr).V4)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(go__go_3_0_54, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v1_4.UnsafePtr).V5)})))})))}
goto end_branch_2
} else {

}
}
{
if (uint32(v2_5_1.IntVal) == 902936544) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(&Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v1_4.UnsafePtr).V0, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v1_4.UnsafePtr).V1, k_1, v_2, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v1_4.UnsafePtr).V4, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v1_4.UnsafePtr).V5})}
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_2:
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__t2))}
goto end_branch_3
} else {

}
}
{
__t3 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_3:
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__t3))}
})
return go__go_3_0_54
}

func Call_insertWith__118979962(dictOrd_0_loop *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value], app_1_loop gopurs_runtime.Value, k_2_loop gopurs_runtime.Value, v_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value] = dictOrd_0_loop
_ = dictOrd_0
var app_1 gopurs_runtime.Value = app_1_loop
_ = app_1
var k_2 gopurs_runtime.Value = k_2_loop
_ = k_2
var v_3 gopurs_runtime.Value = v_3_loop
_ = v_3
var go__go_4_0_55 gopurs_runtime.Value
_ = go__go_4_0_55
go__go_4_0_55 = gopurs_runtime.Func(func(v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
var __t3 gopurs_runtime.Value
{
if (v1_5.Type == 9 && v1_5.IntVal == 324739070 && v1_5.UnsafePtr == nil) {
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(&Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]{1, 1, 1, k_2, v_3, gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(nil))}), gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(nil))})})}
goto end_branch_3
} else {

}
}
{
if (v1_5.Type == 9 && v1_5.IntVal == 324739070 && v1_5.UnsafePtr != nil) {
v2_6_1 := gopurs_runtime.Apply2(dictOrd_0.V1, k_2, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v1_5.UnsafePtr).V2)
_ = v2_6_1
var __t2 gopurs_runtime.Value
{
if (uint32(v2_6_1.IntVal) == 1527465420) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_unsafeBalancedNode(), (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v1_5.UnsafePtr).V2, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v1_5.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(go__go_4_0_55, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v1_5.UnsafePtr).V4)})))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v1_5.UnsafePtr).V5)})))}
goto end_branch_2
} else {

}
}
{
if (uint32(v2_6_1.IntVal) == 380165415) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_unsafeBalancedNode(), (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v1_5.UnsafePtr).V2, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v1_5.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v1_5.UnsafePtr).V4)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(go__go_4_0_55, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v1_5.UnsafePtr).V5)})))})))}
goto end_branch_2
} else {

}
}
{
if (uint32(v2_6_1.IntVal) == 902936544) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(&Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v1_5.UnsafePtr).V0, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v1_5.UnsafePtr).V1, k_2, gopurs_runtime.Apply2(app_1, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v1_5.UnsafePtr).V3, v_3), (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v1_5.UnsafePtr).V4, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v1_5.UnsafePtr).V5})}
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_2:
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__t2))}
goto end_branch_3
} else {

}
}
{
__t3 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_3:
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__t3))}
})
return go__go_4_0_55
}

func Call_intersectionWith__3717755541(dictOrd_0_loop *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictOrd_0 *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value] = dictOrd_0_loop
_ = dictOrd_0
compare_1_0 := dictOrd_0.V1
_ = compare_1_0
return gopurs_runtime.Func(func(app_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m2_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_unsafeIntersectionWith(), compare_1_0, app_2, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](m1_3))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](m2_4))})))}
})
})
})
}

func Call_intersectionWith__4144106805(dictOrd_0_loop *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictOrd_0 *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value] = dictOrd_0_loop
_ = dictOrd_0
compare_1_0 := dictOrd_0.V1
_ = compare_1_0
return gopurs_runtime.Func(func(app_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m2_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_unsafeIntersectionWith(), compare_1_0, app_2, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](m1_3))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](m2_4))})))}
})
})
})
}

func Call_iterMapU__878452066(iter_0_loop gopurs_runtime.Value, v_1_loop *Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var iter_0 gopurs_runtime.Value = iter_0_loop
_ = iter_0
var v_1 *Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value] = v_1_loop
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
var __t5 *Constructor_IterEmit[gopurs_runtime.Value, gopurs_runtime.Value]
{
var __t_tag_2 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_1)}.UnsafePtr).V4)}
if (__t_tag_2.Type == 9 && __t_tag_2.IntVal == 324739070 && __t_tag_2.UnsafePtr == nil) {
var __t4 *Constructor_IterEmit[gopurs_runtime.Value, gopurs_runtime.Value]
{
var __t_tag_3 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_1)}.UnsafePtr).V5)}
if (__t_tag_3.Type == 9 && __t_tag_3.IntVal == 324739070 && __t_tag_3.UnsafePtr == nil) {
__t4 = &Constructor_IterEmit[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_1)}.UnsafePtr).V2, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_1)}.UnsafePtr).V3, iter_0}
goto end_branch_4
} else {

}
}
{
__t4 = &Constructor_IterEmit[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_1)}.UnsafePtr).V2, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_1)}.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 2861335956, UnsafePtr: unsafe.Pointer(&Constructor_IterNode[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_1)}.UnsafePtr).V5, iter_0})}}
}
end_branch_4:
__t5 = __t4
goto end_branch_5
} else {

}
}
{
var __t1 *Constructor_IterEmit[gopurs_runtime.Value, gopurs_runtime.Value]
{
var __t_tag_0 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_1)}.UnsafePtr).V5)}
if (__t_tag_0.Type == 9 && __t_tag_0.IntVal == 324739070 && __t_tag_0.UnsafePtr == nil) {
__t1 = &Constructor_IterEmit[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_1)}.UnsafePtr).V2, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_1)}.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 2861335956, UnsafePtr: unsafe.Pointer(&Constructor_IterNode[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_1)}.UnsafePtr).V4, iter_0})}}
goto end_branch_1
} else {

}
}
{
__t1 = &Constructor_IterEmit[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_1)}.UnsafePtr).V2, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_1)}.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 2861335956, UnsafePtr: unsafe.Pointer(&Constructor_IterNode[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_1)}.UnsafePtr).V4, gopurs_runtime.Value{Type: 9, IntVal: 2861335956, UnsafePtr: unsafe.Pointer(&Constructor_IterNode[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_1)}.UnsafePtr).V5, iter_0})}})}}
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

func Call_lookup__3378638282(dictOrd_0_loop *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value], k_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value] = dictOrd_0_loop
_ = dictOrd_0
var k_1 gopurs_runtime.Value = k_1_loop
_ = k_1
var go__go_2_0_58 gopurs_runtime.Value
go__go_2_0_58 = gopurs_runtime.Func(func(v_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_3_loop gopurs_runtime.Value = v_3_loop_val
go__go_2_0_58:
for {
if false { continue go__go_2_0_58 }
var v_3 gopurs_runtime.Value = v_3_loop
_ = v_3
var __t3 gopurs_runtime.Value
{
if (v_3.Type == 9 && v_3.IntVal == 324739070 && v_3.UnsafePtr == nil) {
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(nil))}
goto end_branch_3
} else {

}
}
{
if (v_3.Type == 9 && v_3.IntVal == 324739070 && v_3.UnsafePtr != nil) {
v1_4_1 := gopurs_runtime.Apply2(dictOrd_0.V1, k_1, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V2)
_ = v1_4_1
var __t2 gopurs_runtime.Value
{
if (uint32(v1_4_1.IntVal) == 1527465420) {
v_3_loop = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V4)}
continue go__go_2_0_58
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](gopurs_runtime.Value{}))}
goto end_branch_2
} else {

}
}
{
if (uint32(v1_4_1.IntVal) == 380165415) {
v_3_loop = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V5)}
continue go__go_2_0_58
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](gopurs_runtime.Value{}))}
goto end_branch_2
} else {

}
}
{
if (uint32(v1_4_1.IntVal) == 902936544) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V3})}
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_2:
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](__t2))}
goto end_branch_3
} else {

}
}
{
__t3 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_3:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](__t3))}
}
}()
})
return go__go_2_0_58
}

func Call_mapMaybe__3426301240(dictOrd_0_loop *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value], x_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value] = dictOrd_0_loop
_ = dictOrd_0
var x_1 gopurs_runtime.Value = x_1_loop
_ = x_1
return Call_mapMaybeWithKey(dictOrd_0, gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return x_1
}))
}

func Call_mapMaybe__1970555288(dictOrd_0_loop *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value], x_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value] = dictOrd_0_loop
_ = dictOrd_0
var x_1 gopurs_runtime.Value = x_1_loop
_ = x_1
return Call_mapMaybeWithKey(dictOrd_0, gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return x_1
}))
}

func Call_mapMaybeWithKey__817660689(dictOrd_0_loop *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value], f_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value] = dictOrd_0_loop
_ = dictOrd_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var go__go_2_0_59 gopurs_runtime.Value
_ = go__go_2_0_59
go__go_2_0_59 = gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t3 gopurs_runtime.Value
{
if (v_3.Type == 9 && v_3.IntVal == 324739070 && v_3.UnsafePtr == nil) {
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(nil))}
goto end_branch_3
} else {

}
}
{
if (v_3.Type == 9 && v_3.IntVal == 324739070 && v_3.UnsafePtr != nil) {
v2_4_1 := gopurs_runtime.Apply2(f_1, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V2, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V3)
_ = v2_4_1
var __t2 gopurs_runtime.Value
{
if (v2_4_1.Type == 9 && v2_4_1.IntVal == 930809136 && v2_4_1.UnsafePtr != nil) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_unsafeBalancedNode(), (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V2, (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(v2_4_1.UnsafePtr).V0, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(go__go_2_0_59, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V4)})))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(go__go_2_0_59, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V5)})))})))}
goto end_branch_2
} else {

}
}
{
if (v2_4_1.Type == 9 && v2_4_1.IntVal == 930809136 && v2_4_1.UnsafePtr == nil) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp2(Get_unsafeJoinNodes(), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(go__go_2_0_59, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V4)})))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(go__go_2_0_59, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V5)})))})))}
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_2:
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__t2))}
goto end_branch_3
} else {

}
}
{
__t3 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_3:
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__t3))}
})
return go__go_2_0_59
}

func Call_singleton__3511563426(k_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) *Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value] {
var k_0 gopurs_runtime.Value = k_0_loop
_ = k_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(&Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]{1, 1, 1, k_0, v_1, gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(nil))}), gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(nil))})})})
}

func Call_singleton__2450056090(k_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) *Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value] {
var k_0 gopurs_runtime.Value = k_0_loop
_ = k_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(&Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]{1, 1, 1, k_0, v_1, gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(nil))}), gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(nil))})})})
}

func Call_stepWith__3186376421(f_0_loop gopurs_runtime.Value, next_1_loop gopurs_runtime.Value, done_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var next_1 gopurs_runtime.Value = next_1_loop
_ = next_1
var done_2 gopurs_runtime.Value = done_2_loop
_ = done_2
var go__go_3_0_60 gopurs_runtime.Value
go__go_3_0_60 = gopurs_runtime.Func(func(v_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_4_loop gopurs_runtime.Value = v_4_loop_val
go__go_3_0_60:
for {
if false { continue go__go_3_0_60 }
var v_4 gopurs_runtime.Value = v_4_loop
_ = v_4
var __t1 gopurs_runtime.Value
{
if (v_4.Type == 9 && v_4.IntVal == 2509360378) {
__t1 = gopurs_runtime.Apply(done_2, pkg_Data_Unit.Get_unit())
goto end_branch_1
} else {

}
}
{
if (v_4.Type == 9 && v_4.IntVal == 1343415489) {
__t1 = gopurs_runtime.UncurriedApp3(next_1, (*Constructor_IterEmit[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V0, (*Constructor_IterEmit[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V1, (*Constructor_IterEmit[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V2)
goto end_branch_1
} else {

}
}
{
if (v_4.Type == 9 && v_4.IntVal == 2861335956) {
v_4_loop = gopurs_runtime.Apply2(f_0, (*Constructor_IterNode[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_IterNode[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V0)})
continue go__go_3_0_60
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
return go__go_3_0_60
}

func Call_toMapIter__2014410513(a_0_loop *Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var a_0 *Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value] = a_0_loop
_ = a_0
return gopurs_runtime.Value{Type: 9, IntVal: 2861335956, UnsafePtr: unsafe.Pointer(&Constructor_IterNode[gopurs_runtime.Value, gopurs_runtime.Value]{1, a_0, gopurs_runtime.Value{Type: 9, IntVal: 2509360378, UnsafePtr: unsafe.Pointer(nil)}})}
}

func Call_toMapIter__1738891721(a_0_loop *Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var a_0 *Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value] = a_0_loop
_ = a_0
return gopurs_runtime.Value{Type: 9, IntVal: 2861335956, UnsafePtr: unsafe.Pointer(&Constructor_IterNode[gopurs_runtime.Value, gopurs_runtime.Value]{1, a_0, gopurs_runtime.Value{Type: 9, IntVal: 2509360378, UnsafePtr: unsafe.Pointer(nil)}})}
}

func Call_toUnfoldable__2183602684(dictUnfoldable_0_loop *pkg_Data_Unfoldable.Constructor_Unfoldable[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictUnfoldable_0 *pkg_Data_Unfoldable.Constructor_Unfoldable[gopurs_runtime.Value] = dictUnfoldable_0_loop
_ = dictUnfoldable_0
__local_var_1_0 := gopurs_runtime.Apply(dictUnfoldable_0.V1, Get_stepUnfoldr())
_ = __local_var_1_0
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_1_0, gopurs_runtime.Value{Type: 9, IntVal: 2861335956, UnsafePtr: unsafe.Pointer(&Constructor_IterNode[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](x_2), gopurs_runtime.Value{Type: 9, IntVal: 2509360378, UnsafePtr: unsafe.Pointer(nil)}})})
})
}

func Call_unionWith__2507192643(dictOrd_0_loop *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictOrd_0 *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value] = dictOrd_0_loop
_ = dictOrd_0
compare_1_0 := dictOrd_0.V1
_ = compare_1_0
return gopurs_runtime.Func(func(app_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m2_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_unsafeUnionWith(), compare_1_0, app_2, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](m1_3))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](m2_4))})))}
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
__t6 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(&Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]{1, 1, 1, __local_var_0, __local_var_1, gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(nil))}), gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(nil))})})}
goto end_branch_6
} else {

}
}
{
if ((__local_var_3.Type == 9 && __local_var_3.IntVal == 324739070 && __local_var_3.UnsafePtr != nil)) && ((gopurs_runtime.Bool(Call_greaterThan__4087042607(gopurs_runtime.Int((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V0), gopurs_runtime.Int(1))).IntVal) != (0)) {
var __t5 *Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]
{
var __t_tag_0 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V4)}
var __t_and_4 bool = false
if (__t_tag_0.Type == 9 && __t_tag_0.IntVal == 324739070 && __t_tag_0.UnsafePtr != nil) {

var __t3 gopurs_runtime.Value
{
var __t_tag_1 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V5)}
if (__t_tag_1.Type == 9 && __t_tag_1.IntVal == 324739070 && __t_tag_1.UnsafePtr == nil) {
__t3 = gopurs_runtime.Int(0)
goto end_branch_3
} else {

}
}
{
var __t_tag_2 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V5)}
if (__t_tag_2.Type == 9 && __t_tag_2.IntVal == 324739070 && __t_tag_2.UnsafePtr != nil) {
__t3 = gopurs_runtime.Int((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V5)}.UnsafePtr).V0)
goto end_branch_3
} else {

}
}
{
__t3 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_3:
__t_and_4 = (gopurs_runtime.Bool(Call_greaterThan__4087042607(gopurs_runtime.Int((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V4)}.UnsafePtr).V0), gopurs_runtime.Int(__t3.IntVal))).IntVal) != (0)
}
if __t_and_4 {
__t5 = gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_unsafeNode(), (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V4)}.UnsafePtr).V2, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V4)}.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_unsafeNode(), __local_var_0, __local_var_1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__local_var_2))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V4)}.UnsafePtr).V4)})))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_unsafeNode(), (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V2, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V4)}.UnsafePtr).V5)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V5)})))}))
goto end_branch_5
} else {

}
}
{
__t5 = gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_unsafeNode(), (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V2, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_unsafeNode(), __local_var_0, __local_var_1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__local_var_2))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V4)})))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V5)}))
}
end_branch_5:
__t6 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(__t5)}
goto end_branch_6
} else {

}
}
{
__t6 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_unsafeNode(), __local_var_0, __local_var_1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__local_var_2))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__local_var_3))})))}
}
end_branch_6:
__t27 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__t6))}
goto end_branch_27
} else {

}
}
{
if (__local_var_2.Type == 9 && __local_var_2.IntVal == 324739070 && __local_var_2.UnsafePtr != nil) {
var __t26 *Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]
{
if (__local_var_3.Type == 9 && __local_var_3.IntVal == 324739070 && __local_var_3.UnsafePtr != nil) {
var __t19 *Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]
{
if (gopurs_runtime.Bool(Call_greaterThan__4087042607(gopurs_runtime.Int((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V0), gopurs_runtime.Int(Call_add__560788792(gopurs_runtime.Int((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V0), gopurs_runtime.Int(1)).IntVal))).IntVal) != (0) {
var __t12 *Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]
{
var __t_tag_7 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V4)}
var __t_and_11 bool = false
if (__t_tag_7.Type == 9 && __t_tag_7.IntVal == 324739070 && __t_tag_7.UnsafePtr != nil) {

var __t10 gopurs_runtime.Value
{
var __t_tag_8 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V5)}
if (__t_tag_8.Type == 9 && __t_tag_8.IntVal == 324739070 && __t_tag_8.UnsafePtr == nil) {
__t10 = gopurs_runtime.Int(0)
goto end_branch_10
} else {

}
}
{
var __t_tag_9 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V5)}
if (__t_tag_9.Type == 9 && __t_tag_9.IntVal == 324739070 && __t_tag_9.UnsafePtr != nil) {
__t10 = gopurs_runtime.Int((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V5)}.UnsafePtr).V0)
goto end_branch_10
} else {

}
}
{
__t10 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_10:
__t_and_11 = (gopurs_runtime.Bool(Call_greaterThan__4087042607(gopurs_runtime.Int((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V4)}.UnsafePtr).V0), gopurs_runtime.Int(__t10.IntVal))).IntVal) != (0)
}
if __t_and_11 {
__t12 = gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_unsafeNode(), (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V4)}.UnsafePtr).V2, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V4)}.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_unsafeNode(), __local_var_0, __local_var_1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__local_var_2))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V4)}.UnsafePtr).V4)})))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_unsafeNode(), (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V2, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V4)}.UnsafePtr).V5)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V5)})))}))
goto end_branch_12
} else {

}
}
{
__t12 = gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_unsafeNode(), (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V2, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_unsafeNode(), __local_var_0, __local_var_1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__local_var_2))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V4)})))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V5)}))
}
end_branch_12:
__t19 = __t12
goto end_branch_19
} else {

}
}
{
if (gopurs_runtime.Bool(Call_greaterThan__4087042607(gopurs_runtime.Int((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V0), gopurs_runtime.Int(Call_add__560788792(gopurs_runtime.Int((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V0), gopurs_runtime.Int(1)).IntVal))).IntVal) != (0) {
var __t18 *Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]
{
var __t_tag_13 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V5)}
var __t_and_17 bool = false
if (__t_tag_13.Type == 9 && __t_tag_13.IntVal == 324739070 && __t_tag_13.UnsafePtr != nil) {

var __t16 gopurs_runtime.Value
{
var __t_tag_14 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V4)}
if (__t_tag_14.Type == 9 && __t_tag_14.IntVal == 324739070 && __t_tag_14.UnsafePtr == nil) {
__t16 = gopurs_runtime.Int(0)
goto end_branch_16
} else {

}
}
{
var __t_tag_15 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V4)}
if (__t_tag_15.Type == 9 && __t_tag_15.IntVal == 324739070 && __t_tag_15.UnsafePtr != nil) {
__t16 = gopurs_runtime.Int((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V4)}.UnsafePtr).V0)
goto end_branch_16
} else {

}
}
{
__t16 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_16:
__t_and_17 = (gopurs_runtime.Bool(Call_lessThanOrEq__4087042607(gopurs_runtime.Int(__t16.IntVal), gopurs_runtime.Int((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V5)}.UnsafePtr).V0))).IntVal) != (0)
}
if __t_and_17 {
__t18 = gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_unsafeNode(), (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V5)}.UnsafePtr).V2, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V5)}.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_unsafeNode(), (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V2, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V4)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V5)}.UnsafePtr).V4)})))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_unsafeNode(), __local_var_0, __local_var_1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V5)}.UnsafePtr).V5)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__local_var_3))})))}))
goto end_branch_18
} else {

}
}
{
__t18 = gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_unsafeNode(), (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V2, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V4)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_unsafeNode(), __local_var_0, __local_var_1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V5)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__local_var_3))})))}))
}
end_branch_18:
__t19 = __t18
goto end_branch_19
} else {

}
}
{
__t19 = gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_unsafeNode(), __local_var_0, __local_var_1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__local_var_2))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__local_var_3))}))
}
end_branch_19:
__t26 = __t19
goto end_branch_26
} else {

}
}
{
if ((__local_var_3.Type == 9 && __local_var_3.IntVal == 324739070 && __local_var_3.UnsafePtr == nil)) && ((gopurs_runtime.Bool(Call_greaterThan__4087042607(gopurs_runtime.Int((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V0), gopurs_runtime.Int(1))).IntVal) != (0)) {
var __t25 *Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]
{
var __t_tag_20 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V5)}
var __t_and_24 bool = false
if (__t_tag_20.Type == 9 && __t_tag_20.IntVal == 324739070 && __t_tag_20.UnsafePtr != nil) {

var __t23 gopurs_runtime.Value
{
var __t_tag_21 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V4)}
if (__t_tag_21.Type == 9 && __t_tag_21.IntVal == 324739070 && __t_tag_21.UnsafePtr == nil) {
__t23 = gopurs_runtime.Int(0)
goto end_branch_23
} else {

}
}
{
var __t_tag_22 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V4)}
if (__t_tag_22.Type == 9 && __t_tag_22.IntVal == 324739070 && __t_tag_22.UnsafePtr != nil) {
__t23 = gopurs_runtime.Int((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V4)}.UnsafePtr).V0)
goto end_branch_23
} else {

}
}
{
__t23 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_23:
__t_and_24 = (gopurs_runtime.Bool(Call_lessThanOrEq__4087042607(gopurs_runtime.Int(__t23.IntVal), gopurs_runtime.Int((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V5)}.UnsafePtr).V0))).IntVal) != (0)
}
if __t_and_24 {
__t25 = gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_unsafeNode(), (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V5)}.UnsafePtr).V2, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V5)}.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_unsafeNode(), (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V2, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V4)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V5)}.UnsafePtr).V4)})))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_unsafeNode(), __local_var_0, __local_var_1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V5)}.UnsafePtr).V5)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__local_var_3))})))}))
goto end_branch_25
} else {

}
}
{
__t25 = gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_unsafeNode(), (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V2, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V4)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_unsafeNode(), __local_var_0, __local_var_1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V5)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__local_var_3))})))}))
}
end_branch_25:
__t26 = __t25
goto end_branch_26
} else {

}
}
{
__t26 = gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_unsafeNode(), __local_var_0, __local_var_1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__local_var_2))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__local_var_3))}))
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
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__t27))}
}

func Call_unsafeDifference__4097927905(__local_var_0_loop gopurs_runtime.Value, __local_var_1_loop gopurs_runtime.Value, __local_var_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __local_var_0 gopurs_runtime.Value = __local_var_0_loop
_ = __local_var_0
var __local_var_1 gopurs_runtime.Value = __local_var_1_loop
_ = __local_var_1
var __local_var_2 gopurs_runtime.Value = __local_var_2_loop
_ = __local_var_2
var __t1 gopurs_runtime.Value
{
if (__local_var_1.Type == 9 && __local_var_1.IntVal == 324739070 && __local_var_1.UnsafePtr == nil) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(nil))}
goto end_branch_1
} else {

}
}
{
if (__local_var_2.Type == 9 && __local_var_2.IntVal == 324739070 && __local_var_2.UnsafePtr == nil) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__local_var_1))}
goto end_branch_1
} else {

}
}
{
if (__local_var_2.Type == 9 && __local_var_2.IntVal == 324739070 && __local_var_2.UnsafePtr != nil) {
v_3_0 := gopurs_runtime.UncurriedApp3(Get_unsafeSplit(), __local_var_0, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V2, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__local_var_1))})
_ = v_3_0
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp2(Get_unsafeJoinNodes(), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp3(Get_unsafeDifference(), __local_var_0, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Split[gopurs_runtime.Value, gopurs_runtime.Value])(v_3_0.UnsafePtr).V1)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V4)})))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp3(Get_unsafeDifference(), __local_var_0, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Split[gopurs_runtime.Value, gopurs_runtime.Value])(v_3_0.UnsafePtr).V2)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V5)})))})))}
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__t1))}
}

func Call_unsafeIntersectionWith__4109280494(__local_var_0_loop gopurs_runtime.Value, __local_var_1_loop gopurs_runtime.Value, __local_var_2_loop gopurs_runtime.Value, __local_var_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
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
__t6 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(nil))}
goto end_branch_6
} else {

}
}
{
if (__local_var_3.Type == 9 && __local_var_3.IntVal == 324739070 && __local_var_3.UnsafePtr == nil) {
__t6 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(nil))}
goto end_branch_6
} else {

}
}
{
if (__local_var_3.Type == 9 && __local_var_3.IntVal == 324739070 && __local_var_3.UnsafePtr != nil) {
v_4_0 := gopurs_runtime.UncurriedApp3(Get_unsafeSplit(), __local_var_0, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V2, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__local_var_2))})
_ = v_4_0
l_prime_5_1 := gopurs_runtime.UncurriedApp4(Get_unsafeIntersectionWith(), __local_var_0, __local_var_1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Split[gopurs_runtime.Value, gopurs_runtime.Value])(v_4_0.UnsafePtr).V1)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V4)})
_ = l_prime_5_1
r_prime_6_2 := gopurs_runtime.UncurriedApp4(Get_unsafeIntersectionWith(), __local_var_0, __local_var_1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Split[gopurs_runtime.Value, gopurs_runtime.Value])(v_4_0.UnsafePtr).V2)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V5)})
_ = r_prime_6_2
var __t5 gopurs_runtime.Value
{
var __t_tag_3 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Split[gopurs_runtime.Value, gopurs_runtime.Value])(v_4_0.UnsafePtr).V0)}
if (__t_tag_3.Type == 9 && __t_tag_3.IntVal == 930809136 && __t_tag_3.UnsafePtr != nil) {
__t5 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_unsafeBalancedNode(), (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V2, gopurs_runtime.Apply2(__local_var_1, (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Split[gopurs_runtime.Value, gopurs_runtime.Value])(v_4_0.UnsafePtr).V0)}.UnsafePtr).V0, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V3), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](l_prime_5_1))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](r_prime_6_2))})))}
goto end_branch_5
} else {

}
}
{
var __t_tag_4 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Split[gopurs_runtime.Value, gopurs_runtime.Value])(v_4_0.UnsafePtr).V0)}
if (__t_tag_4.Type == 9 && __t_tag_4.IntVal == 930809136 && __t_tag_4.UnsafePtr == nil) {
__t5 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp2(Get_unsafeJoinNodes(), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](l_prime_5_1))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](r_prime_6_2))})))}
goto end_branch_5
} else {

}
}
{
__t5 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_5:
__t6 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__t5))}
goto end_branch_6
} else {

}
}
{
__t6 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_6:
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__t6))}
}

func Call_unsafeJoinNodes__3967876672(__local_var_0_loop gopurs_runtime.Value, __local_var_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __local_var_0 gopurs_runtime.Value = __local_var_0_loop
_ = __local_var_0
var __local_var_1 gopurs_runtime.Value = __local_var_1_loop
_ = __local_var_1
var __t1 gopurs_runtime.Value
{
if (__local_var_0.Type == 9 && __local_var_0.IntVal == 324739070 && __local_var_0.UnsafePtr == nil) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__local_var_1))}
goto end_branch_1
} else {

}
}
{
if (__local_var_0.Type == 9 && __local_var_0.IntVal == 324739070 && __local_var_0.UnsafePtr != nil) {
v2_2_0 := gopurs_runtime.UncurriedApp4(Get_unsafeSplitLast(), (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_0.UnsafePtr).V2, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_0.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_0.UnsafePtr).V4)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_0.UnsafePtr).V5)})
_ = v2_2_0
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_unsafeBalancedNode(), (*Constructor_SplitLast[gopurs_runtime.Value, gopurs_runtime.Value])(v2_2_0.UnsafePtr).V0, (*Constructor_SplitLast[gopurs_runtime.Value, gopurs_runtime.Value])(v2_2_0.UnsafePtr).V1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_SplitLast[gopurs_runtime.Value, gopurs_runtime.Value])(v2_2_0.UnsafePtr).V2)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__local_var_1))})))}
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__t1))}
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
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(&Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]{1, 1, 1, __local_var_0, __local_var_1, gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__local_var_2), gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__local_var_3)})}
goto end_branch_0
} else {

}
}
{
if (__local_var_3.Type == 9 && __local_var_3.IntVal == 324739070 && __local_var_3.UnsafePtr != nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(&Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]{1, Call_add__560788792(gopurs_runtime.Int(1), gopurs_runtime.Int((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V0)).IntVal, Call_add__560788792(gopurs_runtime.Int(1), gopurs_runtime.Int((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V1)).IntVal, __local_var_0, __local_var_1, gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__local_var_2), gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__local_var_3)})}
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__t0))}
goto end_branch_3
} else {

}
}
{
if (__local_var_2.Type == 9 && __local_var_2.IntVal == 324739070 && __local_var_2.UnsafePtr != nil) {
var __t2 gopurs_runtime.Value
{
if (__local_var_3.Type == 9 && __local_var_3.IntVal == 324739070 && __local_var_3.UnsafePtr == nil) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(&Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]{1, Call_add__560788792(gopurs_runtime.Int(1), gopurs_runtime.Int((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V0)).IntVal, Call_add__560788792(gopurs_runtime.Int(1), gopurs_runtime.Int((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V1)).IntVal, __local_var_0, __local_var_1, gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__local_var_2), gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__local_var_3)})}
goto end_branch_2
} else {

}
}
{
if (__local_var_3.Type == 9 && __local_var_3.IntVal == 324739070 && __local_var_3.UnsafePtr != nil) {
var __t1 int64
{
if (gopurs_runtime.Bool(Call_greaterThan__4087042607(gopurs_runtime.Int((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V0), gopurs_runtime.Int((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V0))).IntVal) != (0) {
__t1 = (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V0
goto end_branch_1
} else {

}
}
{
__t1 = (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V0
}
end_branch_1:
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(&Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]{1, Call_add__560788792(gopurs_runtime.Int(1), gopurs_runtime.Int(__t1)).IntVal, Call_add__560788792(gopurs_runtime.Int(Call_add__560788792(gopurs_runtime.Int(1), gopurs_runtime.Int((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V1)).IntVal), gopurs_runtime.Int((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V1)).IntVal, __local_var_0, __local_var_1, gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__local_var_2), gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__local_var_3)})}
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_2:
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__t2))}
goto end_branch_3
} else {

}
}
{
__t3 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_3:
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__t3))}
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
__t4 = gopurs_runtime.Value{Type: 9, IntVal: 3373277644, UnsafePtr: unsafe.Pointer(&Constructor_Split[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(nil))}), gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(nil))}), gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(nil))})})}
goto end_branch_4
} else {

}
}
{
if (__local_var_2.Type == 9 && __local_var_2.IntVal == 324739070 && __local_var_2.UnsafePtr != nil) {
v_3_0 := gopurs_runtime.Apply2(__local_var_0, __local_var_1, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V2)
_ = v_3_0
var __t3 gopurs_runtime.Value
{
if (uint32(v_3_0.IntVal) == 1527465420) {
v1_4_1 := gopurs_runtime.UncurriedApp3(Get_unsafeSplit(), __local_var_0, __local_var_1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V4)})
_ = v1_4_1
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 3373277644, UnsafePtr: unsafe.Pointer(&Constructor_Split[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Split[gopurs_runtime.Value, gopurs_runtime.Value])(v1_4_1.UnsafePtr).V0, (*Constructor_Split[gopurs_runtime.Value, gopurs_runtime.Value])(v1_4_1.UnsafePtr).V1, gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_unsafeBalancedNode(), (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V2, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Split[gopurs_runtime.Value, gopurs_runtime.Value])(v1_4_1.UnsafePtr).V2)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V5)}))})}
goto end_branch_3
} else {

}
}
{
if (uint32(v_3_0.IntVal) == 380165415) {
v1_4_2 := gopurs_runtime.UncurriedApp3(Get_unsafeSplit(), __local_var_0, __local_var_1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V5)})
_ = v1_4_2
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 3373277644, UnsafePtr: unsafe.Pointer(&Constructor_Split[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Split[gopurs_runtime.Value, gopurs_runtime.Value])(v1_4_2.UnsafePtr).V0, gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_unsafeBalancedNode(), (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V2, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V4)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Split[gopurs_runtime.Value, gopurs_runtime.Value])(v1_4_2.UnsafePtr).V1)})), (*Constructor_Split[gopurs_runtime.Value, gopurs_runtime.Value])(v1_4_2.UnsafePtr).V2})}
goto end_branch_3
} else {

}
}
{
if (uint32(v_3_0.IntVal) == 902936544) {
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 3373277644, UnsafePtr: unsafe.Pointer(&Constructor_Split[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V3})}), (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V4, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V5})}
goto end_branch_3
} else {

}
}
{
__t3 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_3:
__t4 = gopurs_runtime.Value{Type: 9, IntVal: 3373277644, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Split[gopurs_runtime.Value, gopurs_runtime.Value]](__t3))}
goto end_branch_4
} else {

}
}
{
__t4 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_4:
return gopurs_runtime.Value{Type: 9, IntVal: 3373277644, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Split[gopurs_runtime.Value, gopurs_runtime.Value]](__t4))}
}

func Call_unsafeSplitLast__224676098(__local_var_0_loop gopurs_runtime.Value, __local_var_1_loop gopurs_runtime.Value, __local_var_2_loop gopurs_runtime.Value, __local_var_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
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
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 2668112006, UnsafePtr: unsafe.Pointer(&Constructor_SplitLast[gopurs_runtime.Value, gopurs_runtime.Value]{1, __local_var_0, __local_var_1, gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__local_var_2)})}
goto end_branch_1
} else {

}
}
{
if (__local_var_3.Type == 9 && __local_var_3.IntVal == 324739070 && __local_var_3.UnsafePtr != nil) {
v1_4_0 := gopurs_runtime.UncurriedApp4(Get_unsafeSplitLast(), (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V2, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V4)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V5)})
_ = v1_4_0
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 2668112006, UnsafePtr: unsafe.Pointer(&Constructor_SplitLast[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_SplitLast[gopurs_runtime.Value, gopurs_runtime.Value])(v1_4_0.UnsafePtr).V0, (*Constructor_SplitLast[gopurs_runtime.Value, gopurs_runtime.Value])(v1_4_0.UnsafePtr).V1, gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_unsafeBalancedNode(), __local_var_0, __local_var_1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__local_var_2))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_SplitLast[gopurs_runtime.Value, gopurs_runtime.Value])(v1_4_0.UnsafePtr).V2)}))})}
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
return gopurs_runtime.Value{Type: 9, IntVal: 2668112006, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_SplitLast[gopurs_runtime.Value, gopurs_runtime.Value]](__t1))}
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
__t6 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__local_var_3))}
goto end_branch_6
} else {

}
}
{
if (__local_var_3.Type == 9 && __local_var_3.IntVal == 324739070 && __local_var_3.UnsafePtr == nil) {
__t6 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__local_var_2))}
goto end_branch_6
} else {

}
}
{
if (__local_var_3.Type == 9 && __local_var_3.IntVal == 324739070 && __local_var_3.UnsafePtr != nil) {
v_4_0 := gopurs_runtime.UncurriedApp3(Get_unsafeSplit(), __local_var_0, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V2, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__local_var_2))})
_ = v_4_0
l_prime_5_1 := gopurs_runtime.UncurriedApp4(Get_unsafeUnionWith(), __local_var_0, __local_var_1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Split[gopurs_runtime.Value, gopurs_runtime.Value])(v_4_0.UnsafePtr).V1)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V4)})
_ = l_prime_5_1
r_prime_6_2 := gopurs_runtime.UncurriedApp4(Get_unsafeUnionWith(), __local_var_0, __local_var_1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Split[gopurs_runtime.Value, gopurs_runtime.Value])(v_4_0.UnsafePtr).V2)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V5)})
_ = r_prime_6_2
var __t5 gopurs_runtime.Value
{
var __t_tag_3 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Split[gopurs_runtime.Value, gopurs_runtime.Value])(v_4_0.UnsafePtr).V0)}
if (__t_tag_3.Type == 9 && __t_tag_3.IntVal == 930809136 && __t_tag_3.UnsafePtr != nil) {
__t5 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_unsafeBalancedNode(), (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V2, gopurs_runtime.Apply2(__local_var_1, (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Split[gopurs_runtime.Value, gopurs_runtime.Value])(v_4_0.UnsafePtr).V0)}.UnsafePtr).V0, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V3), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](l_prime_5_1))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](r_prime_6_2))})))}
goto end_branch_5
} else {

}
}
{
var __t_tag_4 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Split[gopurs_runtime.Value, gopurs_runtime.Value])(v_4_0.UnsafePtr).V0)}
if (__t_tag_4.Type == 9 && __t_tag_4.IntVal == 930809136 && __t_tag_4.UnsafePtr == nil) {
__t5 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_unsafeBalancedNode(), (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V2, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](l_prime_5_1))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](r_prime_6_2))})))}
goto end_branch_5
} else {

}
}
{
__t5 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_5:
__t6 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__t5))}
goto end_branch_6
} else {

}
}
{
__t6 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_6:
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__t6))}
}

func Call_fromJust__1791383420(_dollar__unused_0_loop gopurs_runtime.Value, v_1_loop *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]) gopurs_runtime.Value {
var _dollar__unused_0 gopurs_runtime.Value = _dollar__unused_0_loop
_ = _dollar__unused_0
var v_1 *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value] = v_1_loop
_ = v_1
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v_1)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v_1)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v_1)}.UnsafePtr != nil) {
__t0 = (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v_1)}.UnsafePtr).V0
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return __t0
}

func Call_isNothing__2591355336(v2_0_loop *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]) bool {
var v2_0 *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value] = v2_0_loop
_ = v2_0
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_0)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_0)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_0)}.UnsafePtr == nil) {
__t0 = gopurs_runtime.Bool(true)
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_0)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_0)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_0)}.UnsafePtr != nil) {
__t0 = gopurs_runtime.Bool(false)
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return (__t0.IntVal) != (0)
}

func Call_maybe__1594528518(v_0_loop gopurs_runtime.Value, v1_1_loop gopurs_runtime.Value, v2_2_loop *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var v1_1 gopurs_runtime.Value = v1_1_loop
_ = v1_1
var v2_2 *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value] = v2_2_loop
_ = v2_2
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_2)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_2)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_2)}.UnsafePtr == nil) {
__t0 = v_0
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_2)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_2)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_2)}.UnsafePtr != nil) {
__t0 = gopurs_runtime.Apply(v1_1, (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_2)}.UnsafePtr).V0)
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return __t0
}

func Call_maybe__3658316244(v_0_loop gopurs_runtime.Value, v1_1_loop gopurs_runtime.Value, v2_2_loop *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var v1_1 gopurs_runtime.Value = v1_1_loop
_ = v1_1
var v2_2 *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value] = v2_2_loop
_ = v2_2
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_2)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_2)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_2)}.UnsafePtr == nil) {
__t0 = v_0
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_2)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_2)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_2)}.UnsafePtr != nil) {
__t0 = gopurs_runtime.Apply(v1_1, (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_2)}.UnsafePtr).V0)
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return __t0
}

func Call_mempty__2312420373(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "mempty")
}

func Call_abs__1599282999(__eta0_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
var __t2 gopurs_runtime.Value
{
var __t1 bool
{
if (__eta0_0.IntVal) < (gopurs_runtime.Int(0).IntVal) {
__t1 = false
goto end_branch_1
} else {

}
}
{
__t1 = true
}
end_branch_1:
if __t1 {
__t2 = __eta0_0
goto end_branch_2
} else {

}
}
{
__local_var_1_0 := &pkg_Data_Ring.Constructor_Ring[int64]{1, gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict4("add", "mul", "one", "zero", pkg_Data_Semiring.Get_intAdd(), pkg_Data_Semiring.Get_intMul(), gopurs_runtime.Int(1), gopurs_runtime.Int(0))
}), pkg_Data_Ring.Get_intSub()}
_ = __local_var_1_0
__t2 = gopurs_runtime.Apply2(__local_var_1_0.V1, gopurs_runtime.RecordGet(gopurs_runtime.Apply(__local_var_1_0.V0, gopurs_runtime.Value{}), "zero"), __eta0_0)
}
end_branch_2:
return __t2
}

func Call_abs__2515802711(dictOrd_0_loop *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value], dictRing_1_loop *pkg_Data_Ring.Constructor_Ring[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictOrd_0 *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value] = dictOrd_0_loop
_ = dictOrd_0
var dictRing_1 *pkg_Data_Ring.Constructor_Ring[gopurs_runtime.Value] = dictRing_1_loop
_ = dictRing_1
zero_2_0 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(dictRing_1.V0, gopurs_runtime.Value{}), "zero")
_ = zero_2_0
return gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t3 gopurs_runtime.Value
{
var __t2 bool
{
var __t_tag_1 gopurs_runtime.Value = gopurs_runtime.Apply2(dictOrd_0.V1, x_3, zero_2_0)
if (uint32(__t_tag_1.IntVal) == 1527465420) {
__t2 = false
goto end_branch_2
} else {

}
}
{
__t2 = true
}
end_branch_2:
if __t2 {
__t3 = x_3
goto end_branch_3
} else {

}
}
{
__t3 = gopurs_runtime.Apply2(dictRing_1.V1, gopurs_runtime.RecordGet(gopurs_runtime.Apply(dictRing_1.V0, gopurs_runtime.Value{}), "zero"), x_3)
}
end_branch_3:
return __t3
})
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

func Call_compare__2802126154(dict_0_loop *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value]) gopurs_runtime.Value {
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

func Call_greaterThanOrEq__1697837627(dictOrd_0_loop *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value], a1_1_loop gopurs_runtime.Value, a2_2_loop gopurs_runtime.Value) bool {
var dictOrd_0 *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value] = dictOrd_0_loop
_ = dictOrd_0
var a1_1 gopurs_runtime.Value = a1_1_loop
_ = a1_1
var a2_2 gopurs_runtime.Value = a2_2_loop
_ = a2_2
var __t1 bool
{
var __t_tag_0 gopurs_runtime.Value = gopurs_runtime.Apply2(dictOrd_0.V1, a1_1, a2_2)
if (uint32(__t_tag_0.IntVal) == 1527465420) {
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

func Call_greaterThanOrEq__1409282474(dictOrd_0_loop *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value], a1_1_loop gopurs_runtime.Value, a2_2_loop gopurs_runtime.Value) bool {
var dictOrd_0 *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value] = dictOrd_0_loop
_ = dictOrd_0
var a1_1 gopurs_runtime.Value = a1_1_loop
_ = a1_1
var a2_2 gopurs_runtime.Value = a2_2_loop
_ = a2_2
var __t1 bool
{
var __t_tag_0 gopurs_runtime.Value = gopurs_runtime.Apply2(dictOrd_0.V1, a1_1, a2_2)
if (uint32(__t_tag_0.IntVal) == 1527465420) {
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

func Call_lessThan__4087042607(a1_0_loop gopurs_runtime.Value, a2_1_loop gopurs_runtime.Value) bool {
var a1_0 gopurs_runtime.Value = a1_0_loop
_ = a1_0
var a2_1 gopurs_runtime.Value = a2_1_loop
_ = a2_1
var __t0 bool
{
if (a1_0.IntVal) < (a2_1.IntVal) {
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

func Call_lessThan__1409282474(dictOrd_0_loop *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value], a1_1_loop gopurs_runtime.Value, a2_2_loop gopurs_runtime.Value) bool {
var dictOrd_0 *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value] = dictOrd_0_loop
_ = dictOrd_0
var a1_1 gopurs_runtime.Value = a1_1_loop
_ = a1_1
var a2_2 gopurs_runtime.Value = a2_2_loop
_ = a2_2
var __t1 bool
{
var __t_tag_0 gopurs_runtime.Value = gopurs_runtime.Apply2(dictOrd_0.V1, a1_1, a2_2)
if (uint32(__t_tag_0.IntVal) == 1527465420) {
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

func Call_negate__2674436160(dictRing_0_loop *pkg_Data_Ring.Constructor_Ring[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictRing_0 *pkg_Data_Ring.Constructor_Ring[gopurs_runtime.Value] = dictRing_0_loop
_ = dictRing_0
Semiring0_1_0 := gopurs_runtime.CoerceToStruct[pkg_Data_Semiring.Constructor_Semiring[gopurs_runtime.Value]](gopurs_runtime.Apply(dictRing_0.V0, gopurs_runtime.Value{}))
_ = Semiring0_1_0
return gopurs_runtime.Func(func(a_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(dictRing_0.V1, Semiring0_1_0.V3, a_2)
})
}

func Call_negate__1364373265(dictRing_0_loop *pkg_Data_Ring.Constructor_Ring[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictRing_0 *pkg_Data_Ring.Constructor_Ring[gopurs_runtime.Value] = dictRing_0_loop
_ = dictRing_0
Semiring0_1_0 := gopurs_runtime.CoerceToStruct[pkg_Data_Semiring.Constructor_Semiring[gopurs_runtime.Value]](gopurs_runtime.Apply(dictRing_0.V0, gopurs_runtime.Value{}))
_ = Semiring0_1_0
return gopurs_runtime.Func(func(a_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(dictRing_0.V1, Semiring0_1_0.V3, a_2)
})
}

func Call_sub__1043827704(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
_ = __eta1_1
return gopurs_runtime.Int((__eta0_0.IntVal) - (__eta1_1.IntVal))
}

func Call_sub__3675938712(dict_0_loop *pkg_Data_Ring.Constructor_Ring[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Ring.Constructor_Ring[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_append__493084344(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
_ = __eta1_1
return gopurs_runtime.Str((__eta0_0.StrVal()) + (__eta1_1.StrVal()))
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

func Call_zero__1204848985(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "zero")
}

func Call_show__1488465650(__eta0_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
return gopurs_runtime.Str(gopurs_runtime.Apply(pkg_Data_Show.Get_showIntImpl(), __eta0_0).StrVal())
}

func Call_show__2742601362(dict_0_loop *pkg_Data_Show.Constructor_Show[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Show.Constructor_Show[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_show__2696961938(dict_0_loop *pkg_Data_Show.Constructor_Show[[]*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Show.Constructor_Show[[]*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]] = dict_0_loop
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
var go__go_4_2_63 gopurs_runtime.Value
_ = go__go_4_2_63
go__go_4_2_63 = gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
var __t6 gopurs_runtime.Value
{
if (v_5.Type == 9 && v_5.IntVal == 324739070 && v_5.UnsafePtr == nil) {
__t6 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "pure"), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(nil))})
goto end_branch_6
} else {

}
}
{
if (v_5.Type == 9 && v_5.IntVal == 324739070 && v_5.UnsafePtr != nil) {
__local_var_6_3 := (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_5.UnsafePtr).V0
_ = __local_var_6_3
__local_var_7_4 := (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_5.UnsafePtr).V2
_ = __local_var_7_4
__local_var_8_5 := (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_5.UnsafePtr).V1
_ = __local_var_8_5
__t6 = gopurs_runtime.Apply2(Apply0_1_0.V1, gopurs_runtime.Apply2(Apply0_1_0.V1, gopurs_runtime.Apply2(Functor0_2_1.V0, gopurs_runtime.Func(func(l_prime_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_prime_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(r_prime_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(&Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]{1, __local_var_6_3, __local_var_8_5, __local_var_7_4, v_prime_10, gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](l_prime_9), gopurs_runtime.CoerceToStruct[Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](r_prime_11)})}
})
})
}), gopurs_runtime.Apply(go__go_4_2_63, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_5.UnsafePtr).V4)})), gopurs_runtime.Apply(f_3, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_5.UnsafePtr).V3)), gopurs_runtime.Apply(go__go_4_2_63, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_5.UnsafePtr).V5)}))
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
return go__go_4_2_63
})
}

func Call_fst__20422131(v_0_loop *pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var v_0 *pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value] = v_0_loop
_ = v_0
return (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V0
}

func Call_snd__20422131(v_0_loop *pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var v_0 *pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value] = v_0_loop
_ = v_0
return (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V1
}

func Call_unfoldr__1128708256(dict_0_loop *pkg_Data_Unfoldable.Constructor_Unfoldable[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Unfoldable.Constructor_Unfoldable[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_unfoldr__1519733018(dict_0_loop *pkg_Data_Unfoldable.Constructor_Unfoldable[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Unfoldable.Constructor_Unfoldable[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}


