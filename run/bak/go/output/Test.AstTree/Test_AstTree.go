package Test_AstTree

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_Semiring "gopurs/output/Data.Semiring"
	pkg_Data_Ring "gopurs/output/Data.Ring"
	pkg_Effect_Console "gopurs/output/Effect.Console"
	pkg_Data_Show "gopurs/output/Data.Show"
)

var Val gopurs_runtime.Value
var once_Val sync.Once
func Get_Val() gopurs_runtime.Value {
	once_Val.Do(func() {
		Val = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Val"), "value0": value0})
})
	})
	return Val
}

var Add gopurs_runtime.Value
var once_Add sync.Once
func Get_Add() gopurs_runtime.Value {
	once_Add.Do(func() {
		Add = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(value1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Add"), "value0": value0, "value1": value1})
})
})
	})
	return Add
}

var Mul gopurs_runtime.Value
var once_Mul sync.Once
func Get_Mul() gopurs_runtime.Value {
	once_Mul.Do(func() {
		Mul = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(value1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Mul"), "value0": value0, "value1": value1})
})
})
	})
	return Mul
}

var Sub gopurs_runtime.Value
var once_Sub sync.Once
func Get_Sub() gopurs_runtime.Value {
	once_Sub.Do(func() {
		Sub = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(value1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Sub"), "value0": value0, "value1": value1})
})
})
	})
	return Sub
}

var eval gopurs_runtime.Value
var once_eval sync.Once
func Get_eval() gopurs_runtime.Value {
	once_eval.Do(func() {
		eval = gopurs_runtime.Func(func(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
eval:
for {
if false { continue eval }
var v_0 = v_0_loop
_ = v_0
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(v_0.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Val")).IntVal != 0 {
__t0 = v_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"]
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Bool(v_0.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Add")).IntVal != 0 {
__t0 = gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Semiring.Get_intAdd(), gopurs_runtime.Apply(Get_eval(), v_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"])), gopurs_runtime.Apply(Get_eval(), v_0.PtrVal.(map[string]gopurs_runtime.Value)["value1"]))
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Bool(v_0.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Mul")).IntVal != 0 {
__t0 = gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Semiring.Get_intMul(), gopurs_runtime.Apply(Get_eval(), v_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"])), gopurs_runtime.Apply(Get_eval(), v_0.PtrVal.(map[string]gopurs_runtime.Value)["value1"]))
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Bool(v_0.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Sub")).IntVal != 0 {
__t0 = gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Ring.Get_intSub(), gopurs_runtime.Apply(Get_eval(), v_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"])), gopurs_runtime.Apply(Get_eval(), v_0.PtrVal.(map[string]gopurs_runtime.Value)["value1"]))
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
	return eval
}

var describe gopurs_runtime.Value
var once_describe sync.Once
func Get_describe() gopurs_runtime.Value {
	once_describe.Do(func() {
		describe = gopurs_runtime.Apply(pkg_Effect_Console.Get_log(), gopurs_runtime.Str("AST Evaluation:"))
	})
	return describe
}

var buildTree gopurs_runtime.Value
var once_buildTree sync.Once
func Get_buildTree() gopurs_runtime.Value {
	once_buildTree.Do(func() {
		buildTree = gopurs_runtime.Func(func(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
buildTree:
for {
if false { continue buildTree }
var v_0 = v_0_loop
_ = v_0
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(v_0.IntVal == gopurs_runtime.Int(0).IntVal)).IntVal != 0 {
__t0 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Val"), "value0": gopurs_runtime.Int(1)})
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Add"), "value0": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Mul"), "value0": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Val"), "value0": v_0}), "value1": gopurs_runtime.Apply(Get_buildTree(), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Ring.Get_intSub(), v_0), gopurs_runtime.Int(1)))}), "value1": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Sub"), "value0": gopurs_runtime.Apply(Get_buildTree(), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Ring.Get_intSub(), v_0), gopurs_runtime.Int(1))), "value1": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Val"), "value0": gopurs_runtime.Int(1)})})})
}
end_branch_0:
return __t0
}
}()
})
	})
	return buildTree
}

var act gopurs_runtime.Value
var once_act sync.Once
func Get_act() gopurs_runtime.Value {
	once_act.Do(func() {
		act = gopurs_runtime.Apply(pkg_Effect_Console.Get_log(), gopurs_runtime.Apply(pkg_Data_Show.Get_showIntImpl(), gopurs_runtime.Apply(Get_eval(), gopurs_runtime.Apply(Get_buildTree(), gopurs_runtime.Int(3)))))
	})
	return act
}


