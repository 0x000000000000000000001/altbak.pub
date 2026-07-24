package Data_Array_ST_Iterator

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Control_Monad_ST_Internal "gopurs/output/Control.Monad.ST.Internal"
	pkg_Data_Array_ST "gopurs/output/Data.Array.ST"
	pkg_Data_Unit "gopurs/output/Data.Unit"
)

var Iterator gopurs_runtime.Value
var once_Iterator sync.Once
func Get_Iterator() gopurs_runtime.Value {
	once_Iterator.Do(func() {
		Iterator = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(value1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Constructor2("Iterator", value0, value1)
})
})
	})
	return Iterator
}

var peek gopurs_runtime.Value
var once_peek sync.Once
func Get_peek() gopurs_runtime.Value {
	once_peek.Do(func() {
		peek = gopurs_runtime.Func(func(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
__local_var_1_0 := (*[1024]gopurs_runtime.Value)(v_0.UnsafePtr)[1]
_ = __local_var_1_0
return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
i_2_1 := *(__local_var_1_0.PtrVal.(*gopurs_runtime.Value))
_ = i_2_1
return gopurs_runtime.Apply((*[1024]gopurs_runtime.Value)(v_0.UnsafePtr)[0], i_2_1)
})
}()
})
	})
	return peek
}

var next gopurs_runtime.Value
var once_next sync.Once
func Get_next() gopurs_runtime.Value {
	once_next.Do(func() {
		next = gopurs_runtime.Func(func(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
__local_var_1_0 := (*[1024]gopurs_runtime.Value)(v_0.UnsafePtr)[1]
_ = __local_var_1_0
return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
i_2_1 := *(__local_var_1_0.PtrVal.(*gopurs_runtime.Value))
_ = i_2_1
_dollar__unused_3_2 := gopurs_runtime.Apply(gopurs_runtime.Apply2(pkg_Control_Monad_ST_Internal.Get_modifyImpl(), gopurs_runtime.Func(func(s_3 gopurs_runtime.Value) gopurs_runtime.Value {
s_prime_4_3 := s_3.IntVal + 1
_ = s_prime_4_3
return gopurs_runtime.RecordDict2("state", "value", gopurs_runtime.Int(s_prime_4_3), gopurs_runtime.Int(s_prime_4_3))
}), __local_var_1_0), gopurs_runtime.Value{})
_ = _dollar__unused_3_2
return gopurs_runtime.Apply((*[1024]gopurs_runtime.Value)(v_0.UnsafePtr)[0], i_2_1)
})
}()
})
	})
	return next
}

var pushWhile gopurs_runtime.Value
var once_pushWhile sync.Once
func Get_pushWhile() gopurs_runtime.Value {
	once_pushWhile.Do(func() {
		pushWhile = gopurs_runtime.Func3(func(p_0_box gopurs_runtime.Value, iter_1_box gopurs_runtime.Value, array_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_pushWhile(p_0_box, iter_1_box, array_2_box)
})
	})
	return pushWhile
}

var pushAll gopurs_runtime.Value
var once_pushAll sync.Once
func Get_pushAll() gopurs_runtime.Value {
	once_pushAll.Do(func() {
		pushAll = gopurs_runtime.Apply(Get_pushWhile(), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(true)
}))
	})
	return pushAll
}

var iterator gopurs_runtime.Value
var once_iterator sync.Once
func Get_iterator() gopurs_runtime.Value {
	once_iterator.Do(func() {
		iterator = gopurs_runtime.Func(func(f_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
__local_var_1_0 := gopurs_runtime.Apply(Get_Iterator(), f_0)
_ = __local_var_1_0
return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
__local_ref_2 := 0
_ = __local_ref_2
__local_var_2_1 := gopurs_runtime.Value{PtrVal: &__local_ref_2}
_ = __local_var_2_1
return gopurs_runtime.Apply(__local_var_1_0, __local_var_2_1)
})
}()
})
	})
	return iterator
}

var iterate gopurs_runtime.Value
var once_iterate sync.Once
func Get_iterate() gopurs_runtime.Value {
	once_iterate.Do(func() {
		iterate = gopurs_runtime.Func2(func(iter_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_iterate(iter_0_box, f_1_box)
})
	})
	return iterate
}

var exhausted gopurs_runtime.Value
var once_exhausted sync.Once
func Get_exhausted() gopurs_runtime.Value {
	once_exhausted.Do(func() {
		exhausted = gopurs_runtime.Func(func(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
__local_var_1_0 := (*[1024]gopurs_runtime.Value)(x_0.UnsafePtr)[1]
_ = __local_var_1_0
return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
i_2_1 := *(__local_var_1_0.PtrVal.(*gopurs_runtime.Value))
_ = i_2_1
__local_var_3_2 := gopurs_runtime.Apply((*[1024]gopurs_runtime.Value)(x_0.UnsafePtr)[0], i_2_1)
_ = __local_var_3_2
return gopurs_runtime.Apply(gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
var __t3 gopurs_runtime.Value
{
if gopurs_runtime.Bool(__local_var_3_2.StrVal == "Nothing").IntVal != 0 {
__t3 = gopurs_runtime.Bool(true)
goto end_branch_3
} else {

}
}
{
if gopurs_runtime.Bool(__local_var_3_2.StrVal == "Just").IntVal != 0 {
__t3 = gopurs_runtime.Bool(false)
goto end_branch_3
} else {

}
}
{
__t3 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_3:
return __t3
}), gopurs_runtime.Value{})
})
}()
})
	})
	return exhausted
}

