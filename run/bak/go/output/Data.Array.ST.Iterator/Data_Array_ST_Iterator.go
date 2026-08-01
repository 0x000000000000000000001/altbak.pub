package Data_Array_ST_Iterator

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Control_Monad_ST_Internal "gopurs/output/Control.Monad.ST.Internal"
	pkg_Data_Unit "gopurs/output/Data.Unit"
	pkg_Data_Maybe "gopurs/output/Data.Maybe"
	pkg_Data_HeytingAlgebra "gopurs/output/Data.HeytingAlgebra"
	pkg_Data_Array_ST "gopurs/output/Data.Array.ST"
	unsafe "unsafe"
)

var cache_void gopurs_runtime.Value
var once_void sync.Once
func Get_void() gopurs_runtime.Value {
	once_void.Do(func() {
		cache_void = gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Control_Monad_ST_Internal.Get_functorST(), "map"), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Unit.Get_unit()
}))
	})
	return cache_void
}

var cache_Iterator gopurs_runtime.Value
var once_Iterator sync.Once
func Get_Iterator() gopurs_runtime.Value {
	once_Iterator.Do(func() {
		cache_Iterator = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(value1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3127013252, UnsafePtr: unsafe.Pointer(&Constructor_Iterator[gopurs_runtime.Value, gopurs_runtime.Value]{1, value0, value1})}
})
})
	})
	return cache_Iterator
}

var cache_peek gopurs_runtime.Value
var once_peek sync.Once
func Get_peek() gopurs_runtime.Value {
	once_peek.Do(func() {
		cache_peek = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_peek(v_0_box)
})
	})
	return cache_peek
}

var cache_next gopurs_runtime.Value
var once_next sync.Once
func Get_next() gopurs_runtime.Value {
	once_next.Do(func() {
		cache_next = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_next(v_0_box)
})
	})
	return cache_next
}

var cache_pushWhile gopurs_runtime.Value
var once_pushWhile sync.Once
func Get_pushWhile() gopurs_runtime.Value {
	once_pushWhile.Do(func() {
		cache_pushWhile = gopurs_runtime.Func3(func(p_0_box gopurs_runtime.Value, iter_1_box gopurs_runtime.Value, array_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_pushWhile(p_0_box, iter_1_box, array_2_box)
})
	})
	return cache_pushWhile
}

var cache_pushWhile__gopurs_runtime_Value_1584032505 gopurs_runtime.Value
var once_pushWhile__gopurs_runtime_Value_1584032505 sync.Once
func Get_pushWhile__gopurs_runtime_Value_1584032505() gopurs_runtime.Value {
	once_pushWhile__gopurs_runtime_Value_1584032505.Do(func() {
		cache_pushWhile__gopurs_runtime_Value_1584032505 = gopurs_runtime.Func3(func(p_0_box gopurs_runtime.Value, iter_1_box gopurs_runtime.Value, array_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_pushWhile__gopurs_runtime_Value_1584032505(p_0_box, iter_1_box, array_2_box)
})
	})
	return cache_pushWhile__gopurs_runtime_Value_1584032505
}

var cache_pushAll gopurs_runtime.Value
var once_pushAll sync.Once
func Get_pushAll() gopurs_runtime.Value {
	once_pushAll.Do(func() {
		cache_pushAll = gopurs_runtime.Apply(Get_pushWhile(), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(true)
}))
	})
	return cache_pushAll
}

var cache_iterator gopurs_runtime.Value
var once_iterator sync.Once
func Get_iterator() gopurs_runtime.Value {
	once_iterator.Do(func() {
		cache_iterator = gopurs_runtime.Func(func(f_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_iterator(f_0_box)
})
	})
	return cache_iterator
}

var cache_iterate gopurs_runtime.Value
var once_iterate sync.Once
func Get_iterate() gopurs_runtime.Value {
	once_iterate.Do(func() {
		cache_iterate = gopurs_runtime.Func2(func(iter_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_iterate(iter_0_box, f_1_box)
})
	})
	return cache_iterate
}

