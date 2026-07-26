package Data_Map_Internal

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_Ord "gopurs/output/Data.Ord"
	pkg_Data_Maybe "gopurs/output/Data.Maybe"
	pkg_Data_Tuple "gopurs/output/Data.Tuple"
	pkg_Data_Unfoldable "gopurs/output/Data.Unfoldable"
	pkg_Data_List_Types "gopurs/output/Data.List.Types"
	pkg_Data_Unit "gopurs/output/Data.Unit"
	pkg_Data_Function "gopurs/output/Data.Function"
	pkg_Data_Semigroup "gopurs/output/Data.Semigroup"
	pkg_Data_Show "gopurs/output/Data.Show"
	pkg_Data_HeytingAlgebra "gopurs/output/Data.HeytingAlgebra"
	unsafe "unsafe"
)

var cache_greaterThan gopurs_runtime.Value
var once_greaterThan sync.Once
func Get_greaterThan() gopurs_runtime.Value {
	once_greaterThan.Do(func() {
		cache_greaterThan = func() gopurs_runtime.Value {
__local_var_0_0 := gopurs_runtime.Apply3(pkg_Data_Ord.Get_ordIntImpl(), gopurs_runtime.Value{Type: 9, IntVal: 1527465420, UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: 902936544, UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: 380165415, UnsafePtr: nil})
_ = __local_var_0_0
return gopurs_runtime.Func2(func(a1_1 gopurs_runtime.Value, a2_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t_tag_1 gopurs_runtime.Value = gopurs_runtime.Apply2(__local_var_0_0, a1_1, a2_2)
return gopurs_runtime.Bool((__t_tag_1.Type == 9 && __t_tag_1.IntVal == 380165415))
})
}()
	})
	return cache_greaterThan
}

var cache_lessThanOrEq gopurs_runtime.Value
var once_lessThanOrEq sync.Once
func Get_lessThanOrEq() gopurs_runtime.Value {
	once_lessThanOrEq.Do(func() {
		cache_lessThanOrEq = func() gopurs_runtime.Value {
__local_var_0_0 := gopurs_runtime.Apply3(pkg_Data_Ord.Get_ordIntImpl(), gopurs_runtime.Value{Type: 9, IntVal: 1527465420, UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: 902936544, UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: 380165415, UnsafePtr: nil})
_ = __local_var_0_0
return gopurs_runtime.Func2(func(a1_1 gopurs_runtime.Value, a2_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t_tag_1 gopurs_runtime.Value = gopurs_runtime.Apply2(__local_var_0_0, a1_1, a2_2)
return gopurs_runtime.Bool(((__t_tag_1.Type == 9 && __t_tag_1.IntVal == 380165415)) != (true))
})
}()
	})
	return cache_lessThanOrEq
}

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

var cache_lessThan gopurs_runtime.Value
var once_lessThan sync.Once
func Get_lessThan() gopurs_runtime.Value {
	once_lessThan.Do(func() {
		cache_lessThan = func() gopurs_runtime.Value {
__local_var_0_0 := gopurs_runtime.Apply3(pkg_Data_Ord.Get_ordIntImpl(), gopurs_runtime.Value{Type: 9, IntVal: 1527465420, UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: 902936544, UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: 380165415, UnsafePtr: nil})
_ = __local_var_0_0
return gopurs_runtime.Func2(func(a1_1 gopurs_runtime.Value, a2_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t_tag_1 gopurs_runtime.Value = gopurs_runtime.Apply2(__local_var_0_0, a1_1, a2_2)
return gopurs_runtime.Bool((__t_tag_1.Type == 9 && __t_tag_1.IntVal == 1527465420))
})
}()
	})
	return cache_lessThan
}

var cache_abs gopurs_runtime.Value
var once_abs sync.Once
func Get_abs() gopurs_runtime.Value {
	once_abs.Do(func() {
		cache_abs = func() gopurs_runtime.Value {
__local_var_0_0 := gopurs_runtime.Apply3(pkg_Data_Ord.Get_ordIntImpl(), gopurs_runtime.Value{Type: 9, IntVal: 1527465420, UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: 902936544, UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: 380165415, UnsafePtr: nil})
_ = __local_var_0_0
return gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
var __t_tag_2 gopurs_runtime.Value = gopurs_runtime.Apply2(__local_var_0_0, x_1, gopurs_runtime.Int(0))
if ((__t_tag_2.Type == 9 && __t_tag_2.IntVal == 1527465420)) != (true) {
__t1 = x_1
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Int((0) - (x_1.IntVal))
}
end_branch_1:
return __t1
})
}()
	})
	return cache_abs
}

var cache_Leaf gopurs_runtime.Value
var once_Leaf sync.Once
func Get_Leaf() gopurs_runtime.Value {
	once_Leaf.Do(func() {
		cache_Leaf = gopurs_runtime.Value{Type: 9, IntVal: 687041424, UnsafePtr: nil}
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
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(&Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]{value0.IntVal, value1.IntVal, value2, value3, value4, value5})}
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
		cache_IterLeaf = gopurs_runtime.Value{Type: 9, IntVal: 2509360378, UnsafePtr: nil}
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
return gopurs_runtime.Value{Type: 9, IntVal: 1343415489, UnsafePtr: unsafe.Pointer(&Constructor_IterEmit[gopurs_runtime.Value, gopurs_runtime.Value]{value0, value1, value2})}
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
return gopurs_runtime.Value{Type: 9, IntVal: 2861335956, UnsafePtr: unsafe.Pointer(&Constructor_IterNode[gopurs_runtime.Value, gopurs_runtime.Value]{value0, value1})}
})
})
	})
	return cache_IterNode
}

var cache_IterDone gopurs_runtime.Value
var once_IterDone sync.Once
func Get_IterDone() gopurs_runtime.Value {
	once_IterDone.Do(func() {
		cache_IterDone = gopurs_runtime.Value{Type: 9, IntVal: 4236111124, UnsafePtr: nil}
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
return gopurs_runtime.Value{Type: 9, IntVal: 953589075, UnsafePtr: unsafe.Pointer(&Constructor_IterNext[gopurs_runtime.Value, gopurs_runtime.Value]{value0, value1, value2})}
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
return gopurs_runtime.Value{Type: 9, IntVal: 3373277644, UnsafePtr: unsafe.Pointer(&Constructor_Split[gopurs_runtime.Value, gopurs_runtime.Value]{value0, value1, value2})}
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
return gopurs_runtime.Value{Type: 9, IntVal: 2668112006, UnsafePtr: unsafe.Pointer(&Constructor_SplitLast[gopurs_runtime.Value, gopurs_runtime.Value]{value0, value1, value2})}
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
		cache_unsafeNode = gopurs_runtime.Func4(func(k_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value, l_2_box gopurs_runtime.Value, r_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_unsafeNode(k_0_box, v_1_box, l_2_box, r_3_box)
})
	})
	return cache_unsafeNode
}

var cache_toMapIter gopurs_runtime.Value
var once_toMapIter sync.Once
func Get_toMapIter() gopurs_runtime.Value {
	once_toMapIter.Do(func() {
		cache_toMapIter = gopurs_runtime.Func(func(a_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_toMapIter(a_0_box)
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
return Call_size(v_0_box)
})
	})
	return cache_size
}

var cache_singleton gopurs_runtime.Value
var once_singleton sync.Once
func Get_singleton() gopurs_runtime.Value {
	once_singleton.Do(func() {
		cache_singleton = gopurs_runtime.Func2(func(k_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_singleton(k_0_box, v_1_box)
})
	})
	return cache_singleton
}

var cache_unsafeBalancedNode gopurs_runtime.Value
var once_unsafeBalancedNode sync.Once
func Get_unsafeBalancedNode() gopurs_runtime.Value {
	once_unsafeBalancedNode.Do(func() {
		cache_unsafeBalancedNode = gopurs_runtime.Func4(func(k_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value, l_2_box gopurs_runtime.Value, r_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_unsafeBalancedNode(k_0_box, v_1_box, l_2_box, r_3_box)
})
	})
	return cache_unsafeBalancedNode
}

var cache_unsafeSplit gopurs_runtime.Value
var once_unsafeSplit sync.Once
func Get_unsafeSplit() gopurs_runtime.Value {
	once_unsafeSplit.Do(func() {
		cache_unsafeSplit = gopurs_runtime.Func3(func(comp_0_box gopurs_runtime.Value, k_1_box gopurs_runtime.Value, m_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_unsafeSplit(comp_0_box, k_1_box, m_2_box)
})
	})
	return cache_unsafeSplit
}

var cache_unsafeSplitLast gopurs_runtime.Value
var once_unsafeSplitLast sync.Once
func Get_unsafeSplitLast() gopurs_runtime.Value {
	once_unsafeSplitLast.Do(func() {
		cache_unsafeSplitLast = gopurs_runtime.Func4(func(k_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value, l_2_box gopurs_runtime.Value, r_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_unsafeSplitLast(k_0_box, v_1_box, l_2_box, r_3_box)
})
	})
	return cache_unsafeSplitLast
}

var cache_unsafeJoinNodes gopurs_runtime.Value
var once_unsafeJoinNodes sync.Once
func Get_unsafeJoinNodes() gopurs_runtime.Value {
	once_unsafeJoinNodes.Do(func() {
		cache_unsafeJoinNodes = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_unsafeJoinNodes(v_0_box, v1_1_box)
})
	})
	return cache_unsafeJoinNodes
}

var cache_unsafeDifference gopurs_runtime.Value
var once_unsafeDifference sync.Once
func Get_unsafeDifference() gopurs_runtime.Value {
	once_unsafeDifference.Do(func() {
		cache_unsafeDifference = gopurs_runtime.Func3(func(comp_0_box gopurs_runtime.Value, l_1_box gopurs_runtime.Value, r_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_unsafeDifference(comp_0_box, l_1_box, r_2_box)
})
	})
	return cache_unsafeDifference
}

var cache_unsafeIntersectionWith gopurs_runtime.Value
var once_unsafeIntersectionWith sync.Once
func Get_unsafeIntersectionWith() gopurs_runtime.Value {
	once_unsafeIntersectionWith.Do(func() {
		cache_unsafeIntersectionWith = gopurs_runtime.Func4(func(comp_0_box gopurs_runtime.Value, app_1_box gopurs_runtime.Value, l_2_box gopurs_runtime.Value, r_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_unsafeIntersectionWith(comp_0_box, app_1_box, l_2_box, r_3_box)
})
	})
	return cache_unsafeIntersectionWith
}

var cache_unsafeUnionWith gopurs_runtime.Value
var once_unsafeUnionWith sync.Once
func Get_unsafeUnionWith() gopurs_runtime.Value {
	once_unsafeUnionWith.Do(func() {
		cache_unsafeUnionWith = gopurs_runtime.Func4(func(comp_0_box gopurs_runtime.Value, app_1_box gopurs_runtime.Value, l_2_box gopurs_runtime.Value, r_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_unsafeUnionWith(comp_0_box, app_1_box, l_2_box, r_3_box)
})
	})
	return cache_unsafeUnionWith
}

var cache_unionWith gopurs_runtime.Value
var once_unionWith sync.Once
func Get_unionWith() gopurs_runtime.Value {
	once_unionWith.Do(func() {
		cache_unionWith = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_unionWith(dictOrd_0_box)
})
	})
	return cache_unionWith
}

var cache_union gopurs_runtime.Value
var once_union sync.Once
func Get_union() gopurs_runtime.Value {
	once_union.Do(func() {
		cache_union = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_union(dictOrd_0_box)
})
	})
	return cache_union
}

var cache_update gopurs_runtime.Value
var once_update sync.Once
func Get_update() gopurs_runtime.Value {
	once_update.Do(func() {
		cache_update = gopurs_runtime.Func3(func(dictOrd_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, k_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_update(dictOrd_0_box, f_1_box, k_2_box)
})
	})
	return cache_update
}

var cache_showTree gopurs_runtime.Value
var once_showTree sync.Once
func Get_showTree() gopurs_runtime.Value {
	once_showTree.Do(func() {
		cache_showTree = gopurs_runtime.Func2(func(dictShow_0_box gopurs_runtime.Value, dictShow1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_showTree(dictShow_0_box, dictShow1_1_box)
})
	})
	return cache_showTree
}

var cache_semigroupMap gopurs_runtime.Value
var once_semigroupMap sync.Once
func Get_semigroupMap() gopurs_runtime.Value {
	once_semigroupMap.Do(func() {
		cache_semigroupMap = gopurs_runtime.Func2(func(_dollar__unused_0_box gopurs_runtime.Value, dictOrd_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_semigroupMap(_dollar__unused_0_box, dictOrd_1_box)
})
	})
	return cache_semigroupMap
}

var cache_pop gopurs_runtime.Value
var once_pop sync.Once
func Get_pop() gopurs_runtime.Value {
	once_pop.Do(func() {
		cache_pop = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_pop(dictOrd_0_box)
})
	})
	return cache_pop
}

var cache_member gopurs_runtime.Value
var once_member sync.Once
func Get_member() gopurs_runtime.Value {
	once_member.Do(func() {
		cache_member = gopurs_runtime.Func2(func(dictOrd_0_box gopurs_runtime.Value, k_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_member(dictOrd_0_box, k_1_box)
})
	})
	return cache_member
}

var cache_mapMaybeWithKey gopurs_runtime.Value
var once_mapMaybeWithKey sync.Once
func Get_mapMaybeWithKey() gopurs_runtime.Value {
	once_mapMaybeWithKey.Do(func() {
		cache_mapMaybeWithKey = gopurs_runtime.Func2(func(dictOrd_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_mapMaybeWithKey(dictOrd_0_box, f_1_box)
})
	})
	return cache_mapMaybeWithKey
}

var cache_mapMaybe gopurs_runtime.Value
var once_mapMaybe sync.Once
func Get_mapMaybe() gopurs_runtime.Value {
	once_mapMaybe.Do(func() {
		cache_mapMaybe = gopurs_runtime.Func2(func(dictOrd_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_mapMaybe(dictOrd_0_box, x_1_box)
})
	})
	return cache_mapMaybe
}

var cache_lookupLE gopurs_runtime.Value
var once_lookupLE sync.Once
func Get_lookupLE() gopurs_runtime.Value {
	once_lookupLE.Do(func() {
		cache_lookupLE = gopurs_runtime.Func2(func(dictOrd_0_box gopurs_runtime.Value, k_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_lookupLE(dictOrd_0_box, k_1_box)
})
	})
	return cache_lookupLE
}

var cache_lookupGE gopurs_runtime.Value
var once_lookupGE sync.Once
func Get_lookupGE() gopurs_runtime.Value {
	once_lookupGE.Do(func() {
		cache_lookupGE = gopurs_runtime.Func2(func(dictOrd_0_box gopurs_runtime.Value, k_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_lookupGE(dictOrd_0_box, k_1_box)
})
	})
	return cache_lookupGE
}

var cache_lookup gopurs_runtime.Value
var once_lookup sync.Once
func Get_lookup() gopurs_runtime.Value {
	once_lookup.Do(func() {
		cache_lookup = gopurs_runtime.Func2(func(dictOrd_0_box gopurs_runtime.Value, k_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_lookup(dictOrd_0_box, k_1_box)
})
	})
	return cache_lookup
}

var cache_iterMapU gopurs_runtime.Value
var once_iterMapU sync.Once
func Get_iterMapU() gopurs_runtime.Value {
	once_iterMapU.Do(func() {
		cache_iterMapU = gopurs_runtime.Func2(func(iter_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_iterMapU(iter_0_box, v_1_box)
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
		cache_stepUnfoldrUnordered = Call_stepWith(Get_iterMapU(), gopurs_runtime.Func3(func(k_0 gopurs_runtime.Value, v_1 gopurs_runtime.Value, next_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{k_0, v_1})}, next_2})}})}
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3589588149, UnsafePtr: nil}
}))
	})
	return cache_stepUnfoldrUnordered
}

var cache_toUnfoldableUnordered gopurs_runtime.Value
var once_toUnfoldableUnordered sync.Once
func Get_toUnfoldableUnordered() gopurs_runtime.Value {
	once_toUnfoldableUnordered.Do(func() {
		cache_toUnfoldableUnordered = gopurs_runtime.Func(func(dictUnfoldable_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_toUnfoldableUnordered(dictUnfoldable_0_box)
})
	})
	return cache_toUnfoldableUnordered
}

var cache_stepUnordered gopurs_runtime.Value
var once_stepUnordered sync.Once
func Get_stepUnordered() gopurs_runtime.Value {
	once_stepUnordered.Do(func() {
		cache_stepUnordered = Call_stepWith(Get_iterMapU(), gopurs_runtime.Func3(func(k_0 gopurs_runtime.Value, v_1 gopurs_runtime.Value, next_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 953589075, UnsafePtr: unsafe.Pointer(&Constructor_IterNext[gopurs_runtime.Value, gopurs_runtime.Value]{k_0, v_1, next_2})}
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4236111124, UnsafePtr: nil}
}))
	})
	return cache_stepUnordered
}

