package Data_Map_Internal

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_Function_Uncurried "gopurs/output/Data.Function.Uncurried"
	pkg_Data_Semiring "gopurs/output/Data.Semiring"
	pkg_Data_Ord "gopurs/output/Data.Ord"
	pkg_Data_Unit "gopurs/output/Data.Unit"
	pkg_Data_Function "gopurs/output/Data.Function"
	pkg_Data_Semigroup "gopurs/output/Data.Semigroup"
	pkg_Data_Show "gopurs/output/Data.Show"
	pkg_Data_HeytingAlgebra "gopurs/output/Data.HeytingAlgebra"
	pkg_Data_Unfoldable "gopurs/output/Data.Unfoldable"
	pkg_Data_Tuple "gopurs/output/Data.Tuple"
	pkg_Control_Category "gopurs/output/Control.Category"
	pkg_Data_List_Types "gopurs/output/Data.List.Types"
	pkg_Data_Eq "gopurs/output/Data.Eq"
	pkg_Data_Ring "gopurs/output/Data.Ring"
)

var Leaf gopurs_runtime.Value
var once_Leaf sync.Once
func Get_Leaf() gopurs_runtime.Value {
	once_Leaf.Do(func() {
		Leaf = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Leaf")})
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
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Node"), "value0": value0, "value1": value1, "value2": value2, "value3": value3, "value4": value4, "value5": value5})
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
		IterLeaf = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("IterLeaf")})
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
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("IterEmit"), "value0": value0, "value1": value1, "value2": value2})
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
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("IterNode"), "value0": value0, "value1": value1})
})
})
	})
	return IterNode
}

var IterDone gopurs_runtime.Value
var once_IterDone sync.Once
func Get_IterDone() gopurs_runtime.Value {
	once_IterDone.Do(func() {
		IterDone = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("IterDone")})
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
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("IterNext"), "value0": value0, "value1": value1, "value2": value2})
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
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Split"), "value0": value0, "value1": value1, "value2": value2})
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
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("SplitLast"), "value0": value0, "value1": value1, "value2": value2})
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
		unsafeNode = gopurs_runtime.Apply(pkg_Data_Function_Uncurried.Get_mkFn4(), gopurs_runtime.Func(func(k_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(l_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(r_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(l_2.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Leaf")).IntVal != 0 {
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(r_3.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Leaf")).IntVal != 0 {
__t1 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Node"), "value0": gopurs_runtime.Int(1), "value1": gopurs_runtime.Int(1), "value2": k_0, "value3": v_1, "value4": l_2, "value5": r_3})
goto end_branch_1
} else {

}
}
{
if (gopurs_runtime.Bool(r_3.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Node")).IntVal != 0 {
__t1 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Node"), "value0": gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Semiring.Get_intAdd(), gopurs_runtime.Int(1)), r_3.PtrVal.(map[string]gopurs_runtime.Value)["value0"]), "value1": gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Semiring.Get_intAdd(), gopurs_runtime.Int(1)), r_3.PtrVal.(map[string]gopurs_runtime.Value)["value1"]), "value2": k_0, "value3": v_1, "value4": l_2, "value5": r_3})
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
if (gopurs_runtime.Bool(l_2.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Node")).IntVal != 0 {
var __t2 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(r_3.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Leaf")).IntVal != 0 {
__t2 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Node"), "value0": gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Semiring.Get_intAdd(), gopurs_runtime.Int(1)), l_2.PtrVal.(map[string]gopurs_runtime.Value)["value0"]), "value1": gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Semiring.Get_intAdd(), gopurs_runtime.Int(1)), l_2.PtrVal.(map[string]gopurs_runtime.Value)["value1"]), "value2": k_0, "value3": v_1, "value4": l_2, "value5": r_3})
goto end_branch_2
} else {

}
}
{
if (gopurs_runtime.Bool(r_3.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Node")).IntVal != 0 {
var __t3 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Ord.Get_ordInt().PtrVal.(map[string]gopurs_runtime.Value)["compare"], l_2.PtrVal.(map[string]gopurs_runtime.Value)["value0"]), r_3.PtrVal.(map[string]gopurs_runtime.Value)["value0"]).PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "GT")).IntVal != 0 {
__t3 = l_2.PtrVal.(map[string]gopurs_runtime.Value)["value0"]
goto end_branch_3
} else {

}
}
{
__t3 = r_3.PtrVal.(map[string]gopurs_runtime.Value)["value0"]
}
end_branch_3:
__t2 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Node"), "value0": gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Semiring.Get_intAdd(), gopurs_runtime.Int(1)), __t3), "value1": gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Semiring.Get_intAdd(), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Semiring.Get_intAdd(), gopurs_runtime.Int(1)), l_2.PtrVal.(map[string]gopurs_runtime.Value)["value1"])), r_3.PtrVal.(map[string]gopurs_runtime.Value)["value1"]), "value2": k_0, "value3": v_1, "value4": l_2, "value5": r_3})
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
})
})
})
}))
	})
	return unsafeNode
}

var toMapIter gopurs_runtime.Value
var once_toMapIter sync.Once
func Get_toMapIter() gopurs_runtime.Value {
	once_toMapIter.Do(func() {
		toMapIter = gopurs_runtime.Func(func(a_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("IterNode"), "value0": a_0, "value1": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("IterLeaf")})})
})
	})
	return toMapIter
}

var stepWith gopurs_runtime.Value
var once_stepWith sync.Once
func Get_stepWith() gopurs_runtime.Value {
	once_stepWith.Do(func() {
		stepWith = gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(next_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(done_2 gopurs_runtime.Value) gopurs_runtime.Value {
var go__3_0 gopurs_runtime.Value
go__3_0 = gopurs_runtime.Func(func(v_4_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
go__3_0:
for {
if false { continue go__3_0 }
var v_4 = v_4_loop
_ = v_4
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(v_4.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "IterLeaf")).IntVal != 0 {
__t1 = gopurs_runtime.Apply(done_2, pkg_Data_Unit.Get_unit())
goto end_branch_1
} else {

}
}
{
if (gopurs_runtime.Bool(v_4.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "IterEmit")).IntVal != 0 {
__t1 = gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Function_Uncurried.Get_runFn3(), next_1), v_4.PtrVal.(map[string]gopurs_runtime.Value)["value0"]), v_4.PtrVal.(map[string]gopurs_runtime.Value)["value1"]), v_4.PtrVal.(map[string]gopurs_runtime.Value)["value2"])
goto end_branch_1
} else {

}
}
{
if (gopurs_runtime.Bool(v_4.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "IterNode")).IntVal != 0 {
v_4_loop = gopurs_runtime.Apply(gopurs_runtime.Apply(f_0, v_4.PtrVal.(map[string]gopurs_runtime.Value)["value1"]), v_4.PtrVal.(map[string]gopurs_runtime.Value)["value0"])
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
if (gopurs_runtime.Bool(v_0.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Leaf")).IntVal != 0 {
__t0 = gopurs_runtime.Int(0)
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Bool(v_0.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Node")).IntVal != 0 {
__t0 = v_0.PtrVal.(map[string]gopurs_runtime.Value)["value1"]
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
		singleton = gopurs_runtime.Func(func(k_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Node"), "value0": gopurs_runtime.Int(1), "value1": gopurs_runtime.Int(1), "value2": k_0, "value3": v_1, "value4": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Leaf")}), "value5": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Leaf")})})
})
})
	})
	return singleton
}

var unsafeBalancedNode gopurs_runtime.Value
var once_unsafeBalancedNode sync.Once
func Get_unsafeBalancedNode() gopurs_runtime.Value {
	once_unsafeBalancedNode.Do(func() {
		unsafeBalancedNode = func() gopurs_runtime.Value {
rotateLeft_0_0 := gopurs_runtime.Apply(pkg_Data_Function_Uncurried.Get_mkFn7(), gopurs_runtime.Func(func(k_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(l_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(rk_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(rv_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(rl_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(rr_6 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
var __t2 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(rr_6.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Leaf")).IntVal != 0 {
__t2 = gopurs_runtime.Int(0)
goto end_branch_2
} else {

}
}
{
if (gopurs_runtime.Bool(rr_6.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Node")).IntVal != 0 {
__t2 = rr_6.PtrVal.(map[string]gopurs_runtime.Value)["value0"]
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_2:
if (gopurs_runtime.Bool(gopurs_runtime.Bool(rl_5.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Node").IntVal != 0 && gopurs_runtime.Bool(gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Ord.Get_ordInt().PtrVal.(map[string]gopurs_runtime.Value)["compare"], rl_5.PtrVal.(map[string]gopurs_runtime.Value)["value0"]), __t2).PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "GT").IntVal != 0)).IntVal != 0 {
__t1 = gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Function_Uncurried.Get_runFn4(), Get_unsafeNode()), rl_5.PtrVal.(map[string]gopurs_runtime.Value)["value2"]), rl_5.PtrVal.(map[string]gopurs_runtime.Value)["value3"]), gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Function_Uncurried.Get_runFn4(), Get_unsafeNode()), k_0), v_1), l_2), rl_5.PtrVal.(map[string]gopurs_runtime.Value)["value4"])), gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Function_Uncurried.Get_runFn4(), Get_unsafeNode()), rk_3), rv_4), rl_5.PtrVal.(map[string]gopurs_runtime.Value)["value5"]), rr_6))
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Function_Uncurried.Get_runFn4(), Get_unsafeNode()), rk_3), rv_4), gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Function_Uncurried.Get_runFn4(), Get_unsafeNode()), k_0), v_1), l_2), rl_5)), rr_6)
}
end_branch_1:
return __t1
})
})
})
})
})
})
}))
rotateRight_1_3 := gopurs_runtime.Apply(pkg_Data_Function_Uncurried.Get_mkFn7(), gopurs_runtime.Func(func(k_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(lk_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(lv_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(ll_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(lr_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(r_7 gopurs_runtime.Value) gopurs_runtime.Value {
var __t4 gopurs_runtime.Value
{
var __t5 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(ll_5.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Leaf")).IntVal != 0 {
__t5 = gopurs_runtime.Int(0)
goto end_branch_5
} else {

}
}
{
if (gopurs_runtime.Bool(ll_5.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Node")).IntVal != 0 {
__t5 = ll_5.PtrVal.(map[string]gopurs_runtime.Value)["value0"]
goto end_branch_5
} else {

}
}
{
__t5 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_5:
if (gopurs_runtime.Bool(gopurs_runtime.Bool(lr_6.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Node").IntVal != 0 && gopurs_runtime.Bool(gopurs_runtime.Bool(gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Ord.Get_ordInt().PtrVal.(map[string]gopurs_runtime.Value)["compare"], __t5), lr_6.PtrVal.(map[string]gopurs_runtime.Value)["value0"]).PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "GT").IntVal == 0).IntVal != 0)).IntVal != 0 {
__t4 = gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Function_Uncurried.Get_runFn4(), Get_unsafeNode()), lr_6.PtrVal.(map[string]gopurs_runtime.Value)["value2"]), lr_6.PtrVal.(map[string]gopurs_runtime.Value)["value3"]), gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Function_Uncurried.Get_runFn4(), Get_unsafeNode()), lk_3), lv_4), ll_5), lr_6.PtrVal.(map[string]gopurs_runtime.Value)["value4"])), gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Function_Uncurried.Get_runFn4(), Get_unsafeNode()), k_1), v_2), lr_6.PtrVal.(map[string]gopurs_runtime.Value)["value5"]), r_7))
goto end_branch_4
} else {

}
}
{
__t4 = gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Function_Uncurried.Get_runFn4(), Get_unsafeNode()), lk_3), lv_4), ll_5), gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Function_Uncurried.Get_runFn4(), Get_unsafeNode()), k_1), v_2), lr_6), r_7))
}
end_branch_4:
return __t4
})
})
})
})
})
})
}))
return gopurs_runtime.Apply(pkg_Data_Function_Uncurried.Get_mkFn4(), gopurs_runtime.Func(func(k_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(l_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(r_5 gopurs_runtime.Value) gopurs_runtime.Value {
var __t6 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(l_4.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Leaf")).IntVal != 0 {
var __t7 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(r_5.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Leaf")).IntVal != 0 {
__t7 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Node"), "value0": gopurs_runtime.Int(1), "value1": gopurs_runtime.Int(1), "value2": k_2, "value3": v_3, "value4": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Leaf")}), "value5": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Leaf")})})
goto end_branch_7
} else {

}
}
{
if (gopurs_runtime.Bool(gopurs_runtime.Bool(r_5.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Node").IntVal != 0 && gopurs_runtime.Bool(gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Ord.Get_ordInt().PtrVal.(map[string]gopurs_runtime.Value)["compare"], r_5.PtrVal.(map[string]gopurs_runtime.Value)["value0"]), gopurs_runtime.Int(1)).PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "GT").IntVal != 0)).IntVal != 0 {
__t7 = gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Function_Uncurried.Get_runFn7(), rotateLeft_0_0), k_2), v_3), l_4), r_5.PtrVal.(map[string]gopurs_runtime.Value)["value2"]), r_5.PtrVal.(map[string]gopurs_runtime.Value)["value3"]), r_5.PtrVal.(map[string]gopurs_runtime.Value)["value4"]), r_5.PtrVal.(map[string]gopurs_runtime.Value)["value5"])
goto end_branch_7
} else {

}
}
{
__t7 = gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Function_Uncurried.Get_runFn4(), Get_unsafeNode()), k_2), v_3), l_4), r_5)
}
end_branch_7:
__t6 = __t7
goto end_branch_6
} else {

}
}
{
if (gopurs_runtime.Bool(l_4.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Node")).IntVal != 0 {
var __t8 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(r_5.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Node")).IntVal != 0 {
var __t9 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Ord.Get_ordInt().PtrVal.(map[string]gopurs_runtime.Value)["compare"], r_5.PtrVal.(map[string]gopurs_runtime.Value)["value0"]), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Semiring.Get_intAdd(), l_4.PtrVal.(map[string]gopurs_runtime.Value)["value0"]), gopurs_runtime.Int(1))).PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "GT")).IntVal != 0 {
__t9 = gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Function_Uncurried.Get_runFn7(), rotateLeft_0_0), k_2), v_3), l_4), r_5.PtrVal.(map[string]gopurs_runtime.Value)["value2"]), r_5.PtrVal.(map[string]gopurs_runtime.Value)["value3"]), r_5.PtrVal.(map[string]gopurs_runtime.Value)["value4"]), r_5.PtrVal.(map[string]gopurs_runtime.Value)["value5"])
goto end_branch_9
} else {

}
}
{
if (gopurs_runtime.Bool(gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Ord.Get_ordInt().PtrVal.(map[string]gopurs_runtime.Value)["compare"], l_4.PtrVal.(map[string]gopurs_runtime.Value)["value0"]), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Semiring.Get_intAdd(), r_5.PtrVal.(map[string]gopurs_runtime.Value)["value0"]), gopurs_runtime.Int(1))).PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "GT")).IntVal != 0 {
__t9 = gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Function_Uncurried.Get_runFn7(), rotateRight_1_3), k_2), v_3), l_4.PtrVal.(map[string]gopurs_runtime.Value)["value2"]), l_4.PtrVal.(map[string]gopurs_runtime.Value)["value3"]), l_4.PtrVal.(map[string]gopurs_runtime.Value)["value4"]), l_4.PtrVal.(map[string]gopurs_runtime.Value)["value5"]), r_5)
goto end_branch_9
} else {

}
}
{
__t9 = gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Function_Uncurried.Get_runFn4(), Get_unsafeNode()), k_2), v_3), l_4), r_5)
}
end_branch_9:
__t8 = __t9
goto end_branch_8
} else {

}
}
{
if (gopurs_runtime.Bool(gopurs_runtime.Bool(r_5.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Leaf").IntVal != 0 && gopurs_runtime.Bool(gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Ord.Get_ordInt().PtrVal.(map[string]gopurs_runtime.Value)["compare"], l_4.PtrVal.(map[string]gopurs_runtime.Value)["value0"]), gopurs_runtime.Int(1)).PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "GT").IntVal != 0)).IntVal != 0 {
__t8 = gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Function_Uncurried.Get_runFn7(), rotateRight_1_3), k_2), v_3), l_4.PtrVal.(map[string]gopurs_runtime.Value)["value2"]), l_4.PtrVal.(map[string]gopurs_runtime.Value)["value3"]), l_4.PtrVal.(map[string]gopurs_runtime.Value)["value4"]), l_4.PtrVal.(map[string]gopurs_runtime.Value)["value5"]), r_5)
goto end_branch_8
} else {

}
}
{
__t8 = gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Function_Uncurried.Get_runFn4(), Get_unsafeNode()), k_2), v_3), l_4), r_5)
}
end_branch_8:
__t6 = __t8
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
})
})
}))
}()
	})
	return unsafeBalancedNode
}

