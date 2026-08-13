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
		cache_Data_Array_ST_Iterator_not = gopurs_runtime.RecordGet(Get_Data_HeytingAlgebra_heytingAlgebraBoolean(), "not")
	})
	return cache_Data_Array_ST_Iterator_not
}

var cache_Data_Array_ST_Iterator_void gopurs_runtime.Value
var once_Data_Array_ST_Iterator_void sync.Once
func Get_Data_Array_ST_Iterator_void() gopurs_runtime.Value {
	once_Data_Array_ST_Iterator_void.Do(func() {
		cache_Data_Array_ST_Iterator_void = gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Control_Monad_ST_Internal_functorST(), "map"), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Unit_unit()
}))
	})
	return cache_Data_Array_ST_Iterator_void
}

var cache_Data_Array_ST_Iterator_void1 gopurs_runtime.Value
var once_Data_Array_ST_Iterator_void1 sync.Once
func Get_Data_Array_ST_Iterator_void1() gopurs_runtime.Value {
	once_Data_Array_ST_Iterator_void1.Do(func() {
		cache_Data_Array_ST_Iterator_void1 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Control_Monad_ST_Internal_functorST(), "map"), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Unit_unit()
}))
	})
	return cache_Data_Array_ST_Iterator_void1
}

var cache_Data_Array_ST_Iterator_Iterator gopurs_runtime.Value
var once_Data_Array_ST_Iterator_Iterator sync.Once
func Get_Data_Array_ST_Iterator_Iterator() gopurs_runtime.Value {
	once_Data_Array_ST_Iterator_Iterator.Do(func() {
		cache_Data_Array_ST_Iterator_Iterator = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(value1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3127013252, UnsafePtr: unsafe.Pointer(&Constructor_Data_Array_ST_Iterator_Iterator{1, value0, value1})}
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
return Call_Data_Array_ST_Iterator_peek(gopurs_runtime.CoerceToStruct[Constructor_Data_Array_ST_Iterator_Iterator](v_0_box))
})
	})
	return cache_Data_Array_ST_Iterator_peek
}

var cache_Data_Array_ST_Iterator_next gopurs_runtime.Value
var once_Data_Array_ST_Iterator_next sync.Once
func Get_Data_Array_ST_Iterator_next() gopurs_runtime.Value {
	once_Data_Array_ST_Iterator_next.Do(func() {
		cache_Data_Array_ST_Iterator_next = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Array_ST_Iterator_next(gopurs_runtime.CoerceToStruct[Constructor_Data_Array_ST_Iterator_Iterator](v_0_box))
})
	})
	return cache_Data_Array_ST_Iterator_next
}

var cache_Data_Array_ST_Iterator_pushWhile gopurs_runtime.Value
var once_Data_Array_ST_Iterator_pushWhile sync.Once
func Get_Data_Array_ST_Iterator_pushWhile() gopurs_runtime.Value {
	once_Data_Array_ST_Iterator_pushWhile.Do(func() {
		cache_Data_Array_ST_Iterator_pushWhile = gopurs_runtime.Func3(func(p_0_box gopurs_runtime.Value, iter_1_box gopurs_runtime.Value, array_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Array_ST_Iterator_pushWhile(p_0_box, gopurs_runtime.CoerceToStruct[Constructor_Data_Array_ST_Iterator_Iterator](iter_1_box), array_2_box)
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
return Call_Data_Array_ST_Iterator_iterate(gopurs_runtime.CoerceToStruct[Constructor_Data_Array_ST_Iterator_Iterator](iter_0_box), f_1_box)
})
	})
	return cache_Data_Array_ST_Iterator_iterate
}

