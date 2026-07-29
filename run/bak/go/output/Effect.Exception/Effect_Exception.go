package Effect_Exception

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_Maybe "gopurs/output/Data.Maybe"
	pkg_Effect "gopurs/output/Effect"
	pkg_Data_Either "gopurs/output/Data.Either"
	unsafe "unsafe"
)

var cache_try gopurs_runtime.Value
var once_try sync.Once
func Get_try() gopurs_runtime.Value {
	once_try.Do(func() {
		cache_try = gopurs_runtime.Func(func(action_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_try(action_0_box)
})
	})
	return cache_try
}

var cache_throw gopurs_runtime.Value
var once_throw sync.Once
func Get_throw() gopurs_runtime.Value {
	once_throw.Do(func() {
		cache_throw = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_throw(x_0_box.StrVal())
})
	})
	return cache_throw
}

var cache_stack gopurs_runtime.Value
var once_stack sync.Once
func Get_stack() gopurs_runtime.Value {
	once_stack.Do(func() {
		cache_stack = gopurs_runtime.Apply2(Get_stackImpl(), pkg_Data_Maybe.Get_Just(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: nil})
	})
	return cache_stack
}

var cache_showError gopurs_runtime.Value
var once_showError sync.Once
func Get_showError() gopurs_runtime.Value {
	once_showError.Do(func() {
		cache_showError = gopurs_runtime.RecordDict1("show", Get_showErrorImpl())
	})
	return cache_showError
}

func Call_try(action_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var action_0 gopurs_runtime.Value = action_0_loop
_ = action_0
return gopurs_runtime.Apply2(Get_catchException(), gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Effect.Get_applicativeEffect(), "pure"), func() gopurs_runtime.Value {
if ((action_0.UnsafePtr) != (nil)) && (((*struct{Rc uint32})(action_0.UnsafePtr).Rc) == (1)) {
(*pkg_Data_Either.Constructor_Left[gopurs_runtime.Value, gopurs_runtime.Value])(action_0.UnsafePtr).V0 = x_1
return action_0
} else {
return gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&pkg_Data_Either.Constructor_Left[gopurs_runtime.Value, gopurs_runtime.Value]{1, x_1})}
}
}())
}), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Effect.Get_functorEffect(), "map"), pkg_Data_Either.Get_Right(), action_0))
}

func Call_throw(x_0_loop string) gopurs_runtime.Value {
var x_0 string = x_0_loop
_ = x_0
return gopurs_runtime.Apply(Get_throwException(), gopurs_runtime.Apply(Get_error(), gopurs_runtime.Str(x_0)))
}

func Get_catchException() gopurs_runtime.Value {
	return _Gopurs_CatchException
}

func Get_error() gopurs_runtime.Value {
	return _Gopurs_Error
}

func Get_errorWithCause() gopurs_runtime.Value {
	return _Gopurs_ErrorWithCause
}

func Get_errorWithName() gopurs_runtime.Value {
	return _Gopurs_ErrorWithName
}

func Get_message() gopurs_runtime.Value {
	return _Gopurs_Message
}

func Get_name() gopurs_runtime.Value {
	return _Gopurs_Name
}

func Get_showErrorImpl() gopurs_runtime.Value {
	return _Gopurs_ShowErrorImpl
}

func Get_stackImpl() gopurs_runtime.Value {
	return _Gopurs_StackImpl
}

func Get_throwException() gopurs_runtime.Value {
	return _Gopurs_ThrowException
}