var unsafeSplit gopurs_runtime.Value
var once_unsafeSplit sync.Once
func Get_unsafeSplit() gopurs_runtime.Value {
	once_unsafeSplit.Do(func() {
		unsafeSplit = gopurs_runtime.Apply(pkg_Data_Function_Uncurried.Get_mkFn3(), gopurs_runtime.Func(func(comp_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(m_2.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Leaf")).IntVal != 0 {
__t0 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Split"), "value0": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")}), "value1": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Leaf")}), "value2": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Leaf")})})
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Bool(m_2.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Node")).IntVal != 0 {
v_3_1 := gopurs_runtime.Apply(gopurs_runtime.Apply(comp_0, k_1), m_2.PtrVal.(map[string]gopurs_runtime.Value)["value2"])
var __t2 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(v_3_1.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "LT")).IntVal != 0 {
v1_4_3 := gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Function_Uncurried.Get_runFn3(), Get_unsafeSplit()), comp_0), k_1), m_2.PtrVal.(map[string]gopurs_runtime.Value)["value4"])
__t2 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Split"), "value0": v1_4_3.PtrVal.(map[string]gopurs_runtime.Value)["value0"], "value1": v1_4_3.PtrVal.(map[string]gopurs_runtime.Value)["value1"], "value2": gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Function_Uncurried.Get_runFn4(), Get_unsafeBalancedNode()), m_2.PtrVal.(map[string]gopurs_runtime.Value)["value2"]), m_2.PtrVal.(map[string]gopurs_runtime.Value)["value3"]), v1_4_3.PtrVal.(map[string]gopurs_runtime.Value)["value2"]), m_2.PtrVal.(map[string]gopurs_runtime.Value)["value5"])})
goto end_branch_2
} else {

}
}
{
if (gopurs_runtime.Bool(v_3_1.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "GT")).IntVal != 0 {
v1_4_4 := gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Function_Uncurried.Get_runFn3(), Get_unsafeSplit()), comp_0), k_1), m_2.PtrVal.(map[string]gopurs_runtime.Value)["value5"])
__t2 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Split"), "value0": v1_4_4.PtrVal.(map[string]gopurs_runtime.Value)["value0"], "value1": gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Function_Uncurried.Get_runFn4(), Get_unsafeBalancedNode()), m_2.PtrVal.(map[string]gopurs_runtime.Value)["value2"]), m_2.PtrVal.(map[string]gopurs_runtime.Value)["value3"]), m_2.PtrVal.(map[string]gopurs_runtime.Value)["value4"]), v1_4_4.PtrVal.(map[string]gopurs_runtime.Value)["value1"]), "value2": v1_4_4.PtrVal.(map[string]gopurs_runtime.Value)["value2"]})
goto end_branch_2
} else {

}
}
{
if (gopurs_runtime.Bool(v_3_1.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "EQ")).IntVal != 0 {
__t2 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Split"), "value0": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Just"), "value0": m_2.PtrVal.(map[string]gopurs_runtime.Value)["value3"]}), "value1": m_2.PtrVal.(map[string]gopurs_runtime.Value)["value4"], "value2": m_2.PtrVal.(map[string]gopurs_runtime.Value)["value5"]})
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
})
})
}))
	})
	return unsafeSplit
}

var unsafeSplitLast gopurs_runtime.Value
var once_unsafeSplitLast sync.Once
func Get_unsafeSplitLast() gopurs_runtime.Value {
	once_unsafeSplitLast.Do(func() {
		unsafeSplitLast = gopurs_runtime.Apply(pkg_Data_Function_Uncurried.Get_mkFn4(), gopurs_runtime.Func(func(k_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(l_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(r_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(r_3.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Leaf")).IntVal != 0 {
__t0 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("SplitLast"), "value0": k_0, "value1": v_1, "value2": l_2})
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Bool(r_3.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Node")).IntVal != 0 {
v1_4_1 := gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Function_Uncurried.Get_runFn4(), Get_unsafeSplitLast()), r_3.PtrVal.(map[string]gopurs_runtime.Value)["value2"]), r_3.PtrVal.(map[string]gopurs_runtime.Value)["value3"]), r_3.PtrVal.(map[string]gopurs_runtime.Value)["value4"]), r_3.PtrVal.(map[string]gopurs_runtime.Value)["value5"])
__t0 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("SplitLast"), "value0": v1_4_1.PtrVal.(map[string]gopurs_runtime.Value)["value0"], "value1": v1_4_1.PtrVal.(map[string]gopurs_runtime.Value)["value1"], "value2": gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Function_Uncurried.Get_runFn4(), Get_unsafeBalancedNode()), k_0), v_1), l_2), v1_4_1.PtrVal.(map[string]gopurs_runtime.Value)["value2"])})
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
})
}))
	})
	return unsafeSplitLast
}

var unsafeJoinNodes gopurs_runtime.Value
var once_unsafeJoinNodes sync.Once
func Get_unsafeJoinNodes() gopurs_runtime.Value {
	once_unsafeJoinNodes.Do(func() {
		unsafeJoinNodes = gopurs_runtime.Apply(pkg_Data_Function_Uncurried.Get_mkFn2(), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(v_0.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Leaf")).IntVal != 0 {
__t0 = v1_1
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Bool(v_0.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Node")).IntVal != 0 {
v2_2_1 := gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Function_Uncurried.Get_runFn4(), Get_unsafeSplitLast()), v_0.PtrVal.(map[string]gopurs_runtime.Value)["value2"]), v_0.PtrVal.(map[string]gopurs_runtime.Value)["value3"]), v_0.PtrVal.(map[string]gopurs_runtime.Value)["value4"]), v_0.PtrVal.(map[string]gopurs_runtime.Value)["value5"])
__t0 = gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Function_Uncurried.Get_runFn4(), Get_unsafeBalancedNode()), v2_2_1.PtrVal.(map[string]gopurs_runtime.Value)["value0"]), v2_2_1.PtrVal.(map[string]gopurs_runtime.Value)["value1"]), v2_2_1.PtrVal.(map[string]gopurs_runtime.Value)["value2"]), v1_1)
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
}))
	})
	return unsafeJoinNodes
}

var unsafeDifference gopurs_runtime.Value
var once_unsafeDifference sync.Once
func Get_unsafeDifference() gopurs_runtime.Value {
	once_unsafeDifference.Do(func() {
		unsafeDifference = gopurs_runtime.Apply(pkg_Data_Function_Uncurried.Get_mkFn3(), gopurs_runtime.Func(func(comp_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(l_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(r_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(l_1.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Leaf")).IntVal != 0 {
__t0 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Leaf")})
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Bool(r_2.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Leaf")).IntVal != 0 {
__t0 = l_1
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Bool(r_2.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Node")).IntVal != 0 {
v_3_1 := gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Function_Uncurried.Get_runFn3(), Get_unsafeSplit()), comp_0), r_2.PtrVal.(map[string]gopurs_runtime.Value)["value2"]), l_1)
__t0 = gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Function_Uncurried.Get_runFn2(), Get_unsafeJoinNodes()), gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Function_Uncurried.Get_runFn3(), Get_unsafeDifference()), comp_0), v_3_1.PtrVal.(map[string]gopurs_runtime.Value)["value1"]), r_2.PtrVal.(map[string]gopurs_runtime.Value)["value4"])), gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Function_Uncurried.Get_runFn3(), Get_unsafeDifference()), comp_0), v_3_1.PtrVal.(map[string]gopurs_runtime.Value)["value2"]), r_2.PtrVal.(map[string]gopurs_runtime.Value)["value5"]))
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
}))
	})
	return unsafeDifference
}

var unsafeIntersectionWith gopurs_runtime.Value
var once_unsafeIntersectionWith sync.Once
func Get_unsafeIntersectionWith() gopurs_runtime.Value {
	once_unsafeIntersectionWith.Do(func() {
		unsafeIntersectionWith = gopurs_runtime.Apply(pkg_Data_Function_Uncurried.Get_mkFn4(), gopurs_runtime.Func(func(comp_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(app_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(l_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(r_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(l_2.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Leaf")).IntVal != 0 {
__t0 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Leaf")})
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Bool(r_3.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Leaf")).IntVal != 0 {
__t0 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Leaf")})
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Bool(r_3.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Node")).IntVal != 0 {
v_4_1 := gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Function_Uncurried.Get_runFn3(), Get_unsafeSplit()), comp_0), r_3.PtrVal.(map[string]gopurs_runtime.Value)["value2"]), l_2)
l_prime_5_2 := gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Function_Uncurried.Get_runFn4(), Get_unsafeIntersectionWith()), comp_0), app_1), v_4_1.PtrVal.(map[string]gopurs_runtime.Value)["value1"]), r_3.PtrVal.(map[string]gopurs_runtime.Value)["value4"])
r_prime_6_3 := gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Function_Uncurried.Get_runFn4(), Get_unsafeIntersectionWith()), comp_0), app_1), v_4_1.PtrVal.(map[string]gopurs_runtime.Value)["value2"]), r_3.PtrVal.(map[string]gopurs_runtime.Value)["value5"])
var __t4 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(v_4_1.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just")).IntVal != 0 {
__t4 = gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Function_Uncurried.Get_runFn4(), Get_unsafeBalancedNode()), r_3.PtrVal.(map[string]gopurs_runtime.Value)["value2"]), gopurs_runtime.Apply(gopurs_runtime.Apply(app_1, v_4_1.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["value0"]), r_3.PtrVal.(map[string]gopurs_runtime.Value)["value3"])), l_prime_5_2), r_prime_6_3)
goto end_branch_4
} else {

}
}
{
if (gopurs_runtime.Bool(v_4_1.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Nothing")).IntVal != 0 {
__t4 = gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Function_Uncurried.Get_runFn2(), Get_unsafeJoinNodes()), l_prime_5_2), r_prime_6_3)
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
})
})
})
}))
	})
	return unsafeIntersectionWith
}

var unsafeUnionWith gopurs_runtime.Value
var once_unsafeUnionWith sync.Once
func Get_unsafeUnionWith() gopurs_runtime.Value {
	once_unsafeUnionWith.Do(func() {
		unsafeUnionWith = gopurs_runtime.Apply(pkg_Data_Function_Uncurried.Get_mkFn4(), gopurs_runtime.Func(func(comp_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(app_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(l_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(r_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(l_2.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Leaf")).IntVal != 0 {
__t0 = r_3
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Bool(r_3.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Leaf")).IntVal != 0 {
__t0 = l_2
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Bool(r_3.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Node")).IntVal != 0 {
v_4_1 := gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Function_Uncurried.Get_runFn3(), Get_unsafeSplit()), comp_0), r_3.PtrVal.(map[string]gopurs_runtime.Value)["value2"]), l_2)
l_prime_5_2 := gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Function_Uncurried.Get_runFn4(), Get_unsafeUnionWith()), comp_0), app_1), v_4_1.PtrVal.(map[string]gopurs_runtime.Value)["value1"]), r_3.PtrVal.(map[string]gopurs_runtime.Value)["value4"])
r_prime_6_3 := gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Function_Uncurried.Get_runFn4(), Get_unsafeUnionWith()), comp_0), app_1), v_4_1.PtrVal.(map[string]gopurs_runtime.Value)["value2"]), r_3.PtrVal.(map[string]gopurs_runtime.Value)["value5"])
var __t4 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(v_4_1.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just")).IntVal != 0 {
__t4 = gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Function_Uncurried.Get_runFn4(), Get_unsafeBalancedNode()), r_3.PtrVal.(map[string]gopurs_runtime.Value)["value2"]), gopurs_runtime.Apply(gopurs_runtime.Apply(app_1, v_4_1.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["value0"]), r_3.PtrVal.(map[string]gopurs_runtime.Value)["value3"])), l_prime_5_2), r_prime_6_3)
goto end_branch_4
} else {

}
}
{
if (gopurs_runtime.Bool(v_4_1.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Nothing")).IntVal != 0 {
__t4 = gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Function_Uncurried.Get_runFn4(), Get_unsafeBalancedNode()), r_3.PtrVal.(map[string]gopurs_runtime.Value)["value2"]), r_3.PtrVal.(map[string]gopurs_runtime.Value)["value3"]), l_prime_5_2), r_prime_6_3)
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
})
})
})
}))
	})
	return unsafeUnionWith
}

var unionWith gopurs_runtime.Value
var once_unionWith sync.Once
func Get_unionWith() gopurs_runtime.Value {
	once_unionWith.Do(func() {
		unionWith = gopurs_runtime.Func(func(dictOrd_0 gopurs_runtime.Value) gopurs_runtime.Value {
compare_1_0 := dictOrd_0.PtrVal.(map[string]gopurs_runtime.Value)["compare"]
return gopurs_runtime.Func(func(app_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m2_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Function_Uncurried.Get_runFn4(), Get_unsafeUnionWith()), compare_1_0), app_2), m1_3), m2_4)
})
})
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
compare_1_0 := dictOrd_0.PtrVal.(map[string]gopurs_runtime.Value)["compare"]
return gopurs_runtime.Func(func(m1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m2_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Function_Uncurried.Get_runFn4(), Get_unsafeUnionWith()), compare_1_0), pkg_Data_Function.Get_const_()), m1_2), m2_3)
})
})
})
	})
	return union
}

var update gopurs_runtime.Value
var once_update sync.Once
func Get_update() gopurs_runtime.Value {
	once_update.Do(func() {
		update = gopurs_runtime.Func(func(dictOrd_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_2 gopurs_runtime.Value) gopurs_runtime.Value {
var go__3_0 gopurs_runtime.Value
go__3_0 = gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(v_4.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Leaf")).IntVal != 0 {
__t1 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Leaf")})
goto end_branch_1
} else {

}
}
{
if (gopurs_runtime.Bool(v_4.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Node")).IntVal != 0 {
v1_5_2 := gopurs_runtime.Apply(gopurs_runtime.Apply(dictOrd_0.PtrVal.(map[string]gopurs_runtime.Value)["compare"], k_2), v_4.PtrVal.(map[string]gopurs_runtime.Value)["value2"])
var __t3 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(v1_5_2.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "LT")).IntVal != 0 {
__t3 = gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Function_Uncurried.Get_runFn4(), Get_unsafeBalancedNode()), v_4.PtrVal.(map[string]gopurs_runtime.Value)["value2"]), v_4.PtrVal.(map[string]gopurs_runtime.Value)["value3"]), gopurs_runtime.Apply(go__3_0, v_4.PtrVal.(map[string]gopurs_runtime.Value)["value4"])), v_4.PtrVal.(map[string]gopurs_runtime.Value)["value5"])
goto end_branch_3
} else {

}
}
{
if (gopurs_runtime.Bool(v1_5_2.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "GT")).IntVal != 0 {
__t3 = gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Function_Uncurried.Get_runFn4(), Get_unsafeBalancedNode()), v_4.PtrVal.(map[string]gopurs_runtime.Value)["value2"]), v_4.PtrVal.(map[string]gopurs_runtime.Value)["value3"]), v_4.PtrVal.(map[string]gopurs_runtime.Value)["value4"]), gopurs_runtime.Apply(go__3_0, v_4.PtrVal.(map[string]gopurs_runtime.Value)["value5"]))
goto end_branch_3
} else {

}
}
{
if (gopurs_runtime.Bool(v1_5_2.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "EQ")).IntVal != 0 {
v2_6_4 := gopurs_runtime.Apply(f_1, v_4.PtrVal.(map[string]gopurs_runtime.Value)["value3"])
var __t5 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(v2_6_4.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Nothing")).IntVal != 0 {
__t5 = gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Function_Uncurried.Get_runFn2(), Get_unsafeJoinNodes()), v_4.PtrVal.(map[string]gopurs_runtime.Value)["value4"]), v_4.PtrVal.(map[string]gopurs_runtime.Value)["value5"])
goto end_branch_5
} else {

}
}
{
if (gopurs_runtime.Bool(v2_6_4.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just")).IntVal != 0 {
__t5 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Node"), "value0": v_4.PtrVal.(map[string]gopurs_runtime.Value)["value0"], "value1": v_4.PtrVal.(map[string]gopurs_runtime.Value)["value1"], "value2": v_4.PtrVal.(map[string]gopurs_runtime.Value)["value2"], "value3": v2_6_4.PtrVal.(map[string]gopurs_runtime.Value)["value0"], "value4": v_4.PtrVal.(map[string]gopurs_runtime.Value)["value4"], "value5": v_4.PtrVal.(map[string]gopurs_runtime.Value)["value5"]})
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
})
	})
	return update
}

