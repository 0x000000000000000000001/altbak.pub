package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Test_AstTree_logShow gopurs_runtime.Value
var once_Test_AstTree_logShow sync.Once

func Get_Test_AstTree_logShow() gopurs_runtime.Value {
	once_Test_AstTree_logShow.Do(func() {
		cache_Test_AstTree_logShow = gopurs_runtime.Func(func(a_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Test_AstTree_logShow(a_0_box.IntVal)
		})
	})
	return cache_Test_AstTree_logShow
}

var cache_Test_AstTree_Val gopurs_runtime.Value
var once_Test_AstTree_Val sync.Once

func Get_Test_AstTree_Val() gopurs_runtime.Value {
	once_Test_AstTree_Val.Do(func() {
		cache_Test_AstTree_Val = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 245188743, UnsafePtr: unsafe.Pointer((&Constructor_Test_AstTree_Val{1, value0.IntVal}))}
		})
	})
	return cache_Test_AstTree_Val
}

var cache_Test_AstTree_Add gopurs_runtime.Value
var once_Test_AstTree_Add sync.Once

func Get_Test_AstTree_Add() gopurs_runtime.Value {
	once_Test_AstTree_Add.Do(func() {
		cache_Test_AstTree_Add = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(value1 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Value{Type: 9, IntVal: 2937956733, UnsafePtr: unsafe.Pointer((&Constructor_Test_AstTree_Add{1, value0, value1}))}
			})
		})
	})
	return cache_Test_AstTree_Add
}

var cache_Test_AstTree_Mul gopurs_runtime.Value
var once_Test_AstTree_Mul sync.Once

func Get_Test_AstTree_Mul() gopurs_runtime.Value {
	once_Test_AstTree_Mul.Do(func() {
		cache_Test_AstTree_Mul = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(value1 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Value{Type: 9, IntVal: 3406566728, UnsafePtr: unsafe.Pointer((&Constructor_Test_AstTree_Mul{1, value0, value1}))}
			})
		})
	})
	return cache_Test_AstTree_Mul
}

var cache_Test_AstTree_Sub gopurs_runtime.Value
var once_Test_AstTree_Sub sync.Once

func Get_Test_AstTree_Sub() gopurs_runtime.Value {
	once_Test_AstTree_Sub.Do(func() {
		cache_Test_AstTree_Sub = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(value1 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Value{Type: 9, IntVal: 2029887576, UnsafePtr: unsafe.Pointer((&Constructor_Test_AstTree_Sub{1, value0, value1}))}
			})
		})
	})
	return cache_Test_AstTree_Sub
}

var cache_Test_AstTree_eval gopurs_runtime.Value
var once_Test_AstTree_eval sync.Once

func Get_Test_AstTree_eval() gopurs_runtime.Value {
	once_Test_AstTree_eval.Do(func() {
		cache_Test_AstTree_eval = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Int(Call_Test_AstTree_eval(v_0_box))
		})
	})
	return cache_Test_AstTree_eval
}

var cache_Test_AstTree_describe gopurs_runtime.Value
var once_Test_AstTree_describe sync.Once

func Get_Test_AstTree_describe() gopurs_runtime.Value {
	once_Test_AstTree_describe.Do(func() {
		cache_Test_AstTree_describe = gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("AST Evaluation:"))
	})
	return cache_Test_AstTree_describe
}

var cache_Test_AstTree_buildTree gopurs_runtime.Value
var once_Test_AstTree_buildTree sync.Once

func Get_Test_AstTree_buildTree() gopurs_runtime.Value {
	once_Test_AstTree_buildTree.Do(func() {
		cache_Test_AstTree_buildTree = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Test_AstTree_buildTree(v_0_box.IntVal)
		})
	})
	return cache_Test_AstTree_buildTree
}

var cache_Test_AstTree_act gopurs_runtime.Value
var once_Test_AstTree_act sync.Once

func Get_Test_AstTree_act() gopurs_runtime.Value {
	once_Test_AstTree_act.Do(func() {
		cache_Test_AstTree_act = func() gopurs_runtime.Value {
			// TAST (Let): __local_var_0_0 -> gopurs_runtime.Value
			__local_var_0_0 := gopurs_runtime.Apply(Get_Bench_opaque(), gopurs_runtime.Int(3))
			_ = __local_var_0_0
			return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
				dummy_1_1 := gopurs_runtime.Apply(__local_var_0_0, gopurs_runtime.Value{})
				_ = dummy_1_1
				return gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str(gopurs_runtime.Apply(Get_Data_Show_showIntImpl(), gopurs_runtime.Int(Call_Test_AstTree_eval(Call_Test_AstTree_buildTree(dummy_1_1.IntVal)))).StrVal())), gopurs_runtime.Value{})
			})
		}()
	})
	return cache_Test_AstTree_act
}