var cache_Data_Array_ST_Iterator_exhausted gopurs_runtime.Value
var once_Data_Array_ST_Iterator_exhausted sync.Once
func Get_Data_Array_ST_Iterator_exhausted() gopurs_runtime.Value {
	once_Data_Array_ST_Iterator_exhausted.Do(func() {
		cache_Data_Array_ST_Iterator_exhausted = func() gopurs_runtime.Value {
// TAST (Let): __local_var_0_0 -> gopurs_runtime.Value
__local_var_0_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Control_Monad_ST_Internal_functorST(), "map"), Get_Data_Maybe_isNothing())
_ = __local_var_0_0
return gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_0_0, Call_Data_Array_ST_Iterator_peek(gopurs_runtime.CoerceToStruct[Constructor_Data_Array_ST_Iterator_Iterator](x_1)))
})
}()
	})
	return cache_Data_Array_ST_Iterator_exhausted
}

var cache_Data_Array_ST_Iterator_iterate__1936300090 gopurs_runtime.Value
var once_Data_Array_ST_Iterator_iterate__1936300090 sync.Once
func Get_Data_Array_ST_Iterator_iterate__1936300090() gopurs_runtime.Value {
	once_Data_Array_ST_Iterator_iterate__1936300090.Do(func() {
		cache_Data_Array_ST_Iterator_iterate__1936300090 = gopurs_runtime.Func2(func(iter_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Array_ST_Iterator_iterate__1936300090(gopurs_runtime.CoerceToStruct[Constructor_Data_Array_ST_Iterator_Iterator](iter_0_box), f_1_box)
})
	})
	return cache_Data_Array_ST_Iterator_iterate__1936300090
}

var cache_Data_Array_ST_Iterator_iterate__1835948090 gopurs_runtime.Value
var once_Data_Array_ST_Iterator_iterate__1835948090 sync.Once
func Get_Data_Array_ST_Iterator_iterate__1835948090() gopurs_runtime.Value {
	once_Data_Array_ST_Iterator_iterate__1835948090.Do(func() {
		cache_Data_Array_ST_Iterator_iterate__1835948090 = gopurs_runtime.Func2(func(iter_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Array_ST_Iterator_iterate__1835948090(gopurs_runtime.CoerceToStruct[Constructor_Data_Array_ST_Iterator_Iterator](iter_0_box), f_1_box)
})
	})
	return cache_Data_Array_ST_Iterator_iterate__1835948090
}

var cache_Data_Array_ST_Iterator_iterator__1149050118 gopurs_runtime.Value
var once_Data_Array_ST_Iterator_iterator__1149050118 sync.Once
func Get_Data_Array_ST_Iterator_iterator__1149050118() gopurs_runtime.Value {
	once_Data_Array_ST_Iterator_iterator__1149050118.Do(func() {
		cache_Data_Array_ST_Iterator_iterator__1149050118 = gopurs_runtime.Func(func(f_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Array_ST_Iterator_iterator__1149050118(f_0_box)
})
	})
	return cache_Data_Array_ST_Iterator_iterator__1149050118
}

var cache_Data_Array_ST_Iterator_iterator__2947907462 gopurs_runtime.Value
var once_Data_Array_ST_Iterator_iterator__2947907462 sync.Once
func Get_Data_Array_ST_Iterator_iterator__2947907462() gopurs_runtime.Value {
	once_Data_Array_ST_Iterator_iterator__2947907462.Do(func() {
		cache_Data_Array_ST_Iterator_iterator__2947907462 = gopurs_runtime.Func(func(f_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Array_ST_Iterator_iterator__2947907462(f_0_box)
})
	})
	return cache_Data_Array_ST_Iterator_iterator__2947907462
}

var cache_Data_Array_ST_Iterator_next__2731492779 gopurs_runtime.Value
var once_Data_Array_ST_Iterator_next__2731492779 sync.Once
func Get_Data_Array_ST_Iterator_next__2731492779() gopurs_runtime.Value {
	once_Data_Array_ST_Iterator_next__2731492779.Do(func() {
		cache_Data_Array_ST_Iterator_next__2731492779 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Array_ST_Iterator_next__2731492779(gopurs_runtime.CoerceToStruct[Constructor_Data_Array_ST_Iterator_Iterator](v_0_box))
})
	})
	return cache_Data_Array_ST_Iterator_next__2731492779
}