var showTree gopurs_runtime.Value
var once_showTree sync.Once
func Get_showTree() gopurs_runtime.Value {
	once_showTree.Do(func() {
		showTree = gopurs_runtime.Func(func(dictShow_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(dictShow1_1 gopurs_runtime.Value) gopurs_runtime.Value {
var go__2_0 gopurs_runtime.Value
go__2_0 = gopurs_runtime.Func(func(ind_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(v_4.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Leaf")).IntVal != 0 {
__t1 = gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Semigroup.Get_concatString(), ind_3), gopurs_runtime.Str("Leaf"))
goto end_branch_1
} else {

}
}
{
if (gopurs_runtime.Bool(v_4.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Node")).IntVal != 0 {
__t1 = gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Semigroup.Get_concatString(), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Semigroup.Get_concatString(), ind_3), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Semigroup.Get_concatString(), gopurs_runtime.Str("[")), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Semigroup.Get_concatString(), gopurs_runtime.Apply(pkg_Data_Show.Get_showIntImpl(), v_4.PtrVal.(map[string]gopurs_runtime.Value)["value0"])), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Semigroup.Get_concatString(), gopurs_runtime.Str("] ")), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Semigroup.Get_concatString(), gopurs_runtime.Apply(dictShow_0.PtrVal.(map[string]gopurs_runtime.Value)["show"], v_4.PtrVal.(map[string]gopurs_runtime.Value)["value2"])), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Semigroup.Get_concatString(), gopurs_runtime.Str(" => ")), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Semigroup.Get_concatString(), gopurs_runtime.Apply(dictShow1_1.PtrVal.(map[string]gopurs_runtime.Value)["show"], v_4.PtrVal.(map[string]gopurs_runtime.Value)["value3"])), gopurs_runtime.Str("\n"))))))))), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Semigroup.Get_concatString(), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Semigroup.Get_concatString(), gopurs_runtime.Apply(gopurs_runtime.Apply(go__2_0, gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Semigroup.Get_concatString(), ind_3), gopurs_runtime.Str("    "))), v_4.PtrVal.(map[string]gopurs_runtime.Value)["value4"])), gopurs_runtime.Str("\n"))), gopurs_runtime.Apply(gopurs_runtime.Apply(go__2_0, gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Semigroup.Get_concatString(), ind_3), gopurs_runtime.Str("    "))), v_4.PtrVal.(map[string]gopurs_runtime.Value)["value5"])))
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
})
return gopurs_runtime.Apply(go__2_0, gopurs_runtime.Str(""))
})
})
	})
	return showTree
}

var semigroupMap gopurs_runtime.Value
var once_semigroupMap sync.Once
func Get_semigroupMap() gopurs_runtime.Value {
	once_semigroupMap.Do(func() {
		semigroupMap = gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(dictOrd_1 gopurs_runtime.Value) gopurs_runtime.Value {
compare_2_0 := dictOrd_1.PtrVal.(map[string]gopurs_runtime.Value)["compare"]
return gopurs_runtime.Func(func(dictSemigroup_3 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_4_1 := dictSemigroup_3.PtrVal.(map[string]gopurs_runtime.Value)["append"]
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"append": gopurs_runtime.Func(func(m1_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m2_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Function_Uncurried.Get_runFn4(), Get_unsafeUnionWith()), compare_2_0), __local_var_4_1), m1_5), m2_6)
})
})})
})
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
compare_1_0 := dictOrd_0.PtrVal.(map[string]gopurs_runtime.Value)["compare"]
return gopurs_runtime.Func(func(k_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m_3 gopurs_runtime.Value) gopurs_runtime.Value {
v_4_1 := gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Function_Uncurried.Get_runFn3(), Get_unsafeSplit()), compare_1_0), k_2), m_3)
var __t2 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(v_4_1.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just")).IntVal != 0 {
__t2 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Just"), "value0": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Tuple"), "value0": v_4_1.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["value0"], "value1": gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Function_Uncurried.Get_runFn2(), Get_unsafeJoinNodes()), v_4_1.PtrVal.(map[string]gopurs_runtime.Value)["value1"]), v_4_1.PtrVal.(map[string]gopurs_runtime.Value)["value2"])})})
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")})
}
end_branch_2:
return __t2
})
})
})
	})
	return pop
}

var member gopurs_runtime.Value
var once_member sync.Once
func Get_member() gopurs_runtime.Value {
	once_member.Do(func() {
		member = gopurs_runtime.Func(func(dictOrd_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_1 gopurs_runtime.Value) gopurs_runtime.Value {
var go__2_0 gopurs_runtime.Value
go__2_0 = gopurs_runtime.Func(func(v_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
go__2_0:
for {
if false { continue go__2_0 }
var v_3 = v_3_loop
_ = v_3
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(v_3.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Leaf")).IntVal != 0 {
__t1 = gopurs_runtime.Bool(false)
goto end_branch_1
} else {

}
}
{
if (gopurs_runtime.Bool(v_3.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Node")).IntVal != 0 {
v1_4_2 := gopurs_runtime.Apply(gopurs_runtime.Apply(dictOrd_0.PtrVal.(map[string]gopurs_runtime.Value)["compare"], k_1), v_3.PtrVal.(map[string]gopurs_runtime.Value)["value2"])
var __t3 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(v1_4_2.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "LT")).IntVal != 0 {
v_3_loop = v_3.PtrVal.(map[string]gopurs_runtime.Value)["value4"]
continue go__2_0
__t3 = gopurs_runtime.Value{}
goto end_branch_3
} else {

}
}
{
if (gopurs_runtime.Bool(v1_4_2.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "GT")).IntVal != 0 {
v_3_loop = v_3.PtrVal.(map[string]gopurs_runtime.Value)["value5"]
continue go__2_0
__t3 = gopurs_runtime.Value{}
goto end_branch_3
} else {

}
}
{
if (gopurs_runtime.Bool(v1_4_2.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "EQ")).IntVal != 0 {
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
})
})
	})
	return member
}

var mapMaybeWithKey gopurs_runtime.Value
var once_mapMaybeWithKey sync.Once
func Get_mapMaybeWithKey() gopurs_runtime.Value {
	once_mapMaybeWithKey.Do(func() {
		mapMaybeWithKey = gopurs_runtime.Func(func(dictOrd_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
var go__2_0 gopurs_runtime.Value
go__2_0 = gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(v_3.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Leaf")).IntVal != 0 {
__t1 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Leaf")})
goto end_branch_1
} else {

}
}
{
if (gopurs_runtime.Bool(v_3.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Node")).IntVal != 0 {
v2_4_2 := gopurs_runtime.Apply(gopurs_runtime.Apply(f_1, v_3.PtrVal.(map[string]gopurs_runtime.Value)["value2"]), v_3.PtrVal.(map[string]gopurs_runtime.Value)["value3"])
var __t3 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(v2_4_2.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just")).IntVal != 0 {
__t3 = gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Function_Uncurried.Get_runFn4(), Get_unsafeBalancedNode()), v_3.PtrVal.(map[string]gopurs_runtime.Value)["value2"]), v2_4_2.PtrVal.(map[string]gopurs_runtime.Value)["value0"]), gopurs_runtime.Apply(go__2_0, v_3.PtrVal.(map[string]gopurs_runtime.Value)["value4"])), gopurs_runtime.Apply(go__2_0, v_3.PtrVal.(map[string]gopurs_runtime.Value)["value5"]))
goto end_branch_3
} else {

}
}
{
if (gopurs_runtime.Bool(v2_4_2.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Nothing")).IntVal != 0 {
__t3 = gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Function_Uncurried.Get_runFn2(), Get_unsafeJoinNodes()), gopurs_runtime.Apply(go__2_0, v_3.PtrVal.(map[string]gopurs_runtime.Value)["value4"])), gopurs_runtime.Apply(go__2_0, v_3.PtrVal.(map[string]gopurs_runtime.Value)["value5"]))
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
	})
	return mapMaybeWithKey
}

var mapMaybe gopurs_runtime.Value
var once_mapMaybe sync.Once
func Get_mapMaybe() gopurs_runtime.Value {
	once_mapMaybe.Do(func() {
		mapMaybe = gopurs_runtime.Func(func(dictOrd_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(Get_mapMaybeWithKey(), dictOrd_0), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return x_1
}))
})
})
	})
	return mapMaybe
}

var lookupLE gopurs_runtime.Value
var once_lookupLE sync.Once
func Get_lookupLE() gopurs_runtime.Value {
	once_lookupLE.Do(func() {
		lookupLE = gopurs_runtime.Func(func(dictOrd_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_1 gopurs_runtime.Value) gopurs_runtime.Value {
var go__2_0 gopurs_runtime.Value
go__2_0 = gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(v_3.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Leaf")).IntVal != 0 {
__t1 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")})
goto end_branch_1
} else {

}
}
{
if (gopurs_runtime.Bool(v_3.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Node")).IntVal != 0 {
v1_4_2 := gopurs_runtime.Apply(gopurs_runtime.Apply(dictOrd_0.PtrVal.(map[string]gopurs_runtime.Value)["compare"], k_1), v_3.PtrVal.(map[string]gopurs_runtime.Value)["value2"])
var __t3 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(v1_4_2.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "LT")).IntVal != 0 {
__t3 = gopurs_runtime.Apply(go__2_0, v_3.PtrVal.(map[string]gopurs_runtime.Value)["value4"])
goto end_branch_3
} else {

}
}
{
if (gopurs_runtime.Bool(v1_4_2.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "GT")).IntVal != 0 {
v2_5_4 := gopurs_runtime.Apply(go__2_0, v_3.PtrVal.(map[string]gopurs_runtime.Value)["value5"])
var __t5 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(v2_5_4.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Nothing")).IntVal != 0 {
__t5 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Just"), "value0": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"key": v_3.PtrVal.(map[string]gopurs_runtime.Value)["value2"], "value": v_3.PtrVal.(map[string]gopurs_runtime.Value)["value3"]})})
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
if (gopurs_runtime.Bool(v1_4_2.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "EQ")).IntVal != 0 {
__t3 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Just"), "value0": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"key": v_3.PtrVal.(map[string]gopurs_runtime.Value)["value2"], "value": v_3.PtrVal.(map[string]gopurs_runtime.Value)["value3"]})})
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
	})
	return lookupLE
}

var lookupGE gopurs_runtime.Value
var once_lookupGE sync.Once
func Get_lookupGE() gopurs_runtime.Value {
	once_lookupGE.Do(func() {
		lookupGE = gopurs_runtime.Func(func(dictOrd_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_1 gopurs_runtime.Value) gopurs_runtime.Value {
var go__2_0 gopurs_runtime.Value
go__2_0 = gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(v_3.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Leaf")).IntVal != 0 {
__t1 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")})
goto end_branch_1
} else {

}
}
{
if (gopurs_runtime.Bool(v_3.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Node")).IntVal != 0 {
v1_4_2 := gopurs_runtime.Apply(gopurs_runtime.Apply(dictOrd_0.PtrVal.(map[string]gopurs_runtime.Value)["compare"], k_1), v_3.PtrVal.(map[string]gopurs_runtime.Value)["value2"])
var __t3 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(v1_4_2.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "LT")).IntVal != 0 {
v2_5_4 := gopurs_runtime.Apply(go__2_0, v_3.PtrVal.(map[string]gopurs_runtime.Value)["value4"])
var __t5 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(v2_5_4.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Nothing")).IntVal != 0 {
__t5 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Just"), "value0": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"key": v_3.PtrVal.(map[string]gopurs_runtime.Value)["value2"], "value": v_3.PtrVal.(map[string]gopurs_runtime.Value)["value3"]})})
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
if (gopurs_runtime.Bool(v1_4_2.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "GT")).IntVal != 0 {
__t3 = gopurs_runtime.Apply(go__2_0, v_3.PtrVal.(map[string]gopurs_runtime.Value)["value5"])
goto end_branch_3
} else {

}
}
{
if (gopurs_runtime.Bool(v1_4_2.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "EQ")).IntVal != 0 {
__t3 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Just"), "value0": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"key": v_3.PtrVal.(map[string]gopurs_runtime.Value)["value2"], "value": v_3.PtrVal.(map[string]gopurs_runtime.Value)["value3"]})})
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
	})
	return lookupGE
}

var lookup gopurs_runtime.Value
var once_lookup sync.Once
func Get_lookup() gopurs_runtime.Value {
	once_lookup.Do(func() {
		lookup = gopurs_runtime.Func(func(dictOrd_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_1 gopurs_runtime.Value) gopurs_runtime.Value {
var go__2_0 gopurs_runtime.Value
go__2_0 = gopurs_runtime.Func(func(v_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
go__2_0:
for {
if false { continue go__2_0 }
var v_3 = v_3_loop
_ = v_3
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(v_3.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Leaf")).IntVal != 0 {
__t1 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")})
goto end_branch_1
} else {

}
}
{
if (gopurs_runtime.Bool(v_3.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Node")).IntVal != 0 {
v1_4_2 := gopurs_runtime.Apply(gopurs_runtime.Apply(dictOrd_0.PtrVal.(map[string]gopurs_runtime.Value)["compare"], k_1), v_3.PtrVal.(map[string]gopurs_runtime.Value)["value2"])
var __t3 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(v1_4_2.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "LT")).IntVal != 0 {
v_3_loop = v_3.PtrVal.(map[string]gopurs_runtime.Value)["value4"]
continue go__2_0
__t3 = gopurs_runtime.Value{}
goto end_branch_3
} else {

}
}
{
if (gopurs_runtime.Bool(v1_4_2.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "GT")).IntVal != 0 {
v_3_loop = v_3.PtrVal.(map[string]gopurs_runtime.Value)["value5"]
continue go__2_0
__t3 = gopurs_runtime.Value{}
goto end_branch_3
} else {

}
}
{
if (gopurs_runtime.Bool(v1_4_2.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "EQ")).IntVal != 0 {
__t3 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Just"), "value0": v_3.PtrVal.(map[string]gopurs_runtime.Value)["value3"]})
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
	})
	return lookup
}

var iterMapU gopurs_runtime.Value
var once_iterMapU sync.Once
func Get_iterMapU() gopurs_runtime.Value {
	once_iterMapU.Do(func() {
		iterMapU = gopurs_runtime.Func(func(iter_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(v_1.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Leaf")).IntVal != 0 {
__t0 = iter_0
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Bool(v_1.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Node")).IntVal != 0 {
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(v_1.PtrVal.(map[string]gopurs_runtime.Value)["value4"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Leaf")).IntVal != 0 {
var __t2 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(v_1.PtrVal.(map[string]gopurs_runtime.Value)["value5"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Leaf")).IntVal != 0 {
__t2 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("IterEmit"), "value0": v_1.PtrVal.(map[string]gopurs_runtime.Value)["value2"], "value1": v_1.PtrVal.(map[string]gopurs_runtime.Value)["value3"], "value2": iter_0})
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("IterEmit"), "value0": v_1.PtrVal.(map[string]gopurs_runtime.Value)["value2"], "value1": v_1.PtrVal.(map[string]gopurs_runtime.Value)["value3"], "value2": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("IterNode"), "value0": v_1.PtrVal.(map[string]gopurs_runtime.Value)["value5"], "value1": iter_0})})
}
end_branch_2:
__t1 = __t2
goto end_branch_1
} else {

}
}
{
if (gopurs_runtime.Bool(v_1.PtrVal.(map[string]gopurs_runtime.Value)["value5"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Leaf")).IntVal != 0 {
__t1 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("IterEmit"), "value0": v_1.PtrVal.(map[string]gopurs_runtime.Value)["value2"], "value1": v_1.PtrVal.(map[string]gopurs_runtime.Value)["value3"], "value2": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("IterNode"), "value0": v_1.PtrVal.(map[string]gopurs_runtime.Value)["value4"], "value1": iter_0})})
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("IterEmit"), "value0": v_1.PtrVal.(map[string]gopurs_runtime.Value)["value2"], "value1": v_1.PtrVal.(map[string]gopurs_runtime.Value)["value3"], "value2": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("IterNode"), "value0": v_1.PtrVal.(map[string]gopurs_runtime.Value)["value4"], "value1": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("IterNode"), "value0": v_1.PtrVal.(map[string]gopurs_runtime.Value)["value5"], "value1": iter_0})})})
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
		stepUnfoldrUnordered = gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(Get_stepWith(), Get_iterMapU()), gopurs_runtime.Apply(pkg_Data_Function_Uncurried.Get_mkFn3(), gopurs_runtime.Func(func(k_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(next_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Just"), "value0": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Tuple"), "value0": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Tuple"), "value0": k_0, "value1": v_1}), "value1": next_2})})
})
})
}))), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")})
}))
	})
	return stepUnfoldrUnordered
}

var toUnfoldableUnordered gopurs_runtime.Value
var once_toUnfoldableUnordered sync.Once
func Get_toUnfoldableUnordered() gopurs_runtime.Value {
	once_toUnfoldableUnordered.Do(func() {
		toUnfoldableUnordered = gopurs_runtime.Func(func(dictUnfoldable_0 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_1_0 := gopurs_runtime.Apply(dictUnfoldable_0.PtrVal.(map[string]gopurs_runtime.Value)["unfoldr"], Get_stepUnfoldrUnordered())
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_1_0, gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("IterNode"), "value0": x_2, "value1": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("IterLeaf")})}))
})
})
	})
	return toUnfoldableUnordered
}

