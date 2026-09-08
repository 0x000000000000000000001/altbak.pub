package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Data_Array_ST_Iterator_not gopurs_runtime.Value
var once_Data_Array_ST_Iterator_not sync.Once
func Get_Data_Array_ST_Iterator_not() gopurs_runtime.Value {
	once_Data_Array_ST_Iterator_not.Do(func() {
		cache_Data_Array_ST_Iterator_not = Get_Data_HeytingAlgebra_boolNot()
	})
	return cache_Data_Array_ST_Iterator_not
}

var cache_Data_Array_ST_Iterator_void gopurs_runtime.Value
var once_Data_Array_ST_Iterator_void sync.Once
func Get_Data_Array_ST_Iterator_void() gopurs_runtime.Value {
	once_Data_Array_ST_Iterator_void.Do(func() {
		cache_Data_Array_ST_Iterator_void = gopurs_runtime.Func(func(__local_var_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Array_ST_Iterator_void(__local_var_0_box)
})
	})
	return cache_Data_Array_ST_Iterator_void
}

var cache_Data_Array_ST_Iterator_void1 gopurs_runtime.Value
var once_Data_Array_ST_Iterator_void1 sync.Once
func Get_Data_Array_ST_Iterator_void1() gopurs_runtime.Value {
	once_Data_Array_ST_Iterator_void1.Do(func() {
		cache_Data_Array_ST_Iterator_void1 = gopurs_runtime.Func(func(__local_var_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Array_ST_Iterator_void1(__local_var_0_box)
})
	})
	return cache_Data_Array_ST_Iterator_void1
}

var cache_Data_Array_ST_Iterator_Iterator gopurs_runtime.Value
var once_Data_Array_ST_Iterator_Iterator sync.Once
func Get_Data_Array_ST_Iterator_Iterator() gopurs_runtime.Value {
	once_Data_Array_ST_Iterator_Iterator.Do(func() {
		cache_Data_Array_ST_Iterator_Iterator = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(value1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3127013252, UnsafePtr: unsafe.Pointer((&Constructor_Data_Array_ST_Iterator_Iterator[gopurs_runtime.Value, gopurs_runtime.Value]{1, value0, value1}))}
})
})
	})
	return cache_Data_Array_ST_Iterator_Iterator
}

var cache_Data_Array_ST_Iterator_peek gopurs_runtime.Value
var once_Data_Array_ST_Iterator_peek sync.Once
func Get_Data_Array_ST_Iterator_peek() gopurs_runtime.Value {
	once_Data_Array_ST_Iterator_peek.Do(func() {
		cache_Data_Array_ST_Iterator_peek = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Array_ST_Iterator_peek(gopurs_runtime.CoerceToStruct[Constructor_Data_Array_ST_Iterator_Iterator[gopurs_runtime.Value, gopurs_runtime.Value]](v_0_box))
})
	})
	return cache_Data_Array_ST_Iterator_peek
}

var cache_Data_Array_ST_Iterator_next gopurs_runtime.Value
var once_Data_Array_ST_Iterator_next sync.Once
func Get_Data_Array_ST_Iterator_next() gopurs_runtime.Value {
	once_Data_Array_ST_Iterator_next.Do(func() {
		cache_Data_Array_ST_Iterator_next = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Array_ST_Iterator_next(gopurs_runtime.CoerceToStruct[Constructor_Data_Array_ST_Iterator_Iterator[gopurs_runtime.Value, gopurs_runtime.Value]](v_0_box))
})
	})
	return cache_Data_Array_ST_Iterator_next
}

var cache_Data_Array_ST_Iterator_pushWhile gopurs_runtime.Value
var once_Data_Array_ST_Iterator_pushWhile sync.Once
func Get_Data_Array_ST_Iterator_pushWhile() gopurs_runtime.Value {
	once_Data_Array_ST_Iterator_pushWhile.Do(func() {
		cache_Data_Array_ST_Iterator_pushWhile = gopurs_runtime.Func3(func(p_0_box gopurs_runtime.Value, iter_1_box gopurs_runtime.Value, array_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Array_ST_Iterator_pushWhile(p_0_box, gopurs_runtime.CoerceToStruct[Constructor_Data_Array_ST_Iterator_Iterator[gopurs_runtime.Value, gopurs_runtime.Value]](iter_1_box), array_2_box)
})
	})
	return cache_Data_Array_ST_Iterator_pushWhile
}

