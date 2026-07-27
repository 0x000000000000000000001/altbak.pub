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
return gopurs_runtime.Func0(func() gopurs_runtime.Value {
return Call_try(func() interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply(action_0_box, nil))
})()
})
})
	})
	return cache_try
}

var cache_throw gopurs_runtime.Value
var once_throw sync.Once
func Get_throw() gopurs_runtime.Value {
	once_throw.Do(func() {
		cache_throw = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func0(func() gopurs_runtime.Value {
return gopurs_runtime.Any(Call_throw(x_0_box.StrVal())())
})
})
	})
	return cache_throw
}

var cache_stack gopurs_runtime.Value
var once_stack sync.Once
func Get_stack() gopurs_runtime.Value {
	once_stack.Do(func() {
		cache_stack = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(func(inner_arg0 gopurs_runtime.Value) *pkg_Data_Maybe.Constructor_Just[string] {
return (*pkg_Data_Maybe.Constructor_Just[string])(gopurs_runtime.Apply(gopurs_runtime.Apply2(Get_stackImpl(), pkg_Data_Maybe.Get_Just(), gopurs_runtime.Any(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: nil})), inner_arg0).UnsafePtr)
}(arg0))}
})
	})
	return cache_stack
}

var cache_showError gopurs_runtime.Value
var once_showError sync.Once
func Get_showError() gopurs_runtime.Value {
	once_showError.Do(func() {
		cache_showError = gopurs_runtime.Any(gopurs_runtime.UnboxAny(gopurs_runtime.RecordDict1("show", Get_showErrorImpl())))
	})
	return cache_showError
}

var cache_catchException gopurs_runtime.Value
var once_catchException sync.Once
func Get_catchException() gopurs_runtime.Value {
	once_catchException.Do(func() {
		cache_catchException = gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func0(func() gopurs_runtime.Value {
return gopurs_runtime.Any(CatchException(func(inner_arg0 gopurs_runtime.Value) func() interface{} {
return func() interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply(gopurs_runtime.Apply(arg0, inner_arg0), nil))
}
}, func() interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply(arg1, nil))
})())
})
})
	})
	return cache_catchException
}

var cache_error gopurs_runtime.Value
var once_error sync.Once
func Get_error() gopurs_runtime.Value {
	once_error.Do(func() {
		cache_error = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
return Error(arg0.StrVal())
})
	})
	return cache_error
}

var cache_errorWithCause gopurs_runtime.Value
var once_errorWithCause sync.Once
func Get_errorWithCause() gopurs_runtime.Value {
	once_errorWithCause.Do(func() {
		cache_errorWithCause = gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
return ErrorWithCause(arg0.StrVal(), arg1)
})
	})
	return cache_errorWithCause
}

var cache_errorWithName gopurs_runtime.Value
var once_errorWithName sync.Once
func Get_errorWithName() gopurs_runtime.Value {
	once_errorWithName.Do(func() {
		cache_errorWithName = gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
return ErrorWithName(arg0.StrVal(), arg1.StrVal())
})
	})
	return cache_errorWithName
}

var cache_message gopurs_runtime.Value
var once_message sync.Once
func Get_message() gopurs_runtime.Value {
	once_message.Do(func() {
		cache_message = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str(Message(arg0))
})
	})
	return cache_message
}

var cache_name gopurs_runtime.Value
var once_name sync.Once
func Get_name() gopurs_runtime.Value {
	once_name.Do(func() {
		cache_name = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str(Name(arg0))
})
	})
	return cache_name
}

var cache_showErrorImpl gopurs_runtime.Value
var once_showErrorImpl sync.Once
func Get_showErrorImpl() gopurs_runtime.Value {
	once_showErrorImpl.Do(func() {
		cache_showErrorImpl = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str(ShowErrorImpl(arg0))
})
	})
	return cache_showErrorImpl
}

var cache_stackImpl gopurs_runtime.Value
var once_stackImpl sync.Once
func Get_stackImpl() gopurs_runtime.Value {
	once_stackImpl.Do(func() {
		cache_stackImpl = gopurs_runtime.Func3(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(StackImpl(func(inner_arg0 interface{}) *pkg_Data_Maybe.Constructor_Just[interface{}] {
return (*pkg_Data_Maybe.Constructor_Just[interface{}])(gopurs_runtime.Apply(arg0, gopurs_runtime.Any(inner_arg0)).UnsafePtr)
}, (*pkg_Data_Maybe.Constructor_Just[interface{}])(arg1.UnsafePtr), arg2))}
})
	})
	return cache_stackImpl
}

var cache_throwException gopurs_runtime.Value
var once_throwException sync.Once
func Get_throwException() gopurs_runtime.Value {
	once_throwException.Do(func() {
		cache_throwException = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func0(func() gopurs_runtime.Value {
return gopurs_runtime.Any(ThrowException(arg0)())
})
})
	})
	return cache_throwException
}

func Call_try(action_0_loop func() interface{}) func() gopurs_runtime.Value {
var action_0 func() interface{} = action_0_loop
_ = action_0
return func() gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply2(Get_catchException(), gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Effect.Get_applicativeEffect(), "pure"), gopurs_runtime.Any(gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&pkg_Data_Either.Constructor_Left[gopurs_runtime.Value, gopurs_runtime.Value]{gopurs_runtime.UnboxAny(x_1)})}))
}), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Effect.Get_functorEffect(), "map"), pkg_Data_Either.Get_Right(), gopurs_runtime.Func0(func() gopurs_runtime.Value {
return gopurs_runtime.Any(action_0())
}))), nil)
}
}

func Call_throw(x_0_loop string) func() interface{} {
var x_0 string = x_0_loop
_ = x_0
return func() interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply(gopurs_runtime.Apply(Get_throwException(), gopurs_runtime.Apply(Get_error(), gopurs_runtime.Str(x_0))), nil))
}
}