var stepUnordered gopurs_runtime.Value
var once_stepUnordered sync.Once
func Get_stepUnordered() gopurs_runtime.Value {
	once_stepUnordered.Do(func() {
		stepUnordered = gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(Get_stepWith(), Get_iterMapU()), gopurs_runtime.Apply(pkg_Data_Function_Uncurried.Get_mkFn3(), gopurs_runtime.Func(func(k_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(next_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("IterNext"), "value0": k_0, "value1": v_1, "value2": next_2})
})
})
}))), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("IterDone")})
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
var iter_1 = iter_1_loop
_ = iter_1
var v_2 = v_2_loop
_ = v_2
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(v_2.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Leaf")).IntVal != 0 {
__t1 = iter_1
goto end_branch_1
} else {

}
}
{
if (gopurs_runtime.Bool(v_2.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Node")).IntVal != 0 {
var __t2 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(v_2.PtrVal.(map[string]gopurs_runtime.Value)["value5"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Leaf")).IntVal != 0 {
iter_1_loop = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("IterEmit"), "value0": v_2.PtrVal.(map[string]gopurs_runtime.Value)["value2"], "value1": v_2.PtrVal.(map[string]gopurs_runtime.Value)["value3"], "value2": iter_1})
v_2_loop = v_2.PtrVal.(map[string]gopurs_runtime.Value)["value4"]
continue go__0_0
__t2 = gopurs_runtime.Value{}
goto end_branch_2
} else {

}
}
{
iter_1_loop = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("IterEmit"), "value0": v_2.PtrVal.(map[string]gopurs_runtime.Value)["value2"], "value1": v_2.PtrVal.(map[string]gopurs_runtime.Value)["value3"], "value2": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("IterNode"), "value0": v_2.PtrVal.(map[string]gopurs_runtime.Value)["value4"], "value1": iter_1})})
v_2_loop = v_2.PtrVal.(map[string]gopurs_runtime.Value)["value5"]
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
		stepDesc = gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(Get_stepWith(), Get_iterMapR()), gopurs_runtime.Apply(pkg_Data_Function_Uncurried.Get_mkFn3(), gopurs_runtime.Func(func(k_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(next_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("IterNext"), "value0": k_0, "value1": v_1, "value2": next_2})
})
})
}))), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("IterDone")})
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
var iter_1 = iter_1_loop
_ = iter_1
var v_2 = v_2_loop
_ = v_2
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(v_2.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Leaf")).IntVal != 0 {
__t1 = iter_1
goto end_branch_1
} else {

}
}
{
if (gopurs_runtime.Bool(v_2.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Node")).IntVal != 0 {
var __t2 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(v_2.PtrVal.(map[string]gopurs_runtime.Value)["value5"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Leaf")).IntVal != 0 {
iter_1_loop = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("IterEmit"), "value0": v_2.PtrVal.(map[string]gopurs_runtime.Value)["value2"], "value1": v_2.PtrVal.(map[string]gopurs_runtime.Value)["value3"], "value2": iter_1})
v_2_loop = v_2.PtrVal.(map[string]gopurs_runtime.Value)["value4"]
continue go__0_0
__t2 = gopurs_runtime.Value{}
goto end_branch_2
} else {

}
}
{
iter_1_loop = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("IterEmit"), "value0": v_2.PtrVal.(map[string]gopurs_runtime.Value)["value2"], "value1": v_2.PtrVal.(map[string]gopurs_runtime.Value)["value3"], "value2": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("IterNode"), "value0": v_2.PtrVal.(map[string]gopurs_runtime.Value)["value5"], "value1": iter_1})})
v_2_loop = v_2.PtrVal.(map[string]gopurs_runtime.Value)["value4"]
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
		stepAsc = gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(Get_stepWith(), Get_iterMapL()), gopurs_runtime.Apply(pkg_Data_Function_Uncurried.Get_mkFn3(), gopurs_runtime.Func(func(k_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(next_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("IterNext"), "value0": k_0, "value1": v_1, "value2": next_2})
})
})
}))), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("IterDone")})
}))
	})
	return stepAsc
}

var eqMapIter gopurs_runtime.Value
var once_eqMapIter sync.Once
func Get_eqMapIter() gopurs_runtime.Value {
	once_eqMapIter.Do(func() {
		eqMapIter = gopurs_runtime.Func(func(dictEq_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(dictEq1_1 gopurs_runtime.Value) gopurs_runtime.Value {
var go__2_0 gopurs_runtime.Value
go__2_0 = gopurs_runtime.Func(func(a_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_4 gopurs_runtime.Value) gopurs_runtime.Value {
v_5_1 := gopurs_runtime.Apply(Get_stepAsc(), a_3)
var __t2 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(v_5_1.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "IterNext")).IntVal != 0 {
v2_6_3 := gopurs_runtime.Apply(Get_stepAsc(), b_4)
__t2 = gopurs_runtime.Bool(gopurs_runtime.Bool(v2_6_3.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "IterNext").IntVal != 0 && gopurs_runtime.Bool(gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_HeytingAlgebra.Get_boolConj(), gopurs_runtime.Apply(gopurs_runtime.Apply(dictEq_0.PtrVal.(map[string]gopurs_runtime.Value)["eq"], v_5_1.PtrVal.(map[string]gopurs_runtime.Value)["value0"]), v2_6_3.PtrVal.(map[string]gopurs_runtime.Value)["value0"])), gopurs_runtime.Apply(gopurs_runtime.Apply(dictEq1_1.PtrVal.(map[string]gopurs_runtime.Value)["eq"], v_5_1.PtrVal.(map[string]gopurs_runtime.Value)["value1"]), v2_6_3.PtrVal.(map[string]gopurs_runtime.Value)["value1"])).IntVal != 0 && gopurs_runtime.Apply(gopurs_runtime.Apply(go__2_0, v_5_1.PtrVal.(map[string]gopurs_runtime.Value)["value2"]), v2_6_3.PtrVal.(map[string]gopurs_runtime.Value)["value2"]).IntVal != 0).IntVal != 0)
goto end_branch_2
} else {

}
}
{
if (gopurs_runtime.Bool(v_5_1.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "IterDone")).IntVal != 0 {
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
})
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"eq": go__2_0})
})
})
	})
	return eqMapIter
}

var ordMapIter gopurs_runtime.Value
var once_ordMapIter sync.Once
func Get_ordMapIter() gopurs_runtime.Value {
	once_ordMapIter.Do(func() {
		ordMapIter = gopurs_runtime.Func(func(dictOrd_0 gopurs_runtime.Value) gopurs_runtime.Value {
eqMapIter1_1_0 := gopurs_runtime.Apply(Get_eqMapIter(), gopurs_runtime.Apply(dictOrd_0.PtrVal.(map[string]gopurs_runtime.Value)["Eq0"], gopurs_runtime.Value{}))
return gopurs_runtime.Func(func(dictOrd1_2 gopurs_runtime.Value) gopurs_runtime.Value {
eqMapIter2_3_1 := gopurs_runtime.Apply(eqMapIter1_1_0, gopurs_runtime.Apply(dictOrd1_2.PtrVal.(map[string]gopurs_runtime.Value)["Eq0"], gopurs_runtime.Value{}))
var go__4_2 gopurs_runtime.Value
go__4_2 = gopurs_runtime.Func(func(a_5_loop gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_6_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
go__4_2:
for {
if false { continue go__4_2 }
var a_5 = a_5_loop
_ = a_5
var b_6 = b_6_loop
_ = b_6
v_7_3 := gopurs_runtime.Apply(Get_stepAsc(), b_6)
v1_8_4 := gopurs_runtime.Apply(Get_stepAsc(), a_5)
var __t5 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(v1_8_4.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "IterNext")).IntVal != 0 {
var __t6 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(v_7_3.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "IterNext")).IntVal != 0 {
v3_9_7 := gopurs_runtime.Apply(gopurs_runtime.Apply(dictOrd_0.PtrVal.(map[string]gopurs_runtime.Value)["compare"], v1_8_4.PtrVal.(map[string]gopurs_runtime.Value)["value0"]), v_7_3.PtrVal.(map[string]gopurs_runtime.Value)["value0"])
var __t8 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(v3_9_7.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "EQ")).IntVal != 0 {
v4_10_9 := gopurs_runtime.Apply(gopurs_runtime.Apply(dictOrd1_2.PtrVal.(map[string]gopurs_runtime.Value)["compare"], v1_8_4.PtrVal.(map[string]gopurs_runtime.Value)["value1"]), v_7_3.PtrVal.(map[string]gopurs_runtime.Value)["value1"])
var __t10 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(v4_10_9.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "EQ")).IntVal != 0 {
a_5_loop = v1_8_4.PtrVal.(map[string]gopurs_runtime.Value)["value2"]
b_6_loop = v_7_3.PtrVal.(map[string]gopurs_runtime.Value)["value2"]
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
if (gopurs_runtime.Bool(v_7_3.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "IterDone")).IntVal != 0 {
__t6 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("GT")})
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
if (gopurs_runtime.Bool(v1_8_4.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "IterDone")).IntVal != 0 {
var __t11 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(v_7_3.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "IterDone")).IntVal != 0 {
__t11 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("EQ")})
goto end_branch_11
} else {

}
}
{
__t11 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("LT")})
}
end_branch_11:
__t5 = __t11
goto end_branch_5
} else {

}
}
{
if (gopurs_runtime.Bool(v_7_3.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "IterDone")).IntVal != 0 {
__t5 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("GT")})
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
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"compare": go__4_2, "Eq0": gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return eqMapIter2_3_1
})})
})
})
	})
	return ordMapIter
}

var stepUnfoldr gopurs_runtime.Value
var once_stepUnfoldr sync.Once
func Get_stepUnfoldr() gopurs_runtime.Value {
	once_stepUnfoldr.Do(func() {
		stepUnfoldr = gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(Get_stepWith(), Get_iterMapL()), gopurs_runtime.Apply(pkg_Data_Function_Uncurried.Get_mkFn3(), gopurs_runtime.Func(func(k_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(next_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Just"), "value0": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Tuple"), "value0": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Tuple"), "value0": k_0, "value1": v_1}), "value1": next_2})})
})
})
}))), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")})
}))
	})
	return stepUnfoldr
}

var toUnfoldable gopurs_runtime.Value
var once_toUnfoldable sync.Once
func Get_toUnfoldable() gopurs_runtime.Value {
	once_toUnfoldable.Do(func() {
		toUnfoldable = gopurs_runtime.Func(func(dictUnfoldable_0 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_1_0 := gopurs_runtime.Apply(dictUnfoldable_0.PtrVal.(map[string]gopurs_runtime.Value)["unfoldr"], Get_stepUnfoldr())
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_1_0, gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("IterNode"), "value0": x_2, "value1": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("IterLeaf")})}))
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
__local_var_0_0 := gopurs_runtime.Apply(pkg_Data_Unfoldable.Get_unfoldableArray().PtrVal.(map[string]gopurs_runtime.Value)["unfoldr"], Get_stepUnfoldr())
return gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_0_0, gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("IterNode"), "value0": x_1, "value1": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("IterLeaf")})}))
})
}()
	})
	return toUnfoldable1
}

var showMap gopurs_runtime.Value
var once_showMap sync.Once
func Get_showMap() gopurs_runtime.Value {
	once_showMap.Do(func() {
		showMap = gopurs_runtime.Func(func(dictShow_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(dictShow1_1 gopurs_runtime.Value) gopurs_runtime.Value {
show1_2_0 := gopurs_runtime.Apply(pkg_Data_Show.Get_showArrayImpl(), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Tuple.Get_showTuple(), dictShow_0), dictShow1_1).PtrVal.(map[string]gopurs_runtime.Value)["show"])
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"show": gopurs_runtime.Func(func(as_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Semigroup.Get_concatString(), gopurs_runtime.Str("(fromFoldable ")), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Semigroup.Get_concatString(), gopurs_runtime.Apply(show1_2_0, gopurs_runtime.Apply(Get_toUnfoldable1(), as_3))), gopurs_runtime.Str(")")))
})})
})
})
	})
	return showMap
}

var isSubmap gopurs_runtime.Value
var once_isSubmap sync.Once
func Get_isSubmap() gopurs_runtime.Value {
	once_isSubmap.Do(func() {
		isSubmap = gopurs_runtime.Func(func(dictOrd_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(dictEq_1 gopurs_runtime.Value) gopurs_runtime.Value {
var go__2_0 gopurs_runtime.Value
go__2_0 = gopurs_runtime.Func(func(m1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m2_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(m1_3.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Leaf")).IntVal != 0 {
__t1 = gopurs_runtime.Bool(true)
goto end_branch_1
} else {

}
}
{
if (gopurs_runtime.Bool(m1_3.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Node")).IntVal != 0 {
__local_var_5_2 := m1_3.PtrVal.(map[string]gopurs_runtime.Value)["value2"]
var go__6_3 gopurs_runtime.Value
go__6_3 = gopurs_runtime.Func(func(v_7_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
go__6_3:
for {
if false { continue go__6_3 }
var v_7 = v_7_loop
_ = v_7
var __t4 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(v_7.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Leaf")).IntVal != 0 {
__t4 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")})
goto end_branch_4
} else {

}
}
{
if (gopurs_runtime.Bool(v_7.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Node")).IntVal != 0 {
v1_8_5 := gopurs_runtime.Apply(gopurs_runtime.Apply(dictOrd_0.PtrVal.(map[string]gopurs_runtime.Value)["compare"], __local_var_5_2), v_7.PtrVal.(map[string]gopurs_runtime.Value)["value2"])
var __t6 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(v1_8_5.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "LT")).IntVal != 0 {
v_7_loop = v_7.PtrVal.(map[string]gopurs_runtime.Value)["value4"]
continue go__6_3
__t6 = gopurs_runtime.Value{}
goto end_branch_6
} else {

}
}
{
if (gopurs_runtime.Bool(v1_8_5.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "GT")).IntVal != 0 {
v_7_loop = v_7.PtrVal.(map[string]gopurs_runtime.Value)["value5"]
continue go__6_3
__t6 = gopurs_runtime.Value{}
goto end_branch_6
} else {

}
}
{
if (gopurs_runtime.Bool(v1_8_5.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "EQ")).IntVal != 0 {
__t6 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Just"), "value0": v_7.PtrVal.(map[string]gopurs_runtime.Value)["value3"]})
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
var __t8 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(v1_7_7.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Nothing")).IntVal != 0 {
__t8 = gopurs_runtime.Bool(false)
goto end_branch_8
} else {

}
}
{
if (gopurs_runtime.Bool(v1_7_7.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just")).IntVal != 0 {
__t8 = gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_HeytingAlgebra.Get_boolConj(), gopurs_runtime.Apply(gopurs_runtime.Apply(dictEq_1.PtrVal.(map[string]gopurs_runtime.Value)["eq"], m1_3.PtrVal.(map[string]gopurs_runtime.Value)["value3"]), v1_7_7.PtrVal.(map[string]gopurs_runtime.Value)["value0"])), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_HeytingAlgebra.Get_boolConj(), gopurs_runtime.Apply(gopurs_runtime.Apply(go__2_0, m1_3.PtrVal.(map[string]gopurs_runtime.Value)["value4"]), m2_4)), gopurs_runtime.Apply(gopurs_runtime.Apply(go__2_0, m1_3.PtrVal.(map[string]gopurs_runtime.Value)["value5"]), m2_4)))
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
})
return go__2_0
})
})
	})
	return isSubmap
}

var isEmpty gopurs_runtime.Value
var once_isEmpty sync.Once
func Get_isEmpty() gopurs_runtime.Value {
	once_isEmpty.Do(func() {
		isEmpty = gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(v_0.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Leaf")
})
	})
	return isEmpty
}

var intersectionWith gopurs_runtime.Value
var once_intersectionWith sync.Once
func Get_intersectionWith() gopurs_runtime.Value {
	once_intersectionWith.Do(func() {
		intersectionWith = gopurs_runtime.Func(func(dictOrd_0 gopurs_runtime.Value) gopurs_runtime.Value {
compare_1_0 := dictOrd_0.PtrVal.(map[string]gopurs_runtime.Value)["compare"]
return gopurs_runtime.Func(func(app_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m2_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Function_Uncurried.Get_runFn4(), Get_unsafeIntersectionWith()), compare_1_0), app_2), m1_3), m2_4)
})
})
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
compare_1_0 := dictOrd_0.PtrVal.(map[string]gopurs_runtime.Value)["compare"]
return gopurs_runtime.Func(func(m1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m2_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Function_Uncurried.Get_runFn4(), Get_unsafeIntersectionWith()), compare_1_0), pkg_Data_Function.Get_const_()), m1_2), m2_3)
})
})
})
	})
	return intersection
}

