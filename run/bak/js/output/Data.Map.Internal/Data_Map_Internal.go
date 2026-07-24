package Data_Map_Internal

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_Unit "gopurs/output/Data.Unit"
	pkg_Data_Function "gopurs/output/Data.Function"
	pkg_Data_Show "gopurs/output/Data.Show"
	pkg_Data_Unfoldable "gopurs/output/Data.Unfoldable"
	pkg_Control_Category "gopurs/output/Control.Category"
)

var Leaf gopurs_runtime.Value
var once_Leaf sync.Once
func Get_Leaf() gopurs_runtime.Value {
	once_Leaf.Do(func() {
		Leaf = gopurs_runtime.Constructor0("Leaf")
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
return gopurs_runtime.Constructor("Node", []gopurs_runtime.Value{value0, value1, value2, value3, value4, value5})
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
		IterLeaf = gopurs_runtime.Constructor0("IterLeaf")
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
return gopurs_runtime.Constructor3("IterEmit", value0, value1, value2)
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
return gopurs_runtime.Constructor2("IterNode", value0, value1)
})
})
	})
	return IterNode
}

var IterDone gopurs_runtime.Value
var once_IterDone sync.Once
func Get_IterDone() gopurs_runtime.Value {
	once_IterDone.Do(func() {
		IterDone = gopurs_runtime.Constructor0("IterDone")
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
return gopurs_runtime.Constructor3("IterNext", value0, value1, value2)
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
return gopurs_runtime.Constructor3("Split", value0, value1, value2)
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
return gopurs_runtime.Constructor3("SplitLast", value0, value1, value2)
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
		unsafeNode = gopurs_runtime.Func4(Call_unsafeNode)
	})
	return unsafeNode
}

var toMapIter gopurs_runtime.Value
var once_toMapIter sync.Once
func Get_toMapIter() gopurs_runtime.Value {
	once_toMapIter.Do(func() {
		toMapIter = gopurs_runtime.Func(func(a_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Constructor2("IterNode", a_0, gopurs_runtime.Constructor0("IterLeaf"))
})
	})
	return toMapIter
}

