package Data_Map_Internal

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_Function "gopurs/output/Data.Function"
	pkg_Data_Maybe "gopurs/output/Data.Maybe"
	pkg_Data_Tuple "gopurs/output/Data.Tuple"
	pkg_Data_Ordering "gopurs/output/Data.Ordering"
	pkg_Data_Unfoldable "gopurs/output/Data.Unfoldable"
	pkg_Data_List_Types "gopurs/output/Data.List.Types"
	pkg_Control_Category "gopurs/output/Control.Category"
	pkg_Data_Unit "gopurs/output/Data.Unit"
	pkg_Data_Show "gopurs/output/Data.Show"
	unsafe "unsafe"
)

var Leaf gopurs_runtime.Value
var once_Leaf sync.Once
func Get_Leaf() gopurs_runtime.Value {
	once_Leaf.Do(func() {
		Leaf = gopurs_runtime.Value{Type: 9, IntVal: 687041424, UnsafePtr: unsafe.Pointer(&Data_Data_Map_Internal_Leaf{})}
	})
	return Leaf
}

var Node gopurs_runtime.Value
var once_Node sync.Once
func Get_Node() gopurs_runtime.Value {
	once_Node.Do(func() {
		Node = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(value1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(value2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(value3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(value4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(value5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(&Data_Data_Map_Internal_Node{value0, value1, value2, value3, value4, value5})}
})
})
})
})
})
})
	})
	return Node
}

var IterLeaf gopurs_runtime.Value
var once_IterLeaf sync.Once
func Get_IterLeaf() gopurs_runtime.Value {
	once_IterLeaf.Do(func() {
		IterLeaf = gopurs_runtime.Value{Type: 9, IntVal: 2509360378, UnsafePtr: unsafe.Pointer(&Data_Data_Map_Internal_IterLeaf{})}
	})
	return IterLeaf
}

var IterEmit gopurs_runtime.Value
var once_IterEmit sync.Once
func Get_IterEmit() gopurs_runtime.Value {
	once_IterEmit.Do(func() {
		IterEmit = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(value1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(value2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1343415489, UnsafePtr: unsafe.Pointer(&Data_Data_Map_Internal_IterEmit{value0, value1, value2})}
})
})
})
	})
	return IterEmit
}

var IterNode gopurs_runtime.Value
var once_IterNode sync.Once
func Get_IterNode() gopurs_runtime.Value {
	once_IterNode.Do(func() {
		IterNode = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(value1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2861335956, UnsafePtr: unsafe.Pointer(&Data_Data_Map_Internal_IterNode{value0, value1})}
})
})
	})
	return IterNode
}

var IterDone gopurs_runtime.Value
var once_IterDone sync.Once
func Get_IterDone() gopurs_runtime.Value {
	once_IterDone.Do(func() {
		IterDone = gopurs_runtime.Value{Type: 9, IntVal: 4236111124, UnsafePtr: unsafe.Pointer(&Data_Data_Map_Internal_IterDone{})}
	})
	return IterDone
}

var IterNext gopurs_runtime.Value
var once_IterNext sync.Once
func Get_IterNext() gopurs_runtime.Value {
	once_IterNext.Do(func() {
		IterNext = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(value1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(value2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 953589075, UnsafePtr: unsafe.Pointer(&Data_Data_Map_Internal_IterNext{value0, value1, value2})}
})
})
})
	})
	return IterNext
}

var Split gopurs_runtime.Value
var once_Split sync.Once
func Get_Split() gopurs_runtime.Value {
	once_Split.Do(func() {
		Split = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(value1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(value2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3373277644, UnsafePtr: unsafe.Pointer(&Data_Data_Map_Internal_Split{value0, value1, value2})}
})
})
})
	})
	return Split
}

var SplitLast gopurs_runtime.Value
var once_SplitLast sync.Once
func Get_SplitLast() gopurs_runtime.Value {
	once_SplitLast.Do(func() {
		SplitLast = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(value1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(value2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2668112006, UnsafePtr: unsafe.Pointer(&Data_Data_Map_Internal_SplitLast{value0, value1, value2})}
})
})
})
	})
	return SplitLast
}

var unsafeNode gopurs_runtime.Value
var once_unsafeNode sync.Once
func Get_unsafeNode() gopurs_runtime.Value {
	once_unsafeNode.Do(func() {
		unsafeNode = gopurs_runtime.Func4(func(k_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value, l_2_box gopurs_runtime.Value, r_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_unsafeNode(k_0_box, v_1_box, l_2_box, r_3_box)
})
	})
	return unsafeNode
}

var toMapIter gopurs_runtime.Value
var once_toMapIter sync.Once
func Get_toMapIter() gopurs_runtime.Value {
	once_toMapIter.Do(func() {
		toMapIter = gopurs_runtime.Func(func(a_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var a_0 gopurs_runtime.Value = a_0_loop
_ = a_0
return gopurs_runtime.Value{Type: 9, IntVal: 2861335956, UnsafePtr: unsafe.Pointer(&Data_Data_Map_Internal_IterNode{a_0, gopurs_runtime.Value{Type: 9, IntVal: 2509360378, UnsafePtr: unsafe.Pointer(&Data_Data_Map_Internal_IterLeaf{})}})}
}()
})
	})
	return toMapIter
}

var stepWith gopurs_runtime.Value
var once_stepWith sync.Once
func Get_stepWith() gopurs_runtime.Value {
	once_stepWith.Do(func() {
		stepWith = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, next_1_box gopurs_runtime.Value, done_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_stepWith(f_0_box, next_1_box, done_2_box)
})
	})
	return stepWith
}

var size gopurs_runtime.Value
var once_size sync.Once
func Get_size() gopurs_runtime.Value {
	once_size.Do(func() {
		size = gopurs_runtime.Func(func(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var __t0 gopurs_runtime.Value
{
if (v_0.Type == 9 && v_0.IntVal == 687041424) {
__t0 = gopurs_runtime.Int(0)
goto end_branch_0
} else {

}
}
{
if (v_0.Type == 9 && v_0.IntVal == 324739070) {
__t0 = (*Data_Data_Map_Internal_Node)(v_0.UnsafePtr).V1
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return __t0
}()
})
	})
	return size
}

var singleton gopurs_runtime.Value
var once_singleton sync.Once
func Get_singleton() gopurs_runtime.Value {
	once_singleton.Do(func() {
		singleton = gopurs_runtime.Func2(func(k_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_singleton(k_0_box, v_1_box)
})
	})
	return singleton
}

var unsafeBalancedNode gopurs_runtime.Value
var once_unsafeBalancedNode sync.Once
func Get_unsafeBalancedNode() gopurs_runtime.Value {
	once_unsafeBalancedNode.Do(func() {
		unsafeBalancedNode = gopurs_runtime.Func4(func(k_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value, l_2_box gopurs_runtime.Value, r_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_unsafeBalancedNode(k_0_box, v_1_box, l_2_box, r_3_box)
})
	})
	return unsafeBalancedNode
}

var unsafeSplit gopurs_runtime.Value
var once_unsafeSplit sync.Once
func Get_unsafeSplit() gopurs_runtime.Value {
	once_unsafeSplit.Do(func() {
		unsafeSplit = gopurs_runtime.Func3(func(comp_0_box gopurs_runtime.Value, k_1_box gopurs_runtime.Value, m_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_unsafeSplit(comp_0_box, k_1_box, m_2_box)
})
	})
	return unsafeSplit
}

var unsafeSplitLast gopurs_runtime.Value
var once_unsafeSplitLast sync.Once
func Get_unsafeSplitLast() gopurs_runtime.Value {
	once_unsafeSplitLast.Do(func() {
		unsafeSplitLast = gopurs_runtime.Func4(func(k_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value, l_2_box gopurs_runtime.Value, r_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_unsafeSplitLast(k_0_box, v_1_box, l_2_box, r_3_box)
})
	})
	return unsafeSplitLast
}

var unsafeJoinNodes gopurs_runtime.Value
var once_unsafeJoinNodes sync.Once
func Get_unsafeJoinNodes() gopurs_runtime.Value {
	once_unsafeJoinNodes.Do(func() {
		unsafeJoinNodes = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_unsafeJoinNodes(v_0_box, v1_1_box)
})
	})
	return unsafeJoinNodes
}

var unsafeDifference gopurs_runtime.Value
var once_unsafeDifference sync.Once
func Get_unsafeDifference() gopurs_runtime.Value {
	once_unsafeDifference.Do(func() {
		unsafeDifference = gopurs_runtime.Func3(func(comp_0_box gopurs_runtime.Value, l_1_box gopurs_runtime.Value, r_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_unsafeDifference(comp_0_box, l_1_box, r_2_box)
})
	})
	return unsafeDifference
}

var unsafeIntersectionWith gopurs_runtime.Value
var once_unsafeIntersectionWith sync.Once
func Get_unsafeIntersectionWith() gopurs_runtime.Value {
	once_unsafeIntersectionWith.Do(func() {
		unsafeIntersectionWith = gopurs_runtime.Func4(func(comp_0_box gopurs_runtime.Value, app_1_box gopurs_runtime.Value, l_2_box gopurs_runtime.Value, r_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_unsafeIntersectionWith(comp_0_box, app_1_box, l_2_box, r_3_box)
})
	})
	return unsafeIntersectionWith
}

var unsafeUnionWith gopurs_runtime.Value
var once_unsafeUnionWith sync.Once
func Get_unsafeUnionWith() gopurs_runtime.Value {
	once_unsafeUnionWith.Do(func() {
		unsafeUnionWith = gopurs_runtime.Func4(func(comp_0_box gopurs_runtime.Value, app_1_box gopurs_runtime.Value, l_2_box gopurs_runtime.Value, r_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_unsafeUnionWith(comp_0_box, app_1_box, l_2_box, r_3_box)
})
	})
	return unsafeUnionWith
}

var unionWith gopurs_runtime.Value
var once_unionWith sync.Once
func Get_unionWith() gopurs_runtime.Value {
	once_unionWith.Do(func() {
		unionWith = gopurs_runtime.Func(func(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
compare_1_0 := gopurs_runtime.RecordGet(dictOrd_0, "compare")
_ = compare_1_0
return gopurs_runtime.Func3(func(app_2 gopurs_runtime.Value, m1_3 gopurs_runtime.Value, m2_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.UncurriedApp4(Get_unsafeUnionWith(), compare_1_0, app_2, m1_3, m2_4)
})
}()
})
	})
	return unionWith
}

var union gopurs_runtime.Value
var once_union sync.Once
func Get_union() gopurs_runtime.Value {
	once_union.Do(func() {
		union = gopurs_runtime.Func(func(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
compare_1_0 := gopurs_runtime.RecordGet(dictOrd_0, "compare")
_ = compare_1_0
return gopurs_runtime.Func2(func(m1_2 gopurs_runtime.Value, m2_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.UncurriedApp4(Get_unsafeUnionWith(), compare_1_0, pkg_Data_Function.Get_const_(), m1_2, m2_3)
})
}()
})
	})
	return union
}

var update gopurs_runtime.Value
var once_update sync.Once
func Get_update() gopurs_runtime.Value {
	once_update.Do(func() {
		update = gopurs_runtime.Func3(func(dictOrd_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, k_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_update(dictOrd_0_box, f_1_box, k_2_box)
})
	})
	return update
}

var showTree gopurs_runtime.Value
var once_showTree sync.Once
func Get_showTree() gopurs_runtime.Value {
	once_showTree.Do(func() {
		showTree = gopurs_runtime.Func2(func(dictShow_0_box gopurs_runtime.Value, dictShow1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_showTree(dictShow_0_box, dictShow1_1_box)
})
	})
	return showTree
}

var semigroupMap gopurs_runtime.Value
var once_semigroupMap sync.Once
func Get_semigroupMap() gopurs_runtime.Value {
	once_semigroupMap.Do(func() {
		semigroupMap = gopurs_runtime.Func2(func(_dollar__unused_0_box gopurs_runtime.Value, dictOrd_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_semigroupMap(_dollar__unused_0_box, dictOrd_1_box)
})
	})
	return semigroupMap
}

var pop gopurs_runtime.Value
var once_pop sync.Once
func Get_pop() gopurs_runtime.Value {
	once_pop.Do(func() {
		pop = gopurs_runtime.Func(func(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
compare_1_0 := gopurs_runtime.RecordGet(dictOrd_0, "compare")
_ = compare_1_0
return gopurs_runtime.Func2(func(k_2 gopurs_runtime.Value, m_3 gopurs_runtime.Value) gopurs_runtime.Value {
v_4_1 := gopurs_runtime.UncurriedApp3(Get_unsafeSplit(), compare_1_0, k_2, m_3)
_ = v_4_1
var __t2 gopurs_runtime.Value
{
if ((*Data_Data_Map_Internal_Split)(v_4_1.UnsafePtr).V0.Type == 9 && (*Data_Data_Map_Internal_Split)(v_4_1.UnsafePtr).V0.IntVal == 1354639136) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 1354639136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Just{gopurs_runtime.Value{Type: 9, IntVal: 1102100576, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Data_Data_Tuple_Tuple{(*pkg_Data_Maybe.Data_Data_Maybe_Just)((*Data_Data_Map_Internal_Split)(v_4_1.UnsafePtr).V0.UnsafePtr).V0, gopurs_runtime.UncurriedApp2(Get_unsafeJoinNodes(), (*Data_Data_Map_Internal_Split)(v_4_1.UnsafePtr).V1, (*Data_Data_Map_Internal_Split)(v_4_1.UnsafePtr).V2)})}})}
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 42808261, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Nothing{})}
}
end_branch_2:
return __t2
})
}()
})
	})
	return pop
}

var member gopurs_runtime.Value
var once_member sync.Once
func Get_member() gopurs_runtime.Value {
	once_member.Do(func() {
		member = gopurs_runtime.Func2(func(dictOrd_0_box gopurs_runtime.Value, k_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_member(dictOrd_0_box, k_1_box)
})
	})
	return member
}

var mapMaybeWithKey gopurs_runtime.Value
var once_mapMaybeWithKey sync.Once
func Get_mapMaybeWithKey() gopurs_runtime.Value {
	once_mapMaybeWithKey.Do(func() {
		mapMaybeWithKey = gopurs_runtime.Func2(func(dictOrd_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_mapMaybeWithKey(dictOrd_0_box, f_1_box)
})
	})
	return mapMaybeWithKey
}

var mapMaybe gopurs_runtime.Value
var once_mapMaybe sync.Once
func Get_mapMaybe() gopurs_runtime.Value {
	once_mapMaybe.Do(func() {
		mapMaybe = gopurs_runtime.Func2(func(dictOrd_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_mapMaybe(dictOrd_0_box, x_1_box)
})
	})
	return mapMaybe
}

var lookupLE gopurs_runtime.Value
var once_lookupLE sync.Once
func Get_lookupLE() gopurs_runtime.Value {
	once_lookupLE.Do(func() {
		lookupLE = gopurs_runtime.Func2(func(dictOrd_0_box gopurs_runtime.Value, k_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_lookupLE(dictOrd_0_box, k_1_box)
})
	})
	return lookupLE
}

var lookupGE gopurs_runtime.Value
var once_lookupGE sync.Once
func Get_lookupGE() gopurs_runtime.Value {
	once_lookupGE.Do(func() {
		lookupGE = gopurs_runtime.Func2(func(dictOrd_0_box gopurs_runtime.Value, k_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_lookupGE(dictOrd_0_box, k_1_box)
})
	})
	return lookupGE
}

var lookup gopurs_runtime.Value
var once_lookup sync.Once
func Get_lookup() gopurs_runtime.Value {
	once_lookup.Do(func() {
		lookup = gopurs_runtime.Func2(func(dictOrd_0_box gopurs_runtime.Value, k_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_lookup(dictOrd_0_box, k_1_box)
})
	})
	return lookup
}

var iterMapU gopurs_runtime.Value
var once_iterMapU sync.Once
func Get_iterMapU() gopurs_runtime.Value {
	once_iterMapU.Do(func() {
		iterMapU = gopurs_runtime.Func2(func(iter_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_iterMapU(iter_0_box, v_1_box)
})
	})
	return iterMapU
}

var stepUnorderedCps gopurs_runtime.Value
var once_stepUnorderedCps sync.Once
func Get_stepUnorderedCps() gopurs_runtime.Value {
	once_stepUnorderedCps.Do(func() {
		stepUnorderedCps = gopurs_runtime.Apply(Get_stepWith(), Get_iterMapU())
	})
	return stepUnorderedCps
}

var stepUnfoldrUnordered gopurs_runtime.Value
var once_stepUnfoldrUnordered sync.Once
func Get_stepUnfoldrUnordered() gopurs_runtime.Value {
	once_stepUnfoldrUnordered.Do(func() {
		stepUnfoldrUnordered = Call_stepWith(Get_iterMapU(), gopurs_runtime.Func3(func(k_0 gopurs_runtime.Value, v_1 gopurs_runtime.Value, next_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1354639136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Just{gopurs_runtime.Value{Type: 9, IntVal: 1102100576, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Data_Data_Tuple_Tuple{gopurs_runtime.Value{Type: 9, IntVal: 1102100576, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Data_Data_Tuple_Tuple{k_0, v_1})}, next_2})}})}
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 42808261, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Nothing{})}
}))
	})
	return stepUnfoldrUnordered
}

var toUnfoldableUnordered gopurs_runtime.Value
var once_toUnfoldableUnordered sync.Once
func Get_toUnfoldableUnordered() gopurs_runtime.Value {
	once_toUnfoldableUnordered.Do(func() {
		toUnfoldableUnordered = gopurs_runtime.Func(func(dictUnfoldable_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictUnfoldable_0 gopurs_runtime.Value = dictUnfoldable_0_loop
_ = dictUnfoldable_0
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictUnfoldable_0, "unfoldr"), Get_stepUnfoldrUnordered())
_ = __local_var_1_0
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_1_0, gopurs_runtime.Value{Type: 9, IntVal: 2861335956, UnsafePtr: unsafe.Pointer(&Data_Data_Map_Internal_IterNode{x_2, gopurs_runtime.Value{Type: 9, IntVal: 2509360378, UnsafePtr: unsafe.Pointer(&Data_Data_Map_Internal_IterLeaf{})}})})
})
}()
})
	})
	return toUnfoldableUnordered
}

var stepUnordered gopurs_runtime.Value
var once_stepUnordered sync.Once
func Get_stepUnordered() gopurs_runtime.Value {
	once_stepUnordered.Do(func() {
		stepUnordered = Call_stepWith(Get_iterMapU(), gopurs_runtime.Func3(func(k_0 gopurs_runtime.Value, v_1 gopurs_runtime.Value, next_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 953589075, UnsafePtr: unsafe.Pointer(&Data_Data_Map_Internal_IterNext{k_0, v_1, next_2})}
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4236111124, UnsafePtr: unsafe.Pointer(&Data_Data_Map_Internal_IterDone{})}
}))
	})
	return stepUnordered
}