var insertWith gopurs_runtime.Value
var once_insertWith sync.Once
func Get_insertWith() gopurs_runtime.Value {
	once_insertWith.Do(func() {
		insertWith = gopurs_runtime.Func(func(dictOrd_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(app_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
var go__4_0 gopurs_runtime.Value
go__4_0 = gopurs_runtime.Func(func(v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(v1_5.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Leaf")).IntVal != 0 {
__t1 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Node"), "value0": gopurs_runtime.Int(1), "value1": gopurs_runtime.Int(1), "value2": k_2, "value3": v_3, "value4": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Leaf")}), "value5": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Leaf")})})
goto end_branch_1
} else {

}
}
{
if (gopurs_runtime.Bool(v1_5.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Node")).IntVal != 0 {
v2_6_2 := gopurs_runtime.Apply(gopurs_runtime.Apply(dictOrd_0.PtrVal.(map[string]gopurs_runtime.Value)["compare"], k_2), v1_5.PtrVal.(map[string]gopurs_runtime.Value)["value2"])
var __t3 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(v2_6_2.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "LT")).IntVal != 0 {
__t3 = gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Function_Uncurried.Get_runFn4(), Get_unsafeBalancedNode()), v1_5.PtrVal.(map[string]gopurs_runtime.Value)["value2"]), v1_5.PtrVal.(map[string]gopurs_runtime.Value)["value3"]), gopurs_runtime.Apply(go__4_0, v1_5.PtrVal.(map[string]gopurs_runtime.Value)["value4"])), v1_5.PtrVal.(map[string]gopurs_runtime.Value)["value5"])
goto end_branch_3
} else {

}
}
{
if (gopurs_runtime.Bool(v2_6_2.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "GT")).IntVal != 0 {
__t3 = gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Function_Uncurried.Get_runFn4(), Get_unsafeBalancedNode()), v1_5.PtrVal.(map[string]gopurs_runtime.Value)["value2"]), v1_5.PtrVal.(map[string]gopurs_runtime.Value)["value3"]), v1_5.PtrVal.(map[string]gopurs_runtime.Value)["value4"]), gopurs_runtime.Apply(go__4_0, v1_5.PtrVal.(map[string]gopurs_runtime.Value)["value5"]))
goto end_branch_3
} else {

}
}
{
if (gopurs_runtime.Bool(v2_6_2.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "EQ")).IntVal != 0 {
__t3 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Node"), "value0": v1_5.PtrVal.(map[string]gopurs_runtime.Value)["value0"], "value1": v1_5.PtrVal.(map[string]gopurs_runtime.Value)["value1"], "value2": k_2, "value3": gopurs_runtime.Apply(gopurs_runtime.Apply(app_1, v1_5.PtrVal.(map[string]gopurs_runtime.Value)["value3"]), v_3), "value4": v1_5.PtrVal.(map[string]gopurs_runtime.Value)["value4"], "value5": v1_5.PtrVal.(map[string]gopurs_runtime.Value)["value5"]})
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
})
})
	})
	return insertWith
}

var insert gopurs_runtime.Value
var once_insert sync.Once
func Get_insert() gopurs_runtime.Value {
	once_insert.Do(func() {
		insert = gopurs_runtime.Func(func(dictOrd_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
var go__3_0 gopurs_runtime.Value
go__3_0 = gopurs_runtime.Func(func(v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(v1_4.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Leaf")).IntVal != 0 {
__t1 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Node"), "value0": gopurs_runtime.Int(1), "value1": gopurs_runtime.Int(1), "value2": k_1, "value3": v_2, "value4": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Leaf")}), "value5": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Leaf")})})
goto end_branch_1
} else {

}
}
{
if (gopurs_runtime.Bool(v1_4.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Node")).IntVal != 0 {
v2_5_2 := gopurs_runtime.Apply(gopurs_runtime.Apply(dictOrd_0.PtrVal.(map[string]gopurs_runtime.Value)["compare"], k_1), v1_4.PtrVal.(map[string]gopurs_runtime.Value)["value2"])
var __t3 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(v2_5_2.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "LT")).IntVal != 0 {
__t3 = gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Function_Uncurried.Get_runFn4(), Get_unsafeBalancedNode()), v1_4.PtrVal.(map[string]gopurs_runtime.Value)["value2"]), v1_4.PtrVal.(map[string]gopurs_runtime.Value)["value3"]), gopurs_runtime.Apply(go__3_0, v1_4.PtrVal.(map[string]gopurs_runtime.Value)["value4"])), v1_4.PtrVal.(map[string]gopurs_runtime.Value)["value5"])
goto end_branch_3
} else {

}
}
{
if (gopurs_runtime.Bool(v2_5_2.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "GT")).IntVal != 0 {
__t3 = gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Function_Uncurried.Get_runFn4(), Get_unsafeBalancedNode()), v1_4.PtrVal.(map[string]gopurs_runtime.Value)["value2"]), v1_4.PtrVal.(map[string]gopurs_runtime.Value)["value3"]), v1_4.PtrVal.(map[string]gopurs_runtime.Value)["value4"]), gopurs_runtime.Apply(go__3_0, v1_4.PtrVal.(map[string]gopurs_runtime.Value)["value5"]))
goto end_branch_3
} else {

}
}
{
if (gopurs_runtime.Bool(v2_5_2.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "EQ")).IntVal != 0 {
__t3 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Node"), "value0": v1_4.PtrVal.(map[string]gopurs_runtime.Value)["value0"], "value1": v1_4.PtrVal.(map[string]gopurs_runtime.Value)["value1"], "value2": k_1, "value3": v_2, "value4": v1_4.PtrVal.(map[string]gopurs_runtime.Value)["value4"], "value5": v1_4.PtrVal.(map[string]gopurs_runtime.Value)["value5"]})
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
})
	})
	return insert
}

var functorMap gopurs_runtime.Value
var once_functorMap sync.Once
func Get_functorMap() gopurs_runtime.Value {
	once_functorMap.Do(func() {
		functorMap = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"map": gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
var go__1_0 gopurs_runtime.Value
go__1_0 = gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(v_2.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Leaf")).IntVal != 0 {
__t1 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Leaf")})
goto end_branch_1
} else {

}
}
{
if (gopurs_runtime.Bool(v_2.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Node")).IntVal != 0 {
__t1 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Node"), "value0": v_2.PtrVal.(map[string]gopurs_runtime.Value)["value0"], "value1": v_2.PtrVal.(map[string]gopurs_runtime.Value)["value1"], "value2": v_2.PtrVal.(map[string]gopurs_runtime.Value)["value2"], "value3": gopurs_runtime.Apply(f_0, v_2.PtrVal.(map[string]gopurs_runtime.Value)["value3"]), "value4": gopurs_runtime.Apply(go__1_0, v_2.PtrVal.(map[string]gopurs_runtime.Value)["value4"]), "value5": gopurs_runtime.Apply(go__1_0, v_2.PtrVal.(map[string]gopurs_runtime.Value)["value5"])})
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
})})
	})
	return functorMap
}

var functorWithIndexMap gopurs_runtime.Value
var once_functorWithIndexMap sync.Once
func Get_functorWithIndexMap() gopurs_runtime.Value {
	once_functorWithIndexMap.Do(func() {
		functorWithIndexMap = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"mapWithIndex": gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
var go__1_0 gopurs_runtime.Value
go__1_0 = gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(v_2.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Leaf")).IntVal != 0 {
__t1 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Leaf")})
goto end_branch_1
} else {

}
}
{
if (gopurs_runtime.Bool(v_2.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Node")).IntVal != 0 {
__t1 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Node"), "value0": v_2.PtrVal.(map[string]gopurs_runtime.Value)["value0"], "value1": v_2.PtrVal.(map[string]gopurs_runtime.Value)["value1"], "value2": v_2.PtrVal.(map[string]gopurs_runtime.Value)["value2"], "value3": gopurs_runtime.Apply(gopurs_runtime.Apply(f_0, v_2.PtrVal.(map[string]gopurs_runtime.Value)["value2"]), v_2.PtrVal.(map[string]gopurs_runtime.Value)["value3"]), "value4": gopurs_runtime.Apply(go__1_0, v_2.PtrVal.(map[string]gopurs_runtime.Value)["value4"]), "value5": gopurs_runtime.Apply(go__1_0, v_2.PtrVal.(map[string]gopurs_runtime.Value)["value5"])})
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
}), "Functor0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_functorMap()
})})
	})
	return functorWithIndexMap
}

var foldableMap gopurs_runtime.Value
var once_foldableMap sync.Once
func Get_foldableMap() gopurs_runtime.Value {
	once_foldableMap.Do(func() {
		foldableMap = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"foldr": gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(z_1 gopurs_runtime.Value) gopurs_runtime.Value {
var go__2_0 gopurs_runtime.Value
go__2_0 = gopurs_runtime.Apply(pkg_Data_Function_Uncurried.Get_mkFn2(), gopurs_runtime.Func(func(m_prime_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(z_prime_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(m_prime_3.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Leaf")).IntVal != 0 {
__t1 = z_prime_4
goto end_branch_1
} else {

}
}
{
if (gopurs_runtime.Bool(m_prime_3.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Node")).IntVal != 0 {
__t1 = gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Function_Uncurried.Get_runFn2(), go__2_0), m_prime_3.PtrVal.(map[string]gopurs_runtime.Value)["value4"]), gopurs_runtime.Apply(gopurs_runtime.Apply(f_0, m_prime_3.PtrVal.(map[string]gopurs_runtime.Value)["value3"]), gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Function_Uncurried.Get_runFn2(), go__2_0), m_prime_3.PtrVal.(map[string]gopurs_runtime.Value)["value5"]), z_prime_4)))
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
}))
return gopurs_runtime.Func(func(m_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Function_Uncurried.Get_runFn2(), go__2_0), m_3), z_1)
})
})
}), "foldl": gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(z_1 gopurs_runtime.Value) gopurs_runtime.Value {
var go__2_2 gopurs_runtime.Value
go__2_2 = gopurs_runtime.Apply(pkg_Data_Function_Uncurried.Get_mkFn2(), gopurs_runtime.Func(func(z_prime_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m_prime_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t3 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(m_prime_4.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Leaf")).IntVal != 0 {
__t3 = z_prime_3
goto end_branch_3
} else {

}
}
{
if (gopurs_runtime.Bool(m_prime_4.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Node")).IntVal != 0 {
__t3 = gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Function_Uncurried.Get_runFn2(), go__2_2), gopurs_runtime.Apply(gopurs_runtime.Apply(f_0, gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Function_Uncurried.Get_runFn2(), go__2_2), z_prime_3), m_prime_4.PtrVal.(map[string]gopurs_runtime.Value)["value4"])), m_prime_4.PtrVal.(map[string]gopurs_runtime.Value)["value3"])), m_prime_4.PtrVal.(map[string]gopurs_runtime.Value)["value5"])
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
}))
return gopurs_runtime.Func(func(m_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Function_Uncurried.Get_runFn2(), go__2_2), z_1), m_3)
})
})
}), "foldMap": gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
mempty_1_4 := dictMonoid_0.PtrVal.(map[string]gopurs_runtime.Value)["mempty"]
__local_var_2_5 := gopurs_runtime.Apply(dictMonoid_0.PtrVal.(map[string]gopurs_runtime.Value)["Semigroup0"], gopurs_runtime.Value{})
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
var go__4_6 gopurs_runtime.Value
go__4_6 = gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
var __t7 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(v_5.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Leaf")).IntVal != 0 {
__t7 = mempty_1_4
goto end_branch_7
} else {

}
}
{
if (gopurs_runtime.Bool(v_5.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Node")).IntVal != 0 {
__t7 = gopurs_runtime.Apply(gopurs_runtime.Apply(__local_var_2_5.PtrVal.(map[string]gopurs_runtime.Value)["append"], gopurs_runtime.Apply(go__4_6, v_5.PtrVal.(map[string]gopurs_runtime.Value)["value4"])), gopurs_runtime.Apply(gopurs_runtime.Apply(__local_var_2_5.PtrVal.(map[string]gopurs_runtime.Value)["append"], gopurs_runtime.Apply(f_3, v_5.PtrVal.(map[string]gopurs_runtime.Value)["value3"])), gopurs_runtime.Apply(go__4_6, v_5.PtrVal.(map[string]gopurs_runtime.Value)["value5"])))
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
})})
	})
	return foldableMap
}

var foldableWithIndexMap gopurs_runtime.Value
var once_foldableWithIndexMap sync.Once
func Get_foldableWithIndexMap() gopurs_runtime.Value {
	once_foldableWithIndexMap.Do(func() {
		foldableWithIndexMap = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"foldrWithIndex": gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(z_1 gopurs_runtime.Value) gopurs_runtime.Value {
var go__2_0 gopurs_runtime.Value
go__2_0 = gopurs_runtime.Apply(pkg_Data_Function_Uncurried.Get_mkFn2(), gopurs_runtime.Func(func(m_prime_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(z_prime_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(m_prime_3.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Leaf")).IntVal != 0 {
__t1 = z_prime_4
goto end_branch_1
} else {

}
}
{
if (gopurs_runtime.Bool(m_prime_3.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Node")).IntVal != 0 {
__t1 = gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Function_Uncurried.Get_runFn2(), go__2_0), m_prime_3.PtrVal.(map[string]gopurs_runtime.Value)["value4"]), gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(f_0, m_prime_3.PtrVal.(map[string]gopurs_runtime.Value)["value2"]), m_prime_3.PtrVal.(map[string]gopurs_runtime.Value)["value3"]), gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Function_Uncurried.Get_runFn2(), go__2_0), m_prime_3.PtrVal.(map[string]gopurs_runtime.Value)["value5"]), z_prime_4)))
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
}))
return gopurs_runtime.Func(func(m_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Function_Uncurried.Get_runFn2(), go__2_0), m_3), z_1)
})
})
}), "foldlWithIndex": gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(z_1 gopurs_runtime.Value) gopurs_runtime.Value {
var go__2_2 gopurs_runtime.Value
go__2_2 = gopurs_runtime.Apply(pkg_Data_Function_Uncurried.Get_mkFn2(), gopurs_runtime.Func(func(z_prime_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m_prime_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t3 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(m_prime_4.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Leaf")).IntVal != 0 {
__t3 = z_prime_3
goto end_branch_3
} else {

}
}
{
if (gopurs_runtime.Bool(m_prime_4.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Node")).IntVal != 0 {
__t3 = gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Function_Uncurried.Get_runFn2(), go__2_2), gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(f_0, m_prime_4.PtrVal.(map[string]gopurs_runtime.Value)["value2"]), gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Function_Uncurried.Get_runFn2(), go__2_2), z_prime_3), m_prime_4.PtrVal.(map[string]gopurs_runtime.Value)["value4"])), m_prime_4.PtrVal.(map[string]gopurs_runtime.Value)["value3"])), m_prime_4.PtrVal.(map[string]gopurs_runtime.Value)["value5"])
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
}))
return gopurs_runtime.Func(func(m_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Function_Uncurried.Get_runFn2(), go__2_2), z_1), m_3)
})
})
}), "foldMapWithIndex": gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
mempty_1_4 := dictMonoid_0.PtrVal.(map[string]gopurs_runtime.Value)["mempty"]
__local_var_2_5 := gopurs_runtime.Apply(dictMonoid_0.PtrVal.(map[string]gopurs_runtime.Value)["Semigroup0"], gopurs_runtime.Value{})
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
var go__4_6 gopurs_runtime.Value
go__4_6 = gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
var __t7 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(v_5.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Leaf")).IntVal != 0 {
__t7 = mempty_1_4
goto end_branch_7
} else {

}
}
{
if (gopurs_runtime.Bool(v_5.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Node")).IntVal != 0 {
__t7 = gopurs_runtime.Apply(gopurs_runtime.Apply(__local_var_2_5.PtrVal.(map[string]gopurs_runtime.Value)["append"], gopurs_runtime.Apply(go__4_6, v_5.PtrVal.(map[string]gopurs_runtime.Value)["value4"])), gopurs_runtime.Apply(gopurs_runtime.Apply(__local_var_2_5.PtrVal.(map[string]gopurs_runtime.Value)["append"], gopurs_runtime.Apply(gopurs_runtime.Apply(f_3, v_5.PtrVal.(map[string]gopurs_runtime.Value)["value2"]), v_5.PtrVal.(map[string]gopurs_runtime.Value)["value3"])), gopurs_runtime.Apply(go__4_6, v_5.PtrVal.(map[string]gopurs_runtime.Value)["value5"])))
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
}), "Foldable0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_foldableMap()
})})
	})
	return foldableWithIndexMap
}