var cache_iterMapR gopurs_runtime.Value
var once_iterMapR sync.Once
func Get_iterMapR() gopurs_runtime.Value {
	once_iterMapR.Do(func() {
		cache_iterMapR = func() gopurs_runtime.Value {
var go__0_0 gopurs_runtime.Value
go__0_0 = gopurs_runtime.Func(func(iter_1_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var iter_1_loop gopurs_runtime.Value = iter_1_loop_val
var v_2_loop gopurs_runtime.Value = v_2_loop_val
go__0_0:
for {
if false { continue go__0_0 }
var iter_1 gopurs_runtime.Value = iter_1_loop
_ = iter_1
var v_2 gopurs_runtime.Value = v_2_loop
_ = v_2
var __t1 gopurs_runtime.Value
{
if (v_2.Type == 9 && v_2.IntVal == 687041424) {
__t1 = iter_1
goto end_branch_1
} else {

}
}
{
if (v_2.Type == 9 && v_2.IntVal == 324739070) {
var __t2 gopurs_runtime.Value
{
var __t_tag_3 gopurs_runtime.Value = (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V5
if (__t_tag_3.Type == 9 && __t_tag_3.IntVal == 687041424) {
iter_1_loop = gopurs_runtime.Value{Type: 9, IntVal: 1343415489, UnsafePtr: unsafe.Pointer(&Constructor_IterEmit[gopurs_runtime.Value, gopurs_runtime.Value]{(*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V2, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V3, iter_1})}
v_2_loop = (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V4
continue go__0_0
__t2 = gopurs_runtime.Value{}
goto end_branch_2
} else {

}
}
{
iter_1_loop = gopurs_runtime.Value{Type: 9, IntVal: 1343415489, UnsafePtr: unsafe.Pointer(&Constructor_IterEmit[gopurs_runtime.Value, gopurs_runtime.Value]{(*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V2, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 2861335956, UnsafePtr: unsafe.Pointer(&Constructor_IterNode[gopurs_runtime.Value, gopurs_runtime.Value]{(*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V4, iter_1})}})}
v_2_loop = (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V5
continue go__0_0
__t2 = gopurs_runtime.Value{}
}
end_branch_2:
__t1 = __t2
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
})
return go__0_0
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
		cache_stepDesc = Call_stepWith(Get_iterMapR(), gopurs_runtime.Func3(func(k_0 gopurs_runtime.Value, v_1 gopurs_runtime.Value, next_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 953589075, UnsafePtr: unsafe.Pointer(&Constructor_IterNext[gopurs_runtime.Value, gopurs_runtime.Value]{k_0, v_1, next_2})}
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4236111124, UnsafePtr: nil}
}))
	})
	return cache_stepDesc
}

var cache_iterMapL gopurs_runtime.Value
var once_iterMapL sync.Once
func Get_iterMapL() gopurs_runtime.Value {
	once_iterMapL.Do(func() {
		cache_iterMapL = func() gopurs_runtime.Value {
var go__0_0 gopurs_runtime.Value
go__0_0 = gopurs_runtime.Func(func(iter_1_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var iter_1_loop gopurs_runtime.Value = iter_1_loop_val
var v_2_loop gopurs_runtime.Value = v_2_loop_val
go__0_0:
for {
if false { continue go__0_0 }
var iter_1 gopurs_runtime.Value = iter_1_loop
_ = iter_1
var v_2 gopurs_runtime.Value = v_2_loop
_ = v_2
var __t1 gopurs_runtime.Value
{
if (v_2.Type == 9 && v_2.IntVal == 687041424) {
__t1 = iter_1
goto end_branch_1
} else {

}
}
{
if (v_2.Type == 9 && v_2.IntVal == 324739070) {
var __t2 gopurs_runtime.Value
{
var __t_tag_3 gopurs_runtime.Value = (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V5
if (__t_tag_3.Type == 9 && __t_tag_3.IntVal == 687041424) {
iter_1_loop = gopurs_runtime.Value{Type: 9, IntVal: 1343415489, UnsafePtr: unsafe.Pointer(&Constructor_IterEmit[gopurs_runtime.Value, gopurs_runtime.Value]{(*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V2, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V3, iter_1})}
v_2_loop = (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V4
continue go__0_0
__t2 = gopurs_runtime.Value{}
goto end_branch_2
} else {

}
}
{
iter_1_loop = gopurs_runtime.Value{Type: 9, IntVal: 1343415489, UnsafePtr: unsafe.Pointer(&Constructor_IterEmit[gopurs_runtime.Value, gopurs_runtime.Value]{(*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V2, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 2861335956, UnsafePtr: unsafe.Pointer(&Constructor_IterNode[gopurs_runtime.Value, gopurs_runtime.Value]{(*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V5, iter_1})}})}
v_2_loop = (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V4
continue go__0_0
__t2 = gopurs_runtime.Value{}
}
end_branch_2:
__t1 = __t2
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
})
return go__0_0
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
		cache_stepAsc = Call_stepWith(Get_iterMapL(), gopurs_runtime.Func3(func(k_0 gopurs_runtime.Value, v_1 gopurs_runtime.Value, next_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 953589075, UnsafePtr: unsafe.Pointer(&Constructor_IterNext[gopurs_runtime.Value, gopurs_runtime.Value]{k_0, v_1, next_2})}
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4236111124, UnsafePtr: nil}
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
		cache_stepUnfoldr = Call_stepWith(Get_iterMapL(), gopurs_runtime.Func3(func(k_0 gopurs_runtime.Value, v_1 gopurs_runtime.Value, next_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{k_0, v_1})}, next_2})}})}
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3589588149, UnsafePtr: nil}
}))
	})
	return cache_stepUnfoldr
}

var cache_toUnfoldable gopurs_runtime.Value
var once_toUnfoldable sync.Once
func Get_toUnfoldable() gopurs_runtime.Value {
	once_toUnfoldable.Do(func() {
		cache_toUnfoldable = gopurs_runtime.Func(func(dictUnfoldable_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_toUnfoldable(dictUnfoldable_0_box)
})
	})
	return cache_toUnfoldable
}

var cache_toUnfoldable1 gopurs_runtime.Value
var once_toUnfoldable1 sync.Once
func Get_toUnfoldable1() gopurs_runtime.Value {
	once_toUnfoldable1.Do(func() {
		cache_toUnfoldable1 = func() gopurs_runtime.Value {
__local_var_0_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Unfoldable.Get_unfoldableArray(), "unfoldr"), Get_stepUnfoldr())
_ = __local_var_0_0
return gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_0_0, gopurs_runtime.Value{Type: 9, IntVal: 2861335956, UnsafePtr: unsafe.Pointer(&Constructor_IterNode[gopurs_runtime.Value, gopurs_runtime.Value]{x_1, gopurs_runtime.Value{Type: 9, IntVal: 2509360378, UnsafePtr: nil}})})
})
}()
	})
	return cache_toUnfoldable1
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
return Call_isSubmap(dictOrd_0_box, dictEq_1_box)
})
	})
	return cache_isSubmap
}

var cache_isEmpty gopurs_runtime.Value
var once_isEmpty sync.Once
func Get_isEmpty() gopurs_runtime.Value {
	once_isEmpty.Do(func() {
		cache_isEmpty = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_isEmpty(v_0_box))
})
	})
	return cache_isEmpty
}

var cache_intersectionWith gopurs_runtime.Value
var once_intersectionWith sync.Once
func Get_intersectionWith() gopurs_runtime.Value {
	once_intersectionWith.Do(func() {
		cache_intersectionWith = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_intersectionWith(dictOrd_0_box)
})
	})
	return cache_intersectionWith
}

var cache_intersection gopurs_runtime.Value
var once_intersection sync.Once
func Get_intersection() gopurs_runtime.Value {
	once_intersection.Do(func() {
		cache_intersection = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_intersection(dictOrd_0_box)
})
	})
	return cache_intersection
}

var cache_insertWith gopurs_runtime.Value
var once_insertWith sync.Once
func Get_insertWith() gopurs_runtime.Value {
	once_insertWith.Do(func() {
		cache_insertWith = gopurs_runtime.Func4(func(dictOrd_0_box gopurs_runtime.Value, app_1_box gopurs_runtime.Value, k_2_box gopurs_runtime.Value, v_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_insertWith(dictOrd_0_box, app_1_box, k_2_box, v_3_box)
})
	})
	return cache_insertWith
}

var cache_insert gopurs_runtime.Value
var once_insert sync.Once
func Get_insert() gopurs_runtime.Value {
	once_insert.Do(func() {
		cache_insert = gopurs_runtime.Func3(func(dictOrd_0_box gopurs_runtime.Value, k_1_box gopurs_runtime.Value, v_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_insert(dictOrd_0_box, k_1_box, v_2_box)
})
	})
	return cache_insert
}

var cache_functorMap gopurs_runtime.Value
var once_functorMap sync.Once
func Get_functorMap() gopurs_runtime.Value {
	once_functorMap.Do(func() {
		cache_functorMap = gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
var go__1_0 gopurs_runtime.Value
_ = go__1_0
go__1_0 = gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if (v_2.Type == 9 && v_2.IntVal == 687041424) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 687041424, UnsafePtr: nil}
goto end_branch_1
} else {

}
}
{
if (v_2.Type == 9 && v_2.IntVal == 324739070) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(&Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]{gopurs_runtime.Int((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V0).IntVal, gopurs_runtime.Int((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V1).IntVal, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V2, gopurs_runtime.Apply(f_0, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V3), gopurs_runtime.Apply(go__1_0, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V4), gopurs_runtime.Apply(go__1_0, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V5)})}
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
return go__1_0
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
var go__1_0 gopurs_runtime.Value
_ = go__1_0
go__1_0 = gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if (v_2.Type == 9 && v_2.IntVal == 687041424) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 687041424, UnsafePtr: nil}
goto end_branch_1
} else {

}
}
{
if (v_2.Type == 9 && v_2.IntVal == 324739070) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(&Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]{gopurs_runtime.Int((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V0).IntVal, gopurs_runtime.Int((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V1).IntVal, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V2, gopurs_runtime.Apply2(f_0, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V2, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V3), gopurs_runtime.Apply(go__1_0, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V4), gopurs_runtime.Apply(go__1_0, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V5)})}
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
return go__1_0
}))
	})
	return cache_functorWithIndexMap
}

var cache_foldableMap gopurs_runtime.Value
var once_foldableMap sync.Once
func Get_foldableMap() gopurs_runtime.Value {
	once_foldableMap.Do(func() {
		cache_foldableMap = gopurs_runtime.RecordDict3("foldMap", "foldl", "foldr", gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
mempty_1_0 := gopurs_runtime.RecordGet(dictMonoid_0, "mempty")
_ = mempty_1_0
__local_var_2_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{})
_ = __local_var_2_1
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
var go__4_2 gopurs_runtime.Value
_ = go__4_2
go__4_2 = gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
var __t3 gopurs_runtime.Value
{
if (v_5.Type == 9 && v_5.IntVal == 687041424) {
__t3 = mempty_1_0
goto end_branch_3
} else {

}
}
{
if (v_5.Type == 9 && v_5.IntVal == 324739070) {
__t3 = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_1, "append"), gopurs_runtime.Apply(go__4_2, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_5.UnsafePtr).V4), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_1, "append"), gopurs_runtime.Apply(f_3, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_5.UnsafePtr).V3), gopurs_runtime.Apply(go__4_2, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_5.UnsafePtr).V5)))
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
return go__4_2
})
}), gopurs_runtime.Func2(func(f_0 gopurs_runtime.Value, z_1 gopurs_runtime.Value) gopurs_runtime.Value {
var go__2_4 gopurs_runtime.Value
_ = go__2_4
go__2_4 = gopurs_runtime.Func2(Call_go__2_4)
return gopurs_runtime.Func(func(m_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.UncurriedApp2(go__2_4, z_1, m_3)
})
}), gopurs_runtime.Func2(func(f_0 gopurs_runtime.Value, z_1 gopurs_runtime.Value) gopurs_runtime.Value {
var go__2_6 gopurs_runtime.Value
_ = go__2_6
go__2_6 = gopurs_runtime.Func2(Call_go__2_6)
return gopurs_runtime.Func(func(m_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.UncurriedApp2(go__2_6, m_3, z_1)
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
mempty_1_0 := gopurs_runtime.RecordGet(dictMonoid_0, "mempty")
_ = mempty_1_0
__local_var_2_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{})
_ = __local_var_2_1
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
var go__4_2 gopurs_runtime.Value
_ = go__4_2
go__4_2 = gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
var __t3 gopurs_runtime.Value
{
if (v_5.Type == 9 && v_5.IntVal == 687041424) {
__t3 = mempty_1_0
goto end_branch_3
} else {

}
}
{
if (v_5.Type == 9 && v_5.IntVal == 324739070) {
__t3 = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_1, "append"), gopurs_runtime.Apply(go__4_2, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_5.UnsafePtr).V4), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_1, "append"), gopurs_runtime.Apply2(f_3, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_5.UnsafePtr).V2, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_5.UnsafePtr).V3), gopurs_runtime.Apply(go__4_2, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_5.UnsafePtr).V5)))
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
return go__4_2
})
}), gopurs_runtime.Func2(func(f_0 gopurs_runtime.Value, z_1 gopurs_runtime.Value) gopurs_runtime.Value {
var go__2_4 gopurs_runtime.Value
_ = go__2_4
go__2_4 = gopurs_runtime.Func2(Call_go__2_4)
return gopurs_runtime.Func(func(m_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.UncurriedApp2(go__2_4, z_1, m_3)
})
}), gopurs_runtime.Func2(func(f_0 gopurs_runtime.Value, z_1 gopurs_runtime.Value) gopurs_runtime.Value {
var go__2_6 gopurs_runtime.Value
_ = go__2_6
go__2_6 = gopurs_runtime.Func2(Call_go__2_6)
return gopurs_runtime.Func(func(m_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.UncurriedApp2(go__2_6, m_3, z_1)
})
}))
	})
	return cache_foldableWithIndexMap
}

var cache_keys gopurs_runtime.Value
var once_keys sync.Once
func Get_keys() gopurs_runtime.Value {
	once_keys.Do(func() {
		cache_keys = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_foldableWithIndexMap(), "foldrWithIndex"), gopurs_runtime.Func3(func(k_0 gopurs_runtime.Value, v_1 gopurs_runtime.Value, acc_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]{k_0, acc_2})}
}), gopurs_runtime.Value{Type: 9, IntVal: 786377863, UnsafePtr: nil})
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
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_traversableMap(), "traverse"), dictApplicative_0, Get_identity())
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
Apply0_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{})
_ = Apply0_1_0
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
var go__3_1 gopurs_runtime.Value
_ = go__3_1
go__3_1 = gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 gopurs_runtime.Value
{
if (v_4.Type == 9 && v_4.IntVal == 687041424) {
__t2 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "pure"), gopurs_runtime.Value{Type: 9, IntVal: 687041424, UnsafePtr: nil})
goto end_branch_2
} else {

}
}
{
if (v_4.Type == 9 && v_4.IntVal == 324739070) {
__local_var_5_3 := gopurs_runtime.Int((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V0)
_ = __local_var_5_3
__local_var_6_4 := (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V2
_ = __local_var_6_4
__local_var_7_5 := gopurs_runtime.Int((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V1)
_ = __local_var_7_5
__t2 = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Apply0_1_0, "apply"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Apply0_1_0, "apply"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Apply0_1_0, "Functor0"), gopurs_runtime.Value{}), "map"), gopurs_runtime.Func3(func(l_prime_8 gopurs_runtime.Value, v_prime_9 gopurs_runtime.Value, r_prime_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(&Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]{__local_var_5_3.IntVal, __local_var_7_5.IntVal, __local_var_6_4, v_prime_9, l_prime_8, r_prime_10})}
}), gopurs_runtime.Apply(go__3_1, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V4)), gopurs_runtime.Apply(f_2, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V3)), gopurs_runtime.Apply(go__3_1, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V5))
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
return go__3_1
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
Apply0_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{})
_ = Apply0_1_0
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
var go__3_1 gopurs_runtime.Value
_ = go__3_1
go__3_1 = gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 gopurs_runtime.Value
{
if (v_4.Type == 9 && v_4.IntVal == 687041424) {
__t2 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "pure"), gopurs_runtime.Value{Type: 9, IntVal: 687041424, UnsafePtr: nil})
goto end_branch_2
} else {

}
}
{
if (v_4.Type == 9 && v_4.IntVal == 324739070) {
__local_var_5_3 := gopurs_runtime.Int((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V0)
_ = __local_var_5_3
__local_var_6_4 := (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V2
_ = __local_var_6_4
__local_var_7_5 := gopurs_runtime.Int((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V1)
_ = __local_var_7_5
__t2 = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Apply0_1_0, "apply"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Apply0_1_0, "apply"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Apply0_1_0, "Functor0"), gopurs_runtime.Value{}), "map"), gopurs_runtime.Func3(func(l_prime_8 gopurs_runtime.Value, v_prime_9 gopurs_runtime.Value, r_prime_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(&Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]{__local_var_5_3.IntVal, __local_var_7_5.IntVal, __local_var_6_4, v_prime_9, l_prime_8, r_prime_10})}
}), gopurs_runtime.Apply(go__3_1, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V4)), gopurs_runtime.Apply2(f_2, __local_var_6_4, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V3)), gopurs_runtime.Apply(go__3_1, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V5))
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
return go__3_1
})
}))
	})
	return cache_traversableWithIndexMap
}

var cache_values gopurs_runtime.Value
var once_values sync.Once
func Get_values() gopurs_runtime.Value {
	once_values.Do(func() {
		cache_values = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_foldableMap(), "foldr"), pkg_Data_List_Types.Get_Cons(), gopurs_runtime.Value{Type: 9, IntVal: 786377863, UnsafePtr: nil})
	})
	return cache_values
}

var cache_foldSubmapBy gopurs_runtime.Value
var once_foldSubmapBy sync.Once
func Get_foldSubmapBy() gopurs_runtime.Value {
	once_foldSubmapBy.Do(func() {
		cache_foldSubmapBy = gopurs_runtime.Func6(func(dictOrd_0_box gopurs_runtime.Value, appendFn_1_box gopurs_runtime.Value, memptyValue_2_box gopurs_runtime.Value, kmin_3_box gopurs_runtime.Value, kmax_4_box gopurs_runtime.Value, f_5_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldSubmapBy(dictOrd_0_box, appendFn_1_box, memptyValue_2_box, kmin_3_box, kmax_4_box, f_5_box)
})
	})
	return cache_foldSubmapBy
}

var cache_foldSubmap gopurs_runtime.Value
var once_foldSubmap sync.Once
func Get_foldSubmap() gopurs_runtime.Value {
	once_foldSubmap.Do(func() {
		cache_foldSubmap = gopurs_runtime.Func2(func(dictOrd_0_box gopurs_runtime.Value, dictMonoid_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldSubmap(dictOrd_0_box, dictMonoid_1_box)
})
	})
	return cache_foldSubmap
}

var cache_findMin gopurs_runtime.Value
var once_findMin sync.Once
func Get_findMin() gopurs_runtime.Value {
	once_findMin.Do(func() {
		cache_findMin = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_findMin(v_0_box)
})
	})
	return cache_findMin
}

var cache_lookupGT gopurs_runtime.Value
var once_lookupGT sync.Once
func Get_lookupGT() gopurs_runtime.Value {
	once_lookupGT.Do(func() {
		cache_lookupGT = gopurs_runtime.Func2(func(dictOrd_0_box gopurs_runtime.Value, k_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_lookupGT(dictOrd_0_box, k_1_box)
})
	})
	return cache_lookupGT
}

var cache_findMax gopurs_runtime.Value
var once_findMax sync.Once
func Get_findMax() gopurs_runtime.Value {
	once_findMax.Do(func() {
		cache_findMax = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_findMax(v_0_box)
})
	})
	return cache_findMax
}

var cache_lookupLT gopurs_runtime.Value
var once_lookupLT sync.Once
func Get_lookupLT() gopurs_runtime.Value {
	once_lookupLT.Do(func() {
		cache_lookupLT = gopurs_runtime.Func2(func(dictOrd_0_box gopurs_runtime.Value, k_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_lookupLT(dictOrd_0_box, k_1_box)
})
	})
	return cache_lookupLT
}

var cache_filterWithKey gopurs_runtime.Value
var once_filterWithKey sync.Once
func Get_filterWithKey() gopurs_runtime.Value {
	once_filterWithKey.Do(func() {
		cache_filterWithKey = gopurs_runtime.Func2(func(dictOrd_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_filterWithKey(dictOrd_0_box, f_1_box)
})
	})
	return cache_filterWithKey
}

var cache_filterKeys gopurs_runtime.Value
var once_filterKeys sync.Once
func Get_filterKeys() gopurs_runtime.Value {
	once_filterKeys.Do(func() {
		cache_filterKeys = gopurs_runtime.Func2(func(dictOrd_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_filterKeys(dictOrd_0_box, f_1_box)
})
	})
	return cache_filterKeys
}