var iterMapR gopurs_runtime.Value
var once_iterMapR sync.Once
func Get_iterMapR() gopurs_runtime.Value {
	once_iterMapR.Do(func() {
		iterMapR = func() gopurs_runtime.Value {
var go__0_0 gopurs_runtime.Value
go__0_0 = gopurs_runtime.Func(func(iter_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
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
if ((*Data_Data_Map_Internal_Node)(v_2.UnsafePtr).V5.Type == 9 && (*Data_Data_Map_Internal_Node)(v_2.UnsafePtr).V5.IntVal == 687041424) {
iter_1_loop = gopurs_runtime.Value{Type: 9, IntVal: 1343415489, UnsafePtr: unsafe.Pointer(&Data_Data_Map_Internal_IterEmit{(*Data_Data_Map_Internal_Node)(v_2.UnsafePtr).V2, (*Data_Data_Map_Internal_Node)(v_2.UnsafePtr).V3, iter_1})}
v_2_loop = (*Data_Data_Map_Internal_Node)(v_2.UnsafePtr).V4
continue go__0_0
__t2 = gopurs_runtime.Value{}
goto end_branch_2
} else {

}
}
{
iter_1_loop = gopurs_runtime.Value{Type: 9, IntVal: 1343415489, UnsafePtr: unsafe.Pointer(&Data_Data_Map_Internal_IterEmit{(*Data_Data_Map_Internal_Node)(v_2.UnsafePtr).V2, (*Data_Data_Map_Internal_Node)(v_2.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 2861335956, UnsafePtr: unsafe.Pointer(&Data_Data_Map_Internal_IterNode{(*Data_Data_Map_Internal_Node)(v_2.UnsafePtr).V4, iter_1})}})}
v_2_loop = (*Data_Data_Map_Internal_Node)(v_2.UnsafePtr).V5
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
	return iterMapR
}

var stepDescCps gopurs_runtime.Value
var once_stepDescCps sync.Once
func Get_stepDescCps() gopurs_runtime.Value {
	once_stepDescCps.Do(func() {
		stepDescCps = gopurs_runtime.Apply(Get_stepWith(), Get_iterMapR())
	})
	return stepDescCps
}

var stepDesc gopurs_runtime.Value
var once_stepDesc sync.Once
func Get_stepDesc() gopurs_runtime.Value {
	once_stepDesc.Do(func() {
		stepDesc = Call_stepWith(Get_iterMapR(), gopurs_runtime.Func3(func(k_0 gopurs_runtime.Value, v_1 gopurs_runtime.Value, next_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 953589075, UnsafePtr: unsafe.Pointer(&Data_Data_Map_Internal_IterNext{k_0, v_1, next_2})}
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4236111124, UnsafePtr: unsafe.Pointer(&Data_Data_Map_Internal_IterDone{})}
}))
	})
	return stepDesc
}

var iterMapL gopurs_runtime.Value
var once_iterMapL sync.Once
func Get_iterMapL() gopurs_runtime.Value {
	once_iterMapL.Do(func() {
		iterMapL = func() gopurs_runtime.Value {
var go__0_0 gopurs_runtime.Value
go__0_0 = gopurs_runtime.Func(func(iter_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
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
if ((*Data_Data_Map_Internal_Node)(v_2.UnsafePtr).V5.Type == 9 && (*Data_Data_Map_Internal_Node)(v_2.UnsafePtr).V5.IntVal == 687041424) {
iter_1_loop = gopurs_runtime.Value{Type: 9, IntVal: 1343415489, UnsafePtr: unsafe.Pointer(&Data_Data_Map_Internal_IterEmit{(*Data_Data_Map_Internal_Node)(v_2.UnsafePtr).V2, (*Data_Data_Map_Internal_Node)(v_2.UnsafePtr).V3, iter_1})}
v_2_loop = (*Data_Data_Map_Internal_Node)(v_2.UnsafePtr).V4
continue go__0_0
__t2 = gopurs_runtime.Value{}
goto end_branch_2
} else {

}
}
{
iter_1_loop = gopurs_runtime.Value{Type: 9, IntVal: 1343415489, UnsafePtr: unsafe.Pointer(&Data_Data_Map_Internal_IterEmit{(*Data_Data_Map_Internal_Node)(v_2.UnsafePtr).V2, (*Data_Data_Map_Internal_Node)(v_2.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 2861335956, UnsafePtr: unsafe.Pointer(&Data_Data_Map_Internal_IterNode{(*Data_Data_Map_Internal_Node)(v_2.UnsafePtr).V5, iter_1})}})}
v_2_loop = (*Data_Data_Map_Internal_Node)(v_2.UnsafePtr).V4
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
	return iterMapL
}

var stepAscCps gopurs_runtime.Value
var once_stepAscCps sync.Once
func Get_stepAscCps() gopurs_runtime.Value {
	once_stepAscCps.Do(func() {
		stepAscCps = gopurs_runtime.Apply(Get_stepWith(), Get_iterMapL())
	})
	return stepAscCps
}

var stepAsc gopurs_runtime.Value
var once_stepAsc sync.Once
func Get_stepAsc() gopurs_runtime.Value {
	once_stepAsc.Do(func() {
		stepAsc = Call_stepWith(Get_iterMapL(), gopurs_runtime.Func3(func(k_0 gopurs_runtime.Value, v_1 gopurs_runtime.Value, next_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 953589075, UnsafePtr: unsafe.Pointer(&Data_Data_Map_Internal_IterNext{k_0, v_1, next_2})}
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4236111124, UnsafePtr: unsafe.Pointer(&Data_Data_Map_Internal_IterDone{})}
}))
	})
	return stepAsc
}

var eqMapIter gopurs_runtime.Value
var once_eqMapIter sync.Once
func Get_eqMapIter() gopurs_runtime.Value {
	once_eqMapIter.Do(func() {
		eqMapIter = gopurs_runtime.Func2(func(dictEq_0_box gopurs_runtime.Value, dictEq1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_eqMapIter(dictEq_0_box, dictEq1_1_box)
})
	})
	return eqMapIter
}

var ordMapIter gopurs_runtime.Value
var once_ordMapIter sync.Once
func Get_ordMapIter() gopurs_runtime.Value {
	once_ordMapIter.Do(func() {
		ordMapIter = gopurs_runtime.Func(func(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
eqMapIter1_1_0 := gopurs_runtime.Apply(Get_eqMapIter(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd_0, "Eq0"), gopurs_runtime.Value{}))
_ = eqMapIter1_1_0
return gopurs_runtime.Func(func(dictOrd1_2 gopurs_runtime.Value) gopurs_runtime.Value {
eqMapIter2_3_1 := gopurs_runtime.Apply(eqMapIter1_1_0, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd1_2, "Eq0"), gopurs_runtime.Value{}))
_ = eqMapIter2_3_1
var go__4_2 gopurs_runtime.Value
go__4_2 = gopurs_runtime.Func(func(a_5_loop gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_6_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
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
v3_9_7 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictOrd_0, "compare"), (*Data_Data_Map_Internal_IterNext)(v1_8_4.UnsafePtr).V0, (*Data_Data_Map_Internal_IterNext)(v_7_3.UnsafePtr).V0)
_ = v3_9_7
var __t8 gopurs_runtime.Value
{
if (v3_9_7.Type == 9 && v3_9_7.IntVal == 1111389260) {
v4_10_9 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictOrd1_2, "compare"), (*Data_Data_Map_Internal_IterNext)(v1_8_4.UnsafePtr).V1, (*Data_Data_Map_Internal_IterNext)(v_7_3.UnsafePtr).V1)
_ = v4_10_9
var __t10 gopurs_runtime.Value
{
if (v4_10_9.Type == 9 && v4_10_9.IntVal == 1111389260) {
a_5_loop = (*Data_Data_Map_Internal_IterNext)(v1_8_4.UnsafePtr).V2
b_6_loop = (*Data_Data_Map_Internal_IterNext)(v_7_3.UnsafePtr).V2
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
__t6 = gopurs_runtime.Value{Type: 9, IntVal: 2098047435, UnsafePtr: unsafe.Pointer(&pkg_Data_Ordering.Data_Data_Ordering_GT{})}
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
__t11 = gopurs_runtime.Value{Type: 9, IntVal: 1111389260, UnsafePtr: unsafe.Pointer(&pkg_Data_Ordering.Data_Data_Ordering_EQ{})}
goto end_branch_11
} else {

}
}
{
__t11 = gopurs_runtime.Value{Type: 9, IntVal: 3866105248, UnsafePtr: unsafe.Pointer(&pkg_Data_Ordering.Data_Data_Ordering_LT{})}
}
end_branch_11:
__t5 = __t11
goto end_branch_5
} else {

}
}
{
if (v_7_3.Type == 9 && v_7_3.IntVal == 4236111124) {
__t5 = gopurs_runtime.Value{Type: 9, IntVal: 2098047435, UnsafePtr: unsafe.Pointer(&pkg_Data_Ordering.Data_Data_Ordering_GT{})}
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
return gopurs_runtime.RecordDict2("compare", "Eq0", go__4_2, gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return eqMapIter2_3_1
}))
})
}()
})
	})
	return ordMapIter
}

var stepUnfoldr gopurs_runtime.Value
var once_stepUnfoldr sync.Once
func Get_stepUnfoldr() gopurs_runtime.Value {
	once_stepUnfoldr.Do(func() {
		stepUnfoldr = Call_stepWith(Get_iterMapL(), gopurs_runtime.Func3(func(k_0 gopurs_runtime.Value, v_1 gopurs_runtime.Value, next_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1354639136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Just{gopurs_runtime.Value{Type: 9, IntVal: 1102100576, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Data_Data_Tuple_Tuple{gopurs_runtime.Value{Type: 9, IntVal: 1102100576, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Data_Data_Tuple_Tuple{k_0, v_1})}, next_2})}})}
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 42808261, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Nothing{})}
}))
	})
	return stepUnfoldr
}

var toUnfoldable gopurs_runtime.Value
var once_toUnfoldable sync.Once
func Get_toUnfoldable() gopurs_runtime.Value {
	once_toUnfoldable.Do(func() {
		toUnfoldable = gopurs_runtime.Func(func(dictUnfoldable_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictUnfoldable_0 gopurs_runtime.Value = dictUnfoldable_0_loop
_ = dictUnfoldable_0
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictUnfoldable_0, "unfoldr"), Get_stepUnfoldr())
_ = __local_var_1_0
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_1_0, gopurs_runtime.Value{Type: 9, IntVal: 2861335956, UnsafePtr: unsafe.Pointer(&Data_Data_Map_Internal_IterNode{x_2, gopurs_runtime.Value{Type: 9, IntVal: 2509360378, UnsafePtr: unsafe.Pointer(&Data_Data_Map_Internal_IterLeaf{})}})})
})
}()
})
	})
	return toUnfoldable
}

var toUnfoldable1 gopurs_runtime.Value
var once_toUnfoldable1 sync.Once
func Get_toUnfoldable1() gopurs_runtime.Value {
	once_toUnfoldable1.Do(func() {
		toUnfoldable1 = func() gopurs_runtime.Value {
__local_var_0_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Unfoldable.Get_unfoldableArray(), "unfoldr"), Get_stepUnfoldr())
_ = __local_var_0_0
return gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_0_0, gopurs_runtime.Value{Type: 9, IntVal: 2861335956, UnsafePtr: unsafe.Pointer(&Data_Data_Map_Internal_IterNode{x_1, gopurs_runtime.Value{Type: 9, IntVal: 2509360378, UnsafePtr: unsafe.Pointer(&Data_Data_Map_Internal_IterLeaf{})}})})
})
}()
	})
	return toUnfoldable1
}

var showMap gopurs_runtime.Value
var once_showMap sync.Once
func Get_showMap() gopurs_runtime.Value {
	once_showMap.Do(func() {
		showMap = gopurs_runtime.Func2(func(dictShow_0_box gopurs_runtime.Value, dictShow1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_showMap(dictShow_0_box, dictShow1_1_box)
})
	})
	return showMap
}

var isSubmap gopurs_runtime.Value
var once_isSubmap sync.Once
func Get_isSubmap() gopurs_runtime.Value {
	once_isSubmap.Do(func() {
		isSubmap = gopurs_runtime.Func2(func(dictOrd_0_box gopurs_runtime.Value, dictEq_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_isSubmap(dictOrd_0_box, dictEq_1_box)
})
	})
	return isSubmap
}

var isEmpty gopurs_runtime.Value
var once_isEmpty sync.Once
func Get_isEmpty() gopurs_runtime.Value {
	once_isEmpty.Do(func() {
		isEmpty = gopurs_runtime.Func(func(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
return gopurs_runtime.Bool((v_0.Type == 9 && v_0.IntVal == 687041424))
}()
})
	})
	return isEmpty
}

var intersectionWith gopurs_runtime.Value
var once_intersectionWith sync.Once
func Get_intersectionWith() gopurs_runtime.Value {
	once_intersectionWith.Do(func() {
		intersectionWith = gopurs_runtime.Func(func(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
compare_1_0 := gopurs_runtime.RecordGet(dictOrd_0, "compare")
_ = compare_1_0
return gopurs_runtime.Func3(func(app_2 gopurs_runtime.Value, m1_3 gopurs_runtime.Value, m2_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.UncurriedApp4(Get_unsafeIntersectionWith(), compare_1_0, app_2, m1_3, m2_4)
})
}()
})
	})
	return intersectionWith
}

var intersection gopurs_runtime.Value
var once_intersection sync.Once
func Get_intersection() gopurs_runtime.Value {
	once_intersection.Do(func() {
		intersection = gopurs_runtime.Func(func(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
compare_1_0 := gopurs_runtime.RecordGet(dictOrd_0, "compare")
_ = compare_1_0
return gopurs_runtime.Func2(func(m1_2 gopurs_runtime.Value, m2_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.UncurriedApp4(Get_unsafeIntersectionWith(), compare_1_0, pkg_Data_Function.Get_const_(), m1_2, m2_3)
})
}()
})
	})
	return intersection
}

var insertWith gopurs_runtime.Value
var once_insertWith sync.Once
func Get_insertWith() gopurs_runtime.Value {
	once_insertWith.Do(func() {
		insertWith = gopurs_runtime.Func4(func(dictOrd_0_box gopurs_runtime.Value, app_1_box gopurs_runtime.Value, k_2_box gopurs_runtime.Value, v_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_insertWith(dictOrd_0_box, app_1_box, k_2_box, v_3_box)
})
	})
	return insertWith
}

var insert gopurs_runtime.Value
var once_insert sync.Once
func Get_insert() gopurs_runtime.Value {
	once_insert.Do(func() {
		insert = gopurs_runtime.Func3(func(dictOrd_0_box gopurs_runtime.Value, k_1_box gopurs_runtime.Value, v_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_insert(dictOrd_0_box, k_1_box, v_2_box)
})
	})
	return insert
}

var functorMap gopurs_runtime.Value
var once_functorMap sync.Once
func Get_functorMap() gopurs_runtime.Value {
	once_functorMap.Do(func() {
		functorMap = gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
var go__1_0 gopurs_runtime.Value
_ = go__1_0
go__1_0 = gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if (v_2.Type == 9 && v_2.IntVal == 687041424) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 687041424, UnsafePtr: unsafe.Pointer(&Data_Data_Map_Internal_Leaf{})}
goto end_branch_1
} else {

}
}
{
if (v_2.Type == 9 && v_2.IntVal == 324739070) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(&Data_Data_Map_Internal_Node{(*Data_Data_Map_Internal_Node)(v_2.UnsafePtr).V0, (*Data_Data_Map_Internal_Node)(v_2.UnsafePtr).V1, (*Data_Data_Map_Internal_Node)(v_2.UnsafePtr).V2, gopurs_runtime.Apply(f_0, (*Data_Data_Map_Internal_Node)(v_2.UnsafePtr).V3), gopurs_runtime.Apply(go__1_0, (*Data_Data_Map_Internal_Node)(v_2.UnsafePtr).V4), gopurs_runtime.Apply(go__1_0, (*Data_Data_Map_Internal_Node)(v_2.UnsafePtr).V5)})}
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
	return functorMap
}

var functorWithIndexMap gopurs_runtime.Value
var once_functorWithIndexMap sync.Once
func Get_functorWithIndexMap() gopurs_runtime.Value {
	once_functorWithIndexMap.Do(func() {
		functorWithIndexMap = gopurs_runtime.RecordDict2("mapWithIndex", "Functor0", gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
var go__1_0 gopurs_runtime.Value
_ = go__1_0
go__1_0 = gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if (v_2.Type == 9 && v_2.IntVal == 687041424) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 687041424, UnsafePtr: unsafe.Pointer(&Data_Data_Map_Internal_Leaf{})}
goto end_branch_1
} else {

}
}
{
if (v_2.Type == 9 && v_2.IntVal == 324739070) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(&Data_Data_Map_Internal_Node{(*Data_Data_Map_Internal_Node)(v_2.UnsafePtr).V0, (*Data_Data_Map_Internal_Node)(v_2.UnsafePtr).V1, (*Data_Data_Map_Internal_Node)(v_2.UnsafePtr).V2, gopurs_runtime.Apply2(f_0, (*Data_Data_Map_Internal_Node)(v_2.UnsafePtr).V2, (*Data_Data_Map_Internal_Node)(v_2.UnsafePtr).V3), gopurs_runtime.Apply(go__1_0, (*Data_Data_Map_Internal_Node)(v_2.UnsafePtr).V4), gopurs_runtime.Apply(go__1_0, (*Data_Data_Map_Internal_Node)(v_2.UnsafePtr).V5)})}
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
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_functorMap()
}))
	})
	return functorWithIndexMap
}

var foldableMap gopurs_runtime.Value
var once_foldableMap sync.Once
func Get_foldableMap() gopurs_runtime.Value {
	once_foldableMap.Do(func() {
		foldableMap = gopurs_runtime.RecordDict3("foldr", "foldl", "foldMap", gopurs_runtime.Func2(func(f_0 gopurs_runtime.Value, z_1 gopurs_runtime.Value) gopurs_runtime.Value {
var go__2_0 gopurs_runtime.Value
_ = go__2_0
go__2_0 = gopurs_runtime.Func2(Call_go__2_0)
return gopurs_runtime.Func(func(m_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.UncurriedApp2(go__2_0, m_3, z_1)
})
}), gopurs_runtime.Func2(func(f_0 gopurs_runtime.Value, z_1 gopurs_runtime.Value) gopurs_runtime.Value {
var go__2_2 gopurs_runtime.Value
_ = go__2_2
go__2_2 = gopurs_runtime.Func2(Call_go__2_2)
return gopurs_runtime.Func(func(m_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.UncurriedApp2(go__2_2, z_1, m_3)
})
}), gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
mempty_1_4 := gopurs_runtime.RecordGet(dictMonoid_0, "mempty")
_ = mempty_1_4
__local_var_2_5 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{})
_ = __local_var_2_5
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
var go__4_6 gopurs_runtime.Value
_ = go__4_6
go__4_6 = gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
var __t7 gopurs_runtime.Value
{
if (v_5.Type == 9 && v_5.IntVal == 687041424) {
__t7 = mempty_1_4
goto end_branch_7
} else {

}
}
{
if (v_5.Type == 9 && v_5.IntVal == 324739070) {
__t7 = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_5, "append"), gopurs_runtime.Apply(go__4_6, (*Data_Data_Map_Internal_Node)(v_5.UnsafePtr).V4), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_5, "append"), gopurs_runtime.Apply(f_3, (*Data_Data_Map_Internal_Node)(v_5.UnsafePtr).V3), gopurs_runtime.Apply(go__4_6, (*Data_Data_Map_Internal_Node)(v_5.UnsafePtr).V5)))
goto end_branch_7
} else {

}
}
{
__t7 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_7:
return __t7
})
return go__4_6
})
}))
	})
	return foldableMap
}

var foldableWithIndexMap gopurs_runtime.Value
var once_foldableWithIndexMap sync.Once
func Get_foldableWithIndexMap() gopurs_runtime.Value {
	once_foldableWithIndexMap.Do(func() {
		foldableWithIndexMap = gopurs_runtime.RecordDict4("foldrWithIndex", "foldlWithIndex", "foldMapWithIndex", "Foldable0", gopurs_runtime.Func2(func(f_0 gopurs_runtime.Value, z_1 gopurs_runtime.Value) gopurs_runtime.Value {
var go__2_0 gopurs_runtime.Value
_ = go__2_0
go__2_0 = gopurs_runtime.Func2(Call_go__2_0)
return gopurs_runtime.Func(func(m_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.UncurriedApp2(go__2_0, m_3, z_1)
})
}), gopurs_runtime.Func2(func(f_0 gopurs_runtime.Value, z_1 gopurs_runtime.Value) gopurs_runtime.Value {
var go__2_2 gopurs_runtime.Value
_ = go__2_2
go__2_2 = gopurs_runtime.Func2(Call_go__2_2)
return gopurs_runtime.Func(func(m_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.UncurriedApp2(go__2_2, z_1, m_3)
})
}), gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
mempty_1_4 := gopurs_runtime.RecordGet(dictMonoid_0, "mempty")
_ = mempty_1_4
__local_var_2_5 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{})
_ = __local_var_2_5
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
var go__4_6 gopurs_runtime.Value
_ = go__4_6
go__4_6 = gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
var __t7 gopurs_runtime.Value
{
if (v_5.Type == 9 && v_5.IntVal == 687041424) {
__t7 = mempty_1_4
goto end_branch_7
} else {

}
}
{
if (v_5.Type == 9 && v_5.IntVal == 324739070) {
__t7 = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_5, "append"), gopurs_runtime.Apply(go__4_6, (*Data_Data_Map_Internal_Node)(v_5.UnsafePtr).V4), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_5, "append"), gopurs_runtime.Apply2(f_3, (*Data_Data_Map_Internal_Node)(v_5.UnsafePtr).V2, (*Data_Data_Map_Internal_Node)(v_5.UnsafePtr).V3), gopurs_runtime.Apply(go__4_6, (*Data_Data_Map_Internal_Node)(v_5.UnsafePtr).V5)))
goto end_branch_7
} else {

}
}
{
__t7 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_7:
return __t7
})
return go__4_6
})
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_foldableMap()
}))
	})
	return foldableWithIndexMap
}