var cache_Data_Array_ST_Iterator_pushAll gopurs_runtime.Value
var once_Data_Array_ST_Iterator_pushAll sync.Once
func Get_Data_Array_ST_Iterator_pushAll() gopurs_runtime.Value {
	once_Data_Array_ST_Iterator_pushAll.Do(func() {
		cache_Data_Array_ST_Iterator_pushAll = gopurs_runtime.Apply(Get_Data_Array_ST_Iterator_pushWhile(), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(true)
}))
	})
	return cache_Data_Array_ST_Iterator_pushAll
}

var cache_Data_Array_ST_Iterator_iterator gopurs_runtime.Value
var once_Data_Array_ST_Iterator_iterator sync.Once
func Get_Data_Array_ST_Iterator_iterator() gopurs_runtime.Value {
	once_Data_Array_ST_Iterator_iterator.Do(func() {
		cache_Data_Array_ST_Iterator_iterator = gopurs_runtime.Func(func(f_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Array_ST_Iterator_iterator(f_0_box)
})
	})
	return cache_Data_Array_ST_Iterator_iterator
}

var cache_Data_Array_ST_Iterator_iterate gopurs_runtime.Value
var once_Data_Array_ST_Iterator_iterate sync.Once
func Get_Data_Array_ST_Iterator_iterate() gopurs_runtime.Value {
	once_Data_Array_ST_Iterator_iterate.Do(func() {
		cache_Data_Array_ST_Iterator_iterate = gopurs_runtime.Func2(func(iter_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Array_ST_Iterator_iterate(gopurs_runtime.CoerceToStruct[Constructor_Data_Array_ST_Iterator_Iterator[gopurs_runtime.Value, gopurs_runtime.Value]](iter_0_box), f_1_box)
})
	})
	return cache_Data_Array_ST_Iterator_iterate
}

var cache_Data_Array_ST_Iterator_exhausted gopurs_runtime.Value
var once_Data_Array_ST_Iterator_exhausted sync.Once
func Get_Data_Array_ST_Iterator_exhausted() gopurs_runtime.Value {
	once_Data_Array_ST_Iterator_exhausted.Do(func() {
		cache_Data_Array_ST_Iterator_exhausted = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Array_ST_Iterator_exhausted(gopurs_runtime.CoerceToStruct[Constructor_Data_Array_ST_Iterator_Iterator[gopurs_runtime.Value, gopurs_runtime.Value]](x_0_box))
})
	})
	return cache_Data_Array_ST_Iterator_exhausted
}

type Constructor_Data_Array_ST_Iterator_Iterator[T_r any, T_a any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
}


func Call_Data_Array_ST_Iterator_void(__local_var_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __local_var_0 gopurs_runtime.Value = __local_var_0_loop
_ = __local_var_0
return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_1_0 := gopurs_runtime.Apply(__local_var_0, gopurs_runtime.Value{})
_ = __local_var_1_0
return Get_Data_Unit_unit()
})
}

func Call_Data_Array_ST_Iterator_void1(__local_var_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __local_var_0 gopurs_runtime.Value = __local_var_0_loop
_ = __local_var_0
return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_1_0 := gopurs_runtime.Apply(__local_var_0, gopurs_runtime.Value{})
_ = __local_var_1_0
return Get_Data_Unit_unit()
})
}

func Call_Data_Array_ST_Iterator_peek(v_0_loop *Constructor_Data_Array_ST_Iterator_Iterator[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var v_0 *Constructor_Data_Array_ST_Iterator_Iterator[gopurs_runtime.Value, gopurs_runtime.Value] = v_0_loop
_ = v_0
return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_1_0 shape=Other bindingType=Any
__local_var_1_0 := (v_0).V1
_ = __local_var_1_0
i_2_1 := (*(__local_var_1_0.PtrVal().(*interface{}))).(gopurs_runtime.Value)
_ = i_2_1
return gopurs_runtime.Apply((v_0).V0, gopurs_runtime.Int(i_2_1.IntVal))
})
}