var stepWith gopurs_runtime.Value
var once_stepWith sync.Once
func Get_stepWith() gopurs_runtime.Value {
	once_stepWith.Do(func() {
		stepWith = gopurs_runtime.Func3(func(f_0 gopurs_runtime.Value, next_1 gopurs_runtime.Value, done_2 gopurs_runtime.Value) gopurs_runtime.Value {
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
if gopurs_runtime.Bool(v_4.StrVal == "IterLeaf").IntVal != 0 {
__t1 = gopurs_runtime.Apply(done_2, pkg_Data_Unit.Get_unit())
goto end_branch_1
} else {

}
}
{
if gopurs_runtime.Bool(v_4.StrVal == "IterEmit").IntVal != 0 {
__t1 = gopurs_runtime.UncurriedApp3(next_1, (*[1024]gopurs_runtime.Value)(v_4.UnsafePtr)[0], (*[1024]gopurs_runtime.Value)(v_4.UnsafePtr)[1], (*[1024]gopurs_runtime.Value)(v_4.UnsafePtr)[2])
goto end_branch_1
} else {

}
}
{
if gopurs_runtime.Bool(v_4.StrVal == "IterNode").IntVal != 0 {
v_4_loop = gopurs_runtime.Apply2(f_0, (*[1024]gopurs_runtime.Value)(v_4.UnsafePtr)[1], (*[1024]gopurs_runtime.Value)(v_4.UnsafePtr)[0])
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
})
	})
	return stepWith
}

var size gopurs_runtime.Value
var once_size sync.Once
func Get_size() gopurs_runtime.Value {
	once_size.Do(func() {
		size = gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if gopurs_runtime.Bool(v_0.StrVal == "Leaf").IntVal != 0 {
__t0 = gopurs_runtime.Int(0)
goto end_branch_0
} else {

}
}
{
if gopurs_runtime.Bool(v_0.StrVal == "Node").IntVal != 0 {
__t0 = (*[1024]gopurs_runtime.Value)(v_0.UnsafePtr)[1]
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
	})
	return size
}

var singleton gopurs_runtime.Value
var once_singleton sync.Once
func Get_singleton() gopurs_runtime.Value {
	once_singleton.Do(func() {
		singleton = gopurs_runtime.Func2(func(k_0 gopurs_runtime.Value, v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Constructor("Node", []gopurs_runtime.Value{gopurs_runtime.Int(1), gopurs_runtime.Int(1), k_0, v_1, gopurs_runtime.Constructor0("Leaf"), gopurs_runtime.Constructor0("Leaf")})
})
	})
	return singleton
}

var unsafeBalancedNode gopurs_runtime.Value
var once_unsafeBalancedNode sync.Once
func Get_unsafeBalancedNode() gopurs_runtime.Value {
	once_unsafeBalancedNode.Do(func() {
		unsafeBalancedNode = gopurs_runtime.Func4(Call_unsafeBalancedNode)
	})
	return unsafeBalancedNode
}

var unsafeSplit gopurs_runtime.Value
var once_unsafeSplit sync.Once
func Get_unsafeSplit() gopurs_runtime.Value {
	once_unsafeSplit.Do(func() {
		unsafeSplit = gopurs_runtime.Func3(Call_unsafeSplit)
	})
	return unsafeSplit
}

var unsafeSplitLast gopurs_runtime.Value
var once_unsafeSplitLast sync.Once
func Get_unsafeSplitLast() gopurs_runtime.Value {
	once_unsafeSplitLast.Do(func() {
		unsafeSplitLast = gopurs_runtime.Func4(Call_unsafeSplitLast)
	})
	return unsafeSplitLast
}

var unsafeJoinNodes gopurs_runtime.Value
var once_unsafeJoinNodes sync.Once
func Get_unsafeJoinNodes() gopurs_runtime.Value {
	once_unsafeJoinNodes.Do(func() {
		unsafeJoinNodes = gopurs_runtime.Func2(Call_unsafeJoinNodes)
	})
	return unsafeJoinNodes
}

var unsafeDifference gopurs_runtime.Value
var once_unsafeDifference sync.Once
func Get_unsafeDifference() gopurs_runtime.Value {
	once_unsafeDifference.Do(func() {
		unsafeDifference = gopurs_runtime.Func3(Call_unsafeDifference)
	})
	return unsafeDifference
}

var unsafeIntersectionWith gopurs_runtime.Value
var once_unsafeIntersectionWith sync.Once
func Get_unsafeIntersectionWith() gopurs_runtime.Value {
	once_unsafeIntersectionWith.Do(func() {
		unsafeIntersectionWith = gopurs_runtime.Func4(Call_unsafeIntersectionWith)
	})
	return unsafeIntersectionWith
}

var unsafeUnionWith gopurs_runtime.Value
var once_unsafeUnionWith sync.Once
func Get_unsafeUnionWith() gopurs_runtime.Value {
	once_unsafeUnionWith.Do(func() {
		unsafeUnionWith = gopurs_runtime.Func4(Call_unsafeUnionWith)
	})
	return unsafeUnionWith
}

var unionWith gopurs_runtime.Value
var once_unionWith sync.Once
func Get_unionWith() gopurs_runtime.Value {
	once_unionWith.Do(func() {
		unionWith = gopurs_runtime.Func(func(dictOrd_0 gopurs_runtime.Value) gopurs_runtime.Value {
compare_1_0 := gopurs_runtime.RecordGet(dictOrd_0, "compare")
_ = compare_1_0
return gopurs_runtime.Func3(func(app_2 gopurs_runtime.Value, m1_3 gopurs_runtime.Value, m2_4 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Map_Internal.Call_unsafeUnionWith(compare_1_0, app_2, m1_3, m2_4)
})
})
	})
	return unionWith
}

var union gopurs_runtime.Value
var once_union sync.Once
func Get_union() gopurs_runtime.Value {
	once_union.Do(func() {
		union = gopurs_runtime.Func(func(dictOrd_0 gopurs_runtime.Value) gopurs_runtime.Value {
compare_1_0 := gopurs_runtime.RecordGet(dictOrd_0, "compare")
_ = compare_1_0
return gopurs_runtime.Func2(func(m1_2 gopurs_runtime.Value, m2_3 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Map_Internal.Call_unsafeUnionWith(compare_1_0, pkg_Data_Function.Get_const_(), m1_2, m2_3)
})
})
	})
	return union
}

var update gopurs_runtime.Value
var once_update sync.Once
func Get_update() gopurs_runtime.Value {
	once_update.Do(func() {
		update = gopurs_runtime.Func3(func(dictOrd_0 gopurs_runtime.Value, f_1 gopurs_runtime.Value, k_2 gopurs_runtime.Value) gopurs_runtime.Value {
var go__3_0 gopurs_runtime.Value
_ = go__3_0
go__3_0 = gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if gopurs_runtime.Bool(v_4.StrVal == "Leaf").IntVal != 0 {
__t1 = gopurs_runtime.Constructor0("Leaf")
goto end_branch_1
} else {

}
}
{
if gopurs_runtime.Bool(v_4.StrVal == "Node").IntVal != 0 {
v1_5_2 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictOrd_0, "compare"), k_2, (*[1024]gopurs_runtime.Value)(v_4.UnsafePtr)[2])
_ = v1_5_2
var __t3 gopurs_runtime.Value
{
if gopurs_runtime.Bool(v1_5_2.StrVal == "LT").IntVal != 0 {
__t3 = pkg_Data_Map_Internal.Call_unsafeBalancedNode((*[1024]gopurs_runtime.Value)(v_4.UnsafePtr)[2], (*[1024]gopurs_runtime.Value)(v_4.UnsafePtr)[3], gopurs_runtime.Apply(go__3_0, (*[1024]gopurs_runtime.Value)(v_4.UnsafePtr)[4]), (*[1024]gopurs_runtime.Value)(v_4.UnsafePtr)[5])
goto end_branch_3
} else {

}
}
{
if gopurs_runtime.Bool(v1_5_2.StrVal == "GT").IntVal != 0 {
__t3 = pkg_Data_Map_Internal.Call_unsafeBalancedNode((*[1024]gopurs_runtime.Value)(v_4.UnsafePtr)[2], (*[1024]gopurs_runtime.Value)(v_4.UnsafePtr)[3], (*[1024]gopurs_runtime.Value)(v_4.UnsafePtr)[4], gopurs_runtime.Apply(go__3_0, (*[1024]gopurs_runtime.Value)(v_4.UnsafePtr)[5]))
goto end_branch_3
} else {

}
}
{
if gopurs_runtime.Bool(v1_5_2.StrVal == "EQ").IntVal != 0 {
v2_6_4 := gopurs_runtime.Apply(f_1, (*[1024]gopurs_runtime.Value)(v_4.UnsafePtr)[3])
_ = v2_6_4
var __t5 gopurs_runtime.Value
{
if gopurs_runtime.Bool(v2_6_4.StrVal == "Nothing").IntVal != 0 {
__t5 = pkg_Data_Map_Internal.Call_unsafeJoinNodes((*[1024]gopurs_runtime.Value)(v_4.UnsafePtr)[4], (*[1024]gopurs_runtime.Value)(v_4.UnsafePtr)[5])
goto end_branch_5
} else {

}
}
{
if gopurs_runtime.Bool(v2_6_4.StrVal == "Just").IntVal != 0 {
__t5 = gopurs_runtime.Constructor("Node", []gopurs_runtime.Value{(*[1024]gopurs_runtime.Value)(v_4.UnsafePtr)[0], (*[1024]gopurs_runtime.Value)(v_4.UnsafePtr)[1], (*[1024]gopurs_runtime.Value)(v_4.UnsafePtr)[2], (*[1024]gopurs_runtime.Value)(v2_6_4.UnsafePtr)[0], (*[1024]gopurs_runtime.Value)(v_4.UnsafePtr)[4], (*[1024]gopurs_runtime.Value)(v_4.UnsafePtr)[5]})
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
})
	})
	return update
}

var showTree gopurs_runtime.Value
var once_showTree sync.Once
func Get_showTree() gopurs_runtime.Value {
	once_showTree.Do(func() {
		showTree = gopurs_runtime.Func2(func(dictShow_0 gopurs_runtime.Value, dictShow1_1 gopurs_runtime.Value) gopurs_runtime.Value {
var go__2_0 gopurs_runtime.Value
_ = go__2_0
go__2_0 = gopurs_runtime.Func2(func(ind_3 gopurs_runtime.Value, v_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if gopurs_runtime.Bool(v_4.StrVal == "Leaf").IntVal != 0 {
__t1 = gopurs_runtime.Str(ind_3.StrVal + "Leaf")
goto end_branch_1
} else {

}
}
{
if gopurs_runtime.Bool(v_4.StrVal == "Node").IntVal != 0 {
__t1 = gopurs_runtime.Str(ind_3.StrVal + "[" + gopurs_runtime.Apply(pkg_Data_Show.Get_showIntImpl(), (*[1024]gopurs_runtime.Value)(v_4.UnsafePtr)[0]).StrVal + "] " + gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictShow_0, "show"), (*[1024]gopurs_runtime.Value)(v_4.UnsafePtr)[2]).StrVal + " => " + gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictShow1_1, "show"), (*[1024]gopurs_runtime.Value)(v_4.UnsafePtr)[3]).StrVal + "\n" + gopurs_runtime.Apply2(go__2_0, gopurs_runtime.Str(ind_3.StrVal + "    "), (*[1024]gopurs_runtime.Value)(v_4.UnsafePtr)[4]).StrVal + "\n" + gopurs_runtime.Apply2(go__2_0, gopurs_runtime.Str(ind_3.StrVal + "    "), (*[1024]gopurs_runtime.Value)(v_4.UnsafePtr)[5]).StrVal)
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
})
	})
	return showTree
}

var semigroupMap gopurs_runtime.Value
var once_semigroupMap sync.Once
func Get_semigroupMap() gopurs_runtime.Value {
	once_semigroupMap.Do(func() {
		semigroupMap = gopurs_runtime.Func2(func(_dollar__unused_0 gopurs_runtime.Value, dictOrd_1 gopurs_runtime.Value) gopurs_runtime.Value {
compare_2_0 := gopurs_runtime.RecordGet(dictOrd_1, "compare")
_ = compare_2_0
return gopurs_runtime.Func(func(dictSemigroup_3 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_4_1 := gopurs_runtime.RecordGet(dictSemigroup_3, "append")
_ = __local_var_4_1
return gopurs_runtime.RecordDict1("append", gopurs_runtime.Func2(func(m1_5 gopurs_runtime.Value, m2_6 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Map_Internal.Call_unsafeUnionWith(compare_2_0, __local_var_4_1, m1_5, m2_6)
}))
})
})
	})
	return semigroupMap
}

var pop gopurs_runtime.Value
var once_pop sync.Once
func Get_pop() gopurs_runtime.Value {
	once_pop.Do(func() {
		pop = gopurs_runtime.Func(func(dictOrd_0 gopurs_runtime.Value) gopurs_runtime.Value {
compare_1_0 := gopurs_runtime.RecordGet(dictOrd_0, "compare")
_ = compare_1_0
return gopurs_runtime.Func2(func(k_2 gopurs_runtime.Value, m_3 gopurs_runtime.Value) gopurs_runtime.Value {
v_4_1 := pkg_Data_Map_Internal.Call_unsafeSplit(compare_1_0, k_2, m_3)
_ = v_4_1
var __t2 gopurs_runtime.Value
{
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)(v_4_1.UnsafePtr)[0].StrVal == "Just").IntVal != 0 {
__t2 = gopurs_runtime.Constructor1("Just", gopurs_runtime.Constructor2("Tuple", (*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)(v_4_1.UnsafePtr)[0].UnsafePtr)[0], pkg_Data_Map_Internal.Call_unsafeJoinNodes((*[1024]gopurs_runtime.Value)(v_4_1.UnsafePtr)[1], (*[1024]gopurs_runtime.Value)(v_4_1.UnsafePtr)[2])))
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.Constructor0("Nothing")
}
end_branch_2:
return __t2
})
})
	})
	return pop
}

var member gopurs_runtime.Value
var once_member sync.Once
func Get_member() gopurs_runtime.Value {
	once_member.Do(func() {
		member = gopurs_runtime.Func2(func(dictOrd_0 gopurs_runtime.Value, k_1 gopurs_runtime.Value) gopurs_runtime.Value {
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
if gopurs_runtime.Bool(v_3.StrVal == "Leaf").IntVal != 0 {
__t1 = gopurs_runtime.Boolean(false)
goto end_branch_1
} else {

}
}
{
if gopurs_runtime.Bool(v_3.StrVal == "Node").IntVal != 0 {
v1_4_2 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictOrd_0, "compare"), k_1, (*[1024]gopurs_runtime.Value)(v_3.UnsafePtr)[2])
_ = v1_4_2
var __t3 gopurs_runtime.Value
{
if gopurs_runtime.Bool(v1_4_2.StrVal == "LT").IntVal != 0 {
v_3_loop = (*[1024]gopurs_runtime.Value)(v_3.UnsafePtr)[4]
continue go__2_0
__t3 = gopurs_runtime.Value{}
goto end_branch_3
} else {

}
}
{
if gopurs_runtime.Bool(v1_4_2.StrVal == "GT").IntVal != 0 {
v_3_loop = (*[1024]gopurs_runtime.Value)(v_3.UnsafePtr)[5]
continue go__2_0
__t3 = gopurs_runtime.Value{}
goto end_branch_3
} else {

}
}
{
if gopurs_runtime.Bool(v1_4_2.StrVal == "EQ").IntVal != 0 {
__t3 = gopurs_runtime.Boolean(true)
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
})
	})
	return member
}

var mapMaybeWithKey gopurs_runtime.Value
var once_mapMaybeWithKey sync.Once
func Get_mapMaybeWithKey() gopurs_runtime.Value {
	once_mapMaybeWithKey.Do(func() {
		mapMaybeWithKey = gopurs_runtime.Func2(func(dictOrd_0 gopurs_runtime.Value, f_1 gopurs_runtime.Value) gopurs_runtime.Value {
var go__2_0 gopurs_runtime.Value
_ = go__2_0
go__2_0 = gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if gopurs_runtime.Bool(v_3.StrVal == "Leaf").IntVal != 0 {
__t1 = gopurs_runtime.Constructor0("Leaf")
goto end_branch_1
} else {

}
}
{
if gopurs_runtime.Bool(v_3.StrVal == "Node").IntVal != 0 {
v2_4_2 := gopurs_runtime.Apply2(f_1, (*[1024]gopurs_runtime.Value)(v_3.UnsafePtr)[2], (*[1024]gopurs_runtime.Value)(v_3.UnsafePtr)[3])
_ = v2_4_2
var __t3 gopurs_runtime.Value
{
if gopurs_runtime.Bool(v2_4_2.StrVal == "Just").IntVal != 0 {
__t3 = pkg_Data_Map_Internal.Call_unsafeBalancedNode((*[1024]gopurs_runtime.Value)(v_3.UnsafePtr)[2], (*[1024]gopurs_runtime.Value)(v2_4_2.UnsafePtr)[0], gopurs_runtime.Apply(go__2_0, (*[1024]gopurs_runtime.Value)(v_3.UnsafePtr)[4]), gopurs_runtime.Apply(go__2_0, (*[1024]gopurs_runtime.Value)(v_3.UnsafePtr)[5]))
goto end_branch_3
} else {

}
}
{
if gopurs_runtime.Bool(v2_4_2.StrVal == "Nothing").IntVal != 0 {
__t3 = pkg_Data_Map_Internal.Call_unsafeJoinNodes(gopurs_runtime.Apply(go__2_0, (*[1024]gopurs_runtime.Value)(v_3.UnsafePtr)[4]), gopurs_runtime.Apply(go__2_0, (*[1024]gopurs_runtime.Value)(v_3.UnsafePtr)[5]))
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
})
	})
	return mapMaybeWithKey
}

var mapMaybe gopurs_runtime.Value
var once_mapMaybe sync.Once
func Get_mapMaybe() gopurs_runtime.Value {
	once_mapMaybe.Do(func() {
		mapMaybe = gopurs_runtime.Func2(func(dictOrd_0 gopurs_runtime.Value, x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Get_mapMaybeWithKey(), dictOrd_0, gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return x_1
}))
})
	})
	return mapMaybe
}

var lookupLE gopurs_runtime.Value
var once_lookupLE sync.Once
func Get_lookupLE() gopurs_runtime.Value {
	once_lookupLE.Do(func() {
		lookupLE = gopurs_runtime.Func2(func(dictOrd_0 gopurs_runtime.Value, k_1 gopurs_runtime.Value) gopurs_runtime.Value {
var go__2_0 gopurs_runtime.Value
_ = go__2_0
go__2_0 = gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if gopurs_runtime.Bool(v_3.StrVal == "Leaf").IntVal != 0 {
__t1 = gopurs_runtime.Constructor0("Nothing")
goto end_branch_1
} else {

}
}
{
if gopurs_runtime.Bool(v_3.StrVal == "Node").IntVal != 0 {
v1_4_2 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictOrd_0, "compare"), k_1, (*[1024]gopurs_runtime.Value)(v_3.UnsafePtr)[2])
_ = v1_4_2
var __t3 gopurs_runtime.Value
{
if gopurs_runtime.Bool(v1_4_2.StrVal == "LT").IntVal != 0 {
__t3 = gopurs_runtime.Apply(go__2_0, (*[1024]gopurs_runtime.Value)(v_3.UnsafePtr)[4])
goto end_branch_3
} else {

}
}
{
if gopurs_runtime.Bool(v1_4_2.StrVal == "GT").IntVal != 0 {
v2_5_4 := gopurs_runtime.Apply(go__2_0, (*[1024]gopurs_runtime.Value)(v_3.UnsafePtr)[5])
_ = v2_5_4
var __t5 gopurs_runtime.Value
{
if gopurs_runtime.Bool(v2_5_4.StrVal == "Nothing").IntVal != 0 {
__t5 = gopurs_runtime.Constructor1("Just", gopurs_runtime.RecordDict2("key", "value", (*[1024]gopurs_runtime.Value)(v_3.UnsafePtr)[2], (*[1024]gopurs_runtime.Value)(v_3.UnsafePtr)[3]))
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
if gopurs_runtime.Bool(v1_4_2.StrVal == "EQ").IntVal != 0 {
__t3 = gopurs_runtime.Constructor1("Just", gopurs_runtime.RecordDict2("key", "value", (*[1024]gopurs_runtime.Value)(v_3.UnsafePtr)[2], (*[1024]gopurs_runtime.Value)(v_3.UnsafePtr)[3]))
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
})
	})
	return lookupLE
}

var lookupGE gopurs_runtime.Value
var once_lookupGE sync.Once
func Get_lookupGE() gopurs_runtime.Value {
	once_lookupGE.Do(func() {
		lookupGE = gopurs_runtime.Func2(func(dictOrd_0 gopurs_runtime.Value, k_1 gopurs_runtime.Value) gopurs_runtime.Value {
var go__2_0 gopurs_runtime.Value
_ = go__2_0
go__2_0 = gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if gopurs_runtime.Bool(v_3.StrVal == "Leaf").IntVal != 0 {
__t1 = gopurs_runtime.Constructor0("Nothing")
goto end_branch_1
} else {

}
}
{
if gopurs_runtime.Bool(v_3.StrVal == "Node").IntVal != 0 {
v1_4_2 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictOrd_0, "compare"), k_1, (*[1024]gopurs_runtime.Value)(v_3.UnsafePtr)[2])
_ = v1_4_2
var __t3 gopurs_runtime.Value
{
if gopurs_runtime.Bool(v1_4_2.StrVal == "LT").IntVal != 0 {
v2_5_4 := gopurs_runtime.Apply(go__2_0, (*[1024]gopurs_runtime.Value)(v_3.UnsafePtr)[4])
_ = v2_5_4
var __t5 gopurs_runtime.Value
{
if gopurs_runtime.Bool(v2_5_4.StrVal == "Nothing").IntVal != 0 {
__t5 = gopurs_runtime.Constructor1("Just", gopurs_runtime.RecordDict2("key", "value", (*[1024]gopurs_runtime.Value)(v_3.UnsafePtr)[2], (*[1024]gopurs_runtime.Value)(v_3.UnsafePtr)[3]))
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
if gopurs_runtime.Bool(v1_4_2.StrVal == "GT").IntVal != 0 {
__t3 = gopurs_runtime.Apply(go__2_0, (*[1024]gopurs_runtime.Value)(v_3.UnsafePtr)[5])
goto end_branch_3
} else {

}
}
{
if gopurs_runtime.Bool(v1_4_2.StrVal == "EQ").IntVal != 0 {
__t3 = gopurs_runtime.Constructor1("Just", gopurs_runtime.RecordDict2("key", "value", (*[1024]gopurs_runtime.Value)(v_3.UnsafePtr)[2], (*[1024]gopurs_runtime.Value)(v_3.UnsafePtr)[3]))
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
})
	})
	return lookupGE
}

var lookup gopurs_runtime.Value
var once_lookup sync.Once
func Get_lookup() gopurs_runtime.Value {
	once_lookup.Do(func() {
		lookup = gopurs_runtime.Func2(func(dictOrd_0 gopurs_runtime.Value, k_1 gopurs_runtime.Value) gopurs_runtime.Value {
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
if gopurs_runtime.Bool(v_3.StrVal == "Leaf").IntVal != 0 {
__t1 = gopurs_runtime.Constructor0("Nothing")
goto end_branch_1
} else {

}
}
{
if gopurs_runtime.Bool(v_3.StrVal == "Node").IntVal != 0 {
v1_4_2 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictOrd_0, "compare"), k_1, (*[1024]gopurs_runtime.Value)(v_3.UnsafePtr)[2])
_ = v1_4_2
var __t3 gopurs_runtime.Value
{
if gopurs_runtime.Bool(v1_4_2.StrVal == "LT").IntVal != 0 {
v_3_loop = (*[1024]gopurs_runtime.Value)(v_3.UnsafePtr)[4]
continue go__2_0
__t3 = gopurs_runtime.Value{}
goto end_branch_3
} else {

}
}
{
if gopurs_runtime.Bool(v1_4_2.StrVal == "GT").IntVal != 0 {
v_3_loop = (*[1024]gopurs_runtime.Value)(v_3.UnsafePtr)[5]
continue go__2_0
__t3 = gopurs_runtime.Value{}
goto end_branch_3
} else {

}
}
{
if gopurs_runtime.Bool(v1_4_2.StrVal == "EQ").IntVal != 0 {
__t3 = gopurs_runtime.Constructor1("Just", (*[1024]gopurs_runtime.Value)(v_3.UnsafePtr)[3])
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
})
	})
	return lookup
}

var iterMapU gopurs_runtime.Value
var once_iterMapU sync.Once
func Get_iterMapU() gopurs_runtime.Value {
	once_iterMapU.Do(func() {
		iterMapU = gopurs_runtime.Func2(func(iter_0 gopurs_runtime.Value, v_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if gopurs_runtime.Bool(v_1.StrVal == "Leaf").IntVal != 0 {
__t0 = iter_0
goto end_branch_0
} else {

}
}
{
if gopurs_runtime.Bool(v_1.StrVal == "Node").IntVal != 0 {
var __t1 gopurs_runtime.Value
{
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)(v_1.UnsafePtr)[4].StrVal == "Leaf").IntVal != 0 {
var __t2 gopurs_runtime.Value
{
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)(v_1.UnsafePtr)[5].StrVal == "Leaf").IntVal != 0 {
__t2 = gopurs_runtime.Constructor3("IterEmit", (*[1024]gopurs_runtime.Value)(v_1.UnsafePtr)[2], (*[1024]gopurs_runtime.Value)(v_1.UnsafePtr)[3], iter_0)
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.Constructor3("IterEmit", (*[1024]gopurs_runtime.Value)(v_1.UnsafePtr)[2], (*[1024]gopurs_runtime.Value)(v_1.UnsafePtr)[3], gopurs_runtime.Constructor2("IterNode", (*[1024]gopurs_runtime.Value)(v_1.UnsafePtr)[5], iter_0))
}
end_branch_2:
__t1 = __t2
goto end_branch_1
} else {

}
}
{
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)(v_1.UnsafePtr)[5].StrVal == "Leaf").IntVal != 0 {
__t1 = gopurs_runtime.Constructor3("IterEmit", (*[1024]gopurs_runtime.Value)(v_1.UnsafePtr)[2], (*[1024]gopurs_runtime.Value)(v_1.UnsafePtr)[3], gopurs_runtime.Constructor2("IterNode", (*[1024]gopurs_runtime.Value)(v_1.UnsafePtr)[4], iter_0))
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Constructor3("IterEmit", (*[1024]gopurs_runtime.Value)(v_1.UnsafePtr)[2], (*[1024]gopurs_runtime.Value)(v_1.UnsafePtr)[3], gopurs_runtime.Constructor2("IterNode", (*[1024]gopurs_runtime.Value)(v_1.UnsafePtr)[4], gopurs_runtime.Constructor2("IterNode", (*[1024]gopurs_runtime.Value)(v_1.UnsafePtr)[5], iter_0)))
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
		stepUnfoldrUnordered = gopurs_runtime.Apply3(Get_stepWith(), Get_iterMapU(), gopurs_runtime.Func3(func(k_0 gopurs_runtime.Value, v_1 gopurs_runtime.Value, next_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Constructor1("Just", gopurs_runtime.Constructor2("Tuple", gopurs_runtime.Constructor2("Tuple", k_0, v_1), next_2))
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Constructor0("Nothing")
}))
	})
	return stepUnfoldrUnordered
}

var toUnfoldableUnordered gopurs_runtime.Value
var once_toUnfoldableUnordered sync.Once
func Get_toUnfoldableUnordered() gopurs_runtime.Value {
	once_toUnfoldableUnordered.Do(func() {
		toUnfoldableUnordered = gopurs_runtime.Func(func(dictUnfoldable_0 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictUnfoldable_0, "unfoldr"), Get_stepUnfoldrUnordered())
_ = __local_var_1_0
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_1_0, gopurs_runtime.Constructor2("IterNode", x_2, gopurs_runtime.Constructor0("IterLeaf")))
})
})
	})
	return toUnfoldableUnordered
}

var stepUnordered gopurs_runtime.Value
var once_stepUnordered sync.Once
func Get_stepUnordered() gopurs_runtime.Value {
	once_stepUnordered.Do(func() {
		stepUnordered = gopurs_runtime.Apply3(Get_stepWith(), Get_iterMapU(), gopurs_runtime.Func3(func(k_0 gopurs_runtime.Value, v_1 gopurs_runtime.Value, next_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Constructor3("IterNext", k_0, v_1, next_2)
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Constructor0("IterDone")
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
if gopurs_runtime.Bool(v_2.StrVal == "Leaf").IntVal != 0 {
__t1 = iter_1
goto end_branch_1
} else {

}
}
{
if gopurs_runtime.Bool(v_2.StrVal == "Node").IntVal != 0 {
var __t2 gopurs_runtime.Value
{
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)(v_2.UnsafePtr)[5].StrVal == "Leaf").IntVal != 0 {
iter_1_loop = gopurs_runtime.Constructor3("IterEmit", (*[1024]gopurs_runtime.Value)(v_2.UnsafePtr)[2], (*[1024]gopurs_runtime.Value)(v_2.UnsafePtr)[3], iter_1)
v_2_loop = (*[1024]gopurs_runtime.Value)(v_2.UnsafePtr)[4]
continue go__0_0
__t2 = gopurs_runtime.Value{}
goto end_branch_2
} else {

}
}
{
iter_1_loop = gopurs_runtime.Constructor3("IterEmit", (*[1024]gopurs_runtime.Value)(v_2.UnsafePtr)[2], (*[1024]gopurs_runtime.Value)(v_2.UnsafePtr)[3], gopurs_runtime.Constructor2("IterNode", (*[1024]gopurs_runtime.Value)(v_2.UnsafePtr)[4], iter_1))
v_2_loop = (*[1024]gopurs_runtime.Value)(v_2.UnsafePtr)[5]
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
		stepDesc = gopurs_runtime.Apply3(Get_stepWith(), Get_iterMapR(), gopurs_runtime.Func3(func(k_0 gopurs_runtime.Value, v_1 gopurs_runtime.Value, next_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Constructor3("IterNext", k_0, v_1, next_2)
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Constructor0("IterDone")
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
if gopurs_runtime.Bool(v_2.StrVal == "Leaf").IntVal != 0 {
__t1 = iter_1
goto end_branch_1
} else {

}
}
{
if gopurs_runtime.Bool(v_2.StrVal == "Node").IntVal != 0 {
var __t2 gopurs_runtime.Value
{
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)(v_2.UnsafePtr)[5].StrVal == "Leaf").IntVal != 0 {
iter_1_loop = gopurs_runtime.Constructor3("IterEmit", (*[1024]gopurs_runtime.Value)(v_2.UnsafePtr)[2], (*[1024]gopurs_runtime.Value)(v_2.UnsafePtr)[3], iter_1)
v_2_loop = (*[1024]gopurs_runtime.Value)(v_2.UnsafePtr)[4]
continue go__0_0
__t2 = gopurs_runtime.Value{}
goto end_branch_2
} else {

}
}
{
iter_1_loop = gopurs_runtime.Constructor3("IterEmit", (*[1024]gopurs_runtime.Value)(v_2.UnsafePtr)[2], (*[1024]gopurs_runtime.Value)(v_2.UnsafePtr)[3], gopurs_runtime.Constructor2("IterNode", (*[1024]gopurs_runtime.Value)(v_2.UnsafePtr)[5], iter_1))
v_2_loop = (*[1024]gopurs_runtime.Value)(v_2.UnsafePtr)[4]
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
		stepAsc = gopurs_runtime.Apply3(Get_stepWith(), Get_iterMapL(), gopurs_runtime.Func3(func(k_0 gopurs_runtime.Value, v_1 gopurs_runtime.Value, next_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Constructor3("IterNext", k_0, v_1, next_2)
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Constructor0("IterDone")
}))
	})
	return stepAsc
}

var eqMapIter gopurs_runtime.Value
var once_eqMapIter sync.Once
func Get_eqMapIter() gopurs_runtime.Value {
	once_eqMapIter.Do(func() {
		eqMapIter = gopurs_runtime.Func2(func(dictEq_0 gopurs_runtime.Value, dictEq1_1 gopurs_runtime.Value) gopurs_runtime.Value {
var go__2_0 gopurs_runtime.Value
_ = go__2_0
go__2_0 = gopurs_runtime.Func2(func(a_3 gopurs_runtime.Value, b_4 gopurs_runtime.Value) gopurs_runtime.Value {
v_5_1 := gopurs_runtime.Apply(Get_stepAsc(), a_3)
_ = v_5_1
var __t2 gopurs_runtime.Value
{
if gopurs_runtime.Bool(v_5_1.StrVal == "IterNext").IntVal != 0 {
v2_6_3 := gopurs_runtime.Apply(Get_stepAsc(), b_4)
_ = v2_6_3
__t2 = gopurs_runtime.Boolean(gopurs_runtime.Bool(v2_6_3.StrVal == "IterNext").IntVal != 0 && gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq_0, "eq"), (*[1024]gopurs_runtime.Value)(v_5_1.UnsafePtr)[0], (*[1024]gopurs_runtime.Value)(v2_6_3.UnsafePtr)[0]).IntVal != 0 && gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq1_1, "eq"), (*[1024]gopurs_runtime.Value)(v_5_1.UnsafePtr)[1], (*[1024]gopurs_runtime.Value)(v2_6_3.UnsafePtr)[1]).IntVal != 0 && gopurs_runtime.Apply2(go__2_0, (*[1024]gopurs_runtime.Value)(v_5_1.UnsafePtr)[2], (*[1024]gopurs_runtime.Value)(v2_6_3.UnsafePtr)[2]).IntVal != 0)
goto end_branch_2
} else {

}
}
{
if gopurs_runtime.Bool(v_5_1.StrVal == "IterDone").IntVal != 0 {
__t2 = gopurs_runtime.Boolean(true)
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
})
	})
	return eqMapIter
}

var ordMapIter gopurs_runtime.Value
var once_ordMapIter sync.Once
func Get_ordMapIter() gopurs_runtime.Value {
	once_ordMapIter.Do(func() {
		ordMapIter = gopurs_runtime.Func(func(dictOrd_0 gopurs_runtime.Value) gopurs_runtime.Value {
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
if gopurs_runtime.Bool(v1_8_4.StrVal == "IterNext").IntVal != 0 {
var __t6 gopurs_runtime.Value
{
if gopurs_runtime.Bool(v_7_3.StrVal == "IterNext").IntVal != 0 {
v3_9_7 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictOrd_0, "compare"), (*[1024]gopurs_runtime.Value)(v1_8_4.UnsafePtr)[0], (*[1024]gopurs_runtime.Value)(v_7_3.UnsafePtr)[0])
_ = v3_9_7
var __t8 gopurs_runtime.Value
{
if gopurs_runtime.Bool(v3_9_7.StrVal == "EQ").IntVal != 0 {
v4_10_9 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictOrd1_2, "compare"), (*[1024]gopurs_runtime.Value)(v1_8_4.UnsafePtr)[1], (*[1024]gopurs_runtime.Value)(v_7_3.UnsafePtr)[1])
_ = v4_10_9
var __t10 gopurs_runtime.Value
{
if gopurs_runtime.Bool(v4_10_9.StrVal == "EQ").IntVal != 0 {
a_5_loop = (*[1024]gopurs_runtime.Value)(v1_8_4.UnsafePtr)[2]
b_6_loop = (*[1024]gopurs_runtime.Value)(v_7_3.UnsafePtr)[2]
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
if gopurs_runtime.Bool(v_7_3.StrVal == "IterDone").IntVal != 0 {
__t6 = gopurs_runtime.Constructor0("GT")
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
if gopurs_runtime.Bool(v1_8_4.StrVal == "IterDone").IntVal != 0 {
var __t11 gopurs_runtime.Value
{
if gopurs_runtime.Bool(v_7_3.StrVal == "IterDone").IntVal != 0 {
__t11 = gopurs_runtime.Constructor0("EQ")
goto end_branch_11
} else {

}
}
{
__t11 = gopurs_runtime.Constructor0("LT")
}
end_branch_11:
__t5 = __t11
goto end_branch_5
} else {

}
}
{
if gopurs_runtime.Bool(v_7_3.StrVal == "IterDone").IntVal != 0 {
__t5 = gopurs_runtime.Constructor0("GT")
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
})
	})
	return ordMapIter
}

var stepUnfoldr gopurs_runtime.Value
var once_stepUnfoldr sync.Once
func Get_stepUnfoldr() gopurs_runtime.Value {
	once_stepUnfoldr.Do(func() {
		stepUnfoldr = gopurs_runtime.Apply3(Get_stepWith(), Get_iterMapL(), gopurs_runtime.Func3(func(k_0 gopurs_runtime.Value, v_1 gopurs_runtime.Value, next_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Constructor1("Just", gopurs_runtime.Constructor2("Tuple", gopurs_runtime.Constructor2("Tuple", k_0, v_1), next_2))
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Constructor0("Nothing")
}))
	})
	return stepUnfoldr
}

var toUnfoldable gopurs_runtime.Value
var once_toUnfoldable sync.Once
func Get_toUnfoldable() gopurs_runtime.Value {
	once_toUnfoldable.Do(func() {
		toUnfoldable = gopurs_runtime.Func(func(dictUnfoldable_0 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictUnfoldable_0, "unfoldr"), Get_stepUnfoldr())
_ = __local_var_1_0
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_1_0, gopurs_runtime.Constructor2("IterNode", x_2, gopurs_runtime.Constructor0("IterLeaf")))
})
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
return gopurs_runtime.Apply(__local_var_0_0, gopurs_runtime.Constructor2("IterNode", x_1, gopurs_runtime.Constructor0("IterLeaf")))
})
}()
	})
	return toUnfoldable1
}

var showMap gopurs_runtime.Value
var once_showMap sync.Once
func Get_showMap() gopurs_runtime.Value {
	once_showMap.Do(func() {
		showMap = gopurs_runtime.Func2(func(dictShow_0 gopurs_runtime.Value, dictShow1_1 gopurs_runtime.Value) gopurs_runtime.Value {
show1_2_0 := gopurs_runtime.Apply(pkg_Data_Show.Get_showArrayImpl(), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str("(Tuple " + gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictShow_0, "show"), (*[1024]gopurs_runtime.Value)(v_2.UnsafePtr)[0]).StrVal + " " + gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictShow1_1, "show"), (*[1024]gopurs_runtime.Value)(v_2.UnsafePtr)[1]).StrVal + ")")
}))
_ = show1_2_0
return gopurs_runtime.RecordDict1("show", gopurs_runtime.Func(func(as_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str("(fromFoldable " + gopurs_runtime.Apply(show1_2_0, gopurs_runtime.Apply(Get_toUnfoldable1(), as_3)).StrVal + ")")
}))
})
	})
	return showMap
}

var isSubmap gopurs_runtime.Value
var once_isSubmap sync.Once
func Get_isSubmap() gopurs_runtime.Value {
	once_isSubmap.Do(func() {
		isSubmap = gopurs_runtime.Func2(func(dictOrd_0 gopurs_runtime.Value, dictEq_1 gopurs_runtime.Value) gopurs_runtime.Value {
var go__2_0 gopurs_runtime.Value
_ = go__2_0
go__2_0 = gopurs_runtime.Func2(func(m1_3 gopurs_runtime.Value, m2_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if gopurs_runtime.Bool(m1_3.StrVal == "Leaf").IntVal != 0 {
__t1 = gopurs_runtime.Boolean(true)
goto end_branch_1
} else {

}
}
{
if gopurs_runtime.Bool(m1_3.StrVal == "Node").IntVal != 0 {
__local_var_5_2 := (*[1024]gopurs_runtime.Value)(m1_3.UnsafePtr)[2]
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
if gopurs_runtime.Bool(v_7.StrVal == "Leaf").IntVal != 0 {
__t4 = gopurs_runtime.Constructor0("Nothing")
goto end_branch_4
} else {

}
}
{
if gopurs_runtime.Bool(v_7.StrVal == "Node").IntVal != 0 {
v1_8_5 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictOrd_0, "compare"), __local_var_5_2, (*[1024]gopurs_runtime.Value)(v_7.UnsafePtr)[2])
_ = v1_8_5
var __t6 gopurs_runtime.Value
{
if gopurs_runtime.Bool(v1_8_5.StrVal == "LT").IntVal != 0 {
v_7_loop = (*[1024]gopurs_runtime.Value)(v_7.UnsafePtr)[4]
continue go__6_3
__t6 = gopurs_runtime.Value{}
goto end_branch_6
} else {

}
}
{
if gopurs_runtime.Bool(v1_8_5.StrVal == "GT").IntVal != 0 {
v_7_loop = (*[1024]gopurs_runtime.Value)(v_7.UnsafePtr)[5]
continue go__6_3
__t6 = gopurs_runtime.Value{}
goto end_branch_6
} else {

}
}
{
if gopurs_runtime.Bool(v1_8_5.StrVal == "EQ").IntVal != 0 {
__t6 = gopurs_runtime.Constructor1("Just", (*[1024]gopurs_runtime.Value)(v_7.UnsafePtr)[3])
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
if gopurs_runtime.Bool(v1_7_7.StrVal == "Nothing").IntVal != 0 {
__t8 = gopurs_runtime.Boolean(false)
goto end_branch_8
} else {

}
}
{
if gopurs_runtime.Bool(v1_7_7.StrVal == "Just").IntVal != 0 {
__t8 = gopurs_runtime.Boolean(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq_1, "eq"), (*[1024]gopurs_runtime.Value)(m1_3.UnsafePtr)[3], (*[1024]gopurs_runtime.Value)(v1_7_7.UnsafePtr)[0]).IntVal != 0 && gopurs_runtime.Apply2(go__2_0, (*[1024]gopurs_runtime.Value)(m1_3.UnsafePtr)[4], m2_4).IntVal != 0 && gopurs_runtime.Apply2(go__2_0, (*[1024]gopurs_runtime.Value)(m1_3.UnsafePtr)[5], m2_4).IntVal != 0)
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
})
	})
	return isSubmap
}

var isEmpty gopurs_runtime.Value
var once_isEmpty sync.Once
func Get_isEmpty() gopurs_runtime.Value {
	once_isEmpty.Do(func() {
		isEmpty = gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(v_0.StrVal == "Leaf")
})
	})
	return isEmpty
}

var intersectionWith gopurs_runtime.Value
var once_intersectionWith sync.Once
func Get_intersectionWith() gopurs_runtime.Value {
	once_intersectionWith.Do(func() {
		intersectionWith = gopurs_runtime.Func(func(dictOrd_0 gopurs_runtime.Value) gopurs_runtime.Value {
compare_1_0 := gopurs_runtime.RecordGet(dictOrd_0, "compare")
_ = compare_1_0
return gopurs_runtime.Func3(func(app_2 gopurs_runtime.Value, m1_3 gopurs_runtime.Value, m2_4 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Map_Internal.Call_unsafeIntersectionWith(compare_1_0, app_2, m1_3, m2_4)
})
})
	})
	return intersectionWith
}

var intersection gopurs_runtime.Value
var once_intersection sync.Once
func Get_intersection() gopurs_runtime.Value {
	once_intersection.Do(func() {
		intersection = gopurs_runtime.Func(func(dictOrd_0 gopurs_runtime.Value) gopurs_runtime.Value {
compare_1_0 := gopurs_runtime.RecordGet(dictOrd_0, "compare")
_ = compare_1_0
return gopurs_runtime.Func2(func(m1_2 gopurs_runtime.Value, m2_3 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Map_Internal.Call_unsafeIntersectionWith(compare_1_0, pkg_Data_Function.Get_const_(), m1_2, m2_3)
})
})
	})
	return intersection
}

var insertWith gopurs_runtime.Value
var once_insertWith sync.Once
func Get_insertWith() gopurs_runtime.Value {
	once_insertWith.Do(func() {
		insertWith = gopurs_runtime.Func4(func(dictOrd_0 gopurs_runtime.Value, app_1 gopurs_runtime.Value, k_2 gopurs_runtime.Value, v_3 gopurs_runtime.Value) gopurs_runtime.Value {
var go__4_0 gopurs_runtime.Value
_ = go__4_0
go__4_0 = gopurs_runtime.Func(func(v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if gopurs_runtime.Bool(v1_5.StrVal == "Leaf").IntVal != 0 {
__t1 = gopurs_runtime.Constructor("Node", []gopurs_runtime.Value{gopurs_runtime.Int(1), gopurs_runtime.Int(1), k_2, v_3, gopurs_runtime.Constructor0("Leaf"), gopurs_runtime.Constructor0("Leaf")})
goto end_branch_1
} else {

}
}
{
if gopurs_runtime.Bool(v1_5.StrVal == "Node").IntVal != 0 {
v2_6_2 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictOrd_0, "compare"), k_2, (*[1024]gopurs_runtime.Value)(v1_5.UnsafePtr)[2])
_ = v2_6_2
var __t3 gopurs_runtime.Value
{
if gopurs_runtime.Bool(v2_6_2.StrVal == "LT").IntVal != 0 {
__t3 = pkg_Data_Map_Internal.Call_unsafeBalancedNode((*[1024]gopurs_runtime.Value)(v1_5.UnsafePtr)[2], (*[1024]gopurs_runtime.Value)(v1_5.UnsafePtr)[3], gopurs_runtime.Apply(go__4_0, (*[1024]gopurs_runtime.Value)(v1_5.UnsafePtr)[4]), (*[1024]gopurs_runtime.Value)(v1_5.UnsafePtr)[5])
goto end_branch_3
} else {

}
}
{
if gopurs_runtime.Bool(v2_6_2.StrVal == "GT").IntVal != 0 {
__t3 = pkg_Data_Map_Internal.Call_unsafeBalancedNode((*[1024]gopurs_runtime.Value)(v1_5.UnsafePtr)[2], (*[1024]gopurs_runtime.Value)(v1_5.UnsafePtr)[3], (*[1024]gopurs_runtime.Value)(v1_5.UnsafePtr)[4], gopurs_runtime.Apply(go__4_0, (*[1024]gopurs_runtime.Value)(v1_5.UnsafePtr)[5]))
goto end_branch_3
} else {

}
}
{
if gopurs_runtime.Bool(v2_6_2.StrVal == "EQ").IntVal != 0 {
__t3 = gopurs_runtime.Constructor("Node", []gopurs_runtime.Value{(*[1024]gopurs_runtime.Value)(v1_5.UnsafePtr)[0], (*[1024]gopurs_runtime.Value)(v1_5.UnsafePtr)[1], k_2, gopurs_runtime.Apply2(app_1, (*[1024]gopurs_runtime.Value)(v1_5.UnsafePtr)[3], v_3), (*[1024]gopurs_runtime.Value)(v1_5.UnsafePtr)[4], (*[1024]gopurs_runtime.Value)(v1_5.UnsafePtr)[5]})
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
})
	})
	return insertWith
}

var insert gopurs_runtime.Value
var once_insert sync.Once
func Get_insert() gopurs_runtime.Value {
	once_insert.Do(func() {
		insert = gopurs_runtime.Func3(func(dictOrd_0 gopurs_runtime.Value, k_1 gopurs_runtime.Value, v_2 gopurs_runtime.Value) gopurs_runtime.Value {
var go__3_0 gopurs_runtime.Value
_ = go__3_0
go__3_0 = gopurs_runtime.Func(func(v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if gopurs_runtime.Bool(v1_4.StrVal == "Leaf").IntVal != 0 {
__t1 = gopurs_runtime.Constructor("Node", []gopurs_runtime.Value{gopurs_runtime.Int(1), gopurs_runtime.Int(1), k_1, v_2, gopurs_runtime.Constructor0("Leaf"), gopurs_runtime.Constructor0("Leaf")})
goto end_branch_1
} else {

}
}
{
if gopurs_runtime.Bool(v1_4.StrVal == "Node").IntVal != 0 {
v2_5_2 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictOrd_0, "compare"), k_1, (*[1024]gopurs_runtime.Value)(v1_4.UnsafePtr)[2])
_ = v2_5_2
var __t3 gopurs_runtime.Value
{
if gopurs_runtime.Bool(v2_5_2.StrVal == "LT").IntVal != 0 {
__t3 = pkg_Data_Map_Internal.Call_unsafeBalancedNode((*[1024]gopurs_runtime.Value)(v1_4.UnsafePtr)[2], (*[1024]gopurs_runtime.Value)(v1_4.UnsafePtr)[3], gopurs_runtime.Apply(go__3_0, (*[1024]gopurs_runtime.Value)(v1_4.UnsafePtr)[4]), (*[1024]gopurs_runtime.Value)(v1_4.UnsafePtr)[5])
goto end_branch_3
} else {

}
}
{
if gopurs_runtime.Bool(v2_5_2.StrVal == "GT").IntVal != 0 {
__t3 = pkg_Data_Map_Internal.Call_unsafeBalancedNode((*[1024]gopurs_runtime.Value)(v1_4.UnsafePtr)[2], (*[1024]gopurs_runtime.Value)(v1_4.UnsafePtr)[3], (*[1024]gopurs_runtime.Value)(v1_4.UnsafePtr)[4], gopurs_runtime.Apply(go__3_0, (*[1024]gopurs_runtime.Value)(v1_4.UnsafePtr)[5]))
goto end_branch_3
} else {

}
}
{
if gopurs_runtime.Bool(v2_5_2.StrVal == "EQ").IntVal != 0 {
__t3 = gopurs_runtime.Constructor("Node", []gopurs_runtime.Value{(*[1024]gopurs_runtime.Value)(v1_4.UnsafePtr)[0], (*[1024]gopurs_runtime.Value)(v1_4.UnsafePtr)[1], k_1, v_2, (*[1024]gopurs_runtime.Value)(v1_4.UnsafePtr)[4], (*[1024]gopurs_runtime.Value)(v1_4.UnsafePtr)[5]})
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
if gopurs_runtime.Bool(v_2.StrVal == "Leaf").IntVal != 0 {
__t1 = gopurs_runtime.Constructor0("Leaf")
goto end_branch_1
} else {

}
}
{
if gopurs_runtime.Bool(v_2.StrVal == "Node").IntVal != 0 {
__t1 = gopurs_runtime.Constructor("Node", []gopurs_runtime.Value{(*[1024]gopurs_runtime.Value)(v_2.UnsafePtr)[0], (*[1024]gopurs_runtime.Value)(v_2.UnsafePtr)[1], (*[1024]gopurs_runtime.Value)(v_2.UnsafePtr)[2], gopurs_runtime.Apply(f_0, (*[1024]gopurs_runtime.Value)(v_2.UnsafePtr)[3]), gopurs_runtime.Apply(go__1_0, (*[1024]gopurs_runtime.Value)(v_2.UnsafePtr)[4]), gopurs_runtime.Apply(go__1_0, (*[1024]gopurs_runtime.Value)(v_2.UnsafePtr)[5])})
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
if gopurs_runtime.Bool(v_2.StrVal == "Leaf").IntVal != 0 {
__t1 = gopurs_runtime.Constructor0("Leaf")
goto end_branch_1
} else {

}
}
{
if gopurs_runtime.Bool(v_2.StrVal == "Node").IntVal != 0 {
__t1 = gopurs_runtime.Constructor("Node", []gopurs_runtime.Value{(*[1024]gopurs_runtime.Value)(v_2.UnsafePtr)[0], (*[1024]gopurs_runtime.Value)(v_2.UnsafePtr)[1], (*[1024]gopurs_runtime.Value)(v_2.UnsafePtr)[2], gopurs_runtime.Apply2(f_0, (*[1024]gopurs_runtime.Value)(v_2.UnsafePtr)[2], (*[1024]gopurs_runtime.Value)(v_2.UnsafePtr)[3]), gopurs_runtime.Apply(go__1_0, (*[1024]gopurs_runtime.Value)(v_2.UnsafePtr)[4]), gopurs_runtime.Apply(go__1_0, (*[1024]gopurs_runtime.Value)(v_2.UnsafePtr)[5])})
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
if gopurs_runtime.Bool(v_5.StrVal == "Leaf").IntVal != 0 {
__t7 = mempty_1_4
goto end_branch_7
} else {

}
}
{
if gopurs_runtime.Bool(v_5.StrVal == "Node").IntVal != 0 {
__t7 = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_5, "append"), gopurs_runtime.Apply(go__4_6, (*[1024]gopurs_runtime.Value)(v_5.UnsafePtr)[4]), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_5, "append"), gopurs_runtime.Apply(f_3, (*[1024]gopurs_runtime.Value)(v_5.UnsafePtr)[3]), gopurs_runtime.Apply(go__4_6, (*[1024]gopurs_runtime.Value)(v_5.UnsafePtr)[5])))
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
if gopurs_runtime.Bool(v_5.StrVal == "Leaf").IntVal != 0 {
__t7 = mempty_1_4
goto end_branch_7
} else {

}
}
{
if gopurs_runtime.Bool(v_5.StrVal == "Node").IntVal != 0 {
__t7 = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_5, "append"), gopurs_runtime.Apply(go__4_6, (*[1024]gopurs_runtime.Value)(v_5.UnsafePtr)[4]), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_5, "append"), gopurs_runtime.Apply2(f_3, (*[1024]gopurs_runtime.Value)(v_5.UnsafePtr)[2], (*[1024]gopurs_runtime.Value)(v_5.UnsafePtr)[3]), gopurs_runtime.Apply(go__4_6, (*[1024]gopurs_runtime.Value)(v_5.UnsafePtr)[5])))
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
return gopurs_runtime.UncurriedApp2(go__0_0, m_1, gopurs_runtime.Constructor0("Nil"))
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
if gopurs_runtime.Bool(v_4.StrVal == "Leaf").IntVal != 0 {
__t2 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "pure"), gopurs_runtime.Constructor0("Leaf"))
goto end_branch_2
} else {

}
}
{
if gopurs_runtime.Bool(v_4.StrVal == "Node").IntVal != 0 {
__local_var_5_3 := (*[1024]gopurs_runtime.Value)(v_4.UnsafePtr)[0]
_ = __local_var_5_3
__local_var_6_4 := (*[1024]gopurs_runtime.Value)(v_4.UnsafePtr)[2]
_ = __local_var_6_4
__local_var_7_5 := (*[1024]gopurs_runtime.Value)(v_4.UnsafePtr)[1]
_ = __local_var_7_5
__t2 = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Apply0_1_0, "apply"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Apply0_1_0, "apply"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Apply0_1_0, "Functor0"), gopurs_runtime.Value{}), "map"), gopurs_runtime.Func3(func(l_prime_8 gopurs_runtime.Value, v_prime_9 gopurs_runtime.Value, r_prime_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Constructor("Node", []gopurs_runtime.Value{__local_var_5_3, __local_var_7_5, __local_var_6_4, v_prime_9, l_prime_8, r_prime_10})
}), gopurs_runtime.Apply(go__3_1, (*[1024]gopurs_runtime.Value)(v_4.UnsafePtr)[4])), gopurs_runtime.Apply(f_2, (*[1024]gopurs_runtime.Value)(v_4.UnsafePtr)[3])), gopurs_runtime.Apply(go__3_1, (*[1024]gopurs_runtime.Value)(v_4.UnsafePtr)[5]))
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
if gopurs_runtime.Bool(v_4.StrVal == "Leaf").IntVal != 0 {
__t2 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "pure"), gopurs_runtime.Constructor0("Leaf"))
goto end_branch_2
} else {

}
}
{
if gopurs_runtime.Bool(v_4.StrVal == "Node").IntVal != 0 {
__local_var_5_3 := (*[1024]gopurs_runtime.Value)(v_4.UnsafePtr)[0]
_ = __local_var_5_3
__local_var_6_4 := (*[1024]gopurs_runtime.Value)(v_4.UnsafePtr)[2]
_ = __local_var_6_4
__local_var_7_5 := (*[1024]gopurs_runtime.Value)(v_4.UnsafePtr)[1]
_ = __local_var_7_5
__t2 = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Apply0_1_0, "apply"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Apply0_1_0, "apply"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Apply0_1_0, "Functor0"), gopurs_runtime.Value{}), "map"), gopurs_runtime.Func3(func(l_prime_8 gopurs_runtime.Value, v_prime_9 gopurs_runtime.Value, r_prime_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Constructor("Node", []gopurs_runtime.Value{__local_var_5_3, __local_var_7_5, __local_var_6_4, v_prime_9, l_prime_8, r_prime_10})
}), gopurs_runtime.Apply(go__3_1, (*[1024]gopurs_runtime.Value)(v_4.UnsafePtr)[4])), gopurs_runtime.Apply2(f_2, __local_var_6_4, (*[1024]gopurs_runtime.Value)(v_4.UnsafePtr)[3])), gopurs_runtime.Apply(go__3_1, (*[1024]gopurs_runtime.Value)(v_4.UnsafePtr)[5]))
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
return gopurs_runtime.UncurriedApp2(go__0_0, m_1, gopurs_runtime.Constructor0("Nil"))
})
}()
	})
	return values
}