var keys gopurs_runtime.Value
var once_keys sync.Once
func Get_keys() gopurs_runtime.Value {
	once_keys.Do(func() {
		keys = gopurs_runtime.Apply(gopurs_runtime.Apply(Get_foldableWithIndexMap().PtrVal.(map[string]gopurs_runtime.Value)["foldrWithIndex"], gopurs_runtime.Func(func(k_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(acc_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Cons"), "value0": k_0, "value1": acc_2})
})
})
})), gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nil")}))
	})
	return keys
}

var traversableMap gopurs_runtime.Value
var once_traversableMap sync.Once
func Get_traversableMap() gopurs_runtime.Value {
	once_traversableMap.Do(func() {
		traversableMap = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"traverse": gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
Apply0_1_0 := gopurs_runtime.Apply(dictApplicative_0.PtrVal.(map[string]gopurs_runtime.Value)["Apply0"], gopurs_runtime.Value{})
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
var go__3_1 gopurs_runtime.Value
go__3_1 = gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(v_4.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Leaf")).IntVal != 0 {
__t2 = gopurs_runtime.Apply(dictApplicative_0.PtrVal.(map[string]gopurs_runtime.Value)["pure"], gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Leaf")}))
goto end_branch_2
} else {

}
}
{
if (gopurs_runtime.Bool(v_4.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Node")).IntVal != 0 {
__local_var_5_3 := v_4.PtrVal.(map[string]gopurs_runtime.Value)["value0"]
__local_var_6_4 := v_4.PtrVal.(map[string]gopurs_runtime.Value)["value2"]
__local_var_7_5 := v_4.PtrVal.(map[string]gopurs_runtime.Value)["value1"]
__t2 = gopurs_runtime.Apply(gopurs_runtime.Apply(Apply0_1_0.PtrVal.(map[string]gopurs_runtime.Value)["apply"], gopurs_runtime.Apply(gopurs_runtime.Apply(Apply0_1_0.PtrVal.(map[string]gopurs_runtime.Value)["apply"], gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(Apply0_1_0.PtrVal.(map[string]gopurs_runtime.Value)["Functor0"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["map"], gopurs_runtime.Func(func(l_prime_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_prime_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(r_prime_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Node"), "value0": __local_var_5_3, "value1": __local_var_7_5, "value2": __local_var_6_4, "value3": v_prime_9, "value4": l_prime_8, "value5": r_prime_10})
})
})
})), gopurs_runtime.Apply(go__3_1, v_4.PtrVal.(map[string]gopurs_runtime.Value)["value4"]))), gopurs_runtime.Apply(f_2, v_4.PtrVal.(map[string]gopurs_runtime.Value)["value3"]))), gopurs_runtime.Apply(go__3_1, v_4.PtrVal.(map[string]gopurs_runtime.Value)["value5"]))
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
}), "sequence": gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(Get_traversableMap().PtrVal.(map[string]gopurs_runtime.Value)["traverse"], dictApplicative_0), pkg_Control_Category.Get_categoryFn().PtrVal.(map[string]gopurs_runtime.Value)["identity"])
}), "Functor0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_functorMap()
}), "Foldable1": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_foldableMap()
})})
	})
	return traversableMap
}

var traversableWithIndexMap gopurs_runtime.Value
var once_traversableWithIndexMap sync.Once
func Get_traversableWithIndexMap() gopurs_runtime.Value {
	once_traversableWithIndexMap.Do(func() {
		traversableWithIndexMap = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"traverseWithIndex": gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
Apply0_1_0 := gopurs_runtime.Apply(dictApplicative_0.PtrVal.(map[string]gopurs_runtime.Value)["Apply0"], gopurs_runtime.Value{})
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
var go__3_1 gopurs_runtime.Value
go__3_1 = gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(v_4.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Leaf")).IntVal != 0 {
__t2 = gopurs_runtime.Apply(dictApplicative_0.PtrVal.(map[string]gopurs_runtime.Value)["pure"], gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Leaf")}))
goto end_branch_2
} else {

}
}
{
if (gopurs_runtime.Bool(v_4.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Node")).IntVal != 0 {
__local_var_5_3 := v_4.PtrVal.(map[string]gopurs_runtime.Value)["value0"]
__local_var_6_4 := v_4.PtrVal.(map[string]gopurs_runtime.Value)["value2"]
__local_var_7_5 := v_4.PtrVal.(map[string]gopurs_runtime.Value)["value1"]
__t2 = gopurs_runtime.Apply(gopurs_runtime.Apply(Apply0_1_0.PtrVal.(map[string]gopurs_runtime.Value)["apply"], gopurs_runtime.Apply(gopurs_runtime.Apply(Apply0_1_0.PtrVal.(map[string]gopurs_runtime.Value)["apply"], gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(Apply0_1_0.PtrVal.(map[string]gopurs_runtime.Value)["Functor0"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["map"], gopurs_runtime.Func(func(l_prime_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_prime_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(r_prime_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Node"), "value0": __local_var_5_3, "value1": __local_var_7_5, "value2": __local_var_6_4, "value3": v_prime_9, "value4": l_prime_8, "value5": r_prime_10})
})
})
})), gopurs_runtime.Apply(go__3_1, v_4.PtrVal.(map[string]gopurs_runtime.Value)["value4"]))), gopurs_runtime.Apply(gopurs_runtime.Apply(f_2, __local_var_6_4), v_4.PtrVal.(map[string]gopurs_runtime.Value)["value3"]))), gopurs_runtime.Apply(go__3_1, v_4.PtrVal.(map[string]gopurs_runtime.Value)["value5"]))
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
}), "FunctorWithIndex0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_functorWithIndexMap()
}), "FoldableWithIndex1": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_foldableWithIndexMap()
}), "Traversable2": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_traversableMap()
})})
	})
	return traversableWithIndexMap
}