func Call_Data_Array_ST_Iterator_next(v_0_loop *Constructor_Data_Array_ST_Iterator_Iterator[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var v_0 *Constructor_Data_Array_ST_Iterator_Iterator[gopurs_runtime.Value, gopurs_runtime.Value] = v_0_loop
_ = v_0
return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_1_0 shape=Other bindingType=Any
__local_var_1_0 := (v_0).V1
_ = __local_var_1_0
i_2_1 := (*(__local_var_1_0.PtrVal().(*interface{}))).(gopurs_runtime.Value)
_ = i_2_1
_dollar___unused_3_2 := gopurs_runtime.Apply(gopurs_runtime.Apply2(Get_Control_Monad_ST_Internal_modifyImpl(), gopurs_runtime.Func(func(s_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): s_prime__4_3 shape=Other bindingType=(TypeVar a)
s_prime__4_3 := gopurs_runtime.Int((s_3.IntVal) + (int64(1)))
_ = s_prime__4_3
return func() gopurs_runtime.Value {
				orig := struct{
	state gopurs_runtime.Value
	value gopurs_runtime.Value
}{s_prime__4_3, s_prime__4_3}
				_ = orig
				return gopurs_runtime.RecordDict([]string{"state", "value"}, []gopurs_runtime.Value{orig.state, orig.value})
				}()
}), __local_var_1_0), gopurs_runtime.Value{})
_ = _dollar___unused_3_2
return gopurs_runtime.Apply((v_0).V0, gopurs_runtime.Int(i_2_1.IntVal))
})
}

func Call_Data_Array_ST_Iterator_pushWhile(p_0_loop gopurs_runtime.Value, iter_1_loop *Constructor_Data_Array_ST_Iterator_Iterator[gopurs_runtime.Value, gopurs_runtime.Value], array_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var p_0 gopurs_runtime.Value = p_0_loop
_ = p_0
var iter_1 *Constructor_Data_Array_ST_Iterator_Iterator[gopurs_runtime.Value, gopurs_runtime.Value] = iter_1_loop
_ = iter_1
var array_2 gopurs_runtime.Value = array_2_loop
_ = array_2
return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_3_0 shape=App(Var) bindingType=Any
__local_var_3_0 := gopurs_runtime.Apply(Get_Control_Monad_ST_Internal_newImpl(), gopurs_runtime.Bool(false))
_ = __local_var_3_0
go__break_4_1 := gopurs_runtime.Apply(__local_var_3_0, gopurs_runtime.Value{})
_ = go__break_4_1
return gopurs_runtime.Apply(gopurs_runtime.Apply2(Get_Control_Monad_ST_Internal_while(), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_5_2 := (*(go__break_4_1.PtrVal().(*interface{}))).(gopurs_runtime.Value)
_ = __local_var_5_2
return gopurs_runtime.Bool(((__local_var_5_2.IntVal) != (0)) != (true))
}), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_5_3 shape=Other bindingType=Any
__local_var_5_3 := (iter_1).V1
_ = __local_var_5_3
i_6_5 := (*(__local_var_5_3.PtrVal().(*interface{}))).(gopurs_runtime.Value)
_ = i_6_5
mx_6_4 := gopurs_runtime.Apply((iter_1).V0, gopurs_runtime.Int(i_6_5.IntVal))
_ = mx_6_4
var __t10 gopurs_runtime.Value
{
var __t_tag_7 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](mx_6_4)
if ((__t_tag_7 != nil)) && ((gopurs_runtime.Apply(p_0, (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(mx_6_4.UnsafePtr).V0).IntVal) != (0)) {
__t10 = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
_dollar___unused_7_8 := gopurs_runtime.Apply(gopurs_runtime.Apply2(Get_Data_Array_ST_push(), (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(mx_6_4.UnsafePtr).V0, array_2), gopurs_runtime.Value{})
_ = _dollar___unused_7_8
__local_var_8_9 := gopurs_runtime.Apply(Call_Data_Array_ST_Iterator_next(iter_1), gopurs_runtime.Value{})
_ = __local_var_8_9
return Get_Data_Unit_unit()
})
goto end_branch_10
} else {

}
}
{
__t10 = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
*(go__break_4_1.PtrVal().(*interface{})) = gopurs_runtime.Bool(true)
__local_var_7_6 := gopurs_runtime.Bool(true)
_ = __local_var_7_6
return Get_Data_Unit_unit()
})
}
end_branch_10:
return gopurs_runtime.Apply(__t10, gopurs_runtime.Value{})
})), gopurs_runtime.Value{})
})
}