var cache_Data_Array_ST_Iterator_next__1163996811 gopurs_runtime.Value
var once_Data_Array_ST_Iterator_next__1163996811 sync.Once
func Get_Data_Array_ST_Iterator_next__1163996811() gopurs_runtime.Value {
	once_Data_Array_ST_Iterator_next__1163996811.Do(func() {
		cache_Data_Array_ST_Iterator_next__1163996811 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Array_ST_Iterator_next__1163996811(gopurs_runtime.CoerceToStruct[Constructor_Data_Array_ST_Iterator_Iterator](v_0_box))
})
	})
	return cache_Data_Array_ST_Iterator_next__1163996811
}

var cache_Data_Array_ST_Iterator_peek__201669949 gopurs_runtime.Value
var once_Data_Array_ST_Iterator_peek__201669949 sync.Once
func Get_Data_Array_ST_Iterator_peek__201669949() gopurs_runtime.Value {
	once_Data_Array_ST_Iterator_peek__201669949.Do(func() {
		cache_Data_Array_ST_Iterator_peek__201669949 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Array_ST_Iterator_peek__201669949(gopurs_runtime.CoerceToStruct[Constructor_Data_Array_ST_Iterator_Iterator](v_0_box))
})
	})
	return cache_Data_Array_ST_Iterator_peek__201669949
}

var cache_Data_Array_ST_Iterator_peek__2731492779 gopurs_runtime.Value
var once_Data_Array_ST_Iterator_peek__2731492779 sync.Once
func Get_Data_Array_ST_Iterator_peek__2731492779() gopurs_runtime.Value {
	once_Data_Array_ST_Iterator_peek__2731492779.Do(func() {
		cache_Data_Array_ST_Iterator_peek__2731492779 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Array_ST_Iterator_peek__2731492779(gopurs_runtime.CoerceToStruct[Constructor_Data_Array_ST_Iterator_Iterator](v_0_box))
})
	})
	return cache_Data_Array_ST_Iterator_peek__2731492779
}

var cache_Data_Array_ST_Iterator_pushWhile__2298419255 gopurs_runtime.Value
var once_Data_Array_ST_Iterator_pushWhile__2298419255 sync.Once
func Get_Data_Array_ST_Iterator_pushWhile__2298419255() gopurs_runtime.Value {
	once_Data_Array_ST_Iterator_pushWhile__2298419255.Do(func() {
		cache_Data_Array_ST_Iterator_pushWhile__2298419255 = gopurs_runtime.Func3(func(p_0_box gopurs_runtime.Value, iter_1_box gopurs_runtime.Value, array_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Array_ST_Iterator_pushWhile__2298419255(p_0_box, gopurs_runtime.CoerceToStruct[Constructor_Data_Array_ST_Iterator_Iterator](iter_1_box), array_2_box)
})
	})
	return cache_Data_Array_ST_Iterator_pushWhile__2298419255
}

type Constructor_Data_Array_ST_Iterator_Iterator struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
}


func Call_Data_Array_ST_Iterator_peek(v_0_loop *Constructor_Data_Array_ST_Iterator_Iterator) gopurs_runtime.Value {
var v_0 *Constructor_Data_Array_ST_Iterator_Iterator = v_0_loop
_ = v_0
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Control_Monad_ST_Internal_bindST(), "bind"), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
return (*((v_0).V1.PtrVal().(*interface{}))).(gopurs_runtime.Value)
}), gopurs_runtime.Func(func(i_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Control_Monad_ST_Internal_applicativeST(), "pure"), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply((v_0).V0, gopurs_runtime.Int(i_1.IntVal))))})
}))
}