var values gopurs_runtime.Value
var once_values sync.Once
func Get_values() gopurs_runtime.Value {
	once_values.Do(func() {
		values = gopurs_runtime.Apply(gopurs_runtime.Apply(Get_foldableMap().PtrVal.(map[string]gopurs_runtime.Value)["foldr"], pkg_Data_List_Types.Get_Cons()), gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nil")}))
	})
	return values
}

var foldSubmapBy gopurs_runtime.Value
var once_foldSubmapBy sync.Once
func Get_foldSubmapBy() gopurs_runtime.Value {
	once_foldSubmapBy.Do(func() {
		foldSubmapBy = gopurs_runtime.Func(func(dictOrd_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(appendFn_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(memptyValue_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(kmin_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(kmax_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_5 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(kmin_3.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just")).IntVal != 0 {
__local_var_6_2 := kmin_3.PtrVal.(map[string]gopurs_runtime.Value)["value0"]
__t1 = gopurs_runtime.Func(func(k_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(gopurs_runtime.Apply(gopurs_runtime.Apply(dictOrd_0.PtrVal.(map[string]gopurs_runtime.Value)["compare"], k_7), __local_var_6_2).PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "LT")
})
goto end_branch_1
} else {

}
}
{
if (gopurs_runtime.Bool(kmin_3.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Nothing")).IntVal != 0 {
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
var __t4 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(kmax_4.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just")).IntVal != 0 {
__local_var_7_5 := kmax_4.PtrVal.(map[string]gopurs_runtime.Value)["value0"]
__t4 = gopurs_runtime.Func(func(k_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(gopurs_runtime.Apply(gopurs_runtime.Apply(dictOrd_0.PtrVal.(map[string]gopurs_runtime.Value)["compare"], k_8), __local_var_7_5).PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "GT")
})
goto end_branch_4
} else {

}
}
{
if (gopurs_runtime.Bool(kmax_4.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Nothing")).IntVal != 0 {
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
var __t7 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(kmin_3.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just")).IntVal != 0 {
var __t8 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(kmax_4.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just")).IntVal != 0 {
__local_var_8_9 := kmax_4.PtrVal.(map[string]gopurs_runtime.Value)["value0"]
__local_var_9_10 := kmin_3.PtrVal.(map[string]gopurs_runtime.Value)["value0"]
__t8 = gopurs_runtime.Func(func(k_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_HeytingAlgebra.Get_boolConj(), gopurs_runtime.Bool(gopurs_runtime.Bool(gopurs_runtime.Apply(gopurs_runtime.Apply(dictOrd_0.PtrVal.(map[string]gopurs_runtime.Value)["compare"], __local_var_9_10), k_10).PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "GT").IntVal == 0)), gopurs_runtime.Bool(gopurs_runtime.Bool(gopurs_runtime.Apply(gopurs_runtime.Apply(dictOrd_0.PtrVal.(map[string]gopurs_runtime.Value)["compare"], k_10), __local_var_8_9).PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "GT").IntVal == 0))
})
goto end_branch_8
} else {

}
}
{
if (gopurs_runtime.Bool(kmax_4.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Nothing")).IntVal != 0 {
__local_var_8_11 := kmin_3.PtrVal.(map[string]gopurs_runtime.Value)["value0"]
__t8 = gopurs_runtime.Func(func(k_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(gopurs_runtime.Bool(gopurs_runtime.Apply(gopurs_runtime.Apply(dictOrd_0.PtrVal.(map[string]gopurs_runtime.Value)["compare"], __local_var_8_11), k_9).PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "GT").IntVal == 0)
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
if (gopurs_runtime.Bool(kmin_3.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Nothing")).IntVal != 0 {
var __t12 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(kmax_4.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just")).IntVal != 0 {
__local_var_8_13 := kmax_4.PtrVal.(map[string]gopurs_runtime.Value)["value0"]
__t12 = gopurs_runtime.Func(func(k_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(gopurs_runtime.Bool(gopurs_runtime.Apply(gopurs_runtime.Apply(dictOrd_0.PtrVal.(map[string]gopurs_runtime.Value)["compare"], k_9), __local_var_8_13).PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "GT").IntVal == 0)
})
goto end_branch_12
} else {

}
}
{
if (gopurs_runtime.Bool(kmax_4.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Nothing")).IntVal != 0 {
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
var go__9_14 gopurs_runtime.Value
go__9_14 = gopurs_runtime.Func(func(v_10 gopurs_runtime.Value) gopurs_runtime.Value {
var __t15 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(v_10.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Leaf")).IntVal != 0 {
__t15 = memptyValue_2
goto end_branch_15
} else {

}
}
{
if (gopurs_runtime.Bool(v_10.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Node")).IntVal != 0 {
var __t16 gopurs_runtime.Value
{
if (gopurs_runtime.Apply(tooSmall_6_0, v_10.PtrVal.(map[string]gopurs_runtime.Value)["value2"])).IntVal != 0 {
__t16 = memptyValue_2
goto end_branch_16
} else {

}
}
{
__t16 = gopurs_runtime.Apply(go__9_14, v_10.PtrVal.(map[string]gopurs_runtime.Value)["value4"])
}
end_branch_16:
var __t17 gopurs_runtime.Value
{
if (gopurs_runtime.Apply(inBounds_8_6, v_10.PtrVal.(map[string]gopurs_runtime.Value)["value2"])).IntVal != 0 {
__t17 = gopurs_runtime.Apply(gopurs_runtime.Apply(f_5, v_10.PtrVal.(map[string]gopurs_runtime.Value)["value2"]), v_10.PtrVal.(map[string]gopurs_runtime.Value)["value3"])
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
if (gopurs_runtime.Apply(tooLarge_7_3, v_10.PtrVal.(map[string]gopurs_runtime.Value)["value2"])).IntVal != 0 {
__t18 = memptyValue_2
goto end_branch_18
} else {

}
}
{
__t18 = gopurs_runtime.Apply(go__9_14, v_10.PtrVal.(map[string]gopurs_runtime.Value)["value5"])
}
end_branch_18:
__t15 = gopurs_runtime.Apply(gopurs_runtime.Apply(appendFn_1, gopurs_runtime.Apply(gopurs_runtime.Apply(appendFn_1, __t16), __t17)), __t18)
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
})
})
})
	})
	return foldSubmapBy
}

var foldSubmap gopurs_runtime.Value
var once_foldSubmap sync.Once
func Get_foldSubmap() gopurs_runtime.Value {
	once_foldSubmap.Do(func() {
		foldSubmap = gopurs_runtime.Func(func(dictOrd_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(dictMonoid_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(Get_foldSubmapBy(), dictOrd_0), gopurs_runtime.Apply(dictMonoid_1.PtrVal.(map[string]gopurs_runtime.Value)["Semigroup0"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["append"]), dictMonoid_1.PtrVal.(map[string]gopurs_runtime.Value)["mempty"])
})
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
var v_0 = v_0_loop
_ = v_0
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(v_0.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Leaf")).IntVal != 0 {
__t0 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")})
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Bool(v_0.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Node")).IntVal != 0 {
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(v_0.PtrVal.(map[string]gopurs_runtime.Value)["value4"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Leaf")).IntVal != 0 {
__t1 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Just"), "value0": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"key": v_0.PtrVal.(map[string]gopurs_runtime.Value)["value2"], "value": v_0.PtrVal.(map[string]gopurs_runtime.Value)["value3"]})})
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Apply(Get_findMin(), v_0.PtrVal.(map[string]gopurs_runtime.Value)["value4"])
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
		lookupGT = gopurs_runtime.Func(func(dictOrd_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_1 gopurs_runtime.Value) gopurs_runtime.Value {
var go__2_0 gopurs_runtime.Value
go__2_0 = gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(v_3.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Leaf")).IntVal != 0 {
__t1 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")})
goto end_branch_1
} else {

}
}
{
if (gopurs_runtime.Bool(v_3.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Node")).IntVal != 0 {
v1_4_2 := gopurs_runtime.Apply(gopurs_runtime.Apply(dictOrd_0.PtrVal.(map[string]gopurs_runtime.Value)["compare"], k_1), v_3.PtrVal.(map[string]gopurs_runtime.Value)["value2"])
var __t3 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(v1_4_2.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "LT")).IntVal != 0 {
v2_5_4 := gopurs_runtime.Apply(go__2_0, v_3.PtrVal.(map[string]gopurs_runtime.Value)["value4"])
var __t5 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(v2_5_4.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Nothing")).IntVal != 0 {
__t5 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Just"), "value0": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"key": v_3.PtrVal.(map[string]gopurs_runtime.Value)["value2"], "value": v_3.PtrVal.(map[string]gopurs_runtime.Value)["value3"]})})
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
if (gopurs_runtime.Bool(v1_4_2.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "GT")).IntVal != 0 {
__t3 = gopurs_runtime.Apply(go__2_0, v_3.PtrVal.(map[string]gopurs_runtime.Value)["value5"])
goto end_branch_3
} else {

}
}
{
if (gopurs_runtime.Bool(v1_4_2.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "EQ")).IntVal != 0 {
__t3 = gopurs_runtime.Apply(Get_findMin(), v_3.PtrVal.(map[string]gopurs_runtime.Value)["value5"])
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
var v_0 = v_0_loop
_ = v_0
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(v_0.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Leaf")).IntVal != 0 {
__t0 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")})
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Bool(v_0.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Node")).IntVal != 0 {
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(v_0.PtrVal.(map[string]gopurs_runtime.Value)["value5"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Leaf")).IntVal != 0 {
__t1 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Just"), "value0": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"key": v_0.PtrVal.(map[string]gopurs_runtime.Value)["value2"], "value": v_0.PtrVal.(map[string]gopurs_runtime.Value)["value3"]})})
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Apply(Get_findMax(), v_0.PtrVal.(map[string]gopurs_runtime.Value)["value5"])
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
		lookupLT = gopurs_runtime.Func(func(dictOrd_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_1 gopurs_runtime.Value) gopurs_runtime.Value {
var go__2_0 gopurs_runtime.Value
go__2_0 = gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(v_3.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Leaf")).IntVal != 0 {
__t1 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")})
goto end_branch_1
} else {

}
}
{
if (gopurs_runtime.Bool(v_3.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Node")).IntVal != 0 {
v1_4_2 := gopurs_runtime.Apply(gopurs_runtime.Apply(dictOrd_0.PtrVal.(map[string]gopurs_runtime.Value)["compare"], k_1), v_3.PtrVal.(map[string]gopurs_runtime.Value)["value2"])
var __t3 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(v1_4_2.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "LT")).IntVal != 0 {
__t3 = gopurs_runtime.Apply(go__2_0, v_3.PtrVal.(map[string]gopurs_runtime.Value)["value4"])
goto end_branch_3
} else {

}
}
{
if (gopurs_runtime.Bool(v1_4_2.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "GT")).IntVal != 0 {
v2_5_4 := gopurs_runtime.Apply(go__2_0, v_3.PtrVal.(map[string]gopurs_runtime.Value)["value5"])
var __t5 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(v2_5_4.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Nothing")).IntVal != 0 {
__t5 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Just"), "value0": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"key": v_3.PtrVal.(map[string]gopurs_runtime.Value)["value2"], "value": v_3.PtrVal.(map[string]gopurs_runtime.Value)["value3"]})})
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
if (gopurs_runtime.Bool(v1_4_2.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "EQ")).IntVal != 0 {
__t3 = gopurs_runtime.Apply(Get_findMax(), v_3.PtrVal.(map[string]gopurs_runtime.Value)["value4"])
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
	})
	return lookupLT
}

var filterWithKey gopurs_runtime.Value
var once_filterWithKey sync.Once
func Get_filterWithKey() gopurs_runtime.Value {
	once_filterWithKey.Do(func() {
		filterWithKey = gopurs_runtime.Func(func(dictOrd_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
var go__2_0 gopurs_runtime.Value
go__2_0 = gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(v_3.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Leaf")).IntVal != 0 {
__t1 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Leaf")})
goto end_branch_1
} else {

}
}
{
if (gopurs_runtime.Bool(v_3.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Node")).IntVal != 0 {
var __t2 gopurs_runtime.Value
{
if (gopurs_runtime.Apply(gopurs_runtime.Apply(f_1, v_3.PtrVal.(map[string]gopurs_runtime.Value)["value2"]), v_3.PtrVal.(map[string]gopurs_runtime.Value)["value3"])).IntVal != 0 {
__t2 = gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Function_Uncurried.Get_runFn4(), Get_unsafeBalancedNode()), v_3.PtrVal.(map[string]gopurs_runtime.Value)["value2"]), v_3.PtrVal.(map[string]gopurs_runtime.Value)["value3"]), gopurs_runtime.Apply(go__2_0, v_3.PtrVal.(map[string]gopurs_runtime.Value)["value4"])), gopurs_runtime.Apply(go__2_0, v_3.PtrVal.(map[string]gopurs_runtime.Value)["value5"]))
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Function_Uncurried.Get_runFn2(), Get_unsafeJoinNodes()), gopurs_runtime.Apply(go__2_0, v_3.PtrVal.(map[string]gopurs_runtime.Value)["value4"])), gopurs_runtime.Apply(go__2_0, v_3.PtrVal.(map[string]gopurs_runtime.Value)["value5"]))
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
	})
	return filterWithKey
}

var filterKeys gopurs_runtime.Value
var once_filterKeys sync.Once
func Get_filterKeys() gopurs_runtime.Value {
	once_filterKeys.Do(func() {
		filterKeys = gopurs_runtime.Func(func(dictOrd_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
var go__2_0 gopurs_runtime.Value
go__2_0 = gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(v_3.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Leaf")).IntVal != 0 {
__t1 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Leaf")})
goto end_branch_1
} else {

}
}
{
if (gopurs_runtime.Bool(v_3.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Node")).IntVal != 0 {
var __t2 gopurs_runtime.Value
{
if (gopurs_runtime.Apply(f_1, v_3.PtrVal.(map[string]gopurs_runtime.Value)["value2"])).IntVal != 0 {
__t2 = gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Function_Uncurried.Get_runFn4(), Get_unsafeBalancedNode()), v_3.PtrVal.(map[string]gopurs_runtime.Value)["value2"]), v_3.PtrVal.(map[string]gopurs_runtime.Value)["value3"]), gopurs_runtime.Apply(go__2_0, v_3.PtrVal.(map[string]gopurs_runtime.Value)["value4"])), gopurs_runtime.Apply(go__2_0, v_3.PtrVal.(map[string]gopurs_runtime.Value)["value5"]))
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Function_Uncurried.Get_runFn2(), Get_unsafeJoinNodes()), gopurs_runtime.Apply(go__2_0, v_3.PtrVal.(map[string]gopurs_runtime.Value)["value4"])), gopurs_runtime.Apply(go__2_0, v_3.PtrVal.(map[string]gopurs_runtime.Value)["value5"]))
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
	})
	return filterKeys
}

var filter gopurs_runtime.Value
var once_filter sync.Once
func Get_filter() gopurs_runtime.Value {
	once_filter.Do(func() {
		filter = gopurs_runtime.Func(func(dictOrd_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(Get_filterWithKey(), dictOrd_0), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return x_1
}))
})
})
	})
	return filter
}

var eqMap gopurs_runtime.Value
var once_eqMap sync.Once
func Get_eqMap() gopurs_runtime.Value {
	once_eqMap.Do(func() {
		eqMap = gopurs_runtime.Func(func(dictEq_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(dictEq1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"eq": gopurs_runtime.Func(func(xs_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(ys_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(xs_2.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Leaf")).IntVal != 0 {
__t0 = gopurs_runtime.Bool(ys_3.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Leaf")
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Bool(xs_2.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Node")).IntVal != 0 {
__t0 = gopurs_runtime.Bool(gopurs_runtime.Bool(ys_3.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Node").IntVal != 0 && gopurs_runtime.Bool(gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Eq.Get_eqIntImpl(), xs_2.PtrVal.(map[string]gopurs_runtime.Value)["value1"]), ys_3.PtrVal.(map[string]gopurs_runtime.Value)["value1"]).IntVal != 0 && gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(Get_eqMapIter(), dictEq_0), dictEq1_1).PtrVal.(map[string]gopurs_runtime.Value)["eq"], gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("IterNode"), "value0": xs_2, "value1": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("IterLeaf")})})), gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("IterNode"), "value0": ys_3, "value1": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("IterLeaf")})})).IntVal != 0).IntVal != 0)
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
})})
})
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
eqMap1_2_1 := gopurs_runtime.Apply(Get_eqMap(), gopurs_runtime.Apply(dictOrd_0.PtrVal.(map[string]gopurs_runtime.Value)["Eq0"], gopurs_runtime.Value{}))
return gopurs_runtime.Func(func(dictOrd1_3 gopurs_runtime.Value) gopurs_runtime.Value {
eqMap2_4_2 := gopurs_runtime.Apply(eqMap1_2_1, gopurs_runtime.Apply(dictOrd1_3.PtrVal.(map[string]gopurs_runtime.Value)["Eq0"], gopurs_runtime.Value{}))
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"compare": gopurs_runtime.Func(func(xs_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(ys_6 gopurs_runtime.Value) gopurs_runtime.Value {
var __t3 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(xs_5.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Leaf")).IntVal != 0 {
var __t4 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(ys_6.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Leaf")).IntVal != 0 {
__t4 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("EQ")})
goto end_branch_4
} else {

}
}
{
__t4 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("LT")})
}
end_branch_4:
__t3 = __t4
goto end_branch_3
} else {

}
}
{
if (gopurs_runtime.Bool(ys_6.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Leaf")).IntVal != 0 {
__t3 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("GT")})
goto end_branch_3
} else {

}
}
{
__t3 = gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(ordMapIter1_1_0, dictOrd1_3).PtrVal.(map[string]gopurs_runtime.Value)["compare"], gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("IterNode"), "value0": xs_5, "value1": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("IterLeaf")})})), gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("IterNode"), "value0": ys_6, "value1": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("IterLeaf")})}))
}
end_branch_3:
return __t3
})
}), "Eq0": gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return eqMap2_4_2
})})
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
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"eq1": gopurs_runtime.Func(func(dictEq1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(Get_eqMap(), dictEq_0), dictEq1_1).PtrVal.(map[string]gopurs_runtime.Value)["eq"]
})})
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
__local_var_2_1 := gopurs_runtime.Apply(dictOrd_0.PtrVal.(map[string]gopurs_runtime.Value)["Eq0"], gopurs_runtime.Value{})
eq1Map1_3_2 := gopurs_runtime.Record(map[string]gopurs_runtime.Value{"eq1": gopurs_runtime.Func(func(dictEq1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(Get_eqMap(), __local_var_2_1), dictEq1_3).PtrVal.(map[string]gopurs_runtime.Value)["eq"]
})})
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"compare1": gopurs_runtime.Func(func(dictOrd1_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(ordMap1_1_0, dictOrd1_4).PtrVal.(map[string]gopurs_runtime.Value)["compare"]
}), "Eq10": gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return eq1Map1_3_2
})})
})
	})
	return ord1Map
}

var empty gopurs_runtime.Value
var once_empty sync.Once
func Get_empty() gopurs_runtime.Value {
	once_empty.Do(func() {
		empty = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Leaf")})
	})
	return empty
}

var fromFoldable gopurs_runtime.Value
var once_fromFoldable sync.Once
func Get_fromFoldable() gopurs_runtime.Value {
	once_fromFoldable.Do(func() {
		fromFoldable = gopurs_runtime.Func(func(dictOrd_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(dictFoldable_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(dictFoldable_1.PtrVal.(map[string]gopurs_runtime.Value)["foldl"], gopurs_runtime.Func(func(m_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(Get_insert(), dictOrd_0), v_3.PtrVal.(map[string]gopurs_runtime.Value)["value0"]), v_3.PtrVal.(map[string]gopurs_runtime.Value)["value1"]), m_2)
})
})), gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Leaf")}))
})
})
	})
	return fromFoldable
}

var fromFoldableWith gopurs_runtime.Value
var once_fromFoldableWith sync.Once
func Get_fromFoldableWith() gopurs_runtime.Value {
	once_fromFoldableWith.Do(func() {
		fromFoldableWith = gopurs_runtime.Func(func(dictOrd_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(dictFoldable_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
f_prime_3_0 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_insertWith(), dictOrd_0), gopurs_runtime.Func(func(b_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(f_2, a_4), b_3)
})
}))
return gopurs_runtime.Apply(gopurs_runtime.Apply(dictFoldable_1.PtrVal.(map[string]gopurs_runtime.Value)["foldl"], gopurs_runtime.Func(func(m_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(f_prime_3_0, v_5.PtrVal.(map[string]gopurs_runtime.Value)["value0"]), v_5.PtrVal.(map[string]gopurs_runtime.Value)["value1"]), m_4)
})
})), gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Leaf")}))
})
})
})
	})
	return fromFoldableWith
}

var fromFoldableWithIndex gopurs_runtime.Value
var once_fromFoldableWithIndex sync.Once
func Get_fromFoldableWithIndex() gopurs_runtime.Value {
	once_fromFoldableWithIndex.Do(func() {
		fromFoldableWithIndex = gopurs_runtime.Func(func(dictOrd_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(dictFoldableWithIndex_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(dictFoldableWithIndex_1.PtrVal.(map[string]gopurs_runtime.Value)["foldlWithIndex"], gopurs_runtime.Func(func(k_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(Get_insert(), dictOrd_0), k_2), v_4), m_3)
})
})
})), gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Leaf")}))
})
})
	})
	return fromFoldableWithIndex
}

var monoidSemigroupMap gopurs_runtime.Value
var once_monoidSemigroupMap sync.Once
func Get_monoidSemigroupMap() gopurs_runtime.Value {
	once_monoidSemigroupMap.Do(func() {
		monoidSemigroupMap = gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(dictOrd_1 gopurs_runtime.Value) gopurs_runtime.Value {
semigroupMap2_2_0 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_semigroupMap(), gopurs_runtime.Value{}), dictOrd_1)
return gopurs_runtime.Func(func(dictSemigroup_3 gopurs_runtime.Value) gopurs_runtime.Value {
semigroupMap3_4_1 := gopurs_runtime.Apply(semigroupMap2_2_0, dictSemigroup_3)
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"mempty": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Leaf")}), "Semigroup0": gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return semigroupMap3_4_1
})})
})
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
compare_1_0 := dictOrd_0.PtrVal.(map[string]gopurs_runtime.Value)["compare"]
return gopurs_runtime.Func(func(kmin_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(kmax_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(Get_foldSubmapBy(), dictOrd_0), gopurs_runtime.Func(func(m1_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m2_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Function_Uncurried.Get_runFn4(), Get_unsafeUnionWith()), compare_1_0), pkg_Data_Function.Get_const_()), m1_4), m2_5)
})
})), gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Leaf")})), kmin_2), kmax_3), Get_singleton())
})
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
compare_1_0 := dictOrd_0.PtrVal.(map[string]gopurs_runtime.Value)["compare"]
return gopurs_runtime.Func(func(dictFoldable_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(dictFoldable_2.PtrVal.(map[string]gopurs_runtime.Value)["foldl"], gopurs_runtime.Func(func(m1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m2_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Function_Uncurried.Get_runFn4(), Get_unsafeUnionWith()), compare_1_0), pkg_Data_Function.Get_const_()), m1_3), m2_4)
})
})), gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Leaf")}))
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
compare_1_0 := dictOrd_0.PtrVal.(map[string]gopurs_runtime.Value)["compare"]
return gopurs_runtime.Func(func(m1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m2_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Function_Uncurried.Get_runFn3(), Get_unsafeDifference()), compare_1_0), m1_2), m2_3)
})
})
})
	})
	return difference
}

var delete_ gopurs_runtime.Value
var once_delete_ sync.Once
func Get_delete_() gopurs_runtime.Value {
	once_delete_.Do(func() {
		delete_ = gopurs_runtime.Func(func(dictOrd_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_1 gopurs_runtime.Value) gopurs_runtime.Value {
var go__2_0 gopurs_runtime.Value
go__2_0 = gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(v_3.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Leaf")).IntVal != 0 {
__t1 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Leaf")})
goto end_branch_1
} else {

}
}
{
if (gopurs_runtime.Bool(v_3.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Node")).IntVal != 0 {
v1_4_2 := gopurs_runtime.Apply(gopurs_runtime.Apply(dictOrd_0.PtrVal.(map[string]gopurs_runtime.Value)["compare"], k_1), v_3.PtrVal.(map[string]gopurs_runtime.Value)["value2"])
var __t3 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(v1_4_2.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "LT")).IntVal != 0 {
__t3 = gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Function_Uncurried.Get_runFn4(), Get_unsafeBalancedNode()), v_3.PtrVal.(map[string]gopurs_runtime.Value)["value2"]), v_3.PtrVal.(map[string]gopurs_runtime.Value)["value3"]), gopurs_runtime.Apply(go__2_0, v_3.PtrVal.(map[string]gopurs_runtime.Value)["value4"])), v_3.PtrVal.(map[string]gopurs_runtime.Value)["value5"])
goto end_branch_3
} else {

}
}
{
if (gopurs_runtime.Bool(v1_4_2.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "GT")).IntVal != 0 {
__t3 = gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Function_Uncurried.Get_runFn4(), Get_unsafeBalancedNode()), v_3.PtrVal.(map[string]gopurs_runtime.Value)["value2"]), v_3.PtrVal.(map[string]gopurs_runtime.Value)["value3"]), v_3.PtrVal.(map[string]gopurs_runtime.Value)["value4"]), gopurs_runtime.Apply(go__2_0, v_3.PtrVal.(map[string]gopurs_runtime.Value)["value5"]))
goto end_branch_3
} else {

}
}
{
if (gopurs_runtime.Bool(v1_4_2.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "EQ")).IntVal != 0 {
__t3 = gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Function_Uncurried.Get_runFn2(), Get_unsafeJoinNodes()), v_3.PtrVal.(map[string]gopurs_runtime.Value)["value4"]), v_3.PtrVal.(map[string]gopurs_runtime.Value)["value5"])
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
	})
	return delete_
}