var cache_filter gopurs_runtime.Value
var once_filter sync.Once
func Get_filter() gopurs_runtime.Value {
	once_filter.Do(func() {
		cache_filter = gopurs_runtime.Func2(func(dictOrd_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_filter(dictOrd_0_box, x_1_box)
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
		cache_empty = gopurs_runtime.Value{Type: 9, IntVal: 687041424, UnsafePtr: nil}
	})
	return cache_empty
}

var cache_fromFoldable gopurs_runtime.Value
var once_fromFoldable sync.Once
func Get_fromFoldable() gopurs_runtime.Value {
	once_fromFoldable.Do(func() {
		cache_fromFoldable = gopurs_runtime.Func2(func(dictOrd_0_box gopurs_runtime.Value, dictFoldable_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_fromFoldable(dictOrd_0_box, dictFoldable_1_box)
})
	})
	return cache_fromFoldable
}

var cache_fromFoldableWith gopurs_runtime.Value
var once_fromFoldableWith sync.Once
func Get_fromFoldableWith() gopurs_runtime.Value {
	once_fromFoldableWith.Do(func() {
		cache_fromFoldableWith = gopurs_runtime.Func3(func(dictOrd_0_box gopurs_runtime.Value, dictFoldable_1_box gopurs_runtime.Value, f_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_fromFoldableWith(dictOrd_0_box, dictFoldable_1_box, f_2_box)
})
	})
	return cache_fromFoldableWith
}

var cache_fromFoldableWithIndex gopurs_runtime.Value
var once_fromFoldableWithIndex sync.Once
func Get_fromFoldableWithIndex() gopurs_runtime.Value {
	once_fromFoldableWithIndex.Do(func() {
		cache_fromFoldableWithIndex = gopurs_runtime.Func2(func(dictOrd_0_box gopurs_runtime.Value, dictFoldableWithIndex_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_fromFoldableWithIndex(dictOrd_0_box, dictFoldableWithIndex_1_box)
})
	})
	return cache_fromFoldableWithIndex
}

var cache_monoidSemigroupMap gopurs_runtime.Value
var once_monoidSemigroupMap sync.Once
func Get_monoidSemigroupMap() gopurs_runtime.Value {
	once_monoidSemigroupMap.Do(func() {
		cache_monoidSemigroupMap = gopurs_runtime.Func2(func(_dollar__unused_0_box gopurs_runtime.Value, dictOrd_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_monoidSemigroupMap(_dollar__unused_0_box, dictOrd_1_box)
})
	})
	return cache_monoidSemigroupMap
}

var cache_submap gopurs_runtime.Value
var once_submap sync.Once
func Get_submap() gopurs_runtime.Value {
	once_submap.Do(func() {
		cache_submap = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_submap(dictOrd_0_box)
})
	})
	return cache_submap
}

var cache_unions gopurs_runtime.Value
var once_unions sync.Once
func Get_unions() gopurs_runtime.Value {
	once_unions.Do(func() {
		cache_unions = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_unions(dictOrd_0_box)
})
	})
	return cache_unions
}

var cache_difference gopurs_runtime.Value
var once_difference sync.Once
func Get_difference() gopurs_runtime.Value {
	once_difference.Do(func() {
		cache_difference = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_difference(dictOrd_0_box)
})
	})
	return cache_difference
}

var cache_delete_ gopurs_runtime.Value
var once_delete_ sync.Once
func Get_delete_() gopurs_runtime.Value {
	once_delete_.Do(func() {
		cache_delete_ = gopurs_runtime.Func2(func(dictOrd_0_box gopurs_runtime.Value, k_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_delete_(dictOrd_0_box, k_1_box)
})
	})
	return cache_delete_
}

var cache_checkValid gopurs_runtime.Value
var once_checkValid sync.Once
func Get_checkValid() gopurs_runtime.Value {
	once_checkValid.Do(func() {
		cache_checkValid = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_checkValid(dictOrd_0_box)
})
	})
	return cache_checkValid
}

