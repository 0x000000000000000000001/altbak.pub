package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Control_Monad_Except_unwrap gopurs_runtime.Value
var once_Control_Monad_Except_unwrap sync.Once
func Get_Control_Monad_Except_unwrap() gopurs_runtime.Value {
	once_Control_Monad_Except_unwrap.Do(func() {
		cache_Control_Monad_Except_unwrap = Get_Unsafe_Coerce_unsafeCoerce()
	})
	return cache_Control_Monad_Except_unwrap
}

var cache_Control_Monad_Except_withExcept gopurs_runtime.Value
var once_Control_Monad_Except_withExcept sync.Once
func Get_Control_Monad_Except_withExcept() gopurs_runtime.Value {
	once_Control_Monad_Except_withExcept.Do(func() {
		cache_Control_Monad_Except_withExcept = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Except_withExcept(f_0_box, v_1_box)
})
	})
	return cache_Control_Monad_Except_withExcept
}

var cache_Control_Monad_Except_runExcept gopurs_runtime.Value
var once_Control_Monad_Except_runExcept sync.Once
func Get_Control_Monad_Except_runExcept() gopurs_runtime.Value {
	once_Control_Monad_Except_runExcept.Do(func() {
		cache_Control_Monad_Except_runExcept = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Except_runExcept(x_0_box)
})
	})
	return cache_Control_Monad_Except_runExcept
}

var cache_Control_Monad_Except_mapExcept gopurs_runtime.Value
var once_Control_Monad_Except_mapExcept sync.Once
func Get_Control_Monad_Except_mapExcept() gopurs_runtime.Value {
	once_Control_Monad_Except_mapExcept.Do(func() {
		cache_Control_Monad_Except_mapExcept = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Except_mapExcept(f_0_box, v_1_box)
})
	})
	return cache_Control_Monad_Except_mapExcept
}

func Call_Control_Monad_Except_withExcept(f_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
var __t0 gopurs_runtime.Value
{
if (v_1.Type == 9 && v_1.IntVal == 2465973597) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right{1, (*Constructor_Data_Either_Right)(v_1.UnsafePtr).V0})}
goto end_branch_0
} else {

}
}
{
if (v_1.Type == 9 && v_1.IntVal == 3711209382) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, gopurs_runtime.Apply(f_0, (*Constructor_Data_Either_Left)(v_1.UnsafePtr).V0)})}
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

func Call_Control_Monad_Except_runExcept(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_Control_Monad_Except_mapExcept(f_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return gopurs_runtime.Apply(f_0, v_1)
}