func Call_Data_Array_ST_Iterator_next(v_0_loop *Constructor_Data_Array_ST_Iterator_Iterator) gopurs_runtime.Value {
var v_0 *Constructor_Data_Array_ST_Iterator_Iterator = v_0_loop
_ = v_0
// TAST (Let): __local_var_1_0 -> gopurs_runtime.Value
__local_var_1_0 := (v_0).V1
_ = __local_var_1_0
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Control_Monad_ST_Internal_bindST(), "bind"), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
return (*(__local_var_1_0.PtrVal().(*interface{}))).(gopurs_runtime.Value)
}), gopurs_runtime.Func(func(i_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Control_Monad_ST_Internal_bindST(), "bind"), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_3_1 := (*(__local_var_1_0.PtrVal().(*interface{}))).(gopurs_runtime.Value)
_ = __local_var_3_1
*(__local_var_1_0.PtrVal().(*interface{})) = gopurs_runtime.Int((__local_var_3_1.IntVal) + (1))
return gopurs_runtime.Int((__local_var_3_1.IntVal) + (1))
}), gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Control_Monad_ST_Internal_applicativeST(), "pure"), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply((v_0).V0, gopurs_runtime.Int(i_2.IntVal))))})
}))
}))
}

func Call_Data_Array_ST_Iterator_pushWhile(p_0_loop gopurs_runtime.Value, iter_1_loop *Constructor_Data_Array_ST_Iterator_Iterator, array_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var p_0 gopurs_runtime.Value = p_0_loop
_ = p_0
var iter_1 *Constructor_Data_Array_ST_Iterator_Iterator = iter_1_loop
_ = iter_1
var array_2 gopurs_runtime.Value = array_2_loop
_ = array_2
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Control_Monad_ST_Internal_bindST(), "bind"), gopurs_runtime.Apply(Get_Control_Monad_ST_Internal_newImpl(), gopurs_runtime.Bool(false)), gopurs_runtime.Func(func(go__break_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Get_Control_Monad_ST_Internal_while(), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Control_Monad_ST_Internal_functorST(), "map"), gopurs_runtime.RecordGet(Get_Data_HeytingAlgebra_heytingAlgebraBoolean(), "not"), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
return (*(go__break_3.PtrVal().(*interface{}))).(gopurs_runtime.Value)
})), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Control_Monad_ST_Internal_bindST(), "bind"), Call_Data_Array_ST_Iterator_peek(iter_1), gopurs_runtime.Func(func(mx_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if ((mx_4.Type == 9 && mx_4.IntVal == 930809136 && mx_4.UnsafePtr != nil)) && ((gopurs_runtime.Apply(p_0, (*Constructor_Data_Maybe_Just)(mx_4.UnsafePtr).V0).IntVal) != (0)) {
__t0 = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Control_Monad_ST_Internal_bindST(), "bind"), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.UncurriedApp2(Get_Data_Array_ST_pushImpl(), (*Constructor_Data_Maybe_Just)(mx_4.UnsafePtr).V0, array_2)
}), gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_Data_Array_ST_Iterator_void(), Call_Data_Array_ST_Iterator_next(iter_1))
}))
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Apply(Get_Data_Array_ST_Iterator_void1(), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
*(go__break_3.PtrVal().(*interface{})) = gopurs_runtime.Bool(true)
return gopurs_runtime.Bool(true)
}))
}
end_branch_0:
return __t0
})))
}))
}

func Call_Data_Array_ST_Iterator_iterator(f_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Control_Monad_ST_Internal_functorST(), "map"), gopurs_runtime.Apply(Get_Data_Array_ST_Iterator_Iterator(), f_0), gopurs_runtime.Apply(Get_Control_Monad_ST_Internal_newImpl(), gopurs_runtime.Int(0)))
}