var cache_catMaybes gopurs_runtime.Value
var once_catMaybes sync.Once
func Get_catMaybes() gopurs_runtime.Value {
	once_catMaybes.Do(func() {
		cache_catMaybes = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_catMaybes(dictOrd_0_box)
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
return Call_alter(dictOrd_0_box)
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

type Constructor_Leaf[T_k any, T_v any] struct {
	
}


type Constructor_Node[T_k any, T_v any] struct {
	V0 int64
	V1 int64
	V2 T_k
	V3 T_v
	V4 *Constructor_Node
	V5 *Constructor_Node
}


type Constructor_IterLeaf[T_k any, T_v any] struct {
	
}


type Constructor_IterEmit[T_k any, T_v any] struct {
	V0 T_k
	V1 T_v
	V2 gopurs_runtime.Value
}


type Constructor_IterNode[T_k any, T_v any] struct {
	V0 *Constructor_Node
	V1 gopurs_runtime.Value
}


type Constructor_IterDone[T_k any, T_v any] struct {
	
}


type Constructor_IterNext[T_k any, T_v any] struct {
	V0 T_k
	V1 T_v
	V2 gopurs_runtime.Value
}


type Constructor_Split[T_k any, T_v any] struct {
	V0 *pkg_Data_Maybe.Constructor_Just
	V1 *Constructor_Node
	V2 *Constructor_Node
}


type Constructor_SplitLast[T_k any, T_v any] struct {
	V0 T_k
	V1 T_v
	V2 *Constructor_Node
}


func Call_identity(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_unsafeNode(k_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value, l_2_loop gopurs_runtime.Value, r_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
var k_0 gopurs_runtime.Value = k_0_loop
_ = k_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
var l_2 gopurs_runtime.Value = l_2_loop
_ = l_2
var r_3 gopurs_runtime.Value = r_3_loop
_ = r_3
var __t0 gopurs_runtime.Value
{
if (l_2.Type == 9 && l_2.IntVal == 687041424) {
var __t1 gopurs_runtime.Value
{
if (r_3.Type == 9 && r_3.IntVal == 687041424) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(&Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]{1, 1, k_0, v_1, l_2, r_3})}
goto end_branch_1
} else {

}
}
{
if (r_3.Type == 9 && r_3.IntVal == 324739070) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(&Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]{(1) + (gopurs_runtime.Int((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(r_3.UnsafePtr).V0).IntVal), (1) + (gopurs_runtime.Int((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(r_3.UnsafePtr).V1).IntVal), k_0, v_1, l_2, r_3})}
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
__t0 = __t1
goto end_branch_0
} else {

}
}
{
if (l_2.Type == 9 && l_2.IntVal == 324739070) {
var __t2 gopurs_runtime.Value
{
if (r_3.Type == 9 && r_3.IntVal == 687041424) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(&Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]{(1) + (gopurs_runtime.Int((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(l_2.UnsafePtr).V0).IntVal), (1) + (gopurs_runtime.Int((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(l_2.UnsafePtr).V1).IntVal), k_0, v_1, l_2, r_3})}
goto end_branch_2
} else {

}
}
{
if (r_3.Type == 9 && r_3.IntVal == 324739070) {
var __t3 gopurs_runtime.Value
{
if (gopurs_runtime.Apply2(Get_greaterThan(), gopurs_runtime.Int((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(l_2.UnsafePtr).V0), gopurs_runtime.Int((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(r_3.UnsafePtr).V0)).IntVal) != (0) {
__t3 = gopurs_runtime.Int((1) + (gopurs_runtime.Int((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(l_2.UnsafePtr).V0).IntVal))
goto end_branch_3
} else {

}
}
{
__t3 = gopurs_runtime.Int((1) + (gopurs_runtime.Int((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(r_3.UnsafePtr).V0).IntVal))
}
end_branch_3:
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(&Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]{__t3.IntVal, ((1) + (gopurs_runtime.Int((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(l_2.UnsafePtr).V1).IntVal)) + (gopurs_runtime.Int((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(r_3.UnsafePtr).V1).IntVal), k_0, v_1, l_2, r_3})}
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_2:
__t0 = __t2
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

func Call_toMapIter(a_0_loop *Constructor_Node) gopurs_runtime.Value {
var a_0 *Constructor_Node = a_0_loop
_ = a_0
return gopurs_runtime.Value{Type: 9, IntVal: 2861335956, UnsafePtr: unsafe.Pointer(&Constructor_IterNode[gopurs_runtime.Value, gopurs_runtime.Value]{a_0, gopurs_runtime.Value{Type: 9, IntVal: 2509360378, UnsafePtr: nil}})}
}

func Call_stepWith(f_0_loop gopurs_runtime.Value, next_1_loop gopurs_runtime.Value, done_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var next_1 gopurs_runtime.Value = next_1_loop
_ = next_1
var done_2 gopurs_runtime.Value = done_2_loop
_ = done_2
var go__3_0 gopurs_runtime.Value
go__3_0 = gopurs_runtime.Func(func(v_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_4_loop gopurs_runtime.Value = v_4_loop_val
go__3_0:
for {
if false { continue go__3_0 }
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
v_4_loop = gopurs_runtime.Apply2(f_0, (*Constructor_IterNode[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V1, (*Constructor_IterNode[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V0)
continue go__3_0
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
return go__3_0
}

func Call_size(v_0_loop *Constructor_Node) gopurs_runtime.Value {
var v_0 *Constructor_Node = v_0_loop
_ = v_0
var __t0 gopurs_runtime.Value
{
if (v_0 == nil) {
__t0 = gopurs_runtime.Int(0)
goto end_branch_0
} else {

}
}
{
if (v_0 != nil) {
__t0 = gopurs_runtime.Int((v_0).V1)
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

func Call_singleton(k_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var k_0 gopurs_runtime.Value = k_0_loop
_ = k_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(&Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]{1, 1, k_0, v_1, gopurs_runtime.Value{Type: 9, IntVal: 687041424, UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: 687041424, UnsafePtr: nil}})}
}

func Call_unsafeBalancedNode(k_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value, l_2_loop gopurs_runtime.Value, r_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
var k_0 gopurs_runtime.Value = k_0_loop
_ = k_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
var l_2 gopurs_runtime.Value = l_2_loop
_ = l_2
var r_3 gopurs_runtime.Value = r_3_loop
_ = r_3
var __t0 gopurs_runtime.Value
{
if (l_2.Type == 9 && l_2.IntVal == 687041424) {
var __t1 gopurs_runtime.Value
{
if (r_3.Type == 9 && r_3.IntVal == 687041424) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(&Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]{1, 1, k_0, v_1, gopurs_runtime.Value{Type: 9, IntVal: 687041424, UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: 687041424, UnsafePtr: nil}})}
goto end_branch_1
} else {

}
}
{
if ((r_3.Type == 9 && r_3.IntVal == 324739070)) && ((gopurs_runtime.Apply2(Get_greaterThan(), gopurs_runtime.Int((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(r_3.UnsafePtr).V0), gopurs_runtime.Int(1)).IntVal) != (0)) {
var __t2 gopurs_runtime.Value
{
var __t_tag_3 gopurs_runtime.Value = (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(r_3.UnsafePtr).V4
var __t_and_7 bool = false
if (__t_tag_3.Type == 9 && __t_tag_3.IntVal == 324739070) {

var __t4 gopurs_runtime.Value
{
var __t_tag_5 gopurs_runtime.Value = (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(r_3.UnsafePtr).V5
if (__t_tag_5.Type == 9 && __t_tag_5.IntVal == 687041424) {
__t4 = gopurs_runtime.Int(0)
goto end_branch_4
} else {

}
}
{
var __t_tag_6 gopurs_runtime.Value = (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(r_3.UnsafePtr).V5
if (__t_tag_6.Type == 9 && __t_tag_6.IntVal == 324739070) {
__t4 = gopurs_runtime.Int((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(r_3.UnsafePtr).V5.UnsafePtr).V0)
goto end_branch_4
} else {

}
}
{
__t4 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_4:
__t_and_7 = (gopurs_runtime.Apply2(Get_greaterThan(), gopurs_runtime.Int((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(r_3.UnsafePtr).V4.UnsafePtr).V0), __t4).IntVal) != (0)
}
if __t_and_7 {
__t2 = gopurs_runtime.UncurriedApp4(Get_unsafeNode(), (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(r_3.UnsafePtr).V4.UnsafePtr).V2, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(r_3.UnsafePtr).V4.UnsafePtr).V3, gopurs_runtime.UncurriedApp4(Get_unsafeNode(), k_0, v_1, l_2, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(r_3.UnsafePtr).V4.UnsafePtr).V4), gopurs_runtime.UncurriedApp4(Get_unsafeNode(), (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(r_3.UnsafePtr).V2, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(r_3.UnsafePtr).V3, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(r_3.UnsafePtr).V4.UnsafePtr).V5, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(r_3.UnsafePtr).V5))
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.UncurriedApp4(Get_unsafeNode(), (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(r_3.UnsafePtr).V2, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(r_3.UnsafePtr).V3, gopurs_runtime.UncurriedApp4(Get_unsafeNode(), k_0, v_1, l_2, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(r_3.UnsafePtr).V4), (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(r_3.UnsafePtr).V5)
}
end_branch_2:
__t1 = __t2
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.UncurriedApp4(Get_unsafeNode(), k_0, v_1, l_2, r_3)
}
end_branch_1:
__t0 = __t1
goto end_branch_0
} else {

}
}
{
if (l_2.Type == 9 && l_2.IntVal == 324739070) {
var __t8 gopurs_runtime.Value
{
if (r_3.Type == 9 && r_3.IntVal == 324739070) {
var __t9 gopurs_runtime.Value
{
if (gopurs_runtime.Apply2(Get_greaterThan(), gopurs_runtime.Int((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(r_3.UnsafePtr).V0), gopurs_runtime.Int((gopurs_runtime.Int((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(l_2.UnsafePtr).V0).IntVal) + (1))).IntVal) != (0) {
var __t10 gopurs_runtime.Value
{
var __t_tag_11 gopurs_runtime.Value = (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(r_3.UnsafePtr).V4
var __t_and_15 bool = false
if (__t_tag_11.Type == 9 && __t_tag_11.IntVal == 324739070) {

var __t12 gopurs_runtime.Value
{
var __t_tag_13 gopurs_runtime.Value = (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(r_3.UnsafePtr).V5
if (__t_tag_13.Type == 9 && __t_tag_13.IntVal == 687041424) {
__t12 = gopurs_runtime.Int(0)
goto end_branch_12
} else {

}
}
{
var __t_tag_14 gopurs_runtime.Value = (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(r_3.UnsafePtr).V5
if (__t_tag_14.Type == 9 && __t_tag_14.IntVal == 324739070) {
__t12 = gopurs_runtime.Int((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(r_3.UnsafePtr).V5.UnsafePtr).V0)
goto end_branch_12
} else {

}
}
{
__t12 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_12:
__t_and_15 = (gopurs_runtime.Apply2(Get_greaterThan(), gopurs_runtime.Int((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(r_3.UnsafePtr).V4.UnsafePtr).V0), __t12).IntVal) != (0)
}
if __t_and_15 {
__t10 = gopurs_runtime.UncurriedApp4(Get_unsafeNode(), (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(r_3.UnsafePtr).V4.UnsafePtr).V2, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(r_3.UnsafePtr).V4.UnsafePtr).V3, gopurs_runtime.UncurriedApp4(Get_unsafeNode(), k_0, v_1, l_2, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(r_3.UnsafePtr).V4.UnsafePtr).V4), gopurs_runtime.UncurriedApp4(Get_unsafeNode(), (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(r_3.UnsafePtr).V2, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(r_3.UnsafePtr).V3, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(r_3.UnsafePtr).V4.UnsafePtr).V5, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(r_3.UnsafePtr).V5))
goto end_branch_10
} else {

}
}
{
__t10 = gopurs_runtime.UncurriedApp4(Get_unsafeNode(), (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(r_3.UnsafePtr).V2, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(r_3.UnsafePtr).V3, gopurs_runtime.UncurriedApp4(Get_unsafeNode(), k_0, v_1, l_2, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(r_3.UnsafePtr).V4), (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(r_3.UnsafePtr).V5)
}
end_branch_10:
__t9 = __t10
goto end_branch_9
} else {

}
}
{
if (gopurs_runtime.Apply2(Get_greaterThan(), gopurs_runtime.Int((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(l_2.UnsafePtr).V0), gopurs_runtime.Int((gopurs_runtime.Int((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(r_3.UnsafePtr).V0).IntVal) + (1))).IntVal) != (0) {
var __t16 gopurs_runtime.Value
{
var __t_tag_17 gopurs_runtime.Value = (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(l_2.UnsafePtr).V5
var __t_and_21 bool = false
if (__t_tag_17.Type == 9 && __t_tag_17.IntVal == 324739070) {

var __t18 gopurs_runtime.Value
{
var __t_tag_19 gopurs_runtime.Value = (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(l_2.UnsafePtr).V4
if (__t_tag_19.Type == 9 && __t_tag_19.IntVal == 687041424) {
__t18 = gopurs_runtime.Int(0)
goto end_branch_18
} else {

}
}
{
var __t_tag_20 gopurs_runtime.Value = (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(l_2.UnsafePtr).V4
if (__t_tag_20.Type == 9 && __t_tag_20.IntVal == 324739070) {
__t18 = gopurs_runtime.Int((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(l_2.UnsafePtr).V4.UnsafePtr).V0)
goto end_branch_18
} else {

}
}
{
__t18 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_18:
__t_and_21 = (gopurs_runtime.Apply2(Get_lessThanOrEq(), __t18, gopurs_runtime.Int((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(l_2.UnsafePtr).V5.UnsafePtr).V0)).IntVal) != (0)
}
if __t_and_21 {
__t16 = gopurs_runtime.UncurriedApp4(Get_unsafeNode(), (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(l_2.UnsafePtr).V5.UnsafePtr).V2, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(l_2.UnsafePtr).V5.UnsafePtr).V3, gopurs_runtime.UncurriedApp4(Get_unsafeNode(), (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(l_2.UnsafePtr).V2, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(l_2.UnsafePtr).V3, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(l_2.UnsafePtr).V4, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(l_2.UnsafePtr).V5.UnsafePtr).V4), gopurs_runtime.UncurriedApp4(Get_unsafeNode(), k_0, v_1, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(l_2.UnsafePtr).V5.UnsafePtr).V5, r_3))
goto end_branch_16
} else {

}
}
{
__t16 = gopurs_runtime.UncurriedApp4(Get_unsafeNode(), (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(l_2.UnsafePtr).V2, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(l_2.UnsafePtr).V3, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(l_2.UnsafePtr).V4, gopurs_runtime.UncurriedApp4(Get_unsafeNode(), k_0, v_1, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(l_2.UnsafePtr).V5, r_3))
}
end_branch_16:
__t9 = __t16
goto end_branch_9
} else {

}
}
{
__t9 = gopurs_runtime.UncurriedApp4(Get_unsafeNode(), k_0, v_1, l_2, r_3)
}
end_branch_9:
__t8 = __t9
goto end_branch_8
} else {

}
}
{
if ((r_3.Type == 9 && r_3.IntVal == 687041424)) && ((gopurs_runtime.Apply2(Get_greaterThan(), gopurs_runtime.Int((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(l_2.UnsafePtr).V0), gopurs_runtime.Int(1)).IntVal) != (0)) {
var __t22 gopurs_runtime.Value
{
var __t_tag_23 gopurs_runtime.Value = (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(l_2.UnsafePtr).V5
var __t_and_27 bool = false
if (__t_tag_23.Type == 9 && __t_tag_23.IntVal == 324739070) {

var __t24 gopurs_runtime.Value
{
var __t_tag_25 gopurs_runtime.Value = (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(l_2.UnsafePtr).V4
if (__t_tag_25.Type == 9 && __t_tag_25.IntVal == 687041424) {
__t24 = gopurs_runtime.Int(0)
goto end_branch_24
} else {

}
}
{
var __t_tag_26 gopurs_runtime.Value = (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(l_2.UnsafePtr).V4
if (__t_tag_26.Type == 9 && __t_tag_26.IntVal == 324739070) {
__t24 = gopurs_runtime.Int((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(l_2.UnsafePtr).V4.UnsafePtr).V0)
goto end_branch_24
} else {

}
}
{
__t24 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_24:
__t_and_27 = (gopurs_runtime.Apply2(Get_lessThanOrEq(), __t24, gopurs_runtime.Int((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(l_2.UnsafePtr).V5.UnsafePtr).V0)).IntVal) != (0)
}
if __t_and_27 {
__t22 = gopurs_runtime.UncurriedApp4(Get_unsafeNode(), (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(l_2.UnsafePtr).V5.UnsafePtr).V2, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(l_2.UnsafePtr).V5.UnsafePtr).V3, gopurs_runtime.UncurriedApp4(Get_unsafeNode(), (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(l_2.UnsafePtr).V2, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(l_2.UnsafePtr).V3, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(l_2.UnsafePtr).V4, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(l_2.UnsafePtr).V5.UnsafePtr).V4), gopurs_runtime.UncurriedApp4(Get_unsafeNode(), k_0, v_1, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(l_2.UnsafePtr).V5.UnsafePtr).V5, r_3))
goto end_branch_22
} else {

}
}
{
__t22 = gopurs_runtime.UncurriedApp4(Get_unsafeNode(), (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(l_2.UnsafePtr).V2, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(l_2.UnsafePtr).V3, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(l_2.UnsafePtr).V4, gopurs_runtime.UncurriedApp4(Get_unsafeNode(), k_0, v_1, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(l_2.UnsafePtr).V5, r_3))
}
end_branch_22:
__t8 = __t22
goto end_branch_8
} else {

}
}
{
__t8 = gopurs_runtime.UncurriedApp4(Get_unsafeNode(), k_0, v_1, l_2, r_3)
}
end_branch_8:
__t0 = __t8
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

func Call_unsafeSplit(comp_0_loop gopurs_runtime.Value, k_1_loop gopurs_runtime.Value, m_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
unsafeSplit:
for {
if false { continue unsafeSplit }
var comp_0 gopurs_runtime.Value = comp_0_loop
_ = comp_0
var k_1 gopurs_runtime.Value = k_1_loop
_ = k_1
var m_2 gopurs_runtime.Value = m_2_loop
_ = m_2
var __t0 gopurs_runtime.Value
{
if (m_2.Type == 9 && m_2.IntVal == 687041424) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 3373277644, UnsafePtr: unsafe.Pointer(&Constructor_Split[gopurs_runtime.Value, gopurs_runtime.Value]{gopurs_runtime.Value{Type: 9, IntVal: 3589588149, UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: 687041424, UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: 687041424, UnsafePtr: nil}})}
goto end_branch_0
} else {

}
}
{
if (m_2.Type == 9 && m_2.IntVal == 324739070) {
v_3_1 := gopurs_runtime.Apply2(comp_0, k_1, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(m_2.UnsafePtr).V2)
_ = v_3_1
var __t2 gopurs_runtime.Value
{
if (v_3_1.Type == 9 && v_3_1.IntVal == 1527465420) {
v1_4_3 := gopurs_runtime.UncurriedApp3(Get_unsafeSplit(), comp_0, k_1, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(m_2.UnsafePtr).V4)
_ = v1_4_3
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 3373277644, UnsafePtr: unsafe.Pointer(&Constructor_Split[gopurs_runtime.Value, gopurs_runtime.Value]{(*Constructor_Split[gopurs_runtime.Value, gopurs_runtime.Value])(v1_4_3.UnsafePtr).V0, (*Constructor_Split[gopurs_runtime.Value, gopurs_runtime.Value])(v1_4_3.UnsafePtr).V1, gopurs_runtime.UncurriedApp4(Get_unsafeBalancedNode(), (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(m_2.UnsafePtr).V2, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(m_2.UnsafePtr).V3, (*Constructor_Split[gopurs_runtime.Value, gopurs_runtime.Value])(v1_4_3.UnsafePtr).V2, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(m_2.UnsafePtr).V5)})}
goto end_branch_2
} else {

}
}
{
if (v_3_1.Type == 9 && v_3_1.IntVal == 380165415) {
v1_4_4 := gopurs_runtime.UncurriedApp3(Get_unsafeSplit(), comp_0, k_1, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(m_2.UnsafePtr).V5)
_ = v1_4_4
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 3373277644, UnsafePtr: unsafe.Pointer(&Constructor_Split[gopurs_runtime.Value, gopurs_runtime.Value]{(*Constructor_Split[gopurs_runtime.Value, gopurs_runtime.Value])(v1_4_4.UnsafePtr).V0, gopurs_runtime.UncurriedApp4(Get_unsafeBalancedNode(), (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(m_2.UnsafePtr).V2, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(m_2.UnsafePtr).V3, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(m_2.UnsafePtr).V4, (*Constructor_Split[gopurs_runtime.Value, gopurs_runtime.Value])(v1_4_4.UnsafePtr).V1), (*Constructor_Split[gopurs_runtime.Value, gopurs_runtime.Value])(v1_4_4.UnsafePtr).V2})}
goto end_branch_2
} else {

}
}
{
if (v_3_1.Type == 9 && v_3_1.IntVal == 902936544) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 3373277644, UnsafePtr: unsafe.Pointer(&Constructor_Split[gopurs_runtime.Value, gopurs_runtime.Value]{gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{(*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(m_2.UnsafePtr).V3})}, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(m_2.UnsafePtr).V4, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(m_2.UnsafePtr).V5})}
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_2:
__t0 = __t2
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
}

func Call_unsafeSplitLast(k_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value, l_2_loop gopurs_runtime.Value, r_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
unsafeSplitLast:
for {
if false { continue unsafeSplitLast }
var k_0 gopurs_runtime.Value = k_0_loop
_ = k_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
var l_2 gopurs_runtime.Value = l_2_loop
_ = l_2
var r_3 gopurs_runtime.Value = r_3_loop
_ = r_3
var __t0 gopurs_runtime.Value
{
if (r_3.Type == 9 && r_3.IntVal == 687041424) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 2668112006, UnsafePtr: unsafe.Pointer(&Constructor_SplitLast[gopurs_runtime.Value, gopurs_runtime.Value]{k_0, v_1, l_2})}
goto end_branch_0
} else {

}
}
{
if (r_3.Type == 9 && r_3.IntVal == 324739070) {
v1_4_1 := gopurs_runtime.UncurriedApp4(Get_unsafeSplitLast(), (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(r_3.UnsafePtr).V2, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(r_3.UnsafePtr).V3, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(r_3.UnsafePtr).V4, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(r_3.UnsafePtr).V5)
_ = v1_4_1
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 2668112006, UnsafePtr: unsafe.Pointer(&Constructor_SplitLast[gopurs_runtime.Value, gopurs_runtime.Value]{(*Constructor_SplitLast[gopurs_runtime.Value, gopurs_runtime.Value])(v1_4_1.UnsafePtr).V0, (*Constructor_SplitLast[gopurs_runtime.Value, gopurs_runtime.Value])(v1_4_1.UnsafePtr).V1, gopurs_runtime.UncurriedApp4(Get_unsafeBalancedNode(), k_0, v_1, l_2, (*Constructor_SplitLast[gopurs_runtime.Value, gopurs_runtime.Value])(v1_4_1.UnsafePtr).V2)})}
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
}

func Call_unsafeJoinNodes(v_0_loop gopurs_runtime.Value, v1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var v1_1 gopurs_runtime.Value = v1_1_loop
_ = v1_1
var __t0 gopurs_runtime.Value
{
if (v_0.Type == 9 && v_0.IntVal == 687041424) {
__t0 = v1_1
goto end_branch_0
} else {

}
}
{
if (v_0.Type == 9 && v_0.IntVal == 324739070) {
v2_2_1 := gopurs_runtime.UncurriedApp4(Get_unsafeSplitLast(), (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_0.UnsafePtr).V2, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_0.UnsafePtr).V3, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_0.UnsafePtr).V4, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_0.UnsafePtr).V5)
_ = v2_2_1
__t0 = gopurs_runtime.UncurriedApp4(Get_unsafeBalancedNode(), (*Constructor_SplitLast[gopurs_runtime.Value, gopurs_runtime.Value])(v2_2_1.UnsafePtr).V0, (*Constructor_SplitLast[gopurs_runtime.Value, gopurs_runtime.Value])(v2_2_1.UnsafePtr).V1, (*Constructor_SplitLast[gopurs_runtime.Value, gopurs_runtime.Value])(v2_2_1.UnsafePtr).V2, v1_1)
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

func Call_unsafeDifference(comp_0_loop gopurs_runtime.Value, l_1_loop gopurs_runtime.Value, r_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
unsafeDifference:
for {
if false { continue unsafeDifference }
var comp_0 gopurs_runtime.Value = comp_0_loop
_ = comp_0
var l_1 gopurs_runtime.Value = l_1_loop
_ = l_1
var r_2 gopurs_runtime.Value = r_2_loop
_ = r_2
var __t0 gopurs_runtime.Value
{
if (l_1.Type == 9 && l_1.IntVal == 687041424) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 687041424, UnsafePtr: nil}
goto end_branch_0
} else {

}
}
{
if (r_2.Type == 9 && r_2.IntVal == 687041424) {
__t0 = l_1
goto end_branch_0
} else {

}
}
{
if (r_2.Type == 9 && r_2.IntVal == 324739070) {
v_3_1 := gopurs_runtime.UncurriedApp3(Get_unsafeSplit(), comp_0, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(r_2.UnsafePtr).V2, l_1)
_ = v_3_1
__t0 = gopurs_runtime.UncurriedApp2(Get_unsafeJoinNodes(), gopurs_runtime.UncurriedApp3(Get_unsafeDifference(), comp_0, (*Constructor_Split[gopurs_runtime.Value, gopurs_runtime.Value])(v_3_1.UnsafePtr).V1, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(r_2.UnsafePtr).V4), gopurs_runtime.UncurriedApp3(Get_unsafeDifference(), comp_0, (*Constructor_Split[gopurs_runtime.Value, gopurs_runtime.Value])(v_3_1.UnsafePtr).V2, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(r_2.UnsafePtr).V5))
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
}

func Call_unsafeIntersectionWith(comp_0_loop gopurs_runtime.Value, app_1_loop gopurs_runtime.Value, l_2_loop gopurs_runtime.Value, r_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
unsafeIntersectionWith:
for {
if false { continue unsafeIntersectionWith }
var comp_0 gopurs_runtime.Value = comp_0_loop
_ = comp_0
var app_1 gopurs_runtime.Value = app_1_loop
_ = app_1
var l_2 gopurs_runtime.Value = l_2_loop
_ = l_2
var r_3 gopurs_runtime.Value = r_3_loop
_ = r_3
var __t0 gopurs_runtime.Value
{
if (l_2.Type == 9 && l_2.IntVal == 687041424) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 687041424, UnsafePtr: nil}
goto end_branch_0
} else {

}
}
{
if (r_3.Type == 9 && r_3.IntVal == 687041424) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 687041424, UnsafePtr: nil}
goto end_branch_0
} else {

}
}
{
if (r_3.Type == 9 && r_3.IntVal == 324739070) {
v_4_1 := gopurs_runtime.UncurriedApp3(Get_unsafeSplit(), comp_0, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(r_3.UnsafePtr).V2, l_2)
_ = v_4_1
l_prime_5_2 := gopurs_runtime.UncurriedApp4(Get_unsafeIntersectionWith(), comp_0, app_1, (*Constructor_Split[gopurs_runtime.Value, gopurs_runtime.Value])(v_4_1.UnsafePtr).V1, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(r_3.UnsafePtr).V4)
_ = l_prime_5_2
r_prime_6_3 := gopurs_runtime.UncurriedApp4(Get_unsafeIntersectionWith(), comp_0, app_1, (*Constructor_Split[gopurs_runtime.Value, gopurs_runtime.Value])(v_4_1.UnsafePtr).V2, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(r_3.UnsafePtr).V5)
_ = r_prime_6_3
var __t4 gopurs_runtime.Value
{
var __t_tag_5 gopurs_runtime.Value = (*Constructor_Split[gopurs_runtime.Value, gopurs_runtime.Value])(v_4_1.UnsafePtr).V0
if (__t_tag_5.Type == 9 && __t_tag_5.IntVal == 930809136) {
__t4 = gopurs_runtime.UncurriedApp4(Get_unsafeBalancedNode(), (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(r_3.UnsafePtr).V2, gopurs_runtime.Apply2(app_1, (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])((*Constructor_Split[gopurs_runtime.Value, gopurs_runtime.Value])(v_4_1.UnsafePtr).V0.UnsafePtr).V0, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(r_3.UnsafePtr).V3), l_prime_5_2, r_prime_6_3)
goto end_branch_4
} else {

}
}
{
var __t_tag_6 gopurs_runtime.Value = (*Constructor_Split[gopurs_runtime.Value, gopurs_runtime.Value])(v_4_1.UnsafePtr).V0
if (__t_tag_6.Type == 9 && __t_tag_6.IntVal == 3589588149) {
__t4 = gopurs_runtime.UncurriedApp2(Get_unsafeJoinNodes(), l_prime_5_2, r_prime_6_3)
goto end_branch_4
} else {

}
}
{
__t4 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_4:
__t0 = __t4
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
}

func Call_unsafeUnionWith(comp_0_loop gopurs_runtime.Value, app_1_loop gopurs_runtime.Value, l_2_loop gopurs_runtime.Value, r_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
unsafeUnionWith:
for {
if false { continue unsafeUnionWith }
var comp_0 gopurs_runtime.Value = comp_0_loop
_ = comp_0
var app_1 gopurs_runtime.Value = app_1_loop
_ = app_1
var l_2 gopurs_runtime.Value = l_2_loop
_ = l_2
var r_3 gopurs_runtime.Value = r_3_loop
_ = r_3
var __t0 gopurs_runtime.Value
{
if (l_2.Type == 9 && l_2.IntVal == 687041424) {
__t0 = r_3
goto end_branch_0
} else {

}
}
{
if (r_3.Type == 9 && r_3.IntVal == 687041424) {
__t0 = l_2
goto end_branch_0
} else {

}
}
{
if (r_3.Type == 9 && r_3.IntVal == 324739070) {
v_4_1 := gopurs_runtime.UncurriedApp3(Get_unsafeSplit(), comp_0, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(r_3.UnsafePtr).V2, l_2)
_ = v_4_1
l_prime_5_2 := gopurs_runtime.UncurriedApp4(Get_unsafeUnionWith(), comp_0, app_1, (*Constructor_Split[gopurs_runtime.Value, gopurs_runtime.Value])(v_4_1.UnsafePtr).V1, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(r_3.UnsafePtr).V4)
_ = l_prime_5_2
r_prime_6_3 := gopurs_runtime.UncurriedApp4(Get_unsafeUnionWith(), comp_0, app_1, (*Constructor_Split[gopurs_runtime.Value, gopurs_runtime.Value])(v_4_1.UnsafePtr).V2, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(r_3.UnsafePtr).V5)
_ = r_prime_6_3
var __t4 gopurs_runtime.Value
{
var __t_tag_5 gopurs_runtime.Value = (*Constructor_Split[gopurs_runtime.Value, gopurs_runtime.Value])(v_4_1.UnsafePtr).V0
if (__t_tag_5.Type == 9 && __t_tag_5.IntVal == 930809136) {
__t4 = gopurs_runtime.UncurriedApp4(Get_unsafeBalancedNode(), (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(r_3.UnsafePtr).V2, gopurs_runtime.Apply2(app_1, (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])((*Constructor_Split[gopurs_runtime.Value, gopurs_runtime.Value])(v_4_1.UnsafePtr).V0.UnsafePtr).V0, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(r_3.UnsafePtr).V3), l_prime_5_2, r_prime_6_3)
goto end_branch_4
} else {

}
}
{
var __t_tag_6 gopurs_runtime.Value = (*Constructor_Split[gopurs_runtime.Value, gopurs_runtime.Value])(v_4_1.UnsafePtr).V0
if (__t_tag_6.Type == 9 && __t_tag_6.IntVal == 3589588149) {
__t4 = gopurs_runtime.UncurriedApp4(Get_unsafeBalancedNode(), (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(r_3.UnsafePtr).V2, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(r_3.UnsafePtr).V3, l_prime_5_2, r_prime_6_3)
goto end_branch_4
} else {

}
}
{
__t4 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_4:
__t0 = __t4
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
}

func Call_unionWith(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
compare_1_0 := ((*gopurs_runtime.RecordData1)(dictOrd_0.UnsafePtr)).V0
_ = compare_1_0
return gopurs_runtime.Func3(func(app_2 gopurs_runtime.Value, m1_3 gopurs_runtime.Value, m2_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.UncurriedApp4(Get_unsafeUnionWith(), compare_1_0, app_2, m1_3, m2_4)
})
}

func Call_union(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
compare_1_0 := ((*gopurs_runtime.RecordData1)(dictOrd_0.UnsafePtr)).V0
_ = compare_1_0
return gopurs_runtime.Func2(func(m1_2 gopurs_runtime.Value, m2_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.UncurriedApp4(Get_unsafeUnionWith(), compare_1_0, pkg_Data_Function.Get_const_(), m1_2, m2_3)
})
}

func Call_update(dictOrd_0_loop gopurs_runtime.Value, f_1_loop gopurs_runtime.Value, k_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var k_2 gopurs_runtime.Value = k_2_loop
_ = k_2
var go__3_0 gopurs_runtime.Value
_ = go__3_0
go__3_0 = gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if (v_4.Type == 9 && v_4.IntVal == 687041424) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 687041424, UnsafePtr: nil}
goto end_branch_1
} else {

}
}
{
if (v_4.Type == 9 && v_4.IntVal == 324739070) {
v1_5_2 := gopurs_runtime.Apply2(((*gopurs_runtime.RecordData1)(dictOrd_0.UnsafePtr)).V0, k_2, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V2)
_ = v1_5_2
var __t3 gopurs_runtime.Value
{
if (v1_5_2.Type == 9 && v1_5_2.IntVal == 1527465420) {
__t3 = gopurs_runtime.UncurriedApp4(Get_unsafeBalancedNode(), (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V2, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V3, gopurs_runtime.Apply(go__3_0, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V4), (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V5)
goto end_branch_3
} else {

}
}
{
if (v1_5_2.Type == 9 && v1_5_2.IntVal == 380165415) {
__t3 = gopurs_runtime.UncurriedApp4(Get_unsafeBalancedNode(), (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V2, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V3, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V4, gopurs_runtime.Apply(go__3_0, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V5))
goto end_branch_3
} else {

}
}
{
if (v1_5_2.Type == 9 && v1_5_2.IntVal == 902936544) {
v2_6_4 := gopurs_runtime.Apply(f_1, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V3)
_ = v2_6_4
var __t5 gopurs_runtime.Value
{
if (v2_6_4.Type == 9 && v2_6_4.IntVal == 3589588149) {
__t5 = gopurs_runtime.UncurriedApp2(Get_unsafeJoinNodes(), (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V4, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V5)
goto end_branch_5
} else {

}
}
{
if (v2_6_4.Type == 9 && v2_6_4.IntVal == 930809136) {
__t5 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(&Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]{gopurs_runtime.Int((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V0).IntVal, gopurs_runtime.Int((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V1).IntVal, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V2, (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(v2_6_4.UnsafePtr).V0, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V4, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V5})}
goto end_branch_5
} else {

}
}
{
__t5 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_5:
__t3 = __t5
goto end_branch_3
} else {

}
}
{
__t3 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_3:
__t1 = __t3
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
return go__3_0
}

func Call_showTree(dictShow_0_loop gopurs_runtime.Value, dictShow1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictShow_0 gopurs_runtime.Value = dictShow_0_loop
_ = dictShow_0
var dictShow1_1 gopurs_runtime.Value = dictShow1_1_loop
_ = dictShow1_1
var go__2_0 gopurs_runtime.Value
_ = go__2_0
go__2_0 = gopurs_runtime.Func2(func(ind_3 gopurs_runtime.Value, v_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if (v_4.Type == 9 && v_4.IntVal == 687041424) {
__t1 = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semigroup.Get_semigroupString(), "append"), ind_3, gopurs_runtime.Str("Leaf"))
goto end_branch_1
} else {

}
}
{
if (v_4.Type == 9 && v_4.IntVal == 324739070) {
__t1 = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semigroup.Get_semigroupString(), "append"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semigroup.Get_semigroupString(), "append"), ind_3, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semigroup.Get_semigroupString(), "append"), gopurs_runtime.Str("["), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semigroup.Get_semigroupString(), "append"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Show.Get_showInt(), "show"), gopurs_runtime.Int((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V0)), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semigroup.Get_semigroupString(), "append"), gopurs_runtime.Str("] "), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semigroup.Get_semigroupString(), "append"), gopurs_runtime.Apply(((*gopurs_runtime.RecordData1)(dictShow_0.UnsafePtr)).V0, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V2), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semigroup.Get_semigroupString(), "append"), gopurs_runtime.Str(" => "), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semigroup.Get_semigroupString(), "append"), gopurs_runtime.Apply(((*gopurs_runtime.RecordData1)(dictShow1_1.UnsafePtr)).V0, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V3), gopurs_runtime.Str("\n")))))))), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semigroup.Get_semigroupString(), "append"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semigroup.Get_semigroupString(), "append"), gopurs_runtime.Apply2(go__2_0, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semigroup.Get_semigroupString(), "append"), ind_3, gopurs_runtime.Str("    ")), (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V4), gopurs_runtime.Str("\n")), gopurs_runtime.Apply2(go__2_0, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semigroup.Get_semigroupString(), "append"), ind_3, gopurs_runtime.Str("    ")), (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V5)))
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
return gopurs_runtime.Apply(go__2_0, gopurs_runtime.Str(""))
}

func Call_semigroupMap(_dollar__unused_0_loop gopurs_runtime.Value, dictOrd_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var _dollar__unused_0 gopurs_runtime.Value = _dollar__unused_0_loop
_ = _dollar__unused_0
var dictOrd_1 gopurs_runtime.Value = dictOrd_1_loop
_ = dictOrd_1
compare_2_0 := ((*gopurs_runtime.RecordData1)(dictOrd_1.UnsafePtr)).V0
_ = compare_2_0
return gopurs_runtime.Func(func(dictSemigroup_3 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_4_1 := gopurs_runtime.RecordGet(dictSemigroup_3, "append")
_ = __local_var_4_1
return gopurs_runtime.RecordDict1("append", gopurs_runtime.Func2(func(m1_5 gopurs_runtime.Value, m2_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.UncurriedApp4(Get_unsafeUnionWith(), compare_2_0, __local_var_4_1, m1_5, m2_6)
}))
})
}

func Call_pop(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
compare_1_0 := ((*gopurs_runtime.RecordData1)(dictOrd_0.UnsafePtr)).V0
_ = compare_1_0
return gopurs_runtime.Func2(func(k_2 gopurs_runtime.Value, m_3 gopurs_runtime.Value) gopurs_runtime.Value {
v_4_1 := gopurs_runtime.UncurriedApp3(Get_unsafeSplit(), compare_1_0, k_2, m_3)
_ = v_4_1
__local_var_5_2 := (*Constructor_Split[gopurs_runtime.Value, gopurs_runtime.Value])(v_4_1.UnsafePtr).V1
_ = __local_var_5_2
__local_var_6_3 := (*Constructor_Split[gopurs_runtime.Value, gopurs_runtime.Value])(v_4_1.UnsafePtr).V2
_ = __local_var_6_3
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_functorMaybe(), "map"), gopurs_runtime.Func(func(a_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{a_7, gopurs_runtime.UncurriedApp2(Get_unsafeJoinNodes(), __local_var_5_2, __local_var_6_3)})}
}), (*Constructor_Split[gopurs_runtime.Value, gopurs_runtime.Value])(v_4_1.UnsafePtr).V0)
})
}

func Call_member(dictOrd_0_loop gopurs_runtime.Value, k_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
var k_1 gopurs_runtime.Value = k_1_loop
_ = k_1
var go__2_0 gopurs_runtime.Value
go__2_0 = gopurs_runtime.Func(func(v_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_3_loop gopurs_runtime.Value = v_3_loop_val
go__2_0:
for {
if false { continue go__2_0 }
var v_3 gopurs_runtime.Value = v_3_loop
_ = v_3
var __t1 gopurs_runtime.Value
{
if (v_3.Type == 9 && v_3.IntVal == 687041424) {
__t1 = gopurs_runtime.Bool(false)
goto end_branch_1
} else {

}
}
{
if (v_3.Type == 9 && v_3.IntVal == 324739070) {
v1_4_2 := gopurs_runtime.Apply2(((*gopurs_runtime.RecordData1)(dictOrd_0.UnsafePtr)).V0, k_1, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V2)
_ = v1_4_2
var __t3 gopurs_runtime.Value
{
if (v1_4_2.Type == 9 && v1_4_2.IntVal == 1527465420) {
v_3_loop = (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V4
continue go__2_0
__t3 = gopurs_runtime.Value{}
goto end_branch_3
} else {

}
}
{
if (v1_4_2.Type == 9 && v1_4_2.IntVal == 380165415) {
v_3_loop = (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V5
continue go__2_0
__t3 = gopurs_runtime.Value{}
goto end_branch_3
} else {

}
}
{
if (v1_4_2.Type == 9 && v1_4_2.IntVal == 902936544) {
__t3 = gopurs_runtime.Bool(true)
goto end_branch_3
} else {

}
}
{
__t3 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_3:
__t1 = __t3
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
return go__2_0
}

func Call_mapMaybeWithKey(dictOrd_0_loop gopurs_runtime.Value, f_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var go__2_0 gopurs_runtime.Value
_ = go__2_0
go__2_0 = gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if (v_3.Type == 9 && v_3.IntVal == 687041424) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 687041424, UnsafePtr: nil}
goto end_branch_1
} else {

}
}
{
if (v_3.Type == 9 && v_3.IntVal == 324739070) {
v2_4_2 := gopurs_runtime.Apply2(f_1, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V2, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V3)
_ = v2_4_2
var __t3 gopurs_runtime.Value
{
if (v2_4_2.Type == 9 && v2_4_2.IntVal == 930809136) {
__t3 = gopurs_runtime.UncurriedApp4(Get_unsafeBalancedNode(), (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V2, (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(v2_4_2.UnsafePtr).V0, gopurs_runtime.Apply(go__2_0, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V4), gopurs_runtime.Apply(go__2_0, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V5))
goto end_branch_3
} else {

}
}
{
if (v2_4_2.Type == 9 && v2_4_2.IntVal == 3589588149) {
__t3 = gopurs_runtime.UncurriedApp2(Get_unsafeJoinNodes(), gopurs_runtime.Apply(go__2_0, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V4), gopurs_runtime.Apply(go__2_0, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V5))
goto end_branch_3
} else {

}
}
{
__t3 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_3:
__t1 = __t3
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
return go__2_0
}

func Call_mapMaybe(dictOrd_0_loop gopurs_runtime.Value, x_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
var x_1 gopurs_runtime.Value = x_1_loop
_ = x_1
return Call_mapMaybeWithKey(dictOrd_0, gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return x_1
}))
}

func Call_lookupLE(dictOrd_0_loop gopurs_runtime.Value, k_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
var k_1 gopurs_runtime.Value = k_1_loop
_ = k_1
var go__2_0 gopurs_runtime.Value
_ = go__2_0
go__2_0 = gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if (v_3.Type == 9 && v_3.IntVal == 687041424) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 3589588149, UnsafePtr: nil}
goto end_branch_1
} else {

}
}
{
if (v_3.Type == 9 && v_3.IntVal == 324739070) {
v1_4_2 := gopurs_runtime.Apply2(((*gopurs_runtime.RecordData1)(dictOrd_0.UnsafePtr)).V0, k_1, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V2)
_ = v1_4_2
var __t3 gopurs_runtime.Value
{
if (v1_4_2.Type == 9 && v1_4_2.IntVal == 1527465420) {
__t3 = gopurs_runtime.Apply(go__2_0, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V4)
goto end_branch_3
} else {

}
}
{
if (v1_4_2.Type == 9 && v1_4_2.IntVal == 380165415) {
v2_5_4 := gopurs_runtime.Apply(go__2_0, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V5)
_ = v2_5_4
var __t5 gopurs_runtime.Value
{
if (v2_5_4.Type == 9 && v2_5_4.IntVal == 3589588149) {
__t5 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{gopurs_runtime.RecordDict2("key", "value", (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V2, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V3)})}
goto end_branch_5
} else {

}
}
{
__t5 = v2_5_4
}
end_branch_5:
__t3 = __t5
goto end_branch_3
} else {

}
}
{
if (v1_4_2.Type == 9 && v1_4_2.IntVal == 902936544) {
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{gopurs_runtime.RecordDict2("key", "value", (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V2, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V3)})}
goto end_branch_3
} else {

}
}
{
__t3 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_3:
__t1 = __t3
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
return go__2_0
}

func Call_lookupGE(dictOrd_0_loop gopurs_runtime.Value, k_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
var k_1 gopurs_runtime.Value = k_1_loop
_ = k_1
var go__2_0 gopurs_runtime.Value
_ = go__2_0
go__2_0 = gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if (v_3.Type == 9 && v_3.IntVal == 687041424) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 3589588149, UnsafePtr: nil}
goto end_branch_1
} else {

}
}
{
if (v_3.Type == 9 && v_3.IntVal == 324739070) {
v1_4_2 := gopurs_runtime.Apply2(((*gopurs_runtime.RecordData1)(dictOrd_0.UnsafePtr)).V0, k_1, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V2)
_ = v1_4_2
var __t3 gopurs_runtime.Value
{
if (v1_4_2.Type == 9 && v1_4_2.IntVal == 1527465420) {
v2_5_4 := gopurs_runtime.Apply(go__2_0, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V4)
_ = v2_5_4
var __t5 gopurs_runtime.Value
{
if (v2_5_4.Type == 9 && v2_5_4.IntVal == 3589588149) {
__t5 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{gopurs_runtime.RecordDict2("key", "value", (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V2, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V3)})}
goto end_branch_5
} else {

}
}
{
__t5 = v2_5_4
}
end_branch_5:
__t3 = __t5
goto end_branch_3
} else {

}
}
{
if (v1_4_2.Type == 9 && v1_4_2.IntVal == 380165415) {
__t3 = gopurs_runtime.Apply(go__2_0, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V5)
goto end_branch_3
} else {

}
}
{
if (v1_4_2.Type == 9 && v1_4_2.IntVal == 902936544) {
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{gopurs_runtime.RecordDict2("key", "value", (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V2, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V3)})}
goto end_branch_3
} else {

}
}
{
__t3 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_3:
__t1 = __t3
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
return go__2_0
}

func Call_lookup(dictOrd_0_loop gopurs_runtime.Value, k_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
var k_1 gopurs_runtime.Value = k_1_loop
_ = k_1
var go__2_0 gopurs_runtime.Value
go__2_0 = gopurs_runtime.Func(func(v_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_3_loop gopurs_runtime.Value = v_3_loop_val
go__2_0:
for {
if false { continue go__2_0 }
var v_3 gopurs_runtime.Value = v_3_loop
_ = v_3
var __t1 gopurs_runtime.Value
{
if (v_3.Type == 9 && v_3.IntVal == 687041424) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 3589588149, UnsafePtr: nil}
goto end_branch_1
} else {

}
}
{
if (v_3.Type == 9 && v_3.IntVal == 324739070) {
v1_4_2 := gopurs_runtime.Apply2(((*gopurs_runtime.RecordData1)(dictOrd_0.UnsafePtr)).V0, k_1, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V2)
_ = v1_4_2
var __t3 gopurs_runtime.Value
{
if (v1_4_2.Type == 9 && v1_4_2.IntVal == 1527465420) {
v_3_loop = (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V4
continue go__2_0
__t3 = gopurs_runtime.Value{}
goto end_branch_3
} else {

}
}
{
if (v1_4_2.Type == 9 && v1_4_2.IntVal == 380165415) {
v_3_loop = (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V5
continue go__2_0
__t3 = gopurs_runtime.Value{}
goto end_branch_3
} else {

}
}
{
if (v1_4_2.Type == 9 && v1_4_2.IntVal == 902936544) {
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{(*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V3})}
goto end_branch_3
} else {

}
}
{
__t3 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_3:
__t1 = __t3
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
return go__2_0
}

func Call_iterMapU(iter_0_loop gopurs_runtime.Value, v_1_loop *Constructor_Node) gopurs_runtime.Value {
var iter_0 gopurs_runtime.Value = iter_0_loop
_ = iter_0
var v_1 *Constructor_Node = v_1_loop
_ = v_1
var __t0 gopurs_runtime.Value
{
if (v_1 == nil) {
__t0 = iter_0
goto end_branch_0
} else {

}
}
{
if (v_1 != nil) {
var __t1 gopurs_runtime.Value
{
var __t_tag_2 gopurs_runtime.Value = (v_1).V4
if (__t_tag_2.Type == 9 && __t_tag_2.IntVal == 687041424) {
var __t3 gopurs_runtime.Value
{
var __t_tag_4 gopurs_runtime.Value = (v_1).V5
if (__t_tag_4.Type == 9 && __t_tag_4.IntVal == 687041424) {
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 1343415489, UnsafePtr: unsafe.Pointer(&Constructor_IterEmit[gopurs_runtime.Value, gopurs_runtime.Value]{(v_1).V2, (v_1).V3, iter_0})}
goto end_branch_3
} else {

}
}
{
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 1343415489, UnsafePtr: unsafe.Pointer(&Constructor_IterEmit[gopurs_runtime.Value, gopurs_runtime.Value]{(v_1).V2, (v_1).V3, gopurs_runtime.Value{Type: 9, IntVal: 2861335956, UnsafePtr: unsafe.Pointer(&Constructor_IterNode[gopurs_runtime.Value, gopurs_runtime.Value]{(v_1).V5, iter_0})}})}
}
end_branch_3:
__t1 = __t3
goto end_branch_1
} else {

}
}
{
var __t_tag_5 gopurs_runtime.Value = (v_1).V5
if (__t_tag_5.Type == 9 && __t_tag_5.IntVal == 687041424) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 1343415489, UnsafePtr: unsafe.Pointer(&Constructor_IterEmit[gopurs_runtime.Value, gopurs_runtime.Value]{(v_1).V2, (v_1).V3, gopurs_runtime.Value{Type: 9, IntVal: 2861335956, UnsafePtr: unsafe.Pointer(&Constructor_IterNode[gopurs_runtime.Value, gopurs_runtime.Value]{(v_1).V4, iter_0})}})}
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 1343415489, UnsafePtr: unsafe.Pointer(&Constructor_IterEmit[gopurs_runtime.Value, gopurs_runtime.Value]{(v_1).V2, (v_1).V3, gopurs_runtime.Value{Type: 9, IntVal: 2861335956, UnsafePtr: unsafe.Pointer(&Constructor_IterNode[gopurs_runtime.Value, gopurs_runtime.Value]{(v_1).V4, gopurs_runtime.Value{Type: 9, IntVal: 2861335956, UnsafePtr: unsafe.Pointer(&Constructor_IterNode[gopurs_runtime.Value, gopurs_runtime.Value]{(v_1).V5, iter_0})}})}})}
}
end_branch_1:
__t0 = __t1
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

func Call_toUnfoldableUnordered(dictUnfoldable_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictUnfoldable_0 gopurs_runtime.Value = dictUnfoldable_0_loop
_ = dictUnfoldable_0
__local_var_1_0 := gopurs_runtime.Apply(((*gopurs_runtime.RecordData1)(dictUnfoldable_0.UnsafePtr)).V0, Get_stepUnfoldrUnordered())
_ = __local_var_1_0
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_1_0, gopurs_runtime.Value{Type: 9, IntVal: 2861335956, UnsafePtr: unsafe.Pointer(&Constructor_IterNode[gopurs_runtime.Value, gopurs_runtime.Value]{x_2, gopurs_runtime.Value{Type: 9, IntVal: 2509360378, UnsafePtr: nil}})})
})
}

func Call_eqMapIter(dictEq_0_loop gopurs_runtime.Value, dictEq1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictEq_0 gopurs_runtime.Value = dictEq_0_loop
_ = dictEq_0
var dictEq1_1 gopurs_runtime.Value = dictEq1_1_loop
_ = dictEq1_1
var go__2_0 gopurs_runtime.Value
_ = go__2_0
go__2_0 = gopurs_runtime.Func2(func(a_3 gopurs_runtime.Value, b_4 gopurs_runtime.Value) gopurs_runtime.Value {
v_5_1 := gopurs_runtime.Apply(Get_stepAsc(), a_3)
_ = v_5_1
var __t2 gopurs_runtime.Value
{
if (v_5_1.Type == 9 && v_5_1.IntVal == 953589075) {
v2_6_3 := gopurs_runtime.Apply(Get_stepAsc(), b_4)
_ = v2_6_3
__t2 = gopurs_runtime.Bool(((v2_6_3.Type == 9 && v2_6_3.IntVal == 953589075)) && (((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_HeytingAlgebra.Get_heytingAlgebraBoolean(), "conj"), gopurs_runtime.Apply2(((*gopurs_runtime.RecordData1)(dictEq_0.UnsafePtr)).V0, (*Constructor_IterNext[gopurs_runtime.Value, gopurs_runtime.Value])(v_5_1.UnsafePtr).V0, (*Constructor_IterNext[gopurs_runtime.Value, gopurs_runtime.Value])(v2_6_3.UnsafePtr).V0), gopurs_runtime.Apply2(((*gopurs_runtime.RecordData1)(dictEq1_1.UnsafePtr)).V0, (*Constructor_IterNext[gopurs_runtime.Value, gopurs_runtime.Value])(v_5_1.UnsafePtr).V1, (*Constructor_IterNext[gopurs_runtime.Value, gopurs_runtime.Value])(v2_6_3.UnsafePtr).V1)).IntVal) != (0)) && ((gopurs_runtime.Apply2(go__2_0, (*Constructor_IterNext[gopurs_runtime.Value, gopurs_runtime.Value])(v_5_1.UnsafePtr).V2, (*Constructor_IterNext[gopurs_runtime.Value, gopurs_runtime.Value])(v2_6_3.UnsafePtr).V2).IntVal) != (0))))
goto end_branch_2
} else {

}
}
{
if (v_5_1.Type == 9 && v_5_1.IntVal == 4236111124) {
__t2 = gopurs_runtime.Bool(true)
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
return gopurs_runtime.RecordDict1("eq", go__2_0)
}

func Call_ordMapIter(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
eqMapIter1_1_0 := gopurs_runtime.Apply(Get_eqMapIter(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd_0, "Eq0_NOT_FOUND"), gopurs_runtime.Value{}))
_ = eqMapIter1_1_0
return gopurs_runtime.Func(func(dictOrd1_2 gopurs_runtime.Value) gopurs_runtime.Value {
eqMapIter2_3_1 := gopurs_runtime.Apply(eqMapIter1_1_0, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd1_2, "Eq0"), gopurs_runtime.Value{}))
_ = eqMapIter2_3_1
var go__4_2 gopurs_runtime.Value
go__4_2 = gopurs_runtime.Func(func(a_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var a_5_loop gopurs_runtime.Value = a_5_loop_val
var b_6_loop gopurs_runtime.Value = b_6_loop_val
go__4_2:
for {
if false { continue go__4_2 }
var a_5 gopurs_runtime.Value = a_5_loop
_ = a_5
var b_6 gopurs_runtime.Value = b_6_loop
_ = b_6
v_7_3 := gopurs_runtime.Apply(Get_stepAsc(), b_6)
_ = v_7_3
v1_8_4 := gopurs_runtime.Apply(Get_stepAsc(), a_5)
_ = v1_8_4
var __t5 gopurs_runtime.Value
{
if (v1_8_4.Type == 9 && v1_8_4.IntVal == 953589075) {
var __t6 gopurs_runtime.Value
{
if (v_7_3.Type == 9 && v_7_3.IntVal == 953589075) {
v3_9_7 := gopurs_runtime.Apply2(((*gopurs_runtime.RecordData1)(dictOrd_0.UnsafePtr)).V0, (*Constructor_IterNext[gopurs_runtime.Value, gopurs_runtime.Value])(v1_8_4.UnsafePtr).V0, (*Constructor_IterNext[gopurs_runtime.Value, gopurs_runtime.Value])(v_7_3.UnsafePtr).V0)
_ = v3_9_7
var __t8 gopurs_runtime.Value
{
if (v3_9_7.Type == 9 && v3_9_7.IntVal == 902936544) {
v4_10_9 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictOrd1_2, "compare"), (*Constructor_IterNext[gopurs_runtime.Value, gopurs_runtime.Value])(v1_8_4.UnsafePtr).V1, (*Constructor_IterNext[gopurs_runtime.Value, gopurs_runtime.Value])(v_7_3.UnsafePtr).V1)
_ = v4_10_9
var __t10 gopurs_runtime.Value
{
if (v4_10_9.Type == 9 && v4_10_9.IntVal == 902936544) {
a_5_loop = (*Constructor_IterNext[gopurs_runtime.Value, gopurs_runtime.Value])(v1_8_4.UnsafePtr).V2
b_6_loop = (*Constructor_IterNext[gopurs_runtime.Value, gopurs_runtime.Value])(v_7_3.UnsafePtr).V2
continue go__4_2
__t10 = gopurs_runtime.Value{}
goto end_branch_10
} else {

}
}
{
__t10 = v4_10_9
}
end_branch_10:
__t8 = __t10
goto end_branch_8
} else {

}
}
{
__t8 = v3_9_7
}
end_branch_8:
__t6 = __t8
goto end_branch_6
} else {

}
}
{
if (v_7_3.Type == 9 && v_7_3.IntVal == 4236111124) {
__t6 = gopurs_runtime.Value{Type: 9, IntVal: 380165415, UnsafePtr: nil}
goto end_branch_6
} else {

}
}
{
__t6 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_6:
__t5 = __t6
goto end_branch_5
} else {

}
}
{
if (v1_8_4.Type == 9 && v1_8_4.IntVal == 4236111124) {
var __t11 gopurs_runtime.Value
{
if (v_7_3.Type == 9 && v_7_3.IntVal == 4236111124) {
__t11 = gopurs_runtime.Value{Type: 9, IntVal: 902936544, UnsafePtr: nil}
goto end_branch_11
} else {

}
}
{
__t11 = gopurs_runtime.Value{Type: 9, IntVal: 1527465420, UnsafePtr: nil}
}
end_branch_11:
__t5 = __t11
goto end_branch_5
} else {

}
}
{
if (v_7_3.Type == 9 && v_7_3.IntVal == 4236111124) {
__t5 = gopurs_runtime.Value{Type: 9, IntVal: 380165415, UnsafePtr: nil}
goto end_branch_5
} else {

}
}
{
__t5 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_5:
return __t5
}
}()
})
})
return gopurs_runtime.RecordDict2("Eq0", "compare", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return eqMapIter2_3_1
}), go__4_2)
})
}

func Call_toUnfoldable(dictUnfoldable_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictUnfoldable_0 gopurs_runtime.Value = dictUnfoldable_0_loop
_ = dictUnfoldable_0
__local_var_1_0 := gopurs_runtime.Apply(((*gopurs_runtime.RecordData1)(dictUnfoldable_0.UnsafePtr)).V0, Get_stepUnfoldr())
_ = __local_var_1_0
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_1_0, gopurs_runtime.Value{Type: 9, IntVal: 2861335956, UnsafePtr: unsafe.Pointer(&Constructor_IterNode[gopurs_runtime.Value, gopurs_runtime.Value]{x_2, gopurs_runtime.Value{Type: 9, IntVal: 2509360378, UnsafePtr: nil}})})
})
}

func Call_showMap(dictShow_0_loop gopurs_runtime.Value, dictShow1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictShow_0 gopurs_runtime.Value = dictShow_0_loop
_ = dictShow_0
var dictShow1_1 gopurs_runtime.Value = dictShow1_1_loop
_ = dictShow1_1
show1_2_0 := gopurs_runtime.Apply(pkg_Data_Show.Get_showArrayImpl(), gopurs_runtime.RecordGet(gopurs_runtime.Apply2(pkg_Data_Tuple.Get_showTuple(), dictShow_0, dictShow1_1), "show"))
_ = show1_2_0
return gopurs_runtime.RecordDict1("show", gopurs_runtime.Func(func(as_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semigroup.Get_semigroupString(), "append"), gopurs_runtime.Str("(fromFoldable "), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semigroup.Get_semigroupString(), "append"), gopurs_runtime.Apply(show1_2_0, gopurs_runtime.Apply(Get_toUnfoldable1(), as_3)), gopurs_runtime.Str(")")))
}))
}

func Call_isSubmap(dictOrd_0_loop gopurs_runtime.Value, dictEq_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
var dictEq_1 gopurs_runtime.Value = dictEq_1_loop
_ = dictEq_1
var go__2_0 gopurs_runtime.Value
_ = go__2_0
go__2_0 = gopurs_runtime.Func2(func(m1_3 gopurs_runtime.Value, m2_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if (m1_3.Type == 9 && m1_3.IntVal == 687041424) {
__t1 = gopurs_runtime.Bool(true)
goto end_branch_1
} else {

}
}
{
if (m1_3.Type == 9 && m1_3.IntVal == 324739070) {
__local_var_5_2 := (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(m1_3.UnsafePtr).V2
_ = __local_var_5_2
var go__6_3 gopurs_runtime.Value
go__6_3 = gopurs_runtime.Func(func(v_7_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_7_loop gopurs_runtime.Value = v_7_loop_val
go__6_3:
for {
if false { continue go__6_3 }
var v_7 gopurs_runtime.Value = v_7_loop
_ = v_7
var __t4 gopurs_runtime.Value
{
if (v_7.Type == 9 && v_7.IntVal == 687041424) {
__t4 = gopurs_runtime.Value{Type: 9, IntVal: 3589588149, UnsafePtr: nil}
goto end_branch_4
} else {

}
}
{
if (v_7.Type == 9 && v_7.IntVal == 324739070) {
v1_8_5 := gopurs_runtime.Apply2(((*gopurs_runtime.RecordData1)(dictOrd_0.UnsafePtr)).V0, __local_var_5_2, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_7.UnsafePtr).V2)
_ = v1_8_5
var __t6 gopurs_runtime.Value
{
if (v1_8_5.Type == 9 && v1_8_5.IntVal == 1527465420) {
v_7_loop = (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_7.UnsafePtr).V4
continue go__6_3
__t6 = gopurs_runtime.Value{}
goto end_branch_6
} else {

}
}
{
if (v1_8_5.Type == 9 && v1_8_5.IntVal == 380165415) {
v_7_loop = (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_7.UnsafePtr).V5
continue go__6_3
__t6 = gopurs_runtime.Value{}
goto end_branch_6
} else {

}
}
{
if (v1_8_5.Type == 9 && v1_8_5.IntVal == 902936544) {
__t6 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{(*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_7.UnsafePtr).V3})}
goto end_branch_6
} else {

}
}
{
__t6 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_6:
__t4 = __t6
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
v1_7_7 := gopurs_runtime.Apply(go__6_3, m2_4)
_ = v1_7_7
var __t8 gopurs_runtime.Value
{
if (v1_7_7.Type == 9 && v1_7_7.IntVal == 3589588149) {
__t8 = gopurs_runtime.Bool(false)
goto end_branch_8
} else {

}
}
{
if (v1_7_7.Type == 9 && v1_7_7.IntVal == 930809136) {
__t8 = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_HeytingAlgebra.Get_heytingAlgebraBoolean(), "conj"), gopurs_runtime.Apply2(((*gopurs_runtime.RecordData1)(dictEq_1.UnsafePtr)).V0, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(m1_3.UnsafePtr).V3, (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(v1_7_7.UnsafePtr).V0), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_HeytingAlgebra.Get_heytingAlgebraBoolean(), "conj"), gopurs_runtime.Apply2(go__2_0, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(m1_3.UnsafePtr).V4, m2_4), gopurs_runtime.Apply2(go__2_0, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(m1_3.UnsafePtr).V5, m2_4)))
goto end_branch_8
} else {

}
}
{
__t8 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_8:
__t1 = __t8
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
return go__2_0
}

func Call_isEmpty(v_0_loop *Constructor_Node) bool {
var v_0 *Constructor_Node = v_0_loop
_ = v_0
return (v_0 == nil)
}

func Call_intersectionWith(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
compare_1_0 := ((*gopurs_runtime.RecordData1)(dictOrd_0.UnsafePtr)).V0
_ = compare_1_0
return gopurs_runtime.Func3(func(app_2 gopurs_runtime.Value, m1_3 gopurs_runtime.Value, m2_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.UncurriedApp4(Get_unsafeIntersectionWith(), compare_1_0, app_2, m1_3, m2_4)
})
}

func Call_intersection(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
compare_1_0 := ((*gopurs_runtime.RecordData1)(dictOrd_0.UnsafePtr)).V0
_ = compare_1_0
return gopurs_runtime.Func2(func(m1_2 gopurs_runtime.Value, m2_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.UncurriedApp4(Get_unsafeIntersectionWith(), compare_1_0, pkg_Data_Function.Get_const_(), m1_2, m2_3)
})
}

func Call_insertWith(dictOrd_0_loop gopurs_runtime.Value, app_1_loop gopurs_runtime.Value, k_2_loop gopurs_runtime.Value, v_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
var app_1 gopurs_runtime.Value = app_1_loop
_ = app_1
var k_2 gopurs_runtime.Value = k_2_loop
_ = k_2
var v_3 gopurs_runtime.Value = v_3_loop
_ = v_3
var go__4_0 gopurs_runtime.Value
_ = go__4_0
go__4_0 = gopurs_runtime.Func(func(v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if (v1_5.Type == 9 && v1_5.IntVal == 687041424) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(&Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]{1, 1, k_2, v_3, gopurs_runtime.Value{Type: 9, IntVal: 687041424, UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: 687041424, UnsafePtr: nil}})}
goto end_branch_1
} else {

}
}
{
if (v1_5.Type == 9 && v1_5.IntVal == 324739070) {
v2_6_2 := gopurs_runtime.Apply2(((*gopurs_runtime.RecordData1)(dictOrd_0.UnsafePtr)).V0, k_2, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v1_5.UnsafePtr).V2)
_ = v2_6_2
var __t3 gopurs_runtime.Value
{
if (v2_6_2.Type == 9 && v2_6_2.IntVal == 1527465420) {
__t3 = gopurs_runtime.UncurriedApp4(Get_unsafeBalancedNode(), (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v1_5.UnsafePtr).V2, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v1_5.UnsafePtr).V3, gopurs_runtime.Apply(go__4_0, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v1_5.UnsafePtr).V4), (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v1_5.UnsafePtr).V5)
goto end_branch_3
} else {

}
}
{
if (v2_6_2.Type == 9 && v2_6_2.IntVal == 380165415) {
__t3 = gopurs_runtime.UncurriedApp4(Get_unsafeBalancedNode(), (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v1_5.UnsafePtr).V2, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v1_5.UnsafePtr).V3, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v1_5.UnsafePtr).V4, gopurs_runtime.Apply(go__4_0, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v1_5.UnsafePtr).V5))
goto end_branch_3
} else {

}
}
{
if (v2_6_2.Type == 9 && v2_6_2.IntVal == 902936544) {
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(&Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]{gopurs_runtime.Int((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v1_5.UnsafePtr).V0).IntVal, gopurs_runtime.Int((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v1_5.UnsafePtr).V1).IntVal, k_2, gopurs_runtime.Apply2(app_1, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v1_5.UnsafePtr).V3, v_3), (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v1_5.UnsafePtr).V4, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v1_5.UnsafePtr).V5})}
goto end_branch_3
} else {

}
}
{
__t3 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_3:
__t1 = __t3
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
return go__4_0
}

func Call_insert(dictOrd_0_loop gopurs_runtime.Value, k_1_loop gopurs_runtime.Value, v_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
var k_1 gopurs_runtime.Value = k_1_loop
_ = k_1
var v_2 gopurs_runtime.Value = v_2_loop
_ = v_2
var go__3_0 gopurs_runtime.Value
_ = go__3_0
go__3_0 = gopurs_runtime.Func(func(v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if (v1_4.Type == 9 && v1_4.IntVal == 687041424) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(&Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]{1, 1, k_1, v_2, gopurs_runtime.Value{Type: 9, IntVal: 687041424, UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: 687041424, UnsafePtr: nil}})}
goto end_branch_1
} else {

}
}
{
if (v1_4.Type == 9 && v1_4.IntVal == 324739070) {
v2_5_2 := gopurs_runtime.Apply2(((*gopurs_runtime.RecordData1)(dictOrd_0.UnsafePtr)).V0, k_1, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v1_4.UnsafePtr).V2)
_ = v2_5_2
var __t3 gopurs_runtime.Value
{
if (v2_5_2.Type == 9 && v2_5_2.IntVal == 1527465420) {
__t3 = gopurs_runtime.UncurriedApp4(Get_unsafeBalancedNode(), (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v1_4.UnsafePtr).V2, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v1_4.UnsafePtr).V3, gopurs_runtime.Apply(go__3_0, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v1_4.UnsafePtr).V4), (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v1_4.UnsafePtr).V5)
goto end_branch_3
} else {

}
}
{
if (v2_5_2.Type == 9 && v2_5_2.IntVal == 380165415) {
__t3 = gopurs_runtime.UncurriedApp4(Get_unsafeBalancedNode(), (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v1_4.UnsafePtr).V2, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v1_4.UnsafePtr).V3, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v1_4.UnsafePtr).V4, gopurs_runtime.Apply(go__3_0, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v1_4.UnsafePtr).V5))
goto end_branch_3
} else {

}
}
{
if (v2_5_2.Type == 9 && v2_5_2.IntVal == 902936544) {
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(&Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]{gopurs_runtime.Int((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v1_4.UnsafePtr).V0).IntVal, gopurs_runtime.Int((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v1_4.UnsafePtr).V1).IntVal, k_1, v_2, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v1_4.UnsafePtr).V4, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v1_4.UnsafePtr).V5})}
goto end_branch_3
} else {

}
}
{
__t3 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_3:
__t1 = __t3
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
return go__3_0
}

func Call_go__2_4(z_prime_3 gopurs_runtime.Value, m_prime_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t5 gopurs_runtime.Value
{
if (m_prime_4.Type == 9 && m_prime_4.IntVal == 687041424) {
__t5 = z_prime_3
goto end_branch_5
} else {

}
}
{
if (m_prime_4.Type == 9 && m_prime_4.IntVal == 324739070) {
__t5 = gopurs_runtime.UncurriedApp2(go__2_4, gopurs_runtime.Apply2(f_0, gopurs_runtime.UncurriedApp2(go__2_4, z_prime_3, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(m_prime_4.UnsafePtr).V4), (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(m_prime_4.UnsafePtr).V3), (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(m_prime_4.UnsafePtr).V5)
goto end_branch_5
} else {

}
}
{
__t5 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_5:
return __t5
}

func Call_go__2_6(m_prime_3 gopurs_runtime.Value, z_prime_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t7 gopurs_runtime.Value
{
if (m_prime_3.Type == 9 && m_prime_3.IntVal == 687041424) {
__t7 = z_prime_4
goto end_branch_7
} else {

}
}
{
if (m_prime_3.Type == 9 && m_prime_3.IntVal == 324739070) {
__t7 = gopurs_runtime.UncurriedApp2(go__2_6, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(m_prime_3.UnsafePtr).V4, gopurs_runtime.Apply2(f_0, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(m_prime_3.UnsafePtr).V3, gopurs_runtime.UncurriedApp2(go__2_6, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(m_prime_3.UnsafePtr).V5, z_prime_4)))
goto end_branch_7
} else {

}
}
{
__t7 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_7:
return __t7
}

func Call_go__2_4(z_prime_3 gopurs_runtime.Value, m_prime_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t5 gopurs_runtime.Value
{
if (m_prime_4.Type == 9 && m_prime_4.IntVal == 687041424) {
__t5 = z_prime_3
goto end_branch_5
} else {

}
}
{
if (m_prime_4.Type == 9 && m_prime_4.IntVal == 324739070) {
__t5 = gopurs_runtime.UncurriedApp2(go__2_4, gopurs_runtime.Apply3(f_0, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(m_prime_4.UnsafePtr).V2, gopurs_runtime.UncurriedApp2(go__2_4, z_prime_3, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(m_prime_4.UnsafePtr).V4), (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(m_prime_4.UnsafePtr).V3), (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(m_prime_4.UnsafePtr).V5)
goto end_branch_5
} else {

}
}
{
__t5 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_5:
return __t5
}

func Call_go__2_6(m_prime_3 gopurs_runtime.Value, z_prime_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t7 gopurs_runtime.Value
{
if (m_prime_3.Type == 9 && m_prime_3.IntVal == 687041424) {
__t7 = z_prime_4
goto end_branch_7
} else {

}
}
{
if (m_prime_3.Type == 9 && m_prime_3.IntVal == 324739070) {
__t7 = gopurs_runtime.UncurriedApp2(go__2_6, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(m_prime_3.UnsafePtr).V4, gopurs_runtime.Apply3(f_0, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(m_prime_3.UnsafePtr).V2, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(m_prime_3.UnsafePtr).V3, gopurs_runtime.UncurriedApp2(go__2_6, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(m_prime_3.UnsafePtr).V5, z_prime_4)))
goto end_branch_7
} else {

}
}
{
__t7 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_7:
return __t7
}

func Call_foldSubmapBy(dictOrd_0_loop gopurs_runtime.Value, appendFn_1_loop gopurs_runtime.Value, memptyValue_2_loop gopurs_runtime.Value, kmin_3_loop *pkg_Data_Maybe.Constructor_Just, kmax_4_loop *pkg_Data_Maybe.Constructor_Just, f_5_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
var appendFn_1 gopurs_runtime.Value = appendFn_1_loop
_ = appendFn_1
var memptyValue_2 gopurs_runtime.Value = memptyValue_2_loop
_ = memptyValue_2
var kmin_3 *pkg_Data_Maybe.Constructor_Just = kmin_3_loop
_ = kmin_3
var kmax_4 *pkg_Data_Maybe.Constructor_Just = kmax_4_loop
_ = kmax_4
var f_5 gopurs_runtime.Value = f_5_loop
_ = f_5
var __t1 gopurs_runtime.Value
{
if (kmin_3 != nil) {
__local_var_6_2 := (kmin_3).V0
_ = __local_var_6_2
__t1 = gopurs_runtime.Func(func(k_7 gopurs_runtime.Value) gopurs_runtime.Value {
var __t_tag_3 gopurs_runtime.Value = gopurs_runtime.Apply2(((*gopurs_runtime.RecordData1)(dictOrd_0.UnsafePtr)).V0, k_7, __local_var_6_2)
return gopurs_runtime.Bool((__t_tag_3.Type == 9 && __t_tag_3.IntVal == 1527465420))
})
goto end_branch_1
} else {

}
}
{
if (kmin_3 == nil) {
__t1 = gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(false)
})
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
tooSmall_6_0 := __t1
_ = tooSmall_6_0
var __t5 gopurs_runtime.Value
{
if (kmax_4 != nil) {
__local_var_7_6 := (kmax_4).V0
_ = __local_var_7_6
__t5 = gopurs_runtime.Func(func(k_8 gopurs_runtime.Value) gopurs_runtime.Value {
var __t_tag_7 gopurs_runtime.Value = gopurs_runtime.Apply2(((*gopurs_runtime.RecordData1)(dictOrd_0.UnsafePtr)).V0, k_8, __local_var_7_6)
return gopurs_runtime.Bool((__t_tag_7.Type == 9 && __t_tag_7.IntVal == 380165415))
})
goto end_branch_5
} else {

}
}
{
if (kmax_4 == nil) {
__t5 = gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(false)
})
goto end_branch_5
} else {

}
}
{
__t5 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_5:
tooLarge_7_4 := __t5
_ = tooLarge_7_4
var __t9 gopurs_runtime.Value
{
if (kmin_3 != nil) {
var __t10 gopurs_runtime.Value
{
if (kmax_4 != nil) {
__local_var_8_11 := (kmax_4).V0
_ = __local_var_8_11
__local_var_9_12 := (kmin_3).V0
_ = __local_var_9_12
__t10 = gopurs_runtime.Func(func(k_10 gopurs_runtime.Value) gopurs_runtime.Value {
var __t_tag_13 gopurs_runtime.Value = gopurs_runtime.Apply2(((*gopurs_runtime.RecordData1)(dictOrd_0.UnsafePtr)).V0, __local_var_9_12, k_10)
var __t_tag_14 gopurs_runtime.Value = gopurs_runtime.Apply2(((*gopurs_runtime.RecordData1)(dictOrd_0.UnsafePtr)).V0, k_10, __local_var_8_11)
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_HeytingAlgebra.Get_heytingAlgebraBoolean(), "conj"), gopurs_runtime.Bool(((__t_tag_13.Type == 9 && __t_tag_13.IntVal == 380165415)) != (true)), gopurs_runtime.Bool(((__t_tag_14.Type == 9 && __t_tag_14.IntVal == 380165415)) != (true)))
})
goto end_branch_10
} else {

}
}
{
if (kmax_4 == nil) {
__local_var_8_15 := (kmin_3).V0
_ = __local_var_8_15
__t10 = gopurs_runtime.Func(func(k_9 gopurs_runtime.Value) gopurs_runtime.Value {
var __t_tag_16 gopurs_runtime.Value = gopurs_runtime.Apply2(((*gopurs_runtime.RecordData1)(dictOrd_0.UnsafePtr)).V0, __local_var_8_15, k_9)
return gopurs_runtime.Bool(((__t_tag_16.Type == 9 && __t_tag_16.IntVal == 380165415)) != (true))
})
goto end_branch_10
} else {

}
}
{
__t10 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_10:
__t9 = __t10
goto end_branch_9
} else {

}
}
{
if (kmin_3 == nil) {
var __t17 gopurs_runtime.Value
{
if (kmax_4 != nil) {
__local_var_8_18 := (kmax_4).V0
_ = __local_var_8_18
__t17 = gopurs_runtime.Func(func(k_9 gopurs_runtime.Value) gopurs_runtime.Value {
var __t_tag_19 gopurs_runtime.Value = gopurs_runtime.Apply2(((*gopurs_runtime.RecordData1)(dictOrd_0.UnsafePtr)).V0, k_9, __local_var_8_18)
return gopurs_runtime.Bool(((__t_tag_19.Type == 9 && __t_tag_19.IntVal == 380165415)) != (true))
})
goto end_branch_17
} else {

}
}
{
if (kmax_4 == nil) {
__t17 = gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(true)
})
goto end_branch_17
} else {

}
}
{
__t17 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_17:
__t9 = __t17
goto end_branch_9
} else {

}
}
{
__t9 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_9:
inBounds_8_8 := __t9
_ = inBounds_8_8
var go__9_20 gopurs_runtime.Value
_ = go__9_20
go__9_20 = gopurs_runtime.Func(func(v_10 gopurs_runtime.Value) gopurs_runtime.Value {
var __t21 gopurs_runtime.Value
{
if (v_10.Type == 9 && v_10.IntVal == 687041424) {
__t21 = memptyValue_2
goto end_branch_21
} else {

}
}
{
if (v_10.Type == 9 && v_10.IntVal == 324739070) {
var __t22 gopurs_runtime.Value
{
if (gopurs_runtime.Apply(tooSmall_6_0, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_10.UnsafePtr).V2).IntVal) != (0) {
__t22 = memptyValue_2
goto end_branch_22
} else {

}
}
{
__t22 = gopurs_runtime.Apply(go__9_20, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_10.UnsafePtr).V4)
}
end_branch_22:
var __t23 gopurs_runtime.Value
{
if (gopurs_runtime.Apply(inBounds_8_8, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_10.UnsafePtr).V2).IntVal) != (0) {
__t23 = gopurs_runtime.Apply2(f_5, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_10.UnsafePtr).V2, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_10.UnsafePtr).V3)
goto end_branch_23
} else {

}
}
{
__t23 = memptyValue_2
}
end_branch_23:
var __t24 gopurs_runtime.Value
{
if (gopurs_runtime.Apply(tooLarge_7_4, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_10.UnsafePtr).V2).IntVal) != (0) {
__t24 = memptyValue_2
goto end_branch_24
} else {

}
}
{
__t24 = gopurs_runtime.Apply(go__9_20, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_10.UnsafePtr).V5)
}
end_branch_24:
__t21 = gopurs_runtime.Apply2(appendFn_1, gopurs_runtime.Apply2(appendFn_1, __t22, __t23), __t24)
goto end_branch_21
} else {

}
}
{
__t21 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_21:
return __t21
})
return go__9_20
}

func Call_foldSubmap(dictOrd_0_loop gopurs_runtime.Value, dictMonoid_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
var dictMonoid_1 gopurs_runtime.Value = dictMonoid_1_loop
_ = dictMonoid_1
return gopurs_runtime.Apply3(Get_foldSubmapBy(), dictOrd_0, gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_1, "Semigroup0_NOT_FOUND"), gopurs_runtime.Value{}), "append"), ((*gopurs_runtime.RecordData1)(dictMonoid_1.UnsafePtr)).V0)
}

func Call_findMin(v_0_loop *Constructor_Node) gopurs_runtime.Value {
findMin:
for {
if false { continue findMin }
var v_0 *Constructor_Node = v_0_loop
_ = v_0
var __t0 gopurs_runtime.Value
{
if (v_0 == nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 3589588149, UnsafePtr: nil}
goto end_branch_0
} else {

}
}
{
if (v_0 != nil) {
var __t1 gopurs_runtime.Value
{
var __t_tag_2 gopurs_runtime.Value = (v_0).V4
if (__t_tag_2.Type == 9 && __t_tag_2.IntVal == 687041424) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{gopurs_runtime.RecordDict2("key", "value", (v_0).V2, (v_0).V3)})}
goto end_branch_1
} else {

}
}
{
v_0_loop = (v_0).V4
continue findMin
__t1 = gopurs_runtime.Value{}
}
end_branch_1:
__t0 = __t1
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
}

func Call_lookupGT(dictOrd_0_loop gopurs_runtime.Value, k_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
var k_1 gopurs_runtime.Value = k_1_loop
_ = k_1
var go__2_0 gopurs_runtime.Value
_ = go__2_0
go__2_0 = gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if (v_3.Type == 9 && v_3.IntVal == 687041424) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 3589588149, UnsafePtr: nil}
goto end_branch_1
} else {

}
}
{
if (v_3.Type == 9 && v_3.IntVal == 324739070) {
v1_4_2 := gopurs_runtime.Apply2(((*gopurs_runtime.RecordData1)(dictOrd_0.UnsafePtr)).V0, k_1, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V2)
_ = v1_4_2
var __t3 gopurs_runtime.Value
{
if (v1_4_2.Type == 9 && v1_4_2.IntVal == 1527465420) {
v2_5_4 := gopurs_runtime.Apply(go__2_0, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V4)
_ = v2_5_4
var __t5 gopurs_runtime.Value
{
if (v2_5_4.Type == 9 && v2_5_4.IntVal == 3589588149) {
__t5 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{gopurs_runtime.RecordDict2("key", "value", (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V2, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V3)})}
goto end_branch_5
} else {

}
}
{
__t5 = v2_5_4
}
end_branch_5:
__t3 = __t5
goto end_branch_3
} else {

}
}
{
if (v1_4_2.Type == 9 && v1_4_2.IntVal == 380165415) {
__t3 = gopurs_runtime.Apply(go__2_0, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V5)
goto end_branch_3
} else {

}
}
{
if (v1_4_2.Type == 9 && v1_4_2.IntVal == 902936544) {
__t3 = gopurs_runtime.Apply(Get_findMin(), (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V5)
goto end_branch_3
} else {

}
}
{
__t3 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_3:
__t1 = __t3
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
return go__2_0
}

func Call_findMax(v_0_loop *Constructor_Node) gopurs_runtime.Value {
findMax:
for {
if false { continue findMax }
var v_0 *Constructor_Node = v_0_loop
_ = v_0
var __t0 gopurs_runtime.Value
{
if (v_0 == nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 3589588149, UnsafePtr: nil}
goto end_branch_0
} else {

}
}
{
if (v_0 != nil) {
var __t1 gopurs_runtime.Value
{
var __t_tag_2 gopurs_runtime.Value = (v_0).V5
if (__t_tag_2.Type == 9 && __t_tag_2.IntVal == 687041424) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{gopurs_runtime.RecordDict2("key", "value", (v_0).V2, (v_0).V3)})}
goto end_branch_1
} else {

}
}
{
v_0_loop = (v_0).V5
continue findMax
__t1 = gopurs_runtime.Value{}
}
end_branch_1:
__t0 = __t1
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
}

func Call_lookupLT(dictOrd_0_loop gopurs_runtime.Value, k_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
var k_1 gopurs_runtime.Value = k_1_loop
_ = k_1
var go__2_0 gopurs_runtime.Value
_ = go__2_0
go__2_0 = gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if (v_3.Type == 9 && v_3.IntVal == 687041424) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 3589588149, UnsafePtr: nil}
goto end_branch_1
} else {

}
}
{
if (v_3.Type == 9 && v_3.IntVal == 324739070) {
v1_4_2 := gopurs_runtime.Apply2(((*gopurs_runtime.RecordData1)(dictOrd_0.UnsafePtr)).V0, k_1, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V2)
_ = v1_4_2
var __t3 gopurs_runtime.Value
{
if (v1_4_2.Type == 9 && v1_4_2.IntVal == 1527465420) {
__t3 = gopurs_runtime.Apply(go__2_0, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V4)
goto end_branch_3
} else {

}
}
{
if (v1_4_2.Type == 9 && v1_4_2.IntVal == 380165415) {
v2_5_4 := gopurs_runtime.Apply(go__2_0, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V5)
_ = v2_5_4
var __t5 gopurs_runtime.Value
{
if (v2_5_4.Type == 9 && v2_5_4.IntVal == 3589588149) {
__t5 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{gopurs_runtime.RecordDict2("key", "value", (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V2, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V3)})}
goto end_branch_5
} else {

}
}
{
__t5 = v2_5_4
}
end_branch_5:
__t3 = __t5
goto end_branch_3
} else {

}
}
{
if (v1_4_2.Type == 9 && v1_4_2.IntVal == 902936544) {
__t3 = gopurs_runtime.Apply(Get_findMax(), (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V4)
goto end_branch_3
} else {

}
}
{
__t3 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_3:
__t1 = __t3
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
return go__2_0
}

func Call_filterWithKey(dictOrd_0_loop gopurs_runtime.Value, f_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var go__2_0 gopurs_runtime.Value
_ = go__2_0
go__2_0 = gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if (v_3.Type == 9 && v_3.IntVal == 687041424) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 687041424, UnsafePtr: nil}
goto end_branch_1
} else {

}
}
{
if (v_3.Type == 9 && v_3.IntVal == 324739070) {
var __t2 gopurs_runtime.Value
{
if (gopurs_runtime.Apply2(f_1, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V2, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V3).IntVal) != (0) {
__t2 = gopurs_runtime.UncurriedApp4(Get_unsafeBalancedNode(), (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V2, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V3, gopurs_runtime.Apply(go__2_0, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V4), gopurs_runtime.Apply(go__2_0, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V5))
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.UncurriedApp2(Get_unsafeJoinNodes(), gopurs_runtime.Apply(go__2_0, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V4), gopurs_runtime.Apply(go__2_0, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V5))
}
end_branch_2:
__t1 = __t2
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
return go__2_0
}

func Call_filterKeys(dictOrd_0_loop gopurs_runtime.Value, f_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var go__2_0 gopurs_runtime.Value
_ = go__2_0
go__2_0 = gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if (v_3.Type == 9 && v_3.IntVal == 687041424) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 687041424, UnsafePtr: nil}
goto end_branch_1
} else {

}
}
{
if (v_3.Type == 9 && v_3.IntVal == 324739070) {
var __t2 gopurs_runtime.Value
{
if (gopurs_runtime.Apply(f_1, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V2).IntVal) != (0) {
__t2 = gopurs_runtime.UncurriedApp4(Get_unsafeBalancedNode(), (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V2, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V3, gopurs_runtime.Apply(go__2_0, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V4), gopurs_runtime.Apply(go__2_0, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V5))
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.UncurriedApp2(Get_unsafeJoinNodes(), gopurs_runtime.Apply(go__2_0, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V4), gopurs_runtime.Apply(go__2_0, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V5))
}
end_branch_2:
__t1 = __t2
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
return go__2_0
}

func Call_filter(dictOrd_0_loop gopurs_runtime.Value, x_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
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
return gopurs_runtime.RecordDict1("eq", gopurs_runtime.Func2(func(xs_2 gopurs_runtime.Value, ys_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (xs_2.Type == 9 && xs_2.IntVal == 687041424) {
__t0 = gopurs_runtime.Bool((ys_3.Type == 9 && ys_3.IntVal == 687041424))
goto end_branch_0
} else {

}
}
{
if (xs_2.Type == 9 && xs_2.IntVal == 324739070) {
__t0 = gopurs_runtime.Bool(((ys_3.Type == 9 && ys_3.IntVal == 324739070)) && (((gopurs_runtime.Int((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(xs_2.UnsafePtr).V1).IntVal) == (gopurs_runtime.Int((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(ys_3.UnsafePtr).V1).IntVal)) && ((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Call_eqMapIter(dictEq_0, dictEq1_1), "eq"), gopurs_runtime.Value{Type: 9, IntVal: 2861335956, UnsafePtr: unsafe.Pointer(&Constructor_IterNode[gopurs_runtime.Value, gopurs_runtime.Value]{xs_2, gopurs_runtime.Value{Type: 9, IntVal: 2509360378, UnsafePtr: nil}})}, gopurs_runtime.Value{Type: 9, IntVal: 2861335956, UnsafePtr: unsafe.Pointer(&Constructor_IterNode[gopurs_runtime.Value, gopurs_runtime.Value]{ys_3, gopurs_runtime.Value{Type: 9, IntVal: 2509360378, UnsafePtr: nil}})}).IntVal) != (0))))
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return __t0
}))
}

func Call_ordMap(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
ordMapIter1_1_0 := gopurs_runtime.Apply(Get_ordMapIter(), dictOrd_0)
_ = ordMapIter1_1_0
eqMap1_2_1 := gopurs_runtime.Apply(Get_eqMap(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd_0, "Eq0_NOT_FOUND"), gopurs_runtime.Value{}))
_ = eqMap1_2_1
return gopurs_runtime.Func(func(dictOrd1_3 gopurs_runtime.Value) gopurs_runtime.Value {
eqMap2_4_2 := gopurs_runtime.Apply(eqMap1_2_1, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd1_3, "Eq0"), gopurs_runtime.Value{}))
_ = eqMap2_4_2
return gopurs_runtime.RecordDict2("Eq0", "compare", gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return eqMap2_4_2
}), gopurs_runtime.Func2(func(xs_5 gopurs_runtime.Value, ys_6 gopurs_runtime.Value) gopurs_runtime.Value {
var __t3 gopurs_runtime.Value
{
if (xs_5.Type == 9 && xs_5.IntVal == 687041424) {
var __t4 gopurs_runtime.Value
{
if (ys_6.Type == 9 && ys_6.IntVal == 687041424) {
__t4 = gopurs_runtime.Value{Type: 9, IntVal: 902936544, UnsafePtr: nil}
goto end_branch_4
} else {

}
}
{
__t4 = gopurs_runtime.Value{Type: 9, IntVal: 1527465420, UnsafePtr: nil}
}
end_branch_4:
__t3 = __t4
goto end_branch_3
} else {

}
}
{
if (ys_6.Type == 9 && ys_6.IntVal == 687041424) {
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 380165415, UnsafePtr: nil}
goto end_branch_3
} else {

}
}
{
__t3 = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(ordMapIter1_1_0, dictOrd1_3), "compare"), gopurs_runtime.Value{Type: 9, IntVal: 2861335956, UnsafePtr: unsafe.Pointer(&Constructor_IterNode[gopurs_runtime.Value, gopurs_runtime.Value]{xs_5, gopurs_runtime.Value{Type: 9, IntVal: 2509360378, UnsafePtr: nil}})}, gopurs_runtime.Value{Type: 9, IntVal: 2861335956, UnsafePtr: unsafe.Pointer(&Constructor_IterNode[gopurs_runtime.Value, gopurs_runtime.Value]{ys_6, gopurs_runtime.Value{Type: 9, IntVal: 2509360378, UnsafePtr: nil}})})
}
end_branch_3:
return __t3
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
ordMap1_1_0 := gopurs_runtime.Apply(Get_ordMap(), dictOrd_0)
_ = ordMap1_1_0
__local_var_2_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd_0, "Eq0_NOT_FOUND"), gopurs_runtime.Value{})
_ = __local_var_2_1
eq1Map1_3_2 := gopurs_runtime.RecordDict1("eq1", gopurs_runtime.Func(func(dictEq1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(Call_eqMap(__local_var_2_1, dictEq1_3), "eq")
}))
_ = eq1Map1_3_2
return gopurs_runtime.RecordDict2("Eq10", "compare1", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return eq1Map1_3_2
}), gopurs_runtime.Func(func(dictOrd1_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(gopurs_runtime.Apply(ordMap1_1_0, dictOrd1_4), "compare")
}))
}