var cache_iterate__gopurs_runtime_Value_408496612 gopurs_runtime.Value
var once_iterate__gopurs_runtime_Value_408496612 sync.Once
func Get_iterate__gopurs_runtime_Value_408496612() gopurs_runtime.Value {
	once_iterate__gopurs_runtime_Value_408496612.Do(func() {
		cache_iterate__gopurs_runtime_Value_408496612 = gopurs_runtime.Func2(func(iter_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_iterate__gopurs_runtime_Value_408496612(iter_0_box, f_1_box)
})
	})
	return cache_iterate__gopurs_runtime_Value_408496612
}

var cache_iterate__gopurs_runtime_Value_894948428 gopurs_runtime.Value
var once_iterate__gopurs_runtime_Value_894948428 sync.Once
func Get_iterate__gopurs_runtime_Value_894948428() gopurs_runtime.Value {
	once_iterate__gopurs_runtime_Value_894948428.Do(func() {
		cache_iterate__gopurs_runtime_Value_894948428 = gopurs_runtime.Func2(func(iter_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_iterate__gopurs_runtime_Value_894948428(iter_0_box, f_1_box)
})
	})
	return cache_iterate__gopurs_runtime_Value_894948428
}

var cache_exhausted gopurs_runtime.Value
var once_exhausted sync.Once
func Get_exhausted() gopurs_runtime.Value {
	once_exhausted.Do(func() {
		cache_exhausted = func() gopurs_runtime.Value {
__local_var_0_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Control_Monad_ST_Internal.Get_functorST(), "map"), pkg_Data_Maybe.Get_isNothing())
_ = __local_var_0_0
return gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_0_0, Call_peek(x_1))
})
}()
	})
	return cache_exhausted
}

type Constructor_Iterator[T_r any, T_a any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
}


func Call_peek(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
__local_var_1_0 := (*Constructor_Iterator[gopurs_runtime.Value, gopurs_runtime.Value])(v_0.UnsafePtr).V1
_ = __local_var_1_0
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Control_Monad_ST_Internal.Get_bindST(), "bind"), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
return (*(__local_var_1_0.PtrVal().(*interface{}))).(gopurs_runtime.Value)
}), gopurs_runtime.Func(func(i_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Control_Monad_ST_Internal.Get_applicativeST(), "pure"), gopurs_runtime.Apply((*Constructor_Iterator[gopurs_runtime.Value, gopurs_runtime.Value])(v_0.UnsafePtr).V0, i_2))
}))
}

func Call_next(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
__local_var_1_0 := (*Constructor_Iterator[gopurs_runtime.Value, gopurs_runtime.Value])(v_0.UnsafePtr).V1
_ = __local_var_1_0
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Control_Monad_ST_Internal.Get_bindST(), "bind"), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
return (*(__local_var_1_0.PtrVal().(*interface{}))).(gopurs_runtime.Value)
}), gopurs_runtime.Func(func(i_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Control_Monad_ST_Internal.Get_bindST(), "bind"), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_3_1 := (*(__local_var_1_0.PtrVal().(*interface{}))).(gopurs_runtime.Value)
_ = __local_var_3_1
*(__local_var_1_0.PtrVal().(*interface{})) = gopurs_runtime.Int((__local_var_3_1.IntVal) + (1))
return gopurs_runtime.Int((__local_var_3_1.IntVal) + (1))
}), gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Control_Monad_ST_Internal.Get_applicativeST(), "pure"), gopurs_runtime.Apply((*Constructor_Iterator[gopurs_runtime.Value, gopurs_runtime.Value])(v_0.UnsafePtr).V0, i_2))
}))
}))
}

