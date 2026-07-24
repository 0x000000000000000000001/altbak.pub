package Effect_Exception

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_Maybe "gopurs/output/Data.Maybe"
)

var try gopurs_runtime.Value
var once_try sync.Once
func Get_try() gopurs_runtime.Value {
	once_try.Do(func() {
		try = gopurs_runtime.Func(func(action_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var action_0 gopurs_runtime.Value = action_0_loop
_ = action_0
return gopurs_runtime.Apply2(Get_catchException(), gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Constructor1("Left", x_1)
})
}), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
a_prime_1_0 := gopurs_runtime.Apply(action_0_loop, gopurs_runtime.Value{})
_ = a_prime_1_0
return gopurs_runtime.Constructor1("Right", a_prime_1_0)
}))
}()
})
	})
	return try
}

var throw gopurs_runtime.Value
var once_throw sync.Once
func Get_throw() gopurs_runtime.Value {
	once_throw.Do(func() {
		throw = gopurs_runtime.Func(func(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return gopurs_runtime.Apply(Get_throwException(), gopurs_runtime.Apply(Get_error(), x_0_loop))
}()
})
	})
	return throw
}

var stack gopurs_runtime.Value
var once_stack sync.Once
func Get_stack() gopurs_runtime.Value {
	once_stack.Do(func() {
		stack = gopurs_runtime.Apply2(Get_stackImpl(), pkg_Data_Maybe.Get_Just(), gopurs_runtime.Constructor0("Nothing"))
	})
	return stack
}

var showError gopurs_runtime.Value
var once_showError sync.Once
func Get_showError() gopurs_runtime.Value {
	once_showError.Do(func() {
		showError = gopurs_runtime.RecordDict1("show", Get_showErrorImpl())
	})
	return showError
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