var foldSubmapBy gopurs_runtime.Value
var once_foldSubmapBy sync.Once
func Get_foldSubmapBy() gopurs_runtime.Value {
	once_foldSubmapBy.Do(func() {
		foldSubmapBy = gopurs_runtime.Func5(func(dictOrd_0 gopurs_runtime.Value, appendFn_1 gopurs_runtime.Value, memptyValue_2 gopurs_runtime.Value, kmin_3 gopurs_runtime.Value, kmax_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_5 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if gopurs_runtime.Bool(kmin_3.StrVal == "Just").IntVal != 0 {
__local_var_6_2 := (*[1024]gopurs_runtime.Value)(kmin_3.UnsafePtr)[0]
_ = __local_var_6_2
__t1 = gopurs_runtime.Func(func(k_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictOrd_0, "compare"), k_7, __local_var_6_2).StrVal == "LT")
})
goto end_branch_1
} else {

}
}
{
if gopurs_runtime.Bool(kmin_3.StrVal == "Nothing").IntVal != 0 {
__t1 = gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Boolean(false)
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
if gopurs_runtime.Bool(kmax_4.StrVal == "Just").IntVal != 0 {
__local_var_7_5 := (*[1024]gopurs_runtime.Value)(kmax_4.UnsafePtr)[0]
_ = __local_var_7_5
__t4 = gopurs_runtime.Func(func(k_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictOrd_0, "compare"), k_8, __local_var_7_5).StrVal == "GT")
})
goto end_branch_4
} else {

}
}
{
if gopurs_runtime.Bool(kmax_4.StrVal == "Nothing").IntVal != 0 {
__t4 = gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Boolean(false)
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
if gopurs_runtime.Bool(kmin_3.StrVal == "Just").IntVal != 0 {
var __t8 gopurs_runtime.Value
{
if gopurs_runtime.Bool(kmax_4.StrVal == "Just").IntVal != 0 {
__local_var_8_9 := (*[1024]gopurs_runtime.Value)(kmax_4.UnsafePtr)[0]
_ = __local_var_8_9
__local_var_9_10 := (*[1024]gopurs_runtime.Value)(kmin_3.UnsafePtr)[0]
_ = __local_var_9_10
__t8 = gopurs_runtime.Func(func(k_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Boolean(gopurs_runtime.Bool(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictOrd_0, "compare"), __local_var_9_10, k_10).StrVal == "GT").IntVal != 0 != true && gopurs_runtime.Bool(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictOrd_0, "compare"), k_10, __local_var_8_9).StrVal == "GT").IntVal != 0 != true)
})
goto end_branch_8
} else {

}
}
{
if gopurs_runtime.Bool(kmax_4.StrVal == "Nothing").IntVal != 0 {
__local_var_8_11 := (*[1024]gopurs_runtime.Value)(kmin_3.UnsafePtr)[0]
_ = __local_var_8_11
__t8 = gopurs_runtime.Func(func(k_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Boolean(gopurs_runtime.Bool(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictOrd_0, "compare"), __local_var_8_11, k_9).StrVal == "GT").IntVal != 0 != true)
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
if gopurs_runtime.Bool(kmin_3.StrVal == "Nothing").IntVal != 0 {
var __t12 gopurs_runtime.Value
{
if gopurs_runtime.Bool(kmax_4.StrVal == "Just").IntVal != 0 {
__local_var_8_13 := (*[1024]gopurs_runtime.Value)(kmax_4.UnsafePtr)[0]
_ = __local_var_8_13
__t12 = gopurs_runtime.Func(func(k_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Boolean(gopurs_runtime.Bool(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictOrd_0, "compare"), k_9, __local_var_8_13).StrVal == "GT").IntVal != 0 != true)
})
goto end_branch_12
} else {

}
}
{
if gopurs_runtime.Bool(kmax_4.StrVal == "Nothing").IntVal != 0 {
__t12 = gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Boolean(true)
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
if gopurs_runtime.Bool(v_10.StrVal == "Leaf").IntVal != 0 {
__t15 = memptyValue_2
goto end_branch_15
} else {

}
}
{
if gopurs_runtime.Bool(v_10.StrVal == "Node").IntVal != 0 {
var __t16 gopurs_runtime.Value
{
if gopurs_runtime.Apply(tooSmall_6_0, (*[1024]gopurs_runtime.Value)(v_10.UnsafePtr)[2]).IntVal != 0 {
__t16 = memptyValue_2
goto end_branch_16
} else {

}
}
{
__t16 = gopurs_runtime.Apply(go__9_14, (*[1024]gopurs_runtime.Value)(v_10.UnsafePtr)[4])
}
end_branch_16:
var __t17 gopurs_runtime.Value
{
if gopurs_runtime.Apply(inBounds_8_6, (*[1024]gopurs_runtime.Value)(v_10.UnsafePtr)[2]).IntVal != 0 {
__t17 = gopurs_runtime.Apply2(f_5, (*[1024]gopurs_runtime.Value)(v_10.UnsafePtr)[2], (*[1024]gopurs_runtime.Value)(v_10.UnsafePtr)[3])
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
if gopurs_runtime.Apply(tooLarge_7_3, (*[1024]gopurs_runtime.Value)(v_10.UnsafePtr)[2]).IntVal != 0 {
__t18 = memptyValue_2
goto end_branch_18
} else {

}
}
{
__t18 = gopurs_runtime.Apply(go__9_14, (*[1024]gopurs_runtime.Value)(v_10.UnsafePtr)[5])
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
})
})
	})
	return foldSubmapBy
}