func Call_fromFoldable(dictOrd_0_loop gopurs_runtime.Value, dictFoldable_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
var dictFoldable_1 gopurs_runtime.Value = dictFoldable_1_loop
_ = dictFoldable_1
return gopurs_runtime.Apply2(((*gopurs_runtime.RecordData3)(dictFoldable_1.UnsafePtr)).V1, gopurs_runtime.Func2(func(m_2 gopurs_runtime.Value, v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Call_insert(dictOrd_0, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V0, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V1), m_2)
}), gopurs_runtime.Value{Type: 9, IntVal: 687041424, UnsafePtr: nil})
}

func Call_fromFoldableWith(dictOrd_0_loop gopurs_runtime.Value, dictFoldable_1_loop gopurs_runtime.Value, f_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
var dictFoldable_1 gopurs_runtime.Value = dictFoldable_1_loop
_ = dictFoldable_1
var f_2 gopurs_runtime.Value = f_2_loop
_ = f_2
f_prime_3_0 := gopurs_runtime.Apply2(Get_insertWith(), dictOrd_0, gopurs_runtime.Func2(func(b_3 gopurs_runtime.Value, a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_2, a_4, b_3)
}))
_ = f_prime_3_0
return gopurs_runtime.Apply2(((*gopurs_runtime.RecordData3)(dictFoldable_1.UnsafePtr)).V1, gopurs_runtime.Func2(func(m_4 gopurs_runtime.Value, v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(f_prime_3_0, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_5.UnsafePtr).V0, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_5.UnsafePtr).V1, m_4)
}), gopurs_runtime.Value{Type: 9, IntVal: 687041424, UnsafePtr: nil})
}