var keys gopurs_runtime.Value
var once_keys sync.Once
func Get_keys() gopurs_runtime.Value {
	once_keys.Do(func() {
		keys = func() gopurs_runtime.Value {
var go__0_0 gopurs_runtime.Value
_ = go__0_0
go__0_0 = gopurs_runtime.Func2(Call_go__0_0)
return gopurs_runtime.Func(func(m_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.UncurriedApp2(go__0_0, m_1, gopurs_runtime.Value{Type: 9, IntVal: 1536536851, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Data_Data_List_Types_Nil{})})
})
}()
	})
	return keys
}

var traversableMap gopurs_runtime.Value
var once_traversableMap sync.Once
func Get_traversableMap() gopurs_runtime.Value {
	once_traversableMap.Do(func() {
		traversableMap = gopurs_runtime.RecordDict4("traverse", "sequence", "Functor0", "Foldable1", gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
Apply0_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{})
_ = Apply0_1_0
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
var go__3_1 gopurs_runtime.Value
_ = go__3_1
go__3_1 = gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 gopurs_runtime.Value
{
if (v_4.Type == 9 && v_4.IntVal == 687041424) {
__t2 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "pure"), gopurs_runtime.Value{Type: 9, IntVal: 687041424, UnsafePtr: unsafe.Pointer(&Data_Data_Map_Internal_Leaf{})})
goto end_branch_2
} else {

}
}
{
if (v_4.Type == 9 && v_4.IntVal == 324739070) {
__local_var_5_3 := (*Data_Data_Map_Internal_Node)(v_4.UnsafePtr).V0
_ = __local_var_5_3
__local_var_6_4 := (*Data_Data_Map_Internal_Node)(v_4.UnsafePtr).V2
_ = __local_var_6_4
__local_var_7_5 := (*Data_Data_Map_Internal_Node)(v_4.UnsafePtr).V1
_ = __local_var_7_5
__t2 = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Apply0_1_0, "apply"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Apply0_1_0, "apply"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Apply0_1_0, "Functor0"), gopurs_runtime.Value{}), "map"), gopurs_runtime.Func3(func(l_prime_8 gopurs_runtime.Value, v_prime_9 gopurs_runtime.Value, r_prime_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(&Data_Data_Map_Internal_Node{__local_var_5_3, __local_var_7_5, __local_var_6_4, v_prime_9, l_prime_8, r_prime_10})}
}), gopurs_runtime.Apply(go__3_1, (*Data_Data_Map_Internal_Node)(v_4.UnsafePtr).V4)), gopurs_runtime.Apply(f_2, (*Data_Data_Map_Internal_Node)(v_4.UnsafePtr).V3)), gopurs_runtime.Apply(go__3_1, (*Data_Data_Map_Internal_Node)(v_4.UnsafePtr).V5))
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
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_traversableMap(), "traverse"), dictApplicative_0, gopurs_runtime.RecordGet(pkg_Control_Category.Get_categoryFn(), "identity"))
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_functorMap()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_foldableMap()
}))
	})
	return traversableMap
}

var traversableWithIndexMap gopurs_runtime.Value
var once_traversableWithIndexMap sync.Once
func Get_traversableWithIndexMap() gopurs_runtime.Value {
	once_traversableWithIndexMap.Do(func() {
		traversableWithIndexMap = gopurs_runtime.RecordDict4("traverseWithIndex", "FunctorWithIndex0", "FoldableWithIndex1", "Traversable2", gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
Apply0_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{})
_ = Apply0_1_0
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
var go__3_1 gopurs_runtime.Value
_ = go__3_1
go__3_1 = gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 gopurs_runtime.Value
{
if (v_4.Type == 9 && v_4.IntVal == 687041424) {
__t2 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "pure"), gopurs_runtime.Value{Type: 9, IntVal: 687041424, UnsafePtr: unsafe.Pointer(&Data_Data_Map_Internal_Leaf{})})
goto end_branch_2
} else {

}
}
{
if (v_4.Type == 9 && v_4.IntVal == 324739070) {
__local_var_5_3 := (*Data_Data_Map_Internal_Node)(v_4.UnsafePtr).V0
_ = __local_var_5_3
__local_var_6_4 := (*Data_Data_Map_Internal_Node)(v_4.UnsafePtr).V2
_ = __local_var_6_4
__local_var_7_5 := (*Data_Data_Map_Internal_Node)(v_4.UnsafePtr).V1
_ = __local_var_7_5
__t2 = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Apply0_1_0, "apply"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Apply0_1_0, "apply"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Apply0_1_0, "Functor0"), gopurs_runtime.Value{}), "map"), gopurs_runtime.Func3(func(l_prime_8 gopurs_runtime.Value, v_prime_9 gopurs_runtime.Value, r_prime_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(&Data_Data_Map_Internal_Node{__local_var_5_3, __local_var_7_5, __local_var_6_4, v_prime_9, l_prime_8, r_prime_10})}
}), gopurs_runtime.Apply(go__3_1, (*Data_Data_Map_Internal_Node)(v_4.UnsafePtr).V4)), gopurs_runtime.Apply2(f_2, __local_var_6_4, (*Data_Data_Map_Internal_Node)(v_4.UnsafePtr).V3)), gopurs_runtime.Apply(go__3_1, (*Data_Data_Map_Internal_Node)(v_4.UnsafePtr).V5))
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
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_functorWithIndexMap()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_foldableWithIndexMap()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_traversableMap()
}))
	})
	return traversableWithIndexMap
}

var values gopurs_runtime.Value
var once_values sync.Once
func Get_values() gopurs_runtime.Value {
	once_values.Do(func() {
		values = func() gopurs_runtime.Value {
var go__0_0 gopurs_runtime.Value
_ = go__0_0
go__0_0 = gopurs_runtime.Func2(Call_go__0_0)
return gopurs_runtime.Func(func(m_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.UncurriedApp2(go__0_0, m_1, gopurs_runtime.Value{Type: 9, IntVal: 1536536851, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Data_Data_List_Types_Nil{})})
})
}()
	})
	return values
}

var foldSubmapBy gopurs_runtime.Value
var once_foldSubmapBy sync.Once
func Get_foldSubmapBy() gopurs_runtime.Value {
	once_foldSubmapBy.Do(func() {
		foldSubmapBy = gopurs_runtime.Func6(func(dictOrd_0_box gopurs_runtime.Value, appendFn_1_box gopurs_runtime.Value, memptyValue_2_box gopurs_runtime.Value, kmin_3_box gopurs_runtime.Value, kmax_4_box gopurs_runtime.Value, f_5_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldSubmapBy(dictOrd_0_box, appendFn_1_box, memptyValue_2_box, kmin_3_box, kmax_4_box, f_5_box)
})
	})
	return foldSubmapBy
}

var foldSubmap gopurs_runtime.Value
var once_foldSubmap sync.Once
func Get_foldSubmap() gopurs_runtime.Value {
	once_foldSubmap.Do(func() {
		foldSubmap = gopurs_runtime.Func2(func(dictOrd_0_box gopurs_runtime.Value, dictMonoid_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldSubmap(dictOrd_0_box, dictMonoid_1_box)
})
	})
	return foldSubmap
}

var findMin gopurs_runtime.Value
var once_findMin sync.Once
func Get_findMin() gopurs_runtime.Value {
	once_findMin.Do(func() {
		findMin = gopurs_runtime.Func(func(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
findMin:
for {
if false { continue findMin }
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var __t0 gopurs_runtime.Value
{
if (v_0.Type == 9 && v_0.IntVal == 687041424) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 42808261, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Nothing{})}
goto end_branch_0
} else {

}
}
{
if (v_0.Type == 9 && v_0.IntVal == 324739070) {
var __t1 gopurs_runtime.Value
{
if ((*Data_Data_Map_Internal_Node)(v_0.UnsafePtr).V4.Type == 9 && (*Data_Data_Map_Internal_Node)(v_0.UnsafePtr).V4.IntVal == 687041424) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 1354639136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Just{gopurs_runtime.RecordDict2("key", "value", (*Data_Data_Map_Internal_Node)(v_0.UnsafePtr).V2, (*Data_Data_Map_Internal_Node)(v_0.UnsafePtr).V3)})}
goto end_branch_1
} else {

}
}
{
v_0_loop = (*Data_Data_Map_Internal_Node)(v_0.UnsafePtr).V4
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
}()
})
	})
	return findMin
}

var lookupGT gopurs_runtime.Value
var once_lookupGT sync.Once
func Get_lookupGT() gopurs_runtime.Value {
	once_lookupGT.Do(func() {
		lookupGT = gopurs_runtime.Func2(func(dictOrd_0_box gopurs_runtime.Value, k_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_lookupGT(dictOrd_0_box, k_1_box)
})
	})
	return lookupGT
}

var findMax gopurs_runtime.Value
var once_findMax sync.Once
func Get_findMax() gopurs_runtime.Value {
	once_findMax.Do(func() {
		findMax = gopurs_runtime.Func(func(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
findMax:
for {
if false { continue findMax }
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var __t0 gopurs_runtime.Value
{
if (v_0.Type == 9 && v_0.IntVal == 687041424) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 42808261, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Nothing{})}
goto end_branch_0
} else {

}
}
{
if (v_0.Type == 9 && v_0.IntVal == 324739070) {
var __t1 gopurs_runtime.Value
{
if ((*Data_Data_Map_Internal_Node)(v_0.UnsafePtr).V5.Type == 9 && (*Data_Data_Map_Internal_Node)(v_0.UnsafePtr).V5.IntVal == 687041424) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 1354639136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Just{gopurs_runtime.RecordDict2("key", "value", (*Data_Data_Map_Internal_Node)(v_0.UnsafePtr).V2, (*Data_Data_Map_Internal_Node)(v_0.UnsafePtr).V3)})}
goto end_branch_1
} else {

}
}
{
v_0_loop = (*Data_Data_Map_Internal_Node)(v_0.UnsafePtr).V5
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
}()
})
	})
	return findMax
}

var lookupLT gopurs_runtime.Value
var once_lookupLT sync.Once
func Get_lookupLT() gopurs_runtime.Value {
	once_lookupLT.Do(func() {
		lookupLT = gopurs_runtime.Func2(func(dictOrd_0_box gopurs_runtime.Value, k_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_lookupLT(dictOrd_0_box, k_1_box)
})
	})
	return lookupLT
}

var filterWithKey gopurs_runtime.Value
var once_filterWithKey sync.Once
func Get_filterWithKey() gopurs_runtime.Value {
	once_filterWithKey.Do(func() {
		filterWithKey = gopurs_runtime.Func2(func(dictOrd_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_filterWithKey(dictOrd_0_box, f_1_box)
})
	})
	return filterWithKey
}

var filterKeys gopurs_runtime.Value
var once_filterKeys sync.Once
func Get_filterKeys() gopurs_runtime.Value {
	once_filterKeys.Do(func() {
		filterKeys = gopurs_runtime.Func2(func(dictOrd_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_filterKeys(dictOrd_0_box, f_1_box)
})
	})
	return filterKeys
}

var filter gopurs_runtime.Value
var once_filter sync.Once
func Get_filter() gopurs_runtime.Value {
	once_filter.Do(func() {
		filter = gopurs_runtime.Func2(func(dictOrd_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_filter(dictOrd_0_box, x_1_box)
})
	})
	return filter
}

var eqMap gopurs_runtime.Value
var once_eqMap sync.Once
func Get_eqMap() gopurs_runtime.Value {
	once_eqMap.Do(func() {
		eqMap = gopurs_runtime.Func2(func(dictEq_0_box gopurs_runtime.Value, dictEq1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_eqMap(dictEq_0_box, dictEq1_1_box)
})
	})
	return eqMap
}

var ordMap gopurs_runtime.Value
var once_ordMap sync.Once
func Get_ordMap() gopurs_runtime.Value {
	once_ordMap.Do(func() {
		ordMap = gopurs_runtime.Func(func(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
ordMapIter1_1_0 := gopurs_runtime.Apply(Get_ordMapIter(), dictOrd_0)
_ = ordMapIter1_1_0
eqMap1_2_1 := gopurs_runtime.Apply(Get_eqMap(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd_0, "Eq0"), gopurs_runtime.Value{}))
_ = eqMap1_2_1
return gopurs_runtime.Func(func(dictOrd1_3 gopurs_runtime.Value) gopurs_runtime.Value {
eqMap2_4_2 := gopurs_runtime.Apply(eqMap1_2_1, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd1_3, "Eq0"), gopurs_runtime.Value{}))
_ = eqMap2_4_2
return gopurs_runtime.RecordDict2("compare", "Eq0", gopurs_runtime.Func2(func(xs_5 gopurs_runtime.Value, ys_6 gopurs_runtime.Value) gopurs_runtime.Value {
var __t3 gopurs_runtime.Value
{
if (xs_5.Type == 9 && xs_5.IntVal == 687041424) {
var __t4 gopurs_runtime.Value
{
if (ys_6.Type == 9 && ys_6.IntVal == 687041424) {
__t4 = gopurs_runtime.Value{Type: 9, IntVal: 1111389260, UnsafePtr: unsafe.Pointer(&pkg_Data_Ordering.Data_Data_Ordering_EQ{})}
goto end_branch_4
} else {

}
}
{
__t4 = gopurs_runtime.Value{Type: 9, IntVal: 3866105248, UnsafePtr: unsafe.Pointer(&pkg_Data_Ordering.Data_Data_Ordering_LT{})}
}
end_branch_4:
__t3 = __t4
goto end_branch_3
} else {

}
}
{
if (ys_6.Type == 9 && ys_6.IntVal == 687041424) {
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 2098047435, UnsafePtr: unsafe.Pointer(&pkg_Data_Ordering.Data_Data_Ordering_GT{})}
goto end_branch_3
} else {

}
}
{
__t3 = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(ordMapIter1_1_0, dictOrd1_3), "compare"), gopurs_runtime.Value{Type: 9, IntVal: 2861335956, UnsafePtr: unsafe.Pointer(&Data_Data_Map_Internal_IterNode{xs_5, gopurs_runtime.Value{Type: 9, IntVal: 2509360378, UnsafePtr: unsafe.Pointer(&Data_Data_Map_Internal_IterLeaf{})}})}, gopurs_runtime.Value{Type: 9, IntVal: 2861335956, UnsafePtr: unsafe.Pointer(&Data_Data_Map_Internal_IterNode{ys_6, gopurs_runtime.Value{Type: 9, IntVal: 2509360378, UnsafePtr: unsafe.Pointer(&Data_Data_Map_Internal_IterLeaf{})}})})
}
end_branch_3:
return __t3
}), gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return eqMap2_4_2
}))
})
}()
})
	})
	return ordMap
}

var eq1Map gopurs_runtime.Value
var once_eq1Map sync.Once
func Get_eq1Map() gopurs_runtime.Value {
	once_eq1Map.Do(func() {
		eq1Map = gopurs_runtime.Func(func(dictEq_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictEq_0 gopurs_runtime.Value = dictEq_0_loop
_ = dictEq_0
return gopurs_runtime.RecordDict1("eq1", gopurs_runtime.Func(func(dictEq1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(Call_eqMap(dictEq_0, dictEq1_1), "eq")
}))
}()
})
	})
	return eq1Map
}

var ord1Map gopurs_runtime.Value
var once_ord1Map sync.Once
func Get_ord1Map() gopurs_runtime.Value {
	once_ord1Map.Do(func() {
		ord1Map = gopurs_runtime.Func(func(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
ordMap1_1_0 := gopurs_runtime.Apply(Get_ordMap(), dictOrd_0)
_ = ordMap1_1_0
__local_var_2_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd_0, "Eq0"), gopurs_runtime.Value{})
_ = __local_var_2_1
eq1Map1_3_2 := gopurs_runtime.RecordDict1("eq1", gopurs_runtime.Func(func(dictEq1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(Call_eqMap(__local_var_2_1, dictEq1_3), "eq")
}))
_ = eq1Map1_3_2
return gopurs_runtime.RecordDict2("compare1", "Eq10", gopurs_runtime.Func(func(dictOrd1_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(gopurs_runtime.Apply(ordMap1_1_0, dictOrd1_4), "compare")
}), gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return eq1Map1_3_2
}))
}()
})
	})
	return ord1Map
}

var empty gopurs_runtime.Value
var once_empty sync.Once
func Get_empty() gopurs_runtime.Value {
	once_empty.Do(func() {
		empty = gopurs_runtime.Value{Type: 9, IntVal: 687041424, UnsafePtr: unsafe.Pointer(&Data_Data_Map_Internal_Leaf{})}
	})
	return empty
}

var fromFoldable gopurs_runtime.Value
var once_fromFoldable sync.Once
func Get_fromFoldable() gopurs_runtime.Value {
	once_fromFoldable.Do(func() {
		fromFoldable = gopurs_runtime.Func2(func(dictOrd_0_box gopurs_runtime.Value, dictFoldable_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_fromFoldable(dictOrd_0_box, dictFoldable_1_box)
})
	})
	return fromFoldable
}

var fromFoldableWith gopurs_runtime.Value
var once_fromFoldableWith sync.Once
func Get_fromFoldableWith() gopurs_runtime.Value {
	once_fromFoldableWith.Do(func() {
		fromFoldableWith = gopurs_runtime.Func3(func(dictOrd_0_box gopurs_runtime.Value, dictFoldable_1_box gopurs_runtime.Value, f_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_fromFoldableWith(dictOrd_0_box, dictFoldable_1_box, f_2_box)
})
	})
	return fromFoldableWith
}

var fromFoldableWithIndex gopurs_runtime.Value
var once_fromFoldableWithIndex sync.Once
func Get_fromFoldableWithIndex() gopurs_runtime.Value {
	once_fromFoldableWithIndex.Do(func() {
		fromFoldableWithIndex = gopurs_runtime.Func2(func(dictOrd_0_box gopurs_runtime.Value, dictFoldableWithIndex_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_fromFoldableWithIndex(dictOrd_0_box, dictFoldableWithIndex_1_box)
})
	})
	return fromFoldableWithIndex
}

var monoidSemigroupMap gopurs_runtime.Value
var once_monoidSemigroupMap sync.Once
func Get_monoidSemigroupMap() gopurs_runtime.Value {
	once_monoidSemigroupMap.Do(func() {
		monoidSemigroupMap = gopurs_runtime.Func2(func(_dollar__unused_0_box gopurs_runtime.Value, dictOrd_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_monoidSemigroupMap(_dollar__unused_0_box, dictOrd_1_box)
})
	})
	return monoidSemigroupMap
}

var submap gopurs_runtime.Value
var once_submap sync.Once
func Get_submap() gopurs_runtime.Value {
	once_submap.Do(func() {
		submap = gopurs_runtime.Func(func(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
compare_1_0 := gopurs_runtime.RecordGet(dictOrd_0, "compare")
_ = compare_1_0
return gopurs_runtime.Func2(func(kmin_2 gopurs_runtime.Value, kmax_3 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldSubmapBy(dictOrd_0, gopurs_runtime.Func2(func(m1_4 gopurs_runtime.Value, m2_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.UncurriedApp4(Get_unsafeUnionWith(), compare_1_0, pkg_Data_Function.Get_const_(), m1_4, m2_5)
}), gopurs_runtime.Value{Type: 9, IntVal: 687041424, UnsafePtr: unsafe.Pointer(&Data_Data_Map_Internal_Leaf{})}, kmin_2, kmax_3, Get_singleton())
})
}()
})
	})
	return submap
}