var foldSubmap gopurs_runtime.Value
var once_foldSubmap sync.Once
func Get_foldSubmap() gopurs_runtime.Value {
	once_foldSubmap.Do(func() {
		foldSubmap = gopurs_runtime.Func2(func(dictOrd_0 gopurs_runtime.Value, dictMonoid_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(Get_foldSubmapBy(), dictOrd_0, gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_1, "Semigroup0"), gopurs_runtime.Value{}), "append"), gopurs_runtime.RecordGet(dictMonoid_1, "mempty"))
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
if gopurs_runtime.Bool(v_0_loop.StrVal == "Leaf").IntVal != 0 {
__t0 = gopurs_runtime.Constructor0("Nothing")
goto end_branch_0
} else {

}
}
{
if gopurs_runtime.Bool(v_0_loop.StrVal == "Node").IntVal != 0 {
var __t1 gopurs_runtime.Value
{
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)(v_0_loop.UnsafePtr)[4].StrVal == "Leaf").IntVal != 0 {
__t1 = gopurs_runtime.Constructor1("Just", gopurs_runtime.RecordDict2("key", "value", (*[1024]gopurs_runtime.Value)(v_0_loop.UnsafePtr)[2], (*[1024]gopurs_runtime.Value)(v_0_loop.UnsafePtr)[3]))
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Apply(Get_findMin(), (*[1024]gopurs_runtime.Value)(v_0_loop.UnsafePtr)[4])
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
		lookupGT = gopurs_runtime.Func2(func(dictOrd_0 gopurs_runtime.Value, k_1 gopurs_runtime.Value) gopurs_runtime.Value {
var go__2_0 gopurs_runtime.Value
_ = go__2_0
go__2_0 = gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if gopurs_runtime.Bool(v_3.StrVal == "Leaf").IntVal != 0 {
__t1 = gopurs_runtime.Constructor0("Nothing")
goto end_branch_1
} else {

}
}
{
if gopurs_runtime.Bool(v_3.StrVal == "Node").IntVal != 0 {
v1_4_2 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictOrd_0, "compare"), k_1, (*[1024]gopurs_runtime.Value)(v_3.UnsafePtr)[2])
_ = v1_4_2
var __t3 gopurs_runtime.Value
{
if gopurs_runtime.Bool(v1_4_2.StrVal == "LT").IntVal != 0 {
v2_5_4 := gopurs_runtime.Apply(go__2_0, (*[1024]gopurs_runtime.Value)(v_3.UnsafePtr)[4])
_ = v2_5_4
var __t5 gopurs_runtime.Value
{
if gopurs_runtime.Bool(v2_5_4.StrVal == "Nothing").IntVal != 0 {
__t5 = gopurs_runtime.Constructor1("Just", gopurs_runtime.RecordDict2("key", "value", (*[1024]gopurs_runtime.Value)(v_3.UnsafePtr)[2], (*[1024]gopurs_runtime.Value)(v_3.UnsafePtr)[3]))
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
if gopurs_runtime.Bool(v1_4_2.StrVal == "GT").IntVal != 0 {
__t3 = gopurs_runtime.Apply(go__2_0, (*[1024]gopurs_runtime.Value)(v_3.UnsafePtr)[5])
goto end_branch_3
} else {

}
}
{
if gopurs_runtime.Bool(v1_4_2.StrVal == "EQ").IntVal != 0 {
__t3 = gopurs_runtime.Apply(Get_findMin(), (*[1024]gopurs_runtime.Value)(v_3.UnsafePtr)[5])
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
if gopurs_runtime.Bool(v_0_loop.StrVal == "Leaf").IntVal != 0 {
__t0 = gopurs_runtime.Constructor0("Nothing")
goto end_branch_0
} else {

}
}
{
if gopurs_runtime.Bool(v_0_loop.StrVal == "Node").IntVal != 0 {
var __t1 gopurs_runtime.Value
{
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)(v_0_loop.UnsafePtr)[5].StrVal == "Leaf").IntVal != 0 {
__t1 = gopurs_runtime.Constructor1("Just", gopurs_runtime.RecordDict2("key", "value", (*[1024]gopurs_runtime.Value)(v_0_loop.UnsafePtr)[2], (*[1024]gopurs_runtime.Value)(v_0_loop.UnsafePtr)[3]))
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Apply(Get_findMax(), (*[1024]gopurs_runtime.Value)(v_0_loop.UnsafePtr)[5])
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
		lookupLT = gopurs_runtime.Func2(func(dictOrd_0 gopurs_runtime.Value, k_1 gopurs_runtime.Value) gopurs_runtime.Value {
var go__2_0 gopurs_runtime.Value
_ = go__2_0
go__2_0 = gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if gopurs_runtime.Bool(v_3.StrVal == "Leaf").IntVal != 0 {
__t1 = gopurs_runtime.Constructor0("Nothing")
goto end_branch_1
} else {

}
}
{
if gopurs_runtime.Bool(v_3.StrVal == "Node").IntVal != 0 {
v1_4_2 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictOrd_0, "compare"), k_1, (*[1024]gopurs_runtime.Value)(v_3.UnsafePtr)[2])
_ = v1_4_2
var __t3 gopurs_runtime.Value
{
if gopurs_runtime.Bool(v1_4_2.StrVal == "LT").IntVal != 0 {
__t3 = gopurs_runtime.Apply(go__2_0, (*[1024]gopurs_runtime.Value)(v_3.UnsafePtr)[4])
goto end_branch_3
} else {

}
}
{
if gopurs_runtime.Bool(v1_4_2.StrVal == "GT").IntVal != 0 {
v2_5_4 := gopurs_runtime.Apply(go__2_0, (*[1024]gopurs_runtime.Value)(v_3.UnsafePtr)[5])
_ = v2_5_4
var __t5 gopurs_runtime.Value
{
if gopurs_runtime.Bool(v2_5_4.StrVal == "Nothing").IntVal != 0 {
__t5 = gopurs_runtime.Constructor1("Just", gopurs_runtime.RecordDict2("key", "value", (*[1024]gopurs_runtime.Value)(v_3.UnsafePtr)[2], (*[1024]gopurs_runtime.Value)(v_3.UnsafePtr)[3]))
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
if gopurs_runtime.Bool(v1_4_2.StrVal == "EQ").IntVal != 0 {
__t3 = gopurs_runtime.Apply(Get_findMax(), (*[1024]gopurs_runtime.Value)(v_3.UnsafePtr)[4])
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
})
	})
	return lookupLT
}

var filterWithKey gopurs_runtime.Value
var once_filterWithKey sync.Once
func Get_filterWithKey() gopurs_runtime.Value {
	once_filterWithKey.Do(func() {
		filterWithKey = gopurs_runtime.Func2(func(dictOrd_0 gopurs_runtime.Value, f_1 gopurs_runtime.Value) gopurs_runtime.Value {
var go__2_0 gopurs_runtime.Value
_ = go__2_0
go__2_0 = gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if gopurs_runtime.Bool(v_3.StrVal == "Leaf").IntVal != 0 {
__t1 = gopurs_runtime.Constructor0("Leaf")
goto end_branch_1
} else {

}
}
{
if gopurs_runtime.Bool(v_3.StrVal == "Node").IntVal != 0 {
var __t2 gopurs_runtime.Value
{
if gopurs_runtime.Apply2(f_1, (*[1024]gopurs_runtime.Value)(v_3.UnsafePtr)[2], (*[1024]gopurs_runtime.Value)(v_3.UnsafePtr)[3]).IntVal != 0 {
__t2 = pkg_Data_Map_Internal.Call_unsafeBalancedNode((*[1024]gopurs_runtime.Value)(v_3.UnsafePtr)[2], (*[1024]gopurs_runtime.Value)(v_3.UnsafePtr)[3], gopurs_runtime.Apply(go__2_0, (*[1024]gopurs_runtime.Value)(v_3.UnsafePtr)[4]), gopurs_runtime.Apply(go__2_0, (*[1024]gopurs_runtime.Value)(v_3.UnsafePtr)[5]))
goto end_branch_2
} else {

}
}
{
__t2 = pkg_Data_Map_Internal.Call_unsafeJoinNodes(gopurs_runtime.Apply(go__2_0, (*[1024]gopurs_runtime.Value)(v_3.UnsafePtr)[4]), gopurs_runtime.Apply(go__2_0, (*[1024]gopurs_runtime.Value)(v_3.UnsafePtr)[5]))
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
})
	})
	return filterWithKey
}

var filterKeys gopurs_runtime.Value
var once_filterKeys sync.Once
func Get_filterKeys() gopurs_runtime.Value {
	once_filterKeys.Do(func() {
		filterKeys = gopurs_runtime.Func2(func(dictOrd_0 gopurs_runtime.Value, f_1 gopurs_runtime.Value) gopurs_runtime.Value {
var go__2_0 gopurs_runtime.Value
_ = go__2_0
go__2_0 = gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if gopurs_runtime.Bool(v_3.StrVal == "Leaf").IntVal != 0 {
__t1 = gopurs_runtime.Constructor0("Leaf")
goto end_branch_1
} else {

}
}
{
if gopurs_runtime.Bool(v_3.StrVal == "Node").IntVal != 0 {
var __t2 gopurs_runtime.Value
{
if gopurs_runtime.Apply(f_1, (*[1024]gopurs_runtime.Value)(v_3.UnsafePtr)[2]).IntVal != 0 {
__t2 = pkg_Data_Map_Internal.Call_unsafeBalancedNode((*[1024]gopurs_runtime.Value)(v_3.UnsafePtr)[2], (*[1024]gopurs_runtime.Value)(v_3.UnsafePtr)[3], gopurs_runtime.Apply(go__2_0, (*[1024]gopurs_runtime.Value)(v_3.UnsafePtr)[4]), gopurs_runtime.Apply(go__2_0, (*[1024]gopurs_runtime.Value)(v_3.UnsafePtr)[5]))
goto end_branch_2
} else {

}
}
{
__t2 = pkg_Data_Map_Internal.Call_unsafeJoinNodes(gopurs_runtime.Apply(go__2_0, (*[1024]gopurs_runtime.Value)(v_3.UnsafePtr)[4]), gopurs_runtime.Apply(go__2_0, (*[1024]gopurs_runtime.Value)(v_3.UnsafePtr)[5]))
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
})
	})
	return filterKeys
}

var filter gopurs_runtime.Value
var once_filter sync.Once
func Get_filter() gopurs_runtime.Value {
	once_filter.Do(func() {
		filter = gopurs_runtime.Func2(func(dictOrd_0 gopurs_runtime.Value, x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Get_filterWithKey(), dictOrd_0, gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return x_1
}))
})
	})
	return filter
}

var eqMap gopurs_runtime.Value
var once_eqMap sync.Once
func Get_eqMap() gopurs_runtime.Value {
	once_eqMap.Do(func() {
		eqMap = gopurs_runtime.Func2(func(dictEq_0 gopurs_runtime.Value, dictEq1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict1("eq", gopurs_runtime.Func2(func(xs_2 gopurs_runtime.Value, ys_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if gopurs_runtime.Bool(xs_2.StrVal == "Leaf").IntVal != 0 {
__t0 = gopurs_runtime.Bool(ys_3.StrVal == "Leaf")
goto end_branch_0
} else {

}
}
{
if gopurs_runtime.Bool(xs_2.StrVal == "Node").IntVal != 0 {
__t0 = gopurs_runtime.Boolean(gopurs_runtime.Bool(ys_3.StrVal == "Node").IntVal != 0 && (*[1024]gopurs_runtime.Value)(xs_2.UnsafePtr)[1].IntVal == (*[1024]gopurs_runtime.Value)(ys_3.UnsafePtr)[1].IntVal && gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply2(Get_eqMapIter(), dictEq_0, dictEq1_1), "eq"), gopurs_runtime.Constructor2("IterNode", xs_2, gopurs_runtime.Constructor0("IterLeaf")), gopurs_runtime.Constructor2("IterNode", ys_3, gopurs_runtime.Constructor0("IterLeaf"))).IntVal != 0)
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
})
	})
	return eqMap
}

var ordMap gopurs_runtime.Value
var once_ordMap sync.Once
func Get_ordMap() gopurs_runtime.Value {
	once_ordMap.Do(func() {
		ordMap = gopurs_runtime.Func(func(dictOrd_0 gopurs_runtime.Value) gopurs_runtime.Value {
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
if gopurs_runtime.Bool(xs_5.StrVal == "Leaf").IntVal != 0 {
var __t4 gopurs_runtime.Value
{
if gopurs_runtime.Bool(ys_6.StrVal == "Leaf").IntVal != 0 {
__t4 = gopurs_runtime.Constructor0("EQ")
goto end_branch_4
} else {

}
}
{
__t4 = gopurs_runtime.Constructor0("LT")
}
end_branch_4:
__t3 = __t4
goto end_branch_3
} else {

}
}
{
if gopurs_runtime.Bool(ys_6.StrVal == "Leaf").IntVal != 0 {
__t3 = gopurs_runtime.Constructor0("GT")
goto end_branch_3
} else {

}
}
{
__t3 = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(ordMapIter1_1_0, dictOrd1_3), "compare"), gopurs_runtime.Constructor2("IterNode", xs_5, gopurs_runtime.Constructor0("IterLeaf")), gopurs_runtime.Constructor2("IterNode", ys_6, gopurs_runtime.Constructor0("IterLeaf")))
}
end_branch_3:
return __t3
}), gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return eqMap2_4_2
}))
})
})
	})
	return ordMap
}

var eq1Map gopurs_runtime.Value
var once_eq1Map sync.Once
func Get_eq1Map() gopurs_runtime.Value {
	once_eq1Map.Do(func() {
		eq1Map = gopurs_runtime.Func(func(dictEq_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict1("eq1", gopurs_runtime.Func(func(dictEq1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(gopurs_runtime.Apply2(Get_eqMap(), dictEq_0, dictEq1_1), "eq")
}))
})
	})
	return eq1Map
}

var ord1Map gopurs_runtime.Value
var once_ord1Map sync.Once
func Get_ord1Map() gopurs_runtime.Value {
	once_ord1Map.Do(func() {
		ord1Map = gopurs_runtime.Func(func(dictOrd_0 gopurs_runtime.Value) gopurs_runtime.Value {
ordMap1_1_0 := gopurs_runtime.Apply(Get_ordMap(), dictOrd_0)
_ = ordMap1_1_0
__local_var_2_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd_0, "Eq0"), gopurs_runtime.Value{})
_ = __local_var_2_1
eq1Map1_3_2 := gopurs_runtime.RecordDict1("eq1", gopurs_runtime.Func(func(dictEq1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(gopurs_runtime.Apply2(Get_eqMap(), __local_var_2_1, dictEq1_3), "eq")
}))
_ = eq1Map1_3_2
return gopurs_runtime.RecordDict2("compare1", "Eq10", gopurs_runtime.Func(func(dictOrd1_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(gopurs_runtime.Apply(ordMap1_1_0, dictOrd1_4), "compare")
}), gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return eq1Map1_3_2
}))
})
	})
	return ord1Map
}