func Call_fromFoldableWithIndex(dictOrd_0_loop gopurs_runtime.Value, dictFoldableWithIndex_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
var dictFoldableWithIndex_1 gopurs_runtime.Value = dictFoldableWithIndex_1_loop
_ = dictFoldableWithIndex_1
return gopurs_runtime.Apply2(((*gopurs_runtime.RecordData3)(dictFoldableWithIndex_1.UnsafePtr)).V1, gopurs_runtime.Func3(func(k_2 gopurs_runtime.Value, m_3 gopurs_runtime.Value, v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Call_insert(dictOrd_0, k_2, v_4), m_3)
}), gopurs_runtime.Value{Type: 9, IntVal: 687041424, UnsafePtr: nil})
}

func Call_monoidSemigroupMap(_dollar__unused_0_loop gopurs_runtime.Value, dictOrd_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var _dollar__unused_0 gopurs_runtime.Value = _dollar__unused_0_loop
_ = _dollar__unused_0
var dictOrd_1 gopurs_runtime.Value = dictOrd_1_loop
_ = dictOrd_1
semigroupMap2_2_0 := Call_semigroupMap(gopurs_runtime.Value{}, dictOrd_1)
_ = semigroupMap2_2_0
return gopurs_runtime.Func(func(dictSemigroup_3 gopurs_runtime.Value) gopurs_runtime.Value {
semigroupMap3_4_1 := gopurs_runtime.Apply(semigroupMap2_2_0, dictSemigroup_3)
_ = semigroupMap3_4_1
return gopurs_runtime.RecordDict2("Semigroup0", "mempty", gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return semigroupMap3_4_1
}), gopurs_runtime.Value{Type: 9, IntVal: 687041424, UnsafePtr: nil})
})
}