var unions gopurs_runtime.Value
var once_unions sync.Once
func Get_unions() gopurs_runtime.Value {
	once_unions.Do(func() {
		unions = gopurs_runtime.Func(func(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
compare_1_0 := gopurs_runtime.RecordGet(dictOrd_0, "compare")
_ = compare_1_0
return gopurs_runtime.Func(func(dictFoldable_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFoldable_2, "foldl"), gopurs_runtime.Func2(func(m1_3 gopurs_runtime.Value, m2_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.UncurriedApp4(Get_unsafeUnionWith(), compare_1_0, pkg_Data_Function.Get_const_(), m1_3, m2_4)
}), gopurs_runtime.Value{Type: 9, IntVal: 687041424, UnsafePtr: unsafe.Pointer(&Data_Data_Map_Internal_Leaf{})})
})
}()
})
	})
	return unions
}

var difference gopurs_runtime.Value
var once_difference sync.Once
func Get_difference() gopurs_runtime.Value {
	once_difference.Do(func() {
		difference = gopurs_runtime.Func(func(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
compare_1_0 := gopurs_runtime.RecordGet(dictOrd_0, "compare")
_ = compare_1_0
return gopurs_runtime.Func2(func(m1_2 gopurs_runtime.Value, m2_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.UncurriedApp3(Get_unsafeDifference(), compare_1_0, m1_2, m2_3)
})
}()
})
	})
	return difference
}

var delete_ gopurs_runtime.Value
var once_delete_ sync.Once
func Get_delete_() gopurs_runtime.Value {
	once_delete_.Do(func() {
		delete_ = gopurs_runtime.Func2(func(dictOrd_0_box gopurs_runtime.Value, k_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_delete_(dictOrd_0_box, k_1_box)
})
	})
	return delete_
}

var checkValid gopurs_runtime.Value
var once_checkValid sync.Once
func Get_checkValid() gopurs_runtime.Value {
	once_checkValid.Do(func() {
		checkValid = gopurs_runtime.Func(func(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
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
if ((*Data_Data_Map_Internal_Node)(v_2.UnsafePtr).V4.Type == 9 && (*Data_Data_Map_Internal_Node)(v_2.UnsafePtr).V4.IntVal == 687041424) {
var __t3 gopurs_runtime.Value
{
if ((*Data_Data_Map_Internal_Node)(v_2.UnsafePtr).V5.Type == 9 && (*Data_Data_Map_Internal_Node)(v_2.UnsafePtr).V5.IntVal == 687041424) {
__t3 = gopurs_runtime.Bool(true)
goto end_branch_3
} else {

}
}
{
if ((*Data_Data_Map_Internal_Node)(v_2.UnsafePtr).V5.Type == 9 && (*Data_Data_Map_Internal_Node)(v_2.UnsafePtr).V5.IntVal == 324739070) {
__t3 = gopurs_runtime.Bool((*Data_Data_Map_Internal_Node)(v_2.UnsafePtr).V0.IntVal == 2 && (*Data_Data_Map_Internal_Node)((*Data_Data_Map_Internal_Node)(v_2.UnsafePtr).V5.UnsafePtr).V0.IntVal == 1 && (*Data_Data_Map_Internal_Node)(v_2.UnsafePtr).V1.IntVal > (*Data_Data_Map_Internal_Node)((*Data_Data_Map_Internal_Node)(v_2.UnsafePtr).V5.UnsafePtr).V1.IntVal && (gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictOrd_0, "compare"), (*Data_Data_Map_Internal_Node)((*Data_Data_Map_Internal_Node)(v_2.UnsafePtr).V5.UnsafePtr).V2, (*Data_Data_Map_Internal_Node)(v_2.UnsafePtr).V2).Type == 9 && gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictOrd_0, "compare"), (*Data_Data_Map_Internal_Node)((*Data_Data_Map_Internal_Node)(v_2.UnsafePtr).V5.UnsafePtr).V2, (*Data_Data_Map_Internal_Node)(v_2.UnsafePtr).V2).IntVal == 2098047435) && gopurs_runtime.Apply(go__1_0, (*Data_Data_Map_Internal_Node)(v_2.UnsafePtr).V5).IntVal != 0)
goto end_branch_3
} else {

}
}
{
__t3 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_3:
__t2 = __t3
goto end_branch_2
} else {

}
}
{
if ((*Data_Data_Map_Internal_Node)(v_2.UnsafePtr).V4.Type == 9 && (*Data_Data_Map_Internal_Node)(v_2.UnsafePtr).V4.IntVal == 324739070) {
var __t4 gopurs_runtime.Value
{
if ((*Data_Data_Map_Internal_Node)(v_2.UnsafePtr).V5.Type == 9 && (*Data_Data_Map_Internal_Node)(v_2.UnsafePtr).V5.IntVal == 687041424) {
__t4 = gopurs_runtime.Bool((*Data_Data_Map_Internal_Node)(v_2.UnsafePtr).V0.IntVal == 2 && (*Data_Data_Map_Internal_Node)((*Data_Data_Map_Internal_Node)(v_2.UnsafePtr).V4.UnsafePtr).V0.IntVal == 1 && (*Data_Data_Map_Internal_Node)(v_2.UnsafePtr).V1.IntVal > (*Data_Data_Map_Internal_Node)((*Data_Data_Map_Internal_Node)(v_2.UnsafePtr).V4.UnsafePtr).V1.IntVal && (gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictOrd_0, "compare"), (*Data_Data_Map_Internal_Node)((*Data_Data_Map_Internal_Node)(v_2.UnsafePtr).V4.UnsafePtr).V2, (*Data_Data_Map_Internal_Node)(v_2.UnsafePtr).V2).Type == 9 && gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictOrd_0, "compare"), (*Data_Data_Map_Internal_Node)((*Data_Data_Map_Internal_Node)(v_2.UnsafePtr).V4.UnsafePtr).V2, (*Data_Data_Map_Internal_Node)(v_2.UnsafePtr).V2).IntVal == 3866105248) && gopurs_runtime.Apply(go__1_0, (*Data_Data_Map_Internal_Node)(v_2.UnsafePtr).V4).IntVal != 0)
goto end_branch_4
} else {

}
}
{
if ((*Data_Data_Map_Internal_Node)(v_2.UnsafePtr).V5.Type == 9 && (*Data_Data_Map_Internal_Node)(v_2.UnsafePtr).V5.IntVal == 324739070) {
__local_var_3_5 := (*Data_Data_Map_Internal_Node)((*Data_Data_Map_Internal_Node)(v_2.UnsafePtr).V5.UnsafePtr).V0.IntVal - (*Data_Data_Map_Internal_Node)((*Data_Data_Map_Internal_Node)(v_2.UnsafePtr).V4.UnsafePtr).V0.IntVal
_ = __local_var_3_5
var __t6 gopurs_runtime.Value
{
if __local_var_3_5 >= 0 {
__t6 = gopurs_runtime.Bool(__local_var_3_5 < 2)
goto end_branch_6
} else {

}
}
{
__t6 = gopurs_runtime.Bool(0 - __local_var_3_5 < 2)
}
end_branch_6:
__t4 = gopurs_runtime.Bool((*Data_Data_Map_Internal_Node)(v_2.UnsafePtr).V0.IntVal > (*Data_Data_Map_Internal_Node)((*Data_Data_Map_Internal_Node)(v_2.UnsafePtr).V5.UnsafePtr).V0.IntVal && (gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictOrd_0, "compare"), (*Data_Data_Map_Internal_Node)((*Data_Data_Map_Internal_Node)(v_2.UnsafePtr).V5.UnsafePtr).V2, (*Data_Data_Map_Internal_Node)(v_2.UnsafePtr).V2).Type == 9 && gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictOrd_0, "compare"), (*Data_Data_Map_Internal_Node)((*Data_Data_Map_Internal_Node)(v_2.UnsafePtr).V5.UnsafePtr).V2, (*Data_Data_Map_Internal_Node)(v_2.UnsafePtr).V2).IntVal == 2098047435) && (*Data_Data_Map_Internal_Node)(v_2.UnsafePtr).V0.IntVal > (*Data_Data_Map_Internal_Node)((*Data_Data_Map_Internal_Node)(v_2.UnsafePtr).V4.UnsafePtr).V0.IntVal && (gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictOrd_0, "compare"), (*Data_Data_Map_Internal_Node)((*Data_Data_Map_Internal_Node)(v_2.UnsafePtr).V4.UnsafePtr).V2, (*Data_Data_Map_Internal_Node)(v_2.UnsafePtr).V2).Type == 9 && gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictOrd_0, "compare"), (*Data_Data_Map_Internal_Node)((*Data_Data_Map_Internal_Node)(v_2.UnsafePtr).V4.UnsafePtr).V2, (*Data_Data_Map_Internal_Node)(v_2.UnsafePtr).V2).IntVal == 3866105248) && __t6.IntVal != 0 && (*Data_Data_Map_Internal_Node)((*Data_Data_Map_Internal_Node)(v_2.UnsafePtr).V5.UnsafePtr).V1.IntVal + (*Data_Data_Map_Internal_Node)((*Data_Data_Map_Internal_Node)(v_2.UnsafePtr).V4.UnsafePtr).V1.IntVal + 1 == (*Data_Data_Map_Internal_Node)(v_2.UnsafePtr).V1.IntVal && gopurs_runtime.Apply(go__1_0, (*Data_Data_Map_Internal_Node)(v_2.UnsafePtr).V4).IntVal != 0 && gopurs_runtime.Apply(go__1_0, (*Data_Data_Map_Internal_Node)(v_2.UnsafePtr).V5).IntVal != 0)
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
}()
})
	})
	return checkValid
}

var catMaybes gopurs_runtime.Value
var once_catMaybes sync.Once
func Get_catMaybes() gopurs_runtime.Value {
	once_catMaybes.Do(func() {
		catMaybes = gopurs_runtime.Func(func(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
return Call_mapMaybeWithKey(dictOrd_0, gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(pkg_Control_Category.Get_categoryFn(), "identity")
}))
}()
})
	})
	return catMaybes
}

var applyMap gopurs_runtime.Value
var once_applyMap sync.Once
func Get_applyMap() gopurs_runtime.Value {
	once_applyMap.Do(func() {
		applyMap = gopurs_runtime.Func(func(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
compare_1_0 := gopurs_runtime.RecordGet(dictOrd_0, "compare")
_ = compare_1_0
return gopurs_runtime.RecordDict2("apply", "Functor0", gopurs_runtime.Func2(func(m1_2 gopurs_runtime.Value, m2_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.UncurriedApp4(Get_unsafeIntersectionWith(), compare_1_0, gopurs_runtime.RecordGet(pkg_Control_Category.Get_categoryFn(), "identity"), m1_2, m2_3)
}), gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_functorMap()
}))
}()
})
	})
	return applyMap
}

var bindMap gopurs_runtime.Value
var once_bindMap sync.Once
func Get_bindMap() gopurs_runtime.Value {
	once_bindMap.Do(func() {
		bindMap = gopurs_runtime.Func(func(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
compare_1_1 := gopurs_runtime.RecordGet(dictOrd_0, "compare")
_ = compare_1_1
applyMap1_1_0 := gopurs_runtime.RecordDict2("apply", "Functor0", gopurs_runtime.Func2(func(m1_2 gopurs_runtime.Value, m2_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.UncurriedApp4(Get_unsafeIntersectionWith(), compare_1_1, gopurs_runtime.RecordGet(pkg_Control_Category.Get_categoryFn(), "identity"), m1_2, m2_3)
}), gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_functorMap()
}))
_ = applyMap1_1_0
return gopurs_runtime.RecordDict2("bind", "Apply0", gopurs_runtime.Func2(func(m_2 gopurs_runtime.Value, f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Call_mapMaybeWithKey(dictOrd_0, gopurs_runtime.Func(func(k_4 gopurs_runtime.Value) gopurs_runtime.Value {
var go__5_2 gopurs_runtime.Value
go__5_2 = gopurs_runtime.Func(func(v_6_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
go__5_2:
for {
if false { continue go__5_2 }
var v_6 gopurs_runtime.Value = v_6_loop
_ = v_6
var __t3 gopurs_runtime.Value
{
if (v_6.Type == 9 && v_6.IntVal == 687041424) {
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 42808261, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Nothing{})}
goto end_branch_3
} else {

}
}
{
if (v_6.Type == 9 && v_6.IntVal == 324739070) {
v1_7_4 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictOrd_0, "compare"), k_4, (*Data_Data_Map_Internal_Node)(v_6.UnsafePtr).V2)
_ = v1_7_4
var __t5 gopurs_runtime.Value
{
if (v1_7_4.Type == 9 && v1_7_4.IntVal == 3866105248) {
v_6_loop = (*Data_Data_Map_Internal_Node)(v_6.UnsafePtr).V4
continue go__5_2
__t5 = gopurs_runtime.Value{}
goto end_branch_5
} else {

}
}
{
if (v1_7_4.Type == 9 && v1_7_4.IntVal == 2098047435) {
v_6_loop = (*Data_Data_Map_Internal_Node)(v_6.UnsafePtr).V5
continue go__5_2
__t5 = gopurs_runtime.Value{}
goto end_branch_5
} else {

}
}
{
if (v1_7_4.Type == 9 && v1_7_4.IntVal == 1111389260) {
__t5 = gopurs_runtime.Value{Type: 9, IntVal: 1354639136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Just{(*Data_Data_Map_Internal_Node)(v_6.UnsafePtr).V3})}
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
}), gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return applyMap1_1_0
}))
}()
})
	})
	return bindMap
}

var anyWithKey gopurs_runtime.Value
var once_anyWithKey sync.Once
func Get_anyWithKey() gopurs_runtime.Value {
	once_anyWithKey.Do(func() {
		anyWithKey = gopurs_runtime.Func(func(predicate_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
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
__t1 = gopurs_runtime.Bool(gopurs_runtime.Apply2(predicate_0, (*Data_Data_Map_Internal_Node)(v_2.UnsafePtr).V2, (*Data_Data_Map_Internal_Node)(v_2.UnsafePtr).V3).IntVal != 0 || gopurs_runtime.Apply(go__1_0, (*Data_Data_Map_Internal_Node)(v_2.UnsafePtr).V4).IntVal != 0 || gopurs_runtime.Apply(go__1_0, (*Data_Data_Map_Internal_Node)(v_2.UnsafePtr).V5).IntVal != 0)
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
}()
})
	})
	return anyWithKey
}

var any gopurs_runtime.Value
var once_any sync.Once
func Get_any() gopurs_runtime.Value {
	once_any.Do(func() {
		any = gopurs_runtime.Func(func(predicate_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
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
__t1 = gopurs_runtime.Bool(gopurs_runtime.Apply(predicate_0, (*Data_Data_Map_Internal_Node)(v_2.UnsafePtr).V3).IntVal != 0 || gopurs_runtime.Apply(go__1_0, (*Data_Data_Map_Internal_Node)(v_2.UnsafePtr).V4).IntVal != 0 || gopurs_runtime.Apply(go__1_0, (*Data_Data_Map_Internal_Node)(v_2.UnsafePtr).V5).IntVal != 0)
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
}()
})
	})
	return any
}

var alter gopurs_runtime.Value
var once_alter sync.Once
func Get_alter() gopurs_runtime.Value {
	once_alter.Do(func() {
		alter = gopurs_runtime.Func(func(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
compare_1_0 := gopurs_runtime.RecordGet(dictOrd_0, "compare")
_ = compare_1_0
return gopurs_runtime.Func3(func(f_2 gopurs_runtime.Value, k_3 gopurs_runtime.Value, m_4 gopurs_runtime.Value) gopurs_runtime.Value {
v_5_1 := gopurs_runtime.UncurriedApp3(Get_unsafeSplit(), compare_1_0, k_3, m_4)
_ = v_5_1
v2_6_2 := gopurs_runtime.Apply(f_2, (*Data_Data_Map_Internal_Split)(v_5_1.UnsafePtr).V0)
_ = v2_6_2
var __t3 gopurs_runtime.Value
{
if (v2_6_2.Type == 9 && v2_6_2.IntVal == 42808261) {
__t3 = gopurs_runtime.UncurriedApp2(Get_unsafeJoinNodes(), (*Data_Data_Map_Internal_Split)(v_5_1.UnsafePtr).V1, (*Data_Data_Map_Internal_Split)(v_5_1.UnsafePtr).V2)
goto end_branch_3
} else {

}
}
{
if (v2_6_2.Type == 9 && v2_6_2.IntVal == 1354639136) {
__t3 = gopurs_runtime.UncurriedApp4(Get_unsafeBalancedNode(), k_3, (*pkg_Data_Maybe.Data_Data_Maybe_Just)(v2_6_2.UnsafePtr).V0, (*Data_Data_Map_Internal_Split)(v_5_1.UnsafePtr).V1, (*Data_Data_Map_Internal_Split)(v_5_1.UnsafePtr).V2)
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
}()
})
	})
	return alter
}

var altMap gopurs_runtime.Value
var once_altMap sync.Once
func Get_altMap() gopurs_runtime.Value {
	once_altMap.Do(func() {
		altMap = gopurs_runtime.Func(func(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
compare_1_0 := gopurs_runtime.RecordGet(dictOrd_0, "compare")
_ = compare_1_0
return gopurs_runtime.RecordDict2("alt", "Functor0", gopurs_runtime.Func2(func(m1_2 gopurs_runtime.Value, m2_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.UncurriedApp4(Get_unsafeUnionWith(), compare_1_0, pkg_Data_Function.Get_const_(), m1_2, m2_3)
}), gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_functorMap()
}))
}()
})
	})
	return altMap
}

var plusMap gopurs_runtime.Value
var once_plusMap sync.Once
func Get_plusMap() gopurs_runtime.Value {
	once_plusMap.Do(func() {
		plusMap = gopurs_runtime.Func(func(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
compare_1_1 := gopurs_runtime.RecordGet(dictOrd_0, "compare")
_ = compare_1_1
altMap1_1_0 := gopurs_runtime.RecordDict2("alt", "Functor0", gopurs_runtime.Func2(func(m1_2 gopurs_runtime.Value, m2_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.UncurriedApp4(Get_unsafeUnionWith(), compare_1_1, pkg_Data_Function.Get_const_(), m1_2, m2_3)
}), gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_functorMap()
}))
_ = altMap1_1_0
return gopurs_runtime.RecordDict2("empty", "Alt0", gopurs_runtime.Value{Type: 9, IntVal: 687041424, UnsafePtr: unsafe.Pointer(&Data_Data_Map_Internal_Leaf{})}, gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return altMap1_1_0
}))
}()
})
	})
	return plusMap
}

type Data_Data_Map_Internal_Leaf struct {
	
}
func Is_Data_Data_Map_Internal_Leaf(v gopurs_runtime.Value) bool {
	return v.Type == 9 && v.IntVal == 687041424
}

type Data_Data_Map_Internal_Node struct {
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
	V2 gopurs_runtime.Value
	V3 gopurs_runtime.Value
	V4 gopurs_runtime.Value
	V5 gopurs_runtime.Value
}
func Is_Data_Data_Map_Internal_Node(v gopurs_runtime.Value) bool {
	return v.Type == 9 && v.IntVal == 324739070
}

type Data_Data_Map_Internal_IterLeaf struct {
	
}
func Is_Data_Data_Map_Internal_IterLeaf(v gopurs_runtime.Value) bool {
	return v.Type == 9 && v.IntVal == 2509360378
}

type Data_Data_Map_Internal_IterEmit struct {
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
	V2 gopurs_runtime.Value
}
func Is_Data_Data_Map_Internal_IterEmit(v gopurs_runtime.Value) bool {
	return v.Type == 9 && v.IntVal == 1343415489
}