var empty gopurs_runtime.Value
var once_empty sync.Once
func Get_empty() gopurs_runtime.Value {
	once_empty.Do(func() {
		empty = gopurs_runtime.Constructor0("Leaf")
	})
	return empty
}

var fromFoldable gopurs_runtime.Value
var once_fromFoldable sync.Once
func Get_fromFoldable() gopurs_runtime.Value {
	once_fromFoldable.Do(func() {
		fromFoldable = gopurs_runtime.Func2(func(dictOrd_0 gopurs_runtime.Value, dictFoldable_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFoldable_1, "foldl"), gopurs_runtime.Func2(func(m_2 gopurs_runtime.Value, v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply4(Get_insert(), dictOrd_0, (*[1024]gopurs_runtime.Value)(v_3.UnsafePtr)[0], (*[1024]gopurs_runtime.Value)(v_3.UnsafePtr)[1], m_2)
}), gopurs_runtime.Constructor0("Leaf"))
})
	})
	return fromFoldable
}

var fromFoldableWith gopurs_runtime.Value
var once_fromFoldableWith sync.Once
func Get_fromFoldableWith() gopurs_runtime.Value {
	once_fromFoldableWith.Do(func() {
		fromFoldableWith = gopurs_runtime.Func3(func(dictOrd_0 gopurs_runtime.Value, dictFoldable_1 gopurs_runtime.Value, f_2 gopurs_runtime.Value) gopurs_runtime.Value {
f_prime_3_0 := gopurs_runtime.Apply2(Get_insertWith(), dictOrd_0, gopurs_runtime.Func2(func(b_3 gopurs_runtime.Value, a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_2, a_4, b_3)
}))
_ = f_prime_3_0
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFoldable_1, "foldl"), gopurs_runtime.Func2(func(m_4 gopurs_runtime.Value, v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(f_prime_3_0, (*[1024]gopurs_runtime.Value)(v_5.UnsafePtr)[0], (*[1024]gopurs_runtime.Value)(v_5.UnsafePtr)[1], m_4)
}), gopurs_runtime.Constructor0("Leaf"))
})
	})
	return fromFoldableWith
}

var fromFoldableWithIndex gopurs_runtime.Value
var once_fromFoldableWithIndex sync.Once
func Get_fromFoldableWithIndex() gopurs_runtime.Value {
	once_fromFoldableWithIndex.Do(func() {
		fromFoldableWithIndex = gopurs_runtime.Func2(func(dictOrd_0 gopurs_runtime.Value, dictFoldableWithIndex_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFoldableWithIndex_1, "foldlWithIndex"), gopurs_runtime.Func3(func(k_2 gopurs_runtime.Value, m_3 gopurs_runtime.Value, v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply4(Get_insert(), dictOrd_0, k_2, v_4, m_3)
}), gopurs_runtime.Constructor0("Leaf"))
})
	})
	return fromFoldableWithIndex
}

var monoidSemigroupMap gopurs_runtime.Value
var once_monoidSemigroupMap sync.Once
func Get_monoidSemigroupMap() gopurs_runtime.Value {
	once_monoidSemigroupMap.Do(func() {
		monoidSemigroupMap = gopurs_runtime.Func2(func(_dollar__unused_0 gopurs_runtime.Value, dictOrd_1 gopurs_runtime.Value) gopurs_runtime.Value {
semigroupMap2_2_0 := gopurs_runtime.Apply2(Get_semigroupMap(), gopurs_runtime.Value{}, dictOrd_1)
_ = semigroupMap2_2_0
return gopurs_runtime.Func(func(dictSemigroup_3 gopurs_runtime.Value) gopurs_runtime.Value {
semigroupMap3_4_1 := gopurs_runtime.Apply(semigroupMap2_2_0, dictSemigroup_3)
_ = semigroupMap3_4_1
return gopurs_runtime.RecordDict2("mempty", "Semigroup0", gopurs_runtime.Constructor0("Leaf"), gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return semigroupMap3_4_1
}))
})
})
	})
	return monoidSemigroupMap
}