func Call_Data_Array_ST_Iterator_iterate(iter_0_loop *Constructor_Data_Array_ST_Iterator_Iterator, f_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var iter_0 *Constructor_Data_Array_ST_Iterator_Iterator = iter_0_loop
_ = iter_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Control_Monad_ST_Internal_bindST(), "bind"), gopurs_runtime.Apply(Get_Control_Monad_ST_Internal_newImpl(), gopurs_runtime.Bool(false)), gopurs_runtime.Func(func(go__break_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Get_Control_Monad_ST_Internal_while(), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Control_Monad_ST_Internal_functorST(), "map"), gopurs_runtime.RecordGet(Get_Data_HeytingAlgebra_heytingAlgebraBoolean(), "not"), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
return (*(go__break_2.PtrVal().(*interface{}))).(gopurs_runtime.Value)
})), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Control_Monad_ST_Internal_bindST(), "bind"), Call_Data_Array_ST_Iterator_next(iter_0), gopurs_runtime.Func(func(mx_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (mx_3.Type == 9 && mx_3.IntVal == 930809136 && mx_3.UnsafePtr != nil) {
__t0 = gopurs_runtime.Apply(f_1, (*Constructor_Data_Maybe_Just)(mx_3.UnsafePtr).V0)
goto end_branch_0
} else {

}
}
{
if (mx_3.Type == 9 && mx_3.IntVal == 930809136 && mx_3.UnsafePtr == nil) {
__t0 = gopurs_runtime.Apply(Get_Data_Array_ST_Iterator_void1(), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
*(go__break_2.PtrVal().(*interface{})) = gopurs_runtime.Bool(true)
return gopurs_runtime.Bool(true)
}))
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return __t0
})))
}))
}

func Call_Data_Array_ST_Iterator_iterate__1936300090(iter_0_loop *Constructor_Data_Array_ST_Iterator_Iterator, f_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var iter_0 *Constructor_Data_Array_ST_Iterator_Iterator = iter_0_loop
_ = iter_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Control_Monad_ST_Internal_bindST(), "bind"), gopurs_runtime.Apply(Get_Control_Monad_ST_Internal_newImpl(), gopurs_runtime.Bool(false)), gopurs_runtime.Func(func(go__break_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Get_Control_Monad_ST_Internal_while(), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Control_Monad_ST_Internal_functorST(), "map"), gopurs_runtime.RecordGet(Get_Data_HeytingAlgebra_heytingAlgebraBoolean(), "not"), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
return (*(go__break_2.PtrVal().(*interface{}))).(gopurs_runtime.Value)
})), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Control_Monad_ST_Internal_bindST(), "bind"), Call_Data_Array_ST_Iterator_next(iter_0), gopurs_runtime.Func(func(mx_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (mx_3.Type == 9 && mx_3.IntVal == 930809136 && mx_3.UnsafePtr != nil) {
__t0 = gopurs_runtime.Apply(f_1, (*Constructor_Data_Maybe_Just)(mx_3.UnsafePtr).V0)
goto end_branch_0
} else {

}
}
{
if (mx_3.Type == 9 && mx_3.IntVal == 930809136 && mx_3.UnsafePtr == nil) {
__t0 = gopurs_runtime.Apply(Get_Data_Array_ST_Iterator_void1(), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
*(go__break_2.PtrVal().(*interface{})) = gopurs_runtime.Bool(true)
return gopurs_runtime.Bool(true)
}))
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return __t0
})))
}))
}