type Data_Data_Map_Internal_IterNode struct {
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
}
func Is_Data_Data_Map_Internal_IterNode(v gopurs_runtime.Value) bool {
	return v.Type == 9 && v.IntVal == 2861335956
}

type Data_Data_Map_Internal_IterDone struct {
	
}
func Is_Data_Data_Map_Internal_IterDone(v gopurs_runtime.Value) bool {
	return v.Type == 9 && v.IntVal == 4236111124
}

type Data_Data_Map_Internal_IterNext struct {
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
	V2 gopurs_runtime.Value
}
func Is_Data_Data_Map_Internal_IterNext(v gopurs_runtime.Value) bool {
	return v.Type == 9 && v.IntVal == 953589075
}

type Data_Data_Map_Internal_Split struct {
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
	V2 gopurs_runtime.Value
}
func Is_Data_Data_Map_Internal_Split(v gopurs_runtime.Value) bool {
	return v.Type == 9 && v.IntVal == 3373277644
}

type Data_Data_Map_Internal_SplitLast struct {
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
	V2 gopurs_runtime.Value
}
func Is_Data_Data_Map_Internal_SplitLast(v gopurs_runtime.Value) bool {
	return v.Type == 9 && v.IntVal == 2668112006
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
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(&Data_Data_Map_Internal_Node{gopurs_runtime.Int(1), gopurs_runtime.Int(1), k_0, v_1, l_2, r_3})}
goto end_branch_1
} else {

}
}
{
if (r_3.Type == 9 && r_3.IntVal == 324739070) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(&Data_Data_Map_Internal_Node{gopurs_runtime.Int(1 + (*Data_Data_Map_Internal_Node)(r_3.UnsafePtr).V0.IntVal), gopurs_runtime.Int(1 + (*Data_Data_Map_Internal_Node)(r_3.UnsafePtr).V1.IntVal), k_0, v_1, l_2, r_3})}
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
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(&Data_Data_Map_Internal_Node{gopurs_runtime.Int(1 + (*Data_Data_Map_Internal_Node)(l_2.UnsafePtr).V0.IntVal), gopurs_runtime.Int(1 + (*Data_Data_Map_Internal_Node)(l_2.UnsafePtr).V1.IntVal), k_0, v_1, l_2, r_3})}
goto end_branch_2
} else {

}
}
{
if (r_3.Type == 9 && r_3.IntVal == 324739070) {
var __t3 gopurs_runtime.Value
{
if (*Data_Data_Map_Internal_Node)(l_2.UnsafePtr).V0.IntVal > (*Data_Data_Map_Internal_Node)(r_3.UnsafePtr).V0.IntVal {
__t3 = gopurs_runtime.Int(1 + (*Data_Data_Map_Internal_Node)(l_2.UnsafePtr).V0.IntVal)
goto end_branch_3
} else {

}
}
{
__t3 = gopurs_runtime.Int(1 + (*Data_Data_Map_Internal_Node)(r_3.UnsafePtr).V0.IntVal)
}
end_branch_3:
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(&Data_Data_Map_Internal_Node{__t3, gopurs_runtime.Int(1 + (*Data_Data_Map_Internal_Node)(l_2.UnsafePtr).V1.IntVal + (*Data_Data_Map_Internal_Node)(r_3.UnsafePtr).V1.IntVal), k_0, v_1, l_2, r_3})}
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

func Call_stepWith(f_0_loop gopurs_runtime.Value, next_1_loop gopurs_runtime.Value, done_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var next_1 gopurs_runtime.Value = next_1_loop
_ = next_1
var done_2 gopurs_runtime.Value = done_2_loop
_ = done_2
var go__3_0 gopurs_runtime.Value
go__3_0 = gopurs_runtime.Func(func(v_4_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
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
__t1 = gopurs_runtime.UncurriedApp3(next_1, (*Data_Data_Map_Internal_IterEmit)(v_4.UnsafePtr).V0, (*Data_Data_Map_Internal_IterEmit)(v_4.UnsafePtr).V1, (*Data_Data_Map_Internal_IterEmit)(v_4.UnsafePtr).V2)
goto end_branch_1
} else {

}
}
{
if (v_4.Type == 9 && v_4.IntVal == 2861335956) {
v_4_loop = gopurs_runtime.Apply2(f_0, (*Data_Data_Map_Internal_IterNode)(v_4.UnsafePtr).V1, (*Data_Data_Map_Internal_IterNode)(v_4.UnsafePtr).V0)
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

func Call_singleton(k_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var k_0 gopurs_runtime.Value = k_0_loop
_ = k_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(&Data_Data_Map_Internal_Node{gopurs_runtime.Int(1), gopurs_runtime.Int(1), k_0, v_1, gopurs_runtime.Value{Type: 9, IntVal: 687041424, UnsafePtr: unsafe.Pointer(&Data_Data_Map_Internal_Leaf{})}, gopurs_runtime.Value{Type: 9, IntVal: 687041424, UnsafePtr: unsafe.Pointer(&Data_Data_Map_Internal_Leaf{})}})}
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
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(&Data_Data_Map_Internal_Node{gopurs_runtime.Int(1), gopurs_runtime.Int(1), k_0, v_1, gopurs_runtime.Value{Type: 9, IntVal: 687041424, UnsafePtr: unsafe.Pointer(&Data_Data_Map_Internal_Leaf{})}, gopurs_runtime.Value{Type: 9, IntVal: 687041424, UnsafePtr: unsafe.Pointer(&Data_Data_Map_Internal_Leaf{})}})}
goto end_branch_1
} else {

}
}
{
if (r_3.Type == 9 && r_3.IntVal == 324739070) && (*Data_Data_Map_Internal_Node)(r_3.UnsafePtr).V0.IntVal > 1 {
var __t2 gopurs_runtime.Value
{
var __t3 gopurs_runtime.Value
{
if ((*Data_Data_Map_Internal_Node)(r_3.UnsafePtr).V5.Type == 9 && (*Data_Data_Map_Internal_Node)(r_3.UnsafePtr).V5.IntVal == 687041424) {
__t3 = gopurs_runtime.Bool((*Data_Data_Map_Internal_Node)((*Data_Data_Map_Internal_Node)(r_3.UnsafePtr).V4.UnsafePtr).V0.IntVal > 0)
goto end_branch_3
} else {

}
}
{
if ((*Data_Data_Map_Internal_Node)(r_3.UnsafePtr).V5.Type == 9 && (*Data_Data_Map_Internal_Node)(r_3.UnsafePtr).V5.IntVal == 324739070) {
__t3 = gopurs_runtime.Bool((*Data_Data_Map_Internal_Node)((*Data_Data_Map_Internal_Node)(r_3.UnsafePtr).V4.UnsafePtr).V0.IntVal > (*Data_Data_Map_Internal_Node)((*Data_Data_Map_Internal_Node)(r_3.UnsafePtr).V5.UnsafePtr).V0.IntVal)
goto end_branch_3
} else {

}
}
{
__t3 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_3:
if ((*Data_Data_Map_Internal_Node)(r_3.UnsafePtr).V4.Type == 9 && (*Data_Data_Map_Internal_Node)(r_3.UnsafePtr).V4.IntVal == 324739070) && __t3.IntVal != 0 {
__t2 = gopurs_runtime.UncurriedApp4(Get_unsafeNode(), (*Data_Data_Map_Internal_Node)((*Data_Data_Map_Internal_Node)(r_3.UnsafePtr).V4.UnsafePtr).V2, (*Data_Data_Map_Internal_Node)((*Data_Data_Map_Internal_Node)(r_3.UnsafePtr).V4.UnsafePtr).V3, gopurs_runtime.UncurriedApp4(Get_unsafeNode(), k_0, v_1, l_2, (*Data_Data_Map_Internal_Node)((*Data_Data_Map_Internal_Node)(r_3.UnsafePtr).V4.UnsafePtr).V4), gopurs_runtime.UncurriedApp4(Get_unsafeNode(), (*Data_Data_Map_Internal_Node)(r_3.UnsafePtr).V2, (*Data_Data_Map_Internal_Node)(r_3.UnsafePtr).V3, (*Data_Data_Map_Internal_Node)((*Data_Data_Map_Internal_Node)(r_3.UnsafePtr).V4.UnsafePtr).V5, (*Data_Data_Map_Internal_Node)(r_3.UnsafePtr).V5))
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.UncurriedApp4(Get_unsafeNode(), (*Data_Data_Map_Internal_Node)(r_3.UnsafePtr).V2, (*Data_Data_Map_Internal_Node)(r_3.UnsafePtr).V3, gopurs_runtime.UncurriedApp4(Get_unsafeNode(), k_0, v_1, l_2, (*Data_Data_Map_Internal_Node)(r_3.UnsafePtr).V4), (*Data_Data_Map_Internal_Node)(r_3.UnsafePtr).V5)
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
var __t4 gopurs_runtime.Value
{
if (r_3.Type == 9 && r_3.IntVal == 324739070) {
var __t5 gopurs_runtime.Value
{
if (*Data_Data_Map_Internal_Node)(r_3.UnsafePtr).V0.IntVal > (*Data_Data_Map_Internal_Node)(l_2.UnsafePtr).V0.IntVal + 1 {
var __t6 gopurs_runtime.Value
{
var __t7 gopurs_runtime.Value
{
if ((*Data_Data_Map_Internal_Node)(r_3.UnsafePtr).V5.Type == 9 && (*Data_Data_Map_Internal_Node)(r_3.UnsafePtr).V5.IntVal == 687041424) {
__t7 = gopurs_runtime.Bool((*Data_Data_Map_Internal_Node)((*Data_Data_Map_Internal_Node)(r_3.UnsafePtr).V4.UnsafePtr).V0.IntVal > 0)
goto end_branch_7
} else {

}
}
{
if ((*Data_Data_Map_Internal_Node)(r_3.UnsafePtr).V5.Type == 9 && (*Data_Data_Map_Internal_Node)(r_3.UnsafePtr).V5.IntVal == 324739070) {
__t7 = gopurs_runtime.Bool((*Data_Data_Map_Internal_Node)((*Data_Data_Map_Internal_Node)(r_3.UnsafePtr).V4.UnsafePtr).V0.IntVal > (*Data_Data_Map_Internal_Node)((*Data_Data_Map_Internal_Node)(r_3.UnsafePtr).V5.UnsafePtr).V0.IntVal)
goto end_branch_7
} else {

}
}
{
__t7 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_7:
if ((*Data_Data_Map_Internal_Node)(r_3.UnsafePtr).V4.Type == 9 && (*Data_Data_Map_Internal_Node)(r_3.UnsafePtr).V4.IntVal == 324739070) && __t7.IntVal != 0 {
__t6 = gopurs_runtime.UncurriedApp4(Get_unsafeNode(), (*Data_Data_Map_Internal_Node)((*Data_Data_Map_Internal_Node)(r_3.UnsafePtr).V4.UnsafePtr).V2, (*Data_Data_Map_Internal_Node)((*Data_Data_Map_Internal_Node)(r_3.UnsafePtr).V4.UnsafePtr).V3, gopurs_runtime.UncurriedApp4(Get_unsafeNode(), k_0, v_1, l_2, (*Data_Data_Map_Internal_Node)((*Data_Data_Map_Internal_Node)(r_3.UnsafePtr).V4.UnsafePtr).V4), gopurs_runtime.UncurriedApp4(Get_unsafeNode(), (*Data_Data_Map_Internal_Node)(r_3.UnsafePtr).V2, (*Data_Data_Map_Internal_Node)(r_3.UnsafePtr).V3, (*Data_Data_Map_Internal_Node)((*Data_Data_Map_Internal_Node)(r_3.UnsafePtr).V4.UnsafePtr).V5, (*Data_Data_Map_Internal_Node)(r_3.UnsafePtr).V5))
goto end_branch_6
} else {

}
}
{
__t6 = gopurs_runtime.UncurriedApp4(Get_unsafeNode(), (*Data_Data_Map_Internal_Node)(r_3.UnsafePtr).V2, (*Data_Data_Map_Internal_Node)(r_3.UnsafePtr).V3, gopurs_runtime.UncurriedApp4(Get_unsafeNode(), k_0, v_1, l_2, (*Data_Data_Map_Internal_Node)(r_3.UnsafePtr).V4), (*Data_Data_Map_Internal_Node)(r_3.UnsafePtr).V5)
}
end_branch_6:
__t5 = __t6
goto end_branch_5
} else {

}
}
{
if (*Data_Data_Map_Internal_Node)(l_2.UnsafePtr).V0.IntVal > (*Data_Data_Map_Internal_Node)(r_3.UnsafePtr).V0.IntVal + 1 {
var __t8 gopurs_runtime.Value
{
var __t9 gopurs_runtime.Value
{
if ((*Data_Data_Map_Internal_Node)(l_2.UnsafePtr).V4.Type == 9 && (*Data_Data_Map_Internal_Node)(l_2.UnsafePtr).V4.IntVal == 687041424) {
__t9 = gopurs_runtime.Bool(0 <= (*Data_Data_Map_Internal_Node)((*Data_Data_Map_Internal_Node)(l_2.UnsafePtr).V5.UnsafePtr).V0.IntVal)
goto end_branch_9
} else {

}
}
{
if ((*Data_Data_Map_Internal_Node)(l_2.UnsafePtr).V4.Type == 9 && (*Data_Data_Map_Internal_Node)(l_2.UnsafePtr).V4.IntVal == 324739070) {
__t9 = gopurs_runtime.Bool((*Data_Data_Map_Internal_Node)((*Data_Data_Map_Internal_Node)(l_2.UnsafePtr).V4.UnsafePtr).V0.IntVal <= (*Data_Data_Map_Internal_Node)((*Data_Data_Map_Internal_Node)(l_2.UnsafePtr).V5.UnsafePtr).V0.IntVal)
goto end_branch_9
} else {

}
}
{
__t9 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_9:
if ((*Data_Data_Map_Internal_Node)(l_2.UnsafePtr).V5.Type == 9 && (*Data_Data_Map_Internal_Node)(l_2.UnsafePtr).V5.IntVal == 324739070) && __t9.IntVal != 0 {
__t8 = gopurs_runtime.UncurriedApp4(Get_unsafeNode(), (*Data_Data_Map_Internal_Node)((*Data_Data_Map_Internal_Node)(l_2.UnsafePtr).V5.UnsafePtr).V2, (*Data_Data_Map_Internal_Node)((*Data_Data_Map_Internal_Node)(l_2.UnsafePtr).V5.UnsafePtr).V3, gopurs_runtime.UncurriedApp4(Get_unsafeNode(), (*Data_Data_Map_Internal_Node)(l_2.UnsafePtr).V2, (*Data_Data_Map_Internal_Node)(l_2.UnsafePtr).V3, (*Data_Data_Map_Internal_Node)(l_2.UnsafePtr).V4, (*Data_Data_Map_Internal_Node)((*Data_Data_Map_Internal_Node)(l_2.UnsafePtr).V5.UnsafePtr).V4), gopurs_runtime.UncurriedApp4(Get_unsafeNode(), k_0, v_1, (*Data_Data_Map_Internal_Node)((*Data_Data_Map_Internal_Node)(l_2.UnsafePtr).V5.UnsafePtr).V5, r_3))
goto end_branch_8
} else {

}
}
{
__t8 = gopurs_runtime.UncurriedApp4(Get_unsafeNode(), (*Data_Data_Map_Internal_Node)(l_2.UnsafePtr).V2, (*Data_Data_Map_Internal_Node)(l_2.UnsafePtr).V3, (*Data_Data_Map_Internal_Node)(l_2.UnsafePtr).V4, gopurs_runtime.UncurriedApp4(Get_unsafeNode(), k_0, v_1, (*Data_Data_Map_Internal_Node)(l_2.UnsafePtr).V5, r_3))
}
end_branch_8:
__t5 = __t8
goto end_branch_5
} else {

}
}
{
__t5 = gopurs_runtime.UncurriedApp4(Get_unsafeNode(), k_0, v_1, l_2, r_3)
}
end_branch_5:
__t4 = __t5
goto end_branch_4
} else {

}
}
{
if (r_3.Type == 9 && r_3.IntVal == 687041424) && (*Data_Data_Map_Internal_Node)(l_2.UnsafePtr).V0.IntVal > 1 {
var __t10 gopurs_runtime.Value
{
var __t11 gopurs_runtime.Value
{
if ((*Data_Data_Map_Internal_Node)(l_2.UnsafePtr).V4.Type == 9 && (*Data_Data_Map_Internal_Node)(l_2.UnsafePtr).V4.IntVal == 687041424) {
__t11 = gopurs_runtime.Bool(0 <= (*Data_Data_Map_Internal_Node)((*Data_Data_Map_Internal_Node)(l_2.UnsafePtr).V5.UnsafePtr).V0.IntVal)
goto end_branch_11
} else {

}
}
{
if ((*Data_Data_Map_Internal_Node)(l_2.UnsafePtr).V4.Type == 9 && (*Data_Data_Map_Internal_Node)(l_2.UnsafePtr).V4.IntVal == 324739070) {
__t11 = gopurs_runtime.Bool((*Data_Data_Map_Internal_Node)((*Data_Data_Map_Internal_Node)(l_2.UnsafePtr).V4.UnsafePtr).V0.IntVal <= (*Data_Data_Map_Internal_Node)((*Data_Data_Map_Internal_Node)(l_2.UnsafePtr).V5.UnsafePtr).V0.IntVal)
goto end_branch_11
} else {

}
}
{
__t11 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_11:
if ((*Data_Data_Map_Internal_Node)(l_2.UnsafePtr).V5.Type == 9 && (*Data_Data_Map_Internal_Node)(l_2.UnsafePtr).V5.IntVal == 324739070) && __t11.IntVal != 0 {
__t10 = gopurs_runtime.UncurriedApp4(Get_unsafeNode(), (*Data_Data_Map_Internal_Node)((*Data_Data_Map_Internal_Node)(l_2.UnsafePtr).V5.UnsafePtr).V2, (*Data_Data_Map_Internal_Node)((*Data_Data_Map_Internal_Node)(l_2.UnsafePtr).V5.UnsafePtr).V3, gopurs_runtime.UncurriedApp4(Get_unsafeNode(), (*Data_Data_Map_Internal_Node)(l_2.UnsafePtr).V2, (*Data_Data_Map_Internal_Node)(l_2.UnsafePtr).V3, (*Data_Data_Map_Internal_Node)(l_2.UnsafePtr).V4, (*Data_Data_Map_Internal_Node)((*Data_Data_Map_Internal_Node)(l_2.UnsafePtr).V5.UnsafePtr).V4), gopurs_runtime.UncurriedApp4(Get_unsafeNode(), k_0, v_1, (*Data_Data_Map_Internal_Node)((*Data_Data_Map_Internal_Node)(l_2.UnsafePtr).V5.UnsafePtr).V5, r_3))
goto end_branch_10
} else {

}
}
{
__t10 = gopurs_runtime.UncurriedApp4(Get_unsafeNode(), (*Data_Data_Map_Internal_Node)(l_2.UnsafePtr).V2, (*Data_Data_Map_Internal_Node)(l_2.UnsafePtr).V3, (*Data_Data_Map_Internal_Node)(l_2.UnsafePtr).V4, gopurs_runtime.UncurriedApp4(Get_unsafeNode(), k_0, v_1, (*Data_Data_Map_Internal_Node)(l_2.UnsafePtr).V5, r_3))
}
end_branch_10:
__t4 = __t10
goto end_branch_4
} else {

}
}
{
__t4 = gopurs_runtime.UncurriedApp4(Get_unsafeNode(), k_0, v_1, l_2, r_3)
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
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 3373277644, UnsafePtr: unsafe.Pointer(&Data_Data_Map_Internal_Split{gopurs_runtime.Value{Type: 9, IntVal: 42808261, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Nothing{})}, gopurs_runtime.Value{Type: 9, IntVal: 687041424, UnsafePtr: unsafe.Pointer(&Data_Data_Map_Internal_Leaf{})}, gopurs_runtime.Value{Type: 9, IntVal: 687041424, UnsafePtr: unsafe.Pointer(&Data_Data_Map_Internal_Leaf{})}})}
goto end_branch_0
} else {

}
}
{
if (m_2.Type == 9 && m_2.IntVal == 324739070) {
v_3_1 := gopurs_runtime.Apply2(comp_0, k_1, (*Data_Data_Map_Internal_Node)(m_2.UnsafePtr).V2)
_ = v_3_1
var __t2 gopurs_runtime.Value
{
if (v_3_1.Type == 9 && v_3_1.IntVal == 3866105248) {
v1_4_3 := gopurs_runtime.UncurriedApp3(Get_unsafeSplit(), comp_0, k_1, (*Data_Data_Map_Internal_Node)(m_2.UnsafePtr).V4)
_ = v1_4_3
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 3373277644, UnsafePtr: unsafe.Pointer(&Data_Data_Map_Internal_Split{(*Data_Data_Map_Internal_Split)(v1_4_3.UnsafePtr).V0, (*Data_Data_Map_Internal_Split)(v1_4_3.UnsafePtr).V1, gopurs_runtime.UncurriedApp4(Get_unsafeBalancedNode(), (*Data_Data_Map_Internal_Node)(m_2.UnsafePtr).V2, (*Data_Data_Map_Internal_Node)(m_2.UnsafePtr).V3, (*Data_Data_Map_Internal_Split)(v1_4_3.UnsafePtr).V2, (*Data_Data_Map_Internal_Node)(m_2.UnsafePtr).V5)})}
goto end_branch_2
} else {

}
}
{
if (v_3_1.Type == 9 && v_3_1.IntVal == 2098047435) {
v1_4_4 := gopurs_runtime.UncurriedApp3(Get_unsafeSplit(), comp_0, k_1, (*Data_Data_Map_Internal_Node)(m_2.UnsafePtr).V5)
_ = v1_4_4
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 3373277644, UnsafePtr: unsafe.Pointer(&Data_Data_Map_Internal_Split{(*Data_Data_Map_Internal_Split)(v1_4_4.UnsafePtr).V0, gopurs_runtime.UncurriedApp4(Get_unsafeBalancedNode(), (*Data_Data_Map_Internal_Node)(m_2.UnsafePtr).V2, (*Data_Data_Map_Internal_Node)(m_2.UnsafePtr).V3, (*Data_Data_Map_Internal_Node)(m_2.UnsafePtr).V4, (*Data_Data_Map_Internal_Split)(v1_4_4.UnsafePtr).V1), (*Data_Data_Map_Internal_Split)(v1_4_4.UnsafePtr).V2})}
goto end_branch_2
} else {

}
}
{
if (v_3_1.Type == 9 && v_3_1.IntVal == 1111389260) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 3373277644, UnsafePtr: unsafe.Pointer(&Data_Data_Map_Internal_Split{gopurs_runtime.Value{Type: 9, IntVal: 1354639136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Just{(*Data_Data_Map_Internal_Node)(m_2.UnsafePtr).V3})}, (*Data_Data_Map_Internal_Node)(m_2.UnsafePtr).V4, (*Data_Data_Map_Internal_Node)(m_2.UnsafePtr).V5})}
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
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 2668112006, UnsafePtr: unsafe.Pointer(&Data_Data_Map_Internal_SplitLast{k_0, v_1, l_2})}
goto end_branch_0
} else {

}
}
{
if (r_3.Type == 9 && r_3.IntVal == 324739070) {
v1_4_1 := gopurs_runtime.UncurriedApp4(Get_unsafeSplitLast(), (*Data_Data_Map_Internal_Node)(r_3.UnsafePtr).V2, (*Data_Data_Map_Internal_Node)(r_3.UnsafePtr).V3, (*Data_Data_Map_Internal_Node)(r_3.UnsafePtr).V4, (*Data_Data_Map_Internal_Node)(r_3.UnsafePtr).V5)
_ = v1_4_1
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 2668112006, UnsafePtr: unsafe.Pointer(&Data_Data_Map_Internal_SplitLast{(*Data_Data_Map_Internal_SplitLast)(v1_4_1.UnsafePtr).V0, (*Data_Data_Map_Internal_SplitLast)(v1_4_1.UnsafePtr).V1, gopurs_runtime.UncurriedApp4(Get_unsafeBalancedNode(), k_0, v_1, l_2, (*Data_Data_Map_Internal_SplitLast)(v1_4_1.UnsafePtr).V2)})}
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
v2_2_1 := gopurs_runtime.UncurriedApp4(Get_unsafeSplitLast(), (*Data_Data_Map_Internal_Node)(v_0.UnsafePtr).V2, (*Data_Data_Map_Internal_Node)(v_0.UnsafePtr).V3, (*Data_Data_Map_Internal_Node)(v_0.UnsafePtr).V4, (*Data_Data_Map_Internal_Node)(v_0.UnsafePtr).V5)
_ = v2_2_1
__t0 = gopurs_runtime.UncurriedApp4(Get_unsafeBalancedNode(), (*Data_Data_Map_Internal_SplitLast)(v2_2_1.UnsafePtr).V0, (*Data_Data_Map_Internal_SplitLast)(v2_2_1.UnsafePtr).V1, (*Data_Data_Map_Internal_SplitLast)(v2_2_1.UnsafePtr).V2, v1_1)
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
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 687041424, UnsafePtr: unsafe.Pointer(&Data_Data_Map_Internal_Leaf{})}
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
v_3_1 := gopurs_runtime.UncurriedApp3(Get_unsafeSplit(), comp_0, (*Data_Data_Map_Internal_Node)(r_2.UnsafePtr).V2, l_1)
_ = v_3_1
__t0 = gopurs_runtime.UncurriedApp2(Get_unsafeJoinNodes(), gopurs_runtime.UncurriedApp3(Get_unsafeDifference(), comp_0, (*Data_Data_Map_Internal_Split)(v_3_1.UnsafePtr).V1, (*Data_Data_Map_Internal_Node)(r_2.UnsafePtr).V4), gopurs_runtime.UncurriedApp3(Get_unsafeDifference(), comp_0, (*Data_Data_Map_Internal_Split)(v_3_1.UnsafePtr).V2, (*Data_Data_Map_Internal_Node)(r_2.UnsafePtr).V5))
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
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 687041424, UnsafePtr: unsafe.Pointer(&Data_Data_Map_Internal_Leaf{})}
goto end_branch_0
} else {

}
}
{
if (r_3.Type == 9 && r_3.IntVal == 687041424) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 687041424, UnsafePtr: unsafe.Pointer(&Data_Data_Map_Internal_Leaf{})}
goto end_branch_0
} else {

}
}
{
if (r_3.Type == 9 && r_3.IntVal == 324739070) {
v_4_1 := gopurs_runtime.UncurriedApp3(Get_unsafeSplit(), comp_0, (*Data_Data_Map_Internal_Node)(r_3.UnsafePtr).V2, l_2)
_ = v_4_1
l_prime_5_2 := gopurs_runtime.UncurriedApp4(Get_unsafeIntersectionWith(), comp_0, app_1, (*Data_Data_Map_Internal_Split)(v_4_1.UnsafePtr).V1, (*Data_Data_Map_Internal_Node)(r_3.UnsafePtr).V4)
_ = l_prime_5_2
r_prime_6_3 := gopurs_runtime.UncurriedApp4(Get_unsafeIntersectionWith(), comp_0, app_1, (*Data_Data_Map_Internal_Split)(v_4_1.UnsafePtr).V2, (*Data_Data_Map_Internal_Node)(r_3.UnsafePtr).V5)
_ = r_prime_6_3
var __t4 gopurs_runtime.Value
{
if ((*Data_Data_Map_Internal_Split)(v_4_1.UnsafePtr).V0.Type == 9 && (*Data_Data_Map_Internal_Split)(v_4_1.UnsafePtr).V0.IntVal == 1354639136) {
__t4 = gopurs_runtime.UncurriedApp4(Get_unsafeBalancedNode(), (*Data_Data_Map_Internal_Node)(r_3.UnsafePtr).V2, gopurs_runtime.Apply2(app_1, (*pkg_Data_Maybe.Data_Data_Maybe_Just)((*Data_Data_Map_Internal_Split)(v_4_1.UnsafePtr).V0.UnsafePtr).V0, (*Data_Data_Map_Internal_Node)(r_3.UnsafePtr).V3), l_prime_5_2, r_prime_6_3)
goto end_branch_4
} else {

}
}
{
if ((*Data_Data_Map_Internal_Split)(v_4_1.UnsafePtr).V0.Type == 9 && (*Data_Data_Map_Internal_Split)(v_4_1.UnsafePtr).V0.IntVal == 42808261) {
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
v_4_1 := gopurs_runtime.UncurriedApp3(Get_unsafeSplit(), comp_0, (*Data_Data_Map_Internal_Node)(r_3.UnsafePtr).V2, l_2)
_ = v_4_1
l_prime_5_2 := gopurs_runtime.UncurriedApp4(Get_unsafeUnionWith(), comp_0, app_1, (*Data_Data_Map_Internal_Split)(v_4_1.UnsafePtr).V1, (*Data_Data_Map_Internal_Node)(r_3.UnsafePtr).V4)
_ = l_prime_5_2
r_prime_6_3 := gopurs_runtime.UncurriedApp4(Get_unsafeUnionWith(), comp_0, app_1, (*Data_Data_Map_Internal_Split)(v_4_1.UnsafePtr).V2, (*Data_Data_Map_Internal_Node)(r_3.UnsafePtr).V5)
_ = r_prime_6_3
var __t4 gopurs_runtime.Value
{
if ((*Data_Data_Map_Internal_Split)(v_4_1.UnsafePtr).V0.Type == 9 && (*Data_Data_Map_Internal_Split)(v_4_1.UnsafePtr).V0.IntVal == 1354639136) {
__t4 = gopurs_runtime.UncurriedApp4(Get_unsafeBalancedNode(), (*Data_Data_Map_Internal_Node)(r_3.UnsafePtr).V2, gopurs_runtime.Apply2(app_1, (*pkg_Data_Maybe.Data_Data_Maybe_Just)((*Data_Data_Map_Internal_Split)(v_4_1.UnsafePtr).V0.UnsafePtr).V0, (*Data_Data_Map_Internal_Node)(r_3.UnsafePtr).V3), l_prime_5_2, r_prime_6_3)
goto end_branch_4
} else {

}
}
{
if ((*Data_Data_Map_Internal_Split)(v_4_1.UnsafePtr).V0.Type == 9 && (*Data_Data_Map_Internal_Split)(v_4_1.UnsafePtr).V0.IntVal == 42808261) {
__t4 = gopurs_runtime.UncurriedApp4(Get_unsafeBalancedNode(), (*Data_Data_Map_Internal_Node)(r_3.UnsafePtr).V2, (*Data_Data_Map_Internal_Node)(r_3.UnsafePtr).V3, l_prime_5_2, r_prime_6_3)
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
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 687041424, UnsafePtr: unsafe.Pointer(&Data_Data_Map_Internal_Leaf{})}
goto end_branch_1
} else {

}
}
{
if (v_4.Type == 9 && v_4.IntVal == 324739070) {
v1_5_2 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictOrd_0, "compare"), k_2, (*Data_Data_Map_Internal_Node)(v_4.UnsafePtr).V2)
_ = v1_5_2
var __t3 gopurs_runtime.Value
{
if (v1_5_2.Type == 9 && v1_5_2.IntVal == 3866105248) {
__t3 = gopurs_runtime.UncurriedApp4(Get_unsafeBalancedNode(), (*Data_Data_Map_Internal_Node)(v_4.UnsafePtr).V2, (*Data_Data_Map_Internal_Node)(v_4.UnsafePtr).V3, gopurs_runtime.Apply(go__3_0, (*Data_Data_Map_Internal_Node)(v_4.UnsafePtr).V4), (*Data_Data_Map_Internal_Node)(v_4.UnsafePtr).V5)
goto end_branch_3
} else {

}
}
{
if (v1_5_2.Type == 9 && v1_5_2.IntVal == 2098047435) {
__t3 = gopurs_runtime.UncurriedApp4(Get_unsafeBalancedNode(), (*Data_Data_Map_Internal_Node)(v_4.UnsafePtr).V2, (*Data_Data_Map_Internal_Node)(v_4.UnsafePtr).V3, (*Data_Data_Map_Internal_Node)(v_4.UnsafePtr).V4, gopurs_runtime.Apply(go__3_0, (*Data_Data_Map_Internal_Node)(v_4.UnsafePtr).V5))
goto end_branch_3
} else {

}
}
{
if (v1_5_2.Type == 9 && v1_5_2.IntVal == 1111389260) {
v2_6_4 := gopurs_runtime.Apply(f_1, (*Data_Data_Map_Internal_Node)(v_4.UnsafePtr).V3)
_ = v2_6_4
var __t5 gopurs_runtime.Value
{
if (v2_6_4.Type == 9 && v2_6_4.IntVal == 42808261) {
__t5 = gopurs_runtime.UncurriedApp2(Get_unsafeJoinNodes(), (*Data_Data_Map_Internal_Node)(v_4.UnsafePtr).V4, (*Data_Data_Map_Internal_Node)(v_4.UnsafePtr).V5)
goto end_branch_5
} else {

}
}
{
if (v2_6_4.Type == 9 && v2_6_4.IntVal == 1354639136) {
__t5 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(&Data_Data_Map_Internal_Node{(*Data_Data_Map_Internal_Node)(v_4.UnsafePtr).V0, (*Data_Data_Map_Internal_Node)(v_4.UnsafePtr).V1, (*Data_Data_Map_Internal_Node)(v_4.UnsafePtr).V2, (*pkg_Data_Maybe.Data_Data_Maybe_Just)(v2_6_4.UnsafePtr).V0, (*Data_Data_Map_Internal_Node)(v_4.UnsafePtr).V4, (*Data_Data_Map_Internal_Node)(v_4.UnsafePtr).V5})}
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
__t1 = gopurs_runtime.Str(ind_3.StrVal() + "Leaf")
goto end_branch_1
} else {

}
}
{
if (v_4.Type == 9 && v_4.IntVal == 324739070) {
__t1 = gopurs_runtime.Str(ind_3.StrVal() + "[" + gopurs_runtime.Apply(pkg_Data_Show.Get_showIntImpl(), (*Data_Data_Map_Internal_Node)(v_4.UnsafePtr).V0).StrVal() + "] " + gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictShow_0, "show"), (*Data_Data_Map_Internal_Node)(v_4.UnsafePtr).V2).StrVal() + " => " + gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictShow1_1, "show"), (*Data_Data_Map_Internal_Node)(v_4.UnsafePtr).V3).StrVal() + "\n" + gopurs_runtime.Apply2(go__2_0, gopurs_runtime.Str(ind_3.StrVal() + "    "), (*Data_Data_Map_Internal_Node)(v_4.UnsafePtr).V4).StrVal() + "\n" + gopurs_runtime.Apply2(go__2_0, gopurs_runtime.Str(ind_3.StrVal() + "    "), (*Data_Data_Map_Internal_Node)(v_4.UnsafePtr).V5).StrVal())
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
compare_2_0 := gopurs_runtime.RecordGet(dictOrd_1, "compare")
_ = compare_2_0
return gopurs_runtime.Func(func(dictSemigroup_3 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_4_1 := gopurs_runtime.RecordGet(dictSemigroup_3, "append")
_ = __local_var_4_1
return gopurs_runtime.RecordDict1("append", gopurs_runtime.Func2(func(m1_5 gopurs_runtime.Value, m2_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.UncurriedApp4(Get_unsafeUnionWith(), compare_2_0, __local_var_4_1, m1_5, m2_6)
}))
})
}

func Call_member(dictOrd_0_loop gopurs_runtime.Value, k_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
var k_1 gopurs_runtime.Value = k_1_loop
_ = k_1
var go__2_0 gopurs_runtime.Value
go__2_0 = gopurs_runtime.Func(func(v_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
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
v1_4_2 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictOrd_0, "compare"), k_1, (*Data_Data_Map_Internal_Node)(v_3.UnsafePtr).V2)
_ = v1_4_2
var __t3 gopurs_runtime.Value
{
if (v1_4_2.Type == 9 && v1_4_2.IntVal == 3866105248) {
v_3_loop = (*Data_Data_Map_Internal_Node)(v_3.UnsafePtr).V4
continue go__2_0
__t3 = gopurs_runtime.Value{}
goto end_branch_3
} else {

}
}
{
if (v1_4_2.Type == 9 && v1_4_2.IntVal == 2098047435) {
v_3_loop = (*Data_Data_Map_Internal_Node)(v_3.UnsafePtr).V5
continue go__2_0
__t3 = gopurs_runtime.Value{}
goto end_branch_3
} else {

}
}
{
if (v1_4_2.Type == 9 && v1_4_2.IntVal == 1111389260) {
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
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 687041424, UnsafePtr: unsafe.Pointer(&Data_Data_Map_Internal_Leaf{})}
goto end_branch_1
} else {

}
}
{
if (v_3.Type == 9 && v_3.IntVal == 324739070) {
v2_4_2 := gopurs_runtime.Apply2(f_1, (*Data_Data_Map_Internal_Node)(v_3.UnsafePtr).V2, (*Data_Data_Map_Internal_Node)(v_3.UnsafePtr).V3)
_ = v2_4_2
var __t3 gopurs_runtime.Value
{
if (v2_4_2.Type == 9 && v2_4_2.IntVal == 1354639136) {
__t3 = gopurs_runtime.UncurriedApp4(Get_unsafeBalancedNode(), (*Data_Data_Map_Internal_Node)(v_3.UnsafePtr).V2, (*pkg_Data_Maybe.Data_Data_Maybe_Just)(v2_4_2.UnsafePtr).V0, gopurs_runtime.Apply(go__2_0, (*Data_Data_Map_Internal_Node)(v_3.UnsafePtr).V4), gopurs_runtime.Apply(go__2_0, (*Data_Data_Map_Internal_Node)(v_3.UnsafePtr).V5))
goto end_branch_3
} else {

}
}
{
if (v2_4_2.Type == 9 && v2_4_2.IntVal == 42808261) {
__t3 = gopurs_runtime.UncurriedApp2(Get_unsafeJoinNodes(), gopurs_runtime.Apply(go__2_0, (*Data_Data_Map_Internal_Node)(v_3.UnsafePtr).V4), gopurs_runtime.Apply(go__2_0, (*Data_Data_Map_Internal_Node)(v_3.UnsafePtr).V5))
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
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 42808261, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Nothing{})}
goto end_branch_1
} else {

}
}
{
if (v_3.Type == 9 && v_3.IntVal == 324739070) {
v1_4_2 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictOrd_0, "compare"), k_1, (*Data_Data_Map_Internal_Node)(v_3.UnsafePtr).V2)
_ = v1_4_2
var __t3 gopurs_runtime.Value
{
if (v1_4_2.Type == 9 && v1_4_2.IntVal == 3866105248) {
__t3 = gopurs_runtime.Apply(go__2_0, (*Data_Data_Map_Internal_Node)(v_3.UnsafePtr).V4)
goto end_branch_3
} else {

}
}
{
if (v1_4_2.Type == 9 && v1_4_2.IntVal == 2098047435) {
v2_5_4 := gopurs_runtime.Apply(go__2_0, (*Data_Data_Map_Internal_Node)(v_3.UnsafePtr).V5)
_ = v2_5_4
var __t5 gopurs_runtime.Value
{
if (v2_5_4.Type == 9 && v2_5_4.IntVal == 42808261) {
__t5 = gopurs_runtime.Value{Type: 9, IntVal: 1354639136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Just{gopurs_runtime.RecordDict2("key", "value", (*Data_Data_Map_Internal_Node)(v_3.UnsafePtr).V2, (*Data_Data_Map_Internal_Node)(v_3.UnsafePtr).V3)})}
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
if (v1_4_2.Type == 9 && v1_4_2.IntVal == 1111389260) {
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 1354639136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Just{gopurs_runtime.RecordDict2("key", "value", (*Data_Data_Map_Internal_Node)(v_3.UnsafePtr).V2, (*Data_Data_Map_Internal_Node)(v_3.UnsafePtr).V3)})}
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
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 42808261, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Nothing{})}
goto end_branch_1
} else {

}
}
{
if (v_3.Type == 9 && v_3.IntVal == 324739070) {
v1_4_2 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictOrd_0, "compare"), k_1, (*Data_Data_Map_Internal_Node)(v_3.UnsafePtr).V2)
_ = v1_4_2
var __t3 gopurs_runtime.Value
{
if (v1_4_2.Type == 9 && v1_4_2.IntVal == 3866105248) {
v2_5_4 := gopurs_runtime.Apply(go__2_0, (*Data_Data_Map_Internal_Node)(v_3.UnsafePtr).V4)
_ = v2_5_4
var __t5 gopurs_runtime.Value
{
if (v2_5_4.Type == 9 && v2_5_4.IntVal == 42808261) {
__t5 = gopurs_runtime.Value{Type: 9, IntVal: 1354639136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Just{gopurs_runtime.RecordDict2("key", "value", (*Data_Data_Map_Internal_Node)(v_3.UnsafePtr).V2, (*Data_Data_Map_Internal_Node)(v_3.UnsafePtr).V3)})}
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
if (v1_4_2.Type == 9 && v1_4_2.IntVal == 2098047435) {
__t3 = gopurs_runtime.Apply(go__2_0, (*Data_Data_Map_Internal_Node)(v_3.UnsafePtr).V5)
goto end_branch_3
} else {

}
}
{
if (v1_4_2.Type == 9 && v1_4_2.IntVal == 1111389260) {
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 1354639136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Just{gopurs_runtime.RecordDict2("key", "value", (*Data_Data_Map_Internal_Node)(v_3.UnsafePtr).V2, (*Data_Data_Map_Internal_Node)(v_3.UnsafePtr).V3)})}
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
go__2_0 = gopurs_runtime.Func(func(v_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
go__2_0:
for {
if false { continue go__2_0 }
var v_3 gopurs_runtime.Value = v_3_loop
_ = v_3
var __t1 gopurs_runtime.Value
{
if (v_3.Type == 9 && v_3.IntVal == 687041424) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 42808261, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Nothing{})}
goto end_branch_1
} else {

}
}
{
if (v_3.Type == 9 && v_3.IntVal == 324739070) {
v1_4_2 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictOrd_0, "compare"), k_1, (*Data_Data_Map_Internal_Node)(v_3.UnsafePtr).V2)
_ = v1_4_2
var __t3 gopurs_runtime.Value
{
if (v1_4_2.Type == 9 && v1_4_2.IntVal == 3866105248) {
v_3_loop = (*Data_Data_Map_Internal_Node)(v_3.UnsafePtr).V4
continue go__2_0
__t3 = gopurs_runtime.Value{}
goto end_branch_3
} else {

}
}
{
if (v1_4_2.Type == 9 && v1_4_2.IntVal == 2098047435) {
v_3_loop = (*Data_Data_Map_Internal_Node)(v_3.UnsafePtr).V5
continue go__2_0
__t3 = gopurs_runtime.Value{}
goto end_branch_3
} else {

}
}
{
if (v1_4_2.Type == 9 && v1_4_2.IntVal == 1111389260) {
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 1354639136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Just{(*Data_Data_Map_Internal_Node)(v_3.UnsafePtr).V3})}
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

func Call_iterMapU(iter_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var iter_0 gopurs_runtime.Value = iter_0_loop
_ = iter_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
var __t0 gopurs_runtime.Value
{
if (v_1.Type == 9 && v_1.IntVal == 687041424) {
__t0 = iter_0
goto end_branch_0
} else {

}
}
{
if (v_1.Type == 9 && v_1.IntVal == 324739070) {
var __t1 gopurs_runtime.Value
{
if ((*Data_Data_Map_Internal_Node)(v_1.UnsafePtr).V4.Type == 9 && (*Data_Data_Map_Internal_Node)(v_1.UnsafePtr).V4.IntVal == 687041424) {
var __t2 gopurs_runtime.Value
{
if ((*Data_Data_Map_Internal_Node)(v_1.UnsafePtr).V5.Type == 9 && (*Data_Data_Map_Internal_Node)(v_1.UnsafePtr).V5.IntVal == 687041424) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 1343415489, UnsafePtr: unsafe.Pointer(&Data_Data_Map_Internal_IterEmit{(*Data_Data_Map_Internal_Node)(v_1.UnsafePtr).V2, (*Data_Data_Map_Internal_Node)(v_1.UnsafePtr).V3, iter_0})}
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 1343415489, UnsafePtr: unsafe.Pointer(&Data_Data_Map_Internal_IterEmit{(*Data_Data_Map_Internal_Node)(v_1.UnsafePtr).V2, (*Data_Data_Map_Internal_Node)(v_1.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 2861335956, UnsafePtr: unsafe.Pointer(&Data_Data_Map_Internal_IterNode{(*Data_Data_Map_Internal_Node)(v_1.UnsafePtr).V5, iter_0})}})}
}
end_branch_2:
__t1 = __t2
goto end_branch_1
} else {

}
}
{
if ((*Data_Data_Map_Internal_Node)(v_1.UnsafePtr).V5.Type == 9 && (*Data_Data_Map_Internal_Node)(v_1.UnsafePtr).V5.IntVal == 687041424) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 1343415489, UnsafePtr: unsafe.Pointer(&Data_Data_Map_Internal_IterEmit{(*Data_Data_Map_Internal_Node)(v_1.UnsafePtr).V2, (*Data_Data_Map_Internal_Node)(v_1.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 2861335956, UnsafePtr: unsafe.Pointer(&Data_Data_Map_Internal_IterNode{(*Data_Data_Map_Internal_Node)(v_1.UnsafePtr).V4, iter_0})}})}
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 1343415489, UnsafePtr: unsafe.Pointer(&Data_Data_Map_Internal_IterEmit{(*Data_Data_Map_Internal_Node)(v_1.UnsafePtr).V2, (*Data_Data_Map_Internal_Node)(v_1.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 2861335956, UnsafePtr: unsafe.Pointer(&Data_Data_Map_Internal_IterNode{(*Data_Data_Map_Internal_Node)(v_1.UnsafePtr).V4, gopurs_runtime.Value{Type: 9, IntVal: 2861335956, UnsafePtr: unsafe.Pointer(&Data_Data_Map_Internal_IterNode{(*Data_Data_Map_Internal_Node)(v_1.UnsafePtr).V5, iter_0})}})}})}
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
__t2 = gopurs_runtime.Bool((v2_6_3.Type == 9 && v2_6_3.IntVal == 953589075) && gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq_0, "eq"), (*Data_Data_Map_Internal_IterNext)(v_5_1.UnsafePtr).V0, (*Data_Data_Map_Internal_IterNext)(v2_6_3.UnsafePtr).V0).IntVal != 0 && gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq1_1, "eq"), (*Data_Data_Map_Internal_IterNext)(v_5_1.UnsafePtr).V1, (*Data_Data_Map_Internal_IterNext)(v2_6_3.UnsafePtr).V1).IntVal != 0 && gopurs_runtime.Apply2(go__2_0, (*Data_Data_Map_Internal_IterNext)(v_5_1.UnsafePtr).V2, (*Data_Data_Map_Internal_IterNext)(v2_6_3.UnsafePtr).V2).IntVal != 0)
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