var submap gopurs_runtime.Value
var once_submap sync.Once
func Get_submap() gopurs_runtime.Value {
	once_submap.Do(func() {
		submap = gopurs_runtime.Func(func(dictOrd_0 gopurs_runtime.Value) gopurs_runtime.Value {
compare_1_0 := gopurs_runtime.RecordGet(dictOrd_0, "compare")
_ = compare_1_0
return gopurs_runtime.Func2(func(kmin_2 gopurs_runtime.Value, kmax_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply5(Get_foldSubmapBy(), dictOrd_0, gopurs_runtime.Func2(func(m1_4 gopurs_runtime.Value, m2_5 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Map_Internal.Call_unsafeUnionWith(compare_1_0, pkg_Data_Function.Get_const_(), m1_4, m2_5)
}), gopurs_runtime.Constructor0("Leaf"), kmin_2, kmax_3), Get_singleton())
})
})
	})
	return submap
}

var unions gopurs_runtime.Value
var once_unions sync.Once
func Get_unions() gopurs_runtime.Value {
	once_unions.Do(func() {
		unions = gopurs_runtime.Func(func(dictOrd_0 gopurs_runtime.Value) gopurs_runtime.Value {
compare_1_0 := gopurs_runtime.RecordGet(dictOrd_0, "compare")
_ = compare_1_0
return gopurs_runtime.Func(func(dictFoldable_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFoldable_2, "foldl"), gopurs_runtime.Func2(func(m1_3 gopurs_runtime.Value, m2_4 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Map_Internal.Call_unsafeUnionWith(compare_1_0, pkg_Data_Function.Get_const_(), m1_3, m2_4)
}), gopurs_runtime.Constructor0("Leaf"))
})
})
	})
	return unions
}

var difference gopurs_runtime.Value
var once_difference sync.Once
func Get_difference() gopurs_runtime.Value {
	once_difference.Do(func() {
		difference = gopurs_runtime.Func(func(dictOrd_0 gopurs_runtime.Value) gopurs_runtime.Value {
compare_1_0 := gopurs_runtime.RecordGet(dictOrd_0, "compare")
_ = compare_1_0
return gopurs_runtime.Func2(func(m1_2 gopurs_runtime.Value, m2_3 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Map_Internal.Call_unsafeDifference(compare_1_0, m1_2, m2_3)
})
})
	})
	return difference
}

var delete_ gopurs_runtime.Value
var once_delete_ sync.Once
func Get_delete_() gopurs_runtime.Value {
	once_delete_.Do(func() {
		delete_ = gopurs_runtime.Func2(func(dictOrd_0 gopurs_runtime.Value, k_1 gopurs_runtime.Value) gopurs_runtime.Value {
var go__2_0 gopurs_runtime.Value
_ = go__2_0
go__2_0 = gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if gopurs_runtime.Bool(v_3.StrVal == "Leaf").IntVal != 0 {
__t1 = gopurs_runtime.Constructor0("Leaf")
goto end_branch_1
} else {

}
}
{
if gopurs_runtime.Bool(v_3.StrVal == "Node").IntVal != 0 {
v1_4_2 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictOrd_0, "compare"), k_1, (*[1024]gopurs_runtime.Value)(v_3.UnsafePtr)[2])
_ = v1_4_2
var __t3 gopurs_runtime.Value
{
if gopurs_runtime.Bool(v1_4_2.StrVal == "LT").IntVal != 0 {
__t3 = pkg_Data_Map_Internal.Call_unsafeBalancedNode((*[1024]gopurs_runtime.Value)(v_3.UnsafePtr)[2], (*[1024]gopurs_runtime.Value)(v_3.UnsafePtr)[3], gopurs_runtime.Apply(go__2_0, (*[1024]gopurs_runtime.Value)(v_3.UnsafePtr)[4]), (*[1024]gopurs_runtime.Value)(v_3.UnsafePtr)[5])
goto end_branch_3
} else {

}
}
{
if gopurs_runtime.Bool(v1_4_2.StrVal == "GT").IntVal != 0 {
__t3 = pkg_Data_Map_Internal.Call_unsafeBalancedNode((*[1024]gopurs_runtime.Value)(v_3.UnsafePtr)[2], (*[1024]gopurs_runtime.Value)(v_3.UnsafePtr)[3], (*[1024]gopurs_runtime.Value)(v_3.UnsafePtr)[4], gopurs_runtime.Apply(go__2_0, (*[1024]gopurs_runtime.Value)(v_3.UnsafePtr)[5]))
goto end_branch_3
} else {

}
}
{
if gopurs_runtime.Bool(v1_4_2.StrVal == "EQ").IntVal != 0 {
__t3 = pkg_Data_Map_Internal.Call_unsafeJoinNodes((*[1024]gopurs_runtime.Value)(v_3.UnsafePtr)[4], (*[1024]gopurs_runtime.Value)(v_3.UnsafePtr)[5])
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
})
	})
	return delete_
}

var checkValid gopurs_runtime.Value
var once_checkValid sync.Once
func Get_checkValid() gopurs_runtime.Value {
	once_checkValid.Do(func() {
		checkValid = gopurs_runtime.Func(func(dictOrd_0 gopurs_runtime.Value) gopurs_runtime.Value {
var go__1_0 gopurs_runtime.Value
_ = go__1_0
go__1_0 = gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if gopurs_runtime.Bool(v_2.StrVal == "Leaf").IntVal != 0 {
__t1 = gopurs_runtime.Boolean(true)
goto end_branch_1
} else {

}
}
{
if gopurs_runtime.Bool(v_2.StrVal == "Node").IntVal != 0 {
var __t2 gopurs_runtime.Value
{
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)(v_2.UnsafePtr)[4].StrVal == "Leaf").IntVal != 0 {
var __t3 gopurs_runtime.Value
{
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)(v_2.UnsafePtr)[5].StrVal == "Leaf").IntVal != 0 {
__t3 = gopurs_runtime.Boolean(true)
goto end_branch_3
} else {

}
}
{
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)(v_2.UnsafePtr)[5].StrVal == "Node").IntVal != 0 {
__t3 = gopurs_runtime.Boolean((*[1024]gopurs_runtime.Value)(v_2.UnsafePtr)[0].IntVal == 2 && (*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)(v_2.UnsafePtr)[5].UnsafePtr)[0].IntVal == 1 && (*[1024]gopurs_runtime.Value)(v_2.UnsafePtr)[1].IntVal > (*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)(v_2.UnsafePtr)[5].UnsafePtr)[1].IntVal && gopurs_runtime.Bool(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictOrd_0, "compare"), (*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)(v_2.UnsafePtr)[5].UnsafePtr)[2], (*[1024]gopurs_runtime.Value)(v_2.UnsafePtr)[2]).StrVal == "GT").IntVal != 0 && gopurs_runtime.Apply(go__1_0, (*[1024]gopurs_runtime.Value)(v_2.UnsafePtr)[5]).IntVal != 0)
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
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)(v_2.UnsafePtr)[4].StrVal == "Node").IntVal != 0 {
var __t4 gopurs_runtime.Value
{
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)(v_2.UnsafePtr)[5].StrVal == "Leaf").IntVal != 0 {
__t4 = gopurs_runtime.Boolean((*[1024]gopurs_runtime.Value)(v_2.UnsafePtr)[0].IntVal == 2 && (*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)(v_2.UnsafePtr)[4].UnsafePtr)[0].IntVal == 1 && (*[1024]gopurs_runtime.Value)(v_2.UnsafePtr)[1].IntVal > (*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)(v_2.UnsafePtr)[4].UnsafePtr)[1].IntVal && gopurs_runtime.Bool(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictOrd_0, "compare"), (*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)(v_2.UnsafePtr)[4].UnsafePtr)[2], (*[1024]gopurs_runtime.Value)(v_2.UnsafePtr)[2]).StrVal == "LT").IntVal != 0 && gopurs_runtime.Apply(go__1_0, (*[1024]gopurs_runtime.Value)(v_2.UnsafePtr)[4]).IntVal != 0)
goto end_branch_4
} else {

}
}
{
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)(v_2.UnsafePtr)[5].StrVal == "Node").IntVal != 0 {
__local_var_3_5 := (*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)(v_2.UnsafePtr)[5].UnsafePtr)[0].IntVal - (*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)(v_2.UnsafePtr)[4].UnsafePtr)[0].IntVal
_ = __local_var_3_5
var __t6 gopurs_runtime.Value
{
if __local_var_3_5 >= 0 {
__t6 = gopurs_runtime.Boolean(__local_var_3_5 < 2)
goto end_branch_6
} else {

}
}
{
__t6 = gopurs_runtime.Boolean(0 - __local_var_3_5 < 2)
}
end_branch_6:
__t4 = gopurs_runtime.Boolean((*[1024]gopurs_runtime.Value)(v_2.UnsafePtr)[0].IntVal > (*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)(v_2.UnsafePtr)[5].UnsafePtr)[0].IntVal && gopurs_runtime.Bool(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictOrd_0, "compare"), (*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)(v_2.UnsafePtr)[5].UnsafePtr)[2], (*[1024]gopurs_runtime.Value)(v_2.UnsafePtr)[2]).StrVal == "GT").IntVal != 0 && (*[1024]gopurs_runtime.Value)(v_2.UnsafePtr)[0].IntVal > (*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)(v_2.UnsafePtr)[4].UnsafePtr)[0].IntVal && gopurs_runtime.Bool(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictOrd_0, "compare"), (*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)(v_2.UnsafePtr)[4].UnsafePtr)[2], (*[1024]gopurs_runtime.Value)(v_2.UnsafePtr)[2]).StrVal == "LT").IntVal != 0 && __t6.IntVal != 0 && (*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)(v_2.UnsafePtr)[5].UnsafePtr)[1].IntVal + (*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)(v_2.UnsafePtr)[4].UnsafePtr)[1].IntVal + 1 == (*[1024]gopurs_runtime.Value)(v_2.UnsafePtr)[1].IntVal && gopurs_runtime.Apply(go__1_0, (*[1024]gopurs_runtime.Value)(v_2.UnsafePtr)[4]).IntVal != 0 && gopurs_runtime.Apply(go__1_0, (*[1024]gopurs_runtime.Value)(v_2.UnsafePtr)[5]).IntVal != 0)
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
})
	})
	return checkValid
}

var catMaybes gopurs_runtime.Value
var once_catMaybes sync.Once
func Get_catMaybes() gopurs_runtime.Value {
	once_catMaybes.Do(func() {
		catMaybes = gopurs_runtime.Func(func(dictOrd_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Get_mapMaybeWithKey(), dictOrd_0, gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(pkg_Control_Category.Get_categoryFn(), "identity")
}))
})
	})
	return catMaybes
}

var applyMap gopurs_runtime.Value
var once_applyMap sync.Once
func Get_applyMap() gopurs_runtime.Value {
	once_applyMap.Do(func() {
		applyMap = gopurs_runtime.Func(func(dictOrd_0 gopurs_runtime.Value) gopurs_runtime.Value {
compare_1_0 := gopurs_runtime.RecordGet(dictOrd_0, "compare")
_ = compare_1_0
return gopurs_runtime.RecordDict2("apply", "Functor0", gopurs_runtime.Func2(func(m1_2 gopurs_runtime.Value, m2_3 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Map_Internal.Call_unsafeIntersectionWith(compare_1_0, gopurs_runtime.RecordGet(pkg_Control_Category.Get_categoryFn(), "identity"), m1_2, m2_3)
}), gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_functorMap()
}))
})
	})
	return applyMap
}

var bindMap gopurs_runtime.Value
var once_bindMap sync.Once
func Get_bindMap() gopurs_runtime.Value {
	once_bindMap.Do(func() {
		bindMap = gopurs_runtime.Func(func(dictOrd_0 gopurs_runtime.Value) gopurs_runtime.Value {
compare_1_1 := gopurs_runtime.RecordGet(dictOrd_0, "compare")
_ = compare_1_1
applyMap1_1_0 := gopurs_runtime.RecordDict2("apply", "Functor0", gopurs_runtime.Func2(func(m1_2 gopurs_runtime.Value, m2_3 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Map_Internal.Call_unsafeIntersectionWith(compare_1_1, gopurs_runtime.RecordGet(pkg_Control_Category.Get_categoryFn(), "identity"), m1_2, m2_3)
}), gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_functorMap()
}))
_ = applyMap1_1_0
return gopurs_runtime.RecordDict2("bind", "Apply0", gopurs_runtime.Func2(func(m_2 gopurs_runtime.Value, f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(Get_mapMaybeWithKey(), dictOrd_0, gopurs_runtime.Func(func(k_4 gopurs_runtime.Value) gopurs_runtime.Value {
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
if gopurs_runtime.Bool(v_6.StrVal == "Leaf").IntVal != 0 {
__t3 = gopurs_runtime.Constructor0("Nothing")
goto end_branch_3
} else {

}
}
{
if gopurs_runtime.Bool(v_6.StrVal == "Node").IntVal != 0 {
v1_7_4 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictOrd_0, "compare"), k_4, (*[1024]gopurs_runtime.Value)(v_6.UnsafePtr)[2])
_ = v1_7_4
var __t5 gopurs_runtime.Value
{
if gopurs_runtime.Bool(v1_7_4.StrVal == "LT").IntVal != 0 {
v_6_loop = (*[1024]gopurs_runtime.Value)(v_6.UnsafePtr)[4]
continue go__5_2
__t5 = gopurs_runtime.Value{}
goto end_branch_5
} else {

}
}
{
if gopurs_runtime.Bool(v1_7_4.StrVal == "GT").IntVal != 0 {
v_6_loop = (*[1024]gopurs_runtime.Value)(v_6.UnsafePtr)[5]
continue go__5_2
__t5 = gopurs_runtime.Value{}
goto end_branch_5
} else {

}
}
{
if gopurs_runtime.Bool(v1_7_4.StrVal == "EQ").IntVal != 0 {
__t5 = gopurs_runtime.Constructor1("Just", (*[1024]gopurs_runtime.Value)(v_6.UnsafePtr)[3])
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
}), m_2)
}), gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return applyMap1_1_0
}))
})
	})
	return bindMap
}

var anyWithKey gopurs_runtime.Value
var once_anyWithKey sync.Once
func Get_anyWithKey() gopurs_runtime.Value {
	once_anyWithKey.Do(func() {
		anyWithKey = gopurs_runtime.Func(func(predicate_0 gopurs_runtime.Value) gopurs_runtime.Value {
var go__1_0 gopurs_runtime.Value
_ = go__1_0
go__1_0 = gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if gopurs_runtime.Bool(v_2.StrVal == "Leaf").IntVal != 0 {
__t1 = gopurs_runtime.Boolean(false)
goto end_branch_1
} else {

}
}
{
if gopurs_runtime.Bool(v_2.StrVal == "Node").IntVal != 0 {
__t1 = gopurs_runtime.Boolean(gopurs_runtime.Apply2(predicate_0, (*[1024]gopurs_runtime.Value)(v_2.UnsafePtr)[2], (*[1024]gopurs_runtime.Value)(v_2.UnsafePtr)[3]).IntVal != 0 || gopurs_runtime.Apply(go__1_0, (*[1024]gopurs_runtime.Value)(v_2.UnsafePtr)[4]).IntVal != 0 || gopurs_runtime.Apply(go__1_0, (*[1024]gopurs_runtime.Value)(v_2.UnsafePtr)[5]).IntVal != 0)
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
})
	})
	return anyWithKey
}

var any gopurs_runtime.Value
var once_any sync.Once
func Get_any() gopurs_runtime.Value {
	once_any.Do(func() {
		any = gopurs_runtime.Func(func(predicate_0 gopurs_runtime.Value) gopurs_runtime.Value {
var go__1_0 gopurs_runtime.Value
_ = go__1_0
go__1_0 = gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if gopurs_runtime.Bool(v_2.StrVal == "Leaf").IntVal != 0 {
__t1 = gopurs_runtime.Boolean(false)
goto end_branch_1
} else {

}
}
{
if gopurs_runtime.Bool(v_2.StrVal == "Node").IntVal != 0 {
__t1 = gopurs_runtime.Boolean(gopurs_runtime.Apply(predicate_0, (*[1024]gopurs_runtime.Value)(v_2.UnsafePtr)[3]).IntVal != 0 || gopurs_runtime.Apply(go__1_0, (*[1024]gopurs_runtime.Value)(v_2.UnsafePtr)[4]).IntVal != 0 || gopurs_runtime.Apply(go__1_0, (*[1024]gopurs_runtime.Value)(v_2.UnsafePtr)[5]).IntVal != 0)
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
})
	})
	return any
}

var alter gopurs_runtime.Value
var once_alter sync.Once
func Get_alter() gopurs_runtime.Value {
	once_alter.Do(func() {
		alter = gopurs_runtime.Func(func(dictOrd_0 gopurs_runtime.Value) gopurs_runtime.Value {
compare_1_0 := gopurs_runtime.RecordGet(dictOrd_0, "compare")
_ = compare_1_0
return gopurs_runtime.Func3(func(f_2 gopurs_runtime.Value, k_3 gopurs_runtime.Value, m_4 gopurs_runtime.Value) gopurs_runtime.Value {
v_5_1 := pkg_Data_Map_Internal.Call_unsafeSplit(compare_1_0, k_3, m_4)
_ = v_5_1
v2_6_2 := gopurs_runtime.Apply(f_2, (*[1024]gopurs_runtime.Value)(v_5_1.UnsafePtr)[0])
_ = v2_6_2
var __t3 gopurs_runtime.Value
{
if gopurs_runtime.Bool(v2_6_2.StrVal == "Nothing").IntVal != 0 {
__t3 = pkg_Data_Map_Internal.Call_unsafeJoinNodes((*[1024]gopurs_runtime.Value)(v_5_1.UnsafePtr)[1], (*[1024]gopurs_runtime.Value)(v_5_1.UnsafePtr)[2])
goto end_branch_3
} else {

}
}
{
if gopurs_runtime.Bool(v2_6_2.StrVal == "Just").IntVal != 0 {
__t3 = pkg_Data_Map_Internal.Call_unsafeBalancedNode(k_3, (*[1024]gopurs_runtime.Value)(v2_6_2.UnsafePtr)[0], (*[1024]gopurs_runtime.Value)(v_5_1.UnsafePtr)[1], (*[1024]gopurs_runtime.Value)(v_5_1.UnsafePtr)[2])
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
})
	})
	return alter
}

var altMap gopurs_runtime.Value
var once_altMap sync.Once
func Get_altMap() gopurs_runtime.Value {
	once_altMap.Do(func() {
		altMap = gopurs_runtime.Func(func(dictOrd_0 gopurs_runtime.Value) gopurs_runtime.Value {
compare_1_0 := gopurs_runtime.RecordGet(dictOrd_0, "compare")
_ = compare_1_0
return gopurs_runtime.RecordDict2("alt", "Functor0", gopurs_runtime.Func2(func(m1_2 gopurs_runtime.Value, m2_3 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Map_Internal.Call_unsafeUnionWith(compare_1_0, pkg_Data_Function.Get_const_(), m1_2, m2_3)
}), gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_functorMap()
}))
})
	})
	return altMap
}

var plusMap gopurs_runtime.Value
var once_plusMap sync.Once
func Get_plusMap() gopurs_runtime.Value {
	once_plusMap.Do(func() {
		plusMap = gopurs_runtime.Func(func(dictOrd_0 gopurs_runtime.Value) gopurs_runtime.Value {
compare_1_1 := gopurs_runtime.RecordGet(dictOrd_0, "compare")
_ = compare_1_1
altMap1_1_0 := gopurs_runtime.RecordDict2("alt", "Functor0", gopurs_runtime.Func2(func(m1_2 gopurs_runtime.Value, m2_3 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Map_Internal.Call_unsafeUnionWith(compare_1_1, pkg_Data_Function.Get_const_(), m1_2, m2_3)
}), gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_functorMap()
}))
_ = altMap1_1_0
return gopurs_runtime.RecordDict2("empty", "Alt0", gopurs_runtime.Constructor0("Leaf"), gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return altMap1_1_0
}))
})
	})
	return plusMap
}