func Call_Data_Array_ST_Iterator_iterate__1835948090(iter_0_loop *Constructor_Data_Array_ST_Iterator_Iterator, f_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var iter_0 *Constructor_Data_Array_ST_Iterator_Iterator = iter_0_loop
_ = iter_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Control_Monad_ST_Internal_bindST(), "bind"), gopurs_runtime.Apply(Get_Control_Monad_ST_Internal_newImpl(), gopurs_runtime.Bool(false)), gopurs_runtime.Func(func(go__break_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Get_Control_Monad_ST_Internal_while(), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Control_Monad_ST_Internal_functorST(), "map"), gopurs_runtime.RecordGet(Get_Data_HeytingAlgebra_heytingAlgebraBoolean(), "not"), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
return (*(go__break_2.PtrVal().(*interface{}))).(gopurs_runtime.Value)
})), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Control_Monad_ST_Internal_bindST(), "bind"), Call_Data_Array_ST_Iterator_next(gopurs_runtime.CoerceToStruct[Constructor_Data_Array_ST_Iterator_Iterator](gopurs_runtime.Value{Type: 9, IntVal: 3127013252, UnsafePtr: unsafe.Pointer(iter_0)})), gopurs_runtime.Func(func(mx_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (mx_3.Type == 9 && mx_3.IntVal == 930809136 && mx_3.UnsafePtr != nil) {
__t0 = gopurs_runtime.Apply(f_1, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple]((*Constructor_Data_Maybe_Just)(mx_3.UnsafePtr).V0))})
goto end_branch_0
} else {

}
}
{
if (mx_3.Type == 9 && mx_3.IntVal == 930809136 && mx_3.UnsafePtr == nil) {
__t0 = gopurs_runtime.Apply(Get_Data_Array_ST_Iterator_void1(), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
*(go__break_2.PtrVal().(*interface{})) = gopurs_runtime.Bool(true)
return gopurs_runtime.Bool(true)
}))
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return __t0
})))
}))
}

func Call_Data_Array_ST_Iterator_iterator__1149050118(f_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Control_Monad_ST_Internal_functorST(), "map"), gopurs_runtime.Apply(Get_Data_Array_ST_Iterator_Iterator(), f_0), gopurs_runtime.Apply(Get_Control_Monad_ST_Internal_newImpl(), gopurs_runtime.Int(0)))
}

func Call_Data_Array_ST_Iterator_iterator__2947907462(f_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Control_Monad_ST_Internal_functorST(), "map"), gopurs_runtime.Apply(Get_Data_Array_ST_Iterator_Iterator(), f_0), gopurs_runtime.Apply(Get_Control_Monad_ST_Internal_newImpl(), gopurs_runtime.Int(0)))
}

func Call_Data_Array_ST_Iterator_next__2731492779(v_0_loop *Constructor_Data_Array_ST_Iterator_Iterator) gopurs_runtime.Value {
var v_0 *Constructor_Data_Array_ST_Iterator_Iterator = v_0_loop
_ = v_0
// TAST (Let): __local_var_1_0 -> gopurs_runtime.Value
__local_var_1_0 := (v_0).V1
_ = __local_var_1_0
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Control_Monad_ST_Internal_bindST(), "bind"), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
return (*(__local_var_1_0.PtrVal().(*interface{}))).(gopurs_runtime.Value)
}), gopurs_runtime.Func(func(i_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Control_Monad_ST_Internal_bindST(), "bind"), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_3_1 := (*(__local_var_1_0.PtrVal().(*interface{}))).(gopurs_runtime.Value)
_ = __local_var_3_1
*(__local_var_1_0.PtrVal().(*interface{})) = gopurs_runtime.Int((__local_var_3_1.IntVal) + (1))
return gopurs_runtime.Int((__local_var_3_1.IntVal) + (1))
}), gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Control_Monad_ST_Internal_applicativeST(), "pure"), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply((v_0).V0, gopurs_runtime.Int(i_2.IntVal))))})
}))
}))
}

func Call_Data_Array_ST_Iterator_next__1163996811(v_0_loop *Constructor_Data_Array_ST_Iterator_Iterator) gopurs_runtime.Value {
var v_0 *Constructor_Data_Array_ST_Iterator_Iterator = v_0_loop
_ = v_0
// TAST (Let): __local_var_1_0 -> gopurs_runtime.Value
__local_var_1_0 := (v_0).V1
_ = __local_var_1_0
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Control_Monad_ST_Internal_bindST(), "bind"), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
return (*(__local_var_1_0.PtrVal().(*interface{}))).(gopurs_runtime.Value)
}), gopurs_runtime.Func(func(i_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Control_Monad_ST_Internal_bindST(), "bind"), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_3_1 := (*(__local_var_1_0.PtrVal().(*interface{}))).(gopurs_runtime.Value)
_ = __local_var_3_1
*(__local_var_1_0.PtrVal().(*interface{})) = gopurs_runtime.Int((__local_var_3_1.IntVal) + (1))
return gopurs_runtime.Int((__local_var_3_1.IntVal) + (1))
}), gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Control_Monad_ST_Internal_applicativeST(), "pure"), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply((v_0).V0, gopurs_runtime.Int(i_2.IntVal))))})
}))
}))
}