func Call_submap(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
compare_1_0 := ((*gopurs_runtime.RecordData1)(dictOrd_0.UnsafePtr)).V0
_ = compare_1_0
return gopurs_runtime.Func2(func(kmin_2 gopurs_runtime.Value, kmax_3 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldSubmapBy(dictOrd_0, gopurs_runtime.Func2(func(m1_4 gopurs_runtime.Value, m2_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.UncurriedApp4(Get_unsafeUnionWith(), compare_1_0, pkg_Data_Function.Get_const_(), m1_4, m2_5)
}), gopurs_runtime.Value{Type: 9, IntVal: 687041424, UnsafePtr: nil}, kmin_2, kmax_3, Get_singleton())
})
}

func Call_unions(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
compare_1_0 := ((*gopurs_runtime.RecordData1)(dictOrd_0.UnsafePtr)).V0
_ = compare_1_0
return gopurs_runtime.Func(func(dictFoldable_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFoldable_2, "foldl"), gopurs_runtime.Func2(func(m1_3 gopurs_runtime.Value, m2_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.UncurriedApp4(Get_unsafeUnionWith(), compare_1_0, pkg_Data_Function.Get_const_(), m1_3, m2_4)
}), gopurs_runtime.Value{Type: 9, IntVal: 687041424, UnsafePtr: nil})
})
}

