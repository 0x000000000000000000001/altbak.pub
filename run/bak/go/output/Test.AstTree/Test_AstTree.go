package Test_AstTree

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Effect_Console "gopurs/output/Effect.Console"
	pkg_Effect "gopurs/output/Effect"
	pkg_Bench "gopurs/output/Bench"
	pkg_Data_Show "gopurs/output/Data.Show"
	unsafe "unsafe"
)

var cache_Val gopurs_runtime.Value
var once_Val sync.Once
func Get_Val() gopurs_runtime.Value {
	once_Val.Do(func() {
		cache_Val = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 245188743, UnsafePtr: unsafe.Pointer(&Constructor_Val{value0.IntVal})}
})
	})
	return cache_Val
}

var cache_Add gopurs_runtime.Value
var once_Add sync.Once
func Get_Add() gopurs_runtime.Value {
	once_Add.Do(func() {
		cache_Add = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(value1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2937956733, UnsafePtr: unsafe.Pointer(&Constructor_Add{value0, value1})}
})
})
	})
	return cache_Add
}

var cache_Mul gopurs_runtime.Value
var once_Mul sync.Once
func Get_Mul() gopurs_runtime.Value {
	once_Mul.Do(func() {
		cache_Mul = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(value1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3406566728, UnsafePtr: unsafe.Pointer(&Constructor_Mul{value0, value1})}
})
})
	})
	return cache_Mul
}

var cache_Sub gopurs_runtime.Value
var once_Sub sync.Once
func Get_Sub() gopurs_runtime.Value {
	once_Sub.Do(func() {
		cache_Sub = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(value1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2029887576, UnsafePtr: unsafe.Pointer(&Constructor_Sub{value0, value1})}
})
})
	})
	return cache_Sub
}

var cache_eval gopurs_runtime.Value
var once_eval sync.Once
func Get_eval() gopurs_runtime.Value {
	once_eval.Do(func() {
		cache_eval = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int(Call_eval(v_0_box))
})
	})
	return cache_eval
}

var cache_describe gopurs_runtime.Value
var once_describe sync.Once
func Get_describe() gopurs_runtime.Value {
	once_describe.Do(func() {
		cache_describe = gopurs_runtime.Func0(func() gopurs_runtime.Value {
return func() gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Effect_Console.Get_log(), gopurs_runtime.Str("AST Evaluation:")), nil)
}()
})
	})
	return cache_describe
}

var cache_buildTree gopurs_runtime.Value
var once_buildTree sync.Once
func Get_buildTree() gopurs_runtime.Value {
	once_buildTree.Do(func() {
		cache_buildTree = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_buildTree(v_0_box.IntVal)
})
	})
	return cache_buildTree
}

var cache_act gopurs_runtime.Value
var once_act sync.Once
func Get_act() gopurs_runtime.Value {
	once_act.Do(func() {
		cache_act = gopurs_runtime.Func0(func() gopurs_runtime.Value {
return func() gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Effect.Get_bindEffect(), "bind"), gopurs_runtime.Apply(pkg_Bench.Get_opaque(), gopurs_runtime.Int(3)), gopurs_runtime.Func(func(dummy_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pkg_Effect_Console.Get_log(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Show.Get_showInt(), "show"), gopurs_runtime.Int(Call_eval(Call_buildTree(dummy_0.IntVal)))))
})), nil)
}()
})
	})
	return cache_act
}

type Constructor_Val struct {
	V0 int64
}


type Constructor_Add struct {
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
}


type Constructor_Mul struct {
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
}


type Constructor_Sub struct {
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
}


func Call_eval(v_0_loop gopurs_runtime.Value) int64 {
eval:
for {
if false { continue eval }
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var __t0 gopurs_runtime.Value
{
if (v_0.Type == 9 && v_0.IntVal == 245188743) {
__t0 = gopurs_runtime.Int((*Constructor_Val)(v_0.UnsafePtr).V0)
goto end_branch_0
} else {

}
}
{
if (v_0.Type == 9 && v_0.IntVal == 2937956733) {
__t0 = gopurs_runtime.Int((gopurs_runtime.Int(Call_eval((*Constructor_Add)(v_0.UnsafePtr).V0)).IntVal) + (gopurs_runtime.Int(Call_eval((*Constructor_Add)(v_0.UnsafePtr).V1)).IntVal))
goto end_branch_0
} else {

}
}
{
if (v_0.Type == 9 && v_0.IntVal == 3406566728) {
__t0 = gopurs_runtime.Int((gopurs_runtime.Int(Call_eval((*Constructor_Mul)(v_0.UnsafePtr).V0)).IntVal) * (gopurs_runtime.Int(Call_eval((*Constructor_Mul)(v_0.UnsafePtr).V1)).IntVal))
goto end_branch_0
} else {

}
}
{
if (v_0.Type == 9 && v_0.IntVal == 2029887576) {
__t0 = gopurs_runtime.Int((gopurs_runtime.Int(Call_eval((*Constructor_Sub)(v_0.UnsafePtr).V0)).IntVal) - (gopurs_runtime.Int(Call_eval((*Constructor_Sub)(v_0.UnsafePtr).V1)).IntVal))
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
}

func Call_buildTree(v_0_loop int64) gopurs_runtime.Value {
buildTree:
for {
if false { continue buildTree }
var v_0 int64 = v_0_loop
_ = v_0
var __t0 gopurs_runtime.Value
{
if (v_0) == (0) {
__t0 = gopurs_runtime.Any(gopurs_runtime.Value{Type: 9, IntVal: 245188743, UnsafePtr: unsafe.Pointer(&Constructor_Val{1})})
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Any(gopurs_runtime.Value{Type: 9, IntVal: 2937956733, UnsafePtr: unsafe.Pointer(&Constructor_Add{gopurs_runtime.Any(gopurs_runtime.Value{Type: 9, IntVal: 3406566728, UnsafePtr: unsafe.Pointer(&Constructor_Mul{gopurs_runtime.Any(gopurs_runtime.Value{Type: 9, IntVal: 245188743, UnsafePtr: unsafe.Pointer(&Constructor_Val{v_0})}), Call_buildTree((v_0) - (1))})}), gopurs_runtime.Any(gopurs_runtime.Value{Type: 9, IntVal: 2029887576, UnsafePtr: unsafe.Pointer(&Constructor_Sub{Call_buildTree((v_0) - (1)), gopurs_runtime.Any(gopurs_runtime.Value{Type: 9, IntVal: 245188743, UnsafePtr: unsafe.Pointer(&Constructor_Val{1})})})})})})
}
end_branch_0:
return __t0
}
}
