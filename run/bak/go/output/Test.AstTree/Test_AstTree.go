package Test_AstTree

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Effect_Console "gopurs/output/Effect.Console"
	pkg_Bench "gopurs/output/Bench"
	pkg_Data_Show "gopurs/output/Data.Show"
)

var Val gopurs_runtime.Value
var once_Val sync.Once
func Get_Val() gopurs_runtime.Value {
	once_Val.Do(func() {
		Val = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Constructor1("Val", value0)
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
return gopurs_runtime.Constructor2("Add", value0, value1)
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
return gopurs_runtime.Constructor2("Mul", value0, value1)
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
return gopurs_runtime.Constructor2("Sub", value0, value1)
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
if (gopurs_runtime.Bool(v_0.StrVal == "Val")).IntVal != 0 {
__t0 = gopurs_runtime.ConstructorGet(v_0, 0)
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Bool(v_0.StrVal == "Add")).IntVal != 0 {
__t0 = gopurs_runtime.Int(gopurs_runtime.Apply(Get_eval(), gopurs_runtime.ConstructorGet(v_0, 0)).IntVal + gopurs_runtime.Apply(Get_eval(), gopurs_runtime.ConstructorGet(v_0, 1)).IntVal)
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Bool(v_0.StrVal == "Mul")).IntVal != 0 {
__t0 = gopurs_runtime.Int(gopurs_runtime.Apply(Get_eval(), gopurs_runtime.ConstructorGet(v_0, 0)).IntVal * gopurs_runtime.Apply(Get_eval(), gopurs_runtime.ConstructorGet(v_0, 1)).IntVal)
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Bool(v_0.StrVal == "Sub")).IntVal != 0 {
__t0 = gopurs_runtime.Int(gopurs_runtime.Apply(Get_eval(), gopurs_runtime.ConstructorGet(v_0, 0)).IntVal - gopurs_runtime.Apply(Get_eval(), gopurs_runtime.ConstructorGet(v_0, 1)).IntVal)
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
__t0 = gopurs_runtime.Constructor1("Val", gopurs_runtime.Int(1))
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Constructor2("Add", gopurs_runtime.Constructor2("Mul", gopurs_runtime.Constructor1("Val", v_0), gopurs_runtime.Apply(Get_buildTree(), gopurs_runtime.Int(v_0.IntVal - gopurs_runtime.Int(1).IntVal))), gopurs_runtime.Constructor2("Sub", gopurs_runtime.Apply(Get_buildTree(), gopurs_runtime.Int(v_0.IntVal - gopurs_runtime.Int(1).IntVal)), gopurs_runtime.Constructor1("Val", gopurs_runtime.Int(1))))
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
		act = func() gopurs_runtime.Value {
__local_var_0_0 := gopurs_runtime.Apply(pkg_Bench.Get_opaque(), gopurs_runtime.Int(3))
_ = __local_var_0_0
return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
dummy_1_1 := gopurs_runtime.Apply(__local_var_0_0, gopurs_runtime.Value{})
_ = dummy_1_1
return gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Effect_Console.Get_log(), gopurs_runtime.Apply(pkg_Data_Show.Get_showIntImpl(), gopurs_runtime.Apply(Get_eval(), gopurs_runtime.Apply(Get_buildTree(), dummy_1_1)))), gopurs_runtime.Value{})
})
}()
	})
	return act
}