func Call_difference(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
compare_1_0 := ((*gopurs_runtime.RecordData1)(dictOrd_0.UnsafePtr)).V0
_ = compare_1_0
return gopurs_runtime.Func2(func(m1_2 gopurs_runtime.Value, m2_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.UncurriedApp3(Get_unsafeDifference(), compare_1_0, m1_2, m2_3)
})
}

func Call_delete_(dictOrd_0_loop gopurs_runtime.Value, k_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
var k_1 gopurs_runtime.Value = k_1_loop
_ = k_1
var go__2_0 gopurs_runtime.Value
_ = go__2_0
go__2_0 = gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if (v_3.Type == 9 && v_3.IntVal == 687041424) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 687041424, UnsafePtr: nil}
goto end_branch_1
} else {

}
}
{
if (v_3.Type == 9 && v_3.IntVal == 324739070) {
v1_4_2 := gopurs_runtime.Apply2(((*gopurs_runtime.RecordData1)(dictOrd_0.UnsafePtr)).V0, k_1, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V2)
_ = v1_4_2
var __t3 gopurs_runtime.Value
{
if (v1_4_2.Type == 9 && v1_4_2.IntVal == 1527465420) {
__t3 = gopurs_runtime.UncurriedApp4(Get_unsafeBalancedNode(), (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V2, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V3, gopurs_runtime.Apply(go__2_0, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V4), (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V5)
goto end_branch_3
} else {

}
}
{
if (v1_4_2.Type == 9 && v1_4_2.IntVal == 380165415) {
__t3 = gopurs_runtime.UncurriedApp4(Get_unsafeBalancedNode(), (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V2, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V3, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V4, gopurs_runtime.Apply(go__2_0, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V5))
goto end_branch_3
} else {

}
}
{
if (v1_4_2.Type == 9 && v1_4_2.IntVal == 902936544) {
__t3 = gopurs_runtime.UncurriedApp2(Get_unsafeJoinNodes(), (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V4, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V5)
goto end_branch_3
} else {

}
}
{
__t3 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_3:
__t1 = __t3
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
return go__2_0
}

func Call_checkValid(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
var go__1_0 gopurs_runtime.Value
_ = go__1_0
go__1_0 = gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if (v_2.Type == 9 && v_2.IntVal == 687041424) {
__t1 = gopurs_runtime.Bool(true)
goto end_branch_1
} else {

}
}
{
if (v_2.Type == 9 && v_2.IntVal == 324739070) {
var __t2 gopurs_runtime.Value
{
var __t_tag_3 gopurs_runtime.Value = (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V4
if (__t_tag_3.Type == 9 && __t_tag_3.IntVal == 687041424) {
var __t4 gopurs_runtime.Value
{
var __t_tag_5 gopurs_runtime.Value = (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V5
if (__t_tag_5.Type == 9 && __t_tag_5.IntVal == 687041424) {
__t4 = gopurs_runtime.Bool(true)
goto end_branch_4
} else {

}
}
{
var __t_tag_6 gopurs_runtime.Value = (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V5
if (__t_tag_6.Type == 9 && __t_tag_6.IntVal == 324739070) {
var __t_tag_7 gopurs_runtime.Value = gopurs_runtime.Apply2(((*gopurs_runtime.RecordData1)(dictOrd_0.UnsafePtr)).V0, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V5.UnsafePtr).V2, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V2)
__t4 = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_HeytingAlgebra.Get_heytingAlgebraBoolean(), "conj"), gopurs_runtime.Bool((gopurs_runtime.Int((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V0).IntVal) == (2)), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_HeytingAlgebra.Get_heytingAlgebraBoolean(), "conj"), gopurs_runtime.Bool((gopurs_runtime.Int((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V5.UnsafePtr).V0).IntVal) == (1)), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_HeytingAlgebra.Get_heytingAlgebraBoolean(), "conj"), gopurs_runtime.Apply2(Get_greaterThan(), gopurs_runtime.Int((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V1), gopurs_runtime.Int((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V5.UnsafePtr).V1)), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_HeytingAlgebra.Get_heytingAlgebraBoolean(), "conj"), gopurs_runtime.Bool((__t_tag_7.Type == 9 && __t_tag_7.IntVal == 380165415)), gopurs_runtime.Apply(go__1_0, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V5)))))
goto end_branch_4
} else {

}
}
{
__t4 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_4:
__t2 = __t4
goto end_branch_2
} else {

}
}
{
var __t_tag_8 gopurs_runtime.Value = (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V4
if (__t_tag_8.Type == 9 && __t_tag_8.IntVal == 324739070) {
var __t9 gopurs_runtime.Value
{
var __t_tag_10 gopurs_runtime.Value = (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V5
if (__t_tag_10.Type == 9 && __t_tag_10.IntVal == 687041424) {
var __t_tag_11 gopurs_runtime.Value = gopurs_runtime.Apply2(((*gopurs_runtime.RecordData1)(dictOrd_0.UnsafePtr)).V0, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V4.UnsafePtr).V2, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V2)
__t9 = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_HeytingAlgebra.Get_heytingAlgebraBoolean(), "conj"), gopurs_runtime.Bool((gopurs_runtime.Int((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V0).IntVal) == (2)), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_HeytingAlgebra.Get_heytingAlgebraBoolean(), "conj"), gopurs_runtime.Bool((gopurs_runtime.Int((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V4.UnsafePtr).V0).IntVal) == (1)), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_HeytingAlgebra.Get_heytingAlgebraBoolean(), "conj"), gopurs_runtime.Apply2(Get_greaterThan(), gopurs_runtime.Int((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V1), gopurs_runtime.Int((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V4.UnsafePtr).V1)), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_HeytingAlgebra.Get_heytingAlgebraBoolean(), "conj"), gopurs_runtime.Bool((__t_tag_11.Type == 9 && __t_tag_11.IntVal == 1527465420)), gopurs_runtime.Apply(go__1_0, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V4)))))
goto end_branch_9
} else {

}
}
{
var __t_tag_12 gopurs_runtime.Value = (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V5
if (__t_tag_12.Type == 9 && __t_tag_12.IntVal == 324739070) {
var __t_tag_13 gopurs_runtime.Value = gopurs_runtime.Apply2(((*gopurs_runtime.RecordData1)(dictOrd_0.UnsafePtr)).V0, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V5.UnsafePtr).V2, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V2)
var __t_tag_14 gopurs_runtime.Value = gopurs_runtime.Apply2(((*gopurs_runtime.RecordData1)(dictOrd_0.UnsafePtr)).V0, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V4.UnsafePtr).V2, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V2)
__t9 = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_HeytingAlgebra.Get_heytingAlgebraBoolean(), "conj"), gopurs_runtime.Apply2(Get_greaterThan(), gopurs_runtime.Int((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V0), gopurs_runtime.Int((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V5.UnsafePtr).V0)), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_HeytingAlgebra.Get_heytingAlgebraBoolean(), "conj"), gopurs_runtime.Bool((__t_tag_13.Type == 9 && __t_tag_13.IntVal == 380165415)), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_HeytingAlgebra.Get_heytingAlgebraBoolean(), "conj"), gopurs_runtime.Apply2(Get_greaterThan(), gopurs_runtime.Int((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V0), gopurs_runtime.Int((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V4.UnsafePtr).V0)), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_HeytingAlgebra.Get_heytingAlgebraBoolean(), "conj"), gopurs_runtime.Bool((__t_tag_14.Type == 9 && __t_tag_14.IntVal == 1527465420)), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_HeytingAlgebra.Get_heytingAlgebraBoolean(), "conj"), gopurs_runtime.Apply2(Get_lessThan(), gopurs_runtime.Apply(Get_abs(), gopurs_runtime.Int((gopurs_runtime.Int((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V5.UnsafePtr).V0).IntVal) - (gopurs_runtime.Int((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V4.UnsafePtr).V0).IntVal))), gopurs_runtime.Int(2)), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_HeytingAlgebra.Get_heytingAlgebraBoolean(), "conj"), gopurs_runtime.Bool((((gopurs_runtime.Int((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V5.UnsafePtr).V1).IntVal) + (gopurs_runtime.Int((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V4.UnsafePtr).V1).IntVal)) + (1)) == (gopurs_runtime.Int((*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V1).IntVal)), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_HeytingAlgebra.Get_heytingAlgebraBoolean(), "conj"), gopurs_runtime.Apply(go__1_0, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V4), gopurs_runtime.Apply(go__1_0, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V5))))))))
goto end_branch_9
} else {

}
}
{
__t9 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_9:
__t2 = __t9
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_2:
__t1 = __t2
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
return go__1_0
}

func Call_catMaybes(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
return Call_mapMaybeWithKey(dictOrd_0, gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_identity()
}))
}

func Call_applyMap(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
compare_1_0 := ((*gopurs_runtime.RecordData1)(dictOrd_0.UnsafePtr)).V0
_ = compare_1_0
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_functorMap()
}), gopurs_runtime.Func2(func(m1_2 gopurs_runtime.Value, m2_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.UncurriedApp4(Get_unsafeIntersectionWith(), compare_1_0, Get_identity(), m1_2, m2_3)
}))
}

func Call_bindMap(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
compare_1_1 := ((*gopurs_runtime.RecordData1)(dictOrd_0.UnsafePtr)).V0
_ = compare_1_1
applyMap1_1_0 := gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_functorMap()
}), gopurs_runtime.Func2(func(m1_2 gopurs_runtime.Value, m2_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.UncurriedApp4(Get_unsafeIntersectionWith(), compare_1_1, Get_identity(), m1_2, m2_3)
}))
_ = applyMap1_1_0
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return applyMap1_1_0
}), gopurs_runtime.Func2(func(m_2 gopurs_runtime.Value, f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Call_mapMaybeWithKey(dictOrd_0, gopurs_runtime.Func(func(k_4 gopurs_runtime.Value) gopurs_runtime.Value {
var go__5_2 gopurs_runtime.Value
go__5_2 = gopurs_runtime.Func(func(v_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_6_loop gopurs_runtime.Value = v_6_loop_val
go__5_2:
for {
if false { continue go__5_2 }
var v_6 gopurs_runtime.Value = v_6_loop
_ = v_6
var __t3 gopurs_runtime.Value
{
if (v_6.Type == 9 && v_6.IntVal == 687041424) {
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 3589588149, UnsafePtr: nil}
goto end_branch_3
} else {

}
}
{
if (v_6.Type == 9 && v_6.IntVal == 324739070) {
v1_7_4 := gopurs_runtime.Apply2(((*gopurs_runtime.RecordData1)(dictOrd_0.UnsafePtr)).V0, k_4, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_6.UnsafePtr).V2)
_ = v1_7_4
var __t5 gopurs_runtime.Value
{
if (v1_7_4.Type == 9 && v1_7_4.IntVal == 1527465420) {
v_6_loop = (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_6.UnsafePtr).V4
continue go__5_2
__t5 = gopurs_runtime.Value{}
goto end_branch_5
} else {

}
}
{
if (v1_7_4.Type == 9 && v1_7_4.IntVal == 380165415) {
v_6_loop = (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_6.UnsafePtr).V5
continue go__5_2
__t5 = gopurs_runtime.Value{}
goto end_branch_5
} else {

}
}
{
if (v1_7_4.Type == 9 && v1_7_4.IntVal == 902936544) {
__t5 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{(*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_6.UnsafePtr).V3})}
goto end_branch_5
} else {

}
}
{
__t5 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_5:
__t3 = __t5
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
return gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(go__5_2, gopurs_runtime.Apply(f_3, x_6))
})
})), m_2)
}))
}

func Call_anyWithKey(predicate_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var predicate_0 gopurs_runtime.Value = predicate_0_loop
_ = predicate_0
var go__1_0 gopurs_runtime.Value
_ = go__1_0
go__1_0 = gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if (v_2.Type == 9 && v_2.IntVal == 687041424) {
__t1 = gopurs_runtime.Bool(false)
goto end_branch_1
} else {

}
}
{
if (v_2.Type == 9 && v_2.IntVal == 324739070) {
__t1 = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_HeytingAlgebra.Get_heytingAlgebraBoolean(), "disj"), gopurs_runtime.Apply2(predicate_0, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V2, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V3), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_HeytingAlgebra.Get_heytingAlgebraBoolean(), "disj"), gopurs_runtime.Apply(go__1_0, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V4), gopurs_runtime.Apply(go__1_0, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V5)))
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
return go__1_0
}

func Call_any(predicate_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var predicate_0 gopurs_runtime.Value = predicate_0_loop
_ = predicate_0
var go__1_0 gopurs_runtime.Value
_ = go__1_0
go__1_0 = gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if (v_2.Type == 9 && v_2.IntVal == 687041424) {
__t1 = gopurs_runtime.Bool(false)
goto end_branch_1
} else {

}
}
{
if (v_2.Type == 9 && v_2.IntVal == 324739070) {
__t1 = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_HeytingAlgebra.Get_heytingAlgebraBoolean(), "disj"), gopurs_runtime.Apply(predicate_0, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V3), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_HeytingAlgebra.Get_heytingAlgebraBoolean(), "disj"), gopurs_runtime.Apply(go__1_0, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V4), gopurs_runtime.Apply(go__1_0, (*Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V5)))
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
return go__1_0
}

func Call_alter(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
compare_1_0 := ((*gopurs_runtime.RecordData1)(dictOrd_0.UnsafePtr)).V0
_ = compare_1_0
return gopurs_runtime.Func3(func(f_2 gopurs_runtime.Value, k_3 gopurs_runtime.Value, m_4 gopurs_runtime.Value) gopurs_runtime.Value {
v_5_1 := gopurs_runtime.UncurriedApp3(Get_unsafeSplit(), compare_1_0, k_3, m_4)
_ = v_5_1
v2_6_2 := gopurs_runtime.Apply(f_2, (*Constructor_Split[gopurs_runtime.Value, gopurs_runtime.Value])(v_5_1.UnsafePtr).V0)
_ = v2_6_2
var __t3 gopurs_runtime.Value
{
if (v2_6_2.Type == 9 && v2_6_2.IntVal == 3589588149) {
__t3 = gopurs_runtime.UncurriedApp2(Get_unsafeJoinNodes(), (*Constructor_Split[gopurs_runtime.Value, gopurs_runtime.Value])(v_5_1.UnsafePtr).V1, (*Constructor_Split[gopurs_runtime.Value, gopurs_runtime.Value])(v_5_1.UnsafePtr).V2)
goto end_branch_3
} else {

}
}
{
if (v2_6_2.Type == 9 && v2_6_2.IntVal == 930809136) {
__t3 = gopurs_runtime.UncurriedApp4(Get_unsafeBalancedNode(), k_3, (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(v2_6_2.UnsafePtr).V0, (*Constructor_Split[gopurs_runtime.Value, gopurs_runtime.Value])(v_5_1.UnsafePtr).V1, (*Constructor_Split[gopurs_runtime.Value, gopurs_runtime.Value])(v_5_1.UnsafePtr).V2)
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
}

func Call_altMap(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
compare_1_0 := ((*gopurs_runtime.RecordData1)(dictOrd_0.UnsafePtr)).V0
_ = compare_1_0
return gopurs_runtime.RecordDict2("Functor0", "alt", gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_functorMap()
}), gopurs_runtime.Func2(func(m1_2 gopurs_runtime.Value, m2_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.UncurriedApp4(Get_unsafeUnionWith(), compare_1_0, pkg_Data_Function.Get_const_(), m1_2, m2_3)
}))
}

func Call_plusMap(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
compare_1_1 := ((*gopurs_runtime.RecordData1)(dictOrd_0.UnsafePtr)).V0
_ = compare_1_1
altMap1_1_0 := gopurs_runtime.RecordDict2("Functor0", "alt", gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_functorMap()
}), gopurs_runtime.Func2(func(m1_2 gopurs_runtime.Value, m2_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.UncurriedApp4(Get_unsafeUnionWith(), compare_1_1, pkg_Data_Function.Get_const_(), m1_2, m2_3)
}))
_ = altMap1_1_0
return gopurs_runtime.RecordDict2("Alt0", "empty", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return altMap1_1_0
}), gopurs_runtime.Value{Type: 9, IntVal: 687041424, UnsafePtr: nil})
}