func Call_showMap(dictShow_0_loop gopurs_runtime.Value, dictShow1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictShow_0 gopurs_runtime.Value = dictShow_0_loop
_ = dictShow_0
var dictShow1_1 gopurs_runtime.Value = dictShow1_1_loop
_ = dictShow1_1
show1_2_0 := gopurs_runtime.Apply(pkg_Data_Show.Get_showArrayImpl(), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str("(Tuple " + gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictShow_0, "show"), (*pkg_Data_Tuple.Data_Data_Tuple_Tuple)(v_2.UnsafePtr).V0).StrVal() + " " + gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictShow1_1, "show"), (*pkg_Data_Tuple.Data_Data_Tuple_Tuple)(v_2.UnsafePtr).V1).StrVal() + ")")
}))
_ = show1_2_0
return gopurs_runtime.RecordDict1("show", gopurs_runtime.Func(func(as_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str("(fromFoldable " + gopurs_runtime.Apply(show1_2_0, gopurs_runtime.Apply(Get_toUnfoldable1(), as_3)).StrVal() + ")")
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
__local_var_5_2 := (*Data_Data_Map_Internal_Node)(m1_3.UnsafePtr).V2
_ = __local_var_5_2
var go__6_3 gopurs_runtime.Value
go__6_3 = gopurs_runtime.Func(func(v_7_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
go__6_3:
for {
if false { continue go__6_3 }
var v_7 gopurs_runtime.Value = v_7_loop
_ = v_7
var __t4 gopurs_runtime.Value
{
if (v_7.Type == 9 && v_7.IntVal == 687041424) {
__t4 = gopurs_runtime.Value{Type: 9, IntVal: 42808261, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Nothing{})}
goto end_branch_4
} else {

}
}
{
if (v_7.Type == 9 && v_7.IntVal == 324739070) {
v1_8_5 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictOrd_0, "compare"), __local_var_5_2, (*Data_Data_Map_Internal_Node)(v_7.UnsafePtr).V2)
_ = v1_8_5
var __t6 gopurs_runtime.Value
{
if (v1_8_5.Type == 9 && v1_8_5.IntVal == 3866105248) {
v_7_loop = (*Data_Data_Map_Internal_Node)(v_7.UnsafePtr).V4
continue go__6_3
__t6 = gopurs_runtime.Value{}
goto end_branch_6
} else {

}
}
{
if (v1_8_5.Type == 9 && v1_8_5.IntVal == 2098047435) {
v_7_loop = (*Data_Data_Map_Internal_Node)(v_7.UnsafePtr).V5
continue go__6_3
__t6 = gopurs_runtime.Value{}
goto end_branch_6
} else {

}
}
{
if (v1_8_5.Type == 9 && v1_8_5.IntVal == 1111389260) {
__t6 = gopurs_runtime.Value{Type: 9, IntVal: 1354639136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Just{(*Data_Data_Map_Internal_Node)(v_7.UnsafePtr).V3})}
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
if (v1_7_7.Type == 9 && v1_7_7.IntVal == 42808261) {
__t8 = gopurs_runtime.Bool(false)
goto end_branch_8
} else {

}
}
{
if (v1_7_7.Type == 9 && v1_7_7.IntVal == 1354639136) {
__t8 = gopurs_runtime.Bool(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq_1, "eq"), (*Data_Data_Map_Internal_Node)(m1_3.UnsafePtr).V3, (*pkg_Data_Maybe.Data_Data_Maybe_Just)(v1_7_7.UnsafePtr).V0).IntVal != 0 && gopurs_runtime.Apply2(go__2_0, (*Data_Data_Map_Internal_Node)(m1_3.UnsafePtr).V4, m2_4).IntVal != 0 && gopurs_runtime.Apply2(go__2_0, (*Data_Data_Map_Internal_Node)(m1_3.UnsafePtr).V5, m2_4).IntVal != 0)
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
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(&Data_Data_Map_Internal_Node{gopurs_runtime.Int(1), gopurs_runtime.Int(1), k_2, v_3, gopurs_runtime.Value{Type: 9, IntVal: 687041424, UnsafePtr: unsafe.Pointer(&Data_Data_Map_Internal_Leaf{})}, gopurs_runtime.Value{Type: 9, IntVal: 687041424, UnsafePtr: unsafe.Pointer(&Data_Data_Map_Internal_Leaf{})}})}
goto end_branch_1
} else {

}
}
{
if (v1_5.Type == 9 && v1_5.IntVal == 324739070) {
v2_6_2 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictOrd_0, "compare"), k_2, (*Data_Data_Map_Internal_Node)(v1_5.UnsafePtr).V2)
_ = v2_6_2
var __t3 gopurs_runtime.Value
{
if (v2_6_2.Type == 9 && v2_6_2.IntVal == 3866105248) {
__t3 = gopurs_runtime.UncurriedApp4(Get_unsafeBalancedNode(), (*Data_Data_Map_Internal_Node)(v1_5.UnsafePtr).V2, (*Data_Data_Map_Internal_Node)(v1_5.UnsafePtr).V3, gopurs_runtime.Apply(go__4_0, (*Data_Data_Map_Internal_Node)(v1_5.UnsafePtr).V4), (*Data_Data_Map_Internal_Node)(v1_5.UnsafePtr).V5)
goto end_branch_3
} else {

}
}
{
if (v2_6_2.Type == 9 && v2_6_2.IntVal == 2098047435) {
__t3 = gopurs_runtime.UncurriedApp4(Get_unsafeBalancedNode(), (*Data_Data_Map_Internal_Node)(v1_5.UnsafePtr).V2, (*Data_Data_Map_Internal_Node)(v1_5.UnsafePtr).V3, (*Data_Data_Map_Internal_Node)(v1_5.UnsafePtr).V4, gopurs_runtime.Apply(go__4_0, (*Data_Data_Map_Internal_Node)(v1_5.UnsafePtr).V5))
goto end_branch_3
} else {

}
}
{
if (v2_6_2.Type == 9 && v2_6_2.IntVal == 1111389260) {
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(&Data_Data_Map_Internal_Node{(*Data_Data_Map_Internal_Node)(v1_5.UnsafePtr).V0, (*Data_Data_Map_Internal_Node)(v1_5.UnsafePtr).V1, k_2, gopurs_runtime.Apply2(app_1, (*Data_Data_Map_Internal_Node)(v1_5.UnsafePtr).V3, v_3), (*Data_Data_Map_Internal_Node)(v1_5.UnsafePtr).V4, (*Data_Data_Map_Internal_Node)(v1_5.UnsafePtr).V5})}
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
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(&Data_Data_Map_Internal_Node{gopurs_runtime.Int(1), gopurs_runtime.Int(1), k_1, v_2, gopurs_runtime.Value{Type: 9, IntVal: 687041424, UnsafePtr: unsafe.Pointer(&Data_Data_Map_Internal_Leaf{})}, gopurs_runtime.Value{Type: 9, IntVal: 687041424, UnsafePtr: unsafe.Pointer(&Data_Data_Map_Internal_Leaf{})}})}
goto end_branch_1
} else {

}
}
{
if (v1_4.Type == 9 && v1_4.IntVal == 324739070) {
v2_5_2 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictOrd_0, "compare"), k_1, (*Data_Data_Map_Internal_Node)(v1_4.UnsafePtr).V2)
_ = v2_5_2
var __t3 gopurs_runtime.Value
{
if (v2_5_2.Type == 9 && v2_5_2.IntVal == 3866105248) {
__t3 = gopurs_runtime.UncurriedApp4(Get_unsafeBalancedNode(), (*Data_Data_Map_Internal_Node)(v1_4.UnsafePtr).V2, (*Data_Data_Map_Internal_Node)(v1_4.UnsafePtr).V3, gopurs_runtime.Apply(go__3_0, (*Data_Data_Map_Internal_Node)(v1_4.UnsafePtr).V4), (*Data_Data_Map_Internal_Node)(v1_4.UnsafePtr).V5)
goto end_branch_3
} else {

}
}
{
if (v2_5_2.Type == 9 && v2_5_2.IntVal == 2098047435) {
__t3 = gopurs_runtime.UncurriedApp4(Get_unsafeBalancedNode(), (*Data_Data_Map_Internal_Node)(v1_4.UnsafePtr).V2, (*Data_Data_Map_Internal_Node)(v1_4.UnsafePtr).V3, (*Data_Data_Map_Internal_Node)(v1_4.UnsafePtr).V4, gopurs_runtime.Apply(go__3_0, (*Data_Data_Map_Internal_Node)(v1_4.UnsafePtr).V5))
goto end_branch_3
} else {

}
}
{
if (v2_5_2.Type == 9 && v2_5_2.IntVal == 1111389260) {
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(&Data_Data_Map_Internal_Node{(*Data_Data_Map_Internal_Node)(v1_4.UnsafePtr).V0, (*Data_Data_Map_Internal_Node)(v1_4.UnsafePtr).V1, k_1, v_2, (*Data_Data_Map_Internal_Node)(v1_4.UnsafePtr).V4, (*Data_Data_Map_Internal_Node)(v1_4.UnsafePtr).V5})}
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

func Call_go__2_0(m_prime_3 gopurs_runtime.Value, z_prime_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if (m_prime_3.Type == 9 && m_prime_3.IntVal == 687041424) {
__t1 = z_prime_4
goto end_branch_1
} else {

}
}
{
if (m_prime_3.Type == 9 && m_prime_3.IntVal == 324739070) {
__t1 = gopurs_runtime.UncurriedApp2(go__2_0, (*Data_Data_Map_Internal_Node)(m_prime_3.UnsafePtr).V4, gopurs_runtime.Apply2(f_0, (*Data_Data_Map_Internal_Node)(m_prime_3.UnsafePtr).V3, gopurs_runtime.UncurriedApp2(go__2_0, (*Data_Data_Map_Internal_Node)(m_prime_3.UnsafePtr).V5, z_prime_4)))
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

func Call_go__2_2(z_prime_3 gopurs_runtime.Value, m_prime_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t3 gopurs_runtime.Value
{
if (m_prime_4.Type == 9 && m_prime_4.IntVal == 687041424) {
__t3 = z_prime_3
goto end_branch_3
} else {

}
}
{
if (m_prime_4.Type == 9 && m_prime_4.IntVal == 324739070) {
__t3 = gopurs_runtime.UncurriedApp2(go__2_2, gopurs_runtime.Apply2(f_0, gopurs_runtime.UncurriedApp2(go__2_2, z_prime_3, (*Data_Data_Map_Internal_Node)(m_prime_4.UnsafePtr).V4), (*Data_Data_Map_Internal_Node)(m_prime_4.UnsafePtr).V3), (*Data_Data_Map_Internal_Node)(m_prime_4.UnsafePtr).V5)
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

func Call_go__2_0(m_prime_3 gopurs_runtime.Value, z_prime_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if (m_prime_3.Type == 9 && m_prime_3.IntVal == 687041424) {
__t1 = z_prime_4
goto end_branch_1
} else {

}
}
{
if (m_prime_3.Type == 9 && m_prime_3.IntVal == 324739070) {
__t1 = gopurs_runtime.UncurriedApp2(go__2_0, (*Data_Data_Map_Internal_Node)(m_prime_3.UnsafePtr).V4, gopurs_runtime.Apply3(f_0, (*Data_Data_Map_Internal_Node)(m_prime_3.UnsafePtr).V2, (*Data_Data_Map_Internal_Node)(m_prime_3.UnsafePtr).V3, gopurs_runtime.UncurriedApp2(go__2_0, (*Data_Data_Map_Internal_Node)(m_prime_3.UnsafePtr).V5, z_prime_4)))
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

func Call_go__2_2(z_prime_3 gopurs_runtime.Value, m_prime_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t3 gopurs_runtime.Value
{
if (m_prime_4.Type == 9 && m_prime_4.IntVal == 687041424) {
__t3 = z_prime_3
goto end_branch_3
} else {

}
}
{
if (m_prime_4.Type == 9 && m_prime_4.IntVal == 324739070) {
__t3 = gopurs_runtime.UncurriedApp2(go__2_2, gopurs_runtime.Apply3(f_0, (*Data_Data_Map_Internal_Node)(m_prime_4.UnsafePtr).V2, gopurs_runtime.UncurriedApp2(go__2_2, z_prime_3, (*Data_Data_Map_Internal_Node)(m_prime_4.UnsafePtr).V4), (*Data_Data_Map_Internal_Node)(m_prime_4.UnsafePtr).V3), (*Data_Data_Map_Internal_Node)(m_prime_4.UnsafePtr).V5)
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

func Call_go__0_0(m_prime_1 gopurs_runtime.Value, z_prime_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if (m_prime_1.Type == 9 && m_prime_1.IntVal == 687041424) {
__t1 = z_prime_2
goto end_branch_1
} else {

}
}
{
if (m_prime_1.Type == 9 && m_prime_1.IntVal == 324739070) {
__t1 = gopurs_runtime.UncurriedApp2(go__0_0, (*Data_Data_Map_Internal_Node)(m_prime_1.UnsafePtr).V4, gopurs_runtime.Value{Type: 9, IntVal: 2709581417, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Data_Data_List_Types_Cons{(*Data_Data_Map_Internal_Node)(m_prime_1.UnsafePtr).V2, gopurs_runtime.UncurriedApp2(go__0_0, (*Data_Data_Map_Internal_Node)(m_prime_1.UnsafePtr).V5, z_prime_2)})})
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

func Call_go__0_0(m_prime_1 gopurs_runtime.Value, z_prime_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if (m_prime_1.Type == 9 && m_prime_1.IntVal == 687041424) {
__t1 = z_prime_2
goto end_branch_1
} else {

}
}
{
if (m_prime_1.Type == 9 && m_prime_1.IntVal == 324739070) {
__t1 = gopurs_runtime.UncurriedApp2(go__0_0, (*Data_Data_Map_Internal_Node)(m_prime_1.UnsafePtr).V4, gopurs_runtime.Value{Type: 9, IntVal: 2709581417, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Data_Data_List_Types_Cons{(*Data_Data_Map_Internal_Node)(m_prime_1.UnsafePtr).V3, gopurs_runtime.UncurriedApp2(go__0_0, (*Data_Data_Map_Internal_Node)(m_prime_1.UnsafePtr).V5, z_prime_2)})})
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

func Call_foldSubmapBy(dictOrd_0_loop gopurs_runtime.Value, appendFn_1_loop gopurs_runtime.Value, memptyValue_2_loop gopurs_runtime.Value, kmin_3_loop gopurs_runtime.Value, kmax_4_loop gopurs_runtime.Value, f_5_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
var appendFn_1 gopurs_runtime.Value = appendFn_1_loop
_ = appendFn_1
var memptyValue_2 gopurs_runtime.Value = memptyValue_2_loop
_ = memptyValue_2
var kmin_3 gopurs_runtime.Value = kmin_3_loop
_ = kmin_3
var kmax_4 gopurs_runtime.Value = kmax_4_loop
_ = kmax_4
var f_5 gopurs_runtime.Value = f_5_loop
_ = f_5
var __t1 gopurs_runtime.Value
{
if (kmin_3.Type == 9 && kmin_3.IntVal == 1354639136) {
__local_var_6_2 := (*pkg_Data_Maybe.Data_Data_Maybe_Just)(kmin_3.UnsafePtr).V0
_ = __local_var_6_2
__t1 = gopurs_runtime.Func(func(k_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictOrd_0, "compare"), k_7, __local_var_6_2).Type == 9 && gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictOrd_0, "compare"), k_7, __local_var_6_2).IntVal == 3866105248))
})
goto end_branch_1
} else {

}
}
{
if (kmin_3.Type == 9 && kmin_3.IntVal == 42808261) {
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
var __t4 gopurs_runtime.Value
{
if (kmax_4.Type == 9 && kmax_4.IntVal == 1354639136) {
__local_var_7_5 := (*pkg_Data_Maybe.Data_Data_Maybe_Just)(kmax_4.UnsafePtr).V0
_ = __local_var_7_5
__t4 = gopurs_runtime.Func(func(k_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictOrd_0, "compare"), k_8, __local_var_7_5).Type == 9 && gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictOrd_0, "compare"), k_8, __local_var_7_5).IntVal == 2098047435))
})
goto end_branch_4
} else {

}
}
{
if (kmax_4.Type == 9 && kmax_4.IntVal == 42808261) {
__t4 = gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
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
tooLarge_7_3 := __t4
_ = tooLarge_7_3
var __t7 gopurs_runtime.Value
{
if (kmin_3.Type == 9 && kmin_3.IntVal == 1354639136) {
var __t8 gopurs_runtime.Value
{
if (kmax_4.Type == 9 && kmax_4.IntVal == 1354639136) {
__local_var_8_9 := (*pkg_Data_Maybe.Data_Data_Maybe_Just)(kmax_4.UnsafePtr).V0
_ = __local_var_8_9
__local_var_9_10 := (*pkg_Data_Maybe.Data_Data_Maybe_Just)(kmin_3.UnsafePtr).V0
_ = __local_var_9_10
__t8 = gopurs_runtime.Func(func(k_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictOrd_0, "compare"), __local_var_9_10, k_10).Type == 9 && gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictOrd_0, "compare"), __local_var_9_10, k_10).IntVal == 2098047435) != true && (gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictOrd_0, "compare"), k_10, __local_var_8_9).Type == 9 && gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictOrd_0, "compare"), k_10, __local_var_8_9).IntVal == 2098047435) != true)
})
goto end_branch_8
} else {

}
}
{
if (kmax_4.Type == 9 && kmax_4.IntVal == 42808261) {
__local_var_8_11 := (*pkg_Data_Maybe.Data_Data_Maybe_Just)(kmin_3.UnsafePtr).V0
_ = __local_var_8_11
__t8 = gopurs_runtime.Func(func(k_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictOrd_0, "compare"), __local_var_8_11, k_9).Type == 9 && gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictOrd_0, "compare"), __local_var_8_11, k_9).IntVal == 2098047435) != true)
})
goto end_branch_8
} else {

}
}
{
__t8 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_8:
__t7 = __t8
goto end_branch_7
} else {

}
}
{
if (kmin_3.Type == 9 && kmin_3.IntVal == 42808261) {
var __t12 gopurs_runtime.Value
{
if (kmax_4.Type == 9 && kmax_4.IntVal == 1354639136) {
__local_var_8_13 := (*pkg_Data_Maybe.Data_Data_Maybe_Just)(kmax_4.UnsafePtr).V0
_ = __local_var_8_13
__t12 = gopurs_runtime.Func(func(k_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictOrd_0, "compare"), k_9, __local_var_8_13).Type == 9 && gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictOrd_0, "compare"), k_9, __local_var_8_13).IntVal == 2098047435) != true)
})
goto end_branch_12
} else {

}
}
{
if (kmax_4.Type == 9 && kmax_4.IntVal == 42808261) {
__t12 = gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(true)
})
goto end_branch_12
} else {

}
}
{
__t12 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_12:
__t7 = __t12
goto end_branch_7
} else {

}
}
{
__t7 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_7:
inBounds_8_6 := __t7
_ = inBounds_8_6
var go__9_14 gopurs_runtime.Value
_ = go__9_14
go__9_14 = gopurs_runtime.Func(func(v_10 gopurs_runtime.Value) gopurs_runtime.Value {
var __t15 gopurs_runtime.Value
{
if (v_10.Type == 9 && v_10.IntVal == 687041424) {
__t15 = memptyValue_2
goto end_branch_15
} else {

}
}
{
if (v_10.Type == 9 && v_10.IntVal == 324739070) {
var __t16 gopurs_runtime.Value
{
if gopurs_runtime.Apply(tooSmall_6_0, (*Data_Data_Map_Internal_Node)(v_10.UnsafePtr).V2).IntVal != 0 {
__t16 = memptyValue_2
goto end_branch_16
} else {

}
}
{
__t16 = gopurs_runtime.Apply(go__9_14, (*Data_Data_Map_Internal_Node)(v_10.UnsafePtr).V4)
}
end_branch_16:
var __t17 gopurs_runtime.Value
{
if gopurs_runtime.Apply(inBounds_8_6, (*Data_Data_Map_Internal_Node)(v_10.UnsafePtr).V2).IntVal != 0 {
__t17 = gopurs_runtime.Apply2(f_5, (*Data_Data_Map_Internal_Node)(v_10.UnsafePtr).V2, (*Data_Data_Map_Internal_Node)(v_10.UnsafePtr).V3)
goto end_branch_17
} else {

}
}
{
__t17 = memptyValue_2
}
end_branch_17:
var __t18 gopurs_runtime.Value
{
if gopurs_runtime.Apply(tooLarge_7_3, (*Data_Data_Map_Internal_Node)(v_10.UnsafePtr).V2).IntVal != 0 {
__t18 = memptyValue_2
goto end_branch_18
} else {

}
}
{
__t18 = gopurs_runtime.Apply(go__9_14, (*Data_Data_Map_Internal_Node)(v_10.UnsafePtr).V5)
}
end_branch_18:
__t15 = gopurs_runtime.Apply2(appendFn_1, gopurs_runtime.Apply2(appendFn_1, __t16, __t17), __t18)
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
return go__9_14
}

func Call_foldSubmap(dictOrd_0_loop gopurs_runtime.Value, dictMonoid_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
var dictMonoid_1 gopurs_runtime.Value = dictMonoid_1_loop
_ = dictMonoid_1
return gopurs_runtime.Apply3(Get_foldSubmapBy(), dictOrd_0, gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_1, "Semigroup0"), gopurs_runtime.Value{}), "append"), gopurs_runtime.RecordGet(dictMonoid_1, "mempty"))
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
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 42808261, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Nothing{})}
goto end_branch_1
} else {

}
}
{
if (v_3.Type == 9 && v_3.IntVal == 324739070) {
v1_4_2 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictOrd_0, "compare"), k_1, (*Data_Data_Map_Internal_Node)(v_3.UnsafePtr).V2)
_ = v1_4_2
var __t3 gopurs_runtime.Value
{
if (v1_4_2.Type == 9 && v1_4_2.IntVal == 3866105248) {
v2_5_4 := gopurs_runtime.Apply(go__2_0, (*Data_Data_Map_Internal_Node)(v_3.UnsafePtr).V4)
_ = v2_5_4
var __t5 gopurs_runtime.Value
{
if (v2_5_4.Type == 9 && v2_5_4.IntVal == 42808261) {
__t5 = gopurs_runtime.Value{Type: 9, IntVal: 1354639136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Just{gopurs_runtime.RecordDict2("key", "value", (*Data_Data_Map_Internal_Node)(v_3.UnsafePtr).V2, (*Data_Data_Map_Internal_Node)(v_3.UnsafePtr).V3)})}
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
if (v1_4_2.Type == 9 && v1_4_2.IntVal == 2098047435) {
__t3 = gopurs_runtime.Apply(go__2_0, (*Data_Data_Map_Internal_Node)(v_3.UnsafePtr).V5)
goto end_branch_3
} else {

}
}
{
if (v1_4_2.Type == 9 && v1_4_2.IntVal == 1111389260) {
__t3 = gopurs_runtime.Apply(Get_findMin(), (*Data_Data_Map_Internal_Node)(v_3.UnsafePtr).V5)
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
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 42808261, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Nothing{})}
goto end_branch_1
} else {

}
}
{
if (v_3.Type == 9 && v_3.IntVal == 324739070) {
v1_4_2 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictOrd_0, "compare"), k_1, (*Data_Data_Map_Internal_Node)(v_3.UnsafePtr).V2)
_ = v1_4_2
var __t3 gopurs_runtime.Value
{
if (v1_4_2.Type == 9 && v1_4_2.IntVal == 3866105248) {
__t3 = gopurs_runtime.Apply(go__2_0, (*Data_Data_Map_Internal_Node)(v_3.UnsafePtr).V4)
goto end_branch_3
} else {

}
}
{
if (v1_4_2.Type == 9 && v1_4_2.IntVal == 2098047435) {
v2_5_4 := gopurs_runtime.Apply(go__2_0, (*Data_Data_Map_Internal_Node)(v_3.UnsafePtr).V5)
_ = v2_5_4
var __t5 gopurs_runtime.Value
{
if (v2_5_4.Type == 9 && v2_5_4.IntVal == 42808261) {
__t5 = gopurs_runtime.Value{Type: 9, IntVal: 1354639136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Just{gopurs_runtime.RecordDict2("key", "value", (*Data_Data_Map_Internal_Node)(v_3.UnsafePtr).V2, (*Data_Data_Map_Internal_Node)(v_3.UnsafePtr).V3)})}
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
if (v1_4_2.Type == 9 && v1_4_2.IntVal == 1111389260) {
__t3 = gopurs_runtime.Apply(Get_findMax(), (*Data_Data_Map_Internal_Node)(v_3.UnsafePtr).V4)
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
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 687041424, UnsafePtr: unsafe.Pointer(&Data_Data_Map_Internal_Leaf{})}
goto end_branch_1
} else {

}
}
{
if (v_3.Type == 9 && v_3.IntVal == 324739070) {
var __t2 gopurs_runtime.Value
{
if gopurs_runtime.Apply2(f_1, (*Data_Data_Map_Internal_Node)(v_3.UnsafePtr).V2, (*Data_Data_Map_Internal_Node)(v_3.UnsafePtr).V3).IntVal != 0 {
__t2 = gopurs_runtime.UncurriedApp4(Get_unsafeBalancedNode(), (*Data_Data_Map_Internal_Node)(v_3.UnsafePtr).V2, (*Data_Data_Map_Internal_Node)(v_3.UnsafePtr).V3, gopurs_runtime.Apply(go__2_0, (*Data_Data_Map_Internal_Node)(v_3.UnsafePtr).V4), gopurs_runtime.Apply(go__2_0, (*Data_Data_Map_Internal_Node)(v_3.UnsafePtr).V5))
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.UncurriedApp2(Get_unsafeJoinNodes(), gopurs_runtime.Apply(go__2_0, (*Data_Data_Map_Internal_Node)(v_3.UnsafePtr).V4), gopurs_runtime.Apply(go__2_0, (*Data_Data_Map_Internal_Node)(v_3.UnsafePtr).V5))
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
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 687041424, UnsafePtr: unsafe.Pointer(&Data_Data_Map_Internal_Leaf{})}
goto end_branch_1
} else {

}
}
{
if (v_3.Type == 9 && v_3.IntVal == 324739070) {
var __t2 gopurs_runtime.Value
{
if gopurs_runtime.Apply(f_1, (*Data_Data_Map_Internal_Node)(v_3.UnsafePtr).V2).IntVal != 0 {
__t2 = gopurs_runtime.UncurriedApp4(Get_unsafeBalancedNode(), (*Data_Data_Map_Internal_Node)(v_3.UnsafePtr).V2, (*Data_Data_Map_Internal_Node)(v_3.UnsafePtr).V3, gopurs_runtime.Apply(go__2_0, (*Data_Data_Map_Internal_Node)(v_3.UnsafePtr).V4), gopurs_runtime.Apply(go__2_0, (*Data_Data_Map_Internal_Node)(v_3.UnsafePtr).V5))
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.UncurriedApp2(Get_unsafeJoinNodes(), gopurs_runtime.Apply(go__2_0, (*Data_Data_Map_Internal_Node)(v_3.UnsafePtr).V4), gopurs_runtime.Apply(go__2_0, (*Data_Data_Map_Internal_Node)(v_3.UnsafePtr).V5))
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
__t0 = gopurs_runtime.Bool((ys_3.Type == 9 && ys_3.IntVal == 324739070) && (*Data_Data_Map_Internal_Node)(xs_2.UnsafePtr).V1.IntVal == (*Data_Data_Map_Internal_Node)(ys_3.UnsafePtr).V1.IntVal && gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Call_eqMapIter(dictEq_0, dictEq1_1), "eq"), gopurs_runtime.Value{Type: 9, IntVal: 2861335956, UnsafePtr: unsafe.Pointer(&Data_Data_Map_Internal_IterNode{xs_2, gopurs_runtime.Value{Type: 9, IntVal: 2509360378, UnsafePtr: unsafe.Pointer(&Data_Data_Map_Internal_IterLeaf{})}})}, gopurs_runtime.Value{Type: 9, IntVal: 2861335956, UnsafePtr: unsafe.Pointer(&Data_Data_Map_Internal_IterNode{ys_3, gopurs_runtime.Value{Type: 9, IntVal: 2509360378, UnsafePtr: unsafe.Pointer(&Data_Data_Map_Internal_IterLeaf{})}})}).IntVal != 0)
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

func Call_fromFoldable(dictOrd_0_loop gopurs_runtime.Value, dictFoldable_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
var dictFoldable_1 gopurs_runtime.Value = dictFoldable_1_loop
_ = dictFoldable_1
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFoldable_1, "foldl"), gopurs_runtime.Func2(func(m_2 gopurs_runtime.Value, v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Call_insert(dictOrd_0, (*pkg_Data_Tuple.Data_Data_Tuple_Tuple)(v_3.UnsafePtr).V0, (*pkg_Data_Tuple.Data_Data_Tuple_Tuple)(v_3.UnsafePtr).V1), m_2)
}), gopurs_runtime.Value{Type: 9, IntVal: 687041424, UnsafePtr: unsafe.Pointer(&Data_Data_Map_Internal_Leaf{})})
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
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFoldable_1, "foldl"), gopurs_runtime.Func2(func(m_4 gopurs_runtime.Value, v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(f_prime_3_0, (*pkg_Data_Tuple.Data_Data_Tuple_Tuple)(v_5.UnsafePtr).V0, (*pkg_Data_Tuple.Data_Data_Tuple_Tuple)(v_5.UnsafePtr).V1, m_4)
}), gopurs_runtime.Value{Type: 9, IntVal: 687041424, UnsafePtr: unsafe.Pointer(&Data_Data_Map_Internal_Leaf{})})
}

