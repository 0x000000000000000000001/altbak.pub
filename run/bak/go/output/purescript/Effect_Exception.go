package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Effect_Exception_pure gopurs_runtime.Value
var once_Effect_Exception_pure sync.Once
func Get_Effect_Exception_pure() gopurs_runtime.Value {
	once_Effect_Exception_pure.Do(func() {
		cache_Effect_Exception_pure = Get_Effect_pureE()
	})
	return cache_Effect_Exception_pure
}

var cache_Effect_Exception_try gopurs_runtime.Value
var once_Effect_Exception_try sync.Once
func Get_Effect_Exception_try() gopurs_runtime.Value {
	once_Effect_Exception_try.Do(func() {
		cache_Effect_Exception_try = gopurs_runtime.Func(func(action_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Effect_Exception_try(action_0_box)
})
	})
	return cache_Effect_Exception_try
}

var cache_Effect_Exception_throw gopurs_runtime.Value
var once_Effect_Exception_throw sync.Once
func Get_Effect_Exception_throw() gopurs_runtime.Value {
	once_Effect_Exception_throw.Do(func() {
		cache_Effect_Exception_throw = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Effect_Exception_throw(x_0_box.StrVal())
})
	})
	return cache_Effect_Exception_throw
}

var cache_Effect_Exception_stack gopurs_runtime.Value
var once_Effect_Exception_stack sync.Once
func Get_Effect_Exception_stack() gopurs_runtime.Value {
	once_Effect_Exception_stack.Do(func() {
		cache_Effect_Exception_stack = gopurs_runtime.Apply2(Get_Effect_Exception_stackImpl(), Get_Data_Maybe_Just(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))})
	})
	return cache_Effect_Exception_stack
}

var cache_Effect_Exception_showError gopurs_runtime.Value
var once_Effect_Exception_showError sync.Once
func Get_Effect_Exception_showError() gopurs_runtime.Value {
	once_Effect_Exception_showError.Do(func() {
		cache_Effect_Exception_showError = gopurs_runtime.RecordDict1("show", Get_Effect_Exception_showErrorImpl())
	})
	return cache_Effect_Exception_showError
}

func Call_Effect_Exception_try(action_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var action_0 gopurs_runtime.Value = action_0_loop
_ = action_0
return gopurs_runtime.Apply2(Get_Effect_Exception_catchException(), gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, x_1})}
})
}), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_1_0 := Get_Data_Either_Right()
_ = __local_var_1_0
__local_var_2_1 := gopurs_runtime.Apply(action_0, gopurs_runtime.Value{})
_ = __local_var_2_1
return gopurs_runtime.Apply(__local_var_1_0, __local_var_2_1)
}))
}

func Call_Effect_Exception_throw(x_0_loop string) gopurs_runtime.Value {
var x_0 string = x_0_loop
_ = x_0
return gopurs_runtime.Apply(Get_Effect_Exception_throwException(), gopurs_runtime.Apply(Get_Effect_Exception_error(), gopurs_runtime.Str(x_0)))
}

func Get_Effect_Exception_catchException() gopurs_runtime.Value {
	return _Gopurs_Effect_Exception_CatchException
}

func Get_Effect_Exception_error() gopurs_runtime.Value {
	return _Gopurs_Effect_Exception_Error
}

func Get_Effect_Exception_errorWithCause() gopurs_runtime.Value {
	return _Gopurs_Effect_Exception_ErrorWithCause
}

func Get_Effect_Exception_errorWithName() gopurs_runtime.Value {
	return _Gopurs_Effect_Exception_ErrorWithName
}

func Get_Effect_Exception_message() gopurs_runtime.Value {
	return _Gopurs_Effect_Exception_Message
}

func Get_Effect_Exception_name() gopurs_runtime.Value {
	return _Gopurs_Effect_Exception_Name
}

func Get_Effect_Exception_showErrorImpl() gopurs_runtime.Value {
	return _Gopurs_Effect_Exception_ShowErrorImpl
}

func Get_Effect_Exception_stackImpl() gopurs_runtime.Value {
	return _Gopurs_Effect_Exception_StackImpl
}

func Get_Effect_Exception_throwException() gopurs_runtime.Value {
	return _Gopurs_Effect_Exception_ThrowException
}