func Call_Data_Array_ST_Iterator_peek__201669949(v_0_loop *Constructor_Data_Array_ST_Iterator_Iterator) gopurs_runtime.Value {
var v_0 *Constructor_Data_Array_ST_Iterator_Iterator = v_0_loop
_ = v_0
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Control_Monad_ST_Internal_bindST(), "bind"), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
return (*((v_0).V1.PtrVal().(*interface{}))).(gopurs_runtime.Value)
}), gopurs_runtime.Func(func(i_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Control_Monad_ST_Internal_applicativeST(), "pure"), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply((v_0).V0, gopurs_runtime.Int(i_1.IntVal))))})
}))
}

func Call_Data_Array_ST_Iterator_peek__2731492779(v_0_loop *Constructor_Data_Array_ST_Iterator_Iterator) gopurs_runtime.Value {
var v_0 *Constructor_Data_Array_ST_Iterator_Iterator = v_0_loop
_ = v_0
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Control_Monad_ST_Internal_bindST(), "bind"), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
return (*((v_0).V1.PtrVal().(*interface{}))).(gopurs_runtime.Value)
}), gopurs_runtime.Func(func(i_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Control_Monad_ST_Internal_applicativeST(), "pure"), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply((v_0).V0, gopurs_runtime.Int(i_1.IntVal))))})
}))
}

func Call_Data_Array_ST_Iterator_pushWhile__2298419255(p_0_loop gopurs_runtime.Value, iter_1_loop *Constructor_Data_Array_ST_Iterator_Iterator, array_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var p_0 gopurs_runtime.Value = p_0_loop
_ = p_0
var iter_1 *Constructor_Data_Array_ST_Iterator_Iterator = iter_1_loop
_ = iter_1
var array_2 gopurs_runtime.Value = array_2_loop
_ = array_2
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Control_Monad_ST_Internal_bindST(), "bind"), gopurs_runtime.Apply(Get_Control_Monad_ST_Internal_newImpl(), gopurs_runtime.Bool(false)), gopurs_runtime.Func(func(go__break_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Get_Control_Monad_ST_Internal_while(), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Control_Monad_ST_Internal_functorST(), "map"), gopurs_runtime.RecordGet(Get_Data_HeytingAlgebra_heytingAlgebraBoolean(), "not"), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
return (*(go__break_3.PtrVal().(*interface{}))).(gopurs_runtime.Value)
})), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Control_Monad_ST_Internal_bindST(), "bind"), Call_Data_Array_ST_Iterator_peek(iter_1), gopurs_runtime.Func(func(mx_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if ((mx_4.Type == 9 && mx_4.IntVal == 930809136 && mx_4.UnsafePtr != nil)) && ((gopurs_runtime.Apply(p_0, (*Constructor_Data_Maybe_Just)(mx_4.UnsafePtr).V0).IntVal) != (0)) {
__t0 = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Control_Monad_ST_Internal_bindST(), "bind"), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.UncurriedApp2(Get_Data_Array_ST_pushImpl(), (*Constructor_Data_Maybe_Just)(mx_4.UnsafePtr).V0, array_2)
}), gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_Data_Array_ST_Iterator_void(), Call_Data_Array_ST_Iterator_next(iter_1))
}))
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Apply(Get_Data_Array_ST_Iterator_void1(), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
*(go__break_3.PtrVal().(*interface{})) = gopurs_runtime.Bool(true)
return gopurs_runtime.Bool(true)
}))
}
end_branch_0:
return __t0
})))
}))
}