func Call_pushWhile(p_0_loop gopurs_runtime.Value, iter_1_loop gopurs_runtime.Value, array_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var p_0 gopurs_runtime.Value = p_0_loop
_ = p_0
var iter_1 gopurs_runtime.Value = iter_1_loop
_ = iter_1
var array_2 gopurs_runtime.Value = array_2_loop
_ = array_2
return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
__local_ref_1 := false
_ = __local_ref_1
break__3_0 := gopurs_runtime.Value{PtrVal: &__local_ref_1}
_ = break__3_0
__local_var_4_3 := (*[1024]gopurs_runtime.Value)(iter_1.UnsafePtr)[1]
_ = __local_var_4_3
return gopurs_runtime.Apply(gopurs_runtime.Apply2(pkg_Control_Monad_ST_Internal.Get_while(), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_4_2 := *(break__3_0.PtrVal.(*gopurs_runtime.Value))
_ = __local_var_4_2
return gopurs_runtime.Bool(__local_var_4_2.IntVal != 0 != true)
}), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
i_5_4 := *(__local_var_4_3.PtrVal.(*gopurs_runtime.Value))
_ = i_5_4
mx_6_5 := gopurs_runtime.Apply((*[1024]gopurs_runtime.Value)(iter_1.UnsafePtr)[0], i_5_4)
_ = mx_6_5
var __t7 gopurs_runtime.Value
{
if gopurs_runtime.Bool(mx_6_5.StrVal == "Just").IntVal != 0 && gopurs_runtime.Apply(p_0, (*[1024]gopurs_runtime.Value)(mx_6_5.UnsafePtr)[0]).IntVal != 0 {
__t7 = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
_dollar__unused_7_8 := gopurs_runtime.Apply(gopurs_runtime.Apply2(pkg_Data_Array_ST.Get_push(), (*[1024]gopurs_runtime.Value)(mx_6_5.UnsafePtr)[0], array_2), gopurs_runtime.Value{})
_ = _dollar__unused_7_8
__local_var_8_9 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_next(), iter_1), gopurs_runtime.Value{})
_ = __local_var_8_9
return pkg_Data_Unit.Get_unit()
})
goto end_branch_7
} else {

}
}
{
__t7 = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
*(break__3_0.PtrVal.(*gopurs_runtime.Value)) = true
__local_var_7_6 := true
_ = __local_var_7_6
return pkg_Data_Unit.Get_unit()
})
}
end_branch_7:
return gopurs_runtime.Apply(__t7, gopurs_runtime.Value{})
})), gopurs_runtime.Value{})
})
}

func Call_iterate(iter_0_loop gopurs_runtime.Value, f_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var iter_0 gopurs_runtime.Value = iter_0_loop
_ = iter_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
__local_ref_1 := false
_ = __local_ref_1
break__2_0 := gopurs_runtime.Value{PtrVal: &__local_ref_1}
_ = break__2_0
__local_var_3_3 := gopurs_runtime.Apply(Get_next(), iter_0)
_ = __local_var_3_3
return gopurs_runtime.Apply(gopurs_runtime.Apply2(pkg_Control_Monad_ST_Internal.Get_while(), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_3_2 := *(break__2_0.PtrVal.(*gopurs_runtime.Value))
_ = __local_var_3_2
return gopurs_runtime.Bool(__local_var_3_2.IntVal != 0 != true)
}), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
mx_4_4 := gopurs_runtime.Apply(__local_var_3_3, gopurs_runtime.Value{})
_ = mx_4_4
var __t5 gopurs_runtime.Value
{
if gopurs_runtime.Bool(mx_4_4.StrVal == "Just").IntVal != 0 {
__t5 = gopurs_runtime.Apply(f_1, (*[1024]gopurs_runtime.Value)(mx_4_4.UnsafePtr)[0])
goto end_branch_5
} else {

}
}
{
if gopurs_runtime.Bool(mx_4_4.StrVal == "Nothing").IntVal != 0 {
__t5 = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
*(break__2_0.PtrVal.(*gopurs_runtime.Value)) = true
__local_var_5_6 := true
_ = __local_var_5_6
return pkg_Data_Unit.Get_unit()
})
goto end_branch_5
} else {

}
}
{
__t5 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_5:
return gopurs_runtime.Apply(__t5, gopurs_runtime.Value{})
})), gopurs_runtime.Value{})
})
}