func Call_fromFoldableWithIndex(dictOrd_0_loop gopurs_runtime.Value, dictFoldableWithIndex_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
var dictFoldableWithIndex_1 gopurs_runtime.Value = dictFoldableWithIndex_1_loop
_ = dictFoldableWithIndex_1
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFoldableWithIndex_1, "foldlWithIndex"), gopurs_runtime.Func3(func(k_2 gopurs_runtime.Value, m_3 gopurs_runtime.Value, v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Call_insert(dictOrd_0, k_2, v_4), m_3)
}), gopurs_runtime.Value{Type: 9, IntVal: 687041424, UnsafePtr: unsafe.Pointer(&Data_Data_Map_Internal_Leaf{})})
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
return gopurs_runtime.RecordDict2("mempty", "Semigroup0", gopurs_runtime.Value{Type: 9, IntVal: 687041424, UnsafePtr: unsafe.Pointer(&Data_Data_Map_Internal_Leaf{})}, gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return semigroupMap3_4_1
}))
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
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 687041424, UnsafePtr: unsafe.Pointer(&Data_Data_Map_Internal_Leaf{})}
goto end_branch_1
} else {

}
}
{
if (v_3.Type == 9 && v_3.IntVal == 324739070) {
v1_4_2 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictOrd_0, "compare"), k_1, (*Data_Data_Map_Internal_Node)(v_3.UnsafePtr).V2)
_ = v1_4_2
var __t3 gopurs_runtime.Value
{
if (v1_4_2.Type == 9 && v1_4_2.IntVal == 3866105248) {
__t3 = gopurs_runtime.UncurriedApp4(Get_unsafeBalancedNode(), (*Data_Data_Map_Internal_Node)(v_3.UnsafePtr).V2, (*Data_Data_Map_Internal_Node)(v_3.UnsafePtr).V3, gopurs_runtime.Apply(go__2_0, (*Data_Data_Map_Internal_Node)(v_3.UnsafePtr).V4), (*Data_Data_Map_Internal_Node)(v_3.UnsafePtr).V5)
goto end_branch_3
} else {

}
}
{
if (v1_4_2.Type == 9 && v1_4_2.IntVal == 2098047435) {
__t3 = gopurs_runtime.UncurriedApp4(Get_unsafeBalancedNode(), (*Data_Data_Map_Internal_Node)(v_3.UnsafePtr).V2, (*Data_Data_Map_Internal_Node)(v_3.UnsafePtr).V3, (*Data_Data_Map_Internal_Node)(v_3.UnsafePtr).V4, gopurs_runtime.Apply(go__2_0, (*Data_Data_Map_Internal_Node)(v_3.UnsafePtr).V5))
goto end_branch_3
} else {

}
}
{
if (v1_4_2.Type == 9 && v1_4_2.IntVal == 1111389260) {
__t3 = gopurs_runtime.UncurriedApp2(Get_unsafeJoinNodes(), (*Data_Data_Map_Internal_Node)(v_3.UnsafePtr).V4, (*Data_Data_Map_Internal_Node)(v_3.UnsafePtr).V5)
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