var checkValid gopurs_runtime.Value
var once_checkValid sync.Once
func Get_checkValid() gopurs_runtime.Value {
	once_checkValid.Do(func() {
		checkValid = gopurs_runtime.Func(func(dictOrd_0 gopurs_runtime.Value) gopurs_runtime.Value {
var go__1_0 gopurs_runtime.Value
go__1_0 = gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(v_2.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Leaf")).IntVal != 0 {
__t1 = gopurs_runtime.Bool(true)
goto end_branch_1
} else {

}
}
{
if (gopurs_runtime.Bool(v_2.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Node")).IntVal != 0 {
var __t2 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(v_2.PtrVal.(map[string]gopurs_runtime.Value)["value4"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Leaf")).IntVal != 0 {
var __t3 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(v_2.PtrVal.(map[string]gopurs_runtime.Value)["value5"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Leaf")).IntVal != 0 {
__t3 = gopurs_runtime.Bool(true)
goto end_branch_3
} else {

}
}
{
if (gopurs_runtime.Bool(v_2.PtrVal.(map[string]gopurs_runtime.Value)["value5"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Node")).IntVal != 0 {
__t3 = gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_HeytingAlgebra.Get_boolConj(), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Eq.Get_eqIntImpl(), v_2.PtrVal.(map[string]gopurs_runtime.Value)["value0"]), gopurs_runtime.Int(2))), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_HeytingAlgebra.Get_boolConj(), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Eq.Get_eqIntImpl(), v_2.PtrVal.(map[string]gopurs_runtime.Value)["value5"].PtrVal.(map[string]gopurs_runtime.Value)["value0"]), gopurs_runtime.Int(1))), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_HeytingAlgebra.Get_boolConj(), gopurs_runtime.Bool(gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Ord.Get_ordInt().PtrVal.(map[string]gopurs_runtime.Value)["compare"], v_2.PtrVal.(map[string]gopurs_runtime.Value)["value1"]), v_2.PtrVal.(map[string]gopurs_runtime.Value)["value5"].PtrVal.(map[string]gopurs_runtime.Value)["value1"]).PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "GT")), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_HeytingAlgebra.Get_boolConj(), gopurs_runtime.Bool(gopurs_runtime.Apply(gopurs_runtime.Apply(dictOrd_0.PtrVal.(map[string]gopurs_runtime.Value)["compare"], v_2.PtrVal.(map[string]gopurs_runtime.Value)["value5"].PtrVal.(map[string]gopurs_runtime.Value)["value2"]), v_2.PtrVal.(map[string]gopurs_runtime.Value)["value2"]).PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "GT")), gopurs_runtime.Apply(go__1_0, v_2.PtrVal.(map[string]gopurs_runtime.Value)["value5"])))))
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
if (gopurs_runtime.Bool(v_2.PtrVal.(map[string]gopurs_runtime.Value)["value4"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Node")).IntVal != 0 {
var __t4 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(v_2.PtrVal.(map[string]gopurs_runtime.Value)["value5"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Leaf")).IntVal != 0 {
__t4 = gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_HeytingAlgebra.Get_boolConj(), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Eq.Get_eqIntImpl(), v_2.PtrVal.(map[string]gopurs_runtime.Value)["value0"]), gopurs_runtime.Int(2))), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_HeytingAlgebra.Get_boolConj(), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Eq.Get_eqIntImpl(), v_2.PtrVal.(map[string]gopurs_runtime.Value)["value4"].PtrVal.(map[string]gopurs_runtime.Value)["value0"]), gopurs_runtime.Int(1))), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_HeytingAlgebra.Get_boolConj(), gopurs_runtime.Bool(gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Ord.Get_ordInt().PtrVal.(map[string]gopurs_runtime.Value)["compare"], v_2.PtrVal.(map[string]gopurs_runtime.Value)["value1"]), v_2.PtrVal.(map[string]gopurs_runtime.Value)["value4"].PtrVal.(map[string]gopurs_runtime.Value)["value1"]).PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "GT")), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_HeytingAlgebra.Get_boolConj(), gopurs_runtime.Bool(gopurs_runtime.Apply(gopurs_runtime.Apply(dictOrd_0.PtrVal.(map[string]gopurs_runtime.Value)["compare"], v_2.PtrVal.(map[string]gopurs_runtime.Value)["value4"].PtrVal.(map[string]gopurs_runtime.Value)["value2"]), v_2.PtrVal.(map[string]gopurs_runtime.Value)["value2"]).PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "LT")), gopurs_runtime.Apply(go__1_0, v_2.PtrVal.(map[string]gopurs_runtime.Value)["value4"])))))
goto end_branch_4
} else {

}
}
{
if (gopurs_runtime.Bool(v_2.PtrVal.(map[string]gopurs_runtime.Value)["value5"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Node")).IntVal != 0 {
__local_var_3_5 := gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Ring.Get_intSub(), v_2.PtrVal.(map[string]gopurs_runtime.Value)["value5"].PtrVal.(map[string]gopurs_runtime.Value)["value0"]), v_2.PtrVal.(map[string]gopurs_runtime.Value)["value4"].PtrVal.(map[string]gopurs_runtime.Value)["value0"])
var __t6 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(gopurs_runtime.Bool(gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Ord.Get_ordInt().PtrVal.(map[string]gopurs_runtime.Value)["compare"], __local_var_3_5), gopurs_runtime.Int(0)).PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "LT").IntVal == 0)).IntVal != 0 {
__t6 = __local_var_3_5
goto end_branch_6
} else {

}
}
{
__t6 = gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Ring.Get_intSub(), gopurs_runtime.Int(0)), __local_var_3_5)
}
end_branch_6:
__t4 = gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_HeytingAlgebra.Get_boolConj(), gopurs_runtime.Bool(gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Ord.Get_ordInt().PtrVal.(map[string]gopurs_runtime.Value)["compare"], v_2.PtrVal.(map[string]gopurs_runtime.Value)["value0"]), v_2.PtrVal.(map[string]gopurs_runtime.Value)["value5"].PtrVal.(map[string]gopurs_runtime.Value)["value0"]).PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "GT")), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_HeytingAlgebra.Get_boolConj(), gopurs_runtime.Bool(gopurs_runtime.Apply(gopurs_runtime.Apply(dictOrd_0.PtrVal.(map[string]gopurs_runtime.Value)["compare"], v_2.PtrVal.(map[string]gopurs_runtime.Value)["value5"].PtrVal.(map[string]gopurs_runtime.Value)["value2"]), v_2.PtrVal.(map[string]gopurs_runtime.Value)["value2"]).PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "GT")), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_HeytingAlgebra.Get_boolConj(), gopurs_runtime.Bool(gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Ord.Get_ordInt().PtrVal.(map[string]gopurs_runtime.Value)["compare"], v_2.PtrVal.(map[string]gopurs_runtime.Value)["value0"]), v_2.PtrVal.(map[string]gopurs_runtime.Value)["value4"].PtrVal.(map[string]gopurs_runtime.Value)["value0"]).PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "GT")), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_HeytingAlgebra.Get_boolConj(), gopurs_runtime.Bool(gopurs_runtime.Apply(gopurs_runtime.Apply(dictOrd_0.PtrVal.(map[string]gopurs_runtime.Value)["compare"], v_2.PtrVal.(map[string]gopurs_runtime.Value)["value4"].PtrVal.(map[string]gopurs_runtime.Value)["value2"]), v_2.PtrVal.(map[string]gopurs_runtime.Value)["value2"]).PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "LT")), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_HeytingAlgebra.Get_boolConj(), gopurs_runtime.Bool(gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Ord.Get_ordInt().PtrVal.(map[string]gopurs_runtime.Value)["compare"], __t6), gopurs_runtime.Int(2)).PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "LT")), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_HeytingAlgebra.Get_boolConj(), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Eq.Get_eqIntImpl(), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Semiring.Get_intAdd(), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Semiring.Get_intAdd(), v_2.PtrVal.(map[string]gopurs_runtime.Value)["value5"].PtrVal.(map[string]gopurs_runtime.Value)["value1"]), v_2.PtrVal.(map[string]gopurs_runtime.Value)["value4"].PtrVal.(map[string]gopurs_runtime.Value)["value1"])), gopurs_runtime.Int(1))), v_2.PtrVal.(map[string]gopurs_runtime.Value)["value1"])), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_HeytingAlgebra.Get_boolConj(), gopurs_runtime.Apply(go__1_0, v_2.PtrVal.(map[string]gopurs_runtime.Value)["value4"])), gopurs_runtime.Apply(go__1_0, v_2.PtrVal.(map[string]gopurs_runtime.Value)["value5"]))))))))
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
return gopurs_runtime.Apply(gopurs_runtime.Apply(Get_mapMaybeWithKey(), dictOrd_0), gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Control_Category.Get_categoryFn().PtrVal.(map[string]gopurs_runtime.Value)["identity"]
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
compare_1_0 := dictOrd_0.PtrVal.(map[string]gopurs_runtime.Value)["compare"]
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"apply": gopurs_runtime.Func(func(m1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m2_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Function_Uncurried.Get_runFn4(), Get_unsafeIntersectionWith()), compare_1_0), pkg_Control_Category.Get_categoryFn().PtrVal.(map[string]gopurs_runtime.Value)["identity"]), m1_2), m2_3)
})
}), "Functor0": gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_functorMap()
})})
})
	})
	return applyMap
}

var bindMap gopurs_runtime.Value
var once_bindMap sync.Once
func Get_bindMap() gopurs_runtime.Value {
	once_bindMap.Do(func() {
		bindMap = gopurs_runtime.Func(func(dictOrd_0 gopurs_runtime.Value) gopurs_runtime.Value {
applyMap1_1_0 := gopurs_runtime.Apply(Get_applyMap(), dictOrd_0)
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"bind": gopurs_runtime.Func(func(m_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(Get_mapMaybeWithKey(), dictOrd_0), gopurs_runtime.Func(func(k_4 gopurs_runtime.Value) gopurs_runtime.Value {
var go__5_1 gopurs_runtime.Value
go__5_1 = gopurs_runtime.Func(func(v_6_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
go__5_1:
for {
if false { continue go__5_1 }
var v_6 = v_6_loop
_ = v_6
var __t2 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(v_6.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Leaf")).IntVal != 0 {
__t2 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")})
goto end_branch_2
} else {

}
}
{
if (gopurs_runtime.Bool(v_6.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Node")).IntVal != 0 {
v1_7_3 := gopurs_runtime.Apply(gopurs_runtime.Apply(dictOrd_0.PtrVal.(map[string]gopurs_runtime.Value)["compare"], k_4), v_6.PtrVal.(map[string]gopurs_runtime.Value)["value2"])
var __t4 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(v1_7_3.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "LT")).IntVal != 0 {
v_6_loop = v_6.PtrVal.(map[string]gopurs_runtime.Value)["value4"]
continue go__5_1
__t4 = gopurs_runtime.Value{}
goto end_branch_4
} else {

}
}
{
if (gopurs_runtime.Bool(v1_7_3.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "GT")).IntVal != 0 {
v_6_loop = v_6.PtrVal.(map[string]gopurs_runtime.Value)["value5"]
continue go__5_1
__t4 = gopurs_runtime.Value{}
goto end_branch_4
} else {

}
}
{
if (gopurs_runtime.Bool(v1_7_3.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "EQ")).IntVal != 0 {
__t4 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Just"), "value0": v_6.PtrVal.(map[string]gopurs_runtime.Value)["value3"]})
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
return __t2
}
}()
})
return gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(go__5_1, gopurs_runtime.Apply(f_3, x_6))
})
})), m_2)
})
}), "Apply0": gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return applyMap1_1_0
})})
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
go__1_0 = gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(v_2.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Leaf")).IntVal != 0 {
__t1 = gopurs_runtime.Bool(false)
goto end_branch_1
} else {

}
}
{
if (gopurs_runtime.Bool(v_2.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Node")).IntVal != 0 {
__t1 = gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_HeytingAlgebra.Get_boolDisj(), gopurs_runtime.Apply(gopurs_runtime.Apply(predicate_0, v_2.PtrVal.(map[string]gopurs_runtime.Value)["value2"]), v_2.PtrVal.(map[string]gopurs_runtime.Value)["value3"])), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_HeytingAlgebra.Get_boolDisj(), gopurs_runtime.Apply(go__1_0, v_2.PtrVal.(map[string]gopurs_runtime.Value)["value4"])), gopurs_runtime.Apply(go__1_0, v_2.PtrVal.(map[string]gopurs_runtime.Value)["value5"])))
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
go__1_0 = gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(v_2.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Leaf")).IntVal != 0 {
__t1 = gopurs_runtime.Bool(false)
goto end_branch_1
} else {

}
}
{
if (gopurs_runtime.Bool(v_2.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Node")).IntVal != 0 {
__t1 = gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_HeytingAlgebra.Get_boolDisj(), gopurs_runtime.Apply(predicate_0, v_2.PtrVal.(map[string]gopurs_runtime.Value)["value3"])), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_HeytingAlgebra.Get_boolDisj(), gopurs_runtime.Apply(go__1_0, v_2.PtrVal.(map[string]gopurs_runtime.Value)["value4"])), gopurs_runtime.Apply(go__1_0, v_2.PtrVal.(map[string]gopurs_runtime.Value)["value5"])))
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
compare_1_0 := dictOrd_0.PtrVal.(map[string]gopurs_runtime.Value)["compare"]
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m_4 gopurs_runtime.Value) gopurs_runtime.Value {
v_5_1 := gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Function_Uncurried.Get_runFn3(), Get_unsafeSplit()), compare_1_0), k_3), m_4)
v2_6_2 := gopurs_runtime.Apply(f_2, v_5_1.PtrVal.(map[string]gopurs_runtime.Value)["value0"])
var __t3 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(v2_6_2.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Nothing")).IntVal != 0 {
__t3 = gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Function_Uncurried.Get_runFn2(), Get_unsafeJoinNodes()), v_5_1.PtrVal.(map[string]gopurs_runtime.Value)["value1"]), v_5_1.PtrVal.(map[string]gopurs_runtime.Value)["value2"])
goto end_branch_3
} else {

}
}
{
if (gopurs_runtime.Bool(v2_6_2.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just")).IntVal != 0 {
__t3 = gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Function_Uncurried.Get_runFn4(), Get_unsafeBalancedNode()), k_3), v2_6_2.PtrVal.(map[string]gopurs_runtime.Value)["value0"]), v_5_1.PtrVal.(map[string]gopurs_runtime.Value)["value1"]), v_5_1.PtrVal.(map[string]gopurs_runtime.Value)["value2"])
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
})
	})
	return alter
}

var altMap gopurs_runtime.Value
var once_altMap sync.Once
func Get_altMap() gopurs_runtime.Value {
	once_altMap.Do(func() {
		altMap = gopurs_runtime.Func(func(dictOrd_0 gopurs_runtime.Value) gopurs_runtime.Value {
compare_1_0 := dictOrd_0.PtrVal.(map[string]gopurs_runtime.Value)["compare"]
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"alt": gopurs_runtime.Func(func(m1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m2_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Function_Uncurried.Get_runFn4(), Get_unsafeUnionWith()), compare_1_0), pkg_Data_Function.Get_const_()), m1_2), m2_3)
})
}), "Functor0": gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_functorMap()
})})
})
	})
	return altMap
}

var plusMap gopurs_runtime.Value
var once_plusMap sync.Once
func Get_plusMap() gopurs_runtime.Value {
	once_plusMap.Do(func() {
		plusMap = gopurs_runtime.Func(func(dictOrd_0 gopurs_runtime.Value) gopurs_runtime.Value {
altMap1_1_0 := gopurs_runtime.Apply(Get_altMap(), dictOrd_0)
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"empty": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Leaf")}), "Alt0": gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return altMap1_1_0
})})
})
	})
	return plusMap
}