func Call_unsafeNode(k_0 gopurs_runtime.Value, v_1 gopurs_runtime.Value, l_2 gopurs_runtime.Value, r_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if gopurs_runtime.Bool(l_2.StrVal == "Leaf").IntVal != 0 {
var __t1 gopurs_runtime.Value
{
if gopurs_runtime.Bool(r_3.StrVal == "Leaf").IntVal != 0 {
__t1 = gopurs_runtime.Constructor("Node", []gopurs_runtime.Value{gopurs_runtime.Int(1), gopurs_runtime.Int(1), k_0, v_1, l_2, r_3})
goto end_branch_1
} else {

}
}
{
if gopurs_runtime.Bool(r_3.StrVal == "Node").IntVal != 0 {
__t1 = gopurs_runtime.Constructor("Node", []gopurs_runtime.Value{gopurs_runtime.Int(1 + (*[1024]gopurs_runtime.Value)(r_3.UnsafePtr)[0].IntVal), gopurs_runtime.Int(1 + (*[1024]gopurs_runtime.Value)(r_3.UnsafePtr)[1].IntVal), k_0, v_1, l_2, r_3})
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
if gopurs_runtime.Bool(l_2.StrVal == "Node").IntVal != 0 {
var __t2 gopurs_runtime.Value
{
if gopurs_runtime.Bool(r_3.StrVal == "Leaf").IntVal != 0 {
__t2 = gopurs_runtime.Constructor("Node", []gopurs_runtime.Value{gopurs_runtime.Int(1 + (*[1024]gopurs_runtime.Value)(l_2.UnsafePtr)[0].IntVal), gopurs_runtime.Int(1 + (*[1024]gopurs_runtime.Value)(l_2.UnsafePtr)[1].IntVal), k_0, v_1, l_2, r_3})
goto end_branch_2
} else {

}
}
{
if gopurs_runtime.Bool(r_3.StrVal == "Node").IntVal != 0 {
var __t3 gopurs_runtime.Value
{
if (*[1024]gopurs_runtime.Value)(l_2.UnsafePtr)[0].IntVal > (*[1024]gopurs_runtime.Value)(r_3.UnsafePtr)[0].IntVal {
__t3 = gopurs_runtime.Int(1 + (*[1024]gopurs_runtime.Value)(l_2.UnsafePtr)[0].IntVal)
goto end_branch_3
} else {

}
}
{
__t3 = gopurs_runtime.Int(1 + (*[1024]gopurs_runtime.Value)(r_3.UnsafePtr)[0].IntVal)
}
end_branch_3:
__t2 = gopurs_runtime.Constructor("Node", []gopurs_runtime.Value{__t3, gopurs_runtime.Int(1 + (*[1024]gopurs_runtime.Value)(l_2.UnsafePtr)[1].IntVal + (*[1024]gopurs_runtime.Value)(r_3.UnsafePtr)[1].IntVal), k_0, v_1, l_2, r_3})
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

func Call_unsafeBalancedNode(k_0 gopurs_runtime.Value, v_1 gopurs_runtime.Value, l_2 gopurs_runtime.Value, r_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if gopurs_runtime.Bool(l_2.StrVal == "Leaf").IntVal != 0 {
var __t1 gopurs_runtime.Value
{
if gopurs_runtime.Bool(r_3.StrVal == "Leaf").IntVal != 0 {
__t1 = gopurs_runtime.Constructor("Node", []gopurs_runtime.Value{gopurs_runtime.Int(1), gopurs_runtime.Int(1), k_0, v_1, gopurs_runtime.Constructor0("Leaf"), gopurs_runtime.Constructor0("Leaf")})
goto end_branch_1
} else {

}
}
{
if gopurs_runtime.Bool(r_3.StrVal == "Node").IntVal != 0 && (*[1024]gopurs_runtime.Value)(r_3.UnsafePtr)[0].IntVal > 1 {
var __t2 gopurs_runtime.Value
{
var __t3 gopurs_runtime.Value
{
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)(r_3.UnsafePtr)[5].StrVal == "Leaf").IntVal != 0 {
__t3 = gopurs_runtime.Boolean((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)(r_3.UnsafePtr)[4].UnsafePtr)[0].IntVal > 0)
goto end_branch_3
} else {

}
}
{
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)(r_3.UnsafePtr)[5].StrVal == "Node").IntVal != 0 {
__t3 = gopurs_runtime.Boolean((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)(r_3.UnsafePtr)[4].UnsafePtr)[0].IntVal > (*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)(r_3.UnsafePtr)[5].UnsafePtr)[0].IntVal)
goto end_branch_3
} else {

}
}
{
__t3 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_3:
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)(r_3.UnsafePtr)[4].StrVal == "Node").IntVal != 0 && __t3.IntVal != 0 {
__t2 = pkg_Data_Map_Internal.Call_unsafeNode((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)(r_3.UnsafePtr)[4].UnsafePtr)[2], (*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)(r_3.UnsafePtr)[4].UnsafePtr)[3], pkg_Data_Map_Internal.Call_unsafeNode(k_0, v_1, l_2, (*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)(r_3.UnsafePtr)[4].UnsafePtr)[4]), pkg_Data_Map_Internal.Call_unsafeNode((*[1024]gopurs_runtime.Value)(r_3.UnsafePtr)[2], (*[1024]gopurs_runtime.Value)(r_3.UnsafePtr)[3], (*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)(r_3.UnsafePtr)[4].UnsafePtr)[5], (*[1024]gopurs_runtime.Value)(r_3.UnsafePtr)[5]))
goto end_branch_2
} else {

}
}
{
__t2 = pkg_Data_Map_Internal.Call_unsafeNode((*[1024]gopurs_runtime.Value)(r_3.UnsafePtr)[2], (*[1024]gopurs_runtime.Value)(r_3.UnsafePtr)[3], pkg_Data_Map_Internal.Call_unsafeNode(k_0, v_1, l_2, (*[1024]gopurs_runtime.Value)(r_3.UnsafePtr)[4]), (*[1024]gopurs_runtime.Value)(r_3.UnsafePtr)[5])
}
end_branch_2:
__t1 = __t2
goto end_branch_1
} else {

}
}
{
__t1 = pkg_Data_Map_Internal.Call_unsafeNode(k_0, v_1, l_2, r_3)
}
end_branch_1:
__t0 = __t1
goto end_branch_0
} else {

}
}
{
if gopurs_runtime.Bool(l_2.StrVal == "Node").IntVal != 0 {
var __t4 gopurs_runtime.Value
{
if gopurs_runtime.Bool(r_3.StrVal == "Node").IntVal != 0 {
var __t5 gopurs_runtime.Value
{
if (*[1024]gopurs_runtime.Value)(r_3.UnsafePtr)[0].IntVal > (*[1024]gopurs_runtime.Value)(l_2.UnsafePtr)[0].IntVal + 1 {
var __t6 gopurs_runtime.Value
{
var __t7 gopurs_runtime.Value
{
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)(r_3.UnsafePtr)[5].StrVal == "Leaf").IntVal != 0 {
__t7 = gopurs_runtime.Boolean((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)(r_3.UnsafePtr)[4].UnsafePtr)[0].IntVal > 0)
goto end_branch_7
} else {

}
}
{
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)(r_3.UnsafePtr)[5].StrVal == "Node").IntVal != 0 {
__t7 = gopurs_runtime.Boolean((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)(r_3.UnsafePtr)[4].UnsafePtr)[0].IntVal > (*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)(r_3.UnsafePtr)[5].UnsafePtr)[0].IntVal)
goto end_branch_7
} else {

}
}
{
__t7 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_7:
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)(r_3.UnsafePtr)[4].StrVal == "Node").IntVal != 0 && __t7.IntVal != 0 {
__t6 = pkg_Data_Map_Internal.Call_unsafeNode((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)(r_3.UnsafePtr)[4].UnsafePtr)[2], (*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)(r_3.UnsafePtr)[4].UnsafePtr)[3], pkg_Data_Map_Internal.Call_unsafeNode(k_0, v_1, l_2, (*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)(r_3.UnsafePtr)[4].UnsafePtr)[4]), pkg_Data_Map_Internal.Call_unsafeNode((*[1024]gopurs_runtime.Value)(r_3.UnsafePtr)[2], (*[1024]gopurs_runtime.Value)(r_3.UnsafePtr)[3], (*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)(r_3.UnsafePtr)[4].UnsafePtr)[5], (*[1024]gopurs_runtime.Value)(r_3.UnsafePtr)[5]))
goto end_branch_6
} else {

}
}
{
__t6 = pkg_Data_Map_Internal.Call_unsafeNode((*[1024]gopurs_runtime.Value)(r_3.UnsafePtr)[2], (*[1024]gopurs_runtime.Value)(r_3.UnsafePtr)[3], pkg_Data_Map_Internal.Call_unsafeNode(k_0, v_1, l_2, (*[1024]gopurs_runtime.Value)(r_3.UnsafePtr)[4]), (*[1024]gopurs_runtime.Value)(r_3.UnsafePtr)[5])
}
end_branch_6:
__t5 = __t6
goto end_branch_5
} else {

}
}
{
if (*[1024]gopurs_runtime.Value)(l_2.UnsafePtr)[0].IntVal > (*[1024]gopurs_runtime.Value)(r_3.UnsafePtr)[0].IntVal + 1 {
var __t8 gopurs_runtime.Value
{
var __t9 gopurs_runtime.Value
{
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)(l_2.UnsafePtr)[4].StrVal == "Leaf").IntVal != 0 {
__t9 = gopurs_runtime.Boolean(0 <= (*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)(l_2.UnsafePtr)[5].UnsafePtr)[0].IntVal)
goto end_branch_9
} else {

}
}
{
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)(l_2.UnsafePtr)[4].StrVal == "Node").IntVal != 0 {
__t9 = gopurs_runtime.Boolean((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)(l_2.UnsafePtr)[4].UnsafePtr)[0].IntVal <= (*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)(l_2.UnsafePtr)[5].UnsafePtr)[0].IntVal)
goto end_branch_9
} else {

}
}
{
__t9 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_9:
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)(l_2.UnsafePtr)[5].StrVal == "Node").IntVal != 0 && __t9.IntVal != 0 {
__t8 = pkg_Data_Map_Internal.Call_unsafeNode((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)(l_2.UnsafePtr)[5].UnsafePtr)[2], (*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)(l_2.UnsafePtr)[5].UnsafePtr)[3], pkg_Data_Map_Internal.Call_unsafeNode((*[1024]gopurs_runtime.Value)(l_2.UnsafePtr)[2], (*[1024]gopurs_runtime.Value)(l_2.UnsafePtr)[3], (*[1024]gopurs_runtime.Value)(l_2.UnsafePtr)[4], (*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)(l_2.UnsafePtr)[5].UnsafePtr)[4]), pkg_Data_Map_Internal.Call_unsafeNode(k_0, v_1, (*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)(l_2.UnsafePtr)[5].UnsafePtr)[5], r_3))
goto end_branch_8
} else {

}
}
{
__t8 = pkg_Data_Map_Internal.Call_unsafeNode((*[1024]gopurs_runtime.Value)(l_2.UnsafePtr)[2], (*[1024]gopurs_runtime.Value)(l_2.UnsafePtr)[3], (*[1024]gopurs_runtime.Value)(l_2.UnsafePtr)[4], pkg_Data_Map_Internal.Call_unsafeNode(k_0, v_1, (*[1024]gopurs_runtime.Value)(l_2.UnsafePtr)[5], r_3))
}
end_branch_8:
__t5 = __t8
goto end_branch_5
} else {

}
}
{
__t5 = pkg_Data_Map_Internal.Call_unsafeNode(k_0, v_1, l_2, r_3)
}
end_branch_5:
__t4 = __t5
goto end_branch_4
} else {

}
}
{
if gopurs_runtime.Bool(r_3.StrVal == "Leaf").IntVal != 0 && (*[1024]gopurs_runtime.Value)(l_2.UnsafePtr)[0].IntVal > 1 {
var __t10 gopurs_runtime.Value
{
var __t11 gopurs_runtime.Value
{
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)(l_2.UnsafePtr)[4].StrVal == "Leaf").IntVal != 0 {
__t11 = gopurs_runtime.Boolean(0 <= (*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)(l_2.UnsafePtr)[5].UnsafePtr)[0].IntVal)
goto end_branch_11
} else {

}
}
{
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)(l_2.UnsafePtr)[4].StrVal == "Node").IntVal != 0 {
__t11 = gopurs_runtime.Boolean((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)(l_2.UnsafePtr)[4].UnsafePtr)[0].IntVal <= (*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)(l_2.UnsafePtr)[5].UnsafePtr)[0].IntVal)
goto end_branch_11
} else {

}
}
{
__t11 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_11:
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)(l_2.UnsafePtr)[5].StrVal == "Node").IntVal != 0 && __t11.IntVal != 0 {
__t10 = pkg_Data_Map_Internal.Call_unsafeNode((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)(l_2.UnsafePtr)[5].UnsafePtr)[2], (*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)(l_2.UnsafePtr)[5].UnsafePtr)[3], pkg_Data_Map_Internal.Call_unsafeNode((*[1024]gopurs_runtime.Value)(l_2.UnsafePtr)[2], (*[1024]gopurs_runtime.Value)(l_2.UnsafePtr)[3], (*[1024]gopurs_runtime.Value)(l_2.UnsafePtr)[4], (*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)(l_2.UnsafePtr)[5].UnsafePtr)[4]), pkg_Data_Map_Internal.Call_unsafeNode(k_0, v_1, (*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)(l_2.UnsafePtr)[5].UnsafePtr)[5], r_3))
goto end_branch_10
} else {

}
}
{
__t10 = pkg_Data_Map_Internal.Call_unsafeNode((*[1024]gopurs_runtime.Value)(l_2.UnsafePtr)[2], (*[1024]gopurs_runtime.Value)(l_2.UnsafePtr)[3], (*[1024]gopurs_runtime.Value)(l_2.UnsafePtr)[4], pkg_Data_Map_Internal.Call_unsafeNode(k_0, v_1, (*[1024]gopurs_runtime.Value)(l_2.UnsafePtr)[5], r_3))
}
end_branch_10:
__t4 = __t10
goto end_branch_4
} else {

}
}
{
__t4 = pkg_Data_Map_Internal.Call_unsafeNode(k_0, v_1, l_2, r_3)
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
if gopurs_runtime.Bool(m_2_loop.StrVal == "Leaf").IntVal != 0 {
__t0 = gopurs_runtime.Constructor3("Split", gopurs_runtime.Constructor0("Nothing"), gopurs_runtime.Constructor0("Leaf"), gopurs_runtime.Constructor0("Leaf"))
goto end_branch_0
} else {

}
}
{
if gopurs_runtime.Bool(m_2_loop.StrVal == "Node").IntVal != 0 {
v_3_1 := gopurs_runtime.Apply2(comp_0_loop, k_1_loop, (*[1024]gopurs_runtime.Value)(m_2_loop.UnsafePtr)[2])
_ = v_3_1
var __t2 gopurs_runtime.Value
{
if gopurs_runtime.Bool(v_3_1.StrVal == "LT").IntVal != 0 {
v1_4_3 := pkg_Data_Map_Internal.Call_unsafeSplit(comp_0_loop, k_1_loop, (*[1024]gopurs_runtime.Value)(m_2_loop.UnsafePtr)[4])
_ = v1_4_3
__t2 = gopurs_runtime.Constructor3("Split", (*[1024]gopurs_runtime.Value)(v1_4_3.UnsafePtr)[0], (*[1024]gopurs_runtime.Value)(v1_4_3.UnsafePtr)[1], pkg_Data_Map_Internal.Call_unsafeBalancedNode((*[1024]gopurs_runtime.Value)(m_2_loop.UnsafePtr)[2], (*[1024]gopurs_runtime.Value)(m_2_loop.UnsafePtr)[3], (*[1024]gopurs_runtime.Value)(v1_4_3.UnsafePtr)[2], (*[1024]gopurs_runtime.Value)(m_2_loop.UnsafePtr)[5]))
goto end_branch_2
} else {

}
}
{
if gopurs_runtime.Bool(v_3_1.StrVal == "GT").IntVal != 0 {
v1_4_4 := pkg_Data_Map_Internal.Call_unsafeSplit(comp_0_loop, k_1_loop, (*[1024]gopurs_runtime.Value)(m_2_loop.UnsafePtr)[5])
_ = v1_4_4
__t2 = gopurs_runtime.Constructor3("Split", (*[1024]gopurs_runtime.Value)(v1_4_4.UnsafePtr)[0], pkg_Data_Map_Internal.Call_unsafeBalancedNode((*[1024]gopurs_runtime.Value)(m_2_loop.UnsafePtr)[2], (*[1024]gopurs_runtime.Value)(m_2_loop.UnsafePtr)[3], (*[1024]gopurs_runtime.Value)(m_2_loop.UnsafePtr)[4], (*[1024]gopurs_runtime.Value)(v1_4_4.UnsafePtr)[1]), (*[1024]gopurs_runtime.Value)(v1_4_4.UnsafePtr)[2])
goto end_branch_2
} else {

}
}
{
if gopurs_runtime.Bool(v_3_1.StrVal == "EQ").IntVal != 0 {
__t2 = gopurs_runtime.Constructor3("Split", gopurs_runtime.Constructor1("Just", (*[1024]gopurs_runtime.Value)(m_2_loop.UnsafePtr)[3]), (*[1024]gopurs_runtime.Value)(m_2_loop.UnsafePtr)[4], (*[1024]gopurs_runtime.Value)(m_2_loop.UnsafePtr)[5])
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
if gopurs_runtime.Bool(r_3_loop.StrVal == "Leaf").IntVal != 0 {
__t0 = gopurs_runtime.Constructor3("SplitLast", k_0_loop, v_1_loop, l_2_loop)
goto end_branch_0
} else {

}
}
{
if gopurs_runtime.Bool(r_3_loop.StrVal == "Node").IntVal != 0 {
v1_4_1 := pkg_Data_Map_Internal.Call_unsafeSplitLast((*[1024]gopurs_runtime.Value)(r_3_loop.UnsafePtr)[2], (*[1024]gopurs_runtime.Value)(r_3_loop.UnsafePtr)[3], (*[1024]gopurs_runtime.Value)(r_3_loop.UnsafePtr)[4], (*[1024]gopurs_runtime.Value)(r_3_loop.UnsafePtr)[5])
_ = v1_4_1
__t0 = gopurs_runtime.Constructor3("SplitLast", (*[1024]gopurs_runtime.Value)(v1_4_1.UnsafePtr)[0], (*[1024]gopurs_runtime.Value)(v1_4_1.UnsafePtr)[1], pkg_Data_Map_Internal.Call_unsafeBalancedNode(k_0_loop, v_1_loop, l_2_loop, (*[1024]gopurs_runtime.Value)(v1_4_1.UnsafePtr)[2]))
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

func Call_unsafeJoinNodes(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if gopurs_runtime.Bool(v_0.StrVal == "Leaf").IntVal != 0 {
__t0 = v1_1
goto end_branch_0
} else {

}
}
{
if gopurs_runtime.Bool(v_0.StrVal == "Node").IntVal != 0 {
v2_2_1 := pkg_Data_Map_Internal.Call_unsafeSplitLast((*[1024]gopurs_runtime.Value)(v_0.UnsafePtr)[2], (*[1024]gopurs_runtime.Value)(v_0.UnsafePtr)[3], (*[1024]gopurs_runtime.Value)(v_0.UnsafePtr)[4], (*[1024]gopurs_runtime.Value)(v_0.UnsafePtr)[5])
_ = v2_2_1
__t0 = pkg_Data_Map_Internal.Call_unsafeBalancedNode((*[1024]gopurs_runtime.Value)(v2_2_1.UnsafePtr)[0], (*[1024]gopurs_runtime.Value)(v2_2_1.UnsafePtr)[1], (*[1024]gopurs_runtime.Value)(v2_2_1.UnsafePtr)[2], v1_1)
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
if gopurs_runtime.Bool(l_1_loop.StrVal == "Leaf").IntVal != 0 {
__t0 = gopurs_runtime.Constructor0("Leaf")
goto end_branch_0
} else {

}
}
{
if gopurs_runtime.Bool(r_2_loop.StrVal == "Leaf").IntVal != 0 {
__t0 = l_1_loop
goto end_branch_0
} else {

}
}
{
if gopurs_runtime.Bool(r_2_loop.StrVal == "Node").IntVal != 0 {
v_3_1 := pkg_Data_Map_Internal.Call_unsafeSplit(comp_0_loop, (*[1024]gopurs_runtime.Value)(r_2_loop.UnsafePtr)[2], l_1_loop)
_ = v_3_1
__t0 = pkg_Data_Map_Internal.Call_unsafeJoinNodes(pkg_Data_Map_Internal.Call_unsafeDifference(comp_0_loop, (*[1024]gopurs_runtime.Value)(v_3_1.UnsafePtr)[1], (*[1024]gopurs_runtime.Value)(r_2_loop.UnsafePtr)[4]), pkg_Data_Map_Internal.Call_unsafeDifference(comp_0_loop, (*[1024]gopurs_runtime.Value)(v_3_1.UnsafePtr)[2], (*[1024]gopurs_runtime.Value)(r_2_loop.UnsafePtr)[5]))
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
if gopurs_runtime.Bool(l_2_loop.StrVal == "Leaf").IntVal != 0 {
__t0 = gopurs_runtime.Constructor0("Leaf")
goto end_branch_0
} else {

}
}
{
if gopurs_runtime.Bool(r_3_loop.StrVal == "Leaf").IntVal != 0 {
__t0 = gopurs_runtime.Constructor0("Leaf")
goto end_branch_0
} else {

}
}
{
if gopurs_runtime.Bool(r_3_loop.StrVal == "Node").IntVal != 0 {
v_4_1 := pkg_Data_Map_Internal.Call_unsafeSplit(comp_0_loop, (*[1024]gopurs_runtime.Value)(r_3_loop.UnsafePtr)[2], l_2_loop)
_ = v_4_1
l_prime_5_2 := pkg_Data_Map_Internal.Call_unsafeIntersectionWith(comp_0_loop, app_1_loop, (*[1024]gopurs_runtime.Value)(v_4_1.UnsafePtr)[1], (*[1024]gopurs_runtime.Value)(r_3_loop.UnsafePtr)[4])
_ = l_prime_5_2
r_prime_6_3 := pkg_Data_Map_Internal.Call_unsafeIntersectionWith(comp_0_loop, app_1_loop, (*[1024]gopurs_runtime.Value)(v_4_1.UnsafePtr)[2], (*[1024]gopurs_runtime.Value)(r_3_loop.UnsafePtr)[5])
_ = r_prime_6_3
var __t4 gopurs_runtime.Value
{
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)(v_4_1.UnsafePtr)[0].StrVal == "Just").IntVal != 0 {
__t4 = pkg_Data_Map_Internal.Call_unsafeBalancedNode((*[1024]gopurs_runtime.Value)(r_3_loop.UnsafePtr)[2], gopurs_runtime.Apply2(app_1_loop, (*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)(v_4_1.UnsafePtr)[0].UnsafePtr)[0], (*[1024]gopurs_runtime.Value)(r_3_loop.UnsafePtr)[3]), l_prime_5_2, r_prime_6_3)
goto end_branch_4
} else {

}
}
{
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)(v_4_1.UnsafePtr)[0].StrVal == "Nothing").IntVal != 0 {
__t4 = pkg_Data_Map_Internal.Call_unsafeJoinNodes(l_prime_5_2, r_prime_6_3)
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
if gopurs_runtime.Bool(l_2_loop.StrVal == "Leaf").IntVal != 0 {
__t0 = r_3_loop
goto end_branch_0
} else {

}
}
{
if gopurs_runtime.Bool(r_3_loop.StrVal == "Leaf").IntVal != 0 {
__t0 = l_2_loop
goto end_branch_0
} else {

}
}
{
if gopurs_runtime.Bool(r_3_loop.StrVal == "Node").IntVal != 0 {
v_4_1 := pkg_Data_Map_Internal.Call_unsafeSplit(comp_0_loop, (*[1024]gopurs_runtime.Value)(r_3_loop.UnsafePtr)[2], l_2_loop)
_ = v_4_1
l_prime_5_2 := pkg_Data_Map_Internal.Call_unsafeUnionWith(comp_0_loop, app_1_loop, (*[1024]gopurs_runtime.Value)(v_4_1.UnsafePtr)[1], (*[1024]gopurs_runtime.Value)(r_3_loop.UnsafePtr)[4])
_ = l_prime_5_2
r_prime_6_3 := pkg_Data_Map_Internal.Call_unsafeUnionWith(comp_0_loop, app_1_loop, (*[1024]gopurs_runtime.Value)(v_4_1.UnsafePtr)[2], (*[1024]gopurs_runtime.Value)(r_3_loop.UnsafePtr)[5])
_ = r_prime_6_3
var __t4 gopurs_runtime.Value
{
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)(v_4_1.UnsafePtr)[0].StrVal == "Just").IntVal != 0 {
__t4 = pkg_Data_Map_Internal.Call_unsafeBalancedNode((*[1024]gopurs_runtime.Value)(r_3_loop.UnsafePtr)[2], gopurs_runtime.Apply2(app_1_loop, (*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)(v_4_1.UnsafePtr)[0].UnsafePtr)[0], (*[1024]gopurs_runtime.Value)(r_3_loop.UnsafePtr)[3]), l_prime_5_2, r_prime_6_3)
goto end_branch_4
} else {

}
}
{
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)(v_4_1.UnsafePtr)[0].StrVal == "Nothing").IntVal != 0 {
__t4 = pkg_Data_Map_Internal.Call_unsafeBalancedNode((*[1024]gopurs_runtime.Value)(r_3_loop.UnsafePtr)[2], (*[1024]gopurs_runtime.Value)(r_3_loop.UnsafePtr)[3], l_prime_5_2, r_prime_6_3)
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

func Call_go__2_0(m_prime_3 gopurs_runtime.Value, z_prime_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if gopurs_runtime.Bool(m_prime_3.StrVal == "Leaf").IntVal != 0 {
__t1 = z_prime_4
goto end_branch_1
} else {

}
}
{
if gopurs_runtime.Bool(m_prime_3.StrVal == "Node").IntVal != 0 {
__t1 = gopurs_runtime.UncurriedApp2(go__2_0, (*[1024]gopurs_runtime.Value)(m_prime_3.UnsafePtr)[4], gopurs_runtime.Apply2(f_0, (*[1024]gopurs_runtime.Value)(m_prime_3.UnsafePtr)[3], gopurs_runtime.UncurriedApp2(go__2_0, (*[1024]gopurs_runtime.Value)(m_prime_3.UnsafePtr)[5], z_prime_4)))
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
if gopurs_runtime.Bool(m_prime_4.StrVal == "Leaf").IntVal != 0 {
__t3 = z_prime_3
goto end_branch_3
} else {

}
}
{
if gopurs_runtime.Bool(m_prime_4.StrVal == "Node").IntVal != 0 {
__t3 = gopurs_runtime.UncurriedApp2(go__2_2, gopurs_runtime.Apply2(f_0, gopurs_runtime.UncurriedApp2(go__2_2, z_prime_3, (*[1024]gopurs_runtime.Value)(m_prime_4.UnsafePtr)[4]), (*[1024]gopurs_runtime.Value)(m_prime_4.UnsafePtr)[3]), (*[1024]gopurs_runtime.Value)(m_prime_4.UnsafePtr)[5])
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
if gopurs_runtime.Bool(m_prime_3.StrVal == "Leaf").IntVal != 0 {
__t1 = z_prime_4
goto end_branch_1
} else {

}
}
{
if gopurs_runtime.Bool(m_prime_3.StrVal == "Node").IntVal != 0 {
__t1 = gopurs_runtime.UncurriedApp2(go__2_0, (*[1024]gopurs_runtime.Value)(m_prime_3.UnsafePtr)[4], gopurs_runtime.Apply3(f_0, (*[1024]gopurs_runtime.Value)(m_prime_3.UnsafePtr)[2], (*[1024]gopurs_runtime.Value)(m_prime_3.UnsafePtr)[3], gopurs_runtime.UncurriedApp2(go__2_0, (*[1024]gopurs_runtime.Value)(m_prime_3.UnsafePtr)[5], z_prime_4)))
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
if gopurs_runtime.Bool(m_prime_4.StrVal == "Leaf").IntVal != 0 {
__t3 = z_prime_3
goto end_branch_3
} else {

}
}
{
if gopurs_runtime.Bool(m_prime_4.StrVal == "Node").IntVal != 0 {
__t3 = gopurs_runtime.UncurriedApp2(go__2_2, gopurs_runtime.Apply3(f_0, (*[1024]gopurs_runtime.Value)(m_prime_4.UnsafePtr)[2], gopurs_runtime.UncurriedApp2(go__2_2, z_prime_3, (*[1024]gopurs_runtime.Value)(m_prime_4.UnsafePtr)[4]), (*[1024]gopurs_runtime.Value)(m_prime_4.UnsafePtr)[3]), (*[1024]gopurs_runtime.Value)(m_prime_4.UnsafePtr)[5])
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
if gopurs_runtime.Bool(m_prime_1.StrVal == "Leaf").IntVal != 0 {
__t1 = z_prime_2
goto end_branch_1
} else {

}
}
{
if gopurs_runtime.Bool(m_prime_1.StrVal == "Node").IntVal != 0 {
__t1 = gopurs_runtime.UncurriedApp2(go__0_0, (*[1024]gopurs_runtime.Value)(m_prime_1.UnsafePtr)[4], gopurs_runtime.Constructor2("Cons", (*[1024]gopurs_runtime.Value)(m_prime_1.UnsafePtr)[2], gopurs_runtime.UncurriedApp2(go__0_0, (*[1024]gopurs_runtime.Value)(m_prime_1.UnsafePtr)[5], z_prime_2)))
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
if gopurs_runtime.Bool(m_prime_1.StrVal == "Leaf").IntVal != 0 {
__t1 = z_prime_2
goto end_branch_1
} else {

}
}
{
if gopurs_runtime.Bool(m_prime_1.StrVal == "Node").IntVal != 0 {
__t1 = gopurs_runtime.UncurriedApp2(go__0_0, (*[1024]gopurs_runtime.Value)(m_prime_1.UnsafePtr)[4], gopurs_runtime.Constructor2("Cons", (*[1024]gopurs_runtime.Value)(m_prime_1.UnsafePtr)[3], gopurs_runtime.UncurriedApp2(go__0_0, (*[1024]gopurs_runtime.Value)(m_prime_1.UnsafePtr)[5], z_prime_2)))
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