type Constructor_Test_AstTree_Val struct {
	Rc uint32
	V0 int64
}

type Constructor_Test_AstTree_Add struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
}

type Constructor_Test_AstTree_Mul struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
}

type Constructor_Test_AstTree_Sub struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
}

func Call_Test_AstTree_logShow(a_0_loop int64) gopurs_runtime.Value {
	var a_0 int64 = a_0_loop
	_ = a_0
	return gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str(gopurs_runtime.Apply(Get_Data_Show_showIntImpl(), gopurs_runtime.Int(a_0)).StrVal()))
}

func Call_Test_AstTree_eval(v_0_loop gopurs_runtime.Value) int64 {
eval:
	for {
		if false {
			continue eval
		}
		var v_0 gopurs_runtime.Value = v_0_loop
		_ = v_0
		var __t0 int64
		{
			if v_0.Type == 9 && v_0.IntVal == 245188743 {
				__t0 = (*Constructor_Test_AstTree_Val)(v_0.UnsafePtr).V0
				goto end_branch_0
			} else {

			}
		}
		{
			if v_0.Type == 9 && v_0.IntVal == 2937956733 {
				__t0 = (Call_Test_AstTree_eval((*Constructor_Test_AstTree_Add)(v_0.UnsafePtr).V0)) + (Call_Test_AstTree_eval((*Constructor_Test_AstTree_Add)(v_0.UnsafePtr).V1))
				goto end_branch_0
			} else {

			}
		}
		{
			if v_0.Type == 9 && v_0.IntVal == 3406566728 {
				__t0 = (Call_Test_AstTree_eval((*Constructor_Test_AstTree_Mul)(v_0.UnsafePtr).V0)) * (Call_Test_AstTree_eval((*Constructor_Test_AstTree_Mul)(v_0.UnsafePtr).V1))
				goto end_branch_0
			} else {

			}
		}
		{
			if v_0.Type == 9 && v_0.IntVal == 2029887576 {
				__t0 = (Call_Test_AstTree_eval((*Constructor_Test_AstTree_Sub)(v_0.UnsafePtr).V0)) - (Call_Test_AstTree_eval((*Constructor_Test_AstTree_Sub)(v_0.UnsafePtr).V1))
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
}

func Call_Test_AstTree_buildTree(v_0_loop int64) gopurs_runtime.Value {
buildTree:
	for {
		if false {
			continue buildTree
		}
		var v_0 int64 = v_0_loop
		_ = v_0
		var __t0 gopurs_runtime.Value
		{
			if (v_0) == (0) {
				__t0 = gopurs_runtime.Value{Type: 9, IntVal: 245188743, UnsafePtr: unsafe.Pointer((&Constructor_Test_AstTree_Val{1, 1}))}
				goto end_branch_0
			} else {

			}
		}
		{
			__t0 = gopurs_runtime.Value{Type: 9, IntVal: 2937956733, UnsafePtr: unsafe.Pointer((&Constructor_Test_AstTree_Add{1, gopurs_runtime.Value{Type: 9, IntVal: 3406566728, UnsafePtr: unsafe.Pointer((&Constructor_Test_AstTree_Mul{1, gopurs_runtime.Value{Type: 9, IntVal: 245188743, UnsafePtr: unsafe.Pointer((&Constructor_Test_AstTree_Val{1, v_0}))}, Call_Test_AstTree_buildTree((v_0) - (1))}))}, gopurs_runtime.Value{Type: 9, IntVal: 2029887576, UnsafePtr: unsafe.Pointer((&Constructor_Test_AstTree_Sub{1, Call_Test_AstTree_buildTree((v_0) - (1)), gopurs_runtime.Value{Type: 9, IntVal: 245188743, UnsafePtr: unsafe.Pointer((&Constructor_Test_AstTree_Val{1, 1}))}}))}}))}
		}
	end_branch_0:
		return __t0
	}
}