func Call_pushWhile(p_0_loop gopurs_runtime.Value, iter_1_loop gopurs_runtime.Value, array_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var p_0 gopurs_runtime.Value = p_0_loop
_ = p_0
var iter_1 gopurs_runtime.Value = iter_1_loop
_ = iter_1
var array_2 gopurs_runtime.Value = array_2_loop
_ = array_2
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Control_Monad_ST_Internal.Get_bindST(), "bind"), gopurs_runtime.Apply(pkg_Control_Monad_ST_Internal.Get_newImpl(), gopurs_runtime.Bool(false)), gopurs_runtime.Func(func(go__break_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(pkg_Control_Monad_ST_Internal.Get_while(), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Control_Monad_ST_Internal.Get_functorST(), "map"), gopurs_runtime.RecordGet(pkg_Data_HeytingAlgebra.Get_heytingAlgebraBoolean(), "not"), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
return (*(go__break_3.PtrVal().(*interface{}))).(gopurs_runtime.Value)
})), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Control_Monad_ST_Internal.Get_bindST(), "bind"), Call_peek(iter_1), gopurs_runtime.Func(func(mx_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if ((mx_4.Type == 9 && mx_4.IntVal == 930809136 && mx_4.UnsafePtr != nil)) && ((gopurs_runtime.Apply(p_0, (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(mx_4.UnsafePtr).V0).IntVal) != (0)) {
__local_var_5_1 := (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(mx_4.UnsafePtr).V0
_ = __local_var_5_1
__t0 = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Control_Monad_ST_Internal.Get_bindST(), "bind"), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.UncurriedApp2(pkg_Data_Array_ST.Get_pushImpl(), __local_var_5_1, array_2)
}), gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_void(), Call_next(iter_1))
}))
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Apply(Get_void(), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
*(go__break_3.PtrVal().(*interface{})) = gopurs_runtime.Bool(true)
return gopurs_runtime.Bool(true)
}))
}
end_branch_0:
return __t0
})))
}))
}

func Call_pushWhile__gopurs_runtime_Value_1584032505(p_0_loop gopurs_runtime.Value, iter_1_loop gopurs_runtime.Value, array_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var p_0 gopurs_runtime.Value = p_0_loop
_ = p_0
var iter_1 gopurs_runtime.Value = iter_1_loop
_ = iter_1
var array_2 gopurs_runtime.Value = array_2_loop
_ = array_2
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Control_Monad_ST_Internal.Get_bindST(), "bind"), gopurs_runtime.Apply(pkg_Control_Monad_ST_Internal.Get_newImpl(), gopurs_runtime.Bool(false)), gopurs_runtime.Func(func(go__break_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(pkg_Control_Monad_ST_Internal.Get_while(), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Control_Monad_ST_Internal.Get_functorST(), "map"), gopurs_runtime.RecordGet(pkg_Data_HeytingAlgebra.Get_heytingAlgebraBoolean(), "not"), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
return (*(go__break_3.PtrVal().(*interface{}))).(gopurs_runtime.Value)
})), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Control_Monad_ST_Internal.Get_bindST(), "bind"), Call_peek(iter_1), gopurs_runtime.Func(func(mx_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if ((mx_4.Type == 9 && mx_4.IntVal == 930809136 && mx_4.UnsafePtr != nil)) && ((gopurs_runtime.Apply(p_0, (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(mx_4.UnsafePtr).V0).IntVal) != (0)) {
__local_var_5_1 := (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(mx_4.UnsafePtr).V0
_ = __local_var_5_1
__t0 = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Control_Monad_ST_Internal.Get_bindST(), "bind"), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.UncurriedApp2(pkg_Data_Array_ST.Get_pushImpl(), __local_var_5_1, array_2)
}), gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_void(), Call_next(iter_1))
}))
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Apply(Get_void(), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
*(go__break_3.PtrVal().(*interface{})) = gopurs_runtime.Bool(true)
return gopurs_runtime.Bool(true)
}))
}
end_branch_0:
return __t0
})))
}))
}

func Call_iterator(f_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Control_Monad_ST_Internal.Get_functorST(), "map"), gopurs_runtime.Apply(Get_Iterator(), f_0), gopurs_runtime.Apply(pkg_Control_Monad_ST_Internal.Get_newImpl(), gopurs_runtime.Int(0)))
}