func Call_Data_Array_ST_Iterator_iterator(f_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_1_0 shape=App(Var) bindingType=Any
__local_var_1_0 := gopurs_runtime.Apply(Get_Data_Array_ST_Iterator_Iterator(), f_0)
_ = __local_var_1_0
// TAST (Let): __local_var_2_1 shape=App(Var) bindingType=(ADT ["Control","Monad","ST","Internal","ST"] [(TypeVar r), (ADT ["Control","Monad","ST","Internal","STRef"] [(TypeVar r), Int])])
__local_var_2_1 := gopurs_runtime.Apply(Get_Control_Monad_ST_Internal_newImpl(), gopurs_runtime.Int(int64(0)))
_ = __local_var_2_1
__local_var_3_2 := gopurs_runtime.Apply(__local_var_2_1, gopurs_runtime.Value{})
_ = __local_var_3_2
return gopurs_runtime.Apply(__local_var_1_0, __local_var_3_2)
})
}

func Call_Data_Array_ST_Iterator_iterate(iter_0_loop *Constructor_Data_Array_ST_Iterator_Iterator[gopurs_runtime.Value, gopurs_runtime.Value], f_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var iter_0 *Constructor_Data_Array_ST_Iterator_Iterator[gopurs_runtime.Value, gopurs_runtime.Value] = iter_0_loop
_ = iter_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_2_0 shape=App(Var) bindingType=Any
__local_var_2_0 := gopurs_runtime.Apply(Get_Control_Monad_ST_Internal_newImpl(), gopurs_runtime.Bool(false))
_ = __local_var_2_0
go__break_3_1 := gopurs_runtime.Apply(__local_var_2_0, gopurs_runtime.Value{})
_ = go__break_3_1
return gopurs_runtime.Apply(gopurs_runtime.Apply2(Get_Control_Monad_ST_Internal_while(), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_4_2 := (*(go__break_3_1.PtrVal().(*interface{}))).(gopurs_runtime.Value)
_ = __local_var_4_2
return gopurs_runtime.Bool(((__local_var_4_2.IntVal) != (0)) != (true))
}), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_4_3 shape=App(Var) bindingType=Any
__local_var_4_3 := Call_Data_Array_ST_Iterator_next(iter_0)
_ = __local_var_4_3
mx_5_4 := gopurs_runtime.Apply(__local_var_4_3, gopurs_runtime.Value{})
_ = mx_5_4
var __t8 gopurs_runtime.Value
{
var __t_tag_5 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](mx_5_4)
if (__t_tag_5 != nil) {
__t8 = gopurs_runtime.Apply(f_1, (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(mx_5_4.UnsafePtr).V0)
goto end_branch_8
} else {

}
}
{
var __t_tag_6 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](mx_5_4)
if (__t_tag_6 == nil) {
__t8 = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
*(go__break_3_1.PtrVal().(*interface{})) = gopurs_runtime.Bool(true)
__local_var_6_7 := gopurs_runtime.Bool(true)
_ = __local_var_6_7
return Get_Data_Unit_unit()
})
goto end_branch_8
} else {

}
}
{
__t8 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_8:
return gopurs_runtime.Apply(__t8, gopurs_runtime.Value{})
})), gopurs_runtime.Value{})
})
}

func Call_Data_Array_ST_Iterator_exhausted(x_0_loop *Constructor_Data_Array_ST_Iterator_Iterator[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var x_0 *Constructor_Data_Array_ST_Iterator_Iterator[gopurs_runtime.Value, gopurs_runtime.Value] = x_0_loop
_ = x_0
return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_1_0 shape=Let(EffectBind(EffectPure)) bindingType=(TypeVar c)
__local_var_1_0 := gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_1_1 shape=Other bindingType=Any
__local_var_1_1 := (x_0).V1
_ = __local_var_1_1
i_2_2 := (*(__local_var_1_1.PtrVal().(*interface{}))).(gopurs_runtime.Value)
_ = i_2_2
return gopurs_runtime.Apply((x_0).V0, gopurs_runtime.Int(i_2_2.IntVal))
})
_ = __local_var_1_0
__local_var_2_3 := gopurs_runtime.Apply(__local_var_1_0, gopurs_runtime.Value{})
_ = __local_var_2_3
var __t6 gopurs_runtime.Value
{
var __t_tag_4 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](__local_var_2_3)
if (__t_tag_4 == nil) {
__t6 = gopurs_runtime.Bool(true)
goto end_branch_6
} else {

}
}
{
var __t_tag_5 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](__local_var_2_3)
if (__t_tag_5 != nil) {
__t6 = gopurs_runtime.Bool(false)
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
}