func Call_iterate(iter_0_loop gopurs_runtime.Value, f_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var iter_0 gopurs_runtime.Value = iter_0_loop
_ = iter_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Control_Monad_ST_Internal.Get_bindST(), "bind"), gopurs_runtime.Apply(pkg_Control_Monad_ST_Internal.Get_newImpl(), gopurs_runtime.Bool(false)), gopurs_runtime.Func(func(go__break_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(pkg_Control_Monad_ST_Internal.Get_while(), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Control_Monad_ST_Internal.Get_functorST(), "map"), gopurs_runtime.RecordGet(pkg_Data_HeytingAlgebra.Get_heytingAlgebraBoolean(), "not"), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
return (*(go__break_2.PtrVal().(*interface{}))).(gopurs_runtime.Value)
})), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Control_Monad_ST_Internal.Get_bindST(), "bind"), Call_next(iter_0), gopurs_runtime.Func(func(mx_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (mx_3.Type == 9 && mx_3.IntVal == 930809136 && mx_3.UnsafePtr != nil) {
__t0 = gopurs_runtime.Apply(f_1, (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(mx_3.UnsafePtr).V0)
goto end_branch_0
} else {

}
}
{
if (mx_3.Type == 9 && mx_3.IntVal == 930809136 && mx_3.UnsafePtr == nil) {
__t0 = gopurs_runtime.Apply(Get_void(), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
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

func Call_iterate__gopurs_runtime_Value_408496612(iter_0_loop gopurs_runtime.Value, f_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var iter_0 gopurs_runtime.Value = iter_0_loop
_ = iter_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Control_Monad_ST_Internal.Get_bindST(), "bind"), gopurs_runtime.Apply(pkg_Control_Monad_ST_Internal.Get_newImpl(), gopurs_runtime.Bool(false)), gopurs_runtime.Func(func(go__break_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(pkg_Control_Monad_ST_Internal.Get_while(), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Control_Monad_ST_Internal.Get_functorST(), "map"), gopurs_runtime.RecordGet(pkg_Data_HeytingAlgebra.Get_heytingAlgebraBoolean(), "not"), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
return (*(go__break_2.PtrVal().(*interface{}))).(gopurs_runtime.Value)
})), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Control_Monad_ST_Internal.Get_bindST(), "bind"), Call_next(iter_0), gopurs_runtime.Func(func(mx_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (mx_3.Type == 9 && mx_3.IntVal == 930809136 && mx_3.UnsafePtr != nil) {
__t0 = gopurs_runtime.Apply(f_1, (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(mx_3.UnsafePtr).V0)
goto end_branch_0
} else {

}
}
{
if (mx_3.Type == 9 && mx_3.IntVal == 930809136 && mx_3.UnsafePtr == nil) {
__t0 = gopurs_runtime.Apply(Get_void(), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
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

func Call_iterate__gopurs_runtime_Value_894948428(iter_0_loop gopurs_runtime.Value, f_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var iter_0 gopurs_runtime.Value = iter_0_loop
_ = iter_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Control_Monad_ST_Internal.Get_bindST(), "bind"), gopurs_runtime.Apply(pkg_Control_Monad_ST_Internal.Get_newImpl(), gopurs_runtime.Bool(false)), gopurs_runtime.Func(func(go__break_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(pkg_Control_Monad_ST_Internal.Get_while(), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Control_Monad_ST_Internal.Get_functorST(), "map"), gopurs_runtime.RecordGet(pkg_Data_HeytingAlgebra.Get_heytingAlgebraBoolean(), "not"), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
return (*(go__break_2.PtrVal().(*interface{}))).(gopurs_runtime.Value)
})), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Control_Monad_ST_Internal.Get_bindST(), "bind"), Call_next(iter_0), gopurs_runtime.Func(func(mx_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (mx_3.Type == 9 && mx_3.IntVal == 930809136 && mx_3.UnsafePtr != nil) {
__t0 = gopurs_runtime.Apply(f_1, (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(mx_3.UnsafePtr).V0)
goto end_branch_0
} else {

}
}
{
if (mx_3.Type == 9 && mx_3.IntVal == 930809136 && mx_3.UnsafePtr == nil) {
__t0 = gopurs_runtime.Apply(Get_void(), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
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


